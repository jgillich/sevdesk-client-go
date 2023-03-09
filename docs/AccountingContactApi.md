# \AccountingContactApi

All URIs are relative to *https://my.sevdesk.de/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountingContact**](AccountingContactApi.md#CreateAccountingContact) | **Post** /AccountingContact | Create a new accounting contact
[**DeleteAccountingContact**](AccountingContactApi.md#DeleteAccountingContact) | **Delete** /AccountingContact/{accountingContactId} | Deletes an accounting contact
[**GetAccountingContact**](AccountingContactApi.md#GetAccountingContact) | **Get** /AccountingContact | Retrieve accounting contact
[**GetAccountingContactById**](AccountingContactApi.md#GetAccountingContactById) | **Get** /AccountingContact/{accountingContactId} | Find accounting contact by ID
[**UpdateAccountingContact**](AccountingContactApi.md#UpdateAccountingContact) | **Put** /AccountingContact/{accountingContactId} | Update an existing accounting contact



## CreateAccountingContact

> GetAccountingContactById200Response CreateAccountingContact(ctx).ModelAccountingContact(modelAccountingContact).Execute()

Create a new accounting contact



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
    modelAccountingContact := *openapiclient.NewModelAccountingContact(*openapiclient.NewModelAccountingContactContact(int32(123), "Contact")) // ModelAccountingContact | Creation data (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountingContactApi.CreateAccountingContact(context.Background()).ModelAccountingContact(modelAccountingContact).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountingContactApi.CreateAccountingContact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccountingContact`: GetAccountingContactById200Response
    fmt.Fprintf(os.Stdout, "Response from `AccountingContactApi.CreateAccountingContact`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountingContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelAccountingContact** | [**ModelAccountingContact**](ModelAccountingContact.md) | Creation data | 

### Return type

[**GetAccountingContactById200Response**](GetAccountingContactById200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccountingContact

> DeleteAccountingContact200Response DeleteAccountingContact(ctx, accountingContactId).Execute()

Deletes an accounting contact

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
    accountingContactId := int32(56) // int32 | Id of accounting contact resource to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountingContactApi.DeleteAccountingContact(context.Background(), accountingContactId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountingContactApi.DeleteAccountingContact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAccountingContact`: DeleteAccountingContact200Response
    fmt.Fprintf(os.Stdout, "Response from `AccountingContactApi.DeleteAccountingContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountingContactId** | **int32** | Id of accounting contact resource to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountingContactRequest struct via the builder pattern


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


## GetAccountingContact

> GetAccountingContactById200Response GetAccountingContact(ctx).ContactId(contactId).ContactObjectName(contactObjectName).Execute()

Retrieve accounting contact



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
    contactId := "contactId_example" // string | ID of contact for which you want the accounting contact. (optional)
    contactObjectName := "Contact" // string | Object name. Only needed if you also defined the ID of a contact. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountingContactApi.GetAccountingContact(context.Background()).ContactId(contactId).ContactObjectName(contactObjectName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountingContactApi.GetAccountingContact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountingContact`: GetAccountingContactById200Response
    fmt.Fprintf(os.Stdout, "Response from `AccountingContactApi.GetAccountingContact`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountingContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contactId** | **string** | ID of contact for which you want the accounting contact. | 
 **contactObjectName** | **string** | Object name. Only needed if you also defined the ID of a contact. | 

### Return type

[**GetAccountingContactById200Response**](GetAccountingContactById200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountingContactById

> GetAccountingContactById200Response GetAccountingContactById(ctx, accountingContactId).Execute()

Find accounting contact by ID



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
    accountingContactId := int32(56) // int32 | ID of accounting contact to return

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountingContactApi.GetAccountingContactById(context.Background(), accountingContactId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountingContactApi.GetAccountingContactById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountingContactById`: GetAccountingContactById200Response
    fmt.Fprintf(os.Stdout, "Response from `AccountingContactApi.GetAccountingContactById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountingContactId** | **int32** | ID of accounting contact to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountingContactByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAccountingContactById200Response**](GetAccountingContactById200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountingContact

> GetAccountingContactById200Response UpdateAccountingContact(ctx, accountingContactId).ModelAccountingContactUpdate(modelAccountingContactUpdate).Execute()

Update an existing accounting contact



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
    accountingContactId := int32(56) // int32 | ID of accounting contact to update
    modelAccountingContactUpdate := *openapiclient.NewModelAccountingContactUpdate() // ModelAccountingContactUpdate | Update data (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountingContactApi.UpdateAccountingContact(context.Background(), accountingContactId).ModelAccountingContactUpdate(modelAccountingContactUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountingContactApi.UpdateAccountingContact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccountingContact`: GetAccountingContactById200Response
    fmt.Fprintf(os.Stdout, "Response from `AccountingContactApi.UpdateAccountingContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountingContactId** | **int32** | ID of accounting contact to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountingContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelAccountingContactUpdate** | [**ModelAccountingContactUpdate**](ModelAccountingContactUpdate.md) | Update data | 

### Return type

[**GetAccountingContactById200Response**](GetAccountingContactById200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

