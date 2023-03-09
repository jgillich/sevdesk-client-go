# ModelCheckAccountTransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The check account transaction id | [optional] [readonly] 
**ObjectName** | Pointer to **string** | The check account transaction object name | [optional] [readonly] 
**Create** | Pointer to **time.Time** | Date of check account transaction creation | [optional] [readonly] 
**Update** | Pointer to **time.Time** | Date of last check account transaction update | [optional] [readonly] 
**SevClient** | Pointer to [**ModelCheckAccountTransactionResponseSevClient**](ModelCheckAccountTransactionResponseSevClient.md) |  | [optional] 
**ValueDate** | Pointer to **time.Time** | Date the check account transaction was imported | [optional] [readonly] 
**EntryDate** | Pointer to **time.Time** | Date the check account transaction was booked | [optional] [readonly] 
**PaymtPurpose** | Pointer to **string** | the purpose of the transaction | [optional] [readonly] 
**Amount** | Pointer to **string** | Amount of the transaction | [optional] [readonly] 
**PayeePayerName** | Pointer to **string** | Name of the payee/payer | [optional] [readonly] 
**CheckAccount** | Pointer to [**ModelCheckAccountTransactionResponseCheckAccount**](ModelCheckAccountTransactionResponseCheckAccount.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the check account transaction.&lt;br&gt;       100 &lt;-&gt; Created&lt;br&gt;       200 &lt;-&gt; Linked&lt;br&gt;       300 &lt;-&gt; Private&lt;br&gt;       400 &lt;-&gt; Booked | [optional] [readonly] 
**Enshrined** | Pointer to **time.Time** | Defines if the transaction has been enshrined and can not be changed any more. | [optional] [readonly] 
**SourceTransaction** | Pointer to [**ModelCheckAccountTransactionResponseSourceTransaction**](ModelCheckAccountTransactionResponseSourceTransaction.md) |  | [optional] 
**TargetTransaction** | Pointer to [**ModelCheckAccountTransactionResponseTargetTransaction**](ModelCheckAccountTransactionResponseTargetTransaction.md) |  | [optional] 

## Methods

### NewModelCheckAccountTransactionResponse

`func NewModelCheckAccountTransactionResponse() *ModelCheckAccountTransactionResponse`

NewModelCheckAccountTransactionResponse instantiates a new ModelCheckAccountTransactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelCheckAccountTransactionResponseWithDefaults

`func NewModelCheckAccountTransactionResponseWithDefaults() *ModelCheckAccountTransactionResponse`

NewModelCheckAccountTransactionResponseWithDefaults instantiates a new ModelCheckAccountTransactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelCheckAccountTransactionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelCheckAccountTransactionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelCheckAccountTransactionResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelCheckAccountTransactionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectName

`func (o *ModelCheckAccountTransactionResponse) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ModelCheckAccountTransactionResponse) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ModelCheckAccountTransactionResponse) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ModelCheckAccountTransactionResponse) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetCreate

`func (o *ModelCheckAccountTransactionResponse) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ModelCheckAccountTransactionResponse) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ModelCheckAccountTransactionResponse) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ModelCheckAccountTransactionResponse) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUpdate

`func (o *ModelCheckAccountTransactionResponse) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ModelCheckAccountTransactionResponse) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ModelCheckAccountTransactionResponse) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ModelCheckAccountTransactionResponse) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetSevClient

`func (o *ModelCheckAccountTransactionResponse) GetSevClient() ModelCheckAccountTransactionResponseSevClient`

GetSevClient returns the SevClient field if non-nil, zero value otherwise.

### GetSevClientOk

`func (o *ModelCheckAccountTransactionResponse) GetSevClientOk() (*ModelCheckAccountTransactionResponseSevClient, bool)`

GetSevClientOk returns a tuple with the SevClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSevClient

`func (o *ModelCheckAccountTransactionResponse) SetSevClient(v ModelCheckAccountTransactionResponseSevClient)`

SetSevClient sets SevClient field to given value.

### HasSevClient

`func (o *ModelCheckAccountTransactionResponse) HasSevClient() bool`

HasSevClient returns a boolean if a field has been set.

### GetValueDate

`func (o *ModelCheckAccountTransactionResponse) GetValueDate() time.Time`

GetValueDate returns the ValueDate field if non-nil, zero value otherwise.

### GetValueDateOk

`func (o *ModelCheckAccountTransactionResponse) GetValueDateOk() (*time.Time, bool)`

GetValueDateOk returns a tuple with the ValueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueDate

`func (o *ModelCheckAccountTransactionResponse) SetValueDate(v time.Time)`

SetValueDate sets ValueDate field to given value.

### HasValueDate

`func (o *ModelCheckAccountTransactionResponse) HasValueDate() bool`

HasValueDate returns a boolean if a field has been set.

### GetEntryDate

`func (o *ModelCheckAccountTransactionResponse) GetEntryDate() time.Time`

GetEntryDate returns the EntryDate field if non-nil, zero value otherwise.

### GetEntryDateOk

`func (o *ModelCheckAccountTransactionResponse) GetEntryDateOk() (*time.Time, bool)`

GetEntryDateOk returns a tuple with the EntryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryDate

`func (o *ModelCheckAccountTransactionResponse) SetEntryDate(v time.Time)`

SetEntryDate sets EntryDate field to given value.

### HasEntryDate

`func (o *ModelCheckAccountTransactionResponse) HasEntryDate() bool`

HasEntryDate returns a boolean if a field has been set.

### GetPaymtPurpose

`func (o *ModelCheckAccountTransactionResponse) GetPaymtPurpose() string`

GetPaymtPurpose returns the PaymtPurpose field if non-nil, zero value otherwise.

### GetPaymtPurposeOk

`func (o *ModelCheckAccountTransactionResponse) GetPaymtPurposeOk() (*string, bool)`

GetPaymtPurposeOk returns a tuple with the PaymtPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymtPurpose

`func (o *ModelCheckAccountTransactionResponse) SetPaymtPurpose(v string)`

SetPaymtPurpose sets PaymtPurpose field to given value.

### HasPaymtPurpose

`func (o *ModelCheckAccountTransactionResponse) HasPaymtPurpose() bool`

HasPaymtPurpose returns a boolean if a field has been set.

### GetAmount

`func (o *ModelCheckAccountTransactionResponse) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ModelCheckAccountTransactionResponse) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ModelCheckAccountTransactionResponse) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ModelCheckAccountTransactionResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetPayeePayerName

`func (o *ModelCheckAccountTransactionResponse) GetPayeePayerName() string`

GetPayeePayerName returns the PayeePayerName field if non-nil, zero value otherwise.

### GetPayeePayerNameOk

`func (o *ModelCheckAccountTransactionResponse) GetPayeePayerNameOk() (*string, bool)`

GetPayeePayerNameOk returns a tuple with the PayeePayerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeePayerName

`func (o *ModelCheckAccountTransactionResponse) SetPayeePayerName(v string)`

SetPayeePayerName sets PayeePayerName field to given value.

### HasPayeePayerName

`func (o *ModelCheckAccountTransactionResponse) HasPayeePayerName() bool`

HasPayeePayerName returns a boolean if a field has been set.

### GetCheckAccount

`func (o *ModelCheckAccountTransactionResponse) GetCheckAccount() ModelCheckAccountTransactionResponseCheckAccount`

GetCheckAccount returns the CheckAccount field if non-nil, zero value otherwise.

### GetCheckAccountOk

`func (o *ModelCheckAccountTransactionResponse) GetCheckAccountOk() (*ModelCheckAccountTransactionResponseCheckAccount, bool)`

GetCheckAccountOk returns a tuple with the CheckAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckAccount

`func (o *ModelCheckAccountTransactionResponse) SetCheckAccount(v ModelCheckAccountTransactionResponseCheckAccount)`

SetCheckAccount sets CheckAccount field to given value.

### HasCheckAccount

`func (o *ModelCheckAccountTransactionResponse) HasCheckAccount() bool`

HasCheckAccount returns a boolean if a field has been set.

### GetStatus

`func (o *ModelCheckAccountTransactionResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelCheckAccountTransactionResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelCheckAccountTransactionResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelCheckAccountTransactionResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEnshrined

`func (o *ModelCheckAccountTransactionResponse) GetEnshrined() time.Time`

GetEnshrined returns the Enshrined field if non-nil, zero value otherwise.

### GetEnshrinedOk

`func (o *ModelCheckAccountTransactionResponse) GetEnshrinedOk() (*time.Time, bool)`

GetEnshrinedOk returns a tuple with the Enshrined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnshrined

`func (o *ModelCheckAccountTransactionResponse) SetEnshrined(v time.Time)`

SetEnshrined sets Enshrined field to given value.

### HasEnshrined

`func (o *ModelCheckAccountTransactionResponse) HasEnshrined() bool`

HasEnshrined returns a boolean if a field has been set.

### GetSourceTransaction

`func (o *ModelCheckAccountTransactionResponse) GetSourceTransaction() ModelCheckAccountTransactionResponseSourceTransaction`

GetSourceTransaction returns the SourceTransaction field if non-nil, zero value otherwise.

### GetSourceTransactionOk

`func (o *ModelCheckAccountTransactionResponse) GetSourceTransactionOk() (*ModelCheckAccountTransactionResponseSourceTransaction, bool)`

GetSourceTransactionOk returns a tuple with the SourceTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTransaction

`func (o *ModelCheckAccountTransactionResponse) SetSourceTransaction(v ModelCheckAccountTransactionResponseSourceTransaction)`

SetSourceTransaction sets SourceTransaction field to given value.

### HasSourceTransaction

`func (o *ModelCheckAccountTransactionResponse) HasSourceTransaction() bool`

HasSourceTransaction returns a boolean if a field has been set.

### GetTargetTransaction

`func (o *ModelCheckAccountTransactionResponse) GetTargetTransaction() ModelCheckAccountTransactionResponseTargetTransaction`

GetTargetTransaction returns the TargetTransaction field if non-nil, zero value otherwise.

### GetTargetTransactionOk

`func (o *ModelCheckAccountTransactionResponse) GetTargetTransactionOk() (*ModelCheckAccountTransactionResponseTargetTransaction, bool)`

GetTargetTransactionOk returns a tuple with the TargetTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTransaction

`func (o *ModelCheckAccountTransactionResponse) SetTargetTransaction(v ModelCheckAccountTransactionResponseTargetTransaction)`

SetTargetTransaction sets TargetTransaction field to given value.

### HasTargetTransaction

`func (o *ModelCheckAccountTransactionResponse) HasTargetTransaction() bool`

HasTargetTransaction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


