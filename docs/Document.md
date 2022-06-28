# Document

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Associative name of file | [optional] 
**Uri** | Pointer to **string** | Link to the document | [optional] 
**Hash** | Pointer to **string** | Hash of the file | [optional] 
**Encrypted** | Pointer to **bool** | If true then client helper will encrypt document with account key | [optional] 

## Methods

### NewDocument

`func NewDocument() *Document`

NewDocument instantiates a new Document object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentWithDefaults

`func NewDocumentWithDefaults() *Document`

NewDocumentWithDefaults instantiates a new Document object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Document) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Document) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Document) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Document) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUri

`func (o *Document) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *Document) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *Document) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *Document) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetHash

`func (o *Document) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *Document) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *Document) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *Document) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetEncrypted

`func (o *Document) GetEncrypted() bool`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *Document) GetEncryptedOk() (*bool, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *Document) SetEncrypted(v bool)`

SetEncrypted sets Encrypted field to given value.

### HasEncrypted

`func (o *Document) HasEncrypted() bool`

HasEncrypted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


