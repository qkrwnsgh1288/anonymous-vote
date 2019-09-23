package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/qkrwnsgh1288/vote-dapp/x/voteservice/internal/types"
)

type VoteKeeper struct {
	storekey sdk.StoreKey
	cdc      *codec.Codec
}

func NewVoteKeeper(storeKey sdk.StoreKey, cdc *codec.Codec) VoteKeeper {
	return VoteKeeper{
		storekey: storeKey,
		cdc:      cdc,
	}
}

func (v VoteKeeper) setAgenda(ctx sdk.Context, agendaTopic string, agenda types.Agenda) {
	if agenda.AgendaProposer.Empty() {
		return
	}
	store := ctx.KVStore(v.storekey)
	store.Set([]byte(agendaTopic), v.cdc.MustMarshalBinaryBare(agenda))
}
func (v VoteKeeper) getAgenda(ctx sdk.Context, agendaTopic string) types.Agenda {
	store := ctx.KVStore(v.storekey)
	if v.IsTopicPresent(ctx, agendaTopic) {
		return types.NewAgenda()
	}
	bz := store.Get([]byte(agendaTopic))
	var agenda types.Agenda
	v.cdc.MustUnmarshalBinaryBare(bz, &agenda)
	return agenda
}

func (v VoteKeeper) IsTopicPresent(ctx sdk.Context, agendaTopic string) bool {
	store := ctx.KVStore(v.storekey)
	return store.Has([]byte(agendaTopic))
}
