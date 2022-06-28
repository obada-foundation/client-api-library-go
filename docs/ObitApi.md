# \ObitApi

All URIs are relative to *http://obs.node.obada.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](ObitApi.md#Get) | **Get** /obits/{key} | Get Obit by DID or USN
[**History**](ObitApi.md#History) | **Get** /obits/{key}/history | Get Obit history by DID or USN
[**Save**](ObitApi.md#Save) | **Post** /obits | Save Obit
[**Search**](ObitApi.md#Search) | **Get** /obits | Search obits by query



## Get

> Obit Get(ctx, key).Execute()

Get Obit by DID or USN



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    key := "did:obada:fe096095-e0f0-4918-9607-6567bd5756b5" // string | The given ObitDID or USN argument

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObitApi.Get(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObitApi.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: Obit
    fmt.Fprintf(os.Stdout, "Response from `ObitApi.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The given ObitDID or USN argument | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Obit**](Obit.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## History

> History200Response History(ctx, key).Execute()

Get Obit history by DID or USN



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    key := "did:obada:fe096095-e0f0-4918-9607-6567bd5756b5" // string | The given ObitDID or USN argument

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObitApi.History(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObitApi.History``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `History`: History200Response
    fmt.Fprintf(os.Stdout, "Response from `ObitApi.History`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The given ObitDID or USN argument | 

### Other Parameters

Other parameters are passed through a pointer to a apiHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**History200Response**](History200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Save

> Obit Save(ctx).SaveObitRequest(saveObitRequest).Execute()

Save Obit



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    saveObitRequest := *openapiclient.NewSaveObitRequest("Sony", "MWCN2LL/A", "f6fc84c9f21c24907d6bee6eec38cabab5fa9a7be8c4a7827fe9e56f245bd2d5") // SaveObitRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObitApi.Save(context.Background()).SaveObitRequest(saveObitRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObitApi.Save``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Save`: Obit
    fmt.Fprintf(os.Stdout, "Response from `ObitApi.Save`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **saveObitRequest** | [**SaveObitRequest**](SaveObitRequest.md) |  | 

### Return type

[**Obit**](Obit.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json:
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Search

> Obits Search(ctx).Q(q).Offset(offset).Execute()

Search obits by query



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    q := "fe403a1afe16203f4b8bb3a0e72d3e17567897bc15293e4a87b663e441030aea" // string | Query argument that used for a fulltext search (optional)
    offset := int32(56) // int32 | Number of records to skip for pagination. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObitApi.Search(context.Background()).Q(q).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObitApi.Search``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Search`: Obits
    fmt.Fprintf(os.Stdout, "Response from `ObitApi.Search`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Query argument that used for a fulltext search | 
 **offset** | **int32** | Number of records to skip for pagination. | [default to 0]

### Return type

[**Obits**](Obits.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

