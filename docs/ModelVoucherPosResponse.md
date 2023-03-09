# ModelVoucherPosResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The voucher position id | [optional] [readonly] 
**ObjectName** | Pointer to **string** | The voucher position object name | [optional] [readonly] 
**Create** | Pointer to **string** | Date of voucher position creation | [optional] [readonly] 
**Update** | Pointer to **string** | Date of last voucher position update | [optional] [readonly] 
**SevClient** | Pointer to [**ModelVoucherPosResponseSevClient**](ModelVoucherPosResponseSevClient.md) |  | [optional] 
**Voucher** | [**ModelVoucherPosResponseVoucher**](ModelVoucherPosResponseVoucher.md) |  | 
**AccountingType** | [**ModelVoucherPosResponseAccountingType**](ModelVoucherPosResponseAccountingType.md) |  | 
**EstimatedAccountingType** | Pointer to [**ModelVoucherPosResponseEstimatedAccountingType**](ModelVoucherPosResponseEstimatedAccountingType.md) |  | [optional] 
**TaxRate** | **string** | Tax rate of the voucher position. | 
**Net** | **bool** | Determines whether &#39;sumNet&#39; or &#39;sumGross&#39; is regarded.&lt;br&gt;       If both are not given, &#39;sum&#39; is regarded and treated as net or gross depending on &#39;net&#39;.  All positions must be either net or gross, a mixture of the two is not possible. | 
**IsAsset** | Pointer to **bool** | Determines whether position is regarded as an asset which can be depreciated. | [optional] 
**SumNet** | **string** | Net sum of the voucher position.&lt;br&gt;      Only regarded if &#39;net&#39; is &#39;true&#39;, otherwise its readOnly. | 
**SumTax** | Pointer to **string** | Tax sum of the voucher position. | [optional] [readonly] 
**SumGross** | **string** | Gross sum of the voucher position.&lt;br&gt;      Only regarded if &#39;net&#39; is &#39;false&#39;, otherwise its readOnly. | 
**SumNetAccounting** | Pointer to **string** | Net accounting sum. Is equal to sumNet. | [optional] [readonly] 
**SumTaxAccounting** | Pointer to **string** | Tax accounting sum. Is equal to sumTax. | [optional] [readonly] 
**SumGrossAccounting** | Pointer to **string** | Gross accounting sum. Is equal to sumGross. | [optional] [readonly] 
**Comment** | Pointer to **NullableString** | Comment for the voucher position. | [optional] 

## Methods

### NewModelVoucherPosResponse

`func NewModelVoucherPosResponse(voucher ModelVoucherPosResponseVoucher, accountingType ModelVoucherPosResponseAccountingType, taxRate string, net bool, sumNet string, sumGross string, ) *ModelVoucherPosResponse`

NewModelVoucherPosResponse instantiates a new ModelVoucherPosResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelVoucherPosResponseWithDefaults

`func NewModelVoucherPosResponseWithDefaults() *ModelVoucherPosResponse`

NewModelVoucherPosResponseWithDefaults instantiates a new ModelVoucherPosResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelVoucherPosResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelVoucherPosResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelVoucherPosResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelVoucherPosResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectName

`func (o *ModelVoucherPosResponse) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ModelVoucherPosResponse) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ModelVoucherPosResponse) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ModelVoucherPosResponse) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetCreate

`func (o *ModelVoucherPosResponse) GetCreate() string`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ModelVoucherPosResponse) GetCreateOk() (*string, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ModelVoucherPosResponse) SetCreate(v string)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ModelVoucherPosResponse) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUpdate

`func (o *ModelVoucherPosResponse) GetUpdate() string`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ModelVoucherPosResponse) GetUpdateOk() (*string, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ModelVoucherPosResponse) SetUpdate(v string)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ModelVoucherPosResponse) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetSevClient

`func (o *ModelVoucherPosResponse) GetSevClient() ModelVoucherPosResponseSevClient`

GetSevClient returns the SevClient field if non-nil, zero value otherwise.

### GetSevClientOk

`func (o *ModelVoucherPosResponse) GetSevClientOk() (*ModelVoucherPosResponseSevClient, bool)`

GetSevClientOk returns a tuple with the SevClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSevClient

`func (o *ModelVoucherPosResponse) SetSevClient(v ModelVoucherPosResponseSevClient)`

SetSevClient sets SevClient field to given value.

### HasSevClient

`func (o *ModelVoucherPosResponse) HasSevClient() bool`

HasSevClient returns a boolean if a field has been set.

### GetVoucher

`func (o *ModelVoucherPosResponse) GetVoucher() ModelVoucherPosResponseVoucher`

GetVoucher returns the Voucher field if non-nil, zero value otherwise.

### GetVoucherOk

`func (o *ModelVoucherPosResponse) GetVoucherOk() (*ModelVoucherPosResponseVoucher, bool)`

GetVoucherOk returns a tuple with the Voucher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucher

`func (o *ModelVoucherPosResponse) SetVoucher(v ModelVoucherPosResponseVoucher)`

SetVoucher sets Voucher field to given value.


### GetAccountingType

`func (o *ModelVoucherPosResponse) GetAccountingType() ModelVoucherPosResponseAccountingType`

GetAccountingType returns the AccountingType field if non-nil, zero value otherwise.

### GetAccountingTypeOk

`func (o *ModelVoucherPosResponse) GetAccountingTypeOk() (*ModelVoucherPosResponseAccountingType, bool)`

GetAccountingTypeOk returns a tuple with the AccountingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingType

`func (o *ModelVoucherPosResponse) SetAccountingType(v ModelVoucherPosResponseAccountingType)`

SetAccountingType sets AccountingType field to given value.


### GetEstimatedAccountingType

`func (o *ModelVoucherPosResponse) GetEstimatedAccountingType() ModelVoucherPosResponseEstimatedAccountingType`

GetEstimatedAccountingType returns the EstimatedAccountingType field if non-nil, zero value otherwise.

### GetEstimatedAccountingTypeOk

`func (o *ModelVoucherPosResponse) GetEstimatedAccountingTypeOk() (*ModelVoucherPosResponseEstimatedAccountingType, bool)`

GetEstimatedAccountingTypeOk returns a tuple with the EstimatedAccountingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedAccountingType

`func (o *ModelVoucherPosResponse) SetEstimatedAccountingType(v ModelVoucherPosResponseEstimatedAccountingType)`

SetEstimatedAccountingType sets EstimatedAccountingType field to given value.

### HasEstimatedAccountingType

`func (o *ModelVoucherPosResponse) HasEstimatedAccountingType() bool`

HasEstimatedAccountingType returns a boolean if a field has been set.

### GetTaxRate

`func (o *ModelVoucherPosResponse) GetTaxRate() string`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *ModelVoucherPosResponse) GetTaxRateOk() (*string, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *ModelVoucherPosResponse) SetTaxRate(v string)`

SetTaxRate sets TaxRate field to given value.


### GetNet

`func (o *ModelVoucherPosResponse) GetNet() bool`

GetNet returns the Net field if non-nil, zero value otherwise.

### GetNetOk

`func (o *ModelVoucherPosResponse) GetNetOk() (*bool, bool)`

GetNetOk returns a tuple with the Net field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet

`func (o *ModelVoucherPosResponse) SetNet(v bool)`

SetNet sets Net field to given value.


### GetIsAsset

`func (o *ModelVoucherPosResponse) GetIsAsset() bool`

GetIsAsset returns the IsAsset field if non-nil, zero value otherwise.

### GetIsAssetOk

`func (o *ModelVoucherPosResponse) GetIsAssetOk() (*bool, bool)`

GetIsAssetOk returns a tuple with the IsAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAsset

`func (o *ModelVoucherPosResponse) SetIsAsset(v bool)`

SetIsAsset sets IsAsset field to given value.

### HasIsAsset

`func (o *ModelVoucherPosResponse) HasIsAsset() bool`

HasIsAsset returns a boolean if a field has been set.

### GetSumNet

`func (o *ModelVoucherPosResponse) GetSumNet() string`

GetSumNet returns the SumNet field if non-nil, zero value otherwise.

### GetSumNetOk

`func (o *ModelVoucherPosResponse) GetSumNetOk() (*string, bool)`

GetSumNetOk returns a tuple with the SumNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumNet

`func (o *ModelVoucherPosResponse) SetSumNet(v string)`

SetSumNet sets SumNet field to given value.


### GetSumTax

`func (o *ModelVoucherPosResponse) GetSumTax() string`

GetSumTax returns the SumTax field if non-nil, zero value otherwise.

### GetSumTaxOk

`func (o *ModelVoucherPosResponse) GetSumTaxOk() (*string, bool)`

GetSumTaxOk returns a tuple with the SumTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumTax

`func (o *ModelVoucherPosResponse) SetSumTax(v string)`

SetSumTax sets SumTax field to given value.

### HasSumTax

`func (o *ModelVoucherPosResponse) HasSumTax() bool`

HasSumTax returns a boolean if a field has been set.

### GetSumGross

`func (o *ModelVoucherPosResponse) GetSumGross() string`

GetSumGross returns the SumGross field if non-nil, zero value otherwise.

### GetSumGrossOk

`func (o *ModelVoucherPosResponse) GetSumGrossOk() (*string, bool)`

GetSumGrossOk returns a tuple with the SumGross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumGross

`func (o *ModelVoucherPosResponse) SetSumGross(v string)`

SetSumGross sets SumGross field to given value.


### GetSumNetAccounting

`func (o *ModelVoucherPosResponse) GetSumNetAccounting() string`

GetSumNetAccounting returns the SumNetAccounting field if non-nil, zero value otherwise.

### GetSumNetAccountingOk

`func (o *ModelVoucherPosResponse) GetSumNetAccountingOk() (*string, bool)`

GetSumNetAccountingOk returns a tuple with the SumNetAccounting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumNetAccounting

`func (o *ModelVoucherPosResponse) SetSumNetAccounting(v string)`

SetSumNetAccounting sets SumNetAccounting field to given value.

### HasSumNetAccounting

`func (o *ModelVoucherPosResponse) HasSumNetAccounting() bool`

HasSumNetAccounting returns a boolean if a field has been set.

### GetSumTaxAccounting

`func (o *ModelVoucherPosResponse) GetSumTaxAccounting() string`

GetSumTaxAccounting returns the SumTaxAccounting field if non-nil, zero value otherwise.

### GetSumTaxAccountingOk

`func (o *ModelVoucherPosResponse) GetSumTaxAccountingOk() (*string, bool)`

GetSumTaxAccountingOk returns a tuple with the SumTaxAccounting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumTaxAccounting

`func (o *ModelVoucherPosResponse) SetSumTaxAccounting(v string)`

SetSumTaxAccounting sets SumTaxAccounting field to given value.

### HasSumTaxAccounting

`func (o *ModelVoucherPosResponse) HasSumTaxAccounting() bool`

HasSumTaxAccounting returns a boolean if a field has been set.

### GetSumGrossAccounting

`func (o *ModelVoucherPosResponse) GetSumGrossAccounting() string`

GetSumGrossAccounting returns the SumGrossAccounting field if non-nil, zero value otherwise.

### GetSumGrossAccountingOk

`func (o *ModelVoucherPosResponse) GetSumGrossAccountingOk() (*string, bool)`

GetSumGrossAccountingOk returns a tuple with the SumGrossAccounting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumGrossAccounting

`func (o *ModelVoucherPosResponse) SetSumGrossAccounting(v string)`

SetSumGrossAccounting sets SumGrossAccounting field to given value.

### HasSumGrossAccounting

`func (o *ModelVoucherPosResponse) HasSumGrossAccounting() bool`

HasSumGrossAccounting returns a boolean if a field has been set.

### GetComment

`func (o *ModelVoucherPosResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ModelVoucherPosResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ModelVoucherPosResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ModelVoucherPosResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *ModelVoucherPosResponse) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *ModelVoucherPosResponse) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


