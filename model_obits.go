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

// Obits Obits search response
type Obits struct {
	Data []Obit `json:"data,omitempty"`
	Meta *ObitsMeta `json:"meta,omitempty"`
}

// NewObits instantiates a new Obits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObits() *Obits {
	this := Obits{}
	return &this
}

// NewObitsWithDefaults instantiates a new Obits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObitsWithDefaults() *Obits {
	this := Obits{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Obits) GetData() []Obit {
	if o == nil || o.Data == nil {
		var ret []Obit
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Obits) GetDataOk() ([]Obit, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Obits) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Obit and assigns it to the Data field.
func (o *Obits) SetData(v []Obit) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Obits) GetMeta() ObitsMeta {
	if o == nil || o.Meta == nil {
		var ret ObitsMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Obits) GetMetaOk() (*ObitsMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Obits) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ObitsMeta and assigns it to the Meta field.
func (o *Obits) SetMeta(v ObitsMeta) {
	o.Meta = &v
}

func (o Obits) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableObits struct {
	value *Obits
	isSet bool
}

func (v NullableObits) Get() *Obits {
	return v.value
}

func (v *NullableObits) Set(val *Obits) {
	v.value = val
	v.isSet = true
}

func (v NullableObits) IsSet() bool {
	return v.isSet
}

func (v *NullableObits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObits(val *Obits) *NullableObits {
	return &NullableObits{value: val, isSet: true}
}

func (v NullableObits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


