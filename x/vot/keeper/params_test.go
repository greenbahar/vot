package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/greenbahar/vot/testutil/keeper"
	"github.com/greenbahar/vot/x/vot/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.VotKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
