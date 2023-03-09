# ModelCheckAccountUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the check account | [optional] 
**Type** | Pointer to **string** | The type of the check account. Account with a CSV or MT940 import are regarded as online.&lt;br&gt;       Apart from that, created check accounts over the API need to be offline, as online accounts with an active connection       to a bank application can not be managed over the API. | [optional] 
**ImportType** | Pointer to **NullableString** | Import type. Transactions can be imported by this method on the check account. | [optional] 
**Currency** | Pointer to **string** | The currency of the check account. | [optional] 
**DefaultAccount** | Pointer to **int32** | Defines if this check account is the default account. | [optional] [default to 0]
**Status** | Pointer to **int32** | Status of the check account. 0 &lt;-&gt; Archived - 100 &lt;-&gt; Active | [optional] [default to 100]
**AutoMapTransactions** | Pointer to **NullableInt32** | Defines if transactions on this account are automatically mapped to invoice and vouchers when imported if possible. | [optional] [default to 1]

## Methods

### NewModelCheckAccountUpdate

`func NewModelCheckAccountUpdate() *ModelCheckAccountUpdate`

NewModelCheckAccountUpdate instantiates a new ModelCheckAccountUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelCheckAccountUpdateWithDefaults

`func NewModelCheckAccountUpdateWithDefaults() *ModelCheckAccountUpdate`

NewModelCheckAccountUpdateWithDefaults instantiates a new ModelCheckAccountUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModelCheckAccountUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelCheckAccountUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelCheckAccountUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelCheckAccountUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ModelCheckAccountUpdate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelCheckAccountUpdate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelCheckAccountUpdate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ModelCheckAccountUpdate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetImportType

`func (o *ModelCheckAccountUpdate) GetImportType() string`

GetImportType returns the ImportType field if non-nil, zero value otherwise.

### GetImportTypeOk

`func (o *ModelCheckAccountUpdate) GetImportTypeOk() (*string, bool)`

GetImportTypeOk returns a tuple with the ImportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportType

`func (o *ModelCheckAccountUpdate) SetImportType(v string)`

SetImportType sets ImportType field to given value.

### HasImportType

`func (o *ModelCheckAccountUpdate) HasImportType() bool`

HasImportType returns a boolean if a field has been set.

### SetImportTypeNil

`func (o *ModelCheckAccountUpdate) SetImportTypeNil(b bool)`

 SetImportTypeNil sets the value for ImportType to be an explicit nil

### UnsetImportType
`func (o *ModelCheckAccountUpdate) UnsetImportType()`

UnsetImportType ensures that no value is present for ImportType, not even an explicit nil
### GetCurrency

`func (o *ModelCheckAccountUpdate) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ModelCheckAccountUpdate) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ModelCheckAccountUpdate) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ModelCheckAccountUpdate) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDefaultAccount

`func (o *ModelCheckAccountUpdate) GetDefaultAccount() int32`

GetDefaultAccount returns the DefaultAccount field if non-nil, zero value otherwise.

### GetDefaultAccountOk

`func (o *ModelCheckAccountUpdate) GetDefaultAccountOk() (*int32, bool)`

GetDefaultAccountOk returns a tuple with the DefaultAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAccount

`func (o *ModelCheckAccountUpdate) SetDefaultAccount(v int32)`

SetDefaultAccount sets DefaultAccount field to given value.

### HasDefaultAccount

`func (o *ModelCheckAccountUpdate) HasDefaultAccount() bool`

HasDefaultAccount returns a boolean if a field has been set.

### GetStatus

`func (o *ModelCheckAccountUpdate) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelCheckAccountUpdate) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelCheckAccountUpdate) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelCheckAccountUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAutoMapTransactions

`func (o *ModelCheckAccountUpdate) GetAutoMapTransactions() int32`

GetAutoMapTransactions returns the AutoMapTransactions field if non-nil, zero value otherwise.

### GetAutoMapTransactionsOk

`func (o *ModelCheckAccountUpdate) GetAutoMapTransactionsOk() (*int32, bool)`

GetAutoMapTransactionsOk returns a tuple with the AutoMapTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoMapTransactions

`func (o *ModelCheckAccountUpdate) SetAutoMapTransactions(v int32)`

SetAutoMapTransactions sets AutoMapTransactions field to given value.

### HasAutoMapTransactions

`func (o *ModelCheckAccountUpdate) HasAutoMapTransactions() bool`

HasAutoMapTransactions returns a boolean if a field has been set.

### SetAutoMapTransactionsNil

`func (o *ModelCheckAccountUpdate) SetAutoMapTransactionsNil(b bool)`

 SetAutoMapTransactionsNil sets the value for AutoMapTransactions to be an explicit nil

### UnsetAutoMapTransactions
`func (o *ModelCheckAccountUpdate) UnsetAutoMapTransactions()`

UnsetAutoMapTransactions ensures that no value is present for AutoMapTransactions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


