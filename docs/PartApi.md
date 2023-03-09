# \PartApi

All URIs are relative to *https://my.sevdesk.de/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePart**](PartApi.md#CreatePart) | **Post** /Part | Create a new part
[**GetPartById**](PartApi.md#GetPartById) | **Get** /Part/{partId} | Find part by ID
[**GetParts**](PartApi.md#GetParts) | **Get** /Part | Retrieve parts
[**PartGetStock**](PartApi.md#PartGetStock) | **Get** /Part/{partId}/getStock | Get stock of a part
[**UpdatePart**](PartApi.md#UpdatePart) | **Put** /Part/{partId} | Update an existing part



## CreatePart

> GetPartById200Response CreatePart(ctx).ModelPart(modelPart).Execute()

Create a new part



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
    modelPart := *openapiclient.NewModelPart("Dragonglass", "Part-1000", float32(10000), *openapiclient.NewModelPartUnity(int32(1), "Unity"), float32(19)) // ModelPart | Creation data. Please be aware, that you need to provide at least all required parameter      of the part model! (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartApi.CreatePart(context.Background()).ModelPart(modelPart).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartApi.CreatePart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePart`: GetPartById200Response
    fmt.Fprintf(os.Stdout, "Response from `PartApi.CreatePart`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelPart** | [**ModelPart**](ModelPart.md) | Creation data. Please be aware, that you need to provide at least all required parameter      of the part model! | 

### Return type

[**GetPartById200Response**](GetPartById200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartById

> GetPartById200Response GetPartById(ctx, partId).Execute()

Find part by ID



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
    partId := int32(56) // int32 | ID of part to return

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartApi.GetPartById(context.Background(), partId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartApi.GetPartById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPartById`: GetPartById200Response
    fmt.Fprintf(os.Stdout, "Response from `PartApi.GetPartById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partId** | **int32** | ID of part to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetPartById200Response**](GetPartById200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetParts

> GetPartById200Response GetParts(ctx).PartNumber(partNumber).Name(name).Execute()

Retrieve parts



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
    partNumber := "partNumber_example" // string | Retrieve all parts with this part number (optional)
    name := "name_example" // string | Retrieve all parts with this name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartApi.GetParts(context.Background()).PartNumber(partNumber).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartApi.GetParts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetParts`: GetPartById200Response
    fmt.Fprintf(os.Stdout, "Response from `PartApi.GetParts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPartsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **partNumber** | **string** | Retrieve all parts with this part number | 
 **name** | **string** | Retrieve all parts with this name | 

### Return type

[**GetPartById200Response**](GetPartById200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartGetStock

> PartGetStock200Response PartGetStock(ctx, partId).Execute()

Get stock of a part



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
    partId := int32(56) // int32 | ID of part for which you want the current stock.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartApi.PartGetStock(context.Background(), partId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartApi.PartGetStock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartGetStock`: PartGetStock200Response
    fmt.Fprintf(os.Stdout, "Response from `PartApi.PartGetStock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partId** | **int32** | ID of part for which you want the current stock. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartGetStockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PartGetStock200Response**](PartGetStock200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePart

> GetPartById200Response UpdatePart(ctx, partId).ModelPartUpdate(modelPartUpdate).Execute()

Update an existing part



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
    partId := int32(56) // int32 | ID of part to update
    modelPartUpdate := *openapiclient.NewModelPartUpdate() // ModelPartUpdate | Update data (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartApi.UpdatePart(context.Background(), partId).ModelPartUpdate(modelPartUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartApi.UpdatePart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePart`: GetPartById200Response
    fmt.Fprintf(os.Stdout, "Response from `PartApi.UpdatePart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partId** | **int32** | ID of part to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelPartUpdate** | [**ModelPartUpdate**](ModelPartUpdate.md) | Update data | 

### Return type

[**GetPartById200Response**](GetPartById200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

