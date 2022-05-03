package keeper

import (
	"context"
	"github.com/greenbahar/vot/x/vot/rules"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/greenbahar/vot/x/vot/types"
)

func (k msgServer) CreateVote(goCtx context.Context, msg *types.MsgCreateVote) (*types.MsgCreateVoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	nextVote, found := k.Keeper.GetNextVote(ctx)
	if !found {
		panic("NextVote not found")
	}
	newIndex := strconv.FormatUint(nextVote.IdValue, 10)

	options := rules.ParsOptions(msg.Options)
	newVot := rules.NewVote(msg.Question, msg.Days, options...)
	newVoteSerialized, err := newVot.String()
	if err != nil {
		panic("cannot serialize the Vote object")
	}

	storedVot := types.StoredVote{
		Index:   newIndex,
		Vote:    newVoteSerialized,
		Creator: msg.Creator,
	}

	k.Keeper.SetStoredVote(ctx, storedVot)
	nextVote.IdValue++
	k.Keeper.SetNextVote(ctx, nextVote)

	return &types.MsgCreateVoteResponse{IdValue: newIndex}, nil
}
