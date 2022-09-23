package keeper_test

import (
	"context"
	"testing"

	keepertest "elola/testutil/keeper"
	"elola/x/elola/keeper"
	"elola/x/elola/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ElolaKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
