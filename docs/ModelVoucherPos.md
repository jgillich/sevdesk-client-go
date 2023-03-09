# ModelVoucherPos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The voucher position id | [optional] [readonly] 
**ObjectName** | **string** | The voucher position object name | 
**MapAll** | **bool** |  | 
**Create** | Pointer to **string** | Date of voucher position creation | [optional] [readonly] 
**Update** | Pointer to **string** | Date of last voucher position update | [optional] [readonly] 
**SevClient** | Pointer to [**ModelVoucherPosSevClient**](ModelVoucherPosSevClient.md) |  | [optional] 
**Voucher** | [**ModelVoucherPosVoucher**](ModelVoucherPosVoucher.md) |  | 
**AccountingType** | [**ModelVoucherPosAccountingType**](ModelVoucherPosAccountingType.md) |  | 
**EstimatedAccountingType** | Pointer to [**ModelVoucherPosEstimatedAccountingType**](ModelVoucherPosEstimatedAccountingType.md) |  | [optional] 
**TaxRate** | **float32** | Tax rate of the voucher position. | 
**Net** | **bool** | Determines whether &#39;sumNet&#39; or &#39;sumGross&#39; is regarded.&lt;br&gt;       If both are not given, &#39;sum&#39; is regarded and treated as net or gross depending on &#39;net&#39;.   All positions must be either net or gross, a mixture of the two is not possible. | 
**IsAsset** | Pointer to **bool** | Determines whether position is regarded as an asset which can be depreciated. | [optional] 
**SumNet** | **float32** | Net sum of the voucher position.&lt;br&gt;      Only regarded if &#39;net&#39; is &#39;true&#39;, otherwise its readOnly. | 
**SumTax** | Pointer to **float32** | Tax sum of the voucher position. | [optional] [readonly] 
**SumGross** | **float32** | Gross sum of the voucher position.&lt;br&gt;      Only regarded if &#39;net&#39; is &#39;false&#39;, otherwise its readOnly. | 
**SumNetAccounting** | Pointer to **float32** | Net accounting sum. Is equal to sumNet. | [optional] [readonly] 
**SumTaxAccounting** | Pointer to **float32** | Tax accounting sum. Is equal to sumTax. | [optional] [readonly] 
**SumGrossAccounting** | Pointer to **float32** | Gross accounting sum. Is equal to sumGross. | [optional] [readonly] 
**Comment** | Pointer to **NullableString** | Comment for the voucher position. | [optional] 

## Methods

### NewModelVoucherPos

`func NewModelVoucherPos(objectName string, mapAll bool, voucher ModelVoucherPosVoucher, accountingType ModelVoucherPosAccountingType, taxRate float32, net bool, sumNet float32, sumGross float32, ) *ModelVoucherPos`

NewModelVoucherPos instantiates a new ModelVoucherPos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelVoucherPosWithDefaults

`func NewModelVoucherPosWithDefaults() *ModelVoucherPos`

NewModelVoucherPosWithDefaults instantiates a new ModelVoucherPos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelVoucherPos) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelVoucherPos) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelVoucherPos) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelVoucherPos) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectName

`func (o *ModelVoucherPos) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ModelVoucherPos) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ModelVoucherPos) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.


### GetMapAll

`func (o *ModelVoucherPos) GetMapAll() bool`

GetMapAll returns the MapAll field if non-nil, zero value otherwise.

### GetMapAllOk

`func (o *ModelVoucherPos) GetMapAllOk() (*bool, bool)`

GetMapAllOk returns a tuple with the MapAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapAll

`func (o *ModelVoucherPos) SetMapAll(v bool)`

SetMapAll sets MapAll field to given value.


### GetCreate

`func (o *ModelVoucherPos) GetCreate() string`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ModelVoucherPos) GetCreateOk() (*string, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ModelVoucherPos) SetCreate(v string)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ModelVoucherPos) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUpdate

`func (o *ModelVoucherPos) GetUpdate() string`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ModelVoucherPos) GetUpdateOk() (*string, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ModelVoucherPos) SetUpdate(v string)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ModelVoucherPos) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetSevClient

`func (o *ModelVoucherPos) GetSevClient() ModelVoucherPosSevClient`

GetSevClient returns the SevClient field if non-nil, zero value otherwise.

### GetSevClientOk

`func (o *ModelVoucherPos) GetSevClientOk() (*ModelVoucherPosSevClient, bool)`

GetSevClientOk returns a tuple with the SevClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSevClient

`func (o *ModelVoucherPos) SetSevClient(v ModelVoucherPosSevClient)`

SetSevClient sets SevClient field to given value.

### HasSevClient

`func (o *ModelVoucherPos) HasSevClient() bool`

HasSevClient returns a boolean if a field has been set.

### GetVoucher

`func (o *ModelVoucherPos) GetVoucher() ModelVoucherPosVoucher`

GetVoucher returns the Voucher field if non-nil, zero value otherwise.

### GetVoucherOk

`func (o *ModelVoucherPos) GetVoucherOk() (*ModelVoucherPosVoucher, bool)`

GetVoucherOk returns a tuple with the Voucher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucher

`func (o *ModelVoucherPos) SetVoucher(v ModelVoucherPosVoucher)`

SetVoucher sets Voucher field to given value.


### GetAccountingType

`func (o *ModelVoucherPos) GetAccountingType() ModelVoucherPosAccountingType`

GetAccountingType returns the AccountingType field if non-nil, zero value otherwise.

### GetAccountingTypeOk

`func (o *ModelVoucherPos) GetAccountingTypeOk() (*ModelVoucherPosAccountingType, bool)`

GetAccountingTypeOk returns a tuple with the AccountingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingType

`func (o *ModelVoucherPos) SetAccountingType(v ModelVoucherPosAccountingType)`

SetAccountingType sets AccountingType field to given value.


### GetEstimatedAccountingType

`func (o *ModelVoucherPos) GetEstimatedAccountingType() ModelVoucherPosEstimatedAccountingType`

GetEstimatedAccountingType returns the EstimatedAccountingType field if non-nil, zero value otherwise.

### GetEstimatedAccountingTypeOk

`func (o *ModelVoucherPos) GetEstimatedAccountingTypeOk() (*ModelVoucherPosEstimatedAccountingType, bool)`

GetEstimatedAccountingTypeOk returns a tuple with the EstimatedAccountingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedAccountingType

`func (o *ModelVoucherPos) SetEstimatedAccountingType(v ModelVoucherPosEstimatedAccountingType)`

SetEstimatedAccountingType sets EstimatedAccountingType field to given value.

### HasEstimatedAccountingType

`func (o *ModelVoucherPos) HasEstimatedAccountingType() bool`

HasEstimatedAccountingType returns a boolean if a field has been set.

### GetTaxRate

`func (o *ModelVoucherPos) GetTaxRate() float32`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *ModelVoucherPos) GetTaxRateOk() (*float32, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *ModelVoucherPos) SetTaxRate(v float32)`

SetTaxRate sets TaxRate field to given value.


### GetNet

`func (o *ModelVoucherPos) GetNet() bool`

GetNet returns the Net field if non-nil, zero value otherwise.

### GetNetOk

`func (o *ModelVoucherPos) GetNetOk() (*bool, bool)`

GetNetOk returns a tuple with the Net field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet

`func (o *ModelVoucherPos) SetNet(v bool)`

SetNet sets Net field to given value.


### GetIsAsset

`func (o *ModelVoucherPos) GetIsAsset() bool`

GetIsAsset returns the IsAsset field if non-nil, zero value otherwise.

### GetIsAssetOk

`func (o *ModelVoucherPos) GetIsAssetOk() (*bool, bool)`

GetIsAssetOk returns a tuple with the IsAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAsset

`func (o *ModelVoucherPos) SetIsAsset(v bool)`

SetIsAsset sets IsAsset field to given value.

### HasIsAsset

`func (o *ModelVoucherPos) HasIsAsset() bool`

HasIsAsset returns a boolean if a field has been set.

### GetSumNet

`func (o *ModelVoucherPos) GetSumNet() float32`

GetSumNet returns the SumNet field if non-nil, zero value otherwise.

### GetSumNetOk

`func (o *ModelVoucherPos) GetSumNetOk() (*float32, bool)`

GetSumNetOk returns a tuple with the SumNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumNet

`func (o *ModelVoucherPos) SetSumNet(v float32)`

SetSumNet sets SumNet field to given value.


### GetSumTax

`func (o *ModelVoucherPos) GetSumTax() float32`

GetSumTax returns the SumTax field if non-nil, zero value otherwise.

### GetSumTaxOk

`func (o *ModelVoucherPos) GetSumTaxOk() (*float32, bool)`

GetSumTaxOk returns a tuple with the SumTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumTax

`func (o *ModelVoucherPos) SetSumTax(v float32)`

SetSumTax sets SumTax field to given value.

### HasSumTax

`func (o *ModelVoucherPos) HasSumTax() bool`

HasSumTax returns a boolean if a field has been set.

### GetSumGross

`func (o *ModelVoucherPos) GetSumGross() float32`

GetSumGross returns the SumGross field if non-nil, zero value otherwise.

### GetSumGrossOk

`func (o *ModelVoucherPos) GetSumGrossOk() (*float32, bool)`

GetSumGrossOk returns a tuple with the SumGross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumGross

`func (o *ModelVoucherPos) SetSumGross(v float32)`

SetSumGross sets SumGross field to given value.


### GetSumNetAccounting

`func (o *ModelVoucherPos) GetSumNetAccounting() float32`

GetSumNetAccounting returns the SumNetAccounting field if non-nil, zero value otherwise.

### GetSumNetAccountingOk

`func (o *ModelVoucherPos) GetSumNetAccountingOk() (*float32, bool)`

GetSumNetAccountingOk returns a tuple with the SumNetAccounting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumNetAccounting

`func (o *ModelVoucherPos) SetSumNetAccounting(v float32)`

SetSumNetAccounting sets SumNetAccounting field to given value.

### HasSumNetAccounting

`func (o *ModelVoucherPos) HasSumNetAccounting() bool`

HasSumNetAccounting returns a boolean if a field has been set.

### GetSumTaxAccounting

`func (o *ModelVoucherPos) GetSumTaxAccounting() float32`

GetSumTaxAccounting returns the SumTaxAccounting field if non-nil, zero value otherwise.

### GetSumTaxAccountingOk

`func (o *ModelVoucherPos) GetSumTaxAccountingOk() (*float32, bool)`

GetSumTaxAccountingOk returns a tuple with the SumTaxAccounting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumTaxAccounting

`func (o *ModelVoucherPos) SetSumTaxAccounting(v float32)`

SetSumTaxAccounting sets SumTaxAccounting field to given value.

### HasSumTaxAccounting

`func (o *ModelVoucherPos) HasSumTaxAccounting() bool`

HasSumTaxAccounting returns a boolean if a field has been set.

### GetSumGrossAccounting

`func (o *ModelVoucherPos) GetSumGrossAccounting() float32`

GetSumGrossAccounting returns the SumGrossAccounting field if non-nil, zero value otherwise.

### GetSumGrossAccountingOk

`func (o *ModelVoucherPos) GetSumGrossAccountingOk() (*float32, bool)`

GetSumGrossAccountingOk returns a tuple with the SumGrossAccounting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumGrossAccounting

`func (o *ModelVoucherPos) SetSumGrossAccounting(v float32)`

SetSumGrossAccounting sets SumGrossAccounting field to given value.

### HasSumGrossAccounting

`func (o *ModelVoucherPos) HasSumGrossAccounting() bool`

HasSumGrossAccounting returns a boolean if a field has been set.

### GetComment

`func (o *ModelVoucherPos) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ModelVoucherPos) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ModelVoucherPos) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ModelVoucherPos) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *ModelVoucherPos) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *ModelVoucherPos) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


