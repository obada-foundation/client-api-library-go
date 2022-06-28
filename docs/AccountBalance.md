# AccountBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**Balance** | Pointer to **int64** |  | [optional] 

## Methods

### NewAccountBalance

`func NewAccountBalance() *AccountBalance`

NewAccountBalance instantiates a new AccountBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountBalanceWithDefaults

`func NewAccountBalanceWithDefaults() *AccountBalance`

NewAccountBalanceWithDefaults instantiates a new AccountBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *AccountBalance) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AccountBalance) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AccountBalance) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AccountBalance) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetBalance

`func (o *AccountBalance) GetBalance() int64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *AccountBalance) GetBalanceOk() (*int64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *AccountBalance) SetBalance(v int64)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *AccountBalance) HasBalance() bool`

HasBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


