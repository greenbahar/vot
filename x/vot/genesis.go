package vot

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/greenbahar/vot/x/vot/keeper"
	"github.com/greenbahar/vot/x/vot/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.NextVote != nil {
		k.SetNextVote(ctx, *genState.NextVote)
	}
	// Set all the storedVote
	for _, elem := range genState.StoredVoteList {
		k.SetStoredVote(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all nextVote
	nextVote, found := k.GetNextVote(ctx)
	if found {
		genesis.NextVote = &nextVote
	}
	genesis.StoredVoteList = k.GetAllStoredVote(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
