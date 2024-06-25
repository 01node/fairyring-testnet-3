package keeper

import (
	"context"
	"strconv"

	"github.com/Fairblock/fairyring/x/pep/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// LatestHeight returns the last registered height
func (k Keeper) LatestHeight(goCtx context.Context, req *types.QueryLatestHeightRequest) (*types.QueryLatestHeightResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	height, err := strconv.ParseUint(k.GetLatestHeight(ctx), 10, 64)

	if err != nil {
		return nil, err
	}

	return &types.QueryLatestHeightResponse{
		Height: height,
	}, nil
}
