# SaveInvoiceDiscountSave

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Discount** | **bool** | Defines if this is a discount or a surcharge | 
**Text** | **string** | A text for your discount | 
**Percentage** | **bool** | Defines if this is a percentage or an absolute discount | 
**Value** | **float32** | Value of the discount | 
**ObjectName** | **string** | Object name of the discount | 
**MapAll** | **bool** | Internal param | 

## Methods

### NewSaveInvoiceDiscountSave

`func NewSaveInvoiceDiscountSave(discount bool, text string, percentage bool, value float32, objectName string, mapAll bool, ) *SaveInvoiceDiscountSave`

NewSaveInvoiceDiscountSave instantiates a new SaveInvoiceDiscountSave object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveInvoiceDiscountSaveWithDefaults

`func NewSaveInvoiceDiscountSaveWithDefaults() *SaveInvoiceDiscountSave`

NewSaveInvoiceDiscountSaveWithDefaults instantiates a new SaveInvoiceDiscountSave object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscount

`func (o *SaveInvoiceDiscountSave) GetDiscount() bool`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *SaveInvoiceDiscountSave) GetDiscountOk() (*bool, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *SaveInvoiceDiscountSave) SetDiscount(v bool)`

SetDiscount sets Discount field to given value.


### GetText

`func (o *SaveInvoiceDiscountSave) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *SaveInvoiceDiscountSave) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *SaveInvoiceDiscountSave) SetText(v string)`

SetText sets Text field to given value.


### GetPercentage

`func (o *SaveInvoiceDiscountSave) GetPercentage() bool`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *SaveInvoiceDiscountSave) GetPercentageOk() (*bool, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *SaveInvoiceDiscountSave) SetPercentage(v bool)`

SetPercentage sets Percentage field to given value.


### GetValue

`func (o *SaveInvoiceDiscountSave) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SaveInvoiceDiscountSave) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SaveInvoiceDiscountSave) SetValue(v float32)`

SetValue sets Value field to given value.


### GetObjectName

`func (o *SaveInvoiceDiscountSave) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *SaveInvoiceDiscountSave) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *SaveInvoiceDiscountSave) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.


### GetMapAll

`func (o *SaveInvoiceDiscountSave) GetMapAll() bool`

GetMapAll returns the MapAll field if non-nil, zero value otherwise.

### GetMapAllOk

`func (o *SaveInvoiceDiscountSave) GetMapAllOk() (*bool, bool)`

GetMapAllOk returns a tuple with the MapAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapAll

`func (o *SaveInvoiceDiscountSave) SetMapAll(v bool)`

SetMapAll sets MapAll field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


