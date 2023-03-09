# VoucherUploadFileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | Pointer to ***os.File** | The file to upload | [optional] 

## Methods

### NewVoucherUploadFileRequest

`func NewVoucherUploadFileRequest() *VoucherUploadFileRequest`

NewVoucherUploadFileRequest instantiates a new VoucherUploadFileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoucherUploadFileRequestWithDefaults

`func NewVoucherUploadFileRequestWithDefaults() *VoucherUploadFileRequest`

NewVoucherUploadFileRequestWithDefaults instantiates a new VoucherUploadFileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *VoucherUploadFileRequest) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *VoucherUploadFileRequest) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *VoucherUploadFileRequest) SetFile(v *os.File)`

SetFile sets File field to given value.

### HasFile

`func (o *VoucherUploadFileRequest) HasFile() bool`

HasFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


