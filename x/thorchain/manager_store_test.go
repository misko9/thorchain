package thorchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cometbft/cometbft/crypto/tmhash"
	. "gopkg.in/check.v1"

	"gitlab.com/thorchain/thornode/common"
	"gitlab.com/thorchain/thornode/common/cosmos"
)

type StoreManagerTestSuite struct{}

var _ = Suite(&StoreManagerTestSuite{})

func (s *StoreManagerTestSuite) TestFakeTxHash(c *C) {
	recoveryTx := TxOutItem{
		Chain:       common.BSCChain,
		VaultPubKey: common.PubKey("thorpub1addwnpepq0xtamtm6l35efh3f5wlt5fql7cx9j94fywtumz83vvnzagx46h76yk8sa3"),
		ToAddress:   common.Address("bnb1pa6hpjs7qv0vkd5ks5tqa2xtt2gk5n08yw7v7f"),
		Coin:        common.NewCoin(common.BNBBEP20Asset, cosmos.NewUint(100000000)),
		Memo:        "",
		InHash:      "",
		GasRate:     int64(10),
		MaxGas:      common.Gas{common.NewCoin(common.BNBBEP20Asset, cosmos.NewUint(50000))},
	}

	txBytes, err := json.Marshal(recoveryTx)
	c.Assert(err, IsNil)

	hash := strings.ToUpper(hex.EncodeToString(tmhash.Sum(txBytes)))
	c.Assert(hash, Equals, "DC44C3AFAB4D7B9A078726179D58FC7BFCFF8440C1354E72A99D396A1D01EF7D")

	// Change one thing, make sure it produces a different hash
	recoveryTx.GasRate = int64(11)
	txBytes, err = json.Marshal(recoveryTx)
	c.Assert(err, IsNil)

	hash = strings.ToUpper(hex.EncodeToString(tmhash.Sum(txBytes)))
	c.Assert(hash, Equals, "14C10F837126B1E8A796AFCEC377CC6764DD58D7CBE82786E1FFFCB951489E38")
}

func (s *StoreManagerTestSuite) TestRemoveTransactions(c *C) {
	ctx, mgr := setupManagerForTest(c)
	storeMgr := newStoreMgr(mgr)
	vault := NewVault(1024, ActiveVault, AsgardVault, GetRandomPubKey(), []string{
		common.ETHChain.String(),
	}, nil)

	c.Assert(storeMgr.mgr.Keeper().SaveNetworkFee(ctx, common.ETHChain, NetworkFee{
		Chain: common.ETHChain, TransactionSize: 80000, TransactionFeeRate: 30,
	}), IsNil)

	inTxID, err := common.NewTxID("BC68035CE2C8A2C549604FF7DB59E07931F39040758B138190338FA697338DB3")
	c.Assert(err, IsNil)
	tx := common.NewTx(inTxID,
		"0x3a196410a0f5facd08fd7880a4b8551cd085c031",
		"0xf56cBa49337A624E94042e325Ad6Bc864436E370",
		common.NewCoins(common.NewCoin(common.ETHAsset, cosmos.NewUint(200*common.One))),
		common.Gas{
			common.NewCoin(common.RuneNative, cosmos.NewUint(2000000)),
		}, "SWAP:ETH.AAVE-0x7Fc66500c84A76Ad7e9c93437bFc5Ac33E2DDaE9")
	observedTx := NewObservedTx(tx, 1281323, vault.PubKey, 1281323)
	voter := NewObservedTxVoter(inTxID, []ObservedTx{
		observedTx,
	})
	aaveAsset, _ := common.NewAsset("ETH.AAVE-0x7Fc66500c84A76Ad7e9c93437bFc5Ac33E2DDaE9")
	voter.Actions = []TxOutItem{
		{
			Chain:       common.ETHChain,
			ToAddress:   "0x3a196410a0f5facd08fd7880a4b8551cd085c031",
			VaultPubKey: vault.PubKey,
			Coin:        common.NewCoin(aaveAsset, cosmos.NewUint(1422136902)),
			Memo:        "OUT:BC68035CE2C8A2C549604FF7DB59E07931F39040758B138190338FA697338DB3",
			InHash:      "BC68035CE2C8A2C549604FF7DB59E07931F39040758B138190338FA697338DB3",
		},
		{
			Chain:       common.ETHChain,
			ToAddress:   "0x3a196410a0f5facd08fd7880a4b8551cd085c031",
			VaultPubKey: vault.PubKey,
			Coin:        common.NewCoin(aaveAsset, cosmos.NewUint(1330195098)),
			Memo:        "OUT:BC68035CE2C8A2C549604FF7DB59E07931F39040758B138190338FA697338DB3",
			InHash:      "BC68035CE2C8A2C549604FF7DB59E07931F39040758B138190338FA697338DB3",
		},
	}
	voter.Tx = voter.Txs[0]
	storeMgr.mgr.Keeper().SetObservedTxInVoter(ctx, voter)
	hegicAsset, _ := common.NewAsset("ETH.HEGIC-0x584bC13c7D411c00c01A62e8019472dE68768430")
	inTxID1, err := common.NewTxID("5DA19C1C5C2F6BBDF9D4FB0E6FF16A0DF6D6D7FE1F8E95CA755E5B3C6AADA458")
	c.Assert(err, IsNil)
	tx1 := common.NewTx(inTxID1,
		"0x3a196410a0f5facd08fd7880a4b8551cd085c031",
		"0xf56cBa49337A624E94042e325Ad6Bc864436E370",
		common.NewCoins(common.NewCoin(common.ETHAsset, cosmos.NewUint(200*common.One))),
		common.Gas{
			common.NewCoin(common.RuneNative, cosmos.NewUint(2000000)),
		}, "SWAP:ETH.HEGIC-0x584bC13c7D411c00c01A62e8019472dE68768430")
	observedTx1 := NewObservedTx(tx1, 1281323, vault.PubKey, 1281323)
	voter1 := NewObservedTxVoter(inTxID1, []ObservedTx{
		observedTx1,
	})
	voter1.Actions = []TxOutItem{
		{
			Chain:       common.ETHChain,
			ToAddress:   "0x3a196410a0f5facd08fd7880a4b8551cd085c031",
			VaultPubKey: vault.PubKey,
			Coin:        common.NewCoin(hegicAsset, cosmos.NewUint(3083783295390)),
			Memo:        "OUT:5DA19C1C5C2F6BBDF9D4FB0E6FF16A0DF6D6D7FE1F8E95CA755E5B3C6AADA458",
			InHash:      "5DA19C1C5C2F6BBDF9D4FB0E6FF16A0DF6D6D7FE1F8E95CA755E5B3C6AADA458",
		},
		{
			Chain:       common.ETHChain,
			ToAddress:   "0x3a196410a0f5facd08fd7880a4b8551cd085c031",
			VaultPubKey: vault.PubKey,
			Coin:        common.NewCoin(hegicAsset, cosmos.NewUint(2481151780248)),
			Memo:        "OUT:5DA19C1C5C2F6BBDF9D4FB0E6FF16A0DF6D6D7FE1F8E95CA755E5B3C6AADA458",
			InHash:      "5DA19C1C5C2F6BBDF9D4FB0E6FF16A0DF6D6D7FE1F8E95CA755E5B3C6AADA458",
		},
	}
	voter1.Tx = voter.Txs[0]
	storeMgr.mgr.Keeper().SetObservedTxInVoter(ctx, voter1)

	inTxID2, err := common.NewTxID("D58D1EF6D6E49EB99D0524128C16115893396FD05877EF4856FCE474B5BA09A7")
	c.Assert(err, IsNil)
	tx2 := common.NewTx(inTxID2,
		"0x3a196410a0f5facd08fd7880a4b8551cd085c031",
		"0xf56cBa49337A624E94042e325Ad6Bc864436E370",
		common.NewCoins(common.NewCoin(common.ETHAsset, cosmos.NewUint(150005145000))),
		common.Gas{
			common.NewCoin(common.RuneNative, cosmos.NewUint(2000000)),
		}, "SWAP:ETH.ETH")
	observedTx2 := NewObservedTx(tx2, 1281323, vault.PubKey, 1281323)
	voter2 := NewObservedTxVoter(inTxID2, []ObservedTx{
		observedTx2,
	})
	voter2.Actions = []TxOutItem{
		{
			Chain:       common.ETHChain,
			ToAddress:   "0x3a196410a0f5facd08fd7880a4b8551cd085c031",
			VaultPubKey: vault.PubKey,
			Coin:        common.NewCoin(hegicAsset, cosmos.NewUint(150003465000)),
			Memo:        "REFUND:D58D1EF6D6E49EB99D0524128C16115893396FD05877EF4856FCE474B5BA09A7",
			InHash:      "D58D1EF6D6E49EB99D0524128C16115893396FD05877EF4856FCE474B5BA09A7",
		},
	}
	voter2.Tx = voter.Txs[0]
	storeMgr.mgr.Keeper().SetObservedTxInVoter(ctx, voter2)

	allTxIDs := []common.TxID{
		inTxID, inTxID1, inTxID2,
	}
	removeTransactions(ctx, mgr, inTxID.String(), inTxID1.String(), inTxID2.String())
	for _, txID := range allTxIDs {
		voterAfter, err := storeMgr.mgr.Keeper().GetObservedTxInVoter(ctx, txID)
		c.Assert(err, IsNil)
		txAfter := voterAfter.GetTx(NodeAccounts{})
		c.Assert(txAfter.IsDone(len(voterAfter.Actions)), Equals, true)
	}
}

// Check that the hashing behaves as expected.
func (s *StoreManagerTestSuite) TestMemoHash(c *C) {
	inboundTxID := "B07A6B1B40ADBA2E404D9BCE1BEF6EDE6F70AD135E83806E4F4B6863CF637D0B"
	memo := fmt.Sprintf("REFUND:%s", inboundTxID)

	// This is the hash produced if using sha256 instead of Keccak-256
	// (which gave EE31ACC02D631DC3220990A1DD2E9030F4CFC227A61E975B5DEF1037106D1CCD)
	hash := fmt.Sprintf("%X", sha256.Sum256([]byte(memo)))
	fakeTxID, err := common.NewTxID(hash)
	c.Assert(err, IsNil)
	c.Assert(fakeTxID.String(), Equals, "AC0605F714563B3D5A34C64CCB6D90C1EA4EF13E1BA5E8638FE1FC796547332F")
}
