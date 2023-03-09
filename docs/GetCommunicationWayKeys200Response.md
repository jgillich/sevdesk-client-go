# GetCommunicationWayKeys200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the communication way key 1. ID: 1 - Privat 2. ID: 2 - Arbeit 3. ID: 3 - Fax 4. ID: 4 - Mobil 5. ID: 5 - \&quot; \&quot; 6. ID: 6 - Autobox 7. ID: 7 - Newsletter 8. ID: 8 - Rechnungsadresse | [optional] 
**ObjectName** | Pointer to **string** | object name which is &#39;CommunicationWayKey&#39;. | [optional] 
**Create** | Pointer to **time.Time** | Date the communication way key was created | [optional] 
**Upadate** | Pointer to **time.Time** | Date the communication way key was last updated | [optional] 
**Name** | Pointer to **string** | Name of the communication way key | [optional] 
**TranslationCode** | Pointer to **string** |  | [optional] 

## Methods

### NewGetCommunicationWayKeys200Response

`func NewGetCommunicationWayKeys200Response() *GetCommunicationWayKeys200Response`

NewGetCommunicationWayKeys200Response instantiates a new GetCommunicationWayKeys200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCommunicationWayKeys200ResponseWithDefaults

`func NewGetCommunicationWayKeys200ResponseWithDefaults() *GetCommunicationWayKeys200Response`

NewGetCommunicationWayKeys200ResponseWithDefaults instantiates a new GetCommunicationWayKeys200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetCommunicationWayKeys200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCommunicationWayKeys200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCommunicationWayKeys200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetCommunicationWayKeys200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectName

`func (o *GetCommunicationWayKeys200Response) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *GetCommunicationWayKeys200Response) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *GetCommunicationWayKeys200Response) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *GetCommunicationWayKeys200Response) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetCreate

`func (o *GetCommunicationWayKeys200Response) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *GetCommunicationWayKeys200Response) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *GetCommunicationWayKeys200Response) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *GetCommunicationWayKeys200Response) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUpadate

`func (o *GetCommunicationWayKeys200Response) GetUpadate() time.Time`

GetUpadate returns the Upadate field if non-nil, zero value otherwise.

### GetUpadateOk

`func (o *GetCommunicationWayKeys200Response) GetUpadateOk() (*time.Time, bool)`

GetUpadateOk returns a tuple with the Upadate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpadate

`func (o *GetCommunicationWayKeys200Response) SetUpadate(v time.Time)`

SetUpadate sets Upadate field to given value.

### HasUpadate

`func (o *GetCommunicationWayKeys200Response) HasUpadate() bool`

HasUpadate returns a boolean if a field has been set.

### GetName

`func (o *GetCommunicationWayKeys200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCommunicationWayKeys200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCommunicationWayKeys200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetCommunicationWayKeys200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTranslationCode

`func (o *GetCommunicationWayKeys200Response) GetTranslationCode() string`

GetTranslationCode returns the TranslationCode field if non-nil, zero value otherwise.

### GetTranslationCodeOk

`func (o *GetCommunicationWayKeys200Response) GetTranslationCodeOk() (*string, bool)`

GetTranslationCodeOk returns a tuple with the TranslationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslationCode

`func (o *GetCommunicationWayKeys200Response) SetTranslationCode(v string)`

SetTranslationCode sets TranslationCode field to given value.

### HasTranslationCode

`func (o *GetCommunicationWayKeys200Response) HasTranslationCode() bool`

HasTranslationCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


