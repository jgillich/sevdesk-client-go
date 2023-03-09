# ModelInvoicePosUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The invoice position id | [optional] [readonly] 
**ObjectName** | Pointer to **string** | The invoice position object name | [optional] [readonly] 
**Create** | Pointer to **time.Time** | Date of invoice position creation | [optional] [readonly] 
**Update** | Pointer to **time.Time** | Date of last invoice position update | [optional] [readonly] 
**Invoice** | Pointer to [**NullableModelInvoicePosUpdateInvoice**](ModelInvoicePosUpdateInvoice.md) |  | [optional] 
**Part** | Pointer to [**ModelInvoicePosPart**](ModelInvoicePosPart.md) |  | [optional] 
**Quantity** | Pointer to **NullableFloat32** | Quantity of the article/part | [optional] 
**Price** | Pointer to **NullableFloat32** | Price of the article/part. Is either gross or net, depending on the sevDesk account setting. | [optional] 
**Name** | Pointer to **NullableString** | Name of the article/part. | [optional] 
**Unity** | Pointer to [**ModelInvoicePosUnity**](ModelInvoicePosUnity.md) |  | [optional] 
**SevClient** | Pointer to [**ModelInvoicePosSevClient**](ModelInvoicePosSevClient.md) |  | [optional] 
**PositionNumber** | Pointer to **NullableInt32** | Position number of your position. Can be used to order multiple positions. | [optional] 
**Text** | Pointer to **NullableString** | A text describing your position. | [optional] 
**Discount** | Pointer to **NullableFloat32** | An optional discount of the position. | [optional] 
**TaxRate** | Pointer to **NullableFloat32** | Tax rate of the position. | [optional] 
**SumDiscount** | Pointer to **NullableFloat32** | Discount sum of the position | [optional] [readonly] 
**SumNetAccounting** | Pointer to **NullableFloat32** | Net accounting sum of the position | [optional] [readonly] 
**SumTaxAccounting** | Pointer to **NullableFloat32** | Tax accounting sum of the position | [optional] [readonly] 
**SumGrossAccounting** | Pointer to **NullableFloat32** | Gross accounting sum of the position | [optional] [readonly] 
**PriceNet** | Pointer to **NullableFloat32** | Net price of the part | [optional] [readonly] 
**PriceGross** | Pointer to **NullableFloat32** | Gross price of the part | [optional] 
**PriceTax** | Pointer to **NullableFloat32** | Tax on the price of the part | [optional] 

## Methods

### NewModelInvoicePosUpdate

`func NewModelInvoicePosUpdate() *ModelInvoicePosUpdate`

NewModelInvoicePosUpdate instantiates a new ModelInvoicePosUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelInvoicePosUpdateWithDefaults

`func NewModelInvoicePosUpdateWithDefaults() *ModelInvoicePosUpdate`

NewModelInvoicePosUpdateWithDefaults instantiates a new ModelInvoicePosUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelInvoicePosUpdate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelInvoicePosUpdate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelInvoicePosUpdate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelInvoicePosUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectName

`func (o *ModelInvoicePosUpdate) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ModelInvoicePosUpdate) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ModelInvoicePosUpdate) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ModelInvoicePosUpdate) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetCreate

`func (o *ModelInvoicePosUpdate) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ModelInvoicePosUpdate) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ModelInvoicePosUpdate) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ModelInvoicePosUpdate) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUpdate

`func (o *ModelInvoicePosUpdate) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ModelInvoicePosUpdate) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ModelInvoicePosUpdate) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ModelInvoicePosUpdate) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetInvoice

`func (o *ModelInvoicePosUpdate) GetInvoice() ModelInvoicePosUpdateInvoice`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *ModelInvoicePosUpdate) GetInvoiceOk() (*ModelInvoicePosUpdateInvoice, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *ModelInvoicePosUpdate) SetInvoice(v ModelInvoicePosUpdateInvoice)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *ModelInvoicePosUpdate) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### SetInvoiceNil

`func (o *ModelInvoicePosUpdate) SetInvoiceNil(b bool)`

 SetInvoiceNil sets the value for Invoice to be an explicit nil

### UnsetInvoice
`func (o *ModelInvoicePosUpdate) UnsetInvoice()`

UnsetInvoice ensures that no value is present for Invoice, not even an explicit nil
### GetPart

`func (o *ModelInvoicePosUpdate) GetPart() ModelInvoicePosPart`

GetPart returns the Part field if non-nil, zero value otherwise.

### GetPartOk

`func (o *ModelInvoicePosUpdate) GetPartOk() (*ModelInvoicePosPart, bool)`

GetPartOk returns a tuple with the Part field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPart

`func (o *ModelInvoicePosUpdate) SetPart(v ModelInvoicePosPart)`

SetPart sets Part field to given value.

### HasPart

`func (o *ModelInvoicePosUpdate) HasPart() bool`

HasPart returns a boolean if a field has been set.

### GetQuantity

`func (o *ModelInvoicePosUpdate) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ModelInvoicePosUpdate) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ModelInvoicePosUpdate) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ModelInvoicePosUpdate) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *ModelInvoicePosUpdate) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *ModelInvoicePosUpdate) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetPrice

`func (o *ModelInvoicePosUpdate) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ModelInvoicePosUpdate) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ModelInvoicePosUpdate) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ModelInvoicePosUpdate) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *ModelInvoicePosUpdate) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *ModelInvoicePosUpdate) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetName

`func (o *ModelInvoicePosUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelInvoicePosUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelInvoicePosUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelInvoicePosUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ModelInvoicePosUpdate) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ModelInvoicePosUpdate) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUnity

`func (o *ModelInvoicePosUpdate) GetUnity() ModelInvoicePosUnity`

GetUnity returns the Unity field if non-nil, zero value otherwise.

### GetUnityOk

`func (o *ModelInvoicePosUpdate) GetUnityOk() (*ModelInvoicePosUnity, bool)`

GetUnityOk returns a tuple with the Unity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnity

`func (o *ModelInvoicePosUpdate) SetUnity(v ModelInvoicePosUnity)`

SetUnity sets Unity field to given value.

### HasUnity

`func (o *ModelInvoicePosUpdate) HasUnity() bool`

HasUnity returns a boolean if a field has been set.

### GetSevClient

`func (o *ModelInvoicePosUpdate) GetSevClient() ModelInvoicePosSevClient`

GetSevClient returns the SevClient field if non-nil, zero value otherwise.

### GetSevClientOk

`func (o *ModelInvoicePosUpdate) GetSevClientOk() (*ModelInvoicePosSevClient, bool)`

GetSevClientOk returns a tuple with the SevClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSevClient

`func (o *ModelInvoicePosUpdate) SetSevClient(v ModelInvoicePosSevClient)`

SetSevClient sets SevClient field to given value.

### HasSevClient

`func (o *ModelInvoicePosUpdate) HasSevClient() bool`

HasSevClient returns a boolean if a field has been set.

### GetPositionNumber

`func (o *ModelInvoicePosUpdate) GetPositionNumber() int32`

GetPositionNumber returns the PositionNumber field if non-nil, zero value otherwise.

### GetPositionNumberOk

`func (o *ModelInvoicePosUpdate) GetPositionNumberOk() (*int32, bool)`

GetPositionNumberOk returns a tuple with the PositionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionNumber

`func (o *ModelInvoicePosUpdate) SetPositionNumber(v int32)`

SetPositionNumber sets PositionNumber field to given value.

### HasPositionNumber

`func (o *ModelInvoicePosUpdate) HasPositionNumber() bool`

HasPositionNumber returns a boolean if a field has been set.

### SetPositionNumberNil

`func (o *ModelInvoicePosUpdate) SetPositionNumberNil(b bool)`

 SetPositionNumberNil sets the value for PositionNumber to be an explicit nil

### UnsetPositionNumber
`func (o *ModelInvoicePosUpdate) UnsetPositionNumber()`

UnsetPositionNumber ensures that no value is present for PositionNumber, not even an explicit nil
### GetText

`func (o *ModelInvoicePosUpdate) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ModelInvoicePosUpdate) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ModelInvoicePosUpdate) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ModelInvoicePosUpdate) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *ModelInvoicePosUpdate) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *ModelInvoicePosUpdate) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetDiscount

`func (o *ModelInvoicePosUpdate) GetDiscount() float32`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *ModelInvoicePosUpdate) GetDiscountOk() (*float32, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *ModelInvoicePosUpdate) SetDiscount(v float32)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *ModelInvoicePosUpdate) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### SetDiscountNil

`func (o *ModelInvoicePosUpdate) SetDiscountNil(b bool)`

 SetDiscountNil sets the value for Discount to be an explicit nil

### UnsetDiscount
`func (o *ModelInvoicePosUpdate) UnsetDiscount()`

UnsetDiscount ensures that no value is present for Discount, not even an explicit nil
### GetTaxRate

`func (o *ModelInvoicePosUpdate) GetTaxRate() float32`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *ModelInvoicePosUpdate) GetTaxRateOk() (*float32, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *ModelInvoicePosUpdate) SetTaxRate(v float32)`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *ModelInvoicePosUpdate) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### SetTaxRateNil

`func (o *ModelInvoicePosUpdate) SetTaxRateNil(b bool)`

 SetTaxRateNil sets the value for TaxRate to be an explicit nil

### UnsetTaxRate
`func (o *ModelInvoicePosUpdate) UnsetTaxRate()`

UnsetTaxRate ensures that no value is present for TaxRate, not even an explicit nil
### GetSumDiscount

`func (o *ModelInvoicePosUpdate) GetSumDiscount() float32`

GetSumDiscount returns the SumDiscount field if non-nil, zero value otherwise.

### GetSumDiscountOk

`func (o *ModelInvoicePosUpdate) GetSumDiscountOk() (*float32, bool)`

GetSumDiscountOk returns a tuple with the SumDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumDiscount

`func (o *ModelInvoicePosUpdate) SetSumDiscount(v float32)`

SetSumDiscount sets SumDiscount field to given value.

### HasSumDiscount

`func (o *ModelInvoicePosUpdate) HasSumDiscount() bool`

HasSumDiscount returns a boolean if a field has been set.

### SetSumDiscountNil

`func (o *ModelInvoicePosUpdate) SetSumDiscountNil(b bool)`

 SetSumDiscountNil sets the value for SumDiscount to be an explicit nil

### UnsetSumDiscount
`func (o *ModelInvoicePosUpdate) UnsetSumDiscount()`

UnsetSumDiscount ensures that no value is present for SumDiscount, not even an explicit nil
### GetSumNetAccounting

`func (o *ModelInvoicePosUpdate) GetSumNetAccounting() float32`

GetSumNetAccounting returns the SumNetAccounting field if non-nil, zero value otherwise.

### GetSumNetAccountingOk

`func (o *ModelInvoicePosUpdate) GetSumNetAccountingOk() (*float32, bool)`

GetSumNetAccountingOk returns a tuple with the SumNetAccounting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumNetAccounting

`func (o *ModelInvoicePosUpdate) SetSumNetAccounting(v float32)`

SetSumNetAccounting sets SumNetAccounting field to given value.

### HasSumNetAccounting

`func (o *ModelInvoicePosUpdate) HasSumNetAccounting() bool`

HasSumNetAccounting returns a boolean if a field has been set.

### SetSumNetAccountingNil

`func (o *ModelInvoicePosUpdate) SetSumNetAccountingNil(b bool)`

 SetSumNetAccountingNil sets the value for SumNetAccounting to be an explicit nil

### UnsetSumNetAccounting
`func (o *ModelInvoicePosUpdate) UnsetSumNetAccounting()`

UnsetSumNetAccounting ensures that no value is present for SumNetAccounting, not even an explicit nil
### GetSumTaxAccounting

`func (o *ModelInvoicePosUpdate) GetSumTaxAccounting() float32`

GetSumTaxAccounting returns the SumTaxAccounting field if non-nil, zero value otherwise.

### GetSumTaxAccountingOk

`func (o *ModelInvoicePosUpdate) GetSumTaxAccountingOk() (*float32, bool)`

GetSumTaxAccountingOk returns a tuple with the SumTaxAccounting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumTaxAccounting

`func (o *ModelInvoicePosUpdate) SetSumTaxAccounting(v float32)`

SetSumTaxAccounting sets SumTaxAccounting field to given value.

### HasSumTaxAccounting

`func (o *ModelInvoicePosUpdate) HasSumTaxAccounting() bool`

HasSumTaxAccounting returns a boolean if a field has been set.

### SetSumTaxAccountingNil

`func (o *ModelInvoicePosUpdate) SetSumTaxAccountingNil(b bool)`

 SetSumTaxAccountingNil sets the value for SumTaxAccounting to be an explicit nil

### UnsetSumTaxAccounting
`func (o *ModelInvoicePosUpdate) UnsetSumTaxAccounting()`

UnsetSumTaxAccounting ensures that no value is present for SumTaxAccounting, not even an explicit nil
### GetSumGrossAccounting

`func (o *ModelInvoicePosUpdate) GetSumGrossAccounting() float32`

GetSumGrossAccounting returns the SumGrossAccounting field if non-nil, zero value otherwise.

### GetSumGrossAccountingOk

`func (o *ModelInvoicePosUpdate) GetSumGrossAccountingOk() (*float32, bool)`

GetSumGrossAccountingOk returns a tuple with the SumGrossAccounting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumGrossAccounting

`func (o *ModelInvoicePosUpdate) SetSumGrossAccounting(v float32)`

SetSumGrossAccounting sets SumGrossAccounting field to given value.

### HasSumGrossAccounting

`func (o *ModelInvoicePosUpdate) HasSumGrossAccounting() bool`

HasSumGrossAccounting returns a boolean if a field has been set.

### SetSumGrossAccountingNil

`func (o *ModelInvoicePosUpdate) SetSumGrossAccountingNil(b bool)`

 SetSumGrossAccountingNil sets the value for SumGrossAccounting to be an explicit nil

### UnsetSumGrossAccounting
`func (o *ModelInvoicePosUpdate) UnsetSumGrossAccounting()`

UnsetSumGrossAccounting ensures that no value is present for SumGrossAccounting, not even an explicit nil
### GetPriceNet

`func (o *ModelInvoicePosUpdate) GetPriceNet() float32`

GetPriceNet returns the PriceNet field if non-nil, zero value otherwise.

### GetPriceNetOk

`func (o *ModelInvoicePosUpdate) GetPriceNetOk() (*float32, bool)`

GetPriceNetOk returns a tuple with the PriceNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceNet

`func (o *ModelInvoicePosUpdate) SetPriceNet(v float32)`

SetPriceNet sets PriceNet field to given value.

### HasPriceNet

`func (o *ModelInvoicePosUpdate) HasPriceNet() bool`

HasPriceNet returns a boolean if a field has been set.

### SetPriceNetNil

`func (o *ModelInvoicePosUpdate) SetPriceNetNil(b bool)`

 SetPriceNetNil sets the value for PriceNet to be an explicit nil

### UnsetPriceNet
`func (o *ModelInvoicePosUpdate) UnsetPriceNet()`

UnsetPriceNet ensures that no value is present for PriceNet, not even an explicit nil
### GetPriceGross

`func (o *ModelInvoicePosUpdate) GetPriceGross() float32`

GetPriceGross returns the PriceGross field if non-nil, zero value otherwise.

### GetPriceGrossOk

`func (o *ModelInvoicePosUpdate) GetPriceGrossOk() (*float32, bool)`

GetPriceGrossOk returns a tuple with the PriceGross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceGross

`func (o *ModelInvoicePosUpdate) SetPriceGross(v float32)`

SetPriceGross sets PriceGross field to given value.

### HasPriceGross

`func (o *ModelInvoicePosUpdate) HasPriceGross() bool`

HasPriceGross returns a boolean if a field has been set.

### SetPriceGrossNil

`func (o *ModelInvoicePosUpdate) SetPriceGrossNil(b bool)`

 SetPriceGrossNil sets the value for PriceGross to be an explicit nil

### UnsetPriceGross
`func (o *ModelInvoicePosUpdate) UnsetPriceGross()`

UnsetPriceGross ensures that no value is present for PriceGross, not even an explicit nil
### GetPriceTax

`func (o *ModelInvoicePosUpdate) GetPriceTax() float32`

GetPriceTax returns the PriceTax field if non-nil, zero value otherwise.

### GetPriceTaxOk

`func (o *ModelInvoicePosUpdate) GetPriceTaxOk() (*float32, bool)`

GetPriceTaxOk returns a tuple with the PriceTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceTax

`func (o *ModelInvoicePosUpdate) SetPriceTax(v float32)`

SetPriceTax sets PriceTax field to given value.

### HasPriceTax

`func (o *ModelInvoicePosUpdate) HasPriceTax() bool`

HasPriceTax returns a boolean if a field has been set.

### SetPriceTaxNil

`func (o *ModelInvoicePosUpdate) SetPriceTaxNil(b bool)`

 SetPriceTaxNil sets the value for PriceTax to be an explicit nil

### UnsetPriceTax
`func (o *ModelInvoicePosUpdate) UnsetPriceTax()`

UnsetPriceTax ensures that no value is present for PriceTax, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


