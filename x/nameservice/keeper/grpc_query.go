package keeper

import (
	"github.com/uwezukwechibuzor/nameservice/x/nameservice/types"
)

var _ types.QueryServer = Keeper{}
