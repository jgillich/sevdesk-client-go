# ModelContactAddressResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The contact address id | [optional] [readonly] 
**ObjectName** | Pointer to **string** | The contact address object name | [optional] [readonly] 
**Create** | Pointer to **time.Time** | Date of contact address creation | [optional] [readonly] 
**Update** | Pointer to **time.Time** | Date of last contact address update | [optional] [readonly] 
**Contact** | [**ModelContactAddressContact**](ModelContactAddressContact.md) |  | 
**Street** | Pointer to **NullableString** | Street name | [optional] 
**Zip** | Pointer to **NullableString** | Zib code | [optional] 
**City** | Pointer to **NullableString** | City name | [optional] 
**Country** | [**ModelContactAddressCountry**](ModelContactAddressCountry.md) |  | 
**Category** | Pointer to [**NullableModelContactAddressCategory**](ModelContactAddressCategory.md) |  | [optional] 
**Name** | Pointer to **NullableString** | Name in address | [optional] 
**SevClient** | Pointer to [**ModelContactAddressSevClient**](ModelContactAddressSevClient.md) |  | [optional] 
**Name2** | Pointer to **string** | Second name in address | [optional] 
**Name3** | Pointer to **NullableString** | Third name in address | [optional] 
**Name4** | Pointer to **NullableString** | Fourth name in address | [optional] 

## Methods

### NewModelContactAddressResponse

`func NewModelContactAddressResponse(contact ModelContactAddressContact, country ModelContactAddressCountry, ) *ModelContactAddressResponse`

NewModelContactAddressResponse instantiates a new ModelContactAddressResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelContactAddressResponseWithDefaults

`func NewModelContactAddressResponseWithDefaults() *ModelContactAddressResponse`

NewModelContactAddressResponseWithDefaults instantiates a new ModelContactAddressResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelContactAddressResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelContactAddressResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelContactAddressResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelContactAddressResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectName

`func (o *ModelContactAddressResponse) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ModelContactAddressResponse) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ModelContactAddressResponse) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ModelContactAddressResponse) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetCreate

`func (o *ModelContactAddressResponse) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ModelContactAddressResponse) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ModelContactAddressResponse) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ModelContactAddressResponse) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUpdate

`func (o *ModelContactAddressResponse) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ModelContactAddressResponse) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ModelContactAddressResponse) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ModelContactAddressResponse) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetContact

`func (o *ModelContactAddressResponse) GetContact() ModelContactAddressContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ModelContactAddressResponse) GetContactOk() (*ModelContactAddressContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ModelContactAddressResponse) SetContact(v ModelContactAddressContact)`

SetContact sets Contact field to given value.


### GetStreet

`func (o *ModelContactAddressResponse) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *ModelContactAddressResponse) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *ModelContactAddressResponse) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *ModelContactAddressResponse) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### SetStreetNil

`func (o *ModelContactAddressResponse) SetStreetNil(b bool)`

 SetStreetNil sets the value for Street to be an explicit nil

### UnsetStreet
`func (o *ModelContactAddressResponse) UnsetStreet()`

UnsetStreet ensures that no value is present for Street, not even an explicit nil
### GetZip

`func (o *ModelContactAddressResponse) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *ModelContactAddressResponse) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *ModelContactAddressResponse) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *ModelContactAddressResponse) HasZip() bool`

HasZip returns a boolean if a field has been set.

### SetZipNil

`func (o *ModelContactAddressResponse) SetZipNil(b bool)`

 SetZipNil sets the value for Zip to be an explicit nil

### UnsetZip
`func (o *ModelContactAddressResponse) UnsetZip()`

UnsetZip ensures that no value is present for Zip, not even an explicit nil
### GetCity

`func (o *ModelContactAddressResponse) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ModelContactAddressResponse) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ModelContactAddressResponse) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *ModelContactAddressResponse) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *ModelContactAddressResponse) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *ModelContactAddressResponse) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetCountry

`func (o *ModelContactAddressResponse) GetCountry() ModelContactAddressCountry`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ModelContactAddressResponse) GetCountryOk() (*ModelContactAddressCountry, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ModelContactAddressResponse) SetCountry(v ModelContactAddressCountry)`

SetCountry sets Country field to given value.


### GetCategory

`func (o *ModelContactAddressResponse) GetCategory() ModelContactAddressCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ModelContactAddressResponse) GetCategoryOk() (*ModelContactAddressCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ModelContactAddressResponse) SetCategory(v ModelContactAddressCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ModelContactAddressResponse) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *ModelContactAddressResponse) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *ModelContactAddressResponse) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetName

`func (o *ModelContactAddressResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelContactAddressResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelContactAddressResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelContactAddressResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ModelContactAddressResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ModelContactAddressResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSevClient

`func (o *ModelContactAddressResponse) GetSevClient() ModelContactAddressSevClient`

GetSevClient returns the SevClient field if non-nil, zero value otherwise.

### GetSevClientOk

`func (o *ModelContactAddressResponse) GetSevClientOk() (*ModelContactAddressSevClient, bool)`

GetSevClientOk returns a tuple with the SevClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSevClient

`func (o *ModelContactAddressResponse) SetSevClient(v ModelContactAddressSevClient)`

SetSevClient sets SevClient field to given value.

### HasSevClient

`func (o *ModelContactAddressResponse) HasSevClient() bool`

HasSevClient returns a boolean if a field has been set.

### GetName2

`func (o *ModelContactAddressResponse) GetName2() string`

GetName2 returns the Name2 field if non-nil, zero value otherwise.

### GetName2Ok

`func (o *ModelContactAddressResponse) GetName2Ok() (*string, bool)`

GetName2Ok returns a tuple with the Name2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName2

`func (o *ModelContactAddressResponse) SetName2(v string)`

SetName2 sets Name2 field to given value.

### HasName2

`func (o *ModelContactAddressResponse) HasName2() bool`

HasName2 returns a boolean if a field has been set.

### GetName3

`func (o *ModelContactAddressResponse) GetName3() string`

GetName3 returns the Name3 field if non-nil, zero value otherwise.

### GetName3Ok

`func (o *ModelContactAddressResponse) GetName3Ok() (*string, bool)`

GetName3Ok returns a tuple with the Name3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName3

`func (o *ModelContactAddressResponse) SetName3(v string)`

SetName3 sets Name3 field to given value.

### HasName3

`func (o *ModelContactAddressResponse) HasName3() bool`

HasName3 returns a boolean if a field has been set.

### SetName3Nil

`func (o *ModelContactAddressResponse) SetName3Nil(b bool)`

 SetName3Nil sets the value for Name3 to be an explicit nil

### UnsetName3
`func (o *ModelContactAddressResponse) UnsetName3()`

UnsetName3 ensures that no value is present for Name3, not even an explicit nil
### GetName4

`func (o *ModelContactAddressResponse) GetName4() string`

GetName4 returns the Name4 field if non-nil, zero value otherwise.

### GetName4Ok

`func (o *ModelContactAddressResponse) GetName4Ok() (*string, bool)`

GetName4Ok returns a tuple with the Name4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName4

`func (o *ModelContactAddressResponse) SetName4(v string)`

SetName4 sets Name4 field to given value.

### HasName4

`func (o *ModelContactAddressResponse) HasName4() bool`

HasName4 returns a boolean if a field has been set.

### SetName4Nil

`func (o *ModelContactAddressResponse) SetName4Nil(b bool)`

 SetName4Nil sets the value for Name4 to be an explicit nil

### UnsetName4
`func (o *ModelContactAddressResponse) UnsetName4()`

UnsetName4 ensures that no value is present for Name4, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


