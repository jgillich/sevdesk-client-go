# ModelEmailOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The email id | [optional] [readonly] 
**ObjectName** | Pointer to **string** | The email object name | [optional] [readonly] 
**Create** | Pointer to **time.Time** | Date of mail creation | [optional] [readonly] 
**Update** | Pointer to **time.Time** | Date of last mail update | [optional] [readonly] 
**Object** | Pointer to [**NullableModelEmailOrderObject**](ModelEmailOrderObject.md) |  | [optional] 
**From** | **string** | The sender of the email | 
**To** | **string** | The recipient of the email | 
**Subject** | **string** | The subject of the email | 
**Text** | Pointer to **NullableString** | The text of the email | [optional] 
**SevClient** | Pointer to [**ModelEmailOrderSevClient**](ModelEmailOrderSevClient.md) |  | [optional] 
**Cc** | Pointer to **NullableString** | A list of mail addresses which are in the cc | [optional] 
**Bcc** | Pointer to **NullableString** | A list of mail addresses which are in the bcc | [optional] 
**Arrived** | Pointer to **NullableTime** | Date the mail arrived | [optional] 

## Methods

### NewModelEmailOrder

`func NewModelEmailOrder(from string, to string, subject string, ) *ModelEmailOrder`

NewModelEmailOrder instantiates a new ModelEmailOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelEmailOrderWithDefaults

`func NewModelEmailOrderWithDefaults() *ModelEmailOrder`

NewModelEmailOrderWithDefaults instantiates a new ModelEmailOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelEmailOrder) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelEmailOrder) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelEmailOrder) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelEmailOrder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectName

`func (o *ModelEmailOrder) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ModelEmailOrder) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ModelEmailOrder) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ModelEmailOrder) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetCreate

`func (o *ModelEmailOrder) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ModelEmailOrder) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ModelEmailOrder) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ModelEmailOrder) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUpdate

`func (o *ModelEmailOrder) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ModelEmailOrder) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ModelEmailOrder) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ModelEmailOrder) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetObject

`func (o *ModelEmailOrder) GetObject() ModelEmailOrderObject`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ModelEmailOrder) GetObjectOk() (*ModelEmailOrderObject, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ModelEmailOrder) SetObject(v ModelEmailOrderObject)`

SetObject sets Object field to given value.

### HasObject

`func (o *ModelEmailOrder) HasObject() bool`

HasObject returns a boolean if a field has been set.

### SetObjectNil

`func (o *ModelEmailOrder) SetObjectNil(b bool)`

 SetObjectNil sets the value for Object to be an explicit nil

### UnsetObject
`func (o *ModelEmailOrder) UnsetObject()`

UnsetObject ensures that no value is present for Object, not even an explicit nil
### GetFrom

`func (o *ModelEmailOrder) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ModelEmailOrder) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ModelEmailOrder) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTo

`func (o *ModelEmailOrder) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ModelEmailOrder) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ModelEmailOrder) SetTo(v string)`

SetTo sets To field to given value.


### GetSubject

`func (o *ModelEmailOrder) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ModelEmailOrder) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ModelEmailOrder) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetText

`func (o *ModelEmailOrder) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ModelEmailOrder) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ModelEmailOrder) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ModelEmailOrder) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *ModelEmailOrder) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *ModelEmailOrder) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetSevClient

`func (o *ModelEmailOrder) GetSevClient() ModelEmailOrderSevClient`

GetSevClient returns the SevClient field if non-nil, zero value otherwise.

### GetSevClientOk

`func (o *ModelEmailOrder) GetSevClientOk() (*ModelEmailOrderSevClient, bool)`

GetSevClientOk returns a tuple with the SevClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSevClient

`func (o *ModelEmailOrder) SetSevClient(v ModelEmailOrderSevClient)`

SetSevClient sets SevClient field to given value.

### HasSevClient

`func (o *ModelEmailOrder) HasSevClient() bool`

HasSevClient returns a boolean if a field has been set.

### GetCc

`func (o *ModelEmailOrder) GetCc() string`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *ModelEmailOrder) GetCcOk() (*string, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *ModelEmailOrder) SetCc(v string)`

SetCc sets Cc field to given value.

### HasCc

`func (o *ModelEmailOrder) HasCc() bool`

HasCc returns a boolean if a field has been set.

### SetCcNil

`func (o *ModelEmailOrder) SetCcNil(b bool)`

 SetCcNil sets the value for Cc to be an explicit nil

### UnsetCc
`func (o *ModelEmailOrder) UnsetCc()`

UnsetCc ensures that no value is present for Cc, not even an explicit nil
### GetBcc

`func (o *ModelEmailOrder) GetBcc() string`

GetBcc returns the Bcc field if non-nil, zero value otherwise.

### GetBccOk

`func (o *ModelEmailOrder) GetBccOk() (*string, bool)`

GetBccOk returns a tuple with the Bcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcc

`func (o *ModelEmailOrder) SetBcc(v string)`

SetBcc sets Bcc field to given value.

### HasBcc

`func (o *ModelEmailOrder) HasBcc() bool`

HasBcc returns a boolean if a field has been set.

### SetBccNil

`func (o *ModelEmailOrder) SetBccNil(b bool)`

 SetBccNil sets the value for Bcc to be an explicit nil

### UnsetBcc
`func (o *ModelEmailOrder) UnsetBcc()`

UnsetBcc ensures that no value is present for Bcc, not even an explicit nil
### GetArrived

`func (o *ModelEmailOrder) GetArrived() time.Time`

GetArrived returns the Arrived field if non-nil, zero value otherwise.

### GetArrivedOk

`func (o *ModelEmailOrder) GetArrivedOk() (*time.Time, bool)`

GetArrivedOk returns a tuple with the Arrived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrived

`func (o *ModelEmailOrder) SetArrived(v time.Time)`

SetArrived sets Arrived field to given value.

### HasArrived

`func (o *ModelEmailOrder) HasArrived() bool`

HasArrived returns a boolean if a field has been set.

### SetArrivedNil

`func (o *ModelEmailOrder) SetArrivedNil(b bool)`

 SetArrivedNil sets the value for Arrived to be an explicit nil

### UnsetArrived
`func (o *ModelEmailOrder) UnsetArrived()`

UnsetArrived ensures that no value is present for Arrived, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


