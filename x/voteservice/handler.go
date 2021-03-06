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
		case MsgTally:
			return handleMsgTally(ctx, keeper, msg)
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

		WhiteList: msg.WhiteList,
		State:     crypto.SIGNUP,
		//Voters:    msg.Voters,
	}

	keeper.SetAgenda(ctx, msg.AgendaTopic, agenda)
	return sdk.Result{}
}

// 2. MsgRegisterByVoter (Addr, RegisteredKey setting)
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
	// todo: It should consider calling CreateZKP on the front-end and verifying only here.
	v1_r, v1_vG, err := crypto.CreateZKP(addr, zkInfo.X, zkInfo.V, zkInfo.XG)
	if err != nil {
		return types.ErrInvalidPubkeyInCreateZKP(types.DefaultCodespace).Result()
	}
	// verifyZKP
	if !crypto.VerifyZKP(addr, zkInfo.XG, v1_r, v1_vG) {
		return types.ErrInvalidVerifyZKP(types.DefaultCodespace).Result()
	}

	agenda := keeper.GetAgenda(ctx, msg.AgendaTopic)

	// check state
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
	for _, setupAddr := range agenda.WhiteList {
		if setupAddr == addr {
			tmpVoter := types.MakeDefaultSVoter()
			tmpVoter.Addr = addr
			tmpVoter.RegisteredKey = types.SPoint{
				X: zkInfo.XG.X.String(),
				Y: zkInfo.XG.Y.String(),
			}
			agenda.Voters = append(agenda.Voters, tmpVoter)
			agenda.TotalRegistered += 1

			hasSetupList = true
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

	// check state
	if agenda.State != crypto.SIGNUP {
		return types.ErrStateIsNotSIGNUP(types.DefaultCodespace).Result()
	}
	// It is only possible by proposer
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
	agenda.State = crypto.VOTE
	keeper.SetAgenda(ctx, msg.AgendaTopic, agenda)
	return sdk.Result{}
}

// 4. MsgVoteAgenda
func handleMsgVoteAgenda(ctx sdk.Context, keeper Keeper, msg MsgVoteAgenda) sdk.Result {
	// todo: valid check more
	if !keeper.IsTopicPresent(ctx, msg.AgendaTopic) {
		return types.ErrAgendaTopicDoesNotExist(types.DefaultCodespace).Result()
	}
	agenda := keeper.GetAgenda(ctx, msg.AgendaTopic)

	// check state
	if agenda.State != crypto.VOTE {
		return types.ErrStateIsNotVOTE(types.DefaultCodespace).Result()
	}

	zkInfo := crypto.ZkInfo{
		X: common.GetBigInt(msg.ZkInfo[0], 10),
		V: common.GetBigInt(msg.ZkInfo[3], 10),
		W: common.GetBigInt(msg.ZkInfo[4], 10),
		R: common.GetBigInt(msg.ZkInfo[5], 10),
		D: common.GetBigInt(msg.ZkInfo[6], 10),
	}
	addr := msg.VoteAddr.String()

	for i, voter := range agenda.Voters {
		if voter.Addr == addr {
			xG := types.GetPointFromSPoint(voter.RegisteredKey, 10)
			yG := types.GetPointFromSPoint(voter.ReconstructedKey, 10)
			isFirst := false

			switch msg.YesOrNo {
			case "yes":
				// todo: It should consider calling Create1outof2ZKPYesVote on the front-end and verifying only here.
				y, a1, b1, a2, b2, params, _ := crypto.Create1outof2ZKPYesVote(addr, xG, yG, zkInfo.W, zkInfo.R, zkInfo.D, zkInfo.X)
				if !crypto.Verify1outof2ZKP(addr, params, xG, yG, y, a1, b1, a2, b2) {
					return types.ErrInvalidVerify1outof2ZKP(types.DefaultCodespace).Result()
				}
				if voter.Vote.X == "" && voter.Vote.Y == "" {
					isFirst = true
				}
				agenda.Voters[i].Commitment = crypto.CommitToVote(addr, params, xG, yG, y, a1, b1, a2, b2)
				agenda.Voters[i].Vote = types.GetSPointFromPoint(y)
			case "no":
				// todo: It should consider calling Create1outof2ZKPNoVote on the front-end and verifying only here.
				y, a1, b1, a2, b2, params, _ := crypto.Create1outof2ZKPNoVote(addr, xG, yG, zkInfo.W, zkInfo.R, zkInfo.D, zkInfo.X)
				if !crypto.Verify1outof2ZKP(addr, params, xG, yG, y, a1, b1, a2, b2) {
					return types.ErrInvalidVerify1outof2ZKP(types.DefaultCodespace).Result()
				}
				if voter.Vote.X == "" && voter.Vote.Y == "" {
					isFirst = true
				}
				agenda.Voters[i].Commitment = crypto.CommitToVote(addr, params, xG, yG, y, a1, b1, a2, b2)
				agenda.Voters[i].Vote = types.GetSPointFromPoint(y)
			default:
				return types.ErrInvalidAnswer(types.DefaultCodespace).Result()
			}
			if isFirst {
				agenda.TotalVoteComplete += 1
			}
			break
		}
	}

	keeper.SetAgenda(ctx, msg.AgendaTopic, agenda)
	return sdk.Result{}
}

// 5. MsgTally
func handleMsgTally(ctx sdk.Context, keeper Keeper, msg types.MsgTally) sdk.Result {
	// todo: valid check more
	if !keeper.IsTopicPresent(ctx, msg.AgendaTopic) {
		return types.ErrAgendaTopicDoesNotExist(types.DefaultCodespace).Result()
	}
	agenda := keeper.GetAgenda(ctx, msg.AgendaTopic)
	totalRegistered := agenda.TotalRegistered

	// check state
	if agenda.State != crypto.VOTE {
		return types.ErrStateIsNotVOTE(types.DefaultCodespace).Result()
	}
	// It is only possible by proposer
	if agenda.AgendaProposer.String() != msg.ProposerAddr.String() {
		return types.ErrDoNotHavePermission(types.DefaultCodespace).Result()
	}
	// check whether voting is completed
	if agenda.TotalRegistered != agenda.TotalVoteComplete {
		return types.ErrVotingIsNotFinished(types.DefaultCodespace).Result()
	}

	temp := crypto.MakeDefaultJacobianPoint()
	vote := crypto.MakeDefaultPoint()
	G := crypto.MakeDefaultPoint()

	G.X.SetBytes(crypto.Curve.Gx.Bytes())
	G.Y.SetBytes(crypto.Curve.Gy.Bytes())

	tempCurve := crypto.MakeDefaultPoint()
	tempCurve.X.SetBytes(crypto.Curve.Gx.Bytes())
	tempCurve.Y.SetBytes(crypto.Curve.Gy.Bytes())

	// Sum all votes
	for i := 0; i < totalRegistered; i++ {
		vote.X.SetString(agenda.Voters[i].Vote.X, 10)
		vote.Y.SetString(agenda.Voters[i].Vote.Y, 10)

		if i == 0 {
			temp.X.SetBytes(vote.X.Bytes())
			temp.Y.SetBytes(vote.Y.Bytes())
			temp.Z = big.NewInt(1)
		} else {
			crypto.AddMixedM(&temp, vote)
		}
	}

	// Each vote is represented by a G.
	// If there are no votes... then it is 0G = (0,0)...
	if temp.X.Cmp(big.NewInt(0)) == 0 {
		fmt.Println("temp.X = ", temp.X)
	} else {
		crypto.ToZ1(&temp, crypto.Curve.P)

		tempG := crypto.MakeDefaultJacobianPoint()
		tempG.X.SetBytes(G.X.Bytes())
		tempG.Y.SetBytes(G.Y.Bytes())
		tempG.Z = big.NewInt(1)

		// Start adding 'G' and looking for a match
		for i := 1; i <= totalRegistered; i++ {
			if temp.X.Cmp(tempG.X) == 0 {
				agenda.FinalYes = i
				agenda.FinalNo = agenda.TotalRegistered - i

				agenda.State = crypto.FINISHED
				keeper.SetAgenda(ctx, msg.AgendaTopic, agenda)
				return sdk.Result{}
			}
			//tempG.X, tempG.Y = Curve.Add(tempG.X, tempG.Y, Curve.Gx, Curve.Gy)
			crypto.AddMixedM(&tempG, G)
			crypto.ToZ1(&tempG, crypto.Curve.P)
		}
	}
	return types.ErrSomethingIsBad(types.DefaultCodespace).Result()

}
