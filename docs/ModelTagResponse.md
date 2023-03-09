# ModelTagResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Id of the tag | [optional] [readonly] 
**ObjectName** | Pointer to **string** | Internal object name which is &#39;Tag&#39;. | [optional] [readonly] 
**Create** | Pointer to **time.Time** | Date of tag creation | [optional] [readonly] 
**Name** | Pointer to **string** | name of the tag | [optional] [readonly] 
**SevClient** | Pointer to [**ModelTagCreateResponseSevClient**](ModelTagCreateResponseSevClient.md) |  | [optional] 

## Methods

### NewModelTagResponse

`func NewModelTagResponse() *ModelTagResponse`

NewModelTagResponse instantiates a new ModelTagResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelTagResponseWithDefaults

`func NewModelTagResponseWithDefaults() *ModelTagResponse`

NewModelTagResponseWithDefaults instantiates a new ModelTagResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelTagResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelTagResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelTagResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelTagResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectName

`func (o *ModelTagResponse) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ModelTagResponse) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ModelTagResponse) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ModelTagResponse) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetCreate

`func (o *ModelTagResponse) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ModelTagResponse) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ModelTagResponse) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ModelTagResponse) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetName

`func (o *ModelTagResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelTagResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelTagResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelTagResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSevClient

`func (o *ModelTagResponse) GetSevClient() ModelTagCreateResponseSevClient`

GetSevClient returns the SevClient field if non-nil, zero value otherwise.

### GetSevClientOk

`func (o *ModelTagResponse) GetSevClientOk() (*ModelTagCreateResponseSevClient, bool)`

GetSevClientOk returns a tuple with the SevClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSevClient

`func (o *ModelTagResponse) SetSevClient(v ModelTagCreateResponseSevClient)`

SetSevClient sets SevClient field to given value.

### HasSevClient

`func (o *ModelTagResponse) HasSevClient() bool`

HasSevClient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


