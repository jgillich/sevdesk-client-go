# ModelPartUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | The part id | [optional] [readonly] 
**ObjectName** | Pointer to **NullableString** | The part object name | [optional] [readonly] 
**Create** | Pointer to **time.Time** | Date of part creation | [optional] [readonly] 
**Update** | Pointer to **time.Time** | Date of last part update | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the part | [optional] 
**PartNumber** | Pointer to **string** | The part number | [optional] 
**Text** | Pointer to **NullableString** | A text describing the part | [optional] 
**Category** | Pointer to [**NullableModelPartCategory**](ModelPartCategory.md) |  | [optional] 
**Stock** | Pointer to **float32** | The stock of the part | [optional] 
**StockEnabled** | Pointer to **NullableBool** | Defines if the stock should be enabled | [optional] 
**Unity** | Pointer to [**ModelPartUnity**](ModelPartUnity.md) |  | [optional] 
**Price** | Pointer to **NullableFloat32** | Net price for which the part is sold. we will change this parameter so that the gross price is calculated automatically, until then the priceGross parameter must be used. | [optional] 
**PriceNet** | Pointer to **NullableFloat32** | Net price for which the part is sold | [optional] 
**PriceGross** | Pointer to **NullableFloat32** | Gross price for which the part is sold | [optional] 
**SevClient** | Pointer to [**ModelPartSevClient**](ModelPartSevClient.md) |  | [optional] 
**PricePurchase** | Pointer to **NullableFloat32** | Purchase price of the part | [optional] 
**TaxRate** | Pointer to **float32** | Tax rate of the part | [optional] 
**Status** | Pointer to **NullableInt32** | Status of the part. 50 &lt;-&gt; Inactive - 100 &lt;-&gt; Active | [optional] 
**InternalComment** | Pointer to **NullableString** | An internal comment for the part.&lt;br&gt;       Does not appear on invoices and orders. | [optional] 

## Methods

### NewModelPartUpdate

`func NewModelPartUpdate() *ModelPartUpdate`

NewModelPartUpdate instantiates a new ModelPartUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelPartUpdateWithDefaults

`func NewModelPartUpdateWithDefaults() *ModelPartUpdate`

NewModelPartUpdateWithDefaults instantiates a new ModelPartUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelPartUpdate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelPartUpdate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelPartUpdate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelPartUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ModelPartUpdate) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ModelPartUpdate) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectName

`func (o *ModelPartUpdate) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ModelPartUpdate) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ModelPartUpdate) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ModelPartUpdate) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### SetObjectNameNil

`func (o *ModelPartUpdate) SetObjectNameNil(b bool)`

 SetObjectNameNil sets the value for ObjectName to be an explicit nil

### UnsetObjectName
`func (o *ModelPartUpdate) UnsetObjectName()`

UnsetObjectName ensures that no value is present for ObjectName, not even an explicit nil
### GetCreate

`func (o *ModelPartUpdate) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ModelPartUpdate) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ModelPartUpdate) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ModelPartUpdate) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUpdate

`func (o *ModelPartUpdate) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ModelPartUpdate) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ModelPartUpdate) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ModelPartUpdate) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetName

`func (o *ModelPartUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelPartUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelPartUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelPartUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartNumber

`func (o *ModelPartUpdate) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *ModelPartUpdate) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *ModelPartUpdate) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *ModelPartUpdate) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetText

`func (o *ModelPartUpdate) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ModelPartUpdate) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ModelPartUpdate) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ModelPartUpdate) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *ModelPartUpdate) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *ModelPartUpdate) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetCategory

`func (o *ModelPartUpdate) GetCategory() ModelPartCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ModelPartUpdate) GetCategoryOk() (*ModelPartCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ModelPartUpdate) SetCategory(v ModelPartCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ModelPartUpdate) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *ModelPartUpdate) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *ModelPartUpdate) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetStock

`func (o *ModelPartUpdate) GetStock() float32`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *ModelPartUpdate) GetStockOk() (*float32, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *ModelPartUpdate) SetStock(v float32)`

SetStock sets Stock field to given value.

### HasStock

`func (o *ModelPartUpdate) HasStock() bool`

HasStock returns a boolean if a field has been set.

### GetStockEnabled

`func (o *ModelPartUpdate) GetStockEnabled() bool`

GetStockEnabled returns the StockEnabled field if non-nil, zero value otherwise.

### GetStockEnabledOk

`func (o *ModelPartUpdate) GetStockEnabledOk() (*bool, bool)`

GetStockEnabledOk returns a tuple with the StockEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockEnabled

`func (o *ModelPartUpdate) SetStockEnabled(v bool)`

SetStockEnabled sets StockEnabled field to given value.

### HasStockEnabled

`func (o *ModelPartUpdate) HasStockEnabled() bool`

HasStockEnabled returns a boolean if a field has been set.

### SetStockEnabledNil

`func (o *ModelPartUpdate) SetStockEnabledNil(b bool)`

 SetStockEnabledNil sets the value for StockEnabled to be an explicit nil

### UnsetStockEnabled
`func (o *ModelPartUpdate) UnsetStockEnabled()`

UnsetStockEnabled ensures that no value is present for StockEnabled, not even an explicit nil
### GetUnity

`func (o *ModelPartUpdate) GetUnity() ModelPartUnity`

GetUnity returns the Unity field if non-nil, zero value otherwise.

### GetUnityOk

`func (o *ModelPartUpdate) GetUnityOk() (*ModelPartUnity, bool)`

GetUnityOk returns a tuple with the Unity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnity

`func (o *ModelPartUpdate) SetUnity(v ModelPartUnity)`

SetUnity sets Unity field to given value.

### HasUnity

`func (o *ModelPartUpdate) HasUnity() bool`

HasUnity returns a boolean if a field has been set.

### GetPrice

`func (o *ModelPartUpdate) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ModelPartUpdate) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ModelPartUpdate) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ModelPartUpdate) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *ModelPartUpdate) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *ModelPartUpdate) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetPriceNet

`func (o *ModelPartUpdate) GetPriceNet() float32`

GetPriceNet returns the PriceNet field if non-nil, zero value otherwise.

### GetPriceNetOk

`func (o *ModelPartUpdate) GetPriceNetOk() (*float32, bool)`

GetPriceNetOk returns a tuple with the PriceNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceNet

`func (o *ModelPartUpdate) SetPriceNet(v float32)`

SetPriceNet sets PriceNet field to given value.

### HasPriceNet

`func (o *ModelPartUpdate) HasPriceNet() bool`

HasPriceNet returns a boolean if a field has been set.

### SetPriceNetNil

`func (o *ModelPartUpdate) SetPriceNetNil(b bool)`

 SetPriceNetNil sets the value for PriceNet to be an explicit nil

### UnsetPriceNet
`func (o *ModelPartUpdate) UnsetPriceNet()`

UnsetPriceNet ensures that no value is present for PriceNet, not even an explicit nil
### GetPriceGross

`func (o *ModelPartUpdate) GetPriceGross() float32`

GetPriceGross returns the PriceGross field if non-nil, zero value otherwise.

### GetPriceGrossOk

`func (o *ModelPartUpdate) GetPriceGrossOk() (*float32, bool)`

GetPriceGrossOk returns a tuple with the PriceGross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceGross

`func (o *ModelPartUpdate) SetPriceGross(v float32)`

SetPriceGross sets PriceGross field to given value.

### HasPriceGross

`func (o *ModelPartUpdate) HasPriceGross() bool`

HasPriceGross returns a boolean if a field has been set.

### SetPriceGrossNil

`func (o *ModelPartUpdate) SetPriceGrossNil(b bool)`

 SetPriceGrossNil sets the value for PriceGross to be an explicit nil

### UnsetPriceGross
`func (o *ModelPartUpdate) UnsetPriceGross()`

UnsetPriceGross ensures that no value is present for PriceGross, not even an explicit nil
### GetSevClient

`func (o *ModelPartUpdate) GetSevClient() ModelPartSevClient`

GetSevClient returns the SevClient field if non-nil, zero value otherwise.

### GetSevClientOk

`func (o *ModelPartUpdate) GetSevClientOk() (*ModelPartSevClient, bool)`

GetSevClientOk returns a tuple with the SevClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSevClient

`func (o *ModelPartUpdate) SetSevClient(v ModelPartSevClient)`

SetSevClient sets SevClient field to given value.

### HasSevClient

`func (o *ModelPartUpdate) HasSevClient() bool`

HasSevClient returns a boolean if a field has been set.

### GetPricePurchase

`func (o *ModelPartUpdate) GetPricePurchase() float32`

GetPricePurchase returns the PricePurchase field if non-nil, zero value otherwise.

### GetPricePurchaseOk

`func (o *ModelPartUpdate) GetPricePurchaseOk() (*float32, bool)`

GetPricePurchaseOk returns a tuple with the PricePurchase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePurchase

`func (o *ModelPartUpdate) SetPricePurchase(v float32)`

SetPricePurchase sets PricePurchase field to given value.

### HasPricePurchase

`func (o *ModelPartUpdate) HasPricePurchase() bool`

HasPricePurchase returns a boolean if a field has been set.

### SetPricePurchaseNil

`func (o *ModelPartUpdate) SetPricePurchaseNil(b bool)`

 SetPricePurchaseNil sets the value for PricePurchase to be an explicit nil

### UnsetPricePurchase
`func (o *ModelPartUpdate) UnsetPricePurchase()`

UnsetPricePurchase ensures that no value is present for PricePurchase, not even an explicit nil
### GetTaxRate

`func (o *ModelPartUpdate) GetTaxRate() float32`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *ModelPartUpdate) GetTaxRateOk() (*float32, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *ModelPartUpdate) SetTaxRate(v float32)`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *ModelPartUpdate) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### GetStatus

`func (o *ModelPartUpdate) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelPartUpdate) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelPartUpdate) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelPartUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ModelPartUpdate) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ModelPartUpdate) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetInternalComment

`func (o *ModelPartUpdate) GetInternalComment() string`

GetInternalComment returns the InternalComment field if non-nil, zero value otherwise.

### GetInternalCommentOk

`func (o *ModelPartUpdate) GetInternalCommentOk() (*string, bool)`

GetInternalCommentOk returns a tuple with the InternalComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalComment

`func (o *ModelPartUpdate) SetInternalComment(v string)`

SetInternalComment sets InternalComment field to given value.

### HasInternalComment

`func (o *ModelPartUpdate) HasInternalComment() bool`

HasInternalComment returns a boolean if a field has been set.

### SetInternalCommentNil

`func (o *ModelPartUpdate) SetInternalCommentNil(b bool)`

 SetInternalCommentNil sets the value for InternalComment to be an explicit nil

### UnsetInternalComment
`func (o *ModelPartUpdate) UnsetInternalComment()`

UnsetInternalComment ensures that no value is present for InternalComment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


