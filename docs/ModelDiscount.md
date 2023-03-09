# ModelDiscount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | the id of the discount | [optional] [readonly] 
**ObjectName** | Pointer to **string** | Model name, which is &#39;Discounts&#39; | [optional] [readonly] 
**Create** | Pointer to **time.Time** | Date of discount creation | [optional] [readonly] 
**Update** | Pointer to **time.Time** | Date of last discount update | [optional] [readonly] 
**Object** | Pointer to [**ModelDiscountObject**](ModelDiscountObject.md) |  | [optional] 
**SevClient** | Pointer to **string** | Client to which invoice belongs. Will be filled automatically | [optional] [readonly] 
**Text** | Pointer to **string** | A text describing your position. | [optional] [readonly] 
**Percentage** | Pointer to **string** | Defines if this is a percentage or an absolute discount | [optional] 
**Value** | Pointer to **string** | Value of the discount | [optional] 
**IsNet** | Pointer to **string** | Defines is the Discount net or gross 0 - gross 1 - net | [optional] 

## Methods

### NewModelDiscount

`func NewModelDiscount() *ModelDiscount`

NewModelDiscount instantiates a new ModelDiscount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelDiscountWithDefaults

`func NewModelDiscountWithDefaults() *ModelDiscount`

NewModelDiscountWithDefaults instantiates a new ModelDiscount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelDiscount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelDiscount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelDiscount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelDiscount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectName

`func (o *ModelDiscount) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ModelDiscount) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ModelDiscount) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ModelDiscount) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetCreate

`func (o *ModelDiscount) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ModelDiscount) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ModelDiscount) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ModelDiscount) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUpdate

`func (o *ModelDiscount) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ModelDiscount) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ModelDiscount) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ModelDiscount) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetObject

`func (o *ModelDiscount) GetObject() ModelDiscountObject`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ModelDiscount) GetObjectOk() (*ModelDiscountObject, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ModelDiscount) SetObject(v ModelDiscountObject)`

SetObject sets Object field to given value.

### HasObject

`func (o *ModelDiscount) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetSevClient

`func (o *ModelDiscount) GetSevClient() string`

GetSevClient returns the SevClient field if non-nil, zero value otherwise.

### GetSevClientOk

`func (o *ModelDiscount) GetSevClientOk() (*string, bool)`

GetSevClientOk returns a tuple with the SevClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSevClient

`func (o *ModelDiscount) SetSevClient(v string)`

SetSevClient sets SevClient field to given value.

### HasSevClient

`func (o *ModelDiscount) HasSevClient() bool`

HasSevClient returns a boolean if a field has been set.

### GetText

`func (o *ModelDiscount) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ModelDiscount) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ModelDiscount) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ModelDiscount) HasText() bool`

HasText returns a boolean if a field has been set.

### GetPercentage

`func (o *ModelDiscount) GetPercentage() string`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *ModelDiscount) GetPercentageOk() (*string, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *ModelDiscount) SetPercentage(v string)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *ModelDiscount) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### GetValue

`func (o *ModelDiscount) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ModelDiscount) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ModelDiscount) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ModelDiscount) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetIsNet

`func (o *ModelDiscount) GetIsNet() string`

GetIsNet returns the IsNet field if non-nil, zero value otherwise.

### GetIsNetOk

`func (o *ModelDiscount) GetIsNetOk() (*string, bool)`

GetIsNetOk returns a tuple with the IsNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNet

`func (o *ModelDiscount) SetIsNet(v string)`

SetIsNet sets IsNet field to given value.

### HasIsNet

`func (o *ModelDiscount) HasIsNet() bool`

HasIsNet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


