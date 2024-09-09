package thorchain

import (
	"errors"
	"fmt"

	. "gopkg.in/check.v1"

	"gitlab.com/thorchain/thornode/common"
	"gitlab.com/thorchain/thornode/common/cosmos"
	"gitlab.com/thorchain/thornode/constants"
	"gitlab.com/thorchain/thornode/x/thorchain/keeper"
)

type HandlerAddLiquiditySuite struct{}

var _ = Suite(&HandlerAddLiquiditySuite{})

type MockAddLiquidityKeeper struct {
	keeper.KVStoreDummy
	currentPool       Pool
	activeNodeAccount NodeAccount
	failGetPool       bool
	lp                LiquidityProvider
	pol               ProtocolOwnedLiquidity
	polAddress        common.Address
}

func (m *MockAddLiquidityKeeper) PoolExist(_ cosmos.Context, asset common.Asset) bool {
	return m.currentPool.Asset.Equals(asset)
}

func (m *MockAddLiquidityKeeper) GetPools(_ cosmos.Context) (Pools, error) {
	return Pools{m.currentPool}, nil
}

func (m *MockAddLiquidityKeeper) GetPool(_ cosmos.Context, _ common.Asset) (Pool, error) {
	if m.failGetPool {
		return Pool{}, errors.New("fail to get pool")
	}
	return m.currentPool, nil
}

func (m *MockAddLiquidityKeeper) SetPool(_ cosmos.Context, pool Pool) error {
	m.currentPool = pool
	return nil
}

func (m *MockAddLiquidityKeeper) GetModuleAddress(mod string) (common.Address, error) {
	return m.polAddress, nil
}

func (m *MockAddLiquidityKeeper) GetPOL(_ cosmos.Context) (ProtocolOwnedLiquidity, error) {
	return m.pol, nil
}

func (m *MockAddLiquidityKeeper) SetPOL(_ cosmos.Context, pol ProtocolOwnedLiquidity) error {
	m.pol = pol
	return nil
}

func (m *MockAddLiquidityKeeper) ListValidatorsWithBond(_ cosmos.Context) (NodeAccounts, error) {
	return NodeAccounts{m.activeNodeAccount}, nil
}

func (m *MockAddLiquidityKeeper) ListActiveValidators(_ cosmos.Context) (NodeAccounts, error) {
	return NodeAccounts{m.activeNodeAccount}, nil
}

func (m *MockAddLiquidityKeeper) GetNodeAccount(_ cosmos.Context, addr cosmos.AccAddress) (NodeAccount, error) {
	if m.activeNodeAccount.NodeAddress.Equals(addr) {
		return m.activeNodeAccount, nil
	}
	return NodeAccount{}, errors.New("not exist")
}

func (m *MockAddLiquidityKeeper) GetLiquidityProvider(_ cosmos.Context, asset common.Asset, addr common.Address) (LiquidityProvider, error) {
	return m.lp, nil
}

func (m *MockAddLiquidityKeeper) SetLiquidityProvider(ctx cosmos.Context, lp LiquidityProvider) {
	m.lp = lp
}

func (m *MockAddLiquidityKeeper) AddOwnership(ctx cosmos.Context, coin common.Coin, _ cosmos.AccAddress) error {
	m.lp.Units = m.lp.Units.Add(coin.Amount)
	return nil
}

func (m *MockAddLiquidityKeeper) GetAsgardVaults(_ cosmos.Context) (Vaults, error) {
	return Vaults{
		{
			Status: ActiveVault,
			Type:   AsgardVault,
			Coins: common.NewCoins(
				common.NewCoin(m.currentPool.Asset, m.currentPool.BalanceAsset),
			),
		},
	}, nil
}

func (s *HandlerAddLiquiditySuite) SetUpSuite(c *C) {
	SetupConfigForTest()
}

func (s *HandlerAddLiquiditySuite) TestAddLiquidityHandler(c *C) {
	var err error
	ctx, mgr := setupManagerForTest(c)
	activeNodeAccount := GetRandomValidatorNode(NodeActive)
	runeAddr := GetRandomRUNEAddress()
	ethAddr := GetRandomETHAddress()
	pool := NewPool()
	pool.Asset = common.ETHAsset
	pool.Status = PoolAvailable
	k := &MockAddLiquidityKeeper{
		activeNodeAccount: activeNodeAccount,
		currentPool:       pool,
		lp: LiquidityProvider{
			Asset:             common.ETHAsset,
			RuneAddress:       runeAddr,
			AssetAddress:      ethAddr,
			Units:             cosmos.ZeroUint(),
			PendingRune:       cosmos.ZeroUint(),
			PendingAsset:      cosmos.ZeroUint(),
			RuneDepositValue:  cosmos.ZeroUint(),
			AssetDepositValue: cosmos.ZeroUint(),
		},
		pol:        NewProtocolOwnedLiquidity(),
		polAddress: GetRandomRUNEAddress(),
	}
	mgr.K = k
	// happy path
	addHandler := NewAddLiquidityHandler(mgr)
	addTxHash := GetRandomTxHash()
	tx := common.NewTx(
		addTxHash,
		runeAddr,
		GetRandomETHAddress(),
		common.Coins{common.NewCoin(common.ETHAsset, cosmos.NewUint(common.One*5))},
		common.Gas{
			common.NewCoin(common.ETHAsset, cosmos.NewUint(10000)),
		},
		"add:ETH",
	)
	msg := NewMsgAddLiquidity(
		tx,
		common.ETHAsset,
		cosmos.NewUint(100*common.One),
		cosmos.ZeroUint(),
		runeAddr,
		ethAddr,
		common.NoAddress, cosmos.ZeroUint(),
		activeNodeAccount.NodeAddress)
	_, err = addHandler.Run(ctx, msg)
	c.Assert(err, IsNil)

	midLiquidityPool, err := mgr.Keeper().GetPool(ctx, common.ETHAsset)
	c.Assert(err, IsNil)
	c.Assert(midLiquidityPool.PendingInboundRune.String(), Equals, "10000000000")

	msg.RuneAmount = cosmos.ZeroUint()
	msg.AssetAmount = cosmos.NewUint(100 * common.One)
	_, err = addHandler.Run(ctx, msg)
	c.Assert(err, IsNil)

	postLiquidityPool, err := mgr.Keeper().GetPool(ctx, common.ETHAsset)
	c.Assert(err, IsNil)
	c.Assert(postLiquidityPool.BalanceAsset.String(), Equals, "10000000000")
	c.Assert(postLiquidityPool.BalanceRune.String(), Equals, "10000000000")
	c.Assert(postLiquidityPool.PendingInboundAsset.String(), Equals, "0")
	c.Assert(postLiquidityPool.PendingInboundRune.String(), Equals, "0")

	pol, err := mgr.Keeper().GetPOL(ctx)
	c.Assert(err, IsNil)
	c.Check(pol.RuneDeposited.Uint64(), Equals, uint64(0))
}

func (s *HandlerAddLiquiditySuite) TestAddLiquidityHandler_NoPool_ShouldCreateNewPool(c *C) {
	activeNodeAccount := GetRandomValidatorNode(NodeActive)
	activeNodeAccount.Bond = cosmos.NewUint(1000000 * common.One)
	runeAddr := GetRandomRUNEAddress()
	ethAddr := GetRandomETHAddress()
	pool := NewPool()
	pool.Status = PoolAvailable
	k := &MockAddLiquidityKeeper{
		activeNodeAccount: activeNodeAccount,
		currentPool:       pool,
		lp: LiquidityProvider{
			Asset:             common.ETHAsset,
			RuneAddress:       runeAddr,
			AssetAddress:      ethAddr,
			Units:             cosmos.ZeroUint(),
			PendingRune:       cosmos.ZeroUint(),
			PendingAsset:      cosmos.ZeroUint(),
			RuneDepositValue:  cosmos.ZeroUint(),
			AssetDepositValue: cosmos.ZeroUint(),
		},
	}
	// happy path
	ctx, mgr := setupManagerForTest(c)
	mgr.K = k
	addHandler := NewAddLiquidityHandler(mgr)
	preLiquidityPool, err := k.GetPool(ctx, common.ETHAsset)
	c.Assert(err, IsNil)
	c.Assert(preLiquidityPool.IsEmpty(), Equals, true)
	addTxHash := GetRandomTxHash()
	tx := common.NewTx(
		addTxHash,
		runeAddr,
		GetRandomETHAddress(),
		common.Coins{common.NewCoin(common.ETHAsset, cosmos.NewUint(common.One*5))},
		common.Gas{
			common.NewCoin(common.ETHAsset, cosmos.NewUint(10000)),
		},
		"add:ETH",
	)
	mgr.constAccessor = constants.NewDummyConstants(map[constants.ConstantName]int64{
		constants.MaximumLiquidityRune: 600_000_00000000,
	}, map[constants.ConstantName]bool{
		constants.StrictBondLiquidityRatio: true,
	}, map[constants.ConstantName]string{})

	msg := NewMsgAddLiquidity(
		tx,
		common.ETHAsset,
		cosmos.NewUint(100*common.One),
		cosmos.NewUint(100*common.One),
		runeAddr,
		ethAddr,
		common.NoAddress, cosmos.ZeroUint(),
		activeNodeAccount.NodeAddress)
	_, err = addHandler.Run(ctx, msg)
	c.Assert(err, IsNil)
	postLiquidityPool, err := k.GetPool(ctx, common.ETHAsset)
	c.Assert(err, IsNil)
	c.Assert(postLiquidityPool.BalanceAsset.String(), Equals, preLiquidityPool.BalanceAsset.Add(msg.AssetAmount).String())
	c.Assert(postLiquidityPool.BalanceRune.String(), Equals, preLiquidityPool.BalanceRune.Add(msg.RuneAmount).String())
}

func (s *HandlerAddLiquiditySuite) TestAddLiquidityHandlerValidation(c *C) {
	ctx, _ := setupKeeperForTest(c)
	activeNodeAccount := GetRandomValidatorNode(NodeActive)
	runeAddr := GetRandomRUNEAddress()
	ethAddr := GetRandomETHAddress()
	ethSynthAsset, _ := common.NewAsset("ETH/ETH")
	mainnetBTCAddr, err := common.NewAddress("bc1qy0tj9fh0u6fgz0mejjp6776z6kugych0zwrkwr")
	c.Assert(err, IsNil)
	tx := common.NewTx(
		GetRandomTxHash(),
		GetRandomRUNEAddress(),
		GetRandomRUNEAddress(),
		common.Coins{common.NewCoin(ethSynthAsset, cosmos.NewUint(common.One*5))},
		common.Gas{
			{Asset: common.RuneNative, Amount: cosmos.NewUint(1 * common.One)},
		},
		"add:ETH.ETH",
	)

	k := &MockAddLiquidityKeeper{
		activeNodeAccount: activeNodeAccount,
		currentPool: Pool{
			BalanceRune:  cosmos.ZeroUint(),
			BalanceAsset: cosmos.ZeroUint(),
			Asset:        common.ETHAsset,
			LPUnits:      cosmos.ZeroUint(),
			Status:       PoolAvailable,
		},
		lp: LiquidityProvider{
			Asset:             common.ETHAsset,
			RuneAddress:       runeAddr,
			AssetAddress:      ethAddr,
			Units:             cosmos.ZeroUint(),
			PendingRune:       cosmos.ZeroUint(),
			PendingAsset:      cosmos.ZeroUint(),
			RuneDepositValue:  cosmos.ZeroUint(),
			AssetDepositValue: cosmos.ZeroUint(),
		},
	}
	testCases := []struct {
		name           string
		msg            *MsgAddLiquidity
		expectedResult error
	}{
		{
			name:           "empty signer should fail",
			msg:            NewMsgAddLiquidity(GetRandomTx(), common.ETHAsset, cosmos.NewUint(common.One*5), cosmos.NewUint(common.One*5), GetRandomETHAddress(), GetRandomETHAddress(), common.NoAddress, cosmos.ZeroUint(), cosmos.AccAddress{}),
			expectedResult: errAddLiquidityFailValidation,
		},
		{
			name:           "empty asset should fail",
			msg:            NewMsgAddLiquidity(GetRandomTx(), common.Asset{}, cosmos.NewUint(common.One*5), cosmos.NewUint(common.One*5), GetRandomETHAddress(), GetRandomETHAddress(), common.NoAddress, cosmos.ZeroUint(), GetRandomValidatorNode(NodeActive).NodeAddress),
			expectedResult: errAddLiquidityFailValidation,
		},
		{
			name:           "synth asset from coins should fail",
			msg:            NewMsgAddLiquidity(tx, common.ETHAsset, cosmos.NewUint(common.One*5), cosmos.NewUint(common.One*5), GetRandomETHAddress(), GetRandomETHAddress(), common.NoAddress, cosmos.ZeroUint(), GetRandomValidatorNode(NodeActive).NodeAddress),
			expectedResult: errAddLiquidityFailValidation,
		},
		{
			name:           "empty addresses should fail",
			msg:            NewMsgAddLiquidity(GetRandomTx(), common.BTCAsset, cosmos.NewUint(common.One*5), cosmos.NewUint(common.One*5), common.NoAddress, common.NoAddress, common.NoAddress, cosmos.ZeroUint(), GetRandomValidatorNode(NodeActive).NodeAddress),
			expectedResult: errAddLiquidityFailValidation,
		},
		{
			name:           "total liquidity provider is more than total bond should fail",
			msg:            NewMsgAddLiquidity(GetRandomTx(), common.ETHAsset, cosmos.NewUint(common.One*5000), cosmos.NewUint(common.One*5000), GetRandomRUNEAddress(), GetRandomETHAddress(), common.NoAddress, cosmos.ZeroUint(), activeNodeAccount.NodeAddress),
			expectedResult: errAddLiquidityRUNEMoreThanBond,
		},
		{
			name:           "rune address with wrong chain should fail",
			msg:            NewMsgAddLiquidity(GetRandomTx(), common.ETHAsset, cosmos.NewUint(common.One*5), cosmos.NewUint(common.One*5), GetRandomETHAddress(), GetRandomRUNEAddress(), common.NoAddress, cosmos.ZeroUint(), GetRandomValidatorNode(NodeActive).NodeAddress),
			expectedResult: errAddLiquidityFailValidation,
		},
		{
			name:           "asset address with wrong network should fail",
			msg:            NewMsgAddLiquidity(GetRandomTx(), common.ETHAsset, cosmos.NewUint(common.One*5), cosmos.NewUint(common.One*5), GetRandomRUNEAddress(), mainnetBTCAddr, common.NoAddress, cosmos.ZeroUint(), GetRandomValidatorNode(NodeActive).NodeAddress),
			expectedResult: fmt.Errorf("address(%s) is not same network", mainnetBTCAddr),
		},
	}
	constAccessor := constants.NewDummyConstants(map[constants.ConstantName]int64{
		constants.MaximumLiquidityRune: 600_000_00000000,
	}, map[constants.ConstantName]bool{
		constants.StrictBondLiquidityRatio: true,
	}, map[constants.ConstantName]string{})

	for _, item := range testCases {
		mgr := NewDummyMgrWithKeeper(k)
		mgr.constAccessor = constAccessor
		addHandler := NewAddLiquidityHandler(mgr)
		_, err := addHandler.Run(ctx, item.msg)
		c.Assert(err.Error(), Equals, item.expectedResult.Error(), Commentf("name:%s, actual: %w, expected: %w", item.name, err, item.expectedResult))
	}
}

func (s *HandlerAddLiquiditySuite) TestHandlerAddLiquidityFailScenario(c *C) {
	activeNodeAccount := GetRandomValidatorNode(NodeActive)
	emptyPool := Pool{
		BalanceRune:  cosmos.ZeroUint(),
		BalanceAsset: cosmos.ZeroUint(),
		Asset:        common.ETHAsset,
		LPUnits:      cosmos.ZeroUint(),
		Status:       PoolAvailable,
	}

	testCases := []struct {
		name           string
		k              keeper.Keeper
		expectedResult error
	}{
		{
			name: "fail to get pool should fail add liquidity",
			k: &MockAddLiquidityKeeper{
				activeNodeAccount: activeNodeAccount,
				currentPool:       emptyPool,
				failGetPool:       true,
			},
			expectedResult: errInternal,
		},
		{
			name: "suspended pool should fail add liquidity",
			k: &MockAddLiquidityKeeper{
				activeNodeAccount: activeNodeAccount,
				currentPool: Pool{
					BalanceRune:  cosmos.ZeroUint(),
					BalanceAsset: cosmos.ZeroUint(),
					Asset:        common.ETHAsset,
					LPUnits:      cosmos.ZeroUint(),
					Status:       PoolSuspended,
				},
			},
			expectedResult: errInvalidPoolStatus,
		},
	}
	for _, tc := range testCases {
		runeAddr := GetRandomRUNEAddress()
		ethAddr := GetRandomETHAddress()
		addTxHash := GetRandomTxHash()
		tx := common.NewTx(
			addTxHash,
			runeAddr,
			GetRandomETHAddress(),
			common.Coins{common.NewCoin(common.ETHAsset, cosmos.NewUint(common.One*5))},
			common.Gas{
				common.NewCoin(common.ETHAsset, cosmos.NewUint(10000)),
			},
			"add:ETH",
		)
		msg := NewMsgAddLiquidity(
			tx,
			common.ETHAsset,
			cosmos.NewUint(100*common.One),
			cosmos.NewUint(100*common.One),
			runeAddr,
			ethAddr,
			common.NoAddress, cosmos.ZeroUint(),
			activeNodeAccount.NodeAddress)
		ctx, mgr := setupManagerForTest(c)
		mgr.K = tc.k
		addHandler := NewAddLiquidityHandler(mgr)
		_, err := addHandler.Run(ctx, msg)
		c.Assert(errors.Is(err, tc.expectedResult), Equals, true, Commentf(tc.name))
	}
}

type AddLiquidityTestKeeper struct {
	keeper.KVStoreDummy
	store          map[string]interface{}
	liquidityUnits cosmos.Uint
}

// NewAddLiquidityTestKeeper
func NewAddLiquidityTestKeeper() *AddLiquidityTestKeeper {
	return &AddLiquidityTestKeeper{
		store:          make(map[string]interface{}),
		liquidityUnits: cosmos.ZeroUint(),
	}
}

func (p *AddLiquidityTestKeeper) PoolExist(ctx cosmos.Context, asset common.Asset) bool {
	_, ok := p.store[asset.String()]
	return ok
}

var notExistLiquidityProviderAsset, _ = common.NewAsset("ETH.NotExistLiquidityProviderAsset")

func (p *AddLiquidityTestKeeper) GetPool(ctx cosmos.Context, asset common.Asset) (Pool, error) {
	if p, ok := p.store[asset.String()]; ok {
		pool, ok := p.(Pool)
		if !ok {
			return pool, fmt.Errorf("dev error: failed to cast pool")
		}
		return pool, nil
	}
	return NewPool(), nil
}

func (p *AddLiquidityTestKeeper) SetPool(ctx cosmos.Context, ps Pool) error {
	p.store[ps.Asset.String()] = ps
	return nil
}

func (p *AddLiquidityTestKeeper) GetModuleAddress(_ string) (common.Address, error) {
	return common.NoAddress, nil
}

func (p *AddLiquidityTestKeeper) GetPOL(_ cosmos.Context) (ProtocolOwnedLiquidity, error) {
	return NewProtocolOwnedLiquidity(), nil
}

func (p *AddLiquidityTestKeeper) SetPOL(_ cosmos.Context, pol ProtocolOwnedLiquidity) error {
	return nil
}

func (p *AddLiquidityTestKeeper) GetLiquidityProvider(ctx cosmos.Context, asset common.Asset, addr common.Address) (LiquidityProvider, error) {
	if notExistLiquidityProviderAsset.Equals(asset) {
		return LiquidityProvider{}, errors.New("simulate error for test")
	}
	lp := LiquidityProvider{
		Asset:             asset,
		RuneAddress:       addr,
		Units:             cosmos.ZeroUint(),
		PendingRune:       cosmos.ZeroUint(),
		PendingAsset:      cosmos.ZeroUint(),
		RuneDepositValue:  cosmos.ZeroUint(),
		AssetDepositValue: cosmos.ZeroUint(),
	}
	key := p.GetKey(ctx, "lp/", lp.Key())
	if res, ok := p.store[key]; ok {
		lp, ok := res.(LiquidityProvider)
		if !ok {
			return lp, fmt.Errorf("dev error: failed to cast liquidity provider")
		}
		return lp, nil
	}
	lp.Units = p.liquidityUnits
	return lp, nil
}

func (p *AddLiquidityTestKeeper) SetLiquidityProvider(ctx cosmos.Context, lp LiquidityProvider) {
	key := p.GetKey(ctx, "lp/", lp.Key())
	p.store[key] = lp
}

func (p *AddLiquidityTestKeeper) AddOwnership(ctx cosmos.Context, coin common.Coin, addr cosmos.AccAddress) error {
	p.liquidityUnits = p.liquidityUnits.Add(coin.Amount)
	return nil
}

func (s *HandlerAddLiquiditySuite) TestCalculateLPUnits(c *C) {
	inputs := []struct {
		name           string
		oldLPUnits     cosmos.Uint
		poolRune       cosmos.Uint
		poolAsset      cosmos.Uint
		addRune        cosmos.Uint
		addAsset       cosmos.Uint
		poolUnits      cosmos.Uint
		liquidityUnits cosmos.Uint
		expectedErr    error
	}{
		{
			name:           "first-add-zero-rune",
			oldLPUnits:     cosmos.ZeroUint(),
			poolRune:       cosmos.ZeroUint(),
			poolAsset:      cosmos.ZeroUint(),
			addRune:        cosmos.ZeroUint(),
			addAsset:       cosmos.NewUint(100 * common.One),
			poolUnits:      cosmos.ZeroUint(),
			liquidityUnits: cosmos.ZeroUint(),
			expectedErr:    errors.New("total RUNE in the pool is zero"),
		},
		{
			name:           "first-add-zero-asset",
			oldLPUnits:     cosmos.ZeroUint(),
			poolRune:       cosmos.ZeroUint(),
			poolAsset:      cosmos.ZeroUint(),
			addRune:        cosmos.NewUint(100 * common.One),
			addAsset:       cosmos.ZeroUint(),
			poolUnits:      cosmos.ZeroUint(),
			liquidityUnits: cosmos.ZeroUint(),
			expectedErr:    errors.New("total asset in the pool is zero"),
		},
		{
			name:           "first-add",
			oldLPUnits:     cosmos.ZeroUint(),
			poolRune:       cosmos.ZeroUint(),
			poolAsset:      cosmos.ZeroUint(),
			addRune:        cosmos.NewUint(100 * common.One),
			addAsset:       cosmos.NewUint(100 * common.One),
			poolUnits:      cosmos.NewUint(100 * common.One),
			liquidityUnits: cosmos.NewUint(100 * common.One),
			expectedErr:    nil,
		},
		{
			name:           "second-add",
			oldLPUnits:     cosmos.NewUint(789500 * common.One),
			poolRune:       cosmos.NewUint(500 * common.One),
			poolAsset:      cosmos.NewUint(500 * common.One),
			addRune:        cosmos.NewUint(345 * common.One),
			addAsset:       cosmos.NewUint(234 * common.One),
			poolUnits:      cosmos.NewUint(1240460 * common.One),
			liquidityUnits: cosmos.NewUint(450960 * common.One),
			expectedErr:    nil,
		},
		{
			name:           "asym-add",
			oldLPUnits:     cosmos.NewUint(300 * common.One),
			poolRune:       cosmos.NewUint(500 * common.One),
			poolAsset:      cosmos.NewUint(500 * common.One),
			addRune:        cosmos.NewUint(500 * common.One),
			addAsset:       cosmos.ZeroUint(),
			poolUnits:      cosmos.NewUint(400 * common.One),
			liquidityUnits: cosmos.NewUint(100 * common.One),
			expectedErr:    nil,
		},
	}

	for _, item := range inputs {
		c.Logf("Name: %s", item.name)
		poolUnits, liquidityUnits, err := calculatePoolUnits(item.oldLPUnits, item.poolRune, item.poolAsset, item.addRune, item.addAsset)
		if item.expectedErr == nil {
			c.Assert(err, IsNil)
		} else {
			c.Assert(err.Error(), Equals, item.expectedErr.Error())
		}

		c.Check(item.poolUnits.Uint64(), Equals, poolUnits.Uint64(), Commentf("%d / %d", item.poolUnits.Uint64(), poolUnits.Uint64()))
		c.Check(item.liquidityUnits.Uint64(), Equals, liquidityUnits.Uint64(), Commentf("%d / %d", item.liquidityUnits.Uint64(), liquidityUnits.Uint64()))
	}
}

func (s *HandlerAddLiquiditySuite) TestValidateAddLiquidityMessage(c *C) {
	ps := NewAddLiquidityTestKeeper()
	ctx, mgr := setupManagerForTest(c)
	mgr.K = ps
	txID := GetRandomTxHash()
	ethAddress := GetRandomETHAddress()
	assetAddress := GetRandomETHAddress()
	h := NewAddLiquidityHandler(mgr)
	c.Assert(h.validateAddLiquidityMessage(ctx, ps, common.Asset{}, txID, ethAddress, assetAddress), NotNil)
	c.Assert(h.validateAddLiquidityMessage(ctx, ps, common.ETHAsset, txID, ethAddress, assetAddress), NotNil)
	c.Assert(h.validateAddLiquidityMessage(ctx, ps, common.ETHAsset, txID, ethAddress, assetAddress), NotNil)
	c.Assert(h.validateAddLiquidityMessage(ctx, ps, common.ETHAsset, common.TxID(""), ethAddress, assetAddress), NotNil)
	c.Assert(h.validateAddLiquidityMessage(ctx, ps, common.ETHAsset, txID, common.NoAddress, common.NoAddress), NotNil)
	c.Assert(h.validateAddLiquidityMessage(ctx, ps, common.ETHAsset, txID, ethAddress, assetAddress), NotNil)
	c.Assert(h.validateAddLiquidityMessage(ctx, ps, common.ETHAsset, txID, common.NoAddress, assetAddress), NotNil)
	c.Assert(h.validateAddLiquidityMessage(ctx, ps, common.BTCAsset, txID, ethAddress, common.NoAddress), NotNil)
	c.Assert(ps.SetPool(ctx, Pool{
		BalanceRune:  cosmos.NewUint(100 * common.One),
		BalanceAsset: cosmos.NewUint(100 * common.One),
		Asset:        common.ETHAsset,
		LPUnits:      cosmos.NewUint(100 * common.One),
		Status:       PoolAvailable,
	}), IsNil)
	// happy path
	c.Assert(h.validateAddLiquidityMessage(ctx, ps, common.ETHAsset, txID, ethAddress, assetAddress), Equals, nil)
	// Don't accept THOR.ETH
	thorAsset := common.ETHAsset
	var err error // Previously undeclared, so declaration needed
	thorAsset.Chain, err = common.NewChain("THOR")
	c.Assert(err, IsNil)
	c.Assert(h.validateAddLiquidityMessage(ctx, ps, thorAsset, txID, ethAddress, assetAddress), NotNil)
	// Don't accept TEST.ETH
	testAsset := common.ETHAsset
	testAsset.Chain, err = common.NewChain("TEST")
	c.Assert(err, IsNil)
	c.Assert(h.validateAddLiquidityMessage(ctx, ps, testAsset, txID, ethAddress, assetAddress), NotNil)
}

func (s *HandlerAddLiquiditySuite) TestAddLiquidityV1(c *C) {
	ps := NewAddLiquidityTestKeeper()
	ctx, _ := setupKeeperForTest(c)
	txID := GetRandomTxHash()

	runeAddress := GetRandomRUNEAddress()
	assetAddress := GetRandomETHAddress()
	btcAddress, err := common.NewAddress("bc1qwqdg6squsna38e46795at95yu9atm8azzmyvckulcc7kytlcckxswvvzej")
	c.Assert(err, IsNil)
	constAccessor := constants.GetConstantValues(GetCurrentVersion())
	h := NewAddLiquidityHandler(NewDummyMgrWithKeeper(ps))
	err = h.addLiquidity(ctx, common.Asset{}, cosmos.NewUint(100*common.One), cosmos.NewUint(100*common.One), runeAddress, assetAddress, txID, false, constAccessor)
	c.Assert(err, NotNil)
	c.Assert(ps.SetPool(ctx, Pool{
		BalanceRune:         cosmos.ZeroUint(),
		BalanceAsset:        cosmos.NewUint(100 * common.One),
		Asset:               common.ETHAsset,
		LPUnits:             cosmos.NewUint(100 * common.One),
		SynthUnits:          cosmos.ZeroUint(),
		PendingInboundAsset: cosmos.ZeroUint(),
		PendingInboundRune:  cosmos.ZeroUint(),
		Status:              PoolAvailable,
	}), IsNil)
	err = h.addLiquidity(ctx, common.ETHAsset, cosmos.NewUint(100*common.One), cosmos.NewUint(100*common.One), runeAddress, assetAddress, txID, false, constAccessor)
	c.Assert(err, IsNil)
	su, err := ps.GetLiquidityProvider(ctx, common.ETHAsset, runeAddress)
	c.Assert(err, IsNil)
	// c.Assert(su.Units.Equal(cosmos.NewUint(11250000000)), Equals, true, Commentf("%d", su.Units.Uint64()))

	c.Assert(ps.SetPool(ctx, Pool{
		BalanceRune:         cosmos.NewUint(100 * common.One),
		BalanceAsset:        cosmos.NewUint(100 * common.One),
		Asset:               notExistLiquidityProviderAsset,
		LPUnits:             cosmos.NewUint(100 * common.One),
		SynthUnits:          cosmos.ZeroUint(),
		PendingInboundAsset: cosmos.ZeroUint(),
		PendingInboundRune:  cosmos.ZeroUint(),
		Status:              PoolAvailable,
	}), IsNil)
	// add asymmetically
	err = h.addLiquidity(ctx, common.ETHAsset, cosmos.NewUint(100*common.One), cosmos.ZeroUint(), runeAddress, assetAddress, txID, false, constAccessor)
	c.Assert(err, IsNil)
	err = h.addLiquidity(ctx, common.ETHAsset, cosmos.ZeroUint(), cosmos.NewUint(100*common.One), runeAddress, assetAddress, txID, false, constAccessor)
	c.Assert(err, IsNil)

	err = h.addLiquidity(ctx, notExistLiquidityProviderAsset, cosmos.NewUint(100*common.One), cosmos.NewUint(100*common.One), runeAddress, assetAddress, txID, false, constAccessor)
	c.Assert(err, NotNil)
	c.Assert(ps.SetPool(ctx, Pool{
		BalanceRune:         cosmos.NewUint(100 * common.One),
		BalanceAsset:        cosmos.NewUint(100 * common.One),
		Asset:               common.ETHAsset,
		LPUnits:             cosmos.NewUint(100 * common.One),
		SynthUnits:          cosmos.ZeroUint(),
		PendingInboundAsset: cosmos.ZeroUint(),
		PendingInboundRune:  cosmos.ZeroUint(),
		Status:              PoolAvailable,
	}), IsNil)

	for i := 1; i <= 150; i++ {
		lp := LiquidityProvider{Units: cosmos.NewUint(common.One / 5000)}
		ps.SetLiquidityProvider(ctx, lp)
	}
	err = h.addLiquidity(ctx, common.ETHAsset, cosmos.NewUint(common.One), cosmos.NewUint(common.One), runeAddress, assetAddress, txID, false, constAccessor)
	c.Assert(err, IsNil)

	err = h.addLiquidity(ctx, common.ETHAsset, cosmos.NewUint(100*common.One), cosmos.NewUint(100*common.One), runeAddress, assetAddress, txID, false, constAccessor)
	c.Assert(err, IsNil)
	p, err := ps.GetPool(ctx, common.ETHAsset)
	c.Assert(err, IsNil)
	c.Check(p.LPUnits.Equal(cosmos.NewUint(201*common.One)), Equals, true, Commentf("%d", p.LPUnits.Uint64()))

	// Test atomic cross chain liquidity provision
	// create BTC pool
	c.Assert(ps.SetPool(ctx, Pool{
		BalanceRune:         cosmos.ZeroUint(),
		BalanceAsset:        cosmos.ZeroUint(),
		Asset:               common.BTCAsset,
		LPUnits:             cosmos.ZeroUint(),
		SynthUnits:          cosmos.ZeroUint(),
		PendingInboundAsset: cosmos.ZeroUint(),
		PendingInboundRune:  cosmos.ZeroUint(),
		Status:              PoolAvailable,
	}), IsNil)

	// add rune
	err = h.addLiquidity(ctx, common.BTCAsset, cosmos.NewUint(100*common.One), cosmos.ZeroUint(), runeAddress, btcAddress, txID, true, constAccessor)
	c.Assert(err, IsNil)
	_, err = ps.GetLiquidityProvider(ctx, common.BTCAsset, runeAddress)
	c.Assert(err, IsNil)
	// c.Check(su.Units.IsZero(), Equals, true)
	// add btc
	err = h.addLiquidity(ctx, common.BTCAsset, cosmos.ZeroUint(), cosmos.NewUint(100*common.One), runeAddress, btcAddress, txID, false, constAccessor)
	c.Assert(err, IsNil)
	su, err = ps.GetLiquidityProvider(ctx, common.BTCAsset, runeAddress)
	c.Assert(err, IsNil)
	c.Check(su.Units.IsZero(), Equals, false)
	p, err = ps.GetPool(ctx, common.BTCAsset)
	c.Assert(err, IsNil)
	c.Check(p.BalanceAsset.Equal(cosmos.NewUint(100*common.One)), Equals, true, Commentf("%d", p.BalanceAsset.Uint64()))
	c.Check(p.BalanceRune.Equal(cosmos.NewUint(100*common.One)), Equals, true, Commentf("%d", p.BalanceRune.Uint64()))
	c.Check(p.LPUnits.Equal(cosmos.NewUint(100*common.One)), Equals, true, Commentf("%d", p.LPUnits.Uint64()))
}

func (HandlerAddLiquiditySuite) TestRuneOnlyFairMergeProvidedLiquidity(c *C) {
	ctx, k := setupKeeperForTest(c)
	txID := GetRandomTxHash()

	c.Assert(k.SetPool(ctx, Pool{
		BalanceRune:  cosmos.NewUint(100 * common.One),
		BalanceAsset: cosmos.NewUint(100 * common.One),
		Asset:        common.BTCAsset,
		LPUnits:      cosmos.NewUint(100 * common.One),
		SynthUnits:   cosmos.ZeroUint(),
		Status:       PoolAvailable,
	}), IsNil)

	runeAddr := GetRandomRUNEAddress()
	constAccessor := constants.GetConstantValues(GetCurrentVersion())
	h := NewAddLiquidityHandler(NewDummyMgrWithKeeper(k))
	err := h.addLiquidity(ctx, common.BTCAsset, cosmos.NewUint(100*common.One), cosmos.ZeroUint(), runeAddr, common.NoAddress, txID, false, constAccessor)
	c.Assert(err, IsNil)

	su, err := k.GetLiquidityProvider(ctx, common.BTCAsset, runeAddr)
	c.Assert(err, IsNil)
	c.Assert(su.Units.Uint64(), Equals, uint64(3333333333), Commentf("%d", su.Units.Uint64()))

	pool, err := k.GetPool(ctx, common.BTCAsset)
	c.Assert(err, IsNil)
	c.Assert(pool.LPUnits.Uint64(), Equals, uint64(13333333333), Commentf("%d", pool.LPUnits.Uint64()))
}

func (HandlerAddLiquiditySuite) TestAssetOnlyFairMergeProvidedLiquidity(c *C) {
	ctx, k := setupKeeperForTest(c)
	txID := GetRandomTxHash()

	c.Assert(k.SetPool(ctx, Pool{
		BalanceRune:  cosmos.NewUint(100 * common.One),
		BalanceAsset: cosmos.NewUint(100 * common.One),
		Asset:        common.BTCAsset,
		LPUnits:      cosmos.NewUint(100 * common.One),
		SynthUnits:   cosmos.ZeroUint(),
		Status:       PoolAvailable,
	}), IsNil)

	assetAddr := GetRandomBTCAddress()
	constAccessor := constants.GetConstantValues(GetCurrentVersion())
	h := NewAddLiquidityHandler(NewDummyMgrWithKeeper(k))
	err := h.addLiquidity(ctx, common.BTCAsset, cosmos.ZeroUint(), cosmos.NewUint(100*common.One), common.NoAddress, assetAddr, txID, false, constAccessor)
	c.Assert(err, IsNil)

	su, err := k.GetLiquidityProvider(ctx, common.BTCAsset, assetAddr)
	c.Assert(err, IsNil)
	c.Assert(su.Units.Uint64(), Equals, uint64(3333333333), Commentf("%d", su.Units.Uint64()))

	pool, err := k.GetPool(ctx, common.BTCAsset)
	c.Assert(err, IsNil)
	c.Assert(pool.LPUnits.Uint64(), Equals, uint64(13333333333), Commentf("%d", pool.LPUnits.Uint64()))
}

func (HandlerAddLiquiditySuite) TestSynthValidate(c *C) {
	ctx, mgr := setupManagerForTest(c)

	asset := common.BTCAsset.GetSyntheticAsset()

	c.Assert(mgr.Keeper().SetPool(ctx, Pool{
		BalanceRune:  cosmos.NewUint(100 * common.One),
		BalanceAsset: cosmos.NewUint(10 * common.One),
		Asset:        asset,
		LPUnits:      cosmos.ZeroUint(),
		SynthUnits:   cosmos.ZeroUint(),
		Status:       PoolAvailable,
	}), IsNil)

	handler := NewAddLiquidityHandler(mgr)

	addr := GetRandomBTCAddress()
	signer := GetRandomBech32Addr()
	addTxHash := GetRandomTxHash()

	tx := common.NewTx(
		addTxHash,
		addr,
		addr,
		common.Coins{common.NewCoin(asset, cosmos.NewUint(1000*common.One))},
		common.Gas{
			common.NewCoin(common.ETHAsset, cosmos.NewUint(10000)),
		},
		fmt.Sprintf("add:%s", asset.String()),
	)

	// don't allow add liquidity when the gas asset pool doesn't exist
	msg := NewMsgAddLiquidity(tx, asset, cosmos.ZeroUint(), cosmos.NewUint(1000*common.One), common.NoAddress, addr, common.NoAddress, cosmos.ZeroUint(), signer)
	err := handler.validate(ctx, *msg)
	c.Assert(err, NotNil)

	// Set gas pool's Asset to represent existence for IsEmpty
	gasPool := NewPool()
	gasPool.Asset = common.BTCAsset
	c.Assert(mgr.Keeper().SetPool(ctx, gasPool), IsNil)

	// happy path
	err = handler.validate(ctx, *msg)
	c.Assert(err, IsNil)

	// don't accept THOR/BTC
	thorAsset := asset
	thorAsset.Chain, err = common.NewChain("THOR")
	c.Assert(err, IsNil)
	msg.Asset = thorAsset
	err = handler.validate(ctx, *msg)
	c.Assert(err, NotNil)
	// don't accept TEST/BTC
	testAsset := asset
	testAsset.Chain, err = common.NewChain("TEST")
	c.Assert(err, IsNil)
	msg.Asset = testAsset
	err = handler.validate(ctx, *msg)
	c.Assert(err, NotNil)

	// don't allow non-gas assets
	busd, err := common.NewAsset("ETH.BUSD-BD1")
	c.Assert(err, IsNil)
	msg = NewMsgAddLiquidity(tx, busd.GetSyntheticAsset(), cosmos.ZeroUint(), cosmos.NewUint(1000*common.One), addr, common.NoAddress, common.NoAddress, cosmos.ZeroUint(), signer)
	err = handler.validate(ctx, *msg)
	c.Assert(err, NotNil)

	// address mismatch
	msg = NewMsgAddLiquidity(tx, asset, cosmos.ZeroUint(), cosmos.NewUint(1000*common.One), addr, common.NoAddress, common.NoAddress, cosmos.ZeroUint(), signer)
	err = handler.validate(ctx, *msg)
	c.Assert(err, NotNil)
	msg = NewMsgAddLiquidity(tx, asset, cosmos.ZeroUint(), cosmos.NewUint(1000*common.One), common.NoAddress, common.NoAddress, common.NoAddress, cosmos.ZeroUint(), signer)
	err = handler.validate(ctx, *msg)
	c.Assert(err, NotNil)

	// don't allow rune
	msg = NewMsgAddLiquidity(tx, asset, cosmos.NewUint(1000*common.One), cosmos.ZeroUint(), common.NoAddress, addr, common.NoAddress, cosmos.ZeroUint(), signer)
	err = handler.validate(ctx, *msg)
	c.Assert(err, NotNil)
	msg = NewMsgAddLiquidity(tx, asset, cosmos.NewUint(1000*common.One), cosmos.NewUint(1000*common.One), common.NoAddress, addr, common.NoAddress, cosmos.ZeroUint(), signer)
	err = handler.validate(ctx, *msg)
	c.Assert(err, NotNil)
}

func (HandlerAddLiquiditySuite) TestAddSynthNoLPs(c *C) {
	// there is an odd case where its possible in a synth vault to have a
	// balance asset of non-zero BUT have no LPs yet. Testing this edge case.
	ctx, k := setupKeeperForTest(c)
	txID := GetRandomTxHash()

	asset := common.BTCAsset.GetSyntheticAsset()

	pool := NewPool()
	pool.Asset = asset
	pool.Status = PoolAvailable
	pool.BalanceRune = cosmos.NewUint(0)
	pool.BalanceAsset = cosmos.NewUint(10 * common.One)
	c.Assert(k.SetPool(ctx, pool), IsNil)

	coin := common.NewCoin(asset, pool.BalanceAsset)
	c.Assert(k.MintToModule(ctx, ModuleName, coin), IsNil)
	c.Assert(k.SendFromModuleToModule(ctx, ModuleName, AsgardName, common.NewCoins(coin)), IsNil)

	addr := GetRandomBTCAddress()
	constAccessor := constants.GetConstantValues(GetCurrentVersion())
	h := NewAddLiquidityHandler(NewDummyMgrWithKeeper(k))
	addCoin := common.NewCoin(asset, cosmos.NewUint(10*common.One))
	c.Assert(k.MintToModule(ctx, ModuleName, addCoin), IsNil)
	c.Assert(k.SendFromModuleToModule(ctx, ModuleName, AsgardName, common.NewCoins(addCoin)), IsNil)
	err := h.addLiquidity(ctx, asset, cosmos.ZeroUint(), addCoin.Amount, common.NoAddress, addr, txID, false, constAccessor)
	c.Assert(err, IsNil)

	su, err := k.GetLiquidityProvider(ctx, asset, addr)
	c.Assert(err, IsNil)
	c.Check(su.Units.Uint64(), Equals, uint64(10*common.One), Commentf("%d", su.Units.Uint64()))

	pool, err = k.GetPool(ctx, asset)
	c.Assert(err, IsNil)
	c.Check(pool.BalanceRune.Uint64(), Equals, uint64(0), Commentf("%d", pool.BalanceRune.Uint64()))
	c.Check(pool.BalanceAsset.Uint64(), Equals, uint64(20*common.One), Commentf("%d", pool.BalanceAsset.Uint64()))
	c.Check(pool.LPUnits.Uint64(), Equals, uint64(10*common.One), Commentf("%d", pool.LPUnits.Uint64()))
}

func (HandlerAddLiquiditySuite) TestAddSynth(c *C) {
	ctx, k := setupKeeperForTest(c)
	txID := GetRandomTxHash()

	asset := common.BTCAsset.GetSyntheticAsset()

	pool := NewPool()
	pool.Asset = asset
	pool.Status = PoolAvailable
	pool.BalanceRune = cosmos.NewUint(0)
	pool.BalanceAsset = cosmos.NewUint(100 * common.One)
	pool.LPUnits = cosmos.NewUint(100)
	c.Assert(k.SetPool(ctx, pool), IsNil)

	coin := common.NewCoin(asset, pool.BalanceAsset)
	c.Assert(k.MintToModule(ctx, ModuleName, coin), IsNil)
	c.Assert(k.SendFromModuleToModule(ctx, ModuleName, AsgardName, common.NewCoins(coin)), IsNil)

	addr := GetRandomBTCAddress()
	constAccessor := constants.GetConstantValues(GetCurrentVersion())
	h := NewAddLiquidityHandler(NewDummyMgrWithKeeper(k))
	addCoin := common.NewCoin(asset, cosmos.NewUint(100*common.One))
	c.Assert(k.MintToModule(ctx, ModuleName, addCoin), IsNil)
	c.Assert(k.SendFromModuleToModule(ctx, ModuleName, AsgardName, common.NewCoins(addCoin)), IsNil)
	err := h.addLiquidity(ctx, asset, cosmos.ZeroUint(), addCoin.Amount, common.NoAddress, addr, txID, false, constAccessor)
	c.Assert(err, IsNil)

	su, err := k.GetLiquidityProvider(ctx, asset, addr)
	c.Assert(err, IsNil)
	c.Check(su.Units.Uint64(), Equals, uint64(100), Commentf("%d", su.Units.Uint64()))

	pool, err = k.GetPool(ctx, asset)
	c.Assert(err, IsNil)
	c.Check(pool.BalanceRune.Uint64(), Equals, uint64(0), Commentf("%d", pool.BalanceRune.Uint64()))
	c.Check(pool.BalanceAsset.Uint64(), Equals, uint64(200*common.One), Commentf("%d", pool.BalanceAsset.Uint64()))
	c.Check(pool.LPUnits.Uint64(), Equals, uint64(200), Commentf("%d", pool.LPUnits.Uint64()))
}

func (s *HandlerAddLiquiditySuite) TestAddLiquidityPOL(c *C) {
	var err error
	ctx, mgr := setupManagerForTest(c)
	activeNodeAccount := GetRandomValidatorNode(NodeActive)
	polAddr := GetRandomRUNEAddress()
	pool := NewPool()
	pool.Asset = common.ETHAsset
	pool.Status = PoolAvailable
	pool.BalanceAsset = cosmos.NewUint(1)
	pool.BalanceRune = cosmos.NewUint(1)
	k := &MockAddLiquidityKeeper{
		activeNodeAccount: activeNodeAccount,
		currentPool:       pool,
		lp: LiquidityProvider{
			Asset:             common.ETHAsset,
			RuneAddress:       polAddr,
			AssetAddress:      common.NoAddress,
			Units:             cosmos.ZeroUint(),
			PendingRune:       cosmos.ZeroUint(),
			PendingAsset:      cosmos.ZeroUint(),
			RuneDepositValue:  cosmos.ZeroUint(),
			AssetDepositValue: cosmos.ZeroUint(),
		},
		pol:        NewProtocolOwnedLiquidity(),
		polAddress: polAddr,
	}
	mgr.K = k
	addHandler := NewAddLiquidityHandler(mgr)
	addTxHash := GetRandomTxHash()
	tx := common.NewTx(
		addTxHash,
		polAddr,
		polAddr,
		common.Coins{common.NewCoin(common.RuneAsset(), cosmos.NewUint(common.One*100))},
		common.Gas{},
		"add:ETH",
	)
	msg := NewMsgAddLiquidity(
		tx,
		common.ETHAsset,
		cosmos.NewUint(100*common.One),
		cosmos.ZeroUint(),
		polAddr,
		GetRandomETHAddress(),
		common.NoAddress, cosmos.ZeroUint(),
		activeNodeAccount.NodeAddress)

	_, err = addHandler.Run(ctx, msg)
	c.Assert(err, NotNil, Commentf("pol add with asset addr should fail"))

	// happy path
	msg.AssetAddress = common.NoAddress
	_, err = addHandler.Run(ctx, msg)
	c.Assert(err, IsNil)

	postLiquidityPool, err := mgr.Keeper().GetPool(ctx, common.ETHAsset)
	c.Assert(err, IsNil)
	c.Assert(postLiquidityPool.BalanceAsset.String(), Equals, "1")
	c.Assert(postLiquidityPool.BalanceRune.String(), Equals, "10000000001")
	c.Assert(postLiquidityPool.PendingInboundAsset.String(), Equals, "0")
	c.Assert(postLiquidityPool.PendingInboundRune.String(), Equals, "0")

	pol, err := mgr.Keeper().GetPOL(ctx)
	c.Assert(err, IsNil)
	c.Check(pol.RuneDeposited.Uint64(), Equals, uint64(10000000000))
}
