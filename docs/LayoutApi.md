# \LayoutApi

All URIs are relative to *https://my.sevdesk.de/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLetterpapersWithThumb**](LayoutApi.md#GetLetterpapersWithThumb) | **Get** /DocServer/getLetterpapersWithThumb | Retrieve letterpapers
[**GetTemplates**](LayoutApi.md#GetTemplates) | **Get** /DocServer/getTemplatesWithThumb | Retrieve templates
[**UpdateCreditNoteTemplate**](LayoutApi.md#UpdateCreditNoteTemplate) | **Put** /CreditNote/{creditNoteId}/changeParameter | Update an of credit note template
[**UpdateInvoiceTemplate**](LayoutApi.md#UpdateInvoiceTemplate) | **Put** /Invoice/{invoiceId}/changeParameter | Update an invoice template
[**UpdateOrderTemplate**](LayoutApi.md#UpdateOrderTemplate) | **Put** /Order/{orderId}/changeParameter | Update an order template



## GetLetterpapersWithThumb

> GetLetterpapersWithThumb200Response GetLetterpapersWithThumb(ctx).Execute()

Retrieve letterpapers



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
    resp, r, err := apiClient.LayoutApi.GetLetterpapersWithThumb(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LayoutApi.GetLetterpapersWithThumb``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLetterpapersWithThumb`: GetLetterpapersWithThumb200Response
    fmt.Fprintf(os.Stdout, "Response from `LayoutApi.GetLetterpapersWithThumb`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLetterpapersWithThumbRequest struct via the builder pattern


### Return type

[**GetLetterpapersWithThumb200Response**](GetLetterpapersWithThumb200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplates

> GetTemplates200Response GetTemplates(ctx).Type_(type_).Execute()

Retrieve templates



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
    type_ := "type__example" // string | Type of the templates you want to get. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LayoutApi.GetTemplates(context.Background()).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LayoutApi.GetTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTemplates`: GetTemplates200Response
    fmt.Fprintf(os.Stdout, "Response from `LayoutApi.GetTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Type of the templates you want to get. | 

### Return type

[**GetTemplates200Response**](GetTemplates200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCreditNoteTemplate

> UpdateOrderTemplate200Response UpdateCreditNoteTemplate(ctx, creditNoteId).ModelChangeLayout(modelChangeLayout).Execute()

Update an of credit note template



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
    creditNoteId := int32(56) // int32 | ID of credit note to update
    modelChangeLayout := *openapiclient.NewModelChangeLayout("template", "573ef2706bd2d50366786471") // ModelChangeLayout | Change Layout (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LayoutApi.UpdateCreditNoteTemplate(context.Background(), creditNoteId).ModelChangeLayout(modelChangeLayout).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LayoutApi.UpdateCreditNoteTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCreditNoteTemplate`: UpdateOrderTemplate200Response
    fmt.Fprintf(os.Stdout, "Response from `LayoutApi.UpdateCreditNoteTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**creditNoteId** | **int32** | ID of credit note to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCreditNoteTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelChangeLayout** | [**ModelChangeLayout**](ModelChangeLayout.md) | Change Layout | 

### Return type

[**UpdateOrderTemplate200Response**](UpdateOrderTemplate200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInvoiceTemplate

> UpdateOrderTemplate200Response UpdateInvoiceTemplate(ctx, invoiceId).ModelChangeLayout(modelChangeLayout).Execute()

Update an invoice template



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
    invoiceId := int32(56) // int32 | ID of invoice to update
    modelChangeLayout := *openapiclient.NewModelChangeLayout("template", "573ef2706bd2d50366786471") // ModelChangeLayout | Change Layout (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LayoutApi.UpdateInvoiceTemplate(context.Background(), invoiceId).ModelChangeLayout(modelChangeLayout).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LayoutApi.UpdateInvoiceTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInvoiceTemplate`: UpdateOrderTemplate200Response
    fmt.Fprintf(os.Stdout, "Response from `LayoutApi.UpdateInvoiceTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **int32** | ID of invoice to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInvoiceTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelChangeLayout** | [**ModelChangeLayout**](ModelChangeLayout.md) | Change Layout | 

### Return type

[**UpdateOrderTemplate200Response**](UpdateOrderTemplate200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrderTemplate

> UpdateOrderTemplate200Response UpdateOrderTemplate(ctx, orderId).ModelChangeLayout(modelChangeLayout).Execute()

Update an order template



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
    modelChangeLayout := *openapiclient.NewModelChangeLayout("template", "573ef2706bd2d50366786471") // ModelChangeLayout | Change Layout (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LayoutApi.UpdateOrderTemplate(context.Background(), orderId).ModelChangeLayout(modelChangeLayout).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LayoutApi.UpdateOrderTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrderTemplate`: UpdateOrderTemplate200Response
    fmt.Fprintf(os.Stdout, "Response from `LayoutApi.UpdateOrderTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int32** | ID of order to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrderTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelChangeLayout** | [**ModelChangeLayout**](ModelChangeLayout.md) | Change Layout | 

### Return type

[**UpdateOrderTemplate200Response**](UpdateOrderTemplate200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

