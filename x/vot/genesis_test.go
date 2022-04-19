package vot_test

import (
	"testing"

	keepertest "github.com/greenbahar/vot/testutil/keeper"
	"github.com/greenbahar/vot/testutil/nullify"
	"github.com/greenbahar/vot/x/vot"
	"github.com/greenbahar/vot/x/vot/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:	types.DefaultParams(),
		
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VotKeeper(t)
	vot.InitGenesis(ctx, *k, genesisState)
	got := vot.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	

	// this line is used by starport scaffolding # genesis/test/assert
}
