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

// checks if the ModelOrder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelOrder{}

// ModelOrder Order model
type ModelOrder struct {
	// The order id
	Id *int32 `json:"id,omitempty"`
	// The order object name
	ObjectName *string `json:"objectName,omitempty"`
	MapAll bool `json:"mapAll"`
	// Date of order creation
	Create *time.Time `json:"create,omitempty"`
	// Date of last order update
	Update *time.Time `json:"update,omitempty"`
	// The order number
	OrderNumber string `json:"orderNumber"`
	Contact ModelOrderContact `json:"contact"`
	// Needs to be provided as timestamp or dd.mm.yyyy
	OrderDate time.Time `json:"orderDate"`
	// Please have a look in       <a href='https://api.sevdesk.de/#section/Types-and-status-of-orders'>status of orders</a>      to see what the different status codes mean
	Status int32 `json:"status"`
	// Normally consist of prefix plus the order number
	Header string `json:"header"`
	// Certain html tags can be used here to format your text
	HeadText NullableString `json:"headText,omitempty"`
	// Certain html tags can be used here to format your text
	FootText NullableString `json:"footText,omitempty"`
	AddressCountry ModelOrderAddressCountry `json:"addressCountry"`
	// Delivery terms of the order
	DeliveryTerms NullableString `json:"deliveryTerms,omitempty"`
	// Payment terms of the order
	PaymentTerms NullableString `json:"paymentTerms,omitempty"`
	// Version of the order.<br>      Can be used if you have multiple drafts for the same order.<br>      Should start with 0
	Version int32 `json:"version"`
	// Defines if the client uses the small settlement scheme.      If yes, the order must not contain any vat
	SmallSettlement *bool `json:"smallSettlement,omitempty"`
	ContactPerson ModelOrderContactPerson `json:"contactPerson"`
	// Is overwritten by order position tax rates
	TaxRate float32 `json:"taxRate"`
	TaxSet NullableModelOrderTaxSet `json:"taxSet,omitempty"`
	// A common tax text would be 'Umsatzsteuer 19%'
	TaxText string `json:"taxText"`
	// Tax type of the order. There are four tax types: 1. default - Umsatzsteuer ausweisen 2. eu - Steuerfreie innergemeinschaftliche Lieferung (Europäische Union) 3. noteu - Steuerschuldnerschaft des Leistungsempfängers (außerhalb EU, z. B. Schweiz) 4. custom - Using custom tax set 5. ss - Not subject to VAT according to §19 1 UStG Tax rates are heavily connected to the tax type used.
	TaxType string `json:"taxType"`
	// Type of the order. For more information on the different types, check      <a href='https://api.sevdesk.de/#section/Types-and-status-of-orders'>this</a>   
	OrderType *string `json:"orderType,omitempty"`
	// The date the order was sent to the customer
	SendDate NullableTime `json:"sendDate,omitempty"`
	// Complete address of the recipient including name, street, city, zip and country.<br>       Line breaks can be used and will be displayed on the invoice pdf.
	Address NullableString `json:"address,omitempty"`
	// Currency used in the order. Needs to be currency code according to ISO-4217
	Currency string `json:"currency"`
	// Internal note of the customer. Contains data entered into field 'Referenz/Bestellnummer'
	CustomerInternalNote NullableString `json:"customerInternalNote,omitempty"`
	// If true, the net amount of each position will be shown on the order. Otherwise gross amount
	ShowNet *bool `json:"showNet,omitempty"`
	// Type which was used to send the order. IMPORTANT: Please refer to the order section of the       *     API-Overview to understand how this attribute can be used before using it!
	SendType NullableString `json:"sendType,omitempty"`
	Origin NullableModelOrderOrigin `json:"origin,omitempty"`
}

// NewModelOrder instantiates a new ModelOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelOrder(mapAll bool, orderNumber string, contact ModelOrderContact, orderDate time.Time, status int32, header string, addressCountry ModelOrderAddressCountry, version int32, contactPerson ModelOrderContactPerson, taxRate float32, taxText string, taxType string, currency string) *ModelOrder {
	this := ModelOrder{}
	this.MapAll = mapAll
	this.OrderNumber = orderNumber
	this.Contact = contact
	this.OrderDate = orderDate
	this.Status = status
	this.Header = header
	this.AddressCountry = addressCountry
	this.Version = version
	this.ContactPerson = contactPerson
	this.TaxRate = taxRate
	this.TaxText = taxText
	this.TaxType = taxType
	this.Currency = currency
	return &this
}

// NewModelOrderWithDefaults instantiates a new ModelOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelOrderWithDefaults() *ModelOrder {
	this := ModelOrder{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelOrder) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelOrder) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelOrder) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ModelOrder) SetId(v int32) {
	o.Id = &v
}

// GetObjectName returns the ObjectName field value if set, zero value otherwise.
func (o *ModelOrder) GetObjectName() string {
	if o == nil || IsNil(o.ObjectName) {
		var ret string
		return ret
	}
	return *o.ObjectName
}

// GetObjectNameOk returns a tuple with the ObjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelOrder) GetObjectNameOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectName) {
		return nil, false
	}
	return o.ObjectName, true
}

// HasObjectName returns a boolean if a field has been set.
func (o *ModelOrder) HasObjectName() bool {
	if o != nil && !IsNil(o.ObjectName) {
		return true
	}

	return false
}

// SetObjectName gets a reference to the given string and assigns it to the ObjectName field.
func (o *ModelOrder) SetObjectName(v string) {
	o.ObjectName = &v
}

// GetMapAll returns the MapAll field value
func (o *ModelOrder) GetMapAll() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MapAll
}

// GetMapAllOk returns a tuple with the MapAll field value
// and a boolean to check if the value has been set.
func (o *ModelOrder) GetMapAllOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MapAll, true
}

// SetMapAll sets field value
func (o *ModelOrder) SetMapAll(v bool) {
	o.MapAll = v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *ModelOrder) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelOrder) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *ModelOrder) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *ModelOrder) SetCreate(v time.Time) {
	o.Create = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *ModelOrder) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelOrder) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *ModelOrder) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *ModelOrder) SetUpdate(v time.Time) {
	o.Update = &v
}

// GetOrderNumber returns the OrderNumber field value
func (o *ModelOrder) GetOrderNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderNumber
}

// GetOrderNumberOk returns a tuple with the OrderNumber field value
// and a boolean to check if the value has been set.
func (o *ModelOrder) GetOrderNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderNumber, true
}

// SetOrderNumber sets field value
func (o *ModelOrder) SetOrderNumber(v string) {
	o.OrderNumber = v
}

// GetContact returns the Contact field value
func (o *ModelOrder) GetContact() ModelOrderContact {
	if o == nil {
		var ret ModelOrderContact
		return ret
	}

	return o.Contact
}

// GetContactOk returns a tuple with the Contact field value
// and a boolean to check if the value has been set.
func (o *ModelOrder) GetContactOk() (*ModelOrderContact, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contact, true
}

// SetContact sets field value
func (o *ModelOrder) SetContact(v ModelOrderContact) {
	o.Contact = v
}

// GetOrderDate returns the OrderDate field value
func (o *ModelOrder) GetOrderDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.OrderDate
}

// GetOrderDateOk returns a tuple with the OrderDate field value
// and a boolean to check if the value has been set.
func (o *ModelOrder) GetOrderDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderDate, true
}

// SetOrderDate sets field value
func (o *ModelOrder) SetOrderDate(v time.Time) {
	o.OrderDate = v
}

// GetStatus returns the Status field value
func (o *ModelOrder) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ModelOrder) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ModelOrder) SetStatus(v int32) {
	o.Status = v
}

// GetHeader returns the Header field value
func (o *ModelOrder) GetHeader() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Header
}

// GetHeaderOk returns a tuple with the Header field value
// and a boolean to check if the value has been set.
func (o *ModelOrder) GetHeaderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Header, true
}

// SetHeader sets field value
func (o *ModelOrder) SetHeader(v string) {
	o.Header = v
}

// GetHeadText returns the HeadText field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelOrder) GetHeadText() string {
	if o == nil || IsNil(o.HeadText.Get()) {
		var ret string
		return ret
	}
	return *o.HeadText.Get()
}

// GetHeadTextOk returns a tuple with the HeadText field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelOrder) GetHeadTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HeadText.Get(), o.HeadText.IsSet()
}

// HasHeadText returns a boolean if a field has been set.
func (o *ModelOrder) HasHeadText() bool {
	if o != nil && o.HeadText.IsSet() {
		return true
	}

	return false
}

// SetHeadText gets a reference to the given NullableString and assigns it to the HeadText field.
func (o *ModelOrder) SetHeadText(v string) {
	o.HeadText.Set(&v)
}
// SetHeadTextNil sets the value for HeadText to be an explicit nil
func (o *ModelOrder) SetHeadTextNil() {
	o.HeadText.Set(nil)
}

// UnsetHeadText ensures that no value is present for HeadText, not even an explicit nil
func (o *ModelOrder) UnsetHeadText() {
	o.HeadText.Unset()
}

// GetFootText returns the FootText field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelOrder) GetFootText() string {
	if o == nil || IsNil(o.FootText.Get()) {
		var ret string
		return ret
	}
	return *o.FootText.Get()
}

// GetFootTextOk returns a tuple with the FootText field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelOrder) GetFootTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FootText.Get(), o.FootText.IsSet()
}

// HasFootText returns a boolean if a field has been set.
func (o *ModelOrder) HasFootText() bool {
	if o != nil && o.FootText.IsSet() {
		return true
	}

	return false
}

// SetFootText gets a reference to the given NullableString and assigns it to the FootText field.
func (o *ModelOrder) SetFootText(v string) {
	o.FootText.Set(&v)
}
// SetFootTextNil sets the value for FootText to be an explicit nil
func (o *ModelOrder) SetFootTextNil() {
	o.FootText.Set(nil)
}

// UnsetFootText ensures that no value is present for FootText, not even an explicit nil
func (o *ModelOrder) UnsetFootText() {
	o.FootText.Unset()
}

// GetAddressCountry returns the AddressCountry field value
func (o *ModelOrder) GetAddressCountry() ModelOrderAddressCountry {
	if o == nil {
		var ret ModelOrderAddressCountry
		return ret
	}

	return o.AddressCountry
}

// GetAddressCountryOk returns a tuple with the AddressCountry field value
// and a boolean to check if the value has been set.
func (o *ModelOrder) GetAddressCountryOk() (*ModelOrderAddressCountry, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressCountry, true
}

// SetAddressCountry sets field value
func (o *ModelOrder) SetAddressCountry(v ModelOrderAddressCountry) {
	o.AddressCountry = v
}

// GetDeliveryTerms returns the DeliveryTerms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelOrder) GetDeliveryTerms() string {
	if o == nil || IsNil(o.DeliveryTerms.Get()) {
		var ret string
		return ret
	}
	return *o.DeliveryTerms.Get()
}

// GetDeliveryTermsOk returns a tuple with the DeliveryTerms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelOrder) GetDeliveryTermsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeliveryTerms.Get(), o.DeliveryTerms.IsSet()
}

// HasDeliveryTerms returns a boolean if a field has been set.
func (o *ModelOrder) HasDeliveryTerms() bool {
	if o != nil && o.DeliveryTerms.IsSet() {
		return true
	}

	return false
}

// SetDeliveryTerms gets a reference to the given NullableString and assigns it to the DeliveryTerms field.
func (o *ModelOrder) SetDeliveryTerms(v string) {
	o.DeliveryTerms.Set(&v)
}
// SetDeliveryTermsNil sets the value for DeliveryTerms to be an explicit nil
func (o *ModelOrder) SetDeliveryTermsNil() {
	o.DeliveryTerms.Set(nil)
}

// UnsetDeliveryTerms ensures that no value is present for DeliveryTerms, not even an explicit nil
func (o *ModelOrder) UnsetDeliveryTerms() {
	o.DeliveryTerms.Unset()
}

// GetPaymentTerms returns the PaymentTerms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelOrder) GetPaymentTerms() string {
	if o == nil || IsNil(o.PaymentTerms.Get()) {
		var ret string
		return ret
	}
	return *o.PaymentTerms.Get()
}

// GetPaymentTermsOk returns a tuple with the PaymentTerms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelOrder) GetPaymentTermsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentTerms.Get(), o.PaymentTerms.IsSet()
}

// HasPaymentTerms returns a boolean if a field has been set.
func (o *ModelOrder) HasPaymentTerms() bool {
	if o != nil && o.PaymentTerms.IsSet() {
		return true
	}

	return false
}

// SetPaymentTerms gets a reference to the given NullableString and assigns it to the PaymentTerms field.
func (o *ModelOrder) SetPaymentTerms(v string) {
	o.PaymentTerms.Set(&v)
}
// SetPaymentTermsNil sets the value for PaymentTerms to be an explicit nil
func (o *ModelOrder) SetPaymentTermsNil() {
	o.PaymentTerms.Set(nil)
}

// UnsetPaymentTerms ensures that no value is present for PaymentTerms, not even an explicit nil
func (o *ModelOrder) UnsetPaymentTerms() {
	o.PaymentTerms.Unset()
}

// GetVersion returns the Version field value
func (o *ModelOrder) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ModelOrder) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ModelOrder) SetVersion(v int32) {
	o.Version = v
}

// GetSmallSettlement returns the SmallSettlement field value if set, zero value otherwise.
func (o *ModelOrder) GetSmallSettlement() bool {
	if o == nil || IsNil(o.SmallSettlement) {
		var ret bool
		return ret
	}
	return *o.SmallSettlement
}

// GetSmallSettlementOk returns a tuple with the SmallSettlement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelOrder) GetSmallSettlementOk() (*bool, bool) {
	if o == nil || IsNil(o.SmallSettlement) {
		return nil, false
	}
	return o.SmallSettlement, true
}

// HasSmallSettlement returns a boolean if a field has been set.
func (o *ModelOrder) HasSmallSettlement() bool {
	if o != nil && !IsNil(o.SmallSettlement) {
		return true
	}

	return false
}

// SetSmallSettlement gets a reference to the given bool and assigns it to the SmallSettlement field.
func (o *ModelOrder) SetSmallSettlement(v bool) {
	o.SmallSettlement = &v
}

// GetContactPerson returns the ContactPerson field value
func (o *ModelOrder) GetContactPerson() ModelOrderContactPerson {
	if o == nil {
		var ret ModelOrderContactPerson
		return ret
	}

	return o.ContactPerson
}

// GetContactPersonOk returns a tuple with the ContactPerson field value
// and a boolean to check if the value has been set.
func (o *ModelOrder) GetContactPersonOk() (*ModelOrderContactPerson, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContactPerson, true
}

// SetContactPerson sets field value
func (o *ModelOrder) SetContactPerson(v ModelOrderContactPerson) {
	o.ContactPerson = v
}

// GetTaxRate returns the TaxRate field value
func (o *ModelOrder) GetTaxRate() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TaxRate
}

// GetTaxRateOk returns a tuple with the TaxRate field value
// and a boolean to check if the value has been set.
func (o *ModelOrder) GetTaxRateOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaxRate, true
}

// SetTaxRate sets field value
func (o *ModelOrder) SetTaxRate(v float32) {
	o.TaxRate = v
}

// GetTaxSet returns the TaxSet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelOrder) GetTaxSet() ModelOrderTaxSet {
	if o == nil || IsNil(o.TaxSet.Get()) {
		var ret ModelOrderTaxSet
		return ret
	}
	return *o.TaxSet.Get()
}

// GetTaxSetOk returns a tuple with the TaxSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelOrder) GetTaxSetOk() (*ModelOrderTaxSet, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxSet.Get(), o.TaxSet.IsSet()
}

// HasTaxSet returns a boolean if a field has been set.
func (o *ModelOrder) HasTaxSet() bool {
	if o != nil && o.TaxSet.IsSet() {
		return true
	}

	return false
}

// SetTaxSet gets a reference to the given NullableModelOrderTaxSet and assigns it to the TaxSet field.
func (o *ModelOrder) SetTaxSet(v ModelOrderTaxSet) {
	o.TaxSet.Set(&v)
}
// SetTaxSetNil sets the value for TaxSet to be an explicit nil
func (o *ModelOrder) SetTaxSetNil() {
	o.TaxSet.Set(nil)
}

// UnsetTaxSet ensures that no value is present for TaxSet, not even an explicit nil
func (o *ModelOrder) UnsetTaxSet() {
	o.TaxSet.Unset()
}

// GetTaxText returns the TaxText field value
func (o *ModelOrder) GetTaxText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TaxText
}

// GetTaxTextOk returns a tuple with the TaxText field value
// and a boolean to check if the value has been set.
func (o *ModelOrder) GetTaxTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaxText, true
}

// SetTaxText sets field value
func (o *ModelOrder) SetTaxText(v string) {
	o.TaxText = v
}

// GetTaxType returns the TaxType field value
func (o *ModelOrder) GetTaxType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TaxType
}

// GetTaxTypeOk returns a tuple with the TaxType field value
// and a boolean to check if the value has been set.
func (o *ModelOrder) GetTaxTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaxType, true
}

// SetTaxType sets field value
func (o *ModelOrder) SetTaxType(v string) {
	o.TaxType = v
}

// GetOrderType returns the OrderType field value if set, zero value otherwise.
func (o *ModelOrder) GetOrderType() string {
	if o == nil || IsNil(o.OrderType) {
		var ret string
		return ret
	}
	return *o.OrderType
}

// GetOrderTypeOk returns a tuple with the OrderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelOrder) GetOrderTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OrderType) {
		return nil, false
	}
	return o.OrderType, true
}

// HasOrderType returns a boolean if a field has been set.
func (o *ModelOrder) HasOrderType() bool {
	if o != nil && !IsNil(o.OrderType) {
		return true
	}

	return false
}

// SetOrderType gets a reference to the given string and assigns it to the OrderType field.
func (o *ModelOrder) SetOrderType(v string) {
	o.OrderType = &v
}

// GetSendDate returns the SendDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelOrder) GetSendDate() time.Time {
	if o == nil || IsNil(o.SendDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.SendDate.Get()
}

// GetSendDateOk returns a tuple with the SendDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelOrder) GetSendDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.SendDate.Get(), o.SendDate.IsSet()
}

// HasSendDate returns a boolean if a field has been set.
func (o *ModelOrder) HasSendDate() bool {
	if o != nil && o.SendDate.IsSet() {
		return true
	}

	return false
}

// SetSendDate gets a reference to the given NullableTime and assigns it to the SendDate field.
func (o *ModelOrder) SetSendDate(v time.Time) {
	o.SendDate.Set(&v)
}
// SetSendDateNil sets the value for SendDate to be an explicit nil
func (o *ModelOrder) SetSendDateNil() {
	o.SendDate.Set(nil)
}

// UnsetSendDate ensures that no value is present for SendDate, not even an explicit nil
func (o *ModelOrder) UnsetSendDate() {
	o.SendDate.Unset()
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelOrder) GetAddress() string {
	if o == nil || IsNil(o.Address.Get()) {
		var ret string
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelOrder) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *ModelOrder) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableString and assigns it to the Address field.
func (o *ModelOrder) SetAddress(v string) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *ModelOrder) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *ModelOrder) UnsetAddress() {
	o.Address.Unset()
}

// GetCurrency returns the Currency field value
func (o *ModelOrder) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *ModelOrder) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *ModelOrder) SetCurrency(v string) {
	o.Currency = v
}

// GetCustomerInternalNote returns the CustomerInternalNote field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelOrder) GetCustomerInternalNote() string {
	if o == nil || IsNil(o.CustomerInternalNote.Get()) {
		var ret string
		return ret
	}
	return *o.CustomerInternalNote.Get()
}

// GetCustomerInternalNoteOk returns a tuple with the CustomerInternalNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelOrder) GetCustomerInternalNoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerInternalNote.Get(), o.CustomerInternalNote.IsSet()
}

// HasCustomerInternalNote returns a boolean if a field has been set.
func (o *ModelOrder) HasCustomerInternalNote() bool {
	if o != nil && o.CustomerInternalNote.IsSet() {
		return true
	}

	return false
}

// SetCustomerInternalNote gets a reference to the given NullableString and assigns it to the CustomerInternalNote field.
func (o *ModelOrder) SetCustomerInternalNote(v string) {
	o.CustomerInternalNote.Set(&v)
}
// SetCustomerInternalNoteNil sets the value for CustomerInternalNote to be an explicit nil
func (o *ModelOrder) SetCustomerInternalNoteNil() {
	o.CustomerInternalNote.Set(nil)
}

// UnsetCustomerInternalNote ensures that no value is present for CustomerInternalNote, not even an explicit nil
func (o *ModelOrder) UnsetCustomerInternalNote() {
	o.CustomerInternalNote.Unset()
}

// GetShowNet returns the ShowNet field value if set, zero value otherwise.
func (o *ModelOrder) GetShowNet() bool {
	if o == nil || IsNil(o.ShowNet) {
		var ret bool
		return ret
	}
	return *o.ShowNet
}

// GetShowNetOk returns a tuple with the ShowNet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelOrder) GetShowNetOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowNet) {
		return nil, false
	}
	return o.ShowNet, true
}

// HasShowNet returns a boolean if a field has been set.
func (o *ModelOrder) HasShowNet() bool {
	if o != nil && !IsNil(o.ShowNet) {
		return true
	}

	return false
}

// SetShowNet gets a reference to the given bool and assigns it to the ShowNet field.
func (o *ModelOrder) SetShowNet(v bool) {
	o.ShowNet = &v
}

// GetSendType returns the SendType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelOrder) GetSendType() string {
	if o == nil || IsNil(o.SendType.Get()) {
		var ret string
		return ret
	}
	return *o.SendType.Get()
}

// GetSendTypeOk returns a tuple with the SendType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelOrder) GetSendTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SendType.Get(), o.SendType.IsSet()
}

// HasSendType returns a boolean if a field has been set.
func (o *ModelOrder) HasSendType() bool {
	if o != nil && o.SendType.IsSet() {
		return true
	}

	return false
}

// SetSendType gets a reference to the given NullableString and assigns it to the SendType field.
func (o *ModelOrder) SetSendType(v string) {
	o.SendType.Set(&v)
}
// SetSendTypeNil sets the value for SendType to be an explicit nil
func (o *ModelOrder) SetSendTypeNil() {
	o.SendType.Set(nil)
}

// UnsetSendType ensures that no value is present for SendType, not even an explicit nil
func (o *ModelOrder) UnsetSendType() {
	o.SendType.Unset()
}

// GetOrigin returns the Origin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelOrder) GetOrigin() ModelOrderOrigin {
	if o == nil || IsNil(o.Origin.Get()) {
		var ret ModelOrderOrigin
		return ret
	}
	return *o.Origin.Get()
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelOrder) GetOriginOk() (*ModelOrderOrigin, bool) {
	if o == nil {
		return nil, false
	}
	return o.Origin.Get(), o.Origin.IsSet()
}

// HasOrigin returns a boolean if a field has been set.
func (o *ModelOrder) HasOrigin() bool {
	if o != nil && o.Origin.IsSet() {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given NullableModelOrderOrigin and assigns it to the Origin field.
func (o *ModelOrder) SetOrigin(v ModelOrderOrigin) {
	o.Origin.Set(&v)
}
// SetOriginNil sets the value for Origin to be an explicit nil
func (o *ModelOrder) SetOriginNil() {
	o.Origin.Set(nil)
}

// UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
func (o *ModelOrder) UnsetOrigin() {
	o.Origin.Unset()
}

func (o ModelOrder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelOrder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	if !IsNil(o.ObjectName) {
		toSerialize["objectName"] = o.ObjectName
	}
	toSerialize["mapAll"] = o.MapAll
	// skip: create is readOnly
	// skip: update is readOnly
	toSerialize["orderNumber"] = o.OrderNumber
	toSerialize["contact"] = o.Contact
	toSerialize["orderDate"] = o.OrderDate
	toSerialize["status"] = o.Status
	toSerialize["header"] = o.Header
	if o.HeadText.IsSet() {
		toSerialize["headText"] = o.HeadText.Get()
	}
	if o.FootText.IsSet() {
		toSerialize["footText"] = o.FootText.Get()
	}
	toSerialize["addressCountry"] = o.AddressCountry
	if o.DeliveryTerms.IsSet() {
		toSerialize["deliveryTerms"] = o.DeliveryTerms.Get()
	}
	if o.PaymentTerms.IsSet() {
		toSerialize["paymentTerms"] = o.PaymentTerms.Get()
	}
	toSerialize["version"] = o.Version
	if !IsNil(o.SmallSettlement) {
		toSerialize["smallSettlement"] = o.SmallSettlement
	}
	toSerialize["contactPerson"] = o.ContactPerson
	toSerialize["taxRate"] = o.TaxRate
	if o.TaxSet.IsSet() {
		toSerialize["taxSet"] = o.TaxSet.Get()
	}
	toSerialize["taxText"] = o.TaxText
	toSerialize["taxType"] = o.TaxType
	if !IsNil(o.OrderType) {
		toSerialize["orderType"] = o.OrderType
	}
	if o.SendDate.IsSet() {
		toSerialize["sendDate"] = o.SendDate.Get()
	}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	toSerialize["currency"] = o.Currency
	if o.CustomerInternalNote.IsSet() {
		toSerialize["customerInternalNote"] = o.CustomerInternalNote.Get()
	}
	if !IsNil(o.ShowNet) {
		toSerialize["showNet"] = o.ShowNet
	}
	if o.SendType.IsSet() {
		toSerialize["sendType"] = o.SendType.Get()
	}
	if o.Origin.IsSet() {
		toSerialize["origin"] = o.Origin.Get()
	}
	return toSerialize, nil
}

type NullableModelOrder struct {
	value *ModelOrder
	isSet bool
}

func (v NullableModelOrder) Get() *ModelOrder {
	return v.value
}

func (v *NullableModelOrder) Set(val *ModelOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableModelOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableModelOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelOrder(val *ModelOrder) *NullableModelOrder {
	return &NullableModelOrder{value: val, isSet: true}
}

func (v NullableModelOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


