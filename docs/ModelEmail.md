# ModelEmail

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

### NewModelEmail

`func NewModelEmail(from string, to string, subject string, ) *ModelEmail`

NewModelEmail instantiates a new ModelEmail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelEmailWithDefaults

`func NewModelEmailWithDefaults() *ModelEmail`

NewModelEmailWithDefaults instantiates a new ModelEmail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelEmail) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelEmail) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelEmail) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelEmail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectName

`func (o *ModelEmail) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ModelEmail) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ModelEmail) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ModelEmail) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetCreate

`func (o *ModelEmail) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ModelEmail) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ModelEmail) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ModelEmail) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUpdate

`func (o *ModelEmail) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ModelEmail) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ModelEmail) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ModelEmail) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetObject

`func (o *ModelEmail) GetObject() ModelEmailOrderObject`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ModelEmail) GetObjectOk() (*ModelEmailOrderObject, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ModelEmail) SetObject(v ModelEmailOrderObject)`

SetObject sets Object field to given value.

### HasObject

`func (o *ModelEmail) HasObject() bool`

HasObject returns a boolean if a field has been set.

### SetObjectNil

`func (o *ModelEmail) SetObjectNil(b bool)`

 SetObjectNil sets the value for Object to be an explicit nil

### UnsetObject
`func (o *ModelEmail) UnsetObject()`

UnsetObject ensures that no value is present for Object, not even an explicit nil
### GetFrom

`func (o *ModelEmail) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ModelEmail) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ModelEmail) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTo

`func (o *ModelEmail) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ModelEmail) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ModelEmail) SetTo(v string)`

SetTo sets To field to given value.


### GetSubject

`func (o *ModelEmail) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ModelEmail) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ModelEmail) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetText

`func (o *ModelEmail) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ModelEmail) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ModelEmail) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ModelEmail) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *ModelEmail) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *ModelEmail) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetSevClient

`func (o *ModelEmail) GetSevClient() ModelEmailOrderSevClient`

GetSevClient returns the SevClient field if non-nil, zero value otherwise.

### GetSevClientOk

`func (o *ModelEmail) GetSevClientOk() (*ModelEmailOrderSevClient, bool)`

GetSevClientOk returns a tuple with the SevClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSevClient

`func (o *ModelEmail) SetSevClient(v ModelEmailOrderSevClient)`

SetSevClient sets SevClient field to given value.

### HasSevClient

`func (o *ModelEmail) HasSevClient() bool`

HasSevClient returns a boolean if a field has been set.

### GetCc

`func (o *ModelEmail) GetCc() string`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *ModelEmail) GetCcOk() (*string, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *ModelEmail) SetCc(v string)`

SetCc sets Cc field to given value.

### HasCc

`func (o *ModelEmail) HasCc() bool`

HasCc returns a boolean if a field has been set.

### SetCcNil

`func (o *ModelEmail) SetCcNil(b bool)`

 SetCcNil sets the value for Cc to be an explicit nil

### UnsetCc
`func (o *ModelEmail) UnsetCc()`

UnsetCc ensures that no value is present for Cc, not even an explicit nil
### GetBcc

`func (o *ModelEmail) GetBcc() string`

GetBcc returns the Bcc field if non-nil, zero value otherwise.

### GetBccOk

`func (o *ModelEmail) GetBccOk() (*string, bool)`

GetBccOk returns a tuple with the Bcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcc

`func (o *ModelEmail) SetBcc(v string)`

SetBcc sets Bcc field to given value.

### HasBcc

`func (o *ModelEmail) HasBcc() bool`

HasBcc returns a boolean if a field has been set.

### SetBccNil

`func (o *ModelEmail) SetBccNil(b bool)`

 SetBccNil sets the value for Bcc to be an explicit nil

### UnsetBcc
`func (o *ModelEmail) UnsetBcc()`

UnsetBcc ensures that no value is present for Bcc, not even an explicit nil
### GetArrived

`func (o *ModelEmail) GetArrived() time.Time`

GetArrived returns the Arrived field if non-nil, zero value otherwise.

### GetArrivedOk

`func (o *ModelEmail) GetArrivedOk() (*time.Time, bool)`

GetArrivedOk returns a tuple with the Arrived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrived

`func (o *ModelEmail) SetArrived(v time.Time)`

SetArrived sets Arrived field to given value.

### HasArrived

`func (o *ModelEmail) HasArrived() bool`

HasArrived returns a boolean if a field has been set.

### SetArrivedNil

`func (o *ModelEmail) SetArrivedNil(b bool)`

 SetArrivedNil sets the value for Arrived to be an explicit nil

### UnsetArrived
`func (o *ModelEmail) UnsetArrived()`

UnsetArrived ensures that no value is present for Arrived, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


