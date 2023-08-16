/*
OBADA API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
Contact: techops@obada.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github.com/obada-foundation/client-api-library-go

import (
	"encoding/json"
)

// AccountBalance struct for AccountBalance
type AccountBalance struct {
	// Denomination unit
	Denom *string `json:"denom,omitempty"`
	Amount *string `json:"amount,omitempty"`
}

// NewAccountBalance instantiates a new AccountBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountBalance() *AccountBalance {
	this := AccountBalance{}
	return &this
}

// NewAccountBalanceWithDefaults instantiates a new AccountBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountBalanceWithDefaults() *AccountBalance {
	this := AccountBalance{}
	return &this
}

// GetDenom returns the Denom field value if set, zero value otherwise.
func (o *AccountBalance) GetDenom() string {
	if o == nil || o.Denom == nil {
		var ret string
		return ret
	}
	return *o.Denom
}

// GetDenomOk returns a tuple with the Denom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBalance) GetDenomOk() (*string, bool) {
	if o == nil || o.Denom == nil {
		return nil, false
	}
	return o.Denom, true
}

// HasDenom returns a boolean if a field has been set.
func (o *AccountBalance) HasDenom() bool {
	if o != nil && o.Denom != nil {
		return true
	}

	return false
}

// SetDenom gets a reference to the given string and assigns it to the Denom field.
func (o *AccountBalance) SetDenom(v string) {
	o.Denom = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *AccountBalance) GetAmount() string {
	if o == nil || o.Amount == nil {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBalance) GetAmountOk() (*string, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *AccountBalance) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *AccountBalance) SetAmount(v string) {
	o.Amount = &v
}

func (o AccountBalance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Denom != nil {
		toSerialize["denom"] = o.Denom
	}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	return json.Marshal(toSerialize)
}

type NullableAccountBalance struct {
	value *AccountBalance
	isSet bool
}

func (v NullableAccountBalance) Get() *AccountBalance {
	return v.value
}

func (v *NullableAccountBalance) Set(val *AccountBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountBalance(val *AccountBalance) *NullableAccountBalance {
	return &NullableAccountBalance{value: val, isSet: true}
}

func (v NullableAccountBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


