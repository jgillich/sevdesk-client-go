# ModelOrderPosUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The order position id | [optional] [readonly] 
**ObjectName** | Pointer to **string** | The order position object name | [optional] [readonly] 
**Create** | Pointer to **time.Time** | Date of order position creation | [optional] [readonly] 
**Update** | Pointer to **time.Time** | Date of last order position update | [optional] [readonly] 
**Order** | Pointer to [**ModelOrderPosOrder**](ModelOrderPosOrder.md) |  | [optional] 
**Part** | Pointer to [**ModelInvoicePosPart**](ModelInvoicePosPart.md) |  | [optional] 
**Quantity** | Pointer to **NullableFloat32** | Quantity of the article/part | [optional] 
**Price** | Pointer to **NullableFloat32** | Price of the article/part. Is either gross or net, depending on the sevDesk account setting. | [optional] 
**PriceNet** | Pointer to **NullableFloat32** | Net price of the part | [optional] [readonly] 
**PriceTax** | Pointer to **NullableFloat32** | Tax on the price of the part | [optional] 
**PriceGross** | Pointer to **NullableFloat32** | Gross price of the part | [optional] 
**Name** | Pointer to **NullableFloat32** | Name of the article/part. | [optional] 
**Unity** | Pointer to [**ModelCreditNotePosUnity**](ModelCreditNotePosUnity.md) |  | [optional] 
**SevClient** | Pointer to [**ModelOrderPosSevClient**](ModelOrderPosSevClient.md) |  | [optional] 
**PositionNumber** | Pointer to **NullableInt32** | Position number of your position. Can be used to order multiple positions. | [optional] 
**Text** | Pointer to **NullableString** | A text describing your position. | [optional] 
**Discount** | Pointer to **NullableFloat32** | An optional discount of the position. | [optional] 
**Optional** | Pointer to **NullableBool** | Defines if the position is optional. | [optional] 
**TaxRate** | Pointer to **NullableFloat32** | Tax rate of the position. | [optional] 
**SumDiscount** | Pointer to **NullableFloat32** | Discount sum of the position | [optional] [readonly] 

## Methods

### NewModelOrderPosUpdate

`func NewModelOrderPosUpdate() *ModelOrderPosUpdate`

NewModelOrderPosUpdate instantiates a new ModelOrderPosUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelOrderPosUpdateWithDefaults

`func NewModelOrderPosUpdateWithDefaults() *ModelOrderPosUpdate`

NewModelOrderPosUpdateWithDefaults instantiates a new ModelOrderPosUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelOrderPosUpdate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelOrderPosUpdate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelOrderPosUpdate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelOrderPosUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectName

`func (o *ModelOrderPosUpdate) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ModelOrderPosUpdate) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ModelOrderPosUpdate) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ModelOrderPosUpdate) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetCreate

`func (o *ModelOrderPosUpdate) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ModelOrderPosUpdate) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ModelOrderPosUpdate) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ModelOrderPosUpdate) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUpdate

`func (o *ModelOrderPosUpdate) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ModelOrderPosUpdate) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ModelOrderPosUpdate) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ModelOrderPosUpdate) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetOrder

`func (o *ModelOrderPosUpdate) GetOrder() ModelOrderPosOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ModelOrderPosUpdate) GetOrderOk() (*ModelOrderPosOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ModelOrderPosUpdate) SetOrder(v ModelOrderPosOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ModelOrderPosUpdate) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPart

`func (o *ModelOrderPosUpdate) GetPart() ModelInvoicePosPart`

GetPart returns the Part field if non-nil, zero value otherwise.

### GetPartOk

`func (o *ModelOrderPosUpdate) GetPartOk() (*ModelInvoicePosPart, bool)`

GetPartOk returns a tuple with the Part field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPart

`func (o *ModelOrderPosUpdate) SetPart(v ModelInvoicePosPart)`

SetPart sets Part field to given value.

### HasPart

`func (o *ModelOrderPosUpdate) HasPart() bool`

HasPart returns a boolean if a field has been set.

### GetQuantity

`func (o *ModelOrderPosUpdate) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ModelOrderPosUpdate) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ModelOrderPosUpdate) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ModelOrderPosUpdate) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *ModelOrderPosUpdate) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *ModelOrderPosUpdate) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetPrice

`func (o *ModelOrderPosUpdate) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ModelOrderPosUpdate) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ModelOrderPosUpdate) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ModelOrderPosUpdate) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *ModelOrderPosUpdate) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *ModelOrderPosUpdate) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetPriceNet

`func (o *ModelOrderPosUpdate) GetPriceNet() float32`

GetPriceNet returns the PriceNet field if non-nil, zero value otherwise.

### GetPriceNetOk

`func (o *ModelOrderPosUpdate) GetPriceNetOk() (*float32, bool)`

GetPriceNetOk returns a tuple with the PriceNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceNet

`func (o *ModelOrderPosUpdate) SetPriceNet(v float32)`

SetPriceNet sets PriceNet field to given value.

### HasPriceNet

`func (o *ModelOrderPosUpdate) HasPriceNet() bool`

HasPriceNet returns a boolean if a field has been set.

### SetPriceNetNil

`func (o *ModelOrderPosUpdate) SetPriceNetNil(b bool)`

 SetPriceNetNil sets the value for PriceNet to be an explicit nil

### UnsetPriceNet
`func (o *ModelOrderPosUpdate) UnsetPriceNet()`

UnsetPriceNet ensures that no value is present for PriceNet, not even an explicit nil
### GetPriceTax

`func (o *ModelOrderPosUpdate) GetPriceTax() float32`

GetPriceTax returns the PriceTax field if non-nil, zero value otherwise.

### GetPriceTaxOk

`func (o *ModelOrderPosUpdate) GetPriceTaxOk() (*float32, bool)`

GetPriceTaxOk returns a tuple with the PriceTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceTax

`func (o *ModelOrderPosUpdate) SetPriceTax(v float32)`

SetPriceTax sets PriceTax field to given value.

### HasPriceTax

`func (o *ModelOrderPosUpdate) HasPriceTax() bool`

HasPriceTax returns a boolean if a field has been set.

### SetPriceTaxNil

`func (o *ModelOrderPosUpdate) SetPriceTaxNil(b bool)`

 SetPriceTaxNil sets the value for PriceTax to be an explicit nil

### UnsetPriceTax
`func (o *ModelOrderPosUpdate) UnsetPriceTax()`

UnsetPriceTax ensures that no value is present for PriceTax, not even an explicit nil
### GetPriceGross

`func (o *ModelOrderPosUpdate) GetPriceGross() float32`

GetPriceGross returns the PriceGross field if non-nil, zero value otherwise.

### GetPriceGrossOk

`func (o *ModelOrderPosUpdate) GetPriceGrossOk() (*float32, bool)`

GetPriceGrossOk returns a tuple with the PriceGross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceGross

`func (o *ModelOrderPosUpdate) SetPriceGross(v float32)`

SetPriceGross sets PriceGross field to given value.

### HasPriceGross

`func (o *ModelOrderPosUpdate) HasPriceGross() bool`

HasPriceGross returns a boolean if a field has been set.

### SetPriceGrossNil

`func (o *ModelOrderPosUpdate) SetPriceGrossNil(b bool)`

 SetPriceGrossNil sets the value for PriceGross to be an explicit nil

### UnsetPriceGross
`func (o *ModelOrderPosUpdate) UnsetPriceGross()`

UnsetPriceGross ensures that no value is present for PriceGross, not even an explicit nil
### GetName

`func (o *ModelOrderPosUpdate) GetName() float32`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelOrderPosUpdate) GetNameOk() (*float32, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelOrderPosUpdate) SetName(v float32)`

SetName sets Name field to given value.

### HasName

`func (o *ModelOrderPosUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ModelOrderPosUpdate) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ModelOrderPosUpdate) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUnity

`func (o *ModelOrderPosUpdate) GetUnity() ModelCreditNotePosUnity`

GetUnity returns the Unity field if non-nil, zero value otherwise.

### GetUnityOk

`func (o *ModelOrderPosUpdate) GetUnityOk() (*ModelCreditNotePosUnity, bool)`

GetUnityOk returns a tuple with the Unity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnity

`func (o *ModelOrderPosUpdate) SetUnity(v ModelCreditNotePosUnity)`

SetUnity sets Unity field to given value.

### HasUnity

`func (o *ModelOrderPosUpdate) HasUnity() bool`

HasUnity returns a boolean if a field has been set.

### GetSevClient

`func (o *ModelOrderPosUpdate) GetSevClient() ModelOrderPosSevClient`

GetSevClient returns the SevClient field if non-nil, zero value otherwise.

### GetSevClientOk

`func (o *ModelOrderPosUpdate) GetSevClientOk() (*ModelOrderPosSevClient, bool)`

GetSevClientOk returns a tuple with the SevClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSevClient

`func (o *ModelOrderPosUpdate) SetSevClient(v ModelOrderPosSevClient)`

SetSevClient sets SevClient field to given value.

### HasSevClient

`func (o *ModelOrderPosUpdate) HasSevClient() bool`

HasSevClient returns a boolean if a field has been set.

### GetPositionNumber

`func (o *ModelOrderPosUpdate) GetPositionNumber() int32`

GetPositionNumber returns the PositionNumber field if non-nil, zero value otherwise.

### GetPositionNumberOk

`func (o *ModelOrderPosUpdate) GetPositionNumberOk() (*int32, bool)`

GetPositionNumberOk returns a tuple with the PositionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionNumber

`func (o *ModelOrderPosUpdate) SetPositionNumber(v int32)`

SetPositionNumber sets PositionNumber field to given value.

### HasPositionNumber

`func (o *ModelOrderPosUpdate) HasPositionNumber() bool`

HasPositionNumber returns a boolean if a field has been set.

### SetPositionNumberNil

`func (o *ModelOrderPosUpdate) SetPositionNumberNil(b bool)`

 SetPositionNumberNil sets the value for PositionNumber to be an explicit nil

### UnsetPositionNumber
`func (o *ModelOrderPosUpdate) UnsetPositionNumber()`

UnsetPositionNumber ensures that no value is present for PositionNumber, not even an explicit nil
### GetText

`func (o *ModelOrderPosUpdate) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ModelOrderPosUpdate) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ModelOrderPosUpdate) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ModelOrderPosUpdate) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *ModelOrderPosUpdate) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *ModelOrderPosUpdate) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetDiscount

`func (o *ModelOrderPosUpdate) GetDiscount() float32`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *ModelOrderPosUpdate) GetDiscountOk() (*float32, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *ModelOrderPosUpdate) SetDiscount(v float32)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *ModelOrderPosUpdate) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### SetDiscountNil

`func (o *ModelOrderPosUpdate) SetDiscountNil(b bool)`

 SetDiscountNil sets the value for Discount to be an explicit nil

### UnsetDiscount
`func (o *ModelOrderPosUpdate) UnsetDiscount()`

UnsetDiscount ensures that no value is present for Discount, not even an explicit nil
### GetOptional

`func (o *ModelOrderPosUpdate) GetOptional() bool`

GetOptional returns the Optional field if non-nil, zero value otherwise.

### GetOptionalOk

`func (o *ModelOrderPosUpdate) GetOptionalOk() (*bool, bool)`

GetOptionalOk returns a tuple with the Optional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptional

`func (o *ModelOrderPosUpdate) SetOptional(v bool)`

SetOptional sets Optional field to given value.

### HasOptional

`func (o *ModelOrderPosUpdate) HasOptional() bool`

HasOptional returns a boolean if a field has been set.

### SetOptionalNil

`func (o *ModelOrderPosUpdate) SetOptionalNil(b bool)`

 SetOptionalNil sets the value for Optional to be an explicit nil

### UnsetOptional
`func (o *ModelOrderPosUpdate) UnsetOptional()`

UnsetOptional ensures that no value is present for Optional, not even an explicit nil
### GetTaxRate

`func (o *ModelOrderPosUpdate) GetTaxRate() float32`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *ModelOrderPosUpdate) GetTaxRateOk() (*float32, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *ModelOrderPosUpdate) SetTaxRate(v float32)`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *ModelOrderPosUpdate) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### SetTaxRateNil

`func (o *ModelOrderPosUpdate) SetTaxRateNil(b bool)`

 SetTaxRateNil sets the value for TaxRate to be an explicit nil

### UnsetTaxRate
`func (o *ModelOrderPosUpdate) UnsetTaxRate()`

UnsetTaxRate ensures that no value is present for TaxRate, not even an explicit nil
### GetSumDiscount

`func (o *ModelOrderPosUpdate) GetSumDiscount() float32`

GetSumDiscount returns the SumDiscount field if non-nil, zero value otherwise.

### GetSumDiscountOk

`func (o *ModelOrderPosUpdate) GetSumDiscountOk() (*float32, bool)`

GetSumDiscountOk returns a tuple with the SumDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumDiscount

`func (o *ModelOrderPosUpdate) SetSumDiscount(v float32)`

SetSumDiscount sets SumDiscount field to given value.

### HasSumDiscount

`func (o *ModelOrderPosUpdate) HasSumDiscount() bool`

HasSumDiscount returns a boolean if a field has been set.

### SetSumDiscountNil

`func (o *ModelOrderPosUpdate) SetSumDiscountNil(b bool)`

 SetSumDiscountNil sets the value for SumDiscount to be an explicit nil

### UnsetSumDiscount
`func (o *ModelOrderPosUpdate) UnsetSumDiscount()`

UnsetSumDiscount ensures that no value is present for SumDiscount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


