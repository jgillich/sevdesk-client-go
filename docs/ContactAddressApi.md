# \ContactAddressApi

All URIs are relative to *https://my.sevdesk.de/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContactAddressId**](ContactAddressApi.md#ContactAddressId) | **Get** /ContactAddress/{contactAddressId} | Find contact address by ID
[**CreateContactAddress**](ContactAddressApi.md#CreateContactAddress) | **Post** /ContactAddress | Create a new contact address
[**DeleteContactAddress**](ContactAddressApi.md#DeleteContactAddress) | **Delete** /ContactAddress/{contactAddressId} | Deletes a contact address
[**GetContactAddresses**](ContactAddressApi.md#GetContactAddresses) | **Get** /ContactAddress | Retrieve contact addresses
[**UpdateContactAddress**](ContactAddressApi.md#UpdateContactAddress) | **Put** /ContactAddress/{contactAddressId} | update a existing contact address



## ContactAddressId

> ContactAddressId200Response ContactAddressId(ctx, contactAddressId).Execute()

Find contact address by ID



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
    contactAddressId := int32(56) // int32 | ID of contact address to return

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactAddressApi.ContactAddressId(context.Background(), contactAddressId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactAddressApi.ContactAddressId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContactAddressId`: ContactAddressId200Response
    fmt.Fprintf(os.Stdout, "Response from `ContactAddressApi.ContactAddressId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactAddressId** | **int32** | ID of contact address to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiContactAddressIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContactAddressId200Response**](ContactAddressId200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContactAddress

> ContactAddressId200Response CreateContactAddress(ctx).ModelContactAddress(modelContactAddress).Execute()

Create a new contact address



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
    modelContactAddress := *openapiclient.NewModelContactAddress(*openapiclient.NewModelContactAddressContact(int32(123), "Contact"), *openapiclient.NewModelContactAddressCountry(int32(123), "StaticCountry")) // ModelContactAddress | Creation data (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactAddressApi.CreateContactAddress(context.Background()).ModelContactAddress(modelContactAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactAddressApi.CreateContactAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateContactAddress`: ContactAddressId200Response
    fmt.Fprintf(os.Stdout, "Response from `ContactAddressApi.CreateContactAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateContactAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelContactAddress** | [**ModelContactAddress**](ModelContactAddress.md) | Creation data | 

### Return type

[**ContactAddressId200Response**](ContactAddressId200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContactAddress

> DeleteAccountingContact200Response DeleteContactAddress(ctx, contactAddressId).Execute()

Deletes a contact address

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
    contactAddressId := int32(56) // int32 | Id of contact address resource to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactAddressApi.DeleteContactAddress(context.Background(), contactAddressId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactAddressApi.DeleteContactAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteContactAddress`: DeleteAccountingContact200Response
    fmt.Fprintf(os.Stdout, "Response from `ContactAddressApi.DeleteContactAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactAddressId** | **int32** | Id of contact address resource to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContactAddressRequest struct via the builder pattern


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


## GetContactAddresses

> ContactAddressId200Response GetContactAddresses(ctx).Execute()

Retrieve contact addresses



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
    resp, r, err := apiClient.ContactAddressApi.GetContactAddresses(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactAddressApi.GetContactAddresses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContactAddresses`: ContactAddressId200Response
    fmt.Fprintf(os.Stdout, "Response from `ContactAddressApi.GetContactAddresses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetContactAddressesRequest struct via the builder pattern


### Return type

[**ContactAddressId200Response**](ContactAddressId200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContactAddress

> ContactAddressId200Response UpdateContactAddress(ctx, contactAddressId).ModelContactAddressUpdate(modelContactAddressUpdate).Execute()

update a existing contact address



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
    contactAddressId := int32(56) // int32 | ID of contact address to return
    modelContactAddressUpdate := *openapiclient.NewModelContactAddressUpdate() // ModelContactAddressUpdate | Creation data (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactAddressApi.UpdateContactAddress(context.Background(), contactAddressId).ModelContactAddressUpdate(modelContactAddressUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactAddressApi.UpdateContactAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateContactAddress`: ContactAddressId200Response
    fmt.Fprintf(os.Stdout, "Response from `ContactAddressApi.UpdateContactAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactAddressId** | **int32** | ID of contact address to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContactAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelContactAddressUpdate** | [**ModelContactAddressUpdate**](ModelContactAddressUpdate.md) | Creation data | 

### Return type

[**ContactAddressId200Response**](ContactAddressId200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

