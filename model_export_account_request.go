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

// ExportAccountRequest OBADA account export payload
type ExportAccountRequest struct {
	// OBADA account
	Address *string `json:"address,omitempty"`
	// Passphrase to decrypt the account
	Passphrase *string `json:"passphrase,omitempty"`
}

// NewExportAccountRequest instantiates a new ExportAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportAccountRequest() *ExportAccountRequest {
	this := ExportAccountRequest{}
	return &this
}

// NewExportAccountRequestWithDefaults instantiates a new ExportAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportAccountRequestWithDefaults() *ExportAccountRequest {
	this := ExportAccountRequest{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ExportAccountRequest) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportAccountRequest) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ExportAccountRequest) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *ExportAccountRequest) SetAddress(v string) {
	o.Address = &v
}

// GetPassphrase returns the Passphrase field value if set, zero value otherwise.
func (o *ExportAccountRequest) GetPassphrase() string {
	if o == nil || o.Passphrase == nil {
		var ret string
		return ret
	}
	return *o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportAccountRequest) GetPassphraseOk() (*string, bool) {
	if o == nil || o.Passphrase == nil {
		return nil, false
	}
	return o.Passphrase, true
}

// HasPassphrase returns a boolean if a field has been set.
func (o *ExportAccountRequest) HasPassphrase() bool {
	if o != nil && o.Passphrase != nil {
		return true
	}

	return false
}

// SetPassphrase gets a reference to the given string and assigns it to the Passphrase field.
func (o *ExportAccountRequest) SetPassphrase(v string) {
	o.Passphrase = &v
}

func (o ExportAccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Passphrase != nil {
		toSerialize["passphrase"] = o.Passphrase
	}
	return json.Marshal(toSerialize)
}

type NullableExportAccountRequest struct {
	value *ExportAccountRequest
	isSet bool
}

func (v NullableExportAccountRequest) Get() *ExportAccountRequest {
	return v.value
}

func (v *NullableExportAccountRequest) Set(val *ExportAccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableExportAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableExportAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportAccountRequest(val *ExportAccountRequest) *NullableExportAccountRequest {
	return &NullableExportAccountRequest{value: val, isSet: true}
}

func (v NullableExportAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


