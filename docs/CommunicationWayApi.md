# \CommunicationWayApi

All URIs are relative to *https://my.sevdesk.de/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCommunicationWay**](CommunicationWayApi.md#CreateCommunicationWay) | **Post** /CommunicationWay | Create a new contact communication way
[**DeleteCommunicationWay**](CommunicationWayApi.md#DeleteCommunicationWay) | **Delete** /CommunicationWay/{communicationWayId} | Deletes a communication way
[**GetCommunicationWayById**](CommunicationWayApi.md#GetCommunicationWayById) | **Get** /CommunicationWay/{communicationWayId} | Find communication way by ID
[**GetCommunicationWayKeys**](CommunicationWayApi.md#GetCommunicationWayKeys) | **Get** /CommunicationWayKey | Retrieve communication way keys
[**GetCommunicationWays**](CommunicationWayApi.md#GetCommunicationWays) | **Get** /CommunicationWay | Retrieve communication ways
[**UpdateCommunicationWay**](CommunicationWayApi.md#UpdateCommunicationWay) | **Put** /CommunicationWay/{communicationWayId} | Update a existing communication way



## CreateCommunicationWay

> GetCommunicationWayById200Response CreateCommunicationWay(ctx).ModelCommunicationWay(modelCommunicationWay).Execute()

Create a new contact communication way



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
    modelCommunicationWay := *openapiclient.NewModelCommunicationWay("EMAIL", "Value_example", *openapiclient.NewModelCommunicationWayKey(int32(123), "CommunicationWayKey")) // ModelCommunicationWay | Creation data (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CommunicationWayApi.CreateCommunicationWay(context.Background()).ModelCommunicationWay(modelCommunicationWay).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationWayApi.CreateCommunicationWay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCommunicationWay`: GetCommunicationWayById200Response
    fmt.Fprintf(os.Stdout, "Response from `CommunicationWayApi.CreateCommunicationWay`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCommunicationWayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelCommunicationWay** | [**ModelCommunicationWay**](ModelCommunicationWay.md) | Creation data | 

### Return type

[**GetCommunicationWayById200Response**](GetCommunicationWayById200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCommunicationWay

> DeleteAccountingContact200Response DeleteCommunicationWay(ctx, communicationWayId).Execute()

Deletes a communication way

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
    communicationWayId := int32(56) // int32 | Id of communication way resource to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CommunicationWayApi.DeleteCommunicationWay(context.Background(), communicationWayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationWayApi.DeleteCommunicationWay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCommunicationWay`: DeleteAccountingContact200Response
    fmt.Fprintf(os.Stdout, "Response from `CommunicationWayApi.DeleteCommunicationWay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**communicationWayId** | **int32** | Id of communication way resource to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommunicationWayRequest struct via the builder pattern


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


## GetCommunicationWayById

> GetCommunicationWayById200Response GetCommunicationWayById(ctx, communicationWayId).Execute()

Find communication way by ID



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
    communicationWayId := int32(56) // int32 | ID of communication way to return

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CommunicationWayApi.GetCommunicationWayById(context.Background(), communicationWayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationWayApi.GetCommunicationWayById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCommunicationWayById`: GetCommunicationWayById200Response
    fmt.Fprintf(os.Stdout, "Response from `CommunicationWayApi.GetCommunicationWayById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**communicationWayId** | **int32** | ID of communication way to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommunicationWayByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCommunicationWayById200Response**](GetCommunicationWayById200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommunicationWayKeys

> GetCommunicationWayKeys200Response GetCommunicationWayKeys(ctx).Execute()

Retrieve communication way keys



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
    resp, r, err := apiClient.CommunicationWayApi.GetCommunicationWayKeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationWayApi.GetCommunicationWayKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCommunicationWayKeys`: GetCommunicationWayKeys200Response
    fmt.Fprintf(os.Stdout, "Response from `CommunicationWayApi.GetCommunicationWayKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommunicationWayKeysRequest struct via the builder pattern


### Return type

[**GetCommunicationWayKeys200Response**](GetCommunicationWayKeys200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommunicationWays

> GetCommunicationWayById200Response GetCommunicationWays(ctx).ContactId(contactId).ContactObjectName(contactObjectName).Type_(type_).Main(main).Execute()

Retrieve communication ways



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
    contactId := "contactId_example" // string | ID of contact for which you want the communication ways. (optional)
    contactObjectName := "Contact" // string | Object name. Only needed if you also defined the ID of a contact. (optional)
    type_ := "type__example" // string | Type of the communication ways you want to get. (optional)
    main := "main_example" // string | Define if you only want the main communication way. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CommunicationWayApi.GetCommunicationWays(context.Background()).ContactId(contactId).ContactObjectName(contactObjectName).Type_(type_).Main(main).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationWayApi.GetCommunicationWays``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCommunicationWays`: GetCommunicationWayById200Response
    fmt.Fprintf(os.Stdout, "Response from `CommunicationWayApi.GetCommunicationWays`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCommunicationWaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contactId** | **string** | ID of contact for which you want the communication ways. | 
 **contactObjectName** | **string** | Object name. Only needed if you also defined the ID of a contact. | 
 **type_** | **string** | Type of the communication ways you want to get. | 
 **main** | **string** | Define if you only want the main communication way. | 

### Return type

[**GetCommunicationWayById200Response**](GetCommunicationWayById200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCommunicationWay

> GetCommunicationWayById200Response UpdateCommunicationWay(ctx, communicationWayId).ModelCommunicationWayUpdate(modelCommunicationWayUpdate).Execute()

Update a existing communication way



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
    communicationWayId := int32(56) // int32 | ID of CommunicationWay to update
    modelCommunicationWayUpdate := *openapiclient.NewModelCommunicationWayUpdate() // ModelCommunicationWayUpdate | Update data (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CommunicationWayApi.UpdateCommunicationWay(context.Background(), communicationWayId).ModelCommunicationWayUpdate(modelCommunicationWayUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationWayApi.UpdateCommunicationWay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCommunicationWay`: GetCommunicationWayById200Response
    fmt.Fprintf(os.Stdout, "Response from `CommunicationWayApi.UpdateCommunicationWay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**communicationWayId** | **int32** | ID of CommunicationWay to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCommunicationWayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelCommunicationWayUpdate** | [**ModelCommunicationWayUpdate**](ModelCommunicationWayUpdate.md) | Update data | 

### Return type

[**GetCommunicationWayById200Response**](GetCommunicationWayById200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

