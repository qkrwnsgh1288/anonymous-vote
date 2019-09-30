package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"os"
	"path/filepath"
	"testing"
)

func TestLevelDB(t *testing.T) {
	const dbName = "blockstore.db"
	const StoreKey = "voteservice"
	dbPath := filepath.Join(os.Getenv("GOPATH"), "projects", ".voted", "data", dbName)
	fmt.Println(dbPath)

	keys := sdk.NewKVStoreKeys(StoreKey)
	fmt.Println(keys)

	ctx := sdk.Context{}
	store := ctx.KVStore(keys[StoreKey])
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	for ; iterator.Valid(); iterator.Next() {
		fmt.Println(iterator.Key(), iterator.Value())
	}

}
