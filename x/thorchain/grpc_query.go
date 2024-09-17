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

func (s queryServer) RunePool(c context.Context, req *types.QueryRunePoolRequest) (*types.QueryRunePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryRUNEPool(ctx, req)
}

func (s queryServer) RuneProvider(c context.Context, req *types.QueryRuneProviderRequest) (*types.QueryRuneProviderResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryRUNEProvider(ctx, req)
}

func (s queryServer) RuneProviders(c context.Context, req *types.QueryRuneProvidersRequest) (*types.QueryRuneProvidersResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryRUNEProviders(ctx, req)
}

func (s queryServer) MimirValues(c context.Context, req *types.QueryMimirValuesRequest) (*types.QueryMimirValuesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryMimirValues(ctx, req)
}

func (s queryServer) MimirWithKey(c context.Context, req *types.QueryMimirWithKeyRequest) (*types.QueryMimirWithKeyResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryMimirWithKey(ctx, req)
}

func (s queryServer) MimirAdminValues(c context.Context, req *types.QueryMimirAdminValuesRequest) (*types.QueryMimirAdminValuesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryMimirAdminValues(ctx, req)
}

func (s queryServer) MimirNodesAllValues(c context.Context, req *types.QueryMimirNodesAllValuesRequest) (*types.QueryMimirNodesAllValuesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryMimirNodesAllValues(ctx, req)
}

func (s queryServer) MimirNodesValues(c context.Context, req *types.QueryMimirNodesValuesRequest) (*types.QueryMimirNodesValuesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryMimirNodesValues(ctx, req)
}

func (s queryServer) MimirNodeValues(c context.Context, req *types.QueryMimirNodeValuesRequest) (*types.QueryMimirNodeValuesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryMimirNodeValues(ctx, req)
}

func (s queryServer) InboundAddresses(c context.Context, req *types.QueryInboundAddressesRequest) (*types.QueryInboundAddressesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryInboundAddresses(ctx, req)
}

func (s queryServer) Version(c context.Context, req *types.QueryVersionRequest) (*types.QueryVersionResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryVersion(ctx, req)
}

func (s queryServer) Thorname(c context.Context, req *types.QueryThornameRequest) (*types.QueryThornameResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryTHORName(ctx, req)
}

func (s queryServer) Invariant(c context.Context, req *types.QueryInvariantRequest) (*types.QueryInvariantResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryInvariant(ctx, req)
}

func (s queryServer) Invariants(c context.Context, req *types.QueryInvariantsRequest) (*types.QueryInvariantsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryInvariants(ctx, req)
}