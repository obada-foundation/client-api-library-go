# Accounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HdAccounts** | Pointer to [**[]Account**](Account.md) |  | [optional] 
**ImportedAccounts** | Pointer to [**[]Account**](Account.md) |  | [optional] 

## Methods

### NewAccounts

`func NewAccounts() *Accounts`

NewAccounts instantiates a new Accounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountsWithDefaults

`func NewAccountsWithDefaults() *Accounts`

NewAccountsWithDefaults instantiates a new Accounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHdAccounts

`func (o *Accounts) GetHdAccounts() []Account`

GetHdAccounts returns the HdAccounts field if non-nil, zero value otherwise.

### GetHdAccountsOk

`func (o *Accounts) GetHdAccountsOk() (*[]Account, bool)`

GetHdAccountsOk returns a tuple with the HdAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdAccounts

`func (o *Accounts) SetHdAccounts(v []Account)`

SetHdAccounts sets HdAccounts field to given value.

### HasHdAccounts

`func (o *Accounts) HasHdAccounts() bool`

HasHdAccounts returns a boolean if a field has been set.

### GetImportedAccounts

`func (o *Accounts) GetImportedAccounts() []Account`

GetImportedAccounts returns the ImportedAccounts field if non-nil, zero value otherwise.

### GetImportedAccountsOk

`func (o *Accounts) GetImportedAccountsOk() (*[]Account, bool)`

GetImportedAccountsOk returns a tuple with the ImportedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedAccounts

`func (o *Accounts) SetImportedAccounts(v []Account)`

SetImportedAccounts sets ImportedAccounts field to given value.

### HasImportedAccounts

`func (o *Accounts) HasImportedAccounts() bool`

HasImportedAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


