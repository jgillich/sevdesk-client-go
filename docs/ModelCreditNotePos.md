# ModelCreditNotePos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | The creditNote position id. | [optional] [readonly] 
**ObjectName** | **string** | The creditNote position object name | 
**MapAll** | **bool** |  | 
**Create** | Pointer to **string** | Date of creditNote position creation | [optional] [readonly] 
**Update** | Pointer to **string** | Date of last creditNote position update | [optional] [readonly] 
**CreditNote** | Pointer to [**ModelCreditNotePosCreditNote**](ModelCreditNotePosCreditNote.md) |  | [optional] 
**Part** | Pointer to [**ModelInvoicePosPart**](ModelInvoicePosPart.md) |  | [optional] 
**Quantity** | **float32** | Quantity of the article/part | 
**Price** | Pointer to **NullableFloat32** | Price of the article/part. Is either gross or net, depending on the sevDesk account setting. | [optional] 
**PriceNet** | Pointer to **NullableFloat32** | Net price of the part | [optional] [readonly] 
**PriceTax** | Pointer to **NullableFloat32** | Tax on the price of the part | [optional] 
**PriceGross** | Pointer to **NullableFloat32** | Gross price of the part | [optional] 
**Name** | Pointer to **NullableString** | Name of the article/part. | [optional] 
**Unity** | [**ModelCreditNotePosUnity**](ModelCreditNotePosUnity.md) |  | 
**SevClient** | Pointer to [**ModelCreditNotePosSevClient**](ModelCreditNotePosSevClient.md) |  | [optional] 
**PositionNumber** | Pointer to **NullableInt32** | Position number of your position. Can be used to creditNote multiple positions. | [optional] 
**Text** | Pointer to **NullableString** | A text describing your position. | [optional] 
**Discount** | Pointer to **NullableFloat32** | An optional discount of the position. | [optional] 
**Optional** | Pointer to **NullableBool** | Defines if the position is optional. | [optional] 
**TaxRate** | **float32** | Tax rate of the position. | 
**SumDiscount** | Pointer to **NullableFloat32** | Discount sum of the position | [optional] [readonly] 

## Methods

### NewModelCreditNotePos

`func NewModelCreditNotePos(objectName string, mapAll bool, quantity float32, unity ModelCreditNotePosUnity, taxRate float32, ) *ModelCreditNotePos`

NewModelCreditNotePos instantiates a new ModelCreditNotePos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelCreditNotePosWithDefaults

`func NewModelCreditNotePosWithDefaults() *ModelCreditNotePos`

NewModelCreditNotePosWithDefaults instantiates a new ModelCreditNotePos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelCreditNotePos) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelCreditNotePos) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelCreditNotePos) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelCreditNotePos) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ModelCreditNotePos) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ModelCreditNotePos) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectName

`func (o *ModelCreditNotePos) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ModelCreditNotePos) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ModelCreditNotePos) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.


### GetMapAll

`func (o *ModelCreditNotePos) GetMapAll() bool`

GetMapAll returns the MapAll field if non-nil, zero value otherwise.

### GetMapAllOk

`func (o *ModelCreditNotePos) GetMapAllOk() (*bool, bool)`

GetMapAllOk returns a tuple with the MapAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapAll

`func (o *ModelCreditNotePos) SetMapAll(v bool)`

SetMapAll sets MapAll field to given value.


### GetCreate

`func (o *ModelCreditNotePos) GetCreate() string`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ModelCreditNotePos) GetCreateOk() (*string, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ModelCreditNotePos) SetCreate(v string)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ModelCreditNotePos) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUpdate

`func (o *ModelCreditNotePos) GetUpdate() string`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ModelCreditNotePos) GetUpdateOk() (*string, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ModelCreditNotePos) SetUpdate(v string)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ModelCreditNotePos) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetCreditNote

`func (o *ModelCreditNotePos) GetCreditNote() ModelCreditNotePosCreditNote`

GetCreditNote returns the CreditNote field if non-nil, zero value otherwise.

### GetCreditNoteOk

`func (o *ModelCreditNotePos) GetCreditNoteOk() (*ModelCreditNotePosCreditNote, bool)`

GetCreditNoteOk returns a tuple with the CreditNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNote

`func (o *ModelCreditNotePos) SetCreditNote(v ModelCreditNotePosCreditNote)`

SetCreditNote sets CreditNote field to given value.

### HasCreditNote

`func (o *ModelCreditNotePos) HasCreditNote() bool`

HasCreditNote returns a boolean if a field has been set.

### GetPart

`func (o *ModelCreditNotePos) GetPart() ModelInvoicePosPart`

GetPart returns the Part field if non-nil, zero value otherwise.

### GetPartOk

`func (o *ModelCreditNotePos) GetPartOk() (*ModelInvoicePosPart, bool)`

GetPartOk returns a tuple with the Part field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPart

`func (o *ModelCreditNotePos) SetPart(v ModelInvoicePosPart)`

SetPart sets Part field to given value.

### HasPart

`func (o *ModelCreditNotePos) HasPart() bool`

HasPart returns a boolean if a field has been set.

### GetQuantity

`func (o *ModelCreditNotePos) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ModelCreditNotePos) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ModelCreditNotePos) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetPrice

`func (o *ModelCreditNotePos) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ModelCreditNotePos) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ModelCreditNotePos) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ModelCreditNotePos) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *ModelCreditNotePos) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *ModelCreditNotePos) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetPriceNet

`func (o *ModelCreditNotePos) GetPriceNet() float32`

GetPriceNet returns the PriceNet field if non-nil, zero value otherwise.

### GetPriceNetOk

`func (o *ModelCreditNotePos) GetPriceNetOk() (*float32, bool)`

GetPriceNetOk returns a tuple with the PriceNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceNet

`func (o *ModelCreditNotePos) SetPriceNet(v float32)`

SetPriceNet sets PriceNet field to given value.

### HasPriceNet

`func (o *ModelCreditNotePos) HasPriceNet() bool`

HasPriceNet returns a boolean if a field has been set.

### SetPriceNetNil

`func (o *ModelCreditNotePos) SetPriceNetNil(b bool)`

 SetPriceNetNil sets the value for PriceNet to be an explicit nil

### UnsetPriceNet
`func (o *ModelCreditNotePos) UnsetPriceNet()`

UnsetPriceNet ensures that no value is present for PriceNet, not even an explicit nil
### GetPriceTax

`func (o *ModelCreditNotePos) GetPriceTax() float32`

GetPriceTax returns the PriceTax field if non-nil, zero value otherwise.

### GetPriceTaxOk

`func (o *ModelCreditNotePos) GetPriceTaxOk() (*float32, bool)`

GetPriceTaxOk returns a tuple with the PriceTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceTax

`func (o *ModelCreditNotePos) SetPriceTax(v float32)`

SetPriceTax sets PriceTax field to given value.

### HasPriceTax

`func (o *ModelCreditNotePos) HasPriceTax() bool`

HasPriceTax returns a boolean if a field has been set.

### SetPriceTaxNil

`func (o *ModelCreditNotePos) SetPriceTaxNil(b bool)`

 SetPriceTaxNil sets the value for PriceTax to be an explicit nil

### UnsetPriceTax
`func (o *ModelCreditNotePos) UnsetPriceTax()`

UnsetPriceTax ensures that no value is present for PriceTax, not even an explicit nil
### GetPriceGross

`func (o *ModelCreditNotePos) GetPriceGross() float32`

GetPriceGross returns the PriceGross field if non-nil, zero value otherwise.

### GetPriceGrossOk

`func (o *ModelCreditNotePos) GetPriceGrossOk() (*float32, bool)`

GetPriceGrossOk returns a tuple with the PriceGross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceGross

`func (o *ModelCreditNotePos) SetPriceGross(v float32)`

SetPriceGross sets PriceGross field to given value.

### HasPriceGross

`func (o *ModelCreditNotePos) HasPriceGross() bool`

HasPriceGross returns a boolean if a field has been set.

### SetPriceGrossNil

`func (o *ModelCreditNotePos) SetPriceGrossNil(b bool)`

 SetPriceGrossNil sets the value for PriceGross to be an explicit nil

### UnsetPriceGross
`func (o *ModelCreditNotePos) UnsetPriceGross()`

UnsetPriceGross ensures that no value is present for PriceGross, not even an explicit nil
### GetName

`func (o *ModelCreditNotePos) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelCreditNotePos) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelCreditNotePos) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelCreditNotePos) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ModelCreditNotePos) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ModelCreditNotePos) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUnity

`func (o *ModelCreditNotePos) GetUnity() ModelCreditNotePosUnity`

GetUnity returns the Unity field if non-nil, zero value otherwise.

### GetUnityOk

`func (o *ModelCreditNotePos) GetUnityOk() (*ModelCreditNotePosUnity, bool)`

GetUnityOk returns a tuple with the Unity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnity

`func (o *ModelCreditNotePos) SetUnity(v ModelCreditNotePosUnity)`

SetUnity sets Unity field to given value.


### GetSevClient

`func (o *ModelCreditNotePos) GetSevClient() ModelCreditNotePosSevClient`

GetSevClient returns the SevClient field if non-nil, zero value otherwise.

### GetSevClientOk

`func (o *ModelCreditNotePos) GetSevClientOk() (*ModelCreditNotePosSevClient, bool)`

GetSevClientOk returns a tuple with the SevClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSevClient

`func (o *ModelCreditNotePos) SetSevClient(v ModelCreditNotePosSevClient)`

SetSevClient sets SevClient field to given value.

### HasSevClient

`func (o *ModelCreditNotePos) HasSevClient() bool`

HasSevClient returns a boolean if a field has been set.

### GetPositionNumber

`func (o *ModelCreditNotePos) GetPositionNumber() int32`

GetPositionNumber returns the PositionNumber field if non-nil, zero value otherwise.

### GetPositionNumberOk

`func (o *ModelCreditNotePos) GetPositionNumberOk() (*int32, bool)`

GetPositionNumberOk returns a tuple with the PositionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionNumber

`func (o *ModelCreditNotePos) SetPositionNumber(v int32)`

SetPositionNumber sets PositionNumber field to given value.

### HasPositionNumber

`func (o *ModelCreditNotePos) HasPositionNumber() bool`

HasPositionNumber returns a boolean if a field has been set.

### SetPositionNumberNil

`func (o *ModelCreditNotePos) SetPositionNumberNil(b bool)`

 SetPositionNumberNil sets the value for PositionNumber to be an explicit nil

### UnsetPositionNumber
`func (o *ModelCreditNotePos) UnsetPositionNumber()`

UnsetPositionNumber ensures that no value is present for PositionNumber, not even an explicit nil
### GetText

`func (o *ModelCreditNotePos) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ModelCreditNotePos) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ModelCreditNotePos) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ModelCreditNotePos) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *ModelCreditNotePos) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *ModelCreditNotePos) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetDiscount

`func (o *ModelCreditNotePos) GetDiscount() float32`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *ModelCreditNotePos) GetDiscountOk() (*float32, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *ModelCreditNotePos) SetDiscount(v float32)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *ModelCreditNotePos) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### SetDiscountNil

`func (o *ModelCreditNotePos) SetDiscountNil(b bool)`

 SetDiscountNil sets the value for Discount to be an explicit nil

### UnsetDiscount
`func (o *ModelCreditNotePos) UnsetDiscount()`

UnsetDiscount ensures that no value is present for Discount, not even an explicit nil
### GetOptional

`func (o *ModelCreditNotePos) GetOptional() bool`

GetOptional returns the Optional field if non-nil, zero value otherwise.

### GetOptionalOk

`func (o *ModelCreditNotePos) GetOptionalOk() (*bool, bool)`

GetOptionalOk returns a tuple with the Optional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptional

`func (o *ModelCreditNotePos) SetOptional(v bool)`

SetOptional sets Optional field to given value.

### HasOptional

`func (o *ModelCreditNotePos) HasOptional() bool`

HasOptional returns a boolean if a field has been set.

### SetOptionalNil

`func (o *ModelCreditNotePos) SetOptionalNil(b bool)`

 SetOptionalNil sets the value for Optional to be an explicit nil

### UnsetOptional
`func (o *ModelCreditNotePos) UnsetOptional()`

UnsetOptional ensures that no value is present for Optional, not even an explicit nil
### GetTaxRate

`func (o *ModelCreditNotePos) GetTaxRate() float32`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *ModelCreditNotePos) GetTaxRateOk() (*float32, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *ModelCreditNotePos) SetTaxRate(v float32)`

SetTaxRate sets TaxRate field to given value.


### GetSumDiscount

`func (o *ModelCreditNotePos) GetSumDiscount() float32`

GetSumDiscount returns the SumDiscount field if non-nil, zero value otherwise.

### GetSumDiscountOk

`func (o *ModelCreditNotePos) GetSumDiscountOk() (*float32, bool)`

GetSumDiscountOk returns a tuple with the SumDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumDiscount

`func (o *ModelCreditNotePos) SetSumDiscount(v float32)`

SetSumDiscount sets SumDiscount field to given value.

### HasSumDiscount

`func (o *ModelCreditNotePos) HasSumDiscount() bool`

HasSumDiscount returns a boolean if a field has been set.

### SetSumDiscountNil

`func (o *ModelCreditNotePos) SetSumDiscountNil(b bool)`

 SetSumDiscountNil sets the value for SumDiscount to be an explicit nil

### UnsetSumDiscount
`func (o *ModelCreditNotePos) UnsetSumDiscount()`

UnsetSumDiscount ensures that no value is present for SumDiscount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


