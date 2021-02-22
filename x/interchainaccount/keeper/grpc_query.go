package keeper

import (
	"github.com/seantking/interchain-account/x/interchainaccount/types"
)

var _ types.QueryServer = Keeper{}
