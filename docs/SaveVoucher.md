# SaveVoucher

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Voucher** | [**ModelVoucher**](ModelVoucher.md) |  | 
**VoucherPosSave** | Pointer to [**ModelVoucherPos**](ModelVoucherPos.md) |  | [optional] 
**VoucherPosDelete** | Pointer to [**SaveVoucherVoucherPosDelete**](SaveVoucherVoucherPosDelete.md) |  | [optional] [default to null]
**Filename** | Pointer to ***os.File** | Filename of a previously upload file which should be attached. | [optional] 

## Methods

### NewSaveVoucher

`func NewSaveVoucher(voucher ModelVoucher, ) *SaveVoucher`

NewSaveVoucher instantiates a new SaveVoucher object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveVoucherWithDefaults

`func NewSaveVoucherWithDefaults() *SaveVoucher`

NewSaveVoucherWithDefaults instantiates a new SaveVoucher object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVoucher

`func (o *SaveVoucher) GetVoucher() ModelVoucher`

GetVoucher returns the Voucher field if non-nil, zero value otherwise.

### GetVoucherOk

`func (o *SaveVoucher) GetVoucherOk() (*ModelVoucher, bool)`

GetVoucherOk returns a tuple with the Voucher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucher

`func (o *SaveVoucher) SetVoucher(v ModelVoucher)`

SetVoucher sets Voucher field to given value.


### GetVoucherPosSave

`func (o *SaveVoucher) GetVoucherPosSave() ModelVoucherPos`

GetVoucherPosSave returns the VoucherPosSave field if non-nil, zero value otherwise.

### GetVoucherPosSaveOk

`func (o *SaveVoucher) GetVoucherPosSaveOk() (*ModelVoucherPos, bool)`

GetVoucherPosSaveOk returns a tuple with the VoucherPosSave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherPosSave

`func (o *SaveVoucher) SetVoucherPosSave(v ModelVoucherPos)`

SetVoucherPosSave sets VoucherPosSave field to given value.

### HasVoucherPosSave

`func (o *SaveVoucher) HasVoucherPosSave() bool`

HasVoucherPosSave returns a boolean if a field has been set.

### GetVoucherPosDelete

`func (o *SaveVoucher) GetVoucherPosDelete() SaveVoucherVoucherPosDelete`

GetVoucherPosDelete returns the VoucherPosDelete field if non-nil, zero value otherwise.

### GetVoucherPosDeleteOk

`func (o *SaveVoucher) GetVoucherPosDeleteOk() (*SaveVoucherVoucherPosDelete, bool)`

GetVoucherPosDeleteOk returns a tuple with the VoucherPosDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherPosDelete

`func (o *SaveVoucher) SetVoucherPosDelete(v SaveVoucherVoucherPosDelete)`

SetVoucherPosDelete sets VoucherPosDelete field to given value.

### HasVoucherPosDelete

`func (o *SaveVoucher) HasVoucherPosDelete() bool`

HasVoucherPosDelete returns a boolean if a field has been set.

### GetFilename

`func (o *SaveVoucher) GetFilename() *os.File`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *SaveVoucher) GetFilenameOk() (**os.File, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *SaveVoucher) SetFilename(v *os.File)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *SaveVoucher) HasFilename() bool`

HasFilename returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


