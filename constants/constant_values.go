package constants

import (
	"fmt"

	"github.com/blang/semver"
)

// ConstantName the name we used to get constant values.
//
//go:generate stringer -type=ConstantName
type ConstantName int

const (
	EmissionCurve ConstantName = iota
	MaxRuneSupply
	BlocksPerYear
	OutboundTransactionFee
	NativeTransactionFee
	PoolCycle
	MinRunePoolDepth
	MaxAvailablePools
	StagedPoolCost
	PendingLiquidityAgeLimit
	MinimumNodesForBFT
	DesiredValidatorSet
	AsgardSize
	DerivedDepthBasisPts
	DerivedMinDepth
	MaxAnchorSlip
	MaxAnchorBlocks
	DynamicMaxAnchorSlipBlocks
	DynamicMaxAnchorTarget
	DynamicMaxAnchorCalcInterval
	ChurnInterval
	ChurnRetryInterval
	ValidatorsChangeWindow
	LeaveProcessPerBlockHeight
	BadValidatorRedline
	LackOfObservationPenalty
	SigningTransactionPeriod
	DoubleSignMaxAge
	PauseBond
	PauseUnbond
	MinimumBondInRune
	FundMigrationInterval
	ArtificialRagnarokBlockHeight
	MaximumLiquidityRune
	StrictBondLiquidityRatio
	DefaultPoolStatus
	MaxOutboundAttempts
	SlashPenalty
	PauseOnSlashThreshold
	FailKeygenSlashPoints
	FailKeysignSlashPoints
	LiquidityLockUpBlocks
	ObserveSlashPoints
	DoubleBlockSignSlashPoints
	MissBlockSignSlashPoints
	ObservationDelayFlexibility
	JailTimeKeygen
	JailTimeKeysign
	NodePauseChainBlocks
	EnableDerivedAssets
	MinSwapsPerBlock
	MaxSwapsPerBlock
	EnableOrderBooks
	MintSynths
	BurnSynths
	MaxSynthPerPoolDepth
	MaxSynthsForSaversYield
	VirtualMultSynths
	VirtualMultSynthsBasisPoints
	MinSlashPointsForBadValidator
	BondLockupPeriod
	MaxBondProviders
	NumberOfNewNodesPerChurn
	MinTxOutVolumeThreshold
	TxOutDelayRate
	TxOutDelayMax
	MaxTxOutOffset
	TNSRegisterFee
	TNSFeeOnSale
	TNSFeePerBlock
	StreamingSwapPause
	StreamingSwapMinBPFee
	StreamingSwapMaxLength
	StreamingSwapMaxLengthNative
	MinCR
	MaxCR
	LoanStreamingSwapsInterval
	PauseLoans
	LoanRepaymentMaturity
	LendingLever
	PermittedSolvencyGap
	NodeOperatorFee
	ValidatorMaxRewardRatio
	MaxNodeToChurnOutForLowVersion
	ChurnOutForLowVersionBlocks
	POLMaxNetworkDeposit
	POLMaxPoolMovement
	POLTargetSynthPerPoolDepth
	POLBuffer
	RagnarokProcessNumOfLPPerIteration
	SwapOutDexAggregationDisabled
	SynthYieldBasisPoints
	SynthYieldCycle
	MinimumL1OutboundFeeUSD
	MinimumPoolLiquidityFee
	ChurnMigrateRounds
	AllowWideBlame
	MaxAffiliateFeeBasisPoints
	TargetOutboundFeeSurplusRune
	MaxOutboundFeeMultiplierBasisPoints
	MinOutboundFeeMultiplierBasisPoints
	NativeOutboundFeeUSD
	NativeTransactionFeeUSD
	TNSRegisterFeeUSD
	TNSFeePerBlockUSD
	EnableUSDFees
	PreferredAssetOutboundFeeMultiplier
	FeeUSDRoundSignificantDigits
	MigrationVaultSecurityBps
	CloutReset
	CloutLimit
	KeygenRetryInterval
	SaversStreamingSwapsInterval
	RescheduleCoalesceBlocks
	SignerConcurrency
	L1SlipMinBps
	SynthSlipMinBps
	TradeAccountsSlipMinBps
	TradeAccountsEnabled
	EVMDisableContractWhitelist
	OperationalVotesMin
	RUNEPoolEnabled
	RUNEPoolDepositMaturityBlocks
	RUNEPoolMaxReserveBackstop
	SaversEjectInterval
)

// ConstantValues define methods used to get constant values
type ConstantValues interface {
	fmt.Stringer
	GetInt64Value(name ConstantName) int64
	GetBoolValue(name ConstantName) bool
	GetStringValue(name ConstantName) string
	GetConstantValByKeyname() ConstantValsByKeyname
}

// GetConstantValues will return an  implementation of ConstantValues which provide ways to get constant values
// TODO hard fork remove unused version parameter
func GetConstantValues(_ semver.Version) ConstantValues {
	return NewConstantValue()
}
