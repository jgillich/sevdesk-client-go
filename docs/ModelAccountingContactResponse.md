# ModelAccountingContactResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The accounting contact id | [optional] [readonly] 
**ObjectName** | Pointer to **string** | The accounting contact object name | [optional] [readonly] 
**Create** | Pointer to **time.Time** | Date of accounting contact creation | [optional] [readonly] 
**Update** | Pointer to **time.Time** | Date of last accounting contact update | [optional] [readonly] 
**Contact** | Pointer to [**ModelAccountingContactResponseContact**](ModelAccountingContactResponseContact.md) |  | [optional] 
**SevClient** | Pointer to [**ModelAccountingContactResponseSevClient**](ModelAccountingContactResponseSevClient.md) |  | [optional] 
**DebitorNumber** | Pointer to **string** | Debitor number of the accounting contact. | [optional] [readonly] 
**CreditorNumber** | Pointer to **string** | Creditor number of the accounting contact. | [optional] [readonly] 

## Methods

### NewModelAccountingContactResponse

`func NewModelAccountingContactResponse() *ModelAccountingContactResponse`

NewModelAccountingContactResponse instantiates a new ModelAccountingContactResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAccountingContactResponseWithDefaults

`func NewModelAccountingContactResponseWithDefaults() *ModelAccountingContactResponse`

NewModelAccountingContactResponseWithDefaults instantiates a new ModelAccountingContactResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelAccountingContactResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelAccountingContactResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelAccountingContactResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelAccountingContactResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectName

`func (o *ModelAccountingContactResponse) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ModelAccountingContactResponse) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ModelAccountingContactResponse) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ModelAccountingContactResponse) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetCreate

`func (o *ModelAccountingContactResponse) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ModelAccountingContactResponse) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ModelAccountingContactResponse) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ModelAccountingContactResponse) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUpdate

`func (o *ModelAccountingContactResponse) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ModelAccountingContactResponse) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ModelAccountingContactResponse) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ModelAccountingContactResponse) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetContact

`func (o *ModelAccountingContactResponse) GetContact() ModelAccountingContactResponseContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ModelAccountingContactResponse) GetContactOk() (*ModelAccountingContactResponseContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ModelAccountingContactResponse) SetContact(v ModelAccountingContactResponseContact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *ModelAccountingContactResponse) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetSevClient

`func (o *ModelAccountingContactResponse) GetSevClient() ModelAccountingContactResponseSevClient`

GetSevClient returns the SevClient field if non-nil, zero value otherwise.

### GetSevClientOk

`func (o *ModelAccountingContactResponse) GetSevClientOk() (*ModelAccountingContactResponseSevClient, bool)`

GetSevClientOk returns a tuple with the SevClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSevClient

`func (o *ModelAccountingContactResponse) SetSevClient(v ModelAccountingContactResponseSevClient)`

SetSevClient sets SevClient field to given value.

### HasSevClient

`func (o *ModelAccountingContactResponse) HasSevClient() bool`

HasSevClient returns a boolean if a field has been set.

### GetDebitorNumber

`func (o *ModelAccountingContactResponse) GetDebitorNumber() string`

GetDebitorNumber returns the DebitorNumber field if non-nil, zero value otherwise.

### GetDebitorNumberOk

`func (o *ModelAccountingContactResponse) GetDebitorNumberOk() (*string, bool)`

GetDebitorNumberOk returns a tuple with the DebitorNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitorNumber

`func (o *ModelAccountingContactResponse) SetDebitorNumber(v string)`

SetDebitorNumber sets DebitorNumber field to given value.

### HasDebitorNumber

`func (o *ModelAccountingContactResponse) HasDebitorNumber() bool`

HasDebitorNumber returns a boolean if a field has been set.

### GetCreditorNumber

`func (o *ModelAccountingContactResponse) GetCreditorNumber() string`

GetCreditorNumber returns the CreditorNumber field if non-nil, zero value otherwise.

### GetCreditorNumberOk

`func (o *ModelAccountingContactResponse) GetCreditorNumberOk() (*string, bool)`

GetCreditorNumberOk returns a tuple with the CreditorNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditorNumber

`func (o *ModelAccountingContactResponse) SetCreditorNumber(v string)`

SetCreditorNumber sets CreditorNumber field to given value.

### HasCreditorNumber

`func (o *ModelAccountingContactResponse) HasCreditorNumber() bool`

HasCreditorNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


