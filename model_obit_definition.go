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

// ObitDefinition struct for ObitDefinition
type ObitDefinition struct {
	SerialHash *string `json:"serial_hash,omitempty"`
	Usn *string `json:"usn,omitempty"`
	ObitDid *string `json:"obit_did,omitempty"`
	UsnBase58 *string `json:"usn_base58,omitempty"`
}

// NewObitDefinition instantiates a new ObitDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObitDefinition() *ObitDefinition {
	this := ObitDefinition{}
	return &this
}

// NewObitDefinitionWithDefaults instantiates a new ObitDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObitDefinitionWithDefaults() *ObitDefinition {
	this := ObitDefinition{}
	return &this
}

// GetSerialHash returns the SerialHash field value if set, zero value otherwise.
func (o *ObitDefinition) GetSerialHash() string {
	if o == nil || o.SerialHash == nil {
		var ret string
		return ret
	}
	return *o.SerialHash
}

// GetSerialHashOk returns a tuple with the SerialHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObitDefinition) GetSerialHashOk() (*string, bool) {
	if o == nil || o.SerialHash == nil {
		return nil, false
	}
	return o.SerialHash, true
}

// HasSerialHash returns a boolean if a field has been set.
func (o *ObitDefinition) HasSerialHash() bool {
	if o != nil && o.SerialHash != nil {
		return true
	}

	return false
}

// SetSerialHash gets a reference to the given string and assigns it to the SerialHash field.
func (o *ObitDefinition) SetSerialHash(v string) {
	o.SerialHash = &v
}

// GetUsn returns the Usn field value if set, zero value otherwise.
func (o *ObitDefinition) GetUsn() string {
	if o == nil || o.Usn == nil {
		var ret string
		return ret
	}
	return *o.Usn
}

// GetUsnOk returns a tuple with the Usn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObitDefinition) GetUsnOk() (*string, bool) {
	if o == nil || o.Usn == nil {
		return nil, false
	}
	return o.Usn, true
}

// HasUsn returns a boolean if a field has been set.
func (o *ObitDefinition) HasUsn() bool {
	if o != nil && o.Usn != nil {
		return true
	}

	return false
}

// SetUsn gets a reference to the given string and assigns it to the Usn field.
func (o *ObitDefinition) SetUsn(v string) {
	o.Usn = &v
}

// GetObitDid returns the ObitDid field value if set, zero value otherwise.
func (o *ObitDefinition) GetObitDid() string {
	if o == nil || o.ObitDid == nil {
		var ret string
		return ret
	}
	return *o.ObitDid
}

// GetObitDidOk returns a tuple with the ObitDid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObitDefinition) GetObitDidOk() (*string, bool) {
	if o == nil || o.ObitDid == nil {
		return nil, false
	}
	return o.ObitDid, true
}

// HasObitDid returns a boolean if a field has been set.
func (o *ObitDefinition) HasObitDid() bool {
	if o != nil && o.ObitDid != nil {
		return true
	}

	return false
}

// SetObitDid gets a reference to the given string and assigns it to the ObitDid field.
func (o *ObitDefinition) SetObitDid(v string) {
	o.ObitDid = &v
}

// GetUsnBase58 returns the UsnBase58 field value if set, zero value otherwise.
func (o *ObitDefinition) GetUsnBase58() string {
	if o == nil || o.UsnBase58 == nil {
		var ret string
		return ret
	}
	return *o.UsnBase58
}

// GetUsnBase58Ok returns a tuple with the UsnBase58 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObitDefinition) GetUsnBase58Ok() (*string, bool) {
	if o == nil || o.UsnBase58 == nil {
		return nil, false
	}
	return o.UsnBase58, true
}

// HasUsnBase58 returns a boolean if a field has been set.
func (o *ObitDefinition) HasUsnBase58() bool {
	if o != nil && o.UsnBase58 != nil {
		return true
	}

	return false
}

// SetUsnBase58 gets a reference to the given string and assigns it to the UsnBase58 field.
func (o *ObitDefinition) SetUsnBase58(v string) {
	o.UsnBase58 = &v
}

func (o ObitDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SerialHash != nil {
		toSerialize["serial_hash"] = o.SerialHash
	}
	if o.Usn != nil {
		toSerialize["usn"] = o.Usn
	}
	if o.ObitDid != nil {
		toSerialize["obit_did"] = o.ObitDid
	}
	if o.UsnBase58 != nil {
		toSerialize["usn_base58"] = o.UsnBase58
	}
	return json.Marshal(toSerialize)
}

type NullableObitDefinition struct {
	value *ObitDefinition
	isSet bool
}

func (v NullableObitDefinition) Get() *ObitDefinition {
	return v.value
}

func (v *NullableObitDefinition) Set(val *ObitDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableObitDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableObitDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObitDefinition(val *ObitDefinition) *NullableObitDefinition {
	return &NullableObitDefinition{value: val, isSet: true}
}

func (v NullableObitDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObitDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


