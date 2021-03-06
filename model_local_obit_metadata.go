/*
 * OBADA Client Helper API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Contact: techops@obada.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github.com/obada-foundation/client-api-library-go

import (
	"encoding/json"
)

// LocalObitMetadata A record that represent key/value metadata record
type LocalObitMetadata struct {
	// Metadata key
	Key string `json:"key"`
	// Matadata value
	Value string `json:"value"`
}

// NewLocalObitMetadata instantiates a new LocalObitMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocalObitMetadata(key string, value string) *LocalObitMetadata {
	this := LocalObitMetadata{}
	this.Key = key
	this.Value = value
	return &this
}

// NewLocalObitMetadataWithDefaults instantiates a new LocalObitMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocalObitMetadataWithDefaults() *LocalObitMetadata {
	this := LocalObitMetadata{}
	return &this
}

// GetKey returns the Key field value
func (o *LocalObitMetadata) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *LocalObitMetadata) GetKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *LocalObitMetadata) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *LocalObitMetadata) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *LocalObitMetadata) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *LocalObitMetadata) SetValue(v string) {
	o.Value = v
}

func (o LocalObitMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableLocalObitMetadata struct {
	value *LocalObitMetadata
	isSet bool
}

func (v NullableLocalObitMetadata) Get() *LocalObitMetadata {
	return v.value
}

func (v *NullableLocalObitMetadata) Set(val *LocalObitMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalObitMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalObitMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalObitMetadata(val *LocalObitMetadata) *NullableLocalObitMetadata {
	return &NullableLocalObitMetadata{value: val, isSet: true}
}

func (v NullableLocalObitMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalObitMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


