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

// MnemonicRequest Mnemonic for seeding the wallet response
type MnemonicRequest struct {
	// Mnemonic for seeding
	Mnemonic string `json:"mnemonic"`
	// Flag that specify if exisiting wallet should be replaced, if false is send then error will be send back
	Force *bool `json:"force,omitempty"`
}

// NewMnemonicRequest instantiates a new MnemonicRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMnemonicRequest(mnemonic string) *MnemonicRequest {
	this := MnemonicRequest{}
	this.Mnemonic = mnemonic
	var force bool = false
	this.Force = &force
	return &this
}

// NewMnemonicRequestWithDefaults instantiates a new MnemonicRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMnemonicRequestWithDefaults() *MnemonicRequest {
	this := MnemonicRequest{}
	var force bool = false
	this.Force = &force
	return &this
}

// GetMnemonic returns the Mnemonic field value
func (o *MnemonicRequest) GetMnemonic() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mnemonic
}

// GetMnemonicOk returns a tuple with the Mnemonic field value
// and a boolean to check if the value has been set.
func (o *MnemonicRequest) GetMnemonicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mnemonic, true
}

// SetMnemonic sets field value
func (o *MnemonicRequest) SetMnemonic(v string) {
	o.Mnemonic = v
}

// GetForce returns the Force field value if set, zero value otherwise.
func (o *MnemonicRequest) GetForce() bool {
	if o == nil || o.Force == nil {
		var ret bool
		return ret
	}
	return *o.Force
}

// GetForceOk returns a tuple with the Force field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MnemonicRequest) GetForceOk() (*bool, bool) {
	if o == nil || o.Force == nil {
		return nil, false
	}
	return o.Force, true
}

// HasForce returns a boolean if a field has been set.
func (o *MnemonicRequest) HasForce() bool {
	if o != nil && o.Force != nil {
		return true
	}

	return false
}

// SetForce gets a reference to the given bool and assigns it to the Force field.
func (o *MnemonicRequest) SetForce(v bool) {
	o.Force = &v
}

func (o MnemonicRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mnemonic"] = o.Mnemonic
	}
	if o.Force != nil {
		toSerialize["force"] = o.Force
	}
	return json.Marshal(toSerialize)
}

type NullableMnemonicRequest struct {
	value *MnemonicRequest
	isSet bool
}

func (v NullableMnemonicRequest) Get() *MnemonicRequest {
	return v.value
}

func (v *NullableMnemonicRequest) Set(val *MnemonicRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMnemonicRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMnemonicRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMnemonicRequest(val *MnemonicRequest) *NullableMnemonicRequest {
	return &NullableMnemonicRequest{value: val, isSet: true}
}

func (v NullableMnemonicRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMnemonicRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


