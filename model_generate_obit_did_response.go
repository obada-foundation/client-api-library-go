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

// GenerateObitDIDResponse Returns steps that used for DID generation
type GenerateObitDIDResponse struct {
	// Serial Number Hash is sha256 of the Serial Number input
	SerialNumberHash *string `json:"serial_number_hash,omitempty"`
	// Universal Serial Number
	Usn *string `json:"usn,omitempty"`
	// DID
	Did *string `json:"did,omitempty"`
	// Base58
	UsnBase58 *string `json:"usn_base58,omitempty"`
}

// NewGenerateObitDIDResponse instantiates a new GenerateObitDIDResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateObitDIDResponse() *GenerateObitDIDResponse {
	this := GenerateObitDIDResponse{}
	return &this
}

// NewGenerateObitDIDResponseWithDefaults instantiates a new GenerateObitDIDResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateObitDIDResponseWithDefaults() *GenerateObitDIDResponse {
	this := GenerateObitDIDResponse{}
	return &this
}

// GetSerialNumberHash returns the SerialNumberHash field value if set, zero value otherwise.
func (o *GenerateObitDIDResponse) GetSerialNumberHash() string {
	if o == nil || o.SerialNumberHash == nil {
		var ret string
		return ret
	}
	return *o.SerialNumberHash
}

// GetSerialNumberHashOk returns a tuple with the SerialNumberHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateObitDIDResponse) GetSerialNumberHashOk() (*string, bool) {
	if o == nil || o.SerialNumberHash == nil {
		return nil, false
	}
	return o.SerialNumberHash, true
}

// HasSerialNumberHash returns a boolean if a field has been set.
func (o *GenerateObitDIDResponse) HasSerialNumberHash() bool {
	if o != nil && o.SerialNumberHash != nil {
		return true
	}

	return false
}

// SetSerialNumberHash gets a reference to the given string and assigns it to the SerialNumberHash field.
func (o *GenerateObitDIDResponse) SetSerialNumberHash(v string) {
	o.SerialNumberHash = &v
}

// GetUsn returns the Usn field value if set, zero value otherwise.
func (o *GenerateObitDIDResponse) GetUsn() string {
	if o == nil || o.Usn == nil {
		var ret string
		return ret
	}
	return *o.Usn
}

// GetUsnOk returns a tuple with the Usn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateObitDIDResponse) GetUsnOk() (*string, bool) {
	if o == nil || o.Usn == nil {
		return nil, false
	}
	return o.Usn, true
}

// HasUsn returns a boolean if a field has been set.
func (o *GenerateObitDIDResponse) HasUsn() bool {
	if o != nil && o.Usn != nil {
		return true
	}

	return false
}

// SetUsn gets a reference to the given string and assigns it to the Usn field.
func (o *GenerateObitDIDResponse) SetUsn(v string) {
	o.Usn = &v
}

// GetDid returns the Did field value if set, zero value otherwise.
func (o *GenerateObitDIDResponse) GetDid() string {
	if o == nil || o.Did == nil {
		var ret string
		return ret
	}
	return *o.Did
}

// GetDidOk returns a tuple with the Did field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateObitDIDResponse) GetDidOk() (*string, bool) {
	if o == nil || o.Did == nil {
		return nil, false
	}
	return o.Did, true
}

// HasDid returns a boolean if a field has been set.
func (o *GenerateObitDIDResponse) HasDid() bool {
	if o != nil && o.Did != nil {
		return true
	}

	return false
}

// SetDid gets a reference to the given string and assigns it to the Did field.
func (o *GenerateObitDIDResponse) SetDid(v string) {
	o.Did = &v
}

// GetUsnBase58 returns the UsnBase58 field value if set, zero value otherwise.
func (o *GenerateObitDIDResponse) GetUsnBase58() string {
	if o == nil || o.UsnBase58 == nil {
		var ret string
		return ret
	}
	return *o.UsnBase58
}

// GetUsnBase58Ok returns a tuple with the UsnBase58 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateObitDIDResponse) GetUsnBase58Ok() (*string, bool) {
	if o == nil || o.UsnBase58 == nil {
		return nil, false
	}
	return o.UsnBase58, true
}

// HasUsnBase58 returns a boolean if a field has been set.
func (o *GenerateObitDIDResponse) HasUsnBase58() bool {
	if o != nil && o.UsnBase58 != nil {
		return true
	}

	return false
}

// SetUsnBase58 gets a reference to the given string and assigns it to the UsnBase58 field.
func (o *GenerateObitDIDResponse) SetUsnBase58(v string) {
	o.UsnBase58 = &v
}

func (o GenerateObitDIDResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SerialNumberHash != nil {
		toSerialize["serial_number_hash"] = o.SerialNumberHash
	}
	if o.Usn != nil {
		toSerialize["usn"] = o.Usn
	}
	if o.Did != nil {
		toSerialize["did"] = o.Did
	}
	if o.UsnBase58 != nil {
		toSerialize["usn_base58"] = o.UsnBase58
	}
	return json.Marshal(toSerialize)
}

type NullableGenerateObitDIDResponse struct {
	value *GenerateObitDIDResponse
	isSet bool
}

func (v NullableGenerateObitDIDResponse) Get() *GenerateObitDIDResponse {
	return v.value
}

func (v *NullableGenerateObitDIDResponse) Set(val *GenerateObitDIDResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateObitDIDResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateObitDIDResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateObitDIDResponse(val *GenerateObitDIDResponse) *NullableGenerateObitDIDResponse {
	return &NullableGenerateObitDIDResponse{value: val, isSet: true}
}

func (v NullableGenerateObitDIDResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateObitDIDResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

