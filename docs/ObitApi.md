# \ObitApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadObitFromChain**](ObitApi.md#DownloadObitFromChain) | **Post** /api/server/obit/download | Download Obit from Blockchain
[**FetchObitFromChain**](ObitApi.md#FetchObitFromChain) | **Get** /api/server/obit/{obit_did} | Get Obit From Blockchain
[**GenerateObitDef**](ObitApi.md#GenerateObitDef) | **Get** /api/obit/generate | Generate Obit Definition
[**GetClientObit**](ObitApi.md#GetClientObit) | **Get** /api/client/obit/{obit_did} | Get Client Obit
[**SaveClientObit**](ObitApi.md#SaveClientObit) | **Post** /api/client/obit | Save Client Obit
[**UploadObit**](ObitApi.md#UploadObit) | **Post** /api/server/obit/upload | Upload Obit to Blockchain



## DownloadObitFromChain

> ClientObitResponse DownloadObitFromChain(ctx).ObitDid(obitDid).Execute()

Download Obit from Blockchain



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
    obitDid := *openapiclient.NewObitDid("did:obada:fe096095-e0f0-4918-9607-6567bd5756b5") // ObitDid |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObitApi.DownloadObitFromChain(context.Background()).ObitDid(obitDid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObitApi.DownloadObitFromChain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadObitFromChain`: ClientObitResponse
    fmt.Fprintf(os.Stdout, "Response from `ObitApi.DownloadObitFromChain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDownloadObitFromChainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **obitDid** | [**ObitDid**](ObitDid.md) |  | 

### Return type

[**ClientObitResponse**](ClientObitResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchObitFromChain

> BlockChainObitResponse FetchObitFromChain(ctx, obitDid).Execute()

Get Obit From Blockchain



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
    obitDid := "did:obada:81413bc1ad2074a6ae35d1f65f64f1bca2e8a20959f37ef0349a729ddc567d9b" // string | Required.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObitApi.FetchObitFromChain(context.Background(), obitDid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObitApi.FetchObitFromChain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchObitFromChain`: BlockChainObitResponse
    fmt.Fprintf(os.Stdout, "Response from `ObitApi.FetchObitFromChain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**obitDid** | **string** | Required. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchObitFromChainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BlockChainObitResponse**](BlockChainObitResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateObitDef

> ObitDefinitionResponse GenerateObitDef(ctx).Manufacturer(manufacturer).PartNumber(partNumber).SerialNumber(serialNumber).Execute()

Generate Obit Definition



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
    manufacturer := "Apple" // string | Device Id (Required)
    partNumber := "123456789" // string | Part Number (Required)
    serialNumber := "123456789" // string | Serial Number (Required)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObitApi.GenerateObitDef(context.Background()).Manufacturer(manufacturer).PartNumber(partNumber).SerialNumber(serialNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObitApi.GenerateObitDef``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateObitDef`: ObitDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `ObitApi.GenerateObitDef`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateObitDefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **manufacturer** | **string** | Device Id (Required) | 
 **partNumber** | **string** | Part Number (Required) | 
 **serialNumber** | **string** | Serial Number (Required) | 

### Return type

[**ObitDefinitionResponse**](ObitDefinitionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientObit

> ClientObitResponse GetClientObit(ctx, obitDid).Execute()

Get Client Obit

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
    obitDid := "did:obada:81413bc1ad2074a6ae35d1f65f64f1bca2e8a20959f37ef0349a729ddc567d9b" // string | Required.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObitApi.GetClientObit(context.Background(), obitDid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObitApi.GetClientObit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClientObit`: ClientObitResponse
    fmt.Fprintf(os.Stdout, "Response from `ObitApi.GetClientObit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**obitDid** | **string** | Required. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientObitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClientObitResponse**](ClientObitResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveClientObit

> ClientObitResponse SaveClientObit(ctx).LocalObit(localObit).Execute()

Save Client Obit

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    localObit := *openapiclient.NewLocalObit("Tradeloop", "ObitStatus_example", "Sony", "MWCN2LL/A", "f6fc84c9f21c24907d6bee6eec38cabab5fa9a7be8c4a7827fe9e56f245bd2d5", time.Now()) // LocalObit |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObitApi.SaveClientObit(context.Background()).LocalObit(localObit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObitApi.SaveClientObit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SaveClientObit`: ClientObitResponse
    fmt.Fprintf(os.Stdout, "Response from `ObitApi.SaveClientObit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveClientObitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **localObit** | [**LocalObit**](LocalObit.md) |  | 

### Return type

[**ClientObitResponse**](ClientObitResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadObit

> BaseResponse UploadObit(ctx).ObitDid(obitDid).Execute()

Upload Obit to Blockchain



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
    obitDid := *openapiclient.NewObitDid("did:obada:fe096095-e0f0-4918-9607-6567bd5756b5") // ObitDid |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObitApi.UploadObit(context.Background()).ObitDid(obitDid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObitApi.UploadObit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadObit`: BaseResponse
    fmt.Fprintf(os.Stdout, "Response from `ObitApi.UploadObit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadObitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **obitDid** | [**ObitDid**](ObitDid.md) |  | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

