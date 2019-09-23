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
)

type (
	Keeper        = keeper.Keeper
	MsgSetName    = types.MsgSetName
	MsgBuyName    = types.MsgBuyName
	MsgDeleteName = types.MsgDeleteName
	MsgAgenda     = types.MsgAgenda

	QueryResResolve = types.QueryResResolve
	QueryResNames   = types.QueryResNames
	Whois           = types.Whois
)
