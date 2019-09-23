package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/qkrwnsgh1288/vote-dapp/x/voteservice/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints supported by the voteservice Querier
const (
	QueryAgenda = "agenda"
	QueryTopics = "topics"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryAgenda:
			return queryAgenda(ctx, path[1:], req, keeper)
		case QueryTopics:
			return queryTopics(ctx, req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown voteservice query endpoint")
		}
	}
}

func queryAgenda(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	agenda := keeper.GetAgenda(ctx, path[0])
	res, err := codec.MarshalJSONIndent(keeper.cdc, agenda)
	if err != nil {
		panic("could not marshal result to JSON")
	}
	return res, nil
}
func queryTopics(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	var topicList types.QueryResTopics
	iterator := keeper.GetTopicsIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		topicList = append(topicList, string(iterator.Key()))
	}
	res, err := codec.MarshalJSONIndent(keeper.cdc, topicList)
	if err != nil {
		panic("could not marshal result to JSON")
	}
	return res, nil
}
