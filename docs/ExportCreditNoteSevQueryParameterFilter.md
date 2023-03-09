# ExportCreditNoteSevQueryParameterFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | Pointer to **time.Time** | Start date of the credit note | [optional] 
**EndDate** | Pointer to **time.Time** | End date of the credit note | [optional] 
**Contact** | Pointer to [**ExportCreditNoteSevQueryParameterFilterContact**](ExportCreditNoteSevQueryParameterFilterContact.md) |  | [optional] 
**StartAmount** | Pointer to **int32** | filters the credit notes by amount | [optional] 
**EndAmount** | Pointer to **int32** | filters the credit notes by amount | [optional] 

## Methods

### NewExportCreditNoteSevQueryParameterFilter

`func NewExportCreditNoteSevQueryParameterFilter() *ExportCreditNoteSevQueryParameterFilter`

NewExportCreditNoteSevQueryParameterFilter instantiates a new ExportCreditNoteSevQueryParameterFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportCreditNoteSevQueryParameterFilterWithDefaults

`func NewExportCreditNoteSevQueryParameterFilterWithDefaults() *ExportCreditNoteSevQueryParameterFilter`

NewExportCreditNoteSevQueryParameterFilterWithDefaults instantiates a new ExportCreditNoteSevQueryParameterFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *ExportCreditNoteSevQueryParameterFilter) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ExportCreditNoteSevQueryParameterFilter) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ExportCreditNoteSevQueryParameterFilter) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ExportCreditNoteSevQueryParameterFilter) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ExportCreditNoteSevQueryParameterFilter) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ExportCreditNoteSevQueryParameterFilter) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ExportCreditNoteSevQueryParameterFilter) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ExportCreditNoteSevQueryParameterFilter) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetContact

`func (o *ExportCreditNoteSevQueryParameterFilter) GetContact() ExportCreditNoteSevQueryParameterFilterContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ExportCreditNoteSevQueryParameterFilter) GetContactOk() (*ExportCreditNoteSevQueryParameterFilterContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ExportCreditNoteSevQueryParameterFilter) SetContact(v ExportCreditNoteSevQueryParameterFilterContact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *ExportCreditNoteSevQueryParameterFilter) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetStartAmount

`func (o *ExportCreditNoteSevQueryParameterFilter) GetStartAmount() int32`

GetStartAmount returns the StartAmount field if non-nil, zero value otherwise.

### GetStartAmountOk

`func (o *ExportCreditNoteSevQueryParameterFilter) GetStartAmountOk() (*int32, bool)`

GetStartAmountOk returns a tuple with the StartAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAmount

`func (o *ExportCreditNoteSevQueryParameterFilter) SetStartAmount(v int32)`

SetStartAmount sets StartAmount field to given value.

### HasStartAmount

`func (o *ExportCreditNoteSevQueryParameterFilter) HasStartAmount() bool`

HasStartAmount returns a boolean if a field has been set.

### GetEndAmount

`func (o *ExportCreditNoteSevQueryParameterFilter) GetEndAmount() int32`

GetEndAmount returns the EndAmount field if non-nil, zero value otherwise.

### GetEndAmountOk

`func (o *ExportCreditNoteSevQueryParameterFilter) GetEndAmountOk() (*int32, bool)`

GetEndAmountOk returns a tuple with the EndAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAmount

`func (o *ExportCreditNoteSevQueryParameterFilter) SetEndAmount(v int32)`

SetEndAmount sets EndAmount field to given value.

### HasEndAmount

`func (o *ExportCreditNoteSevQueryParameterFilter) HasEndAmount() bool`

HasEndAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


