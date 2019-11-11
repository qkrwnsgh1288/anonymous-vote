package voteservice

import (
	"github.com/qkrwnsgh1288/anonymous-vote/x/voteservice/internal/keeper"
	"github.com/qkrwnsgh1288/anonymous-vote/x/voteservice/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewKeeper     = keeper.NewKeeper
	NewQuerier    = keeper.NewQuerier
	RegisterCodec = types.RegisterCodec
	ModuleCdc     = types.ModuleCdc

	NewMsgMakeAgenda = types.NewMsgMakeAgenda
	NewMsgVoteAgenda = types.NewMsgVoteAgenda
	NewAgenda        = types.NewAgenda
)

type (
	Keeper = keeper.Keeper

	MsgMakeAgenda         = types.MsgMakeAgenda
	MsgRegisterByVoter    = types.MsgRegisterByVoter
	MsgRegisterByProposer = types.MsgRegisterByProposer
	MsgVoteAgenda         = types.MsgVoteAgenda
	MsgTally              = types.MsgTally
	Agenda                = types.Agenda
)
