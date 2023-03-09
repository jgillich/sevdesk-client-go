# ModelTagCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Id of the tag | [optional] [readonly] 
**ObjectName** | Pointer to **string** | Internal object name which is &#39;Tag&#39;. | [optional] [readonly] 
**Create** | Pointer to **time.Time** | Date of tag creation | [optional] [readonly] 
**Tag** | Pointer to [**ModelTagCreateResponseTag**](ModelTagCreateResponseTag.md) |  | [optional] 
**Object** | Pointer to [**ModelTagCreateResponseObject**](ModelTagCreateResponseObject.md) |  | [optional] 
**SevClient** | Pointer to [**ModelTagCreateResponseSevClient**](ModelTagCreateResponseSevClient.md) |  | [optional] 

## Methods

### NewModelTagCreateResponse

`func NewModelTagCreateResponse() *ModelTagCreateResponse`

NewModelTagCreateResponse instantiates a new ModelTagCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelTagCreateResponseWithDefaults

`func NewModelTagCreateResponseWithDefaults() *ModelTagCreateResponse`

NewModelTagCreateResponseWithDefaults instantiates a new ModelTagCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelTagCreateResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelTagCreateResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelTagCreateResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelTagCreateResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectName

`func (o *ModelTagCreateResponse) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ModelTagCreateResponse) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ModelTagCreateResponse) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ModelTagCreateResponse) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetCreate

`func (o *ModelTagCreateResponse) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ModelTagCreateResponse) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ModelTagCreateResponse) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ModelTagCreateResponse) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetTag

`func (o *ModelTagCreateResponse) GetTag() ModelTagCreateResponseTag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ModelTagCreateResponse) GetTagOk() (*ModelTagCreateResponseTag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ModelTagCreateResponse) SetTag(v ModelTagCreateResponseTag)`

SetTag sets Tag field to given value.

### HasTag

`func (o *ModelTagCreateResponse) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetObject

`func (o *ModelTagCreateResponse) GetObject() ModelTagCreateResponseObject`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ModelTagCreateResponse) GetObjectOk() (*ModelTagCreateResponseObject, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ModelTagCreateResponse) SetObject(v ModelTagCreateResponseObject)`

SetObject sets Object field to given value.

### HasObject

`func (o *ModelTagCreateResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetSevClient

`func (o *ModelTagCreateResponse) GetSevClient() ModelTagCreateResponseSevClient`

GetSevClient returns the SevClient field if non-nil, zero value otherwise.

### GetSevClientOk

`func (o *ModelTagCreateResponse) GetSevClientOk() (*ModelTagCreateResponseSevClient, bool)`

GetSevClientOk returns a tuple with the SevClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSevClient

`func (o *ModelTagCreateResponse) SetSevClient(v ModelTagCreateResponseSevClient)`

SetSevClient sets SevClient field to given value.

### HasSevClient

`func (o *ModelTagCreateResponse) HasSevClient() bool`

HasSevClient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


