package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/greenbahar/vot/x/vot/types"
)

// SetNextVote set nextVote in the store
func (k Keeper) SetNextVote(ctx sdk.Context, nextVote types.NextVote) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NextVoteKey))
	b := k.cdc.MustMarshal(&nextVote)
	store.Set([]byte{0}, b)
}

// GetNextVote returns nextVote
func (k Keeper) GetNextVote(ctx sdk.Context) (val types.NextVote, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NextVoteKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNextVote removes nextVote from the store
func (k Keeper) RemoveNextVote(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NextVoteKey))
	store.Delete([]byte{0})
}
