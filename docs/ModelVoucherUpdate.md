# ModelVoucherUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VoucherDate** | Pointer to **NullableTime** | Needs to be provided as timestamp or dd.mm.yyyy | [optional] 
**Supplier** | Pointer to [**NullableModelVoucherUpdateSupplier**](ModelVoucherUpdateSupplier.md) |  | [optional] 
**SupplierName** | Pointer to **NullableString** | The supplier name.&lt;br&gt;       The value you provide here will determine what supplier name is shown for the voucher in case you did not provide a supplier. | [optional] 
**Description** | Pointer to **NullableString** | The description of the voucher. Essentially the voucher number. | [optional] 
**PayDate** | Pointer to **NullableTime** | Needs to be timestamp or dd.mm.yyyy | [optional] 
**Status** | Pointer to **float32** | Please have a look in       &lt;a href&#x3D;&#39;https://api.sevdesk.de/#section/Types-and-status-of-vouchers&#39;&gt;status of vouchers&lt;/a&gt;      to see what the different status codes mean | [optional] 
**PaidAmount** | Pointer to **NullableFloat32** | Amount which has already been paid for this voucher by the customer | [optional] [readonly] 
**TaxType** | Pointer to **string** | Tax type of the voucher. There are four tax types: 1. default - Umsatzsteuer ausweisen 2. eu - Steuerfreie innergemeinschaftliche Lieferung (Europäische Union) 3. noteu - Steuerschuldnerschaft des Leistungsempfängers (außerhalb EU, z. B. Schweiz) 4. custom - Using custom tax set 5. ss - Not subject to VAT according to §19 1 UStG Tax rates are heavily connected to the tax type used. | [optional] 
**CreditDebit** | Pointer to **string** | Defines if your voucher is a credit (C) or debit (D) | [optional] 
**VoucherType** | Pointer to **string** | Type of the voucher. For more information on the different types, check       &lt;a href&#x3D;&#39;https://api.sevdesk.de/#section/Types-and-status-of-vouchers&#39;&gt;this&lt;/a&gt;   | [optional] 
**Currency** | Pointer to **NullableString** | specifies which currency the voucher should have. Attention: If the currency differs from the default currency stored in the account, then either the \&quot;propertyForeignCurrencyDeadline\&quot; or \&quot;propertyExchangeRate\&quot; parameter must be specified. If both parameters are specified, then the \&quot;propertyForeignCurrencyDeadline\&quot; parameter is preferred | [optional] 
**PropertyForeignCurrencyDeadline** | Pointer to **NullableTime** | Defines the exchange rate day and and then the exchange rate is set from sevDesk. Needs to be provided as timestamp or dd.mm.yyyy | [optional] 
**PropertyExchangeRate** | Pointer to **NullableFloat32** | Defines the exchange rate | [optional] 
**TaxSet** | Pointer to [**NullableModelVoucherUpdateTaxSet**](ModelVoucherUpdateTaxSet.md) |  | [optional] 
**PaymentDeadline** | Pointer to **NullableTime** | Payment deadline of the voucher. | [optional] 
**DeliveryDate** | Pointer to **time.Time** | Needs to be provided as timestamp or dd.mm.yyyy | [optional] 
**DeliveryDateUntil** | Pointer to **NullableTime** | Needs to be provided as timestamp or dd.mm.yyyy | [optional] 
**Document** | Pointer to [**NullableModelVoucherUpdateDocument**](ModelVoucherUpdateDocument.md) |  | [optional] 
**CostCentre** | Pointer to [**ModelVoucherUpdateCostCentre**](ModelVoucherUpdateCostCentre.md) |  | [optional] 

## Methods

### NewModelVoucherUpdate

`func NewModelVoucherUpdate() *ModelVoucherUpdate`

NewModelVoucherUpdate instantiates a new ModelVoucherUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelVoucherUpdateWithDefaults

`func NewModelVoucherUpdateWithDefaults() *ModelVoucherUpdate`

NewModelVoucherUpdateWithDefaults instantiates a new ModelVoucherUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVoucherDate

`func (o *ModelVoucherUpdate) GetVoucherDate() time.Time`

GetVoucherDate returns the VoucherDate field if non-nil, zero value otherwise.

### GetVoucherDateOk

`func (o *ModelVoucherUpdate) GetVoucherDateOk() (*time.Time, bool)`

GetVoucherDateOk returns a tuple with the VoucherDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherDate

`func (o *ModelVoucherUpdate) SetVoucherDate(v time.Time)`

SetVoucherDate sets VoucherDate field to given value.

### HasVoucherDate

`func (o *ModelVoucherUpdate) HasVoucherDate() bool`

HasVoucherDate returns a boolean if a field has been set.

### SetVoucherDateNil

`func (o *ModelVoucherUpdate) SetVoucherDateNil(b bool)`

 SetVoucherDateNil sets the value for VoucherDate to be an explicit nil

### UnsetVoucherDate
`func (o *ModelVoucherUpdate) UnsetVoucherDate()`

UnsetVoucherDate ensures that no value is present for VoucherDate, not even an explicit nil
### GetSupplier

`func (o *ModelVoucherUpdate) GetSupplier() ModelVoucherUpdateSupplier`

GetSupplier returns the Supplier field if non-nil, zero value otherwise.

### GetSupplierOk

`func (o *ModelVoucherUpdate) GetSupplierOk() (*ModelVoucherUpdateSupplier, bool)`

GetSupplierOk returns a tuple with the Supplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplier

`func (o *ModelVoucherUpdate) SetSupplier(v ModelVoucherUpdateSupplier)`

SetSupplier sets Supplier field to given value.

### HasSupplier

`func (o *ModelVoucherUpdate) HasSupplier() bool`

HasSupplier returns a boolean if a field has been set.

### SetSupplierNil

`func (o *ModelVoucherUpdate) SetSupplierNil(b bool)`

 SetSupplierNil sets the value for Supplier to be an explicit nil

### UnsetSupplier
`func (o *ModelVoucherUpdate) UnsetSupplier()`

UnsetSupplier ensures that no value is present for Supplier, not even an explicit nil
### GetSupplierName

`func (o *ModelVoucherUpdate) GetSupplierName() string`

GetSupplierName returns the SupplierName field if non-nil, zero value otherwise.

### GetSupplierNameOk

`func (o *ModelVoucherUpdate) GetSupplierNameOk() (*string, bool)`

GetSupplierNameOk returns a tuple with the SupplierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierName

`func (o *ModelVoucherUpdate) SetSupplierName(v string)`

SetSupplierName sets SupplierName field to given value.

### HasSupplierName

`func (o *ModelVoucherUpdate) HasSupplierName() bool`

HasSupplierName returns a boolean if a field has been set.

### SetSupplierNameNil

`func (o *ModelVoucherUpdate) SetSupplierNameNil(b bool)`

 SetSupplierNameNil sets the value for SupplierName to be an explicit nil

### UnsetSupplierName
`func (o *ModelVoucherUpdate) UnsetSupplierName()`

UnsetSupplierName ensures that no value is present for SupplierName, not even an explicit nil
### GetDescription

`func (o *ModelVoucherUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelVoucherUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelVoucherUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelVoucherUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ModelVoucherUpdate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ModelVoucherUpdate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPayDate

`func (o *ModelVoucherUpdate) GetPayDate() time.Time`

GetPayDate returns the PayDate field if non-nil, zero value otherwise.

### GetPayDateOk

`func (o *ModelVoucherUpdate) GetPayDateOk() (*time.Time, bool)`

GetPayDateOk returns a tuple with the PayDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayDate

`func (o *ModelVoucherUpdate) SetPayDate(v time.Time)`

SetPayDate sets PayDate field to given value.

### HasPayDate

`func (o *ModelVoucherUpdate) HasPayDate() bool`

HasPayDate returns a boolean if a field has been set.

### SetPayDateNil

`func (o *ModelVoucherUpdate) SetPayDateNil(b bool)`

 SetPayDateNil sets the value for PayDate to be an explicit nil

### UnsetPayDate
`func (o *ModelVoucherUpdate) UnsetPayDate()`

UnsetPayDate ensures that no value is present for PayDate, not even an explicit nil
### GetStatus

`func (o *ModelVoucherUpdate) GetStatus() float32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelVoucherUpdate) GetStatusOk() (*float32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelVoucherUpdate) SetStatus(v float32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelVoucherUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPaidAmount

`func (o *ModelVoucherUpdate) GetPaidAmount() float32`

GetPaidAmount returns the PaidAmount field if non-nil, zero value otherwise.

### GetPaidAmountOk

`func (o *ModelVoucherUpdate) GetPaidAmountOk() (*float32, bool)`

GetPaidAmountOk returns a tuple with the PaidAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidAmount

`func (o *ModelVoucherUpdate) SetPaidAmount(v float32)`

SetPaidAmount sets PaidAmount field to given value.

### HasPaidAmount

`func (o *ModelVoucherUpdate) HasPaidAmount() bool`

HasPaidAmount returns a boolean if a field has been set.

### SetPaidAmountNil

`func (o *ModelVoucherUpdate) SetPaidAmountNil(b bool)`

 SetPaidAmountNil sets the value for PaidAmount to be an explicit nil

### UnsetPaidAmount
`func (o *ModelVoucherUpdate) UnsetPaidAmount()`

UnsetPaidAmount ensures that no value is present for PaidAmount, not even an explicit nil
### GetTaxType

`func (o *ModelVoucherUpdate) GetTaxType() string`

GetTaxType returns the TaxType field if non-nil, zero value otherwise.

### GetTaxTypeOk

`func (o *ModelVoucherUpdate) GetTaxTypeOk() (*string, bool)`

GetTaxTypeOk returns a tuple with the TaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxType

`func (o *ModelVoucherUpdate) SetTaxType(v string)`

SetTaxType sets TaxType field to given value.

### HasTaxType

`func (o *ModelVoucherUpdate) HasTaxType() bool`

HasTaxType returns a boolean if a field has been set.

### GetCreditDebit

`func (o *ModelVoucherUpdate) GetCreditDebit() string`

GetCreditDebit returns the CreditDebit field if non-nil, zero value otherwise.

### GetCreditDebitOk

`func (o *ModelVoucherUpdate) GetCreditDebitOk() (*string, bool)`

GetCreditDebitOk returns a tuple with the CreditDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditDebit

`func (o *ModelVoucherUpdate) SetCreditDebit(v string)`

SetCreditDebit sets CreditDebit field to given value.

### HasCreditDebit

`func (o *ModelVoucherUpdate) HasCreditDebit() bool`

HasCreditDebit returns a boolean if a field has been set.

### GetVoucherType

`func (o *ModelVoucherUpdate) GetVoucherType() string`

GetVoucherType returns the VoucherType field if non-nil, zero value otherwise.

### GetVoucherTypeOk

`func (o *ModelVoucherUpdate) GetVoucherTypeOk() (*string, bool)`

GetVoucherTypeOk returns a tuple with the VoucherType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherType

`func (o *ModelVoucherUpdate) SetVoucherType(v string)`

SetVoucherType sets VoucherType field to given value.

### HasVoucherType

`func (o *ModelVoucherUpdate) HasVoucherType() bool`

HasVoucherType returns a boolean if a field has been set.

### GetCurrency

`func (o *ModelVoucherUpdate) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ModelVoucherUpdate) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ModelVoucherUpdate) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ModelVoucherUpdate) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *ModelVoucherUpdate) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *ModelVoucherUpdate) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetPropertyForeignCurrencyDeadline

`func (o *ModelVoucherUpdate) GetPropertyForeignCurrencyDeadline() time.Time`

GetPropertyForeignCurrencyDeadline returns the PropertyForeignCurrencyDeadline field if non-nil, zero value otherwise.

### GetPropertyForeignCurrencyDeadlineOk

`func (o *ModelVoucherUpdate) GetPropertyForeignCurrencyDeadlineOk() (*time.Time, bool)`

GetPropertyForeignCurrencyDeadlineOk returns a tuple with the PropertyForeignCurrencyDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyForeignCurrencyDeadline

`func (o *ModelVoucherUpdate) SetPropertyForeignCurrencyDeadline(v time.Time)`

SetPropertyForeignCurrencyDeadline sets PropertyForeignCurrencyDeadline field to given value.

### HasPropertyForeignCurrencyDeadline

`func (o *ModelVoucherUpdate) HasPropertyForeignCurrencyDeadline() bool`

HasPropertyForeignCurrencyDeadline returns a boolean if a field has been set.

### SetPropertyForeignCurrencyDeadlineNil

`func (o *ModelVoucherUpdate) SetPropertyForeignCurrencyDeadlineNil(b bool)`

 SetPropertyForeignCurrencyDeadlineNil sets the value for PropertyForeignCurrencyDeadline to be an explicit nil

### UnsetPropertyForeignCurrencyDeadline
`func (o *ModelVoucherUpdate) UnsetPropertyForeignCurrencyDeadline()`

UnsetPropertyForeignCurrencyDeadline ensures that no value is present for PropertyForeignCurrencyDeadline, not even an explicit nil
### GetPropertyExchangeRate

`func (o *ModelVoucherUpdate) GetPropertyExchangeRate() float32`

GetPropertyExchangeRate returns the PropertyExchangeRate field if non-nil, zero value otherwise.

### GetPropertyExchangeRateOk

`func (o *ModelVoucherUpdate) GetPropertyExchangeRateOk() (*float32, bool)`

GetPropertyExchangeRateOk returns a tuple with the PropertyExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyExchangeRate

`func (o *ModelVoucherUpdate) SetPropertyExchangeRate(v float32)`

SetPropertyExchangeRate sets PropertyExchangeRate field to given value.

### HasPropertyExchangeRate

`func (o *ModelVoucherUpdate) HasPropertyExchangeRate() bool`

HasPropertyExchangeRate returns a boolean if a field has been set.

### SetPropertyExchangeRateNil

`func (o *ModelVoucherUpdate) SetPropertyExchangeRateNil(b bool)`

 SetPropertyExchangeRateNil sets the value for PropertyExchangeRate to be an explicit nil

### UnsetPropertyExchangeRate
`func (o *ModelVoucherUpdate) UnsetPropertyExchangeRate()`

UnsetPropertyExchangeRate ensures that no value is present for PropertyExchangeRate, not even an explicit nil
### GetTaxSet

`func (o *ModelVoucherUpdate) GetTaxSet() ModelVoucherUpdateTaxSet`

GetTaxSet returns the TaxSet field if non-nil, zero value otherwise.

### GetTaxSetOk

`func (o *ModelVoucherUpdate) GetTaxSetOk() (*ModelVoucherUpdateTaxSet, bool)`

GetTaxSetOk returns a tuple with the TaxSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxSet

`func (o *ModelVoucherUpdate) SetTaxSet(v ModelVoucherUpdateTaxSet)`

SetTaxSet sets TaxSet field to given value.

### HasTaxSet

`func (o *ModelVoucherUpdate) HasTaxSet() bool`

HasTaxSet returns a boolean if a field has been set.

### SetTaxSetNil

`func (o *ModelVoucherUpdate) SetTaxSetNil(b bool)`

 SetTaxSetNil sets the value for TaxSet to be an explicit nil

### UnsetTaxSet
`func (o *ModelVoucherUpdate) UnsetTaxSet()`

UnsetTaxSet ensures that no value is present for TaxSet, not even an explicit nil
### GetPaymentDeadline

`func (o *ModelVoucherUpdate) GetPaymentDeadline() time.Time`

GetPaymentDeadline returns the PaymentDeadline field if non-nil, zero value otherwise.

### GetPaymentDeadlineOk

`func (o *ModelVoucherUpdate) GetPaymentDeadlineOk() (*time.Time, bool)`

GetPaymentDeadlineOk returns a tuple with the PaymentDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDeadline

`func (o *ModelVoucherUpdate) SetPaymentDeadline(v time.Time)`

SetPaymentDeadline sets PaymentDeadline field to given value.

### HasPaymentDeadline

`func (o *ModelVoucherUpdate) HasPaymentDeadline() bool`

HasPaymentDeadline returns a boolean if a field has been set.

### SetPaymentDeadlineNil

`func (o *ModelVoucherUpdate) SetPaymentDeadlineNil(b bool)`

 SetPaymentDeadlineNil sets the value for PaymentDeadline to be an explicit nil

### UnsetPaymentDeadline
`func (o *ModelVoucherUpdate) UnsetPaymentDeadline()`

UnsetPaymentDeadline ensures that no value is present for PaymentDeadline, not even an explicit nil
### GetDeliveryDate

`func (o *ModelVoucherUpdate) GetDeliveryDate() time.Time`

GetDeliveryDate returns the DeliveryDate field if non-nil, zero value otherwise.

### GetDeliveryDateOk

`func (o *ModelVoucherUpdate) GetDeliveryDateOk() (*time.Time, bool)`

GetDeliveryDateOk returns a tuple with the DeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDate

`func (o *ModelVoucherUpdate) SetDeliveryDate(v time.Time)`

SetDeliveryDate sets DeliveryDate field to given value.

### HasDeliveryDate

`func (o *ModelVoucherUpdate) HasDeliveryDate() bool`

HasDeliveryDate returns a boolean if a field has been set.

### GetDeliveryDateUntil

`func (o *ModelVoucherUpdate) GetDeliveryDateUntil() time.Time`

GetDeliveryDateUntil returns the DeliveryDateUntil field if non-nil, zero value otherwise.

### GetDeliveryDateUntilOk

`func (o *ModelVoucherUpdate) GetDeliveryDateUntilOk() (*time.Time, bool)`

GetDeliveryDateUntilOk returns a tuple with the DeliveryDateUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDateUntil

`func (o *ModelVoucherUpdate) SetDeliveryDateUntil(v time.Time)`

SetDeliveryDateUntil sets DeliveryDateUntil field to given value.

### HasDeliveryDateUntil

`func (o *ModelVoucherUpdate) HasDeliveryDateUntil() bool`

HasDeliveryDateUntil returns a boolean if a field has been set.

### SetDeliveryDateUntilNil

`func (o *ModelVoucherUpdate) SetDeliveryDateUntilNil(b bool)`

 SetDeliveryDateUntilNil sets the value for DeliveryDateUntil to be an explicit nil

### UnsetDeliveryDateUntil
`func (o *ModelVoucherUpdate) UnsetDeliveryDateUntil()`

UnsetDeliveryDateUntil ensures that no value is present for DeliveryDateUntil, not even an explicit nil
### GetDocument

`func (o *ModelVoucherUpdate) GetDocument() ModelVoucherUpdateDocument`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *ModelVoucherUpdate) GetDocumentOk() (*ModelVoucherUpdateDocument, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *ModelVoucherUpdate) SetDocument(v ModelVoucherUpdateDocument)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *ModelVoucherUpdate) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### SetDocumentNil

`func (o *ModelVoucherUpdate) SetDocumentNil(b bool)`

 SetDocumentNil sets the value for Document to be an explicit nil

### UnsetDocument
`func (o *ModelVoucherUpdate) UnsetDocument()`

UnsetDocument ensures that no value is present for Document, not even an explicit nil
### GetCostCentre

`func (o *ModelVoucherUpdate) GetCostCentre() ModelVoucherUpdateCostCentre`

GetCostCentre returns the CostCentre field if non-nil, zero value otherwise.

### GetCostCentreOk

`func (o *ModelVoucherUpdate) GetCostCentreOk() (*ModelVoucherUpdateCostCentre, bool)`

GetCostCentreOk returns a tuple with the CostCentre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCentre

`func (o *ModelVoucherUpdate) SetCostCentre(v ModelVoucherUpdateCostCentre)`

SetCostCentre sets CostCentre field to given value.

### HasCostCentre

`func (o *ModelVoucherUpdate) HasCostCentre() bool`

HasCostCentre returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


