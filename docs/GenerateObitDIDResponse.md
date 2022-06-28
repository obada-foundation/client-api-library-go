# GenerateObitDIDResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SerialNumberHash** | Pointer to **string** | Serial Number Hash is sha256 of the Serial Number input | [optional] 
**Usn** | Pointer to **string** | Universal Serial Number | [optional] 
**Did** | Pointer to **string** | DID | [optional] 
**UsnBase58** | Pointer to **string** | Base58 | [optional] 

## Methods

### NewGenerateObitDIDResponse

`func NewGenerateObitDIDResponse() *GenerateObitDIDResponse`

NewGenerateObitDIDResponse instantiates a new GenerateObitDIDResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateObitDIDResponseWithDefaults

`func NewGenerateObitDIDResponseWithDefaults() *GenerateObitDIDResponse`

NewGenerateObitDIDResponseWithDefaults instantiates a new GenerateObitDIDResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerialNumberHash

`func (o *GenerateObitDIDResponse) GetSerialNumberHash() string`

GetSerialNumberHash returns the SerialNumberHash field if non-nil, zero value otherwise.

### GetSerialNumberHashOk

`func (o *GenerateObitDIDResponse) GetSerialNumberHashOk() (*string, bool)`

GetSerialNumberHashOk returns a tuple with the SerialNumberHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumberHash

`func (o *GenerateObitDIDResponse) SetSerialNumberHash(v string)`

SetSerialNumberHash sets SerialNumberHash field to given value.

### HasSerialNumberHash

`func (o *GenerateObitDIDResponse) HasSerialNumberHash() bool`

HasSerialNumberHash returns a boolean if a field has been set.

### GetUsn

`func (o *GenerateObitDIDResponse) GetUsn() string`

GetUsn returns the Usn field if non-nil, zero value otherwise.

### GetUsnOk

`func (o *GenerateObitDIDResponse) GetUsnOk() (*string, bool)`

GetUsnOk returns a tuple with the Usn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsn

`func (o *GenerateObitDIDResponse) SetUsn(v string)`

SetUsn sets Usn field to given value.

### HasUsn

`func (o *GenerateObitDIDResponse) HasUsn() bool`

HasUsn returns a boolean if a field has been set.

### GetDid

`func (o *GenerateObitDIDResponse) GetDid() string`

GetDid returns the Did field if non-nil, zero value otherwise.

### GetDidOk

`func (o *GenerateObitDIDResponse) GetDidOk() (*string, bool)`

GetDidOk returns a tuple with the Did field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDid

`func (o *GenerateObitDIDResponse) SetDid(v string)`

SetDid sets Did field to given value.

### HasDid

`func (o *GenerateObitDIDResponse) HasDid() bool`

HasDid returns a boolean if a field has been set.

### GetUsnBase58

`func (o *GenerateObitDIDResponse) GetUsnBase58() string`

GetUsnBase58 returns the UsnBase58 field if non-nil, zero value otherwise.

### GetUsnBase58Ok

`func (o *GenerateObitDIDResponse) GetUsnBase58Ok() (*string, bool)`

GetUsnBase58Ok returns a tuple with the UsnBase58 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsnBase58

`func (o *GenerateObitDIDResponse) SetUsnBase58(v string)`

SetUsnBase58 sets UsnBase58 field to given value.

### HasUsnBase58

`func (o *GenerateObitDIDResponse) HasUsnBase58() bool`

HasUsnBase58 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


