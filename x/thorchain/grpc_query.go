package thorchain

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"gitlab.com/thorchain/thornode/x/thorchain/types"
)

type queryServer struct {
	mgr *Mgrs
}

var _ types.QueryServer = queryServer{}

func NewQueryServerImpl(mgr *Mgrs) types.QueryServer {
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

func (s queryServer) TradeAccount(c context.Context, req *types.QueryTradeAccountRequest) (*types.QueryTradeAccountsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryTradeAccount(ctx, req)
}

func (s queryServer) TradeAccounts(c context.Context, req *types.QueryTradeAccountsRequest) (*types.QueryTradeAccountsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryTradeAccounts(ctx, req)
}

func (s queryServer) Node(c context.Context, req *types.QueryNodeRequest) (*types.QueryNodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryNode(ctx, req)
}

func (s queryServer) Nodes(c context.Context, req *types.QueryNodesRequest) (*types.QueryNodesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryNodes(ctx, req)
}

func (s queryServer) PoolSlip(c context.Context, req *types.QueryPoolSlipRequest) (*types.QueryPoolSlipsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryPoolSlip(ctx, req)
}

func (s queryServer) PoolSlips(c context.Context, req *types.QueryPoolSlipsRequest) (*types.QueryPoolSlipsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryPoolSlips(ctx, req)
}

func (s queryServer) OutboundFee(c context.Context, req *types.QueryOutboundFeeRequest) (*types.QueryOutboundFeesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryOutboundFees(ctx, req.Asset)
}

func (s queryServer) OutboundFees(c context.Context, req *types.QueryOutboundFeesRequest) (*types.QueryOutboundFeesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryOutboundFees(ctx, "")
}

func (s queryServer) StreamingSwap(c context.Context, req *types.QueryStreamingSwapRequest) (*types.QueryStreamingSwapResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryStreamingSwap(ctx, req)
}

func (s queryServer) StreamingSwaps(c context.Context, req *types.QueryStreamingSwapsRequest) (*types.QueryStreamingSwapsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryStreamingSwaps(ctx, req)
}

func (s queryServer) Ban(c context.Context, req *types.QueryBanRequest) (*types.BanVoter, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryBan(ctx, req)
}

func (s queryServer) Ragnarok(c context.Context, req *types.QueryRagnarokRequest) (*types.QueryRagnarokResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryRagnarok(ctx, req)
}