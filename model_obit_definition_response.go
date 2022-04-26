/*
OBADA Client Helper API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
Contact: techops@obada.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github.com/obada-foundation/client-api-library-go

import (
	"encoding/json"
)

// ObitDefinitionResponse struct for ObitDefinitionResponse
type ObitDefinitionResponse struct {
	Status *int32 `json:"status,omitempty"`
	Obit *ObitDefinition `json:"obit,omitempty"`
}

// NewObitDefinitionResponse instantiates a new ObitDefinitionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObitDefinitionResponse() *ObitDefinitionResponse {
	this := ObitDefinitionResponse{}
	return &this
}

// NewObitDefinitionResponseWithDefaults instantiates a new ObitDefinitionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObitDefinitionResponseWithDefaults() *ObitDefinitionResponse {
	this := ObitDefinitionResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ObitDefinitionResponse) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObitDefinitionResponse) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ObitDefinitionResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *ObitDefinitionResponse) SetStatus(v int32) {
	o.Status = &v
}

// GetObit returns the Obit field value if set, zero value otherwise.
func (o *ObitDefinitionResponse) GetObit() ObitDefinition {
	if o == nil || o.Obit == nil {
		var ret ObitDefinition
		return ret
	}
	return *o.Obit
}

// GetObitOk returns a tuple with the Obit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObitDefinitionResponse) GetObitOk() (*ObitDefinition, bool) {
	if o == nil || o.Obit == nil {
		return nil, false
	}
	return o.Obit, true
}

// HasObit returns a boolean if a field has been set.
func (o *ObitDefinitionResponse) HasObit() bool {
	if o != nil && o.Obit != nil {
		return true
	}

	return false
}

// SetObit gets a reference to the given ObitDefinition and assigns it to the Obit field.
func (o *ObitDefinitionResponse) SetObit(v ObitDefinition) {
	o.Obit = &v
}

func (o ObitDefinitionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Obit != nil {
		toSerialize["obit"] = o.Obit
	}
	return json.Marshal(toSerialize)
}

type NullableObitDefinitionResponse struct {
	value *ObitDefinitionResponse
	isSet bool
}

func (v NullableObitDefinitionResponse) Get() *ObitDefinitionResponse {
	return v.value
}

func (v *NullableObitDefinitionResponse) Set(val *ObitDefinitionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableObitDefinitionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableObitDefinitionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObitDefinitionResponse(val *ObitDefinitionResponse) *NullableObitDefinitionResponse {
	return &NullableObitDefinitionResponse{value: val, isSet: true}
}

func (v NullableObitDefinitionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObitDefinitionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


