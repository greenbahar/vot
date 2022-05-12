package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"
)

const TypeMsgCreateVote = "create_vote"

var _ sdk.Msg = &MsgCreateVote{}

func NewMsgCreateVote(creator string, question string, options string, days string) *MsgCreateVote {
	intDays, err := strconv.Atoi(days)
	if err != nil {
		panic(err)
	}
	return &MsgCreateVote{
		Creator:  creator,
		Question: question,
		Options:  options,
		Days:     uint32(intDays),
	}
}

func (msg *MsgCreateVote) Route() string {
	return RouterKey
}

func (msg *MsgCreateVote) Type() string {
	return TypeMsgCreateVote
}

func (msg *MsgCreateVote) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateVote) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateVote) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
