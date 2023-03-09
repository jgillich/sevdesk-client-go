# ModelAccountingContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | [**ModelAccountingContactContact**](ModelAccountingContactContact.md) |  | 
**DebitorNumber** | Pointer to **NullableInt32** | Debitor number of the accounting contact. | [optional] 
**CreditorNumber** | Pointer to **NullableInt32** | Creditor number of the accounting contact. | [optional] 

## Methods

### NewModelAccountingContact

`func NewModelAccountingContact(contact ModelAccountingContactContact, ) *ModelAccountingContact`

NewModelAccountingContact instantiates a new ModelAccountingContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAccountingContactWithDefaults

`func NewModelAccountingContactWithDefaults() *ModelAccountingContact`

NewModelAccountingContactWithDefaults instantiates a new ModelAccountingContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContact

`func (o *ModelAccountingContact) GetContact() ModelAccountingContactContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ModelAccountingContact) GetContactOk() (*ModelAccountingContactContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ModelAccountingContact) SetContact(v ModelAccountingContactContact)`

SetContact sets Contact field to given value.


### GetDebitorNumber

`func (o *ModelAccountingContact) GetDebitorNumber() int32`

GetDebitorNumber returns the DebitorNumber field if non-nil, zero value otherwise.

### GetDebitorNumberOk

`func (o *ModelAccountingContact) GetDebitorNumberOk() (*int32, bool)`

GetDebitorNumberOk returns a tuple with the DebitorNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitorNumber

`func (o *ModelAccountingContact) SetDebitorNumber(v int32)`

SetDebitorNumber sets DebitorNumber field to given value.

### HasDebitorNumber

`func (o *ModelAccountingContact) HasDebitorNumber() bool`

HasDebitorNumber returns a boolean if a field has been set.

### SetDebitorNumberNil

`func (o *ModelAccountingContact) SetDebitorNumberNil(b bool)`

 SetDebitorNumberNil sets the value for DebitorNumber to be an explicit nil

### UnsetDebitorNumber
`func (o *ModelAccountingContact) UnsetDebitorNumber()`

UnsetDebitorNumber ensures that no value is present for DebitorNumber, not even an explicit nil
### GetCreditorNumber

`func (o *ModelAccountingContact) GetCreditorNumber() int32`

GetCreditorNumber returns the CreditorNumber field if non-nil, zero value otherwise.

### GetCreditorNumberOk

`func (o *ModelAccountingContact) GetCreditorNumberOk() (*int32, bool)`

GetCreditorNumberOk returns a tuple with the CreditorNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditorNumber

`func (o *ModelAccountingContact) SetCreditorNumber(v int32)`

SetCreditorNumber sets CreditorNumber field to given value.

### HasCreditorNumber

`func (o *ModelAccountingContact) HasCreditorNumber() bool`

HasCreditorNumber returns a boolean if a field has been set.

### SetCreditorNumberNil

`func (o *ModelAccountingContact) SetCreditorNumberNil(b bool)`

 SetCreditorNumberNil sets the value for CreditorNumber to be an explicit nil

### UnsetCreditorNumber
`func (o *ModelAccountingContact) UnsetCreditorNumber()`

UnsetCreditorNumber ensures that no value is present for CreditorNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


