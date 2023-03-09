# ModelAccountingContactUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | Pointer to [**NullableModelAccountingContactUpdateContact**](ModelAccountingContactUpdateContact.md) |  | [optional] 
**DebitorNumber** | Pointer to **NullableInt32** | Debitor number of the accounting contact. | [optional] 
**CreditorNumber** | Pointer to **NullableInt32** | Creditor number of the accounting contact. | [optional] 

## Methods

### NewModelAccountingContactUpdate

`func NewModelAccountingContactUpdate() *ModelAccountingContactUpdate`

NewModelAccountingContactUpdate instantiates a new ModelAccountingContactUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAccountingContactUpdateWithDefaults

`func NewModelAccountingContactUpdateWithDefaults() *ModelAccountingContactUpdate`

NewModelAccountingContactUpdateWithDefaults instantiates a new ModelAccountingContactUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContact

`func (o *ModelAccountingContactUpdate) GetContact() ModelAccountingContactUpdateContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ModelAccountingContactUpdate) GetContactOk() (*ModelAccountingContactUpdateContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ModelAccountingContactUpdate) SetContact(v ModelAccountingContactUpdateContact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *ModelAccountingContactUpdate) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *ModelAccountingContactUpdate) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *ModelAccountingContactUpdate) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetDebitorNumber

`func (o *ModelAccountingContactUpdate) GetDebitorNumber() int32`

GetDebitorNumber returns the DebitorNumber field if non-nil, zero value otherwise.

### GetDebitorNumberOk

`func (o *ModelAccountingContactUpdate) GetDebitorNumberOk() (*int32, bool)`

GetDebitorNumberOk returns a tuple with the DebitorNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitorNumber

`func (o *ModelAccountingContactUpdate) SetDebitorNumber(v int32)`

SetDebitorNumber sets DebitorNumber field to given value.

### HasDebitorNumber

`func (o *ModelAccountingContactUpdate) HasDebitorNumber() bool`

HasDebitorNumber returns a boolean if a field has been set.

### SetDebitorNumberNil

`func (o *ModelAccountingContactUpdate) SetDebitorNumberNil(b bool)`

 SetDebitorNumberNil sets the value for DebitorNumber to be an explicit nil

### UnsetDebitorNumber
`func (o *ModelAccountingContactUpdate) UnsetDebitorNumber()`

UnsetDebitorNumber ensures that no value is present for DebitorNumber, not even an explicit nil
### GetCreditorNumber

`func (o *ModelAccountingContactUpdate) GetCreditorNumber() int32`

GetCreditorNumber returns the CreditorNumber field if non-nil, zero value otherwise.

### GetCreditorNumberOk

`func (o *ModelAccountingContactUpdate) GetCreditorNumberOk() (*int32, bool)`

GetCreditorNumberOk returns a tuple with the CreditorNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditorNumber

`func (o *ModelAccountingContactUpdate) SetCreditorNumber(v int32)`

SetCreditorNumber sets CreditorNumber field to given value.

### HasCreditorNumber

`func (o *ModelAccountingContactUpdate) HasCreditorNumber() bool`

HasCreditorNumber returns a boolean if a field has been set.

### SetCreditorNumberNil

`func (o *ModelAccountingContactUpdate) SetCreditorNumberNil(b bool)`

 SetCreditorNumberNil sets the value for CreditorNumber to be an explicit nil

### UnsetCreditorNumber
`func (o *ModelAccountingContactUpdate) UnsetCreditorNumber()`

UnsetCreditorNumber ensures that no value is present for CreditorNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


