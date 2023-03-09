# ModelCreditNoteUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The creditNote id | [optional] [readonly] 
**ObjectName** | Pointer to **string** | The creditNote object name | [optional] [readonly] 
**Create** | Pointer to **time.Time** | Date of creditNote creation | [optional] [readonly] 
**Update** | Pointer to **time.Time** | Date of last creditNote update | [optional] [readonly] 
**CreditNoteNumber** | Pointer to **NullableString** | The creditNote number | [optional] 
**Contact** | Pointer to [**NullableModelCreditNoteUpdateContact**](ModelCreditNoteUpdateContact.md) |  | [optional] 
**CreditNoteDate** | Pointer to **NullableTime** | Needs to be provided as timestamp or dd.mm.yyyy | [optional] 
**Status** | Pointer to **string** | Please have a look in       &lt;a href&#x3D;&#39;https://api.sevdesk.de/#section/Types-and-status-of-credit-notes&#39;&gt;status of credit note&lt;/a&gt;      to see what the different status codes mean | [optional] 
**Header** | Pointer to **NullableString** | Normally consist of prefix plus the creditNote number | [optional] 
**HeadText** | Pointer to **NullableString** | Certain html tags can be used here to format your text | [optional] 
**FootText** | Pointer to **NullableString** | Certain html tags can be used here to format your text | [optional] 
**AddressCountry** | Pointer to [**NullableModelCreditNoteAddressCountry**](ModelCreditNoteAddressCountry.md) |  | [optional] 
**CreateUser** | Pointer to [**ModelCreditNoteCreateUser**](ModelCreditNoteCreateUser.md) |  | [optional] 
**SevClient** | Pointer to [**ModelCreditNoteSevClient**](ModelCreditNoteSevClient.md) |  | [optional] 
**DeliveryTerms** | Pointer to **NullableString** | Delivery terms of the creditNote | [optional] 
**DeliveryDate** | Pointer to **time.Time** | Timestamp. This can also be a date range if you also use the attribute deliveryDateUntil | [optional] 
**PaymentTerms** | Pointer to **NullableString** | Payment terms of the creditNote | [optional] 
**Version** | Pointer to **NullableInt32** | Version of the creditNote.&lt;br&gt;      Can be used if you have multiple drafts for the same creditNote.&lt;br&gt;      Should start with 0 | [optional] 
**SmallSettlement** | Pointer to **NullableBool** | Defines if the client uses the small settlement scheme.      If yes, the creditNote must not contain any vat | [optional] 
**ContactPerson** | Pointer to [**NullableModelCreditNoteUpdateContactPerson**](ModelCreditNoteUpdateContactPerson.md) |  | [optional] 
**TaxRate** | Pointer to **NullableFloat32** | Is overwritten by creditNote position tax rates | [optional] 
**TaxSet** | Pointer to [**NullableModelCreditNoteTaxSet**](ModelCreditNoteTaxSet.md) |  | [optional] 
**TaxText** | Pointer to **NullableString** | A common tax text would be &#39;Umsatzsteuer 19%&#39; | [optional] 
**TaxType** | Pointer to **NullableString** | Tax type of the creditNote. There are four tax types: 1. default - Umsatzsteuer ausweisen 2. eu - Steuerfreie innergemeinschaftliche Lieferung (Europäische Union) 3. noteu - Steuerschuldnerschaft des Leistungsempfängers (außerhalb EU, z. B. Schweiz) 4. custom - Using custom tax set 5. ss - Not subject to VAT according to §19 1 UStG Tax rates are heavily connected to the tax type used. | [optional] 
**CreditNoteType** | Pointer to **NullableString** | Type of the creditNote. For more information on the different types, check      &lt;a href&#x3D;&#39;https://api.sevdesk.de/#section/Types-and-status-of-credit-notes&#39;&gt;this&lt;/a&gt;   | [optional] 
**SendDate** | Pointer to **NullableTime** | The date the creditNote was sent to the customer | [optional] 
**Address** | Pointer to **NullableString** | Complete address of the recipient including name, street, city, zip and country.&lt;br&gt;       Line breaks can be used and will be displayed on the invoice pdf. | [optional] 
**Currency** | Pointer to **NullableString** | Currency used in the creditNote. Needs to be currency code according to ISO-4217 | [optional] 
**SumNet** | Pointer to **float32** | Net sum of the creditNote | [optional] [readonly] 
**SumTax** | Pointer to **float32** | Tax sum of the creditNote | [optional] [readonly] 
**SumGross** | Pointer to **float32** | Gross sum of the creditNote | [optional] [readonly] 
**SumDiscounts** | Pointer to **float32** | Sum of all discounts in the creditNote | [optional] [readonly] 
**SumNetForeignCurrency** | Pointer to **float32** | Net sum of the creditNote in the foreign currency | [optional] [readonly] 
**SumTaxForeignCurrency** | Pointer to **float32** | Tax sum of the creditNote in the foreign currency | [optional] [readonly] 
**SumGrossForeignCurrency** | Pointer to **float32** | Gross sum of the creditNote in the foreign currency | [optional] [readonly] 
**SumDiscountsForeignCurrency** | Pointer to **float32** | Discounts sum of the creditNote in the foreign currency | [optional] [readonly] 
**CustomerInternalNote** | Pointer to **NullableString** | Internal note of the customer. Contains data entered into field &#39;Referenz/Bestellnummer&#39; | [optional] 
**ShowNet** | Pointer to **bool** | If true, the net amount of each position will be shown on the creditNote. Otherwise gross amount | [optional] 
**SendType** | Pointer to **NullableString** | Type which was used to send the creditNote. IMPORTANT: Please refer to the creditNote section of the       *     API-Overview to understand how this attribute can be used before using it! | [optional] 

## Methods

### NewModelCreditNoteUpdate

`func NewModelCreditNoteUpdate() *ModelCreditNoteUpdate`

NewModelCreditNoteUpdate instantiates a new ModelCreditNoteUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelCreditNoteUpdateWithDefaults

`func NewModelCreditNoteUpdateWithDefaults() *ModelCreditNoteUpdate`

NewModelCreditNoteUpdateWithDefaults instantiates a new ModelCreditNoteUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelCreditNoteUpdate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelCreditNoteUpdate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelCreditNoteUpdate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelCreditNoteUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectName

`func (o *ModelCreditNoteUpdate) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ModelCreditNoteUpdate) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ModelCreditNoteUpdate) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ModelCreditNoteUpdate) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetCreate

`func (o *ModelCreditNoteUpdate) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ModelCreditNoteUpdate) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ModelCreditNoteUpdate) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ModelCreditNoteUpdate) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUpdate

`func (o *ModelCreditNoteUpdate) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ModelCreditNoteUpdate) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ModelCreditNoteUpdate) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ModelCreditNoteUpdate) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetCreditNoteNumber

`func (o *ModelCreditNoteUpdate) GetCreditNoteNumber() string`

GetCreditNoteNumber returns the CreditNoteNumber field if non-nil, zero value otherwise.

### GetCreditNoteNumberOk

`func (o *ModelCreditNoteUpdate) GetCreditNoteNumberOk() (*string, bool)`

GetCreditNoteNumberOk returns a tuple with the CreditNoteNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNoteNumber

`func (o *ModelCreditNoteUpdate) SetCreditNoteNumber(v string)`

SetCreditNoteNumber sets CreditNoteNumber field to given value.

### HasCreditNoteNumber

`func (o *ModelCreditNoteUpdate) HasCreditNoteNumber() bool`

HasCreditNoteNumber returns a boolean if a field has been set.

### SetCreditNoteNumberNil

`func (o *ModelCreditNoteUpdate) SetCreditNoteNumberNil(b bool)`

 SetCreditNoteNumberNil sets the value for CreditNoteNumber to be an explicit nil

### UnsetCreditNoteNumber
`func (o *ModelCreditNoteUpdate) UnsetCreditNoteNumber()`

UnsetCreditNoteNumber ensures that no value is present for CreditNoteNumber, not even an explicit nil
### GetContact

`func (o *ModelCreditNoteUpdate) GetContact() ModelCreditNoteUpdateContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ModelCreditNoteUpdate) GetContactOk() (*ModelCreditNoteUpdateContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ModelCreditNoteUpdate) SetContact(v ModelCreditNoteUpdateContact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *ModelCreditNoteUpdate) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *ModelCreditNoteUpdate) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *ModelCreditNoteUpdate) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetCreditNoteDate

`func (o *ModelCreditNoteUpdate) GetCreditNoteDate() time.Time`

GetCreditNoteDate returns the CreditNoteDate field if non-nil, zero value otherwise.

### GetCreditNoteDateOk

`func (o *ModelCreditNoteUpdate) GetCreditNoteDateOk() (*time.Time, bool)`

GetCreditNoteDateOk returns a tuple with the CreditNoteDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNoteDate

`func (o *ModelCreditNoteUpdate) SetCreditNoteDate(v time.Time)`

SetCreditNoteDate sets CreditNoteDate field to given value.

### HasCreditNoteDate

`func (o *ModelCreditNoteUpdate) HasCreditNoteDate() bool`

HasCreditNoteDate returns a boolean if a field has been set.

### SetCreditNoteDateNil

`func (o *ModelCreditNoteUpdate) SetCreditNoteDateNil(b bool)`

 SetCreditNoteDateNil sets the value for CreditNoteDate to be an explicit nil

### UnsetCreditNoteDate
`func (o *ModelCreditNoteUpdate) UnsetCreditNoteDate()`

UnsetCreditNoteDate ensures that no value is present for CreditNoteDate, not even an explicit nil
### GetStatus

`func (o *ModelCreditNoteUpdate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelCreditNoteUpdate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelCreditNoteUpdate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelCreditNoteUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetHeader

`func (o *ModelCreditNoteUpdate) GetHeader() string`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *ModelCreditNoteUpdate) GetHeaderOk() (*string, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *ModelCreditNoteUpdate) SetHeader(v string)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *ModelCreditNoteUpdate) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### SetHeaderNil

`func (o *ModelCreditNoteUpdate) SetHeaderNil(b bool)`

 SetHeaderNil sets the value for Header to be an explicit nil

### UnsetHeader
`func (o *ModelCreditNoteUpdate) UnsetHeader()`

UnsetHeader ensures that no value is present for Header, not even an explicit nil
### GetHeadText

`func (o *ModelCreditNoteUpdate) GetHeadText() string`

GetHeadText returns the HeadText field if non-nil, zero value otherwise.

### GetHeadTextOk

`func (o *ModelCreditNoteUpdate) GetHeadTextOk() (*string, bool)`

GetHeadTextOk returns a tuple with the HeadText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadText

`func (o *ModelCreditNoteUpdate) SetHeadText(v string)`

SetHeadText sets HeadText field to given value.

### HasHeadText

`func (o *ModelCreditNoteUpdate) HasHeadText() bool`

HasHeadText returns a boolean if a field has been set.

### SetHeadTextNil

`func (o *ModelCreditNoteUpdate) SetHeadTextNil(b bool)`

 SetHeadTextNil sets the value for HeadText to be an explicit nil

### UnsetHeadText
`func (o *ModelCreditNoteUpdate) UnsetHeadText()`

UnsetHeadText ensures that no value is present for HeadText, not even an explicit nil
### GetFootText

`func (o *ModelCreditNoteUpdate) GetFootText() string`

GetFootText returns the FootText field if non-nil, zero value otherwise.

### GetFootTextOk

`func (o *ModelCreditNoteUpdate) GetFootTextOk() (*string, bool)`

GetFootTextOk returns a tuple with the FootText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFootText

`func (o *ModelCreditNoteUpdate) SetFootText(v string)`

SetFootText sets FootText field to given value.

### HasFootText

`func (o *ModelCreditNoteUpdate) HasFootText() bool`

HasFootText returns a boolean if a field has been set.

### SetFootTextNil

`func (o *ModelCreditNoteUpdate) SetFootTextNil(b bool)`

 SetFootTextNil sets the value for FootText to be an explicit nil

### UnsetFootText
`func (o *ModelCreditNoteUpdate) UnsetFootText()`

UnsetFootText ensures that no value is present for FootText, not even an explicit nil
### GetAddressCountry

`func (o *ModelCreditNoteUpdate) GetAddressCountry() ModelCreditNoteAddressCountry`

GetAddressCountry returns the AddressCountry field if non-nil, zero value otherwise.

### GetAddressCountryOk

`func (o *ModelCreditNoteUpdate) GetAddressCountryOk() (*ModelCreditNoteAddressCountry, bool)`

GetAddressCountryOk returns a tuple with the AddressCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCountry

`func (o *ModelCreditNoteUpdate) SetAddressCountry(v ModelCreditNoteAddressCountry)`

SetAddressCountry sets AddressCountry field to given value.

### HasAddressCountry

`func (o *ModelCreditNoteUpdate) HasAddressCountry() bool`

HasAddressCountry returns a boolean if a field has been set.

### SetAddressCountryNil

`func (o *ModelCreditNoteUpdate) SetAddressCountryNil(b bool)`

 SetAddressCountryNil sets the value for AddressCountry to be an explicit nil

### UnsetAddressCountry
`func (o *ModelCreditNoteUpdate) UnsetAddressCountry()`

UnsetAddressCountry ensures that no value is present for AddressCountry, not even an explicit nil
### GetCreateUser

`func (o *ModelCreditNoteUpdate) GetCreateUser() ModelCreditNoteCreateUser`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *ModelCreditNoteUpdate) GetCreateUserOk() (*ModelCreditNoteCreateUser, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *ModelCreditNoteUpdate) SetCreateUser(v ModelCreditNoteCreateUser)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *ModelCreditNoteUpdate) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetSevClient

`func (o *ModelCreditNoteUpdate) GetSevClient() ModelCreditNoteSevClient`

GetSevClient returns the SevClient field if non-nil, zero value otherwise.

### GetSevClientOk

`func (o *ModelCreditNoteUpdate) GetSevClientOk() (*ModelCreditNoteSevClient, bool)`

GetSevClientOk returns a tuple with the SevClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSevClient

`func (o *ModelCreditNoteUpdate) SetSevClient(v ModelCreditNoteSevClient)`

SetSevClient sets SevClient field to given value.

### HasSevClient

`func (o *ModelCreditNoteUpdate) HasSevClient() bool`

HasSevClient returns a boolean if a field has been set.

### GetDeliveryTerms

`func (o *ModelCreditNoteUpdate) GetDeliveryTerms() string`

GetDeliveryTerms returns the DeliveryTerms field if non-nil, zero value otherwise.

### GetDeliveryTermsOk

`func (o *ModelCreditNoteUpdate) GetDeliveryTermsOk() (*string, bool)`

GetDeliveryTermsOk returns a tuple with the DeliveryTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTerms

`func (o *ModelCreditNoteUpdate) SetDeliveryTerms(v string)`

SetDeliveryTerms sets DeliveryTerms field to given value.

### HasDeliveryTerms

`func (o *ModelCreditNoteUpdate) HasDeliveryTerms() bool`

HasDeliveryTerms returns a boolean if a field has been set.

### SetDeliveryTermsNil

`func (o *ModelCreditNoteUpdate) SetDeliveryTermsNil(b bool)`

 SetDeliveryTermsNil sets the value for DeliveryTerms to be an explicit nil

### UnsetDeliveryTerms
`func (o *ModelCreditNoteUpdate) UnsetDeliveryTerms()`

UnsetDeliveryTerms ensures that no value is present for DeliveryTerms, not even an explicit nil
### GetDeliveryDate

`func (o *ModelCreditNoteUpdate) GetDeliveryDate() time.Time`

GetDeliveryDate returns the DeliveryDate field if non-nil, zero value otherwise.

### GetDeliveryDateOk

`func (o *ModelCreditNoteUpdate) GetDeliveryDateOk() (*time.Time, bool)`

GetDeliveryDateOk returns a tuple with the DeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDate

`func (o *ModelCreditNoteUpdate) SetDeliveryDate(v time.Time)`

SetDeliveryDate sets DeliveryDate field to given value.

### HasDeliveryDate

`func (o *ModelCreditNoteUpdate) HasDeliveryDate() bool`

HasDeliveryDate returns a boolean if a field has been set.

### GetPaymentTerms

`func (o *ModelCreditNoteUpdate) GetPaymentTerms() string`

GetPaymentTerms returns the PaymentTerms field if non-nil, zero value otherwise.

### GetPaymentTermsOk

`func (o *ModelCreditNoteUpdate) GetPaymentTermsOk() (*string, bool)`

GetPaymentTermsOk returns a tuple with the PaymentTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTerms

`func (o *ModelCreditNoteUpdate) SetPaymentTerms(v string)`

SetPaymentTerms sets PaymentTerms field to given value.

### HasPaymentTerms

`func (o *ModelCreditNoteUpdate) HasPaymentTerms() bool`

HasPaymentTerms returns a boolean if a field has been set.

### SetPaymentTermsNil

`func (o *ModelCreditNoteUpdate) SetPaymentTermsNil(b bool)`

 SetPaymentTermsNil sets the value for PaymentTerms to be an explicit nil

### UnsetPaymentTerms
`func (o *ModelCreditNoteUpdate) UnsetPaymentTerms()`

UnsetPaymentTerms ensures that no value is present for PaymentTerms, not even an explicit nil
### GetVersion

`func (o *ModelCreditNoteUpdate) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ModelCreditNoteUpdate) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ModelCreditNoteUpdate) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ModelCreditNoteUpdate) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ModelCreditNoteUpdate) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ModelCreditNoteUpdate) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetSmallSettlement

`func (o *ModelCreditNoteUpdate) GetSmallSettlement() bool`

GetSmallSettlement returns the SmallSettlement field if non-nil, zero value otherwise.

### GetSmallSettlementOk

`func (o *ModelCreditNoteUpdate) GetSmallSettlementOk() (*bool, bool)`

GetSmallSettlementOk returns a tuple with the SmallSettlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallSettlement

`func (o *ModelCreditNoteUpdate) SetSmallSettlement(v bool)`

SetSmallSettlement sets SmallSettlement field to given value.

### HasSmallSettlement

`func (o *ModelCreditNoteUpdate) HasSmallSettlement() bool`

HasSmallSettlement returns a boolean if a field has been set.

### SetSmallSettlementNil

`func (o *ModelCreditNoteUpdate) SetSmallSettlementNil(b bool)`

 SetSmallSettlementNil sets the value for SmallSettlement to be an explicit nil

### UnsetSmallSettlement
`func (o *ModelCreditNoteUpdate) UnsetSmallSettlement()`

UnsetSmallSettlement ensures that no value is present for SmallSettlement, not even an explicit nil
### GetContactPerson

`func (o *ModelCreditNoteUpdate) GetContactPerson() ModelCreditNoteUpdateContactPerson`

GetContactPerson returns the ContactPerson field if non-nil, zero value otherwise.

### GetContactPersonOk

`func (o *ModelCreditNoteUpdate) GetContactPersonOk() (*ModelCreditNoteUpdateContactPerson, bool)`

GetContactPersonOk returns a tuple with the ContactPerson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPerson

`func (o *ModelCreditNoteUpdate) SetContactPerson(v ModelCreditNoteUpdateContactPerson)`

SetContactPerson sets ContactPerson field to given value.

### HasContactPerson

`func (o *ModelCreditNoteUpdate) HasContactPerson() bool`

HasContactPerson returns a boolean if a field has been set.

### SetContactPersonNil

`func (o *ModelCreditNoteUpdate) SetContactPersonNil(b bool)`

 SetContactPersonNil sets the value for ContactPerson to be an explicit nil

### UnsetContactPerson
`func (o *ModelCreditNoteUpdate) UnsetContactPerson()`

UnsetContactPerson ensures that no value is present for ContactPerson, not even an explicit nil
### GetTaxRate

`func (o *ModelCreditNoteUpdate) GetTaxRate() float32`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *ModelCreditNoteUpdate) GetTaxRateOk() (*float32, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *ModelCreditNoteUpdate) SetTaxRate(v float32)`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *ModelCreditNoteUpdate) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### SetTaxRateNil

`func (o *ModelCreditNoteUpdate) SetTaxRateNil(b bool)`

 SetTaxRateNil sets the value for TaxRate to be an explicit nil

### UnsetTaxRate
`func (o *ModelCreditNoteUpdate) UnsetTaxRate()`

UnsetTaxRate ensures that no value is present for TaxRate, not even an explicit nil
### GetTaxSet

`func (o *ModelCreditNoteUpdate) GetTaxSet() ModelCreditNoteTaxSet`

GetTaxSet returns the TaxSet field if non-nil, zero value otherwise.

### GetTaxSetOk

`func (o *ModelCreditNoteUpdate) GetTaxSetOk() (*ModelCreditNoteTaxSet, bool)`

GetTaxSetOk returns a tuple with the TaxSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxSet

`func (o *ModelCreditNoteUpdate) SetTaxSet(v ModelCreditNoteTaxSet)`

SetTaxSet sets TaxSet field to given value.

### HasTaxSet

`func (o *ModelCreditNoteUpdate) HasTaxSet() bool`

HasTaxSet returns a boolean if a field has been set.

### SetTaxSetNil

`func (o *ModelCreditNoteUpdate) SetTaxSetNil(b bool)`

 SetTaxSetNil sets the value for TaxSet to be an explicit nil

### UnsetTaxSet
`func (o *ModelCreditNoteUpdate) UnsetTaxSet()`

UnsetTaxSet ensures that no value is present for TaxSet, not even an explicit nil
### GetTaxText

`func (o *ModelCreditNoteUpdate) GetTaxText() string`

GetTaxText returns the TaxText field if non-nil, zero value otherwise.

### GetTaxTextOk

`func (o *ModelCreditNoteUpdate) GetTaxTextOk() (*string, bool)`

GetTaxTextOk returns a tuple with the TaxText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxText

`func (o *ModelCreditNoteUpdate) SetTaxText(v string)`

SetTaxText sets TaxText field to given value.

### HasTaxText

`func (o *ModelCreditNoteUpdate) HasTaxText() bool`

HasTaxText returns a boolean if a field has been set.

### SetTaxTextNil

`func (o *ModelCreditNoteUpdate) SetTaxTextNil(b bool)`

 SetTaxTextNil sets the value for TaxText to be an explicit nil

### UnsetTaxText
`func (o *ModelCreditNoteUpdate) UnsetTaxText()`

UnsetTaxText ensures that no value is present for TaxText, not even an explicit nil
### GetTaxType

`func (o *ModelCreditNoteUpdate) GetTaxType() string`

GetTaxType returns the TaxType field if non-nil, zero value otherwise.

### GetTaxTypeOk

`func (o *ModelCreditNoteUpdate) GetTaxTypeOk() (*string, bool)`

GetTaxTypeOk returns a tuple with the TaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxType

`func (o *ModelCreditNoteUpdate) SetTaxType(v string)`

SetTaxType sets TaxType field to given value.

### HasTaxType

`func (o *ModelCreditNoteUpdate) HasTaxType() bool`

HasTaxType returns a boolean if a field has been set.

### SetTaxTypeNil

`func (o *ModelCreditNoteUpdate) SetTaxTypeNil(b bool)`

 SetTaxTypeNil sets the value for TaxType to be an explicit nil

### UnsetTaxType
`func (o *ModelCreditNoteUpdate) UnsetTaxType()`

UnsetTaxType ensures that no value is present for TaxType, not even an explicit nil
### GetCreditNoteType

`func (o *ModelCreditNoteUpdate) GetCreditNoteType() string`

GetCreditNoteType returns the CreditNoteType field if non-nil, zero value otherwise.

### GetCreditNoteTypeOk

`func (o *ModelCreditNoteUpdate) GetCreditNoteTypeOk() (*string, bool)`

GetCreditNoteTypeOk returns a tuple with the CreditNoteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNoteType

`func (o *ModelCreditNoteUpdate) SetCreditNoteType(v string)`

SetCreditNoteType sets CreditNoteType field to given value.

### HasCreditNoteType

`func (o *ModelCreditNoteUpdate) HasCreditNoteType() bool`

HasCreditNoteType returns a boolean if a field has been set.

### SetCreditNoteTypeNil

`func (o *ModelCreditNoteUpdate) SetCreditNoteTypeNil(b bool)`

 SetCreditNoteTypeNil sets the value for CreditNoteType to be an explicit nil

### UnsetCreditNoteType
`func (o *ModelCreditNoteUpdate) UnsetCreditNoteType()`

UnsetCreditNoteType ensures that no value is present for CreditNoteType, not even an explicit nil
### GetSendDate

`func (o *ModelCreditNoteUpdate) GetSendDate() time.Time`

GetSendDate returns the SendDate field if non-nil, zero value otherwise.

### GetSendDateOk

`func (o *ModelCreditNoteUpdate) GetSendDateOk() (*time.Time, bool)`

GetSendDateOk returns a tuple with the SendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendDate

`func (o *ModelCreditNoteUpdate) SetSendDate(v time.Time)`

SetSendDate sets SendDate field to given value.

### HasSendDate

`func (o *ModelCreditNoteUpdate) HasSendDate() bool`

HasSendDate returns a boolean if a field has been set.

### SetSendDateNil

`func (o *ModelCreditNoteUpdate) SetSendDateNil(b bool)`

 SetSendDateNil sets the value for SendDate to be an explicit nil

### UnsetSendDate
`func (o *ModelCreditNoteUpdate) UnsetSendDate()`

UnsetSendDate ensures that no value is present for SendDate, not even an explicit nil
### GetAddress

`func (o *ModelCreditNoteUpdate) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ModelCreditNoteUpdate) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ModelCreditNoteUpdate) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ModelCreditNoteUpdate) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *ModelCreditNoteUpdate) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *ModelCreditNoteUpdate) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetCurrency

`func (o *ModelCreditNoteUpdate) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ModelCreditNoteUpdate) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ModelCreditNoteUpdate) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ModelCreditNoteUpdate) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *ModelCreditNoteUpdate) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *ModelCreditNoteUpdate) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetSumNet

`func (o *ModelCreditNoteUpdate) GetSumNet() float32`

GetSumNet returns the SumNet field if non-nil, zero value otherwise.

### GetSumNetOk

`func (o *ModelCreditNoteUpdate) GetSumNetOk() (*float32, bool)`

GetSumNetOk returns a tuple with the SumNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumNet

`func (o *ModelCreditNoteUpdate) SetSumNet(v float32)`

SetSumNet sets SumNet field to given value.

### HasSumNet

`func (o *ModelCreditNoteUpdate) HasSumNet() bool`

HasSumNet returns a boolean if a field has been set.

### GetSumTax

`func (o *ModelCreditNoteUpdate) GetSumTax() float32`

GetSumTax returns the SumTax field if non-nil, zero value otherwise.

### GetSumTaxOk

`func (o *ModelCreditNoteUpdate) GetSumTaxOk() (*float32, bool)`

GetSumTaxOk returns a tuple with the SumTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumTax

`func (o *ModelCreditNoteUpdate) SetSumTax(v float32)`

SetSumTax sets SumTax field to given value.

### HasSumTax

`func (o *ModelCreditNoteUpdate) HasSumTax() bool`

HasSumTax returns a boolean if a field has been set.

### GetSumGross

`func (o *ModelCreditNoteUpdate) GetSumGross() float32`

GetSumGross returns the SumGross field if non-nil, zero value otherwise.

### GetSumGrossOk

`func (o *ModelCreditNoteUpdate) GetSumGrossOk() (*float32, bool)`

GetSumGrossOk returns a tuple with the SumGross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumGross

`func (o *ModelCreditNoteUpdate) SetSumGross(v float32)`

SetSumGross sets SumGross field to given value.

### HasSumGross

`func (o *ModelCreditNoteUpdate) HasSumGross() bool`

HasSumGross returns a boolean if a field has been set.

### GetSumDiscounts

`func (o *ModelCreditNoteUpdate) GetSumDiscounts() float32`

GetSumDiscounts returns the SumDiscounts field if non-nil, zero value otherwise.

### GetSumDiscountsOk

`func (o *ModelCreditNoteUpdate) GetSumDiscountsOk() (*float32, bool)`

GetSumDiscountsOk returns a tuple with the SumDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumDiscounts

`func (o *ModelCreditNoteUpdate) SetSumDiscounts(v float32)`

SetSumDiscounts sets SumDiscounts field to given value.

### HasSumDiscounts

`func (o *ModelCreditNoteUpdate) HasSumDiscounts() bool`

HasSumDiscounts returns a boolean if a field has been set.

### GetSumNetForeignCurrency

`func (o *ModelCreditNoteUpdate) GetSumNetForeignCurrency() float32`

GetSumNetForeignCurrency returns the SumNetForeignCurrency field if non-nil, zero value otherwise.

### GetSumNetForeignCurrencyOk

`func (o *ModelCreditNoteUpdate) GetSumNetForeignCurrencyOk() (*float32, bool)`

GetSumNetForeignCurrencyOk returns a tuple with the SumNetForeignCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumNetForeignCurrency

`func (o *ModelCreditNoteUpdate) SetSumNetForeignCurrency(v float32)`

SetSumNetForeignCurrency sets SumNetForeignCurrency field to given value.

### HasSumNetForeignCurrency

`func (o *ModelCreditNoteUpdate) HasSumNetForeignCurrency() bool`

HasSumNetForeignCurrency returns a boolean if a field has been set.

### GetSumTaxForeignCurrency

`func (o *ModelCreditNoteUpdate) GetSumTaxForeignCurrency() float32`

GetSumTaxForeignCurrency returns the SumTaxForeignCurrency field if non-nil, zero value otherwise.

### GetSumTaxForeignCurrencyOk

`func (o *ModelCreditNoteUpdate) GetSumTaxForeignCurrencyOk() (*float32, bool)`

GetSumTaxForeignCurrencyOk returns a tuple with the SumTaxForeignCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumTaxForeignCurrency

`func (o *ModelCreditNoteUpdate) SetSumTaxForeignCurrency(v float32)`

SetSumTaxForeignCurrency sets SumTaxForeignCurrency field to given value.

### HasSumTaxForeignCurrency

`func (o *ModelCreditNoteUpdate) HasSumTaxForeignCurrency() bool`

HasSumTaxForeignCurrency returns a boolean if a field has been set.

### GetSumGrossForeignCurrency

`func (o *ModelCreditNoteUpdate) GetSumGrossForeignCurrency() float32`

GetSumGrossForeignCurrency returns the SumGrossForeignCurrency field if non-nil, zero value otherwise.

### GetSumGrossForeignCurrencyOk

`func (o *ModelCreditNoteUpdate) GetSumGrossForeignCurrencyOk() (*float32, bool)`

GetSumGrossForeignCurrencyOk returns a tuple with the SumGrossForeignCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumGrossForeignCurrency

`func (o *ModelCreditNoteUpdate) SetSumGrossForeignCurrency(v float32)`

SetSumGrossForeignCurrency sets SumGrossForeignCurrency field to given value.

### HasSumGrossForeignCurrency

`func (o *ModelCreditNoteUpdate) HasSumGrossForeignCurrency() bool`

HasSumGrossForeignCurrency returns a boolean if a field has been set.

### GetSumDiscountsForeignCurrency

`func (o *ModelCreditNoteUpdate) GetSumDiscountsForeignCurrency() float32`

GetSumDiscountsForeignCurrency returns the SumDiscountsForeignCurrency field if non-nil, zero value otherwise.

### GetSumDiscountsForeignCurrencyOk

`func (o *ModelCreditNoteUpdate) GetSumDiscountsForeignCurrencyOk() (*float32, bool)`

GetSumDiscountsForeignCurrencyOk returns a tuple with the SumDiscountsForeignCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumDiscountsForeignCurrency

`func (o *ModelCreditNoteUpdate) SetSumDiscountsForeignCurrency(v float32)`

SetSumDiscountsForeignCurrency sets SumDiscountsForeignCurrency field to given value.

### HasSumDiscountsForeignCurrency

`func (o *ModelCreditNoteUpdate) HasSumDiscountsForeignCurrency() bool`

HasSumDiscountsForeignCurrency returns a boolean if a field has been set.

### GetCustomerInternalNote

`func (o *ModelCreditNoteUpdate) GetCustomerInternalNote() string`

GetCustomerInternalNote returns the CustomerInternalNote field if non-nil, zero value otherwise.

### GetCustomerInternalNoteOk

`func (o *ModelCreditNoteUpdate) GetCustomerInternalNoteOk() (*string, bool)`

GetCustomerInternalNoteOk returns a tuple with the CustomerInternalNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerInternalNote

`func (o *ModelCreditNoteUpdate) SetCustomerInternalNote(v string)`

SetCustomerInternalNote sets CustomerInternalNote field to given value.

### HasCustomerInternalNote

`func (o *ModelCreditNoteUpdate) HasCustomerInternalNote() bool`

HasCustomerInternalNote returns a boolean if a field has been set.

### SetCustomerInternalNoteNil

`func (o *ModelCreditNoteUpdate) SetCustomerInternalNoteNil(b bool)`

 SetCustomerInternalNoteNil sets the value for CustomerInternalNote to be an explicit nil

### UnsetCustomerInternalNote
`func (o *ModelCreditNoteUpdate) UnsetCustomerInternalNote()`

UnsetCustomerInternalNote ensures that no value is present for CustomerInternalNote, not even an explicit nil
### GetShowNet

`func (o *ModelCreditNoteUpdate) GetShowNet() bool`

GetShowNet returns the ShowNet field if non-nil, zero value otherwise.

### GetShowNetOk

`func (o *ModelCreditNoteUpdate) GetShowNetOk() (*bool, bool)`

GetShowNetOk returns a tuple with the ShowNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowNet

`func (o *ModelCreditNoteUpdate) SetShowNet(v bool)`

SetShowNet sets ShowNet field to given value.

### HasShowNet

`func (o *ModelCreditNoteUpdate) HasShowNet() bool`

HasShowNet returns a boolean if a field has been set.

### GetSendType

`func (o *ModelCreditNoteUpdate) GetSendType() string`

GetSendType returns the SendType field if non-nil, zero value otherwise.

### GetSendTypeOk

`func (o *ModelCreditNoteUpdate) GetSendTypeOk() (*string, bool)`

GetSendTypeOk returns a tuple with the SendType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendType

`func (o *ModelCreditNoteUpdate) SetSendType(v string)`

SetSendType sets SendType field to given value.

### HasSendType

`func (o *ModelCreditNoteUpdate) HasSendType() bool`

HasSendType returns a boolean if a field has been set.

### SetSendTypeNil

`func (o *ModelCreditNoteUpdate) SetSendTypeNil(b bool)`

 SetSendTypeNil sets the value for SendType to be an explicit nil

### UnsetSendType
`func (o *ModelCreditNoteUpdate) UnsetSendType()`

UnsetSendType ensures that no value is present for SendType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


