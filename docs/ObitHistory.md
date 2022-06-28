# ObitHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObitDid** | Pointer to **string** | OBADA decentralized identifier (max length Rohi?) | [optional] 
**Event** | Pointer to **string** | History event | [optional] 
**OldValues** | Pointer to **map[string]interface{}** |  | [optional] 
**NewValues** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewObitHistory

`func NewObitHistory() *ObitHistory`

NewObitHistory instantiates a new ObitHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObitHistoryWithDefaults

`func NewObitHistoryWithDefaults() *ObitHistory`

NewObitHistoryWithDefaults instantiates a new ObitHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObitDid

`func (o *ObitHistory) GetObitDid() string`

GetObitDid returns the ObitDid field if non-nil, zero value otherwise.

### GetObitDidOk

`func (o *ObitHistory) GetObitDidOk() (*string, bool)`

GetObitDidOk returns a tuple with the ObitDid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObitDid

`func (o *ObitHistory) SetObitDid(v string)`

SetObitDid sets ObitDid field to given value.

### HasObitDid

`func (o *ObitHistory) HasObitDid() bool`

HasObitDid returns a boolean if a field has been set.

### GetEvent

`func (o *ObitHistory) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *ObitHistory) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *ObitHistory) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *ObitHistory) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetOldValues

`func (o *ObitHistory) GetOldValues() map[string]interface{}`

GetOldValues returns the OldValues field if non-nil, zero value otherwise.

### GetOldValuesOk

`func (o *ObitHistory) GetOldValuesOk() (*map[string]interface{}, bool)`

GetOldValuesOk returns a tuple with the OldValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldValues

`func (o *ObitHistory) SetOldValues(v map[string]interface{})`

SetOldValues sets OldValues field to given value.

### HasOldValues

`func (o *ObitHistory) HasOldValues() bool`

HasOldValues returns a boolean if a field has been set.

### GetNewValues

`func (o *ObitHistory) GetNewValues() map[string]interface{}`

GetNewValues returns the NewValues field if non-nil, zero value otherwise.

### GetNewValuesOk

`func (o *ObitHistory) GetNewValuesOk() (*map[string]interface{}, bool)`

GetNewValuesOk returns a tuple with the NewValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewValues

`func (o *ObitHistory) SetNewValues(v map[string]interface{})`

SetNewValues sets NewValues field to given value.

### HasNewValues

`func (o *ObitHistory) HasNewValues() bool`

HasNewValues returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ObitHistory) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ObitHistory) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ObitHistory) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ObitHistory) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ObitHistory) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ObitHistory) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ObitHistory) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ObitHistory) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


