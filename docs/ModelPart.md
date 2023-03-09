# ModelPart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The part id | [optional] [readonly] 
**ObjectName** | Pointer to **string** | The part object name | [optional] [readonly] 
**Create** | Pointer to **time.Time** | Date of part creation | [optional] [readonly] 
**Update** | Pointer to **time.Time** | Date of last part update | [optional] [readonly] 
**Name** | **string** | Name of the part | 
**PartNumber** | **string** | The part number | 
**Text** | Pointer to **NullableString** | A text describing the part | [optional] 
**Category** | Pointer to [**NullableModelPartCategory**](ModelPartCategory.md) |  | [optional] 
**Stock** | **float32** | The stock of the part | 
**StockEnabled** | Pointer to **bool** | Defines if the stock should be enabled | [optional] 
**Unity** | [**ModelPartUnity**](ModelPartUnity.md) |  | 
**Price** | Pointer to **NullableFloat32** | Net price for which the part is sold. we will change this parameter so that the gross price is calculated automatically, until then the priceGross parameter must be used. | [optional] 
**PriceNet** | Pointer to **NullableFloat32** | Net price for which the part is sold | [optional] 
**PriceGross** | Pointer to **NullableFloat32** | Gross price for which the part is sold | [optional] 
**SevClient** | Pointer to [**ModelPartSevClient**](ModelPartSevClient.md) |  | [optional] 
**PricePurchase** | Pointer to **NullableFloat32** | Purchase price of the part | [optional] 
**TaxRate** | **float32** | Tax rate of the part | 
**Status** | Pointer to **NullableInt32** | Status of the part. 50 &lt;-&gt; Inactive - 100 &lt;-&gt; Active | [optional] 
**InternalComment** | Pointer to **NullableString** | An internal comment for the part.&lt;br&gt;       Does not appear on invoices and orders. | [optional] 

## Methods

### NewModelPart

`func NewModelPart(name string, partNumber string, stock float32, unity ModelPartUnity, taxRate float32, ) *ModelPart`

NewModelPart instantiates a new ModelPart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelPartWithDefaults

`func NewModelPartWithDefaults() *ModelPart`

NewModelPartWithDefaults instantiates a new ModelPart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelPart) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelPart) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelPart) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelPart) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectName

`func (o *ModelPart) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ModelPart) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ModelPart) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ModelPart) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetCreate

`func (o *ModelPart) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ModelPart) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ModelPart) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ModelPart) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUpdate

`func (o *ModelPart) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ModelPart) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ModelPart) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ModelPart) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetName

`func (o *ModelPart) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelPart) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelPart) SetName(v string)`

SetName sets Name field to given value.


### GetPartNumber

`func (o *ModelPart) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *ModelPart) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *ModelPart) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.


### GetText

`func (o *ModelPart) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ModelPart) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ModelPart) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ModelPart) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *ModelPart) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *ModelPart) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetCategory

`func (o *ModelPart) GetCategory() ModelPartCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ModelPart) GetCategoryOk() (*ModelPartCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ModelPart) SetCategory(v ModelPartCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ModelPart) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *ModelPart) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *ModelPart) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetStock

`func (o *ModelPart) GetStock() float32`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *ModelPart) GetStockOk() (*float32, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *ModelPart) SetStock(v float32)`

SetStock sets Stock field to given value.


### GetStockEnabled

`func (o *ModelPart) GetStockEnabled() bool`

GetStockEnabled returns the StockEnabled field if non-nil, zero value otherwise.

### GetStockEnabledOk

`func (o *ModelPart) GetStockEnabledOk() (*bool, bool)`

GetStockEnabledOk returns a tuple with the StockEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockEnabled

`func (o *ModelPart) SetStockEnabled(v bool)`

SetStockEnabled sets StockEnabled field to given value.

### HasStockEnabled

`func (o *ModelPart) HasStockEnabled() bool`

HasStockEnabled returns a boolean if a field has been set.

### GetUnity

`func (o *ModelPart) GetUnity() ModelPartUnity`

GetUnity returns the Unity field if non-nil, zero value otherwise.

### GetUnityOk

`func (o *ModelPart) GetUnityOk() (*ModelPartUnity, bool)`

GetUnityOk returns a tuple with the Unity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnity

`func (o *ModelPart) SetUnity(v ModelPartUnity)`

SetUnity sets Unity field to given value.


### GetPrice

`func (o *ModelPart) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ModelPart) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ModelPart) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ModelPart) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *ModelPart) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *ModelPart) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetPriceNet

`func (o *ModelPart) GetPriceNet() float32`

GetPriceNet returns the PriceNet field if non-nil, zero value otherwise.

### GetPriceNetOk

`func (o *ModelPart) GetPriceNetOk() (*float32, bool)`

GetPriceNetOk returns a tuple with the PriceNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceNet

`func (o *ModelPart) SetPriceNet(v float32)`

SetPriceNet sets PriceNet field to given value.

### HasPriceNet

`func (o *ModelPart) HasPriceNet() bool`

HasPriceNet returns a boolean if a field has been set.

### SetPriceNetNil

`func (o *ModelPart) SetPriceNetNil(b bool)`

 SetPriceNetNil sets the value for PriceNet to be an explicit nil

### UnsetPriceNet
`func (o *ModelPart) UnsetPriceNet()`

UnsetPriceNet ensures that no value is present for PriceNet, not even an explicit nil
### GetPriceGross

`func (o *ModelPart) GetPriceGross() float32`

GetPriceGross returns the PriceGross field if non-nil, zero value otherwise.

### GetPriceGrossOk

`func (o *ModelPart) GetPriceGrossOk() (*float32, bool)`

GetPriceGrossOk returns a tuple with the PriceGross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceGross

`func (o *ModelPart) SetPriceGross(v float32)`

SetPriceGross sets PriceGross field to given value.

### HasPriceGross

`func (o *ModelPart) HasPriceGross() bool`

HasPriceGross returns a boolean if a field has been set.

### SetPriceGrossNil

`func (o *ModelPart) SetPriceGrossNil(b bool)`

 SetPriceGrossNil sets the value for PriceGross to be an explicit nil

### UnsetPriceGross
`func (o *ModelPart) UnsetPriceGross()`

UnsetPriceGross ensures that no value is present for PriceGross, not even an explicit nil
### GetSevClient

`func (o *ModelPart) GetSevClient() ModelPartSevClient`

GetSevClient returns the SevClient field if non-nil, zero value otherwise.

### GetSevClientOk

`func (o *ModelPart) GetSevClientOk() (*ModelPartSevClient, bool)`

GetSevClientOk returns a tuple with the SevClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSevClient

`func (o *ModelPart) SetSevClient(v ModelPartSevClient)`

SetSevClient sets SevClient field to given value.

### HasSevClient

`func (o *ModelPart) HasSevClient() bool`

HasSevClient returns a boolean if a field has been set.

### GetPricePurchase

`func (o *ModelPart) GetPricePurchase() float32`

GetPricePurchase returns the PricePurchase field if non-nil, zero value otherwise.

### GetPricePurchaseOk

`func (o *ModelPart) GetPricePurchaseOk() (*float32, bool)`

GetPricePurchaseOk returns a tuple with the PricePurchase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePurchase

`func (o *ModelPart) SetPricePurchase(v float32)`

SetPricePurchase sets PricePurchase field to given value.

### HasPricePurchase

`func (o *ModelPart) HasPricePurchase() bool`

HasPricePurchase returns a boolean if a field has been set.

### SetPricePurchaseNil

`func (o *ModelPart) SetPricePurchaseNil(b bool)`

 SetPricePurchaseNil sets the value for PricePurchase to be an explicit nil

### UnsetPricePurchase
`func (o *ModelPart) UnsetPricePurchase()`

UnsetPricePurchase ensures that no value is present for PricePurchase, not even an explicit nil
### GetTaxRate

`func (o *ModelPart) GetTaxRate() float32`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *ModelPart) GetTaxRateOk() (*float32, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *ModelPart) SetTaxRate(v float32)`

SetTaxRate sets TaxRate field to given value.


### GetStatus

`func (o *ModelPart) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelPart) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelPart) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelPart) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ModelPart) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ModelPart) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetInternalComment

`func (o *ModelPart) GetInternalComment() string`

GetInternalComment returns the InternalComment field if non-nil, zero value otherwise.

### GetInternalCommentOk

`func (o *ModelPart) GetInternalCommentOk() (*string, bool)`

GetInternalCommentOk returns a tuple with the InternalComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalComment

`func (o *ModelPart) SetInternalComment(v string)`

SetInternalComment sets InternalComment field to given value.

### HasInternalComment

`func (o *ModelPart) HasInternalComment() bool`

HasInternalComment returns a boolean if a field has been set.

### SetInternalCommentNil

`func (o *ModelPart) SetInternalCommentNil(b bool)`

 SetInternalCommentNil sets the value for InternalComment to be an explicit nil

### UnsetInternalComment
`func (o *ModelPart) UnsetInternalComment()`

UnsetInternalComment ensures that no value is present for InternalComment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


