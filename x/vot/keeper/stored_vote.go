package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/greenbahar/vot/x/vot/types"
)

// SetStoredVote set a specific storedVote in the store from its index
func (k Keeper) SetStoredVote(ctx sdk.Context, storedVote types.StoredVote) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredVoteKeyPrefix))
	b := k.cdc.MustMarshal(&storedVote)
	store.Set(types.StoredVoteKey(
		storedVote.Index,
	), b)
}

// GetStoredVote returns a storedVote from its index
func (k Keeper) GetStoredVote(
	ctx sdk.Context,
	index string,

) (val types.StoredVote, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredVoteKeyPrefix))

	b := store.Get(types.StoredVoteKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveStoredVote removes a storedVote from the store
func (k Keeper) RemoveStoredVote(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredVoteKeyPrefix))
	store.Delete(types.StoredVoteKey(
		index,
	))
}

// GetAllStoredVote returns all storedVote
func (k Keeper) GetAllStoredVote(ctx sdk.Context) (list []types.StoredVote) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredVoteKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.StoredVote
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
