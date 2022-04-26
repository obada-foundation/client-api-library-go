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
	"time"
)

// LocalObit struct for LocalObit
type LocalObit struct {
	// Owner is the person/entity that owns the obit and the physical asset it represents.
	Owner string `json:"owner"`
	// Represent available Obit statuses:   - FUNCTIONAL   - NON_FUNCTIONAL   - DISPOSED   - STOLEN   - DISABLED_BY_OWNER 
	ObitStatus string `json:"obit_status"`
	// Waiting more specific details from Rohi
	Manufacturer string `json:"manufacturer"`
	// Manufacturer provided. In cases where no part number is provided for the product, use model, or the most specific ID available from the manufacturer. MWCN2LL/A (an iPhone 11 Pro, Silver, 256GB, model A2160)
	PartNumber string `json:"part_number"`
	// Serial Number
	SerialNumber string `json:"serial_number"`
	// Get description from Rohi
	Metadata []LocalObitMetadataInner `json:"metadata,omitempty"`
	// To generate this link, take an SHA-256 hash of the document, and link to it as https://www.some-website.com?h1=hash-of-document. Note this does not yet adhere to the hashlink standard. 
	Documents []LocalObitDocumentsInner `json:"documents,omitempty"`
	// Same as metadata but bigger. Key (string) => Value (string) (hash per line sha256(key + value))
	StructuredData []LocalObitStructuredDataInner `json:"structured_data,omitempty"`
	ModifiedAt time.Time `json:"modified_at"`
}

// NewLocalObit instantiates a new LocalObit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocalObit(owner string, obitStatus string, manufacturer string, partNumber string, serialNumber string, modifiedAt time.Time) *LocalObit {
	this := LocalObit{}
	this.Owner = owner
	this.ObitStatus = obitStatus
	this.Manufacturer = manufacturer
	this.PartNumber = partNumber
	this.SerialNumber = serialNumber
	this.ModifiedAt = modifiedAt
	return &this
}

// NewLocalObitWithDefaults instantiates a new LocalObit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocalObitWithDefaults() *LocalObit {
	this := LocalObit{}
	return &this
}

// GetOwner returns the Owner field value
func (o *LocalObit) GetOwner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *LocalObit) GetOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *LocalObit) SetOwner(v string) {
	o.Owner = v
}

// GetObitStatus returns the ObitStatus field value
func (o *LocalObit) GetObitStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObitStatus
}

// GetObitStatusOk returns a tuple with the ObitStatus field value
// and a boolean to check if the value has been set.
func (o *LocalObit) GetObitStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObitStatus, true
}

// SetObitStatus sets field value
func (o *LocalObit) SetObitStatus(v string) {
	o.ObitStatus = v
}

// GetManufacturer returns the Manufacturer field value
func (o *LocalObit) GetManufacturer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value
// and a boolean to check if the value has been set.
func (o *LocalObit) GetManufacturerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Manufacturer, true
}

// SetManufacturer sets field value
func (o *LocalObit) SetManufacturer(v string) {
	o.Manufacturer = v
}

// GetPartNumber returns the PartNumber field value
func (o *LocalObit) GetPartNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value
// and a boolean to check if the value has been set.
func (o *LocalObit) GetPartNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartNumber, true
}

// SetPartNumber sets field value
func (o *LocalObit) SetPartNumber(v string) {
	o.PartNumber = v
}

// GetSerialNumber returns the SerialNumber field value
func (o *LocalObit) GetSerialNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value
// and a boolean to check if the value has been set.
func (o *LocalObit) GetSerialNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SerialNumber, true
}

// SetSerialNumber sets field value
func (o *LocalObit) SetSerialNumber(v string) {
	o.SerialNumber = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *LocalObit) GetMetadata() []LocalObitMetadataInner {
	if o == nil || o.Metadata == nil {
		var ret []LocalObitMetadataInner
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalObit) GetMetadataOk() ([]LocalObitMetadataInner, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *LocalObit) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []LocalObitMetadataInner and assigns it to the Metadata field.
func (o *LocalObit) SetMetadata(v []LocalObitMetadataInner) {
	o.Metadata = v
}

// GetDocuments returns the Documents field value if set, zero value otherwise.
func (o *LocalObit) GetDocuments() []LocalObitDocumentsInner {
	if o == nil || o.Documents == nil {
		var ret []LocalObitDocumentsInner
		return ret
	}
	return o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalObit) GetDocumentsOk() ([]LocalObitDocumentsInner, bool) {
	if o == nil || o.Documents == nil {
		return nil, false
	}
	return o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *LocalObit) HasDocuments() bool {
	if o != nil && o.Documents != nil {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given []LocalObitDocumentsInner and assigns it to the Documents field.
func (o *LocalObit) SetDocuments(v []LocalObitDocumentsInner) {
	o.Documents = v
}

// GetStructuredData returns the StructuredData field value if set, zero value otherwise.
func (o *LocalObit) GetStructuredData() []LocalObitStructuredDataInner {
	if o == nil || o.StructuredData == nil {
		var ret []LocalObitStructuredDataInner
		return ret
	}
	return o.StructuredData
}

// GetStructuredDataOk returns a tuple with the StructuredData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalObit) GetStructuredDataOk() ([]LocalObitStructuredDataInner, bool) {
	if o == nil || o.StructuredData == nil {
		return nil, false
	}
	return o.StructuredData, true
}

// HasStructuredData returns a boolean if a field has been set.
func (o *LocalObit) HasStructuredData() bool {
	if o != nil && o.StructuredData != nil {
		return true
	}

	return false
}

// SetStructuredData gets a reference to the given []LocalObitStructuredDataInner and assigns it to the StructuredData field.
func (o *LocalObit) SetStructuredData(v []LocalObitStructuredDataInner) {
	o.StructuredData = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *LocalObit) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *LocalObit) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *LocalObit) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

func (o LocalObit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["owner"] = o.Owner
	}
	if true {
		toSerialize["obit_status"] = o.ObitStatus
	}
	if true {
		toSerialize["manufacturer"] = o.Manufacturer
	}
	if true {
		toSerialize["part_number"] = o.PartNumber
	}
	if true {
		toSerialize["serial_number"] = o.SerialNumber
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Documents != nil {
		toSerialize["documents"] = o.Documents
	}
	if o.StructuredData != nil {
		toSerialize["structured_data"] = o.StructuredData
	}
	if true {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	return json.Marshal(toSerialize)
}

type NullableLocalObit struct {
	value *LocalObit
	isSet bool
}

func (v NullableLocalObit) Get() *LocalObit {
	return v.value
}

func (v *NullableLocalObit) Set(val *LocalObit) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalObit) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalObit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalObit(val *LocalObit) *NullableLocalObit {
	return &NullableLocalObit{value: val, isSet: true}
}

func (v NullableLocalObit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalObit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


