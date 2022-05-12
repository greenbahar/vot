package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/greenbahar/vot/x/vot/types"
)

func (k msgServer) SelectVotingOption(goCtx context.Context, msg *types.MsgSelectVotingOption) (*types.MsgSelectVotingOptionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.Logger(ctx).Info("user: %s created tx select-voting-option", msg.Creator)

	storedVote, found := k.Keeper.GetStoredVote(ctx, msg.IdValue)
	if !found {
		k.Logger(ctx).Error("storedVote object not found", msg.IdValue)
		return nil, sdkerrors.Wrapf(types.ErrNotFound, "storedVote object not found", msg.IdValue)
	}
	k.Logger(ctx).Info("this is stored vote: %v", storedVote)

	for _, v := range storedVote.Voters {
		if msg.Creator == v.Voter {
			return nil, sdkerrors.Wrapf(types.ErrDuplication, "you: %s already have participated in this voting: %s", msg.Creator, msg.IdValue)
		}
	}

	vote, err := storedVote.ParseVote()
	if err != nil {
		k.Logger(ctx).Error("cannot pars storedVote object", msg.IdValue)
		return nil, sdkerrors.Wrapf(types.ErrNotParsable, "cannot pars storedVote object", msg.IdValue)
	}
	k.Logger(ctx).Info("here is the vote object of storedVote: %v", vote)

	if voteErr := vote.VoteFor(uint32(msg.SelectedOption)); voteErr != nil {
		k.Logger(ctx).Error("user %s cannot submit the voting option for vote: %s", msg.Creator, msg.IdValue)
		return nil, sdkerrors.Wrapf(types.ErrInternalError, "user %s cannot submit the voting option for vote: %s", msg.Creator, msg.IdValue)
	}

	serializedStoredVote, SerializeErr := vote.String()
	if SerializeErr != nil {
		k.Logger(ctx).Error("vote object cannot be serialized: %s", msg.IdValue)
		return nil, sdkerrors.Wrapf(types.ErrNotSerialized, "vote object cannot be serialized: %s", msg.IdValue)
	}
	storedVote.Vote = serializedStoredVote
	storedVote.Voters = append(storedVote.Voters, &types.StoredVote_Voter{Voter: msg.Creator})
	k.Keeper.SetStoredVote(ctx, storedVote)
	k.Logger(ctx).Info("modified storedVote object: %v", storedVote)

	result := types.MsgSelectVotingOptionResponse_Result{}
	results := make([]*types.MsgSelectVotingOptionResponse_Result, 0)
	for _, item := range vote.GetVoteOptionsWithCounter() {
		result.Option = item.Option
		result.Counter = int64(item.Counter)
		results = append(results, &result)
	}
	k.Logger(ctx).Info("here is the result to return to the user: %v", result)

	//ctx.EventManager().EmitTypedEvents(msg, &types.MsgSelectVotingOptionResponse{Results: results})

	return &types.MsgSelectVotingOptionResponse{
		IdValue: msg.IdValue,
		Results: results,
	}, nil
}
