package keeper

import (
	"github.com/FairBlock/fairyring/x/pep/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetChannel(
	ctx sdk.Context,
) string {
	store := ctx.KVStore(k.storeKey)
	return string(store.Get(types.ChannelKey))
}

func (k Keeper) SetChannel(
	ctx sdk.Context,
	channelID string,
) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.ChannelKey, []byte(channelID))
}
