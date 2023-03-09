# \ContactFieldApi

All URIs are relative to *https://my.sevdesk.de/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContactField**](ContactFieldApi.md#CreateContactField) | **Post** /ContactCustomField | Create contact field
[**CreateContactFieldSetting**](ContactFieldApi.md#CreateContactFieldSetting) | **Post** /ContactCustomFieldSetting | Create contact field setting
[**DeleteContactCustomFieldId**](ContactFieldApi.md#DeleteContactCustomFieldId) | **Delete** /ContactCustomField/{contactCustomFieldId} | delete a contact field
[**DeleteContactFieldSetting**](ContactFieldApi.md#DeleteContactFieldSetting) | **Delete** /ContactCustomFieldSetting/{contactCustomFieldSettingId} | Deletes a contact field setting
[**GetContactFieldSettingById**](ContactFieldApi.md#GetContactFieldSettingById) | **Get** /ContactCustomFieldSetting/{contactCustomFieldSettingId} | Find contact field setting by ID
[**GetContactFieldSettings**](ContactFieldApi.md#GetContactFieldSettings) | **Get** /ContactCustomFieldSetting | Retrieve contact field settings
[**GetContactFields**](ContactFieldApi.md#GetContactFields) | **Get** /ContactCustomField | Retrieve contact fields
[**GetContactFieldsById**](ContactFieldApi.md#GetContactFieldsById) | **Get** /ContactCustomField/{contactCustomFieldId} | Retrieve contact fields
[**GetPlaceholder**](ContactFieldApi.md#GetPlaceholder) | **Get** /Textparser/fetchDictionaryEntriesByType | Retrieve Placeholders
[**GetReferenceCount**](ContactFieldApi.md#GetReferenceCount) | **Get** /ContactCustomFieldSetting/{contactCustomFieldSettingId}/getReferenceCount | Receive count reference
[**UpdateContactFieldSetting**](ContactFieldApi.md#UpdateContactFieldSetting) | **Put** /ContactCustomFieldSetting/{contactCustomFieldSettingId} | Update contact field setting
[**UpdateContactfield**](ContactFieldApi.md#UpdateContactfield) | **Put** /ContactCustomField/{contactCustomFieldId} | Update a contact field



## CreateContactField

> GetContactFieldsById200Response CreateContactField(ctx).ModelContactCustomField(modelContactCustomField).Execute()

Create contact field



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
    modelContactCustomField := *openapiclient.NewModelContactCustomField(*openapiclient.NewModelContactCustomFieldContact(int32(123), "Contact"), *openapiclient.NewModelContactCustomFieldContactCustomFieldSetting(int32(123), "contactCustomFieldSetting"), "Value_example", "ContactCustomField") // ModelContactCustomField |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactFieldApi.CreateContactField(context.Background()).ModelContactCustomField(modelContactCustomField).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactFieldApi.CreateContactField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateContactField`: GetContactFieldsById200Response
    fmt.Fprintf(os.Stdout, "Response from `ContactFieldApi.CreateContactField`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateContactFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelContactCustomField** | [**ModelContactCustomField**](ModelContactCustomField.md) |  | 

### Return type

[**GetContactFieldsById200Response**](GetContactFieldsById200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContactFieldSetting

> GetContactFieldSettings200Response CreateContactFieldSetting(ctx).ModelContactCustomFieldSetting(modelContactCustomFieldSetting).Execute()

Create contact field setting



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
    modelContactCustomFieldSetting := *openapiclient.NewModelContactCustomFieldSetting("Name_example") // ModelContactCustomFieldSetting |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactFieldApi.CreateContactFieldSetting(context.Background()).ModelContactCustomFieldSetting(modelContactCustomFieldSetting).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactFieldApi.CreateContactFieldSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateContactFieldSetting`: GetContactFieldSettings200Response
    fmt.Fprintf(os.Stdout, "Response from `ContactFieldApi.CreateContactFieldSetting`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateContactFieldSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelContactCustomFieldSetting** | [**ModelContactCustomFieldSetting**](ModelContactCustomFieldSetting.md) |  | 

### Return type

[**GetContactFieldSettings200Response**](GetContactFieldSettings200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContactCustomFieldId

> DeleteAccountingContact200Response DeleteContactCustomFieldId(ctx, contactCustomFieldId).Execute()

delete a contact field

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
    contactCustomFieldId := int32(56) // int32 | Id of contact field

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactFieldApi.DeleteContactCustomFieldId(context.Background(), contactCustomFieldId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactFieldApi.DeleteContactCustomFieldId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteContactCustomFieldId`: DeleteAccountingContact200Response
    fmt.Fprintf(os.Stdout, "Response from `ContactFieldApi.DeleteContactCustomFieldId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactCustomFieldId** | **int32** | Id of contact field | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContactCustomFieldIdRequest struct via the builder pattern


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


## DeleteContactFieldSetting

> DeleteAccountingContact200Response DeleteContactFieldSetting(ctx, contactCustomFieldSettingId).Execute()

Deletes a contact field setting

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
    contactCustomFieldSettingId := int32(56) // int32 | Id of contact field to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactFieldApi.DeleteContactFieldSetting(context.Background(), contactCustomFieldSettingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactFieldApi.DeleteContactFieldSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteContactFieldSetting`: DeleteAccountingContact200Response
    fmt.Fprintf(os.Stdout, "Response from `ContactFieldApi.DeleteContactFieldSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactCustomFieldSettingId** | **int32** | Id of contact field to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContactFieldSettingRequest struct via the builder pattern


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


## GetContactFieldSettingById

> GetContactFieldSettings200Response GetContactFieldSettingById(ctx, contactCustomFieldSettingId).Execute()

Find contact field setting by ID



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
    contactCustomFieldSettingId := int32(56) // int32 | ID of contact field to return

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactFieldApi.GetContactFieldSettingById(context.Background(), contactCustomFieldSettingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactFieldApi.GetContactFieldSettingById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContactFieldSettingById`: GetContactFieldSettings200Response
    fmt.Fprintf(os.Stdout, "Response from `ContactFieldApi.GetContactFieldSettingById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactCustomFieldSettingId** | **int32** | ID of contact field to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContactFieldSettingByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetContactFieldSettings200Response**](GetContactFieldSettings200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContactFieldSettings

> GetContactFieldSettings200Response GetContactFieldSettings(ctx).Execute()

Retrieve contact field settings



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
    resp, r, err := apiClient.ContactFieldApi.GetContactFieldSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactFieldApi.GetContactFieldSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContactFieldSettings`: GetContactFieldSettings200Response
    fmt.Fprintf(os.Stdout, "Response from `ContactFieldApi.GetContactFieldSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetContactFieldSettingsRequest struct via the builder pattern


### Return type

[**GetContactFieldSettings200Response**](GetContactFieldSettings200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContactFields

> GetContactFieldsById200Response GetContactFields(ctx).Execute()

Retrieve contact fields



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
    resp, r, err := apiClient.ContactFieldApi.GetContactFields(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactFieldApi.GetContactFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContactFields`: GetContactFieldsById200Response
    fmt.Fprintf(os.Stdout, "Response from `ContactFieldApi.GetContactFields`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetContactFieldsRequest struct via the builder pattern


### Return type

[**GetContactFieldsById200Response**](GetContactFieldsById200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContactFieldsById

> GetContactFieldsById200Response GetContactFieldsById(ctx, contactCustomFieldId).Execute()

Retrieve contact fields



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
    contactCustomFieldId := float32(8.14) // float32 | id of the contact field

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactFieldApi.GetContactFieldsById(context.Background(), contactCustomFieldId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactFieldApi.GetContactFieldsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContactFieldsById`: GetContactFieldsById200Response
    fmt.Fprintf(os.Stdout, "Response from `ContactFieldApi.GetContactFieldsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactCustomFieldId** | **float32** | id of the contact field | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContactFieldsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetContactFieldsById200Response**](GetContactFieldsById200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlaceholder

> GetPlaceholder200Response GetPlaceholder(ctx).ObjectName(objectName).SubObjectName(subObjectName).Execute()

Retrieve Placeholders



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
    objectName := "objectName_example" // string | Model name
    subObjectName := "subObjectName_example" // string | Sub model name, required if you have \"Email\" at objectName (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactFieldApi.GetPlaceholder(context.Background()).ObjectName(objectName).SubObjectName(subObjectName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactFieldApi.GetPlaceholder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlaceholder`: GetPlaceholder200Response
    fmt.Fprintf(os.Stdout, "Response from `ContactFieldApi.GetPlaceholder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPlaceholderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectName** | **string** | Model name | 
 **subObjectName** | **string** | Sub model name, required if you have \&quot;Email\&quot; at objectName | 

### Return type

[**GetPlaceholder200Response**](GetPlaceholder200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReferenceCount

> GetReferenceCount200Response GetReferenceCount(ctx, contactCustomFieldSettingId).Execute()

Receive count reference



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
    contactCustomFieldSettingId := int32(56) // int32 | ID of contact field you want to get the reference count

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactFieldApi.GetReferenceCount(context.Background(), contactCustomFieldSettingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactFieldApi.GetReferenceCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReferenceCount`: GetReferenceCount200Response
    fmt.Fprintf(os.Stdout, "Response from `ContactFieldApi.GetReferenceCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactCustomFieldSettingId** | **int32** | ID of contact field you want to get the reference count | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReferenceCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetReferenceCount200Response**](GetReferenceCount200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContactFieldSetting

> GetContactFieldSettings200Response UpdateContactFieldSetting(ctx, contactCustomFieldSettingId).ModelContactCustomFieldSettingUpdate(modelContactCustomFieldSettingUpdate).Execute()

Update contact field setting



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
    contactCustomFieldSettingId := int32(56) // int32 | ID of contact field setting you want to update
    modelContactCustomFieldSettingUpdate := *openapiclient.NewModelContactCustomFieldSettingUpdate() // ModelContactCustomFieldSettingUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactFieldApi.UpdateContactFieldSetting(context.Background(), contactCustomFieldSettingId).ModelContactCustomFieldSettingUpdate(modelContactCustomFieldSettingUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactFieldApi.UpdateContactFieldSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateContactFieldSetting`: GetContactFieldSettings200Response
    fmt.Fprintf(os.Stdout, "Response from `ContactFieldApi.UpdateContactFieldSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactCustomFieldSettingId** | **int32** | ID of contact field setting you want to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContactFieldSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelContactCustomFieldSettingUpdate** | [**ModelContactCustomFieldSettingUpdate**](ModelContactCustomFieldSettingUpdate.md) |  | 

### Return type

[**GetContactFieldSettings200Response**](GetContactFieldSettings200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContactfield

> GetContactFieldsById200Response UpdateContactfield(ctx, contactCustomFieldId).ModelContactCustomFieldUpdate(modelContactCustomFieldUpdate).Execute()

Update a contact field



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
    contactCustomFieldId := float32(8.14) // float32 | id of the contact field
    modelContactCustomFieldUpdate := *openapiclient.NewModelContactCustomFieldUpdate() // ModelContactCustomFieldUpdate | Update data (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactFieldApi.UpdateContactfield(context.Background(), contactCustomFieldId).ModelContactCustomFieldUpdate(modelContactCustomFieldUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactFieldApi.UpdateContactfield``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateContactfield`: GetContactFieldsById200Response
    fmt.Fprintf(os.Stdout, "Response from `ContactFieldApi.UpdateContactfield`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactCustomFieldId** | **float32** | id of the contact field | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContactfieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelContactCustomFieldUpdate** | [**ModelContactCustomFieldUpdate**](ModelContactCustomFieldUpdate.md) | Update data | 

### Return type

[**GetContactFieldsById200Response**](GetContactFieldsById200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

