# \CheckAccountTransactionApi

All URIs are relative to *https://my.sevdesk.de/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTransaction**](CheckAccountTransactionApi.md#CreateTransaction) | **Post** /CheckAccountTransaction | Create a new transaction
[**DeleteCheckAccountTransaction**](CheckAccountTransactionApi.md#DeleteCheckAccountTransaction) | **Delete** /CheckAccountTransaction/{checkAccountTransactionId} | Deletes a check account transaction
[**GetCheckAccountTransactionById**](CheckAccountTransactionApi.md#GetCheckAccountTransactionById) | **Get** /CheckAccountTransaction/{checkAccountTransactionId} | Find check account transaction by ID
[**GetTransactions**](CheckAccountTransactionApi.md#GetTransactions) | **Get** /CheckAccountTransaction | Retrieve transactions
[**UpdateCheckAccountTransaction**](CheckAccountTransactionApi.md#UpdateCheckAccountTransaction) | **Put** /CheckAccountTransaction/{checkAccountTransactionId} | Update an existing check account transaction



## CreateTransaction

> GetCheckAccountTransactionById200Response CreateTransaction(ctx).ModelCheckAccountTransaction(modelCheckAccountTransaction).Execute()

Create a new transaction



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
    modelCheckAccountTransaction := *openapiclient.NewModelCheckAccountTransaction(time.Now(), float32(100), "Cercei Lannister", *openapiclient.NewModelCheckAccountTransactionCheckAccount(int32(123), "ObjectName_example"), int32(123)) // ModelCheckAccountTransaction | Creation data. Please be aware, that you need to provide at least all required parameter      of the CheckAccountTransaction model! (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CheckAccountTransactionApi.CreateTransaction(context.Background()).ModelCheckAccountTransaction(modelCheckAccountTransaction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckAccountTransactionApi.CreateTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTransaction`: GetCheckAccountTransactionById200Response
    fmt.Fprintf(os.Stdout, "Response from `CheckAccountTransactionApi.CreateTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelCheckAccountTransaction** | [**ModelCheckAccountTransaction**](ModelCheckAccountTransaction.md) | Creation data. Please be aware, that you need to provide at least all required parameter      of the CheckAccountTransaction model! | 

### Return type

[**GetCheckAccountTransactionById200Response**](GetCheckAccountTransactionById200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCheckAccountTransaction

> DeleteAccountingContact200Response DeleteCheckAccountTransaction(ctx, checkAccountTransactionId).Execute()

Deletes a check account transaction

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
    checkAccountTransactionId := int32(56) // int32 | Id of check account transaction to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CheckAccountTransactionApi.DeleteCheckAccountTransaction(context.Background(), checkAccountTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckAccountTransactionApi.DeleteCheckAccountTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCheckAccountTransaction`: DeleteAccountingContact200Response
    fmt.Fprintf(os.Stdout, "Response from `CheckAccountTransactionApi.DeleteCheckAccountTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkAccountTransactionId** | **int32** | Id of check account transaction to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCheckAccountTransactionRequest struct via the builder pattern


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


## GetCheckAccountTransactionById

> GetCheckAccountTransactionById200Response GetCheckAccountTransactionById(ctx, checkAccountTransactionId).Execute()

Find check account transaction by ID



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
    checkAccountTransactionId := int32(56) // int32 | ID of check account transaction

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CheckAccountTransactionApi.GetCheckAccountTransactionById(context.Background(), checkAccountTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckAccountTransactionApi.GetCheckAccountTransactionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCheckAccountTransactionById`: GetCheckAccountTransactionById200Response
    fmt.Fprintf(os.Stdout, "Response from `CheckAccountTransactionApi.GetCheckAccountTransactionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkAccountTransactionId** | **int32** | ID of check account transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCheckAccountTransactionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCheckAccountTransactionById200Response**](GetCheckAccountTransactionById200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactions

> GetCheckAccountTransactionById200Response GetTransactions(ctx).CheckAccountId(checkAccountId).CheckAccountObjectName(checkAccountObjectName).IsBooked(isBooked).PaymtPurpose(paymtPurpose).StartDate(startDate).EndDate(endDate).PayeePayerName(payeePayerName).OnlyCredit(onlyCredit).OnlyDebit(onlyDebit).Execute()

Retrieve transactions



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
    checkAccountId := int32(56) // int32 | Retrieve all transactions on this check account. Must be provided with checkAccount[objectName] (optional)
    checkAccountObjectName := "checkAccountObjectName_example" // string | Only required if checkAccount[id] was provided. 'CheckAccount' should be used as value. (optional)
    isBooked := true // bool | Only retrieve booked transactions (optional)
    paymtPurpose := "paymtPurpose_example" // string | Only retrieve transactions with this payment purpose (optional)
    startDate := time.Now() // time.Time | Only retrieve transactions from this date on (optional)
    endDate := time.Now() // time.Time | Only retrieve transactions up to this date (optional)
    payeePayerName := "payeePayerName_example" // string | Only retrieve transactions with this payee / payer (optional)
    onlyCredit := true // bool | Only retrieve credit transactions (optional)
    onlyDebit := true // bool | Only retrieve debit transactions (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CheckAccountTransactionApi.GetTransactions(context.Background()).CheckAccountId(checkAccountId).CheckAccountObjectName(checkAccountObjectName).IsBooked(isBooked).PaymtPurpose(paymtPurpose).StartDate(startDate).EndDate(endDate).PayeePayerName(payeePayerName).OnlyCredit(onlyCredit).OnlyDebit(onlyDebit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckAccountTransactionApi.GetTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransactions`: GetCheckAccountTransactionById200Response
    fmt.Fprintf(os.Stdout, "Response from `CheckAccountTransactionApi.GetTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkAccountId** | **int32** | Retrieve all transactions on this check account. Must be provided with checkAccount[objectName] | 
 **checkAccountObjectName** | **string** | Only required if checkAccount[id] was provided. &#39;CheckAccount&#39; should be used as value. | 
 **isBooked** | **bool** | Only retrieve booked transactions | 
 **paymtPurpose** | **string** | Only retrieve transactions with this payment purpose | 
 **startDate** | **time.Time** | Only retrieve transactions from this date on | 
 **endDate** | **time.Time** | Only retrieve transactions up to this date | 
 **payeePayerName** | **string** | Only retrieve transactions with this payee / payer | 
 **onlyCredit** | **bool** | Only retrieve credit transactions | 
 **onlyDebit** | **bool** | Only retrieve debit transactions | 

### Return type

[**GetCheckAccountTransactionById200Response**](GetCheckAccountTransactionById200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCheckAccountTransaction

> GetCheckAccountTransactionById200Response UpdateCheckAccountTransaction(ctx, checkAccountTransactionId).ModelCheckAccountTransactionUpdate(modelCheckAccountTransactionUpdate).Execute()

Update an existing check account transaction



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
    checkAccountTransactionId := int32(56) // int32 | ID of check account to update transaction
    modelCheckAccountTransactionUpdate := *openapiclient.NewModelCheckAccountTransactionUpdate() // ModelCheckAccountTransactionUpdate | Update data (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CheckAccountTransactionApi.UpdateCheckAccountTransaction(context.Background(), checkAccountTransactionId).ModelCheckAccountTransactionUpdate(modelCheckAccountTransactionUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckAccountTransactionApi.UpdateCheckAccountTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCheckAccountTransaction`: GetCheckAccountTransactionById200Response
    fmt.Fprintf(os.Stdout, "Response from `CheckAccountTransactionApi.UpdateCheckAccountTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkAccountTransactionId** | **int32** | ID of check account to update transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCheckAccountTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelCheckAccountTransactionUpdate** | [**ModelCheckAccountTransactionUpdate**](ModelCheckAccountTransactionUpdate.md) | Update data | 

### Return type

[**GetCheckAccountTransactionById200Response**](GetCheckAccountTransactionById200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

