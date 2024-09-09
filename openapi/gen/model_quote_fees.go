/*
Thornode API

Thornode REST API.

Contact: devs@thorchain.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// QuoteFees struct for QuoteFees
type QuoteFees struct {
	// the target asset used for all fees
	Asset string `json:"asset"`
	// affiliate fee in the target asset
	Affiliate *string `json:"affiliate,omitempty"`
	// outbound fee in the target asset
	Outbound *string `json:"outbound,omitempty"`
	// liquidity fees paid to pools in the target asset
	Liquidity string `json:"liquidity"`
	// total fees in the target asset
	Total string `json:"total"`
	// the swap slippage in basis points
	SlippageBps int64 `json:"slippage_bps"`
	// total basis points in fees relative to amount out
	TotalBps int64 `json:"total_bps"`
}

// NewQuoteFees instantiates a new QuoteFees object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteFees(asset string, liquidity string, total string, slippageBps int64, totalBps int64) *QuoteFees {
	this := QuoteFees{}
	this.Asset = asset
	this.Liquidity = liquidity
	this.Total = total
	this.SlippageBps = slippageBps
	this.TotalBps = totalBps
	return &this
}

// NewQuoteFeesWithDefaults instantiates a new QuoteFees object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteFeesWithDefaults() *QuoteFees {
	this := QuoteFees{}
	return &this
}

// GetAsset returns the Asset field value
func (o *QuoteFees) GetAsset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *QuoteFees) GetAssetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *QuoteFees) SetAsset(v string) {
	o.Asset = v
}

// GetAffiliate returns the Affiliate field value if set, zero value otherwise.
func (o *QuoteFees) GetAffiliate() string {
	if o == nil || o.Affiliate == nil {
		var ret string
		return ret
	}
	return *o.Affiliate
}

// GetAffiliateOk returns a tuple with the Affiliate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteFees) GetAffiliateOk() (*string, bool) {
	if o == nil || o.Affiliate == nil {
		return nil, false
	}
	return o.Affiliate, true
}

// HasAffiliate returns a boolean if a field has been set.
func (o *QuoteFees) HasAffiliate() bool {
	if o != nil && o.Affiliate != nil {
		return true
	}

	return false
}

// SetAffiliate gets a reference to the given string and assigns it to the Affiliate field.
func (o *QuoteFees) SetAffiliate(v string) {
	o.Affiliate = &v
}

// GetOutbound returns the Outbound field value if set, zero value otherwise.
func (o *QuoteFees) GetOutbound() string {
	if o == nil || o.Outbound == nil {
		var ret string
		return ret
	}
	return *o.Outbound
}

// GetOutboundOk returns a tuple with the Outbound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteFees) GetOutboundOk() (*string, bool) {
	if o == nil || o.Outbound == nil {
		return nil, false
	}
	return o.Outbound, true
}

// HasOutbound returns a boolean if a field has been set.
func (o *QuoteFees) HasOutbound() bool {
	if o != nil && o.Outbound != nil {
		return true
	}

	return false
}

// SetOutbound gets a reference to the given string and assigns it to the Outbound field.
func (o *QuoteFees) SetOutbound(v string) {
	o.Outbound = &v
}

// GetLiquidity returns the Liquidity field value
func (o *QuoteFees) GetLiquidity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Liquidity
}

// GetLiquidityOk returns a tuple with the Liquidity field value
// and a boolean to check if the value has been set.
func (o *QuoteFees) GetLiquidityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Liquidity, true
}

// SetLiquidity sets field value
func (o *QuoteFees) SetLiquidity(v string) {
	o.Liquidity = v
}

// GetTotal returns the Total field value
func (o *QuoteFees) GetTotal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *QuoteFees) GetTotalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *QuoteFees) SetTotal(v string) {
	o.Total = v
}

// GetSlippageBps returns the SlippageBps field value
func (o *QuoteFees) GetSlippageBps() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SlippageBps
}

// GetSlippageBpsOk returns a tuple with the SlippageBps field value
// and a boolean to check if the value has been set.
func (o *QuoteFees) GetSlippageBpsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlippageBps, true
}

// SetSlippageBps sets field value
func (o *QuoteFees) SetSlippageBps(v int64) {
	o.SlippageBps = v
}

// GetTotalBps returns the TotalBps field value
func (o *QuoteFees) GetTotalBps() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalBps
}

// GetTotalBpsOk returns a tuple with the TotalBps field value
// and a boolean to check if the value has been set.
func (o *QuoteFees) GetTotalBpsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalBps, true
}

// SetTotalBps sets field value
func (o *QuoteFees) SetTotalBps(v int64) {
	o.TotalBps = v
}

func (o QuoteFees) MarshalJSON_deprecated() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["asset"] = o.Asset
	}
	if o.Affiliate != nil {
		toSerialize["affiliate"] = o.Affiliate
	}
	if o.Outbound != nil {
		toSerialize["outbound"] = o.Outbound
	}
	if true {
		toSerialize["liquidity"] = o.Liquidity
	}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["slippage_bps"] = o.SlippageBps
	}
	if true {
		toSerialize["total_bps"] = o.TotalBps
	}
	return json.Marshal(toSerialize)
}

type NullableQuoteFees struct {
	value *QuoteFees
	isSet bool
}

func (v NullableQuoteFees) Get() *QuoteFees {
	return v.value
}

func (v *NullableQuoteFees) Set(val *QuoteFees) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteFees) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteFees) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteFees(val *QuoteFees) *NullableQuoteFees {
	return &NullableQuoteFees{value: val, isSet: true}
}

func (v NullableQuoteFees) MarshalJSON_deprecated() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteFees) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


