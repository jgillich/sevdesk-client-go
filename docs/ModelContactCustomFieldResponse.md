# ModelContactCustomFieldResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | id of the contact field | [optional] 
**ObjectName** | Pointer to **string** | Internal object name which is &#39;ContactCustomField&#39;. | [optional] 
**Create** | Pointer to **time.Time** | Date of contact field creation | [optional] 
**Update** | Pointer to **time.Time** | Date of contact field update | [optional] 
**SevClient** | Pointer to [**ModelContactCustomFieldResponseSevClient**](ModelContactCustomFieldResponseSevClient.md) |  | [optional] 
**Contact** | Pointer to [**ModelContactCustomFieldResponseContact**](ModelContactCustomFieldResponseContact.md) |  | [optional] 
**ContactCustomFieldSetting** | Pointer to [**ModelContactCustomFieldResponseContactCustomFieldSetting**](ModelContactCustomFieldResponseContactCustomFieldSetting.md) |  | [optional] 
**Value** | Pointer to **string** | The value of the contact field | [optional] 

## Methods

### NewModelContactCustomFieldResponse

`func NewModelContactCustomFieldResponse() *ModelContactCustomFieldResponse`

NewModelContactCustomFieldResponse instantiates a new ModelContactCustomFieldResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelContactCustomFieldResponseWithDefaults

`func NewModelContactCustomFieldResponseWithDefaults() *ModelContactCustomFieldResponse`

NewModelContactCustomFieldResponseWithDefaults instantiates a new ModelContactCustomFieldResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelContactCustomFieldResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelContactCustomFieldResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelContactCustomFieldResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelContactCustomFieldResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectName

`func (o *ModelContactCustomFieldResponse) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ModelContactCustomFieldResponse) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ModelContactCustomFieldResponse) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ModelContactCustomFieldResponse) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetCreate

`func (o *ModelContactCustomFieldResponse) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ModelContactCustomFieldResponse) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ModelContactCustomFieldResponse) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ModelContactCustomFieldResponse) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUpdate

`func (o *ModelContactCustomFieldResponse) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ModelContactCustomFieldResponse) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ModelContactCustomFieldResponse) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ModelContactCustomFieldResponse) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetSevClient

`func (o *ModelContactCustomFieldResponse) GetSevClient() ModelContactCustomFieldResponseSevClient`

GetSevClient returns the SevClient field if non-nil, zero value otherwise.

### GetSevClientOk

`func (o *ModelContactCustomFieldResponse) GetSevClientOk() (*ModelContactCustomFieldResponseSevClient, bool)`

GetSevClientOk returns a tuple with the SevClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSevClient

`func (o *ModelContactCustomFieldResponse) SetSevClient(v ModelContactCustomFieldResponseSevClient)`

SetSevClient sets SevClient field to given value.

### HasSevClient

`func (o *ModelContactCustomFieldResponse) HasSevClient() bool`

HasSevClient returns a boolean if a field has been set.

### GetContact

`func (o *ModelContactCustomFieldResponse) GetContact() ModelContactCustomFieldResponseContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ModelContactCustomFieldResponse) GetContactOk() (*ModelContactCustomFieldResponseContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ModelContactCustomFieldResponse) SetContact(v ModelContactCustomFieldResponseContact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *ModelContactCustomFieldResponse) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetContactCustomFieldSetting

`func (o *ModelContactCustomFieldResponse) GetContactCustomFieldSetting() ModelContactCustomFieldResponseContactCustomFieldSetting`

GetContactCustomFieldSetting returns the ContactCustomFieldSetting field if non-nil, zero value otherwise.

### GetContactCustomFieldSettingOk

`func (o *ModelContactCustomFieldResponse) GetContactCustomFieldSettingOk() (*ModelContactCustomFieldResponseContactCustomFieldSetting, bool)`

GetContactCustomFieldSettingOk returns a tuple with the ContactCustomFieldSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactCustomFieldSetting

`func (o *ModelContactCustomFieldResponse) SetContactCustomFieldSetting(v ModelContactCustomFieldResponseContactCustomFieldSetting)`

SetContactCustomFieldSetting sets ContactCustomFieldSetting field to given value.

### HasContactCustomFieldSetting

`func (o *ModelContactCustomFieldResponse) HasContactCustomFieldSetting() bool`

HasContactCustomFieldSetting returns a boolean if a field has been set.

### GetValue

`func (o *ModelContactCustomFieldResponse) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ModelContactCustomFieldResponse) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ModelContactCustomFieldResponse) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ModelContactCustomFieldResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


