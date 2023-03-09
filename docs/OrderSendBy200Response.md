# OrderSendBy200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | Pointer to [**[]ModelOrderResponse**](ModelOrderResponse.md) | The order object which was marked as sent. | [optional] 

## Methods

### NewOrderSendBy200Response

`func NewOrderSendBy200Response() *OrderSendBy200Response`

NewOrderSendBy200Response instantiates a new OrderSendBy200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderSendBy200ResponseWithDefaults

`func NewOrderSendBy200ResponseWithDefaults() *OrderSendBy200Response`

NewOrderSendBy200ResponseWithDefaults instantiates a new OrderSendBy200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *OrderSendBy200Response) GetObjects() []ModelOrderResponse`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *OrderSendBy200Response) GetObjectsOk() (*[]ModelOrderResponse, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *OrderSendBy200Response) SetObjects(v []ModelOrderResponse)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *OrderSendBy200Response) HasObjects() bool`

HasObjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


