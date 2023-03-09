# \InvoiceApi

All URIs are relative to *https://my.sevdesk.de/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BookInvoice**](InvoiceApi.md#BookInvoice) | **Put** /Invoice/{invoiceId}/bookAmount | Book an invoice
[**CancelInvoice**](InvoiceApi.md#CancelInvoice) | **Post** /Invoice/{invoiceId}/cancelInvoice | Cancel an invoice / Create cancellation invoice
[**CreateInvoiceByFactory**](InvoiceApi.md#CreateInvoiceByFactory) | **Post** /Invoice/Factory/saveInvoice | Create a new invoice
[**CreateInvoiceFromOrder**](InvoiceApi.md#CreateInvoiceFromOrder) | **Post** /Invoice/Factory/createInvoiceFromOrder | Create invoice from order
[**CreateInvoiceReminder**](InvoiceApi.md#CreateInvoiceReminder) | **Post** /Invoice/Factory/createInvoiceReminder | Create invoice reminder
[**GetInvoiceById**](InvoiceApi.md#GetInvoiceById) | **Get** /Invoice/{invoiceId} | Find invoice by ID
[**GetInvoicePositionsById**](InvoiceApi.md#GetInvoicePositionsById) | **Get** /Invoice/{invoiceId}/getPositions | Find invoice positions
[**GetInvoices**](InvoiceApi.md#GetInvoices) | **Get** /Invoice | Retrieve invoices
[**GetIsInvoicePartiallyPaid**](InvoiceApi.md#GetIsInvoicePartiallyPaid) | **Get** /Invoice/{invoiceId}/getIsPartiallyPaid | Check if an invoice is already partially paid
[**InvoiceGetPdf**](InvoiceApi.md#InvoiceGetPdf) | **Get** /Invoice/{invoiceId}/getPdf | Retrieve pdf document of an invoice
[**InvoiceRender**](InvoiceApi.md#InvoiceRender) | **Post** /Invoice/{invoiceId}/render | Render the pdf document of an invoice
[**InvoiceSendBy**](InvoiceApi.md#InvoiceSendBy) | **Put** /Invoice/{invoiceId}/sendBy | Mark invoice as sent
[**SendInvoiceViaEMail**](InvoiceApi.md#SendInvoiceViaEMail) | **Post** /Invoice/{invoiceId}/sendViaEmail | Send invoice via email



## BookInvoice

> BookInvoice200Response BookInvoice(ctx, invoiceId).BookInvoiceRequest(bookInvoiceRequest).Execute()

Book an invoice



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
    invoiceId := int32(56) // int32 | ID of invoice to book
    bookInvoiceRequest := *openapiclient.NewBookInvoiceRequest(float32(123), int32(123), "Type_example", *openapiclient.NewBookVoucherRequestCheckAccount(int32(123), "CheckAccount")) // BookInvoiceRequest | Booking data (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoiceApi.BookInvoice(context.Background(), invoiceId).BookInvoiceRequest(bookInvoiceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApi.BookInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BookInvoice`: BookInvoice200Response
    fmt.Fprintf(os.Stdout, "Response from `InvoiceApi.BookInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **int32** | ID of invoice to book | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bookInvoiceRequest** | [**BookInvoiceRequest**](BookInvoiceRequest.md) | Booking data | 

### Return type

[**BookInvoice200Response**](BookInvoice200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelInvoice

> GetInvoiceById200Response CancelInvoice(ctx, invoiceId).Execute()

Cancel an invoice / Create cancellation invoice



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
    invoiceId := int32(56) // int32 | ID of invoice to be cancelled

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoiceApi.CancelInvoice(context.Background(), invoiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApi.CancelInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelInvoice`: GetInvoiceById200Response
    fmt.Fprintf(os.Stdout, "Response from `InvoiceApi.CancelInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **int32** | ID of invoice to be cancelled | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetInvoiceById200Response**](GetInvoiceById200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInvoiceByFactory

> SaveInvoiceResponse CreateInvoiceByFactory(ctx).SaveInvoice(saveInvoice).Execute()

Create a new invoice



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
    saveInvoice := *openapiclient.NewSaveInvoice(*openapiclient.NewModelInvoice(*openapiclient.NewModelInvoiceContact(int32(123), "Contact"), *openapiclient.NewModelInvoiceUpdateContactPerson(int32(123), "SevUser"), "01.01.2022", int32(0), *openapiclient.NewModelInvoiceAddressCountry(int32(1), "StaticCountry"), "100", float32(19), "Umsatzsteuer 19%", "default", "RE", "EUR", false)) // SaveInvoice | Creation data. Please be aware, that you need to provide at least all required parameter      of the invoice model! (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoiceApi.CreateInvoiceByFactory(context.Background()).SaveInvoice(saveInvoice).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApi.CreateInvoiceByFactory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInvoiceByFactory`: SaveInvoiceResponse
    fmt.Fprintf(os.Stdout, "Response from `InvoiceApi.CreateInvoiceByFactory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInvoiceByFactoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **saveInvoice** | [**SaveInvoice**](SaveInvoice.md) | Creation data. Please be aware, that you need to provide at least all required parameter      of the invoice model! | 

### Return type

[**SaveInvoiceResponse**](SaveInvoiceResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInvoiceFromOrder

> GetInvoiceById200Response CreateInvoiceFromOrder(ctx).InvoiceId(invoiceId).InvoiceObjectName(invoiceObjectName).ModelCreateInvoiceFromOrder(modelCreateInvoiceFromOrder).Execute()

Create invoice from order



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
    invoiceId := int32(56) // int32 | the id of the invoice
    invoiceObjectName := "Invoice" // string | Model name, which is 'Invoice'
    modelCreateInvoiceFromOrder := *openapiclient.NewModelCreateInvoiceFromOrder(*openapiclient.NewModelCreateInvoiceFromOrderOrder(int32(123), "Order")) // ModelCreateInvoiceFromOrder | Create invoice (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoiceApi.CreateInvoiceFromOrder(context.Background()).InvoiceId(invoiceId).InvoiceObjectName(invoiceObjectName).ModelCreateInvoiceFromOrder(modelCreateInvoiceFromOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApi.CreateInvoiceFromOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInvoiceFromOrder`: GetInvoiceById200Response
    fmt.Fprintf(os.Stdout, "Response from `InvoiceApi.CreateInvoiceFromOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInvoiceFromOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoiceId** | **int32** | the id of the invoice | 
 **invoiceObjectName** | **string** | Model name, which is &#39;Invoice&#39; | 
 **modelCreateInvoiceFromOrder** | [**ModelCreateInvoiceFromOrder**](ModelCreateInvoiceFromOrder.md) | Create invoice | 

### Return type

[**GetInvoiceById200Response**](GetInvoiceById200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInvoiceReminder

> GetInvoiceById200Response CreateInvoiceReminder(ctx).InvoiceId(invoiceId).InvoiceObjectName(invoiceObjectName).CreateInvoiceReminderRequest(createInvoiceReminderRequest).Execute()

Create invoice reminder



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
    invoiceId := int32(56) // int32 | the id of the invoice
    invoiceObjectName := "Invoice" // string | Model name, which is 'Invoice'
    createInvoiceReminderRequest := *openapiclient.NewCreateInvoiceReminderRequest(*openapiclient.NewCreateInvoiceReminderRequestInvoice(int32(123), "Invoice")) // CreateInvoiceReminderRequest | Create invoice (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoiceApi.CreateInvoiceReminder(context.Background()).InvoiceId(invoiceId).InvoiceObjectName(invoiceObjectName).CreateInvoiceReminderRequest(createInvoiceReminderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApi.CreateInvoiceReminder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInvoiceReminder`: GetInvoiceById200Response
    fmt.Fprintf(os.Stdout, "Response from `InvoiceApi.CreateInvoiceReminder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInvoiceReminderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoiceId** | **int32** | the id of the invoice | 
 **invoiceObjectName** | **string** | Model name, which is &#39;Invoice&#39; | 
 **createInvoiceReminderRequest** | [**CreateInvoiceReminderRequest**](CreateInvoiceReminderRequest.md) | Create invoice | 

### Return type

[**GetInvoiceById200Response**](GetInvoiceById200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoiceById

> GetInvoiceById200Response GetInvoiceById(ctx, invoiceId).Execute()

Find invoice by ID



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
    invoiceId := int32(56) // int32 | ID of invoice to return

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoiceApi.GetInvoiceById(context.Background(), invoiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApi.GetInvoiceById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoiceById`: GetInvoiceById200Response
    fmt.Fprintf(os.Stdout, "Response from `InvoiceApi.GetInvoiceById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **int32** | ID of invoice to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetInvoiceById200Response**](GetInvoiceById200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoicePositionsById

> GetInvoicePos200Response GetInvoicePositionsById(ctx, invoiceId).Limit(limit).Offset(offset).Embed(embed).Execute()

Find invoice positions



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
    invoiceId := int32(56) // int32 | ID of invoice to return the positions
    limit := int32(56) // int32 | limits the number of entries returned (optional)
    offset := int32(56) // int32 | set the index where the returned entries start (optional)
    embed := []string{"Inner_example"} // []string | Get some additional information. Embed can handle multiple values, they must be separated by comma. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoiceApi.GetInvoicePositionsById(context.Background(), invoiceId).Limit(limit).Offset(offset).Embed(embed).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApi.GetInvoicePositionsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoicePositionsById`: GetInvoicePos200Response
    fmt.Fprintf(os.Stdout, "Response from `InvoiceApi.GetInvoicePositionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **int32** | ID of invoice to return the positions | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoicePositionsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | limits the number of entries returned | 
 **offset** | **int32** | set the index where the returned entries start | 
 **embed** | **[]string** | Get some additional information. Embed can handle multiple values, they must be separated by comma. | 

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


## GetInvoices

> GetInvoiceById200Response GetInvoices(ctx).Status(status).InvoiceNumber(invoiceNumber).StartDate(startDate).EndDate(endDate).ContactId(contactId).ContactObjectName(contactObjectName).Execute()

Retrieve invoices



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
    status := float32(8.14) // float32 | Status of the invoices (optional)
    invoiceNumber := "invoiceNumber_example" // string | Retrieve all invoices with this invoice number (optional)
    startDate := int32(56) // int32 | Retrieve all invoices with a date equal or higher (optional)
    endDate := int32(56) // int32 | Retrieve all invoices with a date equal or lower (optional)
    contactId := int32(56) // int32 | Retrieve all invoices with this contact. Must be provided with contact[objectName] (optional)
    contactObjectName := "contactObjectName_example" // string | Only required if contact[id] was provided. 'Contact' should be used as value. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoiceApi.GetInvoices(context.Background()).Status(status).InvoiceNumber(invoiceNumber).StartDate(startDate).EndDate(endDate).ContactId(contactId).ContactObjectName(contactObjectName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApi.GetInvoices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoices`: GetInvoiceById200Response
    fmt.Fprintf(os.Stdout, "Response from `InvoiceApi.GetInvoices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **float32** | Status of the invoices | 
 **invoiceNumber** | **string** | Retrieve all invoices with this invoice number | 
 **startDate** | **int32** | Retrieve all invoices with a date equal or higher | 
 **endDate** | **int32** | Retrieve all invoices with a date equal or lower | 
 **contactId** | **int32** | Retrieve all invoices with this contact. Must be provided with contact[objectName] | 
 **contactObjectName** | **string** | Only required if contact[id] was provided. &#39;Contact&#39; should be used as value. | 

### Return type

[**GetInvoiceById200Response**](GetInvoiceById200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIsInvoicePartiallyPaid

> GetIsInvoicePartiallyPaid200Response GetIsInvoicePartiallyPaid(ctx, invoiceId).Execute()

Check if an invoice is already partially paid



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
    invoiceId := int32(56) // int32 | ID of invoice to return

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoiceApi.GetIsInvoicePartiallyPaid(context.Background(), invoiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApi.GetIsInvoicePartiallyPaid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIsInvoicePartiallyPaid`: GetIsInvoicePartiallyPaid200Response
    fmt.Fprintf(os.Stdout, "Response from `InvoiceApi.GetIsInvoicePartiallyPaid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **int32** | ID of invoice to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIsInvoicePartiallyPaidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetIsInvoicePartiallyPaid200Response**](GetIsInvoicePartiallyPaid200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoiceGetPdf

> InvoiceGetPdf200Response InvoiceGetPdf(ctx, invoiceId).Download(download).PreventSendBy(preventSendBy).Execute()

Retrieve pdf document of an invoice



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
    invoiceId := int32(56) // int32 | ID of invoice from which you want the pdf
    download := true // bool | If u want to download the pdf of the invoice. (optional)
    preventSendBy := true // bool | Defines if u want to send the invoice. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoiceApi.InvoiceGetPdf(context.Background(), invoiceId).Download(download).PreventSendBy(preventSendBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApi.InvoiceGetPdf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvoiceGetPdf`: InvoiceGetPdf200Response
    fmt.Fprintf(os.Stdout, "Response from `InvoiceApi.InvoiceGetPdf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **int32** | ID of invoice from which you want the pdf | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceGetPdfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **download** | **bool** | If u want to download the pdf of the invoice. | 
 **preventSendBy** | **bool** | Defines if u want to send the invoice. | 

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


## InvoiceRender

> InvoiceRender201Response InvoiceRender(ctx, invoiceId).InvoiceRenderRequest(invoiceRenderRequest).Execute()

Render the pdf document of an invoice



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
    invoiceId := int32(56) // int32 | ID of invoice to render
    invoiceRenderRequest := *openapiclient.NewInvoiceRenderRequest() // InvoiceRenderRequest | Define if the document should be forcefully re-rendered. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoiceApi.InvoiceRender(context.Background(), invoiceId).InvoiceRenderRequest(invoiceRenderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApi.InvoiceRender``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvoiceRender`: InvoiceRender201Response
    fmt.Fprintf(os.Stdout, "Response from `InvoiceApi.InvoiceRender`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **int32** | ID of invoice to render | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceRenderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invoiceRenderRequest** | [**InvoiceRenderRequest**](InvoiceRenderRequest.md) | Define if the document should be forcefully re-rendered. | 

### Return type

[**InvoiceRender201Response**](InvoiceRender201Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoiceSendBy

> InvoiceSendBy200Response InvoiceSendBy(ctx, invoiceId).InvoiceSendByRequest(invoiceSendByRequest).Execute()

Mark invoice as sent



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
    invoiceId := int32(56) // int32 | ID of invoice to mark as sent
    invoiceSendByRequest := *openapiclient.NewInvoiceSendByRequest("SendType_example", false) // InvoiceSendByRequest | Specify the send type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoiceApi.InvoiceSendBy(context.Background(), invoiceId).InvoiceSendByRequest(invoiceSendByRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApi.InvoiceSendBy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvoiceSendBy`: InvoiceSendBy200Response
    fmt.Fprintf(os.Stdout, "Response from `InvoiceApi.InvoiceSendBy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **int32** | ID of invoice to mark as sent | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceSendByRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invoiceSendByRequest** | [**InvoiceSendByRequest**](InvoiceSendByRequest.md) | Specify the send type | 

### Return type

[**InvoiceSendBy200Response**](InvoiceSendBy200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendInvoiceViaEMail

> SendInvoiceViaEMail201Response SendInvoiceViaEMail(ctx, invoiceId).SendInvoiceViaEMailRequest(sendInvoiceViaEMailRequest).Execute()

Send invoice via email



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
    invoiceId := int32(56) // int32 | ID of invoice to be sent via email
    sendInvoiceViaEMailRequest := *openapiclient.NewSendInvoiceViaEMailRequest("ToEmail_example", "Subject_example", "Text_example") // SendInvoiceViaEMailRequest | Mail data (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoiceApi.SendInvoiceViaEMail(context.Background(), invoiceId).SendInvoiceViaEMailRequest(sendInvoiceViaEMailRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApi.SendInvoiceViaEMail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendInvoiceViaEMail`: SendInvoiceViaEMail201Response
    fmt.Fprintf(os.Stdout, "Response from `InvoiceApi.SendInvoiceViaEMail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **int32** | ID of invoice to be sent via email | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendInvoiceViaEMailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendInvoiceViaEMailRequest** | [**SendInvoiceViaEMailRequest**](SendInvoiceViaEMailRequest.md) | Mail data | 

### Return type

[**SendInvoiceViaEMail201Response**](SendInvoiceViaEMail201Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

