# \CreditNoteApi

All URIs are relative to *https://my.sevdesk.de/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BookCreditNote**](CreditNoteApi.md#BookCreditNote) | **Put** /CreditNote/{creditNoteId}/bookAmount | Book a credit note
[**CreatecreditNote**](CreditNoteApi.md#CreatecreditNote) | **Post** /CreditNote/Factory/saveCreditNote | Create a new creditNote
[**CreditNoteGetPdf**](CreditNoteApi.md#CreditNoteGetPdf) | **Get** /CreditNote/{creditNoteId}/getPdf | Retrieve pdf document of a credit note
[**CreditNoteSendBy**](CreditNoteApi.md#CreditNoteSendBy) | **Put** /CreditNote/{creditNoteId}/sendBy | Mark credit note as sent
[**DeletecreditNote**](CreditNoteApi.md#DeletecreditNote) | **Delete** /creditNote/{creditNoteId} | Deletes an creditNote
[**GetCreditNotes**](CreditNoteApi.md#GetCreditNotes) | **Get** /CreditNote | Retrieve CreditNote
[**GetcreditNoteById**](CreditNoteApi.md#GetcreditNoteById) | **Get** /creditNote/{creditNoteId} | Find creditNote by ID
[**SendCreditNoteByPrinting**](CreditNoteApi.md#SendCreditNoteByPrinting) | **Get** /creditNote/{creditNoteId}/sendByWithRender | Send credit note by printing
[**SendCreditNoteViaEMail**](CreditNoteApi.md#SendCreditNoteViaEMail) | **Post** /CreditNote/{creditNoteId}/sendViaEmail | Send credit note via email
[**UpdatecreditNote**](CreditNoteApi.md#UpdatecreditNote) | **Put** /creditNote/{creditNoteId} | Update an existing creditNote



## BookCreditNote

> BookInvoice200Response BookCreditNote(ctx, creditNoteId).BookCreditNoteRequest(bookCreditNoteRequest).Execute()

Book a credit note



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
    creditNoteId := int32(56) // int32 | ID of credit note to book
    bookCreditNoteRequest := *openapiclient.NewBookCreditNoteRequest(float32(123), int32(123), "Type_example", *openapiclient.NewBookVoucherRequestCheckAccount(int32(123), "CheckAccount")) // BookCreditNoteRequest | Booking data (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreditNoteApi.BookCreditNote(context.Background(), creditNoteId).BookCreditNoteRequest(bookCreditNoteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreditNoteApi.BookCreditNote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BookCreditNote`: BookInvoice200Response
    fmt.Fprintf(os.Stdout, "Response from `CreditNoteApi.BookCreditNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**creditNoteId** | **int32** | ID of credit note to book | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookCreditNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bookCreditNoteRequest** | [**BookCreditNoteRequest**](BookCreditNoteRequest.md) | Booking data | 

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


## CreatecreditNote

> SaveCreditNoteResponse CreatecreditNote(ctx).SaveCreditNote(saveCreditNote).Execute()

Create a new creditNote



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
    saveCreditNote := *openapiclient.NewSaveCreditNote(*openapiclient.NewModelCreditNote("ObjectName_example", false, "Offer-1000", *openapiclient.NewModelCreditNoteContact(int32(123), "Contact"), time.Now(), "100", "My Offer-1000", "TODO", *openapiclient.NewModelCreditNoteContactPerson(int32(123), "SevUser"), float32(19), "Umsatzsteuer 19%", "default", "EUR")) // SaveCreditNote | Creation data. Please be aware, that you need to provide at least all required parameter      of the credit note model! (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreditNoteApi.CreatecreditNote(context.Background()).SaveCreditNote(saveCreditNote).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreditNoteApi.CreatecreditNote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatecreditNote`: SaveCreditNoteResponse
    fmt.Fprintf(os.Stdout, "Response from `CreditNoteApi.CreatecreditNote`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatecreditNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **saveCreditNote** | [**SaveCreditNote**](SaveCreditNote.md) | Creation data. Please be aware, that you need to provide at least all required parameter      of the credit note model! | 

### Return type

[**SaveCreditNoteResponse**](SaveCreditNoteResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditNoteGetPdf

> InvoiceGetPdf200Response CreditNoteGetPdf(ctx, creditNoteId).Download(download).PreventSendBy(preventSendBy).Execute()

Retrieve pdf document of a credit note



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
    creditNoteId := int32(56) // int32 | ID of credit note from which you want the pdf
    download := true // bool | If u want to download the pdf of the credit note. (optional)
    preventSendBy := true // bool | Defines if u want to send the credit note. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreditNoteApi.CreditNoteGetPdf(context.Background(), creditNoteId).Download(download).PreventSendBy(preventSendBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreditNoteApi.CreditNoteGetPdf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreditNoteGetPdf`: InvoiceGetPdf200Response
    fmt.Fprintf(os.Stdout, "Response from `CreditNoteApi.CreditNoteGetPdf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**creditNoteId** | **int32** | ID of credit note from which you want the pdf | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreditNoteGetPdfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **download** | **bool** | If u want to download the pdf of the credit note. | 
 **preventSendBy** | **bool** | Defines if u want to send the credit note. | 

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


## CreditNoteSendBy

> CreditNoteSendBy200Response CreditNoteSendBy(ctx, creditNoteId).CreditNoteSendByRequest(creditNoteSendByRequest).Execute()

Mark credit note as sent



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
    creditNoteId := int32(56) // int32 | ID of credit note to mark as sent
    creditNoteSendByRequest := *openapiclient.NewCreditNoteSendByRequest("VPDF", false) // CreditNoteSendByRequest | Specify the send type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreditNoteApi.CreditNoteSendBy(context.Background(), creditNoteId).CreditNoteSendByRequest(creditNoteSendByRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreditNoteApi.CreditNoteSendBy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreditNoteSendBy`: CreditNoteSendBy200Response
    fmt.Fprintf(os.Stdout, "Response from `CreditNoteApi.CreditNoteSendBy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**creditNoteId** | **int32** | ID of credit note to mark as sent | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreditNoteSendByRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **creditNoteSendByRequest** | [**CreditNoteSendByRequest**](CreditNoteSendByRequest.md) | Specify the send type | 

### Return type

[**CreditNoteSendBy200Response**](CreditNoteSendBy200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletecreditNote

> SendCreditNoteByPrinting200Response DeletecreditNote(ctx, creditNoteId).Execute()

Deletes an creditNote

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
    creditNoteId := int32(56) // int32 | Id of creditNote resource to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreditNoteApi.DeletecreditNote(context.Background(), creditNoteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreditNoteApi.DeletecreditNote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletecreditNote`: SendCreditNoteByPrinting200Response
    fmt.Fprintf(os.Stdout, "Response from `CreditNoteApi.DeletecreditNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**creditNoteId** | **int32** | Id of creditNote resource to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletecreditNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SendCreditNoteByPrinting200Response**](SendCreditNoteByPrinting200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCreditNotes

> SendCreditNoteByPrinting200Response GetCreditNotes(ctx).Status(status).CreditNoteNumber(creditNoteNumber).StartDate(startDate).EndDate(endDate).ContactId(contactId).ContactObjectName(contactObjectName).Execute()

Retrieve CreditNote



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
    status := "status_example" // string | Status of the CreditNote (optional)
    creditNoteNumber := "creditNoteNumber_example" // string | Retrieve all CreditNotes with this creditNote number (optional)
    startDate := int32(01.01.2020) // int32 | Retrieve all CreditNotes with a date equal or higher (optional)
    endDate := int32(01.01.2021) // int32 | Retrieve all CreditNotes with a date equal or lower (optional)
    contactId := int32(56) // int32 | Retrieve all CreditNotes with this contact. Must be provided with contact[objectName] (optional)
    contactObjectName := "contactObjectName_example" // string | Only required if contact[id] was provided. 'Contact' should be used as value. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreditNoteApi.GetCreditNotes(context.Background()).Status(status).CreditNoteNumber(creditNoteNumber).StartDate(startDate).EndDate(endDate).ContactId(contactId).ContactObjectName(contactObjectName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreditNoteApi.GetCreditNotes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCreditNotes`: SendCreditNoteByPrinting200Response
    fmt.Fprintf(os.Stdout, "Response from `CreditNoteApi.GetCreditNotes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCreditNotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Status of the CreditNote | 
 **creditNoteNumber** | **string** | Retrieve all CreditNotes with this creditNote number | 
 **startDate** | **int32** | Retrieve all CreditNotes with a date equal or higher | 
 **endDate** | **int32** | Retrieve all CreditNotes with a date equal or lower | 
 **contactId** | **int32** | Retrieve all CreditNotes with this contact. Must be provided with contact[objectName] | 
 **contactObjectName** | **string** | Only required if contact[id] was provided. &#39;Contact&#39; should be used as value. | 

### Return type

[**SendCreditNoteByPrinting200Response**](SendCreditNoteByPrinting200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetcreditNoteById

> SendCreditNoteByPrinting200Response GetcreditNoteById(ctx, creditNoteId).Execute()

Find creditNote by ID



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
    creditNoteId := int32(56) // int32 | ID of creditNote to return

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreditNoteApi.GetcreditNoteById(context.Background(), creditNoteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreditNoteApi.GetcreditNoteById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetcreditNoteById`: SendCreditNoteByPrinting200Response
    fmt.Fprintf(os.Stdout, "Response from `CreditNoteApi.GetcreditNoteById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**creditNoteId** | **int32** | ID of creditNote to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetcreditNoteByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SendCreditNoteByPrinting200Response**](SendCreditNoteByPrinting200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendCreditNoteByPrinting

> SendCreditNoteByPrinting200Response SendCreditNoteByPrinting(ctx, creditNoteId).SendType(sendType).Execute()

Send credit note by printing



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
    creditNoteId := int32(56) // int32 | ID of creditNote to return
    sendType := "PRN" // string | the type you want to print.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreditNoteApi.SendCreditNoteByPrinting(context.Background(), creditNoteId).SendType(sendType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreditNoteApi.SendCreditNoteByPrinting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendCreditNoteByPrinting`: SendCreditNoteByPrinting200Response
    fmt.Fprintf(os.Stdout, "Response from `CreditNoteApi.SendCreditNoteByPrinting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**creditNoteId** | **int32** | ID of creditNote to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendCreditNoteByPrintingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendType** | **string** | the type you want to print. | 

### Return type

[**SendCreditNoteByPrinting200Response**](SendCreditNoteByPrinting200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendCreditNoteViaEMail

> SendCreditNoteByPrinting200Response SendCreditNoteViaEMail(ctx, creditNoteId).SendInvoiceViaEMailRequest(sendInvoiceViaEMailRequest).Execute()

Send credit note via email



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
    creditNoteId := int32(56) // int32 | ID of credit note to be sent via email
    sendInvoiceViaEMailRequest := *openapiclient.NewSendInvoiceViaEMailRequest("ToEmail_example", "Subject_example", "Text_example") // SendInvoiceViaEMailRequest | Mail data (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreditNoteApi.SendCreditNoteViaEMail(context.Background(), creditNoteId).SendInvoiceViaEMailRequest(sendInvoiceViaEMailRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreditNoteApi.SendCreditNoteViaEMail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendCreditNoteViaEMail`: SendCreditNoteByPrinting200Response
    fmt.Fprintf(os.Stdout, "Response from `CreditNoteApi.SendCreditNoteViaEMail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**creditNoteId** | **int32** | ID of credit note to be sent via email | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendCreditNoteViaEMailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendInvoiceViaEMailRequest** | [**SendInvoiceViaEMailRequest**](SendInvoiceViaEMailRequest.md) | Mail data | 

### Return type

[**SendCreditNoteByPrinting200Response**](SendCreditNoteByPrinting200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatecreditNote

> SendCreditNoteByPrinting200Response UpdatecreditNote(ctx, creditNoteId).ModelCreditNoteUpdate(modelCreditNoteUpdate).Execute()

Update an existing creditNote



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
    creditNoteId := int32(56) // int32 | ID of creditNote to update
    modelCreditNoteUpdate := *openapiclient.NewModelCreditNoteUpdate() // ModelCreditNoteUpdate | Update data (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreditNoteApi.UpdatecreditNote(context.Background(), creditNoteId).ModelCreditNoteUpdate(modelCreditNoteUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreditNoteApi.UpdatecreditNote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatecreditNote`: SendCreditNoteByPrinting200Response
    fmt.Fprintf(os.Stdout, "Response from `CreditNoteApi.UpdatecreditNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**creditNoteId** | **int32** | ID of creditNote to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatecreditNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelCreditNoteUpdate** | [**ModelCreditNoteUpdate**](ModelCreditNoteUpdate.md) | Update data | 

### Return type

[**SendCreditNoteByPrinting200Response**](SendCreditNoteByPrinting200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

