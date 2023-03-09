# \CreditNotePosApi

All URIs are relative to *https://my.sevdesk.de/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetcreditNotePositions**](CreditNotePosApi.md#GetcreditNotePositions) | **Get** /creditNotePos | Retrieve creditNote positions



## GetcreditNotePositions

> GetcreditNotePositions200Response GetcreditNotePositions(ctx).CreditNoteId(creditNoteId).CreditNoteObjectName(creditNoteObjectName).Execute()

Retrieve creditNote positions



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
    creditNoteId := int32(56) // int32 | Retrieve all creditNote positions belonging to this creditNote. Must be provided with creditNote[objectName] (optional)
    creditNoteObjectName := "creditNoteObjectName_example" // string | Only required if creditNote[id] was provided. 'creditNote' should be used as value. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreditNotePosApi.GetcreditNotePositions(context.Background()).CreditNoteId(creditNoteId).CreditNoteObjectName(creditNoteObjectName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreditNotePosApi.GetcreditNotePositions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetcreditNotePositions`: GetcreditNotePositions200Response
    fmt.Fprintf(os.Stdout, "Response from `CreditNotePosApi.GetcreditNotePositions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetcreditNotePositionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **creditNoteId** | **int32** | Retrieve all creditNote positions belonging to this creditNote. Must be provided with creditNote[objectName] | 
 **creditNoteObjectName** | **string** | Only required if creditNote[id] was provided. &#39;creditNote&#39; should be used as value. | 

### Return type

[**GetcreditNotePositions200Response**](GetcreditNotePositions200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

