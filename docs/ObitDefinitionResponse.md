# ObitDefinitionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** |  | [optional] 
**Obit** | Pointer to [**ObitDefinition**](ObitDefinition.md) |  | [optional] 

## Methods

### NewObitDefinitionResponse

`func NewObitDefinitionResponse() *ObitDefinitionResponse`

NewObitDefinitionResponse instantiates a new ObitDefinitionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObitDefinitionResponseWithDefaults

`func NewObitDefinitionResponseWithDefaults() *ObitDefinitionResponse`

NewObitDefinitionResponseWithDefaults instantiates a new ObitDefinitionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ObitDefinitionResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ObitDefinitionResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ObitDefinitionResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ObitDefinitionResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetObit

`func (o *ObitDefinitionResponse) GetObit() ObitDefinition`

GetObit returns the Obit field if non-nil, zero value otherwise.

### GetObitOk

`func (o *ObitDefinitionResponse) GetObitOk() (*ObitDefinition, bool)`

GetObitOk returns a tuple with the Obit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObit

`func (o *ObitDefinitionResponse) SetObit(v ObitDefinition)`

SetObit sets Obit field to given value.

### HasObit

`func (o *ObitDefinitionResponse) HasObit() bool`

HasObit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


