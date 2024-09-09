package common

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/blang/semver"
	"github.com/gogo/protobuf/jsonpb"
)

var (
	// EmptyAsset empty asset, not valid
	EmptyAsset = Asset{Chain: EmptyChain, Symbol: "", Ticker: "", Synth: false}
	// ATOMAsset ATOM
	ATOMAsset = Asset{Chain: GAIAChain, Symbol: "ATOM", Ticker: "ATOM", Synth: false}
	// BNBBEP20Asset BNB
	BNBBEP20Asset = Asset{Chain: BSCChain, Symbol: "BNB", Ticker: "BNB", Synth: false}
	// BTCAsset BTC
	BTCAsset = Asset{Chain: BTCChain, Symbol: "BTC", Ticker: "BTC", Synth: false}
	// LTCAsset BTC
	LTCAsset = Asset{Chain: LTCChain, Symbol: "LTC", Ticker: "LTC", Synth: false}
	// BCHAsset BCH
	BCHAsset = Asset{Chain: BCHChain, Symbol: "BCH", Ticker: "BCH", Synth: false}
	// DOGEAsset DOGE
	DOGEAsset = Asset{Chain: DOGEChain, Symbol: "DOGE", Ticker: "DOGE", Synth: false}
	// ETHAsset ETH
	ETHAsset = Asset{Chain: ETHChain, Symbol: "ETH", Ticker: "ETH", Synth: false}
	// AVAXAsset AVAX
	AVAXAsset = Asset{Chain: AVAXChain, Symbol: "AVAX", Ticker: "AVAX", Synth: false}
	// RuneNative RUNE on thorchain
	RuneNative = Asset{Chain: THORChain, Symbol: "RUNE", Ticker: "RUNE", Synth: false}
	TOR        = Asset{Chain: THORChain, Symbol: "TOR", Ticker: "TOR", Synth: false}
	THORBTC    = Asset{Chain: THORChain, Symbol: "BTC", Ticker: "BTC", Synth: false}
)

// NewAsset parse the given input into Asset object
func NewAsset(input string) (Asset, error) {
	var err error
	var asset Asset
	var sym string
	var parts []string
	switch {
	case strings.Count(input, "~") > 0:
		parts = strings.SplitN(input, "~", 2)
		asset.Trade = true
	case strings.Count(input, "/") > 0:
		parts = strings.SplitN(input, "/", 2)
		asset.Synth = true
	default:
		parts = strings.SplitN(input, ".", 2)
	}
	if len(parts) == 1 {
		asset.Chain = THORChain
		sym = parts[0]
	} else {
		asset.Chain, err = NewChain(parts[0])
		if err != nil {
			return EmptyAsset, err
		}
		sym = parts[1]
	}

	asset.Symbol, err = NewSymbol(sym)
	if err != nil {
		return EmptyAsset, err
	}

	parts = strings.SplitN(sym, "-", 2)
	asset.Ticker, err = NewTicker(parts[0])
	if err != nil {
		return EmptyAsset, err
	}

	return asset, nil
}

func NewAssetWithShortCodes(version semver.Version, input string) (Asset, error) {
	switch {
	case version.GTE(semver.MustParse("1.124.0")):
		return NewAssetWithShortCodesV124(input)
	default:
		return NewAsset(input)
	}
}

func NewAssetWithShortCodesV124(input string) (Asset, error) {
	shorts := make(map[string]string)

	shorts[ATOMAsset.ShortCode()] = ATOMAsset.String()
	shorts[AVAXAsset.ShortCode()] = AVAXAsset.String()
	shorts[BCHAsset.ShortCode()] = BCHAsset.String()
	shorts[BNBBEP20Asset.ShortCode()] = BNBBEP20Asset.String()
	shorts[BTCAsset.ShortCode()] = BTCAsset.String()
	shorts[DOGEAsset.ShortCode()] = DOGEAsset.String()
	shorts[ETHAsset.ShortCode()] = ETHAsset.String()
	shorts[LTCAsset.ShortCode()] = LTCAsset.String()
	shorts[RuneNative.ShortCode()] = RuneNative.String()

	long, ok := shorts[input]
	if ok {
		input = long
	}

	return NewAsset(input)
}

func (a Asset) Valid() error {
	if err := a.Chain.Valid(); err != nil {
		return fmt.Errorf("invalid chain: %w", err)
	}
	if err := a.Symbol.Valid(); err != nil {
		return fmt.Errorf("invalid symbol: %w", err)
	}
	if a.Synth && a.Trade {
		return fmt.Errorf("trade assets cannot be synth assets")
	}
	if a.Synth && a.Chain.IsTHORChain() {
		return fmt.Errorf("synth asset cannot have chain THOR: %s", a)
	}
	if a.Trade && a.Chain.IsTHORChain() {
		return fmt.Errorf("trade asset cannot have chain THOR: %s", a)
	}
	return nil
}

// Equals determinate whether two assets are equivalent
func (a Asset) Equals(a2 Asset) bool {
	return a.Chain.Equals(a2.Chain) && a.Symbol.Equals(a2.Symbol) && a.Ticker.Equals(a2.Ticker) && a.Synth == a2.Synth && a.Trade == a2.Trade
}

func (a Asset) GetChain() Chain {
	if a.Synth || a.Trade {
		return THORChain
	}
	return a.Chain
}

// Get layer1 asset version
func (a Asset) GetLayer1Asset() Asset {
	if !a.IsSyntheticAsset() && !a.IsTradeAsset() {
		return a
	}
	return Asset{
		Chain:  a.Chain,
		Symbol: a.Symbol,
		Ticker: a.Ticker,
		Synth:  false,
		Trade:  false,
	}
}

// Get synthetic asset of asset
func (a Asset) GetSyntheticAsset() Asset {
	if a.IsSyntheticAsset() {
		return a
	}
	return Asset{
		Chain:  a.Chain,
		Symbol: a.Symbol,
		Ticker: a.Ticker,
		Synth:  true,
	}
}

// Get trade asset of asset
func (a Asset) GetTradeAsset() Asset {
	if a.IsTradeAsset() {
		return a
	}
	return Asset{
		Chain:  a.Chain,
		Symbol: a.Symbol,
		Ticker: a.Ticker,
		Trade:  true,
	}
}

// Get derived asset of asset
func (a Asset) GetDerivedAsset() Asset {
	return Asset{
		Chain:  THORChain,
		Symbol: a.Symbol,
		Ticker: a.Ticker,
		Synth:  false,
	}
}

// Check if asset is a pegged asset
func (a Asset) IsSyntheticAsset() bool {
	return a.Synth
}

func (a Asset) IsTradeAsset() bool {
	return a.Trade
}

func (a Asset) IsVaultAsset() bool {
	return a.IsSyntheticAsset()
}

// Check if asset is a derived asset
func (a Asset) IsDerivedAsset() bool {
	return !a.Synth && !a.Trade && a.GetChain().IsTHORChain() && !a.IsRune()
}

// Native return native asset, only relevant on THORChain
func (a Asset) Native() string {
	if a.IsRune() {
		return "rune"
	}
	if a.Equals(TOR) {
		return "tor"
	}
	return strings.ToLower(a.String())
}

// IsEmpty will be true when any of the field is empty, chain,symbol or ticker
func (a Asset) IsEmpty() bool {
	return a.Chain.IsEmpty() || a.Symbol.IsEmpty() || a.Ticker.IsEmpty()
}

// String implement fmt.Stringer , return the string representation of Asset
func (a Asset) String() string {
	div := "."
	if a.Synth {
		div = "/"
	}
	if a.Trade {
		div = "~"
	}
	return fmt.Sprintf("%s%s%s", a.Chain.String(), div, a.Symbol.String())
}

// ShortCode returns the short code for the asset.
func (a Asset) ShortCode() string {
	switch a.String() {
	case "THOR.RUNE":
		return "r"
	case "BTC.BTC":
		return "b"
	case "ETH.ETH":
		return "e"
	case "GAIA.ATOM":
		return "g"
	case "DOGE.DOGE":
		return "d"
	case "LTC.LTC":
		return "l"
	case "BCH.BCH":
		return "c"
	case "AVAX.AVAX":
		return "a"
	case "BSC.BNB":
		return "s"
	default:
		return ""
	}
}

// IsGasAsset check whether asset is base asset used to pay for gas
func (a Asset) IsGasAsset() bool {
	gasAsset := a.GetChain().GetGasAsset()
	if gasAsset.IsEmpty() {
		return false
	}
	return a.Equals(gasAsset)
}

// IsRune is a helper function ,return true only when the asset represent RUNE
func (a Asset) IsRune() bool {
	return RuneAsset().Equals(a)
}

// IsNative is a helper function, returns true when the asset is a native
// asset to THORChain (ie rune, a synth, etc)
func (a Asset) IsNative() bool {
	return a.GetChain().IsTHORChain()
}

// MarshalJSON implement Marshaler interface
func (a Asset) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.String())
}

// UnmarshalJSON implement Unmarshaler interface
func (a *Asset) UnmarshalJSON(data []byte) error {
	var err error
	var assetStr string
	if err = json.Unmarshal(data, &assetStr); err != nil {
		return err
	}
	if assetStr == "." {
		*a = EmptyAsset
		return nil
	}
	*a, err = NewAsset(assetStr)
	return err
}

// MarshalJSONPB implement jsonpb.Marshaler
func (a Asset) MarshalJSONPB(*jsonpb.Marshaler) ([]byte, error) {
	return a.MarshalJSON()
}

// UnmarshalJSONPB implement jsonpb.Unmarshaler
func (a *Asset) UnmarshalJSONPB(unmarshal *jsonpb.Unmarshaler, content []byte) error {
	return a.UnmarshalJSON(content)
}

// RuneAsset return RUNE Asset depends on different environment
func RuneAsset() Asset {
	return RuneNative
}

// Replace pool name "." with a "-" for Mimir key checking.
func (a Asset) MimirString() string {
	return a.Chain.String() + "-" + a.Symbol.String()
}
