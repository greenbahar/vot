package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/greenbahar/vot/testutil/keeper"
	"github.com/greenbahar/vot/testutil/nullify"
	"github.com/greenbahar/vot/x/vot/keeper"
	"github.com/greenbahar/vot/x/vot/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNStoredVote(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.StoredVote {
	items := make([]types.StoredVote, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetStoredVote(ctx, items[i])
	}
	return items
}

func TestStoredVoteGet(t *testing.T) {
	keeper, ctx := keepertest.VotKeeper(t)
	items := createNStoredVote(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetStoredVote(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestStoredVoteRemove(t *testing.T) {
	keeper, ctx := keepertest.VotKeeper(t)
	items := createNStoredVote(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveStoredVote(ctx,
			item.Index,
		)
		_, found := keeper.GetStoredVote(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestStoredVoteGetAll(t *testing.T) {
	keeper, ctx := keepertest.VotKeeper(t)
	items := createNStoredVote(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllStoredVote(ctx)),
	)
}
