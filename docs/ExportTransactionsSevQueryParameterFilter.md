# ExportTransactionsSevQueryParameterFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymtPurpose** | Pointer to **string** | the payment purpose | [optional] 
**Name** | Pointer to **string** | the name of the payee/payer | [optional] 
**StartDate** | Pointer to **time.Time** | Start date of the transactions | [optional] 
**EndDate** | Pointer to **time.Time** | End date of the transactions | [optional] 
**StartAmount** | Pointer to **int32** | filters the transactions by amount | [optional] 
**EndAmount** | Pointer to **int32** | filters the transactions by amount | [optional] 
**CheckAccount** | Pointer to [**ExportTransactionsSevQueryParameterFilterCheckAccount**](ExportTransactionsSevQueryParameterFilterCheckAccount.md) |  | [optional] 

## Methods

### NewExportTransactionsSevQueryParameterFilter

`func NewExportTransactionsSevQueryParameterFilter() *ExportTransactionsSevQueryParameterFilter`

NewExportTransactionsSevQueryParameterFilter instantiates a new ExportTransactionsSevQueryParameterFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportTransactionsSevQueryParameterFilterWithDefaults

`func NewExportTransactionsSevQueryParameterFilterWithDefaults() *ExportTransactionsSevQueryParameterFilter`

NewExportTransactionsSevQueryParameterFilterWithDefaults instantiates a new ExportTransactionsSevQueryParameterFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymtPurpose

`func (o *ExportTransactionsSevQueryParameterFilter) GetPaymtPurpose() string`

GetPaymtPurpose returns the PaymtPurpose field if non-nil, zero value otherwise.

### GetPaymtPurposeOk

`func (o *ExportTransactionsSevQueryParameterFilter) GetPaymtPurposeOk() (*string, bool)`

GetPaymtPurposeOk returns a tuple with the PaymtPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymtPurpose

`func (o *ExportTransactionsSevQueryParameterFilter) SetPaymtPurpose(v string)`

SetPaymtPurpose sets PaymtPurpose field to given value.

### HasPaymtPurpose

`func (o *ExportTransactionsSevQueryParameterFilter) HasPaymtPurpose() bool`

HasPaymtPurpose returns a boolean if a field has been set.

### GetName

`func (o *ExportTransactionsSevQueryParameterFilter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExportTransactionsSevQueryParameterFilter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExportTransactionsSevQueryParameterFilter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExportTransactionsSevQueryParameterFilter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartDate

`func (o *ExportTransactionsSevQueryParameterFilter) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ExportTransactionsSevQueryParameterFilter) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ExportTransactionsSevQueryParameterFilter) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ExportTransactionsSevQueryParameterFilter) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ExportTransactionsSevQueryParameterFilter) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ExportTransactionsSevQueryParameterFilter) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ExportTransactionsSevQueryParameterFilter) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ExportTransactionsSevQueryParameterFilter) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStartAmount

`func (o *ExportTransactionsSevQueryParameterFilter) GetStartAmount() int32`

GetStartAmount returns the StartAmount field if non-nil, zero value otherwise.

### GetStartAmountOk

`func (o *ExportTransactionsSevQueryParameterFilter) GetStartAmountOk() (*int32, bool)`

GetStartAmountOk returns a tuple with the StartAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAmount

`func (o *ExportTransactionsSevQueryParameterFilter) SetStartAmount(v int32)`

SetStartAmount sets StartAmount field to given value.

### HasStartAmount

`func (o *ExportTransactionsSevQueryParameterFilter) HasStartAmount() bool`

HasStartAmount returns a boolean if a field has been set.

### GetEndAmount

`func (o *ExportTransactionsSevQueryParameterFilter) GetEndAmount() int32`

GetEndAmount returns the EndAmount field if non-nil, zero value otherwise.

### GetEndAmountOk

`func (o *ExportTransactionsSevQueryParameterFilter) GetEndAmountOk() (*int32, bool)`

GetEndAmountOk returns a tuple with the EndAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAmount

`func (o *ExportTransactionsSevQueryParameterFilter) SetEndAmount(v int32)`

SetEndAmount sets EndAmount field to given value.

### HasEndAmount

`func (o *ExportTransactionsSevQueryParameterFilter) HasEndAmount() bool`

HasEndAmount returns a boolean if a field has been set.

### GetCheckAccount

`func (o *ExportTransactionsSevQueryParameterFilter) GetCheckAccount() ExportTransactionsSevQueryParameterFilterCheckAccount`

GetCheckAccount returns the CheckAccount field if non-nil, zero value otherwise.

### GetCheckAccountOk

`func (o *ExportTransactionsSevQueryParameterFilter) GetCheckAccountOk() (*ExportTransactionsSevQueryParameterFilterCheckAccount, bool)`

GetCheckAccountOk returns a tuple with the CheckAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckAccount

`func (o *ExportTransactionsSevQueryParameterFilter) SetCheckAccount(v ExportTransactionsSevQueryParameterFilterCheckAccount)`

SetCheckAccount sets CheckAccount field to given value.

### HasCheckAccount

`func (o *ExportTransactionsSevQueryParameterFilter) HasCheckAccount() bool`

HasCheckAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


