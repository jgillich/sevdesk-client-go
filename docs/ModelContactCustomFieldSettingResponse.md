# ModelContactCustomFieldSettingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Id of the contact field | [optional] [readonly] 
**ObjectName** | Pointer to **string** | Internal object name which is &#39;ContactCustomFieldSetting&#39;. | [optional] [readonly] 
**Create** | Pointer to **time.Time** | Date of contact field creation | [optional] [readonly] 
**Update** | Pointer to **time.Time** | Date of contact field updated | [optional] [readonly] 
**SevClient** | Pointer to [**ModelContactCustomFieldSettingResponseSevClient**](ModelContactCustomFieldSettingResponseSevClient.md) |  | [optional] 
**Name** | Pointer to **string** | name of the contact fields | [optional] [readonly] 
**Identifier** | Pointer to **string** | Unique identifier for the contact field | [optional] [readonly] 
**Description** | Pointer to **string** | The description of the contact field | [optional] [readonly] 

## Methods

### NewModelContactCustomFieldSettingResponse

`func NewModelContactCustomFieldSettingResponse() *ModelContactCustomFieldSettingResponse`

NewModelContactCustomFieldSettingResponse instantiates a new ModelContactCustomFieldSettingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelContactCustomFieldSettingResponseWithDefaults

`func NewModelContactCustomFieldSettingResponseWithDefaults() *ModelContactCustomFieldSettingResponse`

NewModelContactCustomFieldSettingResponseWithDefaults instantiates a new ModelContactCustomFieldSettingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelContactCustomFieldSettingResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelContactCustomFieldSettingResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelContactCustomFieldSettingResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelContactCustomFieldSettingResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectName

`func (o *ModelContactCustomFieldSettingResponse) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ModelContactCustomFieldSettingResponse) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ModelContactCustomFieldSettingResponse) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ModelContactCustomFieldSettingResponse) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetCreate

`func (o *ModelContactCustomFieldSettingResponse) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ModelContactCustomFieldSettingResponse) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ModelContactCustomFieldSettingResponse) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ModelContactCustomFieldSettingResponse) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUpdate

`func (o *ModelContactCustomFieldSettingResponse) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ModelContactCustomFieldSettingResponse) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ModelContactCustomFieldSettingResponse) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ModelContactCustomFieldSettingResponse) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetSevClient

`func (o *ModelContactCustomFieldSettingResponse) GetSevClient() ModelContactCustomFieldSettingResponseSevClient`

GetSevClient returns the SevClient field if non-nil, zero value otherwise.

### GetSevClientOk

`func (o *ModelContactCustomFieldSettingResponse) GetSevClientOk() (*ModelContactCustomFieldSettingResponseSevClient, bool)`

GetSevClientOk returns a tuple with the SevClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSevClient

`func (o *ModelContactCustomFieldSettingResponse) SetSevClient(v ModelContactCustomFieldSettingResponseSevClient)`

SetSevClient sets SevClient field to given value.

### HasSevClient

`func (o *ModelContactCustomFieldSettingResponse) HasSevClient() bool`

HasSevClient returns a boolean if a field has been set.

### GetName

`func (o *ModelContactCustomFieldSettingResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelContactCustomFieldSettingResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelContactCustomFieldSettingResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelContactCustomFieldSettingResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIdentifier

`func (o *ModelContactCustomFieldSettingResponse) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ModelContactCustomFieldSettingResponse) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ModelContactCustomFieldSettingResponse) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *ModelContactCustomFieldSettingResponse) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetDescription

`func (o *ModelContactCustomFieldSettingResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelContactCustomFieldSettingResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelContactCustomFieldSettingResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelContactCustomFieldSettingResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


