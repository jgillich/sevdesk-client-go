# CreateTagRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the tag | [optional] 
**Object** | Pointer to [**CreateTagRequestObject**](CreateTagRequestObject.md) |  | [optional] 

## Methods

### NewCreateTagRequest

`func NewCreateTagRequest() *CreateTagRequest`

NewCreateTagRequest instantiates a new CreateTagRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTagRequestWithDefaults

`func NewCreateTagRequestWithDefaults() *CreateTagRequest`

NewCreateTagRequestWithDefaults instantiates a new CreateTagRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateTagRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTagRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTagRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateTagRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObject

`func (o *CreateTagRequest) GetObject() CreateTagRequestObject`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CreateTagRequest) GetObjectOk() (*CreateTagRequestObject, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CreateTagRequest) SetObject(v CreateTagRequestObject)`

SetObject sets Object field to given value.

### HasObject

`func (o *CreateTagRequest) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


