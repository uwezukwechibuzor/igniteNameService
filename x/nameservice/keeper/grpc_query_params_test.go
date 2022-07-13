package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testkeeper "github.com/uwezukwechibuzor/nameservice/testutil/keeper"
	"github.com/uwezukwechibuzor/nameservice/x/nameservice/types"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := testkeeper.NameserviceKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	keeper.SetParams(ctx, params)

	response, err := keeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
