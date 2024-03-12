package keeper

import (
	"github.com/Fairblock/fairyring/x/pep/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AppendEncryptedTx append a specific encryptedTx in the store
func (k Keeper) AppendEncryptedTx(
	ctx sdk.Context,
	encryptedTx types.EncryptedTx,
) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EncryptedTxKeyPrefix))
	var allTxsFromHeight types.EncryptedTxArray
	b := store.Get(types.EncryptedTxAllFromHeightKey(
		encryptedTx.TargetHeight,
	))

	k.cdc.MustUnmarshal(b, &allTxsFromHeight)

	encryptedTx.Index = uint64(len(allTxsFromHeight.GetEncryptedTx()))

	allTxsFromHeight.EncryptedTx = append(allTxsFromHeight.EncryptedTx, encryptedTx)

	parsedEncryptedTxArr := k.cdc.MustMarshal(&allTxsFromHeight)

	store.Set(types.EncryptedTxAllFromHeightKey(
		encryptedTx.TargetHeight,
	), parsedEncryptedTxArr)

	return encryptedTx.Index
}

// SetEncryptedTx add a specific encryptedTx in the store from its index
func (k Keeper) SetEncryptedTx(
	ctx sdk.Context,
	height uint64,
	encryptedTxArr types.EncryptedTxArray,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EncryptedTxKeyPrefix))

	parsedEncryptedTxArr := k.cdc.MustMarshal(&encryptedTxArr)

	store.Set(types.EncryptedTxAllFromHeightKey(
		height,
	), parsedEncryptedTxArr)
}

func (k Keeper) SetEncryptedTxProcessedHeight(
	ctx sdk.Context,
	height uint64,
	index uint64,
	processedHeight uint64,
) {
	arr := k.GetEncryptedTxAllFromHeight(ctx, height)

	if index >= uint64(len(arr.EncryptedTx)) {
		return
	}

	arr.EncryptedTx[index].ProcessedAtChainHeight = processedHeight

	k.SetEncryptedTx(ctx, height, arr)
}

func (k Keeper) SetAllEncryptedTxExpired(
	ctx sdk.Context,
	height uint64,
) {
	arr := k.GetEncryptedTxAllFromHeight(ctx, height)

	for i := range arr.EncryptedTx {
		arr.EncryptedTx[i].Expired = true
	}

	k.SetEncryptedTx(ctx, height, arr)
}

// GetEncryptedTx returns a encryptedTx from its index
func (k Keeper) GetEncryptedTx(
	ctx sdk.Context,
	targetHeight uint64,
	index uint64,

) (val types.EncryptedTx, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EncryptedTxKeyPrefix))

	b := store.Get(types.EncryptedTxAllFromHeightKey(
		targetHeight,
	))
	if b == nil {
		return val, false
	}

	var arr types.EncryptedTxArray
	k.cdc.MustUnmarshal(b, &arr)

	if uint64(len(arr.GetEncryptedTx())) <= index {
		return val, false
	}

	return arr.GetEncryptedTx()[index], true
}

// GetEncryptedTxAllFromHeight returns all encryptedTx from the height provided
func (k Keeper) GetEncryptedTxAllFromHeight(
	ctx sdk.Context,
	targetHeight uint64,

) types.EncryptedTxArray {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EncryptedTxKeyPrefix))

	b := store.Get(types.EncryptedTxAllFromHeightKey(
		targetHeight,
	))

	var arr types.EncryptedTxArray
	k.cdc.MustUnmarshal(b, &arr)

	return arr
}

// GetAllEncryptedArray returns the list of all encrypted txs
func (k Keeper) GetAllEncryptedArray(ctx sdk.Context) (arr []types.EncryptedTxArray) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EncryptedTxKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.EncryptedTxArray
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		arr = append(arr, val)
	}

	return
}

// RemoveEncryptedTx removes a encryptedTx from the store
func (k Keeper) RemoveEncryptedTx(
	ctx sdk.Context,
	targetHeight uint64,
	index uint64,
) {
	arr := k.GetEncryptedTxAllFromHeight(ctx, targetHeight)

	if index >= uint64(len(arr.EncryptedTx)) {
		return
	}

	arr.EncryptedTx = append(arr.EncryptedTx[:index], arr.EncryptedTx[index+1:]...)

	k.SetEncryptedTx(ctx, targetHeight, arr)
}

// RemoveAllEncryptedTxFromHeight removes all encryptedTx from the store for a particular height
func (k Keeper) RemoveAllEncryptedTxFromHeight(
	ctx sdk.Context,
	targetHeight uint64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EncryptedTxKeyPrefix))
	store.Delete(types.EncryptedTxAllFromHeightKey(
		targetHeight,
	))
}
