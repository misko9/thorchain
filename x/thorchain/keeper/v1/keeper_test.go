package keeperv1

import (
	"testing"

	. "gopkg.in/check.v1"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	"github.com/cosmos/cosmos-sdk/runtime"
	storemetrics "cosmossdk.io/store/metrics"
	"github.com/cosmos/cosmos-sdk/x/auth"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authcodec "github.com/cosmos/cosmos-sdk/x/auth/codec"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/types/module/testutil"
	dbm "github.com/cosmos/cosmos-db"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"gitlab.com/thorchain/thornode/common"
	"gitlab.com/thorchain/thornode/common/cosmos"
)

func TestPackage(t *testing.T) { TestingT(t) }

func FundModule(c *C, ctx cosmos.Context, k KVStore, name string, amt uint64) {
	coin := common.NewCoin(common.RuneNative, cosmos.NewUint(amt*common.One))
	err := k.MintToModule(ctx, ModuleName, coin)
	c.Assert(err, IsNil)
	err = k.SendFromModuleToModule(ctx, ModuleName, name, common.NewCoins(coin))
	c.Assert(err, IsNil)
}

var keyThorchain = cosmos.NewKVStoreKey(StoreKey)

func setupKeeperForTest(c *C) (cosmos.Context, KVStore) {
	SetupConfigForTest()
	keys := cosmos.NewKVStoreKeys(
		authtypes.StoreKey, banktypes.StoreKey, stakingtypes.StoreKey,
	)

	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db, log.NewNopLogger(), storemetrics.NewNoOpMetrics())
	ms.MountStoreWithDB(keys[authtypes.StoreKey], cosmos.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keys[banktypes.StoreKey], cosmos.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyThorchain, cosmos.StoreTypeIAVL, db)
	err := ms.LoadLatestVersion()
	c.Assert(err, IsNil)

	ctx := cosmos.NewContext(ms, tmproto.Header{ChainID: "thorchain"}, false, log.NewNopLogger())
	ctx = ctx.WithBlockHeight(18)

	encodingConfig := testutil.MakeTestEncodingConfig(
		bank.AppModuleBasic{},
		auth.AppModuleBasic{},
	)
	
	maccPerms := map[string][]string{
		authtypes.FeeCollectorName:     nil,
		distrtypes.ModuleName:          nil,
		minttypes.ModuleName:           {authtypes.Minter},
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
		ModuleName:                     {authtypes.Minter, authtypes.Burner},
		ReserveName:                    {},
		AsgardName:                     {},
		TreasuryName:                   {},
		RUNEPoolName:                   {},
		BondName:                       {authtypes.Staking},
	}
	ak := authkeeper.NewAccountKeeper(
		encodingConfig.Codec,
		runtime.NewKVStoreService(keys[authtypes.StoreKey]),
		authtypes.ProtoBaseAccount,
		maccPerms,
		authcodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix()),
		sdk.GetConfig().GetBech32AccountAddrPrefix(),
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	bk := bankkeeper.NewBaseKeeper(
		encodingConfig.Codec,
		runtime.NewKVStoreService(keys[banktypes.StoreKey]),
		ak,
		nil,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		log.NewNopLogger(),
	)

	k := NewKVStore(encodingConfig.Codec, bk, ak, keyThorchain, GetCurrentVersion())

	FundModule(c, ctx, k, AsgardName, common.One)

	return ctx, k
}

type KeeperTestSuit struct{}

var _ = Suite(&KeeperTestSuit{})

func (KeeperTestSuit) TestKeeperVersion(c *C) {
	ctx, k := setupKeeperForTest(c)
	c.Check(k.GetStoreVersion(ctx), Equals, int64(38))

	k.SetStoreVersion(ctx, 2)
	c.Check(k.GetStoreVersion(ctx), Equals, int64(2))

	c.Check(k.GetRuneBalanceOfModule(ctx, AsgardName).Equal(cosmos.NewUint(100000000*common.One)), Equals, true)
	coinsToSend := common.NewCoins(common.NewCoin(common.RuneNative, cosmos.NewUint(1*common.One)))
	c.Check(k.SendFromModuleToModule(ctx, AsgardName, BondName, coinsToSend), IsNil)

	acct := GetRandomBech32Addr()
	c.Check(k.SendFromModuleToAccount(ctx, AsgardName, acct, coinsToSend), IsNil)

	// check get account balance
	coins := k.GetBalance(ctx, acct)
	c.Check(coins, HasLen, 1)

	c.Check(k.SendFromAccountToModule(ctx, acct, AsgardName, coinsToSend), IsNil)

	// check no account balance
	coins = k.GetBalance(ctx, GetRandomBech32Addr())
	c.Check(coins, HasLen, 0)
}

func (KeeperTestSuit) TestMaxMint(c *C) {
	ctx, k := setupKeeperForTest(c)

	max := int64(200000000_00000000)
	k.SetMimir(ctx, "MaxRuneSupply", max)
	maxCoin := common.NewCoin(common.RuneAsset(), cosmos.NewUint(uint64(max)))

	// ship asgard rune to reserve
	c.Assert(k.SendFromModuleToModule(ctx, AsgardName, ReserveName, common.NewCoins(common.NewCoin(common.RuneAsset(), cosmos.NewUint(10000000000000000)))), IsNil)
	// mint more rune into reserve to max the supply
	mintAmt := common.NewCoin(common.RuneAsset(), cosmos.NewUint(uint64(max)-10000000000000000))
	c.Assert(k.MintToModule(ctx, ModuleName, mintAmt), IsNil)
	c.Assert(k.SendFromModuleToModule(ctx, ModuleName, ReserveName, common.NewCoins(mintAmt)), IsNil)

	// mint more rune into another module
	moreCoin := common.NewCoin(common.RuneAsset(), cosmos.NewUint(uint64(max/4)))
	c.Assert(k.MintToModule(ctx, ModuleName, moreCoin), IsNil)

	// fetch module balances
	reserve := k.GetRuneBalanceOfModule(ctx, ReserveName)
	mod := k.GetRuneBalanceOfModule(ctx, ModuleName)

	// check reserve has been reduced
	c.Check(maxCoin.Amount.Sub(moreCoin.Amount).Uint64(), Equals, reserve.Uint64())
	// check total is not surpassed the max supply
	c.Check(reserve.Add(mod).Uint64(), Equals, uint64(max))
}
