# SaveOrderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**ModelOrderResponse**](ModelOrderResponse.md) |  | [optional] 
**OrderPos** | Pointer to [**ModelOrderPosResponse**](ModelOrderPosResponse.md) |  | [optional] 

## Methods

### NewSaveOrderResponse

`func NewSaveOrderResponse() *SaveOrderResponse`

NewSaveOrderResponse instantiates a new SaveOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveOrderResponseWithDefaults

`func NewSaveOrderResponseWithDefaults() *SaveOrderResponse`

NewSaveOrderResponseWithDefaults instantiates a new SaveOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *SaveOrderResponse) GetOrder() ModelOrderResponse`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *SaveOrderResponse) GetOrderOk() (*ModelOrderResponse, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *SaveOrderResponse) SetOrder(v ModelOrderResponse)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *SaveOrderResponse) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetOrderPos

`func (o *SaveOrderResponse) GetOrderPos() ModelOrderPosResponse`

GetOrderPos returns the OrderPos field if non-nil, zero value otherwise.

### GetOrderPosOk

`func (o *SaveOrderResponse) GetOrderPosOk() (*ModelOrderPosResponse, bool)`

GetOrderPosOk returns a tuple with the OrderPos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderPos

`func (o *SaveOrderResponse) SetOrderPos(v ModelOrderPosResponse)`

SetOrderPos sets OrderPos field to given value.

### HasOrderPos

`func (o *SaveOrderResponse) HasOrderPos() bool`

HasOrderPos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


