package keeperv1

import (
	"context"

	"gitlab.com/thorchain/thornode/x/thorchain/keeper"
	"gitlab.com/thorchain/thornode/x/thorchain/types"
)

type msgServer struct {
	k keeper.Keeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the module MsgServer interface.
func NewMsgServerImpl(keeper keeper.Keeper) types.MsgServer {
	return &msgServer{k: keeper}
}

func (ms msgServer) AddLiquidity(ctx context.Context, msg *types.MsgAddLiquidity) (*types.MsgEmpty, error) {
	
}
func (ms msgServer) Ban(ctx context.Context, msg *types.MsgBan) (*types.MsgEmpty, error)
func (ms msgServer) Bond(ctx context.Context, msg *types.MsgBond) (*types.MsgEmpty, error)
func (ms msgServer) Consolidate(ctx context.Context, msg *types.MsgConsolidate) (*types.MsgEmpty, error)
func (ms msgServer) Deposit(ctx context.Context, msg *types.MsgDeposit) (*types.MsgEmpty, error)
func (ms msgServer) Donate(ctx context.Context, msg *types.MsgDonate) (*types.MsgEmpty, error)
func (ms msgServer) ErrataTx(ctx context.Context, msg *types.MsgErrataTx) (*types.MsgEmpty, error)
func (ms msgServer) Leave(ctx context.Context, msg *types.MsgLeave) (*types.MsgEmpty, error)
func (ms msgServer) LoanOpen(ctx context.Context, msg *types.MsgLoanOpen) (*types.MsgEmpty, error)
func (ms msgServer) LoanRepayment(ctx context.Context, msg *types.MsgLoanRepayment) (*types.MsgEmpty, error)
func (ms msgServer) ManageThorname(ctx context.Context, msg *types.MsgManageTHORName) (*types.MsgEmpty, error)
func (ms msgServer) Migrate(ctx context.Context, msg *types.MsgMigrate) (*types.MsgEmpty, error)
func (ms msgServer) Mimir(ctx context.Context, msg *types.MsgMimir) (*types.MsgEmpty, error)
func (ms msgServer) NetworkFee(ctx context.Context, msg *types.MsgNetworkFee) (*types.MsgEmpty, error)
func (ms msgServer) NodePauseChain(ctx context.Context, msg *types.MsgNodePauseChain) (*types.MsgEmpty, error)
func (ms msgServer) Noop(ctx context.Context, msg *types.MsgNoOp) (*types.MsgEmpty, error)
func (ms msgServer) ObservedTxIn(ctx context.Context, msg *types.MsgObservedTxIn) (*types.MsgEmpty, error)
func (ms msgServer) ObservedTxOut(ctx context.Context, msg *types.MsgObservedTxOut) (*types.MsgEmpty, error)
func (ms msgServer) Ragnarok(ctx context.Context, msg *types.MsgRagnarok) (*types.MsgEmpty, error)
func (ms msgServer) RefundTx(ctx context.Context, msg *types.MsgRefundTx) (*types.MsgEmpty, error)
func (ms msgServer) ReserveContributor(ctx context.Context, msg *types.MsgReserveContributor) (*types.MsgEmpty, error)
func (ms msgServer) RunePoolDeposit(ctx context.Context, msg *types.MsgRunePoolDeposit) (*types.MsgEmpty, error)
func (ms msgServer) RunePoolWithdraw(ctx context.Context, msg *types.MsgRunePoolWithdraw) (*types.MsgEmpty, error)
func (ms msgServer) ThorSend(ctx context.Context, msg *types.MsgSend) (*types.MsgEmpty, error)
func (ms msgServer) SetIPAddress(ctx context.Context, msg *types.MsgSetIPAddress) (*types.MsgEmpty, error)
func (ms msgServer) SetNodeKeys(ctx context.Context, msg *types.MsgSetNodeKeys) (*types.MsgEmpty, error)
func (ms msgServer) Solvency(ctx context.Context, msg *types.MsgSolvency) (*types.MsgEmpty, error)
func (ms msgServer) Swap(ctx context.Context, msg *types.MsgSwap) (*types.MsgEmpty, error)
func (ms msgServer) TradeAccountDeposit(ctx context.Context, msg *types.MsgTradeAccountDeposit) (*types.MsgEmpty, error)
func (ms msgServer) TradeAccountWithdrawal(ctx context.Context, msg *types.MsgTradeAccountWithdrawal) (*types.MsgEmpty, error)
func (ms msgServer) TssKeysignFail(ctx context.Context, msg *types.MsgTssKeysignFail) (*types.MsgEmpty, error)
func (ms msgServer) TssPool(ctx context.Context, msg *types.MsgTssPool) (*types.MsgEmpty, error)
func (ms msgServer) OutboundTx(ctx context.Context, msg *types.MsgOutboundTx) (*types.MsgEmpty, error)
func (ms msgServer) UnBond(ctx context.Context, msg *types.MsgUnBond) (*types.MsgEmpty, error)
func (ms msgServer) SetVersion(ctx context.Context, msg *types.MsgSetVersion) (*types.MsgEmpty, error)
func (ms msgServer) WithdrawLiquidity(ctx context.Context, msg *types.MsgWithdrawLiquidity) (*types.MsgEmpty, error)