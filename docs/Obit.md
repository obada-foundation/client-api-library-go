# Obit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Did** | Pointer to **string** | OBADA decentralized identifier | [optional] 
**Usn** | Pointer to **string** | Test An 8-12 character “URL Shortened” obit. Create the USN by Base(58) encoding the OBIT ID. Take the first 8 characters. Check for duplicates. With 16 million OBITs there is a 40% chance of collision. If so, add 4 characters. Repeat as necessary until a unique USN is generated.  | [optional] 
**Owner** | Pointer to **string** | Owner is the person/entity that owns the obit and the physical asset it represents. Format is a DID like did:obada:owner:1234. However in the current version only test numbers will be used. | [optional] 
**Manufacturer** | **string** | Waiting more specific details from Rohi | 
**PartNumber** | **string** | Manufacturer provided. In cases where no part number is provided for the product, use model, or the most specific ID available from the manufacturer. MWCN2LL/A (an iPhone 11 Pro, Silver, 256GB, model A2160) | 
**SerialNumberHash** | Pointer to **string** | Serial number hashed with sha256 hash function | [optional] 
**TrustAnchorToken** | Pointer to **string** | JWT token from the trust anchor | [optional] 
**Documents** | Pointer to [**[]Document**](Document.md) | Documents that are attached to Obit | [optional] 
**Checksum** | Pointer to **string** | Hash calculated by SHA256 (previous Obit checksum + Obit data).  | [optional] 

## Methods

### NewObit

`func NewObit(manufacturer string, partNumber string, ) *Obit`

NewObit instantiates a new Obit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObitWithDefaults

`func NewObitWithDefaults() *Obit`

NewObitWithDefaults instantiates a new Obit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDid

`func (o *Obit) GetDid() string`

GetDid returns the Did field if non-nil, zero value otherwise.

### GetDidOk

`func (o *Obit) GetDidOk() (*string, bool)`

GetDidOk returns a tuple with the Did field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDid

`func (o *Obit) SetDid(v string)`

SetDid sets Did field to given value.

### HasDid

`func (o *Obit) HasDid() bool`

HasDid returns a boolean if a field has been set.

### GetUsn

`func (o *Obit) GetUsn() string`

GetUsn returns the Usn field if non-nil, zero value otherwise.

### GetUsnOk

`func (o *Obit) GetUsnOk() (*string, bool)`

GetUsnOk returns a tuple with the Usn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsn

`func (o *Obit) SetUsn(v string)`

SetUsn sets Usn field to given value.

### HasUsn

`func (o *Obit) HasUsn() bool`

HasUsn returns a boolean if a field has been set.

### GetOwner

`func (o *Obit) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Obit) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Obit) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Obit) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetManufacturer

`func (o *Obit) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *Obit) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *Obit) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.


### GetPartNumber

`func (o *Obit) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *Obit) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *Obit) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.


### GetSerialNumberHash

`func (o *Obit) GetSerialNumberHash() string`

GetSerialNumberHash returns the SerialNumberHash field if non-nil, zero value otherwise.

### GetSerialNumberHashOk

`func (o *Obit) GetSerialNumberHashOk() (*string, bool)`

GetSerialNumberHashOk returns a tuple with the SerialNumberHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumberHash

`func (o *Obit) SetSerialNumberHash(v string)`

SetSerialNumberHash sets SerialNumberHash field to given value.

### HasSerialNumberHash

`func (o *Obit) HasSerialNumberHash() bool`

HasSerialNumberHash returns a boolean if a field has been set.

### GetTrustAnchorToken

`func (o *Obit) GetTrustAnchorToken() string`

GetTrustAnchorToken returns the TrustAnchorToken field if non-nil, zero value otherwise.

### GetTrustAnchorTokenOk

`func (o *Obit) GetTrustAnchorTokenOk() (*string, bool)`

GetTrustAnchorTokenOk returns a tuple with the TrustAnchorToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustAnchorToken

`func (o *Obit) SetTrustAnchorToken(v string)`

SetTrustAnchorToken sets TrustAnchorToken field to given value.

### HasTrustAnchorToken

`func (o *Obit) HasTrustAnchorToken() bool`

HasTrustAnchorToken returns a boolean if a field has been set.

### GetDocuments

`func (o *Obit) GetDocuments() []Document`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *Obit) GetDocumentsOk() (*[]Document, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *Obit) SetDocuments(v []Document)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *Obit) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetChecksum

`func (o *Obit) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *Obit) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *Obit) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *Obit) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


