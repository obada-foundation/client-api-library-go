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

// ImportAccountRequest OBADA account import payload
type ImportAccountRequest struct {
	// OBADA account
	PrivateKey string `json:"private_key"`
	// Associative account name
	AccountName *string `json:"account_name,omitempty"`
}

// NewImportAccountRequest instantiates a new ImportAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportAccountRequest(privateKey string) *ImportAccountRequest {
	this := ImportAccountRequest{}
	this.PrivateKey = privateKey
	return &this
}

// NewImportAccountRequestWithDefaults instantiates a new ImportAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportAccountRequestWithDefaults() *ImportAccountRequest {
	this := ImportAccountRequest{}
	return &this
}

// GetPrivateKey returns the PrivateKey field value
func (o *ImportAccountRequest) GetPrivateKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value
// and a boolean to check if the value has been set.
func (o *ImportAccountRequest) GetPrivateKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateKey, true
}

// SetPrivateKey sets field value
func (o *ImportAccountRequest) SetPrivateKey(v string) {
	o.PrivateKey = v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *ImportAccountRequest) GetAccountName() string {
	if o == nil || o.AccountName == nil {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportAccountRequest) GetAccountNameOk() (*string, bool) {
	if o == nil || o.AccountName == nil {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *ImportAccountRequest) HasAccountName() bool {
	if o != nil && o.AccountName != nil {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *ImportAccountRequest) SetAccountName(v string) {
	o.AccountName = &v
}

func (o ImportAccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["private_key"] = o.PrivateKey
	}
	if o.AccountName != nil {
		toSerialize["account_name"] = o.AccountName
	}
	return json.Marshal(toSerialize)
}

type NullableImportAccountRequest struct {
	value *ImportAccountRequest
	isSet bool
}

func (v NullableImportAccountRequest) Get() *ImportAccountRequest {
	return v.value
}

func (v *NullableImportAccountRequest) Set(val *ImportAccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableImportAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableImportAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportAccountRequest(val *ImportAccountRequest) *NullableImportAccountRequest {
	return &NullableImportAccountRequest{value: val, isSet: true}
}

func (v NullableImportAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


