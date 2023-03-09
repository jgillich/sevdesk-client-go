/*
sevDesk API

<b>Contact:</b> To contact our support click <a href='https://landing.sevdesk.de/service-support-center-technik'>here</a><br><br>   # General information  Welcome to our API!<br>  sevDesk offers you the possibility of retrieving data using an interface, namely the sevDesk API, and making changes without having to use the web UI. The sevDesk interface is a REST-Full API. All sevDesk data and functions that are used in the web UI can also be controlled through the API.   # Cross-Origin Resource Sharing  This API features Cross-Origin Resource Sharing (CORS).<br>  It enables cross-domain communication from the browser.<br>  All responses have a wildcard same-origin which makes them completely public and accessible to everyone, including any code on any site.    # Embedding resources  When retrieving resources by using this API, you might encounter nested resources in the resources you requested.<br>  For example, an invoice always contains a contact, of which you can see the ID and the object name.<br>  This API gives you the possibility to embed these resources completely into the resources you originally requested.<br>  Taking our invoice example, this would mean, that you would not only see the ID and object name of a contact, but rather the complete contact resource.    To embed resources, all you need to do is to add the query parameter 'embed' to your GET request.<br>  As values, you can provide the name of the nested resource.<br>  Multiple nested resources are also possible by providing multiple names, separated by a comma.    # Authentication and Authorization  The sevDesk API uses a token authentication to authorize calls. For this purpose every sevDesk administrator has one API token, which is a <b>hexadecimal string</b> containing <b>32 characters</b>. The following clip shows where you can find the API token if this is your first time with our API.<br><br> <video src='OpenAPI/img/findingTheApiToken.mp4' controls width='640' height='360'> Ihr Browser kann dieses Video nicht wiedergeben.<br/> Dieses Video zeigt wie sie Ihr sevDesk API Token finden. </video> <br> The token will be needed in every request that you want to send and needs to be attached to the request url as a <b>Query Parameter</b><br> or provided as a value of an <b>Authorization Header</b>.<br> For security reasons, we suggest putting the API Token in the Authorization Header and not in the query string.<br> However, in the request examples in this documentation, we will keep it in the query string, as it is easier for you to copy them and try them yourself.<br> The following url is an example that shows where your token needs to be placed as a query parameter.<br> In this case, we used some random API token. <ul> <li><span>ht</span>tps://my.sevdesk.de/api/v1/Contact?token=<span style='color:red'>b7794de0085f5cd00560f160f290af38</span></li> </ul> The next example shows the token in the Authorization Header. <ul> <li>\"Authorization\" :<span style='color:red'>\"b7794de0085f5cd00560f160f290af38\"</span></li> </ul> The api tokens have an infinite lifetime and, in other words, exist as long as the sevDesk user exists.<br> For this reason, the user should <b>NEVER</b> be deleted.<br> If really necessary, it is advisable to save the api token as we will <b>NOT</b> be able to retrieve it afterwards!<br> It is also possible to generate a new API token, for example, if you want to prevent the usage of your sevDesk account by other people who got your current API token.<br> To achieve this, you just need to click on the \"generate new\" symbol to the right of your token and confirm it with your password.  # API News  To never miss API news and updates again, subscribe to our <b>free API newsletter</b> and get all relevant information to keep your sevDesk software running smoothly. To subscribe, simply click <a href = 'https://landing.sevdesk.de/api-newsletter'><b>here</b></a> and confirm the email address to which we may send all updates relevant to you.  # API Requests  In our case, REST API requests need to be build by combining the following components. <table> <tr> <th><b>Component</b></th> <th><b>Description</b></th> </tr> <tr> <td>HTTP-Methods</td> <td> <ul> <li>GET (retrieve a resource)</li> <li>POST (create a resource)</li> <li>PUT (update a resource)</li> <li>DELETE (delete a resource)</li> </ul> </td> </tr> <tr> <td>URL of the API</td> <td><span style='color: #2aa198'>ht</span><span style='color: #2aa198'>tps://my.sevdesk.de/api/v1</span></td> </tr> <tr> <td>URI of the resource</td> <td>The resource to query.<br>For example contacts in sevDesk:<br><br> <span style='color:red'>/Contact</span><br><br> Which will result in the following complete URL:<br><br> <span style='color: #2aa198'>ht</span><span style='color: #2aa198'>tps://my.sevdesk.de/api/v1</span><span style='color:red'>/Contact</span> </td> </tr> <tr> <td>Query parameters</td> <td>Are attached by using the connectives <b>?</b> and <b>&</b> in the URL.<br></td> </tr> <tr> <td>Request headers</td> <td>Typical request headers are for example:<br> <div> <br> <ul> <li>Content-type</li> <li>Authorization</li> <li>Accept-Encoding</li> <li>User-Agent</li> <li>...</li> </ul> </div> </td> </tr> <tr> <td>Request body</td> <td>Mostly required in POST and PUT requests.<br> Often the request body contains json, in our case, it also accepts url-encoded data. </td> </tr> </table><br> <span style='color:red'>Note</span>: please pass a meaningful entry at the header \"User-Agent\". If the \"User-Agent\" is set meaningfully, we can offer better support in case of queries from customers.<br> An example how such a \"User-Agent\" can look like: \"Integration-name by abc\". <br><br> This is a sample request for retrieving existing contacts in sevDesk using curl:<br><br> <img src='OpenAPI/img/GETRequest.PNG' alt='Get Request' height='150'><br><br> As you can see, the request contains all the components mentioned above.<br> It's HTTP method is GET, it has a correct endpoint (<span style='color: #2aa198'>ht</span><span style='color: #2aa198'>tps://my.sevdesk.de/api/v1</span><span style='color:red'>/Contact</span>), query parameters like <b>token</b> and additional <b>header</b> information!<br><br> <b>Query Parameters</b><br><br> As you might have seen in the sample request above, there are several other parameters besides \"token\", located in the url.<br> Those are mostly optional but prove to be very useful for a lot of requests as they can limit, extend, sort or filter the data you will get as a response.<br><br> These are the three most used query parameter for the sevDesk API. <table> <tr> <th><b>Parameter</b></th> <th><b>Description</b></th> </tr> <tr> <td>limit</td> <td>Limits the number of entries that are returned.<br> Most useful in GET requests which will most likely deliver big sets of data like country or currency lists.<br> In this case, you can bypass the default limitation on returned entries by providing a high number. </td> </tr> <tr> <td>offset</td> <td>Specifies a certain offset for the data that will be returned.<br> As an example, you can specify \"offset=2\" if you want all entries except for the first two.</td> </tr> <tr> <td>embed</td> <td>Will extend some of the returned data.<br> A brief example can be found below.</td> </tr> </table> This is an example for the usage of the embed parameter.<br> The following first request will return all company contact entries in sevDesk up to a limit of 100 without any additional information and no offset.<br><br> <img src='OpenAPI/img/ContactQueryWithoutEmbed.PNG' width='900' height='850'><br><br> Now have a look at the category attribute located in the response.<br> Naturally, it just contains the id and the object name of the object but no further information.<br> We will now use the parameter embed with the value \"category\".<br><br> <img src='OpenAPI/img/ContactQueryWithEmbed.PNG' width='900' height='850'><br><br> As you can see, the category object is now extended and shows all the attributes and their corresponding values.<br><br> There are lot of other query parameters that can be used to filter the returned data for objects that match a certain pattern, however, those will not be mentioned here and instead can be found at the detail documentation of the most used API endpoints like contact, invoice or voucher.<br><br> <b>Request Headers</b><br><br> The HTTP request (response) headers allow the client as well as the server to pass additional information with the request.<br> They transfer the parameters and arguments which are important for transmitting data over HTTP.<br> Three headers which are useful / necessary when using the sevDesk API are \"Authorization, \"Accept\" and \"Content-type\".<br> Underneath is a brief description of why and how they should be used.<br><br> <b>Authorization</b><br><br> Can be used if you want to provide your API token in the header instead of having it in the url. <ul> <li>Authorization:<span style='color:red'>yourApiToken</span></li> </ul> <b>Accept</b><br><br> Specifies the format of the response.<br> Required for operations with a response body. <ul> <li>Accept:application/<span style='color:red'>format</span> </li> </ul> In our case, <code><span style='color:red'>format</span></code> could be replaced with <code>json</code> or <code>xml</code><br><br> <b>Content-type</b><br><br> Specifies which format is used in the request.<br> Is required for operations with a request body. <ul> <li>Content-type:application/<span style='color:red'>format</span></li> </ul> In our case,<code><span style='color:red'>format</span></code>could be replaced with <code>json</code> or <code>x-www-form-urlencoded</code> <br><br><b>API Responses</b><br><br> HTTP status codes<br> When calling the sevDesk API it is very likely that you will get a HTTP status code in the response.<br> Some API calls will also return JSON response bodies which will contain information about the resource.<br> Each status code which is returned will either be a <b>success</b> code or an <b>error</b> code.<br><br> Success codes <table> <tr> <th><b>Status code</b></th> <th><b>Description</b></th> </tr> <tr> <td><code>200 OK</code></td> <td>The request was successful</td> </tr> <tr> <td><code>201 Created</code></td> <td>Most likely to be found in the response of a <b>POST</b> request.<br> This code indicates that the desired resource was successfully created.</td> </tr> </table> <br>Error codes <table> <tr> <th><b>Status code</b></th> <th><b>Description</b></th> </tr> <tr> <td><code>400 Bad request</code></td> <td>The request you sent is most likely syntactically incorrect.<br> You should check if the parameters in the request body or the url are correct.</td> </tr> <tr> <td><code>401 Unauthorized</code></td> <td>The authentication failed.<br> Most likely caused by a missing or wrong API token.</td> </tr> <tr> <td><code>403 Forbidden</code></td> <td>You do not have the permission the access the resource which is requested.</td> </tr> <tr> <td><code>404 Not found</code></td> <td>The resource you specified does not exist.</td> </tr> <tr> <td><code>500 Internal server error</code></td> <td>An internal server error has occurred.<br> Normally this means that something went wrong on our side.<br> However, sometimes this error will appear if we missed to catch an error which is normally a 400 status code! </td> </tr> </table>  # Your First Request  After reading the introduction to our API, you should now be able to make your first call.<br> For testing our API, we would always recommend to create a trial account for sevDesk to prevent unwanted changes to your main account.<br> A trial account will be in the highest tariff (materials management), so every sevDesk function can be tested! <br><br>To start testing we would recommend one of the following tools: <ul> <li><a href='https://www.getpostman.com/'>Postman</a></li> <li><a href='https://curl.haxx.se/download.html'>cURL</a></li> </ul> This example will illustrate your first request, which is creating a new Contact in sevDesk.<br> <ol> <li>Download <a href='https://www.getpostman.com/'><b>Postman</b></a> for your desired system and start the application</li> <li>Enter <span><b>ht</span>tps://my.sevdesk.de/api/v1/Contact</b> as the url</li> <li>Use the connective <b>?</b> to append <b>token=</b> to the end of the url, or create an authorization header. Insert your API token as the value</li> <li>For this test, select <b>POST</b> as the HTTP method</li> <li>Go to <b>Headers</b> and enter the key-value pair <b>Content-type</b> + <b>application/x-www-form-urlencoded</b><br> As an alternative, you can just go to <b>Body</b> and select <b>x-www-form-urlencoded</b></li> <li>Now go to <b>Body</b> (if you are not there yet) and enter the key-value pairs as shown in the following picture<br><br> <img src='OpenAPI/img/FirstRequestPostman.PNG' width='900'><br><br></li> <li>Click on <b>Send</b>. Your response should now look like this:<br><br> <img src='OpenAPI/img/FirstRequestResponse.PNG' width='900'></li> </ol> As you can see, a successful response in this case returns a JSON-formatted response body containing the contact you just created.<br> For keeping it simple, this was only a minimal example of creating a contact.<br> There are however numerous combinations of parameters that you can provide which add information to your contact.

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sevdesk

import (
	"encoding/json"
	"time"
)

// checks if the ModelVoucherResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelVoucherResponse{}

// ModelVoucherResponse Voucher model
type ModelVoucherResponse struct {
	// The voucher id
	Id *string `json:"id,omitempty"`
	// The voucher object name
	ObjectName *string `json:"objectName,omitempty"`
	MapAll *bool `json:"mapAll,omitempty"`
	// Date of voucher creation
	Create *time.Time `json:"create,omitempty"`
	// Date of last voucher update
	Update *time.Time `json:"update,omitempty"`
	SevClient *ModelVoucherResponseSevClient `json:"sevClient,omitempty"`
	CreateUser *ModelVoucherResponseCreateUser `json:"createUser,omitempty"`
	// Needs to be provided as timestamp or dd.mm.yyyy
	VoucherDate NullableTime `json:"voucherDate,omitempty"`
	Supplier NullableModelVoucherResponseSupplier `json:"supplier,omitempty"`
	// The supplier name.<br>       The value you provide here will determine what supplier name is shown for the voucher in case you did not provide a supplier.
	SupplierName NullableString `json:"supplierName,omitempty"`
	// The description of the voucher. Essentially the voucher number.
	Description NullableString `json:"description,omitempty"`
	Document NullableModelVoucherResponseDocument `json:"document,omitempty"`
	// Needs to be timestamp or dd.mm.yyyy
	PayDate NullableTime `json:"payDate,omitempty"`
	// Please have a look in       <a href='https://api.sevdesk.de/#section/Types-and-status-of-vouchers'>status of vouchers</a>      to see what the different status codes mean
	Status NullableString `json:"status,omitempty"`
	// Net sum of the voucher
	SumNet *string `json:"sumNet,omitempty"`
	// Tax sum of the voucher
	SumTax *string `json:"sumTax,omitempty"`
	// Gross sum of the voucher
	SumGross *string `json:"sumGross,omitempty"`
	// Net accounting sum of the voucher. Is usually the same as sumNet
	SumNetAccounting *string `json:"sumNetAccounting,omitempty"`
	// Tax accounting sum of the voucher. Is usually the same as sumTax
	SumTaxAccounting *string `json:"sumTaxAccounting,omitempty"`
	// Gross accounting sum of the voucher. Is usually the same as sumGross
	SumGrossAccounting *string `json:"sumGrossAccounting,omitempty"`
	// Sum of all discounts in the voucher
	SumDiscounts *string `json:"sumDiscounts,omitempty"`
	// Discounts sum of the voucher in the foreign currency
	SumDiscountsForeignCurrency *string `json:"sumDiscountsForeignCurrency,omitempty"`
	// Amount which has already been paid for this voucher by the customer
	PaidAmount NullableFloat32 `json:"paidAmount,omitempty"`
	// Tax type of the voucher. There are four tax types: 1. default - Umsatzsteuer ausweisen 2. eu - Steuerfreie innergemeinschaftliche Lieferung (Europäische Union) 3. noteu - Steuerschuldnerschaft des Leistungsempfängers (außerhalb EU, z. B. Schweiz) 4. custom - Using custom tax set 5. ss - Not subject to VAT according to §19 1 UStG  Tax rates are heavily connected to the tax type used.
	TaxType NullableString `json:"taxType,omitempty"`
	// Defines if your voucher is a credit (C) or debit (D)
	CreditDebit NullableString `json:"creditDebit,omitempty"`
	CostCentre *ModelVoucherResponseCostCentre `json:"costCentre,omitempty"`
	// Type of the voucher. For more information on the different types, check       <a href='https://api.sevdesk.de/#section/Types-and-status-of-vouchers'>this</a>  
	VoucherType NullableString `json:"voucherType,omitempty"`
	// specifies which currency the voucher should have. Attention: If the currency differs from the default currency stored in the account, then either the \"propertyForeignCurrencyDeadline\" or \"propertyExchangeRate\" parameter must be specified. If both parameters are specified, then the \"propertyForeignCurrencyDeadline\" parameter is preferred
	Currency NullableString `json:"currency,omitempty"`
	// Defines the exchange rate day and and then the exchange rate is set from sevDesk. Needs to be provided as timestamp or dd.mm.yyyy
	PropertyForeignCurrencyDeadline NullableTime `json:"propertyForeignCurrencyDeadline,omitempty"`
	// Defines the exchange rate
	PropertyExchangeRate NullableString `json:"propertyExchangeRate,omitempty"`
	// The DateInterval in which recurring vouchers are generated.<br>       Necessary attribute for all recurring vouchers.
	RecurringInterval NullableString `json:"recurringInterval,omitempty"`
	// The date when the recurring vouchers start being generated.<br>       Necessary attribute for all recurring vouchers.
	RecurringStartDate NullableTime `json:"recurringStartDate,omitempty"`
	// The date when the next voucher should be generated.<br>       Necessary attribute for all recurring vouchers.
	RecurringNextVoucher NullableTime `json:"recurringNextVoucher,omitempty"`
	// The date when the last voucher was generated.
	RecurringLastVoucher NullableTime `json:"recurringLastVoucher,omitempty"`
	// The date when the recurring vouchers end being generated.<br>      Necessary attribute for all recurring vouchers.
	RecurringEndDate NullableTime `json:"recurringEndDate,omitempty"`
	// Defines if and when voucher was enshrined. Enshrined vouchers can not be manipulated.
	Enshrined NullableTime `json:"enshrined,omitempty"`
	TaxSet NullableModelVoucherResponseTaxSet `json:"taxSet,omitempty"`
	// Payment deadline of the voucher.
	PaymentDeadline NullableTime `json:"paymentDeadline,omitempty"`
	// Needs to be provided as timestamp or dd.mm.yyyy
	DeliveryDate *time.Time `json:"deliveryDate,omitempty"`
	// Needs to be provided as timestamp or dd.mm.yyyy
	DeliveryDateUntil NullableTime `json:"deliveryDateUntil,omitempty"`
}

// NewModelVoucherResponse instantiates a new ModelVoucherResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelVoucherResponse() *ModelVoucherResponse {
	this := ModelVoucherResponse{}
	return &this
}

// NewModelVoucherResponseWithDefaults instantiates a new ModelVoucherResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelVoucherResponseWithDefaults() *ModelVoucherResponse {
	this := ModelVoucherResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelVoucherResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelVoucherResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModelVoucherResponse) SetId(v string) {
	o.Id = &v
}

// GetObjectName returns the ObjectName field value if set, zero value otherwise.
func (o *ModelVoucherResponse) GetObjectName() string {
	if o == nil || IsNil(o.ObjectName) {
		var ret string
		return ret
	}
	return *o.ObjectName
}

// GetObjectNameOk returns a tuple with the ObjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelVoucherResponse) GetObjectNameOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectName) {
		return nil, false
	}
	return o.ObjectName, true
}

// HasObjectName returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasObjectName() bool {
	if o != nil && !IsNil(o.ObjectName) {
		return true
	}

	return false
}

// SetObjectName gets a reference to the given string and assigns it to the ObjectName field.
func (o *ModelVoucherResponse) SetObjectName(v string) {
	o.ObjectName = &v
}

// GetMapAll returns the MapAll field value if set, zero value otherwise.
func (o *ModelVoucherResponse) GetMapAll() bool {
	if o == nil || IsNil(o.MapAll) {
		var ret bool
		return ret
	}
	return *o.MapAll
}

// GetMapAllOk returns a tuple with the MapAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelVoucherResponse) GetMapAllOk() (*bool, bool) {
	if o == nil || IsNil(o.MapAll) {
		return nil, false
	}
	return o.MapAll, true
}

// HasMapAll returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasMapAll() bool {
	if o != nil && !IsNil(o.MapAll) {
		return true
	}

	return false
}

// SetMapAll gets a reference to the given bool and assigns it to the MapAll field.
func (o *ModelVoucherResponse) SetMapAll(v bool) {
	o.MapAll = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *ModelVoucherResponse) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelVoucherResponse) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *ModelVoucherResponse) SetCreate(v time.Time) {
	o.Create = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *ModelVoucherResponse) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelVoucherResponse) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *ModelVoucherResponse) SetUpdate(v time.Time) {
	o.Update = &v
}

// GetSevClient returns the SevClient field value if set, zero value otherwise.
func (o *ModelVoucherResponse) GetSevClient() ModelVoucherResponseSevClient {
	if o == nil || IsNil(o.SevClient) {
		var ret ModelVoucherResponseSevClient
		return ret
	}
	return *o.SevClient
}

// GetSevClientOk returns a tuple with the SevClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelVoucherResponse) GetSevClientOk() (*ModelVoucherResponseSevClient, bool) {
	if o == nil || IsNil(o.SevClient) {
		return nil, false
	}
	return o.SevClient, true
}

// HasSevClient returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasSevClient() bool {
	if o != nil && !IsNil(o.SevClient) {
		return true
	}

	return false
}

// SetSevClient gets a reference to the given ModelVoucherResponseSevClient and assigns it to the SevClient field.
func (o *ModelVoucherResponse) SetSevClient(v ModelVoucherResponseSevClient) {
	o.SevClient = &v
}

// GetCreateUser returns the CreateUser field value if set, zero value otherwise.
func (o *ModelVoucherResponse) GetCreateUser() ModelVoucherResponseCreateUser {
	if o == nil || IsNil(o.CreateUser) {
		var ret ModelVoucherResponseCreateUser
		return ret
	}
	return *o.CreateUser
}

// GetCreateUserOk returns a tuple with the CreateUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelVoucherResponse) GetCreateUserOk() (*ModelVoucherResponseCreateUser, bool) {
	if o == nil || IsNil(o.CreateUser) {
		return nil, false
	}
	return o.CreateUser, true
}

// HasCreateUser returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasCreateUser() bool {
	if o != nil && !IsNil(o.CreateUser) {
		return true
	}

	return false
}

// SetCreateUser gets a reference to the given ModelVoucherResponseCreateUser and assigns it to the CreateUser field.
func (o *ModelVoucherResponse) SetCreateUser(v ModelVoucherResponseCreateUser) {
	o.CreateUser = &v
}

// GetVoucherDate returns the VoucherDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelVoucherResponse) GetVoucherDate() time.Time {
	if o == nil || IsNil(o.VoucherDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.VoucherDate.Get()
}

// GetVoucherDateOk returns a tuple with the VoucherDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelVoucherResponse) GetVoucherDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.VoucherDate.Get(), o.VoucherDate.IsSet()
}

// HasVoucherDate returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasVoucherDate() bool {
	if o != nil && o.VoucherDate.IsSet() {
		return true
	}

	return false
}

// SetVoucherDate gets a reference to the given NullableTime and assigns it to the VoucherDate field.
func (o *ModelVoucherResponse) SetVoucherDate(v time.Time) {
	o.VoucherDate.Set(&v)
}
// SetVoucherDateNil sets the value for VoucherDate to be an explicit nil
func (o *ModelVoucherResponse) SetVoucherDateNil() {
	o.VoucherDate.Set(nil)
}

// UnsetVoucherDate ensures that no value is present for VoucherDate, not even an explicit nil
func (o *ModelVoucherResponse) UnsetVoucherDate() {
	o.VoucherDate.Unset()
}

// GetSupplier returns the Supplier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelVoucherResponse) GetSupplier() ModelVoucherResponseSupplier {
	if o == nil || IsNil(o.Supplier.Get()) {
		var ret ModelVoucherResponseSupplier
		return ret
	}
	return *o.Supplier.Get()
}

// GetSupplierOk returns a tuple with the Supplier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelVoucherResponse) GetSupplierOk() (*ModelVoucherResponseSupplier, bool) {
	if o == nil {
		return nil, false
	}
	return o.Supplier.Get(), o.Supplier.IsSet()
}

// HasSupplier returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasSupplier() bool {
	if o != nil && o.Supplier.IsSet() {
		return true
	}

	return false
}

// SetSupplier gets a reference to the given NullableModelVoucherResponseSupplier and assigns it to the Supplier field.
func (o *ModelVoucherResponse) SetSupplier(v ModelVoucherResponseSupplier) {
	o.Supplier.Set(&v)
}
// SetSupplierNil sets the value for Supplier to be an explicit nil
func (o *ModelVoucherResponse) SetSupplierNil() {
	o.Supplier.Set(nil)
}

// UnsetSupplier ensures that no value is present for Supplier, not even an explicit nil
func (o *ModelVoucherResponse) UnsetSupplier() {
	o.Supplier.Unset()
}

// GetSupplierName returns the SupplierName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelVoucherResponse) GetSupplierName() string {
	if o == nil || IsNil(o.SupplierName.Get()) {
		var ret string
		return ret
	}
	return *o.SupplierName.Get()
}

// GetSupplierNameOk returns a tuple with the SupplierName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelVoucherResponse) GetSupplierNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SupplierName.Get(), o.SupplierName.IsSet()
}

// HasSupplierName returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasSupplierName() bool {
	if o != nil && o.SupplierName.IsSet() {
		return true
	}

	return false
}

// SetSupplierName gets a reference to the given NullableString and assigns it to the SupplierName field.
func (o *ModelVoucherResponse) SetSupplierName(v string) {
	o.SupplierName.Set(&v)
}
// SetSupplierNameNil sets the value for SupplierName to be an explicit nil
func (o *ModelVoucherResponse) SetSupplierNameNil() {
	o.SupplierName.Set(nil)
}

// UnsetSupplierName ensures that no value is present for SupplierName, not even an explicit nil
func (o *ModelVoucherResponse) UnsetSupplierName() {
	o.SupplierName.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelVoucherResponse) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelVoucherResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ModelVoucherResponse) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ModelVoucherResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ModelVoucherResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetDocument returns the Document field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelVoucherResponse) GetDocument() ModelVoucherResponseDocument {
	if o == nil || IsNil(o.Document.Get()) {
		var ret ModelVoucherResponseDocument
		return ret
	}
	return *o.Document.Get()
}

// GetDocumentOk returns a tuple with the Document field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelVoucherResponse) GetDocumentOk() (*ModelVoucherResponseDocument, bool) {
	if o == nil {
		return nil, false
	}
	return o.Document.Get(), o.Document.IsSet()
}

// HasDocument returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasDocument() bool {
	if o != nil && o.Document.IsSet() {
		return true
	}

	return false
}

// SetDocument gets a reference to the given NullableModelVoucherResponseDocument and assigns it to the Document field.
func (o *ModelVoucherResponse) SetDocument(v ModelVoucherResponseDocument) {
	o.Document.Set(&v)
}
// SetDocumentNil sets the value for Document to be an explicit nil
func (o *ModelVoucherResponse) SetDocumentNil() {
	o.Document.Set(nil)
}

// UnsetDocument ensures that no value is present for Document, not even an explicit nil
func (o *ModelVoucherResponse) UnsetDocument() {
	o.Document.Unset()
}

// GetPayDate returns the PayDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelVoucherResponse) GetPayDate() time.Time {
	if o == nil || IsNil(o.PayDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.PayDate.Get()
}

// GetPayDateOk returns a tuple with the PayDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelVoucherResponse) GetPayDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.PayDate.Get(), o.PayDate.IsSet()
}

// HasPayDate returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasPayDate() bool {
	if o != nil && o.PayDate.IsSet() {
		return true
	}

	return false
}

// SetPayDate gets a reference to the given NullableTime and assigns it to the PayDate field.
func (o *ModelVoucherResponse) SetPayDate(v time.Time) {
	o.PayDate.Set(&v)
}
// SetPayDateNil sets the value for PayDate to be an explicit nil
func (o *ModelVoucherResponse) SetPayDateNil() {
	o.PayDate.Set(nil)
}

// UnsetPayDate ensures that no value is present for PayDate, not even an explicit nil
func (o *ModelVoucherResponse) UnsetPayDate() {
	o.PayDate.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelVoucherResponse) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelVoucherResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *ModelVoucherResponse) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *ModelVoucherResponse) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *ModelVoucherResponse) UnsetStatus() {
	o.Status.Unset()
}

// GetSumNet returns the SumNet field value if set, zero value otherwise.
func (o *ModelVoucherResponse) GetSumNet() string {
	if o == nil || IsNil(o.SumNet) {
		var ret string
		return ret
	}
	return *o.SumNet
}

// GetSumNetOk returns a tuple with the SumNet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelVoucherResponse) GetSumNetOk() (*string, bool) {
	if o == nil || IsNil(o.SumNet) {
		return nil, false
	}
	return o.SumNet, true
}

// HasSumNet returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasSumNet() bool {
	if o != nil && !IsNil(o.SumNet) {
		return true
	}

	return false
}

// SetSumNet gets a reference to the given string and assigns it to the SumNet field.
func (o *ModelVoucherResponse) SetSumNet(v string) {
	o.SumNet = &v
}

// GetSumTax returns the SumTax field value if set, zero value otherwise.
func (o *ModelVoucherResponse) GetSumTax() string {
	if o == nil || IsNil(o.SumTax) {
		var ret string
		return ret
	}
	return *o.SumTax
}

// GetSumTaxOk returns a tuple with the SumTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelVoucherResponse) GetSumTaxOk() (*string, bool) {
	if o == nil || IsNil(o.SumTax) {
		return nil, false
	}
	return o.SumTax, true
}

// HasSumTax returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasSumTax() bool {
	if o != nil && !IsNil(o.SumTax) {
		return true
	}

	return false
}

// SetSumTax gets a reference to the given string and assigns it to the SumTax field.
func (o *ModelVoucherResponse) SetSumTax(v string) {
	o.SumTax = &v
}

// GetSumGross returns the SumGross field value if set, zero value otherwise.
func (o *ModelVoucherResponse) GetSumGross() string {
	if o == nil || IsNil(o.SumGross) {
		var ret string
		return ret
	}
	return *o.SumGross
}

// GetSumGrossOk returns a tuple with the SumGross field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelVoucherResponse) GetSumGrossOk() (*string, bool) {
	if o == nil || IsNil(o.SumGross) {
		return nil, false
	}
	return o.SumGross, true
}

// HasSumGross returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasSumGross() bool {
	if o != nil && !IsNil(o.SumGross) {
		return true
	}

	return false
}

// SetSumGross gets a reference to the given string and assigns it to the SumGross field.
func (o *ModelVoucherResponse) SetSumGross(v string) {
	o.SumGross = &v
}

// GetSumNetAccounting returns the SumNetAccounting field value if set, zero value otherwise.
func (o *ModelVoucherResponse) GetSumNetAccounting() string {
	if o == nil || IsNil(o.SumNetAccounting) {
		var ret string
		return ret
	}
	return *o.SumNetAccounting
}

// GetSumNetAccountingOk returns a tuple with the SumNetAccounting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelVoucherResponse) GetSumNetAccountingOk() (*string, bool) {
	if o == nil || IsNil(o.SumNetAccounting) {
		return nil, false
	}
	return o.SumNetAccounting, true
}

// HasSumNetAccounting returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasSumNetAccounting() bool {
	if o != nil && !IsNil(o.SumNetAccounting) {
		return true
	}

	return false
}

// SetSumNetAccounting gets a reference to the given string and assigns it to the SumNetAccounting field.
func (o *ModelVoucherResponse) SetSumNetAccounting(v string) {
	o.SumNetAccounting = &v
}

// GetSumTaxAccounting returns the SumTaxAccounting field value if set, zero value otherwise.
func (o *ModelVoucherResponse) GetSumTaxAccounting() string {
	if o == nil || IsNil(o.SumTaxAccounting) {
		var ret string
		return ret
	}
	return *o.SumTaxAccounting
}

// GetSumTaxAccountingOk returns a tuple with the SumTaxAccounting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelVoucherResponse) GetSumTaxAccountingOk() (*string, bool) {
	if o == nil || IsNil(o.SumTaxAccounting) {
		return nil, false
	}
	return o.SumTaxAccounting, true
}

// HasSumTaxAccounting returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasSumTaxAccounting() bool {
	if o != nil && !IsNil(o.SumTaxAccounting) {
		return true
	}

	return false
}

// SetSumTaxAccounting gets a reference to the given string and assigns it to the SumTaxAccounting field.
func (o *ModelVoucherResponse) SetSumTaxAccounting(v string) {
	o.SumTaxAccounting = &v
}

// GetSumGrossAccounting returns the SumGrossAccounting field value if set, zero value otherwise.
func (o *ModelVoucherResponse) GetSumGrossAccounting() string {
	if o == nil || IsNil(o.SumGrossAccounting) {
		var ret string
		return ret
	}
	return *o.SumGrossAccounting
}

// GetSumGrossAccountingOk returns a tuple with the SumGrossAccounting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelVoucherResponse) GetSumGrossAccountingOk() (*string, bool) {
	if o == nil || IsNil(o.SumGrossAccounting) {
		return nil, false
	}
	return o.SumGrossAccounting, true
}

// HasSumGrossAccounting returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasSumGrossAccounting() bool {
	if o != nil && !IsNil(o.SumGrossAccounting) {
		return true
	}

	return false
}

// SetSumGrossAccounting gets a reference to the given string and assigns it to the SumGrossAccounting field.
func (o *ModelVoucherResponse) SetSumGrossAccounting(v string) {
	o.SumGrossAccounting = &v
}

// GetSumDiscounts returns the SumDiscounts field value if set, zero value otherwise.
func (o *ModelVoucherResponse) GetSumDiscounts() string {
	if o == nil || IsNil(o.SumDiscounts) {
		var ret string
		return ret
	}
	return *o.SumDiscounts
}

// GetSumDiscountsOk returns a tuple with the SumDiscounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelVoucherResponse) GetSumDiscountsOk() (*string, bool) {
	if o == nil || IsNil(o.SumDiscounts) {
		return nil, false
	}
	return o.SumDiscounts, true
}

// HasSumDiscounts returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasSumDiscounts() bool {
	if o != nil && !IsNil(o.SumDiscounts) {
		return true
	}

	return false
}

// SetSumDiscounts gets a reference to the given string and assigns it to the SumDiscounts field.
func (o *ModelVoucherResponse) SetSumDiscounts(v string) {
	o.SumDiscounts = &v
}

// GetSumDiscountsForeignCurrency returns the SumDiscountsForeignCurrency field value if set, zero value otherwise.
func (o *ModelVoucherResponse) GetSumDiscountsForeignCurrency() string {
	if o == nil || IsNil(o.SumDiscountsForeignCurrency) {
		var ret string
		return ret
	}
	return *o.SumDiscountsForeignCurrency
}

// GetSumDiscountsForeignCurrencyOk returns a tuple with the SumDiscountsForeignCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelVoucherResponse) GetSumDiscountsForeignCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.SumDiscountsForeignCurrency) {
		return nil, false
	}
	return o.SumDiscountsForeignCurrency, true
}

// HasSumDiscountsForeignCurrency returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasSumDiscountsForeignCurrency() bool {
	if o != nil && !IsNil(o.SumDiscountsForeignCurrency) {
		return true
	}

	return false
}

// SetSumDiscountsForeignCurrency gets a reference to the given string and assigns it to the SumDiscountsForeignCurrency field.
func (o *ModelVoucherResponse) SetSumDiscountsForeignCurrency(v string) {
	o.SumDiscountsForeignCurrency = &v
}

// GetPaidAmount returns the PaidAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelVoucherResponse) GetPaidAmount() float32 {
	if o == nil || IsNil(o.PaidAmount.Get()) {
		var ret float32
		return ret
	}
	return *o.PaidAmount.Get()
}

// GetPaidAmountOk returns a tuple with the PaidAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelVoucherResponse) GetPaidAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaidAmount.Get(), o.PaidAmount.IsSet()
}

// HasPaidAmount returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasPaidAmount() bool {
	if o != nil && o.PaidAmount.IsSet() {
		return true
	}

	return false
}

// SetPaidAmount gets a reference to the given NullableFloat32 and assigns it to the PaidAmount field.
func (o *ModelVoucherResponse) SetPaidAmount(v float32) {
	o.PaidAmount.Set(&v)
}
// SetPaidAmountNil sets the value for PaidAmount to be an explicit nil
func (o *ModelVoucherResponse) SetPaidAmountNil() {
	o.PaidAmount.Set(nil)
}

// UnsetPaidAmount ensures that no value is present for PaidAmount, not even an explicit nil
func (o *ModelVoucherResponse) UnsetPaidAmount() {
	o.PaidAmount.Unset()
}

// GetTaxType returns the TaxType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelVoucherResponse) GetTaxType() string {
	if o == nil || IsNil(o.TaxType.Get()) {
		var ret string
		return ret
	}
	return *o.TaxType.Get()
}

// GetTaxTypeOk returns a tuple with the TaxType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelVoucherResponse) GetTaxTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxType.Get(), o.TaxType.IsSet()
}

// HasTaxType returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasTaxType() bool {
	if o != nil && o.TaxType.IsSet() {
		return true
	}

	return false
}

// SetTaxType gets a reference to the given NullableString and assigns it to the TaxType field.
func (o *ModelVoucherResponse) SetTaxType(v string) {
	o.TaxType.Set(&v)
}
// SetTaxTypeNil sets the value for TaxType to be an explicit nil
func (o *ModelVoucherResponse) SetTaxTypeNil() {
	o.TaxType.Set(nil)
}

// UnsetTaxType ensures that no value is present for TaxType, not even an explicit nil
func (o *ModelVoucherResponse) UnsetTaxType() {
	o.TaxType.Unset()
}

// GetCreditDebit returns the CreditDebit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelVoucherResponse) GetCreditDebit() string {
	if o == nil || IsNil(o.CreditDebit.Get()) {
		var ret string
		return ret
	}
	return *o.CreditDebit.Get()
}

// GetCreditDebitOk returns a tuple with the CreditDebit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelVoucherResponse) GetCreditDebitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreditDebit.Get(), o.CreditDebit.IsSet()
}

// HasCreditDebit returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasCreditDebit() bool {
	if o != nil && o.CreditDebit.IsSet() {
		return true
	}

	return false
}

// SetCreditDebit gets a reference to the given NullableString and assigns it to the CreditDebit field.
func (o *ModelVoucherResponse) SetCreditDebit(v string) {
	o.CreditDebit.Set(&v)
}
// SetCreditDebitNil sets the value for CreditDebit to be an explicit nil
func (o *ModelVoucherResponse) SetCreditDebitNil() {
	o.CreditDebit.Set(nil)
}

// UnsetCreditDebit ensures that no value is present for CreditDebit, not even an explicit nil
func (o *ModelVoucherResponse) UnsetCreditDebit() {
	o.CreditDebit.Unset()
}

// GetCostCentre returns the CostCentre field value if set, zero value otherwise.
func (o *ModelVoucherResponse) GetCostCentre() ModelVoucherResponseCostCentre {
	if o == nil || IsNil(o.CostCentre) {
		var ret ModelVoucherResponseCostCentre
		return ret
	}
	return *o.CostCentre
}

// GetCostCentreOk returns a tuple with the CostCentre field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelVoucherResponse) GetCostCentreOk() (*ModelVoucherResponseCostCentre, bool) {
	if o == nil || IsNil(o.CostCentre) {
		return nil, false
	}
	return o.CostCentre, true
}

// HasCostCentre returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasCostCentre() bool {
	if o != nil && !IsNil(o.CostCentre) {
		return true
	}

	return false
}

// SetCostCentre gets a reference to the given ModelVoucherResponseCostCentre and assigns it to the CostCentre field.
func (o *ModelVoucherResponse) SetCostCentre(v ModelVoucherResponseCostCentre) {
	o.CostCentre = &v
}

// GetVoucherType returns the VoucherType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelVoucherResponse) GetVoucherType() string {
	if o == nil || IsNil(o.VoucherType.Get()) {
		var ret string
		return ret
	}
	return *o.VoucherType.Get()
}

// GetVoucherTypeOk returns a tuple with the VoucherType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelVoucherResponse) GetVoucherTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VoucherType.Get(), o.VoucherType.IsSet()
}

// HasVoucherType returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasVoucherType() bool {
	if o != nil && o.VoucherType.IsSet() {
		return true
	}

	return false
}

// SetVoucherType gets a reference to the given NullableString and assigns it to the VoucherType field.
func (o *ModelVoucherResponse) SetVoucherType(v string) {
	o.VoucherType.Set(&v)
}
// SetVoucherTypeNil sets the value for VoucherType to be an explicit nil
func (o *ModelVoucherResponse) SetVoucherTypeNil() {
	o.VoucherType.Set(nil)
}

// UnsetVoucherType ensures that no value is present for VoucherType, not even an explicit nil
func (o *ModelVoucherResponse) UnsetVoucherType() {
	o.VoucherType.Unset()
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelVoucherResponse) GetCurrency() string {
	if o == nil || IsNil(o.Currency.Get()) {
		var ret string
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelVoucherResponse) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasCurrency() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableString and assigns it to the Currency field.
func (o *ModelVoucherResponse) SetCurrency(v string) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *ModelVoucherResponse) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *ModelVoucherResponse) UnsetCurrency() {
	o.Currency.Unset()
}

// GetPropertyForeignCurrencyDeadline returns the PropertyForeignCurrencyDeadline field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelVoucherResponse) GetPropertyForeignCurrencyDeadline() time.Time {
	if o == nil || IsNil(o.PropertyForeignCurrencyDeadline.Get()) {
		var ret time.Time
		return ret
	}
	return *o.PropertyForeignCurrencyDeadline.Get()
}

// GetPropertyForeignCurrencyDeadlineOk returns a tuple with the PropertyForeignCurrencyDeadline field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelVoucherResponse) GetPropertyForeignCurrencyDeadlineOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.PropertyForeignCurrencyDeadline.Get(), o.PropertyForeignCurrencyDeadline.IsSet()
}

// HasPropertyForeignCurrencyDeadline returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasPropertyForeignCurrencyDeadline() bool {
	if o != nil && o.PropertyForeignCurrencyDeadline.IsSet() {
		return true
	}

	return false
}

// SetPropertyForeignCurrencyDeadline gets a reference to the given NullableTime and assigns it to the PropertyForeignCurrencyDeadline field.
func (o *ModelVoucherResponse) SetPropertyForeignCurrencyDeadline(v time.Time) {
	o.PropertyForeignCurrencyDeadline.Set(&v)
}
// SetPropertyForeignCurrencyDeadlineNil sets the value for PropertyForeignCurrencyDeadline to be an explicit nil
func (o *ModelVoucherResponse) SetPropertyForeignCurrencyDeadlineNil() {
	o.PropertyForeignCurrencyDeadline.Set(nil)
}

// UnsetPropertyForeignCurrencyDeadline ensures that no value is present for PropertyForeignCurrencyDeadline, not even an explicit nil
func (o *ModelVoucherResponse) UnsetPropertyForeignCurrencyDeadline() {
	o.PropertyForeignCurrencyDeadline.Unset()
}

// GetPropertyExchangeRate returns the PropertyExchangeRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelVoucherResponse) GetPropertyExchangeRate() string {
	if o == nil || IsNil(o.PropertyExchangeRate.Get()) {
		var ret string
		return ret
	}
	return *o.PropertyExchangeRate.Get()
}

// GetPropertyExchangeRateOk returns a tuple with the PropertyExchangeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelVoucherResponse) GetPropertyExchangeRateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PropertyExchangeRate.Get(), o.PropertyExchangeRate.IsSet()
}

// HasPropertyExchangeRate returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasPropertyExchangeRate() bool {
	if o != nil && o.PropertyExchangeRate.IsSet() {
		return true
	}

	return false
}

// SetPropertyExchangeRate gets a reference to the given NullableString and assigns it to the PropertyExchangeRate field.
func (o *ModelVoucherResponse) SetPropertyExchangeRate(v string) {
	o.PropertyExchangeRate.Set(&v)
}
// SetPropertyExchangeRateNil sets the value for PropertyExchangeRate to be an explicit nil
func (o *ModelVoucherResponse) SetPropertyExchangeRateNil() {
	o.PropertyExchangeRate.Set(nil)
}

// UnsetPropertyExchangeRate ensures that no value is present for PropertyExchangeRate, not even an explicit nil
func (o *ModelVoucherResponse) UnsetPropertyExchangeRate() {
	o.PropertyExchangeRate.Unset()
}

// GetRecurringInterval returns the RecurringInterval field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelVoucherResponse) GetRecurringInterval() string {
	if o == nil || IsNil(o.RecurringInterval.Get()) {
		var ret string
		return ret
	}
	return *o.RecurringInterval.Get()
}

// GetRecurringIntervalOk returns a tuple with the RecurringInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelVoucherResponse) GetRecurringIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecurringInterval.Get(), o.RecurringInterval.IsSet()
}

// HasRecurringInterval returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasRecurringInterval() bool {
	if o != nil && o.RecurringInterval.IsSet() {
		return true
	}

	return false
}

// SetRecurringInterval gets a reference to the given NullableString and assigns it to the RecurringInterval field.
func (o *ModelVoucherResponse) SetRecurringInterval(v string) {
	o.RecurringInterval.Set(&v)
}
// SetRecurringIntervalNil sets the value for RecurringInterval to be an explicit nil
func (o *ModelVoucherResponse) SetRecurringIntervalNil() {
	o.RecurringInterval.Set(nil)
}

// UnsetRecurringInterval ensures that no value is present for RecurringInterval, not even an explicit nil
func (o *ModelVoucherResponse) UnsetRecurringInterval() {
	o.RecurringInterval.Unset()
}

// GetRecurringStartDate returns the RecurringStartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelVoucherResponse) GetRecurringStartDate() time.Time {
	if o == nil || IsNil(o.RecurringStartDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.RecurringStartDate.Get()
}

// GetRecurringStartDateOk returns a tuple with the RecurringStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelVoucherResponse) GetRecurringStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecurringStartDate.Get(), o.RecurringStartDate.IsSet()
}

// HasRecurringStartDate returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasRecurringStartDate() bool {
	if o != nil && o.RecurringStartDate.IsSet() {
		return true
	}

	return false
}

// SetRecurringStartDate gets a reference to the given NullableTime and assigns it to the RecurringStartDate field.
func (o *ModelVoucherResponse) SetRecurringStartDate(v time.Time) {
	o.RecurringStartDate.Set(&v)
}
// SetRecurringStartDateNil sets the value for RecurringStartDate to be an explicit nil
func (o *ModelVoucherResponse) SetRecurringStartDateNil() {
	o.RecurringStartDate.Set(nil)
}

// UnsetRecurringStartDate ensures that no value is present for RecurringStartDate, not even an explicit nil
func (o *ModelVoucherResponse) UnsetRecurringStartDate() {
	o.RecurringStartDate.Unset()
}

// GetRecurringNextVoucher returns the RecurringNextVoucher field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelVoucherResponse) GetRecurringNextVoucher() time.Time {
	if o == nil || IsNil(o.RecurringNextVoucher.Get()) {
		var ret time.Time
		return ret
	}
	return *o.RecurringNextVoucher.Get()
}

// GetRecurringNextVoucherOk returns a tuple with the RecurringNextVoucher field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelVoucherResponse) GetRecurringNextVoucherOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecurringNextVoucher.Get(), o.RecurringNextVoucher.IsSet()
}

// HasRecurringNextVoucher returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasRecurringNextVoucher() bool {
	if o != nil && o.RecurringNextVoucher.IsSet() {
		return true
	}

	return false
}

// SetRecurringNextVoucher gets a reference to the given NullableTime and assigns it to the RecurringNextVoucher field.
func (o *ModelVoucherResponse) SetRecurringNextVoucher(v time.Time) {
	o.RecurringNextVoucher.Set(&v)
}
// SetRecurringNextVoucherNil sets the value for RecurringNextVoucher to be an explicit nil
func (o *ModelVoucherResponse) SetRecurringNextVoucherNil() {
	o.RecurringNextVoucher.Set(nil)
}

// UnsetRecurringNextVoucher ensures that no value is present for RecurringNextVoucher, not even an explicit nil
func (o *ModelVoucherResponse) UnsetRecurringNextVoucher() {
	o.RecurringNextVoucher.Unset()
}

// GetRecurringLastVoucher returns the RecurringLastVoucher field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelVoucherResponse) GetRecurringLastVoucher() time.Time {
	if o == nil || IsNil(o.RecurringLastVoucher.Get()) {
		var ret time.Time
		return ret
	}
	return *o.RecurringLastVoucher.Get()
}

// GetRecurringLastVoucherOk returns a tuple with the RecurringLastVoucher field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelVoucherResponse) GetRecurringLastVoucherOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecurringLastVoucher.Get(), o.RecurringLastVoucher.IsSet()
}

// HasRecurringLastVoucher returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasRecurringLastVoucher() bool {
	if o != nil && o.RecurringLastVoucher.IsSet() {
		return true
	}

	return false
}

// SetRecurringLastVoucher gets a reference to the given NullableTime and assigns it to the RecurringLastVoucher field.
func (o *ModelVoucherResponse) SetRecurringLastVoucher(v time.Time) {
	o.RecurringLastVoucher.Set(&v)
}
// SetRecurringLastVoucherNil sets the value for RecurringLastVoucher to be an explicit nil
func (o *ModelVoucherResponse) SetRecurringLastVoucherNil() {
	o.RecurringLastVoucher.Set(nil)
}

// UnsetRecurringLastVoucher ensures that no value is present for RecurringLastVoucher, not even an explicit nil
func (o *ModelVoucherResponse) UnsetRecurringLastVoucher() {
	o.RecurringLastVoucher.Unset()
}

// GetRecurringEndDate returns the RecurringEndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelVoucherResponse) GetRecurringEndDate() time.Time {
	if o == nil || IsNil(o.RecurringEndDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.RecurringEndDate.Get()
}

// GetRecurringEndDateOk returns a tuple with the RecurringEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelVoucherResponse) GetRecurringEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecurringEndDate.Get(), o.RecurringEndDate.IsSet()
}

// HasRecurringEndDate returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasRecurringEndDate() bool {
	if o != nil && o.RecurringEndDate.IsSet() {
		return true
	}

	return false
}

// SetRecurringEndDate gets a reference to the given NullableTime and assigns it to the RecurringEndDate field.
func (o *ModelVoucherResponse) SetRecurringEndDate(v time.Time) {
	o.RecurringEndDate.Set(&v)
}
// SetRecurringEndDateNil sets the value for RecurringEndDate to be an explicit nil
func (o *ModelVoucherResponse) SetRecurringEndDateNil() {
	o.RecurringEndDate.Set(nil)
}

// UnsetRecurringEndDate ensures that no value is present for RecurringEndDate, not even an explicit nil
func (o *ModelVoucherResponse) UnsetRecurringEndDate() {
	o.RecurringEndDate.Unset()
}

// GetEnshrined returns the Enshrined field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelVoucherResponse) GetEnshrined() time.Time {
	if o == nil || IsNil(o.Enshrined.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Enshrined.Get()
}

// GetEnshrinedOk returns a tuple with the Enshrined field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelVoucherResponse) GetEnshrinedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enshrined.Get(), o.Enshrined.IsSet()
}

// HasEnshrined returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasEnshrined() bool {
	if o != nil && o.Enshrined.IsSet() {
		return true
	}

	return false
}

// SetEnshrined gets a reference to the given NullableTime and assigns it to the Enshrined field.
func (o *ModelVoucherResponse) SetEnshrined(v time.Time) {
	o.Enshrined.Set(&v)
}
// SetEnshrinedNil sets the value for Enshrined to be an explicit nil
func (o *ModelVoucherResponse) SetEnshrinedNil() {
	o.Enshrined.Set(nil)
}

// UnsetEnshrined ensures that no value is present for Enshrined, not even an explicit nil
func (o *ModelVoucherResponse) UnsetEnshrined() {
	o.Enshrined.Unset()
}

// GetTaxSet returns the TaxSet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelVoucherResponse) GetTaxSet() ModelVoucherResponseTaxSet {
	if o == nil || IsNil(o.TaxSet.Get()) {
		var ret ModelVoucherResponseTaxSet
		return ret
	}
	return *o.TaxSet.Get()
}

// GetTaxSetOk returns a tuple with the TaxSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelVoucherResponse) GetTaxSetOk() (*ModelVoucherResponseTaxSet, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxSet.Get(), o.TaxSet.IsSet()
}

// HasTaxSet returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasTaxSet() bool {
	if o != nil && o.TaxSet.IsSet() {
		return true
	}

	return false
}

// SetTaxSet gets a reference to the given NullableModelVoucherResponseTaxSet and assigns it to the TaxSet field.
func (o *ModelVoucherResponse) SetTaxSet(v ModelVoucherResponseTaxSet) {
	o.TaxSet.Set(&v)
}
// SetTaxSetNil sets the value for TaxSet to be an explicit nil
func (o *ModelVoucherResponse) SetTaxSetNil() {
	o.TaxSet.Set(nil)
}

// UnsetTaxSet ensures that no value is present for TaxSet, not even an explicit nil
func (o *ModelVoucherResponse) UnsetTaxSet() {
	o.TaxSet.Unset()
}

// GetPaymentDeadline returns the PaymentDeadline field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelVoucherResponse) GetPaymentDeadline() time.Time {
	if o == nil || IsNil(o.PaymentDeadline.Get()) {
		var ret time.Time
		return ret
	}
	return *o.PaymentDeadline.Get()
}

// GetPaymentDeadlineOk returns a tuple with the PaymentDeadline field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelVoucherResponse) GetPaymentDeadlineOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentDeadline.Get(), o.PaymentDeadline.IsSet()
}

// HasPaymentDeadline returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasPaymentDeadline() bool {
	if o != nil && o.PaymentDeadline.IsSet() {
		return true
	}

	return false
}

// SetPaymentDeadline gets a reference to the given NullableTime and assigns it to the PaymentDeadline field.
func (o *ModelVoucherResponse) SetPaymentDeadline(v time.Time) {
	o.PaymentDeadline.Set(&v)
}
// SetPaymentDeadlineNil sets the value for PaymentDeadline to be an explicit nil
func (o *ModelVoucherResponse) SetPaymentDeadlineNil() {
	o.PaymentDeadline.Set(nil)
}

// UnsetPaymentDeadline ensures that no value is present for PaymentDeadline, not even an explicit nil
func (o *ModelVoucherResponse) UnsetPaymentDeadline() {
	o.PaymentDeadline.Unset()
}

// GetDeliveryDate returns the DeliveryDate field value if set, zero value otherwise.
func (o *ModelVoucherResponse) GetDeliveryDate() time.Time {
	if o == nil || IsNil(o.DeliveryDate) {
		var ret time.Time
		return ret
	}
	return *o.DeliveryDate
}

// GetDeliveryDateOk returns a tuple with the DeliveryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelVoucherResponse) GetDeliveryDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeliveryDate) {
		return nil, false
	}
	return o.DeliveryDate, true
}

// HasDeliveryDate returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasDeliveryDate() bool {
	if o != nil && !IsNil(o.DeliveryDate) {
		return true
	}

	return false
}

// SetDeliveryDate gets a reference to the given time.Time and assigns it to the DeliveryDate field.
func (o *ModelVoucherResponse) SetDeliveryDate(v time.Time) {
	o.DeliveryDate = &v
}

// GetDeliveryDateUntil returns the DeliveryDateUntil field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelVoucherResponse) GetDeliveryDateUntil() time.Time {
	if o == nil || IsNil(o.DeliveryDateUntil.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DeliveryDateUntil.Get()
}

// GetDeliveryDateUntilOk returns a tuple with the DeliveryDateUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelVoucherResponse) GetDeliveryDateUntilOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeliveryDateUntil.Get(), o.DeliveryDateUntil.IsSet()
}

// HasDeliveryDateUntil returns a boolean if a field has been set.
func (o *ModelVoucherResponse) HasDeliveryDateUntil() bool {
	if o != nil && o.DeliveryDateUntil.IsSet() {
		return true
	}

	return false
}

// SetDeliveryDateUntil gets a reference to the given NullableTime and assigns it to the DeliveryDateUntil field.
func (o *ModelVoucherResponse) SetDeliveryDateUntil(v time.Time) {
	o.DeliveryDateUntil.Set(&v)
}
// SetDeliveryDateUntilNil sets the value for DeliveryDateUntil to be an explicit nil
func (o *ModelVoucherResponse) SetDeliveryDateUntilNil() {
	o.DeliveryDateUntil.Set(nil)
}

// UnsetDeliveryDateUntil ensures that no value is present for DeliveryDateUntil, not even an explicit nil
func (o *ModelVoucherResponse) UnsetDeliveryDateUntil() {
	o.DeliveryDateUntil.Unset()
}

func (o ModelVoucherResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelVoucherResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	// skip: objectName is readOnly
	if !IsNil(o.MapAll) {
		toSerialize["mapAll"] = o.MapAll
	}
	// skip: create is readOnly
	// skip: update is readOnly
	if !IsNil(o.SevClient) {
		toSerialize["sevClient"] = o.SevClient
	}
	if !IsNil(o.CreateUser) {
		toSerialize["createUser"] = o.CreateUser
	}
	if o.VoucherDate.IsSet() {
		toSerialize["voucherDate"] = o.VoucherDate.Get()
	}
	if o.Supplier.IsSet() {
		toSerialize["supplier"] = o.Supplier.Get()
	}
	if o.SupplierName.IsSet() {
		toSerialize["supplierName"] = o.SupplierName.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Document.IsSet() {
		toSerialize["document"] = o.Document.Get()
	}
	if o.PayDate.IsSet() {
		toSerialize["payDate"] = o.PayDate.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	// skip: sumNet is readOnly
	// skip: sumTax is readOnly
	// skip: sumGross is readOnly
	// skip: sumNetAccounting is readOnly
	// skip: sumTaxAccounting is readOnly
	// skip: sumGrossAccounting is readOnly
	// skip: sumDiscounts is readOnly
	// skip: sumDiscountsForeignCurrency is readOnly
	if o.PaidAmount.IsSet() {
		toSerialize["paidAmount"] = o.PaidAmount.Get()
	}
	if o.TaxType.IsSet() {
		toSerialize["taxType"] = o.TaxType.Get()
	}
	if o.CreditDebit.IsSet() {
		toSerialize["creditDebit"] = o.CreditDebit.Get()
	}
	if !IsNil(o.CostCentre) {
		toSerialize["costCentre"] = o.CostCentre
	}
	if o.VoucherType.IsSet() {
		toSerialize["voucherType"] = o.VoucherType.Get()
	}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	if o.PropertyForeignCurrencyDeadline.IsSet() {
		toSerialize["propertyForeignCurrencyDeadline"] = o.PropertyForeignCurrencyDeadline.Get()
	}
	if o.PropertyExchangeRate.IsSet() {
		toSerialize["propertyExchangeRate"] = o.PropertyExchangeRate.Get()
	}
	if o.RecurringInterval.IsSet() {
		toSerialize["recurringInterval"] = o.RecurringInterval.Get()
	}
	if o.RecurringStartDate.IsSet() {
		toSerialize["recurringStartDate"] = o.RecurringStartDate.Get()
	}
	if o.RecurringNextVoucher.IsSet() {
		toSerialize["recurringNextVoucher"] = o.RecurringNextVoucher.Get()
	}
	if o.RecurringLastVoucher.IsSet() {
		toSerialize["recurringLastVoucher"] = o.RecurringLastVoucher.Get()
	}
	if o.RecurringEndDate.IsSet() {
		toSerialize["recurringEndDate"] = o.RecurringEndDate.Get()
	}
	if o.Enshrined.IsSet() {
		toSerialize["enshrined"] = o.Enshrined.Get()
	}
	if o.TaxSet.IsSet() {
		toSerialize["taxSet"] = o.TaxSet.Get()
	}
	if o.PaymentDeadline.IsSet() {
		toSerialize["paymentDeadline"] = o.PaymentDeadline.Get()
	}
	if !IsNil(o.DeliveryDate) {
		toSerialize["deliveryDate"] = o.DeliveryDate
	}
	if o.DeliveryDateUntil.IsSet() {
		toSerialize["deliveryDateUntil"] = o.DeliveryDateUntil.Get()
	}
	return toSerialize, nil
}

type NullableModelVoucherResponse struct {
	value *ModelVoucherResponse
	isSet bool
}

func (v NullableModelVoucherResponse) Get() *ModelVoucherResponse {
	return v.value
}

func (v *NullableModelVoucherResponse) Set(val *ModelVoucherResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModelVoucherResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModelVoucherResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelVoucherResponse(val *ModelVoucherResponse) *NullableModelVoucherResponse {
	return &NullableModelVoucherResponse{value: val, isSet: true}
}

func (v NullableModelVoucherResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelVoucherResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


