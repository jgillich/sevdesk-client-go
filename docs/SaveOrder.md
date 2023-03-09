# SaveOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | [**ModelOrder**](ModelOrder.md) |  | 
**OrderPosSave** | Pointer to [**ModelOrderPos**](ModelOrderPos.md) |  | [optional] 
**OrderPosDelete** | Pointer to [**SaveOrderOrderPosDelete**](SaveOrderOrderPosDelete.md) |  | [optional] [default to null]

## Methods

### NewSaveOrder

`func NewSaveOrder(order ModelOrder, ) *SaveOrder`

NewSaveOrder instantiates a new SaveOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveOrderWithDefaults

`func NewSaveOrderWithDefaults() *SaveOrder`

NewSaveOrderWithDefaults instantiates a new SaveOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *SaveOrder) GetOrder() ModelOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *SaveOrder) GetOrderOk() (*ModelOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *SaveOrder) SetOrder(v ModelOrder)`

SetOrder sets Order field to given value.


### GetOrderPosSave

`func (o *SaveOrder) GetOrderPosSave() ModelOrderPos`

GetOrderPosSave returns the OrderPosSave field if non-nil, zero value otherwise.

### GetOrderPosSaveOk

`func (o *SaveOrder) GetOrderPosSaveOk() (*ModelOrderPos, bool)`

GetOrderPosSaveOk returns a tuple with the OrderPosSave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderPosSave

`func (o *SaveOrder) SetOrderPosSave(v ModelOrderPos)`

SetOrderPosSave sets OrderPosSave field to given value.

### HasOrderPosSave

`func (o *SaveOrder) HasOrderPosSave() bool`

HasOrderPosSave returns a boolean if a field has been set.

### GetOrderPosDelete

`func (o *SaveOrder) GetOrderPosDelete() SaveOrderOrderPosDelete`

GetOrderPosDelete returns the OrderPosDelete field if non-nil, zero value otherwise.

### GetOrderPosDeleteOk

`func (o *SaveOrder) GetOrderPosDeleteOk() (*SaveOrderOrderPosDelete, bool)`

GetOrderPosDeleteOk returns a tuple with the OrderPosDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderPosDelete

`func (o *SaveOrder) SetOrderPosDelete(v SaveOrderOrderPosDelete)`

SetOrderPosDelete sets OrderPosDelete field to given value.

### HasOrderPosDelete

`func (o *SaveOrder) HasOrderPosDelete() bool`

HasOrderPosDelete returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


