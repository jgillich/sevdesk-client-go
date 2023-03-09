# ModelChangeLayoutResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to **string** |  | [optional] 
**Metadaten** | Pointer to [**ModelChangeLayoutResponseMetadaten**](ModelChangeLayoutResponseMetadaten.md) |  | [optional] 

## Methods

### NewModelChangeLayoutResponse

`func NewModelChangeLayoutResponse() *ModelChangeLayoutResponse`

NewModelChangeLayoutResponse instantiates a new ModelChangeLayoutResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelChangeLayoutResponseWithDefaults

`func NewModelChangeLayoutResponseWithDefaults() *ModelChangeLayoutResponse`

NewModelChangeLayoutResponseWithDefaults instantiates a new ModelChangeLayoutResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ModelChangeLayoutResponse) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ModelChangeLayoutResponse) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ModelChangeLayoutResponse) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *ModelChangeLayoutResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetMetadaten

`func (o *ModelChangeLayoutResponse) GetMetadaten() ModelChangeLayoutResponseMetadaten`

GetMetadaten returns the Metadaten field if non-nil, zero value otherwise.

### GetMetadatenOk

`func (o *ModelChangeLayoutResponse) GetMetadatenOk() (*ModelChangeLayoutResponseMetadaten, bool)`

GetMetadatenOk returns a tuple with the Metadaten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadaten

`func (o *ModelChangeLayoutResponse) SetMetadaten(v ModelChangeLayoutResponseMetadaten)`

SetMetadaten sets Metadaten field to given value.

### HasMetadaten

`func (o *ModelChangeLayoutResponse) HasMetadaten() bool`

HasMetadaten returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


