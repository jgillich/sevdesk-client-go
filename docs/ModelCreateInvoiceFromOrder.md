# ModelCreateInvoiceFromOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | [**ModelCreateInvoiceFromOrderOrder**](ModelCreateInvoiceFromOrderOrder.md) |  | 
**Type** | Pointer to **NullableString** | defines the type of amount | [optional] 
**Amount** | Pointer to **NullableFloat32** | Amount which has already been paid for this Invoice | [optional] 
**PartialType** | Pointer to **NullableString** | defines the type of the invoice 1. RE - Schlussrechnung 2. TR - Teilrechnung 3. AR - Abschlagsrechnung | [optional] 

## Methods

### NewModelCreateInvoiceFromOrder

`func NewModelCreateInvoiceFromOrder(order ModelCreateInvoiceFromOrderOrder, ) *ModelCreateInvoiceFromOrder`

NewModelCreateInvoiceFromOrder instantiates a new ModelCreateInvoiceFromOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelCreateInvoiceFromOrderWithDefaults

`func NewModelCreateInvoiceFromOrderWithDefaults() *ModelCreateInvoiceFromOrder`

NewModelCreateInvoiceFromOrderWithDefaults instantiates a new ModelCreateInvoiceFromOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *ModelCreateInvoiceFromOrder) GetOrder() ModelCreateInvoiceFromOrderOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ModelCreateInvoiceFromOrder) GetOrderOk() (*ModelCreateInvoiceFromOrderOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ModelCreateInvoiceFromOrder) SetOrder(v ModelCreateInvoiceFromOrderOrder)`

SetOrder sets Order field to given value.


### GetType

`func (o *ModelCreateInvoiceFromOrder) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelCreateInvoiceFromOrder) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelCreateInvoiceFromOrder) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ModelCreateInvoiceFromOrder) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ModelCreateInvoiceFromOrder) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ModelCreateInvoiceFromOrder) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAmount

`func (o *ModelCreateInvoiceFromOrder) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ModelCreateInvoiceFromOrder) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ModelCreateInvoiceFromOrder) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ModelCreateInvoiceFromOrder) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *ModelCreateInvoiceFromOrder) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *ModelCreateInvoiceFromOrder) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetPartialType

`func (o *ModelCreateInvoiceFromOrder) GetPartialType() string`

GetPartialType returns the PartialType field if non-nil, zero value otherwise.

### GetPartialTypeOk

`func (o *ModelCreateInvoiceFromOrder) GetPartialTypeOk() (*string, bool)`

GetPartialTypeOk returns a tuple with the PartialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialType

`func (o *ModelCreateInvoiceFromOrder) SetPartialType(v string)`

SetPartialType sets PartialType field to given value.

### HasPartialType

`func (o *ModelCreateInvoiceFromOrder) HasPartialType() bool`

HasPartialType returns a boolean if a field has been set.

### SetPartialTypeNil

`func (o *ModelCreateInvoiceFromOrder) SetPartialTypeNil(b bool)`

 SetPartialTypeNil sets the value for PartialType to be an explicit nil

### UnsetPartialType
`func (o *ModelCreateInvoiceFromOrder) UnsetPartialType()`

UnsetPartialType ensures that no value is present for PartialType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


