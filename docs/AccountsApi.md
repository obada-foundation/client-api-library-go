# \AccountsApi

All URIs are relative to *http://obs.node.obada.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Account**](AccountsApi.md#Account) | **Get** /accounts/{address} | Fetches an information about single account
[**Accounts**](AccountsApi.md#Accounts) | **Get** /accounts | Returns a list of OBADA accounts
[**Balance**](AccountsApi.md#Balance) | **Get** /accounts/my-balance | Shows account balance of OBADA address
[**ExportAccount**](AccountsApi.md#ExportAccount) | **Post** /accounts/export-account | Export OBADA account (private key) from client-helper
[**GetMnemonic**](AccountsApi.md#GetMnemonic) | **Get** /accounts/mnemonic | Fetching an existing mnemonic phrase
[**ImportAccount**](AccountsApi.md#ImportAccount) | **Post** /accounts/import-account | Imports an existing OBADA account (private key) to the client-helper user profile
[**ImportWallet**](AccountsApi.md#ImportWallet) | **Post** /accounts/import-wallet | Imports an existing HD wallet to the client-helper user profile
[**NewAccount**](AccountsApi.md#NewAccount) | **Post** /accounts/new-account | Creates a new OBADA account from HD wallet master key
[**NewMnemonic**](AccountsApi.md#NewMnemonic) | **Get** /accounts/new-mnemonic | Generate a new mnemonic phrase for seeding wallet
[**NewWallet**](AccountsApi.md#NewWallet) | **Post** /accounts/new-wallet | Creates profile HD wallet
[**Register**](AccountsApi.md#Register) | **Post** /accounts/register | Register a new client-helper user profile
[**SendCoins**](AccountsApi.md#SendCoins) | **Post** /accounts/{address}/send-coins | Send coins from selected account
[**UpdateAccount**](AccountsApi.md#UpdateAccount) | **Post** /accounts/{address} | Sets account specific information



## Account

> Account Account(ctx, address).Execute()

Fetches an information about single account

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
    address := "obada1yxxnd624tgwqm3eyv5smdvjrrydfh9h943qptg" // string | OBADA address

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.Account(context.Background(), address).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.Account``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Account`: Account
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.Account`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | OBADA address | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Account**](Account.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Accounts

> Accounts Accounts(ctx).Execute()

Returns a list of OBADA accounts

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
    resp, r, err := apiClient.AccountsApi.Accounts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.Accounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Accounts`: Accounts
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.Accounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsRequest struct via the builder pattern


### Return type

[**Accounts**](Accounts.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## ExportAccount

> ExportAccountResponse ExportAccount(ctx).ExportAccountRequest(exportAccountRequest).Execute()

Export OBADA account (private key) from client-helper

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
    exportAccountRequest := *openapiclient.NewExportAccountRequest() // ExportAccountRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.ExportAccount(context.Background()).ExportAccountRequest(exportAccountRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.ExportAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportAccount`: ExportAccountResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.ExportAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exportAccountRequest** | [**ExportAccountRequest**](ExportAccountRequest.md) |  | 

### Return type

[**ExportAccountResponse**](ExportAccountResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMnemonic

> NewMnemonic GetMnemonic(ctx).Execute()

Fetching an existing mnemonic phrase

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
    resp, r, err := apiClient.AccountsApi.GetMnemonic(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetMnemonic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMnemonic`: NewMnemonic
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetMnemonic`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMnemonicRequest struct via the builder pattern


### Return type

[**NewMnemonic**](NewMnemonic.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportAccount

> ImportAccount(ctx).ImportAccountRequest(importAccountRequest).Execute()

Imports an existing OBADA account (private key) to the client-helper user profile

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
    importAccountRequest := *openapiclient.NewImportAccountRequest("PrivateKey_example") // ImportAccountRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.ImportAccount(context.Background()).ImportAccountRequest(importAccountRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.ImportAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importAccountRequest** | [**ImportAccountRequest**](ImportAccountRequest.md) |  | 

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


## ImportWallet

> ImportWallet(ctx).MnemonicRequest(mnemonicRequest).Execute()

Imports an existing HD wallet to the client-helper user profile

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
    mnemonicRequest := *openapiclient.NewMnemonicRequest("fantasy route flavor zoo laptop rent knife stick fancy flame black fan oval stairs express identify crane truly anxiety wave notable gather toe tag") // MnemonicRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.ImportWallet(context.Background()).MnemonicRequest(mnemonicRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.ImportWallet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mnemonicRequest** | [**MnemonicRequest**](MnemonicRequest.md) |  | 

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


## NewAccount

> NewAccount(ctx).AccountRequest(accountRequest).Execute()

Creates a new OBADA account from HD wallet master key

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
    accountRequest := *openapiclient.NewAccountRequest() // AccountRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.NewAccount(context.Background()).AccountRequest(accountRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.NewAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNewAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountRequest** | [**AccountRequest**](AccountRequest.md) |  | 

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


## NewMnemonic

> NewMnemonic NewMnemonic(ctx).Execute()

Generate a new mnemonic phrase for seeding wallet

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
    resp, r, err := apiClient.AccountsApi.NewMnemonic(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.NewMnemonic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NewMnemonic`: NewMnemonic
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.NewMnemonic`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNewMnemonicRequest struct via the builder pattern


### Return type

[**NewMnemonic**](NewMnemonic.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewWallet

> NewWallet(ctx).MnemonicRequest(mnemonicRequest).Execute()

Creates profile HD wallet

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
    mnemonicRequest := *openapiclient.NewMnemonicRequest("fantasy route flavor zoo laptop rent knife stick fancy flame black fan oval stairs express identify crane truly anxiety wave notable gather toe tag") // MnemonicRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.NewWallet(context.Background()).MnemonicRequest(mnemonicRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.NewWallet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNewWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mnemonicRequest** | [**MnemonicRequest**](MnemonicRequest.md) |  | 

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


## Register

> Profile Register(ctx).RegisterRequest(registerRequest).Execute()

Register a new client-helper user profile



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
    registerRequest := *openapiclient.NewRegisterRequest("john.doe@obada.io") // RegisterRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.Register(context.Background()).RegisterRequest(registerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.Register``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Register`: Profile
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.Register`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registerRequest** | [**RegisterRequest**](RegisterRequest.md) |  | 

### Return type

[**Profile**](Profile.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendCoins

> SendCoins(ctx, address).SendCoinsRequest(sendCoinsRequest).Execute()

Send coins from selected account

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
    address := "obada1yxxnd624tgwqm3eyv5smdvjrrydfh9h943qptg" // string | OBADA address
    sendCoinsRequest := *openapiclient.NewSendCoinsRequest("RecepientAddress_example", "Amount_example", "Denom_example") // SendCoinsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.SendCoins(context.Background(), address).SendCoinsRequest(sendCoinsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.SendCoins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | OBADA address | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendCoinsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendCoinsRequest** | [**SendCoinsRequest**](SendCoinsRequest.md) |  | 

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


## UpdateAccount

> UpdateAccount(ctx, address).AccountRequest(accountRequest).Execute()

Sets account specific information

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
    address := "obada1yxxnd624tgwqm3eyv5smdvjrrydfh9h943qptg" // string | OBADA address
    accountRequest := *openapiclient.NewAccountRequest() // AccountRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.UpdateAccount(context.Background(), address).AccountRequest(accountRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.UpdateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | OBADA address | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountRequest** | [**AccountRequest**](AccountRequest.md) |  | 

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

