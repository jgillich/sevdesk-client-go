# \CheckAccountApi

All URIs are relative to *https://my.sevdesk.de/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCheckAccount**](CheckAccountApi.md#CreateCheckAccount) | **Post** /CheckAccount | Create a new check account
[**DeleteCheckAccount**](CheckAccountApi.md#DeleteCheckAccount) | **Delete** /CheckAccount/{checkAccountId} | Deletes a check account
[**GetBalanceAtDate**](CheckAccountApi.md#GetBalanceAtDate) | **Get** /CheckAccount/{checkAccountId}/getBalanceAtDate | Get the balance at a given date
[**GetCheckAccountById**](CheckAccountApi.md#GetCheckAccountById) | **Get** /CheckAccount/{checkAccountId} | Find check account by ID
[**GetCheckAccounts**](CheckAccountApi.md#GetCheckAccounts) | **Get** /CheckAccount | Retrieve check accounts
[**UpdateCheckAccount**](CheckAccountApi.md#UpdateCheckAccount) | **Put** /CheckAccount/{checkAccountId} | Update an existing check account



## CreateCheckAccount

> GetCheckAccountById200Response CreateCheckAccount(ctx).ModelCheckAccount(modelCheckAccount).Execute()

Create a new check account



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/jgillich/sevdesk-client-go"
)

func main() {
    modelCheckAccount := *openapiclient.NewModelCheckAccount("Iron Bank", "online", "EUR", int32(123)) // ModelCheckAccount | Creation data. Please be aware, that you need to provide at least all required parameter      of the CheckAccount model! (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CheckAccountApi.CreateCheckAccount(context.Background()).ModelCheckAccount(modelCheckAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckAccountApi.CreateCheckAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCheckAccount`: GetCheckAccountById200Response
    fmt.Fprintf(os.Stdout, "Response from `CheckAccountApi.CreateCheckAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCheckAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelCheckAccount** | [**ModelCheckAccount**](ModelCheckAccount.md) | Creation data. Please be aware, that you need to provide at least all required parameter      of the CheckAccount model! | 

### Return type

[**GetCheckAccountById200Response**](GetCheckAccountById200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCheckAccount

> DeleteAccountingContact200Response DeleteCheckAccount(ctx, checkAccountId).Execute()

Deletes a check account

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/jgillich/sevdesk-client-go"
)

func main() {
    checkAccountId := int32(56) // int32 | Id of check account to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CheckAccountApi.DeleteCheckAccount(context.Background(), checkAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckAccountApi.DeleteCheckAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCheckAccount`: DeleteAccountingContact200Response
    fmt.Fprintf(os.Stdout, "Response from `CheckAccountApi.DeleteCheckAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkAccountId** | **int32** | Id of check account to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCheckAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteAccountingContact200Response**](DeleteAccountingContact200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBalanceAtDate

> GetBalanceAtDate200Response GetBalanceAtDate(ctx, checkAccountId).Date(date).Execute()

Get the balance at a given date



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/jgillich/sevdesk-client-go"
)

func main() {
    checkAccountId := int32(56) // int32 | ID of check account
    date := time.Now() // string | Only consider transactions up to this date at 23:59:59

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CheckAccountApi.GetBalanceAtDate(context.Background(), checkAccountId).Date(date).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckAccountApi.GetBalanceAtDate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBalanceAtDate`: GetBalanceAtDate200Response
    fmt.Fprintf(os.Stdout, "Response from `CheckAccountApi.GetBalanceAtDate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkAccountId** | **int32** | ID of check account | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBalanceAtDateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **date** | **string** | Only consider transactions up to this date at 23:59:59 | 

### Return type

[**GetBalanceAtDate200Response**](GetBalanceAtDate200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCheckAccountById

> GetCheckAccountById200Response GetCheckAccountById(ctx, checkAccountId).Execute()

Find check account by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/jgillich/sevdesk-client-go"
)

func main() {
    checkAccountId := int32(56) // int32 | ID of check account

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CheckAccountApi.GetCheckAccountById(context.Background(), checkAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckAccountApi.GetCheckAccountById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCheckAccountById`: GetCheckAccountById200Response
    fmt.Fprintf(os.Stdout, "Response from `CheckAccountApi.GetCheckAccountById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkAccountId** | **int32** | ID of check account | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCheckAccountByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCheckAccountById200Response**](GetCheckAccountById200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCheckAccounts

> GetCheckAccountById200Response GetCheckAccounts(ctx).Execute()

Retrieve check accounts



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/jgillich/sevdesk-client-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CheckAccountApi.GetCheckAccounts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckAccountApi.GetCheckAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCheckAccounts`: GetCheckAccountById200Response
    fmt.Fprintf(os.Stdout, "Response from `CheckAccountApi.GetCheckAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCheckAccountsRequest struct via the builder pattern


### Return type

[**GetCheckAccountById200Response**](GetCheckAccountById200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCheckAccount

> GetCheckAccountById200Response UpdateCheckAccount(ctx, checkAccountId).ModelCheckAccountUpdate(modelCheckAccountUpdate).Execute()

Update an existing check account



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/jgillich/sevdesk-client-go"
)

func main() {
    checkAccountId := int32(56) // int32 | ID of check account to update
    modelCheckAccountUpdate := *openapiclient.NewModelCheckAccountUpdate() // ModelCheckAccountUpdate | Update data (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CheckAccountApi.UpdateCheckAccount(context.Background(), checkAccountId).ModelCheckAccountUpdate(modelCheckAccountUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckAccountApi.UpdateCheckAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCheckAccount`: GetCheckAccountById200Response
    fmt.Fprintf(os.Stdout, "Response from `CheckAccountApi.UpdateCheckAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkAccountId** | **int32** | ID of check account to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCheckAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelCheckAccountUpdate** | [**ModelCheckAccountUpdate**](ModelCheckAccountUpdate.md) | Update data | 

### Return type

[**GetCheckAccountById200Response**](GetCheckAccountById200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

