# ModelChangeLayoutResponseMetadaten

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Thumbs** | Pointer to **map[string]interface{}** | the pdf file | [optional] 
**Pages** | Pointer to **int32** | the number of pages in the document | [optional] 
**DocId** | Pointer to **string** | the id of the document | [optional] [readonly] 

## Methods

### NewModelChangeLayoutResponseMetadaten

`func NewModelChangeLayoutResponseMetadaten() *ModelChangeLayoutResponseMetadaten`

NewModelChangeLayoutResponseMetadaten instantiates a new ModelChangeLayoutResponseMetadaten object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelChangeLayoutResponseMetadatenWithDefaults

`func NewModelChangeLayoutResponseMetadatenWithDefaults() *ModelChangeLayoutResponseMetadaten`

NewModelChangeLayoutResponseMetadatenWithDefaults instantiates a new ModelChangeLayoutResponseMetadaten object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThumbs

`func (o *ModelChangeLayoutResponseMetadaten) GetThumbs() map[string]interface{}`

GetThumbs returns the Thumbs field if non-nil, zero value otherwise.

### GetThumbsOk

`func (o *ModelChangeLayoutResponseMetadaten) GetThumbsOk() (*map[string]interface{}, bool)`

GetThumbsOk returns a tuple with the Thumbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbs

`func (o *ModelChangeLayoutResponseMetadaten) SetThumbs(v map[string]interface{})`

SetThumbs sets Thumbs field to given value.

### HasThumbs

`func (o *ModelChangeLayoutResponseMetadaten) HasThumbs() bool`

HasThumbs returns a boolean if a field has been set.

### GetPages

`func (o *ModelChangeLayoutResponseMetadaten) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *ModelChangeLayoutResponseMetadaten) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *ModelChangeLayoutResponseMetadaten) SetPages(v int32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *ModelChangeLayoutResponseMetadaten) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetDocId

`func (o *ModelChangeLayoutResponseMetadaten) GetDocId() string`

GetDocId returns the DocId field if non-nil, zero value otherwise.

### GetDocIdOk

`func (o *ModelChangeLayoutResponseMetadaten) GetDocIdOk() (*string, bool)`

GetDocIdOk returns a tuple with the DocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocId

`func (o *ModelChangeLayoutResponseMetadaten) SetDocId(v string)`

SetDocId sets DocId field to given value.

### HasDocId

`func (o *ModelChangeLayoutResponseMetadaten) HasDocId() bool`

HasDocId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


