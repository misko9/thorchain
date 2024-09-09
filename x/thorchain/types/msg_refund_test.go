package types

import (
	. "gopkg.in/check.v1"

	"gitlab.com/thorchain/thornode/common"
	"gitlab.com/thorchain/thornode/common/cosmos"
)

type MsgRefundTxSuite struct{}

var _ = Suite(&MsgRefundTxSuite{})

func (MsgRefundTxSuite) TestMsgRefundTx(c *C) {
	txID := GetRandomTxHash()
	inTxID := GetRandomTxHash()
	eth := GetRandomETHAddress()
	acc1 := GetRandomBech32Addr()
	tx := NewObservedTx(common.NewTx(
		txID,
		eth,
		GetRandomETHAddress(),
		common.Coins{common.NewCoin(common.ETHAsset, cosmos.OneUint())},
		common.Gas{common.NewCoin(common.ETHAsset, cosmos.NewUint(common.One))},
		"",
	), 12, GetRandomPubKey(), 12)
	m := NewMsgRefundTx(tx, inTxID, acc1)
	EnsureMsgBasicCorrect(m, c)
	c.Check(m.Type(), Equals, "set_tx_refund")

	inputs := []struct {
		txID   common.TxID
		inTxID common.TxID
		sender common.Address
		signer cosmos.AccAddress
	}{
		{
			txID:   common.TxID(""),
			inTxID: inTxID,
			sender: eth,
			signer: acc1,
		},
		{
			txID:   txID,
			inTxID: common.TxID(""),
			sender: eth,
			signer: acc1,
		},
		{
			txID:   txID,
			inTxID: inTxID,
			sender: common.NoAddress,
			signer: acc1,
		},
		{
			txID:   txID,
			inTxID: inTxID,
			sender: eth,
			signer: cosmos.AccAddress{},
		},
	}

	for _, item := range inputs {
		tx = NewObservedTx(common.NewTx(
			item.txID,
			item.sender,
			GetRandomETHAddress(),
			common.Coins{common.NewCoin(common.ETHAsset, cosmos.OneUint())},
			common.Gas{common.NewCoin(common.ETHAsset, cosmos.NewUint(common.One))},
			"",
		), 12, GetRandomPubKey(), 12)
		m = NewMsgRefundTx(tx, item.inTxID, item.signer)
		err := m.ValidateBasic()
		c.Assert(err, NotNil, Commentf("%s", err.Error()))
	}
}
