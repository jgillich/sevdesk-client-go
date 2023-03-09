# ReportContactSevQueryParameterFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Zip** | Pointer to **int32** | filters the contacts by zip code | [optional] 
**City** | Pointer to **string** | filters the contacts by city | [optional] 
**Country** | Pointer to [**ReportContactSevQueryParameterFilterCountry**](ReportContactSevQueryParameterFilterCountry.md) |  | [optional] 
**Depth** | Pointer to **bool** | export only organisations | [optional] 
**OnlyPeople** | Pointer to **bool** | export only people | [optional] 

## Methods

### NewReportContactSevQueryParameterFilter

`func NewReportContactSevQueryParameterFilter() *ReportContactSevQueryParameterFilter`

NewReportContactSevQueryParameterFilter instantiates a new ReportContactSevQueryParameterFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportContactSevQueryParameterFilterWithDefaults

`func NewReportContactSevQueryParameterFilterWithDefaults() *ReportContactSevQueryParameterFilter`

NewReportContactSevQueryParameterFilterWithDefaults instantiates a new ReportContactSevQueryParameterFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZip

`func (o *ReportContactSevQueryParameterFilter) GetZip() int32`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *ReportContactSevQueryParameterFilter) GetZipOk() (*int32, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *ReportContactSevQueryParameterFilter) SetZip(v int32)`

SetZip sets Zip field to given value.

### HasZip

`func (o *ReportContactSevQueryParameterFilter) HasZip() bool`

HasZip returns a boolean if a field has been set.

### GetCity

`func (o *ReportContactSevQueryParameterFilter) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ReportContactSevQueryParameterFilter) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ReportContactSevQueryParameterFilter) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *ReportContactSevQueryParameterFilter) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *ReportContactSevQueryParameterFilter) GetCountry() ReportContactSevQueryParameterFilterCountry`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ReportContactSevQueryParameterFilter) GetCountryOk() (*ReportContactSevQueryParameterFilterCountry, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ReportContactSevQueryParameterFilter) SetCountry(v ReportContactSevQueryParameterFilterCountry)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ReportContactSevQueryParameterFilter) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetDepth

`func (o *ReportContactSevQueryParameterFilter) GetDepth() bool`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *ReportContactSevQueryParameterFilter) GetDepthOk() (*bool, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *ReportContactSevQueryParameterFilter) SetDepth(v bool)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *ReportContactSevQueryParameterFilter) HasDepth() bool`

HasDepth returns a boolean if a field has been set.

### GetOnlyPeople

`func (o *ReportContactSevQueryParameterFilter) GetOnlyPeople() bool`

GetOnlyPeople returns the OnlyPeople field if non-nil, zero value otherwise.

### GetOnlyPeopleOk

`func (o *ReportContactSevQueryParameterFilter) GetOnlyPeopleOk() (*bool, bool)`

GetOnlyPeopleOk returns a tuple with the OnlyPeople field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyPeople

`func (o *ReportContactSevQueryParameterFilter) SetOnlyPeople(v bool)`

SetOnlyPeople sets OnlyPeople field to given value.

### HasOnlyPeople

`func (o *ReportContactSevQueryParameterFilter) HasOnlyPeople() bool`

HasOnlyPeople returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


