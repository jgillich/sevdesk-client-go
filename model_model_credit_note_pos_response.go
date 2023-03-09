/*
sevDesk API

<b>Contact:</b> To contact our support click <a href='https://landing.sevdesk.de/service-support-center-technik'>here</a><br><br>   # General information  Welcome to our API!<br>  sevDesk offers you the possibility of retrieving data using an interface, namely the sevDesk API, and making changes without having to use the web UI. The sevDesk interface is a REST-Full API. All sevDesk data and functions that are used in the web UI can also be controlled through the API.   # Cross-Origin Resource Sharing  This API features Cross-Origin Resource Sharing (CORS).<br>  It enables cross-domain communication from the browser.<br>  All responses have a wildcard same-origin which makes them completely public and accessible to everyone, including any code on any site.    # Embedding resources  When retrieving resources by using this API, you might encounter nested resources in the resources you requested.<br>  For example, an invoice always contains a contact, of which you can see the ID and the object name.<br>  This API gives you the possibility to embed these resources completely into the resources you originally requested.<br>  Taking our invoice example, this would mean, that you would not only see the ID and object name of a contact, but rather the complete contact resource.    To embed resources, all you need to do is to add the query parameter 'embed' to your GET request.<br>  As values, you can provide the name of the nested resource.<br>  Multiple nested resources are also possible by providing multiple names, separated by a comma.    # Authentication and Authorization  The sevDesk API uses a token authentication to authorize calls. For this purpose every sevDesk administrator has one API token, which is a <b>hexadecimal string</b> containing <b>32 characters</b>. The following clip shows where you can find the API token if this is your first time with our API.<br><br> <video src='OpenAPI/img/findingTheApiToken.mp4' controls width='640' height='360'> Ihr Browser kann dieses Video nicht wiedergeben.<br/> Dieses Video zeigt wie sie Ihr sevDesk API Token finden. </video> <br> The token will be needed in every request that you want to send and needs to be attached to the request url as a <b>Query Parameter</b><br> or provided as a value of an <b>Authorization Header</b>.<br> For security reasons, we suggest putting the API Token in the Authorization Header and not in the query string.<br> However, in the request examples in this documentation, we will keep it in the query string, as it is easier for you to copy them and try them yourself.<br> The following url is an example that shows where your token needs to be placed as a query parameter.<br> In this case, we used some random API token. <ul> <li><span>ht</span>tps://my.sevdesk.de/api/v1/Contact?token=<span style='color:red'>b7794de0085f5cd00560f160f290af38</span></li> </ul> The next example shows the token in the Authorization Header. <ul> <li>\"Authorization\" :<span style='color:red'>\"b7794de0085f5cd00560f160f290af38\"</span></li> </ul> The api tokens have an infinite lifetime and, in other words, exist as long as the sevDesk user exists.<br> For this reason, the user should <b>NEVER</b> be deleted.<br> If really necessary, it is advisable to save the api token as we will <b>NOT</b> be able to retrieve it afterwards!<br> It is also possible to generate a new API token, for example, if you want to prevent the usage of your sevDesk account by other people who got your current API token.<br> To achieve this, you just need to click on the \"generate new\" symbol to the right of your token and confirm it with your password.  # API News  To never miss API news and updates again, subscribe to our <b>free API newsletter</b> and get all relevant information to keep your sevDesk software running smoothly. To subscribe, simply click <a href = 'https://landing.sevdesk.de/api-newsletter'><b>here</b></a> and confirm the email address to which we may send all updates relevant to you.  # API Requests  In our case, REST API requests need to be build by combining the following components. <table> <tr> <th><b>Component</b></th> <th><b>Description</b></th> </tr> <tr> <td>HTTP-Methods</td> <td> <ul> <li>GET (retrieve a resource)</li> <li>POST (create a resource)</li> <li>PUT (update a resource)</li> <li>DELETE (delete a resource)</li> </ul> </td> </tr> <tr> <td>URL of the API</td> <td><span style='color: #2aa198'>ht</span><span style='color: #2aa198'>tps://my.sevdesk.de/api/v1</span></td> </tr> <tr> <td>URI of the resource</td> <td>The resource to query.<br>For example contacts in sevDesk:<br><br> <span style='color:red'>/Contact</span><br><br> Which will result in the following complete URL:<br><br> <span style='color: #2aa198'>ht</span><span style='color: #2aa198'>tps://my.sevdesk.de/api/v1</span><span style='color:red'>/Contact</span> </td> </tr> <tr> <td>Query parameters</td> <td>Are attached by using the connectives <b>?</b> and <b>&</b> in the URL.<br></td> </tr> <tr> <td>Request headers</td> <td>Typical request headers are for example:<br> <div> <br> <ul> <li>Content-type</li> <li>Authorization</li> <li>Accept-Encoding</li> <li>User-Agent</li> <li>...</li> </ul> </div> </td> </tr> <tr> <td>Request body</td> <td>Mostly required in POST and PUT requests.<br> Often the request body contains json, in our case, it also accepts url-encoded data. </td> </tr> </table><br> <span style='color:red'>Note</span>: please pass a meaningful entry at the header \"User-Agent\". If the \"User-Agent\" is set meaningfully, we can offer better support in case of queries from customers.<br> An example how such a \"User-Agent\" can look like: \"Integration-name by abc\". <br><br> This is a sample request for retrieving existing contacts in sevDesk using curl:<br><br> <img src='OpenAPI/img/GETRequest.PNG' alt='Get Request' height='150'><br><br> As you can see, the request contains all the components mentioned above.<br> It's HTTP method is GET, it has a correct endpoint (<span style='color: #2aa198'>ht</span><span style='color: #2aa198'>tps://my.sevdesk.de/api/v1</span><span style='color:red'>/Contact</span>), query parameters like <b>token</b> and additional <b>header</b> information!<br><br> <b>Query Parameters</b><br><br> As you might have seen in the sample request above, there are several other parameters besides \"token\", located in the url.<br> Those are mostly optional but prove to be very useful for a lot of requests as they can limit, extend, sort or filter the data you will get as a response.<br><br> These are the three most used query parameter for the sevDesk API. <table> <tr> <th><b>Parameter</b></th> <th><b>Description</b></th> </tr> <tr> <td>limit</td> <td>Limits the number of entries that are returned.<br> Most useful in GET requests which will most likely deliver big sets of data like country or currency lists.<br> In this case, you can bypass the default limitation on returned entries by providing a high number. </td> </tr> <tr> <td>offset</td> <td>Specifies a certain offset for the data that will be returned.<br> As an example, you can specify \"offset=2\" if you want all entries except for the first two.</td> </tr> <tr> <td>embed</td> <td>Will extend some of the returned data.<br> A brief example can be found below.</td> </tr> </table> This is an example for the usage of the embed parameter.<br> The following first request will return all company contact entries in sevDesk up to a limit of 100 without any additional information and no offset.<br><br> <img src='OpenAPI/img/ContactQueryWithoutEmbed.PNG' width='900' height='850'><br><br> Now have a look at the category attribute located in the response.<br> Naturally, it just contains the id and the object name of the object but no further information.<br> We will now use the parameter embed with the value \"category\".<br><br> <img src='OpenAPI/img/ContactQueryWithEmbed.PNG' width='900' height='850'><br><br> As you can see, the category object is now extended and shows all the attributes and their corresponding values.<br><br> There are lot of other query parameters that can be used to filter the returned data for objects that match a certain pattern, however, those will not be mentioned here and instead can be found at the detail documentation of the most used API endpoints like contact, invoice or voucher.<br><br> <b>Request Headers</b><br><br> The HTTP request (response) headers allow the client as well as the server to pass additional information with the request.<br> They transfer the parameters and arguments which are important for transmitting data over HTTP.<br> Three headers which are useful / necessary when using the sevDesk API are \"Authorization, \"Accept\" and \"Content-type\".<br> Underneath is a brief description of why and how they should be used.<br><br> <b>Authorization</b><br><br> Can be used if you want to provide your API token in the header instead of having it in the url. <ul> <li>Authorization:<span style='color:red'>yourApiToken</span></li> </ul> <b>Accept</b><br><br> Specifies the format of the response.<br> Required for operations with a response body. <ul> <li>Accept:application/<span style='color:red'>format</span> </li> </ul> In our case, <code><span style='color:red'>format</span></code> could be replaced with <code>json</code> or <code>xml</code><br><br> <b>Content-type</b><br><br> Specifies which format is used in the request.<br> Is required for operations with a request body. <ul> <li>Content-type:application/<span style='color:red'>format</span></li> </ul> In our case,<code><span style='color:red'>format</span></code>could be replaced with <code>json</code> or <code>x-www-form-urlencoded</code> <br><br><b>API Responses</b><br><br> HTTP status codes<br> When calling the sevDesk API it is very likely that you will get a HTTP status code in the response.<br> Some API calls will also return JSON response bodies which will contain information about the resource.<br> Each status code which is returned will either be a <b>success</b> code or an <b>error</b> code.<br><br> Success codes <table> <tr> <th><b>Status code</b></th> <th><b>Description</b></th> </tr> <tr> <td><code>200 OK</code></td> <td>The request was successful</td> </tr> <tr> <td><code>201 Created</code></td> <td>Most likely to be found in the response of a <b>POST</b> request.<br> This code indicates that the desired resource was successfully created.</td> </tr> </table> <br>Error codes <table> <tr> <th><b>Status code</b></th> <th><b>Description</b></th> </tr> <tr> <td><code>400 Bad request</code></td> <td>The request you sent is most likely syntactically incorrect.<br> You should check if the parameters in the request body or the url are correct.</td> </tr> <tr> <td><code>401 Unauthorized</code></td> <td>The authentication failed.<br> Most likely caused by a missing or wrong API token.</td> </tr> <tr> <td><code>403 Forbidden</code></td> <td>You do not have the permission the access the resource which is requested.</td> </tr> <tr> <td><code>404 Not found</code></td> <td>The resource you specified does not exist.</td> </tr> <tr> <td><code>500 Internal server error</code></td> <td>An internal server error has occurred.<br> Normally this means that something went wrong on our side.<br> However, sometimes this error will appear if we missed to catch an error which is normally a 400 status code! </td> </tr> </table>  # Your First Request  After reading the introduction to our API, you should now be able to make your first call.<br> For testing our API, we would always recommend to create a trial account for sevDesk to prevent unwanted changes to your main account.<br> A trial account will be in the highest tariff (materials management), so every sevDesk function can be tested! <br><br>To start testing we would recommend one of the following tools: <ul> <li><a href='https://www.getpostman.com/'>Postman</a></li> <li><a href='https://curl.haxx.se/download.html'>cURL</a></li> </ul> This example will illustrate your first request, which is creating a new Contact in sevDesk.<br> <ol> <li>Download <a href='https://www.getpostman.com/'><b>Postman</b></a> for your desired system and start the application</li> <li>Enter <span><b>ht</span>tps://my.sevdesk.de/api/v1/Contact</b> as the url</li> <li>Use the connective <b>?</b> to append <b>token=</b> to the end of the url, or create an authorization header. Insert your API token as the value</li> <li>For this test, select <b>POST</b> as the HTTP method</li> <li>Go to <b>Headers</b> and enter the key-value pair <b>Content-type</b> + <b>application/x-www-form-urlencoded</b><br> As an alternative, you can just go to <b>Body</b> and select <b>x-www-form-urlencoded</b></li> <li>Now go to <b>Body</b> (if you are not there yet) and enter the key-value pairs as shown in the following picture<br><br> <img src='OpenAPI/img/FirstRequestPostman.PNG' width='900'><br><br></li> <li>Click on <b>Send</b>. Your response should now look like this:<br><br> <img src='OpenAPI/img/FirstRequestResponse.PNG' width='900'></li> </ol> As you can see, a successful response in this case returns a JSON-formatted response body containing the contact you just created.<br> For keeping it simple, this was only a minimal example of creating a contact.<br> There are however numerous combinations of parameters that you can provide which add information to your contact.

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sevdesk

import (
	"encoding/json"
)

// checks if the ModelCreditNotePosResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelCreditNotePosResponse{}

// ModelCreditNotePosResponse creditNote position model
type ModelCreditNotePosResponse struct {
	// The creditNote position id
	Id *string `json:"id,omitempty"`
	// The creditNote position object name
	ObjectName *string `json:"objectName,omitempty"`
	// Date of creditNote position creation
	Create *string `json:"create,omitempty"`
	// Date of last creditNote position update
	Update *string `json:"update,omitempty"`
	CreditNote ModelCreditNotePosResponseCreditNote `json:"creditNote"`
	Part *ModelCreditNotePosResponsePart `json:"part,omitempty"`
	// Quantity of the article/part
	Quantity string `json:"quantity"`
	// Price of the article/part. Is either gross or net, depending on the sevDesk account setting.
	Price NullableString `json:"price,omitempty"`
	// Net price of the part
	PriceNet NullableString `json:"priceNet,omitempty"`
	// Tax on the price of the part
	PriceTax NullableString `json:"priceTax,omitempty"`
	// Gross price of the part
	PriceGross NullableString `json:"priceGross,omitempty"`
	// Name of the article/part.
	Name NullableString `json:"name,omitempty"`
	Unity ModelCreditNotePosResponseUnity `json:"unity"`
	SevClient *ModelCreditNotePosResponseSevClient `json:"sevClient,omitempty"`
	// Position number of your position. Can be used to creditNote multiple positions.
	PositionNumber NullableString `json:"positionNumber,omitempty"`
	// A text describing your position.
	Text NullableString `json:"text,omitempty"`
	// An optional discount of the position.
	Discount NullableString `json:"discount,omitempty"`
	// Defines if the position is optional.
	Optional NullableBool `json:"optional,omitempty"`
	// Tax rate of the position.
	TaxRate string `json:"taxRate"`
	// Discount sum of the position
	SumDiscount NullableString `json:"sumDiscount,omitempty"`
}

// NewModelCreditNotePosResponse instantiates a new ModelCreditNotePosResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelCreditNotePosResponse(creditNote ModelCreditNotePosResponseCreditNote, quantity string, unity ModelCreditNotePosResponseUnity, taxRate string) *ModelCreditNotePosResponse {
	this := ModelCreditNotePosResponse{}
	this.CreditNote = creditNote
	this.Quantity = quantity
	this.Unity = unity
	this.TaxRate = taxRate
	return &this
}

// NewModelCreditNotePosResponseWithDefaults instantiates a new ModelCreditNotePosResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelCreditNotePosResponseWithDefaults() *ModelCreditNotePosResponse {
	this := ModelCreditNotePosResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelCreditNotePosResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCreditNotePosResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelCreditNotePosResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModelCreditNotePosResponse) SetId(v string) {
	o.Id = &v
}

// GetObjectName returns the ObjectName field value if set, zero value otherwise.
func (o *ModelCreditNotePosResponse) GetObjectName() string {
	if o == nil || IsNil(o.ObjectName) {
		var ret string
		return ret
	}
	return *o.ObjectName
}

// GetObjectNameOk returns a tuple with the ObjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCreditNotePosResponse) GetObjectNameOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectName) {
		return nil, false
	}
	return o.ObjectName, true
}

// HasObjectName returns a boolean if a field has been set.
func (o *ModelCreditNotePosResponse) HasObjectName() bool {
	if o != nil && !IsNil(o.ObjectName) {
		return true
	}

	return false
}

// SetObjectName gets a reference to the given string and assigns it to the ObjectName field.
func (o *ModelCreditNotePosResponse) SetObjectName(v string) {
	o.ObjectName = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *ModelCreditNotePosResponse) GetCreate() string {
	if o == nil || IsNil(o.Create) {
		var ret string
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCreditNotePosResponse) GetCreateOk() (*string, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *ModelCreditNotePosResponse) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given string and assigns it to the Create field.
func (o *ModelCreditNotePosResponse) SetCreate(v string) {
	o.Create = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *ModelCreditNotePosResponse) GetUpdate() string {
	if o == nil || IsNil(o.Update) {
		var ret string
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCreditNotePosResponse) GetUpdateOk() (*string, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *ModelCreditNotePosResponse) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given string and assigns it to the Update field.
func (o *ModelCreditNotePosResponse) SetUpdate(v string) {
	o.Update = &v
}

// GetCreditNote returns the CreditNote field value
func (o *ModelCreditNotePosResponse) GetCreditNote() ModelCreditNotePosResponseCreditNote {
	if o == nil {
		var ret ModelCreditNotePosResponseCreditNote
		return ret
	}

	return o.CreditNote
}

// GetCreditNoteOk returns a tuple with the CreditNote field value
// and a boolean to check if the value has been set.
func (o *ModelCreditNotePosResponse) GetCreditNoteOk() (*ModelCreditNotePosResponseCreditNote, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreditNote, true
}

// SetCreditNote sets field value
func (o *ModelCreditNotePosResponse) SetCreditNote(v ModelCreditNotePosResponseCreditNote) {
	o.CreditNote = v
}

// GetPart returns the Part field value if set, zero value otherwise.
func (o *ModelCreditNotePosResponse) GetPart() ModelCreditNotePosResponsePart {
	if o == nil || IsNil(o.Part) {
		var ret ModelCreditNotePosResponsePart
		return ret
	}
	return *o.Part
}

// GetPartOk returns a tuple with the Part field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCreditNotePosResponse) GetPartOk() (*ModelCreditNotePosResponsePart, bool) {
	if o == nil || IsNil(o.Part) {
		return nil, false
	}
	return o.Part, true
}

// HasPart returns a boolean if a field has been set.
func (o *ModelCreditNotePosResponse) HasPart() bool {
	if o != nil && !IsNil(o.Part) {
		return true
	}

	return false
}

// SetPart gets a reference to the given ModelCreditNotePosResponsePart and assigns it to the Part field.
func (o *ModelCreditNotePosResponse) SetPart(v ModelCreditNotePosResponsePart) {
	o.Part = &v
}

// GetQuantity returns the Quantity field value
func (o *ModelCreditNotePosResponse) GetQuantity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *ModelCreditNotePosResponse) GetQuantityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *ModelCreditNotePosResponse) SetQuantity(v string) {
	o.Quantity = v
}

// GetPrice returns the Price field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCreditNotePosResponse) GetPrice() string {
	if o == nil || IsNil(o.Price.Get()) {
		var ret string
		return ret
	}
	return *o.Price.Get()
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCreditNotePosResponse) GetPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Price.Get(), o.Price.IsSet()
}

// HasPrice returns a boolean if a field has been set.
func (o *ModelCreditNotePosResponse) HasPrice() bool {
	if o != nil && o.Price.IsSet() {
		return true
	}

	return false
}

// SetPrice gets a reference to the given NullableString and assigns it to the Price field.
func (o *ModelCreditNotePosResponse) SetPrice(v string) {
	o.Price.Set(&v)
}
// SetPriceNil sets the value for Price to be an explicit nil
func (o *ModelCreditNotePosResponse) SetPriceNil() {
	o.Price.Set(nil)
}

// UnsetPrice ensures that no value is present for Price, not even an explicit nil
func (o *ModelCreditNotePosResponse) UnsetPrice() {
	o.Price.Unset()
}

// GetPriceNet returns the PriceNet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCreditNotePosResponse) GetPriceNet() string {
	if o == nil || IsNil(o.PriceNet.Get()) {
		var ret string
		return ret
	}
	return *o.PriceNet.Get()
}

// GetPriceNetOk returns a tuple with the PriceNet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCreditNotePosResponse) GetPriceNetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PriceNet.Get(), o.PriceNet.IsSet()
}

// HasPriceNet returns a boolean if a field has been set.
func (o *ModelCreditNotePosResponse) HasPriceNet() bool {
	if o != nil && o.PriceNet.IsSet() {
		return true
	}

	return false
}

// SetPriceNet gets a reference to the given NullableString and assigns it to the PriceNet field.
func (o *ModelCreditNotePosResponse) SetPriceNet(v string) {
	o.PriceNet.Set(&v)
}
// SetPriceNetNil sets the value for PriceNet to be an explicit nil
func (o *ModelCreditNotePosResponse) SetPriceNetNil() {
	o.PriceNet.Set(nil)
}

// UnsetPriceNet ensures that no value is present for PriceNet, not even an explicit nil
func (o *ModelCreditNotePosResponse) UnsetPriceNet() {
	o.PriceNet.Unset()
}

// GetPriceTax returns the PriceTax field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCreditNotePosResponse) GetPriceTax() string {
	if o == nil || IsNil(o.PriceTax.Get()) {
		var ret string
		return ret
	}
	return *o.PriceTax.Get()
}

// GetPriceTaxOk returns a tuple with the PriceTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCreditNotePosResponse) GetPriceTaxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PriceTax.Get(), o.PriceTax.IsSet()
}

// HasPriceTax returns a boolean if a field has been set.
func (o *ModelCreditNotePosResponse) HasPriceTax() bool {
	if o != nil && o.PriceTax.IsSet() {
		return true
	}

	return false
}

// SetPriceTax gets a reference to the given NullableString and assigns it to the PriceTax field.
func (o *ModelCreditNotePosResponse) SetPriceTax(v string) {
	o.PriceTax.Set(&v)
}
// SetPriceTaxNil sets the value for PriceTax to be an explicit nil
func (o *ModelCreditNotePosResponse) SetPriceTaxNil() {
	o.PriceTax.Set(nil)
}

// UnsetPriceTax ensures that no value is present for PriceTax, not even an explicit nil
func (o *ModelCreditNotePosResponse) UnsetPriceTax() {
	o.PriceTax.Unset()
}

// GetPriceGross returns the PriceGross field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCreditNotePosResponse) GetPriceGross() string {
	if o == nil || IsNil(o.PriceGross.Get()) {
		var ret string
		return ret
	}
	return *o.PriceGross.Get()
}

// GetPriceGrossOk returns a tuple with the PriceGross field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCreditNotePosResponse) GetPriceGrossOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PriceGross.Get(), o.PriceGross.IsSet()
}

// HasPriceGross returns a boolean if a field has been set.
func (o *ModelCreditNotePosResponse) HasPriceGross() bool {
	if o != nil && o.PriceGross.IsSet() {
		return true
	}

	return false
}

// SetPriceGross gets a reference to the given NullableString and assigns it to the PriceGross field.
func (o *ModelCreditNotePosResponse) SetPriceGross(v string) {
	o.PriceGross.Set(&v)
}
// SetPriceGrossNil sets the value for PriceGross to be an explicit nil
func (o *ModelCreditNotePosResponse) SetPriceGrossNil() {
	o.PriceGross.Set(nil)
}

// UnsetPriceGross ensures that no value is present for PriceGross, not even an explicit nil
func (o *ModelCreditNotePosResponse) UnsetPriceGross() {
	o.PriceGross.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCreditNotePosResponse) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCreditNotePosResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ModelCreditNotePosResponse) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ModelCreditNotePosResponse) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ModelCreditNotePosResponse) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ModelCreditNotePosResponse) UnsetName() {
	o.Name.Unset()
}

// GetUnity returns the Unity field value
func (o *ModelCreditNotePosResponse) GetUnity() ModelCreditNotePosResponseUnity {
	if o == nil {
		var ret ModelCreditNotePosResponseUnity
		return ret
	}

	return o.Unity
}

// GetUnityOk returns a tuple with the Unity field value
// and a boolean to check if the value has been set.
func (o *ModelCreditNotePosResponse) GetUnityOk() (*ModelCreditNotePosResponseUnity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unity, true
}

// SetUnity sets field value
func (o *ModelCreditNotePosResponse) SetUnity(v ModelCreditNotePosResponseUnity) {
	o.Unity = v
}

// GetSevClient returns the SevClient field value if set, zero value otherwise.
func (o *ModelCreditNotePosResponse) GetSevClient() ModelCreditNotePosResponseSevClient {
	if o == nil || IsNil(o.SevClient) {
		var ret ModelCreditNotePosResponseSevClient
		return ret
	}
	return *o.SevClient
}

// GetSevClientOk returns a tuple with the SevClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCreditNotePosResponse) GetSevClientOk() (*ModelCreditNotePosResponseSevClient, bool) {
	if o == nil || IsNil(o.SevClient) {
		return nil, false
	}
	return o.SevClient, true
}

// HasSevClient returns a boolean if a field has been set.
func (o *ModelCreditNotePosResponse) HasSevClient() bool {
	if o != nil && !IsNil(o.SevClient) {
		return true
	}

	return false
}

// SetSevClient gets a reference to the given ModelCreditNotePosResponseSevClient and assigns it to the SevClient field.
func (o *ModelCreditNotePosResponse) SetSevClient(v ModelCreditNotePosResponseSevClient) {
	o.SevClient = &v
}

// GetPositionNumber returns the PositionNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCreditNotePosResponse) GetPositionNumber() string {
	if o == nil || IsNil(o.PositionNumber.Get()) {
		var ret string
		return ret
	}
	return *o.PositionNumber.Get()
}

// GetPositionNumberOk returns a tuple with the PositionNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCreditNotePosResponse) GetPositionNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PositionNumber.Get(), o.PositionNumber.IsSet()
}

// HasPositionNumber returns a boolean if a field has been set.
func (o *ModelCreditNotePosResponse) HasPositionNumber() bool {
	if o != nil && o.PositionNumber.IsSet() {
		return true
	}

	return false
}

// SetPositionNumber gets a reference to the given NullableString and assigns it to the PositionNumber field.
func (o *ModelCreditNotePosResponse) SetPositionNumber(v string) {
	o.PositionNumber.Set(&v)
}
// SetPositionNumberNil sets the value for PositionNumber to be an explicit nil
func (o *ModelCreditNotePosResponse) SetPositionNumberNil() {
	o.PositionNumber.Set(nil)
}

// UnsetPositionNumber ensures that no value is present for PositionNumber, not even an explicit nil
func (o *ModelCreditNotePosResponse) UnsetPositionNumber() {
	o.PositionNumber.Unset()
}

// GetText returns the Text field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCreditNotePosResponse) GetText() string {
	if o == nil || IsNil(o.Text.Get()) {
		var ret string
		return ret
	}
	return *o.Text.Get()
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCreditNotePosResponse) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Text.Get(), o.Text.IsSet()
}

// HasText returns a boolean if a field has been set.
func (o *ModelCreditNotePosResponse) HasText() bool {
	if o != nil && o.Text.IsSet() {
		return true
	}

	return false
}

// SetText gets a reference to the given NullableString and assigns it to the Text field.
func (o *ModelCreditNotePosResponse) SetText(v string) {
	o.Text.Set(&v)
}
// SetTextNil sets the value for Text to be an explicit nil
func (o *ModelCreditNotePosResponse) SetTextNil() {
	o.Text.Set(nil)
}

// UnsetText ensures that no value is present for Text, not even an explicit nil
func (o *ModelCreditNotePosResponse) UnsetText() {
	o.Text.Unset()
}

// GetDiscount returns the Discount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCreditNotePosResponse) GetDiscount() string {
	if o == nil || IsNil(o.Discount.Get()) {
		var ret string
		return ret
	}
	return *o.Discount.Get()
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCreditNotePosResponse) GetDiscountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Discount.Get(), o.Discount.IsSet()
}

// HasDiscount returns a boolean if a field has been set.
func (o *ModelCreditNotePosResponse) HasDiscount() bool {
	if o != nil && o.Discount.IsSet() {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given NullableString and assigns it to the Discount field.
func (o *ModelCreditNotePosResponse) SetDiscount(v string) {
	o.Discount.Set(&v)
}
// SetDiscountNil sets the value for Discount to be an explicit nil
func (o *ModelCreditNotePosResponse) SetDiscountNil() {
	o.Discount.Set(nil)
}

// UnsetDiscount ensures that no value is present for Discount, not even an explicit nil
func (o *ModelCreditNotePosResponse) UnsetDiscount() {
	o.Discount.Unset()
}

// GetOptional returns the Optional field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCreditNotePosResponse) GetOptional() bool {
	if o == nil || IsNil(o.Optional.Get()) {
		var ret bool
		return ret
	}
	return *o.Optional.Get()
}

// GetOptionalOk returns a tuple with the Optional field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCreditNotePosResponse) GetOptionalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Optional.Get(), o.Optional.IsSet()
}

// HasOptional returns a boolean if a field has been set.
func (o *ModelCreditNotePosResponse) HasOptional() bool {
	if o != nil && o.Optional.IsSet() {
		return true
	}

	return false
}

// SetOptional gets a reference to the given NullableBool and assigns it to the Optional field.
func (o *ModelCreditNotePosResponse) SetOptional(v bool) {
	o.Optional.Set(&v)
}
// SetOptionalNil sets the value for Optional to be an explicit nil
func (o *ModelCreditNotePosResponse) SetOptionalNil() {
	o.Optional.Set(nil)
}

// UnsetOptional ensures that no value is present for Optional, not even an explicit nil
func (o *ModelCreditNotePosResponse) UnsetOptional() {
	o.Optional.Unset()
}

// GetTaxRate returns the TaxRate field value
func (o *ModelCreditNotePosResponse) GetTaxRate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TaxRate
}

// GetTaxRateOk returns a tuple with the TaxRate field value
// and a boolean to check if the value has been set.
func (o *ModelCreditNotePosResponse) GetTaxRateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaxRate, true
}

// SetTaxRate sets field value
func (o *ModelCreditNotePosResponse) SetTaxRate(v string) {
	o.TaxRate = v
}

// GetSumDiscount returns the SumDiscount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCreditNotePosResponse) GetSumDiscount() string {
	if o == nil || IsNil(o.SumDiscount.Get()) {
		var ret string
		return ret
	}
	return *o.SumDiscount.Get()
}

// GetSumDiscountOk returns a tuple with the SumDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCreditNotePosResponse) GetSumDiscountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SumDiscount.Get(), o.SumDiscount.IsSet()
}

// HasSumDiscount returns a boolean if a field has been set.
func (o *ModelCreditNotePosResponse) HasSumDiscount() bool {
	if o != nil && o.SumDiscount.IsSet() {
		return true
	}

	return false
}

// SetSumDiscount gets a reference to the given NullableString and assigns it to the SumDiscount field.
func (o *ModelCreditNotePosResponse) SetSumDiscount(v string) {
	o.SumDiscount.Set(&v)
}
// SetSumDiscountNil sets the value for SumDiscount to be an explicit nil
func (o *ModelCreditNotePosResponse) SetSumDiscountNil() {
	o.SumDiscount.Set(nil)
}

// UnsetSumDiscount ensures that no value is present for SumDiscount, not even an explicit nil
func (o *ModelCreditNotePosResponse) UnsetSumDiscount() {
	o.SumDiscount.Unset()
}

func (o ModelCreditNotePosResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelCreditNotePosResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	// skip: objectName is readOnly
	// skip: create is readOnly
	// skip: update is readOnly
	toSerialize["creditNote"] = o.CreditNote
	if !IsNil(o.Part) {
		toSerialize["part"] = o.Part
	}
	toSerialize["quantity"] = o.Quantity
	if o.Price.IsSet() {
		toSerialize["price"] = o.Price.Get()
	}
	if o.PriceNet.IsSet() {
		toSerialize["priceNet"] = o.PriceNet.Get()
	}
	if o.PriceTax.IsSet() {
		toSerialize["priceTax"] = o.PriceTax.Get()
	}
	if o.PriceGross.IsSet() {
		toSerialize["priceGross"] = o.PriceGross.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	toSerialize["unity"] = o.Unity
	if !IsNil(o.SevClient) {
		toSerialize["sevClient"] = o.SevClient
	}
	if o.PositionNumber.IsSet() {
		toSerialize["positionNumber"] = o.PositionNumber.Get()
	}
	if o.Text.IsSet() {
		toSerialize["text"] = o.Text.Get()
	}
	if o.Discount.IsSet() {
		toSerialize["discount"] = o.Discount.Get()
	}
	if o.Optional.IsSet() {
		toSerialize["optional"] = o.Optional.Get()
	}
	toSerialize["taxRate"] = o.TaxRate
	if o.SumDiscount.IsSet() {
		toSerialize["sumDiscount"] = o.SumDiscount.Get()
	}
	return toSerialize, nil
}

type NullableModelCreditNotePosResponse struct {
	value *ModelCreditNotePosResponse
	isSet bool
}

func (v NullableModelCreditNotePosResponse) Get() *ModelCreditNotePosResponse {
	return v.value
}

func (v *NullableModelCreditNotePosResponse) Set(val *ModelCreditNotePosResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModelCreditNotePosResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModelCreditNotePosResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelCreditNotePosResponse(val *ModelCreditNotePosResponse) *NullableModelCreditNotePosResponse {
	return &NullableModelCreditNotePosResponse{value: val, isSet: true}
}

func (v NullableModelCreditNotePosResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelCreditNotePosResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


