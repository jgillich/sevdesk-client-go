# SaveVoucherResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Voucher** | Pointer to [**ModelVoucherResponse**](ModelVoucherResponse.md) |  | [optional] 
**VoucherPos** | Pointer to [**ModelVoucherPosResponse**](ModelVoucherPosResponse.md) |  | [optional] 
**Filename** | Pointer to ***os.File** | Filename of a previously upload file which should be attached. | [optional] 

## Methods

### NewSaveVoucherResponse

`func NewSaveVoucherResponse() *SaveVoucherResponse`

NewSaveVoucherResponse instantiates a new SaveVoucherResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveVoucherResponseWithDefaults

`func NewSaveVoucherResponseWithDefaults() *SaveVoucherResponse`

NewSaveVoucherResponseWithDefaults instantiates a new SaveVoucherResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVoucher

`func (o *SaveVoucherResponse) GetVoucher() ModelVoucherResponse`

GetVoucher returns the Voucher field if non-nil, zero value otherwise.

### GetVoucherOk

`func (o *SaveVoucherResponse) GetVoucherOk() (*ModelVoucherResponse, bool)`

GetVoucherOk returns a tuple with the Voucher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucher

`func (o *SaveVoucherResponse) SetVoucher(v ModelVoucherResponse)`

SetVoucher sets Voucher field to given value.

### HasVoucher

`func (o *SaveVoucherResponse) HasVoucher() bool`

HasVoucher returns a boolean if a field has been set.

### GetVoucherPos

`func (o *SaveVoucherResponse) GetVoucherPos() ModelVoucherPosResponse`

GetVoucherPos returns the VoucherPos field if non-nil, zero value otherwise.

### GetVoucherPosOk

`func (o *SaveVoucherResponse) GetVoucherPosOk() (*ModelVoucherPosResponse, bool)`

GetVoucherPosOk returns a tuple with the VoucherPos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherPos

`func (o *SaveVoucherResponse) SetVoucherPos(v ModelVoucherPosResponse)`

SetVoucherPos sets VoucherPos field to given value.

### HasVoucherPos

`func (o *SaveVoucherResponse) HasVoucherPos() bool`

HasVoucherPos returns a boolean if a field has been set.

### GetFilename

`func (o *SaveVoucherResponse) GetFilename() *os.File`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *SaveVoucherResponse) GetFilenameOk() (**os.File, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *SaveVoucherResponse) SetFilename(v *os.File)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *SaveVoucherResponse) HasFilename() bool`

HasFilename returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


