# ClientObitResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** |  | [optional] 
**Obit** | Pointer to [**ClientObit**](ClientObit.md) |  | [optional] 

## Methods

### NewClientObitResponse

`func NewClientObitResponse() *ClientObitResponse`

NewClientObitResponse instantiates a new ClientObitResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientObitResponseWithDefaults

`func NewClientObitResponseWithDefaults() *ClientObitResponse`

NewClientObitResponseWithDefaults instantiates a new ClientObitResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ClientObitResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClientObitResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClientObitResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClientObitResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetObit

`func (o *ClientObitResponse) GetObit() ClientObit`

GetObit returns the Obit field if non-nil, zero value otherwise.

### GetObitOk

`func (o *ClientObitResponse) GetObitOk() (*ClientObit, bool)`

GetObitOk returns a tuple with the Obit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObit

`func (o *ClientObitResponse) SetObit(v ClientObit)`

SetObit sets Obit field to given value.

### HasObit

`func (o *ClientObitResponse) HasObit() bool`

HasObit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


