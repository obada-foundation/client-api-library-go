# LocalObit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | **string** | Owner is the person/entity that owns the obit and the physical asset it represents. | 
**ObitStatus** | **string** | Represent available Obit statuses:   - FUNCTIONAL   - NON_FUNCTIONAL   - DISPOSED   - STOLEN   - DISABLED_BY_OWNER  | 
**Manufacturer** | **string** | Waiting more specific details from Rohi | 
**PartNumber** | **string** | Manufacturer provided. In cases where no part number is provided for the product, use model, or the most specific ID available from the manufacturer. MWCN2LL/A (an iPhone 11 Pro, Silver, 256GB, model A2160) | 
**SerialNumber** | **string** | Serial Number | 
**Metadata** | Pointer to [**[]LocalObitMetadata**](LocalObitMetadata.md) | Get description from Rohi | [optional] 
**Documents** | Pointer to [**[]LocalObitDocuments**](LocalObitDocuments.md) | To generate this link, take an SHA-256 hash of the document, and link to it as https://www.some-website.com?h1&#x3D;hash-of-document. Note this does not yet adhere to the hashlink standard.  | [optional] 
**StructuredData** | Pointer to [**[]LocalObitStructuredData**](LocalObitStructuredData.md) | Same as metadata but bigger. Key (string) &#x3D;&gt; Value (string) (hash per line sha256(key + value)) | [optional] 
**ModifiedAt** | **time.Time** |  | 

## Methods

### NewLocalObit

`func NewLocalObit(owner string, obitStatus string, manufacturer string, partNumber string, serialNumber string, modifiedAt time.Time, ) *LocalObit`

NewLocalObit instantiates a new LocalObit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalObitWithDefaults

`func NewLocalObitWithDefaults() *LocalObit`

NewLocalObitWithDefaults instantiates a new LocalObit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *LocalObit) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *LocalObit) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *LocalObit) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetObitStatus

`func (o *LocalObit) GetObitStatus() string`

GetObitStatus returns the ObitStatus field if non-nil, zero value otherwise.

### GetObitStatusOk

`func (o *LocalObit) GetObitStatusOk() (*string, bool)`

GetObitStatusOk returns a tuple with the ObitStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObitStatus

`func (o *LocalObit) SetObitStatus(v string)`

SetObitStatus sets ObitStatus field to given value.


### GetManufacturer

`func (o *LocalObit) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *LocalObit) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *LocalObit) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.


### GetPartNumber

`func (o *LocalObit) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *LocalObit) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *LocalObit) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.


### GetSerialNumber

`func (o *LocalObit) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *LocalObit) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *LocalObit) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### GetMetadata

`func (o *LocalObit) GetMetadata() []LocalObitMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *LocalObit) GetMetadataOk() (*[]LocalObitMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *LocalObit) SetMetadata(v []LocalObitMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *LocalObit) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDocuments

`func (o *LocalObit) GetDocuments() []LocalObitDocuments`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *LocalObit) GetDocumentsOk() (*[]LocalObitDocuments, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *LocalObit) SetDocuments(v []LocalObitDocuments)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *LocalObit) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetStructuredData

`func (o *LocalObit) GetStructuredData() []LocalObitStructuredData`

GetStructuredData returns the StructuredData field if non-nil, zero value otherwise.

### GetStructuredDataOk

`func (o *LocalObit) GetStructuredDataOk() (*[]LocalObitStructuredData, bool)`

GetStructuredDataOk returns a tuple with the StructuredData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredData

`func (o *LocalObit) SetStructuredData(v []LocalObitStructuredData)`

SetStructuredData sets StructuredData field to given value.

### HasStructuredData

`func (o *LocalObit) HasStructuredData() bool`

HasStructuredData returns a boolean if a field has been set.

### GetModifiedAt

`func (o *LocalObit) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *LocalObit) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *LocalObit) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


