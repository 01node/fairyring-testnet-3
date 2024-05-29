package keeper

import (
	"cosmossdk.io/math"
	"github.com/Fairblock/fairyring/x/keyshare/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.KeyExpiry(ctx),
		k.TrustedAddresses(ctx),
		k.MinimumBonded(ctx),
		k.SlashFractionNoKeyshare(ctx),
		k.SlashFractionWrongKeyshare(ctx),
		k.MaxIdledBlock(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// KeyExpiry returns the KeyExpiry param
func (k Keeper) KeyExpiry(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyKeyExpiry, &res)
	return
}

// TrustedAddresses returns the TrustedAddresses param
func (k Keeper) TrustedAddresses(ctx sdk.Context) (res []string) {
	k.paramstore.Get(ctx, types.KeyTrustedAddresses, &res)
	return
}

// MinimumBonded returns the MinimumBonded param
func (k Keeper) MinimumBonded(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyMinimumBonded, &res)
	return
}

// SlashFractionNoKeyshare returns the SlashFractionNoKeyshare param
func (k Keeper) SlashFractionNoKeyshare(ctx sdk.Context) (res math.LegacyDec) {
	k.paramstore.Get(ctx, types.KeySlashFractionNoKeyShare, &res)
	return
}

// SlashFractionWrongKeyshare returns the SlashFractionWrongKeyshare param
func (k Keeper) SlashFractionWrongKeyshare(ctx sdk.Context) (res math.LegacyDec) {
	k.paramstore.Get(ctx, types.KeySlashFractionWrongKeyShare, &res)
	return
}

// MaxIdledBlock returns the MaxIdledBlock param
func (k Keeper) MaxIdledBlock(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyMaxIdledBlock, &res)
	return
}
