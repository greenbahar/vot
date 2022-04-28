package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/greenbahar/vot/x/vot/types"
)

func (k msgServer) CreateVote(goCtx context.Context, msg *types.MsgCreateVote) (*types.MsgCreateVoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateVoteResponse{}, nil
}
