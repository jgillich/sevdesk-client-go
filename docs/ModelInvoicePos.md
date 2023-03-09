# ModelInvoicePos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The invoice position id. &lt;span style&#x3D;&#39;color:red&#39;&gt;Required&lt;/span&gt; if you want to update an invoice position for an existing invoice | [optional] 
**ObjectName** | **string** | The invoice position object name | 
**MapAll** | **bool** |  | 
**Create** | Pointer to **time.Time** | Date of invoice position creation | [optional] [readonly] 
**Update** | Pointer to **time.Time** | Date of last invoice position update | [optional] [readonly] 
**Invoice** | Pointer to [**ModelInvoicePosInvoice**](ModelInvoicePosInvoice.md) |  | [optional] 
**Part** | Pointer to [**ModelInvoicePosPart**](ModelInvoicePosPart.md) |  | [optional] 
**Quantity** | **float32** | Quantity of the article/part | 
**Price** | Pointer to **NullableFloat32** | Price of the article/part. Is either gross or net, depending on the sevDesk account setting. | [optional] 
**Name** | Pointer to **NullableString** | Name of the article/part. | [optional] 
**Unity** | [**ModelInvoicePosUnity**](ModelInvoicePosUnity.md) |  | 
**SevClient** | Pointer to [**ModelInvoicePosSevClient**](ModelInvoicePosSevClient.md) |  | [optional] 
**PositionNumber** | Pointer to **NullableInt32** | Position number of your position. Can be used to order multiple positions. | [optional] 
**Text** | Pointer to **NullableString** | A text describing your position. | [optional] 
**Discount** | Pointer to **NullableFloat32** | An optional discount of the position. | [optional] 
**TaxRate** | **float32** | Tax rate of the position. | 
**SumDiscount** | Pointer to **NullableFloat32** | Discount sum of the position | [optional] [readonly] 
**SumNetAccounting** | Pointer to **NullableFloat32** | Net accounting sum of the position | [optional] [readonly] 
**SumTaxAccounting** | Pointer to **NullableFloat32** | Tax accounting sum of the position | [optional] [readonly] 
**SumGrossAccounting** | Pointer to **NullableFloat32** | Gross accounting sum of the position | [optional] [readonly] 
**PriceNet** | Pointer to **NullableFloat32** | Net price of the part | [optional] [readonly] 
**PriceGross** | Pointer to **NullableFloat32** | Gross price of the part | [optional] 
**PriceTax** | Pointer to **NullableFloat32** | Tax on the price of the part | [optional] 

## Methods

### NewModelInvoicePos

`func NewModelInvoicePos(objectName string, mapAll bool, quantity float32, unity ModelInvoicePosUnity, taxRate float32, ) *ModelInvoicePos`

NewModelInvoicePos instantiates a new ModelInvoicePos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelInvoicePosWithDefaults

`func NewModelInvoicePosWithDefaults() *ModelInvoicePos`

NewModelInvoicePosWithDefaults instantiates a new ModelInvoicePos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelInvoicePos) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelInvoicePos) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelInvoicePos) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelInvoicePos) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectName

`func (o *ModelInvoicePos) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ModelInvoicePos) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ModelInvoicePos) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.


### GetMapAll

`func (o *ModelInvoicePos) GetMapAll() bool`

GetMapAll returns the MapAll field if non-nil, zero value otherwise.

### GetMapAllOk

`func (o *ModelInvoicePos) GetMapAllOk() (*bool, bool)`

GetMapAllOk returns a tuple with the MapAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapAll

`func (o *ModelInvoicePos) SetMapAll(v bool)`

SetMapAll sets MapAll field to given value.


### GetCreate

`func (o *ModelInvoicePos) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ModelInvoicePos) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ModelInvoicePos) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ModelInvoicePos) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUpdate

`func (o *ModelInvoicePos) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ModelInvoicePos) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ModelInvoicePos) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ModelInvoicePos) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetInvoice

`func (o *ModelInvoicePos) GetInvoice() ModelInvoicePosInvoice`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *ModelInvoicePos) GetInvoiceOk() (*ModelInvoicePosInvoice, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *ModelInvoicePos) SetInvoice(v ModelInvoicePosInvoice)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *ModelInvoicePos) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetPart

`func (o *ModelInvoicePos) GetPart() ModelInvoicePosPart`

GetPart returns the Part field if non-nil, zero value otherwise.

### GetPartOk

`func (o *ModelInvoicePos) GetPartOk() (*ModelInvoicePosPart, bool)`

GetPartOk returns a tuple with the Part field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPart

`func (o *ModelInvoicePos) SetPart(v ModelInvoicePosPart)`

SetPart sets Part field to given value.

### HasPart

`func (o *ModelInvoicePos) HasPart() bool`

HasPart returns a boolean if a field has been set.

### GetQuantity

`func (o *ModelInvoicePos) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ModelInvoicePos) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ModelInvoicePos) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetPrice

`func (o *ModelInvoicePos) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ModelInvoicePos) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ModelInvoicePos) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ModelInvoicePos) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *ModelInvoicePos) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *ModelInvoicePos) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetName

`func (o *ModelInvoicePos) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelInvoicePos) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelInvoicePos) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelInvoicePos) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ModelInvoicePos) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ModelInvoicePos) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUnity

`func (o *ModelInvoicePos) GetUnity() ModelInvoicePosUnity`

GetUnity returns the Unity field if non-nil, zero value otherwise.

### GetUnityOk

`func (o *ModelInvoicePos) GetUnityOk() (*ModelInvoicePosUnity, bool)`

GetUnityOk returns a tuple with the Unity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnity

`func (o *ModelInvoicePos) SetUnity(v ModelInvoicePosUnity)`

SetUnity sets Unity field to given value.


### GetSevClient

`func (o *ModelInvoicePos) GetSevClient() ModelInvoicePosSevClient`

GetSevClient returns the SevClient field if non-nil, zero value otherwise.

### GetSevClientOk

`func (o *ModelInvoicePos) GetSevClientOk() (*ModelInvoicePosSevClient, bool)`

GetSevClientOk returns a tuple with the SevClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSevClient

`func (o *ModelInvoicePos) SetSevClient(v ModelInvoicePosSevClient)`

SetSevClient sets SevClient field to given value.

### HasSevClient

`func (o *ModelInvoicePos) HasSevClient() bool`

HasSevClient returns a boolean if a field has been set.

### GetPositionNumber

`func (o *ModelInvoicePos) GetPositionNumber() int32`

GetPositionNumber returns the PositionNumber field if non-nil, zero value otherwise.

### GetPositionNumberOk

`func (o *ModelInvoicePos) GetPositionNumberOk() (*int32, bool)`

GetPositionNumberOk returns a tuple with the PositionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionNumber

`func (o *ModelInvoicePos) SetPositionNumber(v int32)`

SetPositionNumber sets PositionNumber field to given value.

### HasPositionNumber

`func (o *ModelInvoicePos) HasPositionNumber() bool`

HasPositionNumber returns a boolean if a field has been set.

### SetPositionNumberNil

`func (o *ModelInvoicePos) SetPositionNumberNil(b bool)`

 SetPositionNumberNil sets the value for PositionNumber to be an explicit nil

### UnsetPositionNumber
`func (o *ModelInvoicePos) UnsetPositionNumber()`

UnsetPositionNumber ensures that no value is present for PositionNumber, not even an explicit nil
### GetText

`func (o *ModelInvoicePos) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ModelInvoicePos) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ModelInvoicePos) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ModelInvoicePos) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *ModelInvoicePos) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *ModelInvoicePos) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetDiscount

`func (o *ModelInvoicePos) GetDiscount() float32`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *ModelInvoicePos) GetDiscountOk() (*float32, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *ModelInvoicePos) SetDiscount(v float32)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *ModelInvoicePos) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### SetDiscountNil

`func (o *ModelInvoicePos) SetDiscountNil(b bool)`

 SetDiscountNil sets the value for Discount to be an explicit nil

### UnsetDiscount
`func (o *ModelInvoicePos) UnsetDiscount()`

UnsetDiscount ensures that no value is present for Discount, not even an explicit nil
### GetTaxRate

`func (o *ModelInvoicePos) GetTaxRate() float32`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *ModelInvoicePos) GetTaxRateOk() (*float32, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *ModelInvoicePos) SetTaxRate(v float32)`

SetTaxRate sets TaxRate field to given value.


### GetSumDiscount

`func (o *ModelInvoicePos) GetSumDiscount() float32`

GetSumDiscount returns the SumDiscount field if non-nil, zero value otherwise.

### GetSumDiscountOk

`func (o *ModelInvoicePos) GetSumDiscountOk() (*float32, bool)`

GetSumDiscountOk returns a tuple with the SumDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumDiscount

`func (o *ModelInvoicePos) SetSumDiscount(v float32)`

SetSumDiscount sets SumDiscount field to given value.

### HasSumDiscount

`func (o *ModelInvoicePos) HasSumDiscount() bool`

HasSumDiscount returns a boolean if a field has been set.

### SetSumDiscountNil

`func (o *ModelInvoicePos) SetSumDiscountNil(b bool)`

 SetSumDiscountNil sets the value for SumDiscount to be an explicit nil

### UnsetSumDiscount
`func (o *ModelInvoicePos) UnsetSumDiscount()`

UnsetSumDiscount ensures that no value is present for SumDiscount, not even an explicit nil
### GetSumNetAccounting

`func (o *ModelInvoicePos) GetSumNetAccounting() float32`

GetSumNetAccounting returns the SumNetAccounting field if non-nil, zero value otherwise.

### GetSumNetAccountingOk

`func (o *ModelInvoicePos) GetSumNetAccountingOk() (*float32, bool)`

GetSumNetAccountingOk returns a tuple with the SumNetAccounting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumNetAccounting

`func (o *ModelInvoicePos) SetSumNetAccounting(v float32)`

SetSumNetAccounting sets SumNetAccounting field to given value.

### HasSumNetAccounting

`func (o *ModelInvoicePos) HasSumNetAccounting() bool`

HasSumNetAccounting returns a boolean if a field has been set.

### SetSumNetAccountingNil

`func (o *ModelInvoicePos) SetSumNetAccountingNil(b bool)`

 SetSumNetAccountingNil sets the value for SumNetAccounting to be an explicit nil

### UnsetSumNetAccounting
`func (o *ModelInvoicePos) UnsetSumNetAccounting()`

UnsetSumNetAccounting ensures that no value is present for SumNetAccounting, not even an explicit nil
### GetSumTaxAccounting

`func (o *ModelInvoicePos) GetSumTaxAccounting() float32`

GetSumTaxAccounting returns the SumTaxAccounting field if non-nil, zero value otherwise.

### GetSumTaxAccountingOk

`func (o *ModelInvoicePos) GetSumTaxAccountingOk() (*float32, bool)`

GetSumTaxAccountingOk returns a tuple with the SumTaxAccounting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumTaxAccounting

`func (o *ModelInvoicePos) SetSumTaxAccounting(v float32)`

SetSumTaxAccounting sets SumTaxAccounting field to given value.

### HasSumTaxAccounting

`func (o *ModelInvoicePos) HasSumTaxAccounting() bool`

HasSumTaxAccounting returns a boolean if a field has been set.

### SetSumTaxAccountingNil

`func (o *ModelInvoicePos) SetSumTaxAccountingNil(b bool)`

 SetSumTaxAccountingNil sets the value for SumTaxAccounting to be an explicit nil

### UnsetSumTaxAccounting
`func (o *ModelInvoicePos) UnsetSumTaxAccounting()`

UnsetSumTaxAccounting ensures that no value is present for SumTaxAccounting, not even an explicit nil
### GetSumGrossAccounting

`func (o *ModelInvoicePos) GetSumGrossAccounting() float32`

GetSumGrossAccounting returns the SumGrossAccounting field if non-nil, zero value otherwise.

### GetSumGrossAccountingOk

`func (o *ModelInvoicePos) GetSumGrossAccountingOk() (*float32, bool)`

GetSumGrossAccountingOk returns a tuple with the SumGrossAccounting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumGrossAccounting

`func (o *ModelInvoicePos) SetSumGrossAccounting(v float32)`

SetSumGrossAccounting sets SumGrossAccounting field to given value.

### HasSumGrossAccounting

`func (o *ModelInvoicePos) HasSumGrossAccounting() bool`

HasSumGrossAccounting returns a boolean if a field has been set.

### SetSumGrossAccountingNil

`func (o *ModelInvoicePos) SetSumGrossAccountingNil(b bool)`

 SetSumGrossAccountingNil sets the value for SumGrossAccounting to be an explicit nil

### UnsetSumGrossAccounting
`func (o *ModelInvoicePos) UnsetSumGrossAccounting()`

UnsetSumGrossAccounting ensures that no value is present for SumGrossAccounting, not even an explicit nil
### GetPriceNet

`func (o *ModelInvoicePos) GetPriceNet() float32`

GetPriceNet returns the PriceNet field if non-nil, zero value otherwise.

### GetPriceNetOk

`func (o *ModelInvoicePos) GetPriceNetOk() (*float32, bool)`

GetPriceNetOk returns a tuple with the PriceNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceNet

`func (o *ModelInvoicePos) SetPriceNet(v float32)`

SetPriceNet sets PriceNet field to given value.

### HasPriceNet

`func (o *ModelInvoicePos) HasPriceNet() bool`

HasPriceNet returns a boolean if a field has been set.

### SetPriceNetNil

`func (o *ModelInvoicePos) SetPriceNetNil(b bool)`

 SetPriceNetNil sets the value for PriceNet to be an explicit nil

### UnsetPriceNet
`func (o *ModelInvoicePos) UnsetPriceNet()`

UnsetPriceNet ensures that no value is present for PriceNet, not even an explicit nil
### GetPriceGross

`func (o *ModelInvoicePos) GetPriceGross() float32`

GetPriceGross returns the PriceGross field if non-nil, zero value otherwise.

### GetPriceGrossOk

`func (o *ModelInvoicePos) GetPriceGrossOk() (*float32, bool)`

GetPriceGrossOk returns a tuple with the PriceGross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceGross

`func (o *ModelInvoicePos) SetPriceGross(v float32)`

SetPriceGross sets PriceGross field to given value.

### HasPriceGross

`func (o *ModelInvoicePos) HasPriceGross() bool`

HasPriceGross returns a boolean if a field has been set.

### SetPriceGrossNil

`func (o *ModelInvoicePos) SetPriceGrossNil(b bool)`

 SetPriceGrossNil sets the value for PriceGross to be an explicit nil

### UnsetPriceGross
`func (o *ModelInvoicePos) UnsetPriceGross()`

UnsetPriceGross ensures that no value is present for PriceGross, not even an explicit nil
### GetPriceTax

`func (o *ModelInvoicePos) GetPriceTax() float32`

GetPriceTax returns the PriceTax field if non-nil, zero value otherwise.

### GetPriceTaxOk

`func (o *ModelInvoicePos) GetPriceTaxOk() (*float32, bool)`

GetPriceTaxOk returns a tuple with the PriceTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceTax

`func (o *ModelInvoicePos) SetPriceTax(v float32)`

SetPriceTax sets PriceTax field to given value.

### HasPriceTax

`func (o *ModelInvoicePos) HasPriceTax() bool`

HasPriceTax returns a boolean if a field has been set.

### SetPriceTaxNil

`func (o *ModelInvoicePos) SetPriceTaxNil(b bool)`

 SetPriceTaxNil sets the value for PriceTax to be an explicit nil

### UnsetPriceTax
`func (o *ModelInvoicePos) UnsetPriceTax()`

UnsetPriceTax ensures that no value is present for PriceTax, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


