# ModelOrderPosResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The order position id | [optional] [readonly] 
**ObjectName** | Pointer to **string** | The order position object name | [optional] [readonly] 
**Create** | Pointer to **time.Time** | Date of order position creation | [optional] [readonly] 
**Update** | Pointer to **time.Time** | Date of last order position update | [optional] [readonly] 
**Order** | Pointer to [**ModelOrderPosResponseOrder**](ModelOrderPosResponseOrder.md) |  | [optional] 
**Part** | Pointer to [**ModelCreditNotePosResponsePart**](ModelCreditNotePosResponsePart.md) |  | [optional] 
**Quantity** | Pointer to **string** | Quantity of the article/part | [optional] 
**Price** | Pointer to **NullableString** | Price of the article/part. Is either gross or net, depending on the sevDesk account setting. | [optional] 
**PriceNet** | Pointer to **NullableString** | Net price of the part | [optional] [readonly] 
**PriceTax** | Pointer to **NullableString** | Tax on the price of the part | [optional] 
**PriceGross** | Pointer to **NullableString** | Gross price of the part | [optional] 
**Name** | Pointer to **NullableString** | Name of the article/part. | [optional] 
**Unity** | Pointer to [**ModelCreditNotePosResponseUnity**](ModelCreditNotePosResponseUnity.md) |  | [optional] 
**SevClient** | Pointer to [**ModelOrderPosResponseSevClient**](ModelOrderPosResponseSevClient.md) |  | [optional] 
**PositionNumber** | Pointer to **NullableString** | Position number of your position. Can be used to order multiple positions. | [optional] 
**Text** | Pointer to **NullableString** | A text describing your position. | [optional] 
**Discount** | Pointer to **NullableString** | An optional discount of the position. | [optional] 
**Optional** | Pointer to **NullableBool** | Defines if the position is optional. | [optional] 
**TaxRate** | Pointer to **string** | Tax rate of the position. | [optional] 
**SumDiscount** | Pointer to **NullableString** | Discount sum of the position | [optional] [readonly] 

## Methods

### NewModelOrderPosResponse

`func NewModelOrderPosResponse() *ModelOrderPosResponse`

NewModelOrderPosResponse instantiates a new ModelOrderPosResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelOrderPosResponseWithDefaults

`func NewModelOrderPosResponseWithDefaults() *ModelOrderPosResponse`

NewModelOrderPosResponseWithDefaults instantiates a new ModelOrderPosResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelOrderPosResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelOrderPosResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelOrderPosResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelOrderPosResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectName

`func (o *ModelOrderPosResponse) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ModelOrderPosResponse) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ModelOrderPosResponse) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ModelOrderPosResponse) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetCreate

`func (o *ModelOrderPosResponse) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ModelOrderPosResponse) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ModelOrderPosResponse) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ModelOrderPosResponse) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUpdate

`func (o *ModelOrderPosResponse) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ModelOrderPosResponse) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ModelOrderPosResponse) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ModelOrderPosResponse) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetOrder

`func (o *ModelOrderPosResponse) GetOrder() ModelOrderPosResponseOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ModelOrderPosResponse) GetOrderOk() (*ModelOrderPosResponseOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ModelOrderPosResponse) SetOrder(v ModelOrderPosResponseOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ModelOrderPosResponse) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPart

`func (o *ModelOrderPosResponse) GetPart() ModelCreditNotePosResponsePart`

GetPart returns the Part field if non-nil, zero value otherwise.

### GetPartOk

`func (o *ModelOrderPosResponse) GetPartOk() (*ModelCreditNotePosResponsePart, bool)`

GetPartOk returns a tuple with the Part field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPart

`func (o *ModelOrderPosResponse) SetPart(v ModelCreditNotePosResponsePart)`

SetPart sets Part field to given value.

### HasPart

`func (o *ModelOrderPosResponse) HasPart() bool`

HasPart returns a boolean if a field has been set.

### GetQuantity

`func (o *ModelOrderPosResponse) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ModelOrderPosResponse) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ModelOrderPosResponse) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ModelOrderPosResponse) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetPrice

`func (o *ModelOrderPosResponse) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ModelOrderPosResponse) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ModelOrderPosResponse) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ModelOrderPosResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *ModelOrderPosResponse) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *ModelOrderPosResponse) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetPriceNet

`func (o *ModelOrderPosResponse) GetPriceNet() string`

GetPriceNet returns the PriceNet field if non-nil, zero value otherwise.

### GetPriceNetOk

`func (o *ModelOrderPosResponse) GetPriceNetOk() (*string, bool)`

GetPriceNetOk returns a tuple with the PriceNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceNet

`func (o *ModelOrderPosResponse) SetPriceNet(v string)`

SetPriceNet sets PriceNet field to given value.

### HasPriceNet

`func (o *ModelOrderPosResponse) HasPriceNet() bool`

HasPriceNet returns a boolean if a field has been set.

### SetPriceNetNil

`func (o *ModelOrderPosResponse) SetPriceNetNil(b bool)`

 SetPriceNetNil sets the value for PriceNet to be an explicit nil

### UnsetPriceNet
`func (o *ModelOrderPosResponse) UnsetPriceNet()`

UnsetPriceNet ensures that no value is present for PriceNet, not even an explicit nil
### GetPriceTax

`func (o *ModelOrderPosResponse) GetPriceTax() string`

GetPriceTax returns the PriceTax field if non-nil, zero value otherwise.

### GetPriceTaxOk

`func (o *ModelOrderPosResponse) GetPriceTaxOk() (*string, bool)`

GetPriceTaxOk returns a tuple with the PriceTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceTax

`func (o *ModelOrderPosResponse) SetPriceTax(v string)`

SetPriceTax sets PriceTax field to given value.

### HasPriceTax

`func (o *ModelOrderPosResponse) HasPriceTax() bool`

HasPriceTax returns a boolean if a field has been set.

### SetPriceTaxNil

`func (o *ModelOrderPosResponse) SetPriceTaxNil(b bool)`

 SetPriceTaxNil sets the value for PriceTax to be an explicit nil

### UnsetPriceTax
`func (o *ModelOrderPosResponse) UnsetPriceTax()`

UnsetPriceTax ensures that no value is present for PriceTax, not even an explicit nil
### GetPriceGross

`func (o *ModelOrderPosResponse) GetPriceGross() string`

GetPriceGross returns the PriceGross field if non-nil, zero value otherwise.

### GetPriceGrossOk

`func (o *ModelOrderPosResponse) GetPriceGrossOk() (*string, bool)`

GetPriceGrossOk returns a tuple with the PriceGross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceGross

`func (o *ModelOrderPosResponse) SetPriceGross(v string)`

SetPriceGross sets PriceGross field to given value.

### HasPriceGross

`func (o *ModelOrderPosResponse) HasPriceGross() bool`

HasPriceGross returns a boolean if a field has been set.

### SetPriceGrossNil

`func (o *ModelOrderPosResponse) SetPriceGrossNil(b bool)`

 SetPriceGrossNil sets the value for PriceGross to be an explicit nil

### UnsetPriceGross
`func (o *ModelOrderPosResponse) UnsetPriceGross()`

UnsetPriceGross ensures that no value is present for PriceGross, not even an explicit nil
### GetName

`func (o *ModelOrderPosResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelOrderPosResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelOrderPosResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelOrderPosResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ModelOrderPosResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ModelOrderPosResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUnity

`func (o *ModelOrderPosResponse) GetUnity() ModelCreditNotePosResponseUnity`

GetUnity returns the Unity field if non-nil, zero value otherwise.

### GetUnityOk

`func (o *ModelOrderPosResponse) GetUnityOk() (*ModelCreditNotePosResponseUnity, bool)`

GetUnityOk returns a tuple with the Unity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnity

`func (o *ModelOrderPosResponse) SetUnity(v ModelCreditNotePosResponseUnity)`

SetUnity sets Unity field to given value.

### HasUnity

`func (o *ModelOrderPosResponse) HasUnity() bool`

HasUnity returns a boolean if a field has been set.

### GetSevClient

`func (o *ModelOrderPosResponse) GetSevClient() ModelOrderPosResponseSevClient`

GetSevClient returns the SevClient field if non-nil, zero value otherwise.

### GetSevClientOk

`func (o *ModelOrderPosResponse) GetSevClientOk() (*ModelOrderPosResponseSevClient, bool)`

GetSevClientOk returns a tuple with the SevClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSevClient

`func (o *ModelOrderPosResponse) SetSevClient(v ModelOrderPosResponseSevClient)`

SetSevClient sets SevClient field to given value.

### HasSevClient

`func (o *ModelOrderPosResponse) HasSevClient() bool`

HasSevClient returns a boolean if a field has been set.

### GetPositionNumber

`func (o *ModelOrderPosResponse) GetPositionNumber() string`

GetPositionNumber returns the PositionNumber field if non-nil, zero value otherwise.

### GetPositionNumberOk

`func (o *ModelOrderPosResponse) GetPositionNumberOk() (*string, bool)`

GetPositionNumberOk returns a tuple with the PositionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionNumber

`func (o *ModelOrderPosResponse) SetPositionNumber(v string)`

SetPositionNumber sets PositionNumber field to given value.

### HasPositionNumber

`func (o *ModelOrderPosResponse) HasPositionNumber() bool`

HasPositionNumber returns a boolean if a field has been set.

### SetPositionNumberNil

`func (o *ModelOrderPosResponse) SetPositionNumberNil(b bool)`

 SetPositionNumberNil sets the value for PositionNumber to be an explicit nil

### UnsetPositionNumber
`func (o *ModelOrderPosResponse) UnsetPositionNumber()`

UnsetPositionNumber ensures that no value is present for PositionNumber, not even an explicit nil
### GetText

`func (o *ModelOrderPosResponse) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ModelOrderPosResponse) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ModelOrderPosResponse) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ModelOrderPosResponse) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *ModelOrderPosResponse) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *ModelOrderPosResponse) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetDiscount

`func (o *ModelOrderPosResponse) GetDiscount() string`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *ModelOrderPosResponse) GetDiscountOk() (*string, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *ModelOrderPosResponse) SetDiscount(v string)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *ModelOrderPosResponse) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### SetDiscountNil

`func (o *ModelOrderPosResponse) SetDiscountNil(b bool)`

 SetDiscountNil sets the value for Discount to be an explicit nil

### UnsetDiscount
`func (o *ModelOrderPosResponse) UnsetDiscount()`

UnsetDiscount ensures that no value is present for Discount, not even an explicit nil
### GetOptional

`func (o *ModelOrderPosResponse) GetOptional() bool`

GetOptional returns the Optional field if non-nil, zero value otherwise.

### GetOptionalOk

`func (o *ModelOrderPosResponse) GetOptionalOk() (*bool, bool)`

GetOptionalOk returns a tuple with the Optional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptional

`func (o *ModelOrderPosResponse) SetOptional(v bool)`

SetOptional sets Optional field to given value.

### HasOptional

`func (o *ModelOrderPosResponse) HasOptional() bool`

HasOptional returns a boolean if a field has been set.

### SetOptionalNil

`func (o *ModelOrderPosResponse) SetOptionalNil(b bool)`

 SetOptionalNil sets the value for Optional to be an explicit nil

### UnsetOptional
`func (o *ModelOrderPosResponse) UnsetOptional()`

UnsetOptional ensures that no value is present for Optional, not even an explicit nil
### GetTaxRate

`func (o *ModelOrderPosResponse) GetTaxRate() string`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *ModelOrderPosResponse) GetTaxRateOk() (*string, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *ModelOrderPosResponse) SetTaxRate(v string)`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *ModelOrderPosResponse) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### GetSumDiscount

`func (o *ModelOrderPosResponse) GetSumDiscount() string`

GetSumDiscount returns the SumDiscount field if non-nil, zero value otherwise.

### GetSumDiscountOk

`func (o *ModelOrderPosResponse) GetSumDiscountOk() (*string, bool)`

GetSumDiscountOk returns a tuple with the SumDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumDiscount

`func (o *ModelOrderPosResponse) SetSumDiscount(v string)`

SetSumDiscount sets SumDiscount field to given value.

### HasSumDiscount

`func (o *ModelOrderPosResponse) HasSumDiscount() bool`

HasSumDiscount returns a boolean if a field has been set.

### SetSumDiscountNil

`func (o *ModelOrderPosResponse) SetSumDiscountNil(b bool)`

 SetSumDiscountNil sets the value for SumDiscount to be an explicit nil

### UnsetSumDiscount
`func (o *ModelOrderPosResponse) UnsetSumDiscount()`

UnsetSumDiscount ensures that no value is present for SumDiscount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


