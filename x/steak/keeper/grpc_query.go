package keeper

import (
	"github.com/jemrickrioux/steak/x/steak/types"
)

var _ types.QueryServer = Keeper{}
