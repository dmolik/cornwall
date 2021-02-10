package keeper

import (
	"github.com/dmolik/cornwall/x/cornwall/types"
)

var _ types.QueryServer = Keeper{}
