# \ExportApi

All URIs are relative to *https://my.sevdesk.de/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportContact**](ExportApi.md#ExportContact) | **Get** /Export/contactListCsv | Export contact
[**ExportCreditNote**](ExportApi.md#ExportCreditNote) | **Get** /Export/creditNoteCsv | Export creditNote
[**ExportDatev**](ExportApi.md#ExportDatev) | **Get** /Export/datevCSV | Export datev
[**ExportInvoice**](ExportApi.md#ExportInvoice) | **Get** /Export/invoiceCsv | Export invoice
[**ExportInvoiceZip**](ExportApi.md#ExportInvoiceZip) | **Get** /Export/invoiceZip | Export Invoice as zip
[**ExportTransactions**](ExportApi.md#ExportTransactions) | **Get** /Export/transactionsCsv | Export transaction
[**ExportVoucher**](ExportApi.md#ExportVoucher) | **Get** /Export/voucherListCsv | Export voucher as zip
[**ExportVoucherZip**](ExportApi.md#ExportVoucherZip) | **Get** /Export/voucherZip | Export voucher zip



## ExportContact

> map[string]interface{} ExportContact(ctx).SevQuery(sevQuery).Download(download).Execute()

Export contact



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
    sevQuery := map[string][]openapiclient.ExportContactSevQueryParameter{ ... } // ExportContactSevQueryParameter | 
    download := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportApi.ExportContact(context.Background()).SevQuery(sevQuery).Download(download).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportApi.ExportContact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportContact`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ExportApi.ExportContact`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sevQuery** | [**ExportContactSevQueryParameter**](ExportContactSevQueryParameter.md) |  | 
 **download** | **bool** |  | 

### Return type

**map[string]interface{}**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportCreditNote

> map[string]interface{} ExportCreditNote(ctx).SevQuery(sevQuery).Download(download).Execute()

Export creditNote



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
    sevQuery := map[string][]openapiclient.ExportCreditNoteSevQueryParameter{ ... } // ExportCreditNoteSevQueryParameter | 
    download := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportApi.ExportCreditNote(context.Background()).SevQuery(sevQuery).Download(download).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportApi.ExportCreditNote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportCreditNote`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ExportApi.ExportCreditNote`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportCreditNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sevQuery** | [**ExportCreditNoteSevQueryParameter**](ExportCreditNoteSevQueryParameter.md) |  | 
 **download** | **bool** |  | 

### Return type

**map[string]interface{}**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportDatev

> map[string]interface{} ExportDatev(ctx).StartDate(startDate).EndDate(endDate).Scope(scope).Download(download).WithUnpaidDocuments(withUnpaidDocuments).WithEnshrinedDocuments(withEnshrinedDocuments).Enshrine(enshrine).Execute()

Export datev



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
    startDate := int32(1641032867) // int32 | the start date of the export as timestamp
    endDate := int32(1648805267) // int32 | the end date of the export as timestamp
    scope := "EXTCD" // string | Define what you want to include in the datev export. This parameter takes a string of 5 letters. Each stands for a model that should be included. Possible letters are: ‘E’ (Earnings), ‘X’ (Expenditure), ‘T’ (Transactions), ‘C’ (Cashregister), ‘D’ (Assets). By providing one of those letter you specify that it should be included in the datev export. Some combinations are: ‘EXTCD’, ‘EXTD’ …
    download := true // bool | Specifies if the document is downloaded (optional)
    withUnpaidDocuments := true // bool | include unpaid documents (optional)
    withEnshrinedDocuments := true // bool | include enshrined documents (optional)
    enshrine := false // bool | Specify if you want to enshrine all models which were included in the export (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportApi.ExportDatev(context.Background()).StartDate(startDate).EndDate(endDate).Scope(scope).Download(download).WithUnpaidDocuments(withUnpaidDocuments).WithEnshrinedDocuments(withEnshrinedDocuments).Enshrine(enshrine).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportApi.ExportDatev``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportDatev`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ExportApi.ExportDatev`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportDatevRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **int32** | the start date of the export as timestamp | 
 **endDate** | **int32** | the end date of the export as timestamp | 
 **scope** | **string** | Define what you want to include in the datev export. This parameter takes a string of 5 letters. Each stands for a model that should be included. Possible letters are: ‘E’ (Earnings), ‘X’ (Expenditure), ‘T’ (Transactions), ‘C’ (Cashregister), ‘D’ (Assets). By providing one of those letter you specify that it should be included in the datev export. Some combinations are: ‘EXTCD’, ‘EXTD’ … | 
 **download** | **bool** | Specifies if the document is downloaded | 
 **withUnpaidDocuments** | **bool** | include unpaid documents | 
 **withEnshrinedDocuments** | **bool** | include enshrined documents | 
 **enshrine** | **bool** | Specify if you want to enshrine all models which were included in the export | 

### Return type

**map[string]interface{}**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportInvoice

> map[string]interface{} ExportInvoice(ctx).SevQuery(sevQuery).Download(download).Execute()

Export invoice



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
    sevQuery := map[string][]openapiclient.ExportInvoiceSevQueryParameter{ ... } // ExportInvoiceSevQueryParameter | 
    download := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportApi.ExportInvoice(context.Background()).SevQuery(sevQuery).Download(download).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportApi.ExportInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportInvoice`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ExportApi.ExportInvoice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sevQuery** | [**ExportInvoiceSevQueryParameter**](ExportInvoiceSevQueryParameter.md) |  | 
 **download** | **bool** |  | 

### Return type

**map[string]interface{}**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportInvoiceZip

> map[string]interface{} ExportInvoiceZip(ctx).SevQuery(sevQuery).Download(download).Execute()

Export Invoice as zip



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
    sevQuery := map[string][]openapiclient.ExportInvoiceSevQueryParameter{ ... } // ExportInvoiceSevQueryParameter | 
    download := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportApi.ExportInvoiceZip(context.Background()).SevQuery(sevQuery).Download(download).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportApi.ExportInvoiceZip``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportInvoiceZip`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ExportApi.ExportInvoiceZip`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportInvoiceZipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sevQuery** | [**ExportInvoiceSevQueryParameter**](ExportInvoiceSevQueryParameter.md) |  | 
 **download** | **bool** |  | 

### Return type

**map[string]interface{}**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportTransactions

> map[string]interface{} ExportTransactions(ctx).SevQuery(sevQuery).Download(download).Execute()

Export transaction



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
    sevQuery := map[string][]openapiclient.ExportTransactionsSevQueryParameter{ ... } // ExportTransactionsSevQueryParameter | 
    download := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportApi.ExportTransactions(context.Background()).SevQuery(sevQuery).Download(download).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportApi.ExportTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportTransactions`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ExportApi.ExportTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sevQuery** | [**ExportTransactionsSevQueryParameter**](ExportTransactionsSevQueryParameter.md) |  | 
 **download** | **bool** |  | 

### Return type

**map[string]interface{}**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportVoucher

> map[string]interface{} ExportVoucher(ctx).SevQuery(sevQuery).Download(download).Execute()

Export voucher as zip



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
    sevQuery := map[string][]openapiclient.ExportVoucherSevQueryParameter{ ... } // ExportVoucherSevQueryParameter | 
    download := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportApi.ExportVoucher(context.Background()).SevQuery(sevQuery).Download(download).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportApi.ExportVoucher``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportVoucher`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ExportApi.ExportVoucher`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportVoucherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sevQuery** | [**ExportVoucherSevQueryParameter**](ExportVoucherSevQueryParameter.md) |  | 
 **download** | **bool** |  | 

### Return type

**map[string]interface{}**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportVoucherZip

> map[string]interface{} ExportVoucherZip(ctx).SevQuery(sevQuery).Download(download).Execute()

Export voucher zip



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
    sevQuery := map[string][]openapiclient.ExportVoucherSevQueryParameter{ ... } // ExportVoucherSevQueryParameter | 
    download := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportApi.ExportVoucherZip(context.Background()).SevQuery(sevQuery).Download(download).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportApi.ExportVoucherZip``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportVoucherZip`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ExportApi.ExportVoucherZip`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportVoucherZipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sevQuery** | [**ExportVoucherSevQueryParameter**](ExportVoucherSevQueryParameter.md) |  | 
 **download** | **bool** |  | 

### Return type

**map[string]interface{}**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

