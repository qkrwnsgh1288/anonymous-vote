package voteservice

import (
	"fmt"
	"github.com/qkrwnsgh1288/anonymous-vote/x/voteservice/common"
	"github.com/qkrwnsgh1288/anonymous-vote/x/voteservice/crypto"
	"github.com/qkrwnsgh1288/anonymous-vote/x/voteservice/internal/types"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for "voteservice" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgMakeAgenda:
			return handleMsgMakeAgenda(ctx, keeper, msg)
		case MsgRegisterByVoter:
			return handleMsgRegisterByVoter(ctx, keeper, msg)
		case MsgRegisterByProposer:
			return handleMsgRegisterByProposer(ctx, keeper, msg)
		case MsgVoteAgenda:
			return handleMsgVoteAgenda(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized voteservice Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// 1. MsgMakeAgenda
func handleMsgMakeAgenda(ctx sdk.Context, keeper Keeper, msg MsgMakeAgenda) sdk.Result {
	// todo: more valid check
	if keeper.IsTopicPresent(ctx, msg.AgendaTopic) {
		return types.ErrAgendaTopicAlreadyExist(types.DefaultCodespace).Result()
	}

	agenda := types.Agenda{
		AgendaProposer: msg.AgendaProposer,
		AgendaTopic:    msg.AgendaTopic,
		AgendaContent:  msg.AgendaContent,
		SetupList:      msg.SetupList,
		VoteCheckList:  msg.VoteCheckList,
		Progress:       fmt.Sprintf("%d/%d", 0, len(msg.SetupList)),

		State:  msg.State,
		Voters: msg.Voter,
	}

	keeper.SetAgenda(ctx, msg.AgendaTopic, agenda)
	return sdk.Result{}
}

// 2. MsgRegisterByVoter
func handleMsgRegisterByVoter(ctx sdk.Context, keeper Keeper, msg MsgRegisterByVoter) sdk.Result {
	// todo: valid check more
	if !keeper.IsTopicPresent(ctx, msg.AgendaTopic) {
		return types.ErrAgendaTopicDoesNotExist(types.DefaultCodespace).Result()
	}
	zkInfo := crypto.ZkInfo{
		X: common.GetBigInt(msg.ZkInfo[0], 10),
		XG: crypto.Point{
			X: common.GetBigInt(msg.ZkInfo[1], 10),
			Y: common.GetBigInt(msg.ZkInfo[2], 10),
		},
		V: common.GetBigInt(msg.ZkInfo[3], 10),
		//W: common.GetBigInt(msg.ZkInfo[4], 10),
		//R: common.GetBigInt(msg.ZkInfo[5], 10),
		//D: common.GetBigInt(msg.ZkInfo[6], 10),
	}
	addr := msg.VoteAddr.String()

	// createZKP
	v1_r, v1_vG, err := crypto.CreateZKP(addr, zkInfo.X, zkInfo.V, zkInfo.XG)
	if err != nil {
		return types.ErrInvalidPubkeyInCreateZKP(types.DefaultCodespace).Result()
	}
	// verifyZKP
	if !crypto.VerifyZKP(addr, zkInfo.XG, v1_r, v1_vG) {
		return types.ErrInvalidVerifyZKP(types.DefaultCodespace).Result()
	}

	agenda := keeper.GetAgenda(ctx, msg.AgendaTopic)
	if agenda.State != crypto.SIGNUP {
		return types.ErrStateIsNotSIGNUP(types.DefaultCodespace).Result()
	}

	// check whether already registered
	for _, voter := range agenda.Voters {
		if voter.Addr == addr {
			return types.ErrAlreadyRegisterd(types.DefaultCodespace).Result()
		}
	}

	// setup list check && save address, registeredKey
	hasSetupList := false
	for i, setupAddr := range agenda.SetupList {
		if setupAddr == addr {
			agenda.Voters[i].Addr = addr
			agenda.Voters[i].RegisteredKey = types.SPoint{
				X: zkInfo.XG.X.String(),
				Y: zkInfo.XG.Y.String(),
			}
			hasSetupList = true
			agenda.TotalRegistered += 1
			break
		}
	}
	if !hasSetupList {
		return types.ErrDoesNotRegisterAddress(types.DefaultCodespace).Result()
	}
	keeper.SetAgenda(ctx, msg.AgendaTopic, agenda)
	return sdk.Result{}
}

// 3. MsgRegisterByProposer
func handleMsgRegisterByProposer(ctx sdk.Context, keeper Keeper, msg MsgRegisterByProposer) sdk.Result {
	// todo: valid check more
	if !keeper.IsTopicPresent(ctx, msg.AgendaTopic) {
		return types.ErrAgendaTopicDoesNotExist(types.DefaultCodespace).Result()
	}

	agenda := keeper.GetAgenda(ctx, msg.AgendaTopic)
	totalRegistered := agenda.TotalRegistered

	if agenda.State != crypto.SIGNUP {
		return types.ErrStateIsNotSIGNUP(types.DefaultCodespace).Result()
	}
	if agenda.AgendaProposer.String() != msg.ProposerAddr.String() {
		return types.ErrDoNotHavePermission(types.DefaultCodespace).Result()
	}

	if totalRegistered < crypto.Minimum_voter_count {
		return types.ErrInvalidTotalRegisteredCnt(types.DefaultCodespace).Result()
	}
	temp := crypto.MakeDefaultPoint()
	yG := crypto.MakeDefaultPoint()
	beforei := crypto.MakeDefaultPoint()
	afteri := crypto.MakeDefaultPoint()

	// Step 1 is to compute the index 1 reconstructed key
	afteri.X.SetString(agenda.Voters[1].RegisteredKey.X, 10)
	afteri.Y.SetString(agenda.Voters[1].RegisteredKey.Y, 10)

	for i := 2; i < totalRegistered; i++ {
		afteri.X, afteri.Y = crypto.Curve.Add(afteri.X, afteri.Y,
			common.GetBigInt(agenda.Voters[i].RegisteredKey.X, 10), common.GetBigInt(agenda.Voters[i].RegisteredKey.Y, 10))
	}

	agenda.Voters[0].ReconstructedKey.X = afteri.X.String()
	agenda.Voters[0].ReconstructedKey.Y = new(big.Int).Sub(crypto.Curve.P, afteri.Y).String()

	// Step 2 is to add to beforei, and subtract from afteri.
	for i := 1; i < totalRegistered; i++ {
		if i == 1 {
			beforei.X.SetString(agenda.Voters[0].RegisteredKey.X, 10)
			beforei.Y.SetString(agenda.Voters[0].RegisteredKey.Y, 10)
		} else {
			beforei.X, beforei.Y = crypto.Curve.Add(beforei.X, beforei.Y,
				common.GetBigInt(agenda.Voters[i-1].RegisteredKey.X, 10), common.GetBigInt(agenda.Voters[i-1].RegisteredKey.Y, 10))
		}

		// If we have reached the end... just store beforei
		// Otherwise, we need to compute a key.
		// Counting from 0 to n-1...
		if i == (totalRegistered - 1) {
			agenda.Voters[i].ReconstructedKey.X = beforei.X.String()
			agenda.Voters[i].ReconstructedKey.Y = beforei.Y.String()
		} else {
			// Subtract 'i' from afteri
			temp.X.SetString(agenda.Voters[i].RegisteredKey.X, 10)
			temp.Y.Sub(crypto.Curve.P, common.GetBigInt(agenda.Voters[i].RegisteredKey.Y, 10))

			// Grab negation of afteri (did not seem to work with Jacob co-ordinates)
			afteri.X, afteri.Y = crypto.Curve.Add(afteri.X, afteri.Y, temp.X, temp.Y)

			temp.X.SetBytes(afteri.X.Bytes())
			temp.Y.Sub(crypto.Curve.P, afteri.Y)

			// Now we do beforei - afteri...
			yG.X, yG.Y = crypto.Curve.Add(beforei.X, beforei.Y, temp.X, temp.Y)

			agenda.Voters[i].ReconstructedKey.X = yG.X.String()
			agenda.Voters[i].ReconstructedKey.Y = yG.Y.String()
		}
	}
	agenda.State = crypto.COMMITMENT
	keeper.SetAgenda(ctx, msg.AgendaTopic, agenda)
	return sdk.Result{}
}

//
func handleMsgVoteAgenda(ctx sdk.Context, keeper Keeper, msg MsgVoteAgenda) sdk.Result {
	// todo: valid check more
	if !keeper.IsTopicPresent(ctx, msg.AgendaTopic) {
		return types.ErrAgendaTopicDoesNotExist(types.DefaultCodespace).Result()
	}
	voteCount := 0
	agenda := keeper.GetAgenda(ctx, msg.AgendaTopic)

	for i, val := range agenda.SetupList {
		if msg.VoteAddr.String() == val {
			agenda.VoteCheckList[i] = msg.YesOrNo
		}
		if agenda.VoteCheckList[i] != "empty" {
			voteCount += 1
		}
	}
	agenda.Progress = fmt.Sprintf("%d/%d", voteCount, len(agenda.SetupList))

	keeper.SetAgenda(ctx, msg.AgendaTopic, agenda)
	return sdk.Result{}
}
