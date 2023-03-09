# SaveInvoiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invoice** | Pointer to [**ModelInvoiceResponse**](ModelInvoiceResponse.md) |  | [optional] 
**InvoicePos** | Pointer to [**ModelInvoicePosResponse**](ModelInvoicePosResponse.md) |  | [optional] 
**Filename** | Pointer to ***os.File** | Filename of a previously upload file which should be attached. | [optional] 

## Methods

### NewSaveInvoiceResponse

`func NewSaveInvoiceResponse() *SaveInvoiceResponse`

NewSaveInvoiceResponse instantiates a new SaveInvoiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveInvoiceResponseWithDefaults

`func NewSaveInvoiceResponseWithDefaults() *SaveInvoiceResponse`

NewSaveInvoiceResponseWithDefaults instantiates a new SaveInvoiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoice

`func (o *SaveInvoiceResponse) GetInvoice() ModelInvoiceResponse`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *SaveInvoiceResponse) GetInvoiceOk() (*ModelInvoiceResponse, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *SaveInvoiceResponse) SetInvoice(v ModelInvoiceResponse)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *SaveInvoiceResponse) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetInvoicePos

`func (o *SaveInvoiceResponse) GetInvoicePos() ModelInvoicePosResponse`

GetInvoicePos returns the InvoicePos field if non-nil, zero value otherwise.

### GetInvoicePosOk

`func (o *SaveInvoiceResponse) GetInvoicePosOk() (*ModelInvoicePosResponse, bool)`

GetInvoicePosOk returns a tuple with the InvoicePos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePos

`func (o *SaveInvoiceResponse) SetInvoicePos(v ModelInvoicePosResponse)`

SetInvoicePos sets InvoicePos field to given value.

### HasInvoicePos

`func (o *SaveInvoiceResponse) HasInvoicePos() bool`

HasInvoicePos returns a boolean if a field has been set.

### GetFilename

`func (o *SaveInvoiceResponse) GetFilename() *os.File`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *SaveInvoiceResponse) GetFilenameOk() (**os.File, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *SaveInvoiceResponse) SetFilename(v *os.File)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *SaveInvoiceResponse) HasFilename() bool`

HasFilename returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


