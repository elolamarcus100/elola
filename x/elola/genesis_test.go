package elola_test

import (
	"testing"

	keepertest "elola/testutil/keeper"
	"elola/testutil/nullify"
	"elola/x/elola"
	"elola/x/elola/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ElolaKeeper(t)
	elola.InitGenesis(ctx, *k, genesisState)
	got := elola.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
