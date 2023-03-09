# ExportContactSevQueryParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** | Limit export | [optional] 
**ModelName** | **interface{}** | Model name, which is &#39;Contact&#39; | 
**ObjectName** | **interface{}** | Model name, which is &#39;SevQuery&#39; | 
**Filter** | Pointer to [**ReportContactSevQueryParameterFilter**](ReportContactSevQueryParameterFilter.md) |  | [optional] 

## Methods

### NewExportContactSevQueryParameter

`func NewExportContactSevQueryParameter(modelName interface{}, objectName interface{}, ) *ExportContactSevQueryParameter`

NewExportContactSevQueryParameter instantiates a new ExportContactSevQueryParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportContactSevQueryParameterWithDefaults

`func NewExportContactSevQueryParameterWithDefaults() *ExportContactSevQueryParameter`

NewExportContactSevQueryParameterWithDefaults instantiates a new ExportContactSevQueryParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *ExportContactSevQueryParameter) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ExportContactSevQueryParameter) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ExportContactSevQueryParameter) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ExportContactSevQueryParameter) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetModelName

`func (o *ExportContactSevQueryParameter) GetModelName() interface{}`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *ExportContactSevQueryParameter) GetModelNameOk() (*interface{}, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *ExportContactSevQueryParameter) SetModelName(v interface{})`

SetModelName sets ModelName field to given value.


### SetModelNameNil

`func (o *ExportContactSevQueryParameter) SetModelNameNil(b bool)`

 SetModelNameNil sets the value for ModelName to be an explicit nil

### UnsetModelName
`func (o *ExportContactSevQueryParameter) UnsetModelName()`

UnsetModelName ensures that no value is present for ModelName, not even an explicit nil
### GetObjectName

`func (o *ExportContactSevQueryParameter) GetObjectName() interface{}`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ExportContactSevQueryParameter) GetObjectNameOk() (*interface{}, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ExportContactSevQueryParameter) SetObjectName(v interface{})`

SetObjectName sets ObjectName field to given value.


### SetObjectNameNil

`func (o *ExportContactSevQueryParameter) SetObjectNameNil(b bool)`

 SetObjectNameNil sets the value for ObjectName to be an explicit nil

### UnsetObjectName
`func (o *ExportContactSevQueryParameter) UnsetObjectName()`

UnsetObjectName ensures that no value is present for ObjectName, not even an explicit nil
### GetFilter

`func (o *ExportContactSevQueryParameter) GetFilter() ReportContactSevQueryParameterFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ExportContactSevQueryParameter) GetFilterOk() (*ReportContactSevQueryParameterFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ExportContactSevQueryParameter) SetFilter(v ReportContactSevQueryParameterFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ExportContactSevQueryParameter) HasFilter() bool`

HasFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


