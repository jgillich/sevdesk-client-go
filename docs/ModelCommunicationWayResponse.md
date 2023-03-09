# ModelCommunicationWayResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The communication way id | [optional] [readonly] 
**ObjectName** | Pointer to **string** | The communication way object name | [optional] [readonly] 
**Create** | Pointer to **time.Time** | Date of communication way creation | [optional] [readonly] 
**Update** | Pointer to **time.Time** | Date of last communication way update | [optional] [readonly] 
**Contact** | Pointer to [**ModelCommunicationWayResponseContact**](ModelCommunicationWayResponseContact.md) |  | [optional] 
**Type** | Pointer to **string** | Type of the communication way | [optional] [readonly] 
**Value** | Pointer to **string** | The value of the communication way.&lt;br&gt;       For example the phone number, e-mail address or website. | [optional] [readonly] 
**Key** | Pointer to [**ModelCommunicationWayResponseKey**](ModelCommunicationWayResponseKey.md) |  | [optional] 
**Main** | Pointer to **string** | Defines whether the communication way is the main communication way for the contact. | [optional] [readonly] 
**SevClient** | Pointer to [**ModelCommunicationWayResponseSevClient**](ModelCommunicationWayResponseSevClient.md) |  | [optional] 

## Methods

### NewModelCommunicationWayResponse

`func NewModelCommunicationWayResponse() *ModelCommunicationWayResponse`

NewModelCommunicationWayResponse instantiates a new ModelCommunicationWayResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelCommunicationWayResponseWithDefaults

`func NewModelCommunicationWayResponseWithDefaults() *ModelCommunicationWayResponse`

NewModelCommunicationWayResponseWithDefaults instantiates a new ModelCommunicationWayResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelCommunicationWayResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelCommunicationWayResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelCommunicationWayResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelCommunicationWayResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectName

`func (o *ModelCommunicationWayResponse) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ModelCommunicationWayResponse) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ModelCommunicationWayResponse) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ModelCommunicationWayResponse) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetCreate

`func (o *ModelCommunicationWayResponse) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ModelCommunicationWayResponse) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ModelCommunicationWayResponse) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ModelCommunicationWayResponse) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUpdate

`func (o *ModelCommunicationWayResponse) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ModelCommunicationWayResponse) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ModelCommunicationWayResponse) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ModelCommunicationWayResponse) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetContact

`func (o *ModelCommunicationWayResponse) GetContact() ModelCommunicationWayResponseContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ModelCommunicationWayResponse) GetContactOk() (*ModelCommunicationWayResponseContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ModelCommunicationWayResponse) SetContact(v ModelCommunicationWayResponseContact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *ModelCommunicationWayResponse) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetType

`func (o *ModelCommunicationWayResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelCommunicationWayResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelCommunicationWayResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ModelCommunicationWayResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *ModelCommunicationWayResponse) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ModelCommunicationWayResponse) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ModelCommunicationWayResponse) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ModelCommunicationWayResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetKey

`func (o *ModelCommunicationWayResponse) GetKey() ModelCommunicationWayResponseKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ModelCommunicationWayResponse) GetKeyOk() (*ModelCommunicationWayResponseKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ModelCommunicationWayResponse) SetKey(v ModelCommunicationWayResponseKey)`

SetKey sets Key field to given value.

### HasKey

`func (o *ModelCommunicationWayResponse) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMain

`func (o *ModelCommunicationWayResponse) GetMain() string`

GetMain returns the Main field if non-nil, zero value otherwise.

### GetMainOk

`func (o *ModelCommunicationWayResponse) GetMainOk() (*string, bool)`

GetMainOk returns a tuple with the Main field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMain

`func (o *ModelCommunicationWayResponse) SetMain(v string)`

SetMain sets Main field to given value.

### HasMain

`func (o *ModelCommunicationWayResponse) HasMain() bool`

HasMain returns a boolean if a field has been set.

### GetSevClient

`func (o *ModelCommunicationWayResponse) GetSevClient() ModelCommunicationWayResponseSevClient`

GetSevClient returns the SevClient field if non-nil, zero value otherwise.

### GetSevClientOk

`func (o *ModelCommunicationWayResponse) GetSevClientOk() (*ModelCommunicationWayResponseSevClient, bool)`

GetSevClientOk returns a tuple with the SevClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSevClient

`func (o *ModelCommunicationWayResponse) SetSevClient(v ModelCommunicationWayResponseSevClient)`

SetSevClient sets SevClient field to given value.

### HasSevClient

`func (o *ModelCommunicationWayResponse) HasSevClient() bool`

HasSevClient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


