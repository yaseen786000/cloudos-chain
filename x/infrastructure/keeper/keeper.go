package keeper

import (
	"cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"
)

type Keeper struct {
	storeService store.KVStoreService
	cdc          codec.BinaryCodec
}

func NewKeeper(
	storeService store.KVStoreService,
	cdc codec.BinaryCodec,
) *Keeper {

	return &Keeper{
		storeService: storeService,
		cdc:          cdc,
	}
}
