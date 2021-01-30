# ClientObit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ObitDid** | Pointer to **string** |  | [optional] 
**Usn** | Pointer to **string** |  | [optional] 
**OwnerDid** | Pointer to **string** |  | [optional] 
**ObitStatus** | Pointer to **string** | Represent available Obit statuses:   - FUNCTIONAL   - NON_FUNCTIONAL   - DISPOSED   - STOLEN   - DISABLED_BY_OWNER  | [optional] 
**Manufacturer** | Pointer to **string** |  | [optional] 
**PartNumber** | Pointer to **string** |  | [optional] 
**SerialNumberHash** | Pointer to **string** |  | [optional] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] 
**RootHash** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**[]LocalObitMetadata**](LocalObitMetadata.md) | Get description from Rohi | [optional] 
**Documents** | Pointer to [**[]LocalObitDocuments**](LocalObitDocuments.md) | To generate this link, take an SHA-256 hash of the document, and link to it as https://www.some-website.com?h1&#x3D;hash-of-document. Note this does not yet adhere to the hashlink standard.  | [optional] 
**StructuredData** | Pointer to [**[]LocalObitStructuredData**](LocalObitStructuredData.md) | Same as metadata but bigger. Key (string) &#x3D;&gt; Value (string) (hash per line sha256(key + value)) | [optional] 

## Methods

### NewClientObit

`func NewClientObit() *ClientObit`

NewClientObit instantiates a new ClientObit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientObitWithDefaults

`func NewClientObitWithDefaults() *ClientObit`

NewClientObitWithDefaults instantiates a new ClientObit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClientObit) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientObit) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientObit) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ClientObit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObitDid

`func (o *ClientObit) GetObitDid() string`

GetObitDid returns the ObitDid field if non-nil, zero value otherwise.

### GetObitDidOk

`func (o *ClientObit) GetObitDidOk() (*string, bool)`

GetObitDidOk returns a tuple with the ObitDid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObitDid

`func (o *ClientObit) SetObitDid(v string)`

SetObitDid sets ObitDid field to given value.

### HasObitDid

`func (o *ClientObit) HasObitDid() bool`

HasObitDid returns a boolean if a field has been set.

### GetUsn

`func (o *ClientObit) GetUsn() string`

GetUsn returns the Usn field if non-nil, zero value otherwise.

### GetUsnOk

`func (o *ClientObit) GetUsnOk() (*string, bool)`

GetUsnOk returns a tuple with the Usn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsn

`func (o *ClientObit) SetUsn(v string)`

SetUsn sets Usn field to given value.

### HasUsn

`func (o *ClientObit) HasUsn() bool`

HasUsn returns a boolean if a field has been set.

### GetOwnerDid

`func (o *ClientObit) GetOwnerDid() string`

GetOwnerDid returns the OwnerDid field if non-nil, zero value otherwise.

### GetOwnerDidOk

`func (o *ClientObit) GetOwnerDidOk() (*string, bool)`

GetOwnerDidOk returns a tuple with the OwnerDid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerDid

`func (o *ClientObit) SetOwnerDid(v string)`

SetOwnerDid sets OwnerDid field to given value.

### HasOwnerDid

`func (o *ClientObit) HasOwnerDid() bool`

HasOwnerDid returns a boolean if a field has been set.

### GetObitStatus

`func (o *ClientObit) GetObitStatus() string`

GetObitStatus returns the ObitStatus field if non-nil, zero value otherwise.

### GetObitStatusOk

`func (o *ClientObit) GetObitStatusOk() (*string, bool)`

GetObitStatusOk returns a tuple with the ObitStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObitStatus

`func (o *ClientObit) SetObitStatus(v string)`

SetObitStatus sets ObitStatus field to given value.

### HasObitStatus

`func (o *ClientObit) HasObitStatus() bool`

HasObitStatus returns a boolean if a field has been set.

### GetManufacturer

`func (o *ClientObit) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *ClientObit) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *ClientObit) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *ClientObit) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetPartNumber

`func (o *ClientObit) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *ClientObit) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *ClientObit) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *ClientObit) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetSerialNumberHash

`func (o *ClientObit) GetSerialNumberHash() string`

GetSerialNumberHash returns the SerialNumberHash field if non-nil, zero value otherwise.

### GetSerialNumberHashOk

`func (o *ClientObit) GetSerialNumberHashOk() (*string, bool)`

GetSerialNumberHashOk returns a tuple with the SerialNumberHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumberHash

`func (o *ClientObit) SetSerialNumberHash(v string)`

SetSerialNumberHash sets SerialNumberHash field to given value.

### HasSerialNumberHash

`func (o *ClientObit) HasSerialNumberHash() bool`

HasSerialNumberHash returns a boolean if a field has been set.

### GetModifiedAt

`func (o *ClientObit) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ClientObit) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ClientObit) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *ClientObit) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetRootHash

`func (o *ClientObit) GetRootHash() string`

GetRootHash returns the RootHash field if non-nil, zero value otherwise.

### GetRootHashOk

`func (o *ClientObit) GetRootHashOk() (*string, bool)`

GetRootHashOk returns a tuple with the RootHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootHash

`func (o *ClientObit) SetRootHash(v string)`

SetRootHash sets RootHash field to given value.

### HasRootHash

`func (o *ClientObit) HasRootHash() bool`

HasRootHash returns a boolean if a field has been set.

### GetMetadata

`func (o *ClientObit) GetMetadata() []LocalObitMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ClientObit) GetMetadataOk() (*[]LocalObitMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ClientObit) SetMetadata(v []LocalObitMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ClientObit) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDocuments

`func (o *ClientObit) GetDocuments() []LocalObitDocuments`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *ClientObit) GetDocumentsOk() (*[]LocalObitDocuments, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *ClientObit) SetDocuments(v []LocalObitDocuments)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *ClientObit) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetStructuredData

`func (o *ClientObit) GetStructuredData() []LocalObitStructuredData`

GetStructuredData returns the StructuredData field if non-nil, zero value otherwise.

### GetStructuredDataOk

`func (o *ClientObit) GetStructuredDataOk() (*[]LocalObitStructuredData, bool)`

GetStructuredDataOk returns a tuple with the StructuredData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredData

`func (o *ClientObit) SetStructuredData(v []LocalObitStructuredData)`

SetStructuredData sets StructuredData field to given value.

### HasStructuredData

`func (o *ClientObit) HasStructuredData() bool`

HasStructuredData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


