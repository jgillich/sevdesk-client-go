# \ContactApi

All URIs are relative to *https://my.sevdesk.de/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContactCustomerNumberAvailabilityCheck**](ContactApi.md#ContactCustomerNumberAvailabilityCheck) | **Get** /Contact/Mapper/checkCustomerNumberAvailability | Check if a customer number is available
[**CreateContact**](ContactApi.md#CreateContact) | **Post** /Contact | Create a new contact
[**DeleteContact**](ContactApi.md#DeleteContact) | **Delete** /Contact/{contactId} | Deletes a contact
[**FindContactsByCustomFieldValue**](ContactApi.md#FindContactsByCustomFieldValue) | **Get** /Contact/Factory/findContactsByCustomFieldValue | Find contacts by custom field value
[**GetContactById**](ContactApi.md#GetContactById) | **Get** /Contact/{contactId} | Find contact by ID
[**GetContactTabsItemCountById**](ContactApi.md#GetContactTabsItemCountById) | **Get** /Contact/{contactId}/getTabsItemCount | Get number of all items
[**GetContacts**](ContactApi.md#GetContacts) | **Get** /Contact | Retrieve contacts
[**GetNextCustomerNumber**](ContactApi.md#GetNextCustomerNumber) | **Get** /Contact/Factory/getNextCustomerNumber | Get next free customer number
[**UpdateContact**](ContactApi.md#UpdateContact) | **Put** /Contact/{contactId} | Update a existing contact



## ContactCustomerNumberAvailabilityCheck

> GetIsInvoicePartiallyPaid200Response ContactCustomerNumberAvailabilityCheck(ctx).CustomerNumber(customerNumber).Execute()

Check if a customer number is available



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
    customerNumber := "customerNumber_example" // string | The customer number to be checked. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactApi.ContactCustomerNumberAvailabilityCheck(context.Background()).CustomerNumber(customerNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactApi.ContactCustomerNumberAvailabilityCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContactCustomerNumberAvailabilityCheck`: GetIsInvoicePartiallyPaid200Response
    fmt.Fprintf(os.Stdout, "Response from `ContactApi.ContactCustomerNumberAvailabilityCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContactCustomerNumberAvailabilityCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerNumber** | **string** | The customer number to be checked. | 

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


## CreateContact

> GetContacts200Response CreateContact(ctx).ModelContact(modelContact).Execute()

Create a new contact



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
    modelContact := *openapiclient.NewModelContact(*openapiclient.NewModelContactCategory(int32(3), "Category")) // ModelContact | Creation data (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactApi.CreateContact(context.Background()).ModelContact(modelContact).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactApi.CreateContact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateContact`: GetContacts200Response
    fmt.Fprintf(os.Stdout, "Response from `ContactApi.CreateContact`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelContact** | [**ModelContact**](ModelContact.md) | Creation data | 

### Return type

[**GetContacts200Response**](GetContacts200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContact

> DeleteAccountingContact200Response DeleteContact(ctx, contactId).Execute()

Deletes a contact

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
    contactId := int32(56) // int32 | Id of contact resource to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactApi.DeleteContact(context.Background(), contactId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactApi.DeleteContact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteContact`: DeleteAccountingContact200Response
    fmt.Fprintf(os.Stdout, "Response from `ContactApi.DeleteContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **int32** | Id of contact resource to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContactRequest struct via the builder pattern


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


## FindContactsByCustomFieldValue

> GetContacts200Response FindContactsByCustomFieldValue(ctx).Value(value).CustomFieldSettingId(customFieldSettingId).CustomFieldSettingObjectName(customFieldSettingObjectName).CustomFieldName(customFieldName).Execute()

Find contacts by custom field value



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
    value := "value_example" // string | The value to be checked.
    customFieldSettingId := "customFieldSettingId_example" // string | ID of ContactCustomFieldSetting for which the value has to be checked. (optional)
    customFieldSettingObjectName := "ContactCustomFieldSetting" // string | Object name. Only needed if you also defined the ID of a ContactCustomFieldSetting. (optional)
    customFieldName := "customFieldName_example" // string | The ContactCustomFieldSetting name, if no ContactCustomFieldSetting is provided. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactApi.FindContactsByCustomFieldValue(context.Background()).Value(value).CustomFieldSettingId(customFieldSettingId).CustomFieldSettingObjectName(customFieldSettingObjectName).CustomFieldName(customFieldName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactApi.FindContactsByCustomFieldValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindContactsByCustomFieldValue`: GetContacts200Response
    fmt.Fprintf(os.Stdout, "Response from `ContactApi.FindContactsByCustomFieldValue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindContactsByCustomFieldValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **value** | **string** | The value to be checked. | 
 **customFieldSettingId** | **string** | ID of ContactCustomFieldSetting for which the value has to be checked. | 
 **customFieldSettingObjectName** | **string** | Object name. Only needed if you also defined the ID of a ContactCustomFieldSetting. | 
 **customFieldName** | **string** | The ContactCustomFieldSetting name, if no ContactCustomFieldSetting is provided. | 

### Return type

[**GetContacts200Response**](GetContacts200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContactById

> GetContacts200Response GetContactById(ctx, contactId).Execute()

Find contact by ID



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
    contactId := int32(56) // int32 | ID of contact to return

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactApi.GetContactById(context.Background(), contactId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactApi.GetContactById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContactById`: GetContacts200Response
    fmt.Fprintf(os.Stdout, "Response from `ContactApi.GetContactById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **int32** | ID of contact to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContactByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetContacts200Response**](GetContacts200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContactTabsItemCountById

> GetContactTabsItemCountById200Response GetContactTabsItemCountById(ctx, contactId).Execute()

Get number of all items



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
    contactId := int32(56) // int32 | ID of contact to return

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactApi.GetContactTabsItemCountById(context.Background(), contactId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactApi.GetContactTabsItemCountById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContactTabsItemCountById`: GetContactTabsItemCountById200Response
    fmt.Fprintf(os.Stdout, "Response from `ContactApi.GetContactTabsItemCountById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **int32** | ID of contact to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContactTabsItemCountByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetContactTabsItemCountById200Response**](GetContactTabsItemCountById200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContacts

> GetContacts200Response GetContacts(ctx).Depth(depth).CustomerNumber(customerNumber).Execute()

Retrieve contacts



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
    depth := "depth_example" // string | Defines if both organizations <b>and</b> persons should be returned.<br>       '0' -> only organizations, '1' -> organizations and persons (optional)
    customerNumber := "customerNumber_example" // string | Retrieve all contacts with this customer number (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactApi.GetContacts(context.Background()).Depth(depth).CustomerNumber(customerNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactApi.GetContacts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContacts`: GetContacts200Response
    fmt.Fprintf(os.Stdout, "Response from `ContactApi.GetContacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **depth** | **string** | Defines if both organizations &lt;b&gt;and&lt;/b&gt; persons should be returned.&lt;br&gt;       &#39;0&#39; -&gt; only organizations, &#39;1&#39; -&gt; organizations and persons | 
 **customerNumber** | **string** | Retrieve all contacts with this customer number | 

### Return type

[**GetContacts200Response**](GetContacts200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNextCustomerNumber

> GetNextCustomerNumber200Response GetNextCustomerNumber(ctx).Execute()

Get next free customer number



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
    resp, r, err := apiClient.ContactApi.GetNextCustomerNumber(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactApi.GetNextCustomerNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNextCustomerNumber`: GetNextCustomerNumber200Response
    fmt.Fprintf(os.Stdout, "Response from `ContactApi.GetNextCustomerNumber`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNextCustomerNumberRequest struct via the builder pattern


### Return type

[**GetNextCustomerNumber200Response**](GetNextCustomerNumber200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContact

> GetContacts200Response UpdateContact(ctx, contactId).ModelContactUpdate(modelContactUpdate).Execute()

Update a existing contact



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
    contactId := int32(56) // int32 | ID of contact to update
    modelContactUpdate := *openapiclient.NewModelContactUpdate() // ModelContactUpdate | Update data (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactApi.UpdateContact(context.Background(), contactId).ModelContactUpdate(modelContactUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactApi.UpdateContact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateContact`: GetContacts200Response
    fmt.Fprintf(os.Stdout, "Response from `ContactApi.UpdateContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **int32** | ID of contact to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelContactUpdate** | [**ModelContactUpdate**](ModelContactUpdate.md) | Update data | 

### Return type

[**GetContacts200Response**](GetContacts200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

