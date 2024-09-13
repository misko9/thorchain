package thorchain

import (
	"context"
	"fmt"
	"runtime"
	
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrs "github.com/cosmos/cosmos-sdk/types/errors"
	errorsmod "cosmossdk.io/errors"

	"gitlab.com/thorchain/thornode/x/thorchain/types"
)

type msgServer struct {
	mgr Manager
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the module MsgServer interface.
func NewMsgServerImpl(mgr Manager) types.MsgServer {
	return &msgServer{mgr: mgr}
}

func (ms msgServer) AddLiquidity(goCtx context.Context, msg *types.MsgAddLiquidity) (*types.MsgEmpty, error) {
	handler := NewAddLiquidityHandler(ms.mgr)
	return internalHandler(goCtx, handler, msg)
}

func (ms msgServer) Ban(goCtx context.Context, msg *types.MsgBan) (*types.MsgEmpty, error) {
	handler := NewBanHandler(ms.mgr)
	return externalHandler(goCtx, handler, msg)
}

func (ms msgServer) Bond(goCtx context.Context, msg *types.MsgBond) (*types.MsgEmpty, error) {
	handler := NewBondHandler(ms.mgr)
	return internalHandler(goCtx, handler, msg)
}

func (ms msgServer) Consolidate(goCtx context.Context, msg *types.MsgConsolidate) (*types.MsgEmpty, error) {
	handler := NewConsolidateHandler(ms.mgr)
	return internalHandler(goCtx, handler, msg)
}

func (ms msgServer) Deposit(goCtx context.Context, msg *types.MsgDeposit) (*types.MsgEmpty, error) {
	handler := NewDepositHandler(ms.mgr)
	return externalHandler(goCtx, handler, msg)
}

func (ms msgServer) Donate(goCtx context.Context, msg *types.MsgDonate) (*types.MsgEmpty, error) {
	handler := NewDonateHandler(ms.mgr)
	return internalHandler(goCtx, handler, msg)
}

func (ms msgServer) ErrataTx(goCtx context.Context, msg *types.MsgErrataTx) (*types.MsgEmpty, error) {
	handler := NewErrataTxHandler(ms.mgr)
	return externalHandler(goCtx, handler, msg)
}

func (ms msgServer) Leave(goCtx context.Context, msg *types.MsgLeave) (*types.MsgEmpty, error) {
	handler := NewLeaveHandler(ms.mgr)
	return internalHandler(goCtx, handler, msg)
}

func (ms msgServer) LoanOpen(goCtx context.Context, msg *types.MsgLoanOpen) (*types.MsgEmpty, error) {
	handler := NewLoanOpenHandler(ms.mgr)
	return internalHandler(goCtx, handler, msg)
}

func (ms msgServer) LoanRepayment(goCtx context.Context, msg *types.MsgLoanRepayment) (*types.MsgEmpty, error) {
	handler := NewLoanRepaymentHandler(ms.mgr)
	return internalHandler(goCtx, handler, msg)
}

func (ms msgServer) ManageThorname(goCtx context.Context, msg *types.MsgManageTHORName) (*types.MsgEmpty, error) {
	handler := NewManageTHORNameHandler(ms.mgr)
	return internalHandler(goCtx, handler, msg)
}

func (ms msgServer) Migrate(goCtx context.Context, msg *types.MsgMigrate) (*types.MsgEmpty, error) {
	handler := NewMigrateHandler(ms.mgr)
	return internalHandler(goCtx, handler, msg)
}

func (ms msgServer) Mimir(goCtx context.Context, msg *types.MsgMimir) (*types.MsgEmpty, error) {
	handler := NewMimirHandler(ms.mgr)
	return externalHandler(goCtx, handler, msg)
}

func (ms msgServer) NetworkFee(goCtx context.Context, msg *types.MsgNetworkFee) (*types.MsgEmpty, error) {
	handler := NewNetworkFeeHandler(ms.mgr)
	return externalHandler(goCtx, handler, msg)
}

func (ms msgServer) NodePauseChain(goCtx context.Context, msg *types.MsgNodePauseChain) (*types.MsgEmpty, error) {
	handler := NewNodePauseChainHandler(ms.mgr)
	return externalHandler(goCtx, handler, msg)
}

func (ms msgServer) Noop(goCtx context.Context, msg *types.MsgNoOp) (*types.MsgEmpty, error) {
	handler := NewNoOpHandler(ms.mgr)
	return internalHandler(goCtx, handler, msg)
}

func (ms msgServer) ObservedTxIn(goCtx context.Context, msg *types.MsgObservedTxIn) (*types.MsgEmpty, error) {
	handler := NewObservedTxInHandler(ms.mgr)
	return externalHandler(goCtx, handler, msg)
}

func (ms msgServer) ObservedTxOut(goCtx context.Context, msg *types.MsgObservedTxOut) (*types.MsgEmpty, error) {
	handler := NewObservedTxOutHandler(ms.mgr)
	return externalHandler(goCtx, handler, msg)
}

func (ms msgServer) Ragnarok(goCtx context.Context, msg *types.MsgRagnarok) (*types.MsgEmpty, error) {
	handler := NewRagnarokHandler(ms.mgr)
	return internalHandler(goCtx, handler, msg)
}

func (ms msgServer) RefundTx(goCtx context.Context, msg *types.MsgRefundTx) (*types.MsgEmpty, error) {
	handler := NewRefundHandler(ms.mgr)
	return internalHandler(goCtx, handler, msg)
}

func (ms msgServer) ReserveContributor(goCtx context.Context, msg *types.MsgReserveContributor) (*types.MsgEmpty, error) {
	handler := NewReserveContributorHandler(ms.mgr)
	return internalHandler(goCtx, handler, msg)
}

func (ms msgServer) RunePoolDeposit(goCtx context.Context, msg *types.MsgRunePoolDeposit) (*types.MsgEmpty, error) {
	handler := NewRunePoolDepositHandler(ms.mgr)
	return internalHandler(goCtx, handler, msg)
}

func (ms msgServer) RunePoolWithdraw(goCtx context.Context, msg *types.MsgRunePoolWithdraw) (*types.MsgEmpty, error) {
	handler := NewRunePoolWithdrawHandler(ms.mgr)
	return internalHandler(goCtx, handler, msg)
}

func (ms msgServer) ThorSend(goCtx context.Context, msg *types.MsgSend) (*types.MsgEmpty, error) {
	handler := NewSendHandler(ms.mgr)
	return externalHandler(goCtx, handler, msg)
}

func (ms msgServer) SetIPAddress(goCtx context.Context, msg *types.MsgSetIPAddress) (*types.MsgEmpty, error) {
	handler := NewIPAddressHandler(ms.mgr)
	return externalHandler(goCtx, handler, msg)
}

func (ms msgServer) SetNodeKeys(goCtx context.Context, msg *types.MsgSetNodeKeys) (*types.MsgEmpty, error) {
	handler := NewSetNodeKeysHandler(ms.mgr)
	return externalHandler(goCtx, handler, msg)
}

func (ms msgServer) Solvency(goCtx context.Context, msg *types.MsgSolvency) (*types.MsgEmpty, error) {
	handler := NewSolvencyHandler(ms.mgr)
	return externalHandler(goCtx, handler, msg)
}

func (ms msgServer) Swap(goCtx context.Context, msg *types.MsgSwap) (*types.MsgEmpty, error) {
	handler := NewSwapHandler(ms.mgr)
	return internalHandler(goCtx, handler, msg)
}

func (ms msgServer) TradeAccountDeposit(goCtx context.Context, msg *types.MsgTradeAccountDeposit) (*types.MsgEmpty, error) {
	handler := NewTradeAccountDepositHandler(ms.mgr)
	return internalHandler(goCtx, handler, msg)
}

func (ms msgServer) TradeAccountWithdrawal(goCtx context.Context, msg *types.MsgTradeAccountWithdrawal) (*types.MsgEmpty, error) {
	handler := NewTradeAccountWithdrawalHandler(ms.mgr)
	return internalHandler(goCtx, handler, msg)
}

func (ms msgServer) TssKeysignFail(goCtx context.Context, msg *types.MsgTssKeysignFail) (*types.MsgEmpty, error) {
	handler := NewTssKeysignHandler(ms.mgr)
	return externalHandler(goCtx, handler, msg)
}

func (ms msgServer) TssPool(goCtx context.Context, msg *types.MsgTssPool) (*types.MsgEmpty, error) {
	handler := NewTssHandler(ms.mgr)
	return externalHandler(goCtx, handler, msg)
}

func (ms msgServer) OutboundTx(goCtx context.Context, msg *types.MsgOutboundTx) (*types.MsgEmpty, error) {
	handler := NewOutboundTxHandler(ms.mgr)
	return internalHandler(goCtx, handler, msg)
}

func (ms msgServer) UnBond(goCtx context.Context, msg *types.MsgUnBond) (*types.MsgEmpty, error) {
	handler := NewUnBondHandler(ms.mgr)
	return internalHandler(goCtx, handler, msg)
}

func (ms msgServer) SetVersion(goCtx context.Context, msg *types.MsgSetVersion) (*types.MsgEmpty, error) {
	handler := NewVersionHandler(ms.mgr)
	return externalHandler(goCtx, handler, msg)
}

func (ms msgServer) WithdrawLiquidity(goCtx context.Context, msg *types.MsgWithdrawLiquidity) (*types.MsgEmpty, error) {
	handler := NewWithdrawLiquidityHandler(ms.mgr)
	return internalHandler(goCtx, handler, msg)
}

func internalHandler(goCtx context.Context, handler MsgHandler, msg sdk.Msg) (*types.MsgEmpty, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// CacheContext() returns a context which caches all changes and only forwards
	// to the underlying context when commit() is called. Call commit() only when
	// the handler succeeds, otherwise return error and the changes will be discarded.
	// On commit, cached events also have to be explicitly emitted.
	cacheCtx, commit := ctx.CacheContext()
	result, err := handler.Run(ctx, msg)
	if result != nil && result.Size() > 0 {
		return nil, fmt.Errorf("internal handler, handler returned non-empty result, %s", msg)
	}
	if err == nil {
		// TODO (SAM127): is this still necessary?
		// Success, commit the cached changes and events
		commit()
		ctx.EventManager().EmitEvents(cacheCtx.EventManager().Events())
	}
		
	return &types.MsgEmpty{}, err
}

func externalHandler(goCtx context.Context, handler MsgHandler, msg sdk.Msg) (_ *types.MsgEmpty, err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	defer func() {
		if r := recover(); r != nil {
			// print stack
			stack := make([]byte, 1024)
			length := runtime.Stack(stack, true)
			ctx.Logger().Error("panic", "msg", msg)
			fmt.Println(string(stack[:length]))
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	
	ctx = ctx.WithEventManager(sdk.NewEventManager())
	result, err := handler.Run(ctx, msg)
	
	if result != nil && result.Size() > 0 {
		return nil, fmt.Errorf("external handler, handler returned non-empty result, %s", msg)
	}
	if err != nil {
		if _, code, _ := errorsmod.ABCIInfo(err, false); code == 1 {
			// This would be redacted, so wrap it.
			err = errorsmod.Wrap(errInternal, err.Error())
		}
		return nil, err
	}

	// TODO (SAM127): is this still necessary?
	// if len(ctx.EventManager().Events()) > 0 {
	// 	result.Events = ctx.EventManager().ABCIEvents()
	// }

	return &types.MsgEmpty{}, err
}