# Obits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Obit**](Obit.md) |  | [optional] 
**Meta** | Pointer to [**ObitsMeta**](ObitsMeta.md) |  | [optional] 

## Methods

### NewObits

`func NewObits() *Obits`

NewObits instantiates a new Obits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObitsWithDefaults

`func NewObitsWithDefaults() *Obits`

NewObitsWithDefaults instantiates a new Obits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *Obits) GetData() []Obit`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Obits) GetDataOk() (*[]Obit, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Obits) SetData(v []Obit)`

SetData sets Data field to given value.

### HasData

`func (o *Obits) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *Obits) GetMeta() ObitsMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Obits) GetMetaOk() (*ObitsMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Obits) SetMeta(v ObitsMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Obits) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


