# SaveCreditNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreditNote** | [**ModelCreditNote**](ModelCreditNote.md) |  | 
**CreditNotePosSave** | Pointer to [**ModelCreditNotePos**](ModelCreditNotePos.md) |  | [optional] 
**CreditNotePosDelete** | Pointer to [**SaveCreditNoteCreditNotePosDelete**](SaveCreditNoteCreditNotePosDelete.md) |  | [optional] [default to null]
**Filename** | Pointer to ***os.File** | Filename of a previously upload file which should be attached. | [optional] 
**DiscountSave** | Pointer to [**SaveCreditNoteDiscountSave**](SaveCreditNoteDiscountSave.md) |  | [optional] [default to null]
**DiscountDelete** | Pointer to [**SaveCreditNoteDiscountDelete**](SaveCreditNoteDiscountDelete.md) |  | [optional] [default to null]

## Methods

### NewSaveCreditNote

`func NewSaveCreditNote(creditNote ModelCreditNote, ) *SaveCreditNote`

NewSaveCreditNote instantiates a new SaveCreditNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveCreditNoteWithDefaults

`func NewSaveCreditNoteWithDefaults() *SaveCreditNote`

NewSaveCreditNoteWithDefaults instantiates a new SaveCreditNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreditNote

`func (o *SaveCreditNote) GetCreditNote() ModelCreditNote`

GetCreditNote returns the CreditNote field if non-nil, zero value otherwise.

### GetCreditNoteOk

`func (o *SaveCreditNote) GetCreditNoteOk() (*ModelCreditNote, bool)`

GetCreditNoteOk returns a tuple with the CreditNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNote

`func (o *SaveCreditNote) SetCreditNote(v ModelCreditNote)`

SetCreditNote sets CreditNote field to given value.


### GetCreditNotePosSave

`func (o *SaveCreditNote) GetCreditNotePosSave() ModelCreditNotePos`

GetCreditNotePosSave returns the CreditNotePosSave field if non-nil, zero value otherwise.

### GetCreditNotePosSaveOk

`func (o *SaveCreditNote) GetCreditNotePosSaveOk() (*ModelCreditNotePos, bool)`

GetCreditNotePosSaveOk returns a tuple with the CreditNotePosSave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNotePosSave

`func (o *SaveCreditNote) SetCreditNotePosSave(v ModelCreditNotePos)`

SetCreditNotePosSave sets CreditNotePosSave field to given value.

### HasCreditNotePosSave

`func (o *SaveCreditNote) HasCreditNotePosSave() bool`

HasCreditNotePosSave returns a boolean if a field has been set.

### GetCreditNotePosDelete

`func (o *SaveCreditNote) GetCreditNotePosDelete() SaveCreditNoteCreditNotePosDelete`

GetCreditNotePosDelete returns the CreditNotePosDelete field if non-nil, zero value otherwise.

### GetCreditNotePosDeleteOk

`func (o *SaveCreditNote) GetCreditNotePosDeleteOk() (*SaveCreditNoteCreditNotePosDelete, bool)`

GetCreditNotePosDeleteOk returns a tuple with the CreditNotePosDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNotePosDelete

`func (o *SaveCreditNote) SetCreditNotePosDelete(v SaveCreditNoteCreditNotePosDelete)`

SetCreditNotePosDelete sets CreditNotePosDelete field to given value.

### HasCreditNotePosDelete

`func (o *SaveCreditNote) HasCreditNotePosDelete() bool`

HasCreditNotePosDelete returns a boolean if a field has been set.

### GetFilename

`func (o *SaveCreditNote) GetFilename() *os.File`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *SaveCreditNote) GetFilenameOk() (**os.File, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *SaveCreditNote) SetFilename(v *os.File)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *SaveCreditNote) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetDiscountSave

`func (o *SaveCreditNote) GetDiscountSave() SaveCreditNoteDiscountSave`

GetDiscountSave returns the DiscountSave field if non-nil, zero value otherwise.

### GetDiscountSaveOk

`func (o *SaveCreditNote) GetDiscountSaveOk() (*SaveCreditNoteDiscountSave, bool)`

GetDiscountSaveOk returns a tuple with the DiscountSave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountSave

`func (o *SaveCreditNote) SetDiscountSave(v SaveCreditNoteDiscountSave)`

SetDiscountSave sets DiscountSave field to given value.

### HasDiscountSave

`func (o *SaveCreditNote) HasDiscountSave() bool`

HasDiscountSave returns a boolean if a field has been set.

### GetDiscountDelete

`func (o *SaveCreditNote) GetDiscountDelete() SaveCreditNoteDiscountDelete`

GetDiscountDelete returns the DiscountDelete field if non-nil, zero value otherwise.

### GetDiscountDeleteOk

`func (o *SaveCreditNote) GetDiscountDeleteOk() (*SaveCreditNoteDiscountDelete, bool)`

GetDiscountDeleteOk returns a tuple with the DiscountDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountDelete

`func (o *SaveCreditNote) SetDiscountDelete(v SaveCreditNoteDiscountDelete)`

SetDiscountDelete sets DiscountDelete field to given value.

### HasDiscountDelete

`func (o *SaveCreditNote) HasDiscountDelete() bool`

HasDiscountDelete returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


