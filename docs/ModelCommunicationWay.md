# ModelCommunicationWay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The communication way id | [optional] [readonly] 
**ObjectName** | Pointer to **string** | The communication way object name | [optional] [readonly] 
**Create** | Pointer to **time.Time** | Date of communication way creation | [optional] [readonly] 
**Update** | Pointer to **time.Time** | Date of last communication way update | [optional] [readonly] 
**Contact** | Pointer to [**ModelCommunicationWayContact**](ModelCommunicationWayContact.md) |  | [optional] 
**Type** | **string** | Type of the communication way | 
**Value** | **string** | The value of the communication way.&lt;br&gt;       For example the phone number, e-mail address or website. | 
**Key** | [**ModelCommunicationWayKey**](ModelCommunicationWayKey.md) |  | 
**Main** | Pointer to **NullableBool** | Defines whether the communication way is the main communication way for the contact. | [optional] 
**SevClient** | Pointer to [**ModelCommunicationWaySevClient**](ModelCommunicationWaySevClient.md) |  | [optional] 

## Methods

### NewModelCommunicationWay

`func NewModelCommunicationWay(type_ string, value string, key ModelCommunicationWayKey, ) *ModelCommunicationWay`

NewModelCommunicationWay instantiates a new ModelCommunicationWay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelCommunicationWayWithDefaults

`func NewModelCommunicationWayWithDefaults() *ModelCommunicationWay`

NewModelCommunicationWayWithDefaults instantiates a new ModelCommunicationWay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelCommunicationWay) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelCommunicationWay) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelCommunicationWay) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelCommunicationWay) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectName

`func (o *ModelCommunicationWay) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ModelCommunicationWay) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ModelCommunicationWay) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ModelCommunicationWay) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetCreate

`func (o *ModelCommunicationWay) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ModelCommunicationWay) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ModelCommunicationWay) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ModelCommunicationWay) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUpdate

`func (o *ModelCommunicationWay) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ModelCommunicationWay) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ModelCommunicationWay) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ModelCommunicationWay) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetContact

`func (o *ModelCommunicationWay) GetContact() ModelCommunicationWayContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ModelCommunicationWay) GetContactOk() (*ModelCommunicationWayContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ModelCommunicationWay) SetContact(v ModelCommunicationWayContact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *ModelCommunicationWay) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetType

`func (o *ModelCommunicationWay) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelCommunicationWay) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelCommunicationWay) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *ModelCommunicationWay) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ModelCommunicationWay) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ModelCommunicationWay) SetValue(v string)`

SetValue sets Value field to given value.


### GetKey

`func (o *ModelCommunicationWay) GetKey() ModelCommunicationWayKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ModelCommunicationWay) GetKeyOk() (*ModelCommunicationWayKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ModelCommunicationWay) SetKey(v ModelCommunicationWayKey)`

SetKey sets Key field to given value.


### GetMain

`func (o *ModelCommunicationWay) GetMain() bool`

GetMain returns the Main field if non-nil, zero value otherwise.

### GetMainOk

`func (o *ModelCommunicationWay) GetMainOk() (*bool, bool)`

GetMainOk returns a tuple with the Main field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMain

`func (o *ModelCommunicationWay) SetMain(v bool)`

SetMain sets Main field to given value.

### HasMain

`func (o *ModelCommunicationWay) HasMain() bool`

HasMain returns a boolean if a field has been set.

### SetMainNil

`func (o *ModelCommunicationWay) SetMainNil(b bool)`

 SetMainNil sets the value for Main to be an explicit nil

### UnsetMain
`func (o *ModelCommunicationWay) UnsetMain()`

UnsetMain ensures that no value is present for Main, not even an explicit nil
### GetSevClient

`func (o *ModelCommunicationWay) GetSevClient() ModelCommunicationWaySevClient`

GetSevClient returns the SevClient field if non-nil, zero value otherwise.

### GetSevClientOk

`func (o *ModelCommunicationWay) GetSevClientOk() (*ModelCommunicationWaySevClient, bool)`

GetSevClientOk returns a tuple with the SevClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSevClient

`func (o *ModelCommunicationWay) SetSevClient(v ModelCommunicationWaySevClient)`

SetSevClient sets SevClient field to given value.

### HasSevClient

`func (o *ModelCommunicationWay) HasSevClient() bool`

HasSevClient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


