# ModelCreditNotePosResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The creditNote position id | [optional] [readonly] 
**ObjectName** | Pointer to **string** | The creditNote position object name | [optional] [readonly] 
**Create** | Pointer to **string** | Date of creditNote position creation | [optional] [readonly] 
**Update** | Pointer to **string** | Date of last creditNote position update | [optional] [readonly] 
**CreditNote** | [**ModelCreditNotePosResponseCreditNote**](ModelCreditNotePosResponseCreditNote.md) |  | 
**Part** | Pointer to [**ModelCreditNotePosResponsePart**](ModelCreditNotePosResponsePart.md) |  | [optional] 
**Quantity** | **string** | Quantity of the article/part | 
**Price** | Pointer to **NullableString** | Price of the article/part. Is either gross or net, depending on the sevDesk account setting. | [optional] 
**PriceNet** | Pointer to **NullableString** | Net price of the part | [optional] [readonly] 
**PriceTax** | Pointer to **NullableString** | Tax on the price of the part | [optional] 
**PriceGross** | Pointer to **NullableString** | Gross price of the part | [optional] 
**Name** | Pointer to **NullableString** | Name of the article/part. | [optional] 
**Unity** | [**ModelCreditNotePosResponseUnity**](ModelCreditNotePosResponseUnity.md) |  | 
**SevClient** | Pointer to [**ModelCreditNotePosResponseSevClient**](ModelCreditNotePosResponseSevClient.md) |  | [optional] 
**PositionNumber** | Pointer to **NullableString** | Position number of your position. Can be used to creditNote multiple positions. | [optional] 
**Text** | Pointer to **NullableString** | A text describing your position. | [optional] 
**Discount** | Pointer to **NullableString** | An optional discount of the position. | [optional] 
**Optional** | Pointer to **NullableBool** | Defines if the position is optional. | [optional] 
**TaxRate** | **string** | Tax rate of the position. | 
**SumDiscount** | Pointer to **NullableString** | Discount sum of the position | [optional] [readonly] 

## Methods

### NewModelCreditNotePosResponse

`func NewModelCreditNotePosResponse(creditNote ModelCreditNotePosResponseCreditNote, quantity string, unity ModelCreditNotePosResponseUnity, taxRate string, ) *ModelCreditNotePosResponse`

NewModelCreditNotePosResponse instantiates a new ModelCreditNotePosResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelCreditNotePosResponseWithDefaults

`func NewModelCreditNotePosResponseWithDefaults() *ModelCreditNotePosResponse`

NewModelCreditNotePosResponseWithDefaults instantiates a new ModelCreditNotePosResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelCreditNotePosResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelCreditNotePosResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelCreditNotePosResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelCreditNotePosResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectName

`func (o *ModelCreditNotePosResponse) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ModelCreditNotePosResponse) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ModelCreditNotePosResponse) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ModelCreditNotePosResponse) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetCreate

`func (o *ModelCreditNotePosResponse) GetCreate() string`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ModelCreditNotePosResponse) GetCreateOk() (*string, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ModelCreditNotePosResponse) SetCreate(v string)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ModelCreditNotePosResponse) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUpdate

`func (o *ModelCreditNotePosResponse) GetUpdate() string`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ModelCreditNotePosResponse) GetUpdateOk() (*string, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ModelCreditNotePosResponse) SetUpdate(v string)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ModelCreditNotePosResponse) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetCreditNote

`func (o *ModelCreditNotePosResponse) GetCreditNote() ModelCreditNotePosResponseCreditNote`

GetCreditNote returns the CreditNote field if non-nil, zero value otherwise.

### GetCreditNoteOk

`func (o *ModelCreditNotePosResponse) GetCreditNoteOk() (*ModelCreditNotePosResponseCreditNote, bool)`

GetCreditNoteOk returns a tuple with the CreditNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNote

`func (o *ModelCreditNotePosResponse) SetCreditNote(v ModelCreditNotePosResponseCreditNote)`

SetCreditNote sets CreditNote field to given value.


### GetPart

`func (o *ModelCreditNotePosResponse) GetPart() ModelCreditNotePosResponsePart`

GetPart returns the Part field if non-nil, zero value otherwise.

### GetPartOk

`func (o *ModelCreditNotePosResponse) GetPartOk() (*ModelCreditNotePosResponsePart, bool)`

GetPartOk returns a tuple with the Part field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPart

`func (o *ModelCreditNotePosResponse) SetPart(v ModelCreditNotePosResponsePart)`

SetPart sets Part field to given value.

### HasPart

`func (o *ModelCreditNotePosResponse) HasPart() bool`

HasPart returns a boolean if a field has been set.

### GetQuantity

`func (o *ModelCreditNotePosResponse) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ModelCreditNotePosResponse) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ModelCreditNotePosResponse) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.


### GetPrice

`func (o *ModelCreditNotePosResponse) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ModelCreditNotePosResponse) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ModelCreditNotePosResponse) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ModelCreditNotePosResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *ModelCreditNotePosResponse) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *ModelCreditNotePosResponse) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetPriceNet

`func (o *ModelCreditNotePosResponse) GetPriceNet() string`

GetPriceNet returns the PriceNet field if non-nil, zero value otherwise.

### GetPriceNetOk

`func (o *ModelCreditNotePosResponse) GetPriceNetOk() (*string, bool)`

GetPriceNetOk returns a tuple with the PriceNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceNet

`func (o *ModelCreditNotePosResponse) SetPriceNet(v string)`

SetPriceNet sets PriceNet field to given value.

### HasPriceNet

`func (o *ModelCreditNotePosResponse) HasPriceNet() bool`

HasPriceNet returns a boolean if a field has been set.

### SetPriceNetNil

`func (o *ModelCreditNotePosResponse) SetPriceNetNil(b bool)`

 SetPriceNetNil sets the value for PriceNet to be an explicit nil

### UnsetPriceNet
`func (o *ModelCreditNotePosResponse) UnsetPriceNet()`

UnsetPriceNet ensures that no value is present for PriceNet, not even an explicit nil
### GetPriceTax

`func (o *ModelCreditNotePosResponse) GetPriceTax() string`

GetPriceTax returns the PriceTax field if non-nil, zero value otherwise.

### GetPriceTaxOk

`func (o *ModelCreditNotePosResponse) GetPriceTaxOk() (*string, bool)`

GetPriceTaxOk returns a tuple with the PriceTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceTax

`func (o *ModelCreditNotePosResponse) SetPriceTax(v string)`

SetPriceTax sets PriceTax field to given value.

### HasPriceTax

`func (o *ModelCreditNotePosResponse) HasPriceTax() bool`

HasPriceTax returns a boolean if a field has been set.

### SetPriceTaxNil

`func (o *ModelCreditNotePosResponse) SetPriceTaxNil(b bool)`

 SetPriceTaxNil sets the value for PriceTax to be an explicit nil

### UnsetPriceTax
`func (o *ModelCreditNotePosResponse) UnsetPriceTax()`

UnsetPriceTax ensures that no value is present for PriceTax, not even an explicit nil
### GetPriceGross

`func (o *ModelCreditNotePosResponse) GetPriceGross() string`

GetPriceGross returns the PriceGross field if non-nil, zero value otherwise.

### GetPriceGrossOk

`func (o *ModelCreditNotePosResponse) GetPriceGrossOk() (*string, bool)`

GetPriceGrossOk returns a tuple with the PriceGross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceGross

`func (o *ModelCreditNotePosResponse) SetPriceGross(v string)`

SetPriceGross sets PriceGross field to given value.

### HasPriceGross

`func (o *ModelCreditNotePosResponse) HasPriceGross() bool`

HasPriceGross returns a boolean if a field has been set.

### SetPriceGrossNil

`func (o *ModelCreditNotePosResponse) SetPriceGrossNil(b bool)`

 SetPriceGrossNil sets the value for PriceGross to be an explicit nil

### UnsetPriceGross
`func (o *ModelCreditNotePosResponse) UnsetPriceGross()`

UnsetPriceGross ensures that no value is present for PriceGross, not even an explicit nil
### GetName

`func (o *ModelCreditNotePosResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelCreditNotePosResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelCreditNotePosResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelCreditNotePosResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ModelCreditNotePosResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ModelCreditNotePosResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUnity

`func (o *ModelCreditNotePosResponse) GetUnity() ModelCreditNotePosResponseUnity`

GetUnity returns the Unity field if non-nil, zero value otherwise.

### GetUnityOk

`func (o *ModelCreditNotePosResponse) GetUnityOk() (*ModelCreditNotePosResponseUnity, bool)`

GetUnityOk returns a tuple with the Unity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnity

`func (o *ModelCreditNotePosResponse) SetUnity(v ModelCreditNotePosResponseUnity)`

SetUnity sets Unity field to given value.


### GetSevClient

`func (o *ModelCreditNotePosResponse) GetSevClient() ModelCreditNotePosResponseSevClient`

GetSevClient returns the SevClient field if non-nil, zero value otherwise.

### GetSevClientOk

`func (o *ModelCreditNotePosResponse) GetSevClientOk() (*ModelCreditNotePosResponseSevClient, bool)`

GetSevClientOk returns a tuple with the SevClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSevClient

`func (o *ModelCreditNotePosResponse) SetSevClient(v ModelCreditNotePosResponseSevClient)`

SetSevClient sets SevClient field to given value.

### HasSevClient

`func (o *ModelCreditNotePosResponse) HasSevClient() bool`

HasSevClient returns a boolean if a field has been set.

### GetPositionNumber

`func (o *ModelCreditNotePosResponse) GetPositionNumber() string`

GetPositionNumber returns the PositionNumber field if non-nil, zero value otherwise.

### GetPositionNumberOk

`func (o *ModelCreditNotePosResponse) GetPositionNumberOk() (*string, bool)`

GetPositionNumberOk returns a tuple with the PositionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionNumber

`func (o *ModelCreditNotePosResponse) SetPositionNumber(v string)`

SetPositionNumber sets PositionNumber field to given value.

### HasPositionNumber

`func (o *ModelCreditNotePosResponse) HasPositionNumber() bool`

HasPositionNumber returns a boolean if a field has been set.

### SetPositionNumberNil

`func (o *ModelCreditNotePosResponse) SetPositionNumberNil(b bool)`

 SetPositionNumberNil sets the value for PositionNumber to be an explicit nil

### UnsetPositionNumber
`func (o *ModelCreditNotePosResponse) UnsetPositionNumber()`

UnsetPositionNumber ensures that no value is present for PositionNumber, not even an explicit nil
### GetText

`func (o *ModelCreditNotePosResponse) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ModelCreditNotePosResponse) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ModelCreditNotePosResponse) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ModelCreditNotePosResponse) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *ModelCreditNotePosResponse) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *ModelCreditNotePosResponse) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetDiscount

`func (o *ModelCreditNotePosResponse) GetDiscount() string`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *ModelCreditNotePosResponse) GetDiscountOk() (*string, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *ModelCreditNotePosResponse) SetDiscount(v string)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *ModelCreditNotePosResponse) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### SetDiscountNil

`func (o *ModelCreditNotePosResponse) SetDiscountNil(b bool)`

 SetDiscountNil sets the value for Discount to be an explicit nil

### UnsetDiscount
`func (o *ModelCreditNotePosResponse) UnsetDiscount()`

UnsetDiscount ensures that no value is present for Discount, not even an explicit nil
### GetOptional

`func (o *ModelCreditNotePosResponse) GetOptional() bool`

GetOptional returns the Optional field if non-nil, zero value otherwise.

### GetOptionalOk

`func (o *ModelCreditNotePosResponse) GetOptionalOk() (*bool, bool)`

GetOptionalOk returns a tuple with the Optional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptional

`func (o *ModelCreditNotePosResponse) SetOptional(v bool)`

SetOptional sets Optional field to given value.

### HasOptional

`func (o *ModelCreditNotePosResponse) HasOptional() bool`

HasOptional returns a boolean if a field has been set.

### SetOptionalNil

`func (o *ModelCreditNotePosResponse) SetOptionalNil(b bool)`

 SetOptionalNil sets the value for Optional to be an explicit nil

### UnsetOptional
`func (o *ModelCreditNotePosResponse) UnsetOptional()`

UnsetOptional ensures that no value is present for Optional, not even an explicit nil
### GetTaxRate

`func (o *ModelCreditNotePosResponse) GetTaxRate() string`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *ModelCreditNotePosResponse) GetTaxRateOk() (*string, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *ModelCreditNotePosResponse) SetTaxRate(v string)`

SetTaxRate sets TaxRate field to given value.


### GetSumDiscount

`func (o *ModelCreditNotePosResponse) GetSumDiscount() string`

GetSumDiscount returns the SumDiscount field if non-nil, zero value otherwise.

### GetSumDiscountOk

`func (o *ModelCreditNotePosResponse) GetSumDiscountOk() (*string, bool)`

GetSumDiscountOk returns a tuple with the SumDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumDiscount

`func (o *ModelCreditNotePosResponse) SetSumDiscount(v string)`

SetSumDiscount sets SumDiscount field to given value.

### HasSumDiscount

`func (o *ModelCreditNotePosResponse) HasSumDiscount() bool`

HasSumDiscount returns a boolean if a field has been set.

### SetSumDiscountNil

`func (o *ModelCreditNotePosResponse) SetSumDiscountNil(b bool)`

 SetSumDiscountNil sets the value for SumDiscount to be an explicit nil

### UnsetSumDiscount
`func (o *ModelCreditNotePosResponse) UnsetSumDiscount()`

UnsetSumDiscount ensures that no value is present for SumDiscount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


