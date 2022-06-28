# \AccountsApi

All URIs are relative to *http://obs.node.obada.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Balance**](AccountsApi.md#Balance) | **Get** /accounts/my-balance | Shows account balance of OBADA address
[**CreateAccount**](AccountsApi.md#CreateAccount) | **Post** /accounts | Creates a new Account



## Balance

> AccountBalance Balance(ctx).Execute()

Shows account balance of OBADA address

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.Balance(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.Balance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Balance`: AccountBalance
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.Balance`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBalanceRequest struct via the builder pattern


### Return type

[**AccountBalance**](AccountBalance.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccount

> Account CreateAccount(ctx).NewAccountRequest(newAccountRequest).Execute()

Creates a new Account



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
    newAccountRequest := *openapiclient.NewNewAccountRequest("john.doe@obada.io") // NewAccountRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.CreateAccount(context.Background()).NewAccountRequest(newAccountRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.CreateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccount`: Account
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.CreateAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newAccountRequest** | [**NewAccountRequest**](NewAccountRequest.md) |  | 

### Return type

[**Account**](Account.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

