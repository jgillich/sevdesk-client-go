# ExportInvoiceSevQueryParameterFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceType** | Pointer to **[]string** | Type of invoices you want to export 1. RE - Rechnung 2. SR - Stornorechnung 3. TR - Teilrechnung 4. AR - Abschlagsrechnung 5. ER - Endrechnung 6. WKR - Wiederkehrende Rechnung 7. MA - Mahnung  | [optional] 
**StartDate** | Pointer to **time.Time** | Start date of the invoice | [optional] 
**EndDate** | Pointer to **time.Time** | End date of the invoice | [optional] 
**Contact** | Pointer to [**ExportInvoiceSevQueryParameterFilterContact**](ExportInvoiceSevQueryParameterFilterContact.md) |  | [optional] 
**StartAmount** | Pointer to **int32** | filters the invoices by amount | [optional] 
**EndAmount** | Pointer to **int32** | filters the invoices by amount | [optional] 

## Methods

### NewExportInvoiceSevQueryParameterFilter

`func NewExportInvoiceSevQueryParameterFilter() *ExportInvoiceSevQueryParameterFilter`

NewExportInvoiceSevQueryParameterFilter instantiates a new ExportInvoiceSevQueryParameterFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportInvoiceSevQueryParameterFilterWithDefaults

`func NewExportInvoiceSevQueryParameterFilterWithDefaults() *ExportInvoiceSevQueryParameterFilter`

NewExportInvoiceSevQueryParameterFilterWithDefaults instantiates a new ExportInvoiceSevQueryParameterFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceType

`func (o *ExportInvoiceSevQueryParameterFilter) GetInvoiceType() []string`

GetInvoiceType returns the InvoiceType field if non-nil, zero value otherwise.

### GetInvoiceTypeOk

`func (o *ExportInvoiceSevQueryParameterFilter) GetInvoiceTypeOk() (*[]string, bool)`

GetInvoiceTypeOk returns a tuple with the InvoiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceType

`func (o *ExportInvoiceSevQueryParameterFilter) SetInvoiceType(v []string)`

SetInvoiceType sets InvoiceType field to given value.

### HasInvoiceType

`func (o *ExportInvoiceSevQueryParameterFilter) HasInvoiceType() bool`

HasInvoiceType returns a boolean if a field has been set.

### GetStartDate

`func (o *ExportInvoiceSevQueryParameterFilter) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ExportInvoiceSevQueryParameterFilter) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ExportInvoiceSevQueryParameterFilter) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ExportInvoiceSevQueryParameterFilter) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ExportInvoiceSevQueryParameterFilter) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ExportInvoiceSevQueryParameterFilter) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ExportInvoiceSevQueryParameterFilter) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ExportInvoiceSevQueryParameterFilter) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetContact

`func (o *ExportInvoiceSevQueryParameterFilter) GetContact() ExportInvoiceSevQueryParameterFilterContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ExportInvoiceSevQueryParameterFilter) GetContactOk() (*ExportInvoiceSevQueryParameterFilterContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ExportInvoiceSevQueryParameterFilter) SetContact(v ExportInvoiceSevQueryParameterFilterContact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *ExportInvoiceSevQueryParameterFilter) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetStartAmount

`func (o *ExportInvoiceSevQueryParameterFilter) GetStartAmount() int32`

GetStartAmount returns the StartAmount field if non-nil, zero value otherwise.

### GetStartAmountOk

`func (o *ExportInvoiceSevQueryParameterFilter) GetStartAmountOk() (*int32, bool)`

GetStartAmountOk returns a tuple with the StartAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAmount

`func (o *ExportInvoiceSevQueryParameterFilter) SetStartAmount(v int32)`

SetStartAmount sets StartAmount field to given value.

### HasStartAmount

`func (o *ExportInvoiceSevQueryParameterFilter) HasStartAmount() bool`

HasStartAmount returns a boolean if a field has been set.

### GetEndAmount

`func (o *ExportInvoiceSevQueryParameterFilter) GetEndAmount() int32`

GetEndAmount returns the EndAmount field if non-nil, zero value otherwise.

### GetEndAmountOk

`func (o *ExportInvoiceSevQueryParameterFilter) GetEndAmountOk() (*int32, bool)`

GetEndAmountOk returns a tuple with the EndAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAmount

`func (o *ExportInvoiceSevQueryParameterFilter) SetEndAmount(v int32)`

SetEndAmount sets EndAmount field to given value.

### HasEndAmount

`func (o *ExportInvoiceSevQueryParameterFilter) HasEndAmount() bool`

HasEndAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


