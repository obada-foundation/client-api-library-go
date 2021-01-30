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

// ClientObitResponse struct for ClientObitResponse
type ClientObitResponse struct {
	Status *int32 `json:"status,omitempty"`
	Obit *ClientObit `json:"obit,omitempty"`
}

// NewClientObitResponse instantiates a new ClientObitResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientObitResponse() *ClientObitResponse {
	this := ClientObitResponse{}
	return &this
}

// NewClientObitResponseWithDefaults instantiates a new ClientObitResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientObitResponseWithDefaults() *ClientObitResponse {
	this := ClientObitResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ClientObitResponse) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientObitResponse) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ClientObitResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *ClientObitResponse) SetStatus(v int32) {
	o.Status = &v
}

// GetObit returns the Obit field value if set, zero value otherwise.
func (o *ClientObitResponse) GetObit() ClientObit {
	if o == nil || o.Obit == nil {
		var ret ClientObit
		return ret
	}
	return *o.Obit
}

// GetObitOk returns a tuple with the Obit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientObitResponse) GetObitOk() (*ClientObit, bool) {
	if o == nil || o.Obit == nil {
		return nil, false
	}
	return o.Obit, true
}

// HasObit returns a boolean if a field has been set.
func (o *ClientObitResponse) HasObit() bool {
	if o != nil && o.Obit != nil {
		return true
	}

	return false
}

// SetObit gets a reference to the given ClientObit and assigns it to the Obit field.
func (o *ClientObitResponse) SetObit(v ClientObit) {
	o.Obit = &v
}

func (o ClientObitResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Obit != nil {
		toSerialize["obit"] = o.Obit
	}
	return json.Marshal(toSerialize)
}

type NullableClientObitResponse struct {
	value *ClientObitResponse
	isSet bool
}

func (v NullableClientObitResponse) Get() *ClientObitResponse {
	return v.value
}

func (v *NullableClientObitResponse) Set(val *ClientObitResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableClientObitResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableClientObitResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientObitResponse(val *ClientObitResponse) *NullableClientObitResponse {
	return &NullableClientObitResponse{value: val, isSet: true}
}

func (v NullableClientObitResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientObitResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


