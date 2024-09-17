package thorchain

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/url"
	"strconv"
	"strings"
	"time"

	sdkmath "cosmossdk.io/math"
	"cosmossdk.io/log"

	"gitlab.com/thorchain/thornode/common"
	"gitlab.com/thorchain/thornode/common/cosmos"
	"gitlab.com/thorchain/thornode/constants"
	mem "gitlab.com/thorchain/thornode/x/thorchain/memo"
	"gitlab.com/thorchain/thornode/x/thorchain/types"
)

// -------------------------------------------------------------------------------------
// Config
// -------------------------------------------------------------------------------------

const (
	quoteWarning         = "Do not cache this response. Do not send funds after the expiry."
	quoteExpiration      = 15 * time.Minute
	ethBlockRewardAndFee = 3 * 1e18
)

var nullLogger = log.NewNopLogger()

// -------------------------------------------------------------------------------------
// Helpers
// -------------------------------------------------------------------------------------

func quoteErrorResponse(err error) ([]byte, error) {
	return json.Marshal(map[string]string{"error": err.Error()})
}

func quoteParseParams(data []byte) (params url.Values, err error) {
	// parse the query parameters
	u, err := url.ParseRequestURI(string(data))
	if err != nil {
		return nil, fmt.Errorf("bad params: %w", err)
	}

	// error if parameters were not provided
	if len(u.Query()) == 0 {
		return nil, fmt.Errorf("no parameters provided")
	}

	return u.Query(), nil
}

func quoteParseAddress(ctx cosmos.Context, mgr *Mgrs, addrString string, chain common.Chain) (common.Address, error) {
	if addrString == "" {
		return common.NoAddress, nil
	}

	// attempt to parse a raw address
	addr, err := common.NewAddress(addrString)
	if err == nil {
		return addr, nil
	}

	// attempt to lookup a thorname address
	name, err := mgr.Keeper().GetTHORName(ctx, addrString)
	if err != nil {
		return common.NoAddress, fmt.Errorf("unable to parse address: %w", err)
	}

	// find the address for the correct chain
	for _, alias := range name.Aliases {
		if alias.Chain.Equals(chain) {
			return alias.Address, nil
		}
	}

	return common.NoAddress, fmt.Errorf("no thorname alias for chain %s", chain)
}

func quoteHandleAffiliate(ctx cosmos.Context, mgr *Mgrs, affiliateParam string, affiliateBpsParam string, amount sdkmath.Uint) (affiliate common.Address, memo string, bps, newAmount, affiliateAmt sdkmath.Uint, err error) {
	// parse affiliate
	affAmt := cosmos.ZeroUint()
	memo = "" // do not resolve thorname for the memo
	if len(affiliateParam) > 0 {
		affiliate, err = quoteParseAddress(ctx, mgr, affiliateParam, common.THORChain)
		if err != nil {
			err = fmt.Errorf("bad affiliate address: %w", err)
			return
		}
		memo = affiliateParam
	}

	// parse affiliate fee
	bps = sdkmath.NewUint(0)
	if len(affiliateBpsParam) > 0 {
		bps, err = sdkmath.ParseUint(affiliateBpsParam)
		if err != nil {
			err = fmt.Errorf("bad affiliate fee: %w", err)
			return
		}
	}

	// verify affiliate fee
	if bps.GT(sdkmath.NewUint(10000)) {
		err = fmt.Errorf("affiliate fee must be less than 10000 bps")
		return
	}

	// compute the new swap amount if an affiliate fee will be taken first
	if affiliate != common.NoAddress && !bps.IsZero() {
		// calculate the affiliate amount
		affAmt = common.GetSafeShare(
			bps,
			cosmos.NewUint(10000),
			amount,
		)

		// affiliate fee modifies amount at observation before the swap
		amount = amount.Sub(affAmt)
	}

	return affiliate, memo, bps, amount, affAmt, nil
}

func hasSuffixMatch(suffix string, values []string) bool {
	for _, value := range values {
		if strings.HasSuffix(value, suffix) {
			return true
		}
	}
	return false
}

// quoteConvertAsset - converts amount to target asset using THORChain pools
func quoteConvertAsset(ctx cosmos.Context, mgr *Mgrs, fromAsset common.Asset, amount sdkmath.Uint, toAsset common.Asset) (sdkmath.Uint, error) {
	// no conversion necessary
	if fromAsset.Equals(toAsset) {
		return amount, nil
	}

	// convert to rune
	if !fromAsset.IsRune() {
		// get the fromPool for the from asset
		fromPool, err := mgr.Keeper().GetPool(ctx, fromAsset.GetLayer1Asset())
		if err != nil {
			return sdkmath.ZeroUint(), fmt.Errorf("failed to get pool: %w", err)
		}

		// ensure pool exists
		if fromPool.IsEmpty() {
			return sdkmath.ZeroUint(), fmt.Errorf("pool does not exist")
		}

		amount = fromPool.AssetValueInRune(amount)
	}

	// convert to target asset
	if !toAsset.IsRune() {

		toPool, err := mgr.Keeper().GetPool(ctx, toAsset.GetLayer1Asset())
		if err != nil {
			return sdkmath.ZeroUint(), fmt.Errorf("failed to get pool: %w", err)
		}

		// ensure pool exists
		if toPool.IsEmpty() {
			return sdkmath.ZeroUint(), fmt.Errorf("pool does not exist")
		}

		amount = toPool.RuneValueInAsset(amount)
	}

	return amount, nil
}

func quoteReverseFuzzyAsset(ctx cosmos.Context, mgr *Mgrs, asset common.Asset) (common.Asset, error) {
	// get all pools
	pools, err := mgr.Keeper().GetPools(ctx)
	if err != nil {
		return asset, fmt.Errorf("failed to get pools: %w", err)
	}

	// return the asset if no symbol to shorten
	aSplit := strings.Split(asset.Symbol.String(), "-")
	if len(aSplit) == 1 {
		return asset, nil
	}

	// find all other assets that match the chain and ticker
	// (without exactly matching the symbol)
	addressMatches := []string{}
	for _, p := range pools {
		if p.IsAvailable() && !p.IsEmpty() && !p.Asset.IsVaultAsset() &&
			!p.Asset.Symbol.Equals(asset.Symbol) &&
			p.Asset.Chain.Equals(asset.Chain) && p.Asset.Ticker.Equals(asset.Ticker) {
			pSplit := strings.Split(p.Asset.Symbol.String(), "-")
			if len(pSplit) != 2 {
				return asset, fmt.Errorf("ambiguous match: %s", p.Asset.Symbol)
			}
			addressMatches = append(addressMatches, pSplit[1])
		}
	}

	if len(addressMatches) == 0 { // if only one match, drop the address
		asset.Symbol = common.Symbol(asset.Ticker)
	} else { // find the shortest unique suffix of the asset symbol
		address := aSplit[1]

		for i := len(address) - 1; i > 0; i-- {
			if !hasSuffixMatch(address[i:], addressMatches) {
				asset.Symbol = common.Symbol(
					fmt.Sprintf("%s-%s", asset.Ticker, address[i:]),
				)
				break
			}
		}
	}

	return asset, nil
}

// NOTE: streamingQuantity > 0 is a precondition.
func quoteSimulateSwap(ctx cosmos.Context, mgr *Mgrs, amount sdkmath.Uint, msg *MsgSwap, streamingQuantity uint64) (
	res *types.QueryQuoteSwapResponse, emitAmount, outboundFeeAmount sdkmath.Uint, err error,
) {
	// should be unreachable
	if streamingQuantity == 0 {
		return nil, sdkmath.ZeroUint(), sdkmath.ZeroUint(), fmt.Errorf("streaming quantity must be greater than zero")
	}

	msg.Tx.Coins[0].Amount = msg.Tx.Coins[0].Amount.QuoUint64(streamingQuantity)

	// if the generated memo is too long for the source chain send error
	maxMemoLength := msg.Tx.Coins[0].Asset.Chain.MaxMemoLength()
	if !msg.Tx.Coins[0].Asset.Synth && len(msg.Tx.Memo) > maxMemoLength {
		return nil, sdkmath.ZeroUint(), sdkmath.ZeroUint(), fmt.Errorf("generated memo too long for source chain")
	}

	// use the first active node account as the signer
	nodeAccounts, err := mgr.Keeper().ListActiveValidators(ctx)
	if err != nil {
		return nil, sdkmath.ZeroUint(), sdkmath.ZeroUint(), fmt.Errorf("no active node accounts: %w", err)
	}
	msg.Signer = nodeAccounts[0].NodeAddress

	// simulate the swap
	events, err := simulateInternal(ctx, mgr, msg)
	if err != nil {
		return nil, sdkmath.ZeroUint(), sdkmath.ZeroUint(), err
	}

	// extract events
	var swaps []map[string]string
	var fee map[string]string
	for _, e := range events {
		switch e.Type {
		case "swap":
			swaps = append(swaps, eventMap(e))
		case "fee":
			fee = eventMap(e)
		}
	}
	finalSwap := swaps[len(swaps)-1]

	// parse outbound fee from event (except on trade assets with no outbound fee)
	outboundFeeAmount = sdkmath.ZeroUint()
	if !msg.TargetAsset.IsTradeAsset() {
		outboundFeeCoin, err := common.ParseCoin(fee["coins"])
		if err != nil {
			return nil, sdkmath.ZeroUint(), sdkmath.ZeroUint(), fmt.Errorf("unable to parse outbound fee coin: %w", err)
		}
		outboundFeeAmount = outboundFeeCoin.Amount
	}

	// parse outbound amount from event
	emitCoin, err := common.ParseCoin(finalSwap["emit_asset"])
	if err != nil {
		return nil, sdkmath.ZeroUint(), sdkmath.ZeroUint(), fmt.Errorf("unable to parse emit coin: %w", err)
	}
	emitAmount = emitCoin.Amount.MulUint64(streamingQuantity)

	// sum the liquidity fees and convert to target asset
	liquidityFee := sdkmath.ZeroUint()
	for _, s := range swaps {
		liquidityFee = liquidityFee.Add(sdkmath.NewUintFromString(s["liquidity_fee_in_rune"]))
	}
	var targetPool types.Pool
	if !msg.TargetAsset.IsRune() {
		targetPool, err = mgr.Keeper().GetPool(ctx, msg.TargetAsset.GetLayer1Asset())
		if err != nil {
			return nil, sdkmath.ZeroUint(), sdkmath.ZeroUint(), fmt.Errorf("unable to get pool: %w", err)
		}
		liquidityFee = targetPool.RuneValueInAsset(liquidityFee)
	}
	liquidityFee = liquidityFee.MulUint64(streamingQuantity)

	// approximate the affiliate fee in the target asset
	affiliateFee := sdkmath.ZeroUint()
	if msg.AffiliateAddress != common.NoAddress && !msg.AffiliateBasisPoints.IsZero() {
		inAsset := msg.Tx.Coins[0].Asset.GetLayer1Asset()
		if !inAsset.IsRune() {
			pool, err := mgr.Keeper().GetPool(ctx, msg.Tx.Coins[0].Asset.GetLayer1Asset())
			if err != nil {
				return nil, sdkmath.ZeroUint(), sdkmath.ZeroUint(), fmt.Errorf("unable to get pool: %w", err)
			}
			amount = pool.AssetValueInRune(amount)
		}
		affiliateFee = common.GetUncappedShare(msg.AffiliateBasisPoints, cosmos.NewUint(10_000), amount)
		if !msg.TargetAsset.IsRune() {
			affiliateFee = targetPool.RuneValueInAsset(affiliateFee)
		}
	}

	// compute slip based on emit amount instead of slip in event to handle double swap
	slippageBps := liquidityFee.MulUint64(10000).Quo(emitAmount.Add(liquidityFee))

	// build fees
	totalFees := affiliateFee.Add(liquidityFee).Add(outboundFeeAmount)
	fees := types.QuoteFees{
		Asset:       msg.TargetAsset.String(),
		Affiliate:   affiliateFee.String(),
		Liquidity:   liquidityFee.String(),
		Outbound:    outboundFeeAmount.String(),
		Total:       totalFees.String(),
		SlippageBps: slippageBps.BigInt().Int64(),
		TotalBps:    totalFees.MulUint64(10000).Quo(emitAmount.Add(totalFees)).BigInt().Int64(),
	}

	// build response from simulation result events
	return &types.QueryQuoteSwapResponse{
		ExpectedAmountOut: emitAmount.String(),
		Fees:              &fees,
	}, emitAmount, outboundFeeAmount, nil
}

func convertThorchainAmountToWei(amt *big.Int) *big.Int {
	return big.NewInt(0).Mul(amt, big.NewInt(common.One*100))
}

func quoteInboundInfo(ctx cosmos.Context, mgr *Mgrs, amount sdkmath.Uint, chain common.Chain, asset common.Asset) (address, router common.Address, confirmations int64, err error) {
	// If inbound chain is THORChain there is no inbound address
	if chain.IsTHORChain() {
		address = common.NoAddress
		router = common.NoAddress
	} else {
		// get the most secure vault for inbound
		active, err := mgr.Keeper().GetAsgardVaultsByStatus(ctx, ActiveVault)
		if err != nil {
			return common.NoAddress, common.NoAddress, 0, err
		}
		constAccessor := mgr.GetConstants()
		signingTransactionPeriod := constAccessor.GetInt64Value(constants.SigningTransactionPeriod)
		vault := mgr.Keeper().GetMostSecure(ctx, active, signingTransactionPeriod)
		address, err = vault.PubKey.GetAddress(chain)
		if err != nil {
			return common.NoAddress, common.NoAddress, 0, err
		}

		router = common.NoAddress
		if chain.IsEVM() {
			router = vault.GetContract(chain).Router
		}
	}

	// estimate the inbound confirmation count blocks: ceil(amount/coinbase * conf adjustment)
	confMul, err := mgr.Keeper().GetMimirWithRef(ctx, constants.MimirTemplateConfMultiplierBasisPoints, chain.String())
	if confMul < 0 || err != nil {
		confMul = int64(constants.MaxBasisPts)
	}
	if chain.DefaultCoinbase() > 0 {
		confValue := common.GetUncappedShare(cosmos.NewUint(uint64(confMul)), cosmos.NewUint(constants.MaxBasisPts), cosmos.NewUint(uint64(chain.DefaultCoinbase())*common.One))
		confirmations = amount.Quo(confValue).BigInt().Int64()
		if !amount.Mod(confValue).IsZero() {
			confirmations++
		}
	} else if chain.Equals(common.ETHChain) {
		// copying logic from getBlockRequiredConfirmation of ethereum.go
		// convert amount to ETH
		gasAssetAmount, err := quoteConvertAsset(ctx, mgr, asset, amount, chain.GetGasAsset())
		if err != nil {
			return common.NoAddress, common.NoAddress, 0, fmt.Errorf("unable to convert asset: %w", err)
		}
		gasAssetAmountWei := convertThorchainAmountToWei(gasAssetAmount.BigInt())
		confValue := common.GetUncappedShare(cosmos.NewUint(uint64(confMul)), cosmos.NewUint(constants.MaxBasisPts), cosmos.NewUintFromBigInt(big.NewInt(ethBlockRewardAndFee)))
		confirmations = int64(cosmos.NewUintFromBigInt(gasAssetAmountWei).MulUint64(2).Quo(confValue).Uint64())
	}

	// max confirmation adjustment for btc and eth
	if chain.Equals(common.BTCChain) || chain.Equals(common.ETHChain) {
		maxConfirmations, err := mgr.Keeper().GetMimirWithRef(ctx, constants.MimirTemplateMaxConfirmations, chain.String())
		if maxConfirmations < 0 || err != nil {
			maxConfirmations = 0
		}
		if maxConfirmations > 0 && confirmations > maxConfirmations {
			confirmations = maxConfirmations
		}
	}

	// min confirmation adjustment
	confFloor := map[common.Chain]int64{
		common.ETHChain:  2,
		common.DOGEChain: 2,
	}
	if floor := confFloor[chain]; confirmations < floor {
		confirmations = floor
	}

	return address, router, confirmations, nil
}

func quoteOutboundInfo(ctx cosmos.Context, mgr *Mgrs, coin common.Coin) (int64, error) {
	toi := TxOutItem{
		Memo: "OUT:-",
		Coin: coin,
	}
	outboundHeight, _, err := mgr.txOutStore.CalcTxOutHeight(ctx, mgr.GetVersion(), toi)
	if err != nil {
		return 0, err
	}
	return outboundHeight - ctx.BlockHeight(), nil
}

// -------------------------------------------------------------------------------------
// Swap
// -------------------------------------------------------------------------------------

// calculateMinSwapAmount returns the recommended minimum swap amount The recommended
// min swap amount is: - MAX(
//
//	  outbound_fee(src_chain) * 4,
//	  outbound_fee(dest_chain) * 4,
//	  (native_tx_fee_rune * 2) * 10,000 / affiliateBps
//	)
//
// The reason the base value is the MAX of the outbound fees of each chain is because if
// the swap is refunded the input amount will need to cover the outbound fee of the
// source chain. A 4x buffer is applied because outbound fees can spike quickly, meaning
// the original input amount could be less than the new outbound fee. If this happens
// and the swap is refunded, the refund will fail, and the user will lose the entire
// input amount. The min amount could also be determined by the affiliate bps of the
// swap. The affiliate bps of the input amount needs to be enough to cover the native tx fee for the
// affiliate swap to RUNE. In this case, we give a 2x buffer on the native_tx_fee so the
// affiliate receives some amount after the fee is deducted.
func calculateMinSwapAmount(ctx cosmos.Context, mgr *Mgrs, fromAsset, toAsset common.Asset, affiliateBps cosmos.Uint) (cosmos.Uint, error) {
	srcOutboundFee, err := mgr.GasMgr().GetAssetOutboundFee(ctx, fromAsset, false)
	if err != nil {
		return cosmos.ZeroUint(), fmt.Errorf("fail to get outbound fee for source chain gas asset %s: %w", fromAsset, err)
	}
	destOutboundFee, err := mgr.GasMgr().GetAssetOutboundFee(ctx, toAsset, false)
	if err != nil {
		return cosmos.ZeroUint(), fmt.Errorf("fail to get outbound fee for destination chain gas asset %s: %w", toAsset, err)
	}

	if fromAsset.GetChain().IsTHORChain() && toAsset.GetChain().IsTHORChain() {
		// If this is a purely THORChain swap, no need to give a 4x buffer since outbound fees do not change
		// 2x buffer should suffice
		return srcOutboundFee.Mul(cosmos.NewUint(2)), nil
	}

	destInSrcAsset, err := quoteConvertAsset(ctx, mgr, toAsset, destOutboundFee, fromAsset)
	if err != nil {
		return cosmos.ZeroUint(), fmt.Errorf("fail to convert dest fee to src asset %w", err)
	}

	minSwapAmount := srcOutboundFee
	if destInSrcAsset.GT(srcOutboundFee) {
		minSwapAmount = destInSrcAsset
	}

	minSwapAmount = minSwapAmount.Mul(cosmos.NewUint(4))

	if affiliateBps.GT(cosmos.ZeroUint()) {
		nativeTxFeeRune, err := mgr.GasMgr().GetAssetOutboundFee(ctx, common.RuneNative, true)
		if err != nil {
			return cosmos.ZeroUint(), fmt.Errorf("fail to get native tx fee for rune: %w", err)
		}
		affSwapAmountRune := nativeTxFeeRune.Mul(cosmos.NewUint(2))
		mainSwapAmountRune := affSwapAmountRune.Mul(cosmos.NewUint(10_000)).Quo(affiliateBps)

		mainSwapAmount, err := quoteConvertAsset(ctx, mgr, common.RuneAsset(), mainSwapAmountRune, fromAsset)
		if err != nil {
			return cosmos.ZeroUint(), fmt.Errorf("fail to convert main swap amount to src asset %w", err)
		}

		if mainSwapAmount.GT(minSwapAmount) {
			minSwapAmount = mainSwapAmount
		}
	}

	return minSwapAmount, nil
}

func (qs queryServer) queryQuoteSwap(ctx cosmos.Context, req *types.QueryQuoteSwapRequest) (*types.QueryQuoteSwapResponse, error) {
	// validate required parameters
	if len(req.FromAsset) == 0 {
		return nil, fmt.Errorf("missing from_asset parameter")
	}

	if len(req.ToAsset) == 0 {
		return nil, fmt.Errorf("missing to_asset parameter")
	}

	if len(req.Amount) == 0 {
		return nil, fmt.Errorf("missing Amount parameter")
	}

	// parse assets
	fromAsset, err := common.NewAssetWithShortCodes(qs.mgr.GetVersion(), req.FromAsset)
	if err != nil {
		return nil, fmt.Errorf("bad from asset: %w", err)
	}
	fromAsset = fuzzyAssetMatch(ctx, qs.mgr.Keeper(), fromAsset)
	toAsset, err := common.NewAssetWithShortCodes(qs.mgr.GetVersion(), req.ToAsset)
	if err != nil {
		return nil, fmt.Errorf("bad to asset: %w", err)
	}
	toAsset = fuzzyAssetMatch(ctx, qs.mgr.Keeper(), toAsset)

	// parse amount
	amount, err := cosmos.ParseUint(req.Amount)
	if err != nil {
		return nil, fmt.Errorf("bad amount: %w", err)
	}

	// parse streaming interval
	streamingInterval := uint64(0) // default value
	if len(req.StreamingInterval) > 0 {
		streamingInterval, err = strconv.ParseUint(req.StreamingInterval, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("bad streaming interval amount: %w", err)
		}
	}
	streamingQuantity := uint64(0) // default value
	if len(req.StreamingQuantity) > 0 {
		streamingQuantity, err = strconv.ParseUint(req.StreamingQuantity, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("bad streaming quantity amount: %w", err)
		}
	}
	swp := StreamingSwap{
		Interval: streamingInterval,
		Deposit:  amount,
	}
	maxSwapQuantity, err := getMaxSwapQuantity(ctx, qs.mgr, fromAsset, toAsset, swp)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate max streaming swap quantity: %w", err)
	}

	// cap the streaming quantity to the max swap quantity
	if streamingQuantity > maxSwapQuantity {
		streamingQuantity = maxSwapQuantity
	}

	// if from asset is a synth, transfer asset to asgard module
	if fromAsset.IsSyntheticAsset() {
		// mint required coins to asgard so swap can be simulated
		err = qs.mgr.Keeper().MintToModule(ctx, ModuleName, common.NewCoin(fromAsset, amount))
		if err != nil {
			return nil, fmt.Errorf("failed to mint coins to module: %w", err)
		}

		err = qs.mgr.Keeper().SendFromModuleToModule(ctx, ModuleName, AsgardName, common.NewCoins(common.NewCoin(fromAsset, amount)))
		if err != nil {
			return nil, fmt.Errorf("failed to send coins to asgard: %w", err)
		}
	}

	// parse affiliate
	affiliate, affiliateMemo, affiliateBps, swapAmount, affAmt, err := quoteHandleAffiliate(ctx, qs.mgr, req.Affiliate, req.AffiliateBps, amount)
	if err != nil {
		return nil, err
	}

	// simulate/validate the affiliate swap
	if affAmt.GT(sdkmath.ZeroUint()) {
		if fromAsset.IsRune() {
			fee := qs.mgr.Keeper().GetNativeTxFee(ctx)
			if affAmt.LTE(fee) {
				return nil, fmt.Errorf("affiliate amount must be greater than native fee %s", fee)
			}
		} else {
			// validate affiliate address
			affiliateSwapMsg := &types.MsgSwap{
				Tx: common.Tx{
					ID:          common.BlankTxID,
					Chain:       fromAsset.Chain,
					FromAddress: common.NoopAddress,
					ToAddress:   common.NoopAddress,
					Coins: []common.Coin{
						{
							Asset:  fromAsset,
							Amount: affAmt,
						},
					},
					Gas: []common.Coin{{
						Asset:  common.RuneAsset(),
						Amount: sdkmath.NewUint(1),
					}},
					Memo: "",
				},
				TargetAsset:          common.RuneAsset(),
				TradeTarget:          cosmos.ZeroUint(),
				Destination:          affiliate,
				AffiliateAddress:     common.NoAddress,
				AffiliateBasisPoints: cosmos.ZeroUint(),
			}

			nodeAccounts, err := qs.mgr.Keeper().ListActiveValidators(ctx)
			if err != nil {
				return nil, fmt.Errorf("no active node accounts: %w", err)
			}
			affiliateSwapMsg.Signer = nodeAccounts[0].NodeAddress

			// simulate the swap
			_, err = simulateInternal(ctx, qs.mgr, affiliateSwapMsg)
			if err != nil {
				return nil, fmt.Errorf("affiliate swap failed: %w", err)
			}
		}
	}

	// parse destination address or generate a random one
	sendMemo := true
	var destination common.Address
	if len(req.Destination) > 0 {
		destination, err = quoteParseAddress(ctx, qs.mgr, req.Destination, toAsset.Chain)
		if err != nil {
			return nil, fmt.Errorf("bad destination address: %w", err)
		}

	} else {
		chain := common.THORChain
		if !toAsset.IsSyntheticAsset() {
			chain = toAsset.Chain
		}
		destination, err = types.GetRandomPubKey().GetAddress(chain)
		if err != nil {
			return nil, fmt.Errorf("failed to generate address: %w", err)
		}
		sendMemo = false // do not send memo if destination was random
	}

	// parse tolerance basis points
	limit := sdkmath.ZeroUint()
	if len(req.ToleranceBps) > 0 {
		// validate tolerance basis points
		toleranceBasisPoints, err := sdkmath.ParseUint(req.ToleranceBps)
		if err != nil {
			return nil, fmt.Errorf("bad tolerance basis points: %w", err)
		}
		if toleranceBasisPoints.GT(sdkmath.NewUint(10000)) {
			return nil, fmt.Errorf("tolerance basis points must be less than 10000")
		}

		// convert to a limit of target asset amount assuming zero fees and slip
		feelessEmit, err := quoteConvertAsset(ctx, qs.mgr, fromAsset, swapAmount, toAsset)
		if err != nil {
			return nil, err
		}

		limit = feelessEmit.MulUint64(10000 - toleranceBasisPoints.Uint64()).QuoUint64(10000)
	}

	// custom refund addr
	refundAddress := common.NoAddress
	if len(req.RefundAddress) > 0 {
		refundAddress, err = quoteParseAddress(ctx, qs.mgr, req.RefundAddress, fromAsset.Chain)
		if err != nil {
			return nil, fmt.Errorf("bad refund address: %w", err)
		}
	}

	// create the memo
	memo := &SwapMemo{
		MemoBase: mem.MemoBase{
			TxType: TxSwap,
			Asset:  toAsset,
		},
		Destination:          destination,
		SlipLimit:            limit,
		AffiliateAddress:     common.Address(affiliateMemo),
		AffiliateBasisPoints: affiliateBps,
		StreamInterval:       streamingInterval,
		StreamQuantity:       streamingQuantity,
		RefundAddress:        refundAddress,
	}

	// if from asset chain has memo length restrictions use a prefix
	memoString := memo.String()
	if !fromAsset.Synth && len(memoString) > fromAsset.Chain.MaxMemoLength() {
		if len(memo.ShortString()) < len(memoString) { // use short codes if available
			memoString = memo.ShortString()
		} else { // otherwise attempt to shorten
			fuzzyAsset, err := quoteReverseFuzzyAsset(ctx, qs.mgr, toAsset)
			if err == nil {
				memo.Asset = fuzzyAsset
				memoString = memo.String()
			}
		}

		// this is the shortest we can make it
		if len(memoString) > fromAsset.Chain.MaxMemoLength() {
			return nil, fmt.Errorf("generated memo too long for source chain")
		}
	}

	// trade assets must have from address on the source tx
	fromChain := fromAsset.Chain
	if fromAsset.IsSyntheticAsset() || fromAsset.IsDerivedAsset() || fromAsset.IsTradeAsset() {
		fromChain = common.THORChain
	}
	fromPubkey := types.GetRandomPubKey()
	fromAddress, err := fromPubkey.GetAddress(fromChain)
	if err != nil {
		return nil, fmt.Errorf("bad from address: %w", err)
	}

	// if from asset is a trade asset, create fake balance
	if fromAsset.IsTradeAsset() {
		thorAddr, err := fromPubkey.GetThorAddress()
		if err != nil {
			return nil, fmt.Errorf("failed to get thor address: %w", err)
		}
		_, err = qs.mgr.TradeAccountManager().Deposit(ctx, fromAsset, amount, thorAddr, common.NoAddress, common.BlankTxID)
		if err != nil {
			return nil, fmt.Errorf("failed to deposit trade asset: %w", err)
		}
	}

	// create the swap message
	msg := &types.MsgSwap{
		Tx: common.Tx{
			ID:          common.BlankTxID,
			Chain:       fromAsset.Chain,
			FromAddress: fromAddress,
			ToAddress:   common.NoopAddress,
			Coins: []common.Coin{
				{
					Asset:  fromAsset,
					Amount: swapAmount,
				},
			},
			Gas: []common.Coin{{
				Asset:  common.RuneAsset(),
				Amount: sdkmath.NewUint(1),
			}},
			Memo: memoString,
		},
		TargetAsset:          toAsset,
		TradeTarget:          limit,
		Destination:          destination,
		AffiliateAddress:     affiliate,
		AffiliateBasisPoints: affiliateBps,
	}

	// simulate the swap
	res, emitAmount, outboundFeeAmount, err := quoteSimulateSwap(ctx, qs.mgr, amount, msg, 1)
	if err != nil {
		return nil, fmt.Errorf("failed to simulate swap: %w", err)
	}

	// if we're using a streaming swap, calculate emit amount by a sub-swap amount instead
	// of the full amount, then multiply the result by the swap count
	if streamingInterval > 0 && streamingQuantity == 0 {
		streamingQuantity = maxSwapQuantity
	}
	if streamingInterval > 0 && streamingQuantity > 0 {
		msg.TradeTarget = msg.TradeTarget.QuoUint64(streamingQuantity)
		// simulate the swap
		var streamRes *types.QueryQuoteSwapResponse
		streamRes, emitAmount, _, err = quoteSimulateSwap(ctx, qs.mgr, amount, msg, streamingQuantity)
		if err != nil {
			return nil, fmt.Errorf("failed to simulate swap: %w", err)
		}
		res.Fees = streamRes.Fees
	}

	// TODO: After UIs have transitioned everything below the message definition above
	// should reduce to the following:
	//
	// if streamingInterval > 0 && streamingQuantity == 0 {
	//   streamingQuantity = maxSwapQuantity
	// }
	// if streamingInterval > 0 && streamingQuantity > 0 {
	//   msg.TradeTarget = msg.TradeTarget.QuoUint64(streamingQuantity)
	// }
	// res, emitAmount, outboundFeeAmount, err := quoteSimulateSwap(ctx, mgr, amount, msg, streamingQuantity)
	// if err != nil {
	//   return quoteErrorResponse(fmt.Errorf("failed to simulate swap: %w", err))
	// }

	// check invariant
	if emitAmount.LT(outboundFeeAmount) {
		return nil, fmt.Errorf("invariant broken: emit %s less than outbound fee %s", emitAmount, outboundFeeAmount)
	}

	// the amount out will deduct the outbound fee
	res.ExpectedAmountOut = emitAmount.Sub(outboundFeeAmount).String()

	maxQ := int64(maxSwapQuantity)
	res.MaxStreamingQuantity = maxQ
	var streamSwapBlocks int64
	if streamingQuantity > 0 {
		streamSwapBlocks = int64(streamingInterval) * int64(streamingQuantity-1)
	}
	res.StreamingSwapBlocks = streamSwapBlocks
	res.StreamingSwapSeconds = streamSwapBlocks * common.THORChain.ApproximateBlockMilliseconds() / 1000

	// estimate the inbound info
	inboundAddress, routerAddress, inboundConfirmations, err := quoteInboundInfo(ctx, qs.mgr, amount, fromAsset.GetChain(), fromAsset)
	if err != nil {
		return nil, err
	}
	res.InboundAddress = inboundAddress.String()
	if inboundConfirmations > 0 {
		res.InboundConfirmationBlocks = inboundConfirmations
		res.InboundConfirmationSeconds = inboundConfirmations * msg.Tx.Chain.ApproximateBlockMilliseconds() / 1000
	}

	res.OutboundDelayBlocks = 0
	res.OutboundDelaySeconds = 0
	if !toAsset.Chain.IsTHORChain() {
		// estimate the outbound info
		outboundDelay, err := quoteOutboundInfo(ctx, qs.mgr, common.Coin{Asset: toAsset, Amount: emitAmount})
		if err != nil {
			return nil, err
		}
		res.OutboundDelayBlocks = outboundDelay
		res.OutboundDelaySeconds = outboundDelay * common.THORChain.ApproximateBlockMilliseconds() / 1000
	}

	totalSeconds := res.OutboundDelaySeconds
	// TODO: can outbound delay seconds be negative?
	if res.StreamingSwapSeconds != 0 && res.OutboundDelaySeconds < res.StreamingSwapSeconds {
		totalSeconds = res.StreamingSwapSeconds
	}
	if inboundConfirmations > 0 {
		totalSeconds += res.InboundConfirmationSeconds
	}
	res.TotalSwapSeconds = totalSeconds

	// send memo if the destination was provided
	if sendMemo {
		res.Memo = memoString
	}

	// set info fields
	if fromAsset.Chain.IsEVM() {
		res.Router = routerAddress.String()
	}
	if !fromAsset.Chain.DustThreshold().IsZero() {
		res.DustThreshold = fromAsset.Chain.DustThreshold().String()
	}

	res.Notes = fromAsset.GetChain().InboundNotes()
	res.Warning = quoteWarning
	res.Expiry = time.Now().Add(quoteExpiration).Unix()
	minSwapAmount, err := calculateMinSwapAmount(ctx, qs.mgr, fromAsset, toAsset, affiliateBps)
	if err != nil {
		return nil, fmt.Errorf("Failed to calculate min amount in: %s", err.Error())
	}
	res.RecommendedMinAmountIn = minSwapAmount.String()

	// set inbound recommended gas for non-native swaps
	if !fromAsset.Chain.IsTHORChain() {
		inboundGas := qs.mgr.GasMgr().GetGasRate(ctx, fromAsset.Chain)
		res.RecommendedGasRate = inboundGas.String()
		res.GasRateUnits = fromAsset.Chain.GetGasUnits()
	}

	return res, nil
}

// -------------------------------------------------------------------------------------
// Saver Deposit
// -------------------------------------------------------------------------------------

func (qs queryServer) queryQuoteSaverDeposit(ctx cosmos.Context,  req *types.QueryQuoteSaverDepositRequest) (*types.QueryQuoteSaverDepositResponse, error) {
	// validate required parameters
	if len(req.Asset) == 0 {
		return nil, fmt.Errorf("missing asset parameter")
	}
	if len(req.Amount) == 0 {
		return nil, fmt.Errorf("missing amount parameter")
	}

	// parse asset
	asset, err := common.NewAssetWithShortCodes(qs.mgr.GetVersion(), req.Asset)
	if err != nil {
		return nil, fmt.Errorf("bad asset: %w", err)
	}
	asset = fuzzyAssetMatch(ctx, qs.mgr.Keeper(), asset)

	// parse amount
	amount, err := cosmos.ParseUint(req.Amount)
	if err != nil {
		return nil, fmt.Errorf("bad amount: %w", err)
	}

	// parse affiliate
	affiliate, affiliateMemo, affiliateBps, depositAmount, _, err := quoteHandleAffiliate(ctx, qs.mgr, req.Affiliate, req.AffiliateBps, amount)
	if err != nil {
		return nil, err
	}

	// generate deposit memo
	depositMemoComponents := []string{
		"+",
		asset.GetSyntheticAsset().String(),
		"",
		affiliateMemo,
		affiliateBps.String(),
	}
	depositMemo := strings.Join(depositMemoComponents[:2], ":")
	if affiliate != common.NoAddress && !affiliateBps.IsZero() {
		depositMemo = strings.Join(depositMemoComponents, ":")
	}

	swapMsg := types.QueryQuoteSwapRequest{
		FromAsset: asset.String(),
		ToAsset: asset.GetSyntheticAsset().String(),
		Amount: depositAmount.String(),
		Destination: string(GetRandomTHORAddress()),
	}

	ssInterval := qs.mgr.Keeper().GetConfigInt64(ctx, constants.SaversStreamingSwapsInterval)
	if ssInterval > 0 {
		swapMsg.StreamingInterval = fmt.Sprintf("%d", ssInterval)
		swapMsg.StreamingQuantity = fmt.Sprintf("%d", 0)
	}

	swapRes, err := qs.queryQuoteSwap(ctx, &swapMsg)
	if err != nil {
		return nil, fmt.Errorf("unable to queryQuoteSwap: %w", err)
	}

	expectedAmountOut, _ := sdkmath.ParseUint(swapRes.ExpectedAmountOut)
	outboundFee, _ := sdkmath.ParseUint(swapRes.Fees.Outbound)
	depositAmount = expectedAmountOut.Add(outboundFee)

	// use the swap result info to generate the deposit quote
	res := &types.QueryQuoteSaverDepositResponse{
		// TODO: deprecate ExpectedAmountOut in future version
		ExpectedAmountOut:          depositAmount.String(),
		ExpectedAmountDeposit:      depositAmount.String(),
		Fees:                       swapRes.Fees,
		InboundConfirmationBlocks:  swapRes.InboundConfirmationBlocks,
		InboundConfirmationSeconds: swapRes.InboundConfirmationSeconds,
		Memo:                       depositMemo,
	}

	// estimate the inbound info
	inboundAddress, _, inboundConfirmations, err := quoteInboundInfo(ctx, qs.mgr, amount, asset.GetLayer1Asset().Chain, asset)
	if err != nil {
		return nil, err
	}
	res.InboundAddress = inboundAddress.String()
	res.InboundConfirmationBlocks = inboundConfirmations

	// set info fields
	chain := asset.GetLayer1Asset().Chain
	if !chain.DustThreshold().IsZero() {
		res.DustThreshold = chain.DustThreshold().String()
		res.RecommendedMinAmountIn = res.DustThreshold
	}
	res.Notes = chain.InboundNotes()
	res.Warning = quoteWarning
	res.Expiry = time.Now().Add(quoteExpiration).Unix()

	// set inbound recommended gas
	inboundGas := qs.mgr.GasMgr().GetGasRate(ctx, chain)
	res.RecommendedGasRate = inboundGas.String()
	res.GasRateUnits = chain.GetGasUnits()

	return res, nil
}

// -------------------------------------------------------------------------------------
// Saver Withdraw
// -------------------------------------------------------------------------------------

func (qs queryServer) queryQuoteSaverWithdraw(ctx cosmos.Context,  req *types.QueryQuoteSaverWithdrawRequest) (*types.QueryQuoteSaverWithdrawResponse, error) {
	// validate required parameters
	if len(req.Asset) == 0 {
		return nil, fmt.Errorf("missing asset parameter")
	}
	if len(req.Address) == 0 {
		return nil, fmt.Errorf("missing address parameter")
	}
	if len(req.WithdrawBps) == 0 {
		return nil, fmt.Errorf("missing withdraw_bps paramter")
	}

	// parse asset
	asset, err := common.NewAssetWithShortCodes(qs.mgr.GetVersion(), req.Asset)
	if err != nil {
		return nil, fmt.Errorf("bad asset: %w", err)
	}
	asset = fuzzyAssetMatch(ctx, qs.mgr.Keeper(), asset)
	asset = asset.GetSyntheticAsset() // always use the vault asset

	// parse address
	address, err := common.NewAddress(req.Address)
	if err != nil {
		return nil, fmt.Errorf("bad address: %w", err)
	}

	// parse basis points
	basisPoints, err := cosmos.ParseUint(req.WithdrawBps)
	if err != nil {
		return nil, fmt.Errorf("bad basis points: %w", err)
	}

	// validate basis points
	if basisPoints.GT(sdkmath.NewUint(10_000)) {
		return nil, fmt.Errorf("basis points must be less than 10000")
	}

	// get liquidity provider
	lp, err := qs.mgr.Keeper().GetLiquidityProvider(ctx, asset, address)
	if err != nil {
		return nil, fmt.Errorf("failed to get liquidity provider: %w", err)
	}

	// get the pool
	pool, err := qs.mgr.Keeper().GetPool(ctx, asset)
	if err != nil {
		return nil, fmt.Errorf("failed to get pool: %w", err)
	}

	// get the liquidity provider share of the pool
	lpShare := lp.GetSaversAssetRedeemValue(pool)

	// calculate the withdraw amount
	amount := common.GetSafeShare(basisPoints, sdkmath.NewUint(10_000), lpShare)

	swapMsg := types.QueryQuoteSwapRequest{
		FromAsset: asset.String(),
		ToAsset: asset.GetLayer1Asset().String(),
		Amount: amount.String(),
		Destination: string(GetRandomTHORAddress()), // required param, not actually used, spoof it
	}

	ssInterval := qs.mgr.Keeper().GetConfigInt64(ctx, constants.SaversStreamingSwapsInterval)
	if ssInterval > 0 {
		swapMsg.StreamingInterval = fmt.Sprintf("%d", ssInterval)
		swapMsg.StreamingQuantity = fmt.Sprintf("%d", 0)
	}

	swapRes, err := qs.queryQuoteSwap(ctx, &swapMsg)
	if err != nil {
		return nil, fmt.Errorf("unable to queryQuoteSwap: %w", err)
	}

	// use the swap result info to generate the withdraw quote
	res := &types.QueryQuoteSaverWithdrawResponse{
		ExpectedAmountOut: swapRes.ExpectedAmountOut,
		Fees:              swapRes.Fees,
		Memo:              fmt.Sprintf("-:%s:%s", asset.String(), basisPoints.String()),
		DustAmount:        asset.GetLayer1Asset().Chain.DustThreshold().Add(basisPoints).String(),
	}

	// estimate the inbound info
	inboundAddress, _, _, err := quoteInboundInfo(ctx, qs.mgr, amount, asset.GetLayer1Asset().Chain, asset)
	if err != nil {
		return nil, err
	}
	res.InboundAddress = inboundAddress.String()

	// estimate the outbound info
	expectedAmountOut, _ := sdkmath.ParseUint(swapRes.ExpectedAmountOut)
	outboundCoin := common.Coin{Asset: asset.GetLayer1Asset(), Amount: expectedAmountOut}
	outboundDelay, err := quoteOutboundInfo(ctx, qs.mgr, outboundCoin)
	if err != nil {
		return nil, err
	}
	res.OutboundDelayBlocks = outboundDelay
	res.OutboundDelaySeconds = outboundDelay * common.THORChain.ApproximateBlockMilliseconds() / 1000

	// set info fields
	chain := asset.GetLayer1Asset().Chain
	if !chain.DustThreshold().IsZero() {
		res.DustThreshold = chain.DustThreshold().String()
	}
	res.Notes = chain.InboundNotes()
	res.Warning = quoteWarning
	res.Expiry = time.Now().Add(quoteExpiration).Unix()

	// set inbound recommended gas
	inboundGas := qs.mgr.GasMgr().GetGasRate(ctx, chain)
	res.RecommendedGasRate = inboundGas.String()
	res.GasRateUnits = chain.GetGasUnits()

	return res, nil
}

// -------------------------------------------------------------------------------------
// Loan Open
// -------------------------------------------------------------------------------------

func (qs queryServer) queryQuoteLoanOpen(ctx cosmos.Context, req *types.QueryQuoteLoanOpenRequest) (*types.QueryQuoteLoanOpenResponse, error) {
	// validate required parameters
	if len(req.FromAsset) == 0 {
		return nil, fmt.Errorf("missing from_asset parameter")
	}
	if len(req.ToAsset) == 0 {
		return nil, fmt.Errorf("missing to_asset parameter")
	}
	if len(req.Amount) == 0 {
		return nil, fmt.Errorf("missing amount parameter")
	}

	// parse asset
	asset, err := common.NewAssetWithShortCodes(qs.mgr.GetVersion(), req.FromAsset)
	if err != nil {
		return nil, fmt.Errorf("bad asset: %w", err)
	}
	asset = fuzzyAssetMatch(ctx, qs.mgr.Keeper(), asset)

	// parse amount
	amount, err := cosmos.ParseUint(req.Amount)
	if err != nil {
		return nil, fmt.Errorf("bad amount: %w", err)
	}

	// parse min out
	minOut := sdkmath.ZeroUint()
	if len(req.MinOut) > 0 {
		minOut, err = cosmos.ParseUint(req.MinOut)
		if err != nil {
			return nil, fmt.Errorf("bad min out: %w", err)
		}
	}

	// Affiliate fee in RUNE
	affiliateRuneAmt := sdkmath.ZeroUint()

	// parse affiliate
	affiliate, affiliateMemo, affiliateBps, amt, affiliateAmt, err := quoteHandleAffiliate(ctx, qs.mgr, req.Affiliate, req.AffiliateBps, amount)
	if err != nil {
		return nil, err
	}

	if affiliate != common.NoAddress && !affiliateBps.IsZero() {
		affCoin := common.NewCoin(asset, affiliateAmt)
		gasCoin := common.NewCoin(asset.GetChain().GetGasAsset(), cosmos.OneUint())
		fakeTx := common.NewTx(common.BlankTxID, common.NoopAddress, common.NoopAddress, common.NewCoins(affCoin), common.Gas{gasCoin}, "noop")
		affiliateSwap := NewMsgSwap(fakeTx, common.RuneAsset(), affiliate, cosmos.ZeroUint(), common.NoAddress, cosmos.ZeroUint(), "", "", nil, 0, 0, 0, nil)

		_, affiliateRuneAmt, _, err = quoteSimulateSwap(ctx, qs.mgr, affiliateAmt, affiliateSwap, 1)
		if err == nil {
			// skim fee off collateral amount
			amount = amt
		} else {
			affiliateRuneAmt = sdkmath.ZeroUint()
		}
	}

	// parse target asset
	targetAsset, err := common.NewAssetWithShortCodes(qs.mgr.GetVersion(), req.ToAsset)
	if err != nil {
		return nil, fmt.Errorf("bad target asset: %w", err)
	}
	targetAsset = fuzzyAssetMatch(ctx, qs.mgr.Keeper(), targetAsset)

	// parse destination address or generate a random one
	sendMemo := true
	var destination common.Address
	if len(req.Destination) > 0 {
		destination, err = quoteParseAddress(ctx, qs.mgr, req.Destination, targetAsset.Chain)
		if err != nil {
			return nil, fmt.Errorf("bad destination address: %w", err)
		}

	} else {
		destination, err = types.GetRandomPubKey().GetAddress(targetAsset.Chain)
		if err != nil {
			return nil, fmt.Errorf("failed to generate address: %w", err)
		}
		sendMemo = false // do not send memo if destination was random
	}

	// check that destination and affiliate are not the same
	if destination.Equals(affiliate) {
		return nil, fmt.Errorf("destination and affiliate should not be the same")
	}

	// generate random address for collateral owner
	collateralOwner, err := types.GetRandomPubKey().GetAddress(asset.Chain)
	if err != nil {
		return nil, fmt.Errorf("failed to generate address: %w", err)
	}

	// create message for simulation
	msg := &types.MsgLoanOpen{
		Owner:            collateralOwner,
		CollateralAsset:  asset,
		CollateralAmount: amount,
		TargetAddress:    destination,
		TargetAsset:      targetAsset,
		MinOut:           minOut,

		// We calculate the affiliate fee manually as handler_open_loan expects a TxVoter to
		// get the affiliate params from the memo
		AffiliateBasisPoints: cosmos.ZeroUint(),

		// TODO: support aggregator
		Aggregator:              "",
		AggregatorTargetAddress: "",
		AggregatorTargetLimit:   sdkmath.ZeroUint(),
	}

	// simulate message handling
	events, err := simulate(ctx, qs.mgr, msg)
	if err != nil {
		return nil, err
	}

	// create response
	res := &types.QueryQuoteLoanOpenResponse{
		Fees: &types.QuoteFees{
			Asset: targetAsset.String(),
		},
		Expiry:  time.Now().Add(quoteExpiration).Unix(),
		Warning: quoteWarning,
		Notes:   asset.Chain.InboundNotes(),
	}

	// estimate the inbound info
	inboundAddress, routerAddress, inboundConfirmations, err := quoteInboundInfo(ctx, qs.mgr, amount, asset.Chain, asset)
	if err != nil {
		return nil, err
	}
	res.InboundAddress = inboundAddress.String()
	if inboundConfirmations > 0 {
		res.InboundConfirmationBlocks = inboundConfirmations
		res.InboundConfirmationSeconds = inboundConfirmations * asset.Chain.ApproximateBlockMilliseconds() / 1000
	}

	// set info fields
	if asset.Chain.IsEVM() {
		res.Router = routerAddress.String()
	}
	if !asset.Chain.DustThreshold().IsZero() {
		res.DustThreshold = asset.Chain.DustThreshold().String()
	}

	// sum liquidity fees in rune from all swap events
	outboundFee := sdkmath.ZeroUint()
	liquidityFee := sdkmath.ZeroUint()
	affiliateFee := affiliateRuneAmt
	expectedAmountOut := sdkmath.ZeroUint()
	finalEmitAmount := sdkmath.ZeroUint() // used to calculate slippage
	streamingSwapBlocks := int64(0)
	streamingSwapSeconds := int64(0)

	// iterate events in reverse order
	for i := len(events) - 1; i >= 0; i-- {
		e := events[i]
		em := eventMap(e)

		switch e.Type {

		// use final outbound event as expected amount - scheduled_outbound (L1) or outbound (native)
		case "scheduled_outbound":
			if res.ExpectedAmountOut == "" { // if not empty we already saw the last outbound event
				res.ExpectedAmountOut = em["coin_amount"]
				expectedAmountOut = sdkmath.NewUintFromString(em["coin_amount"])
				if em["coin_asset"] != targetAsset.String() { // should be unreachable
					return nil, fmt.Errorf("unexpected outbound asset: %s", em["coin_asset"])
				}

				// estimate the outbound info
				outboundDelay, err := quoteOutboundInfo(ctx, qs.mgr, common.NewCoin(targetAsset, sdkmath.NewUintFromString(res.ExpectedAmountOut)))
				if err != nil {
					return nil, err
				}
				res.OutboundDelayBlocks = outboundDelay
				res.OutboundDelaySeconds = outboundDelay * common.THORChain.ApproximateBlockMilliseconds() / 1000
			}
		case "outbound":
			coin, err := common.ParseCoin(em["coin"])
			if err != nil {
				return nil, fmt.Errorf("failed to parse coin: %w", err)
			}
			toAddress, _ := common.NewAddress(em["to"])

			// check for the outbound event
			if toAddress.Equals(destination) {
				res.ExpectedAmountOut = coin.Amount.String()
				expectedAmountOut = coin.Amount

				if !coin.Asset.Equals(targetAsset) { // should be unreachable
					return nil, fmt.Errorf("unexpected outbound asset: %s", coin.Asset)
				}
			}

		// sum liquidity fee in rune for all swap events
		case "swap":
			liquidityFee = liquidityFee.Add(sdkmath.NewUintFromString(em["liquidity_fee_in_rune"]))
			coin, err := common.ParseCoin(em["emit_asset"])
			if err != nil {
				return nil, fmt.Errorf("failed to parse coin: %w", err)
			}
			if coin.Asset.Equals(targetAsset) {
				finalEmitAmount = coin.Amount
			}
			swapQuantity, err := cosmos.ParseUint(em["streaming_swap_quantity"])
			if err != nil {
				return nil, fmt.Errorf("bad quantity: %w", err)
			}
			streamingSwapBlocks += swapQuantity.BigInt().Int64()

		// extract loan data from loan open event
		case "loan_open":
			res.ExpectedCollateralizationRatio = em["collateralization_ratio"]
			res.ExpectedCollateralDeposited = em["collateral_deposited"]
			res.ExpectedDebtIssued = em["debt_issued"]

		// catch refund if there was an issue
		case "refund":
			if em["reason"] != "" {
				return nil, fmt.Errorf("failed to simulate swap: %s", em["reason"])
			}

		// set outbound fee from fee event
		case "fee":
			coin, err := common.ParseCoin(em["coins"])
			if err != nil {
				return nil, fmt.Errorf("failed to parse coin: %w", err)
			}
			res.Fees.Outbound = coin.Amount.String() // already in target asset
			res.Fees.Asset = coin.Asset.String()
			outboundFee = coin.Amount

			if !coin.Asset.Equals(targetAsset) { // should be unreachable
				return nil, fmt.Errorf("unexpected fee asset: %s", coin.Asset)
			}
		}
	}

	// convert fees to target asset if it is not rune
	if !targetAsset.Equals(common.RuneNative) {
		targetPool, err := qs.mgr.Keeper().GetPool(ctx, targetAsset)
		if err != nil {
			return nil, fmt.Errorf("failed to get pool: %w", err)
		}
		affiliateFee = targetPool.RuneValueInAsset(affiliateRuneAmt)
		liquidityFee = targetPool.RuneValueInAsset(liquidityFee)
	}
	slippageBps := liquidityFee.MulUint64(10000).Quo(finalEmitAmount.Add(liquidityFee))

	// set fee info
	res.Fees.Liquidity = liquidityFee.String()
	totalFees := liquidityFee.Add(outboundFee).Add(affiliateFee)
	res.Fees.Total = totalFees.String()
	res.Fees.SlippageBps = slippageBps.BigInt().Int64()
	res.Fees.TotalBps = totalFees.MulUint64(10000).Quo(expectedAmountOut.Add(totalFees)).BigInt().Int64()
	if !affiliateFee.IsZero() {
		res.Fees.Affiliate = affiliateFee.String()
	}

	// generate memo
	if sendMemo {
		memo := &mem.LoanOpenMemo{
			MemoBase: mem.MemoBase{
				TxType: TxLoanOpen,
			},
			TargetAsset:          targetAsset,
			TargetAddress:        destination,
			MinOut:               minOut,
			AffiliateAddress:     common.Address(affiliateMemo),
			AffiliateBasisPoints: affiliateBps,
			DexTargetLimit:       sdkmath.ZeroUint(),
		}

		// if from asset chain has memo length restrictions use a prefix
		memoString := memo.String()
		if len(memoString) > asset.Chain.MaxMemoLength() {
			if len(memo.ShortString()) < len(memoString) { // use short codes if available
				memoString = memo.ShortString()
			} else { // otherwise attempt to shorten
				fuzzyAsset, err := quoteReverseFuzzyAsset(ctx, qs.mgr, targetAsset)
				if err == nil {
					memo.TargetAsset = fuzzyAsset
					memoString = memo.String()
				}
			}

			// this is the shortest we can make it
			if len(memoString) > asset.Chain.MaxMemoLength() {
				return nil, fmt.Errorf("generated memo too long for source chain")
			}
		}

		res.Memo = memoString
	}

	minLoanOpenAmount, err := calculateMinSwapAmount(ctx, qs.mgr, asset, targetAsset, cosmos.ZeroUint())
	if err != nil {
		return nil, fmt.Errorf("Failed to calculate min amount in: %s", err.Error())
	}
	res.RecommendedMinAmountIn = minLoanOpenAmount.String()

	streamingSwapSeconds += streamingSwapBlocks * common.THORChain.ApproximateBlockMilliseconds() / 1000

	if res.InboundConfirmationSeconds != 0 {
		value := res.InboundConfirmationSeconds
		res.TotalOpenLoanSeconds = streamingSwapSeconds + res.OutboundDelaySeconds + value
	} else {
		res.TotalOpenLoanSeconds = streamingSwapSeconds + res.OutboundDelaySeconds
	}

	res.StreamingSwapBlocks = streamingSwapBlocks
	res.StreamingSwapSeconds = streamingSwapSeconds

	// set inbound recommended gas
	inboundGas := qs.mgr.GasMgr().GetGasRate(ctx, asset.Chain)
	res.RecommendedGasRate = inboundGas.String()
	res.GasRateUnits = asset.Chain.GetGasUnits()

	return res, nil
}

// -------------------------------------------------------------------------------------
// Loan Close
// -------------------------------------------------------------------------------------

func quoteSimulateCloseLoan(ctx cosmos.Context, mgr *Mgrs, msg *MsgLoanRepayment) (
	res *types.QueryQuoteLoanCloseResponse, err error,
) {
	res = &types.QueryQuoteLoanCloseResponse{
		Fees: &types.QuoteFees{
			Asset: msg.CollateralAsset.String(),
		},
		Expiry:  time.Now().Add(quoteExpiration).Unix(),
		Warning: quoteWarning,
		Notes:   msg.Coin.Asset.Chain.InboundNotes(),
	}

	// simulate message handling
	events, err := simulate(ctx, mgr, msg)
	if err != nil {
		return nil, err
	}

	// estimate the inbound info
	inboundAddress, routerAddress, inboundConfirmations, err := quoteInboundInfo(ctx, mgr, msg.Coin.Amount, msg.Coin.Asset.GetChain(), msg.Coin.Asset)
	if err != nil {
		return nil, err
	}
	res.InboundAddress = inboundAddress.String()
	if inboundConfirmations > 0 {
		res.InboundConfirmationBlocks = inboundConfirmations
		res.InboundConfirmationSeconds = inboundConfirmations * msg.Coin.Asset.GetChain().ApproximateBlockMilliseconds() / 1000
	}

	// set info fields
	if msg.Coin.Asset.Chain.IsEVM() {
		res.Router = routerAddress.String()
	}
	if !msg.Coin.Asset.Chain.DustThreshold().IsZero() {
		res.DustThreshold = msg.Coin.Asset.Chain.DustThreshold().String()
	}

	// sum liquidity fees in rune from all swap events
	outboundFee := sdkmath.ZeroUint()
	repaymentLiquidityFee := sdkmath.ZeroUint()
	outboundLiquidityFee := sdkmath.ZeroUint()
	affiliateFee := sdkmath.ZeroUint()
	expectedAmountOut := sdkmath.ZeroUint()
	streamingSwapBlocks := int64(0)
	streamingSwapSeconds := int64(0)
	var repaymentEmit, outboundEmit common.Coin

	// iterate events in reverse order
	for i := len(events) - 1; i >= 0; i-- {
		e := events[i]
		em := eventMap(e)

		switch e.Type {

		// use final outbound event as expected amount - scheduled_outbound (L1) or outbound (native)
		case "scheduled_outbound":
			if res.ExpectedAmountOut == "" { // if not empty we already saw the last outbound event
				res.ExpectedAmountOut = em["coin_amount"]
				expectedAmountOut = sdkmath.NewUintFromString(em["coin_amount"])
				if em["coin_asset"] != msg.CollateralAsset.String() { // should be unreachable
					return nil, fmt.Errorf("unexpected outbound asset: %s", em["coin_asset"])
				}

				// estimate the outbound info
				outboundDelay, err := quoteOutboundInfo(ctx, mgr, common.NewCoin(msg.CollateralAsset, sdkmath.NewUintFromString(res.ExpectedAmountOut)))
				if err != nil {
					return nil, err
				}
				res.OutboundDelayBlocks = outboundDelay
				res.OutboundDelaySeconds = outboundDelay * common.THORChain.ApproximateBlockMilliseconds() / 1000
			}
		case "outbound":
			// track coin and to address
			coin, err := common.ParseCoin(em["coin"])
			if err != nil {
				return nil, fmt.Errorf("failed to parse coin: %w", err)
			}
			toAddress, _ := common.NewAddress(em["to"])

			// check for the outbound event
			if toAddress.Equals(msg.Owner) {
				res.ExpectedAmountOut = coin.Amount.String()
				expectedAmountOut = coin.Amount

				if !coin.Asset.Equals(msg.CollateralAsset) { // should be unreachable
					return nil, fmt.Errorf("unexpected outbound asset: %s", coin.Asset)
				}
			}

		// sum liquidity fee in rune for all swap events
		case "swap":
			coin, err := common.ParseCoin(em["emit_asset"])
			if err != nil {
				return nil, fmt.Errorf("failed to parse coin: %w", err)
			}
			swapQuantity, err := cosmos.ParseUint(em["streaming_swap_quantity"])
			if err != nil {
				return nil, fmt.Errorf("bad amount: %w", err)
			}
			streamingSwapBlocks += swapQuantity.BigInt().Int64()
			switch {
			case coin.Asset.Equals(common.TOR):
				repaymentEmit = coin
				repaymentLiquidityFee = repaymentLiquidityFee.Add(sdkmath.NewUintFromString(em["liquidity_fee_in_rune"]))
			case !coin.IsRune():
				outboundEmit = coin
				outboundLiquidityFee = outboundLiquidityFee.Add(sdkmath.NewUintFromString(em["liquidity_fee_in_rune"]))
			default:
				inCoin, err := common.ParseCoin(em["coin"])
				if err != nil {
					return nil, fmt.Errorf("failed to parse coin: %w", err)
				}
				if inCoin.Asset.IsDerivedAsset() {
					outboundLiquidityFee = outboundLiquidityFee.Add(sdkmath.NewUintFromString(em["liquidity_fee_in_rune"]))
				} else {
					repaymentLiquidityFee = repaymentLiquidityFee.Add(sdkmath.NewUintFromString(em["liquidity_fee_in_rune"]))
				}
			}

		// extract loan data from loan close event
		case "loan_repayment":
			res.ExpectedCollateralWithdrawn = em["collateral_withdrawn"]
			res.ExpectedDebtRepaid = em["debt_repaid"]

		// catch refund if there was an issue
		case "refund":
			if em["reason"] != "" {
				return nil, fmt.Errorf("failed to simulate loan close: %s", em["reason"])
			}

		// set outbound fee from fee event
		case "fee":
			coin, err := common.ParseCoin(em["coins"])
			if err != nil {
				return nil, fmt.Errorf("failed to parse coin: %w", err)
			}
			res.Fees.Outbound = coin.Amount.String() // already in collateral asset
			res.Fees.Asset = coin.Asset.String()
			outboundFee = coin.Amount

			if !coin.Asset.Equals(msg.CollateralAsset) { // should be unreachable
				return nil, fmt.Errorf("unexpected fee asset: %s", coin.Asset)
			}

		}
	}

	// calculate emit values in rune
	torPool, err := mgr.Keeper().GetPool(ctx, common.TOR)
	if err != nil {
		return nil, fmt.Errorf("failed to get pool: %w", err)
	}
	repaymentEmitRune := torPool.RuneValueInAsset(repaymentEmit.Amount)
	outPool, err := mgr.Keeper().GetPool(ctx, outboundEmit.Asset)
	if err != nil {
		return nil, fmt.Errorf("failed to get pool: %w", err)
	}
	outboundEmitRune := outPool.RuneValueInAsset(outboundEmit.Amount)

	// slippage calculation is weighted to repayment and outbound amounts
	outboundSlip := sdkmath.ZeroUint()
	if !outboundEmitRune.IsZero() {
		outboundSlip = outboundLiquidityFee.MulUint64(10000).Quo(outboundEmitRune.Add(outboundLiquidityFee))
	}
	repaymentSlip := repaymentLiquidityFee.MulUint64(10000).Quo(repaymentEmitRune.Add(repaymentLiquidityFee))
	slippageBps := repaymentSlip.Mul(repaymentEmitRune).Add(outboundSlip.Mul(outboundEmitRune)).Quo(repaymentEmitRune.Add(outboundEmitRune))

	// convert fees to target asset if it is not rune
	liquidityFee := repaymentLiquidityFee.Add(outboundLiquidityFee)
	if !msg.CollateralAsset.Equals(common.RuneNative) {
		loanPool, err := mgr.Keeper().GetPool(ctx, msg.CollateralAsset)
		if err != nil {
			return nil, fmt.Errorf("failed to get pool: %w", err)
		}
		affiliateFee = loanPool.RuneValueInAsset(affiliateFee)
		liquidityFee = loanPool.RuneValueInAsset(liquidityFee)
	}

	// set fee info
	res.Fees.Liquidity = liquidityFee.String()
	totalFees := liquidityFee.Add(outboundFee).Add(affiliateFee)
	res.Fees.Total = totalFees.String()
	res.Fees.SlippageBps = slippageBps.BigInt().Int64()
	if !expectedAmountOut.IsZero() {
		res.Fees.TotalBps = totalFees.MulUint64(10000).Quo(expectedAmountOut).BigInt().Int64()
	} else {
		res.Fees.TotalBps = res.Fees.SlippageBps
	}
	if !affiliateFee.IsZero() {
		res.Fees.Affiliate = affiliateFee.String()
	}

	// generate memo
	memo := &mem.LoanRepaymentMemo{
		MemoBase: mem.MemoBase{
			TxType: TxLoanRepayment,
			Asset:  msg.CollateralAsset,
		},
		Owner:  msg.Owner,
		MinOut: msg.MinOut,
	}
	res.Memo = memo.String()

	minLoanCloseAmount, err := calculateMinSwapAmount(ctx, mgr, msg.Coin.Asset, msg.CollateralAsset, cosmos.ZeroUint())
	if err != nil {
		return nil, fmt.Errorf("Failed to calculate min amount in: %s", err.Error())
	}
	res.RecommendedMinAmountIn = minLoanCloseAmount.String()

	streamingSwapSeconds += streamingSwapBlocks * common.THORChain.ApproximateBlockMilliseconds() / 1000

	if res.InboundConfirmationSeconds != 0 {
		value := res.InboundConfirmationSeconds
		res.TotalRepaySeconds = streamingSwapSeconds + res.OutboundDelaySeconds + value
	} else {
		res.TotalRepaySeconds = streamingSwapSeconds + res.OutboundDelaySeconds
	}

	res.StreamingSwapBlocks = streamingSwapBlocks
	res.StreamingSwapSeconds = streamingSwapSeconds
	res.ExpectedAmountIn = msg.Coin.Amount.String()

	return res, nil
}

func (qs queryServer) queryQuoteLoanClose(ctx cosmos.Context, req *types.QueryQuoteLoanCloseRequest) (*types.QueryQuoteLoanCloseResponse, error) {
	// validate required parameters
	if len(req.FromAsset) == 0 {
		return nil, fmt.Errorf("missing from_asset parameter")
	}
	if len(req.ToAsset) == 0 {
		return nil, fmt.Errorf("missing to_asset parameter")
	}
	if len(req.RepayBps) == 0 {
		return nil, fmt.Errorf("missing repay_bps parameter")
	}
	if len(req.LoanOwner) == 0 {
		return nil, fmt.Errorf("missing loan_owner parameter")
	}

	// parse asset
	asset, err := common.NewAssetWithShortCodes(qs.mgr.GetVersion(), req.FromAsset)
	if err != nil {
		return nil, fmt.Errorf("bad asset: %w", err)
	}
	asset = fuzzyAssetMatch(ctx, qs.mgr.Keeper(), asset)

	// parse repayment bps
	repayBps, err := cosmos.ParseUint(req.RepayBps)
	if err != nil {
		return nil, fmt.Errorf("bad amount: %w", err)
	}

	// parse min out
	minOut := sdkmath.ZeroUint()
	if len(req.MinOut) > 0 {
		minOut, err = cosmos.ParseUint(req.MinOut)
		if err != nil {
			return nil, fmt.Errorf("bad min out: %w", err)
		}
	}

	// parse loan asset
	loanAsset, err := common.NewAssetWithShortCodes(qs.mgr.GetVersion(), req.ToAsset)
	if err != nil {
		return nil, fmt.Errorf("bad loan asset: %w", err)
	}
	loanAsset = fuzzyAssetMatch(ctx, qs.mgr.Keeper(), loanAsset)

	// parse loan owner
	loanOwner, err := common.NewAddress(req.LoanOwner)
	if err != nil {
		return nil, fmt.Errorf("bad loan owner: %w", err)
	}

	// generate random from address
	fromAddress, err := types.GetRandomPubKey().GetAddress(asset.Chain)
	if err != nil {
		return nil, fmt.Errorf("bad from address: %w", err)
	}

	// validate if it is valid collateral asset
	key := "LENDING-" + loanAsset.GetDerivedAsset().MimirString()
	val, err := qs.mgr.Keeper().GetMimir(ctx, key)
	if err != nil {
		return nil, fmt.Errorf("ail to fetch LENDING key: %w", err)
	}
	if val <= 0 {
		return nil, fmt.Errorf("Lending is not available for this collateral asset")
	}

	loan, err := qs.mgr.Keeper().GetLoan(ctx, loanAsset, loanOwner)
	if err != nil {
		return nil, fmt.Errorf("failed to get loan: %w", err)
	}

	poolRepayment, err := qs.mgr.Keeper().GetPool(ctx, asset)
	if err != nil {
		return nil, fmt.Errorf("failed to get pool: %w", err)
	}

	poolThorAsset, err := qs.mgr.Keeper().GetPool(ctx, common.TOR)
	if err != nil {
		return nil, fmt.Errorf("failed to get pool: %w", err)
	}

	pendingDebt := loan.DebtIssued.Sub(loan.DebtRepaid)
	totalPendingDebtInRune := poolThorAsset.AssetValueInRune(pendingDebt)
	totalPendingDebtInRepaymentAsset := totalPendingDebtInRune

	if !asset.IsRune() {
		totalPendingDebtInRepaymentAsset = poolRepayment.RuneValueInAsset(totalPendingDebtInRune)
	}

	minBP := qs.mgr.Keeper().GetConfigInt64(ctx, constants.StreamingSwapMinBPFee)
	initialThresholdBasisPoints := sdkmath.NewUint(uint64(minBP)) // Initial threshold to start looking for the target amount
	amountInTorToRepay := pendingDebt.Mul(repayBps).Quo(sdkmath.NewUint(10_000))
	amountToRepay := totalPendingDebtInRepaymentAsset.Mul(repayBps).Quo(sdkmath.NewUint(10_000))
	incrementBasedOnThreshold := amountToRepay.Mul(initialThresholdBasisPoints).Quo(sdkmath.NewUint(10_000))
	amountPlusThresholdToRepay := amountToRepay.Add(incrementBasedOnThreshold)

	msg := &types.MsgLoanRepayment{
		Owner:           loanOwner,
		CollateralAsset: loanAsset,
		Coin:            common.NewCoin(asset, amountPlusThresholdToRepay),
		From:            fromAddress,
		MinOut:          minOut,
	}

	res, err := quoteSimulateCloseLoan(ctx, qs.mgr, msg)
	if err != nil {
		return nil, err
	}

	thresholdBasisPoint := initialThresholdBasisPoints

	for thresholdBasisPoint.LTE(sdkmath.NewUint(1500)) { // Arbitrary cap for the threshold of 1500 BPS to avoid harmful requests.

		exptectedDebtRepaid, err := cosmos.ParseUint(res.ExpectedDebtRepaid)
		if err != nil {
			return nil, fmt.Errorf("bad exptectedDebtRepaid: %w", err)
		}

		if exptectedDebtRepaid.GTE(amountInTorToRepay) {
			break
		}

		// Arbitrarily increment by 10 BPS per iteration until the target is met. A higher amount results in less server load but also less accurate calculations
		thresholdBasisPoint = thresholdBasisPoint.Add(sdkmath.NewUint(10))

		// Resimulate with new threshold
		increment := amountToRepay.Mul(thresholdBasisPoint).Quo(sdkmath.NewUint(10_000))
		newAmount := amountToRepay.Add(increment)
		msg.Coin.Amount = newAmount
		res, err = quoteSimulateCloseLoan(ctx, qs.mgr, msg)
		if err != nil {
			return nil, err
		}
	}

	// set inbound recommended gas for non-native in asset
	if !asset.Chain.IsTHORChain() {
		inboundGas := qs.mgr.GasMgr().GetGasRate(ctx, asset.Chain)
		res.RecommendedGasRate = inboundGas.String()
		res.GasRateUnits = asset.Chain.GetGasUnits()
	}

	return res, nil
}
