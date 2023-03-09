# ModelContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The organization name.&lt;br&gt; Be aware that the type of contact will depend on this attribute.&lt;br&gt; If it holds a value, the contact will be regarded as an organization. | [optional] 
**Status** | Pointer to **NullableInt32** | Defines the status of the contact. 100 &lt;-&gt; Lead - 500 &lt;-&gt; Pending - 1000 &lt;-&gt; Active. | [optional] [default to 100]
**CustomerNumber** | Pointer to **NullableString** | The customer number | [optional] 
**Parent** | Pointer to [**NullableModelContactParent**](ModelContactParent.md) |  | [optional] 
**Surename** | Pointer to **NullableString** | The &lt;b&gt;first&lt;/b&gt; name of the contact.&lt;br&gt; Yeah... not quite right in literally every way. We know.&lt;br&gt; Not to be used for organizations. | [optional] 
**Familyname** | Pointer to **NullableString** | The last name of the contact.&lt;br&gt; Not to be used for organizations. | [optional] 
**Titel** | Pointer to **NullableString** | A non-academic title for the contact. Not to be used for organizations. | [optional] 
**Category** | [**ModelContactCategory**](ModelContactCategory.md) |  | 
**Description** | Pointer to **NullableString** | A description for the contact. | [optional] 
**AcademicTitle** | Pointer to **NullableString** | A academic title for the contact. Not to be used for organizations. | [optional] 
**Gender** | Pointer to **NullableString** | Gender of the contact.&lt;br&gt; Not to be used for organizations. | [optional] 
**Name2** | Pointer to **NullableString** | Second name of the contact.&lt;br&gt; Not to be used for organizations. | [optional] 
**Birthday** | Pointer to **NullableString** | Birthday of the contact.&lt;br&gt; Not to be used for organizations. | [optional] 
**VatNumber** | Pointer to **NullableString** | Vat number of the contact. | [optional] 
**BankAccount** | Pointer to **NullableString** | Bank account number (IBAN) of the contact. | [optional] 
**BankNumber** | Pointer to **NullableString** | Bank number of the bank used by the contact. | [optional] 
**DefaultCashbackTime** | Pointer to **NullableInt32** | Absolute time in days which the contact has to pay his invoices and subsequently get a cashback. | [optional] 
**DefaultCashbackPercent** | Pointer to **NullableFloat32** | Percentage of the invoice sum the contact gets back if he payed invoices in time. | [optional] 
**DefaultTimeToPay** | Pointer to **NullableInt32** | The payment goal in days which is set for every invoice of the contact. | [optional] 
**TaxNumber** | Pointer to **NullableString** | The tax number of the contact. | [optional] 
**TaxOffice** | Pointer to **NullableString** | The tax office of the contact (only for greek customers). | [optional] 
**ExemptVat** | Pointer to **NullableBool** | Defines if the contact is freed from paying vat. | [optional] 
**TaxType** | Pointer to **NullableString** | Defines which tax regulation the contact is using. | [optional] 
**TaxSet** | Pointer to [**NullableModelContactTaxSet**](ModelContactTaxSet.md) |  | [optional] 
**DefaultDiscountAmount** | Pointer to **NullableFloat32** | The default discount the contact gets for every invoice.&lt;br&gt; Depending on defaultDiscountPercentage attribute, in percent or absolute value. | [optional] 
**DefaultDiscountPercentage** | Pointer to **NullableBool** | Defines if the discount is a percentage (true) or an absolute value (false). | [optional] 
**BuyerReference** | Pointer to **NullableString** | Buyer reference of the contact. | [optional] 
**GovernmentAgency** | Pointer to **NullableBool** | Defines whether the contact is a government agency (true) or not (false). | [optional] 

## Methods

### NewModelContact

`func NewModelContact(category ModelContactCategory, ) *ModelContact`

NewModelContact instantiates a new ModelContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelContactWithDefaults

`func NewModelContactWithDefaults() *ModelContact`

NewModelContactWithDefaults instantiates a new ModelContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModelContact) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelContact) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelContact) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelContact) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ModelContact) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ModelContact) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStatus

`func (o *ModelContact) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelContact) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelContact) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelContact) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ModelContact) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ModelContact) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetCustomerNumber

`func (o *ModelContact) GetCustomerNumber() string`

GetCustomerNumber returns the CustomerNumber field if non-nil, zero value otherwise.

### GetCustomerNumberOk

`func (o *ModelContact) GetCustomerNumberOk() (*string, bool)`

GetCustomerNumberOk returns a tuple with the CustomerNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerNumber

`func (o *ModelContact) SetCustomerNumber(v string)`

SetCustomerNumber sets CustomerNumber field to given value.

### HasCustomerNumber

`func (o *ModelContact) HasCustomerNumber() bool`

HasCustomerNumber returns a boolean if a field has been set.

### SetCustomerNumberNil

`func (o *ModelContact) SetCustomerNumberNil(b bool)`

 SetCustomerNumberNil sets the value for CustomerNumber to be an explicit nil

### UnsetCustomerNumber
`func (o *ModelContact) UnsetCustomerNumber()`

UnsetCustomerNumber ensures that no value is present for CustomerNumber, not even an explicit nil
### GetParent

`func (o *ModelContact) GetParent() ModelContactParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ModelContact) GetParentOk() (*ModelContactParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ModelContact) SetParent(v ModelContactParent)`

SetParent sets Parent field to given value.

### HasParent

`func (o *ModelContact) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *ModelContact) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *ModelContact) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetSurename

`func (o *ModelContact) GetSurename() string`

GetSurename returns the Surename field if non-nil, zero value otherwise.

### GetSurenameOk

`func (o *ModelContact) GetSurenameOk() (*string, bool)`

GetSurenameOk returns a tuple with the Surename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurename

`func (o *ModelContact) SetSurename(v string)`

SetSurename sets Surename field to given value.

### HasSurename

`func (o *ModelContact) HasSurename() bool`

HasSurename returns a boolean if a field has been set.

### SetSurenameNil

`func (o *ModelContact) SetSurenameNil(b bool)`

 SetSurenameNil sets the value for Surename to be an explicit nil

### UnsetSurename
`func (o *ModelContact) UnsetSurename()`

UnsetSurename ensures that no value is present for Surename, not even an explicit nil
### GetFamilyname

`func (o *ModelContact) GetFamilyname() string`

GetFamilyname returns the Familyname field if non-nil, zero value otherwise.

### GetFamilynameOk

`func (o *ModelContact) GetFamilynameOk() (*string, bool)`

GetFamilynameOk returns a tuple with the Familyname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyname

`func (o *ModelContact) SetFamilyname(v string)`

SetFamilyname sets Familyname field to given value.

### HasFamilyname

`func (o *ModelContact) HasFamilyname() bool`

HasFamilyname returns a boolean if a field has been set.

### SetFamilynameNil

`func (o *ModelContact) SetFamilynameNil(b bool)`

 SetFamilynameNil sets the value for Familyname to be an explicit nil

### UnsetFamilyname
`func (o *ModelContact) UnsetFamilyname()`

UnsetFamilyname ensures that no value is present for Familyname, not even an explicit nil
### GetTitel

`func (o *ModelContact) GetTitel() string`

GetTitel returns the Titel field if non-nil, zero value otherwise.

### GetTitelOk

`func (o *ModelContact) GetTitelOk() (*string, bool)`

GetTitelOk returns a tuple with the Titel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitel

`func (o *ModelContact) SetTitel(v string)`

SetTitel sets Titel field to given value.

### HasTitel

`func (o *ModelContact) HasTitel() bool`

HasTitel returns a boolean if a field has been set.

### SetTitelNil

`func (o *ModelContact) SetTitelNil(b bool)`

 SetTitelNil sets the value for Titel to be an explicit nil

### UnsetTitel
`func (o *ModelContact) UnsetTitel()`

UnsetTitel ensures that no value is present for Titel, not even an explicit nil
### GetCategory

`func (o *ModelContact) GetCategory() ModelContactCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ModelContact) GetCategoryOk() (*ModelContactCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ModelContact) SetCategory(v ModelContactCategory)`

SetCategory sets Category field to given value.


### GetDescription

`func (o *ModelContact) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelContact) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelContact) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelContact) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ModelContact) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ModelContact) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAcademicTitle

`func (o *ModelContact) GetAcademicTitle() string`

GetAcademicTitle returns the AcademicTitle field if non-nil, zero value otherwise.

### GetAcademicTitleOk

`func (o *ModelContact) GetAcademicTitleOk() (*string, bool)`

GetAcademicTitleOk returns a tuple with the AcademicTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcademicTitle

`func (o *ModelContact) SetAcademicTitle(v string)`

SetAcademicTitle sets AcademicTitle field to given value.

### HasAcademicTitle

`func (o *ModelContact) HasAcademicTitle() bool`

HasAcademicTitle returns a boolean if a field has been set.

### SetAcademicTitleNil

`func (o *ModelContact) SetAcademicTitleNil(b bool)`

 SetAcademicTitleNil sets the value for AcademicTitle to be an explicit nil

### UnsetAcademicTitle
`func (o *ModelContact) UnsetAcademicTitle()`

UnsetAcademicTitle ensures that no value is present for AcademicTitle, not even an explicit nil
### GetGender

`func (o *ModelContact) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *ModelContact) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *ModelContact) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *ModelContact) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *ModelContact) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *ModelContact) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetName2

`func (o *ModelContact) GetName2() string`

GetName2 returns the Name2 field if non-nil, zero value otherwise.

### GetName2Ok

`func (o *ModelContact) GetName2Ok() (*string, bool)`

GetName2Ok returns a tuple with the Name2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName2

`func (o *ModelContact) SetName2(v string)`

SetName2 sets Name2 field to given value.

### HasName2

`func (o *ModelContact) HasName2() bool`

HasName2 returns a boolean if a field has been set.

### SetName2Nil

`func (o *ModelContact) SetName2Nil(b bool)`

 SetName2Nil sets the value for Name2 to be an explicit nil

### UnsetName2
`func (o *ModelContact) UnsetName2()`

UnsetName2 ensures that no value is present for Name2, not even an explicit nil
### GetBirthday

`func (o *ModelContact) GetBirthday() string`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *ModelContact) GetBirthdayOk() (*string, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *ModelContact) SetBirthday(v string)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *ModelContact) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### SetBirthdayNil

`func (o *ModelContact) SetBirthdayNil(b bool)`

 SetBirthdayNil sets the value for Birthday to be an explicit nil

### UnsetBirthday
`func (o *ModelContact) UnsetBirthday()`

UnsetBirthday ensures that no value is present for Birthday, not even an explicit nil
### GetVatNumber

`func (o *ModelContact) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *ModelContact) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *ModelContact) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *ModelContact) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.

### SetVatNumberNil

`func (o *ModelContact) SetVatNumberNil(b bool)`

 SetVatNumberNil sets the value for VatNumber to be an explicit nil

### UnsetVatNumber
`func (o *ModelContact) UnsetVatNumber()`

UnsetVatNumber ensures that no value is present for VatNumber, not even an explicit nil
### GetBankAccount

`func (o *ModelContact) GetBankAccount() string`

GetBankAccount returns the BankAccount field if non-nil, zero value otherwise.

### GetBankAccountOk

`func (o *ModelContact) GetBankAccountOk() (*string, bool)`

GetBankAccountOk returns a tuple with the BankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccount

`func (o *ModelContact) SetBankAccount(v string)`

SetBankAccount sets BankAccount field to given value.

### HasBankAccount

`func (o *ModelContact) HasBankAccount() bool`

HasBankAccount returns a boolean if a field has been set.

### SetBankAccountNil

`func (o *ModelContact) SetBankAccountNil(b bool)`

 SetBankAccountNil sets the value for BankAccount to be an explicit nil

### UnsetBankAccount
`func (o *ModelContact) UnsetBankAccount()`

UnsetBankAccount ensures that no value is present for BankAccount, not even an explicit nil
### GetBankNumber

`func (o *ModelContact) GetBankNumber() string`

GetBankNumber returns the BankNumber field if non-nil, zero value otherwise.

### GetBankNumberOk

`func (o *ModelContact) GetBankNumberOk() (*string, bool)`

GetBankNumberOk returns a tuple with the BankNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankNumber

`func (o *ModelContact) SetBankNumber(v string)`

SetBankNumber sets BankNumber field to given value.

### HasBankNumber

`func (o *ModelContact) HasBankNumber() bool`

HasBankNumber returns a boolean if a field has been set.

### SetBankNumberNil

`func (o *ModelContact) SetBankNumberNil(b bool)`

 SetBankNumberNil sets the value for BankNumber to be an explicit nil

### UnsetBankNumber
`func (o *ModelContact) UnsetBankNumber()`

UnsetBankNumber ensures that no value is present for BankNumber, not even an explicit nil
### GetDefaultCashbackTime

`func (o *ModelContact) GetDefaultCashbackTime() int32`

GetDefaultCashbackTime returns the DefaultCashbackTime field if non-nil, zero value otherwise.

### GetDefaultCashbackTimeOk

`func (o *ModelContact) GetDefaultCashbackTimeOk() (*int32, bool)`

GetDefaultCashbackTimeOk returns a tuple with the DefaultCashbackTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCashbackTime

`func (o *ModelContact) SetDefaultCashbackTime(v int32)`

SetDefaultCashbackTime sets DefaultCashbackTime field to given value.

### HasDefaultCashbackTime

`func (o *ModelContact) HasDefaultCashbackTime() bool`

HasDefaultCashbackTime returns a boolean if a field has been set.

### SetDefaultCashbackTimeNil

`func (o *ModelContact) SetDefaultCashbackTimeNil(b bool)`

 SetDefaultCashbackTimeNil sets the value for DefaultCashbackTime to be an explicit nil

### UnsetDefaultCashbackTime
`func (o *ModelContact) UnsetDefaultCashbackTime()`

UnsetDefaultCashbackTime ensures that no value is present for DefaultCashbackTime, not even an explicit nil
### GetDefaultCashbackPercent

`func (o *ModelContact) GetDefaultCashbackPercent() float32`

GetDefaultCashbackPercent returns the DefaultCashbackPercent field if non-nil, zero value otherwise.

### GetDefaultCashbackPercentOk

`func (o *ModelContact) GetDefaultCashbackPercentOk() (*float32, bool)`

GetDefaultCashbackPercentOk returns a tuple with the DefaultCashbackPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCashbackPercent

`func (o *ModelContact) SetDefaultCashbackPercent(v float32)`

SetDefaultCashbackPercent sets DefaultCashbackPercent field to given value.

### HasDefaultCashbackPercent

`func (o *ModelContact) HasDefaultCashbackPercent() bool`

HasDefaultCashbackPercent returns a boolean if a field has been set.

### SetDefaultCashbackPercentNil

`func (o *ModelContact) SetDefaultCashbackPercentNil(b bool)`

 SetDefaultCashbackPercentNil sets the value for DefaultCashbackPercent to be an explicit nil

### UnsetDefaultCashbackPercent
`func (o *ModelContact) UnsetDefaultCashbackPercent()`

UnsetDefaultCashbackPercent ensures that no value is present for DefaultCashbackPercent, not even an explicit nil
### GetDefaultTimeToPay

`func (o *ModelContact) GetDefaultTimeToPay() int32`

GetDefaultTimeToPay returns the DefaultTimeToPay field if non-nil, zero value otherwise.

### GetDefaultTimeToPayOk

`func (o *ModelContact) GetDefaultTimeToPayOk() (*int32, bool)`

GetDefaultTimeToPayOk returns a tuple with the DefaultTimeToPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTimeToPay

`func (o *ModelContact) SetDefaultTimeToPay(v int32)`

SetDefaultTimeToPay sets DefaultTimeToPay field to given value.

### HasDefaultTimeToPay

`func (o *ModelContact) HasDefaultTimeToPay() bool`

HasDefaultTimeToPay returns a boolean if a field has been set.

### SetDefaultTimeToPayNil

`func (o *ModelContact) SetDefaultTimeToPayNil(b bool)`

 SetDefaultTimeToPayNil sets the value for DefaultTimeToPay to be an explicit nil

### UnsetDefaultTimeToPay
`func (o *ModelContact) UnsetDefaultTimeToPay()`

UnsetDefaultTimeToPay ensures that no value is present for DefaultTimeToPay, not even an explicit nil
### GetTaxNumber

`func (o *ModelContact) GetTaxNumber() string`

GetTaxNumber returns the TaxNumber field if non-nil, zero value otherwise.

### GetTaxNumberOk

`func (o *ModelContact) GetTaxNumberOk() (*string, bool)`

GetTaxNumberOk returns a tuple with the TaxNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxNumber

`func (o *ModelContact) SetTaxNumber(v string)`

SetTaxNumber sets TaxNumber field to given value.

### HasTaxNumber

`func (o *ModelContact) HasTaxNumber() bool`

HasTaxNumber returns a boolean if a field has been set.

### SetTaxNumberNil

`func (o *ModelContact) SetTaxNumberNil(b bool)`

 SetTaxNumberNil sets the value for TaxNumber to be an explicit nil

### UnsetTaxNumber
`func (o *ModelContact) UnsetTaxNumber()`

UnsetTaxNumber ensures that no value is present for TaxNumber, not even an explicit nil
### GetTaxOffice

`func (o *ModelContact) GetTaxOffice() string`

GetTaxOffice returns the TaxOffice field if non-nil, zero value otherwise.

### GetTaxOfficeOk

`func (o *ModelContact) GetTaxOfficeOk() (*string, bool)`

GetTaxOfficeOk returns a tuple with the TaxOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxOffice

`func (o *ModelContact) SetTaxOffice(v string)`

SetTaxOffice sets TaxOffice field to given value.

### HasTaxOffice

`func (o *ModelContact) HasTaxOffice() bool`

HasTaxOffice returns a boolean if a field has been set.

### SetTaxOfficeNil

`func (o *ModelContact) SetTaxOfficeNil(b bool)`

 SetTaxOfficeNil sets the value for TaxOffice to be an explicit nil

### UnsetTaxOffice
`func (o *ModelContact) UnsetTaxOffice()`

UnsetTaxOffice ensures that no value is present for TaxOffice, not even an explicit nil
### GetExemptVat

`func (o *ModelContact) GetExemptVat() bool`

GetExemptVat returns the ExemptVat field if non-nil, zero value otherwise.

### GetExemptVatOk

`func (o *ModelContact) GetExemptVatOk() (*bool, bool)`

GetExemptVatOk returns a tuple with the ExemptVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptVat

`func (o *ModelContact) SetExemptVat(v bool)`

SetExemptVat sets ExemptVat field to given value.

### HasExemptVat

`func (o *ModelContact) HasExemptVat() bool`

HasExemptVat returns a boolean if a field has been set.

### SetExemptVatNil

`func (o *ModelContact) SetExemptVatNil(b bool)`

 SetExemptVatNil sets the value for ExemptVat to be an explicit nil

### UnsetExemptVat
`func (o *ModelContact) UnsetExemptVat()`

UnsetExemptVat ensures that no value is present for ExemptVat, not even an explicit nil
### GetTaxType

`func (o *ModelContact) GetTaxType() string`

GetTaxType returns the TaxType field if non-nil, zero value otherwise.

### GetTaxTypeOk

`func (o *ModelContact) GetTaxTypeOk() (*string, bool)`

GetTaxTypeOk returns a tuple with the TaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxType

`func (o *ModelContact) SetTaxType(v string)`

SetTaxType sets TaxType field to given value.

### HasTaxType

`func (o *ModelContact) HasTaxType() bool`

HasTaxType returns a boolean if a field has been set.

### SetTaxTypeNil

`func (o *ModelContact) SetTaxTypeNil(b bool)`

 SetTaxTypeNil sets the value for TaxType to be an explicit nil

### UnsetTaxType
`func (o *ModelContact) UnsetTaxType()`

UnsetTaxType ensures that no value is present for TaxType, not even an explicit nil
### GetTaxSet

`func (o *ModelContact) GetTaxSet() ModelContactTaxSet`

GetTaxSet returns the TaxSet field if non-nil, zero value otherwise.

### GetTaxSetOk

`func (o *ModelContact) GetTaxSetOk() (*ModelContactTaxSet, bool)`

GetTaxSetOk returns a tuple with the TaxSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxSet

`func (o *ModelContact) SetTaxSet(v ModelContactTaxSet)`

SetTaxSet sets TaxSet field to given value.

### HasTaxSet

`func (o *ModelContact) HasTaxSet() bool`

HasTaxSet returns a boolean if a field has been set.

### SetTaxSetNil

`func (o *ModelContact) SetTaxSetNil(b bool)`

 SetTaxSetNil sets the value for TaxSet to be an explicit nil

### UnsetTaxSet
`func (o *ModelContact) UnsetTaxSet()`

UnsetTaxSet ensures that no value is present for TaxSet, not even an explicit nil
### GetDefaultDiscountAmount

`func (o *ModelContact) GetDefaultDiscountAmount() float32`

GetDefaultDiscountAmount returns the DefaultDiscountAmount field if non-nil, zero value otherwise.

### GetDefaultDiscountAmountOk

`func (o *ModelContact) GetDefaultDiscountAmountOk() (*float32, bool)`

GetDefaultDiscountAmountOk returns a tuple with the DefaultDiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDiscountAmount

`func (o *ModelContact) SetDefaultDiscountAmount(v float32)`

SetDefaultDiscountAmount sets DefaultDiscountAmount field to given value.

### HasDefaultDiscountAmount

`func (o *ModelContact) HasDefaultDiscountAmount() bool`

HasDefaultDiscountAmount returns a boolean if a field has been set.

### SetDefaultDiscountAmountNil

`func (o *ModelContact) SetDefaultDiscountAmountNil(b bool)`

 SetDefaultDiscountAmountNil sets the value for DefaultDiscountAmount to be an explicit nil

### UnsetDefaultDiscountAmount
`func (o *ModelContact) UnsetDefaultDiscountAmount()`

UnsetDefaultDiscountAmount ensures that no value is present for DefaultDiscountAmount, not even an explicit nil
### GetDefaultDiscountPercentage

`func (o *ModelContact) GetDefaultDiscountPercentage() bool`

GetDefaultDiscountPercentage returns the DefaultDiscountPercentage field if non-nil, zero value otherwise.

### GetDefaultDiscountPercentageOk

`func (o *ModelContact) GetDefaultDiscountPercentageOk() (*bool, bool)`

GetDefaultDiscountPercentageOk returns a tuple with the DefaultDiscountPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDiscountPercentage

`func (o *ModelContact) SetDefaultDiscountPercentage(v bool)`

SetDefaultDiscountPercentage sets DefaultDiscountPercentage field to given value.

### HasDefaultDiscountPercentage

`func (o *ModelContact) HasDefaultDiscountPercentage() bool`

HasDefaultDiscountPercentage returns a boolean if a field has been set.

### SetDefaultDiscountPercentageNil

`func (o *ModelContact) SetDefaultDiscountPercentageNil(b bool)`

 SetDefaultDiscountPercentageNil sets the value for DefaultDiscountPercentage to be an explicit nil

### UnsetDefaultDiscountPercentage
`func (o *ModelContact) UnsetDefaultDiscountPercentage()`

UnsetDefaultDiscountPercentage ensures that no value is present for DefaultDiscountPercentage, not even an explicit nil
### GetBuyerReference

`func (o *ModelContact) GetBuyerReference() string`

GetBuyerReference returns the BuyerReference field if non-nil, zero value otherwise.

### GetBuyerReferenceOk

`func (o *ModelContact) GetBuyerReferenceOk() (*string, bool)`

GetBuyerReferenceOk returns a tuple with the BuyerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerReference

`func (o *ModelContact) SetBuyerReference(v string)`

SetBuyerReference sets BuyerReference field to given value.

### HasBuyerReference

`func (o *ModelContact) HasBuyerReference() bool`

HasBuyerReference returns a boolean if a field has been set.

### SetBuyerReferenceNil

`func (o *ModelContact) SetBuyerReferenceNil(b bool)`

 SetBuyerReferenceNil sets the value for BuyerReference to be an explicit nil

### UnsetBuyerReference
`func (o *ModelContact) UnsetBuyerReference()`

UnsetBuyerReference ensures that no value is present for BuyerReference, not even an explicit nil
### GetGovernmentAgency

`func (o *ModelContact) GetGovernmentAgency() bool`

GetGovernmentAgency returns the GovernmentAgency field if non-nil, zero value otherwise.

### GetGovernmentAgencyOk

`func (o *ModelContact) GetGovernmentAgencyOk() (*bool, bool)`

GetGovernmentAgencyOk returns a tuple with the GovernmentAgency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGovernmentAgency

`func (o *ModelContact) SetGovernmentAgency(v bool)`

SetGovernmentAgency sets GovernmentAgency field to given value.

### HasGovernmentAgency

`func (o *ModelContact) HasGovernmentAgency() bool`

HasGovernmentAgency returns a boolean if a field has been set.

### SetGovernmentAgencyNil

`func (o *ModelContact) SetGovernmentAgencyNil(b bool)`

 SetGovernmentAgencyNil sets the value for GovernmentAgency to be an explicit nil

### UnsetGovernmentAgency
`func (o *ModelContact) UnsetGovernmentAgency()`

UnsetGovernmentAgency ensures that no value is present for GovernmentAgency, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


