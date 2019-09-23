package voteservice

import (
	"github.com/qkrwnsgh1288/vote-dapp/x/voteservice/internal/keeper"
	"github.com/qkrwnsgh1288/vote-dapp/x/voteservice/internal/types"
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
	NewAgenda        = types.NewAgenda
)

type (
	Keeper = keeper.Keeper

	MsgMakeAgenda = types.MsgMakeAgenda
	Agenda        = types.Agenda
)
