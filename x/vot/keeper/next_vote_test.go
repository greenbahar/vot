package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/greenbahar/vot/testutil/keeper"
	"github.com/greenbahar/vot/testutil/nullify"
	"github.com/greenbahar/vot/x/vot/keeper"
	"github.com/greenbahar/vot/x/vot/types"
)

func createTestNextVote(keeper *keeper.Keeper, ctx sdk.Context) types.NextVote {
	item := types.NextVote{}
	keeper.SetNextVote(ctx, item)
	return item
}

func TestNextVoteGet(t *testing.T) {
	keeper, ctx := keepertest.VotKeeper(t)
	item := createTestNextVote(keeper, ctx)
	rst, found := keeper.GetNextVote(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestNextVoteRemove(t *testing.T) {
	keeper, ctx := keepertest.VotKeeper(t)
	createTestNextVote(keeper, ctx)
	keeper.RemoveNextVote(ctx)
	_, found := keeper.GetNextVote(ctx)
	require.False(t, found)
}
