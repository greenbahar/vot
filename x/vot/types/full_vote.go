package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/greenbahar/vot/x/vot/rules"
)

// this full_vote file contains helper function related to rules of the voting application

func (sv *StoredVote) ParseVote() (*rules.Vote, error) {
	vote, errVote := rules.Parse(sv.Vote)
	if errVote != nil {
		return nil, sdkerrors.Wrapf(errVote, ErrNotParsable.Error())
	}

	return vote, nil
}
