# ModelVoucher

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The voucher id | [optional] [readonly] 
**ObjectName** | **string** | The voucher object name | 
**MapAll** | **bool** |  | 
**Create** | Pointer to **time.Time** | Date of voucher creation | [optional] [readonly] 
**Update** | Pointer to **time.Time** | Date of last voucher update | [optional] [readonly] 
**SevClient** | Pointer to [**ModelVoucherSevClient**](ModelVoucherSevClient.md) |  | [optional] 
**CreateUser** | Pointer to [**ModelVoucherCreateUser**](ModelVoucherCreateUser.md) |  | [optional] 
**VoucherDate** | Pointer to **NullableTime** | Needs to be provided as timestamp or dd.mm.yyyy | [optional] 
**Supplier** | Pointer to [**NullableModelVoucherSupplier**](ModelVoucherSupplier.md) |  | [optional] 
**SupplierName** | Pointer to **NullableString** | The supplier name.&lt;br&gt;       The value you provide here will determine what supplier name is shown for the voucher in case you did not provide a supplier. | [optional] 
**Description** | Pointer to **NullableString** | The description of the voucher. Essentially the voucher number. | [optional] 
**PayDate** | Pointer to **NullableTime** | Needs to be timestamp or dd.mm.yyyy | [optional] 
**Status** | **float32** | Please have a look in       &lt;a href&#x3D;&#39;https://api.sevdesk.de/#section/Types-and-status-of-vouchers&#39;&gt;status of vouchers&lt;/a&gt;      to see what the different status codes mean | 
**SumNet** | Pointer to **float32** | Net sum of the voucher | [optional] [readonly] 
**SumTax** | Pointer to **float32** | Tax sum of the voucher | [optional] [readonly] 
**SumGross** | Pointer to **float32** | Gross sum of the voucher | [optional] [readonly] 
**SumNetAccounting** | Pointer to **float32** | Net accounting sum of the voucher. Is usually the same as sumNet | [optional] [readonly] 
**SumTaxAccounting** | Pointer to **float32** | Tax accounting sum of the voucher. Is usually the same as sumTax | [optional] [readonly] 
**SumGrossAccounting** | Pointer to **float32** | Gross accounting sum of the voucher. Is usually the same as sumGross | [optional] [readonly] 
**SumDiscounts** | Pointer to **float32** | Sum of all discounts in the voucher | [optional] [readonly] 
**SumDiscountsForeignCurrency** | Pointer to **float32** | Discounts sum of the voucher in the foreign currency | [optional] [readonly] 
**PaidAmount** | Pointer to **NullableFloat32** | Amount which has already been paid for this voucher by the customer | [optional] [readonly] 
**TaxType** | **string** | Tax type of the voucher. There are four tax types: 1. default - Umsatzsteuer ausweisen 2. eu - Steuerfreie innergemeinschaftliche Lieferung (Europäische Union) 3. noteu - Steuerschuldnerschaft des Leistungsempfängers (außerhalb EU, z. B. Schweiz) 4. custom - Using custom tax set 5. ss - Not subject to VAT according to §19 1 UStG Tax rates are heavily connected to the tax type used. | 
**CreditDebit** | **string** | Defines if your voucher is a credit (C) or debit (D) | 
**VoucherType** | **string** | Type of the voucher. For more information on the different types, check       &lt;a href&#x3D;&#39;https://api.sevdesk.de/#section/Types-and-status-of-vouchers&#39;&gt;this&lt;/a&gt;   | 
**Currency** | Pointer to **NullableString** | specifies which currency the voucher should have. Attention: If the currency differs from the default currency stored in the account, then either the \&quot;propertyForeignCurrencyDeadline\&quot; or \&quot;propertyExchangeRate\&quot; parameter must be specified. If both parameters are specified, then the \&quot;propertyForeignCurrencyDeadline\&quot; parameter is preferred | [optional] 
**PropertyForeignCurrencyDeadline** | Pointer to **NullableTime** | Defines the exchange rate day and and then the exchange rate is set from sevDesk. Needs to be provided as timestamp or dd.mm.yyyy | [optional] 
**PropertyExchangeRate** | Pointer to **NullableFloat32** | Defines the exchange rate | [optional] 
**RecurringInterval** | Pointer to **NullableString** | The DateInterval in which recurring vouchers are generated.&lt;br&gt;       Necessary attribute for all recurring vouchers. | [optional] [readonly] 
**RecurringStartDate** | Pointer to **NullableTime** | The date when the recurring vouchers start being generated.&lt;br&gt;       Necessary attribute for all recurring vouchers. | [optional] [readonly] 
**RecurringNextVoucher** | Pointer to **NullableTime** | The date when the next voucher should be generated.&lt;br&gt;       Necessary attribute for all recurring vouchers. | [optional] [readonly] 
**RecurringLastVoucher** | Pointer to **NullableTime** | The date when the last voucher was generated. | [optional] [readonly] 
**RecurringEndDate** | Pointer to **NullableTime** | The date when the recurring vouchers end being generated.&lt;br&gt;      Necessary attribute for all recurring vouchers. | [optional] [readonly] 
**Enshrined** | Pointer to **NullableTime** | Defines if and when voucher was enshrined. Enshrined vouchers can not be manipulated. | [optional] [readonly] 
**TaxSet** | Pointer to [**NullableModelVoucherUpdateTaxSet**](ModelVoucherUpdateTaxSet.md) |  | [optional] 
**PaymentDeadline** | Pointer to **NullableTime** | Payment deadline of the voucher. | [optional] 
**DeliveryDate** | Pointer to **time.Time** | Needs to be provided as timestamp or dd.mm.yyyy | [optional] 
**DeliveryDateUntil** | Pointer to **NullableTime** | Needs to be provided as timestamp or dd.mm.yyyy | [optional] 
**Document** | Pointer to [**NullableModelVoucherUpdateDocument**](ModelVoucherUpdateDocument.md) |  | [optional] 
**CostCentre** | Pointer to [**ModelVoucherUpdateCostCentre**](ModelVoucherUpdateCostCentre.md) |  | [optional] 

## Methods

### NewModelVoucher

`func NewModelVoucher(objectName string, mapAll bool, status float32, taxType string, creditDebit string, voucherType string, ) *ModelVoucher`

NewModelVoucher instantiates a new ModelVoucher object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelVoucherWithDefaults

`func NewModelVoucherWithDefaults() *ModelVoucher`

NewModelVoucherWithDefaults instantiates a new ModelVoucher object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelVoucher) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelVoucher) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelVoucher) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelVoucher) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectName

`func (o *ModelVoucher) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ModelVoucher) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ModelVoucher) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.


### GetMapAll

`func (o *ModelVoucher) GetMapAll() bool`

GetMapAll returns the MapAll field if non-nil, zero value otherwise.

### GetMapAllOk

`func (o *ModelVoucher) GetMapAllOk() (*bool, bool)`

GetMapAllOk returns a tuple with the MapAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapAll

`func (o *ModelVoucher) SetMapAll(v bool)`

SetMapAll sets MapAll field to given value.


### GetCreate

`func (o *ModelVoucher) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ModelVoucher) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ModelVoucher) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ModelVoucher) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUpdate

`func (o *ModelVoucher) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ModelVoucher) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ModelVoucher) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ModelVoucher) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetSevClient

`func (o *ModelVoucher) GetSevClient() ModelVoucherSevClient`

GetSevClient returns the SevClient field if non-nil, zero value otherwise.

### GetSevClientOk

`func (o *ModelVoucher) GetSevClientOk() (*ModelVoucherSevClient, bool)`

GetSevClientOk returns a tuple with the SevClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSevClient

`func (o *ModelVoucher) SetSevClient(v ModelVoucherSevClient)`

SetSevClient sets SevClient field to given value.

### HasSevClient

`func (o *ModelVoucher) HasSevClient() bool`

HasSevClient returns a boolean if a field has been set.

### GetCreateUser

`func (o *ModelVoucher) GetCreateUser() ModelVoucherCreateUser`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *ModelVoucher) GetCreateUserOk() (*ModelVoucherCreateUser, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *ModelVoucher) SetCreateUser(v ModelVoucherCreateUser)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *ModelVoucher) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetVoucherDate

`func (o *ModelVoucher) GetVoucherDate() time.Time`

GetVoucherDate returns the VoucherDate field if non-nil, zero value otherwise.

### GetVoucherDateOk

`func (o *ModelVoucher) GetVoucherDateOk() (*time.Time, bool)`

GetVoucherDateOk returns a tuple with the VoucherDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherDate

`func (o *ModelVoucher) SetVoucherDate(v time.Time)`

SetVoucherDate sets VoucherDate field to given value.

### HasVoucherDate

`func (o *ModelVoucher) HasVoucherDate() bool`

HasVoucherDate returns a boolean if a field has been set.

### SetVoucherDateNil

`func (o *ModelVoucher) SetVoucherDateNil(b bool)`

 SetVoucherDateNil sets the value for VoucherDate to be an explicit nil

### UnsetVoucherDate
`func (o *ModelVoucher) UnsetVoucherDate()`

UnsetVoucherDate ensures that no value is present for VoucherDate, not even an explicit nil
### GetSupplier

`func (o *ModelVoucher) GetSupplier() ModelVoucherSupplier`

GetSupplier returns the Supplier field if non-nil, zero value otherwise.

### GetSupplierOk

`func (o *ModelVoucher) GetSupplierOk() (*ModelVoucherSupplier, bool)`

GetSupplierOk returns a tuple with the Supplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplier

`func (o *ModelVoucher) SetSupplier(v ModelVoucherSupplier)`

SetSupplier sets Supplier field to given value.

### HasSupplier

`func (o *ModelVoucher) HasSupplier() bool`

HasSupplier returns a boolean if a field has been set.

### SetSupplierNil

`func (o *ModelVoucher) SetSupplierNil(b bool)`

 SetSupplierNil sets the value for Supplier to be an explicit nil

### UnsetSupplier
`func (o *ModelVoucher) UnsetSupplier()`

UnsetSupplier ensures that no value is present for Supplier, not even an explicit nil
### GetSupplierName

`func (o *ModelVoucher) GetSupplierName() string`

GetSupplierName returns the SupplierName field if non-nil, zero value otherwise.

### GetSupplierNameOk

`func (o *ModelVoucher) GetSupplierNameOk() (*string, bool)`

GetSupplierNameOk returns a tuple with the SupplierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierName

`func (o *ModelVoucher) SetSupplierName(v string)`

SetSupplierName sets SupplierName field to given value.

### HasSupplierName

`func (o *ModelVoucher) HasSupplierName() bool`

HasSupplierName returns a boolean if a field has been set.

### SetSupplierNameNil

`func (o *ModelVoucher) SetSupplierNameNil(b bool)`

 SetSupplierNameNil sets the value for SupplierName to be an explicit nil

### UnsetSupplierName
`func (o *ModelVoucher) UnsetSupplierName()`

UnsetSupplierName ensures that no value is present for SupplierName, not even an explicit nil
### GetDescription

`func (o *ModelVoucher) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelVoucher) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelVoucher) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelVoucher) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ModelVoucher) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ModelVoucher) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPayDate

`func (o *ModelVoucher) GetPayDate() time.Time`

GetPayDate returns the PayDate field if non-nil, zero value otherwise.

### GetPayDateOk

`func (o *ModelVoucher) GetPayDateOk() (*time.Time, bool)`

GetPayDateOk returns a tuple with the PayDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayDate

`func (o *ModelVoucher) SetPayDate(v time.Time)`

SetPayDate sets PayDate field to given value.

### HasPayDate

`func (o *ModelVoucher) HasPayDate() bool`

HasPayDate returns a boolean if a field has been set.

### SetPayDateNil

`func (o *ModelVoucher) SetPayDateNil(b bool)`

 SetPayDateNil sets the value for PayDate to be an explicit nil

### UnsetPayDate
`func (o *ModelVoucher) UnsetPayDate()`

UnsetPayDate ensures that no value is present for PayDate, not even an explicit nil
### GetStatus

`func (o *ModelVoucher) GetStatus() float32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelVoucher) GetStatusOk() (*float32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelVoucher) SetStatus(v float32)`

SetStatus sets Status field to given value.


### GetSumNet

`func (o *ModelVoucher) GetSumNet() float32`

GetSumNet returns the SumNet field if non-nil, zero value otherwise.

### GetSumNetOk

`func (o *ModelVoucher) GetSumNetOk() (*float32, bool)`

GetSumNetOk returns a tuple with the SumNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumNet

`func (o *ModelVoucher) SetSumNet(v float32)`

SetSumNet sets SumNet field to given value.

### HasSumNet

`func (o *ModelVoucher) HasSumNet() bool`

HasSumNet returns a boolean if a field has been set.

### GetSumTax

`func (o *ModelVoucher) GetSumTax() float32`

GetSumTax returns the SumTax field if non-nil, zero value otherwise.

### GetSumTaxOk

`func (o *ModelVoucher) GetSumTaxOk() (*float32, bool)`

GetSumTaxOk returns a tuple with the SumTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumTax

`func (o *ModelVoucher) SetSumTax(v float32)`

SetSumTax sets SumTax field to given value.

### HasSumTax

`func (o *ModelVoucher) HasSumTax() bool`

HasSumTax returns a boolean if a field has been set.

### GetSumGross

`func (o *ModelVoucher) GetSumGross() float32`

GetSumGross returns the SumGross field if non-nil, zero value otherwise.

### GetSumGrossOk

`func (o *ModelVoucher) GetSumGrossOk() (*float32, bool)`

GetSumGrossOk returns a tuple with the SumGross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumGross

`func (o *ModelVoucher) SetSumGross(v float32)`

SetSumGross sets SumGross field to given value.

### HasSumGross

`func (o *ModelVoucher) HasSumGross() bool`

HasSumGross returns a boolean if a field has been set.

### GetSumNetAccounting

`func (o *ModelVoucher) GetSumNetAccounting() float32`

GetSumNetAccounting returns the SumNetAccounting field if non-nil, zero value otherwise.

### GetSumNetAccountingOk

`func (o *ModelVoucher) GetSumNetAccountingOk() (*float32, bool)`

GetSumNetAccountingOk returns a tuple with the SumNetAccounting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumNetAccounting

`func (o *ModelVoucher) SetSumNetAccounting(v float32)`

SetSumNetAccounting sets SumNetAccounting field to given value.

### HasSumNetAccounting

`func (o *ModelVoucher) HasSumNetAccounting() bool`

HasSumNetAccounting returns a boolean if a field has been set.

### GetSumTaxAccounting

`func (o *ModelVoucher) GetSumTaxAccounting() float32`

GetSumTaxAccounting returns the SumTaxAccounting field if non-nil, zero value otherwise.

### GetSumTaxAccountingOk

`func (o *ModelVoucher) GetSumTaxAccountingOk() (*float32, bool)`

GetSumTaxAccountingOk returns a tuple with the SumTaxAccounting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumTaxAccounting

`func (o *ModelVoucher) SetSumTaxAccounting(v float32)`

SetSumTaxAccounting sets SumTaxAccounting field to given value.

### HasSumTaxAccounting

`func (o *ModelVoucher) HasSumTaxAccounting() bool`

HasSumTaxAccounting returns a boolean if a field has been set.

### GetSumGrossAccounting

`func (o *ModelVoucher) GetSumGrossAccounting() float32`

GetSumGrossAccounting returns the SumGrossAccounting field if non-nil, zero value otherwise.

### GetSumGrossAccountingOk

`func (o *ModelVoucher) GetSumGrossAccountingOk() (*float32, bool)`

GetSumGrossAccountingOk returns a tuple with the SumGrossAccounting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumGrossAccounting

`func (o *ModelVoucher) SetSumGrossAccounting(v float32)`

SetSumGrossAccounting sets SumGrossAccounting field to given value.

### HasSumGrossAccounting

`func (o *ModelVoucher) HasSumGrossAccounting() bool`

HasSumGrossAccounting returns a boolean if a field has been set.

### GetSumDiscounts

`func (o *ModelVoucher) GetSumDiscounts() float32`

GetSumDiscounts returns the SumDiscounts field if non-nil, zero value otherwise.

### GetSumDiscountsOk

`func (o *ModelVoucher) GetSumDiscountsOk() (*float32, bool)`

GetSumDiscountsOk returns a tuple with the SumDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumDiscounts

`func (o *ModelVoucher) SetSumDiscounts(v float32)`

SetSumDiscounts sets SumDiscounts field to given value.

### HasSumDiscounts

`func (o *ModelVoucher) HasSumDiscounts() bool`

HasSumDiscounts returns a boolean if a field has been set.

### GetSumDiscountsForeignCurrency

`func (o *ModelVoucher) GetSumDiscountsForeignCurrency() float32`

GetSumDiscountsForeignCurrency returns the SumDiscountsForeignCurrency field if non-nil, zero value otherwise.

### GetSumDiscountsForeignCurrencyOk

`func (o *ModelVoucher) GetSumDiscountsForeignCurrencyOk() (*float32, bool)`

GetSumDiscountsForeignCurrencyOk returns a tuple with the SumDiscountsForeignCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumDiscountsForeignCurrency

`func (o *ModelVoucher) SetSumDiscountsForeignCurrency(v float32)`

SetSumDiscountsForeignCurrency sets SumDiscountsForeignCurrency field to given value.

### HasSumDiscountsForeignCurrency

`func (o *ModelVoucher) HasSumDiscountsForeignCurrency() bool`

HasSumDiscountsForeignCurrency returns a boolean if a field has been set.

### GetPaidAmount

`func (o *ModelVoucher) GetPaidAmount() float32`

GetPaidAmount returns the PaidAmount field if non-nil, zero value otherwise.

### GetPaidAmountOk

`func (o *ModelVoucher) GetPaidAmountOk() (*float32, bool)`

GetPaidAmountOk returns a tuple with the PaidAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidAmount

`func (o *ModelVoucher) SetPaidAmount(v float32)`

SetPaidAmount sets PaidAmount field to given value.

### HasPaidAmount

`func (o *ModelVoucher) HasPaidAmount() bool`

HasPaidAmount returns a boolean if a field has been set.

### SetPaidAmountNil

`func (o *ModelVoucher) SetPaidAmountNil(b bool)`

 SetPaidAmountNil sets the value for PaidAmount to be an explicit nil

### UnsetPaidAmount
`func (o *ModelVoucher) UnsetPaidAmount()`

UnsetPaidAmount ensures that no value is present for PaidAmount, not even an explicit nil
### GetTaxType

`func (o *ModelVoucher) GetTaxType() string`

GetTaxType returns the TaxType field if non-nil, zero value otherwise.

### GetTaxTypeOk

`func (o *ModelVoucher) GetTaxTypeOk() (*string, bool)`

GetTaxTypeOk returns a tuple with the TaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxType

`func (o *ModelVoucher) SetTaxType(v string)`

SetTaxType sets TaxType field to given value.


### GetCreditDebit

`func (o *ModelVoucher) GetCreditDebit() string`

GetCreditDebit returns the CreditDebit field if non-nil, zero value otherwise.

### GetCreditDebitOk

`func (o *ModelVoucher) GetCreditDebitOk() (*string, bool)`

GetCreditDebitOk returns a tuple with the CreditDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditDebit

`func (o *ModelVoucher) SetCreditDebit(v string)`

SetCreditDebit sets CreditDebit field to given value.


### GetVoucherType

`func (o *ModelVoucher) GetVoucherType() string`

GetVoucherType returns the VoucherType field if non-nil, zero value otherwise.

### GetVoucherTypeOk

`func (o *ModelVoucher) GetVoucherTypeOk() (*string, bool)`

GetVoucherTypeOk returns a tuple with the VoucherType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherType

`func (o *ModelVoucher) SetVoucherType(v string)`

SetVoucherType sets VoucherType field to given value.


### GetCurrency

`func (o *ModelVoucher) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ModelVoucher) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ModelVoucher) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ModelVoucher) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *ModelVoucher) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *ModelVoucher) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetPropertyForeignCurrencyDeadline

`func (o *ModelVoucher) GetPropertyForeignCurrencyDeadline() time.Time`

GetPropertyForeignCurrencyDeadline returns the PropertyForeignCurrencyDeadline field if non-nil, zero value otherwise.

### GetPropertyForeignCurrencyDeadlineOk

`func (o *ModelVoucher) GetPropertyForeignCurrencyDeadlineOk() (*time.Time, bool)`

GetPropertyForeignCurrencyDeadlineOk returns a tuple with the PropertyForeignCurrencyDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyForeignCurrencyDeadline

`func (o *ModelVoucher) SetPropertyForeignCurrencyDeadline(v time.Time)`

SetPropertyForeignCurrencyDeadline sets PropertyForeignCurrencyDeadline field to given value.

### HasPropertyForeignCurrencyDeadline

`func (o *ModelVoucher) HasPropertyForeignCurrencyDeadline() bool`

HasPropertyForeignCurrencyDeadline returns a boolean if a field has been set.

### SetPropertyForeignCurrencyDeadlineNil

`func (o *ModelVoucher) SetPropertyForeignCurrencyDeadlineNil(b bool)`

 SetPropertyForeignCurrencyDeadlineNil sets the value for PropertyForeignCurrencyDeadline to be an explicit nil

### UnsetPropertyForeignCurrencyDeadline
`func (o *ModelVoucher) UnsetPropertyForeignCurrencyDeadline()`

UnsetPropertyForeignCurrencyDeadline ensures that no value is present for PropertyForeignCurrencyDeadline, not even an explicit nil
### GetPropertyExchangeRate

`func (o *ModelVoucher) GetPropertyExchangeRate() float32`

GetPropertyExchangeRate returns the PropertyExchangeRate field if non-nil, zero value otherwise.

### GetPropertyExchangeRateOk

`func (o *ModelVoucher) GetPropertyExchangeRateOk() (*float32, bool)`

GetPropertyExchangeRateOk returns a tuple with the PropertyExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyExchangeRate

`func (o *ModelVoucher) SetPropertyExchangeRate(v float32)`

SetPropertyExchangeRate sets PropertyExchangeRate field to given value.

### HasPropertyExchangeRate

`func (o *ModelVoucher) HasPropertyExchangeRate() bool`

HasPropertyExchangeRate returns a boolean if a field has been set.

### SetPropertyExchangeRateNil

`func (o *ModelVoucher) SetPropertyExchangeRateNil(b bool)`

 SetPropertyExchangeRateNil sets the value for PropertyExchangeRate to be an explicit nil

### UnsetPropertyExchangeRate
`func (o *ModelVoucher) UnsetPropertyExchangeRate()`

UnsetPropertyExchangeRate ensures that no value is present for PropertyExchangeRate, not even an explicit nil
### GetRecurringInterval

`func (o *ModelVoucher) GetRecurringInterval() string`

GetRecurringInterval returns the RecurringInterval field if non-nil, zero value otherwise.

### GetRecurringIntervalOk

`func (o *ModelVoucher) GetRecurringIntervalOk() (*string, bool)`

GetRecurringIntervalOk returns a tuple with the RecurringInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringInterval

`func (o *ModelVoucher) SetRecurringInterval(v string)`

SetRecurringInterval sets RecurringInterval field to given value.

### HasRecurringInterval

`func (o *ModelVoucher) HasRecurringInterval() bool`

HasRecurringInterval returns a boolean if a field has been set.

### SetRecurringIntervalNil

`func (o *ModelVoucher) SetRecurringIntervalNil(b bool)`

 SetRecurringIntervalNil sets the value for RecurringInterval to be an explicit nil

### UnsetRecurringInterval
`func (o *ModelVoucher) UnsetRecurringInterval()`

UnsetRecurringInterval ensures that no value is present for RecurringInterval, not even an explicit nil
### GetRecurringStartDate

`func (o *ModelVoucher) GetRecurringStartDate() time.Time`

GetRecurringStartDate returns the RecurringStartDate field if non-nil, zero value otherwise.

### GetRecurringStartDateOk

`func (o *ModelVoucher) GetRecurringStartDateOk() (*time.Time, bool)`

GetRecurringStartDateOk returns a tuple with the RecurringStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringStartDate

`func (o *ModelVoucher) SetRecurringStartDate(v time.Time)`

SetRecurringStartDate sets RecurringStartDate field to given value.

### HasRecurringStartDate

`func (o *ModelVoucher) HasRecurringStartDate() bool`

HasRecurringStartDate returns a boolean if a field has been set.

### SetRecurringStartDateNil

`func (o *ModelVoucher) SetRecurringStartDateNil(b bool)`

 SetRecurringStartDateNil sets the value for RecurringStartDate to be an explicit nil

### UnsetRecurringStartDate
`func (o *ModelVoucher) UnsetRecurringStartDate()`

UnsetRecurringStartDate ensures that no value is present for RecurringStartDate, not even an explicit nil
### GetRecurringNextVoucher

`func (o *ModelVoucher) GetRecurringNextVoucher() time.Time`

GetRecurringNextVoucher returns the RecurringNextVoucher field if non-nil, zero value otherwise.

### GetRecurringNextVoucherOk

`func (o *ModelVoucher) GetRecurringNextVoucherOk() (*time.Time, bool)`

GetRecurringNextVoucherOk returns a tuple with the RecurringNextVoucher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringNextVoucher

`func (o *ModelVoucher) SetRecurringNextVoucher(v time.Time)`

SetRecurringNextVoucher sets RecurringNextVoucher field to given value.

### HasRecurringNextVoucher

`func (o *ModelVoucher) HasRecurringNextVoucher() bool`

HasRecurringNextVoucher returns a boolean if a field has been set.

### SetRecurringNextVoucherNil

`func (o *ModelVoucher) SetRecurringNextVoucherNil(b bool)`

 SetRecurringNextVoucherNil sets the value for RecurringNextVoucher to be an explicit nil

### UnsetRecurringNextVoucher
`func (o *ModelVoucher) UnsetRecurringNextVoucher()`

UnsetRecurringNextVoucher ensures that no value is present for RecurringNextVoucher, not even an explicit nil
### GetRecurringLastVoucher

`func (o *ModelVoucher) GetRecurringLastVoucher() time.Time`

GetRecurringLastVoucher returns the RecurringLastVoucher field if non-nil, zero value otherwise.

### GetRecurringLastVoucherOk

`func (o *ModelVoucher) GetRecurringLastVoucherOk() (*time.Time, bool)`

GetRecurringLastVoucherOk returns a tuple with the RecurringLastVoucher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringLastVoucher

`func (o *ModelVoucher) SetRecurringLastVoucher(v time.Time)`

SetRecurringLastVoucher sets RecurringLastVoucher field to given value.

### HasRecurringLastVoucher

`func (o *ModelVoucher) HasRecurringLastVoucher() bool`

HasRecurringLastVoucher returns a boolean if a field has been set.

### SetRecurringLastVoucherNil

`func (o *ModelVoucher) SetRecurringLastVoucherNil(b bool)`

 SetRecurringLastVoucherNil sets the value for RecurringLastVoucher to be an explicit nil

### UnsetRecurringLastVoucher
`func (o *ModelVoucher) UnsetRecurringLastVoucher()`

UnsetRecurringLastVoucher ensures that no value is present for RecurringLastVoucher, not even an explicit nil
### GetRecurringEndDate

`func (o *ModelVoucher) GetRecurringEndDate() time.Time`

GetRecurringEndDate returns the RecurringEndDate field if non-nil, zero value otherwise.

### GetRecurringEndDateOk

`func (o *ModelVoucher) GetRecurringEndDateOk() (*time.Time, bool)`

GetRecurringEndDateOk returns a tuple with the RecurringEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringEndDate

`func (o *ModelVoucher) SetRecurringEndDate(v time.Time)`

SetRecurringEndDate sets RecurringEndDate field to given value.

### HasRecurringEndDate

`func (o *ModelVoucher) HasRecurringEndDate() bool`

HasRecurringEndDate returns a boolean if a field has been set.

### SetRecurringEndDateNil

`func (o *ModelVoucher) SetRecurringEndDateNil(b bool)`

 SetRecurringEndDateNil sets the value for RecurringEndDate to be an explicit nil

### UnsetRecurringEndDate
`func (o *ModelVoucher) UnsetRecurringEndDate()`

UnsetRecurringEndDate ensures that no value is present for RecurringEndDate, not even an explicit nil
### GetEnshrined

`func (o *ModelVoucher) GetEnshrined() time.Time`

GetEnshrined returns the Enshrined field if non-nil, zero value otherwise.

### GetEnshrinedOk

`func (o *ModelVoucher) GetEnshrinedOk() (*time.Time, bool)`

GetEnshrinedOk returns a tuple with the Enshrined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnshrined

`func (o *ModelVoucher) SetEnshrined(v time.Time)`

SetEnshrined sets Enshrined field to given value.

### HasEnshrined

`func (o *ModelVoucher) HasEnshrined() bool`

HasEnshrined returns a boolean if a field has been set.

### SetEnshrinedNil

`func (o *ModelVoucher) SetEnshrinedNil(b bool)`

 SetEnshrinedNil sets the value for Enshrined to be an explicit nil

### UnsetEnshrined
`func (o *ModelVoucher) UnsetEnshrined()`

UnsetEnshrined ensures that no value is present for Enshrined, not even an explicit nil
### GetTaxSet

`func (o *ModelVoucher) GetTaxSet() ModelVoucherUpdateTaxSet`

GetTaxSet returns the TaxSet field if non-nil, zero value otherwise.

### GetTaxSetOk

`func (o *ModelVoucher) GetTaxSetOk() (*ModelVoucherUpdateTaxSet, bool)`

GetTaxSetOk returns a tuple with the TaxSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxSet

`func (o *ModelVoucher) SetTaxSet(v ModelVoucherUpdateTaxSet)`

SetTaxSet sets TaxSet field to given value.

### HasTaxSet

`func (o *ModelVoucher) HasTaxSet() bool`

HasTaxSet returns a boolean if a field has been set.

### SetTaxSetNil

`func (o *ModelVoucher) SetTaxSetNil(b bool)`

 SetTaxSetNil sets the value for TaxSet to be an explicit nil

### UnsetTaxSet
`func (o *ModelVoucher) UnsetTaxSet()`

UnsetTaxSet ensures that no value is present for TaxSet, not even an explicit nil
### GetPaymentDeadline

`func (o *ModelVoucher) GetPaymentDeadline() time.Time`

GetPaymentDeadline returns the PaymentDeadline field if non-nil, zero value otherwise.

### GetPaymentDeadlineOk

`func (o *ModelVoucher) GetPaymentDeadlineOk() (*time.Time, bool)`

GetPaymentDeadlineOk returns a tuple with the PaymentDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDeadline

`func (o *ModelVoucher) SetPaymentDeadline(v time.Time)`

SetPaymentDeadline sets PaymentDeadline field to given value.

### HasPaymentDeadline

`func (o *ModelVoucher) HasPaymentDeadline() bool`

HasPaymentDeadline returns a boolean if a field has been set.

### SetPaymentDeadlineNil

`func (o *ModelVoucher) SetPaymentDeadlineNil(b bool)`

 SetPaymentDeadlineNil sets the value for PaymentDeadline to be an explicit nil

### UnsetPaymentDeadline
`func (o *ModelVoucher) UnsetPaymentDeadline()`

UnsetPaymentDeadline ensures that no value is present for PaymentDeadline, not even an explicit nil
### GetDeliveryDate

`func (o *ModelVoucher) GetDeliveryDate() time.Time`

GetDeliveryDate returns the DeliveryDate field if non-nil, zero value otherwise.

### GetDeliveryDateOk

`func (o *ModelVoucher) GetDeliveryDateOk() (*time.Time, bool)`

GetDeliveryDateOk returns a tuple with the DeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDate

`func (o *ModelVoucher) SetDeliveryDate(v time.Time)`

SetDeliveryDate sets DeliveryDate field to given value.

### HasDeliveryDate

`func (o *ModelVoucher) HasDeliveryDate() bool`

HasDeliveryDate returns a boolean if a field has been set.

### GetDeliveryDateUntil

`func (o *ModelVoucher) GetDeliveryDateUntil() time.Time`

GetDeliveryDateUntil returns the DeliveryDateUntil field if non-nil, zero value otherwise.

### GetDeliveryDateUntilOk

`func (o *ModelVoucher) GetDeliveryDateUntilOk() (*time.Time, bool)`

GetDeliveryDateUntilOk returns a tuple with the DeliveryDateUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDateUntil

`func (o *ModelVoucher) SetDeliveryDateUntil(v time.Time)`

SetDeliveryDateUntil sets DeliveryDateUntil field to given value.

### HasDeliveryDateUntil

`func (o *ModelVoucher) HasDeliveryDateUntil() bool`

HasDeliveryDateUntil returns a boolean if a field has been set.

### SetDeliveryDateUntilNil

`func (o *ModelVoucher) SetDeliveryDateUntilNil(b bool)`

 SetDeliveryDateUntilNil sets the value for DeliveryDateUntil to be an explicit nil

### UnsetDeliveryDateUntil
`func (o *ModelVoucher) UnsetDeliveryDateUntil()`

UnsetDeliveryDateUntil ensures that no value is present for DeliveryDateUntil, not even an explicit nil
### GetDocument

`func (o *ModelVoucher) GetDocument() ModelVoucherUpdateDocument`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *ModelVoucher) GetDocumentOk() (*ModelVoucherUpdateDocument, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *ModelVoucher) SetDocument(v ModelVoucherUpdateDocument)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *ModelVoucher) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### SetDocumentNil

`func (o *ModelVoucher) SetDocumentNil(b bool)`

 SetDocumentNil sets the value for Document to be an explicit nil

### UnsetDocument
`func (o *ModelVoucher) UnsetDocument()`

UnsetDocument ensures that no value is present for Document, not even an explicit nil
### GetCostCentre

`func (o *ModelVoucher) GetCostCentre() ModelVoucherUpdateCostCentre`

GetCostCentre returns the CostCentre field if non-nil, zero value otherwise.

### GetCostCentreOk

`func (o *ModelVoucher) GetCostCentreOk() (*ModelVoucherUpdateCostCentre, bool)`

GetCostCentreOk returns a tuple with the CostCentre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCentre

`func (o *ModelVoucher) SetCostCentre(v ModelVoucherUpdateCostCentre)`

SetCostCentre sets CostCentre field to given value.

### HasCostCentre

`func (o *ModelVoucher) HasCostCentre() bool`

HasCostCentre returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


