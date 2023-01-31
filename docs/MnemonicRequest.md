# MnemonicRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mnemonic** | **string** | Mnemonic for seeding | 
**Force** | Pointer to **bool** | Flag that specify if exisiting wallet should be replaced, if false is send then error will be send back | [optional] [default to false]

## Methods

### NewMnemonicRequest

`func NewMnemonicRequest(mnemonic string, ) *MnemonicRequest`

NewMnemonicRequest instantiates a new MnemonicRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMnemonicRequestWithDefaults

`func NewMnemonicRequestWithDefaults() *MnemonicRequest`

NewMnemonicRequestWithDefaults instantiates a new MnemonicRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMnemonic

`func (o *MnemonicRequest) GetMnemonic() string`

GetMnemonic returns the Mnemonic field if non-nil, zero value otherwise.

### GetMnemonicOk

`func (o *MnemonicRequest) GetMnemonicOk() (*string, bool)`

GetMnemonicOk returns a tuple with the Mnemonic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnemonic

`func (o *MnemonicRequest) SetMnemonic(v string)`

SetMnemonic sets Mnemonic field to given value.


### GetForce

`func (o *MnemonicRequest) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *MnemonicRequest) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *MnemonicRequest) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *MnemonicRequest) HasForce() bool`

HasForce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


