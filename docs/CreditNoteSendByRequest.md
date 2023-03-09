# CreditNoteSendByRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SendType** | **string** | Specifies the way in which the credit note was sent to the customer.&lt;br&gt;       Accepts &#39;VPR&#39; (print), &#39;VP&#39; (postal), &#39;VM&#39; (mail) and &#39;VPDF&#39; (downloaded pfd). | 
**SendDraft** | **bool** | Specify if the should be enshrined after marking it as sent. | 

## Methods

### NewCreditNoteSendByRequest

`func NewCreditNoteSendByRequest(sendType string, sendDraft bool, ) *CreditNoteSendByRequest`

NewCreditNoteSendByRequest instantiates a new CreditNoteSendByRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditNoteSendByRequestWithDefaults

`func NewCreditNoteSendByRequestWithDefaults() *CreditNoteSendByRequest`

NewCreditNoteSendByRequestWithDefaults instantiates a new CreditNoteSendByRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSendType

`func (o *CreditNoteSendByRequest) GetSendType() string`

GetSendType returns the SendType field if non-nil, zero value otherwise.

### GetSendTypeOk

`func (o *CreditNoteSendByRequest) GetSendTypeOk() (*string, bool)`

GetSendTypeOk returns a tuple with the SendType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendType

`func (o *CreditNoteSendByRequest) SetSendType(v string)`

SetSendType sets SendType field to given value.


### GetSendDraft

`func (o *CreditNoteSendByRequest) GetSendDraft() bool`

GetSendDraft returns the SendDraft field if non-nil, zero value otherwise.

### GetSendDraftOk

`func (o *CreditNoteSendByRequest) GetSendDraftOk() (*bool, bool)`

GetSendDraftOk returns a tuple with the SendDraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendDraft

`func (o *CreditNoteSendByRequest) SetSendDraft(v bool)`

SetSendDraft sets SendDraft field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


