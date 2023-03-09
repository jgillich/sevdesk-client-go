# ModelContactCustomFieldUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | Pointer to [**ModelContactCustomFieldContact**](ModelContactCustomFieldContact.md) |  | [optional] 
**ContactCustomFieldSetting** | Pointer to [**ModelContactCustomFieldContactCustomFieldSetting**](ModelContactCustomFieldContactCustomFieldSetting.md) |  | [optional] 
**Value** | Pointer to **string** | The value of the contact field | [optional] 
**ObjectName** | Pointer to **string** | Internal object name which is &#39;ContactCustomField&#39;. | [optional] 

## Methods

### NewModelContactCustomFieldUpdate

`func NewModelContactCustomFieldUpdate() *ModelContactCustomFieldUpdate`

NewModelContactCustomFieldUpdate instantiates a new ModelContactCustomFieldUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelContactCustomFieldUpdateWithDefaults

`func NewModelContactCustomFieldUpdateWithDefaults() *ModelContactCustomFieldUpdate`

NewModelContactCustomFieldUpdateWithDefaults instantiates a new ModelContactCustomFieldUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContact

`func (o *ModelContactCustomFieldUpdate) GetContact() ModelContactCustomFieldContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ModelContactCustomFieldUpdate) GetContactOk() (*ModelContactCustomFieldContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ModelContactCustomFieldUpdate) SetContact(v ModelContactCustomFieldContact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *ModelContactCustomFieldUpdate) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetContactCustomFieldSetting

`func (o *ModelContactCustomFieldUpdate) GetContactCustomFieldSetting() ModelContactCustomFieldContactCustomFieldSetting`

GetContactCustomFieldSetting returns the ContactCustomFieldSetting field if non-nil, zero value otherwise.

### GetContactCustomFieldSettingOk

`func (o *ModelContactCustomFieldUpdate) GetContactCustomFieldSettingOk() (*ModelContactCustomFieldContactCustomFieldSetting, bool)`

GetContactCustomFieldSettingOk returns a tuple with the ContactCustomFieldSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactCustomFieldSetting

`func (o *ModelContactCustomFieldUpdate) SetContactCustomFieldSetting(v ModelContactCustomFieldContactCustomFieldSetting)`

SetContactCustomFieldSetting sets ContactCustomFieldSetting field to given value.

### HasContactCustomFieldSetting

`func (o *ModelContactCustomFieldUpdate) HasContactCustomFieldSetting() bool`

HasContactCustomFieldSetting returns a boolean if a field has been set.

### GetValue

`func (o *ModelContactCustomFieldUpdate) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ModelContactCustomFieldUpdate) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ModelContactCustomFieldUpdate) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ModelContactCustomFieldUpdate) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetObjectName

`func (o *ModelContactCustomFieldUpdate) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ModelContactCustomFieldUpdate) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ModelContactCustomFieldUpdate) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ModelContactCustomFieldUpdate) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


