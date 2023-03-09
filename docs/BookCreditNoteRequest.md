# BookCreditNoteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | Amount which should be booked. Can also be a partial amount. | 
**Date** | **int32** | The booking date. Most likely the current date. | 
**Type** | **string** | Define a type for the booking.&lt;br&gt;      The following type abbreviations are available (abbreviation &lt;-&gt; meaning).&lt;br&gt;      &lt;ul&gt;      &lt;li&gt;N &lt;-&gt; Normal booking / partial booking&lt;/li&gt;      &lt;li&gt;CB &lt;-&gt; Reduced amount due to discount (skonto)&lt;/li&gt;      &lt;li&gt;CF &lt;-&gt; Reduced/Higher amount due to currency fluctuations&lt;/li&gt;      &lt;li&gt;O &lt;-&gt; Reduced/Higher amount due to other reasons&lt;/li&gt;      &lt;li&gt;OF &lt;-&gt; Higher amount due to reminder charges&lt;/li&gt;      &lt;li&gt;MTC &lt;-&gt; Reduced amount due to the monetary traffic costs&lt;/li&gt;      &lt;/ul&gt; | 
**CheckAccount** | [**BookVoucherRequestCheckAccount**](BookVoucherRequestCheckAccount.md) |  | 
**CheckAccountTransaction** | Pointer to [**BookCreditNoteRequestCheckAccountTransaction**](BookCreditNoteRequestCheckAccountTransaction.md) |  | [optional] 
**CreateFeed** | Pointer to **bool** | Determines if a feed is created for the booking process. | [optional] 

## Methods

### NewBookCreditNoteRequest

`func NewBookCreditNoteRequest(amount float32, date int32, type_ string, checkAccount BookVoucherRequestCheckAccount, ) *BookCreditNoteRequest`

NewBookCreditNoteRequest instantiates a new BookCreditNoteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookCreditNoteRequestWithDefaults

`func NewBookCreditNoteRequestWithDefaults() *BookCreditNoteRequest`

NewBookCreditNoteRequestWithDefaults instantiates a new BookCreditNoteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *BookCreditNoteRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BookCreditNoteRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BookCreditNoteRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetDate

`func (o *BookCreditNoteRequest) GetDate() int32`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *BookCreditNoteRequest) GetDateOk() (*int32, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *BookCreditNoteRequest) SetDate(v int32)`

SetDate sets Date field to given value.


### GetType

`func (o *BookCreditNoteRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BookCreditNoteRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BookCreditNoteRequest) SetType(v string)`

SetType sets Type field to given value.


### GetCheckAccount

`func (o *BookCreditNoteRequest) GetCheckAccount() BookVoucherRequestCheckAccount`

GetCheckAccount returns the CheckAccount field if non-nil, zero value otherwise.

### GetCheckAccountOk

`func (o *BookCreditNoteRequest) GetCheckAccountOk() (*BookVoucherRequestCheckAccount, bool)`

GetCheckAccountOk returns a tuple with the CheckAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckAccount

`func (o *BookCreditNoteRequest) SetCheckAccount(v BookVoucherRequestCheckAccount)`

SetCheckAccount sets CheckAccount field to given value.


### GetCheckAccountTransaction

`func (o *BookCreditNoteRequest) GetCheckAccountTransaction() BookCreditNoteRequestCheckAccountTransaction`

GetCheckAccountTransaction returns the CheckAccountTransaction field if non-nil, zero value otherwise.

### GetCheckAccountTransactionOk

`func (o *BookCreditNoteRequest) GetCheckAccountTransactionOk() (*BookCreditNoteRequestCheckAccountTransaction, bool)`

GetCheckAccountTransactionOk returns a tuple with the CheckAccountTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckAccountTransaction

`func (o *BookCreditNoteRequest) SetCheckAccountTransaction(v BookCreditNoteRequestCheckAccountTransaction)`

SetCheckAccountTransaction sets CheckAccountTransaction field to given value.

### HasCheckAccountTransaction

`func (o *BookCreditNoteRequest) HasCheckAccountTransaction() bool`

HasCheckAccountTransaction returns a boolean if a field has been set.

### GetCreateFeed

`func (o *BookCreditNoteRequest) GetCreateFeed() bool`

GetCreateFeed returns the CreateFeed field if non-nil, zero value otherwise.

### GetCreateFeedOk

`func (o *BookCreditNoteRequest) GetCreateFeedOk() (*bool, bool)`

GetCreateFeedOk returns a tuple with the CreateFeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateFeed

`func (o *BookCreditNoteRequest) SetCreateFeed(v bool)`

SetCreateFeed sets CreateFeed field to given value.

### HasCreateFeed

`func (o *BookCreditNoteRequest) HasCreateFeed() bool`

HasCreateFeed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


