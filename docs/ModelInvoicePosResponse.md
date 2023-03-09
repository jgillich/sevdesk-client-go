# ModelInvoicePosResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The invoice position id | [optional] [readonly] 
**ObjectName** | Pointer to **string** | The invoice position object name | [optional] [readonly] 
**Create** | Pointer to **time.Time** | Date of invoice position creation | [optional] [readonly] 
**Update** | Pointer to **time.Time** | Date of last invoice position update | [optional] [readonly] 
**Invoice** | Pointer to [**ModelInvoicePosResponseInvoice**](ModelInvoicePosResponseInvoice.md) |  | [optional] 
**Part** | Pointer to [**ModelInvoicePosResponsePart**](ModelInvoicePosResponsePart.md) |  | [optional] 
**Quantity** | Pointer to **string** | Quantity of the article/part | [optional] [readonly] 
**Price** | Pointer to **string** | Price of the article/part. Is either gross or net, depending on the sevDesk account setting. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the article/part. | [optional] [readonly] 
**Unity** | Pointer to [**ModelInvoicePosResponseUnity**](ModelInvoicePosResponseUnity.md) |  | [optional] 
**SevClient** | Pointer to [**ModelInvoicePosResponseSevClient**](ModelInvoicePosResponseSevClient.md) |  | [optional] 
**PositionNumber** | Pointer to **string** | Position number of your position. Can be used to order multiple positions. | [optional] [readonly] 
**Text** | Pointer to **string** | A text describing your position. | [optional] [readonly] 
**Discount** | Pointer to **string** | An optional discount of the position. | [optional] [readonly] 
**TaxRate** | Pointer to **string** | Tax rate of the position. | [optional] [readonly] 
**SumDiscount** | Pointer to **string** | Discount sum of the position | [optional] [readonly] 
**SumNetAccounting** | Pointer to **string** | Net accounting sum of the position | [optional] [readonly] 
**SumTaxAccounting** | Pointer to **string** | Tax accounting sum of the position | [optional] [readonly] 
**SumGrossAccounting** | Pointer to **string** | Gross accounting sum of the position | [optional] [readonly] 
**PriceNet** | Pointer to **string** | Net price of the part | [optional] [readonly] 
**PriceGross** | Pointer to **string** | Gross price of the part | [optional] [readonly] 
**PriceTax** | Pointer to **string** | Tax on the price of the part | [optional] [readonly] 

## Methods

### NewModelInvoicePosResponse

`func NewModelInvoicePosResponse() *ModelInvoicePosResponse`

NewModelInvoicePosResponse instantiates a new ModelInvoicePosResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelInvoicePosResponseWithDefaults

`func NewModelInvoicePosResponseWithDefaults() *ModelInvoicePosResponse`

NewModelInvoicePosResponseWithDefaults instantiates a new ModelInvoicePosResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelInvoicePosResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelInvoicePosResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelInvoicePosResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelInvoicePosResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectName

`func (o *ModelInvoicePosResponse) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ModelInvoicePosResponse) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ModelInvoicePosResponse) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ModelInvoicePosResponse) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetCreate

`func (o *ModelInvoicePosResponse) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ModelInvoicePosResponse) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ModelInvoicePosResponse) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ModelInvoicePosResponse) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUpdate

`func (o *ModelInvoicePosResponse) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ModelInvoicePosResponse) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ModelInvoicePosResponse) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ModelInvoicePosResponse) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetInvoice

`func (o *ModelInvoicePosResponse) GetInvoice() ModelInvoicePosResponseInvoice`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *ModelInvoicePosResponse) GetInvoiceOk() (*ModelInvoicePosResponseInvoice, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *ModelInvoicePosResponse) SetInvoice(v ModelInvoicePosResponseInvoice)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *ModelInvoicePosResponse) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetPart

`func (o *ModelInvoicePosResponse) GetPart() ModelInvoicePosResponsePart`

GetPart returns the Part field if non-nil, zero value otherwise.

### GetPartOk

`func (o *ModelInvoicePosResponse) GetPartOk() (*ModelInvoicePosResponsePart, bool)`

GetPartOk returns a tuple with the Part field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPart

`func (o *ModelInvoicePosResponse) SetPart(v ModelInvoicePosResponsePart)`

SetPart sets Part field to given value.

### HasPart

`func (o *ModelInvoicePosResponse) HasPart() bool`

HasPart returns a boolean if a field has been set.

### GetQuantity

`func (o *ModelInvoicePosResponse) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ModelInvoicePosResponse) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ModelInvoicePosResponse) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ModelInvoicePosResponse) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetPrice

`func (o *ModelInvoicePosResponse) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ModelInvoicePosResponse) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ModelInvoicePosResponse) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ModelInvoicePosResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetName

`func (o *ModelInvoicePosResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelInvoicePosResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelInvoicePosResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelInvoicePosResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUnity

`func (o *ModelInvoicePosResponse) GetUnity() ModelInvoicePosResponseUnity`

GetUnity returns the Unity field if non-nil, zero value otherwise.

### GetUnityOk

`func (o *ModelInvoicePosResponse) GetUnityOk() (*ModelInvoicePosResponseUnity, bool)`

GetUnityOk returns a tuple with the Unity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnity

`func (o *ModelInvoicePosResponse) SetUnity(v ModelInvoicePosResponseUnity)`

SetUnity sets Unity field to given value.

### HasUnity

`func (o *ModelInvoicePosResponse) HasUnity() bool`

HasUnity returns a boolean if a field has been set.

### GetSevClient

`func (o *ModelInvoicePosResponse) GetSevClient() ModelInvoicePosResponseSevClient`

GetSevClient returns the SevClient field if non-nil, zero value otherwise.

### GetSevClientOk

`func (o *ModelInvoicePosResponse) GetSevClientOk() (*ModelInvoicePosResponseSevClient, bool)`

GetSevClientOk returns a tuple with the SevClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSevClient

`func (o *ModelInvoicePosResponse) SetSevClient(v ModelInvoicePosResponseSevClient)`

SetSevClient sets SevClient field to given value.

### HasSevClient

`func (o *ModelInvoicePosResponse) HasSevClient() bool`

HasSevClient returns a boolean if a field has been set.

### GetPositionNumber

`func (o *ModelInvoicePosResponse) GetPositionNumber() string`

GetPositionNumber returns the PositionNumber field if non-nil, zero value otherwise.

### GetPositionNumberOk

`func (o *ModelInvoicePosResponse) GetPositionNumberOk() (*string, bool)`

GetPositionNumberOk returns a tuple with the PositionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionNumber

`func (o *ModelInvoicePosResponse) SetPositionNumber(v string)`

SetPositionNumber sets PositionNumber field to given value.

### HasPositionNumber

`func (o *ModelInvoicePosResponse) HasPositionNumber() bool`

HasPositionNumber returns a boolean if a field has been set.

### GetText

`func (o *ModelInvoicePosResponse) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ModelInvoicePosResponse) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ModelInvoicePosResponse) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ModelInvoicePosResponse) HasText() bool`

HasText returns a boolean if a field has been set.

### GetDiscount

`func (o *ModelInvoicePosResponse) GetDiscount() string`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *ModelInvoicePosResponse) GetDiscountOk() (*string, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *ModelInvoicePosResponse) SetDiscount(v string)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *ModelInvoicePosResponse) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetTaxRate

`func (o *ModelInvoicePosResponse) GetTaxRate() string`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *ModelInvoicePosResponse) GetTaxRateOk() (*string, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *ModelInvoicePosResponse) SetTaxRate(v string)`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *ModelInvoicePosResponse) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### GetSumDiscount

`func (o *ModelInvoicePosResponse) GetSumDiscount() string`

GetSumDiscount returns the SumDiscount field if non-nil, zero value otherwise.

### GetSumDiscountOk

`func (o *ModelInvoicePosResponse) GetSumDiscountOk() (*string, bool)`

GetSumDiscountOk returns a tuple with the SumDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumDiscount

`func (o *ModelInvoicePosResponse) SetSumDiscount(v string)`

SetSumDiscount sets SumDiscount field to given value.

### HasSumDiscount

`func (o *ModelInvoicePosResponse) HasSumDiscount() bool`

HasSumDiscount returns a boolean if a field has been set.

### GetSumNetAccounting

`func (o *ModelInvoicePosResponse) GetSumNetAccounting() string`

GetSumNetAccounting returns the SumNetAccounting field if non-nil, zero value otherwise.

### GetSumNetAccountingOk

`func (o *ModelInvoicePosResponse) GetSumNetAccountingOk() (*string, bool)`

GetSumNetAccountingOk returns a tuple with the SumNetAccounting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumNetAccounting

`func (o *ModelInvoicePosResponse) SetSumNetAccounting(v string)`

SetSumNetAccounting sets SumNetAccounting field to given value.

### HasSumNetAccounting

`func (o *ModelInvoicePosResponse) HasSumNetAccounting() bool`

HasSumNetAccounting returns a boolean if a field has been set.

### GetSumTaxAccounting

`func (o *ModelInvoicePosResponse) GetSumTaxAccounting() string`

GetSumTaxAccounting returns the SumTaxAccounting field if non-nil, zero value otherwise.

### GetSumTaxAccountingOk

`func (o *ModelInvoicePosResponse) GetSumTaxAccountingOk() (*string, bool)`

GetSumTaxAccountingOk returns a tuple with the SumTaxAccounting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumTaxAccounting

`func (o *ModelInvoicePosResponse) SetSumTaxAccounting(v string)`

SetSumTaxAccounting sets SumTaxAccounting field to given value.

### HasSumTaxAccounting

`func (o *ModelInvoicePosResponse) HasSumTaxAccounting() bool`

HasSumTaxAccounting returns a boolean if a field has been set.

### GetSumGrossAccounting

`func (o *ModelInvoicePosResponse) GetSumGrossAccounting() string`

GetSumGrossAccounting returns the SumGrossAccounting field if non-nil, zero value otherwise.

### GetSumGrossAccountingOk

`func (o *ModelInvoicePosResponse) GetSumGrossAccountingOk() (*string, bool)`

GetSumGrossAccountingOk returns a tuple with the SumGrossAccounting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumGrossAccounting

`func (o *ModelInvoicePosResponse) SetSumGrossAccounting(v string)`

SetSumGrossAccounting sets SumGrossAccounting field to given value.

### HasSumGrossAccounting

`func (o *ModelInvoicePosResponse) HasSumGrossAccounting() bool`

HasSumGrossAccounting returns a boolean if a field has been set.

### GetPriceNet

`func (o *ModelInvoicePosResponse) GetPriceNet() string`

GetPriceNet returns the PriceNet field if non-nil, zero value otherwise.

### GetPriceNetOk

`func (o *ModelInvoicePosResponse) GetPriceNetOk() (*string, bool)`

GetPriceNetOk returns a tuple with the PriceNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceNet

`func (o *ModelInvoicePosResponse) SetPriceNet(v string)`

SetPriceNet sets PriceNet field to given value.

### HasPriceNet

`func (o *ModelInvoicePosResponse) HasPriceNet() bool`

HasPriceNet returns a boolean if a field has been set.

### GetPriceGross

`func (o *ModelInvoicePosResponse) GetPriceGross() string`

GetPriceGross returns the PriceGross field if non-nil, zero value otherwise.

### GetPriceGrossOk

`func (o *ModelInvoicePosResponse) GetPriceGrossOk() (*string, bool)`

GetPriceGrossOk returns a tuple with the PriceGross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceGross

`func (o *ModelInvoicePosResponse) SetPriceGross(v string)`

SetPriceGross sets PriceGross field to given value.

### HasPriceGross

`func (o *ModelInvoicePosResponse) HasPriceGross() bool`

HasPriceGross returns a boolean if a field has been set.

### GetPriceTax

`func (o *ModelInvoicePosResponse) GetPriceTax() string`

GetPriceTax returns the PriceTax field if non-nil, zero value otherwise.

### GetPriceTaxOk

`func (o *ModelInvoicePosResponse) GetPriceTaxOk() (*string, bool)`

GetPriceTaxOk returns a tuple with the PriceTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceTax

`func (o *ModelInvoicePosResponse) SetPriceTax(v string)`

SetPriceTax sets PriceTax field to given value.

### HasPriceTax

`func (o *ModelInvoicePosResponse) HasPriceTax() bool`

HasPriceTax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


