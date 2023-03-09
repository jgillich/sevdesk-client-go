# ExportVoucherSevQueryParameterFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | Pointer to **time.Time** | Start date of the voucher | [optional] 
**EndDate** | Pointer to **time.Time** | End date of the voucher | [optional] 
**StartPayDate** | Pointer to **time.Time** | Start pay date of the voucher | [optional] 
**EndPayDate** | Pointer to **time.Time** | End pay date of the voucher | [optional] 
**Contact** | Pointer to [**ExportVoucherSevQueryParameterFilterContact**](ExportVoucherSevQueryParameterFilterContact.md) |  | [optional] 
**StartAmount** | Pointer to **int32** | filters the vouchers by amount | [optional] 
**EndAmount** | Pointer to **int32** | filters the vouchers by amount | [optional] 

## Methods

### NewExportVoucherSevQueryParameterFilter

`func NewExportVoucherSevQueryParameterFilter() *ExportVoucherSevQueryParameterFilter`

NewExportVoucherSevQueryParameterFilter instantiates a new ExportVoucherSevQueryParameterFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportVoucherSevQueryParameterFilterWithDefaults

`func NewExportVoucherSevQueryParameterFilterWithDefaults() *ExportVoucherSevQueryParameterFilter`

NewExportVoucherSevQueryParameterFilterWithDefaults instantiates a new ExportVoucherSevQueryParameterFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *ExportVoucherSevQueryParameterFilter) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ExportVoucherSevQueryParameterFilter) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ExportVoucherSevQueryParameterFilter) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ExportVoucherSevQueryParameterFilter) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ExportVoucherSevQueryParameterFilter) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ExportVoucherSevQueryParameterFilter) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ExportVoucherSevQueryParameterFilter) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ExportVoucherSevQueryParameterFilter) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStartPayDate

`func (o *ExportVoucherSevQueryParameterFilter) GetStartPayDate() time.Time`

GetStartPayDate returns the StartPayDate field if non-nil, zero value otherwise.

### GetStartPayDateOk

`func (o *ExportVoucherSevQueryParameterFilter) GetStartPayDateOk() (*time.Time, bool)`

GetStartPayDateOk returns a tuple with the StartPayDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPayDate

`func (o *ExportVoucherSevQueryParameterFilter) SetStartPayDate(v time.Time)`

SetStartPayDate sets StartPayDate field to given value.

### HasStartPayDate

`func (o *ExportVoucherSevQueryParameterFilter) HasStartPayDate() bool`

HasStartPayDate returns a boolean if a field has been set.

### GetEndPayDate

`func (o *ExportVoucherSevQueryParameterFilter) GetEndPayDate() time.Time`

GetEndPayDate returns the EndPayDate field if non-nil, zero value otherwise.

### GetEndPayDateOk

`func (o *ExportVoucherSevQueryParameterFilter) GetEndPayDateOk() (*time.Time, bool)`

GetEndPayDateOk returns a tuple with the EndPayDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPayDate

`func (o *ExportVoucherSevQueryParameterFilter) SetEndPayDate(v time.Time)`

SetEndPayDate sets EndPayDate field to given value.

### HasEndPayDate

`func (o *ExportVoucherSevQueryParameterFilter) HasEndPayDate() bool`

HasEndPayDate returns a boolean if a field has been set.

### GetContact

`func (o *ExportVoucherSevQueryParameterFilter) GetContact() ExportVoucherSevQueryParameterFilterContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ExportVoucherSevQueryParameterFilter) GetContactOk() (*ExportVoucherSevQueryParameterFilterContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ExportVoucherSevQueryParameterFilter) SetContact(v ExportVoucherSevQueryParameterFilterContact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *ExportVoucherSevQueryParameterFilter) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetStartAmount

`func (o *ExportVoucherSevQueryParameterFilter) GetStartAmount() int32`

GetStartAmount returns the StartAmount field if non-nil, zero value otherwise.

### GetStartAmountOk

`func (o *ExportVoucherSevQueryParameterFilter) GetStartAmountOk() (*int32, bool)`

GetStartAmountOk returns a tuple with the StartAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAmount

`func (o *ExportVoucherSevQueryParameterFilter) SetStartAmount(v int32)`

SetStartAmount sets StartAmount field to given value.

### HasStartAmount

`func (o *ExportVoucherSevQueryParameterFilter) HasStartAmount() bool`

HasStartAmount returns a boolean if a field has been set.

### GetEndAmount

`func (o *ExportVoucherSevQueryParameterFilter) GetEndAmount() int32`

GetEndAmount returns the EndAmount field if non-nil, zero value otherwise.

### GetEndAmountOk

`func (o *ExportVoucherSevQueryParameterFilter) GetEndAmountOk() (*int32, bool)`

GetEndAmountOk returns a tuple with the EndAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAmount

`func (o *ExportVoucherSevQueryParameterFilter) SetEndAmount(v int32)`

SetEndAmount sets EndAmount field to given value.

### HasEndAmount

`func (o *ExportVoucherSevQueryParameterFilter) HasEndAmount() bool`

HasEndAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


