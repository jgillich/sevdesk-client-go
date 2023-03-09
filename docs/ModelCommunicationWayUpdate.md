# ModelCommunicationWayUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | Pointer to [**NullableModelCommunicationWayUpdateContact**](ModelCommunicationWayUpdateContact.md) |  | [optional] 
**Type** | Pointer to **string** | Type of the communication way | [optional] 
**Value** | Pointer to **string** | The value of the communication way.&lt;br&gt;       For example the phone number, e-mail address or website. | [optional] 
**Key** | Pointer to [**NullableModelCommunicationWayUpdateKey**](ModelCommunicationWayUpdateKey.md) |  | [optional] 
**Main** | Pointer to **NullableBool** | Defines whether the communication way is the main communication way for the contact. | [optional] 

## Methods

### NewModelCommunicationWayUpdate

`func NewModelCommunicationWayUpdate() *ModelCommunicationWayUpdate`

NewModelCommunicationWayUpdate instantiates a new ModelCommunicationWayUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelCommunicationWayUpdateWithDefaults

`func NewModelCommunicationWayUpdateWithDefaults() *ModelCommunicationWayUpdate`

NewModelCommunicationWayUpdateWithDefaults instantiates a new ModelCommunicationWayUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContact

`func (o *ModelCommunicationWayUpdate) GetContact() ModelCommunicationWayUpdateContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ModelCommunicationWayUpdate) GetContactOk() (*ModelCommunicationWayUpdateContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ModelCommunicationWayUpdate) SetContact(v ModelCommunicationWayUpdateContact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *ModelCommunicationWayUpdate) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *ModelCommunicationWayUpdate) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *ModelCommunicationWayUpdate) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetType

`func (o *ModelCommunicationWayUpdate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelCommunicationWayUpdate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelCommunicationWayUpdate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ModelCommunicationWayUpdate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *ModelCommunicationWayUpdate) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ModelCommunicationWayUpdate) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ModelCommunicationWayUpdate) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ModelCommunicationWayUpdate) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetKey

`func (o *ModelCommunicationWayUpdate) GetKey() ModelCommunicationWayUpdateKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ModelCommunicationWayUpdate) GetKeyOk() (*ModelCommunicationWayUpdateKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ModelCommunicationWayUpdate) SetKey(v ModelCommunicationWayUpdateKey)`

SetKey sets Key field to given value.

### HasKey

`func (o *ModelCommunicationWayUpdate) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *ModelCommunicationWayUpdate) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *ModelCommunicationWayUpdate) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetMain

`func (o *ModelCommunicationWayUpdate) GetMain() bool`

GetMain returns the Main field if non-nil, zero value otherwise.

### GetMainOk

`func (o *ModelCommunicationWayUpdate) GetMainOk() (*bool, bool)`

GetMainOk returns a tuple with the Main field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMain

`func (o *ModelCommunicationWayUpdate) SetMain(v bool)`

SetMain sets Main field to given value.

### HasMain

`func (o *ModelCommunicationWayUpdate) HasMain() bool`

HasMain returns a boolean if a field has been set.

### SetMainNil

`func (o *ModelCommunicationWayUpdate) SetMainNil(b bool)`

 SetMainNil sets the value for Main to be an explicit nil

### UnsetMain
`func (o *ModelCommunicationWayUpdate) UnsetMain()`

UnsetMain ensures that no value is present for Main, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


