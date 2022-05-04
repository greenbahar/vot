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

	// the error isn't handled because the data is set in blockchain. Having error or not we should return the results.
	// returning this is not correct in golang principles:
	// &types.MsgCreateVoteResponse{IdValue: newIndex}, err
	// because in returns in go we either have an error or we return the results. That's the way we decide what to do
	// with the results
	ctx.EventManager().EmitTypedEvents(msg, &types.MsgCreateVoteResponse{IdValue: newIndex})

	return &types.MsgCreateVoteResponse{IdValue: newIndex}, nil
}
