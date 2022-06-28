# \UtilsApi

All URIs are relative to *http://obs.node.obada.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateChecksum**](UtilsApi.md#GenerateChecksum) | **Post** /obit/checksum | Generates Obit checksum
[**GenerateDID**](UtilsApi.md#GenerateDID) | **Post** /obit/did | Generate Obit DID



## GenerateChecksum

> GenerateObitChecksumResponse GenerateChecksum(ctx).GenerateObitChecksumRequest(generateObitChecksumRequest).Execute()

Generates Obit checksum



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
    generateObitChecksumRequest := *openapiclient.NewGenerateObitChecksumRequest("Apple", "PN123456789", "SN123456789") // GenerateObitChecksumRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UtilsApi.GenerateChecksum(context.Background()).GenerateObitChecksumRequest(generateObitChecksumRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UtilsApi.GenerateChecksum``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateChecksum`: GenerateObitChecksumResponse
    fmt.Fprintf(os.Stdout, "Response from `UtilsApi.GenerateChecksum`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateChecksumRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateObitChecksumRequest** | [**GenerateObitChecksumRequest**](GenerateObitChecksumRequest.md) |  | 

### Return type

[**GenerateObitChecksumResponse**](GenerateObitChecksumResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateDID

> GenerateObitDIDResponse GenerateDID(ctx).GenerateObitDIDRequest(generateObitDIDRequest).Execute()

Generate Obit DID



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
    generateObitDIDRequest := *openapiclient.NewGenerateObitDIDRequest("Apple", "PN123456789", "SN123456789") // GenerateObitDIDRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UtilsApi.GenerateDID(context.Background()).GenerateObitDIDRequest(generateObitDIDRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UtilsApi.GenerateDID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateDID`: GenerateObitDIDResponse
    fmt.Fprintf(os.Stdout, "Response from `UtilsApi.GenerateDID`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateDIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateObitDIDRequest** | [**GenerateObitDIDRequest**](GenerateObitDIDRequest.md) |  | 

### Return type

[**GenerateObitDIDResponse**](GenerateObitDIDResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

