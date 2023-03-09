# ModelCheckAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The check account id | [optional] [readonly] 
**ObjectName** | Pointer to **string** | The check account object name | [optional] [readonly] 
**Create** | Pointer to **time.Time** | Date of check account creation | [optional] [readonly] 
**Update** | Pointer to **time.Time** | Date of last check account update | [optional] [readonly] 
**SevClient** | Pointer to [**ModelCheckAccountResponseSevClient**](ModelCheckAccountResponseSevClient.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the check account | [optional] 
**Type** | Pointer to **string** | The type of the check account. Account with a CSV or MT940 import are regarded as online.&lt;br&gt;       Apart from that, created check accounts over the API need to be offline, as online accounts with an active connection       to a bank application can not be managed over the API. | [optional] 
**ImportType** | Pointer to **NullableString** | Import type. Transactions can be imported by this method on the check account. | [optional] 
**Currency** | Pointer to **string** | The currency of the check account. | [optional] 
**DefaultAccount** | Pointer to **string** | Defines if this check account is the default account. | [optional] [default to "0"]
**Status** | Pointer to **string** | Status of the check account. 0 &lt;-&gt; Archived - 100 &lt;-&gt; Active | [optional] [default to "100"]
**BankServer** | Pointer to **string** | Bank server of check account | [optional] [readonly] 
**AutoMapTransactions** | Pointer to **NullableString** | Defines if transactions on this account are automatically mapped to invoice and vouchers when imported if possible. | [optional] [default to "1"]

## Methods

### NewModelCheckAccountResponse

`func NewModelCheckAccountResponse() *ModelCheckAccountResponse`

NewModelCheckAccountResponse instantiates a new ModelCheckAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelCheckAccountResponseWithDefaults

`func NewModelCheckAccountResponseWithDefaults() *ModelCheckAccountResponse`

NewModelCheckAccountResponseWithDefaults instantiates a new ModelCheckAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelCheckAccountResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelCheckAccountResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelCheckAccountResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelCheckAccountResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectName

`func (o *ModelCheckAccountResponse) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ModelCheckAccountResponse) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ModelCheckAccountResponse) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ModelCheckAccountResponse) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetCreate

`func (o *ModelCheckAccountResponse) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ModelCheckAccountResponse) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ModelCheckAccountResponse) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ModelCheckAccountResponse) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUpdate

`func (o *ModelCheckAccountResponse) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ModelCheckAccountResponse) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ModelCheckAccountResponse) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ModelCheckAccountResponse) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetSevClient

`func (o *ModelCheckAccountResponse) GetSevClient() ModelCheckAccountResponseSevClient`

GetSevClient returns the SevClient field if non-nil, zero value otherwise.

### GetSevClientOk

`func (o *ModelCheckAccountResponse) GetSevClientOk() (*ModelCheckAccountResponseSevClient, bool)`

GetSevClientOk returns a tuple with the SevClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSevClient

`func (o *ModelCheckAccountResponse) SetSevClient(v ModelCheckAccountResponseSevClient)`

SetSevClient sets SevClient field to given value.

### HasSevClient

`func (o *ModelCheckAccountResponse) HasSevClient() bool`

HasSevClient returns a boolean if a field has been set.

### GetName

`func (o *ModelCheckAccountResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelCheckAccountResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelCheckAccountResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelCheckAccountResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ModelCheckAccountResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelCheckAccountResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelCheckAccountResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ModelCheckAccountResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetImportType

`func (o *ModelCheckAccountResponse) GetImportType() string`

GetImportType returns the ImportType field if non-nil, zero value otherwise.

### GetImportTypeOk

`func (o *ModelCheckAccountResponse) GetImportTypeOk() (*string, bool)`

GetImportTypeOk returns a tuple with the ImportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportType

`func (o *ModelCheckAccountResponse) SetImportType(v string)`

SetImportType sets ImportType field to given value.

### HasImportType

`func (o *ModelCheckAccountResponse) HasImportType() bool`

HasImportType returns a boolean if a field has been set.

### SetImportTypeNil

`func (o *ModelCheckAccountResponse) SetImportTypeNil(b bool)`

 SetImportTypeNil sets the value for ImportType to be an explicit nil

### UnsetImportType
`func (o *ModelCheckAccountResponse) UnsetImportType()`

UnsetImportType ensures that no value is present for ImportType, not even an explicit nil
### GetCurrency

`func (o *ModelCheckAccountResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ModelCheckAccountResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ModelCheckAccountResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ModelCheckAccountResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDefaultAccount

`func (o *ModelCheckAccountResponse) GetDefaultAccount() string`

GetDefaultAccount returns the DefaultAccount field if non-nil, zero value otherwise.

### GetDefaultAccountOk

`func (o *ModelCheckAccountResponse) GetDefaultAccountOk() (*string, bool)`

GetDefaultAccountOk returns a tuple with the DefaultAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAccount

`func (o *ModelCheckAccountResponse) SetDefaultAccount(v string)`

SetDefaultAccount sets DefaultAccount field to given value.

### HasDefaultAccount

`func (o *ModelCheckAccountResponse) HasDefaultAccount() bool`

HasDefaultAccount returns a boolean if a field has been set.

### GetStatus

`func (o *ModelCheckAccountResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelCheckAccountResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelCheckAccountResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelCheckAccountResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetBankServer

`func (o *ModelCheckAccountResponse) GetBankServer() string`

GetBankServer returns the BankServer field if non-nil, zero value otherwise.

### GetBankServerOk

`func (o *ModelCheckAccountResponse) GetBankServerOk() (*string, bool)`

GetBankServerOk returns a tuple with the BankServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankServer

`func (o *ModelCheckAccountResponse) SetBankServer(v string)`

SetBankServer sets BankServer field to given value.

### HasBankServer

`func (o *ModelCheckAccountResponse) HasBankServer() bool`

HasBankServer returns a boolean if a field has been set.

### GetAutoMapTransactions

`func (o *ModelCheckAccountResponse) GetAutoMapTransactions() string`

GetAutoMapTransactions returns the AutoMapTransactions field if non-nil, zero value otherwise.

### GetAutoMapTransactionsOk

`func (o *ModelCheckAccountResponse) GetAutoMapTransactionsOk() (*string, bool)`

GetAutoMapTransactionsOk returns a tuple with the AutoMapTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoMapTransactions

`func (o *ModelCheckAccountResponse) SetAutoMapTransactions(v string)`

SetAutoMapTransactions sets AutoMapTransactions field to given value.

### HasAutoMapTransactions

`func (o *ModelCheckAccountResponse) HasAutoMapTransactions() bool`

HasAutoMapTransactions returns a boolean if a field has been set.

### SetAutoMapTransactionsNil

`func (o *ModelCheckAccountResponse) SetAutoMapTransactionsNil(b bool)`

 SetAutoMapTransactionsNil sets the value for AutoMapTransactions to be an explicit nil

### UnsetAutoMapTransactions
`func (o *ModelCheckAccountResponse) UnsetAutoMapTransactions()`

UnsetAutoMapTransactions ensures that no value is present for AutoMapTransactions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


