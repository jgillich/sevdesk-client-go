# SendInvoiceViaEMailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToEmail** | **string** | The recipient of the email. | 
**Subject** | **string** | The subject of the email. | 
**Text** | **string** | The text of the email. Can contain html. | 
**Copy** | Pointer to **bool** | Should a copy of this email be sent to you? | [optional] 
**AdditionalAttachments** | Pointer to **string** | Additional attachments to the mail. String of IDs of existing documents in your       *                      sevdesk account separated by &#39;,&#39; | [optional] 
**CcEmail** | Pointer to **string** | String of mail addresses to be put as cc separated by &#39;,&#39; | [optional] 
**BccEmail** | Pointer to **string** | String of mail addresses to be put as bcc separated by &#39;,&#39; | [optional] 

## Methods

### NewSendInvoiceViaEMailRequest

`func NewSendInvoiceViaEMailRequest(toEmail string, subject string, text string, ) *SendInvoiceViaEMailRequest`

NewSendInvoiceViaEMailRequest instantiates a new SendInvoiceViaEMailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendInvoiceViaEMailRequestWithDefaults

`func NewSendInvoiceViaEMailRequestWithDefaults() *SendInvoiceViaEMailRequest`

NewSendInvoiceViaEMailRequestWithDefaults instantiates a new SendInvoiceViaEMailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToEmail

`func (o *SendInvoiceViaEMailRequest) GetToEmail() string`

GetToEmail returns the ToEmail field if non-nil, zero value otherwise.

### GetToEmailOk

`func (o *SendInvoiceViaEMailRequest) GetToEmailOk() (*string, bool)`

GetToEmailOk returns a tuple with the ToEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToEmail

`func (o *SendInvoiceViaEMailRequest) SetToEmail(v string)`

SetToEmail sets ToEmail field to given value.


### GetSubject

`func (o *SendInvoiceViaEMailRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *SendInvoiceViaEMailRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *SendInvoiceViaEMailRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetText

`func (o *SendInvoiceViaEMailRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *SendInvoiceViaEMailRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *SendInvoiceViaEMailRequest) SetText(v string)`

SetText sets Text field to given value.


### GetCopy

`func (o *SendInvoiceViaEMailRequest) GetCopy() bool`

GetCopy returns the Copy field if non-nil, zero value otherwise.

### GetCopyOk

`func (o *SendInvoiceViaEMailRequest) GetCopyOk() (*bool, bool)`

GetCopyOk returns a tuple with the Copy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopy

`func (o *SendInvoiceViaEMailRequest) SetCopy(v bool)`

SetCopy sets Copy field to given value.

### HasCopy

`func (o *SendInvoiceViaEMailRequest) HasCopy() bool`

HasCopy returns a boolean if a field has been set.

### GetAdditionalAttachments

`func (o *SendInvoiceViaEMailRequest) GetAdditionalAttachments() string`

GetAdditionalAttachments returns the AdditionalAttachments field if non-nil, zero value otherwise.

### GetAdditionalAttachmentsOk

`func (o *SendInvoiceViaEMailRequest) GetAdditionalAttachmentsOk() (*string, bool)`

GetAdditionalAttachmentsOk returns a tuple with the AdditionalAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAttachments

`func (o *SendInvoiceViaEMailRequest) SetAdditionalAttachments(v string)`

SetAdditionalAttachments sets AdditionalAttachments field to given value.

### HasAdditionalAttachments

`func (o *SendInvoiceViaEMailRequest) HasAdditionalAttachments() bool`

HasAdditionalAttachments returns a boolean if a field has been set.

### GetCcEmail

`func (o *SendInvoiceViaEMailRequest) GetCcEmail() string`

GetCcEmail returns the CcEmail field if non-nil, zero value otherwise.

### GetCcEmailOk

`func (o *SendInvoiceViaEMailRequest) GetCcEmailOk() (*string, bool)`

GetCcEmailOk returns a tuple with the CcEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcEmail

`func (o *SendInvoiceViaEMailRequest) SetCcEmail(v string)`

SetCcEmail sets CcEmail field to given value.

### HasCcEmail

`func (o *SendInvoiceViaEMailRequest) HasCcEmail() bool`

HasCcEmail returns a boolean if a field has been set.

### GetBccEmail

`func (o *SendInvoiceViaEMailRequest) GetBccEmail() string`

GetBccEmail returns the BccEmail field if non-nil, zero value otherwise.

### GetBccEmailOk

`func (o *SendInvoiceViaEMailRequest) GetBccEmailOk() (*string, bool)`

GetBccEmailOk returns a tuple with the BccEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBccEmail

`func (o *SendInvoiceViaEMailRequest) SetBccEmail(v string)`

SetBccEmail sets BccEmail field to given value.

### HasBccEmail

`func (o *SendInvoiceViaEMailRequest) HasBccEmail() bool`

HasBccEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


