package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/greenbahar/vot/x/vot/types"
)

func (k msgServer) SelectVotingOption(goCtx context.Context, msg *types.MsgSelectVotingOption) (*types.MsgSelectVotingOptionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	storedVote, found := k.Keeper.GetStoredVote(ctx, msg.IdValue)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrNotFound, "storedVote object not found", msg.IdValue)
	}

	for _, v := range storedVote.Voters {
		if msg.GetCreator() == v.Voter {
			return nil, sdkerrors.Wrapf(types.ErrDuplication, "you: %s already have participated in this voting: %s", msg.Creator, msg.IdValue)
		}
	}

	vote, err := storedVote.ParseVote()
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrNotParsable, "cannot pars storedVote object", msg.IdValue)
	}

	if voteErr := vote.VoteFor(uint32(msg.SelectedOption)); voteErr != nil {
		return nil, sdkerrors.Wrapf(types.ErrInternalError, "user %s cannot submit the voting option for vote: %s", msg.Creator, msg.IdValue)
	}

	storedVote.Voters = append(storedVote.Voters, &types.StoredVote_Voter{Voter: msg.Creator})
	storedVote.Vote, err = vote.String()
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrNotSerialized, "vote object cannot be serialized: %s", msg.IdValue)
	}

	k.Keeper.SetStoredVote(ctx, storedVote)

	result := types.MsgSelectVotingOptionResponse_Result{}
	results := make([]*types.MsgSelectVotingOptionResponse_Result, 0)
	for _, item := range vote.GetVoteOptionsWithCounter() {
		result.Option = item.Option
		result.Counter = int64(item.Counter)
		results = append(results, &result)
	}

	ctx.EventManager().EmitTypedEvents(msg, &types.MsgSelectVotingOptionResponse{Results: results})

	return &types.MsgSelectVotingOptionResponse{
		IdValue: msg.IdValue,
		Results: results,
	}, nil
}
