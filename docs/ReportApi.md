# \ReportApi

All URIs are relative to *https://my.sevdesk.de/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReportContact**](ReportApi.md#ReportContact) | **Get** /Report/contactlist | Export contact list
[**ReportInvoice**](ReportApi.md#ReportInvoice) | **Get** /Report/invoicelist | Export invoice list
[**ReportOrder**](ReportApi.md#ReportOrder) | **Get** /Report/orderlist | Export order list
[**ReportVoucher**](ReportApi.md#ReportVoucher) | **Get** /Report/voucherlist | Export voucher list



## ReportContact

> map[string]interface{} ReportContact(ctx).SevQuery(sevQuery).Download(download).Execute()

Export contact list



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
    sevQuery := map[string][]openapiclient.ReportContactSevQueryParameter{ ... } // ReportContactSevQueryParameter | 
    download := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportApi.ReportContact(context.Background()).SevQuery(sevQuery).Download(download).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportApi.ReportContact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportContact`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ReportApi.ReportContact`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sevQuery** | [**ReportContactSevQueryParameter**](ReportContactSevQueryParameter.md) |  | 
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


## ReportInvoice

> map[string]interface{} ReportInvoice(ctx).SevQuery(sevQuery).Download(download).Execute()

Export invoice list



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
    sevQuery := map[string][]openapiclient.ReportInvoiceSevQueryParameter{ ... } // ReportInvoiceSevQueryParameter | 
    download := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportApi.ReportInvoice(context.Background()).SevQuery(sevQuery).Download(download).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportApi.ReportInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportInvoice`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ReportApi.ReportInvoice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sevQuery** | [**ReportInvoiceSevQueryParameter**](ReportInvoiceSevQueryParameter.md) |  | 
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


## ReportOrder

> map[string]interface{} ReportOrder(ctx).SevQuery(sevQuery).Download(download).Execute()

Export order list



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
    sevQuery := map[string][]openapiclient.ReportOrderSevQueryParameter{ ... } // ReportOrderSevQueryParameter | 
    download := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportApi.ReportOrder(context.Background()).SevQuery(sevQuery).Download(download).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportApi.ReportOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportOrder`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ReportApi.ReportOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sevQuery** | [**ReportOrderSevQueryParameter**](ReportOrderSevQueryParameter.md) |  | 
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


## ReportVoucher

> map[string]interface{} ReportVoucher(ctx).SevQuery(sevQuery).Download(download).Execute()

Export voucher list



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
    sevQuery := map[string][]openapiclient.ReportVoucherSevQueryParameter{ ... } // ReportVoucherSevQueryParameter | 
    download := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportApi.ReportVoucher(context.Background()).SevQuery(sevQuery).Download(download).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportApi.ReportVoucher``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportVoucher`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ReportApi.ReportVoucher`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportVoucherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sevQuery** | [**ReportVoucherSevQueryParameter**](ReportVoucherSevQueryParameter.md) |  | 
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

