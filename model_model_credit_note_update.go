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

// checks if the ModelCreditNoteUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelCreditNoteUpdate{}

// ModelCreditNoteUpdate creditNote model
type ModelCreditNoteUpdate struct {
	// The creditNote id
	Id *int32 `json:"id,omitempty"`
	// The creditNote object name
	ObjectName *string `json:"objectName,omitempty"`
	// Date of creditNote creation
	Create *time.Time `json:"create,omitempty"`
	// Date of last creditNote update
	Update *time.Time `json:"update,omitempty"`
	// The creditNote number
	CreditNoteNumber NullableString `json:"creditNoteNumber,omitempty"`
	Contact NullableModelCreditNoteUpdateContact `json:"contact,omitempty"`
	// Needs to be provided as timestamp or dd.mm.yyyy
	CreditNoteDate NullableTime `json:"creditNoteDate,omitempty"`
	// Please have a look in       <a href='https://api.sevdesk.de/#section/Types-and-status-of-credit-notes'>status of credit note</a>      to see what the different status codes mean
	Status *string `json:"status,omitempty"`
	// Normally consist of prefix plus the creditNote number
	Header NullableString `json:"header,omitempty"`
	// Certain html tags can be used here to format your text
	HeadText NullableString `json:"headText,omitempty"`
	// Certain html tags can be used here to format your text
	FootText NullableString `json:"footText,omitempty"`
	AddressCountry NullableModelCreditNoteAddressCountry `json:"addressCountry,omitempty"`
	CreateUser *ModelCreditNoteCreateUser `json:"createUser,omitempty"`
	SevClient *ModelCreditNoteSevClient `json:"sevClient,omitempty"`
	// Delivery terms of the creditNote
	DeliveryTerms NullableString `json:"deliveryTerms,omitempty"`
	// Timestamp. This can also be a date range if you also use the attribute deliveryDateUntil
	DeliveryDate *time.Time `json:"deliveryDate,omitempty"`
	// Payment terms of the creditNote
	PaymentTerms NullableString `json:"paymentTerms,omitempty"`
	// Version of the creditNote.<br>      Can be used if you have multiple drafts for the same creditNote.<br>      Should start with 0
	Version NullableInt32 `json:"version,omitempty"`
	// Defines if the client uses the small settlement scheme.      If yes, the creditNote must not contain any vat
	SmallSettlement NullableBool `json:"smallSettlement,omitempty"`
	ContactPerson NullableModelCreditNoteUpdateContactPerson `json:"contactPerson,omitempty"`
	// Is overwritten by creditNote position tax rates
	TaxRate NullableFloat32 `json:"taxRate,omitempty"`
	TaxSet NullableModelCreditNoteTaxSet `json:"taxSet,omitempty"`
	// A common tax text would be 'Umsatzsteuer 19%'
	TaxText NullableString `json:"taxText,omitempty"`
	// Tax type of the creditNote. There are four tax types: 1. default - Umsatzsteuer ausweisen 2. eu - Steuerfreie innergemeinschaftliche Lieferung (Europäische Union) 3. noteu - Steuerschuldnerschaft des Leistungsempfängers (außerhalb EU, z. B. Schweiz) 4. custom - Using custom tax set 5. ss - Not subject to VAT according to §19 1 UStG Tax rates are heavily connected to the tax type used.
	TaxType NullableString `json:"taxType,omitempty"`
	// Type of the creditNote. For more information on the different types, check      <a href='https://api.sevdesk.de/#section/Types-and-status-of-credit-notes'>this</a>  
	CreditNoteType NullableString `json:"creditNoteType,omitempty"`
	// The date the creditNote was sent to the customer
	SendDate NullableTime `json:"sendDate,omitempty"`
	// Complete address of the recipient including name, street, city, zip and country.<br>       Line breaks can be used and will be displayed on the invoice pdf.
	Address NullableString `json:"address,omitempty"`
	// Currency used in the creditNote. Needs to be currency code according to ISO-4217
	Currency NullableString `json:"currency,omitempty"`
	// Net sum of the creditNote
	SumNet *float32 `json:"sumNet,omitempty"`
	// Tax sum of the creditNote
	SumTax *float32 `json:"sumTax,omitempty"`
	// Gross sum of the creditNote
	SumGross *float32 `json:"sumGross,omitempty"`
	// Sum of all discounts in the creditNote
	SumDiscounts *float32 `json:"sumDiscounts,omitempty"`
	// Net sum of the creditNote in the foreign currency
	SumNetForeignCurrency *float32 `json:"sumNetForeignCurrency,omitempty"`
	// Tax sum of the creditNote in the foreign currency
	SumTaxForeignCurrency *float32 `json:"sumTaxForeignCurrency,omitempty"`
	// Gross sum of the creditNote in the foreign currency
	SumGrossForeignCurrency *float32 `json:"sumGrossForeignCurrency,omitempty"`
	// Discounts sum of the creditNote in the foreign currency
	SumDiscountsForeignCurrency *float32 `json:"sumDiscountsForeignCurrency,omitempty"`
	// Internal note of the customer. Contains data entered into field 'Referenz/Bestellnummer'
	CustomerInternalNote NullableString `json:"customerInternalNote,omitempty"`
	// If true, the net amount of each position will be shown on the creditNote. Otherwise gross amount
	ShowNet *bool `json:"showNet,omitempty"`
	// Type which was used to send the creditNote. IMPORTANT: Please refer to the creditNote section of the       *     API-Overview to understand how this attribute can be used before using it!
	SendType NullableString `json:"sendType,omitempty"`
}

// NewModelCreditNoteUpdate instantiates a new ModelCreditNoteUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelCreditNoteUpdate() *ModelCreditNoteUpdate {
	this := ModelCreditNoteUpdate{}
	return &this
}

// NewModelCreditNoteUpdateWithDefaults instantiates a new ModelCreditNoteUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelCreditNoteUpdateWithDefaults() *ModelCreditNoteUpdate {
	this := ModelCreditNoteUpdate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelCreditNoteUpdate) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCreditNoteUpdate) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelCreditNoteUpdate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ModelCreditNoteUpdate) SetId(v int32) {
	o.Id = &v
}

// GetObjectName returns the ObjectName field value if set, zero value otherwise.
func (o *ModelCreditNoteUpdate) GetObjectName() string {
	if o == nil || IsNil(o.ObjectName) {
		var ret string
		return ret
	}
	return *o.ObjectName
}

// GetObjectNameOk returns a tuple with the ObjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCreditNoteUpdate) GetObjectNameOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectName) {
		return nil, false
	}
	return o.ObjectName, true
}

// HasObjectName returns a boolean if a field has been set.
func (o *ModelCreditNoteUpdate) HasObjectName() bool {
	if o != nil && !IsNil(o.ObjectName) {
		return true
	}

	return false
}

// SetObjectName gets a reference to the given string and assigns it to the ObjectName field.
func (o *ModelCreditNoteUpdate) SetObjectName(v string) {
	o.ObjectName = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *ModelCreditNoteUpdate) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCreditNoteUpdate) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *ModelCreditNoteUpdate) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *ModelCreditNoteUpdate) SetCreate(v time.Time) {
	o.Create = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *ModelCreditNoteUpdate) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCreditNoteUpdate) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *ModelCreditNoteUpdate) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *ModelCreditNoteUpdate) SetUpdate(v time.Time) {
	o.Update = &v
}

// GetCreditNoteNumber returns the CreditNoteNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCreditNoteUpdate) GetCreditNoteNumber() string {
	if o == nil || IsNil(o.CreditNoteNumber.Get()) {
		var ret string
		return ret
	}
	return *o.CreditNoteNumber.Get()
}

// GetCreditNoteNumberOk returns a tuple with the CreditNoteNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCreditNoteUpdate) GetCreditNoteNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreditNoteNumber.Get(), o.CreditNoteNumber.IsSet()
}

// HasCreditNoteNumber returns a boolean if a field has been set.
func (o *ModelCreditNoteUpdate) HasCreditNoteNumber() bool {
	if o != nil && o.CreditNoteNumber.IsSet() {
		return true
	}

	return false
}

// SetCreditNoteNumber gets a reference to the given NullableString and assigns it to the CreditNoteNumber field.
func (o *ModelCreditNoteUpdate) SetCreditNoteNumber(v string) {
	o.CreditNoteNumber.Set(&v)
}
// SetCreditNoteNumberNil sets the value for CreditNoteNumber to be an explicit nil
func (o *ModelCreditNoteUpdate) SetCreditNoteNumberNil() {
	o.CreditNoteNumber.Set(nil)
}

// UnsetCreditNoteNumber ensures that no value is present for CreditNoteNumber, not even an explicit nil
func (o *ModelCreditNoteUpdate) UnsetCreditNoteNumber() {
	o.CreditNoteNumber.Unset()
}

// GetContact returns the Contact field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCreditNoteUpdate) GetContact() ModelCreditNoteUpdateContact {
	if o == nil || IsNil(o.Contact.Get()) {
		var ret ModelCreditNoteUpdateContact
		return ret
	}
	return *o.Contact.Get()
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCreditNoteUpdate) GetContactOk() (*ModelCreditNoteUpdateContact, bool) {
	if o == nil {
		return nil, false
	}
	return o.Contact.Get(), o.Contact.IsSet()
}

// HasContact returns a boolean if a field has been set.
func (o *ModelCreditNoteUpdate) HasContact() bool {
	if o != nil && o.Contact.IsSet() {
		return true
	}

	return false
}

// SetContact gets a reference to the given NullableModelCreditNoteUpdateContact and assigns it to the Contact field.
func (o *ModelCreditNoteUpdate) SetContact(v ModelCreditNoteUpdateContact) {
	o.Contact.Set(&v)
}
// SetContactNil sets the value for Contact to be an explicit nil
func (o *ModelCreditNoteUpdate) SetContactNil() {
	o.Contact.Set(nil)
}

// UnsetContact ensures that no value is present for Contact, not even an explicit nil
func (o *ModelCreditNoteUpdate) UnsetContact() {
	o.Contact.Unset()
}

// GetCreditNoteDate returns the CreditNoteDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCreditNoteUpdate) GetCreditNoteDate() time.Time {
	if o == nil || IsNil(o.CreditNoteDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CreditNoteDate.Get()
}

// GetCreditNoteDateOk returns a tuple with the CreditNoteDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCreditNoteUpdate) GetCreditNoteDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreditNoteDate.Get(), o.CreditNoteDate.IsSet()
}

// HasCreditNoteDate returns a boolean if a field has been set.
func (o *ModelCreditNoteUpdate) HasCreditNoteDate() bool {
	if o != nil && o.CreditNoteDate.IsSet() {
		return true
	}

	return false
}

// SetCreditNoteDate gets a reference to the given NullableTime and assigns it to the CreditNoteDate field.
func (o *ModelCreditNoteUpdate) SetCreditNoteDate(v time.Time) {
	o.CreditNoteDate.Set(&v)
}
// SetCreditNoteDateNil sets the value for CreditNoteDate to be an explicit nil
func (o *ModelCreditNoteUpdate) SetCreditNoteDateNil() {
	o.CreditNoteDate.Set(nil)
}

// UnsetCreditNoteDate ensures that no value is present for CreditNoteDate, not even an explicit nil
func (o *ModelCreditNoteUpdate) UnsetCreditNoteDate() {
	o.CreditNoteDate.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ModelCreditNoteUpdate) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCreditNoteUpdate) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ModelCreditNoteUpdate) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ModelCreditNoteUpdate) SetStatus(v string) {
	o.Status = &v
}

// GetHeader returns the Header field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCreditNoteUpdate) GetHeader() string {
	if o == nil || IsNil(o.Header.Get()) {
		var ret string
		return ret
	}
	return *o.Header.Get()
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCreditNoteUpdate) GetHeaderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Header.Get(), o.Header.IsSet()
}

// HasHeader returns a boolean if a field has been set.
func (o *ModelCreditNoteUpdate) HasHeader() bool {
	if o != nil && o.Header.IsSet() {
		return true
	}

	return false
}

// SetHeader gets a reference to the given NullableString and assigns it to the Header field.
func (o *ModelCreditNoteUpdate) SetHeader(v string) {
	o.Header.Set(&v)
}
// SetHeaderNil sets the value for Header to be an explicit nil
func (o *ModelCreditNoteUpdate) SetHeaderNil() {
	o.Header.Set(nil)
}

// UnsetHeader ensures that no value is present for Header, not even an explicit nil
func (o *ModelCreditNoteUpdate) UnsetHeader() {
	o.Header.Unset()
}

// GetHeadText returns the HeadText field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCreditNoteUpdate) GetHeadText() string {
	if o == nil || IsNil(o.HeadText.Get()) {
		var ret string
		return ret
	}
	return *o.HeadText.Get()
}

// GetHeadTextOk returns a tuple with the HeadText field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCreditNoteUpdate) GetHeadTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HeadText.Get(), o.HeadText.IsSet()
}

// HasHeadText returns a boolean if a field has been set.
func (o *ModelCreditNoteUpdate) HasHeadText() bool {
	if o != nil && o.HeadText.IsSet() {
		return true
	}

	return false
}

// SetHeadText gets a reference to the given NullableString and assigns it to the HeadText field.
func (o *ModelCreditNoteUpdate) SetHeadText(v string) {
	o.HeadText.Set(&v)
}
// SetHeadTextNil sets the value for HeadText to be an explicit nil
func (o *ModelCreditNoteUpdate) SetHeadTextNil() {
	o.HeadText.Set(nil)
}

// UnsetHeadText ensures that no value is present for HeadText, not even an explicit nil
func (o *ModelCreditNoteUpdate) UnsetHeadText() {
	o.HeadText.Unset()
}

// GetFootText returns the FootText field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCreditNoteUpdate) GetFootText() string {
	if o == nil || IsNil(o.FootText.Get()) {
		var ret string
		return ret
	}
	return *o.FootText.Get()
}

// GetFootTextOk returns a tuple with the FootText field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCreditNoteUpdate) GetFootTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FootText.Get(), o.FootText.IsSet()
}

// HasFootText returns a boolean if a field has been set.
func (o *ModelCreditNoteUpdate) HasFootText() bool {
	if o != nil && o.FootText.IsSet() {
		return true
	}

	return false
}

// SetFootText gets a reference to the given NullableString and assigns it to the FootText field.
func (o *ModelCreditNoteUpdate) SetFootText(v string) {
	o.FootText.Set(&v)
}
// SetFootTextNil sets the value for FootText to be an explicit nil
func (o *ModelCreditNoteUpdate) SetFootTextNil() {
	o.FootText.Set(nil)
}

// UnsetFootText ensures that no value is present for FootText, not even an explicit nil
func (o *ModelCreditNoteUpdate) UnsetFootText() {
	o.FootText.Unset()
}

// GetAddressCountry returns the AddressCountry field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCreditNoteUpdate) GetAddressCountry() ModelCreditNoteAddressCountry {
	if o == nil || IsNil(o.AddressCountry.Get()) {
		var ret ModelCreditNoteAddressCountry
		return ret
	}
	return *o.AddressCountry.Get()
}

// GetAddressCountryOk returns a tuple with the AddressCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCreditNoteUpdate) GetAddressCountryOk() (*ModelCreditNoteAddressCountry, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddressCountry.Get(), o.AddressCountry.IsSet()
}

// HasAddressCountry returns a boolean if a field has been set.
func (o *ModelCreditNoteUpdate) HasAddressCountry() bool {
	if o != nil && o.AddressCountry.IsSet() {
		return true
	}

	return false
}

// SetAddressCountry gets a reference to the given NullableModelCreditNoteAddressCountry and assigns it to the AddressCountry field.
func (o *ModelCreditNoteUpdate) SetAddressCountry(v ModelCreditNoteAddressCountry) {
	o.AddressCountry.Set(&v)
}
// SetAddressCountryNil sets the value for AddressCountry to be an explicit nil
func (o *ModelCreditNoteUpdate) SetAddressCountryNil() {
	o.AddressCountry.Set(nil)
}

// UnsetAddressCountry ensures that no value is present for AddressCountry, not even an explicit nil
func (o *ModelCreditNoteUpdate) UnsetAddressCountry() {
	o.AddressCountry.Unset()
}

// GetCreateUser returns the CreateUser field value if set, zero value otherwise.
func (o *ModelCreditNoteUpdate) GetCreateUser() ModelCreditNoteCreateUser {
	if o == nil || IsNil(o.CreateUser) {
		var ret ModelCreditNoteCreateUser
		return ret
	}
	return *o.CreateUser
}

// GetCreateUserOk returns a tuple with the CreateUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCreditNoteUpdate) GetCreateUserOk() (*ModelCreditNoteCreateUser, bool) {
	if o == nil || IsNil(o.CreateUser) {
		return nil, false
	}
	return o.CreateUser, true
}

// HasCreateUser returns a boolean if a field has been set.
func (o *ModelCreditNoteUpdate) HasCreateUser() bool {
	if o != nil && !IsNil(o.CreateUser) {
		return true
	}

	return false
}

// SetCreateUser gets a reference to the given ModelCreditNoteCreateUser and assigns it to the CreateUser field.
func (o *ModelCreditNoteUpdate) SetCreateUser(v ModelCreditNoteCreateUser) {
	o.CreateUser = &v
}

// GetSevClient returns the SevClient field value if set, zero value otherwise.
func (o *ModelCreditNoteUpdate) GetSevClient() ModelCreditNoteSevClient {
	if o == nil || IsNil(o.SevClient) {
		var ret ModelCreditNoteSevClient
		return ret
	}
	return *o.SevClient
}

// GetSevClientOk returns a tuple with the SevClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCreditNoteUpdate) GetSevClientOk() (*ModelCreditNoteSevClient, bool) {
	if o == nil || IsNil(o.SevClient) {
		return nil, false
	}
	return o.SevClient, true
}

// HasSevClient returns a boolean if a field has been set.
func (o *ModelCreditNoteUpdate) HasSevClient() bool {
	if o != nil && !IsNil(o.SevClient) {
		return true
	}

	return false
}

// SetSevClient gets a reference to the given ModelCreditNoteSevClient and assigns it to the SevClient field.
func (o *ModelCreditNoteUpdate) SetSevClient(v ModelCreditNoteSevClient) {
	o.SevClient = &v
}

// GetDeliveryTerms returns the DeliveryTerms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCreditNoteUpdate) GetDeliveryTerms() string {
	if o == nil || IsNil(o.DeliveryTerms.Get()) {
		var ret string
		return ret
	}
	return *o.DeliveryTerms.Get()
}

// GetDeliveryTermsOk returns a tuple with the DeliveryTerms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCreditNoteUpdate) GetDeliveryTermsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeliveryTerms.Get(), o.DeliveryTerms.IsSet()
}

// HasDeliveryTerms returns a boolean if a field has been set.
func (o *ModelCreditNoteUpdate) HasDeliveryTerms() bool {
	if o != nil && o.DeliveryTerms.IsSet() {
		return true
	}

	return false
}

// SetDeliveryTerms gets a reference to the given NullableString and assigns it to the DeliveryTerms field.
func (o *ModelCreditNoteUpdate) SetDeliveryTerms(v string) {
	o.DeliveryTerms.Set(&v)
}
// SetDeliveryTermsNil sets the value for DeliveryTerms to be an explicit nil
func (o *ModelCreditNoteUpdate) SetDeliveryTermsNil() {
	o.DeliveryTerms.Set(nil)
}

// UnsetDeliveryTerms ensures that no value is present for DeliveryTerms, not even an explicit nil
func (o *ModelCreditNoteUpdate) UnsetDeliveryTerms() {
	o.DeliveryTerms.Unset()
}

// GetDeliveryDate returns the DeliveryDate field value if set, zero value otherwise.
func (o *ModelCreditNoteUpdate) GetDeliveryDate() time.Time {
	if o == nil || IsNil(o.DeliveryDate) {
		var ret time.Time
		return ret
	}
	return *o.DeliveryDate
}

// GetDeliveryDateOk returns a tuple with the DeliveryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCreditNoteUpdate) GetDeliveryDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeliveryDate) {
		return nil, false
	}
	return o.DeliveryDate, true
}

// HasDeliveryDate returns a boolean if a field has been set.
func (o *ModelCreditNoteUpdate) HasDeliveryDate() bool {
	if o != nil && !IsNil(o.DeliveryDate) {
		return true
	}

	return false
}

// SetDeliveryDate gets a reference to the given time.Time and assigns it to the DeliveryDate field.
func (o *ModelCreditNoteUpdate) SetDeliveryDate(v time.Time) {
	o.DeliveryDate = &v
}

// GetPaymentTerms returns the PaymentTerms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCreditNoteUpdate) GetPaymentTerms() string {
	if o == nil || IsNil(o.PaymentTerms.Get()) {
		var ret string
		return ret
	}
	return *o.PaymentTerms.Get()
}

// GetPaymentTermsOk returns a tuple with the PaymentTerms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCreditNoteUpdate) GetPaymentTermsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentTerms.Get(), o.PaymentTerms.IsSet()
}

// HasPaymentTerms returns a boolean if a field has been set.
func (o *ModelCreditNoteUpdate) HasPaymentTerms() bool {
	if o != nil && o.PaymentTerms.IsSet() {
		return true
	}

	return false
}

// SetPaymentTerms gets a reference to the given NullableString and assigns it to the PaymentTerms field.
func (o *ModelCreditNoteUpdate) SetPaymentTerms(v string) {
	o.PaymentTerms.Set(&v)
}
// SetPaymentTermsNil sets the value for PaymentTerms to be an explicit nil
func (o *ModelCreditNoteUpdate) SetPaymentTermsNil() {
	o.PaymentTerms.Set(nil)
}

// UnsetPaymentTerms ensures that no value is present for PaymentTerms, not even an explicit nil
func (o *ModelCreditNoteUpdate) UnsetPaymentTerms() {
	o.PaymentTerms.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCreditNoteUpdate) GetVersion() int32 {
	if o == nil || IsNil(o.Version.Get()) {
		var ret int32
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCreditNoteUpdate) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *ModelCreditNoteUpdate) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableInt32 and assigns it to the Version field.
func (o *ModelCreditNoteUpdate) SetVersion(v int32) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *ModelCreditNoteUpdate) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *ModelCreditNoteUpdate) UnsetVersion() {
	o.Version.Unset()
}

// GetSmallSettlement returns the SmallSettlement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCreditNoteUpdate) GetSmallSettlement() bool {
	if o == nil || IsNil(o.SmallSettlement.Get()) {
		var ret bool
		return ret
	}
	return *o.SmallSettlement.Get()
}

// GetSmallSettlementOk returns a tuple with the SmallSettlement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCreditNoteUpdate) GetSmallSettlementOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SmallSettlement.Get(), o.SmallSettlement.IsSet()
}

// HasSmallSettlement returns a boolean if a field has been set.
func (o *ModelCreditNoteUpdate) HasSmallSettlement() bool {
	if o != nil && o.SmallSettlement.IsSet() {
		return true
	}

	return false
}

// SetSmallSettlement gets a reference to the given NullableBool and assigns it to the SmallSettlement field.
func (o *ModelCreditNoteUpdate) SetSmallSettlement(v bool) {
	o.SmallSettlement.Set(&v)
}
// SetSmallSettlementNil sets the value for SmallSettlement to be an explicit nil
func (o *ModelCreditNoteUpdate) SetSmallSettlementNil() {
	o.SmallSettlement.Set(nil)
}

// UnsetSmallSettlement ensures that no value is present for SmallSettlement, not even an explicit nil
func (o *ModelCreditNoteUpdate) UnsetSmallSettlement() {
	o.SmallSettlement.Unset()
}

// GetContactPerson returns the ContactPerson field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCreditNoteUpdate) GetContactPerson() ModelCreditNoteUpdateContactPerson {
	if o == nil || IsNil(o.ContactPerson.Get()) {
		var ret ModelCreditNoteUpdateContactPerson
		return ret
	}
	return *o.ContactPerson.Get()
}

// GetContactPersonOk returns a tuple with the ContactPerson field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCreditNoteUpdate) GetContactPersonOk() (*ModelCreditNoteUpdateContactPerson, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContactPerson.Get(), o.ContactPerson.IsSet()
}

// HasContactPerson returns a boolean if a field has been set.
func (o *ModelCreditNoteUpdate) HasContactPerson() bool {
	if o != nil && o.ContactPerson.IsSet() {
		return true
	}

	return false
}

// SetContactPerson gets a reference to the given NullableModelCreditNoteUpdateContactPerson and assigns it to the ContactPerson field.
func (o *ModelCreditNoteUpdate) SetContactPerson(v ModelCreditNoteUpdateContactPerson) {
	o.ContactPerson.Set(&v)
}
// SetContactPersonNil sets the value for ContactPerson to be an explicit nil
func (o *ModelCreditNoteUpdate) SetContactPersonNil() {
	o.ContactPerson.Set(nil)
}

// UnsetContactPerson ensures that no value is present for ContactPerson, not even an explicit nil
func (o *ModelCreditNoteUpdate) UnsetContactPerson() {
	o.ContactPerson.Unset()
}

// GetTaxRate returns the TaxRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCreditNoteUpdate) GetTaxRate() float32 {
	if o == nil || IsNil(o.TaxRate.Get()) {
		var ret float32
		return ret
	}
	return *o.TaxRate.Get()
}

// GetTaxRateOk returns a tuple with the TaxRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCreditNoteUpdate) GetTaxRateOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxRate.Get(), o.TaxRate.IsSet()
}

// HasTaxRate returns a boolean if a field has been set.
func (o *ModelCreditNoteUpdate) HasTaxRate() bool {
	if o != nil && o.TaxRate.IsSet() {
		return true
	}

	return false
}

// SetTaxRate gets a reference to the given NullableFloat32 and assigns it to the TaxRate field.
func (o *ModelCreditNoteUpdate) SetTaxRate(v float32) {
	o.TaxRate.Set(&v)
}
// SetTaxRateNil sets the value for TaxRate to be an explicit nil
func (o *ModelCreditNoteUpdate) SetTaxRateNil() {
	o.TaxRate.Set(nil)
}

// UnsetTaxRate ensures that no value is present for TaxRate, not even an explicit nil
func (o *ModelCreditNoteUpdate) UnsetTaxRate() {
	o.TaxRate.Unset()
}

// GetTaxSet returns the TaxSet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCreditNoteUpdate) GetTaxSet() ModelCreditNoteTaxSet {
	if o == nil || IsNil(o.TaxSet.Get()) {
		var ret ModelCreditNoteTaxSet
		return ret
	}
	return *o.TaxSet.Get()
}

// GetTaxSetOk returns a tuple with the TaxSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCreditNoteUpdate) GetTaxSetOk() (*ModelCreditNoteTaxSet, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxSet.Get(), o.TaxSet.IsSet()
}

// HasTaxSet returns a boolean if a field has been set.
func (o *ModelCreditNoteUpdate) HasTaxSet() bool {
	if o != nil && o.TaxSet.IsSet() {
		return true
	}

	return false
}

// SetTaxSet gets a reference to the given NullableModelCreditNoteTaxSet and assigns it to the TaxSet field.
func (o *ModelCreditNoteUpdate) SetTaxSet(v ModelCreditNoteTaxSet) {
	o.TaxSet.Set(&v)
}
// SetTaxSetNil sets the value for TaxSet to be an explicit nil
func (o *ModelCreditNoteUpdate) SetTaxSetNil() {
	o.TaxSet.Set(nil)
}

// UnsetTaxSet ensures that no value is present for TaxSet, not even an explicit nil
func (o *ModelCreditNoteUpdate) UnsetTaxSet() {
	o.TaxSet.Unset()
}

// GetTaxText returns the TaxText field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCreditNoteUpdate) GetTaxText() string {
	if o == nil || IsNil(o.TaxText.Get()) {
		var ret string
		return ret
	}
	return *o.TaxText.Get()
}

// GetTaxTextOk returns a tuple with the TaxText field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCreditNoteUpdate) GetTaxTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxText.Get(), o.TaxText.IsSet()
}

// HasTaxText returns a boolean if a field has been set.
func (o *ModelCreditNoteUpdate) HasTaxText() bool {
	if o != nil && o.TaxText.IsSet() {
		return true
	}

	return false
}

// SetTaxText gets a reference to the given NullableString and assigns it to the TaxText field.
func (o *ModelCreditNoteUpdate) SetTaxText(v string) {
	o.TaxText.Set(&v)
}
// SetTaxTextNil sets the value for TaxText to be an explicit nil
func (o *ModelCreditNoteUpdate) SetTaxTextNil() {
	o.TaxText.Set(nil)
}

// UnsetTaxText ensures that no value is present for TaxText, not even an explicit nil
func (o *ModelCreditNoteUpdate) UnsetTaxText() {
	o.TaxText.Unset()
}

// GetTaxType returns the TaxType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCreditNoteUpdate) GetTaxType() string {
	if o == nil || IsNil(o.TaxType.Get()) {
		var ret string
		return ret
	}
	return *o.TaxType.Get()
}

// GetTaxTypeOk returns a tuple with the TaxType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCreditNoteUpdate) GetTaxTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxType.Get(), o.TaxType.IsSet()
}

// HasTaxType returns a boolean if a field has been set.
func (o *ModelCreditNoteUpdate) HasTaxType() bool {
	if o != nil && o.TaxType.IsSet() {
		return true
	}

	return false
}

// SetTaxType gets a reference to the given NullableString and assigns it to the TaxType field.
func (o *ModelCreditNoteUpdate) SetTaxType(v string) {
	o.TaxType.Set(&v)
}
// SetTaxTypeNil sets the value for TaxType to be an explicit nil
func (o *ModelCreditNoteUpdate) SetTaxTypeNil() {
	o.TaxType.Set(nil)
}

// UnsetTaxType ensures that no value is present for TaxType, not even an explicit nil
func (o *ModelCreditNoteUpdate) UnsetTaxType() {
	o.TaxType.Unset()
}

// GetCreditNoteType returns the CreditNoteType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCreditNoteUpdate) GetCreditNoteType() string {
	if o == nil || IsNil(o.CreditNoteType.Get()) {
		var ret string
		return ret
	}
	return *o.CreditNoteType.Get()
}

// GetCreditNoteTypeOk returns a tuple with the CreditNoteType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCreditNoteUpdate) GetCreditNoteTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreditNoteType.Get(), o.CreditNoteType.IsSet()
}

// HasCreditNoteType returns a boolean if a field has been set.
func (o *ModelCreditNoteUpdate) HasCreditNoteType() bool {
	if o != nil && o.CreditNoteType.IsSet() {
		return true
	}

	return false
}

// SetCreditNoteType gets a reference to the given NullableString and assigns it to the CreditNoteType field.
func (o *ModelCreditNoteUpdate) SetCreditNoteType(v string) {
	o.CreditNoteType.Set(&v)
}
// SetCreditNoteTypeNil sets the value for CreditNoteType to be an explicit nil
func (o *ModelCreditNoteUpdate) SetCreditNoteTypeNil() {
	o.CreditNoteType.Set(nil)
}

// UnsetCreditNoteType ensures that no value is present for CreditNoteType, not even an explicit nil
func (o *ModelCreditNoteUpdate) UnsetCreditNoteType() {
	o.CreditNoteType.Unset()
}

// GetSendDate returns the SendDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCreditNoteUpdate) GetSendDate() time.Time {
	if o == nil || IsNil(o.SendDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.SendDate.Get()
}

// GetSendDateOk returns a tuple with the SendDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCreditNoteUpdate) GetSendDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.SendDate.Get(), o.SendDate.IsSet()
}

// HasSendDate returns a boolean if a field has been set.
func (o *ModelCreditNoteUpdate) HasSendDate() bool {
	if o != nil && o.SendDate.IsSet() {
		return true
	}

	return false
}

// SetSendDate gets a reference to the given NullableTime and assigns it to the SendDate field.
func (o *ModelCreditNoteUpdate) SetSendDate(v time.Time) {
	o.SendDate.Set(&v)
}
// SetSendDateNil sets the value for SendDate to be an explicit nil
func (o *ModelCreditNoteUpdate) SetSendDateNil() {
	o.SendDate.Set(nil)
}

// UnsetSendDate ensures that no value is present for SendDate, not even an explicit nil
func (o *ModelCreditNoteUpdate) UnsetSendDate() {
	o.SendDate.Unset()
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCreditNoteUpdate) GetAddress() string {
	if o == nil || IsNil(o.Address.Get()) {
		var ret string
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCreditNoteUpdate) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *ModelCreditNoteUpdate) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableString and assigns it to the Address field.
func (o *ModelCreditNoteUpdate) SetAddress(v string) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *ModelCreditNoteUpdate) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *ModelCreditNoteUpdate) UnsetAddress() {
	o.Address.Unset()
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCreditNoteUpdate) GetCurrency() string {
	if o == nil || IsNil(o.Currency.Get()) {
		var ret string
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCreditNoteUpdate) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *ModelCreditNoteUpdate) HasCurrency() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableString and assigns it to the Currency field.
func (o *ModelCreditNoteUpdate) SetCurrency(v string) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *ModelCreditNoteUpdate) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *ModelCreditNoteUpdate) UnsetCurrency() {
	o.Currency.Unset()
}

// GetSumNet returns the SumNet field value if set, zero value otherwise.
func (o *ModelCreditNoteUpdate) GetSumNet() float32 {
	if o == nil || IsNil(o.SumNet) {
		var ret float32
		return ret
	}
	return *o.SumNet
}

// GetSumNetOk returns a tuple with the SumNet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCreditNoteUpdate) GetSumNetOk() (*float32, bool) {
	if o == nil || IsNil(o.SumNet) {
		return nil, false
	}
	return o.SumNet, true
}

// HasSumNet returns a boolean if a field has been set.
func (o *ModelCreditNoteUpdate) HasSumNet() bool {
	if o != nil && !IsNil(o.SumNet) {
		return true
	}

	return false
}

// SetSumNet gets a reference to the given float32 and assigns it to the SumNet field.
func (o *ModelCreditNoteUpdate) SetSumNet(v float32) {
	o.SumNet = &v
}

// GetSumTax returns the SumTax field value if set, zero value otherwise.
func (o *ModelCreditNoteUpdate) GetSumTax() float32 {
	if o == nil || IsNil(o.SumTax) {
		var ret float32
		return ret
	}
	return *o.SumTax
}

// GetSumTaxOk returns a tuple with the SumTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCreditNoteUpdate) GetSumTaxOk() (*float32, bool) {
	if o == nil || IsNil(o.SumTax) {
		return nil, false
	}
	return o.SumTax, true
}

// HasSumTax returns a boolean if a field has been set.
func (o *ModelCreditNoteUpdate) HasSumTax() bool {
	if o != nil && !IsNil(o.SumTax) {
		return true
	}

	return false
}

// SetSumTax gets a reference to the given float32 and assigns it to the SumTax field.
func (o *ModelCreditNoteUpdate) SetSumTax(v float32) {
	o.SumTax = &v
}

// GetSumGross returns the SumGross field value if set, zero value otherwise.
func (o *ModelCreditNoteUpdate) GetSumGross() float32 {
	if o == nil || IsNil(o.SumGross) {
		var ret float32
		return ret
	}
	return *o.SumGross
}

// GetSumGrossOk returns a tuple with the SumGross field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCreditNoteUpdate) GetSumGrossOk() (*float32, bool) {
	if o == nil || IsNil(o.SumGross) {
		return nil, false
	}
	return o.SumGross, true
}

// HasSumGross returns a boolean if a field has been set.
func (o *ModelCreditNoteUpdate) HasSumGross() bool {
	if o != nil && !IsNil(o.SumGross) {
		return true
	}

	return false
}

// SetSumGross gets a reference to the given float32 and assigns it to the SumGross field.
func (o *ModelCreditNoteUpdate) SetSumGross(v float32) {
	o.SumGross = &v
}

// GetSumDiscounts returns the SumDiscounts field value if set, zero value otherwise.
func (o *ModelCreditNoteUpdate) GetSumDiscounts() float32 {
	if o == nil || IsNil(o.SumDiscounts) {
		var ret float32
		return ret
	}
	return *o.SumDiscounts
}

// GetSumDiscountsOk returns a tuple with the SumDiscounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCreditNoteUpdate) GetSumDiscountsOk() (*float32, bool) {
	if o == nil || IsNil(o.SumDiscounts) {
		return nil, false
	}
	return o.SumDiscounts, true
}

// HasSumDiscounts returns a boolean if a field has been set.
func (o *ModelCreditNoteUpdate) HasSumDiscounts() bool {
	if o != nil && !IsNil(o.SumDiscounts) {
		return true
	}

	return false
}

// SetSumDiscounts gets a reference to the given float32 and assigns it to the SumDiscounts field.
func (o *ModelCreditNoteUpdate) SetSumDiscounts(v float32) {
	o.SumDiscounts = &v
}

// GetSumNetForeignCurrency returns the SumNetForeignCurrency field value if set, zero value otherwise.
func (o *ModelCreditNoteUpdate) GetSumNetForeignCurrency() float32 {
	if o == nil || IsNil(o.SumNetForeignCurrency) {
		var ret float32
		return ret
	}
	return *o.SumNetForeignCurrency
}

// GetSumNetForeignCurrencyOk returns a tuple with the SumNetForeignCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCreditNoteUpdate) GetSumNetForeignCurrencyOk() (*float32, bool) {
	if o == nil || IsNil(o.SumNetForeignCurrency) {
		return nil, false
	}
	return o.SumNetForeignCurrency, true
}

// HasSumNetForeignCurrency returns a boolean if a field has been set.
func (o *ModelCreditNoteUpdate) HasSumNetForeignCurrency() bool {
	if o != nil && !IsNil(o.SumNetForeignCurrency) {
		return true
	}

	return false
}

// SetSumNetForeignCurrency gets a reference to the given float32 and assigns it to the SumNetForeignCurrency field.
func (o *ModelCreditNoteUpdate) SetSumNetForeignCurrency(v float32) {
	o.SumNetForeignCurrency = &v
}

// GetSumTaxForeignCurrency returns the SumTaxForeignCurrency field value if set, zero value otherwise.
func (o *ModelCreditNoteUpdate) GetSumTaxForeignCurrency() float32 {
	if o == nil || IsNil(o.SumTaxForeignCurrency) {
		var ret float32
		return ret
	}
	return *o.SumTaxForeignCurrency
}

// GetSumTaxForeignCurrencyOk returns a tuple with the SumTaxForeignCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCreditNoteUpdate) GetSumTaxForeignCurrencyOk() (*float32, bool) {
	if o == nil || IsNil(o.SumTaxForeignCurrency) {
		return nil, false
	}
	return o.SumTaxForeignCurrency, true
}

// HasSumTaxForeignCurrency returns a boolean if a field has been set.
func (o *ModelCreditNoteUpdate) HasSumTaxForeignCurrency() bool {
	if o != nil && !IsNil(o.SumTaxForeignCurrency) {
		return true
	}

	return false
}

// SetSumTaxForeignCurrency gets a reference to the given float32 and assigns it to the SumTaxForeignCurrency field.
func (o *ModelCreditNoteUpdate) SetSumTaxForeignCurrency(v float32) {
	o.SumTaxForeignCurrency = &v
}

// GetSumGrossForeignCurrency returns the SumGrossForeignCurrency field value if set, zero value otherwise.
func (o *ModelCreditNoteUpdate) GetSumGrossForeignCurrency() float32 {
	if o == nil || IsNil(o.SumGrossForeignCurrency) {
		var ret float32
		return ret
	}
	return *o.SumGrossForeignCurrency
}

// GetSumGrossForeignCurrencyOk returns a tuple with the SumGrossForeignCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCreditNoteUpdate) GetSumGrossForeignCurrencyOk() (*float32, bool) {
	if o == nil || IsNil(o.SumGrossForeignCurrency) {
		return nil, false
	}
	return o.SumGrossForeignCurrency, true
}

// HasSumGrossForeignCurrency returns a boolean if a field has been set.
func (o *ModelCreditNoteUpdate) HasSumGrossForeignCurrency() bool {
	if o != nil && !IsNil(o.SumGrossForeignCurrency) {
		return true
	}

	return false
}

// SetSumGrossForeignCurrency gets a reference to the given float32 and assigns it to the SumGrossForeignCurrency field.
func (o *ModelCreditNoteUpdate) SetSumGrossForeignCurrency(v float32) {
	o.SumGrossForeignCurrency = &v
}

// GetSumDiscountsForeignCurrency returns the SumDiscountsForeignCurrency field value if set, zero value otherwise.
func (o *ModelCreditNoteUpdate) GetSumDiscountsForeignCurrency() float32 {
	if o == nil || IsNil(o.SumDiscountsForeignCurrency) {
		var ret float32
		return ret
	}
	return *o.SumDiscountsForeignCurrency
}

// GetSumDiscountsForeignCurrencyOk returns a tuple with the SumDiscountsForeignCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCreditNoteUpdate) GetSumDiscountsForeignCurrencyOk() (*float32, bool) {
	if o == nil || IsNil(o.SumDiscountsForeignCurrency) {
		return nil, false
	}
	return o.SumDiscountsForeignCurrency, true
}

// HasSumDiscountsForeignCurrency returns a boolean if a field has been set.
func (o *ModelCreditNoteUpdate) HasSumDiscountsForeignCurrency() bool {
	if o != nil && !IsNil(o.SumDiscountsForeignCurrency) {
		return true
	}

	return false
}

// SetSumDiscountsForeignCurrency gets a reference to the given float32 and assigns it to the SumDiscountsForeignCurrency field.
func (o *ModelCreditNoteUpdate) SetSumDiscountsForeignCurrency(v float32) {
	o.SumDiscountsForeignCurrency = &v
}

// GetCustomerInternalNote returns the CustomerInternalNote field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCreditNoteUpdate) GetCustomerInternalNote() string {
	if o == nil || IsNil(o.CustomerInternalNote.Get()) {
		var ret string
		return ret
	}
	return *o.CustomerInternalNote.Get()
}

// GetCustomerInternalNoteOk returns a tuple with the CustomerInternalNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCreditNoteUpdate) GetCustomerInternalNoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerInternalNote.Get(), o.CustomerInternalNote.IsSet()
}

// HasCustomerInternalNote returns a boolean if a field has been set.
func (o *ModelCreditNoteUpdate) HasCustomerInternalNote() bool {
	if o != nil && o.CustomerInternalNote.IsSet() {
		return true
	}

	return false
}

// SetCustomerInternalNote gets a reference to the given NullableString and assigns it to the CustomerInternalNote field.
func (o *ModelCreditNoteUpdate) SetCustomerInternalNote(v string) {
	o.CustomerInternalNote.Set(&v)
}
// SetCustomerInternalNoteNil sets the value for CustomerInternalNote to be an explicit nil
func (o *ModelCreditNoteUpdate) SetCustomerInternalNoteNil() {
	o.CustomerInternalNote.Set(nil)
}

// UnsetCustomerInternalNote ensures that no value is present for CustomerInternalNote, not even an explicit nil
func (o *ModelCreditNoteUpdate) UnsetCustomerInternalNote() {
	o.CustomerInternalNote.Unset()
}

// GetShowNet returns the ShowNet field value if set, zero value otherwise.
func (o *ModelCreditNoteUpdate) GetShowNet() bool {
	if o == nil || IsNil(o.ShowNet) {
		var ret bool
		return ret
	}
	return *o.ShowNet
}

// GetShowNetOk returns a tuple with the ShowNet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCreditNoteUpdate) GetShowNetOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowNet) {
		return nil, false
	}
	return o.ShowNet, true
}

// HasShowNet returns a boolean if a field has been set.
func (o *ModelCreditNoteUpdate) HasShowNet() bool {
	if o != nil && !IsNil(o.ShowNet) {
		return true
	}

	return false
}

// SetShowNet gets a reference to the given bool and assigns it to the ShowNet field.
func (o *ModelCreditNoteUpdate) SetShowNet(v bool) {
	o.ShowNet = &v
}

// GetSendType returns the SendType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCreditNoteUpdate) GetSendType() string {
	if o == nil || IsNil(o.SendType.Get()) {
		var ret string
		return ret
	}
	return *o.SendType.Get()
}

// GetSendTypeOk returns a tuple with the SendType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCreditNoteUpdate) GetSendTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SendType.Get(), o.SendType.IsSet()
}

// HasSendType returns a boolean if a field has been set.
func (o *ModelCreditNoteUpdate) HasSendType() bool {
	if o != nil && o.SendType.IsSet() {
		return true
	}

	return false
}

// SetSendType gets a reference to the given NullableString and assigns it to the SendType field.
func (o *ModelCreditNoteUpdate) SetSendType(v string) {
	o.SendType.Set(&v)
}
// SetSendTypeNil sets the value for SendType to be an explicit nil
func (o *ModelCreditNoteUpdate) SetSendTypeNil() {
	o.SendType.Set(nil)
}

// UnsetSendType ensures that no value is present for SendType, not even an explicit nil
func (o *ModelCreditNoteUpdate) UnsetSendType() {
	o.SendType.Unset()
}

func (o ModelCreditNoteUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelCreditNoteUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	// skip: objectName is readOnly
	// skip: create is readOnly
	// skip: update is readOnly
	if o.CreditNoteNumber.IsSet() {
		toSerialize["creditNoteNumber"] = o.CreditNoteNumber.Get()
	}
	if o.Contact.IsSet() {
		toSerialize["contact"] = o.Contact.Get()
	}
	if o.CreditNoteDate.IsSet() {
		toSerialize["creditNoteDate"] = o.CreditNoteDate.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
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
	if o.AddressCountry.IsSet() {
		toSerialize["addressCountry"] = o.AddressCountry.Get()
	}
	if !IsNil(o.CreateUser) {
		toSerialize["createUser"] = o.CreateUser
	}
	if !IsNil(o.SevClient) {
		toSerialize["sevClient"] = o.SevClient
	}
	if o.DeliveryTerms.IsSet() {
		toSerialize["deliveryTerms"] = o.DeliveryTerms.Get()
	}
	if !IsNil(o.DeliveryDate) {
		toSerialize["deliveryDate"] = o.DeliveryDate
	}
	if o.PaymentTerms.IsSet() {
		toSerialize["paymentTerms"] = o.PaymentTerms.Get()
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	if o.SmallSettlement.IsSet() {
		toSerialize["smallSettlement"] = o.SmallSettlement.Get()
	}
	if o.ContactPerson.IsSet() {
		toSerialize["contactPerson"] = o.ContactPerson.Get()
	}
	if o.TaxRate.IsSet() {
		toSerialize["taxRate"] = o.TaxRate.Get()
	}
	if o.TaxSet.IsSet() {
		toSerialize["taxSet"] = o.TaxSet.Get()
	}
	if o.TaxText.IsSet() {
		toSerialize["taxText"] = o.TaxText.Get()
	}
	if o.TaxType.IsSet() {
		toSerialize["taxType"] = o.TaxType.Get()
	}
	if o.CreditNoteType.IsSet() {
		toSerialize["creditNoteType"] = o.CreditNoteType.Get()
	}
	if o.SendDate.IsSet() {
		toSerialize["sendDate"] = o.SendDate.Get()
	}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	// skip: sumNet is readOnly
	// skip: sumTax is readOnly
	// skip: sumGross is readOnly
	// skip: sumDiscounts is readOnly
	// skip: sumNetForeignCurrency is readOnly
	// skip: sumTaxForeignCurrency is readOnly
	// skip: sumGrossForeignCurrency is readOnly
	// skip: sumDiscountsForeignCurrency is readOnly
	if o.CustomerInternalNote.IsSet() {
		toSerialize["customerInternalNote"] = o.CustomerInternalNote.Get()
	}
	if !IsNil(o.ShowNet) {
		toSerialize["showNet"] = o.ShowNet
	}
	if o.SendType.IsSet() {
		toSerialize["sendType"] = o.SendType.Get()
	}
	return toSerialize, nil
}

type NullableModelCreditNoteUpdate struct {
	value *ModelCreditNoteUpdate
	isSet bool
}

func (v NullableModelCreditNoteUpdate) Get() *ModelCreditNoteUpdate {
	return v.value
}

func (v *NullableModelCreditNoteUpdate) Set(val *ModelCreditNoteUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableModelCreditNoteUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableModelCreditNoteUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelCreditNoteUpdate(val *ModelCreditNoteUpdate) *NullableModelCreditNoteUpdate {
	return &NullableModelCreditNoteUpdate{value: val, isSet: true}
}

func (v NullableModelCreditNoteUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelCreditNoteUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


