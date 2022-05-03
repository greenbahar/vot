package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSelectVotingOption = "select_voting_option"

var _ sdk.Msg = &MsgSelectVotingOption{}

func NewMsgSelectVotingOption(creator string, idValue string, selectedOption uint64) *MsgSelectVotingOption {
	return &MsgSelectVotingOption{
		Creator:        creator,
		IdValue:        idValue,
		SelectedOption: selectedOption,
	}
}

func (msg *MsgSelectVotingOption) Route() string {
	return RouterKey
}

func (msg *MsgSelectVotingOption) Type() string {
	return TypeMsgSelectVotingOption
}

func (msg *MsgSelectVotingOption) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSelectVotingOption) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSelectVotingOption) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
