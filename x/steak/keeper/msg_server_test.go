package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/jemrickrioux/steak/testutil/keeper"
	"github.com/jemrickrioux/steak/x/steak/keeper"
	"github.com/jemrickrioux/steak/x/steak/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.SteakKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
