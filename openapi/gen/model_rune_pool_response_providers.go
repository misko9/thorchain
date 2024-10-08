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

// RUNEPoolResponseProviders struct for RUNEPoolResponseProviders
type RUNEPoolResponseProviders struct {
	// the units of RUNEPool owned by providers (including pending)
	Units string `json:"units"`
	// the units of RUNEPool owned by providers that remain pending
	PendingUnits string `json:"pending_units"`
	// the amount of RUNE pending
	PendingRune string `json:"pending_rune"`
	// the value of the provider share of the RUNEPool (includes pending RUNE)
	Value string `json:"value"`
	// the profit and loss of the provider share of the RUNEPool
	Pnl string `json:"pnl"`
	// the current RUNE deposited by providers
	CurrentDeposit string `json:"current_deposit"`
}

// NewRUNEPoolResponseProviders instantiates a new RUNEPoolResponseProviders object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRUNEPoolResponseProviders(units string, pendingUnits string, pendingRune string, value string, pnl string, currentDeposit string) *RUNEPoolResponseProviders {
	this := RUNEPoolResponseProviders{}
	this.Units = units
	this.PendingUnits = pendingUnits
	this.PendingRune = pendingRune
	this.Value = value
	this.Pnl = pnl
	this.CurrentDeposit = currentDeposit
	return &this
}

// NewRUNEPoolResponseProvidersWithDefaults instantiates a new RUNEPoolResponseProviders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRUNEPoolResponseProvidersWithDefaults() *RUNEPoolResponseProviders {
	this := RUNEPoolResponseProviders{}
	return &this
}

// GetUnits returns the Units field value
func (o *RUNEPoolResponseProviders) GetUnits() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Units
}

// GetUnitsOk returns a tuple with the Units field value
// and a boolean to check if the value has been set.
func (o *RUNEPoolResponseProviders) GetUnitsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Units, true
}

// SetUnits sets field value
func (o *RUNEPoolResponseProviders) SetUnits(v string) {
	o.Units = v
}

// GetPendingUnits returns the PendingUnits field value
func (o *RUNEPoolResponseProviders) GetPendingUnits() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PendingUnits
}

// GetPendingUnitsOk returns a tuple with the PendingUnits field value
// and a boolean to check if the value has been set.
func (o *RUNEPoolResponseProviders) GetPendingUnitsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingUnits, true
}

// SetPendingUnits sets field value
func (o *RUNEPoolResponseProviders) SetPendingUnits(v string) {
	o.PendingUnits = v
}

// GetPendingRune returns the PendingRune field value
func (o *RUNEPoolResponseProviders) GetPendingRune() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PendingRune
}

// GetPendingRuneOk returns a tuple with the PendingRune field value
// and a boolean to check if the value has been set.
func (o *RUNEPoolResponseProviders) GetPendingRuneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingRune, true
}

// SetPendingRune sets field value
func (o *RUNEPoolResponseProviders) SetPendingRune(v string) {
	o.PendingRune = v
}

// GetValue returns the Value field value
func (o *RUNEPoolResponseProviders) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *RUNEPoolResponseProviders) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *RUNEPoolResponseProviders) SetValue(v string) {
	o.Value = v
}

// GetPnl returns the Pnl field value
func (o *RUNEPoolResponseProviders) GetPnl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pnl
}

// GetPnlOk returns a tuple with the Pnl field value
// and a boolean to check if the value has been set.
func (o *RUNEPoolResponseProviders) GetPnlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pnl, true
}

// SetPnl sets field value
func (o *RUNEPoolResponseProviders) SetPnl(v string) {
	o.Pnl = v
}

// GetCurrentDeposit returns the CurrentDeposit field value
func (o *RUNEPoolResponseProviders) GetCurrentDeposit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrentDeposit
}

// GetCurrentDepositOk returns a tuple with the CurrentDeposit field value
// and a boolean to check if the value has been set.
func (o *RUNEPoolResponseProviders) GetCurrentDepositOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentDeposit, true
}

// SetCurrentDeposit sets field value
func (o *RUNEPoolResponseProviders) SetCurrentDeposit(v string) {
	o.CurrentDeposit = v
}

func (o RUNEPoolResponseProviders) MarshalJSON_deprecated() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["units"] = o.Units
	}
	if true {
		toSerialize["pending_units"] = o.PendingUnits
	}
	if true {
		toSerialize["pending_rune"] = o.PendingRune
	}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["pnl"] = o.Pnl
	}
	if true {
		toSerialize["current_deposit"] = o.CurrentDeposit
	}
	return json.Marshal(toSerialize)
}

type NullableRUNEPoolResponseProviders struct {
	value *RUNEPoolResponseProviders
	isSet bool
}

func (v NullableRUNEPoolResponseProviders) Get() *RUNEPoolResponseProviders {
	return v.value
}

func (v *NullableRUNEPoolResponseProviders) Set(val *RUNEPoolResponseProviders) {
	v.value = val
	v.isSet = true
}

func (v NullableRUNEPoolResponseProviders) IsSet() bool {
	return v.isSet
}

func (v *NullableRUNEPoolResponseProviders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRUNEPoolResponseProviders(val *RUNEPoolResponseProviders) *NullableRUNEPoolResponseProviders {
	return &NullableRUNEPoolResponseProviders{value: val, isSet: true}
}

func (v NullableRUNEPoolResponseProviders) MarshalJSON_deprecated() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRUNEPoolResponseProviders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


