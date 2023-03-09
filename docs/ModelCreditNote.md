# ModelCreditNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The creditNote id. &lt;span style&#x3D;&#39;color:red&#39;&gt;Required&lt;/span&gt; if you want to create/update an credit note position for an existing credit note\&quot; | [optional] 
**ObjectName** | **string** | The creditNote object name | 
**MapAll** | **bool** |  | 
**Create** | Pointer to **time.Time** | Date of creditNote creation | [optional] [readonly] 
**Update** | Pointer to **time.Time** | Date of last creditNote update | [optional] [readonly] 
**CreditNoteNumber** | **string** | The creditNote number | 
**Contact** | [**ModelCreditNoteContact**](ModelCreditNoteContact.md) |  | 
**CreditNoteDate** | **time.Time** | Needs to be provided as timestamp or dd.mm.yyyy | 
**Status** | **string** | Please have a look in       &lt;a href&#x3D;&#39;https://api.sevdesk.de/#section/Types-and-status-of-credit-notes&#39;&gt;status of credit note&lt;/a&gt;      to see what the different status codes mean | 
**Header** | **string** | Normally consist of prefix plus the creditNote number | 
**HeadText** | Pointer to **NullableString** | Certain html tags can be used here to format your text | [optional] 
**FootText** | Pointer to **NullableString** | Certain html tags can be used here to format your text | [optional] 
**AddressCountry** | [**NullableModelCreditNoteAddressCountry**](ModelCreditNoteAddressCountry.md) |  | 
**CreateUser** | Pointer to [**ModelCreditNoteCreateUser**](ModelCreditNoteCreateUser.md) |  | [optional] 
**SevClient** | Pointer to [**ModelCreditNoteSevClient**](ModelCreditNoteSevClient.md) |  | [optional] 
**DeliveryTerms** | Pointer to **NullableString** | Delivery terms of the creditNote | [optional] 
**PaymentTerms** | Pointer to **NullableString** | Payment terms of the creditNote | [optional] 
**Version** | Pointer to **NullableInt32** | Version of the creditNote.&lt;br&gt;      Can be used if you have multiple drafts for the same creditNote.&lt;br&gt;      Should start with 0 | [optional] 
**SmallSettlement** | Pointer to **NullableBool** | Defines if the client uses the small settlement scheme.      If yes, the creditNote must not contain any vat | [optional] 
**ContactPerson** | [**ModelCreditNoteContactPerson**](ModelCreditNoteContactPerson.md) |  | 
**TaxRate** | **float32** | Is overwritten by creditNote position tax rates | 
**TaxSet** | Pointer to [**NullableModelCreditNoteTaxSet**](ModelCreditNoteTaxSet.md) |  | [optional] 
**TaxText** | **string** | A common tax text would be &#39;Umsatzsteuer 19%&#39; | 
**TaxType** | **string** | Tax type of the creditNote. There are four tax types: 1. default - Umsatzsteuer ausweisen 2. eu - Steuerfreie innergemeinschaftliche Lieferung (Europäische Union) 3. noteu - Steuerschuldnerschaft des Leistungsempfängers (außerhalb EU, z. B. Schweiz) 4. custom - Using custom tax set 5. ss - Not subject to VAT according to §19 1 UStG Tax rates are heavily connected to the tax type used. | 
**CreditNoteType** | Pointer to **string** | Type of the creditNote. | [optional] 
**SendDate** | Pointer to **NullableTime** | The date the creditNote was sent to the customer | [optional] 
**Address** | Pointer to **NullableString** | Complete address of the recipient including name, street, city, zip and country.&lt;br&gt;       Line breaks can be used and will be displayed on the invoice pdf. | [optional] 
**BookingCategory** | Pointer to **NullableString** | defines the booking category, for more information see the section \&quot;&lt;a href&#x3D;&#39;https://api.sevdesk.de/#section/Credit-note-booking-categories&#39;&gt;Credit note booking categories&lt;/a&gt;\&quot; | [optional] 
**Currency** | **string** | Currency used in the creditNote. Needs to be currency code according to ISO-4217 | 
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

### NewModelCreditNote

`func NewModelCreditNote(objectName string, mapAll bool, creditNoteNumber string, contact ModelCreditNoteContact, creditNoteDate time.Time, status string, header string, addressCountry NullableModelCreditNoteAddressCountry, contactPerson ModelCreditNoteContactPerson, taxRate float32, taxText string, taxType string, currency string, ) *ModelCreditNote`

NewModelCreditNote instantiates a new ModelCreditNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelCreditNoteWithDefaults

`func NewModelCreditNoteWithDefaults() *ModelCreditNote`

NewModelCreditNoteWithDefaults instantiates a new ModelCreditNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelCreditNote) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelCreditNote) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelCreditNote) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelCreditNote) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectName

`func (o *ModelCreditNote) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ModelCreditNote) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ModelCreditNote) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.


### GetMapAll

`func (o *ModelCreditNote) GetMapAll() bool`

GetMapAll returns the MapAll field if non-nil, zero value otherwise.

### GetMapAllOk

`func (o *ModelCreditNote) GetMapAllOk() (*bool, bool)`

GetMapAllOk returns a tuple with the MapAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapAll

`func (o *ModelCreditNote) SetMapAll(v bool)`

SetMapAll sets MapAll field to given value.


### GetCreate

`func (o *ModelCreditNote) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ModelCreditNote) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ModelCreditNote) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ModelCreditNote) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUpdate

`func (o *ModelCreditNote) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ModelCreditNote) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ModelCreditNote) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ModelCreditNote) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetCreditNoteNumber

`func (o *ModelCreditNote) GetCreditNoteNumber() string`

GetCreditNoteNumber returns the CreditNoteNumber field if non-nil, zero value otherwise.

### GetCreditNoteNumberOk

`func (o *ModelCreditNote) GetCreditNoteNumberOk() (*string, bool)`

GetCreditNoteNumberOk returns a tuple with the CreditNoteNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNoteNumber

`func (o *ModelCreditNote) SetCreditNoteNumber(v string)`

SetCreditNoteNumber sets CreditNoteNumber field to given value.


### GetContact

`func (o *ModelCreditNote) GetContact() ModelCreditNoteContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ModelCreditNote) GetContactOk() (*ModelCreditNoteContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ModelCreditNote) SetContact(v ModelCreditNoteContact)`

SetContact sets Contact field to given value.


### GetCreditNoteDate

`func (o *ModelCreditNote) GetCreditNoteDate() time.Time`

GetCreditNoteDate returns the CreditNoteDate field if non-nil, zero value otherwise.

### GetCreditNoteDateOk

`func (o *ModelCreditNote) GetCreditNoteDateOk() (*time.Time, bool)`

GetCreditNoteDateOk returns a tuple with the CreditNoteDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNoteDate

`func (o *ModelCreditNote) SetCreditNoteDate(v time.Time)`

SetCreditNoteDate sets CreditNoteDate field to given value.


### GetStatus

`func (o *ModelCreditNote) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelCreditNote) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelCreditNote) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetHeader

`func (o *ModelCreditNote) GetHeader() string`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *ModelCreditNote) GetHeaderOk() (*string, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *ModelCreditNote) SetHeader(v string)`

SetHeader sets Header field to given value.


### GetHeadText

`func (o *ModelCreditNote) GetHeadText() string`

GetHeadText returns the HeadText field if non-nil, zero value otherwise.

### GetHeadTextOk

`func (o *ModelCreditNote) GetHeadTextOk() (*string, bool)`

GetHeadTextOk returns a tuple with the HeadText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadText

`func (o *ModelCreditNote) SetHeadText(v string)`

SetHeadText sets HeadText field to given value.

### HasHeadText

`func (o *ModelCreditNote) HasHeadText() bool`

HasHeadText returns a boolean if a field has been set.

### SetHeadTextNil

`func (o *ModelCreditNote) SetHeadTextNil(b bool)`

 SetHeadTextNil sets the value for HeadText to be an explicit nil

### UnsetHeadText
`func (o *ModelCreditNote) UnsetHeadText()`

UnsetHeadText ensures that no value is present for HeadText, not even an explicit nil
### GetFootText

`func (o *ModelCreditNote) GetFootText() string`

GetFootText returns the FootText field if non-nil, zero value otherwise.

### GetFootTextOk

`func (o *ModelCreditNote) GetFootTextOk() (*string, bool)`

GetFootTextOk returns a tuple with the FootText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFootText

`func (o *ModelCreditNote) SetFootText(v string)`

SetFootText sets FootText field to given value.

### HasFootText

`func (o *ModelCreditNote) HasFootText() bool`

HasFootText returns a boolean if a field has been set.

### SetFootTextNil

`func (o *ModelCreditNote) SetFootTextNil(b bool)`

 SetFootTextNil sets the value for FootText to be an explicit nil

### UnsetFootText
`func (o *ModelCreditNote) UnsetFootText()`

UnsetFootText ensures that no value is present for FootText, not even an explicit nil
### GetAddressCountry

`func (o *ModelCreditNote) GetAddressCountry() ModelCreditNoteAddressCountry`

GetAddressCountry returns the AddressCountry field if non-nil, zero value otherwise.

### GetAddressCountryOk

`func (o *ModelCreditNote) GetAddressCountryOk() (*ModelCreditNoteAddressCountry, bool)`

GetAddressCountryOk returns a tuple with the AddressCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCountry

`func (o *ModelCreditNote) SetAddressCountry(v ModelCreditNoteAddressCountry)`

SetAddressCountry sets AddressCountry field to given value.


### SetAddressCountryNil

`func (o *ModelCreditNote) SetAddressCountryNil(b bool)`

 SetAddressCountryNil sets the value for AddressCountry to be an explicit nil

### UnsetAddressCountry
`func (o *ModelCreditNote) UnsetAddressCountry()`

UnsetAddressCountry ensures that no value is present for AddressCountry, not even an explicit nil
### GetCreateUser

`func (o *ModelCreditNote) GetCreateUser() ModelCreditNoteCreateUser`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *ModelCreditNote) GetCreateUserOk() (*ModelCreditNoteCreateUser, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *ModelCreditNote) SetCreateUser(v ModelCreditNoteCreateUser)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *ModelCreditNote) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetSevClient

`func (o *ModelCreditNote) GetSevClient() ModelCreditNoteSevClient`

GetSevClient returns the SevClient field if non-nil, zero value otherwise.

### GetSevClientOk

`func (o *ModelCreditNote) GetSevClientOk() (*ModelCreditNoteSevClient, bool)`

GetSevClientOk returns a tuple with the SevClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSevClient

`func (o *ModelCreditNote) SetSevClient(v ModelCreditNoteSevClient)`

SetSevClient sets SevClient field to given value.

### HasSevClient

`func (o *ModelCreditNote) HasSevClient() bool`

HasSevClient returns a boolean if a field has been set.

### GetDeliveryTerms

`func (o *ModelCreditNote) GetDeliveryTerms() string`

GetDeliveryTerms returns the DeliveryTerms field if non-nil, zero value otherwise.

### GetDeliveryTermsOk

`func (o *ModelCreditNote) GetDeliveryTermsOk() (*string, bool)`

GetDeliveryTermsOk returns a tuple with the DeliveryTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTerms

`func (o *ModelCreditNote) SetDeliveryTerms(v string)`

SetDeliveryTerms sets DeliveryTerms field to given value.

### HasDeliveryTerms

`func (o *ModelCreditNote) HasDeliveryTerms() bool`

HasDeliveryTerms returns a boolean if a field has been set.

### SetDeliveryTermsNil

`func (o *ModelCreditNote) SetDeliveryTermsNil(b bool)`

 SetDeliveryTermsNil sets the value for DeliveryTerms to be an explicit nil

### UnsetDeliveryTerms
`func (o *ModelCreditNote) UnsetDeliveryTerms()`

UnsetDeliveryTerms ensures that no value is present for DeliveryTerms, not even an explicit nil
### GetPaymentTerms

`func (o *ModelCreditNote) GetPaymentTerms() string`

GetPaymentTerms returns the PaymentTerms field if non-nil, zero value otherwise.

### GetPaymentTermsOk

`func (o *ModelCreditNote) GetPaymentTermsOk() (*string, bool)`

GetPaymentTermsOk returns a tuple with the PaymentTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTerms

`func (o *ModelCreditNote) SetPaymentTerms(v string)`

SetPaymentTerms sets PaymentTerms field to given value.

### HasPaymentTerms

`func (o *ModelCreditNote) HasPaymentTerms() bool`

HasPaymentTerms returns a boolean if a field has been set.

### SetPaymentTermsNil

`func (o *ModelCreditNote) SetPaymentTermsNil(b bool)`

 SetPaymentTermsNil sets the value for PaymentTerms to be an explicit nil

### UnsetPaymentTerms
`func (o *ModelCreditNote) UnsetPaymentTerms()`

UnsetPaymentTerms ensures that no value is present for PaymentTerms, not even an explicit nil
### GetVersion

`func (o *ModelCreditNote) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ModelCreditNote) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ModelCreditNote) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ModelCreditNote) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ModelCreditNote) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ModelCreditNote) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetSmallSettlement

`func (o *ModelCreditNote) GetSmallSettlement() bool`

GetSmallSettlement returns the SmallSettlement field if non-nil, zero value otherwise.

### GetSmallSettlementOk

`func (o *ModelCreditNote) GetSmallSettlementOk() (*bool, bool)`

GetSmallSettlementOk returns a tuple with the SmallSettlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallSettlement

`func (o *ModelCreditNote) SetSmallSettlement(v bool)`

SetSmallSettlement sets SmallSettlement field to given value.

### HasSmallSettlement

`func (o *ModelCreditNote) HasSmallSettlement() bool`

HasSmallSettlement returns a boolean if a field has been set.

### SetSmallSettlementNil

`func (o *ModelCreditNote) SetSmallSettlementNil(b bool)`

 SetSmallSettlementNil sets the value for SmallSettlement to be an explicit nil

### UnsetSmallSettlement
`func (o *ModelCreditNote) UnsetSmallSettlement()`

UnsetSmallSettlement ensures that no value is present for SmallSettlement, not even an explicit nil
### GetContactPerson

`func (o *ModelCreditNote) GetContactPerson() ModelCreditNoteContactPerson`

GetContactPerson returns the ContactPerson field if non-nil, zero value otherwise.

### GetContactPersonOk

`func (o *ModelCreditNote) GetContactPersonOk() (*ModelCreditNoteContactPerson, bool)`

GetContactPersonOk returns a tuple with the ContactPerson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPerson

`func (o *ModelCreditNote) SetContactPerson(v ModelCreditNoteContactPerson)`

SetContactPerson sets ContactPerson field to given value.


### GetTaxRate

`func (o *ModelCreditNote) GetTaxRate() float32`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *ModelCreditNote) GetTaxRateOk() (*float32, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *ModelCreditNote) SetTaxRate(v float32)`

SetTaxRate sets TaxRate field to given value.


### GetTaxSet

`func (o *ModelCreditNote) GetTaxSet() ModelCreditNoteTaxSet`

GetTaxSet returns the TaxSet field if non-nil, zero value otherwise.

### GetTaxSetOk

`func (o *ModelCreditNote) GetTaxSetOk() (*ModelCreditNoteTaxSet, bool)`

GetTaxSetOk returns a tuple with the TaxSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxSet

`func (o *ModelCreditNote) SetTaxSet(v ModelCreditNoteTaxSet)`

SetTaxSet sets TaxSet field to given value.

### HasTaxSet

`func (o *ModelCreditNote) HasTaxSet() bool`

HasTaxSet returns a boolean if a field has been set.

### SetTaxSetNil

`func (o *ModelCreditNote) SetTaxSetNil(b bool)`

 SetTaxSetNil sets the value for TaxSet to be an explicit nil

### UnsetTaxSet
`func (o *ModelCreditNote) UnsetTaxSet()`

UnsetTaxSet ensures that no value is present for TaxSet, not even an explicit nil
### GetTaxText

`func (o *ModelCreditNote) GetTaxText() string`

GetTaxText returns the TaxText field if non-nil, zero value otherwise.

### GetTaxTextOk

`func (o *ModelCreditNote) GetTaxTextOk() (*string, bool)`

GetTaxTextOk returns a tuple with the TaxText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxText

`func (o *ModelCreditNote) SetTaxText(v string)`

SetTaxText sets TaxText field to given value.


### GetTaxType

`func (o *ModelCreditNote) GetTaxType() string`

GetTaxType returns the TaxType field if non-nil, zero value otherwise.

### GetTaxTypeOk

`func (o *ModelCreditNote) GetTaxTypeOk() (*string, bool)`

GetTaxTypeOk returns a tuple with the TaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxType

`func (o *ModelCreditNote) SetTaxType(v string)`

SetTaxType sets TaxType field to given value.


### GetCreditNoteType

`func (o *ModelCreditNote) GetCreditNoteType() string`

GetCreditNoteType returns the CreditNoteType field if non-nil, zero value otherwise.

### GetCreditNoteTypeOk

`func (o *ModelCreditNote) GetCreditNoteTypeOk() (*string, bool)`

GetCreditNoteTypeOk returns a tuple with the CreditNoteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNoteType

`func (o *ModelCreditNote) SetCreditNoteType(v string)`

SetCreditNoteType sets CreditNoteType field to given value.

### HasCreditNoteType

`func (o *ModelCreditNote) HasCreditNoteType() bool`

HasCreditNoteType returns a boolean if a field has been set.

### GetSendDate

`func (o *ModelCreditNote) GetSendDate() time.Time`

GetSendDate returns the SendDate field if non-nil, zero value otherwise.

### GetSendDateOk

`func (o *ModelCreditNote) GetSendDateOk() (*time.Time, bool)`

GetSendDateOk returns a tuple with the SendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendDate

`func (o *ModelCreditNote) SetSendDate(v time.Time)`

SetSendDate sets SendDate field to given value.

### HasSendDate

`func (o *ModelCreditNote) HasSendDate() bool`

HasSendDate returns a boolean if a field has been set.

### SetSendDateNil

`func (o *ModelCreditNote) SetSendDateNil(b bool)`

 SetSendDateNil sets the value for SendDate to be an explicit nil

### UnsetSendDate
`func (o *ModelCreditNote) UnsetSendDate()`

UnsetSendDate ensures that no value is present for SendDate, not even an explicit nil
### GetAddress

`func (o *ModelCreditNote) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ModelCreditNote) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ModelCreditNote) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ModelCreditNote) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *ModelCreditNote) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *ModelCreditNote) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetBookingCategory

`func (o *ModelCreditNote) GetBookingCategory() string`

GetBookingCategory returns the BookingCategory field if non-nil, zero value otherwise.

### GetBookingCategoryOk

`func (o *ModelCreditNote) GetBookingCategoryOk() (*string, bool)`

GetBookingCategoryOk returns a tuple with the BookingCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookingCategory

`func (o *ModelCreditNote) SetBookingCategory(v string)`

SetBookingCategory sets BookingCategory field to given value.

### HasBookingCategory

`func (o *ModelCreditNote) HasBookingCategory() bool`

HasBookingCategory returns a boolean if a field has been set.

### SetBookingCategoryNil

`func (o *ModelCreditNote) SetBookingCategoryNil(b bool)`

 SetBookingCategoryNil sets the value for BookingCategory to be an explicit nil

### UnsetBookingCategory
`func (o *ModelCreditNote) UnsetBookingCategory()`

UnsetBookingCategory ensures that no value is present for BookingCategory, not even an explicit nil
### GetCurrency

`func (o *ModelCreditNote) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ModelCreditNote) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ModelCreditNote) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetSumNet

`func (o *ModelCreditNote) GetSumNet() float32`

GetSumNet returns the SumNet field if non-nil, zero value otherwise.

### GetSumNetOk

`func (o *ModelCreditNote) GetSumNetOk() (*float32, bool)`

GetSumNetOk returns a tuple with the SumNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumNet

`func (o *ModelCreditNote) SetSumNet(v float32)`

SetSumNet sets SumNet field to given value.

### HasSumNet

`func (o *ModelCreditNote) HasSumNet() bool`

HasSumNet returns a boolean if a field has been set.

### GetSumTax

`func (o *ModelCreditNote) GetSumTax() float32`

GetSumTax returns the SumTax field if non-nil, zero value otherwise.

### GetSumTaxOk

`func (o *ModelCreditNote) GetSumTaxOk() (*float32, bool)`

GetSumTaxOk returns a tuple with the SumTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumTax

`func (o *ModelCreditNote) SetSumTax(v float32)`

SetSumTax sets SumTax field to given value.

### HasSumTax

`func (o *ModelCreditNote) HasSumTax() bool`

HasSumTax returns a boolean if a field has been set.

### GetSumGross

`func (o *ModelCreditNote) GetSumGross() float32`

GetSumGross returns the SumGross field if non-nil, zero value otherwise.

### GetSumGrossOk

`func (o *ModelCreditNote) GetSumGrossOk() (*float32, bool)`

GetSumGrossOk returns a tuple with the SumGross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumGross

`func (o *ModelCreditNote) SetSumGross(v float32)`

SetSumGross sets SumGross field to given value.

### HasSumGross

`func (o *ModelCreditNote) HasSumGross() bool`

HasSumGross returns a boolean if a field has been set.

### GetSumDiscounts

`func (o *ModelCreditNote) GetSumDiscounts() float32`

GetSumDiscounts returns the SumDiscounts field if non-nil, zero value otherwise.

### GetSumDiscountsOk

`func (o *ModelCreditNote) GetSumDiscountsOk() (*float32, bool)`

GetSumDiscountsOk returns a tuple with the SumDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumDiscounts

`func (o *ModelCreditNote) SetSumDiscounts(v float32)`

SetSumDiscounts sets SumDiscounts field to given value.

### HasSumDiscounts

`func (o *ModelCreditNote) HasSumDiscounts() bool`

HasSumDiscounts returns a boolean if a field has been set.

### GetSumNetForeignCurrency

`func (o *ModelCreditNote) GetSumNetForeignCurrency() float32`

GetSumNetForeignCurrency returns the SumNetForeignCurrency field if non-nil, zero value otherwise.

### GetSumNetForeignCurrencyOk

`func (o *ModelCreditNote) GetSumNetForeignCurrencyOk() (*float32, bool)`

GetSumNetForeignCurrencyOk returns a tuple with the SumNetForeignCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumNetForeignCurrency

`func (o *ModelCreditNote) SetSumNetForeignCurrency(v float32)`

SetSumNetForeignCurrency sets SumNetForeignCurrency field to given value.

### HasSumNetForeignCurrency

`func (o *ModelCreditNote) HasSumNetForeignCurrency() bool`

HasSumNetForeignCurrency returns a boolean if a field has been set.

### GetSumTaxForeignCurrency

`func (o *ModelCreditNote) GetSumTaxForeignCurrency() float32`

GetSumTaxForeignCurrency returns the SumTaxForeignCurrency field if non-nil, zero value otherwise.

### GetSumTaxForeignCurrencyOk

`func (o *ModelCreditNote) GetSumTaxForeignCurrencyOk() (*float32, bool)`

GetSumTaxForeignCurrencyOk returns a tuple with the SumTaxForeignCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumTaxForeignCurrency

`func (o *ModelCreditNote) SetSumTaxForeignCurrency(v float32)`

SetSumTaxForeignCurrency sets SumTaxForeignCurrency field to given value.

### HasSumTaxForeignCurrency

`func (o *ModelCreditNote) HasSumTaxForeignCurrency() bool`

HasSumTaxForeignCurrency returns a boolean if a field has been set.

### GetSumGrossForeignCurrency

`func (o *ModelCreditNote) GetSumGrossForeignCurrency() float32`

GetSumGrossForeignCurrency returns the SumGrossForeignCurrency field if non-nil, zero value otherwise.

### GetSumGrossForeignCurrencyOk

`func (o *ModelCreditNote) GetSumGrossForeignCurrencyOk() (*float32, bool)`

GetSumGrossForeignCurrencyOk returns a tuple with the SumGrossForeignCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumGrossForeignCurrency

`func (o *ModelCreditNote) SetSumGrossForeignCurrency(v float32)`

SetSumGrossForeignCurrency sets SumGrossForeignCurrency field to given value.

### HasSumGrossForeignCurrency

`func (o *ModelCreditNote) HasSumGrossForeignCurrency() bool`

HasSumGrossForeignCurrency returns a boolean if a field has been set.

### GetSumDiscountsForeignCurrency

`func (o *ModelCreditNote) GetSumDiscountsForeignCurrency() float32`

GetSumDiscountsForeignCurrency returns the SumDiscountsForeignCurrency field if non-nil, zero value otherwise.

### GetSumDiscountsForeignCurrencyOk

`func (o *ModelCreditNote) GetSumDiscountsForeignCurrencyOk() (*float32, bool)`

GetSumDiscountsForeignCurrencyOk returns a tuple with the SumDiscountsForeignCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumDiscountsForeignCurrency

`func (o *ModelCreditNote) SetSumDiscountsForeignCurrency(v float32)`

SetSumDiscountsForeignCurrency sets SumDiscountsForeignCurrency field to given value.

### HasSumDiscountsForeignCurrency

`func (o *ModelCreditNote) HasSumDiscountsForeignCurrency() bool`

HasSumDiscountsForeignCurrency returns a boolean if a field has been set.

### GetCustomerInternalNote

`func (o *ModelCreditNote) GetCustomerInternalNote() string`

GetCustomerInternalNote returns the CustomerInternalNote field if non-nil, zero value otherwise.

### GetCustomerInternalNoteOk

`func (o *ModelCreditNote) GetCustomerInternalNoteOk() (*string, bool)`

GetCustomerInternalNoteOk returns a tuple with the CustomerInternalNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerInternalNote

`func (o *ModelCreditNote) SetCustomerInternalNote(v string)`

SetCustomerInternalNote sets CustomerInternalNote field to given value.

### HasCustomerInternalNote

`func (o *ModelCreditNote) HasCustomerInternalNote() bool`

HasCustomerInternalNote returns a boolean if a field has been set.

### SetCustomerInternalNoteNil

`func (o *ModelCreditNote) SetCustomerInternalNoteNil(b bool)`

 SetCustomerInternalNoteNil sets the value for CustomerInternalNote to be an explicit nil

### UnsetCustomerInternalNote
`func (o *ModelCreditNote) UnsetCustomerInternalNote()`

UnsetCustomerInternalNote ensures that no value is present for CustomerInternalNote, not even an explicit nil
### GetShowNet

`func (o *ModelCreditNote) GetShowNet() bool`

GetShowNet returns the ShowNet field if non-nil, zero value otherwise.

### GetShowNetOk

`func (o *ModelCreditNote) GetShowNetOk() (*bool, bool)`

GetShowNetOk returns a tuple with the ShowNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowNet

`func (o *ModelCreditNote) SetShowNet(v bool)`

SetShowNet sets ShowNet field to given value.

### HasShowNet

`func (o *ModelCreditNote) HasShowNet() bool`

HasShowNet returns a boolean if a field has been set.

### GetSendType

`func (o *ModelCreditNote) GetSendType() string`

GetSendType returns the SendType field if non-nil, zero value otherwise.

### GetSendTypeOk

`func (o *ModelCreditNote) GetSendTypeOk() (*string, bool)`

GetSendTypeOk returns a tuple with the SendType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendType

`func (o *ModelCreditNote) SetSendType(v string)`

SetSendType sets SendType field to given value.

### HasSendType

`func (o *ModelCreditNote) HasSendType() bool`

HasSendType returns a boolean if a field has been set.

### SetSendTypeNil

`func (o *ModelCreditNote) SetSendTypeNil(b bool)`

 SetSendTypeNil sets the value for SendType to be an explicit nil

### UnsetSendType
`func (o *ModelCreditNote) UnsetSendType()`

UnsetSendType ensures that no value is present for SendType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


