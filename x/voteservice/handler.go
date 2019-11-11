package voteservice

import (
	"fmt"
	"github.com/qkrwnsgh1288/anonymous-vote/x/voteservice/common"
	"github.com/qkrwnsgh1288/anonymous-vote/x/voteservice/crypto"
	"github.com/qkrwnsgh1288/anonymous-vote/x/voteservice/internal/types"

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

		State: msg.State,
		Voter: msg.Voter,
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

	// setup list check && save address, registeredKey
	hasSetupList := false
	for i, setupAddr := range agenda.SetupList {
		if setupAddr == addr {
			agenda.Voter[i].Addr = addr
			agenda.Voter[i].RegisteredKey = types.SPoint{
				X: zkInfo.XG.X.String(),
				Y: zkInfo.XG.Y.String(),
			}
			hasSetupList = true
		}
	}
	if !hasSetupList {
		return types.ErrDoesNotRegisterAddress(types.DefaultCodespace).Result()
	}

	keeper.SetAgenda(ctx, msg.AgendaTopic, agenda)

	return sdk.Result{}
}
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
