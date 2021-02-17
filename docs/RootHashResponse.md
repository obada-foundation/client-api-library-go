# RootHashResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** |  | [optional] 
**RootHash** | Pointer to **string** |  | [optional] 

## Methods

### NewRootHashResponse

`func NewRootHashResponse() *RootHashResponse`

NewRootHashResponse instantiates a new RootHashResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRootHashResponseWithDefaults

`func NewRootHashResponseWithDefaults() *RootHashResponse`

NewRootHashResponseWithDefaults instantiates a new RootHashResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RootHashResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RootHashResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RootHashResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RootHashResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRootHash

`func (o *RootHashResponse) GetRootHash() string`

GetRootHash returns the RootHash field if non-nil, zero value otherwise.

### GetRootHashOk

`func (o *RootHashResponse) GetRootHashOk() (*string, bool)`

GetRootHashOk returns a tuple with the RootHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootHash

`func (o *RootHashResponse) SetRootHash(v string)`

SetRootHash sets RootHash field to given value.

### HasRootHash

`func (o *RootHashResponse) HasRootHash() bool`

HasRootHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


