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
	"time"
)

// ObitHistory New Obit request payload.
type ObitHistory struct {
	// OBADA decentralized identifier (max length Rohi?)
	ObitDid *string `json:"obit_did,omitempty"`
	// History event
	Event *string `json:"event,omitempty"`
	OldValues map[string]interface{} `json:"old_values,omitempty"`
	NewValues map[string]interface{} `json:"new_values,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewObitHistory instantiates a new ObitHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObitHistory() *ObitHistory {
	this := ObitHistory{}
	return &this
}

// NewObitHistoryWithDefaults instantiates a new ObitHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObitHistoryWithDefaults() *ObitHistory {
	this := ObitHistory{}
	return &this
}

// GetObitDid returns the ObitDid field value if set, zero value otherwise.
func (o *ObitHistory) GetObitDid() string {
	if o == nil || o.ObitDid == nil {
		var ret string
		return ret
	}
	return *o.ObitDid
}

// GetObitDidOk returns a tuple with the ObitDid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObitHistory) GetObitDidOk() (*string, bool) {
	if o == nil || o.ObitDid == nil {
		return nil, false
	}
	return o.ObitDid, true
}

// HasObitDid returns a boolean if a field has been set.
func (o *ObitHistory) HasObitDid() bool {
	if o != nil && o.ObitDid != nil {
		return true
	}

	return false
}

// SetObitDid gets a reference to the given string and assigns it to the ObitDid field.
func (o *ObitHistory) SetObitDid(v string) {
	o.ObitDid = &v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *ObitHistory) GetEvent() string {
	if o == nil || o.Event == nil {
		var ret string
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObitHistory) GetEventOk() (*string, bool) {
	if o == nil || o.Event == nil {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *ObitHistory) HasEvent() bool {
	if o != nil && o.Event != nil {
		return true
	}

	return false
}

// SetEvent gets a reference to the given string and assigns it to the Event field.
func (o *ObitHistory) SetEvent(v string) {
	o.Event = &v
}

// GetOldValues returns the OldValues field value if set, zero value otherwise.
func (o *ObitHistory) GetOldValues() map[string]interface{} {
	if o == nil || o.OldValues == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.OldValues
}

// GetOldValuesOk returns a tuple with the OldValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObitHistory) GetOldValuesOk() (map[string]interface{}, bool) {
	if o == nil || o.OldValues == nil {
		return nil, false
	}
	return o.OldValues, true
}

// HasOldValues returns a boolean if a field has been set.
func (o *ObitHistory) HasOldValues() bool {
	if o != nil && o.OldValues != nil {
		return true
	}

	return false
}

// SetOldValues gets a reference to the given map[string]interface{} and assigns it to the OldValues field.
func (o *ObitHistory) SetOldValues(v map[string]interface{}) {
	o.OldValues = v
}

// GetNewValues returns the NewValues field value if set, zero value otherwise.
func (o *ObitHistory) GetNewValues() map[string]interface{} {
	if o == nil || o.NewValues == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.NewValues
}

// GetNewValuesOk returns a tuple with the NewValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObitHistory) GetNewValuesOk() (map[string]interface{}, bool) {
	if o == nil || o.NewValues == nil {
		return nil, false
	}
	return o.NewValues, true
}

// HasNewValues returns a boolean if a field has been set.
func (o *ObitHistory) HasNewValues() bool {
	if o != nil && o.NewValues != nil {
		return true
	}

	return false
}

// SetNewValues gets a reference to the given map[string]interface{} and assigns it to the NewValues field.
func (o *ObitHistory) SetNewValues(v map[string]interface{}) {
	o.NewValues = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ObitHistory) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObitHistory) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ObitHistory) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ObitHistory) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ObitHistory) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObitHistory) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ObitHistory) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ObitHistory) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o ObitHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObitDid != nil {
		toSerialize["obit_did"] = o.ObitDid
	}
	if o.Event != nil {
		toSerialize["event"] = o.Event
	}
	if o.OldValues != nil {
		toSerialize["old_values"] = o.OldValues
	}
	if o.NewValues != nil {
		toSerialize["new_values"] = o.NewValues
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableObitHistory struct {
	value *ObitHistory
	isSet bool
}

func (v NullableObitHistory) Get() *ObitHistory {
	return v.value
}

func (v *NullableObitHistory) Set(val *ObitHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableObitHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableObitHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObitHistory(val *ObitHistory) *NullableObitHistory {
	return &NullableObitHistory{value: val, isSet: true}
}

func (v NullableObitHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObitHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


