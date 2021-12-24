package keeper_test

import (
	"testing"

	testkeeper "github.com/jemrickrioux/steak/testutil/keeper"
	"github.com/jemrickrioux/steak/x/steak/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.SteakKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
