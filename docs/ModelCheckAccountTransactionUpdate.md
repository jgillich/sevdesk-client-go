# ModelCheckAccountTransactionUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValueDate** | Pointer to **time.Time** | Date the check account transaction was booked | [optional] 
**EntryDate** | Pointer to **time.Time** | Date the check account transaction was imported | [optional] 
**PaymtPurpose** | Pointer to **string** | the purpose of the transaction | [optional] 
**Amount** | Pointer to **NullableFloat32** | Amount of the transaction | [optional] 
**PayeePayerName** | Pointer to **NullableString** | Name of the payee/payer | [optional] 
**CheckAccount** | Pointer to [**ModelCheckAccountTransactionUpdateCheckAccount**](ModelCheckAccountTransactionUpdateCheckAccount.md) |  | [optional] 
**Status** | Pointer to **int32** | Status of the check account transaction.&lt;br&gt;       100 &lt;-&gt; Created&lt;br&gt;       200 &lt;-&gt; Linked&lt;br&gt;       300 &lt;-&gt; Private&lt;br&gt;       400 &lt;-&gt; Booked | [optional] 
**Enshrined** | Pointer to **NullableTime** | Defines if the transaction has been enshrined and can not be changed any more. | [optional] 
**SourceTransaction** | Pointer to [**NullableModelCheckAccountTransactionSourceTransaction**](ModelCheckAccountTransactionSourceTransaction.md) |  | [optional] 
**TargetTransaction** | Pointer to [**NullableModelCheckAccountTransactionTargetTransaction**](ModelCheckAccountTransactionTargetTransaction.md) |  | [optional] 

## Methods

### NewModelCheckAccountTransactionUpdate

`func NewModelCheckAccountTransactionUpdate() *ModelCheckAccountTransactionUpdate`

NewModelCheckAccountTransactionUpdate instantiates a new ModelCheckAccountTransactionUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelCheckAccountTransactionUpdateWithDefaults

`func NewModelCheckAccountTransactionUpdateWithDefaults() *ModelCheckAccountTransactionUpdate`

NewModelCheckAccountTransactionUpdateWithDefaults instantiates a new ModelCheckAccountTransactionUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValueDate

`func (o *ModelCheckAccountTransactionUpdate) GetValueDate() time.Time`

GetValueDate returns the ValueDate field if non-nil, zero value otherwise.

### GetValueDateOk

`func (o *ModelCheckAccountTransactionUpdate) GetValueDateOk() (*time.Time, bool)`

GetValueDateOk returns a tuple with the ValueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueDate

`func (o *ModelCheckAccountTransactionUpdate) SetValueDate(v time.Time)`

SetValueDate sets ValueDate field to given value.

### HasValueDate

`func (o *ModelCheckAccountTransactionUpdate) HasValueDate() bool`

HasValueDate returns a boolean if a field has been set.

### GetEntryDate

`func (o *ModelCheckAccountTransactionUpdate) GetEntryDate() time.Time`

GetEntryDate returns the EntryDate field if non-nil, zero value otherwise.

### GetEntryDateOk

`func (o *ModelCheckAccountTransactionUpdate) GetEntryDateOk() (*time.Time, bool)`

GetEntryDateOk returns a tuple with the EntryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryDate

`func (o *ModelCheckAccountTransactionUpdate) SetEntryDate(v time.Time)`

SetEntryDate sets EntryDate field to given value.

### HasEntryDate

`func (o *ModelCheckAccountTransactionUpdate) HasEntryDate() bool`

HasEntryDate returns a boolean if a field has been set.

### GetPaymtPurpose

`func (o *ModelCheckAccountTransactionUpdate) GetPaymtPurpose() string`

GetPaymtPurpose returns the PaymtPurpose field if non-nil, zero value otherwise.

### GetPaymtPurposeOk

`func (o *ModelCheckAccountTransactionUpdate) GetPaymtPurposeOk() (*string, bool)`

GetPaymtPurposeOk returns a tuple with the PaymtPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymtPurpose

`func (o *ModelCheckAccountTransactionUpdate) SetPaymtPurpose(v string)`

SetPaymtPurpose sets PaymtPurpose field to given value.

### HasPaymtPurpose

`func (o *ModelCheckAccountTransactionUpdate) HasPaymtPurpose() bool`

HasPaymtPurpose returns a boolean if a field has been set.

### GetAmount

`func (o *ModelCheckAccountTransactionUpdate) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ModelCheckAccountTransactionUpdate) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ModelCheckAccountTransactionUpdate) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ModelCheckAccountTransactionUpdate) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *ModelCheckAccountTransactionUpdate) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *ModelCheckAccountTransactionUpdate) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetPayeePayerName

`func (o *ModelCheckAccountTransactionUpdate) GetPayeePayerName() string`

GetPayeePayerName returns the PayeePayerName field if non-nil, zero value otherwise.

### GetPayeePayerNameOk

`func (o *ModelCheckAccountTransactionUpdate) GetPayeePayerNameOk() (*string, bool)`

GetPayeePayerNameOk returns a tuple with the PayeePayerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeePayerName

`func (o *ModelCheckAccountTransactionUpdate) SetPayeePayerName(v string)`

SetPayeePayerName sets PayeePayerName field to given value.

### HasPayeePayerName

`func (o *ModelCheckAccountTransactionUpdate) HasPayeePayerName() bool`

HasPayeePayerName returns a boolean if a field has been set.

### SetPayeePayerNameNil

`func (o *ModelCheckAccountTransactionUpdate) SetPayeePayerNameNil(b bool)`

 SetPayeePayerNameNil sets the value for PayeePayerName to be an explicit nil

### UnsetPayeePayerName
`func (o *ModelCheckAccountTransactionUpdate) UnsetPayeePayerName()`

UnsetPayeePayerName ensures that no value is present for PayeePayerName, not even an explicit nil
### GetCheckAccount

`func (o *ModelCheckAccountTransactionUpdate) GetCheckAccount() ModelCheckAccountTransactionUpdateCheckAccount`

GetCheckAccount returns the CheckAccount field if non-nil, zero value otherwise.

### GetCheckAccountOk

`func (o *ModelCheckAccountTransactionUpdate) GetCheckAccountOk() (*ModelCheckAccountTransactionUpdateCheckAccount, bool)`

GetCheckAccountOk returns a tuple with the CheckAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckAccount

`func (o *ModelCheckAccountTransactionUpdate) SetCheckAccount(v ModelCheckAccountTransactionUpdateCheckAccount)`

SetCheckAccount sets CheckAccount field to given value.

### HasCheckAccount

`func (o *ModelCheckAccountTransactionUpdate) HasCheckAccount() bool`

HasCheckAccount returns a boolean if a field has been set.

### GetStatus

`func (o *ModelCheckAccountTransactionUpdate) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelCheckAccountTransactionUpdate) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelCheckAccountTransactionUpdate) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelCheckAccountTransactionUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEnshrined

`func (o *ModelCheckAccountTransactionUpdate) GetEnshrined() time.Time`

GetEnshrined returns the Enshrined field if non-nil, zero value otherwise.

### GetEnshrinedOk

`func (o *ModelCheckAccountTransactionUpdate) GetEnshrinedOk() (*time.Time, bool)`

GetEnshrinedOk returns a tuple with the Enshrined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnshrined

`func (o *ModelCheckAccountTransactionUpdate) SetEnshrined(v time.Time)`

SetEnshrined sets Enshrined field to given value.

### HasEnshrined

`func (o *ModelCheckAccountTransactionUpdate) HasEnshrined() bool`

HasEnshrined returns a boolean if a field has been set.

### SetEnshrinedNil

`func (o *ModelCheckAccountTransactionUpdate) SetEnshrinedNil(b bool)`

 SetEnshrinedNil sets the value for Enshrined to be an explicit nil

### UnsetEnshrined
`func (o *ModelCheckAccountTransactionUpdate) UnsetEnshrined()`

UnsetEnshrined ensures that no value is present for Enshrined, not even an explicit nil
### GetSourceTransaction

`func (o *ModelCheckAccountTransactionUpdate) GetSourceTransaction() ModelCheckAccountTransactionSourceTransaction`

GetSourceTransaction returns the SourceTransaction field if non-nil, zero value otherwise.

### GetSourceTransactionOk

`func (o *ModelCheckAccountTransactionUpdate) GetSourceTransactionOk() (*ModelCheckAccountTransactionSourceTransaction, bool)`

GetSourceTransactionOk returns a tuple with the SourceTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTransaction

`func (o *ModelCheckAccountTransactionUpdate) SetSourceTransaction(v ModelCheckAccountTransactionSourceTransaction)`

SetSourceTransaction sets SourceTransaction field to given value.

### HasSourceTransaction

`func (o *ModelCheckAccountTransactionUpdate) HasSourceTransaction() bool`

HasSourceTransaction returns a boolean if a field has been set.

### SetSourceTransactionNil

`func (o *ModelCheckAccountTransactionUpdate) SetSourceTransactionNil(b bool)`

 SetSourceTransactionNil sets the value for SourceTransaction to be an explicit nil

### UnsetSourceTransaction
`func (o *ModelCheckAccountTransactionUpdate) UnsetSourceTransaction()`

UnsetSourceTransaction ensures that no value is present for SourceTransaction, not even an explicit nil
### GetTargetTransaction

`func (o *ModelCheckAccountTransactionUpdate) GetTargetTransaction() ModelCheckAccountTransactionTargetTransaction`

GetTargetTransaction returns the TargetTransaction field if non-nil, zero value otherwise.

### GetTargetTransactionOk

`func (o *ModelCheckAccountTransactionUpdate) GetTargetTransactionOk() (*ModelCheckAccountTransactionTargetTransaction, bool)`

GetTargetTransactionOk returns a tuple with the TargetTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTransaction

`func (o *ModelCheckAccountTransactionUpdate) SetTargetTransaction(v ModelCheckAccountTransactionTargetTransaction)`

SetTargetTransaction sets TargetTransaction field to given value.

### HasTargetTransaction

`func (o *ModelCheckAccountTransactionUpdate) HasTargetTransaction() bool`

HasTargetTransaction returns a boolean if a field has been set.

### SetTargetTransactionNil

`func (o *ModelCheckAccountTransactionUpdate) SetTargetTransactionNil(b bool)`

 SetTargetTransactionNil sets the value for TargetTransaction to be an explicit nil

### UnsetTargetTransaction
`func (o *ModelCheckAccountTransactionUpdate) UnsetTargetTransaction()`

UnsetTargetTransaction ensures that no value is present for TargetTransaction, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


