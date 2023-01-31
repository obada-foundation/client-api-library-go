# ImportAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateKey** | **string** | OBADA account | 
**AccountName** | Pointer to **string** | Associative account name | [optional] 

## Methods

### NewImportAccountRequest

`func NewImportAccountRequest(privateKey string, ) *ImportAccountRequest`

NewImportAccountRequest instantiates a new ImportAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportAccountRequestWithDefaults

`func NewImportAccountRequestWithDefaults() *ImportAccountRequest`

NewImportAccountRequestWithDefaults instantiates a new ImportAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateKey

`func (o *ImportAccountRequest) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *ImportAccountRequest) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *ImportAccountRequest) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetAccountName

`func (o *ImportAccountRequest) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *ImportAccountRequest) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *ImportAccountRequest) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *ImportAccountRequest) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


