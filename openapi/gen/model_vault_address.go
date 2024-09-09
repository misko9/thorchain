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

// VaultAddress struct for VaultAddress
type VaultAddress struct {
	Chain string `json:"chain"`
	Address string `json:"address"`
}

// NewVaultAddress instantiates a new VaultAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVaultAddress(chain string, address string) *VaultAddress {
	this := VaultAddress{}
	this.Chain = chain
	this.Address = address
	return &this
}

// NewVaultAddressWithDefaults instantiates a new VaultAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVaultAddressWithDefaults() *VaultAddress {
	this := VaultAddress{}
	return &this
}

// GetChain returns the Chain field value
func (o *VaultAddress) GetChain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Chain
}

// GetChainOk returns a tuple with the Chain field value
// and a boolean to check if the value has been set.
func (o *VaultAddress) GetChainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chain, true
}

// SetChain sets field value
func (o *VaultAddress) SetChain(v string) {
	o.Chain = v
}

// GetAddress returns the Address field value
func (o *VaultAddress) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *VaultAddress) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *VaultAddress) SetAddress(v string) {
	o.Address = v
}

func (o VaultAddress) MarshalJSON_deprecated() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["chain"] = o.Chain
	}
	if true {
		toSerialize["address"] = o.Address
	}
	return json.Marshal(toSerialize)
}

type NullableVaultAddress struct {
	value *VaultAddress
	isSet bool
}

func (v NullableVaultAddress) Get() *VaultAddress {
	return v.value
}

func (v *NullableVaultAddress) Set(val *VaultAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableVaultAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableVaultAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVaultAddress(val *VaultAddress) *NullableVaultAddress {
	return &NullableVaultAddress{value: val, isSet: true}
}

func (v NullableVaultAddress) MarshalJSON_deprecated() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVaultAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


