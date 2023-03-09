# ModelOrderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The order id | [optional] [readonly] 
**ObjectName** | Pointer to **string** | The order object name | [optional] [readonly] 
**Create** | Pointer to **time.Time** | Date of order creation | [optional] [readonly] 
**Update** | Pointer to **time.Time** | Date of last order update | [optional] [readonly] 
**OrderNumber** | Pointer to **string** | The order number | [optional] 
**Contact** | Pointer to [**ModelOrderResponseContact**](ModelOrderResponseContact.md) |  | [optional] 
**OrderDate** | Pointer to **time.Time** | Needs to be provided as timestamp or dd.mm.yyyy | [optional] 
**Status** | Pointer to **string** | Please have a look in       &lt;a href&#x3D;&#39;https://api.sevdesk.de/#section/Types-and-status-of-orders&#39;&gt;status of orders&lt;/a&gt;      to see what the different status codes mean | [optional] 
**Header** | Pointer to **string** | Normally consist of prefix plus the order number | [optional] 
**HeadText** | Pointer to **NullableString** | Certain html tags can be used here to format your text | [optional] 
**FootText** | Pointer to **NullableString** | Certain html tags can be used here to format your text | [optional] 
**AddressCountry** | Pointer to [**NullableModelOrderResponseAddressCountry**](ModelOrderResponseAddressCountry.md) |  | [optional] 
**CreateUser** | Pointer to [**ModelInvoiceResponseCreateUser**](ModelInvoiceResponseCreateUser.md) |  | [optional] 
**SevClient** | Pointer to [**ModelOrderResponseSevClient**](ModelOrderResponseSevClient.md) |  | [optional] 
**DeliveryTerms** | Pointer to **NullableString** | Delivery terms of the order | [optional] 
**PaymentTerms** | Pointer to **NullableString** | Payment terms of the order | [optional] 
**Origin** | Pointer to [**NullableModelOrderResponseOrigin**](ModelOrderResponseOrigin.md) |  | [optional] 
**Version** | Pointer to **string** | Version of the order.&lt;br&gt;      Can be used if you have multiple drafts for the same order.&lt;br&gt;      Should start with 0 | [optional] 
**SmallSettlement** | Pointer to **bool** | Defines if the client uses the small settlement scheme.      If yes, the order must not contain any vat | [optional] 
**ContactPerson** | Pointer to [**ModelOrderResponseContactPerson**](ModelOrderResponseContactPerson.md) |  | [optional] 
**TaxRate** | Pointer to **string** | Is overwritten by order position tax rates | [optional] 
**TaxSet** | Pointer to [**NullableModelOrderResponseTaxSet**](ModelOrderResponseTaxSet.md) |  | [optional] 
**TaxText** | Pointer to **string** | A common tax text would be &#39;Umsatzsteuer 19%&#39; | [optional] 
**TaxType** | Pointer to **string** | Tax type of the order. There are four tax types: 1. default - Umsatzsteuer ausweisen 2. eu - Steuerfreie innergemeinschaftliche Lieferung (Europäische Union) 3. noteu - Steuerschuldnerschaft des Leistungsempfängers (außerhalb EU, z. B. Schweiz) 4. custom - Using custom tax set 5. ss - Not subject to VAT according to §19 1 UStG Tax rates are heavily connected to the tax type used. | [optional] 
**OrderType** | Pointer to **string** | Type of the order. For more information on the different types, check      &lt;a href&#x3D;&#39;https://api.sevdesk.de/#section/Types-and-status-of-orders&#39;&gt;this&lt;/a&gt;   | [optional] 
**SendDate** | Pointer to **NullableTime** | The date the order was sent to the customer | [optional] 
**Address** | Pointer to **NullableString** | Complete address of the recipient including name, street, city, zip and country.&lt;br&gt;       Line breaks can be used and will be displayed on the invoice pdf. | [optional] 
**Currency** | Pointer to **string** | Currency used in the order. Needs to be currency code according to ISO-4217 | [optional] 
**SumNet** | Pointer to **string** | Net sum of the order | [optional] [readonly] 
**SumTax** | Pointer to **string** | Tax sum of the order | [optional] [readonly] 
**SumGross** | Pointer to **string** | Gross sum of the order | [optional] [readonly] 
**SumDiscounts** | Pointer to **string** | Sum of all discounts in the order | [optional] [readonly] 
**SumNetForeignCurrency** | Pointer to **string** | Net sum of the order in the foreign currency | [optional] [readonly] 
**SumTaxForeignCurrency** | Pointer to **string** | Tax sum of the order in the foreign currency | [optional] [readonly] 
**SumGrossForeignCurrency** | Pointer to **string** | Gross sum of the order in the foreign currency | [optional] [readonly] 
**SumDiscountsForeignCurrency** | Pointer to **string** | Discounts sum of the order in the foreign currency | [optional] [readonly] 
**CustomerInternalNote** | Pointer to **NullableString** | Internal note of the customer. Contains data entered into field &#39;Referenz/Bestellnummer&#39; | [optional] 
**ShowNet** | Pointer to **bool** | If true, the net amount of each position will be shown on the order. Otherwise gross amount | [optional] 
**SendType** | Pointer to **NullableString** | Type which was used to send the order. IMPORTANT: Please refer to the order section of the       *     API-Overview to understand how this attribute can be used before using it! | [optional] 

## Methods

### NewModelOrderResponse

`func NewModelOrderResponse() *ModelOrderResponse`

NewModelOrderResponse instantiates a new ModelOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelOrderResponseWithDefaults

`func NewModelOrderResponseWithDefaults() *ModelOrderResponse`

NewModelOrderResponseWithDefaults instantiates a new ModelOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelOrderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelOrderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelOrderResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelOrderResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectName

`func (o *ModelOrderResponse) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ModelOrderResponse) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ModelOrderResponse) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ModelOrderResponse) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetCreate

`func (o *ModelOrderResponse) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ModelOrderResponse) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ModelOrderResponse) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ModelOrderResponse) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUpdate

`func (o *ModelOrderResponse) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ModelOrderResponse) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ModelOrderResponse) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ModelOrderResponse) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetOrderNumber

`func (o *ModelOrderResponse) GetOrderNumber() string`

GetOrderNumber returns the OrderNumber field if non-nil, zero value otherwise.

### GetOrderNumberOk

`func (o *ModelOrderResponse) GetOrderNumberOk() (*string, bool)`

GetOrderNumberOk returns a tuple with the OrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNumber

`func (o *ModelOrderResponse) SetOrderNumber(v string)`

SetOrderNumber sets OrderNumber field to given value.

### HasOrderNumber

`func (o *ModelOrderResponse) HasOrderNumber() bool`

HasOrderNumber returns a boolean if a field has been set.

### GetContact

`func (o *ModelOrderResponse) GetContact() ModelOrderResponseContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ModelOrderResponse) GetContactOk() (*ModelOrderResponseContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ModelOrderResponse) SetContact(v ModelOrderResponseContact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *ModelOrderResponse) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetOrderDate

`func (o *ModelOrderResponse) GetOrderDate() time.Time`

GetOrderDate returns the OrderDate field if non-nil, zero value otherwise.

### GetOrderDateOk

`func (o *ModelOrderResponse) GetOrderDateOk() (*time.Time, bool)`

GetOrderDateOk returns a tuple with the OrderDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDate

`func (o *ModelOrderResponse) SetOrderDate(v time.Time)`

SetOrderDate sets OrderDate field to given value.

### HasOrderDate

`func (o *ModelOrderResponse) HasOrderDate() bool`

HasOrderDate returns a boolean if a field has been set.

### GetStatus

`func (o *ModelOrderResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelOrderResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelOrderResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelOrderResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetHeader

`func (o *ModelOrderResponse) GetHeader() string`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *ModelOrderResponse) GetHeaderOk() (*string, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *ModelOrderResponse) SetHeader(v string)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *ModelOrderResponse) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetHeadText

`func (o *ModelOrderResponse) GetHeadText() string`

GetHeadText returns the HeadText field if non-nil, zero value otherwise.

### GetHeadTextOk

`func (o *ModelOrderResponse) GetHeadTextOk() (*string, bool)`

GetHeadTextOk returns a tuple with the HeadText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadText

`func (o *ModelOrderResponse) SetHeadText(v string)`

SetHeadText sets HeadText field to given value.

### HasHeadText

`func (o *ModelOrderResponse) HasHeadText() bool`

HasHeadText returns a boolean if a field has been set.

### SetHeadTextNil

`func (o *ModelOrderResponse) SetHeadTextNil(b bool)`

 SetHeadTextNil sets the value for HeadText to be an explicit nil

### UnsetHeadText
`func (o *ModelOrderResponse) UnsetHeadText()`

UnsetHeadText ensures that no value is present for HeadText, not even an explicit nil
### GetFootText

`func (o *ModelOrderResponse) GetFootText() string`

GetFootText returns the FootText field if non-nil, zero value otherwise.

### GetFootTextOk

`func (o *ModelOrderResponse) GetFootTextOk() (*string, bool)`

GetFootTextOk returns a tuple with the FootText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFootText

`func (o *ModelOrderResponse) SetFootText(v string)`

SetFootText sets FootText field to given value.

### HasFootText

`func (o *ModelOrderResponse) HasFootText() bool`

HasFootText returns a boolean if a field has been set.

### SetFootTextNil

`func (o *ModelOrderResponse) SetFootTextNil(b bool)`

 SetFootTextNil sets the value for FootText to be an explicit nil

### UnsetFootText
`func (o *ModelOrderResponse) UnsetFootText()`

UnsetFootText ensures that no value is present for FootText, not even an explicit nil
### GetAddressCountry

`func (o *ModelOrderResponse) GetAddressCountry() ModelOrderResponseAddressCountry`

GetAddressCountry returns the AddressCountry field if non-nil, zero value otherwise.

### GetAddressCountryOk

`func (o *ModelOrderResponse) GetAddressCountryOk() (*ModelOrderResponseAddressCountry, bool)`

GetAddressCountryOk returns a tuple with the AddressCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCountry

`func (o *ModelOrderResponse) SetAddressCountry(v ModelOrderResponseAddressCountry)`

SetAddressCountry sets AddressCountry field to given value.

### HasAddressCountry

`func (o *ModelOrderResponse) HasAddressCountry() bool`

HasAddressCountry returns a boolean if a field has been set.

### SetAddressCountryNil

`func (o *ModelOrderResponse) SetAddressCountryNil(b bool)`

 SetAddressCountryNil sets the value for AddressCountry to be an explicit nil

### UnsetAddressCountry
`func (o *ModelOrderResponse) UnsetAddressCountry()`

UnsetAddressCountry ensures that no value is present for AddressCountry, not even an explicit nil
### GetCreateUser

`func (o *ModelOrderResponse) GetCreateUser() ModelInvoiceResponseCreateUser`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *ModelOrderResponse) GetCreateUserOk() (*ModelInvoiceResponseCreateUser, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *ModelOrderResponse) SetCreateUser(v ModelInvoiceResponseCreateUser)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *ModelOrderResponse) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetSevClient

`func (o *ModelOrderResponse) GetSevClient() ModelOrderResponseSevClient`

GetSevClient returns the SevClient field if non-nil, zero value otherwise.

### GetSevClientOk

`func (o *ModelOrderResponse) GetSevClientOk() (*ModelOrderResponseSevClient, bool)`

GetSevClientOk returns a tuple with the SevClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSevClient

`func (o *ModelOrderResponse) SetSevClient(v ModelOrderResponseSevClient)`

SetSevClient sets SevClient field to given value.

### HasSevClient

`func (o *ModelOrderResponse) HasSevClient() bool`

HasSevClient returns a boolean if a field has been set.

### GetDeliveryTerms

`func (o *ModelOrderResponse) GetDeliveryTerms() string`

GetDeliveryTerms returns the DeliveryTerms field if non-nil, zero value otherwise.

### GetDeliveryTermsOk

`func (o *ModelOrderResponse) GetDeliveryTermsOk() (*string, bool)`

GetDeliveryTermsOk returns a tuple with the DeliveryTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTerms

`func (o *ModelOrderResponse) SetDeliveryTerms(v string)`

SetDeliveryTerms sets DeliveryTerms field to given value.

### HasDeliveryTerms

`func (o *ModelOrderResponse) HasDeliveryTerms() bool`

HasDeliveryTerms returns a boolean if a field has been set.

### SetDeliveryTermsNil

`func (o *ModelOrderResponse) SetDeliveryTermsNil(b bool)`

 SetDeliveryTermsNil sets the value for DeliveryTerms to be an explicit nil

### UnsetDeliveryTerms
`func (o *ModelOrderResponse) UnsetDeliveryTerms()`

UnsetDeliveryTerms ensures that no value is present for DeliveryTerms, not even an explicit nil
### GetPaymentTerms

`func (o *ModelOrderResponse) GetPaymentTerms() string`

GetPaymentTerms returns the PaymentTerms field if non-nil, zero value otherwise.

### GetPaymentTermsOk

`func (o *ModelOrderResponse) GetPaymentTermsOk() (*string, bool)`

GetPaymentTermsOk returns a tuple with the PaymentTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTerms

`func (o *ModelOrderResponse) SetPaymentTerms(v string)`

SetPaymentTerms sets PaymentTerms field to given value.

### HasPaymentTerms

`func (o *ModelOrderResponse) HasPaymentTerms() bool`

HasPaymentTerms returns a boolean if a field has been set.

### SetPaymentTermsNil

`func (o *ModelOrderResponse) SetPaymentTermsNil(b bool)`

 SetPaymentTermsNil sets the value for PaymentTerms to be an explicit nil

### UnsetPaymentTerms
`func (o *ModelOrderResponse) UnsetPaymentTerms()`

UnsetPaymentTerms ensures that no value is present for PaymentTerms, not even an explicit nil
### GetOrigin

`func (o *ModelOrderResponse) GetOrigin() ModelOrderResponseOrigin`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *ModelOrderResponse) GetOriginOk() (*ModelOrderResponseOrigin, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *ModelOrderResponse) SetOrigin(v ModelOrderResponseOrigin)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *ModelOrderResponse) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### SetOriginNil

`func (o *ModelOrderResponse) SetOriginNil(b bool)`

 SetOriginNil sets the value for Origin to be an explicit nil

### UnsetOrigin
`func (o *ModelOrderResponse) UnsetOrigin()`

UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
### GetVersion

`func (o *ModelOrderResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ModelOrderResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ModelOrderResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ModelOrderResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSmallSettlement

`func (o *ModelOrderResponse) GetSmallSettlement() bool`

GetSmallSettlement returns the SmallSettlement field if non-nil, zero value otherwise.

### GetSmallSettlementOk

`func (o *ModelOrderResponse) GetSmallSettlementOk() (*bool, bool)`

GetSmallSettlementOk returns a tuple with the SmallSettlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallSettlement

`func (o *ModelOrderResponse) SetSmallSettlement(v bool)`

SetSmallSettlement sets SmallSettlement field to given value.

### HasSmallSettlement

`func (o *ModelOrderResponse) HasSmallSettlement() bool`

HasSmallSettlement returns a boolean if a field has been set.

### GetContactPerson

`func (o *ModelOrderResponse) GetContactPerson() ModelOrderResponseContactPerson`

GetContactPerson returns the ContactPerson field if non-nil, zero value otherwise.

### GetContactPersonOk

`func (o *ModelOrderResponse) GetContactPersonOk() (*ModelOrderResponseContactPerson, bool)`

GetContactPersonOk returns a tuple with the ContactPerson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPerson

`func (o *ModelOrderResponse) SetContactPerson(v ModelOrderResponseContactPerson)`

SetContactPerson sets ContactPerson field to given value.

### HasContactPerson

`func (o *ModelOrderResponse) HasContactPerson() bool`

HasContactPerson returns a boolean if a field has been set.

### GetTaxRate

`func (o *ModelOrderResponse) GetTaxRate() string`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *ModelOrderResponse) GetTaxRateOk() (*string, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *ModelOrderResponse) SetTaxRate(v string)`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *ModelOrderResponse) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### GetTaxSet

`func (o *ModelOrderResponse) GetTaxSet() ModelOrderResponseTaxSet`

GetTaxSet returns the TaxSet field if non-nil, zero value otherwise.

### GetTaxSetOk

`func (o *ModelOrderResponse) GetTaxSetOk() (*ModelOrderResponseTaxSet, bool)`

GetTaxSetOk returns a tuple with the TaxSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxSet

`func (o *ModelOrderResponse) SetTaxSet(v ModelOrderResponseTaxSet)`

SetTaxSet sets TaxSet field to given value.

### HasTaxSet

`func (o *ModelOrderResponse) HasTaxSet() bool`

HasTaxSet returns a boolean if a field has been set.

### SetTaxSetNil

`func (o *ModelOrderResponse) SetTaxSetNil(b bool)`

 SetTaxSetNil sets the value for TaxSet to be an explicit nil

### UnsetTaxSet
`func (o *ModelOrderResponse) UnsetTaxSet()`

UnsetTaxSet ensures that no value is present for TaxSet, not even an explicit nil
### GetTaxText

`func (o *ModelOrderResponse) GetTaxText() string`

GetTaxText returns the TaxText field if non-nil, zero value otherwise.

### GetTaxTextOk

`func (o *ModelOrderResponse) GetTaxTextOk() (*string, bool)`

GetTaxTextOk returns a tuple with the TaxText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxText

`func (o *ModelOrderResponse) SetTaxText(v string)`

SetTaxText sets TaxText field to given value.

### HasTaxText

`func (o *ModelOrderResponse) HasTaxText() bool`

HasTaxText returns a boolean if a field has been set.

### GetTaxType

`func (o *ModelOrderResponse) GetTaxType() string`

GetTaxType returns the TaxType field if non-nil, zero value otherwise.

### GetTaxTypeOk

`func (o *ModelOrderResponse) GetTaxTypeOk() (*string, bool)`

GetTaxTypeOk returns a tuple with the TaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxType

`func (o *ModelOrderResponse) SetTaxType(v string)`

SetTaxType sets TaxType field to given value.

### HasTaxType

`func (o *ModelOrderResponse) HasTaxType() bool`

HasTaxType returns a boolean if a field has been set.

### GetOrderType

`func (o *ModelOrderResponse) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *ModelOrderResponse) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *ModelOrderResponse) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *ModelOrderResponse) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetSendDate

`func (o *ModelOrderResponse) GetSendDate() time.Time`

GetSendDate returns the SendDate field if non-nil, zero value otherwise.

### GetSendDateOk

`func (o *ModelOrderResponse) GetSendDateOk() (*time.Time, bool)`

GetSendDateOk returns a tuple with the SendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendDate

`func (o *ModelOrderResponse) SetSendDate(v time.Time)`

SetSendDate sets SendDate field to given value.

### HasSendDate

`func (o *ModelOrderResponse) HasSendDate() bool`

HasSendDate returns a boolean if a field has been set.

### SetSendDateNil

`func (o *ModelOrderResponse) SetSendDateNil(b bool)`

 SetSendDateNil sets the value for SendDate to be an explicit nil

### UnsetSendDate
`func (o *ModelOrderResponse) UnsetSendDate()`

UnsetSendDate ensures that no value is present for SendDate, not even an explicit nil
### GetAddress

`func (o *ModelOrderResponse) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ModelOrderResponse) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ModelOrderResponse) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ModelOrderResponse) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *ModelOrderResponse) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *ModelOrderResponse) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetCurrency

`func (o *ModelOrderResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ModelOrderResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ModelOrderResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ModelOrderResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetSumNet

`func (o *ModelOrderResponse) GetSumNet() string`

GetSumNet returns the SumNet field if non-nil, zero value otherwise.

### GetSumNetOk

`func (o *ModelOrderResponse) GetSumNetOk() (*string, bool)`

GetSumNetOk returns a tuple with the SumNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumNet

`func (o *ModelOrderResponse) SetSumNet(v string)`

SetSumNet sets SumNet field to given value.

### HasSumNet

`func (o *ModelOrderResponse) HasSumNet() bool`

HasSumNet returns a boolean if a field has been set.

### GetSumTax

`func (o *ModelOrderResponse) GetSumTax() string`

GetSumTax returns the SumTax field if non-nil, zero value otherwise.

### GetSumTaxOk

`func (o *ModelOrderResponse) GetSumTaxOk() (*string, bool)`

GetSumTaxOk returns a tuple with the SumTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumTax

`func (o *ModelOrderResponse) SetSumTax(v string)`

SetSumTax sets SumTax field to given value.

### HasSumTax

`func (o *ModelOrderResponse) HasSumTax() bool`

HasSumTax returns a boolean if a field has been set.

### GetSumGross

`func (o *ModelOrderResponse) GetSumGross() string`

GetSumGross returns the SumGross field if non-nil, zero value otherwise.

### GetSumGrossOk

`func (o *ModelOrderResponse) GetSumGrossOk() (*string, bool)`

GetSumGrossOk returns a tuple with the SumGross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumGross

`func (o *ModelOrderResponse) SetSumGross(v string)`

SetSumGross sets SumGross field to given value.

### HasSumGross

`func (o *ModelOrderResponse) HasSumGross() bool`

HasSumGross returns a boolean if a field has been set.

### GetSumDiscounts

`func (o *ModelOrderResponse) GetSumDiscounts() string`

GetSumDiscounts returns the SumDiscounts field if non-nil, zero value otherwise.

### GetSumDiscountsOk

`func (o *ModelOrderResponse) GetSumDiscountsOk() (*string, bool)`

GetSumDiscountsOk returns a tuple with the SumDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumDiscounts

`func (o *ModelOrderResponse) SetSumDiscounts(v string)`

SetSumDiscounts sets SumDiscounts field to given value.

### HasSumDiscounts

`func (o *ModelOrderResponse) HasSumDiscounts() bool`

HasSumDiscounts returns a boolean if a field has been set.

### GetSumNetForeignCurrency

`func (o *ModelOrderResponse) GetSumNetForeignCurrency() string`

GetSumNetForeignCurrency returns the SumNetForeignCurrency field if non-nil, zero value otherwise.

### GetSumNetForeignCurrencyOk

`func (o *ModelOrderResponse) GetSumNetForeignCurrencyOk() (*string, bool)`

GetSumNetForeignCurrencyOk returns a tuple with the SumNetForeignCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumNetForeignCurrency

`func (o *ModelOrderResponse) SetSumNetForeignCurrency(v string)`

SetSumNetForeignCurrency sets SumNetForeignCurrency field to given value.

### HasSumNetForeignCurrency

`func (o *ModelOrderResponse) HasSumNetForeignCurrency() bool`

HasSumNetForeignCurrency returns a boolean if a field has been set.

### GetSumTaxForeignCurrency

`func (o *ModelOrderResponse) GetSumTaxForeignCurrency() string`

GetSumTaxForeignCurrency returns the SumTaxForeignCurrency field if non-nil, zero value otherwise.

### GetSumTaxForeignCurrencyOk

`func (o *ModelOrderResponse) GetSumTaxForeignCurrencyOk() (*string, bool)`

GetSumTaxForeignCurrencyOk returns a tuple with the SumTaxForeignCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumTaxForeignCurrency

`func (o *ModelOrderResponse) SetSumTaxForeignCurrency(v string)`

SetSumTaxForeignCurrency sets SumTaxForeignCurrency field to given value.

### HasSumTaxForeignCurrency

`func (o *ModelOrderResponse) HasSumTaxForeignCurrency() bool`

HasSumTaxForeignCurrency returns a boolean if a field has been set.

### GetSumGrossForeignCurrency

`func (o *ModelOrderResponse) GetSumGrossForeignCurrency() string`

GetSumGrossForeignCurrency returns the SumGrossForeignCurrency field if non-nil, zero value otherwise.

### GetSumGrossForeignCurrencyOk

`func (o *ModelOrderResponse) GetSumGrossForeignCurrencyOk() (*string, bool)`

GetSumGrossForeignCurrencyOk returns a tuple with the SumGrossForeignCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumGrossForeignCurrency

`func (o *ModelOrderResponse) SetSumGrossForeignCurrency(v string)`

SetSumGrossForeignCurrency sets SumGrossForeignCurrency field to given value.

### HasSumGrossForeignCurrency

`func (o *ModelOrderResponse) HasSumGrossForeignCurrency() bool`

HasSumGrossForeignCurrency returns a boolean if a field has been set.

### GetSumDiscountsForeignCurrency

`func (o *ModelOrderResponse) GetSumDiscountsForeignCurrency() string`

GetSumDiscountsForeignCurrency returns the SumDiscountsForeignCurrency field if non-nil, zero value otherwise.

### GetSumDiscountsForeignCurrencyOk

`func (o *ModelOrderResponse) GetSumDiscountsForeignCurrencyOk() (*string, bool)`

GetSumDiscountsForeignCurrencyOk returns a tuple with the SumDiscountsForeignCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumDiscountsForeignCurrency

`func (o *ModelOrderResponse) SetSumDiscountsForeignCurrency(v string)`

SetSumDiscountsForeignCurrency sets SumDiscountsForeignCurrency field to given value.

### HasSumDiscountsForeignCurrency

`func (o *ModelOrderResponse) HasSumDiscountsForeignCurrency() bool`

HasSumDiscountsForeignCurrency returns a boolean if a field has been set.

### GetCustomerInternalNote

`func (o *ModelOrderResponse) GetCustomerInternalNote() string`

GetCustomerInternalNote returns the CustomerInternalNote field if non-nil, zero value otherwise.

### GetCustomerInternalNoteOk

`func (o *ModelOrderResponse) GetCustomerInternalNoteOk() (*string, bool)`

GetCustomerInternalNoteOk returns a tuple with the CustomerInternalNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerInternalNote

`func (o *ModelOrderResponse) SetCustomerInternalNote(v string)`

SetCustomerInternalNote sets CustomerInternalNote field to given value.

### HasCustomerInternalNote

`func (o *ModelOrderResponse) HasCustomerInternalNote() bool`

HasCustomerInternalNote returns a boolean if a field has been set.

### SetCustomerInternalNoteNil

`func (o *ModelOrderResponse) SetCustomerInternalNoteNil(b bool)`

 SetCustomerInternalNoteNil sets the value for CustomerInternalNote to be an explicit nil

### UnsetCustomerInternalNote
`func (o *ModelOrderResponse) UnsetCustomerInternalNote()`

UnsetCustomerInternalNote ensures that no value is present for CustomerInternalNote, not even an explicit nil
### GetShowNet

`func (o *ModelOrderResponse) GetShowNet() bool`

GetShowNet returns the ShowNet field if non-nil, zero value otherwise.

### GetShowNetOk

`func (o *ModelOrderResponse) GetShowNetOk() (*bool, bool)`

GetShowNetOk returns a tuple with the ShowNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowNet

`func (o *ModelOrderResponse) SetShowNet(v bool)`

SetShowNet sets ShowNet field to given value.

### HasShowNet

`func (o *ModelOrderResponse) HasShowNet() bool`

HasShowNet returns a boolean if a field has been set.

### GetSendType

`func (o *ModelOrderResponse) GetSendType() string`

GetSendType returns the SendType field if non-nil, zero value otherwise.

### GetSendTypeOk

`func (o *ModelOrderResponse) GetSendTypeOk() (*string, bool)`

GetSendTypeOk returns a tuple with the SendType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendType

`func (o *ModelOrderResponse) SetSendType(v string)`

SetSendType sets SendType field to given value.

### HasSendType

`func (o *ModelOrderResponse) HasSendType() bool`

HasSendType returns a boolean if a field has been set.

### SetSendTypeNil

`func (o *ModelOrderResponse) SetSendTypeNil(b bool)`

 SetSendTypeNil sets the value for SendType to be an explicit nil

### UnsetSendType
`func (o *ModelOrderResponse) UnsetSendType()`

UnsetSendType ensures that no value is present for SendType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


