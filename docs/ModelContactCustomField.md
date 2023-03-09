# ModelContactCustomField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | [**ModelContactCustomFieldContact**](ModelContactCustomFieldContact.md) |  | 
**ContactCustomFieldSetting** | [**ModelContactCustomFieldContactCustomFieldSetting**](ModelContactCustomFieldContactCustomFieldSetting.md) |  | 
**Value** | **string** | The value of the contact field | 
**ObjectName** | **string** | Internal object name which is &#39;ContactCustomField&#39;. | 

## Methods

### NewModelContactCustomField

`func NewModelContactCustomField(contact ModelContactCustomFieldContact, contactCustomFieldSetting ModelContactCustomFieldContactCustomFieldSetting, value string, objectName string, ) *ModelContactCustomField`

NewModelContactCustomField instantiates a new ModelContactCustomField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelContactCustomFieldWithDefaults

`func NewModelContactCustomFieldWithDefaults() *ModelContactCustomField`

NewModelContactCustomFieldWithDefaults instantiates a new ModelContactCustomField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContact

`func (o *ModelContactCustomField) GetContact() ModelContactCustomFieldContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ModelContactCustomField) GetContactOk() (*ModelContactCustomFieldContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ModelContactCustomField) SetContact(v ModelContactCustomFieldContact)`

SetContact sets Contact field to given value.


### GetContactCustomFieldSetting

`func (o *ModelContactCustomField) GetContactCustomFieldSetting() ModelContactCustomFieldContactCustomFieldSetting`

GetContactCustomFieldSetting returns the ContactCustomFieldSetting field if non-nil, zero value otherwise.

### GetContactCustomFieldSettingOk

`func (o *ModelContactCustomField) GetContactCustomFieldSettingOk() (*ModelContactCustomFieldContactCustomFieldSetting, bool)`

GetContactCustomFieldSettingOk returns a tuple with the ContactCustomFieldSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactCustomFieldSetting

`func (o *ModelContactCustomField) SetContactCustomFieldSetting(v ModelContactCustomFieldContactCustomFieldSetting)`

SetContactCustomFieldSetting sets ContactCustomFieldSetting field to given value.


### GetValue

`func (o *ModelContactCustomField) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ModelContactCustomField) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ModelContactCustomField) SetValue(v string)`

SetValue sets Value field to given value.


### GetObjectName

`func (o *ModelContactCustomField) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ModelContactCustomField) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ModelContactCustomField) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


