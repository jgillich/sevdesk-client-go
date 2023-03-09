# ModelCheckAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The check account id | [optional] [readonly] 
**ObjectName** | Pointer to **string** | The check account object name | [optional] [readonly] 
**Create** | Pointer to **time.Time** | Date of check account creation | [optional] [readonly] 
**Update** | Pointer to **time.Time** | Date of last check account update | [optional] [readonly] 
**SevClient** | Pointer to [**ModelCheckAccountSevClient**](ModelCheckAccountSevClient.md) |  | [optional] 
**Name** | **string** | Name of the check account | 
**Type** | **string** | The type of the check account. Account with a CSV or MT940 import are regarded as online.&lt;br&gt;       Apart from that, created check accounts over the API need to be offline, as online accounts with an active connection       to a bank application can not be managed over the API. | 
**ImportType** | Pointer to **NullableString** | Import type. Transactions can be imported by this method on the check account. | [optional] 
**Currency** | **string** | The currency of the check account. | 
**DefaultAccount** | Pointer to **int32** | Defines if this check account is the default account. | [optional] [default to 0]
**Status** | **int32** | Status of the check account. 0 &lt;-&gt; Archived - 100 &lt;-&gt; Active | [default to 100]
**BankServer** | Pointer to **string** | Bank server of check account | [optional] [readonly] 
**AutoMapTransactions** | Pointer to **NullableInt32** | Defines if transactions on this account are automatically mapped to invoice and vouchers when imported if possible. | [optional] [default to 1]

## Methods

### NewModelCheckAccount

`func NewModelCheckAccount(name string, type_ string, currency string, status int32, ) *ModelCheckAccount`

NewModelCheckAccount instantiates a new ModelCheckAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelCheckAccountWithDefaults

`func NewModelCheckAccountWithDefaults() *ModelCheckAccount`

NewModelCheckAccountWithDefaults instantiates a new ModelCheckAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelCheckAccount) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelCheckAccount) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelCheckAccount) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelCheckAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectName

`func (o *ModelCheckAccount) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ModelCheckAccount) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ModelCheckAccount) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ModelCheckAccount) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetCreate

`func (o *ModelCheckAccount) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ModelCheckAccount) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ModelCheckAccount) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ModelCheckAccount) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUpdate

`func (o *ModelCheckAccount) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ModelCheckAccount) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ModelCheckAccount) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ModelCheckAccount) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetSevClient

`func (o *ModelCheckAccount) GetSevClient() ModelCheckAccountSevClient`

GetSevClient returns the SevClient field if non-nil, zero value otherwise.

### GetSevClientOk

`func (o *ModelCheckAccount) GetSevClientOk() (*ModelCheckAccountSevClient, bool)`

GetSevClientOk returns a tuple with the SevClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSevClient

`func (o *ModelCheckAccount) SetSevClient(v ModelCheckAccountSevClient)`

SetSevClient sets SevClient field to given value.

### HasSevClient

`func (o *ModelCheckAccount) HasSevClient() bool`

HasSevClient returns a boolean if a field has been set.

### GetName

`func (o *ModelCheckAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelCheckAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelCheckAccount) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ModelCheckAccount) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelCheckAccount) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelCheckAccount) SetType(v string)`

SetType sets Type field to given value.


### GetImportType

`func (o *ModelCheckAccount) GetImportType() string`

GetImportType returns the ImportType field if non-nil, zero value otherwise.

### GetImportTypeOk

`func (o *ModelCheckAccount) GetImportTypeOk() (*string, bool)`

GetImportTypeOk returns a tuple with the ImportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportType

`func (o *ModelCheckAccount) SetImportType(v string)`

SetImportType sets ImportType field to given value.

### HasImportType

`func (o *ModelCheckAccount) HasImportType() bool`

HasImportType returns a boolean if a field has been set.

### SetImportTypeNil

`func (o *ModelCheckAccount) SetImportTypeNil(b bool)`

 SetImportTypeNil sets the value for ImportType to be an explicit nil

### UnsetImportType
`func (o *ModelCheckAccount) UnsetImportType()`

UnsetImportType ensures that no value is present for ImportType, not even an explicit nil
### GetCurrency

`func (o *ModelCheckAccount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ModelCheckAccount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ModelCheckAccount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDefaultAccount

`func (o *ModelCheckAccount) GetDefaultAccount() int32`

GetDefaultAccount returns the DefaultAccount field if non-nil, zero value otherwise.

### GetDefaultAccountOk

`func (o *ModelCheckAccount) GetDefaultAccountOk() (*int32, bool)`

GetDefaultAccountOk returns a tuple with the DefaultAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAccount

`func (o *ModelCheckAccount) SetDefaultAccount(v int32)`

SetDefaultAccount sets DefaultAccount field to given value.

### HasDefaultAccount

`func (o *ModelCheckAccount) HasDefaultAccount() bool`

HasDefaultAccount returns a boolean if a field has been set.

### GetStatus

`func (o *ModelCheckAccount) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelCheckAccount) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelCheckAccount) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetBankServer

`func (o *ModelCheckAccount) GetBankServer() string`

GetBankServer returns the BankServer field if non-nil, zero value otherwise.

### GetBankServerOk

`func (o *ModelCheckAccount) GetBankServerOk() (*string, bool)`

GetBankServerOk returns a tuple with the BankServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankServer

`func (o *ModelCheckAccount) SetBankServer(v string)`

SetBankServer sets BankServer field to given value.

### HasBankServer

`func (o *ModelCheckAccount) HasBankServer() bool`

HasBankServer returns a boolean if a field has been set.

### GetAutoMapTransactions

`func (o *ModelCheckAccount) GetAutoMapTransactions() int32`

GetAutoMapTransactions returns the AutoMapTransactions field if non-nil, zero value otherwise.

### GetAutoMapTransactionsOk

`func (o *ModelCheckAccount) GetAutoMapTransactionsOk() (*int32, bool)`

GetAutoMapTransactionsOk returns a tuple with the AutoMapTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoMapTransactions

`func (o *ModelCheckAccount) SetAutoMapTransactions(v int32)`

SetAutoMapTransactions sets AutoMapTransactions field to given value.

### HasAutoMapTransactions

`func (o *ModelCheckAccount) HasAutoMapTransactions() bool`

HasAutoMapTransactions returns a boolean if a field has been set.

### SetAutoMapTransactionsNil

`func (o *ModelCheckAccount) SetAutoMapTransactionsNil(b bool)`

 SetAutoMapTransactionsNil sets the value for AutoMapTransactions to be an explicit nil

### UnsetAutoMapTransactions
`func (o *ModelCheckAccount) UnsetAutoMapTransactions()`

UnsetAutoMapTransactions ensures that no value is present for AutoMapTransactions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


