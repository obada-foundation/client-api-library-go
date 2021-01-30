/*
 * OBADA Client Helper API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Contact: techops@obada.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// BlockChainObit struct for BlockChainObit
type BlockChainObit struct {
	RootHash *string `json:"root_hash,omitempty"`
	ObitDid *string `json:"obit_did,omitempty"`
}

// NewBlockChainObit instantiates a new BlockChainObit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockChainObit() *BlockChainObit {
	this := BlockChainObit{}
	return &this
}

// NewBlockChainObitWithDefaults instantiates a new BlockChainObit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockChainObitWithDefaults() *BlockChainObit {
	this := BlockChainObit{}
	return &this
}

// GetRootHash returns the RootHash field value if set, zero value otherwise.
func (o *BlockChainObit) GetRootHash() string {
	if o == nil || o.RootHash == nil {
		var ret string
		return ret
	}
	return *o.RootHash
}

// GetRootHashOk returns a tuple with the RootHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockChainObit) GetRootHashOk() (*string, bool) {
	if o == nil || o.RootHash == nil {
		return nil, false
	}
	return o.RootHash, true
}

// HasRootHash returns a boolean if a field has been set.
func (o *BlockChainObit) HasRootHash() bool {
	if o != nil && o.RootHash != nil {
		return true
	}

	return false
}

// SetRootHash gets a reference to the given string and assigns it to the RootHash field.
func (o *BlockChainObit) SetRootHash(v string) {
	o.RootHash = &v
}

// GetObitDid returns the ObitDid field value if set, zero value otherwise.
func (o *BlockChainObit) GetObitDid() string {
	if o == nil || o.ObitDid == nil {
		var ret string
		return ret
	}
	return *o.ObitDid
}

// GetObitDidOk returns a tuple with the ObitDid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockChainObit) GetObitDidOk() (*string, bool) {
	if o == nil || o.ObitDid == nil {
		return nil, false
	}
	return o.ObitDid, true
}

// HasObitDid returns a boolean if a field has been set.
func (o *BlockChainObit) HasObitDid() bool {
	if o != nil && o.ObitDid != nil {
		return true
	}

	return false
}

// SetObitDid gets a reference to the given string and assigns it to the ObitDid field.
func (o *BlockChainObit) SetObitDid(v string) {
	o.ObitDid = &v
}

func (o BlockChainObit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RootHash != nil {
		toSerialize["root_hash"] = o.RootHash
	}
	if o.ObitDid != nil {
		toSerialize["obit_did"] = o.ObitDid
	}
	return json.Marshal(toSerialize)
}

type NullableBlockChainObit struct {
	value *BlockChainObit
	isSet bool
}

func (v NullableBlockChainObit) Get() *BlockChainObit {
	return v.value
}

func (v *NullableBlockChainObit) Set(val *BlockChainObit) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockChainObit) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockChainObit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockChainObit(val *BlockChainObit) *NullableBlockChainObit {
	return &NullableBlockChainObit{value: val, isSet: true}
}

func (v NullableBlockChainObit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockChainObit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


