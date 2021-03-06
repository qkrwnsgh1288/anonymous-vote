package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc is the codec for the module
var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgMakeAgenda{}, "voteservice/MakeAgenda", nil)
	cdc.RegisterConcrete(MsgRegisterByVoter{}, "voteservice/MsgRegisterByVoter", nil)
	cdc.RegisterConcrete(MsgRegisterByProposer{}, "voteservice/MsgRegisterByProposer", nil)
	cdc.RegisterConcrete(MsgVoteAgenda{}, "voteservice/VoteAgenda", nil)
	cdc.RegisterConcrete(MsgTally{}, "voteservice/Tally", nil)
}
