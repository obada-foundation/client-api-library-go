# SaveObitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Manufacturer** | **string** | Waiting more specific details from Rohi | 
**PartNumber** | **string** | Manufacturer provided. In cases where no part number is provided for the product, use model, or the most specific ID available from the manufacturer. MWCN2LL/A (an iPhone 11 Pro, Silver, 256GB, model A2160) | 
**SerialNumber** | **string** | Serial number hashed with sha256 hash function | 
**Documents** | Pointer to [**[]DeviceDocument**](DeviceDocument.md) |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 

## Methods

### NewSaveObitRequest

`func NewSaveObitRequest(manufacturer string, partNumber string, serialNumber string, ) *SaveObitRequest`

NewSaveObitRequest instantiates a new SaveObitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveObitRequestWithDefaults

`func NewSaveObitRequestWithDefaults() *SaveObitRequest`

NewSaveObitRequestWithDefaults instantiates a new SaveObitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManufacturer

`func (o *SaveObitRequest) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *SaveObitRequest) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *SaveObitRequest) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.


### GetPartNumber

`func (o *SaveObitRequest) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *SaveObitRequest) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *SaveObitRequest) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.


### GetSerialNumber

`func (o *SaveObitRequest) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *SaveObitRequest) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *SaveObitRequest) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### GetDocuments

`func (o *SaveObitRequest) GetDocuments() []DeviceDocument`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *SaveObitRequest) GetDocumentsOk() (*[]DeviceDocument, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *SaveObitRequest) SetDocuments(v []DeviceDocument)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *SaveObitRequest) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetAddress

`func (o *SaveObitRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SaveObitRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SaveObitRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *SaveObitRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


