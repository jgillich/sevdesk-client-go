# \OrderPosApi

All URIs are relative to *https://my.sevdesk.de/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOrderPos**](OrderPosApi.md#DeleteOrderPos) | **Delete** /OrderPos/{orderPosId} | Deletes an order Position
[**GetOrderPositionById**](OrderPosApi.md#GetOrderPositionById) | **Get** /OrderPos/{orderPosId} | Find order position by ID
[**GetOrderPositions**](OrderPosApi.md#GetOrderPositions) | **Get** /OrderPos | Retrieve order positions
[**UpdateOrderPosition**](OrderPosApi.md#UpdateOrderPosition) | **Put** /OrderPos/{orderPosId} | Update an existing order position



## DeleteOrderPos

> DeleteAccountingContact200Response DeleteOrderPos(ctx, orderPosId).Execute()

Deletes an order Position

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
    orderPosId := int32(56) // int32 | Id of order position resource to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderPosApi.DeleteOrderPos(context.Background(), orderPosId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderPosApi.DeleteOrderPos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOrderPos`: DeleteAccountingContact200Response
    fmt.Fprintf(os.Stdout, "Response from `OrderPosApi.DeleteOrderPos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderPosId** | **int32** | Id of order position resource to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrderPosRequest struct via the builder pattern


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


## GetOrderPositionById

> GetOrderPositions200Response GetOrderPositionById(ctx, orderPosId).Execute()

Find order position by ID



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
    orderPosId := int32(56) // int32 | ID of order position to return

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderPosApi.GetOrderPositionById(context.Background(), orderPosId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderPosApi.GetOrderPositionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderPositionById`: GetOrderPositions200Response
    fmt.Fprintf(os.Stdout, "Response from `OrderPosApi.GetOrderPositionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderPosId** | **int32** | ID of order position to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderPositionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetOrderPositions200Response**](GetOrderPositions200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderPositions

> GetOrderPositions200Response GetOrderPositions(ctx).OrderId(orderId).OrderObjectName(orderObjectName).Execute()

Retrieve order positions



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
    orderId := int32(56) // int32 | Retrieve all order positions belonging to this order. Must be provided with voucher[objectName] (optional)
    orderObjectName := "orderObjectName_example" // string | Only required if order[id] was provided. 'Order' should be used as value. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderPosApi.GetOrderPositions(context.Background()).OrderId(orderId).OrderObjectName(orderObjectName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderPosApi.GetOrderPositions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderPositions`: GetOrderPositions200Response
    fmt.Fprintf(os.Stdout, "Response from `OrderPosApi.GetOrderPositions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderPositionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **int32** | Retrieve all order positions belonging to this order. Must be provided with voucher[objectName] | 
 **orderObjectName** | **string** | Only required if order[id] was provided. &#39;Order&#39; should be used as value. | 

### Return type

[**GetOrderPositions200Response**](GetOrderPositions200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrderPosition

> GetOrderPositions200Response UpdateOrderPosition(ctx, orderPosId).ModelOrderPosUpdate(modelOrderPosUpdate).Execute()

Update an existing order position



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
    orderPosId := int32(56) // int32 | ID of order position to update
    modelOrderPosUpdate := *openapiclient.NewModelOrderPosUpdate() // ModelOrderPosUpdate | Update data (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderPosApi.UpdateOrderPosition(context.Background(), orderPosId).ModelOrderPosUpdate(modelOrderPosUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderPosApi.UpdateOrderPosition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrderPosition`: GetOrderPositions200Response
    fmt.Fprintf(os.Stdout, "Response from `OrderPosApi.UpdateOrderPosition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderPosId** | **int32** | ID of order position to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrderPositionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelOrderPosUpdate** | [**ModelOrderPosUpdate**](ModelOrderPosUpdate.md) | Update data | 

### Return type

[**GetOrderPositions200Response**](GetOrderPositions200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

