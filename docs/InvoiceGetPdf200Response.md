# InvoiceGetPdf200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | Pointer to **string** |  | [optional] 
**MimeType** | Pointer to **string** |  | [optional] 
**Content** | Pointer to ***os.File** |  | [optional] 
**Base64encoded** | Pointer to **string** |  | [optional] 

## Methods

### NewInvoiceGetPdf200Response

`func NewInvoiceGetPdf200Response() *InvoiceGetPdf200Response`

NewInvoiceGetPdf200Response instantiates a new InvoiceGetPdf200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceGetPdf200ResponseWithDefaults

`func NewInvoiceGetPdf200ResponseWithDefaults() *InvoiceGetPdf200Response`

NewInvoiceGetPdf200ResponseWithDefaults instantiates a new InvoiceGetPdf200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *InvoiceGetPdf200Response) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *InvoiceGetPdf200Response) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *InvoiceGetPdf200Response) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *InvoiceGetPdf200Response) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetMimeType

`func (o *InvoiceGetPdf200Response) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *InvoiceGetPdf200Response) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *InvoiceGetPdf200Response) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *InvoiceGetPdf200Response) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetContent

`func (o *InvoiceGetPdf200Response) GetContent() *os.File`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *InvoiceGetPdf200Response) GetContentOk() (**os.File, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *InvoiceGetPdf200Response) SetContent(v *os.File)`

SetContent sets Content field to given value.

### HasContent

`func (o *InvoiceGetPdf200Response) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetBase64encoded

`func (o *InvoiceGetPdf200Response) GetBase64encoded() string`

GetBase64encoded returns the Base64encoded field if non-nil, zero value otherwise.

### GetBase64encodedOk

`func (o *InvoiceGetPdf200Response) GetBase64encodedOk() (*string, bool)`

GetBase64encodedOk returns a tuple with the Base64encoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase64encoded

`func (o *InvoiceGetPdf200Response) SetBase64encoded(v string)`

SetBase64encoded sets Base64encoded field to given value.

### HasBase64encoded

`func (o *InvoiceGetPdf200Response) HasBase64encoded() bool`

HasBase64encoded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


