# SaveInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invoice** | [**ModelInvoice**](ModelInvoice.md) |  | 
**InvoicePosSave** | Pointer to [**ModelInvoicePos**](ModelInvoicePos.md) |  | [optional] 
**InvoicePosDelete** | Pointer to [**SaveInvoiceInvoicePosDelete**](SaveInvoiceInvoicePosDelete.md) |  | [optional] [default to null]
**Filename** | Pointer to ***os.File** | Filename of a previously upload file which should be attached. | [optional] 
**DiscountSave** | Pointer to [**SaveInvoiceDiscountSave**](SaveInvoiceDiscountSave.md) |  | [optional] 
**DiscountDelete** | Pointer to [**SaveInvoiceDiscountDelete**](SaveInvoiceDiscountDelete.md) |  | [optional] 

## Methods

### NewSaveInvoice

`func NewSaveInvoice(invoice ModelInvoice, ) *SaveInvoice`

NewSaveInvoice instantiates a new SaveInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveInvoiceWithDefaults

`func NewSaveInvoiceWithDefaults() *SaveInvoice`

NewSaveInvoiceWithDefaults instantiates a new SaveInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoice

`func (o *SaveInvoice) GetInvoice() ModelInvoice`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *SaveInvoice) GetInvoiceOk() (*ModelInvoice, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *SaveInvoice) SetInvoice(v ModelInvoice)`

SetInvoice sets Invoice field to given value.


### GetInvoicePosSave

`func (o *SaveInvoice) GetInvoicePosSave() ModelInvoicePos`

GetInvoicePosSave returns the InvoicePosSave field if non-nil, zero value otherwise.

### GetInvoicePosSaveOk

`func (o *SaveInvoice) GetInvoicePosSaveOk() (*ModelInvoicePos, bool)`

GetInvoicePosSaveOk returns a tuple with the InvoicePosSave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePosSave

`func (o *SaveInvoice) SetInvoicePosSave(v ModelInvoicePos)`

SetInvoicePosSave sets InvoicePosSave field to given value.

### HasInvoicePosSave

`func (o *SaveInvoice) HasInvoicePosSave() bool`

HasInvoicePosSave returns a boolean if a field has been set.

### GetInvoicePosDelete

`func (o *SaveInvoice) GetInvoicePosDelete() SaveInvoiceInvoicePosDelete`

GetInvoicePosDelete returns the InvoicePosDelete field if non-nil, zero value otherwise.

### GetInvoicePosDeleteOk

`func (o *SaveInvoice) GetInvoicePosDeleteOk() (*SaveInvoiceInvoicePosDelete, bool)`

GetInvoicePosDeleteOk returns a tuple with the InvoicePosDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePosDelete

`func (o *SaveInvoice) SetInvoicePosDelete(v SaveInvoiceInvoicePosDelete)`

SetInvoicePosDelete sets InvoicePosDelete field to given value.

### HasInvoicePosDelete

`func (o *SaveInvoice) HasInvoicePosDelete() bool`

HasInvoicePosDelete returns a boolean if a field has been set.

### GetFilename

`func (o *SaveInvoice) GetFilename() *os.File`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *SaveInvoice) GetFilenameOk() (**os.File, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *SaveInvoice) SetFilename(v *os.File)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *SaveInvoice) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetDiscountSave

`func (o *SaveInvoice) GetDiscountSave() SaveInvoiceDiscountSave`

GetDiscountSave returns the DiscountSave field if non-nil, zero value otherwise.

### GetDiscountSaveOk

`func (o *SaveInvoice) GetDiscountSaveOk() (*SaveInvoiceDiscountSave, bool)`

GetDiscountSaveOk returns a tuple with the DiscountSave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountSave

`func (o *SaveInvoice) SetDiscountSave(v SaveInvoiceDiscountSave)`

SetDiscountSave sets DiscountSave field to given value.

### HasDiscountSave

`func (o *SaveInvoice) HasDiscountSave() bool`

HasDiscountSave returns a boolean if a field has been set.

### GetDiscountDelete

`func (o *SaveInvoice) GetDiscountDelete() SaveInvoiceDiscountDelete`

GetDiscountDelete returns the DiscountDelete field if non-nil, zero value otherwise.

### GetDiscountDeleteOk

`func (o *SaveInvoice) GetDiscountDeleteOk() (*SaveInvoiceDiscountDelete, bool)`

GetDiscountDeleteOk returns a tuple with the DiscountDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountDelete

`func (o *SaveInvoice) SetDiscountDelete(v SaveInvoiceDiscountDelete)`

SetDiscountDelete sets DiscountDelete field to given value.

### HasDiscountDelete

`func (o *SaveInvoice) HasDiscountDelete() bool`

HasDiscountDelete returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


