# \InvoicePosApi

All URIs are relative to *https://my.sevdesk.de/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInvoicePos**](InvoicePosApi.md#GetInvoicePos) | **Get** /InvoicePos | Retrieve InvoicePos



## GetInvoicePos

> GetInvoicePos200Response GetInvoicePos(ctx).Id(id).InvoiceId(invoiceId).InvoiceObjectName(invoiceObjectName).PartId(partId).PartObjectName(partObjectName).Execute()

Retrieve InvoicePos



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
    id := float32(8.14) // float32 | Retrieve all InvoicePos with this InvoicePos id (optional)
    invoiceId := float32(8.14) // float32 | Retrieve all invoices positions with this invoice. Must be provided with invoice[objectName] (optional)
    invoiceObjectName := "invoiceObjectName_example" // string | Only required if invoice[id] was provided. 'Invoice' should be used as value. (optional)
    partId := float32(8.14) // float32 | Retrieve all invoices positions with this part. Must be provided with part[objectName] (optional)
    partObjectName := "partObjectName_example" // string | Only required if part[id] was provided. 'Part' should be used as value. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoicePosApi.GetInvoicePos(context.Background()).Id(id).InvoiceId(invoiceId).InvoiceObjectName(invoiceObjectName).PartId(partId).PartObjectName(partObjectName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicePosApi.GetInvoicePos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoicePos`: GetInvoicePos200Response
    fmt.Fprintf(os.Stdout, "Response from `InvoicePosApi.GetInvoicePos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoicePosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **float32** | Retrieve all InvoicePos with this InvoicePos id | 
 **invoiceId** | **float32** | Retrieve all invoices positions with this invoice. Must be provided with invoice[objectName] | 
 **invoiceObjectName** | **string** | Only required if invoice[id] was provided. &#39;Invoice&#39; should be used as value. | 
 **partId** | **float32** | Retrieve all invoices positions with this part. Must be provided with part[objectName] | 
 **partObjectName** | **string** | Only required if part[id] was provided. &#39;Part&#39; should be used as value. | 

### Return type

[**GetInvoicePos200Response**](GetInvoicePos200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

