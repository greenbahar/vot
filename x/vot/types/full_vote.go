package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/greenbahar/vot/x/vot/rules"
)

func (sv *StoredVote) ParseVote() (*rules.Vote, error) {
	vote, errVote := rules.Parse(sv.Vote)
	if errVote != nil {
		return nil, sdkerrors.Wrapf(errVote, ErrVoteNotParsable.Error())
	}

	return vote, nil
}
