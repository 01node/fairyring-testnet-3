package keeper

import (
	"context"
	"github.com/Fairblock/fairyring/x/pep/types"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// EncryptedTxAll returns the paginated list of all encrypted Txs
func (k Keeper) EncryptedTxAll(c context.Context, req *types.QueryAllEncryptedTxRequest) (*types.QueryAllEncryptedTxResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var encryptedTxs []types.EncryptedTxArray
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	encryptedTxStore := prefix.NewStore(store, types.KeyPrefix(types.EncryptedTxKeyPrefix))

	pageRes, err := query.Paginate(encryptedTxStore, req.Pagination, func(key []byte, value []byte) error {
		var encryptedTxArr types.EncryptedTxArray
		if err := k.cdc.Unmarshal(value, &encryptedTxArr); err != nil {
			return err
		}

		encryptedTxs = append(encryptedTxs, encryptedTxArr)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllEncryptedTxResponse{EncryptedTxArray: encryptedTxs, Pagination: pageRes}, nil
}

// EncryptedTxAllFromHeight returns all the encrypted TXs for a particular height
func (k Keeper) EncryptedTxAllFromHeight(c context.Context, req *types.QueryAllEncryptedTxFromHeightRequest) (*types.QueryAllEncryptedTxFromHeightResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	val := k.GetEncryptedTxAllFromHeight(ctx, req.TargetHeight)

	return &types.QueryAllEncryptedTxFromHeightResponse{EncryptedTxArray: val}, nil
}

// EncryptedTx returns a singe encrypted Tx by index
func (k Keeper) EncryptedTx(c context.Context, req *types.QueryGetEncryptedTxRequest) (*types.QueryGetEncryptedTxResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetEncryptedTx(
		ctx,
		req.TargetHeight,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetEncryptedTxResponse{EncryptedTx: val}, nil
}
