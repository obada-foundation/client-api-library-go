# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Account address associated name | [optional] 
**PubKey** | Pointer to **string** | Public key | [optional] 
**Address** | Pointer to **string** | OBADA address | [optional] 
**Balance** | Pointer to [**AccountBalance**](AccountBalance.md) |  | [optional] 
**NftCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewAccount

`func NewAccount() *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Account) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Account) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Account) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Account) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPubKey

`func (o *Account) GetPubKey() string`

GetPubKey returns the PubKey field if non-nil, zero value otherwise.

### GetPubKeyOk

`func (o *Account) GetPubKeyOk() (*string, bool)`

GetPubKeyOk returns a tuple with the PubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubKey

`func (o *Account) SetPubKey(v string)`

SetPubKey sets PubKey field to given value.

### HasPubKey

`func (o *Account) HasPubKey() bool`

HasPubKey returns a boolean if a field has been set.

### GetAddress

`func (o *Account) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Account) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Account) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Account) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetBalance

`func (o *Account) GetBalance() AccountBalance`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *Account) GetBalanceOk() (*AccountBalance, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *Account) SetBalance(v AccountBalance)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *Account) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetNftCount

`func (o *Account) GetNftCount() int64`

GetNftCount returns the NftCount field if non-nil, zero value otherwise.

### GetNftCountOk

`func (o *Account) GetNftCountOk() (*int64, bool)`

GetNftCountOk returns a tuple with the NftCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftCount

`func (o *Account) SetNftCount(v int64)`

SetNftCount sets NftCount field to given value.

### HasNftCount

`func (o *Account) HasNftCount() bool`

HasNftCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


