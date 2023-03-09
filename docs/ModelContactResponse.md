# ModelContactResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The contact id | [optional] [readonly] 
**ObjectName** | Pointer to **string** | The contact object name | [optional] [readonly] 
**Create** | Pointer to **time.Time** | Date of contact creation | [optional] [readonly] 
**Update** | Pointer to **time.Time** | Date of last contact update | [optional] [readonly] 
**Name** | Pointer to **string** | The organization name.&lt;br&gt; Be aware that the type of contact will depend on this attribute.&lt;br&gt; If it holds a value, the contact will be regarded as an organization. | [optional] [readonly] 
**Status** | Pointer to **string** | Defines the status of the contact. 100 &lt;-&gt; Lead - 500 &lt;-&gt; Pending - 1000 &lt;-&gt; Active. | [optional] [readonly] 
**CustomerNumber** | Pointer to **string** | The customer number | [optional] [readonly] 
**Parent** | Pointer to [**ModelContactResponseParent**](ModelContactResponseParent.md) |  | [optional] 
**Surename** | Pointer to **string** | The &lt;b&gt;first&lt;/b&gt; name of the contact.&lt;br&gt; Yeah... not quite right in literally every way. We know.&lt;br&gt; Not to be used for organizations. | [optional] [readonly] 
**Familyname** | Pointer to **string** | The last name of the contact.&lt;br&gt; Not to be used for organizations. | [optional] [readonly] 
**Titel** | Pointer to **string** | A non-academic title for the contact. Not to be used for organizations. | [optional] [readonly] 
**Category** | Pointer to [**ModelContactResponseCategory**](ModelContactResponseCategory.md) |  | [optional] 
**Description** | Pointer to **string** | A description for the contact. | [optional] [readonly] 
**AcademicTitle** | Pointer to **string** | A academic title for the contact. Not to be used for organizations. | [optional] [readonly] 
**Gender** | Pointer to **string** | Gender of the contact.&lt;br&gt; Not to be used for organizations. | [optional] [readonly] 
**SevClient** | Pointer to [**ModelContactResponseSevClient**](ModelContactResponseSevClient.md) |  | [optional] 
**Name2** | Pointer to **string** | Second name of the contact.&lt;br&gt; Not to be used for organizations. | [optional] [readonly] 
**Birthday** | Pointer to **string** | Birthday of the contact.&lt;br&gt; Not to be used for organizations. | [optional] [readonly] 
**VatNumber** | Pointer to **string** | Vat number of the contact. | [optional] [readonly] 
**BankAccount** | Pointer to **string** | Bank account number (IBAN) of the contact. | [optional] [readonly] 
**BankNumber** | Pointer to **string** | Bank number of the bank used by the contact. | [optional] [readonly] 
**DefaultCashbackTime** | Pointer to **string** | Absolute time in days which the contact has to pay his invoices and subsequently get a cashback. | [optional] [readonly] 
**DefaultCashbackPercent** | Pointer to **float32** | Percentage of the invoice sum the contact gets back if he payed invoices in time. | [optional] [readonly] 
**DefaultTimeToPay** | Pointer to **string** | The payment goal in days which is set for every invoice of the contact. | [optional] [readonly] 
**TaxNumber** | Pointer to **string** | The tax number of the contact. | [optional] [readonly] 
**TaxOffice** | Pointer to **string** | The tax office of the contact (only for greek customers). | [optional] [readonly] 
**ExemptVat** | Pointer to **string** | Defines if the contact is freed from paying vat. | [optional] [readonly] 
**TaxType** | Pointer to **string** | Defines which tax regulation the contact is using. | [optional] [readonly] 
**TaxSet** | Pointer to [**ModelContactResponseTaxSet**](ModelContactResponseTaxSet.md) |  | [optional] 
**DefaultDiscountAmount** | Pointer to **float32** | The default discount the contact gets for every invoice.&lt;br&gt; Depending on defaultDiscountPercentage attribute, in percent or absolute value. | [optional] [readonly] 
**DefaultDiscountPercentage** | Pointer to **string** | Defines if the discount is a percentage (true) or an absolute value (false). | [optional] [readonly] 
**BuyerReference** | Pointer to **string** | Buyer reference of the contact. | [optional] [readonly] 
**GovernmentAgency** | Pointer to **string** | Defines whether the contact is a government agency (true) or not (false). | [optional] [readonly] 
**AdditionalInformation** | Pointer to **string** | Additional information stored for the contact. | [optional] [readonly] 

## Methods

### NewModelContactResponse

`func NewModelContactResponse() *ModelContactResponse`

NewModelContactResponse instantiates a new ModelContactResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelContactResponseWithDefaults

`func NewModelContactResponseWithDefaults() *ModelContactResponse`

NewModelContactResponseWithDefaults instantiates a new ModelContactResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelContactResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelContactResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelContactResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelContactResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectName

`func (o *ModelContactResponse) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ModelContactResponse) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ModelContactResponse) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ModelContactResponse) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetCreate

`func (o *ModelContactResponse) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ModelContactResponse) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ModelContactResponse) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ModelContactResponse) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUpdate

`func (o *ModelContactResponse) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ModelContactResponse) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ModelContactResponse) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ModelContactResponse) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetName

`func (o *ModelContactResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelContactResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelContactResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelContactResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *ModelContactResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelContactResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelContactResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelContactResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCustomerNumber

`func (o *ModelContactResponse) GetCustomerNumber() string`

GetCustomerNumber returns the CustomerNumber field if non-nil, zero value otherwise.

### GetCustomerNumberOk

`func (o *ModelContactResponse) GetCustomerNumberOk() (*string, bool)`

GetCustomerNumberOk returns a tuple with the CustomerNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerNumber

`func (o *ModelContactResponse) SetCustomerNumber(v string)`

SetCustomerNumber sets CustomerNumber field to given value.

### HasCustomerNumber

`func (o *ModelContactResponse) HasCustomerNumber() bool`

HasCustomerNumber returns a boolean if a field has been set.

### GetParent

`func (o *ModelContactResponse) GetParent() ModelContactResponseParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ModelContactResponse) GetParentOk() (*ModelContactResponseParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ModelContactResponse) SetParent(v ModelContactResponseParent)`

SetParent sets Parent field to given value.

### HasParent

`func (o *ModelContactResponse) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetSurename

`func (o *ModelContactResponse) GetSurename() string`

GetSurename returns the Surename field if non-nil, zero value otherwise.

### GetSurenameOk

`func (o *ModelContactResponse) GetSurenameOk() (*string, bool)`

GetSurenameOk returns a tuple with the Surename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurename

`func (o *ModelContactResponse) SetSurename(v string)`

SetSurename sets Surename field to given value.

### HasSurename

`func (o *ModelContactResponse) HasSurename() bool`

HasSurename returns a boolean if a field has been set.

### GetFamilyname

`func (o *ModelContactResponse) GetFamilyname() string`

GetFamilyname returns the Familyname field if non-nil, zero value otherwise.

### GetFamilynameOk

`func (o *ModelContactResponse) GetFamilynameOk() (*string, bool)`

GetFamilynameOk returns a tuple with the Familyname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyname

`func (o *ModelContactResponse) SetFamilyname(v string)`

SetFamilyname sets Familyname field to given value.

### HasFamilyname

`func (o *ModelContactResponse) HasFamilyname() bool`

HasFamilyname returns a boolean if a field has been set.

### GetTitel

`func (o *ModelContactResponse) GetTitel() string`

GetTitel returns the Titel field if non-nil, zero value otherwise.

### GetTitelOk

`func (o *ModelContactResponse) GetTitelOk() (*string, bool)`

GetTitelOk returns a tuple with the Titel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitel

`func (o *ModelContactResponse) SetTitel(v string)`

SetTitel sets Titel field to given value.

### HasTitel

`func (o *ModelContactResponse) HasTitel() bool`

HasTitel returns a boolean if a field has been set.

### GetCategory

`func (o *ModelContactResponse) GetCategory() ModelContactResponseCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ModelContactResponse) GetCategoryOk() (*ModelContactResponseCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ModelContactResponse) SetCategory(v ModelContactResponseCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ModelContactResponse) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *ModelContactResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelContactResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelContactResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelContactResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAcademicTitle

`func (o *ModelContactResponse) GetAcademicTitle() string`

GetAcademicTitle returns the AcademicTitle field if non-nil, zero value otherwise.

### GetAcademicTitleOk

`func (o *ModelContactResponse) GetAcademicTitleOk() (*string, bool)`

GetAcademicTitleOk returns a tuple with the AcademicTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcademicTitle

`func (o *ModelContactResponse) SetAcademicTitle(v string)`

SetAcademicTitle sets AcademicTitle field to given value.

### HasAcademicTitle

`func (o *ModelContactResponse) HasAcademicTitle() bool`

HasAcademicTitle returns a boolean if a field has been set.

### GetGender

`func (o *ModelContactResponse) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *ModelContactResponse) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *ModelContactResponse) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *ModelContactResponse) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetSevClient

`func (o *ModelContactResponse) GetSevClient() ModelContactResponseSevClient`

GetSevClient returns the SevClient field if non-nil, zero value otherwise.

### GetSevClientOk

`func (o *ModelContactResponse) GetSevClientOk() (*ModelContactResponseSevClient, bool)`

GetSevClientOk returns a tuple with the SevClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSevClient

`func (o *ModelContactResponse) SetSevClient(v ModelContactResponseSevClient)`

SetSevClient sets SevClient field to given value.

### HasSevClient

`func (o *ModelContactResponse) HasSevClient() bool`

HasSevClient returns a boolean if a field has been set.

### GetName2

`func (o *ModelContactResponse) GetName2() string`

GetName2 returns the Name2 field if non-nil, zero value otherwise.

### GetName2Ok

`func (o *ModelContactResponse) GetName2Ok() (*string, bool)`

GetName2Ok returns a tuple with the Name2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName2

`func (o *ModelContactResponse) SetName2(v string)`

SetName2 sets Name2 field to given value.

### HasName2

`func (o *ModelContactResponse) HasName2() bool`

HasName2 returns a boolean if a field has been set.

### GetBirthday

`func (o *ModelContactResponse) GetBirthday() string`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *ModelContactResponse) GetBirthdayOk() (*string, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *ModelContactResponse) SetBirthday(v string)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *ModelContactResponse) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### GetVatNumber

`func (o *ModelContactResponse) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *ModelContactResponse) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *ModelContactResponse) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *ModelContactResponse) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.

### GetBankAccount

`func (o *ModelContactResponse) GetBankAccount() string`

GetBankAccount returns the BankAccount field if non-nil, zero value otherwise.

### GetBankAccountOk

`func (o *ModelContactResponse) GetBankAccountOk() (*string, bool)`

GetBankAccountOk returns a tuple with the BankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccount

`func (o *ModelContactResponse) SetBankAccount(v string)`

SetBankAccount sets BankAccount field to given value.

### HasBankAccount

`func (o *ModelContactResponse) HasBankAccount() bool`

HasBankAccount returns a boolean if a field has been set.

### GetBankNumber

`func (o *ModelContactResponse) GetBankNumber() string`

GetBankNumber returns the BankNumber field if non-nil, zero value otherwise.

### GetBankNumberOk

`func (o *ModelContactResponse) GetBankNumberOk() (*string, bool)`

GetBankNumberOk returns a tuple with the BankNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankNumber

`func (o *ModelContactResponse) SetBankNumber(v string)`

SetBankNumber sets BankNumber field to given value.

### HasBankNumber

`func (o *ModelContactResponse) HasBankNumber() bool`

HasBankNumber returns a boolean if a field has been set.

### GetDefaultCashbackTime

`func (o *ModelContactResponse) GetDefaultCashbackTime() string`

GetDefaultCashbackTime returns the DefaultCashbackTime field if non-nil, zero value otherwise.

### GetDefaultCashbackTimeOk

`func (o *ModelContactResponse) GetDefaultCashbackTimeOk() (*string, bool)`

GetDefaultCashbackTimeOk returns a tuple with the DefaultCashbackTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCashbackTime

`func (o *ModelContactResponse) SetDefaultCashbackTime(v string)`

SetDefaultCashbackTime sets DefaultCashbackTime field to given value.

### HasDefaultCashbackTime

`func (o *ModelContactResponse) HasDefaultCashbackTime() bool`

HasDefaultCashbackTime returns a boolean if a field has been set.

### GetDefaultCashbackPercent

`func (o *ModelContactResponse) GetDefaultCashbackPercent() float32`

GetDefaultCashbackPercent returns the DefaultCashbackPercent field if non-nil, zero value otherwise.

### GetDefaultCashbackPercentOk

`func (o *ModelContactResponse) GetDefaultCashbackPercentOk() (*float32, bool)`

GetDefaultCashbackPercentOk returns a tuple with the DefaultCashbackPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCashbackPercent

`func (o *ModelContactResponse) SetDefaultCashbackPercent(v float32)`

SetDefaultCashbackPercent sets DefaultCashbackPercent field to given value.

### HasDefaultCashbackPercent

`func (o *ModelContactResponse) HasDefaultCashbackPercent() bool`

HasDefaultCashbackPercent returns a boolean if a field has been set.

### GetDefaultTimeToPay

`func (o *ModelContactResponse) GetDefaultTimeToPay() string`

GetDefaultTimeToPay returns the DefaultTimeToPay field if non-nil, zero value otherwise.

### GetDefaultTimeToPayOk

`func (o *ModelContactResponse) GetDefaultTimeToPayOk() (*string, bool)`

GetDefaultTimeToPayOk returns a tuple with the DefaultTimeToPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTimeToPay

`func (o *ModelContactResponse) SetDefaultTimeToPay(v string)`

SetDefaultTimeToPay sets DefaultTimeToPay field to given value.

### HasDefaultTimeToPay

`func (o *ModelContactResponse) HasDefaultTimeToPay() bool`

HasDefaultTimeToPay returns a boolean if a field has been set.

### GetTaxNumber

`func (o *ModelContactResponse) GetTaxNumber() string`

GetTaxNumber returns the TaxNumber field if non-nil, zero value otherwise.

### GetTaxNumberOk

`func (o *ModelContactResponse) GetTaxNumberOk() (*string, bool)`

GetTaxNumberOk returns a tuple with the TaxNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxNumber

`func (o *ModelContactResponse) SetTaxNumber(v string)`

SetTaxNumber sets TaxNumber field to given value.

### HasTaxNumber

`func (o *ModelContactResponse) HasTaxNumber() bool`

HasTaxNumber returns a boolean if a field has been set.

### GetTaxOffice

`func (o *ModelContactResponse) GetTaxOffice() string`

GetTaxOffice returns the TaxOffice field if non-nil, zero value otherwise.

### GetTaxOfficeOk

`func (o *ModelContactResponse) GetTaxOfficeOk() (*string, bool)`

GetTaxOfficeOk returns a tuple with the TaxOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxOffice

`func (o *ModelContactResponse) SetTaxOffice(v string)`

SetTaxOffice sets TaxOffice field to given value.

### HasTaxOffice

`func (o *ModelContactResponse) HasTaxOffice() bool`

HasTaxOffice returns a boolean if a field has been set.

### GetExemptVat

`func (o *ModelContactResponse) GetExemptVat() string`

GetExemptVat returns the ExemptVat field if non-nil, zero value otherwise.

### GetExemptVatOk

`func (o *ModelContactResponse) GetExemptVatOk() (*string, bool)`

GetExemptVatOk returns a tuple with the ExemptVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptVat

`func (o *ModelContactResponse) SetExemptVat(v string)`

SetExemptVat sets ExemptVat field to given value.

### HasExemptVat

`func (o *ModelContactResponse) HasExemptVat() bool`

HasExemptVat returns a boolean if a field has been set.

### GetTaxType

`func (o *ModelContactResponse) GetTaxType() string`

GetTaxType returns the TaxType field if non-nil, zero value otherwise.

### GetTaxTypeOk

`func (o *ModelContactResponse) GetTaxTypeOk() (*string, bool)`

GetTaxTypeOk returns a tuple with the TaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxType

`func (o *ModelContactResponse) SetTaxType(v string)`

SetTaxType sets TaxType field to given value.

### HasTaxType

`func (o *ModelContactResponse) HasTaxType() bool`

HasTaxType returns a boolean if a field has been set.

### GetTaxSet

`func (o *ModelContactResponse) GetTaxSet() ModelContactResponseTaxSet`

GetTaxSet returns the TaxSet field if non-nil, zero value otherwise.

### GetTaxSetOk

`func (o *ModelContactResponse) GetTaxSetOk() (*ModelContactResponseTaxSet, bool)`

GetTaxSetOk returns a tuple with the TaxSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxSet

`func (o *ModelContactResponse) SetTaxSet(v ModelContactResponseTaxSet)`

SetTaxSet sets TaxSet field to given value.

### HasTaxSet

`func (o *ModelContactResponse) HasTaxSet() bool`

HasTaxSet returns a boolean if a field has been set.

### GetDefaultDiscountAmount

`func (o *ModelContactResponse) GetDefaultDiscountAmount() float32`

GetDefaultDiscountAmount returns the DefaultDiscountAmount field if non-nil, zero value otherwise.

### GetDefaultDiscountAmountOk

`func (o *ModelContactResponse) GetDefaultDiscountAmountOk() (*float32, bool)`

GetDefaultDiscountAmountOk returns a tuple with the DefaultDiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDiscountAmount

`func (o *ModelContactResponse) SetDefaultDiscountAmount(v float32)`

SetDefaultDiscountAmount sets DefaultDiscountAmount field to given value.

### HasDefaultDiscountAmount

`func (o *ModelContactResponse) HasDefaultDiscountAmount() bool`

HasDefaultDiscountAmount returns a boolean if a field has been set.

### GetDefaultDiscountPercentage

`func (o *ModelContactResponse) GetDefaultDiscountPercentage() string`

GetDefaultDiscountPercentage returns the DefaultDiscountPercentage field if non-nil, zero value otherwise.

### GetDefaultDiscountPercentageOk

`func (o *ModelContactResponse) GetDefaultDiscountPercentageOk() (*string, bool)`

GetDefaultDiscountPercentageOk returns a tuple with the DefaultDiscountPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDiscountPercentage

`func (o *ModelContactResponse) SetDefaultDiscountPercentage(v string)`

SetDefaultDiscountPercentage sets DefaultDiscountPercentage field to given value.

### HasDefaultDiscountPercentage

`func (o *ModelContactResponse) HasDefaultDiscountPercentage() bool`

HasDefaultDiscountPercentage returns a boolean if a field has been set.

### GetBuyerReference

`func (o *ModelContactResponse) GetBuyerReference() string`

GetBuyerReference returns the BuyerReference field if non-nil, zero value otherwise.

### GetBuyerReferenceOk

`func (o *ModelContactResponse) GetBuyerReferenceOk() (*string, bool)`

GetBuyerReferenceOk returns a tuple with the BuyerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerReference

`func (o *ModelContactResponse) SetBuyerReference(v string)`

SetBuyerReference sets BuyerReference field to given value.

### HasBuyerReference

`func (o *ModelContactResponse) HasBuyerReference() bool`

HasBuyerReference returns a boolean if a field has been set.

### GetGovernmentAgency

`func (o *ModelContactResponse) GetGovernmentAgency() string`

GetGovernmentAgency returns the GovernmentAgency field if non-nil, zero value otherwise.

### GetGovernmentAgencyOk

`func (o *ModelContactResponse) GetGovernmentAgencyOk() (*string, bool)`

GetGovernmentAgencyOk returns a tuple with the GovernmentAgency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGovernmentAgency

`func (o *ModelContactResponse) SetGovernmentAgency(v string)`

SetGovernmentAgency sets GovernmentAgency field to given value.

### HasGovernmentAgency

`func (o *ModelContactResponse) HasGovernmentAgency() bool`

HasGovernmentAgency returns a boolean if a field has been set.

### GetAdditionalInformation

`func (o *ModelContactResponse) GetAdditionalInformation() string`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *ModelContactResponse) GetAdditionalInformationOk() (*string, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *ModelContactResponse) SetAdditionalInformation(v string)`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *ModelContactResponse) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


