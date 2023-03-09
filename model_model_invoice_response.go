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

// checks if the ModelInvoiceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelInvoiceResponse{}

// ModelInvoiceResponse Invoice model
type ModelInvoiceResponse struct {
	// The invoice id
	Id *string `json:"id,omitempty"`
	// The invoice object name
	ObjectName *string `json:"objectName,omitempty"`
	// The invoice number
	InvoiceNumber *string `json:"invoiceNumber,omitempty"`
	Contact *ModelInvoiceResponseContact `json:"contact,omitempty"`
	// Date of invoice creation
	Create *time.Time `json:"create,omitempty"`
	// Date of last invoice update
	Update *time.Time `json:"update,omitempty"`
	SevClient *ModelContactCustomFieldSettingResponseSevClient `json:"sevClient,omitempty"`
	// Needs to be provided as timestamp or dd.mm.yyyy
	InvoiceDate *string `json:"invoiceDate,omitempty"`
	// Normally consist of prefix plus the invoice number
	Header *string `json:"header,omitempty"`
	// Certain html tags can be used here to format your text
	HeadText *string `json:"headText,omitempty"`
	// Certain html tags can be used here to format your text
	FootText *string `json:"footText,omitempty"`
	// The time the customer has to pay the invoice in days
	TimeToPay *string `json:"timeToPay,omitempty"`
	// If a value other than zero is used for the discount attribute,      you need to specify the amount of days for which the discount is granted.
	DiscountTime *string `json:"discountTime,omitempty"`
	// If you want to give a discount, define the percentage here. Otherwise provide zero as value
	Discount *string `json:"discount,omitempty"`
	AddressCountry *ModelInvoiceResponseAddressCountry `json:"addressCountry,omitempty"`
	// Needs to be timestamp or dd.mm.yyyy
	PayDate *time.Time `json:"payDate,omitempty"`
	CreateUser *ModelInvoiceResponseCreateUser `json:"createUser,omitempty"`
	// Timestamp. This can also be a date range if you also use the attribute deliveryDateUntil
	DeliveryDate *time.Time `json:"deliveryDate,omitempty"`
	// Please have a look in our       <a href='https://api.sevdesk.de/#section/Types-and-status-of-invoices'>Types and status of invoices</a>       to see what the different status codes mean
	Status *string `json:"status,omitempty"`
	// Defines if the client uses the small settlement scheme.      If yes, the invoice must not contain any vat
	SmallSettlement *bool `json:"smallSettlement,omitempty"`
	ContactPerson *ModelInvoiceResponseContactPerson `json:"contactPerson,omitempty"`
	// Is overwritten by invoice position tax rates
	TaxRate *string `json:"taxRate,omitempty"`
	// A common tax text would be 'Umsatzsteuer 19%'
	TaxText *string `json:"taxText,omitempty"`
	// Defines how many reminders have already been sent for the invoice.      Starts with 1 (Payment reminder) and should be incremented by one every time another reminder is sent.
	DunningLevel *string `json:"dunningLevel,omitempty"`
	// Tax type of the invoice. There are four tax types: 1. default - Umsatzsteuer ausweisen 2. eu - Steuerfreie innergemeinschaftliche Lieferung (Europäische Union) 3. noteu - Steuerschuldnerschaft des Leistungsempfängers (außerhalb EU, z. B. Schweiz) 4. custom - Using custom tax set 5. ss - Not subject to VAT according to §19 1 UStG Tax rates are heavily connected to the tax type used.
	TaxType *string `json:"taxType,omitempty"`
	PaymentMethod *ModelInvoiceResponsePaymentMethod `json:"paymentMethod,omitempty"`
	CostCentre *ModelInvoiceResponseCostCentre `json:"costCentre,omitempty"`
	// The date the invoice was sent to the customer
	SendDate *time.Time `json:"sendDate,omitempty"`
	Origin *ModelInvoiceResponseOrigin `json:"origin,omitempty"`
	// Type of the invoice. For more information on the different types, check       <a href='https://api.sevdesk.de/#section/Types-and-status-of-invoices'>this</a> section  
	InvoiceType *string `json:"invoiceType,omitempty"`
	// The interval in which recurring invoices are due as ISO-8601 duration.<br>       Necessary attribute for all recurring invoices.
	AccountIntervall *string `json:"accountIntervall,omitempty"`
	// Timestamp when the next invoice will be generated by this recurring invoice.
	AccountNextInvoice *string `json:"accountNextInvoice,omitempty"`
	// Total reminder amount
	ReminderTotal *string `json:"reminderTotal,omitempty"`
	// Debit of the reminder
	ReminderDebit *string `json:"reminderDebit,omitempty"`
	// Deadline of the reminder as timestamp
	ReminderDeadline *time.Time `json:"reminderDeadline,omitempty"`
	// The additional reminder charge
	ReminderCharge *string `json:"reminderCharge,omitempty"`
	TaxSet *ModelInvoiceResponseTaxSet `json:"taxSet,omitempty"`
	// Complete address of the recipient including name, street, city, zip and country.       * Line breaks can be used and will be displayed on the invoice pdf.
	Address *string `json:"address,omitempty"`
	// Currency used in the invoice. Needs to be currency code according to ISO-4217
	Currency *string `json:"currency,omitempty"`
	// Net sum of the invoice
	SumNet *string `json:"sumNet,omitempty"`
	// Tax sum of the invoice
	SumTax *string `json:"sumTax,omitempty"`
	// Gross sum of the invoice
	SumGross *string `json:"sumGross,omitempty"`
	// Sum of all discounts in the invoice
	SumDiscounts *string `json:"sumDiscounts,omitempty"`
	// Net sum of the invoice in the foreign currency
	SumNetForeignCurrency *string `json:"sumNetForeignCurrency,omitempty"`
	// Tax sum of the invoice in the foreign currency
	SumTaxForeignCurrency *string `json:"sumTaxForeignCurrency,omitempty"`
	// Gross sum of the invoice in the foreign currency
	SumGrossForeignCurrency *string `json:"sumGrossForeignCurrency,omitempty"`
	// Discounts sum of the invoice in the foreign currency
	SumDiscountsForeignCurrency *string `json:"sumDiscountsForeignCurrency,omitempty"`
	// Net accounting sum of the invoice. Is usually the same as sumNet
	SumNetAccounting *string `json:"sumNetAccounting,omitempty"`
	// Tax accounting sum of the invoice. Is usually the same as sumTax
	SumTaxAccounting *string `json:"sumTaxAccounting,omitempty"`
	// Gross accounting sum of the invoice. Is usually the same as sumGross
	SumGrossAccounting *string `json:"sumGrossAccounting,omitempty"`
	// Amount which has already been paid for this invoice by the customer
	PaidAmount *float32 `json:"paidAmount,omitempty"`
	// Internal note of the customer. Contains data entered into field 'Referenz/Bestellnummer'
	CustomerInternalNote *string `json:"customerInternalNote,omitempty"`
	// If true, the net amount of each position will be shown on the invoice. Otherwise gross amount
	ShowNet *bool `json:"showNet,omitempty"`
	// Defines if and when invoice was enshrined. Enshrined invoices can not be manipulated.
	Enshrined *time.Time `json:"enshrined,omitempty"`
	// Type which was used to send the invoice. IMPORTANT: Please refer to the invoice section of the       *     API-Overview to understand how this attribute can be used before using it!
	SendType *string `json:"sendType,omitempty"`
	// If the delivery date should be a time range, another timestamp can be provided in this attribute       * to define a range from timestamp used in deliveryDate attribute to the timestamp used here.
	DeliveryDateUntil *string `json:"deliveryDateUntil,omitempty"`
	// Internal attribute
	DatevConnectOnline map[string]interface{} `json:"datevConnectOnline,omitempty"`
	// Internal attribute
	SendPaymentReceivedNotificationDate *string `json:"sendPaymentReceivedNotificationDate,omitempty"`
}

// NewModelInvoiceResponse instantiates a new ModelInvoiceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelInvoiceResponse() *ModelInvoiceResponse {
	this := ModelInvoiceResponse{}
	return &this
}

// NewModelInvoiceResponseWithDefaults instantiates a new ModelInvoiceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelInvoiceResponseWithDefaults() *ModelInvoiceResponse {
	this := ModelInvoiceResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModelInvoiceResponse) SetId(v string) {
	o.Id = &v
}

// GetObjectName returns the ObjectName field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetObjectName() string {
	if o == nil || IsNil(o.ObjectName) {
		var ret string
		return ret
	}
	return *o.ObjectName
}

// GetObjectNameOk returns a tuple with the ObjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetObjectNameOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectName) {
		return nil, false
	}
	return o.ObjectName, true
}

// HasObjectName returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasObjectName() bool {
	if o != nil && !IsNil(o.ObjectName) {
		return true
	}

	return false
}

// SetObjectName gets a reference to the given string and assigns it to the ObjectName field.
func (o *ModelInvoiceResponse) SetObjectName(v string) {
	o.ObjectName = &v
}

// GetInvoiceNumber returns the InvoiceNumber field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetInvoiceNumber() string {
	if o == nil || IsNil(o.InvoiceNumber) {
		var ret string
		return ret
	}
	return *o.InvoiceNumber
}

// GetInvoiceNumberOk returns a tuple with the InvoiceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetInvoiceNumberOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceNumber) {
		return nil, false
	}
	return o.InvoiceNumber, true
}

// HasInvoiceNumber returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasInvoiceNumber() bool {
	if o != nil && !IsNil(o.InvoiceNumber) {
		return true
	}

	return false
}

// SetInvoiceNumber gets a reference to the given string and assigns it to the InvoiceNumber field.
func (o *ModelInvoiceResponse) SetInvoiceNumber(v string) {
	o.InvoiceNumber = &v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetContact() ModelInvoiceResponseContact {
	if o == nil || IsNil(o.Contact) {
		var ret ModelInvoiceResponseContact
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetContactOk() (*ModelInvoiceResponseContact, bool) {
	if o == nil || IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasContact() bool {
	if o != nil && !IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given ModelInvoiceResponseContact and assigns it to the Contact field.
func (o *ModelInvoiceResponse) SetContact(v ModelInvoiceResponseContact) {
	o.Contact = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *ModelInvoiceResponse) SetCreate(v time.Time) {
	o.Create = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *ModelInvoiceResponse) SetUpdate(v time.Time) {
	o.Update = &v
}

// GetSevClient returns the SevClient field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetSevClient() ModelContactCustomFieldSettingResponseSevClient {
	if o == nil || IsNil(o.SevClient) {
		var ret ModelContactCustomFieldSettingResponseSevClient
		return ret
	}
	return *o.SevClient
}

// GetSevClientOk returns a tuple with the SevClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetSevClientOk() (*ModelContactCustomFieldSettingResponseSevClient, bool) {
	if o == nil || IsNil(o.SevClient) {
		return nil, false
	}
	return o.SevClient, true
}

// HasSevClient returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasSevClient() bool {
	if o != nil && !IsNil(o.SevClient) {
		return true
	}

	return false
}

// SetSevClient gets a reference to the given ModelContactCustomFieldSettingResponseSevClient and assigns it to the SevClient field.
func (o *ModelInvoiceResponse) SetSevClient(v ModelContactCustomFieldSettingResponseSevClient) {
	o.SevClient = &v
}

// GetInvoiceDate returns the InvoiceDate field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetInvoiceDate() string {
	if o == nil || IsNil(o.InvoiceDate) {
		var ret string
		return ret
	}
	return *o.InvoiceDate
}

// GetInvoiceDateOk returns a tuple with the InvoiceDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetInvoiceDateOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceDate) {
		return nil, false
	}
	return o.InvoiceDate, true
}

// HasInvoiceDate returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasInvoiceDate() bool {
	if o != nil && !IsNil(o.InvoiceDate) {
		return true
	}

	return false
}

// SetInvoiceDate gets a reference to the given string and assigns it to the InvoiceDate field.
func (o *ModelInvoiceResponse) SetInvoiceDate(v string) {
	o.InvoiceDate = &v
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetHeader() string {
	if o == nil || IsNil(o.Header) {
		var ret string
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetHeaderOk() (*string, bool) {
	if o == nil || IsNil(o.Header) {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasHeader() bool {
	if o != nil && !IsNil(o.Header) {
		return true
	}

	return false
}

// SetHeader gets a reference to the given string and assigns it to the Header field.
func (o *ModelInvoiceResponse) SetHeader(v string) {
	o.Header = &v
}

// GetHeadText returns the HeadText field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetHeadText() string {
	if o == nil || IsNil(o.HeadText) {
		var ret string
		return ret
	}
	return *o.HeadText
}

// GetHeadTextOk returns a tuple with the HeadText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetHeadTextOk() (*string, bool) {
	if o == nil || IsNil(o.HeadText) {
		return nil, false
	}
	return o.HeadText, true
}

// HasHeadText returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasHeadText() bool {
	if o != nil && !IsNil(o.HeadText) {
		return true
	}

	return false
}

// SetHeadText gets a reference to the given string and assigns it to the HeadText field.
func (o *ModelInvoiceResponse) SetHeadText(v string) {
	o.HeadText = &v
}

// GetFootText returns the FootText field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetFootText() string {
	if o == nil || IsNil(o.FootText) {
		var ret string
		return ret
	}
	return *o.FootText
}

// GetFootTextOk returns a tuple with the FootText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetFootTextOk() (*string, bool) {
	if o == nil || IsNil(o.FootText) {
		return nil, false
	}
	return o.FootText, true
}

// HasFootText returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasFootText() bool {
	if o != nil && !IsNil(o.FootText) {
		return true
	}

	return false
}

// SetFootText gets a reference to the given string and assigns it to the FootText field.
func (o *ModelInvoiceResponse) SetFootText(v string) {
	o.FootText = &v
}

// GetTimeToPay returns the TimeToPay field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetTimeToPay() string {
	if o == nil || IsNil(o.TimeToPay) {
		var ret string
		return ret
	}
	return *o.TimeToPay
}

// GetTimeToPayOk returns a tuple with the TimeToPay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetTimeToPayOk() (*string, bool) {
	if o == nil || IsNil(o.TimeToPay) {
		return nil, false
	}
	return o.TimeToPay, true
}

// HasTimeToPay returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasTimeToPay() bool {
	if o != nil && !IsNil(o.TimeToPay) {
		return true
	}

	return false
}

// SetTimeToPay gets a reference to the given string and assigns it to the TimeToPay field.
func (o *ModelInvoiceResponse) SetTimeToPay(v string) {
	o.TimeToPay = &v
}

// GetDiscountTime returns the DiscountTime field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetDiscountTime() string {
	if o == nil || IsNil(o.DiscountTime) {
		var ret string
		return ret
	}
	return *o.DiscountTime
}

// GetDiscountTimeOk returns a tuple with the DiscountTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetDiscountTimeOk() (*string, bool) {
	if o == nil || IsNil(o.DiscountTime) {
		return nil, false
	}
	return o.DiscountTime, true
}

// HasDiscountTime returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasDiscountTime() bool {
	if o != nil && !IsNil(o.DiscountTime) {
		return true
	}

	return false
}

// SetDiscountTime gets a reference to the given string and assigns it to the DiscountTime field.
func (o *ModelInvoiceResponse) SetDiscountTime(v string) {
	o.DiscountTime = &v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetDiscount() string {
	if o == nil || IsNil(o.Discount) {
		var ret string
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetDiscountOk() (*string, bool) {
	if o == nil || IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasDiscount() bool {
	if o != nil && !IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given string and assigns it to the Discount field.
func (o *ModelInvoiceResponse) SetDiscount(v string) {
	o.Discount = &v
}

// GetAddressCountry returns the AddressCountry field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetAddressCountry() ModelInvoiceResponseAddressCountry {
	if o == nil || IsNil(o.AddressCountry) {
		var ret ModelInvoiceResponseAddressCountry
		return ret
	}
	return *o.AddressCountry
}

// GetAddressCountryOk returns a tuple with the AddressCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetAddressCountryOk() (*ModelInvoiceResponseAddressCountry, bool) {
	if o == nil || IsNil(o.AddressCountry) {
		return nil, false
	}
	return o.AddressCountry, true
}

// HasAddressCountry returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasAddressCountry() bool {
	if o != nil && !IsNil(o.AddressCountry) {
		return true
	}

	return false
}

// SetAddressCountry gets a reference to the given ModelInvoiceResponseAddressCountry and assigns it to the AddressCountry field.
func (o *ModelInvoiceResponse) SetAddressCountry(v ModelInvoiceResponseAddressCountry) {
	o.AddressCountry = &v
}

// GetPayDate returns the PayDate field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetPayDate() time.Time {
	if o == nil || IsNil(o.PayDate) {
		var ret time.Time
		return ret
	}
	return *o.PayDate
}

// GetPayDateOk returns a tuple with the PayDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetPayDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PayDate) {
		return nil, false
	}
	return o.PayDate, true
}

// HasPayDate returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasPayDate() bool {
	if o != nil && !IsNil(o.PayDate) {
		return true
	}

	return false
}

// SetPayDate gets a reference to the given time.Time and assigns it to the PayDate field.
func (o *ModelInvoiceResponse) SetPayDate(v time.Time) {
	o.PayDate = &v
}

// GetCreateUser returns the CreateUser field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetCreateUser() ModelInvoiceResponseCreateUser {
	if o == nil || IsNil(o.CreateUser) {
		var ret ModelInvoiceResponseCreateUser
		return ret
	}
	return *o.CreateUser
}

// GetCreateUserOk returns a tuple with the CreateUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetCreateUserOk() (*ModelInvoiceResponseCreateUser, bool) {
	if o == nil || IsNil(o.CreateUser) {
		return nil, false
	}
	return o.CreateUser, true
}

// HasCreateUser returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasCreateUser() bool {
	if o != nil && !IsNil(o.CreateUser) {
		return true
	}

	return false
}

// SetCreateUser gets a reference to the given ModelInvoiceResponseCreateUser and assigns it to the CreateUser field.
func (o *ModelInvoiceResponse) SetCreateUser(v ModelInvoiceResponseCreateUser) {
	o.CreateUser = &v
}

// GetDeliveryDate returns the DeliveryDate field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetDeliveryDate() time.Time {
	if o == nil || IsNil(o.DeliveryDate) {
		var ret time.Time
		return ret
	}
	return *o.DeliveryDate
}

// GetDeliveryDateOk returns a tuple with the DeliveryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetDeliveryDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeliveryDate) {
		return nil, false
	}
	return o.DeliveryDate, true
}

// HasDeliveryDate returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasDeliveryDate() bool {
	if o != nil && !IsNil(o.DeliveryDate) {
		return true
	}

	return false
}

// SetDeliveryDate gets a reference to the given time.Time and assigns it to the DeliveryDate field.
func (o *ModelInvoiceResponse) SetDeliveryDate(v time.Time) {
	o.DeliveryDate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ModelInvoiceResponse) SetStatus(v string) {
	o.Status = &v
}

// GetSmallSettlement returns the SmallSettlement field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetSmallSettlement() bool {
	if o == nil || IsNil(o.SmallSettlement) {
		var ret bool
		return ret
	}
	return *o.SmallSettlement
}

// GetSmallSettlementOk returns a tuple with the SmallSettlement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetSmallSettlementOk() (*bool, bool) {
	if o == nil || IsNil(o.SmallSettlement) {
		return nil, false
	}
	return o.SmallSettlement, true
}

// HasSmallSettlement returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasSmallSettlement() bool {
	if o != nil && !IsNil(o.SmallSettlement) {
		return true
	}

	return false
}

// SetSmallSettlement gets a reference to the given bool and assigns it to the SmallSettlement field.
func (o *ModelInvoiceResponse) SetSmallSettlement(v bool) {
	o.SmallSettlement = &v
}

// GetContactPerson returns the ContactPerson field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetContactPerson() ModelInvoiceResponseContactPerson {
	if o == nil || IsNil(o.ContactPerson) {
		var ret ModelInvoiceResponseContactPerson
		return ret
	}
	return *o.ContactPerson
}

// GetContactPersonOk returns a tuple with the ContactPerson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetContactPersonOk() (*ModelInvoiceResponseContactPerson, bool) {
	if o == nil || IsNil(o.ContactPerson) {
		return nil, false
	}
	return o.ContactPerson, true
}

// HasContactPerson returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasContactPerson() bool {
	if o != nil && !IsNil(o.ContactPerson) {
		return true
	}

	return false
}

// SetContactPerson gets a reference to the given ModelInvoiceResponseContactPerson and assigns it to the ContactPerson field.
func (o *ModelInvoiceResponse) SetContactPerson(v ModelInvoiceResponseContactPerson) {
	o.ContactPerson = &v
}

// GetTaxRate returns the TaxRate field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetTaxRate() string {
	if o == nil || IsNil(o.TaxRate) {
		var ret string
		return ret
	}
	return *o.TaxRate
}

// GetTaxRateOk returns a tuple with the TaxRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetTaxRateOk() (*string, bool) {
	if o == nil || IsNil(o.TaxRate) {
		return nil, false
	}
	return o.TaxRate, true
}

// HasTaxRate returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasTaxRate() bool {
	if o != nil && !IsNil(o.TaxRate) {
		return true
	}

	return false
}

// SetTaxRate gets a reference to the given string and assigns it to the TaxRate field.
func (o *ModelInvoiceResponse) SetTaxRate(v string) {
	o.TaxRate = &v
}

// GetTaxText returns the TaxText field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetTaxText() string {
	if o == nil || IsNil(o.TaxText) {
		var ret string
		return ret
	}
	return *o.TaxText
}

// GetTaxTextOk returns a tuple with the TaxText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetTaxTextOk() (*string, bool) {
	if o == nil || IsNil(o.TaxText) {
		return nil, false
	}
	return o.TaxText, true
}

// HasTaxText returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasTaxText() bool {
	if o != nil && !IsNil(o.TaxText) {
		return true
	}

	return false
}

// SetTaxText gets a reference to the given string and assigns it to the TaxText field.
func (o *ModelInvoiceResponse) SetTaxText(v string) {
	o.TaxText = &v
}

// GetDunningLevel returns the DunningLevel field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetDunningLevel() string {
	if o == nil || IsNil(o.DunningLevel) {
		var ret string
		return ret
	}
	return *o.DunningLevel
}

// GetDunningLevelOk returns a tuple with the DunningLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetDunningLevelOk() (*string, bool) {
	if o == nil || IsNil(o.DunningLevel) {
		return nil, false
	}
	return o.DunningLevel, true
}

// HasDunningLevel returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasDunningLevel() bool {
	if o != nil && !IsNil(o.DunningLevel) {
		return true
	}

	return false
}

// SetDunningLevel gets a reference to the given string and assigns it to the DunningLevel field.
func (o *ModelInvoiceResponse) SetDunningLevel(v string) {
	o.DunningLevel = &v
}

// GetTaxType returns the TaxType field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetTaxType() string {
	if o == nil || IsNil(o.TaxType) {
		var ret string
		return ret
	}
	return *o.TaxType
}

// GetTaxTypeOk returns a tuple with the TaxType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetTaxTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TaxType) {
		return nil, false
	}
	return o.TaxType, true
}

// HasTaxType returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasTaxType() bool {
	if o != nil && !IsNil(o.TaxType) {
		return true
	}

	return false
}

// SetTaxType gets a reference to the given string and assigns it to the TaxType field.
func (o *ModelInvoiceResponse) SetTaxType(v string) {
	o.TaxType = &v
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetPaymentMethod() ModelInvoiceResponsePaymentMethod {
	if o == nil || IsNil(o.PaymentMethod) {
		var ret ModelInvoiceResponsePaymentMethod
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetPaymentMethodOk() (*ModelInvoiceResponsePaymentMethod, bool) {
	if o == nil || IsNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasPaymentMethod() bool {
	if o != nil && !IsNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given ModelInvoiceResponsePaymentMethod and assigns it to the PaymentMethod field.
func (o *ModelInvoiceResponse) SetPaymentMethod(v ModelInvoiceResponsePaymentMethod) {
	o.PaymentMethod = &v
}

// GetCostCentre returns the CostCentre field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetCostCentre() ModelInvoiceResponseCostCentre {
	if o == nil || IsNil(o.CostCentre) {
		var ret ModelInvoiceResponseCostCentre
		return ret
	}
	return *o.CostCentre
}

// GetCostCentreOk returns a tuple with the CostCentre field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetCostCentreOk() (*ModelInvoiceResponseCostCentre, bool) {
	if o == nil || IsNil(o.CostCentre) {
		return nil, false
	}
	return o.CostCentre, true
}

// HasCostCentre returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasCostCentre() bool {
	if o != nil && !IsNil(o.CostCentre) {
		return true
	}

	return false
}

// SetCostCentre gets a reference to the given ModelInvoiceResponseCostCentre and assigns it to the CostCentre field.
func (o *ModelInvoiceResponse) SetCostCentre(v ModelInvoiceResponseCostCentre) {
	o.CostCentre = &v
}

// GetSendDate returns the SendDate field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetSendDate() time.Time {
	if o == nil || IsNil(o.SendDate) {
		var ret time.Time
		return ret
	}
	return *o.SendDate
}

// GetSendDateOk returns a tuple with the SendDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetSendDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SendDate) {
		return nil, false
	}
	return o.SendDate, true
}

// HasSendDate returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasSendDate() bool {
	if o != nil && !IsNil(o.SendDate) {
		return true
	}

	return false
}

// SetSendDate gets a reference to the given time.Time and assigns it to the SendDate field.
func (o *ModelInvoiceResponse) SetSendDate(v time.Time) {
	o.SendDate = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetOrigin() ModelInvoiceResponseOrigin {
	if o == nil || IsNil(o.Origin) {
		var ret ModelInvoiceResponseOrigin
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetOriginOk() (*ModelInvoiceResponseOrigin, bool) {
	if o == nil || IsNil(o.Origin) {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasOrigin() bool {
	if o != nil && !IsNil(o.Origin) {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given ModelInvoiceResponseOrigin and assigns it to the Origin field.
func (o *ModelInvoiceResponse) SetOrigin(v ModelInvoiceResponseOrigin) {
	o.Origin = &v
}

// GetInvoiceType returns the InvoiceType field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetInvoiceType() string {
	if o == nil || IsNil(o.InvoiceType) {
		var ret string
		return ret
	}
	return *o.InvoiceType
}

// GetInvoiceTypeOk returns a tuple with the InvoiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetInvoiceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceType) {
		return nil, false
	}
	return o.InvoiceType, true
}

// HasInvoiceType returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasInvoiceType() bool {
	if o != nil && !IsNil(o.InvoiceType) {
		return true
	}

	return false
}

// SetInvoiceType gets a reference to the given string and assigns it to the InvoiceType field.
func (o *ModelInvoiceResponse) SetInvoiceType(v string) {
	o.InvoiceType = &v
}

// GetAccountIntervall returns the AccountIntervall field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetAccountIntervall() string {
	if o == nil || IsNil(o.AccountIntervall) {
		var ret string
		return ret
	}
	return *o.AccountIntervall
}

// GetAccountIntervallOk returns a tuple with the AccountIntervall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetAccountIntervallOk() (*string, bool) {
	if o == nil || IsNil(o.AccountIntervall) {
		return nil, false
	}
	return o.AccountIntervall, true
}

// HasAccountIntervall returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasAccountIntervall() bool {
	if o != nil && !IsNil(o.AccountIntervall) {
		return true
	}

	return false
}

// SetAccountIntervall gets a reference to the given string and assigns it to the AccountIntervall field.
func (o *ModelInvoiceResponse) SetAccountIntervall(v string) {
	o.AccountIntervall = &v
}

// GetAccountNextInvoice returns the AccountNextInvoice field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetAccountNextInvoice() string {
	if o == nil || IsNil(o.AccountNextInvoice) {
		var ret string
		return ret
	}
	return *o.AccountNextInvoice
}

// GetAccountNextInvoiceOk returns a tuple with the AccountNextInvoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetAccountNextInvoiceOk() (*string, bool) {
	if o == nil || IsNil(o.AccountNextInvoice) {
		return nil, false
	}
	return o.AccountNextInvoice, true
}

// HasAccountNextInvoice returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasAccountNextInvoice() bool {
	if o != nil && !IsNil(o.AccountNextInvoice) {
		return true
	}

	return false
}

// SetAccountNextInvoice gets a reference to the given string and assigns it to the AccountNextInvoice field.
func (o *ModelInvoiceResponse) SetAccountNextInvoice(v string) {
	o.AccountNextInvoice = &v
}

// GetReminderTotal returns the ReminderTotal field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetReminderTotal() string {
	if o == nil || IsNil(o.ReminderTotal) {
		var ret string
		return ret
	}
	return *o.ReminderTotal
}

// GetReminderTotalOk returns a tuple with the ReminderTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetReminderTotalOk() (*string, bool) {
	if o == nil || IsNil(o.ReminderTotal) {
		return nil, false
	}
	return o.ReminderTotal, true
}

// HasReminderTotal returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasReminderTotal() bool {
	if o != nil && !IsNil(o.ReminderTotal) {
		return true
	}

	return false
}

// SetReminderTotal gets a reference to the given string and assigns it to the ReminderTotal field.
func (o *ModelInvoiceResponse) SetReminderTotal(v string) {
	o.ReminderTotal = &v
}

// GetReminderDebit returns the ReminderDebit field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetReminderDebit() string {
	if o == nil || IsNil(o.ReminderDebit) {
		var ret string
		return ret
	}
	return *o.ReminderDebit
}

// GetReminderDebitOk returns a tuple with the ReminderDebit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetReminderDebitOk() (*string, bool) {
	if o == nil || IsNil(o.ReminderDebit) {
		return nil, false
	}
	return o.ReminderDebit, true
}

// HasReminderDebit returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasReminderDebit() bool {
	if o != nil && !IsNil(o.ReminderDebit) {
		return true
	}

	return false
}

// SetReminderDebit gets a reference to the given string and assigns it to the ReminderDebit field.
func (o *ModelInvoiceResponse) SetReminderDebit(v string) {
	o.ReminderDebit = &v
}

// GetReminderDeadline returns the ReminderDeadline field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetReminderDeadline() time.Time {
	if o == nil || IsNil(o.ReminderDeadline) {
		var ret time.Time
		return ret
	}
	return *o.ReminderDeadline
}

// GetReminderDeadlineOk returns a tuple with the ReminderDeadline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetReminderDeadlineOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ReminderDeadline) {
		return nil, false
	}
	return o.ReminderDeadline, true
}

// HasReminderDeadline returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasReminderDeadline() bool {
	if o != nil && !IsNil(o.ReminderDeadline) {
		return true
	}

	return false
}

// SetReminderDeadline gets a reference to the given time.Time and assigns it to the ReminderDeadline field.
func (o *ModelInvoiceResponse) SetReminderDeadline(v time.Time) {
	o.ReminderDeadline = &v
}

// GetReminderCharge returns the ReminderCharge field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetReminderCharge() string {
	if o == nil || IsNil(o.ReminderCharge) {
		var ret string
		return ret
	}
	return *o.ReminderCharge
}

// GetReminderChargeOk returns a tuple with the ReminderCharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetReminderChargeOk() (*string, bool) {
	if o == nil || IsNil(o.ReminderCharge) {
		return nil, false
	}
	return o.ReminderCharge, true
}

// HasReminderCharge returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasReminderCharge() bool {
	if o != nil && !IsNil(o.ReminderCharge) {
		return true
	}

	return false
}

// SetReminderCharge gets a reference to the given string and assigns it to the ReminderCharge field.
func (o *ModelInvoiceResponse) SetReminderCharge(v string) {
	o.ReminderCharge = &v
}

// GetTaxSet returns the TaxSet field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetTaxSet() ModelInvoiceResponseTaxSet {
	if o == nil || IsNil(o.TaxSet) {
		var ret ModelInvoiceResponseTaxSet
		return ret
	}
	return *o.TaxSet
}

// GetTaxSetOk returns a tuple with the TaxSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetTaxSetOk() (*ModelInvoiceResponseTaxSet, bool) {
	if o == nil || IsNil(o.TaxSet) {
		return nil, false
	}
	return o.TaxSet, true
}

// HasTaxSet returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasTaxSet() bool {
	if o != nil && !IsNil(o.TaxSet) {
		return true
	}

	return false
}

// SetTaxSet gets a reference to the given ModelInvoiceResponseTaxSet and assigns it to the TaxSet field.
func (o *ModelInvoiceResponse) SetTaxSet(v ModelInvoiceResponseTaxSet) {
	o.TaxSet = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *ModelInvoiceResponse) SetAddress(v string) {
	o.Address = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *ModelInvoiceResponse) SetCurrency(v string) {
	o.Currency = &v
}

// GetSumNet returns the SumNet field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetSumNet() string {
	if o == nil || IsNil(o.SumNet) {
		var ret string
		return ret
	}
	return *o.SumNet
}

// GetSumNetOk returns a tuple with the SumNet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetSumNetOk() (*string, bool) {
	if o == nil || IsNil(o.SumNet) {
		return nil, false
	}
	return o.SumNet, true
}

// HasSumNet returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasSumNet() bool {
	if o != nil && !IsNil(o.SumNet) {
		return true
	}

	return false
}

// SetSumNet gets a reference to the given string and assigns it to the SumNet field.
func (o *ModelInvoiceResponse) SetSumNet(v string) {
	o.SumNet = &v
}

// GetSumTax returns the SumTax field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetSumTax() string {
	if o == nil || IsNil(o.SumTax) {
		var ret string
		return ret
	}
	return *o.SumTax
}

// GetSumTaxOk returns a tuple with the SumTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetSumTaxOk() (*string, bool) {
	if o == nil || IsNil(o.SumTax) {
		return nil, false
	}
	return o.SumTax, true
}

// HasSumTax returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasSumTax() bool {
	if o != nil && !IsNil(o.SumTax) {
		return true
	}

	return false
}

// SetSumTax gets a reference to the given string and assigns it to the SumTax field.
func (o *ModelInvoiceResponse) SetSumTax(v string) {
	o.SumTax = &v
}

// GetSumGross returns the SumGross field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetSumGross() string {
	if o == nil || IsNil(o.SumGross) {
		var ret string
		return ret
	}
	return *o.SumGross
}

// GetSumGrossOk returns a tuple with the SumGross field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetSumGrossOk() (*string, bool) {
	if o == nil || IsNil(o.SumGross) {
		return nil, false
	}
	return o.SumGross, true
}

// HasSumGross returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasSumGross() bool {
	if o != nil && !IsNil(o.SumGross) {
		return true
	}

	return false
}

// SetSumGross gets a reference to the given string and assigns it to the SumGross field.
func (o *ModelInvoiceResponse) SetSumGross(v string) {
	o.SumGross = &v
}

// GetSumDiscounts returns the SumDiscounts field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetSumDiscounts() string {
	if o == nil || IsNil(o.SumDiscounts) {
		var ret string
		return ret
	}
	return *o.SumDiscounts
}

// GetSumDiscountsOk returns a tuple with the SumDiscounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetSumDiscountsOk() (*string, bool) {
	if o == nil || IsNil(o.SumDiscounts) {
		return nil, false
	}
	return o.SumDiscounts, true
}

// HasSumDiscounts returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasSumDiscounts() bool {
	if o != nil && !IsNil(o.SumDiscounts) {
		return true
	}

	return false
}

// SetSumDiscounts gets a reference to the given string and assigns it to the SumDiscounts field.
func (o *ModelInvoiceResponse) SetSumDiscounts(v string) {
	o.SumDiscounts = &v
}

// GetSumNetForeignCurrency returns the SumNetForeignCurrency field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetSumNetForeignCurrency() string {
	if o == nil || IsNil(o.SumNetForeignCurrency) {
		var ret string
		return ret
	}
	return *o.SumNetForeignCurrency
}

// GetSumNetForeignCurrencyOk returns a tuple with the SumNetForeignCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetSumNetForeignCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.SumNetForeignCurrency) {
		return nil, false
	}
	return o.SumNetForeignCurrency, true
}

// HasSumNetForeignCurrency returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasSumNetForeignCurrency() bool {
	if o != nil && !IsNil(o.SumNetForeignCurrency) {
		return true
	}

	return false
}

// SetSumNetForeignCurrency gets a reference to the given string and assigns it to the SumNetForeignCurrency field.
func (o *ModelInvoiceResponse) SetSumNetForeignCurrency(v string) {
	o.SumNetForeignCurrency = &v
}

// GetSumTaxForeignCurrency returns the SumTaxForeignCurrency field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetSumTaxForeignCurrency() string {
	if o == nil || IsNil(o.SumTaxForeignCurrency) {
		var ret string
		return ret
	}
	return *o.SumTaxForeignCurrency
}

// GetSumTaxForeignCurrencyOk returns a tuple with the SumTaxForeignCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetSumTaxForeignCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.SumTaxForeignCurrency) {
		return nil, false
	}
	return o.SumTaxForeignCurrency, true
}

// HasSumTaxForeignCurrency returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasSumTaxForeignCurrency() bool {
	if o != nil && !IsNil(o.SumTaxForeignCurrency) {
		return true
	}

	return false
}

// SetSumTaxForeignCurrency gets a reference to the given string and assigns it to the SumTaxForeignCurrency field.
func (o *ModelInvoiceResponse) SetSumTaxForeignCurrency(v string) {
	o.SumTaxForeignCurrency = &v
}

// GetSumGrossForeignCurrency returns the SumGrossForeignCurrency field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetSumGrossForeignCurrency() string {
	if o == nil || IsNil(o.SumGrossForeignCurrency) {
		var ret string
		return ret
	}
	return *o.SumGrossForeignCurrency
}

// GetSumGrossForeignCurrencyOk returns a tuple with the SumGrossForeignCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetSumGrossForeignCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.SumGrossForeignCurrency) {
		return nil, false
	}
	return o.SumGrossForeignCurrency, true
}

// HasSumGrossForeignCurrency returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasSumGrossForeignCurrency() bool {
	if o != nil && !IsNil(o.SumGrossForeignCurrency) {
		return true
	}

	return false
}

// SetSumGrossForeignCurrency gets a reference to the given string and assigns it to the SumGrossForeignCurrency field.
func (o *ModelInvoiceResponse) SetSumGrossForeignCurrency(v string) {
	o.SumGrossForeignCurrency = &v
}

// GetSumDiscountsForeignCurrency returns the SumDiscountsForeignCurrency field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetSumDiscountsForeignCurrency() string {
	if o == nil || IsNil(o.SumDiscountsForeignCurrency) {
		var ret string
		return ret
	}
	return *o.SumDiscountsForeignCurrency
}

// GetSumDiscountsForeignCurrencyOk returns a tuple with the SumDiscountsForeignCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetSumDiscountsForeignCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.SumDiscountsForeignCurrency) {
		return nil, false
	}
	return o.SumDiscountsForeignCurrency, true
}

// HasSumDiscountsForeignCurrency returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasSumDiscountsForeignCurrency() bool {
	if o != nil && !IsNil(o.SumDiscountsForeignCurrency) {
		return true
	}

	return false
}

// SetSumDiscountsForeignCurrency gets a reference to the given string and assigns it to the SumDiscountsForeignCurrency field.
func (o *ModelInvoiceResponse) SetSumDiscountsForeignCurrency(v string) {
	o.SumDiscountsForeignCurrency = &v
}

// GetSumNetAccounting returns the SumNetAccounting field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetSumNetAccounting() string {
	if o == nil || IsNil(o.SumNetAccounting) {
		var ret string
		return ret
	}
	return *o.SumNetAccounting
}

// GetSumNetAccountingOk returns a tuple with the SumNetAccounting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetSumNetAccountingOk() (*string, bool) {
	if o == nil || IsNil(o.SumNetAccounting) {
		return nil, false
	}
	return o.SumNetAccounting, true
}

// HasSumNetAccounting returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasSumNetAccounting() bool {
	if o != nil && !IsNil(o.SumNetAccounting) {
		return true
	}

	return false
}

// SetSumNetAccounting gets a reference to the given string and assigns it to the SumNetAccounting field.
func (o *ModelInvoiceResponse) SetSumNetAccounting(v string) {
	o.SumNetAccounting = &v
}

// GetSumTaxAccounting returns the SumTaxAccounting field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetSumTaxAccounting() string {
	if o == nil || IsNil(o.SumTaxAccounting) {
		var ret string
		return ret
	}
	return *o.SumTaxAccounting
}

// GetSumTaxAccountingOk returns a tuple with the SumTaxAccounting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetSumTaxAccountingOk() (*string, bool) {
	if o == nil || IsNil(o.SumTaxAccounting) {
		return nil, false
	}
	return o.SumTaxAccounting, true
}

// HasSumTaxAccounting returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasSumTaxAccounting() bool {
	if o != nil && !IsNil(o.SumTaxAccounting) {
		return true
	}

	return false
}

// SetSumTaxAccounting gets a reference to the given string and assigns it to the SumTaxAccounting field.
func (o *ModelInvoiceResponse) SetSumTaxAccounting(v string) {
	o.SumTaxAccounting = &v
}

// GetSumGrossAccounting returns the SumGrossAccounting field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetSumGrossAccounting() string {
	if o == nil || IsNil(o.SumGrossAccounting) {
		var ret string
		return ret
	}
	return *o.SumGrossAccounting
}

// GetSumGrossAccountingOk returns a tuple with the SumGrossAccounting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetSumGrossAccountingOk() (*string, bool) {
	if o == nil || IsNil(o.SumGrossAccounting) {
		return nil, false
	}
	return o.SumGrossAccounting, true
}

// HasSumGrossAccounting returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasSumGrossAccounting() bool {
	if o != nil && !IsNil(o.SumGrossAccounting) {
		return true
	}

	return false
}

// SetSumGrossAccounting gets a reference to the given string and assigns it to the SumGrossAccounting field.
func (o *ModelInvoiceResponse) SetSumGrossAccounting(v string) {
	o.SumGrossAccounting = &v
}

// GetPaidAmount returns the PaidAmount field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetPaidAmount() float32 {
	if o == nil || IsNil(o.PaidAmount) {
		var ret float32
		return ret
	}
	return *o.PaidAmount
}

// GetPaidAmountOk returns a tuple with the PaidAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetPaidAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.PaidAmount) {
		return nil, false
	}
	return o.PaidAmount, true
}

// HasPaidAmount returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasPaidAmount() bool {
	if o != nil && !IsNil(o.PaidAmount) {
		return true
	}

	return false
}

// SetPaidAmount gets a reference to the given float32 and assigns it to the PaidAmount field.
func (o *ModelInvoiceResponse) SetPaidAmount(v float32) {
	o.PaidAmount = &v
}

// GetCustomerInternalNote returns the CustomerInternalNote field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetCustomerInternalNote() string {
	if o == nil || IsNil(o.CustomerInternalNote) {
		var ret string
		return ret
	}
	return *o.CustomerInternalNote
}

// GetCustomerInternalNoteOk returns a tuple with the CustomerInternalNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetCustomerInternalNoteOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerInternalNote) {
		return nil, false
	}
	return o.CustomerInternalNote, true
}

// HasCustomerInternalNote returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasCustomerInternalNote() bool {
	if o != nil && !IsNil(o.CustomerInternalNote) {
		return true
	}

	return false
}

// SetCustomerInternalNote gets a reference to the given string and assigns it to the CustomerInternalNote field.
func (o *ModelInvoiceResponse) SetCustomerInternalNote(v string) {
	o.CustomerInternalNote = &v
}

// GetShowNet returns the ShowNet field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetShowNet() bool {
	if o == nil || IsNil(o.ShowNet) {
		var ret bool
		return ret
	}
	return *o.ShowNet
}

// GetShowNetOk returns a tuple with the ShowNet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetShowNetOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowNet) {
		return nil, false
	}
	return o.ShowNet, true
}

// HasShowNet returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasShowNet() bool {
	if o != nil && !IsNil(o.ShowNet) {
		return true
	}

	return false
}

// SetShowNet gets a reference to the given bool and assigns it to the ShowNet field.
func (o *ModelInvoiceResponse) SetShowNet(v bool) {
	o.ShowNet = &v
}

// GetEnshrined returns the Enshrined field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetEnshrined() time.Time {
	if o == nil || IsNil(o.Enshrined) {
		var ret time.Time
		return ret
	}
	return *o.Enshrined
}

// GetEnshrinedOk returns a tuple with the Enshrined field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetEnshrinedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Enshrined) {
		return nil, false
	}
	return o.Enshrined, true
}

// HasEnshrined returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasEnshrined() bool {
	if o != nil && !IsNil(o.Enshrined) {
		return true
	}

	return false
}

// SetEnshrined gets a reference to the given time.Time and assigns it to the Enshrined field.
func (o *ModelInvoiceResponse) SetEnshrined(v time.Time) {
	o.Enshrined = &v
}

// GetSendType returns the SendType field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetSendType() string {
	if o == nil || IsNil(o.SendType) {
		var ret string
		return ret
	}
	return *o.SendType
}

// GetSendTypeOk returns a tuple with the SendType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetSendTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SendType) {
		return nil, false
	}
	return o.SendType, true
}

// HasSendType returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasSendType() bool {
	if o != nil && !IsNil(o.SendType) {
		return true
	}

	return false
}

// SetSendType gets a reference to the given string and assigns it to the SendType field.
func (o *ModelInvoiceResponse) SetSendType(v string) {
	o.SendType = &v
}

// GetDeliveryDateUntil returns the DeliveryDateUntil field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetDeliveryDateUntil() string {
	if o == nil || IsNil(o.DeliveryDateUntil) {
		var ret string
		return ret
	}
	return *o.DeliveryDateUntil
}

// GetDeliveryDateUntilOk returns a tuple with the DeliveryDateUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetDeliveryDateUntilOk() (*string, bool) {
	if o == nil || IsNil(o.DeliveryDateUntil) {
		return nil, false
	}
	return o.DeliveryDateUntil, true
}

// HasDeliveryDateUntil returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasDeliveryDateUntil() bool {
	if o != nil && !IsNil(o.DeliveryDateUntil) {
		return true
	}

	return false
}

// SetDeliveryDateUntil gets a reference to the given string and assigns it to the DeliveryDateUntil field.
func (o *ModelInvoiceResponse) SetDeliveryDateUntil(v string) {
	o.DeliveryDateUntil = &v
}

// GetDatevConnectOnline returns the DatevConnectOnline field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetDatevConnectOnline() map[string]interface{} {
	if o == nil || IsNil(o.DatevConnectOnline) {
		var ret map[string]interface{}
		return ret
	}
	return o.DatevConnectOnline
}

// GetDatevConnectOnlineOk returns a tuple with the DatevConnectOnline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetDatevConnectOnlineOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.DatevConnectOnline) {
		return map[string]interface{}{}, false
	}
	return o.DatevConnectOnline, true
}

// HasDatevConnectOnline returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasDatevConnectOnline() bool {
	if o != nil && !IsNil(o.DatevConnectOnline) {
		return true
	}

	return false
}

// SetDatevConnectOnline gets a reference to the given map[string]interface{} and assigns it to the DatevConnectOnline field.
func (o *ModelInvoiceResponse) SetDatevConnectOnline(v map[string]interface{}) {
	o.DatevConnectOnline = v
}

// GetSendPaymentReceivedNotificationDate returns the SendPaymentReceivedNotificationDate field value if set, zero value otherwise.
func (o *ModelInvoiceResponse) GetSendPaymentReceivedNotificationDate() string {
	if o == nil || IsNil(o.SendPaymentReceivedNotificationDate) {
		var ret string
		return ret
	}
	return *o.SendPaymentReceivedNotificationDate
}

// GetSendPaymentReceivedNotificationDateOk returns a tuple with the SendPaymentReceivedNotificationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceResponse) GetSendPaymentReceivedNotificationDateOk() (*string, bool) {
	if o == nil || IsNil(o.SendPaymentReceivedNotificationDate) {
		return nil, false
	}
	return o.SendPaymentReceivedNotificationDate, true
}

// HasSendPaymentReceivedNotificationDate returns a boolean if a field has been set.
func (o *ModelInvoiceResponse) HasSendPaymentReceivedNotificationDate() bool {
	if o != nil && !IsNil(o.SendPaymentReceivedNotificationDate) {
		return true
	}

	return false
}

// SetSendPaymentReceivedNotificationDate gets a reference to the given string and assigns it to the SendPaymentReceivedNotificationDate field.
func (o *ModelInvoiceResponse) SetSendPaymentReceivedNotificationDate(v string) {
	o.SendPaymentReceivedNotificationDate = &v
}

func (o ModelInvoiceResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelInvoiceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	// skip: objectName is readOnly
	// skip: invoiceNumber is readOnly
	if !IsNil(o.Contact) {
		toSerialize["contact"] = o.Contact
	}
	// skip: create is readOnly
	// skip: update is readOnly
	if !IsNil(o.SevClient) {
		toSerialize["sevClient"] = o.SevClient
	}
	// skip: invoiceDate is readOnly
	// skip: header is readOnly
	// skip: headText is readOnly
	// skip: footText is readOnly
	// skip: timeToPay is readOnly
	// skip: discountTime is readOnly
	// skip: discount is readOnly
	if !IsNil(o.AddressCountry) {
		toSerialize["addressCountry"] = o.AddressCountry
	}
	// skip: payDate is readOnly
	if !IsNil(o.CreateUser) {
		toSerialize["createUser"] = o.CreateUser
	}
	// skip: deliveryDate is readOnly
	// skip: status is readOnly
	// skip: smallSettlement is readOnly
	if !IsNil(o.ContactPerson) {
		toSerialize["contactPerson"] = o.ContactPerson
	}
	// skip: taxRate is readOnly
	// skip: taxText is readOnly
	// skip: dunningLevel is readOnly
	// skip: taxType is readOnly
	if !IsNil(o.PaymentMethod) {
		toSerialize["paymentMethod"] = o.PaymentMethod
	}
	if !IsNil(o.CostCentre) {
		toSerialize["costCentre"] = o.CostCentre
	}
	// skip: sendDate is readOnly
	if !IsNil(o.Origin) {
		toSerialize["origin"] = o.Origin
	}
	// skip: invoiceType is readOnly
	// skip: accountIntervall is readOnly
	// skip: accountNextInvoice is readOnly
	// skip: reminderTotal is readOnly
	// skip: reminderDebit is readOnly
	// skip: reminderDeadline is readOnly
	// skip: reminderCharge is readOnly
	if !IsNil(o.TaxSet) {
		toSerialize["taxSet"] = o.TaxSet
	}
	// skip: address is readOnly
	// skip: currency is readOnly
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
	// skip: paidAmount is readOnly
	// skip: customerInternalNote is readOnly
	// skip: showNet is readOnly
	// skip: enshrined is readOnly
	// skip: sendType is readOnly
	// skip: deliveryDateUntil is readOnly
	// skip: datevConnectOnline is readOnly
	// skip: sendPaymentReceivedNotificationDate is readOnly
	return toSerialize, nil
}

type NullableModelInvoiceResponse struct {
	value *ModelInvoiceResponse
	isSet bool
}

func (v NullableModelInvoiceResponse) Get() *ModelInvoiceResponse {
	return v.value
}

func (v *NullableModelInvoiceResponse) Set(val *ModelInvoiceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModelInvoiceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModelInvoiceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelInvoiceResponse(val *ModelInvoiceResponse) *NullableModelInvoiceResponse {
	return &NullableModelInvoiceResponse{value: val, isSet: true}
}

func (v NullableModelInvoiceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelInvoiceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


