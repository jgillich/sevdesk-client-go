# SaveCreditNoteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Voucher** | Pointer to [**ModelCreditNoteResponse**](ModelCreditNoteResponse.md) |  | [optional] 
**VoucherPos** | Pointer to [**ModelCreditNotePosResponse**](ModelCreditNotePosResponse.md) |  | [optional] 
**Filename** | Pointer to ***os.File** | Filename of a previously upload file which should be attached. | [optional] 

## Methods

### NewSaveCreditNoteResponse

`func NewSaveCreditNoteResponse() *SaveCreditNoteResponse`

NewSaveCreditNoteResponse instantiates a new SaveCreditNoteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveCreditNoteResponseWithDefaults

`func NewSaveCreditNoteResponseWithDefaults() *SaveCreditNoteResponse`

NewSaveCreditNoteResponseWithDefaults instantiates a new SaveCreditNoteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVoucher

`func (o *SaveCreditNoteResponse) GetVoucher() ModelCreditNoteResponse`

GetVoucher returns the Voucher field if non-nil, zero value otherwise.

### GetVoucherOk

`func (o *SaveCreditNoteResponse) GetVoucherOk() (*ModelCreditNoteResponse, bool)`

GetVoucherOk returns a tuple with the Voucher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucher

`func (o *SaveCreditNoteResponse) SetVoucher(v ModelCreditNoteResponse)`

SetVoucher sets Voucher field to given value.

### HasVoucher

`func (o *SaveCreditNoteResponse) HasVoucher() bool`

HasVoucher returns a boolean if a field has been set.

### GetVoucherPos

`func (o *SaveCreditNoteResponse) GetVoucherPos() ModelCreditNotePosResponse`

GetVoucherPos returns the VoucherPos field if non-nil, zero value otherwise.

### GetVoucherPosOk

`func (o *SaveCreditNoteResponse) GetVoucherPosOk() (*ModelCreditNotePosResponse, bool)`

GetVoucherPosOk returns a tuple with the VoucherPos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherPos

`func (o *SaveCreditNoteResponse) SetVoucherPos(v ModelCreditNotePosResponse)`

SetVoucherPos sets VoucherPos field to given value.

### HasVoucherPos

`func (o *SaveCreditNoteResponse) HasVoucherPos() bool`

HasVoucherPos returns a boolean if a field has been set.

### GetFilename

`func (o *SaveCreditNoteResponse) GetFilename() *os.File`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *SaveCreditNoteResponse) GetFilenameOk() (**os.File, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *SaveCreditNoteResponse) SetFilename(v *os.File)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *SaveCreditNoteResponse) HasFilename() bool`

HasFilename returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


