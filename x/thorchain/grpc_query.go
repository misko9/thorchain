package thorchain

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"gitlab.com/thorchain/thornode/common/cosmos"
	"gitlab.com/thorchain/thornode/x/thorchain/types"
)

type queryServer struct {
	mgr *Mgrs
	kbs cosmos.KeybaseStore
}

var _ types.QueryServer = queryServer{}

func NewQueryServerImpl(mgr *Mgrs, kbs cosmos.KeybaseStore) types.QueryServer {
	return queryServer{mgr: mgr, kbs: kbs}
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

func (s queryServer) Network(c context.Context, req *types.QueryNetworkRequest) (*types.QueryNetworkResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryNetwork(ctx, req)
}

func (s queryServer) BalanceModule(c context.Context, req *types.QueryBalanceModuleRequest) (*types.QueryBalanceModuleResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryBalanceModule(ctx, req)
}

func (s queryServer) QuoteSwap(c context.Context, req *types.QueryQuoteSwapRequest) (*types.QueryQuoteSwapResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryQuoteSwap(ctx, req)
}

func (s queryServer) QuoteSaverDeposit(c context.Context, req *types.QueryQuoteSaverDepositRequest) (*types.QueryQuoteSaverDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryQuoteSaverDeposit(ctx, req)
}

func (s queryServer) QuoteSaverWithdraw(c context.Context, req *types.QueryQuoteSaverWithdrawRequest) (*types.QueryQuoteSaverWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryQuoteSaverWithdraw(ctx, req)
}

func (s queryServer) QuoteLoanOpen(c context.Context, req *types.QueryQuoteLoanOpenRequest) (*types.QueryQuoteLoanOpenResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryQuoteLoanOpen(ctx, req)
}

func (s queryServer) QuoteLoanClose(c context.Context, req *types.QueryQuoteLoanCloseRequest) (*types.QueryQuoteLoanCloseResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryQuoteLoanClose(ctx, req)
}

func (s queryServer) ConstantValues(c context.Context, req *types.QueryConstantValuesRequest) (*types.QueryConstantValuesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryConstantValues(ctx, req)
}

func (s queryServer) SwapQueue(c context.Context, req *types.QuerySwapQueueRequest) (*types.QuerySwapQueueResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.querySwapQueue(ctx, req)
}

func (s queryServer) LastBlocks(c context.Context, req *types.QueryLastBlocksRequest) (*types.QueryLastBlocksResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryLastBlockHeights(ctx, "")
}

func (s queryServer) ChainsLastBlock(c context.Context, req *types.QueryChainsLastBlockRequest) (*types.QueryLastBlocksResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryLastBlockHeights(ctx, req.Chain)
}

func (s queryServer) Vault(c context.Context, req *types.QueryVaultRequest) (*types.QueryVaultResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryVault(ctx, req)
}

func (s queryServer) AsgardVaults(c context.Context, req *types.QueryAsgardVaultsRequest) (*types.QueryAsgardVaultsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryAsgardVaults(ctx, req)
}

func (s queryServer) VaultsPubkeys(c context.Context, req *types.QueryVaultsPubkeysRequest) (*types.QueryVaultsPubkeysResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryVaultsPubkeys(ctx, req)
}

func (s queryServer) TxStages(c context.Context, req *types.QueryTxStagesRequest) (*types.QueryTxStagesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryTxStages(ctx, req)
}

func (s queryServer) TxStatus(c context.Context, req *types.QueryTxStatusRequest) (*types.QueryTxStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryTxStatus(ctx, req)
}

func (s queryServer) TxVoters(c context.Context, req *types.QueryTxVotersRequest) (*types.ObservedTxVoter, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryTxVoters(ctx, req)
}

func (s queryServer) TxVotersOld(c context.Context, req *types.QueryTxVotersRequest) (*types.ObservedTxVoter, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryTxVoters(ctx, req)
}

func (s queryServer) Tx(c context.Context, req *types.QueryTxRequest) (*types.QueryTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryTx(ctx, req)
}

func (s queryServer) Clout(c context.Context, req *types.QuerySwapperCloutRequest) (*types.SwapperClout, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.querySwapperClout(ctx, req)
}

func (s queryServer) Queue(c context.Context, req *types.QueryQueueRequest) (*types.QueryQueueResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryQueue(ctx, req)
}

func (s queryServer) ScheduledOutbound(c context.Context, req *types.QueryScheduledOutboundRequest) (*types.QueryOutboundResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryScheduledOutbound(ctx, req)
}

func (s queryServer) PendingOutbound(c context.Context, req *types.QueryPendingOutboundRequest) (*types.QueryOutboundResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryPendingOutbound(ctx, req)
}

func (s queryServer) Block(c context.Context, req *types.QueryBlockRequest) (*types.QueryBlockResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryBlock(ctx, req)
}

func (s queryServer) TssKeygenMetric(c context.Context, req *types.QueryTssKeygenMetricRequest) (*types.QueryTssKeygenMetricResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryTssKeygenMetric(ctx, req)
}

func (s queryServer) TssMetric(c context.Context, req *types.QueryTssMetricRequest) (*types.QueryTssMetricResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryTssMetric(ctx, req)
}

func (s queryServer) Keysign(c context.Context, req *types.QueryKeysignRequest) (*types.QueryKeysignResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryKeysign(ctx, req.Height, "")
}

func (s queryServer) KeysignPubkey(c context.Context, req *types.QueryKeysignPubkeyRequest) (*types.QueryKeysignResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryKeysign(ctx, req.Height, req.PubKey)
}