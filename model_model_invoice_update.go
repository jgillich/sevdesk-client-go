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

// checks if the ModelInvoiceUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelInvoiceUpdate{}

// ModelInvoiceUpdate Invoice model
type ModelInvoiceUpdate struct {
	// The invoice id
	Id *int32 `json:"id,omitempty"`
	// The invoice object name
	ObjectName *string `json:"objectName,omitempty"`
	// The invoice number
	InvoiceNumber NullableString `json:"invoiceNumber,omitempty"`
	Contact NullableModelInvoiceUpdateContact `json:"contact,omitempty"`
	// Date of invoice creation
	Create *time.Time `json:"create,omitempty"`
	// Date of last invoice update
	Update *time.Time `json:"update,omitempty"`
	SevClient *ModelInvoiceUpdateSevClient `json:"sevClient,omitempty"`
	// Needs to be provided as timestamp or dd.mm.yyyy
	InvoiceDate NullableString `json:"invoiceDate,omitempty"`
	// Normally consist of prefix plus the invoice number
	Header NullableString `json:"header,omitempty"`
	// Certain html tags can be used here to format your text
	HeadText NullableString `json:"headText,omitempty"`
	// Certain html tags can be used here to format your text
	FootText NullableString `json:"footText,omitempty"`
	// The time the customer has to pay the invoice in days
	TimeToPay NullableInt32 `json:"timeToPay,omitempty"`
	// If a value other than zero is used for the discount attribute,      you need to specify the amount of days for which the discount is granted.
	DiscountTime NullableInt32 `json:"discountTime,omitempty"`
	// If you want to give a discount, define the percentage here. Otherwise provide zero as value
	Discount NullableInt32 `json:"discount,omitempty"`
	AddressCountry NullableModelCreditNoteAddressCountry `json:"addressCountry,omitempty"`
	// Needs to be timestamp or dd.mm.yyyy
	PayDate NullableTime `json:"payDate,omitempty"`
	CreateUser *ModelCreditNoteCreateUser `json:"createUser,omitempty"`
	// Timestamp. This can also be a date range if you also use the attribute deliveryDateUntil
	DeliveryDate NullableTime `json:"deliveryDate,omitempty"`
	// Please have a look in our       <a href='https://api.sevdesk.de/#section/Types-and-status-of-invoices'>Types and status of invoice</a>       to see what the different status codes mean
	Status NullableString `json:"status,omitempty"`
	// Defines if the client uses the small settlement scheme.      If yes, the invoice must not contain any vat
	SmallSettlement NullableBool `json:"smallSettlement,omitempty"`
	ContactPerson *ModelInvoiceUpdateContactPerson `json:"contactPerson,omitempty"`
	// Is overwritten by invoice position tax rates
	TaxRate NullableFloat32 `json:"taxRate,omitempty"`
	// A common tax text would be 'Umsatzsteuer 19%'
	TaxText NullableString `json:"taxText,omitempty"`
	// Defines how many reminders have already been sent for the invoice.      Starts with 1 (Payment reminder) and should be incremented by one every time another reminder is sent.
	DunningLevel NullableInt32 `json:"dunningLevel,omitempty"`
	// Tax type of the invoice. There are four tax types: 1. default - Umsatzsteuer ausweisen 2. eu - Steuerfreie innergemeinschaftliche Lieferung (Europäische Union) 3. noteu - Steuerschuldnerschaft des Leistungsempfängers (außerhalb EU, z. B. Schweiz) 4. custom - Using custom tax set 5. ss - Kleinunternehmer Tax rates are heavily connected to the tax type used.
	TaxType NullableString `json:"taxType,omitempty"`
	PaymentMethod *ModelInvoiceUpdatePaymentMethod `json:"paymentMethod,omitempty"`
	CostCentre *ModelInvoiceUpdateCostCentre `json:"costCentre,omitempty"`
	// The date the invoice was sent to the customer
	SendDate NullableTime `json:"sendDate,omitempty"`
	Origin NullableModelInvoiceUpdateOrigin `json:"origin,omitempty"`
	// Type of the invoice. For more information on the different types, check       <a href='https://api.sevdesk.de/#section/Types-and-status-of-invoices'>this</a> section  
	InvoiceType NullableString `json:"invoiceType,omitempty"`
	// The interval in which recurring invoices are due as ISO-8601 duration.<br>       Necessary attribute for all recurring invoices.
	AccountIntervall NullableString `json:"accountIntervall,omitempty"`
	// Timestamp when the next invoice will be generated by this recurring invoice.
	AccountNextInvoice NullableInt32 `json:"accountNextInvoice,omitempty"`
	// Total reminder amount
	ReminderTotal NullableFloat32 `json:"reminderTotal,omitempty"`
	// Debit of the reminder
	ReminderDebit NullableFloat32 `json:"reminderDebit,omitempty"`
	// Deadline of the reminder as timestamp
	ReminderDeadline NullableInt32 `json:"reminderDeadline,omitempty"`
	// The additional reminder charge
	ReminderCharge NullableFloat32 `json:"reminderCharge,omitempty"`
	TaxSet NullableModelInvoiceUpdateTaxSet `json:"taxSet,omitempty"`
	// Complete address of the recipient including name, street, city, zip and country.       * Line breaks can be used and will be displayed on the invoice pdf.
	Address NullableString `json:"address,omitempty"`
	// Currency used in the invoice. Needs to be currency code according to ISO-4217
	Currency NullableString `json:"currency,omitempty"`
	// Net sum of the invoice
	SumNet NullableFloat32 `json:"sumNet,omitempty"`
	// Tax sum of the invoice
	SumTax NullableFloat32 `json:"sumTax,omitempty"`
	// Gross sum of the invoice
	SumGross NullableFloat32 `json:"sumGross,omitempty"`
	// Sum of all discounts in the invoice
	SumDiscounts NullableFloat32 `json:"sumDiscounts,omitempty"`
	// Net sum of the invoice in the foreign currency
	SumNetForeignCurrency NullableFloat32 `json:"sumNetForeignCurrency,omitempty"`
	// Tax sum of the invoice in the foreign currency
	SumTaxForeignCurrency NullableFloat32 `json:"sumTaxForeignCurrency,omitempty"`
	// Gross sum of the invoice in the foreign currency
	SumGrossForeignCurrency NullableFloat32 `json:"sumGrossForeignCurrency,omitempty"`
	// Discounts sum of the invoice in the foreign currency
	SumDiscountsForeignCurrency NullableFloat32 `json:"sumDiscountsForeignCurrency,omitempty"`
	// Net accounting sum of the invoice. Is usually the same as sumNet
	SumNetAccounting NullableFloat32 `json:"sumNetAccounting,omitempty"`
	// Tax accounting sum of the invoice. Is usually the same as sumTax
	SumTaxAccounting NullableFloat32 `json:"sumTaxAccounting,omitempty"`
	// Gross accounting sum of the invoice. Is usually the same as sumGross
	SumGrossAccounting NullableFloat32 `json:"sumGrossAccounting,omitempty"`
	// Amount which has already been paid for this invoice by the customer
	PaidAmount NullableFloat32 `json:"paidAmount,omitempty"`
	// Internal note of the customer. Contains data entered into field 'Referenz/Bestellnummer'
	CustomerInternalNote NullableString `json:"customerInternalNote,omitempty"`
	// If true, the net amount of each position will be shown on the invoice. Otherwise gross amount
	ShowNet NullableBool `json:"showNet,omitempty"`
	// Defines if and when invoice was enshrined. Enshrined invoices can not be manipulated.
	Enshrined NullableTime `json:"enshrined,omitempty"`
	// Type which was used to send the invoice. IMPORTANT: Please refer to the invoice section of the       *     API-Overview to understand how this attribute can be used before using it!
	SendType NullableString `json:"sendType,omitempty"`
	// If the delivery date should be a time range, another timestamp can be provided in this attribute       * to define a range from timestamp used in deliveryDate attribute to the timestamp used here.
	DeliveryDateUntil NullableInt32 `json:"deliveryDateUntil,omitempty"`
	// Internal attribute
	DatevConnectOnline map[string]interface{} `json:"datevConnectOnline,omitempty"`
	// Internal attribute
	SendPaymentReceivedNotificationDate NullableInt32 `json:"sendPaymentReceivedNotificationDate,omitempty"`
}

// NewModelInvoiceUpdate instantiates a new ModelInvoiceUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelInvoiceUpdate() *ModelInvoiceUpdate {
	this := ModelInvoiceUpdate{}
	return &this
}

// NewModelInvoiceUpdateWithDefaults instantiates a new ModelInvoiceUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelInvoiceUpdateWithDefaults() *ModelInvoiceUpdate {
	this := ModelInvoiceUpdate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelInvoiceUpdate) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceUpdate) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ModelInvoiceUpdate) SetId(v int32) {
	o.Id = &v
}

// GetObjectName returns the ObjectName field value if set, zero value otherwise.
func (o *ModelInvoiceUpdate) GetObjectName() string {
	if o == nil || IsNil(o.ObjectName) {
		var ret string
		return ret
	}
	return *o.ObjectName
}

// GetObjectNameOk returns a tuple with the ObjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceUpdate) GetObjectNameOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectName) {
		return nil, false
	}
	return o.ObjectName, true
}

// HasObjectName returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasObjectName() bool {
	if o != nil && !IsNil(o.ObjectName) {
		return true
	}

	return false
}

// SetObjectName gets a reference to the given string and assigns it to the ObjectName field.
func (o *ModelInvoiceUpdate) SetObjectName(v string) {
	o.ObjectName = &v
}

// GetInvoiceNumber returns the InvoiceNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetInvoiceNumber() string {
	if o == nil || IsNil(o.InvoiceNumber.Get()) {
		var ret string
		return ret
	}
	return *o.InvoiceNumber.Get()
}

// GetInvoiceNumberOk returns a tuple with the InvoiceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetInvoiceNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceNumber.Get(), o.InvoiceNumber.IsSet()
}

// HasInvoiceNumber returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasInvoiceNumber() bool {
	if o != nil && o.InvoiceNumber.IsSet() {
		return true
	}

	return false
}

// SetInvoiceNumber gets a reference to the given NullableString and assigns it to the InvoiceNumber field.
func (o *ModelInvoiceUpdate) SetInvoiceNumber(v string) {
	o.InvoiceNumber.Set(&v)
}
// SetInvoiceNumberNil sets the value for InvoiceNumber to be an explicit nil
func (o *ModelInvoiceUpdate) SetInvoiceNumberNil() {
	o.InvoiceNumber.Set(nil)
}

// UnsetInvoiceNumber ensures that no value is present for InvoiceNumber, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetInvoiceNumber() {
	o.InvoiceNumber.Unset()
}

// GetContact returns the Contact field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetContact() ModelInvoiceUpdateContact {
	if o == nil || IsNil(o.Contact.Get()) {
		var ret ModelInvoiceUpdateContact
		return ret
	}
	return *o.Contact.Get()
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetContactOk() (*ModelInvoiceUpdateContact, bool) {
	if o == nil {
		return nil, false
	}
	return o.Contact.Get(), o.Contact.IsSet()
}

// HasContact returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasContact() bool {
	if o != nil && o.Contact.IsSet() {
		return true
	}

	return false
}

// SetContact gets a reference to the given NullableModelInvoiceUpdateContact and assigns it to the Contact field.
func (o *ModelInvoiceUpdate) SetContact(v ModelInvoiceUpdateContact) {
	o.Contact.Set(&v)
}
// SetContactNil sets the value for Contact to be an explicit nil
func (o *ModelInvoiceUpdate) SetContactNil() {
	o.Contact.Set(nil)
}

// UnsetContact ensures that no value is present for Contact, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetContact() {
	o.Contact.Unset()
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *ModelInvoiceUpdate) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceUpdate) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *ModelInvoiceUpdate) SetCreate(v time.Time) {
	o.Create = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *ModelInvoiceUpdate) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceUpdate) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *ModelInvoiceUpdate) SetUpdate(v time.Time) {
	o.Update = &v
}

// GetSevClient returns the SevClient field value if set, zero value otherwise.
func (o *ModelInvoiceUpdate) GetSevClient() ModelInvoiceUpdateSevClient {
	if o == nil || IsNil(o.SevClient) {
		var ret ModelInvoiceUpdateSevClient
		return ret
	}
	return *o.SevClient
}

// GetSevClientOk returns a tuple with the SevClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceUpdate) GetSevClientOk() (*ModelInvoiceUpdateSevClient, bool) {
	if o == nil || IsNil(o.SevClient) {
		return nil, false
	}
	return o.SevClient, true
}

// HasSevClient returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasSevClient() bool {
	if o != nil && !IsNil(o.SevClient) {
		return true
	}

	return false
}

// SetSevClient gets a reference to the given ModelInvoiceUpdateSevClient and assigns it to the SevClient field.
func (o *ModelInvoiceUpdate) SetSevClient(v ModelInvoiceUpdateSevClient) {
	o.SevClient = &v
}

// GetInvoiceDate returns the InvoiceDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetInvoiceDate() string {
	if o == nil || IsNil(o.InvoiceDate.Get()) {
		var ret string
		return ret
	}
	return *o.InvoiceDate.Get()
}

// GetInvoiceDateOk returns a tuple with the InvoiceDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetInvoiceDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceDate.Get(), o.InvoiceDate.IsSet()
}

// HasInvoiceDate returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasInvoiceDate() bool {
	if o != nil && o.InvoiceDate.IsSet() {
		return true
	}

	return false
}

// SetInvoiceDate gets a reference to the given NullableString and assigns it to the InvoiceDate field.
func (o *ModelInvoiceUpdate) SetInvoiceDate(v string) {
	o.InvoiceDate.Set(&v)
}
// SetInvoiceDateNil sets the value for InvoiceDate to be an explicit nil
func (o *ModelInvoiceUpdate) SetInvoiceDateNil() {
	o.InvoiceDate.Set(nil)
}

// UnsetInvoiceDate ensures that no value is present for InvoiceDate, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetInvoiceDate() {
	o.InvoiceDate.Unset()
}

// GetHeader returns the Header field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetHeader() string {
	if o == nil || IsNil(o.Header.Get()) {
		var ret string
		return ret
	}
	return *o.Header.Get()
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetHeaderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Header.Get(), o.Header.IsSet()
}

// HasHeader returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasHeader() bool {
	if o != nil && o.Header.IsSet() {
		return true
	}

	return false
}

// SetHeader gets a reference to the given NullableString and assigns it to the Header field.
func (o *ModelInvoiceUpdate) SetHeader(v string) {
	o.Header.Set(&v)
}
// SetHeaderNil sets the value for Header to be an explicit nil
func (o *ModelInvoiceUpdate) SetHeaderNil() {
	o.Header.Set(nil)
}

// UnsetHeader ensures that no value is present for Header, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetHeader() {
	o.Header.Unset()
}

// GetHeadText returns the HeadText field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetHeadText() string {
	if o == nil || IsNil(o.HeadText.Get()) {
		var ret string
		return ret
	}
	return *o.HeadText.Get()
}

// GetHeadTextOk returns a tuple with the HeadText field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetHeadTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HeadText.Get(), o.HeadText.IsSet()
}

// HasHeadText returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasHeadText() bool {
	if o != nil && o.HeadText.IsSet() {
		return true
	}

	return false
}

// SetHeadText gets a reference to the given NullableString and assigns it to the HeadText field.
func (o *ModelInvoiceUpdate) SetHeadText(v string) {
	o.HeadText.Set(&v)
}
// SetHeadTextNil sets the value for HeadText to be an explicit nil
func (o *ModelInvoiceUpdate) SetHeadTextNil() {
	o.HeadText.Set(nil)
}

// UnsetHeadText ensures that no value is present for HeadText, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetHeadText() {
	o.HeadText.Unset()
}

// GetFootText returns the FootText field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetFootText() string {
	if o == nil || IsNil(o.FootText.Get()) {
		var ret string
		return ret
	}
	return *o.FootText.Get()
}

// GetFootTextOk returns a tuple with the FootText field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetFootTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FootText.Get(), o.FootText.IsSet()
}

// HasFootText returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasFootText() bool {
	if o != nil && o.FootText.IsSet() {
		return true
	}

	return false
}

// SetFootText gets a reference to the given NullableString and assigns it to the FootText field.
func (o *ModelInvoiceUpdate) SetFootText(v string) {
	o.FootText.Set(&v)
}
// SetFootTextNil sets the value for FootText to be an explicit nil
func (o *ModelInvoiceUpdate) SetFootTextNil() {
	o.FootText.Set(nil)
}

// UnsetFootText ensures that no value is present for FootText, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetFootText() {
	o.FootText.Unset()
}

// GetTimeToPay returns the TimeToPay field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetTimeToPay() int32 {
	if o == nil || IsNil(o.TimeToPay.Get()) {
		var ret int32
		return ret
	}
	return *o.TimeToPay.Get()
}

// GetTimeToPayOk returns a tuple with the TimeToPay field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetTimeToPayOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimeToPay.Get(), o.TimeToPay.IsSet()
}

// HasTimeToPay returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasTimeToPay() bool {
	if o != nil && o.TimeToPay.IsSet() {
		return true
	}

	return false
}

// SetTimeToPay gets a reference to the given NullableInt32 and assigns it to the TimeToPay field.
func (o *ModelInvoiceUpdate) SetTimeToPay(v int32) {
	o.TimeToPay.Set(&v)
}
// SetTimeToPayNil sets the value for TimeToPay to be an explicit nil
func (o *ModelInvoiceUpdate) SetTimeToPayNil() {
	o.TimeToPay.Set(nil)
}

// UnsetTimeToPay ensures that no value is present for TimeToPay, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetTimeToPay() {
	o.TimeToPay.Unset()
}

// GetDiscountTime returns the DiscountTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetDiscountTime() int32 {
	if o == nil || IsNil(o.DiscountTime.Get()) {
		var ret int32
		return ret
	}
	return *o.DiscountTime.Get()
}

// GetDiscountTimeOk returns a tuple with the DiscountTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetDiscountTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiscountTime.Get(), o.DiscountTime.IsSet()
}

// HasDiscountTime returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasDiscountTime() bool {
	if o != nil && o.DiscountTime.IsSet() {
		return true
	}

	return false
}

// SetDiscountTime gets a reference to the given NullableInt32 and assigns it to the DiscountTime field.
func (o *ModelInvoiceUpdate) SetDiscountTime(v int32) {
	o.DiscountTime.Set(&v)
}
// SetDiscountTimeNil sets the value for DiscountTime to be an explicit nil
func (o *ModelInvoiceUpdate) SetDiscountTimeNil() {
	o.DiscountTime.Set(nil)
}

// UnsetDiscountTime ensures that no value is present for DiscountTime, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetDiscountTime() {
	o.DiscountTime.Unset()
}

// GetDiscount returns the Discount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetDiscount() int32 {
	if o == nil || IsNil(o.Discount.Get()) {
		var ret int32
		return ret
	}
	return *o.Discount.Get()
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetDiscountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Discount.Get(), o.Discount.IsSet()
}

// HasDiscount returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasDiscount() bool {
	if o != nil && o.Discount.IsSet() {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given NullableInt32 and assigns it to the Discount field.
func (o *ModelInvoiceUpdate) SetDiscount(v int32) {
	o.Discount.Set(&v)
}
// SetDiscountNil sets the value for Discount to be an explicit nil
func (o *ModelInvoiceUpdate) SetDiscountNil() {
	o.Discount.Set(nil)
}

// UnsetDiscount ensures that no value is present for Discount, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetDiscount() {
	o.Discount.Unset()
}

// GetAddressCountry returns the AddressCountry field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetAddressCountry() ModelCreditNoteAddressCountry {
	if o == nil || IsNil(o.AddressCountry.Get()) {
		var ret ModelCreditNoteAddressCountry
		return ret
	}
	return *o.AddressCountry.Get()
}

// GetAddressCountryOk returns a tuple with the AddressCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetAddressCountryOk() (*ModelCreditNoteAddressCountry, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddressCountry.Get(), o.AddressCountry.IsSet()
}

// HasAddressCountry returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasAddressCountry() bool {
	if o != nil && o.AddressCountry.IsSet() {
		return true
	}

	return false
}

// SetAddressCountry gets a reference to the given NullableModelCreditNoteAddressCountry and assigns it to the AddressCountry field.
func (o *ModelInvoiceUpdate) SetAddressCountry(v ModelCreditNoteAddressCountry) {
	o.AddressCountry.Set(&v)
}
// SetAddressCountryNil sets the value for AddressCountry to be an explicit nil
func (o *ModelInvoiceUpdate) SetAddressCountryNil() {
	o.AddressCountry.Set(nil)
}

// UnsetAddressCountry ensures that no value is present for AddressCountry, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetAddressCountry() {
	o.AddressCountry.Unset()
}

// GetPayDate returns the PayDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetPayDate() time.Time {
	if o == nil || IsNil(o.PayDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.PayDate.Get()
}

// GetPayDateOk returns a tuple with the PayDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetPayDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.PayDate.Get(), o.PayDate.IsSet()
}

// HasPayDate returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasPayDate() bool {
	if o != nil && o.PayDate.IsSet() {
		return true
	}

	return false
}

// SetPayDate gets a reference to the given NullableTime and assigns it to the PayDate field.
func (o *ModelInvoiceUpdate) SetPayDate(v time.Time) {
	o.PayDate.Set(&v)
}
// SetPayDateNil sets the value for PayDate to be an explicit nil
func (o *ModelInvoiceUpdate) SetPayDateNil() {
	o.PayDate.Set(nil)
}

// UnsetPayDate ensures that no value is present for PayDate, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetPayDate() {
	o.PayDate.Unset()
}

// GetCreateUser returns the CreateUser field value if set, zero value otherwise.
func (o *ModelInvoiceUpdate) GetCreateUser() ModelCreditNoteCreateUser {
	if o == nil || IsNil(o.CreateUser) {
		var ret ModelCreditNoteCreateUser
		return ret
	}
	return *o.CreateUser
}

// GetCreateUserOk returns a tuple with the CreateUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceUpdate) GetCreateUserOk() (*ModelCreditNoteCreateUser, bool) {
	if o == nil || IsNil(o.CreateUser) {
		return nil, false
	}
	return o.CreateUser, true
}

// HasCreateUser returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasCreateUser() bool {
	if o != nil && !IsNil(o.CreateUser) {
		return true
	}

	return false
}

// SetCreateUser gets a reference to the given ModelCreditNoteCreateUser and assigns it to the CreateUser field.
func (o *ModelInvoiceUpdate) SetCreateUser(v ModelCreditNoteCreateUser) {
	o.CreateUser = &v
}

// GetDeliveryDate returns the DeliveryDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetDeliveryDate() time.Time {
	if o == nil || IsNil(o.DeliveryDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DeliveryDate.Get()
}

// GetDeliveryDateOk returns a tuple with the DeliveryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetDeliveryDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeliveryDate.Get(), o.DeliveryDate.IsSet()
}

// HasDeliveryDate returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasDeliveryDate() bool {
	if o != nil && o.DeliveryDate.IsSet() {
		return true
	}

	return false
}

// SetDeliveryDate gets a reference to the given NullableTime and assigns it to the DeliveryDate field.
func (o *ModelInvoiceUpdate) SetDeliveryDate(v time.Time) {
	o.DeliveryDate.Set(&v)
}
// SetDeliveryDateNil sets the value for DeliveryDate to be an explicit nil
func (o *ModelInvoiceUpdate) SetDeliveryDateNil() {
	o.DeliveryDate.Set(nil)
}

// UnsetDeliveryDate ensures that no value is present for DeliveryDate, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetDeliveryDate() {
	o.DeliveryDate.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *ModelInvoiceUpdate) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *ModelInvoiceUpdate) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetStatus() {
	o.Status.Unset()
}

// GetSmallSettlement returns the SmallSettlement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetSmallSettlement() bool {
	if o == nil || IsNil(o.SmallSettlement.Get()) {
		var ret bool
		return ret
	}
	return *o.SmallSettlement.Get()
}

// GetSmallSettlementOk returns a tuple with the SmallSettlement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetSmallSettlementOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SmallSettlement.Get(), o.SmallSettlement.IsSet()
}

// HasSmallSettlement returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasSmallSettlement() bool {
	if o != nil && o.SmallSettlement.IsSet() {
		return true
	}

	return false
}

// SetSmallSettlement gets a reference to the given NullableBool and assigns it to the SmallSettlement field.
func (o *ModelInvoiceUpdate) SetSmallSettlement(v bool) {
	o.SmallSettlement.Set(&v)
}
// SetSmallSettlementNil sets the value for SmallSettlement to be an explicit nil
func (o *ModelInvoiceUpdate) SetSmallSettlementNil() {
	o.SmallSettlement.Set(nil)
}

// UnsetSmallSettlement ensures that no value is present for SmallSettlement, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetSmallSettlement() {
	o.SmallSettlement.Unset()
}

// GetContactPerson returns the ContactPerson field value if set, zero value otherwise.
func (o *ModelInvoiceUpdate) GetContactPerson() ModelInvoiceUpdateContactPerson {
	if o == nil || IsNil(o.ContactPerson) {
		var ret ModelInvoiceUpdateContactPerson
		return ret
	}
	return *o.ContactPerson
}

// GetContactPersonOk returns a tuple with the ContactPerson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceUpdate) GetContactPersonOk() (*ModelInvoiceUpdateContactPerson, bool) {
	if o == nil || IsNil(o.ContactPerson) {
		return nil, false
	}
	return o.ContactPerson, true
}

// HasContactPerson returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasContactPerson() bool {
	if o != nil && !IsNil(o.ContactPerson) {
		return true
	}

	return false
}

// SetContactPerson gets a reference to the given ModelInvoiceUpdateContactPerson and assigns it to the ContactPerson field.
func (o *ModelInvoiceUpdate) SetContactPerson(v ModelInvoiceUpdateContactPerson) {
	o.ContactPerson = &v
}

// GetTaxRate returns the TaxRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetTaxRate() float32 {
	if o == nil || IsNil(o.TaxRate.Get()) {
		var ret float32
		return ret
	}
	return *o.TaxRate.Get()
}

// GetTaxRateOk returns a tuple with the TaxRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetTaxRateOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxRate.Get(), o.TaxRate.IsSet()
}

// HasTaxRate returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasTaxRate() bool {
	if o != nil && o.TaxRate.IsSet() {
		return true
	}

	return false
}

// SetTaxRate gets a reference to the given NullableFloat32 and assigns it to the TaxRate field.
func (o *ModelInvoiceUpdate) SetTaxRate(v float32) {
	o.TaxRate.Set(&v)
}
// SetTaxRateNil sets the value for TaxRate to be an explicit nil
func (o *ModelInvoiceUpdate) SetTaxRateNil() {
	o.TaxRate.Set(nil)
}

// UnsetTaxRate ensures that no value is present for TaxRate, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetTaxRate() {
	o.TaxRate.Unset()
}

// GetTaxText returns the TaxText field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetTaxText() string {
	if o == nil || IsNil(o.TaxText.Get()) {
		var ret string
		return ret
	}
	return *o.TaxText.Get()
}

// GetTaxTextOk returns a tuple with the TaxText field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetTaxTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxText.Get(), o.TaxText.IsSet()
}

// HasTaxText returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasTaxText() bool {
	if o != nil && o.TaxText.IsSet() {
		return true
	}

	return false
}

// SetTaxText gets a reference to the given NullableString and assigns it to the TaxText field.
func (o *ModelInvoiceUpdate) SetTaxText(v string) {
	o.TaxText.Set(&v)
}
// SetTaxTextNil sets the value for TaxText to be an explicit nil
func (o *ModelInvoiceUpdate) SetTaxTextNil() {
	o.TaxText.Set(nil)
}

// UnsetTaxText ensures that no value is present for TaxText, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetTaxText() {
	o.TaxText.Unset()
}

// GetDunningLevel returns the DunningLevel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetDunningLevel() int32 {
	if o == nil || IsNil(o.DunningLevel.Get()) {
		var ret int32
		return ret
	}
	return *o.DunningLevel.Get()
}

// GetDunningLevelOk returns a tuple with the DunningLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetDunningLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DunningLevel.Get(), o.DunningLevel.IsSet()
}

// HasDunningLevel returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasDunningLevel() bool {
	if o != nil && o.DunningLevel.IsSet() {
		return true
	}

	return false
}

// SetDunningLevel gets a reference to the given NullableInt32 and assigns it to the DunningLevel field.
func (o *ModelInvoiceUpdate) SetDunningLevel(v int32) {
	o.DunningLevel.Set(&v)
}
// SetDunningLevelNil sets the value for DunningLevel to be an explicit nil
func (o *ModelInvoiceUpdate) SetDunningLevelNil() {
	o.DunningLevel.Set(nil)
}

// UnsetDunningLevel ensures that no value is present for DunningLevel, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetDunningLevel() {
	o.DunningLevel.Unset()
}

// GetTaxType returns the TaxType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetTaxType() string {
	if o == nil || IsNil(o.TaxType.Get()) {
		var ret string
		return ret
	}
	return *o.TaxType.Get()
}

// GetTaxTypeOk returns a tuple with the TaxType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetTaxTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxType.Get(), o.TaxType.IsSet()
}

// HasTaxType returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasTaxType() bool {
	if o != nil && o.TaxType.IsSet() {
		return true
	}

	return false
}

// SetTaxType gets a reference to the given NullableString and assigns it to the TaxType field.
func (o *ModelInvoiceUpdate) SetTaxType(v string) {
	o.TaxType.Set(&v)
}
// SetTaxTypeNil sets the value for TaxType to be an explicit nil
func (o *ModelInvoiceUpdate) SetTaxTypeNil() {
	o.TaxType.Set(nil)
}

// UnsetTaxType ensures that no value is present for TaxType, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetTaxType() {
	o.TaxType.Unset()
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *ModelInvoiceUpdate) GetPaymentMethod() ModelInvoiceUpdatePaymentMethod {
	if o == nil || IsNil(o.PaymentMethod) {
		var ret ModelInvoiceUpdatePaymentMethod
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceUpdate) GetPaymentMethodOk() (*ModelInvoiceUpdatePaymentMethod, bool) {
	if o == nil || IsNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasPaymentMethod() bool {
	if o != nil && !IsNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given ModelInvoiceUpdatePaymentMethod and assigns it to the PaymentMethod field.
func (o *ModelInvoiceUpdate) SetPaymentMethod(v ModelInvoiceUpdatePaymentMethod) {
	o.PaymentMethod = &v
}

// GetCostCentre returns the CostCentre field value if set, zero value otherwise.
func (o *ModelInvoiceUpdate) GetCostCentre() ModelInvoiceUpdateCostCentre {
	if o == nil || IsNil(o.CostCentre) {
		var ret ModelInvoiceUpdateCostCentre
		return ret
	}
	return *o.CostCentre
}

// GetCostCentreOk returns a tuple with the CostCentre field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelInvoiceUpdate) GetCostCentreOk() (*ModelInvoiceUpdateCostCentre, bool) {
	if o == nil || IsNil(o.CostCentre) {
		return nil, false
	}
	return o.CostCentre, true
}

// HasCostCentre returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasCostCentre() bool {
	if o != nil && !IsNil(o.CostCentre) {
		return true
	}

	return false
}

// SetCostCentre gets a reference to the given ModelInvoiceUpdateCostCentre and assigns it to the CostCentre field.
func (o *ModelInvoiceUpdate) SetCostCentre(v ModelInvoiceUpdateCostCentre) {
	o.CostCentre = &v
}

// GetSendDate returns the SendDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetSendDate() time.Time {
	if o == nil || IsNil(o.SendDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.SendDate.Get()
}

// GetSendDateOk returns a tuple with the SendDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetSendDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.SendDate.Get(), o.SendDate.IsSet()
}

// HasSendDate returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasSendDate() bool {
	if o != nil && o.SendDate.IsSet() {
		return true
	}

	return false
}

// SetSendDate gets a reference to the given NullableTime and assigns it to the SendDate field.
func (o *ModelInvoiceUpdate) SetSendDate(v time.Time) {
	o.SendDate.Set(&v)
}
// SetSendDateNil sets the value for SendDate to be an explicit nil
func (o *ModelInvoiceUpdate) SetSendDateNil() {
	o.SendDate.Set(nil)
}

// UnsetSendDate ensures that no value is present for SendDate, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetSendDate() {
	o.SendDate.Unset()
}

// GetOrigin returns the Origin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetOrigin() ModelInvoiceUpdateOrigin {
	if o == nil || IsNil(o.Origin.Get()) {
		var ret ModelInvoiceUpdateOrigin
		return ret
	}
	return *o.Origin.Get()
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetOriginOk() (*ModelInvoiceUpdateOrigin, bool) {
	if o == nil {
		return nil, false
	}
	return o.Origin.Get(), o.Origin.IsSet()
}

// HasOrigin returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasOrigin() bool {
	if o != nil && o.Origin.IsSet() {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given NullableModelInvoiceUpdateOrigin and assigns it to the Origin field.
func (o *ModelInvoiceUpdate) SetOrigin(v ModelInvoiceUpdateOrigin) {
	o.Origin.Set(&v)
}
// SetOriginNil sets the value for Origin to be an explicit nil
func (o *ModelInvoiceUpdate) SetOriginNil() {
	o.Origin.Set(nil)
}

// UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetOrigin() {
	o.Origin.Unset()
}

// GetInvoiceType returns the InvoiceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetInvoiceType() string {
	if o == nil || IsNil(o.InvoiceType.Get()) {
		var ret string
		return ret
	}
	return *o.InvoiceType.Get()
}

// GetInvoiceTypeOk returns a tuple with the InvoiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetInvoiceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceType.Get(), o.InvoiceType.IsSet()
}

// HasInvoiceType returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasInvoiceType() bool {
	if o != nil && o.InvoiceType.IsSet() {
		return true
	}

	return false
}

// SetInvoiceType gets a reference to the given NullableString and assigns it to the InvoiceType field.
func (o *ModelInvoiceUpdate) SetInvoiceType(v string) {
	o.InvoiceType.Set(&v)
}
// SetInvoiceTypeNil sets the value for InvoiceType to be an explicit nil
func (o *ModelInvoiceUpdate) SetInvoiceTypeNil() {
	o.InvoiceType.Set(nil)
}

// UnsetInvoiceType ensures that no value is present for InvoiceType, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetInvoiceType() {
	o.InvoiceType.Unset()
}

// GetAccountIntervall returns the AccountIntervall field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetAccountIntervall() string {
	if o == nil || IsNil(o.AccountIntervall.Get()) {
		var ret string
		return ret
	}
	return *o.AccountIntervall.Get()
}

// GetAccountIntervallOk returns a tuple with the AccountIntervall field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetAccountIntervallOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountIntervall.Get(), o.AccountIntervall.IsSet()
}

// HasAccountIntervall returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasAccountIntervall() bool {
	if o != nil && o.AccountIntervall.IsSet() {
		return true
	}

	return false
}

// SetAccountIntervall gets a reference to the given NullableString and assigns it to the AccountIntervall field.
func (o *ModelInvoiceUpdate) SetAccountIntervall(v string) {
	o.AccountIntervall.Set(&v)
}
// SetAccountIntervallNil sets the value for AccountIntervall to be an explicit nil
func (o *ModelInvoiceUpdate) SetAccountIntervallNil() {
	o.AccountIntervall.Set(nil)
}

// UnsetAccountIntervall ensures that no value is present for AccountIntervall, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetAccountIntervall() {
	o.AccountIntervall.Unset()
}

// GetAccountNextInvoice returns the AccountNextInvoice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetAccountNextInvoice() int32 {
	if o == nil || IsNil(o.AccountNextInvoice.Get()) {
		var ret int32
		return ret
	}
	return *o.AccountNextInvoice.Get()
}

// GetAccountNextInvoiceOk returns a tuple with the AccountNextInvoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetAccountNextInvoiceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountNextInvoice.Get(), o.AccountNextInvoice.IsSet()
}

// HasAccountNextInvoice returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasAccountNextInvoice() bool {
	if o != nil && o.AccountNextInvoice.IsSet() {
		return true
	}

	return false
}

// SetAccountNextInvoice gets a reference to the given NullableInt32 and assigns it to the AccountNextInvoice field.
func (o *ModelInvoiceUpdate) SetAccountNextInvoice(v int32) {
	o.AccountNextInvoice.Set(&v)
}
// SetAccountNextInvoiceNil sets the value for AccountNextInvoice to be an explicit nil
func (o *ModelInvoiceUpdate) SetAccountNextInvoiceNil() {
	o.AccountNextInvoice.Set(nil)
}

// UnsetAccountNextInvoice ensures that no value is present for AccountNextInvoice, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetAccountNextInvoice() {
	o.AccountNextInvoice.Unset()
}

// GetReminderTotal returns the ReminderTotal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetReminderTotal() float32 {
	if o == nil || IsNil(o.ReminderTotal.Get()) {
		var ret float32
		return ret
	}
	return *o.ReminderTotal.Get()
}

// GetReminderTotalOk returns a tuple with the ReminderTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetReminderTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReminderTotal.Get(), o.ReminderTotal.IsSet()
}

// HasReminderTotal returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasReminderTotal() bool {
	if o != nil && o.ReminderTotal.IsSet() {
		return true
	}

	return false
}

// SetReminderTotal gets a reference to the given NullableFloat32 and assigns it to the ReminderTotal field.
func (o *ModelInvoiceUpdate) SetReminderTotal(v float32) {
	o.ReminderTotal.Set(&v)
}
// SetReminderTotalNil sets the value for ReminderTotal to be an explicit nil
func (o *ModelInvoiceUpdate) SetReminderTotalNil() {
	o.ReminderTotal.Set(nil)
}

// UnsetReminderTotal ensures that no value is present for ReminderTotal, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetReminderTotal() {
	o.ReminderTotal.Unset()
}

// GetReminderDebit returns the ReminderDebit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetReminderDebit() float32 {
	if o == nil || IsNil(o.ReminderDebit.Get()) {
		var ret float32
		return ret
	}
	return *o.ReminderDebit.Get()
}

// GetReminderDebitOk returns a tuple with the ReminderDebit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetReminderDebitOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReminderDebit.Get(), o.ReminderDebit.IsSet()
}

// HasReminderDebit returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasReminderDebit() bool {
	if o != nil && o.ReminderDebit.IsSet() {
		return true
	}

	return false
}

// SetReminderDebit gets a reference to the given NullableFloat32 and assigns it to the ReminderDebit field.
func (o *ModelInvoiceUpdate) SetReminderDebit(v float32) {
	o.ReminderDebit.Set(&v)
}
// SetReminderDebitNil sets the value for ReminderDebit to be an explicit nil
func (o *ModelInvoiceUpdate) SetReminderDebitNil() {
	o.ReminderDebit.Set(nil)
}

// UnsetReminderDebit ensures that no value is present for ReminderDebit, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetReminderDebit() {
	o.ReminderDebit.Unset()
}

// GetReminderDeadline returns the ReminderDeadline field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetReminderDeadline() int32 {
	if o == nil || IsNil(o.ReminderDeadline.Get()) {
		var ret int32
		return ret
	}
	return *o.ReminderDeadline.Get()
}

// GetReminderDeadlineOk returns a tuple with the ReminderDeadline field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetReminderDeadlineOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReminderDeadline.Get(), o.ReminderDeadline.IsSet()
}

// HasReminderDeadline returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasReminderDeadline() bool {
	if o != nil && o.ReminderDeadline.IsSet() {
		return true
	}

	return false
}

// SetReminderDeadline gets a reference to the given NullableInt32 and assigns it to the ReminderDeadline field.
func (o *ModelInvoiceUpdate) SetReminderDeadline(v int32) {
	o.ReminderDeadline.Set(&v)
}
// SetReminderDeadlineNil sets the value for ReminderDeadline to be an explicit nil
func (o *ModelInvoiceUpdate) SetReminderDeadlineNil() {
	o.ReminderDeadline.Set(nil)
}

// UnsetReminderDeadline ensures that no value is present for ReminderDeadline, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetReminderDeadline() {
	o.ReminderDeadline.Unset()
}

// GetReminderCharge returns the ReminderCharge field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetReminderCharge() float32 {
	if o == nil || IsNil(o.ReminderCharge.Get()) {
		var ret float32
		return ret
	}
	return *o.ReminderCharge.Get()
}

// GetReminderChargeOk returns a tuple with the ReminderCharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetReminderChargeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReminderCharge.Get(), o.ReminderCharge.IsSet()
}

// HasReminderCharge returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasReminderCharge() bool {
	if o != nil && o.ReminderCharge.IsSet() {
		return true
	}

	return false
}

// SetReminderCharge gets a reference to the given NullableFloat32 and assigns it to the ReminderCharge field.
func (o *ModelInvoiceUpdate) SetReminderCharge(v float32) {
	o.ReminderCharge.Set(&v)
}
// SetReminderChargeNil sets the value for ReminderCharge to be an explicit nil
func (o *ModelInvoiceUpdate) SetReminderChargeNil() {
	o.ReminderCharge.Set(nil)
}

// UnsetReminderCharge ensures that no value is present for ReminderCharge, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetReminderCharge() {
	o.ReminderCharge.Unset()
}

// GetTaxSet returns the TaxSet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetTaxSet() ModelInvoiceUpdateTaxSet {
	if o == nil || IsNil(o.TaxSet.Get()) {
		var ret ModelInvoiceUpdateTaxSet
		return ret
	}
	return *o.TaxSet.Get()
}

// GetTaxSetOk returns a tuple with the TaxSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetTaxSetOk() (*ModelInvoiceUpdateTaxSet, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxSet.Get(), o.TaxSet.IsSet()
}

// HasTaxSet returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasTaxSet() bool {
	if o != nil && o.TaxSet.IsSet() {
		return true
	}

	return false
}

// SetTaxSet gets a reference to the given NullableModelInvoiceUpdateTaxSet and assigns it to the TaxSet field.
func (o *ModelInvoiceUpdate) SetTaxSet(v ModelInvoiceUpdateTaxSet) {
	o.TaxSet.Set(&v)
}
// SetTaxSetNil sets the value for TaxSet to be an explicit nil
func (o *ModelInvoiceUpdate) SetTaxSetNil() {
	o.TaxSet.Set(nil)
}

// UnsetTaxSet ensures that no value is present for TaxSet, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetTaxSet() {
	o.TaxSet.Unset()
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetAddress() string {
	if o == nil || IsNil(o.Address.Get()) {
		var ret string
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableString and assigns it to the Address field.
func (o *ModelInvoiceUpdate) SetAddress(v string) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *ModelInvoiceUpdate) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetAddress() {
	o.Address.Unset()
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetCurrency() string {
	if o == nil || IsNil(o.Currency.Get()) {
		var ret string
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasCurrency() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableString and assigns it to the Currency field.
func (o *ModelInvoiceUpdate) SetCurrency(v string) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *ModelInvoiceUpdate) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetCurrency() {
	o.Currency.Unset()
}

// GetSumNet returns the SumNet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetSumNet() float32 {
	if o == nil || IsNil(o.SumNet.Get()) {
		var ret float32
		return ret
	}
	return *o.SumNet.Get()
}

// GetSumNetOk returns a tuple with the SumNet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetSumNetOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SumNet.Get(), o.SumNet.IsSet()
}

// HasSumNet returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasSumNet() bool {
	if o != nil && o.SumNet.IsSet() {
		return true
	}

	return false
}

// SetSumNet gets a reference to the given NullableFloat32 and assigns it to the SumNet field.
func (o *ModelInvoiceUpdate) SetSumNet(v float32) {
	o.SumNet.Set(&v)
}
// SetSumNetNil sets the value for SumNet to be an explicit nil
func (o *ModelInvoiceUpdate) SetSumNetNil() {
	o.SumNet.Set(nil)
}

// UnsetSumNet ensures that no value is present for SumNet, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetSumNet() {
	o.SumNet.Unset()
}

// GetSumTax returns the SumTax field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetSumTax() float32 {
	if o == nil || IsNil(o.SumTax.Get()) {
		var ret float32
		return ret
	}
	return *o.SumTax.Get()
}

// GetSumTaxOk returns a tuple with the SumTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetSumTaxOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SumTax.Get(), o.SumTax.IsSet()
}

// HasSumTax returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasSumTax() bool {
	if o != nil && o.SumTax.IsSet() {
		return true
	}

	return false
}

// SetSumTax gets a reference to the given NullableFloat32 and assigns it to the SumTax field.
func (o *ModelInvoiceUpdate) SetSumTax(v float32) {
	o.SumTax.Set(&v)
}
// SetSumTaxNil sets the value for SumTax to be an explicit nil
func (o *ModelInvoiceUpdate) SetSumTaxNil() {
	o.SumTax.Set(nil)
}

// UnsetSumTax ensures that no value is present for SumTax, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetSumTax() {
	o.SumTax.Unset()
}

// GetSumGross returns the SumGross field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetSumGross() float32 {
	if o == nil || IsNil(o.SumGross.Get()) {
		var ret float32
		return ret
	}
	return *o.SumGross.Get()
}

// GetSumGrossOk returns a tuple with the SumGross field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetSumGrossOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SumGross.Get(), o.SumGross.IsSet()
}

// HasSumGross returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasSumGross() bool {
	if o != nil && o.SumGross.IsSet() {
		return true
	}

	return false
}

// SetSumGross gets a reference to the given NullableFloat32 and assigns it to the SumGross field.
func (o *ModelInvoiceUpdate) SetSumGross(v float32) {
	o.SumGross.Set(&v)
}
// SetSumGrossNil sets the value for SumGross to be an explicit nil
func (o *ModelInvoiceUpdate) SetSumGrossNil() {
	o.SumGross.Set(nil)
}

// UnsetSumGross ensures that no value is present for SumGross, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetSumGross() {
	o.SumGross.Unset()
}

// GetSumDiscounts returns the SumDiscounts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetSumDiscounts() float32 {
	if o == nil || IsNil(o.SumDiscounts.Get()) {
		var ret float32
		return ret
	}
	return *o.SumDiscounts.Get()
}

// GetSumDiscountsOk returns a tuple with the SumDiscounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetSumDiscountsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SumDiscounts.Get(), o.SumDiscounts.IsSet()
}

// HasSumDiscounts returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasSumDiscounts() bool {
	if o != nil && o.SumDiscounts.IsSet() {
		return true
	}

	return false
}

// SetSumDiscounts gets a reference to the given NullableFloat32 and assigns it to the SumDiscounts field.
func (o *ModelInvoiceUpdate) SetSumDiscounts(v float32) {
	o.SumDiscounts.Set(&v)
}
// SetSumDiscountsNil sets the value for SumDiscounts to be an explicit nil
func (o *ModelInvoiceUpdate) SetSumDiscountsNil() {
	o.SumDiscounts.Set(nil)
}

// UnsetSumDiscounts ensures that no value is present for SumDiscounts, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetSumDiscounts() {
	o.SumDiscounts.Unset()
}

// GetSumNetForeignCurrency returns the SumNetForeignCurrency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetSumNetForeignCurrency() float32 {
	if o == nil || IsNil(o.SumNetForeignCurrency.Get()) {
		var ret float32
		return ret
	}
	return *o.SumNetForeignCurrency.Get()
}

// GetSumNetForeignCurrencyOk returns a tuple with the SumNetForeignCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetSumNetForeignCurrencyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SumNetForeignCurrency.Get(), o.SumNetForeignCurrency.IsSet()
}

// HasSumNetForeignCurrency returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasSumNetForeignCurrency() bool {
	if o != nil && o.SumNetForeignCurrency.IsSet() {
		return true
	}

	return false
}

// SetSumNetForeignCurrency gets a reference to the given NullableFloat32 and assigns it to the SumNetForeignCurrency field.
func (o *ModelInvoiceUpdate) SetSumNetForeignCurrency(v float32) {
	o.SumNetForeignCurrency.Set(&v)
}
// SetSumNetForeignCurrencyNil sets the value for SumNetForeignCurrency to be an explicit nil
func (o *ModelInvoiceUpdate) SetSumNetForeignCurrencyNil() {
	o.SumNetForeignCurrency.Set(nil)
}

// UnsetSumNetForeignCurrency ensures that no value is present for SumNetForeignCurrency, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetSumNetForeignCurrency() {
	o.SumNetForeignCurrency.Unset()
}

// GetSumTaxForeignCurrency returns the SumTaxForeignCurrency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetSumTaxForeignCurrency() float32 {
	if o == nil || IsNil(o.SumTaxForeignCurrency.Get()) {
		var ret float32
		return ret
	}
	return *o.SumTaxForeignCurrency.Get()
}

// GetSumTaxForeignCurrencyOk returns a tuple with the SumTaxForeignCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetSumTaxForeignCurrencyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SumTaxForeignCurrency.Get(), o.SumTaxForeignCurrency.IsSet()
}

// HasSumTaxForeignCurrency returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasSumTaxForeignCurrency() bool {
	if o != nil && o.SumTaxForeignCurrency.IsSet() {
		return true
	}

	return false
}

// SetSumTaxForeignCurrency gets a reference to the given NullableFloat32 and assigns it to the SumTaxForeignCurrency field.
func (o *ModelInvoiceUpdate) SetSumTaxForeignCurrency(v float32) {
	o.SumTaxForeignCurrency.Set(&v)
}
// SetSumTaxForeignCurrencyNil sets the value for SumTaxForeignCurrency to be an explicit nil
func (o *ModelInvoiceUpdate) SetSumTaxForeignCurrencyNil() {
	o.SumTaxForeignCurrency.Set(nil)
}

// UnsetSumTaxForeignCurrency ensures that no value is present for SumTaxForeignCurrency, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetSumTaxForeignCurrency() {
	o.SumTaxForeignCurrency.Unset()
}

// GetSumGrossForeignCurrency returns the SumGrossForeignCurrency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetSumGrossForeignCurrency() float32 {
	if o == nil || IsNil(o.SumGrossForeignCurrency.Get()) {
		var ret float32
		return ret
	}
	return *o.SumGrossForeignCurrency.Get()
}

// GetSumGrossForeignCurrencyOk returns a tuple with the SumGrossForeignCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetSumGrossForeignCurrencyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SumGrossForeignCurrency.Get(), o.SumGrossForeignCurrency.IsSet()
}

// HasSumGrossForeignCurrency returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasSumGrossForeignCurrency() bool {
	if o != nil && o.SumGrossForeignCurrency.IsSet() {
		return true
	}

	return false
}

// SetSumGrossForeignCurrency gets a reference to the given NullableFloat32 and assigns it to the SumGrossForeignCurrency field.
func (o *ModelInvoiceUpdate) SetSumGrossForeignCurrency(v float32) {
	o.SumGrossForeignCurrency.Set(&v)
}
// SetSumGrossForeignCurrencyNil sets the value for SumGrossForeignCurrency to be an explicit nil
func (o *ModelInvoiceUpdate) SetSumGrossForeignCurrencyNil() {
	o.SumGrossForeignCurrency.Set(nil)
}

// UnsetSumGrossForeignCurrency ensures that no value is present for SumGrossForeignCurrency, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetSumGrossForeignCurrency() {
	o.SumGrossForeignCurrency.Unset()
}

// GetSumDiscountsForeignCurrency returns the SumDiscountsForeignCurrency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetSumDiscountsForeignCurrency() float32 {
	if o == nil || IsNil(o.SumDiscountsForeignCurrency.Get()) {
		var ret float32
		return ret
	}
	return *o.SumDiscountsForeignCurrency.Get()
}

// GetSumDiscountsForeignCurrencyOk returns a tuple with the SumDiscountsForeignCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetSumDiscountsForeignCurrencyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SumDiscountsForeignCurrency.Get(), o.SumDiscountsForeignCurrency.IsSet()
}

// HasSumDiscountsForeignCurrency returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasSumDiscountsForeignCurrency() bool {
	if o != nil && o.SumDiscountsForeignCurrency.IsSet() {
		return true
	}

	return false
}

// SetSumDiscountsForeignCurrency gets a reference to the given NullableFloat32 and assigns it to the SumDiscountsForeignCurrency field.
func (o *ModelInvoiceUpdate) SetSumDiscountsForeignCurrency(v float32) {
	o.SumDiscountsForeignCurrency.Set(&v)
}
// SetSumDiscountsForeignCurrencyNil sets the value for SumDiscountsForeignCurrency to be an explicit nil
func (o *ModelInvoiceUpdate) SetSumDiscountsForeignCurrencyNil() {
	o.SumDiscountsForeignCurrency.Set(nil)
}

// UnsetSumDiscountsForeignCurrency ensures that no value is present for SumDiscountsForeignCurrency, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetSumDiscountsForeignCurrency() {
	o.SumDiscountsForeignCurrency.Unset()
}

// GetSumNetAccounting returns the SumNetAccounting field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetSumNetAccounting() float32 {
	if o == nil || IsNil(o.SumNetAccounting.Get()) {
		var ret float32
		return ret
	}
	return *o.SumNetAccounting.Get()
}

// GetSumNetAccountingOk returns a tuple with the SumNetAccounting field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetSumNetAccountingOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SumNetAccounting.Get(), o.SumNetAccounting.IsSet()
}

// HasSumNetAccounting returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasSumNetAccounting() bool {
	if o != nil && o.SumNetAccounting.IsSet() {
		return true
	}

	return false
}

// SetSumNetAccounting gets a reference to the given NullableFloat32 and assigns it to the SumNetAccounting field.
func (o *ModelInvoiceUpdate) SetSumNetAccounting(v float32) {
	o.SumNetAccounting.Set(&v)
}
// SetSumNetAccountingNil sets the value for SumNetAccounting to be an explicit nil
func (o *ModelInvoiceUpdate) SetSumNetAccountingNil() {
	o.SumNetAccounting.Set(nil)
}

// UnsetSumNetAccounting ensures that no value is present for SumNetAccounting, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetSumNetAccounting() {
	o.SumNetAccounting.Unset()
}

// GetSumTaxAccounting returns the SumTaxAccounting field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetSumTaxAccounting() float32 {
	if o == nil || IsNil(o.SumTaxAccounting.Get()) {
		var ret float32
		return ret
	}
	return *o.SumTaxAccounting.Get()
}

// GetSumTaxAccountingOk returns a tuple with the SumTaxAccounting field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetSumTaxAccountingOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SumTaxAccounting.Get(), o.SumTaxAccounting.IsSet()
}

// HasSumTaxAccounting returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasSumTaxAccounting() bool {
	if o != nil && o.SumTaxAccounting.IsSet() {
		return true
	}

	return false
}

// SetSumTaxAccounting gets a reference to the given NullableFloat32 and assigns it to the SumTaxAccounting field.
func (o *ModelInvoiceUpdate) SetSumTaxAccounting(v float32) {
	o.SumTaxAccounting.Set(&v)
}
// SetSumTaxAccountingNil sets the value for SumTaxAccounting to be an explicit nil
func (o *ModelInvoiceUpdate) SetSumTaxAccountingNil() {
	o.SumTaxAccounting.Set(nil)
}

// UnsetSumTaxAccounting ensures that no value is present for SumTaxAccounting, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetSumTaxAccounting() {
	o.SumTaxAccounting.Unset()
}

// GetSumGrossAccounting returns the SumGrossAccounting field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetSumGrossAccounting() float32 {
	if o == nil || IsNil(o.SumGrossAccounting.Get()) {
		var ret float32
		return ret
	}
	return *o.SumGrossAccounting.Get()
}

// GetSumGrossAccountingOk returns a tuple with the SumGrossAccounting field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetSumGrossAccountingOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SumGrossAccounting.Get(), o.SumGrossAccounting.IsSet()
}

// HasSumGrossAccounting returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasSumGrossAccounting() bool {
	if o != nil && o.SumGrossAccounting.IsSet() {
		return true
	}

	return false
}

// SetSumGrossAccounting gets a reference to the given NullableFloat32 and assigns it to the SumGrossAccounting field.
func (o *ModelInvoiceUpdate) SetSumGrossAccounting(v float32) {
	o.SumGrossAccounting.Set(&v)
}
// SetSumGrossAccountingNil sets the value for SumGrossAccounting to be an explicit nil
func (o *ModelInvoiceUpdate) SetSumGrossAccountingNil() {
	o.SumGrossAccounting.Set(nil)
}

// UnsetSumGrossAccounting ensures that no value is present for SumGrossAccounting, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetSumGrossAccounting() {
	o.SumGrossAccounting.Unset()
}

// GetPaidAmount returns the PaidAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetPaidAmount() float32 {
	if o == nil || IsNil(o.PaidAmount.Get()) {
		var ret float32
		return ret
	}
	return *o.PaidAmount.Get()
}

// GetPaidAmountOk returns a tuple with the PaidAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetPaidAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaidAmount.Get(), o.PaidAmount.IsSet()
}

// HasPaidAmount returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasPaidAmount() bool {
	if o != nil && o.PaidAmount.IsSet() {
		return true
	}

	return false
}

// SetPaidAmount gets a reference to the given NullableFloat32 and assigns it to the PaidAmount field.
func (o *ModelInvoiceUpdate) SetPaidAmount(v float32) {
	o.PaidAmount.Set(&v)
}
// SetPaidAmountNil sets the value for PaidAmount to be an explicit nil
func (o *ModelInvoiceUpdate) SetPaidAmountNil() {
	o.PaidAmount.Set(nil)
}

// UnsetPaidAmount ensures that no value is present for PaidAmount, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetPaidAmount() {
	o.PaidAmount.Unset()
}

// GetCustomerInternalNote returns the CustomerInternalNote field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetCustomerInternalNote() string {
	if o == nil || IsNil(o.CustomerInternalNote.Get()) {
		var ret string
		return ret
	}
	return *o.CustomerInternalNote.Get()
}

// GetCustomerInternalNoteOk returns a tuple with the CustomerInternalNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetCustomerInternalNoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerInternalNote.Get(), o.CustomerInternalNote.IsSet()
}

// HasCustomerInternalNote returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasCustomerInternalNote() bool {
	if o != nil && o.CustomerInternalNote.IsSet() {
		return true
	}

	return false
}

// SetCustomerInternalNote gets a reference to the given NullableString and assigns it to the CustomerInternalNote field.
func (o *ModelInvoiceUpdate) SetCustomerInternalNote(v string) {
	o.CustomerInternalNote.Set(&v)
}
// SetCustomerInternalNoteNil sets the value for CustomerInternalNote to be an explicit nil
func (o *ModelInvoiceUpdate) SetCustomerInternalNoteNil() {
	o.CustomerInternalNote.Set(nil)
}

// UnsetCustomerInternalNote ensures that no value is present for CustomerInternalNote, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetCustomerInternalNote() {
	o.CustomerInternalNote.Unset()
}

// GetShowNet returns the ShowNet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetShowNet() bool {
	if o == nil || IsNil(o.ShowNet.Get()) {
		var ret bool
		return ret
	}
	return *o.ShowNet.Get()
}

// GetShowNetOk returns a tuple with the ShowNet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetShowNetOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShowNet.Get(), o.ShowNet.IsSet()
}

// HasShowNet returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasShowNet() bool {
	if o != nil && o.ShowNet.IsSet() {
		return true
	}

	return false
}

// SetShowNet gets a reference to the given NullableBool and assigns it to the ShowNet field.
func (o *ModelInvoiceUpdate) SetShowNet(v bool) {
	o.ShowNet.Set(&v)
}
// SetShowNetNil sets the value for ShowNet to be an explicit nil
func (o *ModelInvoiceUpdate) SetShowNetNil() {
	o.ShowNet.Set(nil)
}

// UnsetShowNet ensures that no value is present for ShowNet, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetShowNet() {
	o.ShowNet.Unset()
}

// GetEnshrined returns the Enshrined field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetEnshrined() time.Time {
	if o == nil || IsNil(o.Enshrined.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Enshrined.Get()
}

// GetEnshrinedOk returns a tuple with the Enshrined field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetEnshrinedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enshrined.Get(), o.Enshrined.IsSet()
}

// HasEnshrined returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasEnshrined() bool {
	if o != nil && o.Enshrined.IsSet() {
		return true
	}

	return false
}

// SetEnshrined gets a reference to the given NullableTime and assigns it to the Enshrined field.
func (o *ModelInvoiceUpdate) SetEnshrined(v time.Time) {
	o.Enshrined.Set(&v)
}
// SetEnshrinedNil sets the value for Enshrined to be an explicit nil
func (o *ModelInvoiceUpdate) SetEnshrinedNil() {
	o.Enshrined.Set(nil)
}

// UnsetEnshrined ensures that no value is present for Enshrined, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetEnshrined() {
	o.Enshrined.Unset()
}

// GetSendType returns the SendType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetSendType() string {
	if o == nil || IsNil(o.SendType.Get()) {
		var ret string
		return ret
	}
	return *o.SendType.Get()
}

// GetSendTypeOk returns a tuple with the SendType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetSendTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SendType.Get(), o.SendType.IsSet()
}

// HasSendType returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasSendType() bool {
	if o != nil && o.SendType.IsSet() {
		return true
	}

	return false
}

// SetSendType gets a reference to the given NullableString and assigns it to the SendType field.
func (o *ModelInvoiceUpdate) SetSendType(v string) {
	o.SendType.Set(&v)
}
// SetSendTypeNil sets the value for SendType to be an explicit nil
func (o *ModelInvoiceUpdate) SetSendTypeNil() {
	o.SendType.Set(nil)
}

// UnsetSendType ensures that no value is present for SendType, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetSendType() {
	o.SendType.Unset()
}

// GetDeliveryDateUntil returns the DeliveryDateUntil field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetDeliveryDateUntil() int32 {
	if o == nil || IsNil(o.DeliveryDateUntil.Get()) {
		var ret int32
		return ret
	}
	return *o.DeliveryDateUntil.Get()
}

// GetDeliveryDateUntilOk returns a tuple with the DeliveryDateUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetDeliveryDateUntilOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeliveryDateUntil.Get(), o.DeliveryDateUntil.IsSet()
}

// HasDeliveryDateUntil returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasDeliveryDateUntil() bool {
	if o != nil && o.DeliveryDateUntil.IsSet() {
		return true
	}

	return false
}

// SetDeliveryDateUntil gets a reference to the given NullableInt32 and assigns it to the DeliveryDateUntil field.
func (o *ModelInvoiceUpdate) SetDeliveryDateUntil(v int32) {
	o.DeliveryDateUntil.Set(&v)
}
// SetDeliveryDateUntilNil sets the value for DeliveryDateUntil to be an explicit nil
func (o *ModelInvoiceUpdate) SetDeliveryDateUntilNil() {
	o.DeliveryDateUntil.Set(nil)
}

// UnsetDeliveryDateUntil ensures that no value is present for DeliveryDateUntil, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetDeliveryDateUntil() {
	o.DeliveryDateUntil.Unset()
}

// GetDatevConnectOnline returns the DatevConnectOnline field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetDatevConnectOnline() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.DatevConnectOnline
}

// GetDatevConnectOnlineOk returns a tuple with the DatevConnectOnline field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetDatevConnectOnlineOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.DatevConnectOnline) {
		return map[string]interface{}{}, false
	}
	return o.DatevConnectOnline, true
}

// HasDatevConnectOnline returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasDatevConnectOnline() bool {
	if o != nil && IsNil(o.DatevConnectOnline) {
		return true
	}

	return false
}

// SetDatevConnectOnline gets a reference to the given map[string]interface{} and assigns it to the DatevConnectOnline field.
func (o *ModelInvoiceUpdate) SetDatevConnectOnline(v map[string]interface{}) {
	o.DatevConnectOnline = v
}

// GetSendPaymentReceivedNotificationDate returns the SendPaymentReceivedNotificationDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInvoiceUpdate) GetSendPaymentReceivedNotificationDate() int32 {
	if o == nil || IsNil(o.SendPaymentReceivedNotificationDate.Get()) {
		var ret int32
		return ret
	}
	return *o.SendPaymentReceivedNotificationDate.Get()
}

// GetSendPaymentReceivedNotificationDateOk returns a tuple with the SendPaymentReceivedNotificationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInvoiceUpdate) GetSendPaymentReceivedNotificationDateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SendPaymentReceivedNotificationDate.Get(), o.SendPaymentReceivedNotificationDate.IsSet()
}

// HasSendPaymentReceivedNotificationDate returns a boolean if a field has been set.
func (o *ModelInvoiceUpdate) HasSendPaymentReceivedNotificationDate() bool {
	if o != nil && o.SendPaymentReceivedNotificationDate.IsSet() {
		return true
	}

	return false
}

// SetSendPaymentReceivedNotificationDate gets a reference to the given NullableInt32 and assigns it to the SendPaymentReceivedNotificationDate field.
func (o *ModelInvoiceUpdate) SetSendPaymentReceivedNotificationDate(v int32) {
	o.SendPaymentReceivedNotificationDate.Set(&v)
}
// SetSendPaymentReceivedNotificationDateNil sets the value for SendPaymentReceivedNotificationDate to be an explicit nil
func (o *ModelInvoiceUpdate) SetSendPaymentReceivedNotificationDateNil() {
	o.SendPaymentReceivedNotificationDate.Set(nil)
}

// UnsetSendPaymentReceivedNotificationDate ensures that no value is present for SendPaymentReceivedNotificationDate, not even an explicit nil
func (o *ModelInvoiceUpdate) UnsetSendPaymentReceivedNotificationDate() {
	o.SendPaymentReceivedNotificationDate.Unset()
}

func (o ModelInvoiceUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelInvoiceUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	// skip: objectName is readOnly
	if o.InvoiceNumber.IsSet() {
		toSerialize["invoiceNumber"] = o.InvoiceNumber.Get()
	}
	if o.Contact.IsSet() {
		toSerialize["contact"] = o.Contact.Get()
	}
	// skip: create is readOnly
	// skip: update is readOnly
	if !IsNil(o.SevClient) {
		toSerialize["sevClient"] = o.SevClient
	}
	if o.InvoiceDate.IsSet() {
		toSerialize["invoiceDate"] = o.InvoiceDate.Get()
	}
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
	if o.DiscountTime.IsSet() {
		toSerialize["discountTime"] = o.DiscountTime.Get()
	}
	if o.Discount.IsSet() {
		toSerialize["discount"] = o.Discount.Get()
	}
	if o.AddressCountry.IsSet() {
		toSerialize["addressCountry"] = o.AddressCountry.Get()
	}
	if o.PayDate.IsSet() {
		toSerialize["payDate"] = o.PayDate.Get()
	}
	if !IsNil(o.CreateUser) {
		toSerialize["createUser"] = o.CreateUser
	}
	if o.DeliveryDate.IsSet() {
		toSerialize["deliveryDate"] = o.DeliveryDate.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.SmallSettlement.IsSet() {
		toSerialize["smallSettlement"] = o.SmallSettlement.Get()
	}
	if !IsNil(o.ContactPerson) {
		toSerialize["contactPerson"] = o.ContactPerson
	}
	if o.TaxRate.IsSet() {
		toSerialize["taxRate"] = o.TaxRate.Get()
	}
	if o.TaxText.IsSet() {
		toSerialize["taxText"] = o.TaxText.Get()
	}
	if o.DunningLevel.IsSet() {
		toSerialize["dunningLevel"] = o.DunningLevel.Get()
	}
	if o.TaxType.IsSet() {
		toSerialize["taxType"] = o.TaxType.Get()
	}
	if !IsNil(o.PaymentMethod) {
		toSerialize["paymentMethod"] = o.PaymentMethod
	}
	if !IsNil(o.CostCentre) {
		toSerialize["costCentre"] = o.CostCentre
	}
	if o.SendDate.IsSet() {
		toSerialize["sendDate"] = o.SendDate.Get()
	}
	if o.Origin.IsSet() {
		toSerialize["origin"] = o.Origin.Get()
	}
	if o.InvoiceType.IsSet() {
		toSerialize["invoiceType"] = o.InvoiceType.Get()
	}
	if o.AccountIntervall.IsSet() {
		toSerialize["accountIntervall"] = o.AccountIntervall.Get()
	}
	if o.AccountNextInvoice.IsSet() {
		toSerialize["accountNextInvoice"] = o.AccountNextInvoice.Get()
	}
	if o.ReminderTotal.IsSet() {
		toSerialize["reminderTotal"] = o.ReminderTotal.Get()
	}
	if o.ReminderDebit.IsSet() {
		toSerialize["reminderDebit"] = o.ReminderDebit.Get()
	}
	if o.ReminderDeadline.IsSet() {
		toSerialize["reminderDeadline"] = o.ReminderDeadline.Get()
	}
	if o.ReminderCharge.IsSet() {
		toSerialize["reminderCharge"] = o.ReminderCharge.Get()
	}
	if o.TaxSet.IsSet() {
		toSerialize["taxSet"] = o.TaxSet.Get()
	}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	if o.SumNet.IsSet() {
		toSerialize["sumNet"] = o.SumNet.Get()
	}
	if o.SumTax.IsSet() {
		toSerialize["sumTax"] = o.SumTax.Get()
	}
	if o.SumGross.IsSet() {
		toSerialize["sumGross"] = o.SumGross.Get()
	}
	if o.SumDiscounts.IsSet() {
		toSerialize["sumDiscounts"] = o.SumDiscounts.Get()
	}
	if o.SumNetForeignCurrency.IsSet() {
		toSerialize["sumNetForeignCurrency"] = o.SumNetForeignCurrency.Get()
	}
	if o.SumTaxForeignCurrency.IsSet() {
		toSerialize["sumTaxForeignCurrency"] = o.SumTaxForeignCurrency.Get()
	}
	if o.SumGrossForeignCurrency.IsSet() {
		toSerialize["sumGrossForeignCurrency"] = o.SumGrossForeignCurrency.Get()
	}
	if o.SumDiscountsForeignCurrency.IsSet() {
		toSerialize["sumDiscountsForeignCurrency"] = o.SumDiscountsForeignCurrency.Get()
	}
	if o.SumNetAccounting.IsSet() {
		toSerialize["sumNetAccounting"] = o.SumNetAccounting.Get()
	}
	if o.SumTaxAccounting.IsSet() {
		toSerialize["sumTaxAccounting"] = o.SumTaxAccounting.Get()
	}
	if o.SumGrossAccounting.IsSet() {
		toSerialize["sumGrossAccounting"] = o.SumGrossAccounting.Get()
	}
	if o.PaidAmount.IsSet() {
		toSerialize["paidAmount"] = o.PaidAmount.Get()
	}
	if o.CustomerInternalNote.IsSet() {
		toSerialize["customerInternalNote"] = o.CustomerInternalNote.Get()
	}
	if o.ShowNet.IsSet() {
		toSerialize["showNet"] = o.ShowNet.Get()
	}
	if o.Enshrined.IsSet() {
		toSerialize["enshrined"] = o.Enshrined.Get()
	}
	if o.SendType.IsSet() {
		toSerialize["sendType"] = o.SendType.Get()
	}
	if o.DeliveryDateUntil.IsSet() {
		toSerialize["deliveryDateUntil"] = o.DeliveryDateUntil.Get()
	}
	if o.DatevConnectOnline != nil {
		toSerialize["datevConnectOnline"] = o.DatevConnectOnline
	}
	if o.SendPaymentReceivedNotificationDate.IsSet() {
		toSerialize["sendPaymentReceivedNotificationDate"] = o.SendPaymentReceivedNotificationDate.Get()
	}
	return toSerialize, nil
}

type NullableModelInvoiceUpdate struct {
	value *ModelInvoiceUpdate
	isSet bool
}

func (v NullableModelInvoiceUpdate) Get() *ModelInvoiceUpdate {
	return v.value
}

func (v *NullableModelInvoiceUpdate) Set(val *ModelInvoiceUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableModelInvoiceUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableModelInvoiceUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelInvoiceUpdate(val *ModelInvoiceUpdate) *NullableModelInvoiceUpdate {
	return &NullableModelInvoiceUpdate{value: val, isSet: true}
}

func (v NullableModelInvoiceUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelInvoiceUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


