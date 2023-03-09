# \OrderApi

All URIs are relative to *https://my.sevdesk.de/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContractNoteFromOrder**](OrderApi.md#CreateContractNoteFromOrder) | **Post** /Order/Factory/createContractNoteFromOrder | Create contract note from order
[**CreateOrder**](OrderApi.md#CreateOrder) | **Post** /Order/Factory/saveOrder | Create a new order
[**CreatePackingListFromOrder**](OrderApi.md#CreatePackingListFromOrder) | **Post** /Order/Factory/createPackingListFromOrder | Create packing list from order
[**DeleteOrder**](OrderApi.md#DeleteOrder) | **Delete** /Order/{orderId} | Deletes an order
[**GetDiscounts**](OrderApi.md#GetDiscounts) | **Get** /Order/{orderId}/getDiscounts | Find order discounts
[**GetOrderById**](OrderApi.md#GetOrderById) | **Get** /Order/{orderId} | Find order by ID
[**GetOrderPositionsById**](OrderApi.md#GetOrderPositionsById) | **Get** /Order/{orderId}/getPositions | Find order positions
[**GetOrders**](OrderApi.md#GetOrders) | **Get** /Order | Retrieve orders
[**GetRelatedObjects**](OrderApi.md#GetRelatedObjects) | **Get** /Order/{orderId}/getRelatedObjects | Find related objects
[**OrderGetPdf**](OrderApi.md#OrderGetPdf) | **Get** /Order/{orderId}/getPdf | Retrieve pdf document of an order
[**OrderSendBy**](OrderApi.md#OrderSendBy) | **Put** /Order/{orderId}/sendBy | Mark order as sent
[**SendorderViaEMail**](OrderApi.md#SendorderViaEMail) | **Post** /Order/{orderId}/sendViaEmail | Send order via email
[**UpdateOrder**](OrderApi.md#UpdateOrder) | **Put** /Order/{orderId} | Update an existing order



## CreateContractNoteFromOrder

> CreatePackingListFromOrder200Response CreateContractNoteFromOrder(ctx).OrderId(orderId).OrderObjectName(orderObjectName).ModelCreatePackingListFromOrder(modelCreatePackingListFromOrder).Execute()

Create contract note from order



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
    orderId := int32(56) // int32 | the id of the order
    orderObjectName := "Order" // string | Model name, which is 'Order'
    modelCreatePackingListFromOrder := *openapiclient.NewModelCreatePackingListFromOrder(int32(123), "Order") // ModelCreatePackingListFromOrder | Create contract note (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.CreateContractNoteFromOrder(context.Background()).OrderId(orderId).OrderObjectName(orderObjectName).ModelCreatePackingListFromOrder(modelCreatePackingListFromOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.CreateContractNoteFromOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateContractNoteFromOrder`: CreatePackingListFromOrder200Response
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.CreateContractNoteFromOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateContractNoteFromOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **int32** | the id of the order | 
 **orderObjectName** | **string** | Model name, which is &#39;Order&#39; | 
 **modelCreatePackingListFromOrder** | [**ModelCreatePackingListFromOrder**](ModelCreatePackingListFromOrder.md) | Create contract note | 

### Return type

[**CreatePackingListFromOrder200Response**](CreatePackingListFromOrder200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrder

> SaveOrderResponse CreateOrder(ctx).SaveOrder(saveOrder).Execute()

Create a new order



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
    saveOrder := *openapiclient.NewSaveOrder(*openapiclient.NewModelOrder(false, "Offer-1000", *openapiclient.NewModelOrderContact(int32(123), "Contact"), time.Now(), int32(100), "My Offer-1000", *openapiclient.NewModelOrderAddressCountry(int32(123), "StaticCountry"), int32(0), *openapiclient.NewModelOrderContactPerson(int32(0), "SevUser"), float32(19), "Umsatzsteuer 19%", "default", "EUR")) // SaveOrder | Creation data. Please be aware, that you need to provide at least all required parameter      of the order model! (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.CreateOrder(context.Background()).SaveOrder(saveOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.CreateOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrder`: SaveOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.CreateOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **saveOrder** | [**SaveOrder**](SaveOrder.md) | Creation data. Please be aware, that you need to provide at least all required parameter      of the order model! | 

### Return type

[**SaveOrderResponse**](SaveOrderResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePackingListFromOrder

> CreatePackingListFromOrder200Response CreatePackingListFromOrder(ctx).OrderId(orderId).OrderObjectName(orderObjectName).ModelCreatePackingListFromOrder(modelCreatePackingListFromOrder).Execute()

Create packing list from order



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
    orderId := int32(56) // int32 | the id of the order
    orderObjectName := "Order" // string | Model name, which is 'Order'
    modelCreatePackingListFromOrder := *openapiclient.NewModelCreatePackingListFromOrder(int32(123), "Order") // ModelCreatePackingListFromOrder | Create packing list (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.CreatePackingListFromOrder(context.Background()).OrderId(orderId).OrderObjectName(orderObjectName).ModelCreatePackingListFromOrder(modelCreatePackingListFromOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.CreatePackingListFromOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePackingListFromOrder`: CreatePackingListFromOrder200Response
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.CreatePackingListFromOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePackingListFromOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **int32** | the id of the order | 
 **orderObjectName** | **string** | Model name, which is &#39;Order&#39; | 
 **modelCreatePackingListFromOrder** | [**ModelCreatePackingListFromOrder**](ModelCreatePackingListFromOrder.md) | Create packing list | 

### Return type

[**CreatePackingListFromOrder200Response**](CreatePackingListFromOrder200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrder

> DeleteAccountingContact200Response DeleteOrder(ctx, orderId).Execute()

Deletes an order

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
    orderId := int32(56) // int32 | Id of order resource to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.DeleteOrder(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.DeleteOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOrder`: DeleteAccountingContact200Response
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.DeleteOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int32** | Id of order resource to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrderRequest struct via the builder pattern


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


## GetDiscounts

> GetDiscounts200Response GetDiscounts(ctx, orderId).Limit(limit).Offset(offset).Embed(embed).Execute()

Find order discounts



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
    orderId := int32(56) // int32 | ID of order to return the positions
    limit := int32(56) // int32 | limits the number of entries returned (optional)
    offset := int32(56) // int32 | set the index where the returned entries start (optional)
    embed := []string{"Inner_example"} // []string | Get some additional information. Embed can handle multiple values, they must be separated by comma. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.GetDiscounts(context.Background(), orderId).Limit(limit).Offset(offset).Embed(embed).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.GetDiscounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDiscounts`: GetDiscounts200Response
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.GetDiscounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int32** | ID of order to return the positions | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiscountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | limits the number of entries returned | 
 **offset** | **int32** | set the index where the returned entries start | 
 **embed** | **[]string** | Get some additional information. Embed can handle multiple values, they must be separated by comma. | 

### Return type

[**GetDiscounts200Response**](GetDiscounts200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderById

> CreatePackingListFromOrder200Response GetOrderById(ctx, orderId).Execute()

Find order by ID



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
    orderId := int32(56) // int32 | ID of order to return

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.GetOrderById(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.GetOrderById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderById`: CreatePackingListFromOrder200Response
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.GetOrderById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int32** | ID of order to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreatePackingListFromOrder200Response**](CreatePackingListFromOrder200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderPositionsById

> GetOrderPositions200Response GetOrderPositionsById(ctx, orderId).Limit(limit).Offset(offset).Embed(embed).Execute()

Find order positions



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
    orderId := int32(56) // int32 | ID of order to return the positions
    limit := int32(56) // int32 | limits the number of entries returned (optional)
    offset := int32(56) // int32 | set the index where the returned entries start (optional)
    embed := []string{"Inner_example"} // []string | Get some additional information. Embed can handle multiple values, they must be separated by comma. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.GetOrderPositionsById(context.Background(), orderId).Limit(limit).Offset(offset).Embed(embed).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.GetOrderPositionsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderPositionsById`: GetOrderPositions200Response
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.GetOrderPositionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int32** | ID of order to return the positions | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderPositionsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | limits the number of entries returned | 
 **offset** | **int32** | set the index where the returned entries start | 
 **embed** | **[]string** | Get some additional information. Embed can handle multiple values, they must be separated by comma. | 

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


## GetOrders

> CreatePackingListFromOrder200Response GetOrders(ctx).Status(status).OrderNumber(orderNumber).StartDate(startDate).EndDate(endDate).ContactId(contactId).ContactObjectName(contactObjectName).Execute()

Retrieve orders



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
    status := int32(56) // int32 | Status of the order (optional)
    orderNumber := "orderNumber_example" // string | Retrieve all orders with this order number (optional)
    startDate := int32(56) // int32 | Retrieve all orders with a date equal or higher (optional)
    endDate := int32(56) // int32 | Retrieve all orders with a date equal or lower (optional)
    contactId := int32(56) // int32 | Retrieve all orders with this contact. Must be provided with contact[objectName] (optional)
    contactObjectName := "contactObjectName_example" // string | Only required if contact[id] was provided. 'Contact' should be used as value. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.GetOrders(context.Background()).Status(status).OrderNumber(orderNumber).StartDate(startDate).EndDate(endDate).ContactId(contactId).ContactObjectName(contactObjectName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.GetOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrders`: CreatePackingListFromOrder200Response
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.GetOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **int32** | Status of the order | 
 **orderNumber** | **string** | Retrieve all orders with this order number | 
 **startDate** | **int32** | Retrieve all orders with a date equal or higher | 
 **endDate** | **int32** | Retrieve all orders with a date equal or lower | 
 **contactId** | **int32** | Retrieve all orders with this contact. Must be provided with contact[objectName] | 
 **contactObjectName** | **string** | Only required if contact[id] was provided. &#39;Contact&#39; should be used as value. | 

### Return type

[**CreatePackingListFromOrder200Response**](CreatePackingListFromOrder200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRelatedObjects

> GetRelatedObjects200Response GetRelatedObjects(ctx, orderId).IncludeItself(includeItself).SortByType(sortByType).Embed(embed).Execute()

Find related objects



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
    orderId := int32(56) // int32 | ID of order to return the positions
    includeItself := true // bool | Define if the related objects include the order itself (optional)
    sortByType := true // bool | Define if you want the related objects sorted by type (optional)
    embed := []string{"Inner_example"} // []string | Get some additional information. Embed can handle multiple values, they must be separated by comma. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.GetRelatedObjects(context.Background(), orderId).IncludeItself(includeItself).SortByType(sortByType).Embed(embed).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.GetRelatedObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRelatedObjects`: GetRelatedObjects200Response
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.GetRelatedObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int32** | ID of order to return the positions | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRelatedObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeItself** | **bool** | Define if the related objects include the order itself | 
 **sortByType** | **bool** | Define if you want the related objects sorted by type | 
 **embed** | **[]string** | Get some additional information. Embed can handle multiple values, they must be separated by comma. | 

### Return type

[**GetRelatedObjects200Response**](GetRelatedObjects200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderGetPdf

> InvoiceGetPdf200Response OrderGetPdf(ctx, orderId).Download(download).PreventSendBy(preventSendBy).Execute()

Retrieve pdf document of an order



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
    orderId := int32(56) // int32 | ID of order from which you want the pdf
    download := true // bool | If u want to download the pdf of the order. (optional)
    preventSendBy := true // bool | Defines if u want to send the order. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderGetPdf(context.Background(), orderId).Download(download).PreventSendBy(preventSendBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderGetPdf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderGetPdf`: InvoiceGetPdf200Response
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderGetPdf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int32** | ID of order from which you want the pdf | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrderGetPdfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **download** | **bool** | If u want to download the pdf of the order. | 
 **preventSendBy** | **bool** | Defines if u want to send the order. | 

### Return type

[**InvoiceGetPdf200Response**](InvoiceGetPdf200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderSendBy

> OrderSendBy200Response OrderSendBy(ctx, orderId).OrderSendByRequest(orderSendByRequest).Execute()

Mark order as sent



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
    orderId := int32(56) // int32 | ID of order to mark as sent
    orderSendByRequest := *openapiclient.NewOrderSendByRequest("SendType_example", false) // OrderSendByRequest | Specify the send type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderSendBy(context.Background(), orderId).OrderSendByRequest(orderSendByRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderSendBy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderSendBy`: OrderSendBy200Response
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderSendBy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int32** | ID of order to mark as sent | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrderSendByRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orderSendByRequest** | [**OrderSendByRequest**](OrderSendByRequest.md) | Specify the send type | 

### Return type

[**OrderSendBy200Response**](OrderSendBy200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendorderViaEMail

> SendorderViaEMail201Response SendorderViaEMail(ctx, orderId).SendInvoiceViaEMailRequest(sendInvoiceViaEMailRequest).Execute()

Send order via email



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
    orderId := int32(56) // int32 | ID of order to be sent via email
    sendInvoiceViaEMailRequest := *openapiclient.NewSendInvoiceViaEMailRequest("ToEmail_example", "Subject_example", "Text_example") // SendInvoiceViaEMailRequest | Mail data (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.SendorderViaEMail(context.Background(), orderId).SendInvoiceViaEMailRequest(sendInvoiceViaEMailRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.SendorderViaEMail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendorderViaEMail`: SendorderViaEMail201Response
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.SendorderViaEMail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int32** | ID of order to be sent via email | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendorderViaEMailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendInvoiceViaEMailRequest** | [**SendInvoiceViaEMailRequest**](SendInvoiceViaEMailRequest.md) | Mail data | 

### Return type

[**SendorderViaEMail201Response**](SendorderViaEMail201Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrder

> CreatePackingListFromOrder200Response UpdateOrder(ctx, orderId).ModelOrderUpdate(modelOrderUpdate).Execute()

Update an existing order



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
    orderId := int32(56) // int32 | ID of order to update
    modelOrderUpdate := *openapiclient.NewModelOrderUpdate() // ModelOrderUpdate | Update data (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.UpdateOrder(context.Background(), orderId).ModelOrderUpdate(modelOrderUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.UpdateOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrder`: CreatePackingListFromOrder200Response
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.UpdateOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int32** | ID of order to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelOrderUpdate** | [**ModelOrderUpdate**](ModelOrderUpdate.md) | Update data | 

### Return type

[**CreatePackingListFromOrder200Response**](CreatePackingListFromOrder200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

