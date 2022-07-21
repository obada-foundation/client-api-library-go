# \NFTApi

All URIs are relative to *http://obs.node.obada.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Mint**](NFTApi.md#Mint) | **Post** /nft/{key}/mint | Mints NFT
[**Nft**](NFTApi.md#Nft) | **Get** /nft/{key} | Fetch NFT from OBADA blockchain Node
[**Send**](NFTApi.md#Send) | **Post** /nft/{key}/send | Send NFT to another address
[**UpdateMetadata**](NFTApi.md#UpdateMetadata) | **Post** /nft/{key}/metadata | Update NFT metadata



## Mint

> Mint(ctx, key).Execute()

Mints NFT

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
    resp, r, err := apiClient.NFTApi.Mint(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NFTApi.Mint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The given ObitDID or USN argument | 

### Other Parameters

Other parameters are passed through a pointer to a apiMintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Nft

> NFT Nft(ctx, key).Execute()

Fetch NFT from OBADA blockchain Node

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
    resp, r, err := apiClient.NFTApi.Nft(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NFTApi.Nft``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Nft`: NFT
    fmt.Fprintf(os.Stdout, "Response from `NFTApi.Nft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The given ObitDID or USN argument | 

### Other Parameters

Other parameters are passed through a pointer to a apiNftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NFT**](NFT.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Send

> Send(ctx, key).SendNFTRequest(sendNFTRequest).Execute()

Send NFT to another address

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
    sendNFTRequest := *openapiclient.NewSendNFTRequest("Receiver_example") // SendNFTRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NFTApi.Send(context.Background(), key).SendNFTRequest(sendNFTRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NFTApi.Send``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The given ObitDID or USN argument | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendNFTRequest** | [**SendNFTRequest**](SendNFTRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMetadata

> UpdateMetadata(ctx, key).Execute()

Update NFT metadata

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
    resp, r, err := apiClient.NFTApi.UpdateMetadata(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NFTApi.UpdateMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The given ObitDID or USN argument | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

