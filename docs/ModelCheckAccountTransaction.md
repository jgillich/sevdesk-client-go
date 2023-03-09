# ModelCheckAccountTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The check account transaction id | [optional] [readonly] 
**ObjectName** | Pointer to **string** | The check account transaction object name | [optional] [readonly] 
**Create** | Pointer to **time.Time** | Date of check account transaction creation | [optional] [readonly] 
**Update** | Pointer to **time.Time** | Date of last check account transaction update | [optional] [readonly] 
**SevClient** | Pointer to [**ModelCheckAccountTransactionSevClient**](ModelCheckAccountTransactionSevClient.md) |  | [optional] 
**ValueDate** | **time.Time** | Date the check account transaction was booked | 
**EntryDate** | Pointer to **time.Time** | Date the check account transaction was imported | [optional] 
**PaymtPurpose** | Pointer to **string** | the purpose of the transaction | [optional] 
**Amount** | **float32** | Amount of the transaction | 
**PayeePayerName** | **string** | Name of the payee/payer | 
**CheckAccount** | [**ModelCheckAccountTransactionCheckAccount**](ModelCheckAccountTransactionCheckAccount.md) |  | 
**Status** | **int32** | Status of the check account transaction.&lt;br&gt;       100 &lt;-&gt; Created&lt;br&gt;       200 &lt;-&gt; Linked&lt;br&gt;       300 &lt;-&gt; Private&lt;br&gt;       400 &lt;-&gt; Booked | 
**Enshrined** | Pointer to **NullableTime** | Defines if the transaction has been enshrined and can not be changed any more. | [optional] 
**SourceTransaction** | Pointer to [**NullableModelCheckAccountTransactionSourceTransaction**](ModelCheckAccountTransactionSourceTransaction.md) |  | [optional] 
**TargetTransaction** | Pointer to [**NullableModelCheckAccountTransactionTargetTransaction**](ModelCheckAccountTransactionTargetTransaction.md) |  | [optional] 

## Methods

### NewModelCheckAccountTransaction

`func NewModelCheckAccountTransaction(valueDate time.Time, amount float32, payeePayerName string, checkAccount ModelCheckAccountTransactionCheckAccount, status int32, ) *ModelCheckAccountTransaction`

NewModelCheckAccountTransaction instantiates a new ModelCheckAccountTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelCheckAccountTransactionWithDefaults

`func NewModelCheckAccountTransactionWithDefaults() *ModelCheckAccountTransaction`

NewModelCheckAccountTransactionWithDefaults instantiates a new ModelCheckAccountTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelCheckAccountTransaction) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelCheckAccountTransaction) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelCheckAccountTransaction) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelCheckAccountTransaction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectName

`func (o *ModelCheckAccountTransaction) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ModelCheckAccountTransaction) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ModelCheckAccountTransaction) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ModelCheckAccountTransaction) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetCreate

`func (o *ModelCheckAccountTransaction) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ModelCheckAccountTransaction) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ModelCheckAccountTransaction) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ModelCheckAccountTransaction) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUpdate

`func (o *ModelCheckAccountTransaction) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ModelCheckAccountTransaction) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ModelCheckAccountTransaction) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ModelCheckAccountTransaction) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetSevClient

`func (o *ModelCheckAccountTransaction) GetSevClient() ModelCheckAccountTransactionSevClient`

GetSevClient returns the SevClient field if non-nil, zero value otherwise.

### GetSevClientOk

`func (o *ModelCheckAccountTransaction) GetSevClientOk() (*ModelCheckAccountTransactionSevClient, bool)`

GetSevClientOk returns a tuple with the SevClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSevClient

`func (o *ModelCheckAccountTransaction) SetSevClient(v ModelCheckAccountTransactionSevClient)`

SetSevClient sets SevClient field to given value.

### HasSevClient

`func (o *ModelCheckAccountTransaction) HasSevClient() bool`

HasSevClient returns a boolean if a field has been set.

### GetValueDate

`func (o *ModelCheckAccountTransaction) GetValueDate() time.Time`

GetValueDate returns the ValueDate field if non-nil, zero value otherwise.

### GetValueDateOk

`func (o *ModelCheckAccountTransaction) GetValueDateOk() (*time.Time, bool)`

GetValueDateOk returns a tuple with the ValueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueDate

`func (o *ModelCheckAccountTransaction) SetValueDate(v time.Time)`

SetValueDate sets ValueDate field to given value.


### GetEntryDate

`func (o *ModelCheckAccountTransaction) GetEntryDate() time.Time`

GetEntryDate returns the EntryDate field if non-nil, zero value otherwise.

### GetEntryDateOk

`func (o *ModelCheckAccountTransaction) GetEntryDateOk() (*time.Time, bool)`

GetEntryDateOk returns a tuple with the EntryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryDate

`func (o *ModelCheckAccountTransaction) SetEntryDate(v time.Time)`

SetEntryDate sets EntryDate field to given value.

### HasEntryDate

`func (o *ModelCheckAccountTransaction) HasEntryDate() bool`

HasEntryDate returns a boolean if a field has been set.

### GetPaymtPurpose

`func (o *ModelCheckAccountTransaction) GetPaymtPurpose() string`

GetPaymtPurpose returns the PaymtPurpose field if non-nil, zero value otherwise.

### GetPaymtPurposeOk

`func (o *ModelCheckAccountTransaction) GetPaymtPurposeOk() (*string, bool)`

GetPaymtPurposeOk returns a tuple with the PaymtPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymtPurpose

`func (o *ModelCheckAccountTransaction) SetPaymtPurpose(v string)`

SetPaymtPurpose sets PaymtPurpose field to given value.

### HasPaymtPurpose

`func (o *ModelCheckAccountTransaction) HasPaymtPurpose() bool`

HasPaymtPurpose returns a boolean if a field has been set.

### GetAmount

`func (o *ModelCheckAccountTransaction) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ModelCheckAccountTransaction) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ModelCheckAccountTransaction) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetPayeePayerName

`func (o *ModelCheckAccountTransaction) GetPayeePayerName() string`

GetPayeePayerName returns the PayeePayerName field if non-nil, zero value otherwise.

### GetPayeePayerNameOk

`func (o *ModelCheckAccountTransaction) GetPayeePayerNameOk() (*string, bool)`

GetPayeePayerNameOk returns a tuple with the PayeePayerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeePayerName

`func (o *ModelCheckAccountTransaction) SetPayeePayerName(v string)`

SetPayeePayerName sets PayeePayerName field to given value.


### GetCheckAccount

`func (o *ModelCheckAccountTransaction) GetCheckAccount() ModelCheckAccountTransactionCheckAccount`

GetCheckAccount returns the CheckAccount field if non-nil, zero value otherwise.

### GetCheckAccountOk

`func (o *ModelCheckAccountTransaction) GetCheckAccountOk() (*ModelCheckAccountTransactionCheckAccount, bool)`

GetCheckAccountOk returns a tuple with the CheckAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckAccount

`func (o *ModelCheckAccountTransaction) SetCheckAccount(v ModelCheckAccountTransactionCheckAccount)`

SetCheckAccount sets CheckAccount field to given value.


### GetStatus

`func (o *ModelCheckAccountTransaction) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelCheckAccountTransaction) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelCheckAccountTransaction) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetEnshrined

`func (o *ModelCheckAccountTransaction) GetEnshrined() time.Time`

GetEnshrined returns the Enshrined field if non-nil, zero value otherwise.

### GetEnshrinedOk

`func (o *ModelCheckAccountTransaction) GetEnshrinedOk() (*time.Time, bool)`

GetEnshrinedOk returns a tuple with the Enshrined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnshrined

`func (o *ModelCheckAccountTransaction) SetEnshrined(v time.Time)`

SetEnshrined sets Enshrined field to given value.

### HasEnshrined

`func (o *ModelCheckAccountTransaction) HasEnshrined() bool`

HasEnshrined returns a boolean if a field has been set.

### SetEnshrinedNil

`func (o *ModelCheckAccountTransaction) SetEnshrinedNil(b bool)`

 SetEnshrinedNil sets the value for Enshrined to be an explicit nil

### UnsetEnshrined
`func (o *ModelCheckAccountTransaction) UnsetEnshrined()`

UnsetEnshrined ensures that no value is present for Enshrined, not even an explicit nil
### GetSourceTransaction

`func (o *ModelCheckAccountTransaction) GetSourceTransaction() ModelCheckAccountTransactionSourceTransaction`

GetSourceTransaction returns the SourceTransaction field if non-nil, zero value otherwise.

### GetSourceTransactionOk

`func (o *ModelCheckAccountTransaction) GetSourceTransactionOk() (*ModelCheckAccountTransactionSourceTransaction, bool)`

GetSourceTransactionOk returns a tuple with the SourceTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTransaction

`func (o *ModelCheckAccountTransaction) SetSourceTransaction(v ModelCheckAccountTransactionSourceTransaction)`

SetSourceTransaction sets SourceTransaction field to given value.

### HasSourceTransaction

`func (o *ModelCheckAccountTransaction) HasSourceTransaction() bool`

HasSourceTransaction returns a boolean if a field has been set.

### SetSourceTransactionNil

`func (o *ModelCheckAccountTransaction) SetSourceTransactionNil(b bool)`

 SetSourceTransactionNil sets the value for SourceTransaction to be an explicit nil

### UnsetSourceTransaction
`func (o *ModelCheckAccountTransaction) UnsetSourceTransaction()`

UnsetSourceTransaction ensures that no value is present for SourceTransaction, not even an explicit nil
### GetTargetTransaction

`func (o *ModelCheckAccountTransaction) GetTargetTransaction() ModelCheckAccountTransactionTargetTransaction`

GetTargetTransaction returns the TargetTransaction field if non-nil, zero value otherwise.

### GetTargetTransactionOk

`func (o *ModelCheckAccountTransaction) GetTargetTransactionOk() (*ModelCheckAccountTransactionTargetTransaction, bool)`

GetTargetTransactionOk returns a tuple with the TargetTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTransaction

`func (o *ModelCheckAccountTransaction) SetTargetTransaction(v ModelCheckAccountTransactionTargetTransaction)`

SetTargetTransaction sets TargetTransaction field to given value.

### HasTargetTransaction

`func (o *ModelCheckAccountTransaction) HasTargetTransaction() bool`

HasTargetTransaction returns a boolean if a field has been set.

### SetTargetTransactionNil

`func (o *ModelCheckAccountTransaction) SetTargetTransactionNil(b bool)`

 SetTargetTransactionNil sets the value for TargetTransaction to be an explicit nil

### UnsetTargetTransaction
`func (o *ModelCheckAccountTransaction) UnsetTargetTransaction()`

UnsetTargetTransaction ensures that no value is present for TargetTransaction, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


