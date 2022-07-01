# DeviceDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Associative name of device document | 
**DocumentFile** | **string** |  | 
**ShouldEncrypt** | Pointer to **bool** | If true then client helper will encrypt document with account key | [optional] [default to false]

## Methods

### NewDeviceDocument

`func NewDeviceDocument(name string, documentFile string, ) *DeviceDocument`

NewDeviceDocument instantiates a new DeviceDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceDocumentWithDefaults

`func NewDeviceDocumentWithDefaults() *DeviceDocument`

NewDeviceDocumentWithDefaults instantiates a new DeviceDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeviceDocument) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceDocument) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceDocument) SetName(v string)`

SetName sets Name field to given value.


### GetDocumentFile

`func (o *DeviceDocument) GetDocumentFile() string`

GetDocumentFile returns the DocumentFile field if non-nil, zero value otherwise.

### GetDocumentFileOk

`func (o *DeviceDocument) GetDocumentFileOk() (*string, bool)`

GetDocumentFileOk returns a tuple with the DocumentFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentFile

`func (o *DeviceDocument) SetDocumentFile(v string)`

SetDocumentFile sets DocumentFile field to given value.


### GetShouldEncrypt

`func (o *DeviceDocument) GetShouldEncrypt() bool`

GetShouldEncrypt returns the ShouldEncrypt field if non-nil, zero value otherwise.

### GetShouldEncryptOk

`func (o *DeviceDocument) GetShouldEncryptOk() (*bool, bool)`

GetShouldEncryptOk returns a tuple with the ShouldEncrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldEncrypt

`func (o *DeviceDocument) SetShouldEncrypt(v bool)`

SetShouldEncrypt sets ShouldEncrypt field to given value.

### HasShouldEncrypt

`func (o *DeviceDocument) HasShouldEncrypt() bool`

HasShouldEncrypt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


