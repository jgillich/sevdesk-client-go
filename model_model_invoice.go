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

// checks if the ModelInvoice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelInvoice{}

// ModelInvoice Invoice model
type ModelInvoice struct {
	// The invoice id. <span style='color:red'>Required</span> if you want to create or update an invoice position for an existing invoice
	Id NullableInt32 `json:"id,omitempty"`
	// The invoice object name.
	ObjectName *string `json:"objectName,omitempty"`
	// The invoice number
	InvoiceNumber NullableString `json:"invoiceNumber,omitempty"`
	Contact ModelInvoiceContact `json:"contact"`
	ContactPerson ModelInvoiceUpdateContactPerson `json:"contactPerson"`
	// Date of invoice creation
	Create *time.Time `json:"create,omitempty"`
	// Date of last invoice update
	Update *time.Time `json:"update,omitempty"`
	SevClient *ModelInvoiceUpdateSevClient `json:"sevClient,omitempty"`
	// Needs to be provided as timestamp or dd.mm.yyyy
	InvoiceDate string `json:"invoiceDate"`
	// Normally consist of prefix plus the invoice number
	Header NullableString `json:"header,omitempty"`
	// Certain html tags can be used here to format your text
	HeadText NullableString `json:"headText,omitempty"`
	// Certain html tags can be used here to format your text
	FootText NullableString `json:"footText,omitempty"`
	// The time the customer has to pay the invoice in days
	TimeToPay NullableInt32 `json:"timeToPay,omitempty"`
	// If you want to give a discount, define the percentage here. Otherwise provide zero as value
	Discount int32 `json:"discount"`
	// Complete address of the recipient including name, street, city, zip and country.       * Line breaks can be used and will be displayed on the invoice pdf.
	Address NullableString `json:"address,omitempty"`
	AddressCountry ModelInvoiceAddressCountry `json:"addressCountry"`
	// Needs to be timestamp or dd.mm.yyyy
	PayDate NullableTime `json:"payDate,omitempty"`
	CreateUser *ModelCreditNoteCreateUser `json:"createUser,omitempty"`
	// Timestamp. This can also be a date range if you also use the attribute deliveryDateUntil
	DeliveryDate NullableTime `json:"deliveryDate,omitempty"`
	// If the delivery date should be a time range, another timestamp can be provided in this attribute       * to define a range from timestamp used in deliveryDate attribute to the timestamp used here.
	DeliveryDateUntil NullableInt32 `json:"deliveryDateUntil,omitempty"`
	// Please have a look in our       <a href='https://api.sevdesk.de/#section/Types-and-status-of-invoices'>Types and status of invoices</a>       to see what the different status codes mean
	Status string `json:"status"`
	// Defines if the client uses the small settlement scheme.      If yes, the invoice must not contain any vat
	SmallSettlement NullableBool `json:"smallSettlement,omitempty"`
	// Is overwritten by invoice position tax rates
	TaxRate float32 `json:"taxRate"`
	// A common tax text would be 'Umsatzsteuer 19%'
	TaxText string `json:"taxText"`
	// Tax type of the invoice. There are four tax types: 1. default - Umsatzsteuer ausweisen 2. eu - Steuerfreie innergemeinschaftliche Lieferung (Europäische Union) 3. noteu - Steuerschuldnerschaft des Leistungsempfängers (außerhalb EU, z. B. Schweiz) 4. custom - Using custom tax set 5. ss - Not subject to VAT according to §19 1 UStG Tax rates are heavily connected to the tax type used.
	TaxType string `json:"taxType"`
	TaxSet NullableModelInvoiceTaxSet `json:"taxSet,omitempty"`
	// Defines how many reminders have already been sent for the invoice.      Starts with 1 (Payment reminder) and should be incremented by one every time another reminder is sent.
	DunningLevel NullableInt32 `json:"dunningLevel,omitempty"`
	PaymentMethod *ModelInvoicePaymentMethod `json:"paymentMethod,omitempty"`
	// The date the invoice was sent to the customer
	SendDate NullableTime `json:"sendDate,omitempty"`
	// Type of the invoice. For more information on the different types, check       <a href='https://api.sevdesk.de/#section/Types-and-status-of-invoices'>this</a> section  
	InvoiceType string `json:"invoiceType"`
	// The interval in which recurring invoices are due as ISO-8601 duration.<br>       Necessary attribute for all recurring invoices.
	AccountIntervall NullableString `json:"accountIntervall,omitempty"`
	// Timestamp when the next invoice will be generated by this recurring invoice.
	AccountNextInvoice NullableInt32 `json:"accountNextInvoice,omitempty"`
	// Currency used in the invoice. Needs to be currency code according to ISO-4217
	Currency string `json:"currency"`
	// Net sum of the invoice
	SumNet *float32 `json:"sumNet,omitempty"`
	// Tax sum of the invoice
	SumTax *float32 `json:"sumTax,omitempty"`
	// Gross sum of the invoice
	SumGross *float32 `json:"sumGross,omitempty"`
	// Sum of all discounts in the invoice
	SumDiscounts *float32 `json:"sumDiscounts,omitempty"`
	// Net sum of the invoice in the foreign currency
	SumNetForeignCurrency *float32 `json:"sumNetForeignCurrency,omitempty"`
	// Tax sum of the invoice in the foreign currency
	SumTaxForeignCurrency *float32 `json:"sumTaxForeignCurrency,omitempty"`
	// Gross sum of the invoice in the foreign currency
	SumGrossForeignCurrency *float32 `json:"sumGrossForeignCurrency,omitempty"`
	// Discounts sum of the invoice in the foreign currency
	SumDiscountsForeignCurrency *float32 `json:"sumDiscountsForeignCurrency,omitempty"`
	// Net accounting sum of the invoice. Is usually the same as sumNet
	SumNetAccounting *float32 `json:"sumNetAccounting,omitempty"`
	// Tax accounting sum of the invoice. Is usually the same as sumTax
	SumTaxAccounting *float32 `json:"sumTaxAccounting,omitempty"`
	// Gross accounting sum of the invoice. Is usually the same as sumGross
	SumGrossAccounting *float32 `json:"sumGrossAccounting,omitempty"`
	// Amount which has already been paid for this invoice by the customer
	PaidAmount NullableFloat32 `json:"paidAmount,omitempty"`
	// If true, the net amount of each position will be shown on the invoice. Otherwise gross amount
	ShowNet *bool `json:"showNet,omitempty"`
	// Defines if and when invoice was enshrined. Enshrined invoices can not be manipulated.
	Enshrined NullableTime `json:"enshrined,omitempty"`
	// Type which was used to send the invoice.
	SendType NullableString `json:"sendType,omitempty"`
	Origin NullableModelInvoiceOrigin `json:"origin,omitempty"`
	// Internal note of the customer. Contains data entered into field 'Referenz/Bestellnummer'
	CustomerInternalNote NullableString `json:"customerInternalNote,omitempty"`
	MapAll bool `json:"mapAll"`
}

// NewModelInvoice instantiates a new ModelInvoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelInvoice(contact ModelInvoiceContact, contactPerson ModelInvoiceUpdateContactPerson, invoiceDate string, discount int32, addressCountry ModelInvoiceAddressCountry, status string, taxRate float32, taxText string, taxType string, invoiceType string, currency string, mapAll bool) *ModelInvoice {
	this := ModelInvoice{}
	this.Contact = contact
	this.ContactPerson = contactPerson
	this.InvoiceDate = invoiceDate
	this.Discount = discount
	this.AddressCountry = addressCountry
	this.Status = status
	this.TaxRate = taxRate
	this.TaxText = taxText
	this.TaxType = taxType
	this.InvoiceType = invoiceType
	this.Currency = currency
	this.MapAll = mapAll
	return &this
}

// NewModelInvoiceWithDefaults instantiates a new ModelInvoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelInvoiceWithDefaults() *ModelInvoice {
	this := ModelInvoice{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoice) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoice) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ModelInvoice) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *ModelInvoice) SetId(v int32) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ModelInvoice) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ModelInvoice) UnsetId() {
	o.Id.Unset()
}

// GetObjectName returns the ObjectName field value if set, zero value otherwise.
func (o *ModelInvoice) GetObjectName() string {
	if o == nil || IsNil(o.ObjectName) {
		var ret string
		return ret
	}
	return *o.ObjectName
}

// GetObjectNameOk returns a tuple with the ObjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoice) GetObjectNameOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectName) {
		return nil, false
	}
	return o.ObjectName, true
}

// HasObjectName returns a boolean if a field has been set.
func (o *ModelInvoice) HasObjectName() bool {
	if o != nil && !IsNil(o.ObjectName) {
		return true
	}

	return false
}

// SetObjectName gets a reference to the given string and assigns it to the ObjectName field.
func (o *ModelInvoice) SetObjectName(v string) {
	o.ObjectName = &v
}

// GetInvoiceNumber returns the InvoiceNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoice) GetInvoiceNumber() string {
	if o == nil || IsNil(o.InvoiceNumber.Get()) {
		var ret string
		return ret
	}
	return *o.InvoiceNumber.Get()
}

// GetInvoiceNumberOk returns a tuple with the InvoiceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoice) GetInvoiceNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceNumber.Get(), o.InvoiceNumber.IsSet()
}

// HasInvoiceNumber returns a boolean if a field has been set.
func (o *ModelInvoice) HasInvoiceNumber() bool {
	if o != nil && o.InvoiceNumber.IsSet() {
		return true
	}

	return false
}

// SetInvoiceNumber gets a reference to the given NullableString and assigns it to the InvoiceNumber field.
func (o *ModelInvoice) SetInvoiceNumber(v string) {
	o.InvoiceNumber.Set(&v)
}
// SetInvoiceNumberNil sets the value for InvoiceNumber to be an explicit nil
func (o *ModelInvoice) SetInvoiceNumberNil() {
	o.InvoiceNumber.Set(nil)
}

// UnsetInvoiceNumber ensures that no value is present for InvoiceNumber, not even an explicit nil
func (o *ModelInvoice) UnsetInvoiceNumber() {
	o.InvoiceNumber.Unset()
}

// GetContact returns the Contact field value
func (o *ModelInvoice) GetContact() ModelInvoiceContact {
	if o == nil {
		var ret ModelInvoiceContact
		return ret
	}

	return o.Contact
}

// GetContactOk returns a tuple with the Contact field value
// and a boolean to check if the value has been set.
func (o *ModelInvoice) GetContactOk() (*ModelInvoiceContact, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contact, true
}

// SetContact sets field value
func (o *ModelInvoice) SetContact(v ModelInvoiceContact) {
	o.Contact = v
}

// GetContactPerson returns the ContactPerson field value
func (o *ModelInvoice) GetContactPerson() ModelInvoiceUpdateContactPerson {
	if o == nil {
		var ret ModelInvoiceUpdateContactPerson
		return ret
	}

	return o.ContactPerson
}

// GetContactPersonOk returns a tuple with the ContactPerson field value
// and a boolean to check if the value has been set.
func (o *ModelInvoice) GetContactPersonOk() (*ModelInvoiceUpdateContactPerson, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContactPerson, true
}

// SetContactPerson sets field value
func (o *ModelInvoice) SetContactPerson(v ModelInvoiceUpdateContactPerson) {
	o.ContactPerson = v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *ModelInvoice) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoice) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *ModelInvoice) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *ModelInvoice) SetCreate(v time.Time) {
	o.Create = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *ModelInvoice) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoice) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *ModelInvoice) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *ModelInvoice) SetUpdate(v time.Time) {
	o.Update = &v
}

// GetSevClient returns the SevClient field value if set, zero value otherwise.
func (o *ModelInvoice) GetSevClient() ModelInvoiceUpdateSevClient {
	if o == nil || IsNil(o.SevClient) {
		var ret ModelInvoiceUpdateSevClient
		return ret
	}
	return *o.SevClient
}

// GetSevClientOk returns a tuple with the SevClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoice) GetSevClientOk() (*ModelInvoiceUpdateSevClient, bool) {
	if o == nil || IsNil(o.SevClient) {
		return nil, false
	}
	return o.SevClient, true
}

// HasSevClient returns a boolean if a field has been set.
func (o *ModelInvoice) HasSevClient() bool {
	if o != nil && !IsNil(o.SevClient) {
		return true
	}

	return false
}

// SetSevClient gets a reference to the given ModelInvoiceUpdateSevClient and assigns it to the SevClient field.
func (o *ModelInvoice) SetSevClient(v ModelInvoiceUpdateSevClient) {
	o.SevClient = &v
}

// GetInvoiceDate returns the InvoiceDate field value
func (o *ModelInvoice) GetInvoiceDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceDate
}

// GetInvoiceDateOk returns a tuple with the InvoiceDate field value
// and a boolean to check if the value has been set.
func (o *ModelInvoice) GetInvoiceDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceDate, true
}

// SetInvoiceDate sets field value
func (o *ModelInvoice) SetInvoiceDate(v string) {
	o.InvoiceDate = v
}

// GetHeader returns the Header field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoice) GetHeader() string {
	if o == nil || IsNil(o.Header.Get()) {
		var ret string
		return ret
	}
	return *o.Header.Get()
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoice) GetHeaderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Header.Get(), o.Header.IsSet()
}

// HasHeader returns a boolean if a field has been set.
func (o *ModelInvoice) HasHeader() bool {
	if o != nil && o.Header.IsSet() {
		return true
	}

	return false
}

// SetHeader gets a reference to the given NullableString and assigns it to the Header field.
func (o *ModelInvoice) SetHeader(v string) {
	o.Header.Set(&v)
}
// SetHeaderNil sets the value for Header to be an explicit nil
func (o *ModelInvoice) SetHeaderNil() {
	o.Header.Set(nil)
}

// UnsetHeader ensures that no value is present for Header, not even an explicit nil
func (o *ModelInvoice) UnsetHeader() {
	o.Header.Unset()
}

// GetHeadText returns the HeadText field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoice) GetHeadText() string {
	if o == nil || IsNil(o.HeadText.Get()) {
		var ret string
		return ret
	}
	return *o.HeadText.Get()
}

// GetHeadTextOk returns a tuple with the HeadText field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoice) GetHeadTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HeadText.Get(), o.HeadText.IsSet()
}

// HasHeadText returns a boolean if a field has been set.
func (o *ModelInvoice) HasHeadText() bool {
	if o != nil && o.HeadText.IsSet() {
		return true
	}

	return false
}

// SetHeadText gets a reference to the given NullableString and assigns it to the HeadText field.
func (o *ModelInvoice) SetHeadText(v string) {
	o.HeadText.Set(&v)
}
// SetHeadTextNil sets the value for HeadText to be an explicit nil
func (o *ModelInvoice) SetHeadTextNil() {
	o.HeadText.Set(nil)
}

// UnsetHeadText ensures that no value is present for HeadText, not even an explicit nil
func (o *ModelInvoice) UnsetHeadText() {
	o.HeadText.Unset()
}

// GetFootText returns the FootText field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoice) GetFootText() string {
	if o == nil || IsNil(o.FootText.Get()) {
		var ret string
		return ret
	}
	return *o.FootText.Get()
}

// GetFootTextOk returns a tuple with the FootText field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoice) GetFootTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FootText.Get(), o.FootText.IsSet()
}

// HasFootText returns a boolean if a field has been set.
func (o *ModelInvoice) HasFootText() bool {
	if o != nil && o.FootText.IsSet() {
		return true
	}

	return false
}

// SetFootText gets a reference to the given NullableString and assigns it to the FootText field.
func (o *ModelInvoice) SetFootText(v string) {
	o.FootText.Set(&v)
}
// SetFootTextNil sets the value for FootText to be an explicit nil
func (o *ModelInvoice) SetFootTextNil() {
	o.FootText.Set(nil)
}

// UnsetFootText ensures that no value is present for FootText, not even an explicit nil
func (o *ModelInvoice) UnsetFootText() {
	o.FootText.Unset()
}

// GetTimeToPay returns the TimeToPay field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoice) GetTimeToPay() int32 {
	if o == nil || IsNil(o.TimeToPay.Get()) {
		var ret int32
		return ret
	}
	return *o.TimeToPay.Get()
}

// GetTimeToPayOk returns a tuple with the TimeToPay field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoice) GetTimeToPayOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimeToPay.Get(), o.TimeToPay.IsSet()
}

// HasTimeToPay returns a boolean if a field has been set.
func (o *ModelInvoice) HasTimeToPay() bool {
	if o != nil && o.TimeToPay.IsSet() {
		return true
	}

	return false
}

// SetTimeToPay gets a reference to the given NullableInt32 and assigns it to the TimeToPay field.
func (o *ModelInvoice) SetTimeToPay(v int32) {
	o.TimeToPay.Set(&v)
}
// SetTimeToPayNil sets the value for TimeToPay to be an explicit nil
func (o *ModelInvoice) SetTimeToPayNil() {
	o.TimeToPay.Set(nil)
}

// UnsetTimeToPay ensures that no value is present for TimeToPay, not even an explicit nil
func (o *ModelInvoice) UnsetTimeToPay() {
	o.TimeToPay.Unset()
}

// GetDiscount returns the Discount field value
func (o *ModelInvoice) GetDiscount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value
// and a boolean to check if the value has been set.
func (o *ModelInvoice) GetDiscountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Discount, true
}

// SetDiscount sets field value
func (o *ModelInvoice) SetDiscount(v int32) {
	o.Discount = v
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoice) GetAddress() string {
	if o == nil || IsNil(o.Address.Get()) {
		var ret string
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoice) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *ModelInvoice) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableString and assigns it to the Address field.
func (o *ModelInvoice) SetAddress(v string) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *ModelInvoice) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *ModelInvoice) UnsetAddress() {
	o.Address.Unset()
}

// GetAddressCountry returns the AddressCountry field value
func (o *ModelInvoice) GetAddressCountry() ModelInvoiceAddressCountry {
	if o == nil {
		var ret ModelInvoiceAddressCountry
		return ret
	}

	return o.AddressCountry
}

// GetAddressCountryOk returns a tuple with the AddressCountry field value
// and a boolean to check if the value has been set.
func (o *ModelInvoice) GetAddressCountryOk() (*ModelInvoiceAddressCountry, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressCountry, true
}

// SetAddressCountry sets field value
func (o *ModelInvoice) SetAddressCountry(v ModelInvoiceAddressCountry) {
	o.AddressCountry = v
}

// GetPayDate returns the PayDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoice) GetPayDate() time.Time {
	if o == nil || IsNil(o.PayDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.PayDate.Get()
}

// GetPayDateOk returns a tuple with the PayDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoice) GetPayDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.PayDate.Get(), o.PayDate.IsSet()
}

// HasPayDate returns a boolean if a field has been set.
func (o *ModelInvoice) HasPayDate() bool {
	if o != nil && o.PayDate.IsSet() {
		return true
	}

	return false
}

// SetPayDate gets a reference to the given NullableTime and assigns it to the PayDate field.
func (o *ModelInvoice) SetPayDate(v time.Time) {
	o.PayDate.Set(&v)
}
// SetPayDateNil sets the value for PayDate to be an explicit nil
func (o *ModelInvoice) SetPayDateNil() {
	o.PayDate.Set(nil)
}

// UnsetPayDate ensures that no value is present for PayDate, not even an explicit nil
func (o *ModelInvoice) UnsetPayDate() {
	o.PayDate.Unset()
}

// GetCreateUser returns the CreateUser field value if set, zero value otherwise.
func (o *ModelInvoice) GetCreateUser() ModelCreditNoteCreateUser {
	if o == nil || IsNil(o.CreateUser) {
		var ret ModelCreditNoteCreateUser
		return ret
	}
	return *o.CreateUser
}

// GetCreateUserOk returns a tuple with the CreateUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoice) GetCreateUserOk() (*ModelCreditNoteCreateUser, bool) {
	if o == nil || IsNil(o.CreateUser) {
		return nil, false
	}
	return o.CreateUser, true
}

// HasCreateUser returns a boolean if a field has been set.
func (o *ModelInvoice) HasCreateUser() bool {
	if o != nil && !IsNil(o.CreateUser) {
		return true
	}

	return false
}

// SetCreateUser gets a reference to the given ModelCreditNoteCreateUser and assigns it to the CreateUser field.
func (o *ModelInvoice) SetCreateUser(v ModelCreditNoteCreateUser) {
	o.CreateUser = &v
}

// GetDeliveryDate returns the DeliveryDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoice) GetDeliveryDate() time.Time {
	if o == nil || IsNil(o.DeliveryDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DeliveryDate.Get()
}

// GetDeliveryDateOk returns a tuple with the DeliveryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoice) GetDeliveryDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeliveryDate.Get(), o.DeliveryDate.IsSet()
}

// HasDeliveryDate returns a boolean if a field has been set.
func (o *ModelInvoice) HasDeliveryDate() bool {
	if o != nil && o.DeliveryDate.IsSet() {
		return true
	}

	return false
}

// SetDeliveryDate gets a reference to the given NullableTime and assigns it to the DeliveryDate field.
func (o *ModelInvoice) SetDeliveryDate(v time.Time) {
	o.DeliveryDate.Set(&v)
}
// SetDeliveryDateNil sets the value for DeliveryDate to be an explicit nil
func (o *ModelInvoice) SetDeliveryDateNil() {
	o.DeliveryDate.Set(nil)
}

// UnsetDeliveryDate ensures that no value is present for DeliveryDate, not even an explicit nil
func (o *ModelInvoice) UnsetDeliveryDate() {
	o.DeliveryDate.Unset()
}

// GetDeliveryDateUntil returns the DeliveryDateUntil field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoice) GetDeliveryDateUntil() int32 {
	if o == nil || IsNil(o.DeliveryDateUntil.Get()) {
		var ret int32
		return ret
	}
	return *o.DeliveryDateUntil.Get()
}

// GetDeliveryDateUntilOk returns a tuple with the DeliveryDateUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoice) GetDeliveryDateUntilOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeliveryDateUntil.Get(), o.DeliveryDateUntil.IsSet()
}

// HasDeliveryDateUntil returns a boolean if a field has been set.
func (o *ModelInvoice) HasDeliveryDateUntil() bool {
	if o != nil && o.DeliveryDateUntil.IsSet() {
		return true
	}

	return false
}

// SetDeliveryDateUntil gets a reference to the given NullableInt32 and assigns it to the DeliveryDateUntil field.
func (o *ModelInvoice) SetDeliveryDateUntil(v int32) {
	o.DeliveryDateUntil.Set(&v)
}
// SetDeliveryDateUntilNil sets the value for DeliveryDateUntil to be an explicit nil
func (o *ModelInvoice) SetDeliveryDateUntilNil() {
	o.DeliveryDateUntil.Set(nil)
}

// UnsetDeliveryDateUntil ensures that no value is present for DeliveryDateUntil, not even an explicit nil
func (o *ModelInvoice) UnsetDeliveryDateUntil() {
	o.DeliveryDateUntil.Unset()
}

// GetStatus returns the Status field value
func (o *ModelInvoice) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ModelInvoice) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ModelInvoice) SetStatus(v string) {
	o.Status = v
}

// GetSmallSettlement returns the SmallSettlement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoice) GetSmallSettlement() bool {
	if o == nil || IsNil(o.SmallSettlement.Get()) {
		var ret bool
		return ret
	}
	return *o.SmallSettlement.Get()
}

// GetSmallSettlementOk returns a tuple with the SmallSettlement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoice) GetSmallSettlementOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SmallSettlement.Get(), o.SmallSettlement.IsSet()
}

// HasSmallSettlement returns a boolean if a field has been set.
func (o *ModelInvoice) HasSmallSettlement() bool {
	if o != nil && o.SmallSettlement.IsSet() {
		return true
	}

	return false
}

// SetSmallSettlement gets a reference to the given NullableBool and assigns it to the SmallSettlement field.
func (o *ModelInvoice) SetSmallSettlement(v bool) {
	o.SmallSettlement.Set(&v)
}
// SetSmallSettlementNil sets the value for SmallSettlement to be an explicit nil
func (o *ModelInvoice) SetSmallSettlementNil() {
	o.SmallSettlement.Set(nil)
}

// UnsetSmallSettlement ensures that no value is present for SmallSettlement, not even an explicit nil
func (o *ModelInvoice) UnsetSmallSettlement() {
	o.SmallSettlement.Unset()
}

// GetTaxRate returns the TaxRate field value
func (o *ModelInvoice) GetTaxRate() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TaxRate
}

// GetTaxRateOk returns a tuple with the TaxRate field value
// and a boolean to check if the value has been set.
func (o *ModelInvoice) GetTaxRateOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaxRate, true
}

// SetTaxRate sets field value
func (o *ModelInvoice) SetTaxRate(v float32) {
	o.TaxRate = v
}

// GetTaxText returns the TaxText field value
func (o *ModelInvoice) GetTaxText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TaxText
}

// GetTaxTextOk returns a tuple with the TaxText field value
// and a boolean to check if the value has been set.
func (o *ModelInvoice) GetTaxTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaxText, true
}

// SetTaxText sets field value
func (o *ModelInvoice) SetTaxText(v string) {
	o.TaxText = v
}

// GetTaxType returns the TaxType field value
func (o *ModelInvoice) GetTaxType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TaxType
}

// GetTaxTypeOk returns a tuple with the TaxType field value
// and a boolean to check if the value has been set.
func (o *ModelInvoice) GetTaxTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaxType, true
}

// SetTaxType sets field value
func (o *ModelInvoice) SetTaxType(v string) {
	o.TaxType = v
}

// GetTaxSet returns the TaxSet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoice) GetTaxSet() ModelInvoiceTaxSet {
	if o == nil || IsNil(o.TaxSet.Get()) {
		var ret ModelInvoiceTaxSet
		return ret
	}
	return *o.TaxSet.Get()
}

// GetTaxSetOk returns a tuple with the TaxSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoice) GetTaxSetOk() (*ModelInvoiceTaxSet, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxSet.Get(), o.TaxSet.IsSet()
}

// HasTaxSet returns a boolean if a field has been set.
func (o *ModelInvoice) HasTaxSet() bool {
	if o != nil && o.TaxSet.IsSet() {
		return true
	}

	return false
}

// SetTaxSet gets a reference to the given NullableModelInvoiceTaxSet and assigns it to the TaxSet field.
func (o *ModelInvoice) SetTaxSet(v ModelInvoiceTaxSet) {
	o.TaxSet.Set(&v)
}
// SetTaxSetNil sets the value for TaxSet to be an explicit nil
func (o *ModelInvoice) SetTaxSetNil() {
	o.TaxSet.Set(nil)
}

// UnsetTaxSet ensures that no value is present for TaxSet, not even an explicit nil
func (o *ModelInvoice) UnsetTaxSet() {
	o.TaxSet.Unset()
}

// GetDunningLevel returns the DunningLevel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoice) GetDunningLevel() int32 {
	if o == nil || IsNil(o.DunningLevel.Get()) {
		var ret int32
		return ret
	}
	return *o.DunningLevel.Get()
}

// GetDunningLevelOk returns a tuple with the DunningLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoice) GetDunningLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DunningLevel.Get(), o.DunningLevel.IsSet()
}

// HasDunningLevel returns a boolean if a field has been set.
func (o *ModelInvoice) HasDunningLevel() bool {
	if o != nil && o.DunningLevel.IsSet() {
		return true
	}

	return false
}

// SetDunningLevel gets a reference to the given NullableInt32 and assigns it to the DunningLevel field.
func (o *ModelInvoice) SetDunningLevel(v int32) {
	o.DunningLevel.Set(&v)
}
// SetDunningLevelNil sets the value for DunningLevel to be an explicit nil
func (o *ModelInvoice) SetDunningLevelNil() {
	o.DunningLevel.Set(nil)
}

// UnsetDunningLevel ensures that no value is present for DunningLevel, not even an explicit nil
func (o *ModelInvoice) UnsetDunningLevel() {
	o.DunningLevel.Unset()
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *ModelInvoice) GetPaymentMethod() ModelInvoicePaymentMethod {
	if o == nil || IsNil(o.PaymentMethod) {
		var ret ModelInvoicePaymentMethod
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoice) GetPaymentMethodOk() (*ModelInvoicePaymentMethod, bool) {
	if o == nil || IsNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *ModelInvoice) HasPaymentMethod() bool {
	if o != nil && !IsNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given ModelInvoicePaymentMethod and assigns it to the PaymentMethod field.
func (o *ModelInvoice) SetPaymentMethod(v ModelInvoicePaymentMethod) {
	o.PaymentMethod = &v
}

// GetSendDate returns the SendDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoice) GetSendDate() time.Time {
	if o == nil || IsNil(o.SendDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.SendDate.Get()
}

// GetSendDateOk returns a tuple with the SendDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoice) GetSendDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.SendDate.Get(), o.SendDate.IsSet()
}

// HasSendDate returns a boolean if a field has been set.
func (o *ModelInvoice) HasSendDate() bool {
	if o != nil && o.SendDate.IsSet() {
		return true
	}

	return false
}

// SetSendDate gets a reference to the given NullableTime and assigns it to the SendDate field.
func (o *ModelInvoice) SetSendDate(v time.Time) {
	o.SendDate.Set(&v)
}
// SetSendDateNil sets the value for SendDate to be an explicit nil
func (o *ModelInvoice) SetSendDateNil() {
	o.SendDate.Set(nil)
}

// UnsetSendDate ensures that no value is present for SendDate, not even an explicit nil
func (o *ModelInvoice) UnsetSendDate() {
	o.SendDate.Unset()
}

// GetInvoiceType returns the InvoiceType field value
func (o *ModelInvoice) GetInvoiceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceType
}

// GetInvoiceTypeOk returns a tuple with the InvoiceType field value
// and a boolean to check if the value has been set.
func (o *ModelInvoice) GetInvoiceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceType, true
}

// SetInvoiceType sets field value
func (o *ModelInvoice) SetInvoiceType(v string) {
	o.InvoiceType = v
}

// GetAccountIntervall returns the AccountIntervall field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoice) GetAccountIntervall() string {
	if o == nil || IsNil(o.AccountIntervall.Get()) {
		var ret string
		return ret
	}
	return *o.AccountIntervall.Get()
}

// GetAccountIntervallOk returns a tuple with the AccountIntervall field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoice) GetAccountIntervallOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountIntervall.Get(), o.AccountIntervall.IsSet()
}

// HasAccountIntervall returns a boolean if a field has been set.
func (o *ModelInvoice) HasAccountIntervall() bool {
	if o != nil && o.AccountIntervall.IsSet() {
		return true
	}

	return false
}

// SetAccountIntervall gets a reference to the given NullableString and assigns it to the AccountIntervall field.
func (o *ModelInvoice) SetAccountIntervall(v string) {
	o.AccountIntervall.Set(&v)
}
// SetAccountIntervallNil sets the value for AccountIntervall to be an explicit nil
func (o *ModelInvoice) SetAccountIntervallNil() {
	o.AccountIntervall.Set(nil)
}

// UnsetAccountIntervall ensures that no value is present for AccountIntervall, not even an explicit nil
func (o *ModelInvoice) UnsetAccountIntervall() {
	o.AccountIntervall.Unset()
}

// GetAccountNextInvoice returns the AccountNextInvoice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoice) GetAccountNextInvoice() int32 {
	if o == nil || IsNil(o.AccountNextInvoice.Get()) {
		var ret int32
		return ret
	}
	return *o.AccountNextInvoice.Get()
}

// GetAccountNextInvoiceOk returns a tuple with the AccountNextInvoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoice) GetAccountNextInvoiceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountNextInvoice.Get(), o.AccountNextInvoice.IsSet()
}

// HasAccountNextInvoice returns a boolean if a field has been set.
func (o *ModelInvoice) HasAccountNextInvoice() bool {
	if o != nil && o.AccountNextInvoice.IsSet() {
		return true
	}

	return false
}

// SetAccountNextInvoice gets a reference to the given NullableInt32 and assigns it to the AccountNextInvoice field.
func (o *ModelInvoice) SetAccountNextInvoice(v int32) {
	o.AccountNextInvoice.Set(&v)
}
// SetAccountNextInvoiceNil sets the value for AccountNextInvoice to be an explicit nil
func (o *ModelInvoice) SetAccountNextInvoiceNil() {
	o.AccountNextInvoice.Set(nil)
}

// UnsetAccountNextInvoice ensures that no value is present for AccountNextInvoice, not even an explicit nil
func (o *ModelInvoice) UnsetAccountNextInvoice() {
	o.AccountNextInvoice.Unset()
}

// GetCurrency returns the Currency field value
func (o *ModelInvoice) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *ModelInvoice) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *ModelInvoice) SetCurrency(v string) {
	o.Currency = v
}

// GetSumNet returns the SumNet field value if set, zero value otherwise.
func (o *ModelInvoice) GetSumNet() float32 {
	if o == nil || IsNil(o.SumNet) {
		var ret float32
		return ret
	}
	return *o.SumNet
}

// GetSumNetOk returns a tuple with the SumNet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoice) GetSumNetOk() (*float32, bool) {
	if o == nil || IsNil(o.SumNet) {
		return nil, false
	}
	return o.SumNet, true
}

// HasSumNet returns a boolean if a field has been set.
func (o *ModelInvoice) HasSumNet() bool {
	if o != nil && !IsNil(o.SumNet) {
		return true
	}

	return false
}

// SetSumNet gets a reference to the given float32 and assigns it to the SumNet field.
func (o *ModelInvoice) SetSumNet(v float32) {
	o.SumNet = &v
}

// GetSumTax returns the SumTax field value if set, zero value otherwise.
func (o *ModelInvoice) GetSumTax() float32 {
	if o == nil || IsNil(o.SumTax) {
		var ret float32
		return ret
	}
	return *o.SumTax
}

// GetSumTaxOk returns a tuple with the SumTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoice) GetSumTaxOk() (*float32, bool) {
	if o == nil || IsNil(o.SumTax) {
		return nil, false
	}
	return o.SumTax, true
}

// HasSumTax returns a boolean if a field has been set.
func (o *ModelInvoice) HasSumTax() bool {
	if o != nil && !IsNil(o.SumTax) {
		return true
	}

	return false
}

// SetSumTax gets a reference to the given float32 and assigns it to the SumTax field.
func (o *ModelInvoice) SetSumTax(v float32) {
	o.SumTax = &v
}

// GetSumGross returns the SumGross field value if set, zero value otherwise.
func (o *ModelInvoice) GetSumGross() float32 {
	if o == nil || IsNil(o.SumGross) {
		var ret float32
		return ret
	}
	return *o.SumGross
}

// GetSumGrossOk returns a tuple with the SumGross field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoice) GetSumGrossOk() (*float32, bool) {
	if o == nil || IsNil(o.SumGross) {
		return nil, false
	}
	return o.SumGross, true
}

// HasSumGross returns a boolean if a field has been set.
func (o *ModelInvoice) HasSumGross() bool {
	if o != nil && !IsNil(o.SumGross) {
		return true
	}

	return false
}

// SetSumGross gets a reference to the given float32 and assigns it to the SumGross field.
func (o *ModelInvoice) SetSumGross(v float32) {
	o.SumGross = &v
}

// GetSumDiscounts returns the SumDiscounts field value if set, zero value otherwise.
func (o *ModelInvoice) GetSumDiscounts() float32 {
	if o == nil || IsNil(o.SumDiscounts) {
		var ret float32
		return ret
	}
	return *o.SumDiscounts
}

// GetSumDiscountsOk returns a tuple with the SumDiscounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoice) GetSumDiscountsOk() (*float32, bool) {
	if o == nil || IsNil(o.SumDiscounts) {
		return nil, false
	}
	return o.SumDiscounts, true
}

// HasSumDiscounts returns a boolean if a field has been set.
func (o *ModelInvoice) HasSumDiscounts() bool {
	if o != nil && !IsNil(o.SumDiscounts) {
		return true
	}

	return false
}

// SetSumDiscounts gets a reference to the given float32 and assigns it to the SumDiscounts field.
func (o *ModelInvoice) SetSumDiscounts(v float32) {
	o.SumDiscounts = &v
}

// GetSumNetForeignCurrency returns the SumNetForeignCurrency field value if set, zero value otherwise.
func (o *ModelInvoice) GetSumNetForeignCurrency() float32 {
	if o == nil || IsNil(o.SumNetForeignCurrency) {
		var ret float32
		return ret
	}
	return *o.SumNetForeignCurrency
}

// GetSumNetForeignCurrencyOk returns a tuple with the SumNetForeignCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoice) GetSumNetForeignCurrencyOk() (*float32, bool) {
	if o == nil || IsNil(o.SumNetForeignCurrency) {
		return nil, false
	}
	return o.SumNetForeignCurrency, true
}

// HasSumNetForeignCurrency returns a boolean if a field has been set.
func (o *ModelInvoice) HasSumNetForeignCurrency() bool {
	if o != nil && !IsNil(o.SumNetForeignCurrency) {
		return true
	}

	return false
}

// SetSumNetForeignCurrency gets a reference to the given float32 and assigns it to the SumNetForeignCurrency field.
func (o *ModelInvoice) SetSumNetForeignCurrency(v float32) {
	o.SumNetForeignCurrency = &v
}

// GetSumTaxForeignCurrency returns the SumTaxForeignCurrency field value if set, zero value otherwise.
func (o *ModelInvoice) GetSumTaxForeignCurrency() float32 {
	if o == nil || IsNil(o.SumTaxForeignCurrency) {
		var ret float32
		return ret
	}
	return *o.SumTaxForeignCurrency
}

// GetSumTaxForeignCurrencyOk returns a tuple with the SumTaxForeignCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoice) GetSumTaxForeignCurrencyOk() (*float32, bool) {
	if o == nil || IsNil(o.SumTaxForeignCurrency) {
		return nil, false
	}
	return o.SumTaxForeignCurrency, true
}

// HasSumTaxForeignCurrency returns a boolean if a field has been set.
func (o *ModelInvoice) HasSumTaxForeignCurrency() bool {
	if o != nil && !IsNil(o.SumTaxForeignCurrency) {
		return true
	}

	return false
}

// SetSumTaxForeignCurrency gets a reference to the given float32 and assigns it to the SumTaxForeignCurrency field.
func (o *ModelInvoice) SetSumTaxForeignCurrency(v float32) {
	o.SumTaxForeignCurrency = &v
}

// GetSumGrossForeignCurrency returns the SumGrossForeignCurrency field value if set, zero value otherwise.
func (o *ModelInvoice) GetSumGrossForeignCurrency() float32 {
	if o == nil || IsNil(o.SumGrossForeignCurrency) {
		var ret float32
		return ret
	}
	return *o.SumGrossForeignCurrency
}

// GetSumGrossForeignCurrencyOk returns a tuple with the SumGrossForeignCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoice) GetSumGrossForeignCurrencyOk() (*float32, bool) {
	if o == nil || IsNil(o.SumGrossForeignCurrency) {
		return nil, false
	}
	return o.SumGrossForeignCurrency, true
}

// HasSumGrossForeignCurrency returns a boolean if a field has been set.
func (o *ModelInvoice) HasSumGrossForeignCurrency() bool {
	if o != nil && !IsNil(o.SumGrossForeignCurrency) {
		return true
	}

	return false
}

// SetSumGrossForeignCurrency gets a reference to the given float32 and assigns it to the SumGrossForeignCurrency field.
func (o *ModelInvoice) SetSumGrossForeignCurrency(v float32) {
	o.SumGrossForeignCurrency = &v
}

// GetSumDiscountsForeignCurrency returns the SumDiscountsForeignCurrency field value if set, zero value otherwise.
func (o *ModelInvoice) GetSumDiscountsForeignCurrency() float32 {
	if o == nil || IsNil(o.SumDiscountsForeignCurrency) {
		var ret float32
		return ret
	}
	return *o.SumDiscountsForeignCurrency
}

// GetSumDiscountsForeignCurrencyOk returns a tuple with the SumDiscountsForeignCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoice) GetSumDiscountsForeignCurrencyOk() (*float32, bool) {
	if o == nil || IsNil(o.SumDiscountsForeignCurrency) {
		return nil, false
	}
	return o.SumDiscountsForeignCurrency, true
}

// HasSumDiscountsForeignCurrency returns a boolean if a field has been set.
func (o *ModelInvoice) HasSumDiscountsForeignCurrency() bool {
	if o != nil && !IsNil(o.SumDiscountsForeignCurrency) {
		return true
	}

	return false
}

// SetSumDiscountsForeignCurrency gets a reference to the given float32 and assigns it to the SumDiscountsForeignCurrency field.
func (o *ModelInvoice) SetSumDiscountsForeignCurrency(v float32) {
	o.SumDiscountsForeignCurrency = &v
}

// GetSumNetAccounting returns the SumNetAccounting field value if set, zero value otherwise.
func (o *ModelInvoice) GetSumNetAccounting() float32 {
	if o == nil || IsNil(o.SumNetAccounting) {
		var ret float32
		return ret
	}
	return *o.SumNetAccounting
}

// GetSumNetAccountingOk returns a tuple with the SumNetAccounting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoice) GetSumNetAccountingOk() (*float32, bool) {
	if o == nil || IsNil(o.SumNetAccounting) {
		return nil, false
	}
	return o.SumNetAccounting, true
}

// HasSumNetAccounting returns a boolean if a field has been set.
func (o *ModelInvoice) HasSumNetAccounting() bool {
	if o != nil && !IsNil(o.SumNetAccounting) {
		return true
	}

	return false
}

// SetSumNetAccounting gets a reference to the given float32 and assigns it to the SumNetAccounting field.
func (o *ModelInvoice) SetSumNetAccounting(v float32) {
	o.SumNetAccounting = &v
}

// GetSumTaxAccounting returns the SumTaxAccounting field value if set, zero value otherwise.
func (o *ModelInvoice) GetSumTaxAccounting() float32 {
	if o == nil || IsNil(o.SumTaxAccounting) {
		var ret float32
		return ret
	}
	return *o.SumTaxAccounting
}

// GetSumTaxAccountingOk returns a tuple with the SumTaxAccounting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoice) GetSumTaxAccountingOk() (*float32, bool) {
	if o == nil || IsNil(o.SumTaxAccounting) {
		return nil, false
	}
	return o.SumTaxAccounting, true
}

// HasSumTaxAccounting returns a boolean if a field has been set.
func (o *ModelInvoice) HasSumTaxAccounting() bool {
	if o != nil && !IsNil(o.SumTaxAccounting) {
		return true
	}

	return false
}

// SetSumTaxAccounting gets a reference to the given float32 and assigns it to the SumTaxAccounting field.
func (o *ModelInvoice) SetSumTaxAccounting(v float32) {
	o.SumTaxAccounting = &v
}

// GetSumGrossAccounting returns the SumGrossAccounting field value if set, zero value otherwise.
func (o *ModelInvoice) GetSumGrossAccounting() float32 {
	if o == nil || IsNil(o.SumGrossAccounting) {
		var ret float32
		return ret
	}
	return *o.SumGrossAccounting
}

// GetSumGrossAccountingOk returns a tuple with the SumGrossAccounting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoice) GetSumGrossAccountingOk() (*float32, bool) {
	if o == nil || IsNil(o.SumGrossAccounting) {
		return nil, false
	}
	return o.SumGrossAccounting, true
}

// HasSumGrossAccounting returns a boolean if a field has been set.
func (o *ModelInvoice) HasSumGrossAccounting() bool {
	if o != nil && !IsNil(o.SumGrossAccounting) {
		return true
	}

	return false
}

// SetSumGrossAccounting gets a reference to the given float32 and assigns it to the SumGrossAccounting field.
func (o *ModelInvoice) SetSumGrossAccounting(v float32) {
	o.SumGrossAccounting = &v
}

// GetPaidAmount returns the PaidAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoice) GetPaidAmount() float32 {
	if o == nil || IsNil(o.PaidAmount.Get()) {
		var ret float32
		return ret
	}
	return *o.PaidAmount.Get()
}

// GetPaidAmountOk returns a tuple with the PaidAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoice) GetPaidAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaidAmount.Get(), o.PaidAmount.IsSet()
}

// HasPaidAmount returns a boolean if a field has been set.
func (o *ModelInvoice) HasPaidAmount() bool {
	if o != nil && o.PaidAmount.IsSet() {
		return true
	}

	return false
}

// SetPaidAmount gets a reference to the given NullableFloat32 and assigns it to the PaidAmount field.
func (o *ModelInvoice) SetPaidAmount(v float32) {
	o.PaidAmount.Set(&v)
}
// SetPaidAmountNil sets the value for PaidAmount to be an explicit nil
func (o *ModelInvoice) SetPaidAmountNil() {
	o.PaidAmount.Set(nil)
}

// UnsetPaidAmount ensures that no value is present for PaidAmount, not even an explicit nil
func (o *ModelInvoice) UnsetPaidAmount() {
	o.PaidAmount.Unset()
}

// GetShowNet returns the ShowNet field value if set, zero value otherwise.
func (o *ModelInvoice) GetShowNet() bool {
	if o == nil || IsNil(o.ShowNet) {
		var ret bool
		return ret
	}
	return *o.ShowNet
}

// GetShowNetOk returns a tuple with the ShowNet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoice) GetShowNetOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowNet) {
		return nil, false
	}
	return o.ShowNet, true
}

// HasShowNet returns a boolean if a field has been set.
func (o *ModelInvoice) HasShowNet() bool {
	if o != nil && !IsNil(o.ShowNet) {
		return true
	}

	return false
}

// SetShowNet gets a reference to the given bool and assigns it to the ShowNet field.
func (o *ModelInvoice) SetShowNet(v bool) {
	o.ShowNet = &v
}

// GetEnshrined returns the Enshrined field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoice) GetEnshrined() time.Time {
	if o == nil || IsNil(o.Enshrined.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Enshrined.Get()
}

// GetEnshrinedOk returns a tuple with the Enshrined field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoice) GetEnshrinedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enshrined.Get(), o.Enshrined.IsSet()
}

// HasEnshrined returns a boolean if a field has been set.
func (o *ModelInvoice) HasEnshrined() bool {
	if o != nil && o.Enshrined.IsSet() {
		return true
	}

	return false
}

// SetEnshrined gets a reference to the given NullableTime and assigns it to the Enshrined field.
func (o *ModelInvoice) SetEnshrined(v time.Time) {
	o.Enshrined.Set(&v)
}
// SetEnshrinedNil sets the value for Enshrined to be an explicit nil
func (o *ModelInvoice) SetEnshrinedNil() {
	o.Enshrined.Set(nil)
}

// UnsetEnshrined ensures that no value is present for Enshrined, not even an explicit nil
func (o *ModelInvoice) UnsetEnshrined() {
	o.Enshrined.Unset()
}

// GetSendType returns the SendType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoice) GetSendType() string {
	if o == nil || IsNil(o.SendType.Get()) {
		var ret string
		return ret
	}
	return *o.SendType.Get()
}

// GetSendTypeOk returns a tuple with the SendType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoice) GetSendTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SendType.Get(), o.SendType.IsSet()
}

// HasSendType returns a boolean if a field has been set.
func (o *ModelInvoice) HasSendType() bool {
	if o != nil && o.SendType.IsSet() {
		return true
	}

	return false
}

// SetSendType gets a reference to the given NullableString and assigns it to the SendType field.
func (o *ModelInvoice) SetSendType(v string) {
	o.SendType.Set(&v)
}
// SetSendTypeNil sets the value for SendType to be an explicit nil
func (o *ModelInvoice) SetSendTypeNil() {
	o.SendType.Set(nil)
}

// UnsetSendType ensures that no value is present for SendType, not even an explicit nil
func (o *ModelInvoice) UnsetSendType() {
	o.SendType.Unset()
}

// GetOrigin returns the Origin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoice) GetOrigin() ModelInvoiceOrigin {
	if o == nil || IsNil(o.Origin.Get()) {
		var ret ModelInvoiceOrigin
		return ret
	}
	return *o.Origin.Get()
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoice) GetOriginOk() (*ModelInvoiceOrigin, bool) {
	if o == nil {
		return nil, false
	}
	return o.Origin.Get(), o.Origin.IsSet()
}

// HasOrigin returns a boolean if a field has been set.
func (o *ModelInvoice) HasOrigin() bool {
	if o != nil && o.Origin.IsSet() {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given NullableModelInvoiceOrigin and assigns it to the Origin field.
func (o *ModelInvoice) SetOrigin(v ModelInvoiceOrigin) {
	o.Origin.Set(&v)
}
// SetOriginNil sets the value for Origin to be an explicit nil
func (o *ModelInvoice) SetOriginNil() {
	o.Origin.Set(nil)
}

// UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
func (o *ModelInvoice) UnsetOrigin() {
	o.Origin.Unset()
}

// GetCustomerInternalNote returns the CustomerInternalNote field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoice) GetCustomerInternalNote() string {
	if o == nil || IsNil(o.CustomerInternalNote.Get()) {
		var ret string
		return ret
	}
	return *o.CustomerInternalNote.Get()
}

// GetCustomerInternalNoteOk returns a tuple with the CustomerInternalNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoice) GetCustomerInternalNoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerInternalNote.Get(), o.CustomerInternalNote.IsSet()
}

// HasCustomerInternalNote returns a boolean if a field has been set.
func (o *ModelInvoice) HasCustomerInternalNote() bool {
	if o != nil && o.CustomerInternalNote.IsSet() {
		return true
	}

	return false
}

// SetCustomerInternalNote gets a reference to the given NullableString and assigns it to the CustomerInternalNote field.
func (o *ModelInvoice) SetCustomerInternalNote(v string) {
	o.CustomerInternalNote.Set(&v)
}
// SetCustomerInternalNoteNil sets the value for CustomerInternalNote to be an explicit nil
func (o *ModelInvoice) SetCustomerInternalNoteNil() {
	o.CustomerInternalNote.Set(nil)
}

// UnsetCustomerInternalNote ensures that no value is present for CustomerInternalNote, not even an explicit nil
func (o *ModelInvoice) UnsetCustomerInternalNote() {
	o.CustomerInternalNote.Unset()
}

// GetMapAll returns the MapAll field value
func (o *ModelInvoice) GetMapAll() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MapAll
}

// GetMapAllOk returns a tuple with the MapAll field value
// and a boolean to check if the value has been set.
func (o *ModelInvoice) GetMapAllOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MapAll, true
}

// SetMapAll sets field value
func (o *ModelInvoice) SetMapAll(v bool) {
	o.MapAll = v
}

func (o ModelInvoice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelInvoice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if !IsNil(o.ObjectName) {
		toSerialize["objectName"] = o.ObjectName
	}
	if o.InvoiceNumber.IsSet() {
		toSerialize["invoiceNumber"] = o.InvoiceNumber.Get()
	}
	toSerialize["contact"] = o.Contact
	toSerialize["contactPerson"] = o.ContactPerson
	// skip: create is readOnly
	// skip: update is readOnly
	if !IsNil(o.SevClient) {
		toSerialize["sevClient"] = o.SevClient
	}
	toSerialize["invoiceDate"] = o.InvoiceDate
	if o.Header.IsSet() {
		toSerialize["header"] = o.Header.Get()
	}
	if o.HeadText.IsSet() {
		toSerialize["headText"] = o.HeadText.Get()
	}
	if o.FootText.IsSet() {
		toSerialize["footText"] = o.FootText.Get()
	}
	if o.TimeToPay.IsSet() {
		toSerialize["timeToPay"] = o.TimeToPay.Get()
	}
	toSerialize["discount"] = o.Discount
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	toSerialize["addressCountry"] = o.AddressCountry
	if o.PayDate.IsSet() {
		toSerialize["payDate"] = o.PayDate.Get()
	}
	if !IsNil(o.CreateUser) {
		toSerialize["createUser"] = o.CreateUser
	}
	if o.DeliveryDate.IsSet() {
		toSerialize["deliveryDate"] = o.DeliveryDate.Get()
	}
	if o.DeliveryDateUntil.IsSet() {
		toSerialize["deliveryDateUntil"] = o.DeliveryDateUntil.Get()
	}
	toSerialize["status"] = o.Status
	if o.SmallSettlement.IsSet() {
		toSerialize["smallSettlement"] = o.SmallSettlement.Get()
	}
	toSerialize["taxRate"] = o.TaxRate
	toSerialize["taxText"] = o.TaxText
	toSerialize["taxType"] = o.TaxType
	if o.TaxSet.IsSet() {
		toSerialize["taxSet"] = o.TaxSet.Get()
	}
	if o.DunningLevel.IsSet() {
		toSerialize["dunningLevel"] = o.DunningLevel.Get()
	}
	if !IsNil(o.PaymentMethod) {
		toSerialize["paymentMethod"] = o.PaymentMethod
	}
	if o.SendDate.IsSet() {
		toSerialize["sendDate"] = o.SendDate.Get()
	}
	toSerialize["invoiceType"] = o.InvoiceType
	if o.AccountIntervall.IsSet() {
		toSerialize["accountIntervall"] = o.AccountIntervall.Get()
	}
	if o.AccountNextInvoice.IsSet() {
		toSerialize["accountNextInvoice"] = o.AccountNextInvoice.Get()
	}
	toSerialize["currency"] = o.Currency
	// skip: sumNet is readOnly
	// skip: sumTax is readOnly
	// skip: sumGross is readOnly
	// skip: sumDiscounts is readOnly
	// skip: sumNetForeignCurrency is readOnly
	// skip: sumTaxForeignCurrency is readOnly
	// skip: sumGrossForeignCurrency is readOnly
	// skip: sumDiscountsForeignCurrency is readOnly
	// skip: sumNetAccounting is readOnly
	// skip: sumTaxAccounting is readOnly
	// skip: sumGrossAccounting is readOnly
	if o.PaidAmount.IsSet() {
		toSerialize["paidAmount"] = o.PaidAmount.Get()
	}
	if !IsNil(o.ShowNet) {
		toSerialize["showNet"] = o.ShowNet
	}
	if o.Enshrined.IsSet() {
		toSerialize["enshrined"] = o.Enshrined.Get()
	}
	if o.SendType.IsSet() {
		toSerialize["sendType"] = o.SendType.Get()
	}
	if o.Origin.IsSet() {
		toSerialize["origin"] = o.Origin.Get()
	}
	if o.CustomerInternalNote.IsSet() {
		toSerialize["customerInternalNote"] = o.CustomerInternalNote.Get()
	}
	toSerialize["mapAll"] = o.MapAll
	return toSerialize, nil
}

type NullableModelInvoice struct {
	value *ModelInvoice
	isSet bool
}

func (v NullableModelInvoice) Get() *ModelInvoice {
	return v.value
}

func (v *NullableModelInvoice) Set(val *ModelInvoice) {
	v.value = val
	v.isSet = true
}

func (v NullableModelInvoice) IsSet() bool {
	return v.isSet
}

func (v *NullableModelInvoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelInvoice(val *ModelInvoice) *NullableModelInvoice {
	return &NullableModelInvoice{value: val, isSet: true}
}

func (v NullableModelInvoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelInvoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


