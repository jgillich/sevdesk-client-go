# ModelContactAddressUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | Pointer to [**NullableModelContactAddressUpdateContact**](ModelContactAddressUpdateContact.md) |  | [optional] 
**Street** | Pointer to **NullableString** | Street name | [optional] 
**Zip** | Pointer to **NullableString** | Zib code | [optional] 
**City** | Pointer to **NullableString** | City name | [optional] 
**Country** | Pointer to [**NullableModelContactAddressUpdateCountry**](ModelContactAddressUpdateCountry.md) |  | [optional] 
**Category** | Pointer to [**NullableModelContactAddressCategory**](ModelContactAddressCategory.md) |  | [optional] 
**Name** | Pointer to **NullableString** | Name in address | [optional] 
**Name2** | Pointer to **string** | Second name in address | [optional] 
**Name3** | Pointer to **NullableString** | Third name in address | [optional] 
**Name4** | Pointer to **NullableString** | Fourth name in address | [optional] 

## Methods

### NewModelContactAddressUpdate

`func NewModelContactAddressUpdate() *ModelContactAddressUpdate`

NewModelContactAddressUpdate instantiates a new ModelContactAddressUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelContactAddressUpdateWithDefaults

`func NewModelContactAddressUpdateWithDefaults() *ModelContactAddressUpdate`

NewModelContactAddressUpdateWithDefaults instantiates a new ModelContactAddressUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContact

`func (o *ModelContactAddressUpdate) GetContact() ModelContactAddressUpdateContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ModelContactAddressUpdate) GetContactOk() (*ModelContactAddressUpdateContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ModelContactAddressUpdate) SetContact(v ModelContactAddressUpdateContact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *ModelContactAddressUpdate) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *ModelContactAddressUpdate) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *ModelContactAddressUpdate) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetStreet

`func (o *ModelContactAddressUpdate) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *ModelContactAddressUpdate) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *ModelContactAddressUpdate) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *ModelContactAddressUpdate) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### SetStreetNil

`func (o *ModelContactAddressUpdate) SetStreetNil(b bool)`

 SetStreetNil sets the value for Street to be an explicit nil

### UnsetStreet
`func (o *ModelContactAddressUpdate) UnsetStreet()`

UnsetStreet ensures that no value is present for Street, not even an explicit nil
### GetZip

`func (o *ModelContactAddressUpdate) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *ModelContactAddressUpdate) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *ModelContactAddressUpdate) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *ModelContactAddressUpdate) HasZip() bool`

HasZip returns a boolean if a field has been set.

### SetZipNil

`func (o *ModelContactAddressUpdate) SetZipNil(b bool)`

 SetZipNil sets the value for Zip to be an explicit nil

### UnsetZip
`func (o *ModelContactAddressUpdate) UnsetZip()`

UnsetZip ensures that no value is present for Zip, not even an explicit nil
### GetCity

`func (o *ModelContactAddressUpdate) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ModelContactAddressUpdate) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ModelContactAddressUpdate) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *ModelContactAddressUpdate) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *ModelContactAddressUpdate) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *ModelContactAddressUpdate) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetCountry

`func (o *ModelContactAddressUpdate) GetCountry() ModelContactAddressUpdateCountry`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ModelContactAddressUpdate) GetCountryOk() (*ModelContactAddressUpdateCountry, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ModelContactAddressUpdate) SetCountry(v ModelContactAddressUpdateCountry)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ModelContactAddressUpdate) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *ModelContactAddressUpdate) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *ModelContactAddressUpdate) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetCategory

`func (o *ModelContactAddressUpdate) GetCategory() ModelContactAddressCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ModelContactAddressUpdate) GetCategoryOk() (*ModelContactAddressCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ModelContactAddressUpdate) SetCategory(v ModelContactAddressCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ModelContactAddressUpdate) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *ModelContactAddressUpdate) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *ModelContactAddressUpdate) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetName

`func (o *ModelContactAddressUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelContactAddressUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelContactAddressUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelContactAddressUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ModelContactAddressUpdate) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ModelContactAddressUpdate) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetName2

`func (o *ModelContactAddressUpdate) GetName2() string`

GetName2 returns the Name2 field if non-nil, zero value otherwise.

### GetName2Ok

`func (o *ModelContactAddressUpdate) GetName2Ok() (*string, bool)`

GetName2Ok returns a tuple with the Name2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName2

`func (o *ModelContactAddressUpdate) SetName2(v string)`

SetName2 sets Name2 field to given value.

### HasName2

`func (o *ModelContactAddressUpdate) HasName2() bool`

HasName2 returns a boolean if a field has been set.

### GetName3

`func (o *ModelContactAddressUpdate) GetName3() string`

GetName3 returns the Name3 field if non-nil, zero value otherwise.

### GetName3Ok

`func (o *ModelContactAddressUpdate) GetName3Ok() (*string, bool)`

GetName3Ok returns a tuple with the Name3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName3

`func (o *ModelContactAddressUpdate) SetName3(v string)`

SetName3 sets Name3 field to given value.

### HasName3

`func (o *ModelContactAddressUpdate) HasName3() bool`

HasName3 returns a boolean if a field has been set.

### SetName3Nil

`func (o *ModelContactAddressUpdate) SetName3Nil(b bool)`

 SetName3Nil sets the value for Name3 to be an explicit nil

### UnsetName3
`func (o *ModelContactAddressUpdate) UnsetName3()`

UnsetName3 ensures that no value is present for Name3, not even an explicit nil
### GetName4

`func (o *ModelContactAddressUpdate) GetName4() string`

GetName4 returns the Name4 field if non-nil, zero value otherwise.

### GetName4Ok

`func (o *ModelContactAddressUpdate) GetName4Ok() (*string, bool)`

GetName4Ok returns a tuple with the Name4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName4

`func (o *ModelContactAddressUpdate) SetName4(v string)`

SetName4 sets Name4 field to given value.

### HasName4

`func (o *ModelContactAddressUpdate) HasName4() bool`

HasName4 returns a boolean if a field has been set.

### SetName4Nil

`func (o *ModelContactAddressUpdate) SetName4Nil(b bool)`

 SetName4Nil sets the value for Name4 to be an explicit nil

### UnsetName4
`func (o *ModelContactAddressUpdate) UnsetName4()`

UnsetName4 ensures that no value is present for Name4, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


