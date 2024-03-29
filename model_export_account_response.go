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

// ExportAccountResponse OBADA account export response
type ExportAccountResponse struct {
	// OBADA account
	PrivateKey *string `json:"private_key,omitempty"`
}

// NewExportAccountResponse instantiates a new ExportAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportAccountResponse() *ExportAccountResponse {
	this := ExportAccountResponse{}
	return &this
}

// NewExportAccountResponseWithDefaults instantiates a new ExportAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportAccountResponseWithDefaults() *ExportAccountResponse {
	this := ExportAccountResponse{}
	return &this
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *ExportAccountResponse) GetPrivateKey() string {
	if o == nil || o.PrivateKey == nil {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportAccountResponse) GetPrivateKeyOk() (*string, bool) {
	if o == nil || o.PrivateKey == nil {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *ExportAccountResponse) HasPrivateKey() bool {
	if o != nil && o.PrivateKey != nil {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *ExportAccountResponse) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

func (o ExportAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PrivateKey != nil {
		toSerialize["private_key"] = o.PrivateKey
	}
	return json.Marshal(toSerialize)
}

type NullableExportAccountResponse struct {
	value *ExportAccountResponse
	isSet bool
}

func (v NullableExportAccountResponse) Get() *ExportAccountResponse {
	return v.value
}

func (v *NullableExportAccountResponse) Set(val *ExportAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExportAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExportAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportAccountResponse(val *ExportAccountResponse) *NullableExportAccountResponse {
	return &NullableExportAccountResponse{value: val, isSet: true}
}

func (v NullableExportAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


