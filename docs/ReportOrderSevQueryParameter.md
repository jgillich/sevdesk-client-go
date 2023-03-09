# ReportOrderSevQueryParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** | Limit export | [optional] 
**ModelName** | **interface{}** | Model name which is exported | 
**ObjectName** | **interface{}** | SevQuery object name | 
**Filter** | Pointer to [**ReportOrderSevQueryParameterFilter**](ReportOrderSevQueryParameterFilter.md) |  | [optional] 

## Methods

### NewReportOrderSevQueryParameter

`func NewReportOrderSevQueryParameter(modelName interface{}, objectName interface{}, ) *ReportOrderSevQueryParameter`

NewReportOrderSevQueryParameter instantiates a new ReportOrderSevQueryParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportOrderSevQueryParameterWithDefaults

`func NewReportOrderSevQueryParameterWithDefaults() *ReportOrderSevQueryParameter`

NewReportOrderSevQueryParameterWithDefaults instantiates a new ReportOrderSevQueryParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *ReportOrderSevQueryParameter) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ReportOrderSevQueryParameter) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ReportOrderSevQueryParameter) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ReportOrderSevQueryParameter) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetModelName

`func (o *ReportOrderSevQueryParameter) GetModelName() interface{}`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *ReportOrderSevQueryParameter) GetModelNameOk() (*interface{}, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *ReportOrderSevQueryParameter) SetModelName(v interface{})`

SetModelName sets ModelName field to given value.


### SetModelNameNil

`func (o *ReportOrderSevQueryParameter) SetModelNameNil(b bool)`

 SetModelNameNil sets the value for ModelName to be an explicit nil

### UnsetModelName
`func (o *ReportOrderSevQueryParameter) UnsetModelName()`

UnsetModelName ensures that no value is present for ModelName, not even an explicit nil
### GetObjectName

`func (o *ReportOrderSevQueryParameter) GetObjectName() interface{}`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ReportOrderSevQueryParameter) GetObjectNameOk() (*interface{}, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ReportOrderSevQueryParameter) SetObjectName(v interface{})`

SetObjectName sets ObjectName field to given value.


### SetObjectNameNil

`func (o *ReportOrderSevQueryParameter) SetObjectNameNil(b bool)`

 SetObjectNameNil sets the value for ObjectName to be an explicit nil

### UnsetObjectName
`func (o *ReportOrderSevQueryParameter) UnsetObjectName()`

UnsetObjectName ensures that no value is present for ObjectName, not even an explicit nil
### GetFilter

`func (o *ReportOrderSevQueryParameter) GetFilter() ReportOrderSevQueryParameterFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ReportOrderSevQueryParameter) GetFilterOk() (*ReportOrderSevQueryParameterFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ReportOrderSevQueryParameter) SetFilter(v ReportOrderSevQueryParameterFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ReportOrderSevQueryParameter) HasFilter() bool`

HasFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


