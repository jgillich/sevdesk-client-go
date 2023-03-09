# ModelOrderUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The order id | [optional] [readonly] 
**ObjectName** | Pointer to **string** | The order object name | [optional] [readonly] 
**Create** | Pointer to **time.Time** | Date of order creation | [optional] [readonly] 
**Update** | Pointer to **time.Time** | Date of last order update | [optional] [readonly] 
**OrderNumber** | Pointer to **string** | The order number | [optional] 
**Contact** | Pointer to [**NullableModelOrderUpdateContact**](ModelOrderUpdateContact.md) |  | [optional] 
**OrderDate** | Pointer to **NullableTime** | Needs to be provided as timestamp or dd.mm.yyyy | [optional] 
**Status** | Pointer to **NullableInt32** | Please have a look in       &lt;a href&#x3D;&#39;https://api.sevdesk.de/#section/Types-and-status-of-orders&#39;&gt;status of orders&lt;/a&gt;      to see what the different status codes mean | [optional] 
**Header** | Pointer to **NullableString** | Normally consist of prefix plus the order number | [optional] 
**HeadText** | Pointer to **NullableString** | Certain html tags can be used here to format your text | [optional] 
**FootText** | Pointer to **NullableString** | Certain html tags can be used here to format your text | [optional] 
**AddressCountry** | Pointer to [**NullableModelOrderUpdateAddressCountry**](ModelOrderUpdateAddressCountry.md) |  | [optional] 
**CreateUser** | Pointer to [**ModelOrderUpdateCreateUser**](ModelOrderUpdateCreateUser.md) |  | [optional] 
**SevClient** | Pointer to [**ModelOrderUpdateSevClient**](ModelOrderUpdateSevClient.md) |  | [optional] 
**DeliveryTerms** | Pointer to **NullableString** | Delivery terms of the order | [optional] 
**PaymentTerms** | Pointer to **NullableString** | Payment terms of the order | [optional] 
**Origin** | Pointer to [**NullableModelOrderOrigin**](ModelOrderOrigin.md) |  | [optional] 
**Version** | Pointer to **NullableInt32** | Version of the order.&lt;br&gt;      Can be used if you have multiple drafts for the same order.&lt;br&gt;      Should start with 0 | [optional] 
**SmallSettlement** | Pointer to **NullableBool** | Defines if the client uses the small settlement scheme.      If yes, the order must not contain any vat | [optional] 
**ContactPerson** | Pointer to [**ModelOrderUpdateContactPerson**](ModelOrderUpdateContactPerson.md) |  | [optional] 
**TaxRate** | Pointer to **NullableFloat32** | Is overwritten by order position tax rates | [optional] 
**TaxSet** | Pointer to [**NullableModelOrderUpdateTaxSet**](ModelOrderUpdateTaxSet.md) |  | [optional] 
**TaxText** | Pointer to **NullableString** | A common tax text would be &#39;Umsatzsteuer 19%&#39; | [optional] 
**TaxType** | Pointer to **NullableString** | Tax type of the order. There are four tax types: 1. default - Umsatzsteuer ausweisen 2. eu - Steuerfreie innergemeinschaftliche Lieferung (Europäische Union) 3. noteu - Steuerschuldnerschaft des Leistungsempfängers (außerhalb EU, z. B. Schweiz) 4. custom - Using custom tax set 5. ss - Not subject to VAT according to §19 1 UStG Tax rates are heavily connected to the tax type used. | [optional] 
**OrderType** | Pointer to **NullableString** | Type of the order. For more information on the different types, check      &lt;a href&#x3D;&#39;https://api.sevdesk.de/#section/Types-and-status-of-orders&#39;&gt;this&lt;/a&gt;   | [optional] 
**SendDate** | Pointer to **NullableTime** | The date the order was sent to the customer | [optional] 
**Address** | Pointer to **NullableString** | Complete address of the recipient including name, street, city, zip and country.&lt;br&gt;       Line breaks can be used and will be displayed on the invoice pdf. | [optional] 
**Currency** | Pointer to **NullableString** | Currency used in the order. Needs to be currency code according to ISO-4217 | [optional] 
**SumNet** | Pointer to **float32** | Net sum of the order | [optional] [readonly] 
**SumTax** | Pointer to **float32** | Tax sum of the order | [optional] [readonly] 
**SumGross** | Pointer to **float32** | Gross sum of the order | [optional] [readonly] 
**SumDiscounts** | Pointer to **float32** | Sum of all discounts in the order | [optional] [readonly] 
**SumNetForeignCurrency** | Pointer to **float32** | Net sum of the order in the foreign currency | [optional] [readonly] 
**SumTaxForeignCurrency** | Pointer to **float32** | Tax sum of the order in the foreign currency | [optional] [readonly] 
**SumGrossForeignCurrency** | Pointer to **float32** | Gross sum of the order in the foreign currency | [optional] [readonly] 
**SumDiscountsForeignCurrency** | Pointer to **float32** | Discounts sum of the order in the foreign currency | [optional] [readonly] 
**CustomerInternalNote** | Pointer to **NullableString** | Internal note of the customer. Contains data entered into field &#39;Referenz/Bestellnummer&#39; | [optional] 
**ShowNet** | Pointer to **NullableBool** | If true, the net amount of each position will be shown on the order. Otherwise gross amount | [optional] 
**SendType** | Pointer to **NullableString** | Type which was used to send the order. IMPORTANT: Please refer to the order section of the       *     API-Overview to understand how this attribute can be used before using it! | [optional] 

## Methods

### NewModelOrderUpdate

`func NewModelOrderUpdate() *ModelOrderUpdate`

NewModelOrderUpdate instantiates a new ModelOrderUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelOrderUpdateWithDefaults

`func NewModelOrderUpdateWithDefaults() *ModelOrderUpdate`

NewModelOrderUpdateWithDefaults instantiates a new ModelOrderUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelOrderUpdate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelOrderUpdate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelOrderUpdate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelOrderUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectName

`func (o *ModelOrderUpdate) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ModelOrderUpdate) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ModelOrderUpdate) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ModelOrderUpdate) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetCreate

`func (o *ModelOrderUpdate) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ModelOrderUpdate) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ModelOrderUpdate) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ModelOrderUpdate) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUpdate

`func (o *ModelOrderUpdate) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ModelOrderUpdate) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ModelOrderUpdate) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ModelOrderUpdate) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetOrderNumber

`func (o *ModelOrderUpdate) GetOrderNumber() string`

GetOrderNumber returns the OrderNumber field if non-nil, zero value otherwise.

### GetOrderNumberOk

`func (o *ModelOrderUpdate) GetOrderNumberOk() (*string, bool)`

GetOrderNumberOk returns a tuple with the OrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNumber

`func (o *ModelOrderUpdate) SetOrderNumber(v string)`

SetOrderNumber sets OrderNumber field to given value.

### HasOrderNumber

`func (o *ModelOrderUpdate) HasOrderNumber() bool`

HasOrderNumber returns a boolean if a field has been set.

### GetContact

`func (o *ModelOrderUpdate) GetContact() ModelOrderUpdateContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ModelOrderUpdate) GetContactOk() (*ModelOrderUpdateContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ModelOrderUpdate) SetContact(v ModelOrderUpdateContact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *ModelOrderUpdate) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *ModelOrderUpdate) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *ModelOrderUpdate) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetOrderDate

`func (o *ModelOrderUpdate) GetOrderDate() time.Time`

GetOrderDate returns the OrderDate field if non-nil, zero value otherwise.

### GetOrderDateOk

`func (o *ModelOrderUpdate) GetOrderDateOk() (*time.Time, bool)`

GetOrderDateOk returns a tuple with the OrderDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDate

`func (o *ModelOrderUpdate) SetOrderDate(v time.Time)`

SetOrderDate sets OrderDate field to given value.

### HasOrderDate

`func (o *ModelOrderUpdate) HasOrderDate() bool`

HasOrderDate returns a boolean if a field has been set.

### SetOrderDateNil

`func (o *ModelOrderUpdate) SetOrderDateNil(b bool)`

 SetOrderDateNil sets the value for OrderDate to be an explicit nil

### UnsetOrderDate
`func (o *ModelOrderUpdate) UnsetOrderDate()`

UnsetOrderDate ensures that no value is present for OrderDate, not even an explicit nil
### GetStatus

`func (o *ModelOrderUpdate) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelOrderUpdate) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelOrderUpdate) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelOrderUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ModelOrderUpdate) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ModelOrderUpdate) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetHeader

`func (o *ModelOrderUpdate) GetHeader() string`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *ModelOrderUpdate) GetHeaderOk() (*string, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *ModelOrderUpdate) SetHeader(v string)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *ModelOrderUpdate) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### SetHeaderNil

`func (o *ModelOrderUpdate) SetHeaderNil(b bool)`

 SetHeaderNil sets the value for Header to be an explicit nil

### UnsetHeader
`func (o *ModelOrderUpdate) UnsetHeader()`

UnsetHeader ensures that no value is present for Header, not even an explicit nil
### GetHeadText

`func (o *ModelOrderUpdate) GetHeadText() string`

GetHeadText returns the HeadText field if non-nil, zero value otherwise.

### GetHeadTextOk

`func (o *ModelOrderUpdate) GetHeadTextOk() (*string, bool)`

GetHeadTextOk returns a tuple with the HeadText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadText

`func (o *ModelOrderUpdate) SetHeadText(v string)`

SetHeadText sets HeadText field to given value.

### HasHeadText

`func (o *ModelOrderUpdate) HasHeadText() bool`

HasHeadText returns a boolean if a field has been set.

### SetHeadTextNil

`func (o *ModelOrderUpdate) SetHeadTextNil(b bool)`

 SetHeadTextNil sets the value for HeadText to be an explicit nil

### UnsetHeadText
`func (o *ModelOrderUpdate) UnsetHeadText()`

UnsetHeadText ensures that no value is present for HeadText, not even an explicit nil
### GetFootText

`func (o *ModelOrderUpdate) GetFootText() string`

GetFootText returns the FootText field if non-nil, zero value otherwise.

### GetFootTextOk

`func (o *ModelOrderUpdate) GetFootTextOk() (*string, bool)`

GetFootTextOk returns a tuple with the FootText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFootText

`func (o *ModelOrderUpdate) SetFootText(v string)`

SetFootText sets FootText field to given value.

### HasFootText

`func (o *ModelOrderUpdate) HasFootText() bool`

HasFootText returns a boolean if a field has been set.

### SetFootTextNil

`func (o *ModelOrderUpdate) SetFootTextNil(b bool)`

 SetFootTextNil sets the value for FootText to be an explicit nil

### UnsetFootText
`func (o *ModelOrderUpdate) UnsetFootText()`

UnsetFootText ensures that no value is present for FootText, not even an explicit nil
### GetAddressCountry

`func (o *ModelOrderUpdate) GetAddressCountry() ModelOrderUpdateAddressCountry`

GetAddressCountry returns the AddressCountry field if non-nil, zero value otherwise.

### GetAddressCountryOk

`func (o *ModelOrderUpdate) GetAddressCountryOk() (*ModelOrderUpdateAddressCountry, bool)`

GetAddressCountryOk returns a tuple with the AddressCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCountry

`func (o *ModelOrderUpdate) SetAddressCountry(v ModelOrderUpdateAddressCountry)`

SetAddressCountry sets AddressCountry field to given value.

### HasAddressCountry

`func (o *ModelOrderUpdate) HasAddressCountry() bool`

HasAddressCountry returns a boolean if a field has been set.

### SetAddressCountryNil

`func (o *ModelOrderUpdate) SetAddressCountryNil(b bool)`

 SetAddressCountryNil sets the value for AddressCountry to be an explicit nil

### UnsetAddressCountry
`func (o *ModelOrderUpdate) UnsetAddressCountry()`

UnsetAddressCountry ensures that no value is present for AddressCountry, not even an explicit nil
### GetCreateUser

`func (o *ModelOrderUpdate) GetCreateUser() ModelOrderUpdateCreateUser`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *ModelOrderUpdate) GetCreateUserOk() (*ModelOrderUpdateCreateUser, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *ModelOrderUpdate) SetCreateUser(v ModelOrderUpdateCreateUser)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *ModelOrderUpdate) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetSevClient

`func (o *ModelOrderUpdate) GetSevClient() ModelOrderUpdateSevClient`

GetSevClient returns the SevClient field if non-nil, zero value otherwise.

### GetSevClientOk

`func (o *ModelOrderUpdate) GetSevClientOk() (*ModelOrderUpdateSevClient, bool)`

GetSevClientOk returns a tuple with the SevClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSevClient

`func (o *ModelOrderUpdate) SetSevClient(v ModelOrderUpdateSevClient)`

SetSevClient sets SevClient field to given value.

### HasSevClient

`func (o *ModelOrderUpdate) HasSevClient() bool`

HasSevClient returns a boolean if a field has been set.

### GetDeliveryTerms

`func (o *ModelOrderUpdate) GetDeliveryTerms() string`

GetDeliveryTerms returns the DeliveryTerms field if non-nil, zero value otherwise.

### GetDeliveryTermsOk

`func (o *ModelOrderUpdate) GetDeliveryTermsOk() (*string, bool)`

GetDeliveryTermsOk returns a tuple with the DeliveryTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTerms

`func (o *ModelOrderUpdate) SetDeliveryTerms(v string)`

SetDeliveryTerms sets DeliveryTerms field to given value.

### HasDeliveryTerms

`func (o *ModelOrderUpdate) HasDeliveryTerms() bool`

HasDeliveryTerms returns a boolean if a field has been set.

### SetDeliveryTermsNil

`func (o *ModelOrderUpdate) SetDeliveryTermsNil(b bool)`

 SetDeliveryTermsNil sets the value for DeliveryTerms to be an explicit nil

### UnsetDeliveryTerms
`func (o *ModelOrderUpdate) UnsetDeliveryTerms()`

UnsetDeliveryTerms ensures that no value is present for DeliveryTerms, not even an explicit nil
### GetPaymentTerms

`func (o *ModelOrderUpdate) GetPaymentTerms() string`

GetPaymentTerms returns the PaymentTerms field if non-nil, zero value otherwise.

### GetPaymentTermsOk

`func (o *ModelOrderUpdate) GetPaymentTermsOk() (*string, bool)`

GetPaymentTermsOk returns a tuple with the PaymentTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTerms

`func (o *ModelOrderUpdate) SetPaymentTerms(v string)`

SetPaymentTerms sets PaymentTerms field to given value.

### HasPaymentTerms

`func (o *ModelOrderUpdate) HasPaymentTerms() bool`

HasPaymentTerms returns a boolean if a field has been set.

### SetPaymentTermsNil

`func (o *ModelOrderUpdate) SetPaymentTermsNil(b bool)`

 SetPaymentTermsNil sets the value for PaymentTerms to be an explicit nil

### UnsetPaymentTerms
`func (o *ModelOrderUpdate) UnsetPaymentTerms()`

UnsetPaymentTerms ensures that no value is present for PaymentTerms, not even an explicit nil
### GetOrigin

`func (o *ModelOrderUpdate) GetOrigin() ModelOrderOrigin`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *ModelOrderUpdate) GetOriginOk() (*ModelOrderOrigin, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *ModelOrderUpdate) SetOrigin(v ModelOrderOrigin)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *ModelOrderUpdate) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### SetOriginNil

`func (o *ModelOrderUpdate) SetOriginNil(b bool)`

 SetOriginNil sets the value for Origin to be an explicit nil

### UnsetOrigin
`func (o *ModelOrderUpdate) UnsetOrigin()`

UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
### GetVersion

`func (o *ModelOrderUpdate) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ModelOrderUpdate) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ModelOrderUpdate) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ModelOrderUpdate) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ModelOrderUpdate) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ModelOrderUpdate) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetSmallSettlement

`func (o *ModelOrderUpdate) GetSmallSettlement() bool`

GetSmallSettlement returns the SmallSettlement field if non-nil, zero value otherwise.

### GetSmallSettlementOk

`func (o *ModelOrderUpdate) GetSmallSettlementOk() (*bool, bool)`

GetSmallSettlementOk returns a tuple with the SmallSettlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallSettlement

`func (o *ModelOrderUpdate) SetSmallSettlement(v bool)`

SetSmallSettlement sets SmallSettlement field to given value.

### HasSmallSettlement

`func (o *ModelOrderUpdate) HasSmallSettlement() bool`

HasSmallSettlement returns a boolean if a field has been set.

### SetSmallSettlementNil

`func (o *ModelOrderUpdate) SetSmallSettlementNil(b bool)`

 SetSmallSettlementNil sets the value for SmallSettlement to be an explicit nil

### UnsetSmallSettlement
`func (o *ModelOrderUpdate) UnsetSmallSettlement()`

UnsetSmallSettlement ensures that no value is present for SmallSettlement, not even an explicit nil
### GetContactPerson

`func (o *ModelOrderUpdate) GetContactPerson() ModelOrderUpdateContactPerson`

GetContactPerson returns the ContactPerson field if non-nil, zero value otherwise.

### GetContactPersonOk

`func (o *ModelOrderUpdate) GetContactPersonOk() (*ModelOrderUpdateContactPerson, bool)`

GetContactPersonOk returns a tuple with the ContactPerson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPerson

`func (o *ModelOrderUpdate) SetContactPerson(v ModelOrderUpdateContactPerson)`

SetContactPerson sets ContactPerson field to given value.

### HasContactPerson

`func (o *ModelOrderUpdate) HasContactPerson() bool`

HasContactPerson returns a boolean if a field has been set.

### GetTaxRate

`func (o *ModelOrderUpdate) GetTaxRate() float32`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *ModelOrderUpdate) GetTaxRateOk() (*float32, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *ModelOrderUpdate) SetTaxRate(v float32)`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *ModelOrderUpdate) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### SetTaxRateNil

`func (o *ModelOrderUpdate) SetTaxRateNil(b bool)`

 SetTaxRateNil sets the value for TaxRate to be an explicit nil

### UnsetTaxRate
`func (o *ModelOrderUpdate) UnsetTaxRate()`

UnsetTaxRate ensures that no value is present for TaxRate, not even an explicit nil
### GetTaxSet

`func (o *ModelOrderUpdate) GetTaxSet() ModelOrderUpdateTaxSet`

GetTaxSet returns the TaxSet field if non-nil, zero value otherwise.

### GetTaxSetOk

`func (o *ModelOrderUpdate) GetTaxSetOk() (*ModelOrderUpdateTaxSet, bool)`

GetTaxSetOk returns a tuple with the TaxSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxSet

`func (o *ModelOrderUpdate) SetTaxSet(v ModelOrderUpdateTaxSet)`

SetTaxSet sets TaxSet field to given value.

### HasTaxSet

`func (o *ModelOrderUpdate) HasTaxSet() bool`

HasTaxSet returns a boolean if a field has been set.

### SetTaxSetNil

`func (o *ModelOrderUpdate) SetTaxSetNil(b bool)`

 SetTaxSetNil sets the value for TaxSet to be an explicit nil

### UnsetTaxSet
`func (o *ModelOrderUpdate) UnsetTaxSet()`

UnsetTaxSet ensures that no value is present for TaxSet, not even an explicit nil
### GetTaxText

`func (o *ModelOrderUpdate) GetTaxText() string`

GetTaxText returns the TaxText field if non-nil, zero value otherwise.

### GetTaxTextOk

`func (o *ModelOrderUpdate) GetTaxTextOk() (*string, bool)`

GetTaxTextOk returns a tuple with the TaxText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxText

`func (o *ModelOrderUpdate) SetTaxText(v string)`

SetTaxText sets TaxText field to given value.

### HasTaxText

`func (o *ModelOrderUpdate) HasTaxText() bool`

HasTaxText returns a boolean if a field has been set.

### SetTaxTextNil

`func (o *ModelOrderUpdate) SetTaxTextNil(b bool)`

 SetTaxTextNil sets the value for TaxText to be an explicit nil

### UnsetTaxText
`func (o *ModelOrderUpdate) UnsetTaxText()`

UnsetTaxText ensures that no value is present for TaxText, not even an explicit nil
### GetTaxType

`func (o *ModelOrderUpdate) GetTaxType() string`

GetTaxType returns the TaxType field if non-nil, zero value otherwise.

### GetTaxTypeOk

`func (o *ModelOrderUpdate) GetTaxTypeOk() (*string, bool)`

GetTaxTypeOk returns a tuple with the TaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxType

`func (o *ModelOrderUpdate) SetTaxType(v string)`

SetTaxType sets TaxType field to given value.

### HasTaxType

`func (o *ModelOrderUpdate) HasTaxType() bool`

HasTaxType returns a boolean if a field has been set.

### SetTaxTypeNil

`func (o *ModelOrderUpdate) SetTaxTypeNil(b bool)`

 SetTaxTypeNil sets the value for TaxType to be an explicit nil

### UnsetTaxType
`func (o *ModelOrderUpdate) UnsetTaxType()`

UnsetTaxType ensures that no value is present for TaxType, not even an explicit nil
### GetOrderType

`func (o *ModelOrderUpdate) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *ModelOrderUpdate) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *ModelOrderUpdate) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *ModelOrderUpdate) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### SetOrderTypeNil

`func (o *ModelOrderUpdate) SetOrderTypeNil(b bool)`

 SetOrderTypeNil sets the value for OrderType to be an explicit nil

### UnsetOrderType
`func (o *ModelOrderUpdate) UnsetOrderType()`

UnsetOrderType ensures that no value is present for OrderType, not even an explicit nil
### GetSendDate

`func (o *ModelOrderUpdate) GetSendDate() time.Time`

GetSendDate returns the SendDate field if non-nil, zero value otherwise.

### GetSendDateOk

`func (o *ModelOrderUpdate) GetSendDateOk() (*time.Time, bool)`

GetSendDateOk returns a tuple with the SendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendDate

`func (o *ModelOrderUpdate) SetSendDate(v time.Time)`

SetSendDate sets SendDate field to given value.

### HasSendDate

`func (o *ModelOrderUpdate) HasSendDate() bool`

HasSendDate returns a boolean if a field has been set.

### SetSendDateNil

`func (o *ModelOrderUpdate) SetSendDateNil(b bool)`

 SetSendDateNil sets the value for SendDate to be an explicit nil

### UnsetSendDate
`func (o *ModelOrderUpdate) UnsetSendDate()`

UnsetSendDate ensures that no value is present for SendDate, not even an explicit nil
### GetAddress

`func (o *ModelOrderUpdate) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ModelOrderUpdate) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ModelOrderUpdate) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ModelOrderUpdate) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *ModelOrderUpdate) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *ModelOrderUpdate) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetCurrency

`func (o *ModelOrderUpdate) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ModelOrderUpdate) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ModelOrderUpdate) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ModelOrderUpdate) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *ModelOrderUpdate) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *ModelOrderUpdate) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetSumNet

`func (o *ModelOrderUpdate) GetSumNet() float32`

GetSumNet returns the SumNet field if non-nil, zero value otherwise.

### GetSumNetOk

`func (o *ModelOrderUpdate) GetSumNetOk() (*float32, bool)`

GetSumNetOk returns a tuple with the SumNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumNet

`func (o *ModelOrderUpdate) SetSumNet(v float32)`

SetSumNet sets SumNet field to given value.

### HasSumNet

`func (o *ModelOrderUpdate) HasSumNet() bool`

HasSumNet returns a boolean if a field has been set.

### GetSumTax

`func (o *ModelOrderUpdate) GetSumTax() float32`

GetSumTax returns the SumTax field if non-nil, zero value otherwise.

### GetSumTaxOk

`func (o *ModelOrderUpdate) GetSumTaxOk() (*float32, bool)`

GetSumTaxOk returns a tuple with the SumTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumTax

`func (o *ModelOrderUpdate) SetSumTax(v float32)`

SetSumTax sets SumTax field to given value.

### HasSumTax

`func (o *ModelOrderUpdate) HasSumTax() bool`

HasSumTax returns a boolean if a field has been set.

### GetSumGross

`func (o *ModelOrderUpdate) GetSumGross() float32`

GetSumGross returns the SumGross field if non-nil, zero value otherwise.

### GetSumGrossOk

`func (o *ModelOrderUpdate) GetSumGrossOk() (*float32, bool)`

GetSumGrossOk returns a tuple with the SumGross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumGross

`func (o *ModelOrderUpdate) SetSumGross(v float32)`

SetSumGross sets SumGross field to given value.

### HasSumGross

`func (o *ModelOrderUpdate) HasSumGross() bool`

HasSumGross returns a boolean if a field has been set.

### GetSumDiscounts

`func (o *ModelOrderUpdate) GetSumDiscounts() float32`

GetSumDiscounts returns the SumDiscounts field if non-nil, zero value otherwise.

### GetSumDiscountsOk

`func (o *ModelOrderUpdate) GetSumDiscountsOk() (*float32, bool)`

GetSumDiscountsOk returns a tuple with the SumDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumDiscounts

`func (o *ModelOrderUpdate) SetSumDiscounts(v float32)`

SetSumDiscounts sets SumDiscounts field to given value.

### HasSumDiscounts

`func (o *ModelOrderUpdate) HasSumDiscounts() bool`

HasSumDiscounts returns a boolean if a field has been set.

### GetSumNetForeignCurrency

`func (o *ModelOrderUpdate) GetSumNetForeignCurrency() float32`

GetSumNetForeignCurrency returns the SumNetForeignCurrency field if non-nil, zero value otherwise.

### GetSumNetForeignCurrencyOk

`func (o *ModelOrderUpdate) GetSumNetForeignCurrencyOk() (*float32, bool)`

GetSumNetForeignCurrencyOk returns a tuple with the SumNetForeignCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumNetForeignCurrency

`func (o *ModelOrderUpdate) SetSumNetForeignCurrency(v float32)`

SetSumNetForeignCurrency sets SumNetForeignCurrency field to given value.

### HasSumNetForeignCurrency

`func (o *ModelOrderUpdate) HasSumNetForeignCurrency() bool`

HasSumNetForeignCurrency returns a boolean if a field has been set.

### GetSumTaxForeignCurrency

`func (o *ModelOrderUpdate) GetSumTaxForeignCurrency() float32`

GetSumTaxForeignCurrency returns the SumTaxForeignCurrency field if non-nil, zero value otherwise.

### GetSumTaxForeignCurrencyOk

`func (o *ModelOrderUpdate) GetSumTaxForeignCurrencyOk() (*float32, bool)`

GetSumTaxForeignCurrencyOk returns a tuple with the SumTaxForeignCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumTaxForeignCurrency

`func (o *ModelOrderUpdate) SetSumTaxForeignCurrency(v float32)`

SetSumTaxForeignCurrency sets SumTaxForeignCurrency field to given value.

### HasSumTaxForeignCurrency

`func (o *ModelOrderUpdate) HasSumTaxForeignCurrency() bool`

HasSumTaxForeignCurrency returns a boolean if a field has been set.

### GetSumGrossForeignCurrency

`func (o *ModelOrderUpdate) GetSumGrossForeignCurrency() float32`

GetSumGrossForeignCurrency returns the SumGrossForeignCurrency field if non-nil, zero value otherwise.

### GetSumGrossForeignCurrencyOk

`func (o *ModelOrderUpdate) GetSumGrossForeignCurrencyOk() (*float32, bool)`

GetSumGrossForeignCurrencyOk returns a tuple with the SumGrossForeignCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumGrossForeignCurrency

`func (o *ModelOrderUpdate) SetSumGrossForeignCurrency(v float32)`

SetSumGrossForeignCurrency sets SumGrossForeignCurrency field to given value.

### HasSumGrossForeignCurrency

`func (o *ModelOrderUpdate) HasSumGrossForeignCurrency() bool`

HasSumGrossForeignCurrency returns a boolean if a field has been set.

### GetSumDiscountsForeignCurrency

`func (o *ModelOrderUpdate) GetSumDiscountsForeignCurrency() float32`

GetSumDiscountsForeignCurrency returns the SumDiscountsForeignCurrency field if non-nil, zero value otherwise.

### GetSumDiscountsForeignCurrencyOk

`func (o *ModelOrderUpdate) GetSumDiscountsForeignCurrencyOk() (*float32, bool)`

GetSumDiscountsForeignCurrencyOk returns a tuple with the SumDiscountsForeignCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumDiscountsForeignCurrency

`func (o *ModelOrderUpdate) SetSumDiscountsForeignCurrency(v float32)`

SetSumDiscountsForeignCurrency sets SumDiscountsForeignCurrency field to given value.

### HasSumDiscountsForeignCurrency

`func (o *ModelOrderUpdate) HasSumDiscountsForeignCurrency() bool`

HasSumDiscountsForeignCurrency returns a boolean if a field has been set.

### GetCustomerInternalNote

`func (o *ModelOrderUpdate) GetCustomerInternalNote() string`

GetCustomerInternalNote returns the CustomerInternalNote field if non-nil, zero value otherwise.

### GetCustomerInternalNoteOk

`func (o *ModelOrderUpdate) GetCustomerInternalNoteOk() (*string, bool)`

GetCustomerInternalNoteOk returns a tuple with the CustomerInternalNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerInternalNote

`func (o *ModelOrderUpdate) SetCustomerInternalNote(v string)`

SetCustomerInternalNote sets CustomerInternalNote field to given value.

### HasCustomerInternalNote

`func (o *ModelOrderUpdate) HasCustomerInternalNote() bool`

HasCustomerInternalNote returns a boolean if a field has been set.

### SetCustomerInternalNoteNil

`func (o *ModelOrderUpdate) SetCustomerInternalNoteNil(b bool)`

 SetCustomerInternalNoteNil sets the value for CustomerInternalNote to be an explicit nil

### UnsetCustomerInternalNote
`func (o *ModelOrderUpdate) UnsetCustomerInternalNote()`

UnsetCustomerInternalNote ensures that no value is present for CustomerInternalNote, not even an explicit nil
### GetShowNet

`func (o *ModelOrderUpdate) GetShowNet() bool`

GetShowNet returns the ShowNet field if non-nil, zero value otherwise.

### GetShowNetOk

`func (o *ModelOrderUpdate) GetShowNetOk() (*bool, bool)`

GetShowNetOk returns a tuple with the ShowNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowNet

`func (o *ModelOrderUpdate) SetShowNet(v bool)`

SetShowNet sets ShowNet field to given value.

### HasShowNet

`func (o *ModelOrderUpdate) HasShowNet() bool`

HasShowNet returns a boolean if a field has been set.

### SetShowNetNil

`func (o *ModelOrderUpdate) SetShowNetNil(b bool)`

 SetShowNetNil sets the value for ShowNet to be an explicit nil

### UnsetShowNet
`func (o *ModelOrderUpdate) UnsetShowNet()`

UnsetShowNet ensures that no value is present for ShowNet, not even an explicit nil
### GetSendType

`func (o *ModelOrderUpdate) GetSendType() string`

GetSendType returns the SendType field if non-nil, zero value otherwise.

### GetSendTypeOk

`func (o *ModelOrderUpdate) GetSendTypeOk() (*string, bool)`

GetSendTypeOk returns a tuple with the SendType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendType

`func (o *ModelOrderUpdate) SetSendType(v string)`

SetSendType sets SendType field to given value.

### HasSendType

`func (o *ModelOrderUpdate) HasSendType() bool`

HasSendType returns a boolean if a field has been set.

### SetSendTypeNil

`func (o *ModelOrderUpdate) SetSendTypeNil(b bool)`

 SetSendTypeNil sets the value for SendType to be an explicit nil

### UnsetSendType
`func (o *ModelOrderUpdate) UnsetSendType()`

UnsetSendType ensures that no value is present for SendType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


