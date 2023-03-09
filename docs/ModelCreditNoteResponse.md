# ModelCreditNoteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The creditNote id | [optional] [readonly] 
**ObjectName** | Pointer to **string** | The creditNote object name | [optional] [readonly] 
**Create** | Pointer to **time.Time** | Date of creditNote creation | [optional] [readonly] 
**Update** | Pointer to **time.Time** | Date of last creditNote update | [optional] [readonly] 
**CreditNoteNumber** | Pointer to **NullableString** | The creditNote number | [optional] 
**Contact** | Pointer to [**NullableModelCreditNoteResponseContact**](ModelCreditNoteResponseContact.md) |  | [optional] 
**CreditNoteDate** | Pointer to **NullableTime** | Needs to be provided as timestamp or dd.mm.yyyy | [optional] 
**Status** | Pointer to **string** | Please have a look in       &lt;a href&#x3D;&#39;https://api.sevdesk.de/#section/Types-and-status-of-credit-notes&#39;&gt;status of credit note&lt;/a&gt;      to see what the different status codes mean | [optional] 
**Header** | Pointer to **NullableString** | Normally consist of prefix plus the creditNote number | [optional] 
**HeadText** | Pointer to **NullableString** | Certain html tags can be used here to format your text | [optional] 
**FootText** | Pointer to **NullableString** | Certain html tags can be used here to format your text | [optional] 
**AddressCountry** | Pointer to [**NullableModelCreditNoteResponseAddressCountry**](ModelCreditNoteResponseAddressCountry.md) |  | [optional] 
**CreateUser** | Pointer to [**ModelInvoiceResponseCreateUser**](ModelInvoiceResponseCreateUser.md) |  | [optional] 
**SevClient** | Pointer to [**ModelCreditNoteResponseSevClient**](ModelCreditNoteResponseSevClient.md) |  | [optional] 
**DeliveryTerms** | Pointer to **NullableString** | Delivery terms of the creditNote | [optional] 
**DeliveryDate** | Pointer to **time.Time** | Timestamp. This can also be a date range if you also use the attribute deliveryDateUntil | [optional] 
**PaymentTerms** | Pointer to **NullableString** | Payment terms of the creditNote | [optional] 
**Version** | Pointer to **NullableString** | Version of the creditNote.&lt;br&gt;      Can be used if you have multiple drafts for the same creditNote.&lt;br&gt;      Should start with 0 | [optional] 
**SmallSettlement** | Pointer to **NullableBool** | Defines if the client uses the small settlement scheme.      If yes, the creditNote must not contain any vat | [optional] 
**ContactPerson** | Pointer to [**NullableModelCreditNoteResponseContactPerson**](ModelCreditNoteResponseContactPerson.md) |  | [optional] 
**TaxRate** | Pointer to **NullableString** | Is overwritten by creditNote position tax rates | [optional] 
**TaxSet** | Pointer to [**NullableModelCreditNoteResponseTaxSet**](ModelCreditNoteResponseTaxSet.md) |  | [optional] 
**TaxText** | Pointer to **NullableString** | A common tax text would be &#39;Umsatzsteuer 19%&#39; | [optional] 
**TaxType** | Pointer to **NullableString** | Tax type of the creditNote. There are four tax types: 1. default - Umsatzsteuer ausweisen 2. eu - Steuerfreie innergemeinschaftliche Lieferung (Europäische Union) 3. noteu - Steuerschuldnerschaft des Leistungsempfängers (außerhalb EU, z. B. Schweiz) 4. custom - Using custom tax set 5. ss - Not subject to VAT according to §19 1 UStG Tax rates are heavily connected to the tax type used. | [optional] 
**CreditNoteType** | Pointer to **NullableString** | Type of the creditNote. For more information on the different types, check      &lt;a href&#x3D;&#39;https://api.sevdesk.de/#section/Types-and-status-of-credit-notes&#39;&gt;this&lt;/a&gt;  . | [optional] 
**SendDate** | Pointer to **NullableTime** | The date the creditNote was sent to the customer | [optional] 
**Address** | Pointer to **NullableString** | Complete address of the recipient including name, street, city, zip and country.&lt;br&gt;       Line breaks can be used and will be displayed on the invoice pdf. | [optional] 
**Currency** | Pointer to **NullableString** | Currency used in the creditNote. Needs to be currency code according to ISO-4217 | [optional] 
**SumNet** | Pointer to **string** | Net sum of the creditNote | [optional] [readonly] 
**SumTax** | Pointer to **string** | Tax sum of the creditNote | [optional] [readonly] 
**SumGross** | Pointer to **string** | Gross sum of the creditNote | [optional] [readonly] 
**SumDiscounts** | Pointer to **string** | Sum of all discounts in the creditNote | [optional] [readonly] 
**SumNetForeignCurrency** | Pointer to **string** | Net sum of the creditNote in the foreign currency | [optional] [readonly] 
**SumTaxForeignCurrency** | Pointer to **string** | Tax sum of the creditNote in the foreign currency | [optional] [readonly] 
**SumGrossForeignCurrency** | Pointer to **string** | Gross sum of the creditNote in the foreign currency | [optional] [readonly] 
**SumDiscountsForeignCurrency** | Pointer to **string** | Discounts sum of the creditNote in the foreign currency | [optional] [readonly] 
**CustomerInternalNote** | Pointer to **NullableString** | Internal note of the customer. Contains data entered into field &#39;Referenz/Bestellnummer&#39; | [optional] 
**ShowNet** | Pointer to **bool** | If true, the net amount of each position will be shown on the creditNote. Otherwise gross amount | [optional] 
**SendType** | Pointer to **NullableString** | Type which was used to send the creditNote. IMPORTANT: Please refer to the creditNote section of the       *     API-Overview to understand how this attribute can be used before using it! | [optional] 

## Methods

### NewModelCreditNoteResponse

`func NewModelCreditNoteResponse() *ModelCreditNoteResponse`

NewModelCreditNoteResponse instantiates a new ModelCreditNoteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelCreditNoteResponseWithDefaults

`func NewModelCreditNoteResponseWithDefaults() *ModelCreditNoteResponse`

NewModelCreditNoteResponseWithDefaults instantiates a new ModelCreditNoteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelCreditNoteResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelCreditNoteResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelCreditNoteResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelCreditNoteResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectName

`func (o *ModelCreditNoteResponse) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ModelCreditNoteResponse) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ModelCreditNoteResponse) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ModelCreditNoteResponse) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetCreate

`func (o *ModelCreditNoteResponse) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ModelCreditNoteResponse) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ModelCreditNoteResponse) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ModelCreditNoteResponse) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUpdate

`func (o *ModelCreditNoteResponse) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ModelCreditNoteResponse) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ModelCreditNoteResponse) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ModelCreditNoteResponse) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetCreditNoteNumber

`func (o *ModelCreditNoteResponse) GetCreditNoteNumber() string`

GetCreditNoteNumber returns the CreditNoteNumber field if non-nil, zero value otherwise.

### GetCreditNoteNumberOk

`func (o *ModelCreditNoteResponse) GetCreditNoteNumberOk() (*string, bool)`

GetCreditNoteNumberOk returns a tuple with the CreditNoteNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNoteNumber

`func (o *ModelCreditNoteResponse) SetCreditNoteNumber(v string)`

SetCreditNoteNumber sets CreditNoteNumber field to given value.

### HasCreditNoteNumber

`func (o *ModelCreditNoteResponse) HasCreditNoteNumber() bool`

HasCreditNoteNumber returns a boolean if a field has been set.

### SetCreditNoteNumberNil

`func (o *ModelCreditNoteResponse) SetCreditNoteNumberNil(b bool)`

 SetCreditNoteNumberNil sets the value for CreditNoteNumber to be an explicit nil

### UnsetCreditNoteNumber
`func (o *ModelCreditNoteResponse) UnsetCreditNoteNumber()`

UnsetCreditNoteNumber ensures that no value is present for CreditNoteNumber, not even an explicit nil
### GetContact

`func (o *ModelCreditNoteResponse) GetContact() ModelCreditNoteResponseContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ModelCreditNoteResponse) GetContactOk() (*ModelCreditNoteResponseContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ModelCreditNoteResponse) SetContact(v ModelCreditNoteResponseContact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *ModelCreditNoteResponse) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *ModelCreditNoteResponse) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *ModelCreditNoteResponse) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetCreditNoteDate

`func (o *ModelCreditNoteResponse) GetCreditNoteDate() time.Time`

GetCreditNoteDate returns the CreditNoteDate field if non-nil, zero value otherwise.

### GetCreditNoteDateOk

`func (o *ModelCreditNoteResponse) GetCreditNoteDateOk() (*time.Time, bool)`

GetCreditNoteDateOk returns a tuple with the CreditNoteDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNoteDate

`func (o *ModelCreditNoteResponse) SetCreditNoteDate(v time.Time)`

SetCreditNoteDate sets CreditNoteDate field to given value.

### HasCreditNoteDate

`func (o *ModelCreditNoteResponse) HasCreditNoteDate() bool`

HasCreditNoteDate returns a boolean if a field has been set.

### SetCreditNoteDateNil

`func (o *ModelCreditNoteResponse) SetCreditNoteDateNil(b bool)`

 SetCreditNoteDateNil sets the value for CreditNoteDate to be an explicit nil

### UnsetCreditNoteDate
`func (o *ModelCreditNoteResponse) UnsetCreditNoteDate()`

UnsetCreditNoteDate ensures that no value is present for CreditNoteDate, not even an explicit nil
### GetStatus

`func (o *ModelCreditNoteResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelCreditNoteResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelCreditNoteResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelCreditNoteResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetHeader

`func (o *ModelCreditNoteResponse) GetHeader() string`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *ModelCreditNoteResponse) GetHeaderOk() (*string, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *ModelCreditNoteResponse) SetHeader(v string)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *ModelCreditNoteResponse) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### SetHeaderNil

`func (o *ModelCreditNoteResponse) SetHeaderNil(b bool)`

 SetHeaderNil sets the value for Header to be an explicit nil

### UnsetHeader
`func (o *ModelCreditNoteResponse) UnsetHeader()`

UnsetHeader ensures that no value is present for Header, not even an explicit nil
### GetHeadText

`func (o *ModelCreditNoteResponse) GetHeadText() string`

GetHeadText returns the HeadText field if non-nil, zero value otherwise.

### GetHeadTextOk

`func (o *ModelCreditNoteResponse) GetHeadTextOk() (*string, bool)`

GetHeadTextOk returns a tuple with the HeadText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadText

`func (o *ModelCreditNoteResponse) SetHeadText(v string)`

SetHeadText sets HeadText field to given value.

### HasHeadText

`func (o *ModelCreditNoteResponse) HasHeadText() bool`

HasHeadText returns a boolean if a field has been set.

### SetHeadTextNil

`func (o *ModelCreditNoteResponse) SetHeadTextNil(b bool)`

 SetHeadTextNil sets the value for HeadText to be an explicit nil

### UnsetHeadText
`func (o *ModelCreditNoteResponse) UnsetHeadText()`

UnsetHeadText ensures that no value is present for HeadText, not even an explicit nil
### GetFootText

`func (o *ModelCreditNoteResponse) GetFootText() string`

GetFootText returns the FootText field if non-nil, zero value otherwise.

### GetFootTextOk

`func (o *ModelCreditNoteResponse) GetFootTextOk() (*string, bool)`

GetFootTextOk returns a tuple with the FootText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFootText

`func (o *ModelCreditNoteResponse) SetFootText(v string)`

SetFootText sets FootText field to given value.

### HasFootText

`func (o *ModelCreditNoteResponse) HasFootText() bool`

HasFootText returns a boolean if a field has been set.

### SetFootTextNil

`func (o *ModelCreditNoteResponse) SetFootTextNil(b bool)`

 SetFootTextNil sets the value for FootText to be an explicit nil

### UnsetFootText
`func (o *ModelCreditNoteResponse) UnsetFootText()`

UnsetFootText ensures that no value is present for FootText, not even an explicit nil
### GetAddressCountry

`func (o *ModelCreditNoteResponse) GetAddressCountry() ModelCreditNoteResponseAddressCountry`

GetAddressCountry returns the AddressCountry field if non-nil, zero value otherwise.

### GetAddressCountryOk

`func (o *ModelCreditNoteResponse) GetAddressCountryOk() (*ModelCreditNoteResponseAddressCountry, bool)`

GetAddressCountryOk returns a tuple with the AddressCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCountry

`func (o *ModelCreditNoteResponse) SetAddressCountry(v ModelCreditNoteResponseAddressCountry)`

SetAddressCountry sets AddressCountry field to given value.

### HasAddressCountry

`func (o *ModelCreditNoteResponse) HasAddressCountry() bool`

HasAddressCountry returns a boolean if a field has been set.

### SetAddressCountryNil

`func (o *ModelCreditNoteResponse) SetAddressCountryNil(b bool)`

 SetAddressCountryNil sets the value for AddressCountry to be an explicit nil

### UnsetAddressCountry
`func (o *ModelCreditNoteResponse) UnsetAddressCountry()`

UnsetAddressCountry ensures that no value is present for AddressCountry, not even an explicit nil
### GetCreateUser

`func (o *ModelCreditNoteResponse) GetCreateUser() ModelInvoiceResponseCreateUser`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *ModelCreditNoteResponse) GetCreateUserOk() (*ModelInvoiceResponseCreateUser, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *ModelCreditNoteResponse) SetCreateUser(v ModelInvoiceResponseCreateUser)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *ModelCreditNoteResponse) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetSevClient

`func (o *ModelCreditNoteResponse) GetSevClient() ModelCreditNoteResponseSevClient`

GetSevClient returns the SevClient field if non-nil, zero value otherwise.

### GetSevClientOk

`func (o *ModelCreditNoteResponse) GetSevClientOk() (*ModelCreditNoteResponseSevClient, bool)`

GetSevClientOk returns a tuple with the SevClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSevClient

`func (o *ModelCreditNoteResponse) SetSevClient(v ModelCreditNoteResponseSevClient)`

SetSevClient sets SevClient field to given value.

### HasSevClient

`func (o *ModelCreditNoteResponse) HasSevClient() bool`

HasSevClient returns a boolean if a field has been set.

### GetDeliveryTerms

`func (o *ModelCreditNoteResponse) GetDeliveryTerms() string`

GetDeliveryTerms returns the DeliveryTerms field if non-nil, zero value otherwise.

### GetDeliveryTermsOk

`func (o *ModelCreditNoteResponse) GetDeliveryTermsOk() (*string, bool)`

GetDeliveryTermsOk returns a tuple with the DeliveryTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTerms

`func (o *ModelCreditNoteResponse) SetDeliveryTerms(v string)`

SetDeliveryTerms sets DeliveryTerms field to given value.

### HasDeliveryTerms

`func (o *ModelCreditNoteResponse) HasDeliveryTerms() bool`

HasDeliveryTerms returns a boolean if a field has been set.

### SetDeliveryTermsNil

`func (o *ModelCreditNoteResponse) SetDeliveryTermsNil(b bool)`

 SetDeliveryTermsNil sets the value for DeliveryTerms to be an explicit nil

### UnsetDeliveryTerms
`func (o *ModelCreditNoteResponse) UnsetDeliveryTerms()`

UnsetDeliveryTerms ensures that no value is present for DeliveryTerms, not even an explicit nil
### GetDeliveryDate

`func (o *ModelCreditNoteResponse) GetDeliveryDate() time.Time`

GetDeliveryDate returns the DeliveryDate field if non-nil, zero value otherwise.

### GetDeliveryDateOk

`func (o *ModelCreditNoteResponse) GetDeliveryDateOk() (*time.Time, bool)`

GetDeliveryDateOk returns a tuple with the DeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDate

`func (o *ModelCreditNoteResponse) SetDeliveryDate(v time.Time)`

SetDeliveryDate sets DeliveryDate field to given value.

### HasDeliveryDate

`func (o *ModelCreditNoteResponse) HasDeliveryDate() bool`

HasDeliveryDate returns a boolean if a field has been set.

### GetPaymentTerms

`func (o *ModelCreditNoteResponse) GetPaymentTerms() string`

GetPaymentTerms returns the PaymentTerms field if non-nil, zero value otherwise.

### GetPaymentTermsOk

`func (o *ModelCreditNoteResponse) GetPaymentTermsOk() (*string, bool)`

GetPaymentTermsOk returns a tuple with the PaymentTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTerms

`func (o *ModelCreditNoteResponse) SetPaymentTerms(v string)`

SetPaymentTerms sets PaymentTerms field to given value.

### HasPaymentTerms

`func (o *ModelCreditNoteResponse) HasPaymentTerms() bool`

HasPaymentTerms returns a boolean if a field has been set.

### SetPaymentTermsNil

`func (o *ModelCreditNoteResponse) SetPaymentTermsNil(b bool)`

 SetPaymentTermsNil sets the value for PaymentTerms to be an explicit nil

### UnsetPaymentTerms
`func (o *ModelCreditNoteResponse) UnsetPaymentTerms()`

UnsetPaymentTerms ensures that no value is present for PaymentTerms, not even an explicit nil
### GetVersion

`func (o *ModelCreditNoteResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ModelCreditNoteResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ModelCreditNoteResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ModelCreditNoteResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ModelCreditNoteResponse) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ModelCreditNoteResponse) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetSmallSettlement

`func (o *ModelCreditNoteResponse) GetSmallSettlement() bool`

GetSmallSettlement returns the SmallSettlement field if non-nil, zero value otherwise.

### GetSmallSettlementOk

`func (o *ModelCreditNoteResponse) GetSmallSettlementOk() (*bool, bool)`

GetSmallSettlementOk returns a tuple with the SmallSettlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallSettlement

`func (o *ModelCreditNoteResponse) SetSmallSettlement(v bool)`

SetSmallSettlement sets SmallSettlement field to given value.

### HasSmallSettlement

`func (o *ModelCreditNoteResponse) HasSmallSettlement() bool`

HasSmallSettlement returns a boolean if a field has been set.

### SetSmallSettlementNil

`func (o *ModelCreditNoteResponse) SetSmallSettlementNil(b bool)`

 SetSmallSettlementNil sets the value for SmallSettlement to be an explicit nil

### UnsetSmallSettlement
`func (o *ModelCreditNoteResponse) UnsetSmallSettlement()`

UnsetSmallSettlement ensures that no value is present for SmallSettlement, not even an explicit nil
### GetContactPerson

`func (o *ModelCreditNoteResponse) GetContactPerson() ModelCreditNoteResponseContactPerson`

GetContactPerson returns the ContactPerson field if non-nil, zero value otherwise.

### GetContactPersonOk

`func (o *ModelCreditNoteResponse) GetContactPersonOk() (*ModelCreditNoteResponseContactPerson, bool)`

GetContactPersonOk returns a tuple with the ContactPerson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPerson

`func (o *ModelCreditNoteResponse) SetContactPerson(v ModelCreditNoteResponseContactPerson)`

SetContactPerson sets ContactPerson field to given value.

### HasContactPerson

`func (o *ModelCreditNoteResponse) HasContactPerson() bool`

HasContactPerson returns a boolean if a field has been set.

### SetContactPersonNil

`func (o *ModelCreditNoteResponse) SetContactPersonNil(b bool)`

 SetContactPersonNil sets the value for ContactPerson to be an explicit nil

### UnsetContactPerson
`func (o *ModelCreditNoteResponse) UnsetContactPerson()`

UnsetContactPerson ensures that no value is present for ContactPerson, not even an explicit nil
### GetTaxRate

`func (o *ModelCreditNoteResponse) GetTaxRate() string`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *ModelCreditNoteResponse) GetTaxRateOk() (*string, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *ModelCreditNoteResponse) SetTaxRate(v string)`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *ModelCreditNoteResponse) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### SetTaxRateNil

`func (o *ModelCreditNoteResponse) SetTaxRateNil(b bool)`

 SetTaxRateNil sets the value for TaxRate to be an explicit nil

### UnsetTaxRate
`func (o *ModelCreditNoteResponse) UnsetTaxRate()`

UnsetTaxRate ensures that no value is present for TaxRate, not even an explicit nil
### GetTaxSet

`func (o *ModelCreditNoteResponse) GetTaxSet() ModelCreditNoteResponseTaxSet`

GetTaxSet returns the TaxSet field if non-nil, zero value otherwise.

### GetTaxSetOk

`func (o *ModelCreditNoteResponse) GetTaxSetOk() (*ModelCreditNoteResponseTaxSet, bool)`

GetTaxSetOk returns a tuple with the TaxSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxSet

`func (o *ModelCreditNoteResponse) SetTaxSet(v ModelCreditNoteResponseTaxSet)`

SetTaxSet sets TaxSet field to given value.

### HasTaxSet

`func (o *ModelCreditNoteResponse) HasTaxSet() bool`

HasTaxSet returns a boolean if a field has been set.

### SetTaxSetNil

`func (o *ModelCreditNoteResponse) SetTaxSetNil(b bool)`

 SetTaxSetNil sets the value for TaxSet to be an explicit nil

### UnsetTaxSet
`func (o *ModelCreditNoteResponse) UnsetTaxSet()`

UnsetTaxSet ensures that no value is present for TaxSet, not even an explicit nil
### GetTaxText

`func (o *ModelCreditNoteResponse) GetTaxText() string`

GetTaxText returns the TaxText field if non-nil, zero value otherwise.

### GetTaxTextOk

`func (o *ModelCreditNoteResponse) GetTaxTextOk() (*string, bool)`

GetTaxTextOk returns a tuple with the TaxText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxText

`func (o *ModelCreditNoteResponse) SetTaxText(v string)`

SetTaxText sets TaxText field to given value.

### HasTaxText

`func (o *ModelCreditNoteResponse) HasTaxText() bool`

HasTaxText returns a boolean if a field has been set.

### SetTaxTextNil

`func (o *ModelCreditNoteResponse) SetTaxTextNil(b bool)`

 SetTaxTextNil sets the value for TaxText to be an explicit nil

### UnsetTaxText
`func (o *ModelCreditNoteResponse) UnsetTaxText()`

UnsetTaxText ensures that no value is present for TaxText, not even an explicit nil
### GetTaxType

`func (o *ModelCreditNoteResponse) GetTaxType() string`

GetTaxType returns the TaxType field if non-nil, zero value otherwise.

### GetTaxTypeOk

`func (o *ModelCreditNoteResponse) GetTaxTypeOk() (*string, bool)`

GetTaxTypeOk returns a tuple with the TaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxType

`func (o *ModelCreditNoteResponse) SetTaxType(v string)`

SetTaxType sets TaxType field to given value.

### HasTaxType

`func (o *ModelCreditNoteResponse) HasTaxType() bool`

HasTaxType returns a boolean if a field has been set.

### SetTaxTypeNil

`func (o *ModelCreditNoteResponse) SetTaxTypeNil(b bool)`

 SetTaxTypeNil sets the value for TaxType to be an explicit nil

### UnsetTaxType
`func (o *ModelCreditNoteResponse) UnsetTaxType()`

UnsetTaxType ensures that no value is present for TaxType, not even an explicit nil
### GetCreditNoteType

`func (o *ModelCreditNoteResponse) GetCreditNoteType() string`

GetCreditNoteType returns the CreditNoteType field if non-nil, zero value otherwise.

### GetCreditNoteTypeOk

`func (o *ModelCreditNoteResponse) GetCreditNoteTypeOk() (*string, bool)`

GetCreditNoteTypeOk returns a tuple with the CreditNoteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNoteType

`func (o *ModelCreditNoteResponse) SetCreditNoteType(v string)`

SetCreditNoteType sets CreditNoteType field to given value.

### HasCreditNoteType

`func (o *ModelCreditNoteResponse) HasCreditNoteType() bool`

HasCreditNoteType returns a boolean if a field has been set.

### SetCreditNoteTypeNil

`func (o *ModelCreditNoteResponse) SetCreditNoteTypeNil(b bool)`

 SetCreditNoteTypeNil sets the value for CreditNoteType to be an explicit nil

### UnsetCreditNoteType
`func (o *ModelCreditNoteResponse) UnsetCreditNoteType()`

UnsetCreditNoteType ensures that no value is present for CreditNoteType, not even an explicit nil
### GetSendDate

`func (o *ModelCreditNoteResponse) GetSendDate() time.Time`

GetSendDate returns the SendDate field if non-nil, zero value otherwise.

### GetSendDateOk

`func (o *ModelCreditNoteResponse) GetSendDateOk() (*time.Time, bool)`

GetSendDateOk returns a tuple with the SendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendDate

`func (o *ModelCreditNoteResponse) SetSendDate(v time.Time)`

SetSendDate sets SendDate field to given value.

### HasSendDate

`func (o *ModelCreditNoteResponse) HasSendDate() bool`

HasSendDate returns a boolean if a field has been set.

### SetSendDateNil

`func (o *ModelCreditNoteResponse) SetSendDateNil(b bool)`

 SetSendDateNil sets the value for SendDate to be an explicit nil

### UnsetSendDate
`func (o *ModelCreditNoteResponse) UnsetSendDate()`

UnsetSendDate ensures that no value is present for SendDate, not even an explicit nil
### GetAddress

`func (o *ModelCreditNoteResponse) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ModelCreditNoteResponse) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ModelCreditNoteResponse) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ModelCreditNoteResponse) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *ModelCreditNoteResponse) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *ModelCreditNoteResponse) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetCurrency

`func (o *ModelCreditNoteResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ModelCreditNoteResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ModelCreditNoteResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ModelCreditNoteResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *ModelCreditNoteResponse) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *ModelCreditNoteResponse) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetSumNet

`func (o *ModelCreditNoteResponse) GetSumNet() string`

GetSumNet returns the SumNet field if non-nil, zero value otherwise.

### GetSumNetOk

`func (o *ModelCreditNoteResponse) GetSumNetOk() (*string, bool)`

GetSumNetOk returns a tuple with the SumNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumNet

`func (o *ModelCreditNoteResponse) SetSumNet(v string)`

SetSumNet sets SumNet field to given value.

### HasSumNet

`func (o *ModelCreditNoteResponse) HasSumNet() bool`

HasSumNet returns a boolean if a field has been set.

### GetSumTax

`func (o *ModelCreditNoteResponse) GetSumTax() string`

GetSumTax returns the SumTax field if non-nil, zero value otherwise.

### GetSumTaxOk

`func (o *ModelCreditNoteResponse) GetSumTaxOk() (*string, bool)`

GetSumTaxOk returns a tuple with the SumTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumTax

`func (o *ModelCreditNoteResponse) SetSumTax(v string)`

SetSumTax sets SumTax field to given value.

### HasSumTax

`func (o *ModelCreditNoteResponse) HasSumTax() bool`

HasSumTax returns a boolean if a field has been set.

### GetSumGross

`func (o *ModelCreditNoteResponse) GetSumGross() string`

GetSumGross returns the SumGross field if non-nil, zero value otherwise.

### GetSumGrossOk

`func (o *ModelCreditNoteResponse) GetSumGrossOk() (*string, bool)`

GetSumGrossOk returns a tuple with the SumGross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumGross

`func (o *ModelCreditNoteResponse) SetSumGross(v string)`

SetSumGross sets SumGross field to given value.

### HasSumGross

`func (o *ModelCreditNoteResponse) HasSumGross() bool`

HasSumGross returns a boolean if a field has been set.

### GetSumDiscounts

`func (o *ModelCreditNoteResponse) GetSumDiscounts() string`

GetSumDiscounts returns the SumDiscounts field if non-nil, zero value otherwise.

### GetSumDiscountsOk

`func (o *ModelCreditNoteResponse) GetSumDiscountsOk() (*string, bool)`

GetSumDiscountsOk returns a tuple with the SumDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumDiscounts

`func (o *ModelCreditNoteResponse) SetSumDiscounts(v string)`

SetSumDiscounts sets SumDiscounts field to given value.

### HasSumDiscounts

`func (o *ModelCreditNoteResponse) HasSumDiscounts() bool`

HasSumDiscounts returns a boolean if a field has been set.

### GetSumNetForeignCurrency

`func (o *ModelCreditNoteResponse) GetSumNetForeignCurrency() string`

GetSumNetForeignCurrency returns the SumNetForeignCurrency field if non-nil, zero value otherwise.

### GetSumNetForeignCurrencyOk

`func (o *ModelCreditNoteResponse) GetSumNetForeignCurrencyOk() (*string, bool)`

GetSumNetForeignCurrencyOk returns a tuple with the SumNetForeignCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumNetForeignCurrency

`func (o *ModelCreditNoteResponse) SetSumNetForeignCurrency(v string)`

SetSumNetForeignCurrency sets SumNetForeignCurrency field to given value.

### HasSumNetForeignCurrency

`func (o *ModelCreditNoteResponse) HasSumNetForeignCurrency() bool`

HasSumNetForeignCurrency returns a boolean if a field has been set.

### GetSumTaxForeignCurrency

`func (o *ModelCreditNoteResponse) GetSumTaxForeignCurrency() string`

GetSumTaxForeignCurrency returns the SumTaxForeignCurrency field if non-nil, zero value otherwise.

### GetSumTaxForeignCurrencyOk

`func (o *ModelCreditNoteResponse) GetSumTaxForeignCurrencyOk() (*string, bool)`

GetSumTaxForeignCurrencyOk returns a tuple with the SumTaxForeignCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumTaxForeignCurrency

`func (o *ModelCreditNoteResponse) SetSumTaxForeignCurrency(v string)`

SetSumTaxForeignCurrency sets SumTaxForeignCurrency field to given value.

### HasSumTaxForeignCurrency

`func (o *ModelCreditNoteResponse) HasSumTaxForeignCurrency() bool`

HasSumTaxForeignCurrency returns a boolean if a field has been set.

### GetSumGrossForeignCurrency

`func (o *ModelCreditNoteResponse) GetSumGrossForeignCurrency() string`

GetSumGrossForeignCurrency returns the SumGrossForeignCurrency field if non-nil, zero value otherwise.

### GetSumGrossForeignCurrencyOk

`func (o *ModelCreditNoteResponse) GetSumGrossForeignCurrencyOk() (*string, bool)`

GetSumGrossForeignCurrencyOk returns a tuple with the SumGrossForeignCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumGrossForeignCurrency

`func (o *ModelCreditNoteResponse) SetSumGrossForeignCurrency(v string)`

SetSumGrossForeignCurrency sets SumGrossForeignCurrency field to given value.

### HasSumGrossForeignCurrency

`func (o *ModelCreditNoteResponse) HasSumGrossForeignCurrency() bool`

HasSumGrossForeignCurrency returns a boolean if a field has been set.

### GetSumDiscountsForeignCurrency

`func (o *ModelCreditNoteResponse) GetSumDiscountsForeignCurrency() string`

GetSumDiscountsForeignCurrency returns the SumDiscountsForeignCurrency field if non-nil, zero value otherwise.

### GetSumDiscountsForeignCurrencyOk

`func (o *ModelCreditNoteResponse) GetSumDiscountsForeignCurrencyOk() (*string, bool)`

GetSumDiscountsForeignCurrencyOk returns a tuple with the SumDiscountsForeignCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumDiscountsForeignCurrency

`func (o *ModelCreditNoteResponse) SetSumDiscountsForeignCurrency(v string)`

SetSumDiscountsForeignCurrency sets SumDiscountsForeignCurrency field to given value.

### HasSumDiscountsForeignCurrency

`func (o *ModelCreditNoteResponse) HasSumDiscountsForeignCurrency() bool`

HasSumDiscountsForeignCurrency returns a boolean if a field has been set.

### GetCustomerInternalNote

`func (o *ModelCreditNoteResponse) GetCustomerInternalNote() string`

GetCustomerInternalNote returns the CustomerInternalNote field if non-nil, zero value otherwise.

### GetCustomerInternalNoteOk

`func (o *ModelCreditNoteResponse) GetCustomerInternalNoteOk() (*string, bool)`

GetCustomerInternalNoteOk returns a tuple with the CustomerInternalNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerInternalNote

`func (o *ModelCreditNoteResponse) SetCustomerInternalNote(v string)`

SetCustomerInternalNote sets CustomerInternalNote field to given value.

### HasCustomerInternalNote

`func (o *ModelCreditNoteResponse) HasCustomerInternalNote() bool`

HasCustomerInternalNote returns a boolean if a field has been set.

### SetCustomerInternalNoteNil

`func (o *ModelCreditNoteResponse) SetCustomerInternalNoteNil(b bool)`

 SetCustomerInternalNoteNil sets the value for CustomerInternalNote to be an explicit nil

### UnsetCustomerInternalNote
`func (o *ModelCreditNoteResponse) UnsetCustomerInternalNote()`

UnsetCustomerInternalNote ensures that no value is present for CustomerInternalNote, not even an explicit nil
### GetShowNet

`func (o *ModelCreditNoteResponse) GetShowNet() bool`

GetShowNet returns the ShowNet field if non-nil, zero value otherwise.

### GetShowNetOk

`func (o *ModelCreditNoteResponse) GetShowNetOk() (*bool, bool)`

GetShowNetOk returns a tuple with the ShowNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowNet

`func (o *ModelCreditNoteResponse) SetShowNet(v bool)`

SetShowNet sets ShowNet field to given value.

### HasShowNet

`func (o *ModelCreditNoteResponse) HasShowNet() bool`

HasShowNet returns a boolean if a field has been set.

### GetSendType

`func (o *ModelCreditNoteResponse) GetSendType() string`

GetSendType returns the SendType field if non-nil, zero value otherwise.

### GetSendTypeOk

`func (o *ModelCreditNoteResponse) GetSendTypeOk() (*string, bool)`

GetSendTypeOk returns a tuple with the SendType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendType

`func (o *ModelCreditNoteResponse) SetSendType(v string)`

SetSendType sets SendType field to given value.

### HasSendType

`func (o *ModelCreditNoteResponse) HasSendType() bool`

HasSendType returns a boolean if a field has been set.

### SetSendTypeNil

`func (o *ModelCreditNoteResponse) SetSendTypeNil(b bool)`

 SetSendTypeNil sets the value for SendType to be an explicit nil

### UnsetSendType
`func (o *ModelCreditNoteResponse) UnsetSendType()`

UnsetSendType ensures that no value is present for SendType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


