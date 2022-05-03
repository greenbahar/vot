package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/vot module sentinel errors
var (
	ErrSample        = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrNotParsable   = sdkerrors.Register(ModuleName, 1101, "object cannot be parsed: %S")
	ErrNotSerialized = sdkerrors.Register(ModuleName, 1102, "object cannot be serialized: %s")
	ErrNotFound      = sdkerrors.Register(ModuleName, 1103, "object by id not found: %s")
	ErrDuplication   = sdkerrors.Register(ModuleName, 1104, "Duplicate operation: %s")
	ErrInternalError = sdkerrors.Register(ModuleName, 1105, "internal error: %s")
)
