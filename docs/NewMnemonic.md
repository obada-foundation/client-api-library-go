# NewMnemonic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mnemonic** | **string** | Mnemonic for seeding | 
**Force** | Pointer to **bool** | Flag that specify if exisiting wallet should be replaced, if false is send then error will be send back | [optional] [default to false]

## Methods

### NewNewMnemonic

`func NewNewMnemonic(mnemonic string, ) *NewMnemonic`

NewNewMnemonic instantiates a new NewMnemonic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewMnemonicWithDefaults

`func NewNewMnemonicWithDefaults() *NewMnemonic`

NewNewMnemonicWithDefaults instantiates a new NewMnemonic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMnemonic

`func (o *NewMnemonic) GetMnemonic() string`

GetMnemonic returns the Mnemonic field if non-nil, zero value otherwise.

### GetMnemonicOk

`func (o *NewMnemonic) GetMnemonicOk() (*string, bool)`

GetMnemonicOk returns a tuple with the Mnemonic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnemonic

`func (o *NewMnemonic) SetMnemonic(v string)`

SetMnemonic sets Mnemonic field to given value.


### GetForce

`func (o *NewMnemonic) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *NewMnemonic) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *NewMnemonic) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *NewMnemonic) HasForce() bool`

HasForce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


