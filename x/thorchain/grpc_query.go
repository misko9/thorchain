package thorchain

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"gitlab.com/thorchain/thornode/x/thorchain/types"
)

type queryServer struct {
	mgr Manager
}

var _ types.QueryServer = queryServer{}

func NewQueryServerImpl(mgr Manager) types.QueryServer {
	return queryServer{mgr: mgr}
}

func (s queryServer) Pool(c context.Context, req *types.QueryPoolRequest) (*types.QueryPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryPool(ctx, req)
}

func (s queryServer) Pools(c context.Context, req *types.QueryPoolsRequest) (*types.QueryPoolsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryPools(ctx, req)
}

func (s queryServer) DerivedPool(c context.Context, req *types.QueryDerivedPoolRequest) (*types.QueryDerivedPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryDerivedPool(ctx, req)
}

func (s queryServer) DerivedPools(c context.Context, req *types.QueryDerivedPoolsRequest) (*types.QueryDerivedPoolsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryDerivedPools(ctx, req)
}

func (s queryServer) LiquidityProvider(c context.Context, req *types.QueryLiquidityProviderRequest) (*types.QueryLiquidityProviderResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryLiquidityProvider(ctx, req)
}

func (s queryServer) LiquidityProviders(c context.Context, req *types.QueryLiquidityProvidersRequest) (*types.QueryLiquidityProvidersResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryLiquidityProviders(ctx, req)
}

func (s queryServer) Saver(c context.Context, req *types.QuerySaverRequest) (*types.QuerySaverResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.querySaver(ctx, req)
}

func (s queryServer) Savers(c context.Context, req *types.QuerySaversRequest) (*types.QuerySaversResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.querySavers(ctx, req)
}

func (s queryServer) Borrower(c context.Context, req *types.QueryBorrowerRequest) (*types.QueryBorrowerResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryBorrower(ctx, req)
}

func (s queryServer) Borrowers(c context.Context, req *types.QueryBorrowersRequest) (*types.QueryBorrowersResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryBorrowers(ctx, req)
}

func (s queryServer) TradeUnit(c context.Context, req *types.QueryTradeUnitRequest) (*types.QueryTradeUnitResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryTradeUnit(ctx, req)
}

func (s queryServer) TradeUnits(c context.Context, req *types.QueryTradeUnitsRequest) (*types.QueryTradeUnitsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryTradeUnits(ctx, req)
}