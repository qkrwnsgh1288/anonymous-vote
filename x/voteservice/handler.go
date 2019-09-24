package voteservice

import (
	"fmt"
	"github.com/qkrwnsgh1288/vote-dapp/x/voteservice/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for "voteservice" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgMakeAgenda:
			return handleMsgMakeAgenda(ctx, keeper, msg)
		case MsgVoteAgenda:
			return handleMsgVoteAgenda(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized voteservice Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgMakeAgenda(ctx sdk.Context, keeper Keeper, msg MsgMakeAgenda) sdk.Result {
	// todo: more valid check
	if keeper.GetAgendaTopic(ctx, msg.AgendaTopic) != "" {
		return types.ErrAgendaTopicAlreadyExist(types.DefaultCodespace).Result()
	}
	agenda := types.Agenda{
		AgendaProposer: msg.AgendaProposer,
		AgendaTopic:    msg.AgendaTopic,
		AgendaContent:  msg.AgendaContent,
	}
	keeper.SetAgenda(ctx, msg.AgendaTopic, agenda)
	return sdk.Result{}
}
func handleMsgVoteAgenda(ctx sdk.Context, keeper Keeper, msg MsgVoteAgenda) sdk.Result {
	// todo: valid check more
	if !keeper.IsTopicPresent(ctx, msg.AgendaTopic) {
		return types.ErrAgendaTopicDoesNotExist(types.DefaultCodespace).Result()
	}

	agenda := keeper.GetAgenda(ctx, msg.AgendaTopic)
	agenda.Voters = append(agenda.Voters, msg.VoteAddr.String()) // todo: string -> sdk.AccAddress
	if msg.YesOrNo {
		agenda.ProCount += 1
	} else {
		agenda.NegCount += 1
	}
	keeper.SetAgenda(ctx, msg.AgendaTopic, agenda)
	return sdk.Result{}
}
