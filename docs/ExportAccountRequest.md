# ExportAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | OBADA account | [optional] 
**Passphrase** | Pointer to **string** | Passphrase to decrypt the account | [optional] 

## Methods

### NewExportAccountRequest

`func NewExportAccountRequest() *ExportAccountRequest`

NewExportAccountRequest instantiates a new ExportAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportAccountRequestWithDefaults

`func NewExportAccountRequestWithDefaults() *ExportAccountRequest`

NewExportAccountRequestWithDefaults instantiates a new ExportAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ExportAccountRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ExportAccountRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ExportAccountRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ExportAccountRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPassphrase

`func (o *ExportAccountRequest) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *ExportAccountRequest) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *ExportAccountRequest) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *ExportAccountRequest) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


