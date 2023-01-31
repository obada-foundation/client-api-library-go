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

// Account OBADA account
type Account struct {
	// Account address associated name
	Name *string `json:"name,omitempty"`
	// Public key
	PubKey *string `json:"pub_key,omitempty"`
	// OBADA address
	Address *string `json:"address,omitempty"`
	Balance *int64 `json:"balance,omitempty"`
	NftCount *int64 `json:"nft_count,omitempty"`
}

// NewAccount instantiates a new Account object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccount() *Account {
	this := Account{}
	return &this
}

// NewAccountWithDefaults instantiates a new Account object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountWithDefaults() *Account {
	this := Account{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Account) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Account) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Account) SetName(v string) {
	o.Name = &v
}

// GetPubKey returns the PubKey field value if set, zero value otherwise.
func (o *Account) GetPubKey() string {
	if o == nil || o.PubKey == nil {
		var ret string
		return ret
	}
	return *o.PubKey
}

// GetPubKeyOk returns a tuple with the PubKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetPubKeyOk() (*string, bool) {
	if o == nil || o.PubKey == nil {
		return nil, false
	}
	return o.PubKey, true
}

// HasPubKey returns a boolean if a field has been set.
func (o *Account) HasPubKey() bool {
	if o != nil && o.PubKey != nil {
		return true
	}

	return false
}

// SetPubKey gets a reference to the given string and assigns it to the PubKey field.
func (o *Account) SetPubKey(v string) {
	o.PubKey = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *Account) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *Account) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *Account) SetAddress(v string) {
	o.Address = &v
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *Account) GetBalance() int64 {
	if o == nil || o.Balance == nil {
		var ret int64
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetBalanceOk() (*int64, bool) {
	if o == nil || o.Balance == nil {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *Account) HasBalance() bool {
	if o != nil && o.Balance != nil {
		return true
	}

	return false
}

// SetBalance gets a reference to the given int64 and assigns it to the Balance field.
func (o *Account) SetBalance(v int64) {
	o.Balance = &v
}

// GetNftCount returns the NftCount field value if set, zero value otherwise.
func (o *Account) GetNftCount() int64 {
	if o == nil || o.NftCount == nil {
		var ret int64
		return ret
	}
	return *o.NftCount
}

// GetNftCountOk returns a tuple with the NftCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetNftCountOk() (*int64, bool) {
	if o == nil || o.NftCount == nil {
		return nil, false
	}
	return o.NftCount, true
}

// HasNftCount returns a boolean if a field has been set.
func (o *Account) HasNftCount() bool {
	if o != nil && o.NftCount != nil {
		return true
	}

	return false
}

// SetNftCount gets a reference to the given int64 and assigns it to the NftCount field.
func (o *Account) SetNftCount(v int64) {
	o.NftCount = &v
}

func (o Account) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PubKey != nil {
		toSerialize["pub_key"] = o.PubKey
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Balance != nil {
		toSerialize["balance"] = o.Balance
	}
	if o.NftCount != nil {
		toSerialize["nft_count"] = o.NftCount
	}
	return json.Marshal(toSerialize)
}

type NullableAccount struct {
	value *Account
	isSet bool
}

func (v NullableAccount) Get() *Account {
	return v.value
}

func (v *NullableAccount) Set(val *Account) {
	v.value = val
	v.isSet = true
}

func (v NullableAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccount(val *Account) *NullableAccount {
	return &NullableAccount{value: val, isSet: true}
}

func (v NullableAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


