# UnprocessableEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] [default to "data validation error"]
**Fields** | Pointer to [**[]UnprocessableEntityFieldsInner**](UnprocessableEntityFieldsInner.md) |  | [optional] 

## Methods

### NewUnprocessableEntity

`func NewUnprocessableEntity() *UnprocessableEntity`

NewUnprocessableEntity instantiates a new UnprocessableEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnprocessableEntityWithDefaults

`func NewUnprocessableEntityWithDefaults() *UnprocessableEntity`

NewUnprocessableEntityWithDefaults instantiates a new UnprocessableEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *UnprocessableEntity) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *UnprocessableEntity) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *UnprocessableEntity) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *UnprocessableEntity) HasError() bool`

HasError returns a boolean if a field has been set.

### GetFields

`func (o *UnprocessableEntity) GetFields() []UnprocessableEntityFieldsInner`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *UnprocessableEntity) GetFieldsOk() (*[]UnprocessableEntityFieldsInner, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *UnprocessableEntity) SetFields(v []UnprocessableEntityFieldsInner)`

SetFields sets Fields field to given value.

### HasFields

`func (o *UnprocessableEntity) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


