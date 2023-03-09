# \VoucherApi

All URIs are relative to *https://my.sevdesk.de/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BookVoucher**](VoucherApi.md#BookVoucher) | **Put** /Voucher/{voucherId}/bookAmount | Book a voucher
[**CreateVoucherByFactory**](VoucherApi.md#CreateVoucherByFactory) | **Post** /Voucher/Factory/saveVoucher | Create a new voucher
[**GetVoucherById**](VoucherApi.md#GetVoucherById) | **Get** /Voucher/{voucherId} | Find voucher by ID
[**GetVouchers**](VoucherApi.md#GetVouchers) | **Get** /Voucher | Retrieve vouchers
[**UpdateVoucher**](VoucherApi.md#UpdateVoucher) | **Put** /Voucher/{voucherId} | Update an existing voucher
[**VoucherUploadFile**](VoucherApi.md#VoucherUploadFile) | **Post** /Voucher/Factory/uploadTempFile | Upload voucher file



## BookVoucher

> BookVoucher200Response BookVoucher(ctx, voucherId).BookVoucherRequest(bookVoucherRequest).Execute()

Book a voucher



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
    voucherId := int32(56) // int32 | ID of voucher to book
    bookVoucherRequest := *openapiclient.NewBookVoucherRequest(float32(123), time.Now(), "Type_example", *openapiclient.NewBookVoucherRequestCheckAccount(int32(123), "CheckAccount")) // BookVoucherRequest | Booking data (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VoucherApi.BookVoucher(context.Background(), voucherId).BookVoucherRequest(bookVoucherRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VoucherApi.BookVoucher``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BookVoucher`: BookVoucher200Response
    fmt.Fprintf(os.Stdout, "Response from `VoucherApi.BookVoucher`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**voucherId** | **int32** | ID of voucher to book | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookVoucherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bookVoucherRequest** | [**BookVoucherRequest**](BookVoucherRequest.md) | Booking data | 

### Return type

[**BookVoucher200Response**](BookVoucher200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVoucherByFactory

> SaveVoucherResponse CreateVoucherByFactory(ctx).SaveVoucher(saveVoucher).Execute()

Create a new voucher



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
    saveVoucher := *openapiclient.NewSaveVoucher(*openapiclient.NewModelVoucher("ObjectName_example", true, float32(50), "default", "C", "VOU")) // SaveVoucher | Creation data. Please be aware, that you need to provide at least all required parameter      of the voucher and voucher position model! (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VoucherApi.CreateVoucherByFactory(context.Background()).SaveVoucher(saveVoucher).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VoucherApi.CreateVoucherByFactory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVoucherByFactory`: SaveVoucherResponse
    fmt.Fprintf(os.Stdout, "Response from `VoucherApi.CreateVoucherByFactory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVoucherByFactoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **saveVoucher** | [**SaveVoucher**](SaveVoucher.md) | Creation data. Please be aware, that you need to provide at least all required parameter      of the voucher and voucher position model! | 

### Return type

[**SaveVoucherResponse**](SaveVoucherResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVoucherById

> GetVoucherById200Response GetVoucherById(ctx, voucherId).Execute()

Find voucher by ID



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
    voucherId := int32(56) // int32 | ID of voucher to return

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VoucherApi.GetVoucherById(context.Background(), voucherId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VoucherApi.GetVoucherById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVoucherById`: GetVoucherById200Response
    fmt.Fprintf(os.Stdout, "Response from `VoucherApi.GetVoucherById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**voucherId** | **int32** | ID of voucher to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVoucherByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetVoucherById200Response**](GetVoucherById200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVouchers

> GetVoucherById200Response GetVouchers(ctx).Status(status).CreditDebit(creditDebit).DescriptionLike(descriptionLike).StartDate(startDate).EndDate(endDate).ContactId(contactId).ContactObjectName(contactObjectName).Execute()

Retrieve vouchers



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
    status := float32(8.14) // float32 | Status of the vouchers to retrieve. (optional)
    creditDebit := "creditDebit_example" // string | Define if you only want credit or debit vouchers. (optional)
    descriptionLike := "descriptionLike_example" // string | Retrieve all vouchers with a description like this. (optional)
    startDate := int32(01.01.2020) // int32 | Retrieve all vouchers with a date equal or higher (optional)
    endDate := int32(01.01.2020) // int32 | Retrieve all vouchers with a date equal or lower (optional)
    contactId := int32(56) // int32 | Retrieve all vouchers with this contact. Must be provided with contact[objectName] (optional)
    contactObjectName := "contactObjectName_example" // string | Only required if contact[id] was provided. 'Contact' should be used as value. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VoucherApi.GetVouchers(context.Background()).Status(status).CreditDebit(creditDebit).DescriptionLike(descriptionLike).StartDate(startDate).EndDate(endDate).ContactId(contactId).ContactObjectName(contactObjectName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VoucherApi.GetVouchers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVouchers`: GetVoucherById200Response
    fmt.Fprintf(os.Stdout, "Response from `VoucherApi.GetVouchers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVouchersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **float32** | Status of the vouchers to retrieve. | 
 **creditDebit** | **string** | Define if you only want credit or debit vouchers. | 
 **descriptionLike** | **string** | Retrieve all vouchers with a description like this. | 
 **startDate** | **int32** | Retrieve all vouchers with a date equal or higher | 
 **endDate** | **int32** | Retrieve all vouchers with a date equal or lower | 
 **contactId** | **int32** | Retrieve all vouchers with this contact. Must be provided with contact[objectName] | 
 **contactObjectName** | **string** | Only required if contact[id] was provided. &#39;Contact&#39; should be used as value. | 

### Return type

[**GetVoucherById200Response**](GetVoucherById200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVoucher

> GetVoucherById200Response UpdateVoucher(ctx, voucherId).ModelVoucherUpdate(modelVoucherUpdate).Execute()

Update an existing voucher



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
    voucherId := int32(56) // int32 | ID of voucher to update
    modelVoucherUpdate := *openapiclient.NewModelVoucherUpdate() // ModelVoucherUpdate | Update data (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VoucherApi.UpdateVoucher(context.Background(), voucherId).ModelVoucherUpdate(modelVoucherUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VoucherApi.UpdateVoucher``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVoucher`: GetVoucherById200Response
    fmt.Fprintf(os.Stdout, "Response from `VoucherApi.UpdateVoucher`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**voucherId** | **int32** | ID of voucher to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVoucherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelVoucherUpdate** | [**ModelVoucherUpdate**](ModelVoucherUpdate.md) | Update data | 

### Return type

[**GetVoucherById200Response**](GetVoucherById200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VoucherUploadFile

> VoucherUploadFile201Response VoucherUploadFile(ctx).VoucherUploadFileRequest(voucherUploadFileRequest).Execute()

Upload voucher file



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
    voucherUploadFileRequest := *openapiclient.NewVoucherUploadFileRequest() // VoucherUploadFileRequest | File to upload (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VoucherApi.VoucherUploadFile(context.Background()).VoucherUploadFileRequest(voucherUploadFileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VoucherApi.VoucherUploadFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VoucherUploadFile`: VoucherUploadFile201Response
    fmt.Fprintf(os.Stdout, "Response from `VoucherApi.VoucherUploadFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVoucherUploadFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **voucherUploadFileRequest** | [**VoucherUploadFileRequest**](VoucherUploadFileRequest.md) | File to upload | 

### Return type

[**VoucherUploadFile201Response**](VoucherUploadFile201Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

