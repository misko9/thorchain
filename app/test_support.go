package app

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
)

func (app *ChainApp) GetBaseApp() *baseapp.BaseApp {
	return app.BaseApp
}

func (app *ChainApp) GetBankKeeper() bankkeeper.Keeper {
	return app.BankKeeper
}

func (app *ChainApp) GetStakingKeeper() *stakingkeeper.Keeper {
	return app.StakingKeeper
}

func (app *ChainApp) GetAccountKeeper() authkeeper.AccountKeeper {
	return app.AccountKeeper
}
