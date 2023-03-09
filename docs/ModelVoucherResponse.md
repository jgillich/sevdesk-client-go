# ModelVoucherResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The voucher id | [optional] [readonly] 
**ObjectName** | Pointer to **string** | The voucher object name | [optional] [readonly] 
**MapAll** | Pointer to **bool** |  | [optional] 
**Create** | Pointer to **time.Time** | Date of voucher creation | [optional] [readonly] 
**Update** | Pointer to **time.Time** | Date of last voucher update | [optional] [readonly] 
**SevClient** | Pointer to [**ModelVoucherResponseSevClient**](ModelVoucherResponseSevClient.md) |  | [optional] 
**CreateUser** | Pointer to [**ModelVoucherResponseCreateUser**](ModelVoucherResponseCreateUser.md) |  | [optional] 
**VoucherDate** | Pointer to **NullableTime** | Needs to be provided as timestamp or dd.mm.yyyy | [optional] 
**Supplier** | Pointer to [**NullableModelVoucherResponseSupplier**](ModelVoucherResponseSupplier.md) |  | [optional] 
**SupplierName** | Pointer to **NullableString** | The supplier name.&lt;br&gt;       The value you provide here will determine what supplier name is shown for the voucher in case you did not provide a supplier. | [optional] 
**Description** | Pointer to **NullableString** | The description of the voucher. Essentially the voucher number. | [optional] 
**Document** | Pointer to [**NullableModelVoucherResponseDocument**](ModelVoucherResponseDocument.md) |  | [optional] 
**PayDate** | Pointer to **NullableTime** | Needs to be timestamp or dd.mm.yyyy | [optional] 
**Status** | Pointer to **NullableString** | Please have a look in       &lt;a href&#x3D;&#39;https://api.sevdesk.de/#section/Types-and-status-of-vouchers&#39;&gt;status of vouchers&lt;/a&gt;      to see what the different status codes mean | [optional] 
**SumNet** | Pointer to **string** | Net sum of the voucher | [optional] [readonly] 
**SumTax** | Pointer to **string** | Tax sum of the voucher | [optional] [readonly] 
**SumGross** | Pointer to **string** | Gross sum of the voucher | [optional] [readonly] 
**SumNetAccounting** | Pointer to **string** | Net accounting sum of the voucher. Is usually the same as sumNet | [optional] [readonly] 
**SumTaxAccounting** | Pointer to **string** | Tax accounting sum of the voucher. Is usually the same as sumTax | [optional] [readonly] 
**SumGrossAccounting** | Pointer to **string** | Gross accounting sum of the voucher. Is usually the same as sumGross | [optional] [readonly] 
**SumDiscounts** | Pointer to **string** | Sum of all discounts in the voucher | [optional] [readonly] 
**SumDiscountsForeignCurrency** | Pointer to **string** | Discounts sum of the voucher in the foreign currency | [optional] [readonly] 
**PaidAmount** | Pointer to **NullableFloat32** | Amount which has already been paid for this voucher by the customer | [optional] [readonly] 
**TaxType** | Pointer to **NullableString** | Tax type of the voucher. There are four tax types: 1. default - Umsatzsteuer ausweisen 2. eu - Steuerfreie innergemeinschaftliche Lieferung (Europäische Union) 3. noteu - Steuerschuldnerschaft des Leistungsempfängers (außerhalb EU, z. B. Schweiz) 4. custom - Using custom tax set 5. ss - Not subject to VAT according to §19 1 UStG  Tax rates are heavily connected to the tax type used. | [optional] 
**CreditDebit** | Pointer to **NullableString** | Defines if your voucher is a credit (C) or debit (D) | [optional] 
**CostCentre** | Pointer to [**ModelVoucherResponseCostCentre**](ModelVoucherResponseCostCentre.md) |  | [optional] 
**VoucherType** | Pointer to **NullableString** | Type of the voucher. For more information on the different types, check       &lt;a href&#x3D;&#39;https://api.sevdesk.de/#section/Types-and-status-of-vouchers&#39;&gt;this&lt;/a&gt;   | [optional] 
**Currency** | Pointer to **NullableString** | specifies which currency the voucher should have. Attention: If the currency differs from the default currency stored in the account, then either the \&quot;propertyForeignCurrencyDeadline\&quot; or \&quot;propertyExchangeRate\&quot; parameter must be specified. If both parameters are specified, then the \&quot;propertyForeignCurrencyDeadline\&quot; parameter is preferred | [optional] 
**PropertyForeignCurrencyDeadline** | Pointer to **NullableTime** | Defines the exchange rate day and and then the exchange rate is set from sevDesk. Needs to be provided as timestamp or dd.mm.yyyy | [optional] 
**PropertyExchangeRate** | Pointer to **NullableString** | Defines the exchange rate | [optional] 
**RecurringInterval** | Pointer to **NullableString** | The DateInterval in which recurring vouchers are generated.&lt;br&gt;       Necessary attribute for all recurring vouchers. | [optional] 
**RecurringStartDate** | Pointer to **NullableTime** | The date when the recurring vouchers start being generated.&lt;br&gt;       Necessary attribute for all recurring vouchers. | [optional] 
**RecurringNextVoucher** | Pointer to **NullableTime** | The date when the next voucher should be generated.&lt;br&gt;       Necessary attribute for all recurring vouchers. | [optional] 
**RecurringLastVoucher** | Pointer to **NullableTime** | The date when the last voucher was generated. | [optional] 
**RecurringEndDate** | Pointer to **NullableTime** | The date when the recurring vouchers end being generated.&lt;br&gt;      Necessary attribute for all recurring vouchers. | [optional] 
**Enshrined** | Pointer to **NullableTime** | Defines if and when voucher was enshrined. Enshrined vouchers can not be manipulated. | [optional] 
**TaxSet** | Pointer to [**NullableModelVoucherResponseTaxSet**](ModelVoucherResponseTaxSet.md) |  | [optional] 
**PaymentDeadline** | Pointer to **NullableTime** | Payment deadline of the voucher. | [optional] 
**DeliveryDate** | Pointer to **time.Time** | Needs to be provided as timestamp or dd.mm.yyyy | [optional] 
**DeliveryDateUntil** | Pointer to **NullableTime** | Needs to be provided as timestamp or dd.mm.yyyy | [optional] 

## Methods

### NewModelVoucherResponse

`func NewModelVoucherResponse() *ModelVoucherResponse`

NewModelVoucherResponse instantiates a new ModelVoucherResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelVoucherResponseWithDefaults

`func NewModelVoucherResponseWithDefaults() *ModelVoucherResponse`

NewModelVoucherResponseWithDefaults instantiates a new ModelVoucherResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelVoucherResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelVoucherResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelVoucherResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelVoucherResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectName

`func (o *ModelVoucherResponse) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ModelVoucherResponse) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ModelVoucherResponse) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ModelVoucherResponse) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetMapAll

`func (o *ModelVoucherResponse) GetMapAll() bool`

GetMapAll returns the MapAll field if non-nil, zero value otherwise.

### GetMapAllOk

`func (o *ModelVoucherResponse) GetMapAllOk() (*bool, bool)`

GetMapAllOk returns a tuple with the MapAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapAll

`func (o *ModelVoucherResponse) SetMapAll(v bool)`

SetMapAll sets MapAll field to given value.

### HasMapAll

`func (o *ModelVoucherResponse) HasMapAll() bool`

HasMapAll returns a boolean if a field has been set.

### GetCreate

`func (o *ModelVoucherResponse) GetCreate() time.Time`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ModelVoucherResponse) GetCreateOk() (*time.Time, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ModelVoucherResponse) SetCreate(v time.Time)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ModelVoucherResponse) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUpdate

`func (o *ModelVoucherResponse) GetUpdate() time.Time`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ModelVoucherResponse) GetUpdateOk() (*time.Time, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ModelVoucherResponse) SetUpdate(v time.Time)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ModelVoucherResponse) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetSevClient

`func (o *ModelVoucherResponse) GetSevClient() ModelVoucherResponseSevClient`

GetSevClient returns the SevClient field if non-nil, zero value otherwise.

### GetSevClientOk

`func (o *ModelVoucherResponse) GetSevClientOk() (*ModelVoucherResponseSevClient, bool)`

GetSevClientOk returns a tuple with the SevClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSevClient

`func (o *ModelVoucherResponse) SetSevClient(v ModelVoucherResponseSevClient)`

SetSevClient sets SevClient field to given value.

### HasSevClient

`func (o *ModelVoucherResponse) HasSevClient() bool`

HasSevClient returns a boolean if a field has been set.

### GetCreateUser

`func (o *ModelVoucherResponse) GetCreateUser() ModelVoucherResponseCreateUser`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *ModelVoucherResponse) GetCreateUserOk() (*ModelVoucherResponseCreateUser, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *ModelVoucherResponse) SetCreateUser(v ModelVoucherResponseCreateUser)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *ModelVoucherResponse) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetVoucherDate

`func (o *ModelVoucherResponse) GetVoucherDate() time.Time`

GetVoucherDate returns the VoucherDate field if non-nil, zero value otherwise.

### GetVoucherDateOk

`func (o *ModelVoucherResponse) GetVoucherDateOk() (*time.Time, bool)`

GetVoucherDateOk returns a tuple with the VoucherDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherDate

`func (o *ModelVoucherResponse) SetVoucherDate(v time.Time)`

SetVoucherDate sets VoucherDate field to given value.

### HasVoucherDate

`func (o *ModelVoucherResponse) HasVoucherDate() bool`

HasVoucherDate returns a boolean if a field has been set.

### SetVoucherDateNil

`func (o *ModelVoucherResponse) SetVoucherDateNil(b bool)`

 SetVoucherDateNil sets the value for VoucherDate to be an explicit nil

### UnsetVoucherDate
`func (o *ModelVoucherResponse) UnsetVoucherDate()`

UnsetVoucherDate ensures that no value is present for VoucherDate, not even an explicit nil
### GetSupplier

`func (o *ModelVoucherResponse) GetSupplier() ModelVoucherResponseSupplier`

GetSupplier returns the Supplier field if non-nil, zero value otherwise.

### GetSupplierOk

`func (o *ModelVoucherResponse) GetSupplierOk() (*ModelVoucherResponseSupplier, bool)`

GetSupplierOk returns a tuple with the Supplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplier

`func (o *ModelVoucherResponse) SetSupplier(v ModelVoucherResponseSupplier)`

SetSupplier sets Supplier field to given value.

### HasSupplier

`func (o *ModelVoucherResponse) HasSupplier() bool`

HasSupplier returns a boolean if a field has been set.

### SetSupplierNil

`func (o *ModelVoucherResponse) SetSupplierNil(b bool)`

 SetSupplierNil sets the value for Supplier to be an explicit nil

### UnsetSupplier
`func (o *ModelVoucherResponse) UnsetSupplier()`

UnsetSupplier ensures that no value is present for Supplier, not even an explicit nil
### GetSupplierName

`func (o *ModelVoucherResponse) GetSupplierName() string`

GetSupplierName returns the SupplierName field if non-nil, zero value otherwise.

### GetSupplierNameOk

`func (o *ModelVoucherResponse) GetSupplierNameOk() (*string, bool)`

GetSupplierNameOk returns a tuple with the SupplierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierName

`func (o *ModelVoucherResponse) SetSupplierName(v string)`

SetSupplierName sets SupplierName field to given value.

### HasSupplierName

`func (o *ModelVoucherResponse) HasSupplierName() bool`

HasSupplierName returns a boolean if a field has been set.

### SetSupplierNameNil

`func (o *ModelVoucherResponse) SetSupplierNameNil(b bool)`

 SetSupplierNameNil sets the value for SupplierName to be an explicit nil

### UnsetSupplierName
`func (o *ModelVoucherResponse) UnsetSupplierName()`

UnsetSupplierName ensures that no value is present for SupplierName, not even an explicit nil
### GetDescription

`func (o *ModelVoucherResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelVoucherResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelVoucherResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelVoucherResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ModelVoucherResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ModelVoucherResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDocument

`func (o *ModelVoucherResponse) GetDocument() ModelVoucherResponseDocument`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *ModelVoucherResponse) GetDocumentOk() (*ModelVoucherResponseDocument, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *ModelVoucherResponse) SetDocument(v ModelVoucherResponseDocument)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *ModelVoucherResponse) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### SetDocumentNil

`func (o *ModelVoucherResponse) SetDocumentNil(b bool)`

 SetDocumentNil sets the value for Document to be an explicit nil

### UnsetDocument
`func (o *ModelVoucherResponse) UnsetDocument()`

UnsetDocument ensures that no value is present for Document, not even an explicit nil
### GetPayDate

`func (o *ModelVoucherResponse) GetPayDate() time.Time`

GetPayDate returns the PayDate field if non-nil, zero value otherwise.

### GetPayDateOk

`func (o *ModelVoucherResponse) GetPayDateOk() (*time.Time, bool)`

GetPayDateOk returns a tuple with the PayDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayDate

`func (o *ModelVoucherResponse) SetPayDate(v time.Time)`

SetPayDate sets PayDate field to given value.

### HasPayDate

`func (o *ModelVoucherResponse) HasPayDate() bool`

HasPayDate returns a boolean if a field has been set.

### SetPayDateNil

`func (o *ModelVoucherResponse) SetPayDateNil(b bool)`

 SetPayDateNil sets the value for PayDate to be an explicit nil

### UnsetPayDate
`func (o *ModelVoucherResponse) UnsetPayDate()`

UnsetPayDate ensures that no value is present for PayDate, not even an explicit nil
### GetStatus

`func (o *ModelVoucherResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelVoucherResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelVoucherResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelVoucherResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ModelVoucherResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ModelVoucherResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetSumNet

`func (o *ModelVoucherResponse) GetSumNet() string`

GetSumNet returns the SumNet field if non-nil, zero value otherwise.

### GetSumNetOk

`func (o *ModelVoucherResponse) GetSumNetOk() (*string, bool)`

GetSumNetOk returns a tuple with the SumNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumNet

`func (o *ModelVoucherResponse) SetSumNet(v string)`

SetSumNet sets SumNet field to given value.

### HasSumNet

`func (o *ModelVoucherResponse) HasSumNet() bool`

HasSumNet returns a boolean if a field has been set.

### GetSumTax

`func (o *ModelVoucherResponse) GetSumTax() string`

GetSumTax returns the SumTax field if non-nil, zero value otherwise.

### GetSumTaxOk

`func (o *ModelVoucherResponse) GetSumTaxOk() (*string, bool)`

GetSumTaxOk returns a tuple with the SumTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumTax

`func (o *ModelVoucherResponse) SetSumTax(v string)`

SetSumTax sets SumTax field to given value.

### HasSumTax

`func (o *ModelVoucherResponse) HasSumTax() bool`

HasSumTax returns a boolean if a field has been set.

### GetSumGross

`func (o *ModelVoucherResponse) GetSumGross() string`

GetSumGross returns the SumGross field if non-nil, zero value otherwise.

### GetSumGrossOk

`func (o *ModelVoucherResponse) GetSumGrossOk() (*string, bool)`

GetSumGrossOk returns a tuple with the SumGross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumGross

`func (o *ModelVoucherResponse) SetSumGross(v string)`

SetSumGross sets SumGross field to given value.

### HasSumGross

`func (o *ModelVoucherResponse) HasSumGross() bool`

HasSumGross returns a boolean if a field has been set.

### GetSumNetAccounting

`func (o *ModelVoucherResponse) GetSumNetAccounting() string`

GetSumNetAccounting returns the SumNetAccounting field if non-nil, zero value otherwise.

### GetSumNetAccountingOk

`func (o *ModelVoucherResponse) GetSumNetAccountingOk() (*string, bool)`

GetSumNetAccountingOk returns a tuple with the SumNetAccounting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumNetAccounting

`func (o *ModelVoucherResponse) SetSumNetAccounting(v string)`

SetSumNetAccounting sets SumNetAccounting field to given value.

### HasSumNetAccounting

`func (o *ModelVoucherResponse) HasSumNetAccounting() bool`

HasSumNetAccounting returns a boolean if a field has been set.

### GetSumTaxAccounting

`func (o *ModelVoucherResponse) GetSumTaxAccounting() string`

GetSumTaxAccounting returns the SumTaxAccounting field if non-nil, zero value otherwise.

### GetSumTaxAccountingOk

`func (o *ModelVoucherResponse) GetSumTaxAccountingOk() (*string, bool)`

GetSumTaxAccountingOk returns a tuple with the SumTaxAccounting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumTaxAccounting

`func (o *ModelVoucherResponse) SetSumTaxAccounting(v string)`

SetSumTaxAccounting sets SumTaxAccounting field to given value.

### HasSumTaxAccounting

`func (o *ModelVoucherResponse) HasSumTaxAccounting() bool`

HasSumTaxAccounting returns a boolean if a field has been set.

### GetSumGrossAccounting

`func (o *ModelVoucherResponse) GetSumGrossAccounting() string`

GetSumGrossAccounting returns the SumGrossAccounting field if non-nil, zero value otherwise.

### GetSumGrossAccountingOk

`func (o *ModelVoucherResponse) GetSumGrossAccountingOk() (*string, bool)`

GetSumGrossAccountingOk returns a tuple with the SumGrossAccounting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumGrossAccounting

`func (o *ModelVoucherResponse) SetSumGrossAccounting(v string)`

SetSumGrossAccounting sets SumGrossAccounting field to given value.

### HasSumGrossAccounting

`func (o *ModelVoucherResponse) HasSumGrossAccounting() bool`

HasSumGrossAccounting returns a boolean if a field has been set.

### GetSumDiscounts

`func (o *ModelVoucherResponse) GetSumDiscounts() string`

GetSumDiscounts returns the SumDiscounts field if non-nil, zero value otherwise.

### GetSumDiscountsOk

`func (o *ModelVoucherResponse) GetSumDiscountsOk() (*string, bool)`

GetSumDiscountsOk returns a tuple with the SumDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumDiscounts

`func (o *ModelVoucherResponse) SetSumDiscounts(v string)`

SetSumDiscounts sets SumDiscounts field to given value.

### HasSumDiscounts

`func (o *ModelVoucherResponse) HasSumDiscounts() bool`

HasSumDiscounts returns a boolean if a field has been set.

### GetSumDiscountsForeignCurrency

`func (o *ModelVoucherResponse) GetSumDiscountsForeignCurrency() string`

GetSumDiscountsForeignCurrency returns the SumDiscountsForeignCurrency field if non-nil, zero value otherwise.

### GetSumDiscountsForeignCurrencyOk

`func (o *ModelVoucherResponse) GetSumDiscountsForeignCurrencyOk() (*string, bool)`

GetSumDiscountsForeignCurrencyOk returns a tuple with the SumDiscountsForeignCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumDiscountsForeignCurrency

`func (o *ModelVoucherResponse) SetSumDiscountsForeignCurrency(v string)`

SetSumDiscountsForeignCurrency sets SumDiscountsForeignCurrency field to given value.

### HasSumDiscountsForeignCurrency

`func (o *ModelVoucherResponse) HasSumDiscountsForeignCurrency() bool`

HasSumDiscountsForeignCurrency returns a boolean if a field has been set.

### GetPaidAmount

`func (o *ModelVoucherResponse) GetPaidAmount() float32`

GetPaidAmount returns the PaidAmount field if non-nil, zero value otherwise.

### GetPaidAmountOk

`func (o *ModelVoucherResponse) GetPaidAmountOk() (*float32, bool)`

GetPaidAmountOk returns a tuple with the PaidAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidAmount

`func (o *ModelVoucherResponse) SetPaidAmount(v float32)`

SetPaidAmount sets PaidAmount field to given value.

### HasPaidAmount

`func (o *ModelVoucherResponse) HasPaidAmount() bool`

HasPaidAmount returns a boolean if a field has been set.

### SetPaidAmountNil

`func (o *ModelVoucherResponse) SetPaidAmountNil(b bool)`

 SetPaidAmountNil sets the value for PaidAmount to be an explicit nil

### UnsetPaidAmount
`func (o *ModelVoucherResponse) UnsetPaidAmount()`

UnsetPaidAmount ensures that no value is present for PaidAmount, not even an explicit nil
### GetTaxType

`func (o *ModelVoucherResponse) GetTaxType() string`

GetTaxType returns the TaxType field if non-nil, zero value otherwise.

### GetTaxTypeOk

`func (o *ModelVoucherResponse) GetTaxTypeOk() (*string, bool)`

GetTaxTypeOk returns a tuple with the TaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxType

`func (o *ModelVoucherResponse) SetTaxType(v string)`

SetTaxType sets TaxType field to given value.

### HasTaxType

`func (o *ModelVoucherResponse) HasTaxType() bool`

HasTaxType returns a boolean if a field has been set.

### SetTaxTypeNil

`func (o *ModelVoucherResponse) SetTaxTypeNil(b bool)`

 SetTaxTypeNil sets the value for TaxType to be an explicit nil

### UnsetTaxType
`func (o *ModelVoucherResponse) UnsetTaxType()`

UnsetTaxType ensures that no value is present for TaxType, not even an explicit nil
### GetCreditDebit

`func (o *ModelVoucherResponse) GetCreditDebit() string`

GetCreditDebit returns the CreditDebit field if non-nil, zero value otherwise.

### GetCreditDebitOk

`func (o *ModelVoucherResponse) GetCreditDebitOk() (*string, bool)`

GetCreditDebitOk returns a tuple with the CreditDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditDebit

`func (o *ModelVoucherResponse) SetCreditDebit(v string)`

SetCreditDebit sets CreditDebit field to given value.

### HasCreditDebit

`func (o *ModelVoucherResponse) HasCreditDebit() bool`

HasCreditDebit returns a boolean if a field has been set.

### SetCreditDebitNil

`func (o *ModelVoucherResponse) SetCreditDebitNil(b bool)`

 SetCreditDebitNil sets the value for CreditDebit to be an explicit nil

### UnsetCreditDebit
`func (o *ModelVoucherResponse) UnsetCreditDebit()`

UnsetCreditDebit ensures that no value is present for CreditDebit, not even an explicit nil
### GetCostCentre

`func (o *ModelVoucherResponse) GetCostCentre() ModelVoucherResponseCostCentre`

GetCostCentre returns the CostCentre field if non-nil, zero value otherwise.

### GetCostCentreOk

`func (o *ModelVoucherResponse) GetCostCentreOk() (*ModelVoucherResponseCostCentre, bool)`

GetCostCentreOk returns a tuple with the CostCentre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCentre

`func (o *ModelVoucherResponse) SetCostCentre(v ModelVoucherResponseCostCentre)`

SetCostCentre sets CostCentre field to given value.

### HasCostCentre

`func (o *ModelVoucherResponse) HasCostCentre() bool`

HasCostCentre returns a boolean if a field has been set.

### GetVoucherType

`func (o *ModelVoucherResponse) GetVoucherType() string`

GetVoucherType returns the VoucherType field if non-nil, zero value otherwise.

### GetVoucherTypeOk

`func (o *ModelVoucherResponse) GetVoucherTypeOk() (*string, bool)`

GetVoucherTypeOk returns a tuple with the VoucherType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherType

`func (o *ModelVoucherResponse) SetVoucherType(v string)`

SetVoucherType sets VoucherType field to given value.

### HasVoucherType

`func (o *ModelVoucherResponse) HasVoucherType() bool`

HasVoucherType returns a boolean if a field has been set.

### SetVoucherTypeNil

`func (o *ModelVoucherResponse) SetVoucherTypeNil(b bool)`

 SetVoucherTypeNil sets the value for VoucherType to be an explicit nil

### UnsetVoucherType
`func (o *ModelVoucherResponse) UnsetVoucherType()`

UnsetVoucherType ensures that no value is present for VoucherType, not even an explicit nil
### GetCurrency

`func (o *ModelVoucherResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ModelVoucherResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ModelVoucherResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ModelVoucherResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *ModelVoucherResponse) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *ModelVoucherResponse) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetPropertyForeignCurrencyDeadline

`func (o *ModelVoucherResponse) GetPropertyForeignCurrencyDeadline() time.Time`

GetPropertyForeignCurrencyDeadline returns the PropertyForeignCurrencyDeadline field if non-nil, zero value otherwise.

### GetPropertyForeignCurrencyDeadlineOk

`func (o *ModelVoucherResponse) GetPropertyForeignCurrencyDeadlineOk() (*time.Time, bool)`

GetPropertyForeignCurrencyDeadlineOk returns a tuple with the PropertyForeignCurrencyDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyForeignCurrencyDeadline

`func (o *ModelVoucherResponse) SetPropertyForeignCurrencyDeadline(v time.Time)`

SetPropertyForeignCurrencyDeadline sets PropertyForeignCurrencyDeadline field to given value.

### HasPropertyForeignCurrencyDeadline

`func (o *ModelVoucherResponse) HasPropertyForeignCurrencyDeadline() bool`

HasPropertyForeignCurrencyDeadline returns a boolean if a field has been set.

### SetPropertyForeignCurrencyDeadlineNil

`func (o *ModelVoucherResponse) SetPropertyForeignCurrencyDeadlineNil(b bool)`

 SetPropertyForeignCurrencyDeadlineNil sets the value for PropertyForeignCurrencyDeadline to be an explicit nil

### UnsetPropertyForeignCurrencyDeadline
`func (o *ModelVoucherResponse) UnsetPropertyForeignCurrencyDeadline()`

UnsetPropertyForeignCurrencyDeadline ensures that no value is present for PropertyForeignCurrencyDeadline, not even an explicit nil
### GetPropertyExchangeRate

`func (o *ModelVoucherResponse) GetPropertyExchangeRate() string`

GetPropertyExchangeRate returns the PropertyExchangeRate field if non-nil, zero value otherwise.

### GetPropertyExchangeRateOk

`func (o *ModelVoucherResponse) GetPropertyExchangeRateOk() (*string, bool)`

GetPropertyExchangeRateOk returns a tuple with the PropertyExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyExchangeRate

`func (o *ModelVoucherResponse) SetPropertyExchangeRate(v string)`

SetPropertyExchangeRate sets PropertyExchangeRate field to given value.

### HasPropertyExchangeRate

`func (o *ModelVoucherResponse) HasPropertyExchangeRate() bool`

HasPropertyExchangeRate returns a boolean if a field has been set.

### SetPropertyExchangeRateNil

`func (o *ModelVoucherResponse) SetPropertyExchangeRateNil(b bool)`

 SetPropertyExchangeRateNil sets the value for PropertyExchangeRate to be an explicit nil

### UnsetPropertyExchangeRate
`func (o *ModelVoucherResponse) UnsetPropertyExchangeRate()`

UnsetPropertyExchangeRate ensures that no value is present for PropertyExchangeRate, not even an explicit nil
### GetRecurringInterval

`func (o *ModelVoucherResponse) GetRecurringInterval() string`

GetRecurringInterval returns the RecurringInterval field if non-nil, zero value otherwise.

### GetRecurringIntervalOk

`func (o *ModelVoucherResponse) GetRecurringIntervalOk() (*string, bool)`

GetRecurringIntervalOk returns a tuple with the RecurringInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringInterval

`func (o *ModelVoucherResponse) SetRecurringInterval(v string)`

SetRecurringInterval sets RecurringInterval field to given value.

### HasRecurringInterval

`func (o *ModelVoucherResponse) HasRecurringInterval() bool`

HasRecurringInterval returns a boolean if a field has been set.

### SetRecurringIntervalNil

`func (o *ModelVoucherResponse) SetRecurringIntervalNil(b bool)`

 SetRecurringIntervalNil sets the value for RecurringInterval to be an explicit nil

### UnsetRecurringInterval
`func (o *ModelVoucherResponse) UnsetRecurringInterval()`

UnsetRecurringInterval ensures that no value is present for RecurringInterval, not even an explicit nil
### GetRecurringStartDate

`func (o *ModelVoucherResponse) GetRecurringStartDate() time.Time`

GetRecurringStartDate returns the RecurringStartDate field if non-nil, zero value otherwise.

### GetRecurringStartDateOk

`func (o *ModelVoucherResponse) GetRecurringStartDateOk() (*time.Time, bool)`

GetRecurringStartDateOk returns a tuple with the RecurringStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringStartDate

`func (o *ModelVoucherResponse) SetRecurringStartDate(v time.Time)`

SetRecurringStartDate sets RecurringStartDate field to given value.

### HasRecurringStartDate

`func (o *ModelVoucherResponse) HasRecurringStartDate() bool`

HasRecurringStartDate returns a boolean if a field has been set.

### SetRecurringStartDateNil

`func (o *ModelVoucherResponse) SetRecurringStartDateNil(b bool)`

 SetRecurringStartDateNil sets the value for RecurringStartDate to be an explicit nil

### UnsetRecurringStartDate
`func (o *ModelVoucherResponse) UnsetRecurringStartDate()`

UnsetRecurringStartDate ensures that no value is present for RecurringStartDate, not even an explicit nil
### GetRecurringNextVoucher

`func (o *ModelVoucherResponse) GetRecurringNextVoucher() time.Time`

GetRecurringNextVoucher returns the RecurringNextVoucher field if non-nil, zero value otherwise.

### GetRecurringNextVoucherOk

`func (o *ModelVoucherResponse) GetRecurringNextVoucherOk() (*time.Time, bool)`

GetRecurringNextVoucherOk returns a tuple with the RecurringNextVoucher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringNextVoucher

`func (o *ModelVoucherResponse) SetRecurringNextVoucher(v time.Time)`

SetRecurringNextVoucher sets RecurringNextVoucher field to given value.

### HasRecurringNextVoucher

`func (o *ModelVoucherResponse) HasRecurringNextVoucher() bool`

HasRecurringNextVoucher returns a boolean if a field has been set.

### SetRecurringNextVoucherNil

`func (o *ModelVoucherResponse) SetRecurringNextVoucherNil(b bool)`

 SetRecurringNextVoucherNil sets the value for RecurringNextVoucher to be an explicit nil

### UnsetRecurringNextVoucher
`func (o *ModelVoucherResponse) UnsetRecurringNextVoucher()`

UnsetRecurringNextVoucher ensures that no value is present for RecurringNextVoucher, not even an explicit nil
### GetRecurringLastVoucher

`func (o *ModelVoucherResponse) GetRecurringLastVoucher() time.Time`

GetRecurringLastVoucher returns the RecurringLastVoucher field if non-nil, zero value otherwise.

### GetRecurringLastVoucherOk

`func (o *ModelVoucherResponse) GetRecurringLastVoucherOk() (*time.Time, bool)`

GetRecurringLastVoucherOk returns a tuple with the RecurringLastVoucher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringLastVoucher

`func (o *ModelVoucherResponse) SetRecurringLastVoucher(v time.Time)`

SetRecurringLastVoucher sets RecurringLastVoucher field to given value.

### HasRecurringLastVoucher

`func (o *ModelVoucherResponse) HasRecurringLastVoucher() bool`

HasRecurringLastVoucher returns a boolean if a field has been set.

### SetRecurringLastVoucherNil

`func (o *ModelVoucherResponse) SetRecurringLastVoucherNil(b bool)`

 SetRecurringLastVoucherNil sets the value for RecurringLastVoucher to be an explicit nil

### UnsetRecurringLastVoucher
`func (o *ModelVoucherResponse) UnsetRecurringLastVoucher()`

UnsetRecurringLastVoucher ensures that no value is present for RecurringLastVoucher, not even an explicit nil
### GetRecurringEndDate

`func (o *ModelVoucherResponse) GetRecurringEndDate() time.Time`

GetRecurringEndDate returns the RecurringEndDate field if non-nil, zero value otherwise.

### GetRecurringEndDateOk

`func (o *ModelVoucherResponse) GetRecurringEndDateOk() (*time.Time, bool)`

GetRecurringEndDateOk returns a tuple with the RecurringEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringEndDate

`func (o *ModelVoucherResponse) SetRecurringEndDate(v time.Time)`

SetRecurringEndDate sets RecurringEndDate field to given value.

### HasRecurringEndDate

`func (o *ModelVoucherResponse) HasRecurringEndDate() bool`

HasRecurringEndDate returns a boolean if a field has been set.

### SetRecurringEndDateNil

`func (o *ModelVoucherResponse) SetRecurringEndDateNil(b bool)`

 SetRecurringEndDateNil sets the value for RecurringEndDate to be an explicit nil

### UnsetRecurringEndDate
`func (o *ModelVoucherResponse) UnsetRecurringEndDate()`

UnsetRecurringEndDate ensures that no value is present for RecurringEndDate, not even an explicit nil
### GetEnshrined

`func (o *ModelVoucherResponse) GetEnshrined() time.Time`

GetEnshrined returns the Enshrined field if non-nil, zero value otherwise.

### GetEnshrinedOk

`func (o *ModelVoucherResponse) GetEnshrinedOk() (*time.Time, bool)`

GetEnshrinedOk returns a tuple with the Enshrined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnshrined

`func (o *ModelVoucherResponse) SetEnshrined(v time.Time)`

SetEnshrined sets Enshrined field to given value.

### HasEnshrined

`func (o *ModelVoucherResponse) HasEnshrined() bool`

HasEnshrined returns a boolean if a field has been set.

### SetEnshrinedNil

`func (o *ModelVoucherResponse) SetEnshrinedNil(b bool)`

 SetEnshrinedNil sets the value for Enshrined to be an explicit nil

### UnsetEnshrined
`func (o *ModelVoucherResponse) UnsetEnshrined()`

UnsetEnshrined ensures that no value is present for Enshrined, not even an explicit nil
### GetTaxSet

`func (o *ModelVoucherResponse) GetTaxSet() ModelVoucherResponseTaxSet`

GetTaxSet returns the TaxSet field if non-nil, zero value otherwise.

### GetTaxSetOk

`func (o *ModelVoucherResponse) GetTaxSetOk() (*ModelVoucherResponseTaxSet, bool)`

GetTaxSetOk returns a tuple with the TaxSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxSet

`func (o *ModelVoucherResponse) SetTaxSet(v ModelVoucherResponseTaxSet)`

SetTaxSet sets TaxSet field to given value.

### HasTaxSet

`func (o *ModelVoucherResponse) HasTaxSet() bool`

HasTaxSet returns a boolean if a field has been set.

### SetTaxSetNil

`func (o *ModelVoucherResponse) SetTaxSetNil(b bool)`

 SetTaxSetNil sets the value for TaxSet to be an explicit nil

### UnsetTaxSet
`func (o *ModelVoucherResponse) UnsetTaxSet()`

UnsetTaxSet ensures that no value is present for TaxSet, not even an explicit nil
### GetPaymentDeadline

`func (o *ModelVoucherResponse) GetPaymentDeadline() time.Time`

GetPaymentDeadline returns the PaymentDeadline field if non-nil, zero value otherwise.

### GetPaymentDeadlineOk

`func (o *ModelVoucherResponse) GetPaymentDeadlineOk() (*time.Time, bool)`

GetPaymentDeadlineOk returns a tuple with the PaymentDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDeadline

`func (o *ModelVoucherResponse) SetPaymentDeadline(v time.Time)`

SetPaymentDeadline sets PaymentDeadline field to given value.

### HasPaymentDeadline

`func (o *ModelVoucherResponse) HasPaymentDeadline() bool`

HasPaymentDeadline returns a boolean if a field has been set.

### SetPaymentDeadlineNil

`func (o *ModelVoucherResponse) SetPaymentDeadlineNil(b bool)`

 SetPaymentDeadlineNil sets the value for PaymentDeadline to be an explicit nil

### UnsetPaymentDeadline
`func (o *ModelVoucherResponse) UnsetPaymentDeadline()`

UnsetPaymentDeadline ensures that no value is present for PaymentDeadline, not even an explicit nil
### GetDeliveryDate

`func (o *ModelVoucherResponse) GetDeliveryDate() time.Time`

GetDeliveryDate returns the DeliveryDate field if non-nil, zero value otherwise.

### GetDeliveryDateOk

`func (o *ModelVoucherResponse) GetDeliveryDateOk() (*time.Time, bool)`

GetDeliveryDateOk returns a tuple with the DeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDate

`func (o *ModelVoucherResponse) SetDeliveryDate(v time.Time)`

SetDeliveryDate sets DeliveryDate field to given value.

### HasDeliveryDate

`func (o *ModelVoucherResponse) HasDeliveryDate() bool`

HasDeliveryDate returns a boolean if a field has been set.

### GetDeliveryDateUntil

`func (o *ModelVoucherResponse) GetDeliveryDateUntil() time.Time`

GetDeliveryDateUntil returns the DeliveryDateUntil field if non-nil, zero value otherwise.

### GetDeliveryDateUntilOk

`func (o *ModelVoucherResponse) GetDeliveryDateUntilOk() (*time.Time, bool)`

GetDeliveryDateUntilOk returns a tuple with the DeliveryDateUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDateUntil

`func (o *ModelVoucherResponse) SetDeliveryDateUntil(v time.Time)`

SetDeliveryDateUntil sets DeliveryDateUntil field to given value.

### HasDeliveryDateUntil

`func (o *ModelVoucherResponse) HasDeliveryDateUntil() bool`

HasDeliveryDateUntil returns a boolean if a field has been set.

### SetDeliveryDateUntilNil

`func (o *ModelVoucherResponse) SetDeliveryDateUntilNil(b bool)`

 SetDeliveryDateUntilNil sets the value for DeliveryDateUntil to be an explicit nil

### UnsetDeliveryDateUntil
`func (o *ModelVoucherResponse) UnsetDeliveryDateUntil()`

UnsetDeliveryDateUntil ensures that no value is present for DeliveryDateUntil, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


