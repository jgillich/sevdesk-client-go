# \VoucherPosApi

All URIs are relative to *https://my.sevdesk.de/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVoucherPositions**](VoucherPosApi.md#GetVoucherPositions) | **Get** /VoucherPos | Retrieve voucher positions



## GetVoucherPositions

> GetVoucherPositions200Response GetVoucherPositions(ctx).VoucherId(voucherId).VoucherObjectName(voucherObjectName).Execute()

Retrieve voucher positions



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
    voucherId := int32(56) // int32 | Retrieve all vouchers positions belonging to this voucher. Must be provided with voucher[objectName] (optional)
    voucherObjectName := "voucherObjectName_example" // string | Only required if voucher[id] was provided. 'Voucher' should be used as value. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VoucherPosApi.GetVoucherPositions(context.Background()).VoucherId(voucherId).VoucherObjectName(voucherObjectName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VoucherPosApi.GetVoucherPositions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVoucherPositions`: GetVoucherPositions200Response
    fmt.Fprintf(os.Stdout, "Response from `VoucherPosApi.GetVoucherPositions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVoucherPositionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **voucherId** | **int32** | Retrieve all vouchers positions belonging to this voucher. Must be provided with voucher[objectName] | 
 **voucherObjectName** | **string** | Only required if voucher[id] was provided. &#39;Voucher&#39; should be used as value. | 

### Return type

[**GetVoucherPositions200Response**](GetVoucherPositions200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

