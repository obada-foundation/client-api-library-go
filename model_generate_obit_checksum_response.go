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

// GenerateObitChecksumResponse Obit checksum with logged data
type GenerateObitChecksumResponse struct {
	// Obit checksum
	Checksum *string `json:"checksum,omitempty"`
	// Log that contain captures from SDK
	ComputeLog *string `json:"compute_log,omitempty"`
}

// NewGenerateObitChecksumResponse instantiates a new GenerateObitChecksumResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateObitChecksumResponse() *GenerateObitChecksumResponse {
	this := GenerateObitChecksumResponse{}
	return &this
}

// NewGenerateObitChecksumResponseWithDefaults instantiates a new GenerateObitChecksumResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateObitChecksumResponseWithDefaults() *GenerateObitChecksumResponse {
	this := GenerateObitChecksumResponse{}
	return &this
}

// GetChecksum returns the Checksum field value if set, zero value otherwise.
func (o *GenerateObitChecksumResponse) GetChecksum() string {
	if o == nil || o.Checksum == nil {
		var ret string
		return ret
	}
	return *o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateObitChecksumResponse) GetChecksumOk() (*string, bool) {
	if o == nil || o.Checksum == nil {
		return nil, false
	}
	return o.Checksum, true
}

// HasChecksum returns a boolean if a field has been set.
func (o *GenerateObitChecksumResponse) HasChecksum() bool {
	if o != nil && o.Checksum != nil {
		return true
	}

	return false
}

// SetChecksum gets a reference to the given string and assigns it to the Checksum field.
func (o *GenerateObitChecksumResponse) SetChecksum(v string) {
	o.Checksum = &v
}

// GetComputeLog returns the ComputeLog field value if set, zero value otherwise.
func (o *GenerateObitChecksumResponse) GetComputeLog() string {
	if o == nil || o.ComputeLog == nil {
		var ret string
		return ret
	}
	return *o.ComputeLog
}

// GetComputeLogOk returns a tuple with the ComputeLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateObitChecksumResponse) GetComputeLogOk() (*string, bool) {
	if o == nil || o.ComputeLog == nil {
		return nil, false
	}
	return o.ComputeLog, true
}

// HasComputeLog returns a boolean if a field has been set.
func (o *GenerateObitChecksumResponse) HasComputeLog() bool {
	if o != nil && o.ComputeLog != nil {
		return true
	}

	return false
}

// SetComputeLog gets a reference to the given string and assigns it to the ComputeLog field.
func (o *GenerateObitChecksumResponse) SetComputeLog(v string) {
	o.ComputeLog = &v
}

func (o GenerateObitChecksumResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Checksum != nil {
		toSerialize["checksum"] = o.Checksum
	}
	if o.ComputeLog != nil {
		toSerialize["compute_log"] = o.ComputeLog
	}
	return json.Marshal(toSerialize)
}

type NullableGenerateObitChecksumResponse struct {
	value *GenerateObitChecksumResponse
	isSet bool
}

func (v NullableGenerateObitChecksumResponse) Get() *GenerateObitChecksumResponse {
	return v.value
}

func (v *NullableGenerateObitChecksumResponse) Set(val *GenerateObitChecksumResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateObitChecksumResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateObitChecksumResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateObitChecksumResponse(val *GenerateObitChecksumResponse) *NullableGenerateObitChecksumResponse {
	return &NullableGenerateObitChecksumResponse{value: val, isSet: true}
}

func (v NullableGenerateObitChecksumResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateObitChecksumResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


