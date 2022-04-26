package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/greenbahar/vot/x/vot/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) StoredVoteAll(c context.Context, req *types.QueryAllStoredVoteRequest) (*types.QueryAllStoredVoteResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var storedVotes []types.StoredVote
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	storedVoteStore := prefix.NewStore(store, types.KeyPrefix(types.StoredVoteKeyPrefix))

	pageRes, err := query.Paginate(storedVoteStore, req.Pagination, func(key []byte, value []byte) error {
		var storedVote types.StoredVote
		if err := k.cdc.Unmarshal(value, &storedVote); err != nil {
			return err
		}

		storedVotes = append(storedVotes, storedVote)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllStoredVoteResponse{StoredVote: storedVotes, Pagination: pageRes}, nil
}

func (k Keeper) StoredVote(c context.Context, req *types.QueryGetStoredVoteRequest) (*types.QueryGetStoredVoteResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetStoredVote(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetStoredVoteResponse{StoredVote: val}, nil
}
