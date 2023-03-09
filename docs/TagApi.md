# \TagApi

All URIs are relative to *https://my.sevdesk.de/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTag**](TagApi.md#CreateTag) | **Post** /Tag/Factory/create | Create a new tag
[**DeleteTag**](TagApi.md#DeleteTag) | **Delete** /Tag/{tagId} | Deletes a tag
[**GetTagById**](TagApi.md#GetTagById) | **Get** /Tag/{tagId} | Find tag by ID
[**GetTagRelations**](TagApi.md#GetTagRelations) | **Get** /TagRelation | Retrieve tag relations
[**GetTags**](TagApi.md#GetTags) | **Get** /Tag | Retrieve tags
[**UpdateTag**](TagApi.md#UpdateTag) | **Put** /Tag/{tagId} | Update tag



## CreateTag

> GetTagRelations200Response CreateTag(ctx).CreateTagRequest(createTagRequest).Execute()

Create a new tag



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
    createTagRequest := *openapiclient.NewCreateTagRequest() // CreateTagRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagApi.CreateTag(context.Background()).CreateTagRequest(createTagRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagApi.CreateTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTag`: GetTagRelations200Response
    fmt.Fprintf(os.Stdout, "Response from `TagApi.CreateTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTagRequest** | [**CreateTagRequest**](CreateTagRequest.md) |  | 

### Return type

[**GetTagRelations200Response**](GetTagRelations200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTag

> DeleteAccountingContact200Response DeleteTag(ctx, tagId).Execute()

Deletes a tag

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
    tagId := int32(56) // int32 | Id of tag to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagApi.DeleteTag(context.Background(), tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagApi.DeleteTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTag`: DeleteAccountingContact200Response
    fmt.Fprintf(os.Stdout, "Response from `TagApi.DeleteTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **int32** | Id of tag to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagRequest struct via the builder pattern


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


## GetTagById

> GetTagById200Response GetTagById(ctx, tagId).Execute()

Find tag by ID



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
    tagId := int32(56) // int32 | ID of tag to return

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagApi.GetTagById(context.Background(), tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagApi.GetTagById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTagById`: GetTagById200Response
    fmt.Fprintf(os.Stdout, "Response from `TagApi.GetTagById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **int32** | ID of tag to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetTagById200Response**](GetTagById200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTagRelations

> GetTagRelations200Response GetTagRelations(ctx).Execute()

Retrieve tag relations



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
    resp, r, err := apiClient.TagApi.GetTagRelations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagApi.GetTagRelations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTagRelations`: GetTagRelations200Response
    fmt.Fprintf(os.Stdout, "Response from `TagApi.GetTagRelations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagRelationsRequest struct via the builder pattern


### Return type

[**GetTagRelations200Response**](GetTagRelations200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTags

> GetTagById200Response GetTags(ctx).Id(id).Name(name).Execute()

Retrieve tags



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
    id := float32(8.14) // float32 | ID of the Tag (optional)
    name := "name_example" // string | Name of the Tag (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagApi.GetTags(context.Background()).Id(id).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagApi.GetTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTags`: GetTagById200Response
    fmt.Fprintf(os.Stdout, "Response from `TagApi.GetTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **float32** | ID of the Tag | 
 **name** | **string** | Name of the Tag | 

### Return type

[**GetTagById200Response**](GetTagById200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTag

> GetTagById200Response UpdateTag(ctx, tagId).UpdateTagRequest(updateTagRequest).Execute()

Update tag



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
    tagId := int32(56) // int32 | ID of tag you want to update
    updateTagRequest := *openapiclient.NewUpdateTagRequest() // UpdateTagRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagApi.UpdateTag(context.Background(), tagId).UpdateTagRequest(updateTagRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagApi.UpdateTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTag`: GetTagById200Response
    fmt.Fprintf(os.Stdout, "Response from `TagApi.UpdateTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **int32** | ID of tag you want to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTagRequest** | [**UpdateTagRequest**](UpdateTagRequest.md) |  | 

### Return type

[**GetTagById200Response**](GetTagById200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

