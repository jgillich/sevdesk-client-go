# BookVoucherRequestCheckAccountTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The id of the check account transaction on which should be booked. | 
**ObjectName** | **string** | Internal object name which is &#39;CheckAccountTransaction&#39;. | 

## Methods

### NewBookVoucherRequestCheckAccountTransaction

`func NewBookVoucherRequestCheckAccountTransaction(id int32, objectName string, ) *BookVoucherRequestCheckAccountTransaction`

NewBookVoucherRequestCheckAccountTransaction instantiates a new BookVoucherRequestCheckAccountTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookVoucherRequestCheckAccountTransactionWithDefaults

`func NewBookVoucherRequestCheckAccountTransactionWithDefaults() *BookVoucherRequestCheckAccountTransaction`

NewBookVoucherRequestCheckAccountTransactionWithDefaults instantiates a new BookVoucherRequestCheckAccountTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BookVoucherRequestCheckAccountTransaction) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BookVoucherRequestCheckAccountTransaction) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BookVoucherRequestCheckAccountTransaction) SetId(v int32)`

SetId sets Id field to given value.


### GetObjectName

`func (o *BookVoucherRequestCheckAccountTransaction) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *BookVoucherRequestCheckAccountTransaction) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *BookVoucherRequestCheckAccountTransaction) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


