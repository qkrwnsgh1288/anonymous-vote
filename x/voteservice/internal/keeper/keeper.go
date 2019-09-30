package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/qkrwnsgh1288/anonymous-vote/x/voteservice/internal/types"
)

type Keeper struct {
	storekey sdk.StoreKey
	cdc      *codec.Codec
}

func NewKeeper(storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		storekey: storeKey,
		cdc:      cdc,
	}
}

func (v Keeper) SetAgenda(ctx sdk.Context, agendaTopic string, agenda types.Agenda) {
	if agenda.AgendaProposer.Empty() {
		return
	}
	store := ctx.KVStore(v.storekey)
	store.Set([]byte(agendaTopic), v.cdc.MustMarshalBinaryBare(agenda))
}
func (v Keeper) GetAgenda(ctx sdk.Context, agendaTopic string) types.Agenda {
	store := ctx.KVStore(v.storekey)
	if !v.IsTopicPresent(ctx, agendaTopic) {
		return types.NewAgenda()
	}
	bz := store.Get([]byte(agendaTopic))
	var agenda types.Agenda
	v.cdc.MustUnmarshalBinaryBare(bz, &agenda)
	return agenda
}

func (v Keeper) GetAgendaTopic(ctx sdk.Context, agendaTopic string) string {
	return v.GetAgenda(ctx, agendaTopic).AgendaTopic
}

func (v Keeper) IsTopicPresent(ctx sdk.Context, agendaTopic string) bool {
	store := ctx.KVStore(v.storekey)
	return store.Has([]byte(agendaTopic))
}
func (v Keeper) GetTopicsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(v.storekey)
	return sdk.KVStorePrefixIterator(store, []byte{})
}
