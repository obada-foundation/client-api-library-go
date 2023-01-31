# SendCoinsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecepientAddress** | **string** |  | 
**Amount** | **string** |  | 
**Denom** | **string** |  | 

## Methods

### NewSendCoinsRequest

`func NewSendCoinsRequest(recepientAddress string, amount string, denom string, ) *SendCoinsRequest`

NewSendCoinsRequest instantiates a new SendCoinsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendCoinsRequestWithDefaults

`func NewSendCoinsRequestWithDefaults() *SendCoinsRequest`

NewSendCoinsRequestWithDefaults instantiates a new SendCoinsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecepientAddress

`func (o *SendCoinsRequest) GetRecepientAddress() string`

GetRecepientAddress returns the RecepientAddress field if non-nil, zero value otherwise.

### GetRecepientAddressOk

`func (o *SendCoinsRequest) GetRecepientAddressOk() (*string, bool)`

GetRecepientAddressOk returns a tuple with the RecepientAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecepientAddress

`func (o *SendCoinsRequest) SetRecepientAddress(v string)`

SetRecepientAddress sets RecepientAddress field to given value.


### GetAmount

`func (o *SendCoinsRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SendCoinsRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SendCoinsRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetDenom

`func (o *SendCoinsRequest) GetDenom() string`

GetDenom returns the Denom field if non-nil, zero value otherwise.

### GetDenomOk

`func (o *SendCoinsRequest) GetDenomOk() (*string, bool)`

GetDenomOk returns a tuple with the Denom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenom

`func (o *SendCoinsRequest) SetDenom(v string)`

SetDenom sets Denom field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


