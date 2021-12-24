package steak_test

import (
	"testing"

	keepertest "github.com/jemrickrioux/steak/testutil/keeper"
	"github.com/jemrickrioux/steak/testutil/nullify"
	"github.com/jemrickrioux/steak/x/steak"
	"github.com/jemrickrioux/steak/x/steak/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SteakKeeper(t)
	steak.InitGenesis(ctx, *k, genesisState)
	got := steak.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
