# GenerateObitChecksumRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Manufacturer** | **string** | Manufacturer (Required) | 
**PartNumber** | **string** | Part Number (Required) | 
**SerialNumber** | **string** | Serial Number (Required) | 
**MetadataUri** | Pointer to **string** | Metadata URI | [optional] 
**MetadataUriHash** | Pointer to **string** | Metadata URI Hash | [optional] 
**TrustAnchorToken** | Pointer to **string** | Trust Token | [optional] 

## Methods

### NewGenerateObitChecksumRequest

`func NewGenerateObitChecksumRequest(manufacturer string, partNumber string, serialNumber string, ) *GenerateObitChecksumRequest`

NewGenerateObitChecksumRequest instantiates a new GenerateObitChecksumRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateObitChecksumRequestWithDefaults

`func NewGenerateObitChecksumRequestWithDefaults() *GenerateObitChecksumRequest`

NewGenerateObitChecksumRequestWithDefaults instantiates a new GenerateObitChecksumRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManufacturer

`func (o *GenerateObitChecksumRequest) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *GenerateObitChecksumRequest) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *GenerateObitChecksumRequest) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.


### GetPartNumber

`func (o *GenerateObitChecksumRequest) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *GenerateObitChecksumRequest) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *GenerateObitChecksumRequest) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.


### GetSerialNumber

`func (o *GenerateObitChecksumRequest) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *GenerateObitChecksumRequest) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *GenerateObitChecksumRequest) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### GetMetadataUri

`func (o *GenerateObitChecksumRequest) GetMetadataUri() string`

GetMetadataUri returns the MetadataUri field if non-nil, zero value otherwise.

### GetMetadataUriOk

`func (o *GenerateObitChecksumRequest) GetMetadataUriOk() (*string, bool)`

GetMetadataUriOk returns a tuple with the MetadataUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataUri

`func (o *GenerateObitChecksumRequest) SetMetadataUri(v string)`

SetMetadataUri sets MetadataUri field to given value.

### HasMetadataUri

`func (o *GenerateObitChecksumRequest) HasMetadataUri() bool`

HasMetadataUri returns a boolean if a field has been set.

### GetMetadataUriHash

`func (o *GenerateObitChecksumRequest) GetMetadataUriHash() string`

GetMetadataUriHash returns the MetadataUriHash field if non-nil, zero value otherwise.

### GetMetadataUriHashOk

`func (o *GenerateObitChecksumRequest) GetMetadataUriHashOk() (*string, bool)`

GetMetadataUriHashOk returns a tuple with the MetadataUriHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataUriHash

`func (o *GenerateObitChecksumRequest) SetMetadataUriHash(v string)`

SetMetadataUriHash sets MetadataUriHash field to given value.

### HasMetadataUriHash

`func (o *GenerateObitChecksumRequest) HasMetadataUriHash() bool`

HasMetadataUriHash returns a boolean if a field has been set.

### GetTrustAnchorToken

`func (o *GenerateObitChecksumRequest) GetTrustAnchorToken() string`

GetTrustAnchorToken returns the TrustAnchorToken field if non-nil, zero value otherwise.

### GetTrustAnchorTokenOk

`func (o *GenerateObitChecksumRequest) GetTrustAnchorTokenOk() (*string, bool)`

GetTrustAnchorTokenOk returns a tuple with the TrustAnchorToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustAnchorToken

`func (o *GenerateObitChecksumRequest) SetTrustAnchorToken(v string)`

SetTrustAnchorToken sets TrustAnchorToken field to given value.

### HasTrustAnchorToken

`func (o *GenerateObitChecksumRequest) HasTrustAnchorToken() bool`

HasTrustAnchorToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


