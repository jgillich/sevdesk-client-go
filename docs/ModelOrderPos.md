# ModelOrderPos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The order position id | [optional] [readonly] 
**ObjectName** | Pointer to **string** | The order position object name | [optional] [readonly] 
**Create** | Pointer to **string** | Date of order position creation | [optional] [readonly] 
**Update** | Pointer to **string** | Date of last order position update | [optional] [readonly] 
**Order** | Pointer to [**ModelOrderPosOrder**](ModelOrderPosOrder.md) |  | [optional] 
**Part** | Pointer to [**ModelInvoicePosPart**](ModelInvoicePosPart.md) |  | [optional] 
**Quantity** | **float32** | Quantity of the article/part | 
**Price** | Pointer to **NullableFloat32** | Price of the article/part. Is either gross or net, depending on the sevDesk account setting. | [optional] 
**PriceNet** | Pointer to **NullableFloat32** | Net price of the part | [optional] [readonly] 
**PriceTax** | Pointer to **NullableFloat32** | Tax on the price of the part | [optional] 
**PriceGross** | Pointer to **NullableFloat32** | Gross price of the part | [optional] 
**Name** | Pointer to **NullableFloat32** | Name of the article/part. | [optional] 
**Unity** | [**ModelCreditNotePosUnity**](ModelCreditNotePosUnity.md) |  | 
**SevClient** | Pointer to [**ModelOrderPosSevClient**](ModelOrderPosSevClient.md) |  | [optional] 
**PositionNumber** | Pointer to **NullableInt32** | Position number of your position. Can be used to order multiple positions. | [optional] 
**Text** | Pointer to **NullableString** | A text describing your position. | [optional] 
**Discount** | Pointer to **NullableFloat32** | An optional discount of the position. | [optional] 
**Optional** | Pointer to **NullableBool** | Defines if the position is optional. | [optional] 
**TaxRate** | **float32** | Tax rate of the position. | 
**SumDiscount** | Pointer to **NullableFloat32** | Discount sum of the position | [optional] [readonly] 

## Methods

### NewModelOrderPos

`func NewModelOrderPos(quantity float32, unity ModelCreditNotePosUnity, taxRate float32, ) *ModelOrderPos`

NewModelOrderPos instantiates a new ModelOrderPos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelOrderPosWithDefaults

`func NewModelOrderPosWithDefaults() *ModelOrderPos`

NewModelOrderPosWithDefaults instantiates a new ModelOrderPos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelOrderPos) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelOrderPos) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelOrderPos) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelOrderPos) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectName

`func (o *ModelOrderPos) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ModelOrderPos) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ModelOrderPos) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ModelOrderPos) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetCreate

`func (o *ModelOrderPos) GetCreate() string`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ModelOrderPos) GetCreateOk() (*string, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ModelOrderPos) SetCreate(v string)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ModelOrderPos) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUpdate

`func (o *ModelOrderPos) GetUpdate() string`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ModelOrderPos) GetUpdateOk() (*string, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ModelOrderPos) SetUpdate(v string)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ModelOrderPos) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetOrder

`func (o *ModelOrderPos) GetOrder() ModelOrderPosOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ModelOrderPos) GetOrderOk() (*ModelOrderPosOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ModelOrderPos) SetOrder(v ModelOrderPosOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ModelOrderPos) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPart

`func (o *ModelOrderPos) GetPart() ModelInvoicePosPart`

GetPart returns the Part field if non-nil, zero value otherwise.

### GetPartOk

`func (o *ModelOrderPos) GetPartOk() (*ModelInvoicePosPart, bool)`

GetPartOk returns a tuple with the Part field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPart

`func (o *ModelOrderPos) SetPart(v ModelInvoicePosPart)`

SetPart sets Part field to given value.

### HasPart

`func (o *ModelOrderPos) HasPart() bool`

HasPart returns a boolean if a field has been set.

### GetQuantity

`func (o *ModelOrderPos) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ModelOrderPos) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ModelOrderPos) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetPrice

`func (o *ModelOrderPos) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ModelOrderPos) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ModelOrderPos) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ModelOrderPos) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *ModelOrderPos) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *ModelOrderPos) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetPriceNet

`func (o *ModelOrderPos) GetPriceNet() float32`

GetPriceNet returns the PriceNet field if non-nil, zero value otherwise.

### GetPriceNetOk

`func (o *ModelOrderPos) GetPriceNetOk() (*float32, bool)`

GetPriceNetOk returns a tuple with the PriceNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceNet

`func (o *ModelOrderPos) SetPriceNet(v float32)`

SetPriceNet sets PriceNet field to given value.

### HasPriceNet

`func (o *ModelOrderPos) HasPriceNet() bool`

HasPriceNet returns a boolean if a field has been set.

### SetPriceNetNil

`func (o *ModelOrderPos) SetPriceNetNil(b bool)`

 SetPriceNetNil sets the value for PriceNet to be an explicit nil

### UnsetPriceNet
`func (o *ModelOrderPos) UnsetPriceNet()`

UnsetPriceNet ensures that no value is present for PriceNet, not even an explicit nil
### GetPriceTax

`func (o *ModelOrderPos) GetPriceTax() float32`

GetPriceTax returns the PriceTax field if non-nil, zero value otherwise.

### GetPriceTaxOk

`func (o *ModelOrderPos) GetPriceTaxOk() (*float32, bool)`

GetPriceTaxOk returns a tuple with the PriceTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceTax

`func (o *ModelOrderPos) SetPriceTax(v float32)`

SetPriceTax sets PriceTax field to given value.

### HasPriceTax

`func (o *ModelOrderPos) HasPriceTax() bool`

HasPriceTax returns a boolean if a field has been set.

### SetPriceTaxNil

`func (o *ModelOrderPos) SetPriceTaxNil(b bool)`

 SetPriceTaxNil sets the value for PriceTax to be an explicit nil

### UnsetPriceTax
`func (o *ModelOrderPos) UnsetPriceTax()`

UnsetPriceTax ensures that no value is present for PriceTax, not even an explicit nil
### GetPriceGross

`func (o *ModelOrderPos) GetPriceGross() float32`

GetPriceGross returns the PriceGross field if non-nil, zero value otherwise.

### GetPriceGrossOk

`func (o *ModelOrderPos) GetPriceGrossOk() (*float32, bool)`

GetPriceGrossOk returns a tuple with the PriceGross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceGross

`func (o *ModelOrderPos) SetPriceGross(v float32)`

SetPriceGross sets PriceGross field to given value.

### HasPriceGross

`func (o *ModelOrderPos) HasPriceGross() bool`

HasPriceGross returns a boolean if a field has been set.

### SetPriceGrossNil

`func (o *ModelOrderPos) SetPriceGrossNil(b bool)`

 SetPriceGrossNil sets the value for PriceGross to be an explicit nil

### UnsetPriceGross
`func (o *ModelOrderPos) UnsetPriceGross()`

UnsetPriceGross ensures that no value is present for PriceGross, not even an explicit nil
### GetName

`func (o *ModelOrderPos) GetName() float32`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelOrderPos) GetNameOk() (*float32, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelOrderPos) SetName(v float32)`

SetName sets Name field to given value.

### HasName

`func (o *ModelOrderPos) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ModelOrderPos) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ModelOrderPos) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUnity

`func (o *ModelOrderPos) GetUnity() ModelCreditNotePosUnity`

GetUnity returns the Unity field if non-nil, zero value otherwise.

### GetUnityOk

`func (o *ModelOrderPos) GetUnityOk() (*ModelCreditNotePosUnity, bool)`

GetUnityOk returns a tuple with the Unity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnity

`func (o *ModelOrderPos) SetUnity(v ModelCreditNotePosUnity)`

SetUnity sets Unity field to given value.


### GetSevClient

`func (o *ModelOrderPos) GetSevClient() ModelOrderPosSevClient`

GetSevClient returns the SevClient field if non-nil, zero value otherwise.

### GetSevClientOk

`func (o *ModelOrderPos) GetSevClientOk() (*ModelOrderPosSevClient, bool)`

GetSevClientOk returns a tuple with the SevClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSevClient

`func (o *ModelOrderPos) SetSevClient(v ModelOrderPosSevClient)`

SetSevClient sets SevClient field to given value.

### HasSevClient

`func (o *ModelOrderPos) HasSevClient() bool`

HasSevClient returns a boolean if a field has been set.

### GetPositionNumber

`func (o *ModelOrderPos) GetPositionNumber() int32`

GetPositionNumber returns the PositionNumber field if non-nil, zero value otherwise.

### GetPositionNumberOk

`func (o *ModelOrderPos) GetPositionNumberOk() (*int32, bool)`

GetPositionNumberOk returns a tuple with the PositionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionNumber

`func (o *ModelOrderPos) SetPositionNumber(v int32)`

SetPositionNumber sets PositionNumber field to given value.

### HasPositionNumber

`func (o *ModelOrderPos) HasPositionNumber() bool`

HasPositionNumber returns a boolean if a field has been set.

### SetPositionNumberNil

`func (o *ModelOrderPos) SetPositionNumberNil(b bool)`

 SetPositionNumberNil sets the value for PositionNumber to be an explicit nil

### UnsetPositionNumber
`func (o *ModelOrderPos) UnsetPositionNumber()`

UnsetPositionNumber ensures that no value is present for PositionNumber, not even an explicit nil
### GetText

`func (o *ModelOrderPos) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ModelOrderPos) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ModelOrderPos) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ModelOrderPos) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *ModelOrderPos) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *ModelOrderPos) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetDiscount

`func (o *ModelOrderPos) GetDiscount() float32`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *ModelOrderPos) GetDiscountOk() (*float32, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *ModelOrderPos) SetDiscount(v float32)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *ModelOrderPos) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### SetDiscountNil

`func (o *ModelOrderPos) SetDiscountNil(b bool)`

 SetDiscountNil sets the value for Discount to be an explicit nil

### UnsetDiscount
`func (o *ModelOrderPos) UnsetDiscount()`

UnsetDiscount ensures that no value is present for Discount, not even an explicit nil
### GetOptional

`func (o *ModelOrderPos) GetOptional() bool`

GetOptional returns the Optional field if non-nil, zero value otherwise.

### GetOptionalOk

`func (o *ModelOrderPos) GetOptionalOk() (*bool, bool)`

GetOptionalOk returns a tuple with the Optional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptional

`func (o *ModelOrderPos) SetOptional(v bool)`

SetOptional sets Optional field to given value.

### HasOptional

`func (o *ModelOrderPos) HasOptional() bool`

HasOptional returns a boolean if a field has been set.

### SetOptionalNil

`func (o *ModelOrderPos) SetOptionalNil(b bool)`

 SetOptionalNil sets the value for Optional to be an explicit nil

### UnsetOptional
`func (o *ModelOrderPos) UnsetOptional()`

UnsetOptional ensures that no value is present for Optional, not even an explicit nil
### GetTaxRate

`func (o *ModelOrderPos) GetTaxRate() float32`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *ModelOrderPos) GetTaxRateOk() (*float32, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *ModelOrderPos) SetTaxRate(v float32)`

SetTaxRate sets TaxRate field to given value.


### GetSumDiscount

`func (o *ModelOrderPos) GetSumDiscount() float32`

GetSumDiscount returns the SumDiscount field if non-nil, zero value otherwise.

### GetSumDiscountOk

`func (o *ModelOrderPos) GetSumDiscountOk() (*float32, bool)`

GetSumDiscountOk returns a tuple with the SumDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumDiscount

`func (o *ModelOrderPos) SetSumDiscount(v float32)`

SetSumDiscount sets SumDiscount field to given value.

### HasSumDiscount

`func (o *ModelOrderPos) HasSumDiscount() bool`

HasSumDiscount returns a boolean if a field has been set.

### SetSumDiscountNil

`func (o *ModelOrderPos) SetSumDiscountNil(b bool)`

 SetSumDiscountNil sets the value for SumDiscount to be an explicit nil

### UnsetSumDiscount
`func (o *ModelOrderPos) UnsetSumDiscount()`

UnsetSumDiscount ensures that no value is present for SumDiscount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


