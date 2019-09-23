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
	NewKeeper        = keeper.NewKeeper
	NewQuerier       = keeper.NewQuerier
	NewMsgBuyName    = types.NewMsgBuyName
	NewMsgSetName    = types.NewMsgSetName
	NewMsgDeleteName = types.NewMsgDeleteName
	NewWhois         = types.NewWhois
	RegisterCodec    = types.RegisterCodec
	ModuleCdc        = types.ModuleCdc

	NewVoteKeeper    = keeper.NewVoteKeeper
	NewVoteQuerier   = keeper.NewVoteQuerier
	NewMsgMakeAgenda = types.NewMsgMakeAgenda
	NewAgenda        = types.NewAgenda
)

type (
	Keeper = keeper.Keeper

	MsgSetName    = types.MsgSetName
	MsgBuyName    = types.MsgBuyName
	MsgDeleteName = types.MsgDeleteName

	QueryResResolve = types.QueryResResolve
	QueryResNames   = types.QueryResNames
	Whois           = types.Whois

	VoteKeeper    = keeper.VoteKeeper
	MsgMakeAgenda = types.MsgMakeAgenda
	Agenda        = types.Agenda
)
