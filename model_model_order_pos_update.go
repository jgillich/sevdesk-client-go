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

// checks if the ModelOrderPosUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelOrderPosUpdate{}

// ModelOrderPosUpdate Order position model
type ModelOrderPosUpdate struct {
	// The order position id
	Id *int32 `json:"id,omitempty"`
	// The order position object name
	ObjectName *string `json:"objectName,omitempty"`
	// Date of order position creation
	Create *time.Time `json:"create,omitempty"`
	// Date of last order position update
	Update *time.Time `json:"update,omitempty"`
	Order *ModelOrderPosOrder `json:"order,omitempty"`
	Part *ModelInvoicePosPart `json:"part,omitempty"`
	// Quantity of the article/part
	Quantity NullableFloat32 `json:"quantity,omitempty"`
	// Price of the article/part. Is either gross or net, depending on the sevDesk account setting.
	Price NullableFloat32 `json:"price,omitempty"`
	// Net price of the part
	PriceNet NullableFloat32 `json:"priceNet,omitempty"`
	// Tax on the price of the part
	PriceTax NullableFloat32 `json:"priceTax,omitempty"`
	// Gross price of the part
	PriceGross NullableFloat32 `json:"priceGross,omitempty"`
	// Name of the article/part.
	Name NullableFloat32 `json:"name,omitempty"`
	Unity *ModelCreditNotePosUnity `json:"unity,omitempty"`
	SevClient *ModelOrderPosSevClient `json:"sevClient,omitempty"`
	// Position number of your position. Can be used to order multiple positions.
	PositionNumber NullableInt32 `json:"positionNumber,omitempty"`
	// A text describing your position.
	Text NullableString `json:"text,omitempty"`
	// An optional discount of the position.
	Discount NullableFloat32 `json:"discount,omitempty"`
	// Defines if the position is optional.
	Optional NullableBool `json:"optional,omitempty"`
	// Tax rate of the position.
	TaxRate NullableFloat32 `json:"taxRate,omitempty"`
	// Discount sum of the position
	SumDiscount NullableFloat32 `json:"sumDiscount,omitempty"`
}

// NewModelOrderPosUpdate instantiates a new ModelOrderPosUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelOrderPosUpdate() *ModelOrderPosUpdate {
	this := ModelOrderPosUpdate{}
	return &this
}

// NewModelOrderPosUpdateWithDefaults instantiates a new ModelOrderPosUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelOrderPosUpdateWithDefaults() *ModelOrderPosUpdate {
	this := ModelOrderPosUpdate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelOrderPosUpdate) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelOrderPosUpdate) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelOrderPosUpdate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ModelOrderPosUpdate) SetId(v int32) {
	o.Id = &v
}

// GetObjectName returns the ObjectName field value if set, zero value otherwise.
func (o *ModelOrderPosUpdate) GetObjectName() string {
	if o == nil || IsNil(o.ObjectName) {
		var ret string
		return ret
	}
	return *o.ObjectName
}

// GetObjectNameOk returns a tuple with the ObjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelOrderPosUpdate) GetObjectNameOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectName) {
		return nil, false
	}
	return o.ObjectName, true
}

// HasObjectName returns a boolean if a field has been set.
func (o *ModelOrderPosUpdate) HasObjectName() bool {
	if o != nil && !IsNil(o.ObjectName) {
		return true
	}

	return false
}

// SetObjectName gets a reference to the given string and assigns it to the ObjectName field.
func (o *ModelOrderPosUpdate) SetObjectName(v string) {
	o.ObjectName = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *ModelOrderPosUpdate) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelOrderPosUpdate) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *ModelOrderPosUpdate) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *ModelOrderPosUpdate) SetCreate(v time.Time) {
	o.Create = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *ModelOrderPosUpdate) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelOrderPosUpdate) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *ModelOrderPosUpdate) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *ModelOrderPosUpdate) SetUpdate(v time.Time) {
	o.Update = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *ModelOrderPosUpdate) GetOrder() ModelOrderPosOrder {
	if o == nil || IsNil(o.Order) {
		var ret ModelOrderPosOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelOrderPosUpdate) GetOrderOk() (*ModelOrderPosOrder, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *ModelOrderPosUpdate) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given ModelOrderPosOrder and assigns it to the Order field.
func (o *ModelOrderPosUpdate) SetOrder(v ModelOrderPosOrder) {
	o.Order = &v
}

// GetPart returns the Part field value if set, zero value otherwise.
func (o *ModelOrderPosUpdate) GetPart() ModelInvoicePosPart {
	if o == nil || IsNil(o.Part) {
		var ret ModelInvoicePosPart
		return ret
	}
	return *o.Part
}

// GetPartOk returns a tuple with the Part field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelOrderPosUpdate) GetPartOk() (*ModelInvoicePosPart, bool) {
	if o == nil || IsNil(o.Part) {
		return nil, false
	}
	return o.Part, true
}

// HasPart returns a boolean if a field has been set.
func (o *ModelOrderPosUpdate) HasPart() bool {
	if o != nil && !IsNil(o.Part) {
		return true
	}

	return false
}

// SetPart gets a reference to the given ModelInvoicePosPart and assigns it to the Part field.
func (o *ModelOrderPosUpdate) SetPart(v ModelInvoicePosPart) {
	o.Part = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelOrderPosUpdate) GetQuantity() float32 {
	if o == nil || IsNil(o.Quantity.Get()) {
		var ret float32
		return ret
	}
	return *o.Quantity.Get()
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelOrderPosUpdate) GetQuantityOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Quantity.Get(), o.Quantity.IsSet()
}

// HasQuantity returns a boolean if a field has been set.
func (o *ModelOrderPosUpdate) HasQuantity() bool {
	if o != nil && o.Quantity.IsSet() {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given NullableFloat32 and assigns it to the Quantity field.
func (o *ModelOrderPosUpdate) SetQuantity(v float32) {
	o.Quantity.Set(&v)
}
// SetQuantityNil sets the value for Quantity to be an explicit nil
func (o *ModelOrderPosUpdate) SetQuantityNil() {
	o.Quantity.Set(nil)
}

// UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
func (o *ModelOrderPosUpdate) UnsetQuantity() {
	o.Quantity.Unset()
}

// GetPrice returns the Price field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelOrderPosUpdate) GetPrice() float32 {
	if o == nil || IsNil(o.Price.Get()) {
		var ret float32
		return ret
	}
	return *o.Price.Get()
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelOrderPosUpdate) GetPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Price.Get(), o.Price.IsSet()
}

// HasPrice returns a boolean if a field has been set.
func (o *ModelOrderPosUpdate) HasPrice() bool {
	if o != nil && o.Price.IsSet() {
		return true
	}

	return false
}

// SetPrice gets a reference to the given NullableFloat32 and assigns it to the Price field.
func (o *ModelOrderPosUpdate) SetPrice(v float32) {
	o.Price.Set(&v)
}
// SetPriceNil sets the value for Price to be an explicit nil
func (o *ModelOrderPosUpdate) SetPriceNil() {
	o.Price.Set(nil)
}

// UnsetPrice ensures that no value is present for Price, not even an explicit nil
func (o *ModelOrderPosUpdate) UnsetPrice() {
	o.Price.Unset()
}

// GetPriceNet returns the PriceNet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelOrderPosUpdate) GetPriceNet() float32 {
	if o == nil || IsNil(o.PriceNet.Get()) {
		var ret float32
		return ret
	}
	return *o.PriceNet.Get()
}

// GetPriceNetOk returns a tuple with the PriceNet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelOrderPosUpdate) GetPriceNetOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PriceNet.Get(), o.PriceNet.IsSet()
}

// HasPriceNet returns a boolean if a field has been set.
func (o *ModelOrderPosUpdate) HasPriceNet() bool {
	if o != nil && o.PriceNet.IsSet() {
		return true
	}

	return false
}

// SetPriceNet gets a reference to the given NullableFloat32 and assigns it to the PriceNet field.
func (o *ModelOrderPosUpdate) SetPriceNet(v float32) {
	o.PriceNet.Set(&v)
}
// SetPriceNetNil sets the value for PriceNet to be an explicit nil
func (o *ModelOrderPosUpdate) SetPriceNetNil() {
	o.PriceNet.Set(nil)
}

// UnsetPriceNet ensures that no value is present for PriceNet, not even an explicit nil
func (o *ModelOrderPosUpdate) UnsetPriceNet() {
	o.PriceNet.Unset()
}

// GetPriceTax returns the PriceTax field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelOrderPosUpdate) GetPriceTax() float32 {
	if o == nil || IsNil(o.PriceTax.Get()) {
		var ret float32
		return ret
	}
	return *o.PriceTax.Get()
}

// GetPriceTaxOk returns a tuple with the PriceTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelOrderPosUpdate) GetPriceTaxOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PriceTax.Get(), o.PriceTax.IsSet()
}

// HasPriceTax returns a boolean if a field has been set.
func (o *ModelOrderPosUpdate) HasPriceTax() bool {
	if o != nil && o.PriceTax.IsSet() {
		return true
	}

	return false
}

// SetPriceTax gets a reference to the given NullableFloat32 and assigns it to the PriceTax field.
func (o *ModelOrderPosUpdate) SetPriceTax(v float32) {
	o.PriceTax.Set(&v)
}
// SetPriceTaxNil sets the value for PriceTax to be an explicit nil
func (o *ModelOrderPosUpdate) SetPriceTaxNil() {
	o.PriceTax.Set(nil)
}

// UnsetPriceTax ensures that no value is present for PriceTax, not even an explicit nil
func (o *ModelOrderPosUpdate) UnsetPriceTax() {
	o.PriceTax.Unset()
}

// GetPriceGross returns the PriceGross field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelOrderPosUpdate) GetPriceGross() float32 {
	if o == nil || IsNil(o.PriceGross.Get()) {
		var ret float32
		return ret
	}
	return *o.PriceGross.Get()
}

// GetPriceGrossOk returns a tuple with the PriceGross field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelOrderPosUpdate) GetPriceGrossOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PriceGross.Get(), o.PriceGross.IsSet()
}

// HasPriceGross returns a boolean if a field has been set.
func (o *ModelOrderPosUpdate) HasPriceGross() bool {
	if o != nil && o.PriceGross.IsSet() {
		return true
	}

	return false
}

// SetPriceGross gets a reference to the given NullableFloat32 and assigns it to the PriceGross field.
func (o *ModelOrderPosUpdate) SetPriceGross(v float32) {
	o.PriceGross.Set(&v)
}
// SetPriceGrossNil sets the value for PriceGross to be an explicit nil
func (o *ModelOrderPosUpdate) SetPriceGrossNil() {
	o.PriceGross.Set(nil)
}

// UnsetPriceGross ensures that no value is present for PriceGross, not even an explicit nil
func (o *ModelOrderPosUpdate) UnsetPriceGross() {
	o.PriceGross.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelOrderPosUpdate) GetName() float32 {
	if o == nil || IsNil(o.Name.Get()) {
		var ret float32
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelOrderPosUpdate) GetNameOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ModelOrderPosUpdate) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableFloat32 and assigns it to the Name field.
func (o *ModelOrderPosUpdate) SetName(v float32) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ModelOrderPosUpdate) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ModelOrderPosUpdate) UnsetName() {
	o.Name.Unset()
}

// GetUnity returns the Unity field value if set, zero value otherwise.
func (o *ModelOrderPosUpdate) GetUnity() ModelCreditNotePosUnity {
	if o == nil || IsNil(o.Unity) {
		var ret ModelCreditNotePosUnity
		return ret
	}
	return *o.Unity
}

// GetUnityOk returns a tuple with the Unity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelOrderPosUpdate) GetUnityOk() (*ModelCreditNotePosUnity, bool) {
	if o == nil || IsNil(o.Unity) {
		return nil, false
	}
	return o.Unity, true
}

// HasUnity returns a boolean if a field has been set.
func (o *ModelOrderPosUpdate) HasUnity() bool {
	if o != nil && !IsNil(o.Unity) {
		return true
	}

	return false
}

// SetUnity gets a reference to the given ModelCreditNotePosUnity and assigns it to the Unity field.
func (o *ModelOrderPosUpdate) SetUnity(v ModelCreditNotePosUnity) {
	o.Unity = &v
}

// GetSevClient returns the SevClient field value if set, zero value otherwise.
func (o *ModelOrderPosUpdate) GetSevClient() ModelOrderPosSevClient {
	if o == nil || IsNil(o.SevClient) {
		var ret ModelOrderPosSevClient
		return ret
	}
	return *o.SevClient
}

// GetSevClientOk returns a tuple with the SevClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelOrderPosUpdate) GetSevClientOk() (*ModelOrderPosSevClient, bool) {
	if o == nil || IsNil(o.SevClient) {
		return nil, false
	}
	return o.SevClient, true
}

// HasSevClient returns a boolean if a field has been set.
func (o *ModelOrderPosUpdate) HasSevClient() bool {
	if o != nil && !IsNil(o.SevClient) {
		return true
	}

	return false
}

// SetSevClient gets a reference to the given ModelOrderPosSevClient and assigns it to the SevClient field.
func (o *ModelOrderPosUpdate) SetSevClient(v ModelOrderPosSevClient) {
	o.SevClient = &v
}

// GetPositionNumber returns the PositionNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelOrderPosUpdate) GetPositionNumber() int32 {
	if o == nil || IsNil(o.PositionNumber.Get()) {
		var ret int32
		return ret
	}
	return *o.PositionNumber.Get()
}

// GetPositionNumberOk returns a tuple with the PositionNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelOrderPosUpdate) GetPositionNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PositionNumber.Get(), o.PositionNumber.IsSet()
}

// HasPositionNumber returns a boolean if a field has been set.
func (o *ModelOrderPosUpdate) HasPositionNumber() bool {
	if o != nil && o.PositionNumber.IsSet() {
		return true
	}

	return false
}

// SetPositionNumber gets a reference to the given NullableInt32 and assigns it to the PositionNumber field.
func (o *ModelOrderPosUpdate) SetPositionNumber(v int32) {
	o.PositionNumber.Set(&v)
}
// SetPositionNumberNil sets the value for PositionNumber to be an explicit nil
func (o *ModelOrderPosUpdate) SetPositionNumberNil() {
	o.PositionNumber.Set(nil)
}

// UnsetPositionNumber ensures that no value is present for PositionNumber, not even an explicit nil
func (o *ModelOrderPosUpdate) UnsetPositionNumber() {
	o.PositionNumber.Unset()
}

// GetText returns the Text field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelOrderPosUpdate) GetText() string {
	if o == nil || IsNil(o.Text.Get()) {
		var ret string
		return ret
	}
	return *o.Text.Get()
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelOrderPosUpdate) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Text.Get(), o.Text.IsSet()
}

// HasText returns a boolean if a field has been set.
func (o *ModelOrderPosUpdate) HasText() bool {
	if o != nil && o.Text.IsSet() {
		return true
	}

	return false
}

// SetText gets a reference to the given NullableString and assigns it to the Text field.
func (o *ModelOrderPosUpdate) SetText(v string) {
	o.Text.Set(&v)
}
// SetTextNil sets the value for Text to be an explicit nil
func (o *ModelOrderPosUpdate) SetTextNil() {
	o.Text.Set(nil)
}

// UnsetText ensures that no value is present for Text, not even an explicit nil
func (o *ModelOrderPosUpdate) UnsetText() {
	o.Text.Unset()
}

// GetDiscount returns the Discount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelOrderPosUpdate) GetDiscount() float32 {
	if o == nil || IsNil(o.Discount.Get()) {
		var ret float32
		return ret
	}
	return *o.Discount.Get()
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelOrderPosUpdate) GetDiscountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Discount.Get(), o.Discount.IsSet()
}

// HasDiscount returns a boolean if a field has been set.
func (o *ModelOrderPosUpdate) HasDiscount() bool {
	if o != nil && o.Discount.IsSet() {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given NullableFloat32 and assigns it to the Discount field.
func (o *ModelOrderPosUpdate) SetDiscount(v float32) {
	o.Discount.Set(&v)
}
// SetDiscountNil sets the value for Discount to be an explicit nil
func (o *ModelOrderPosUpdate) SetDiscountNil() {
	o.Discount.Set(nil)
}

// UnsetDiscount ensures that no value is present for Discount, not even an explicit nil
func (o *ModelOrderPosUpdate) UnsetDiscount() {
	o.Discount.Unset()
}

// GetOptional returns the Optional field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelOrderPosUpdate) GetOptional() bool {
	if o == nil || IsNil(o.Optional.Get()) {
		var ret bool
		return ret
	}
	return *o.Optional.Get()
}

// GetOptionalOk returns a tuple with the Optional field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelOrderPosUpdate) GetOptionalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Optional.Get(), o.Optional.IsSet()
}

// HasOptional returns a boolean if a field has been set.
func (o *ModelOrderPosUpdate) HasOptional() bool {
	if o != nil && o.Optional.IsSet() {
		return true
	}

	return false
}

// SetOptional gets a reference to the given NullableBool and assigns it to the Optional field.
func (o *ModelOrderPosUpdate) SetOptional(v bool) {
	o.Optional.Set(&v)
}
// SetOptionalNil sets the value for Optional to be an explicit nil
func (o *ModelOrderPosUpdate) SetOptionalNil() {
	o.Optional.Set(nil)
}

// UnsetOptional ensures that no value is present for Optional, not even an explicit nil
func (o *ModelOrderPosUpdate) UnsetOptional() {
	o.Optional.Unset()
}

// GetTaxRate returns the TaxRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelOrderPosUpdate) GetTaxRate() float32 {
	if o == nil || IsNil(o.TaxRate.Get()) {
		var ret float32
		return ret
	}
	return *o.TaxRate.Get()
}

// GetTaxRateOk returns a tuple with the TaxRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelOrderPosUpdate) GetTaxRateOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxRate.Get(), o.TaxRate.IsSet()
}

// HasTaxRate returns a boolean if a field has been set.
func (o *ModelOrderPosUpdate) HasTaxRate() bool {
	if o != nil && o.TaxRate.IsSet() {
		return true
	}

	return false
}

// SetTaxRate gets a reference to the given NullableFloat32 and assigns it to the TaxRate field.
func (o *ModelOrderPosUpdate) SetTaxRate(v float32) {
	o.TaxRate.Set(&v)
}
// SetTaxRateNil sets the value for TaxRate to be an explicit nil
func (o *ModelOrderPosUpdate) SetTaxRateNil() {
	o.TaxRate.Set(nil)
}

// UnsetTaxRate ensures that no value is present for TaxRate, not even an explicit nil
func (o *ModelOrderPosUpdate) UnsetTaxRate() {
	o.TaxRate.Unset()
}

// GetSumDiscount returns the SumDiscount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelOrderPosUpdate) GetSumDiscount() float32 {
	if o == nil || IsNil(o.SumDiscount.Get()) {
		var ret float32
		return ret
	}
	return *o.SumDiscount.Get()
}

// GetSumDiscountOk returns a tuple with the SumDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelOrderPosUpdate) GetSumDiscountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SumDiscount.Get(), o.SumDiscount.IsSet()
}

// HasSumDiscount returns a boolean if a field has been set.
func (o *ModelOrderPosUpdate) HasSumDiscount() bool {
	if o != nil && o.SumDiscount.IsSet() {
		return true
	}

	return false
}

// SetSumDiscount gets a reference to the given NullableFloat32 and assigns it to the SumDiscount field.
func (o *ModelOrderPosUpdate) SetSumDiscount(v float32) {
	o.SumDiscount.Set(&v)
}
// SetSumDiscountNil sets the value for SumDiscount to be an explicit nil
func (o *ModelOrderPosUpdate) SetSumDiscountNil() {
	o.SumDiscount.Set(nil)
}

// UnsetSumDiscount ensures that no value is present for SumDiscount, not even an explicit nil
func (o *ModelOrderPosUpdate) UnsetSumDiscount() {
	o.SumDiscount.Unset()
}

func (o ModelOrderPosUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelOrderPosUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	// skip: objectName is readOnly
	// skip: create is readOnly
	// skip: update is readOnly
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.Part) {
		toSerialize["part"] = o.Part
	}
	if o.Quantity.IsSet() {
		toSerialize["quantity"] = o.Quantity.Get()
	}
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
	if !IsNil(o.Unity) {
		toSerialize["unity"] = o.Unity
	}
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
	if o.TaxRate.IsSet() {
		toSerialize["taxRate"] = o.TaxRate.Get()
	}
	if o.SumDiscount.IsSet() {
		toSerialize["sumDiscount"] = o.SumDiscount.Get()
	}
	return toSerialize, nil
}

type NullableModelOrderPosUpdate struct {
	value *ModelOrderPosUpdate
	isSet bool
}

func (v NullableModelOrderPosUpdate) Get() *ModelOrderPosUpdate {
	return v.value
}

func (v *NullableModelOrderPosUpdate) Set(val *ModelOrderPosUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableModelOrderPosUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableModelOrderPosUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelOrderPosUpdate(val *ModelOrderPosUpdate) *NullableModelOrderPosUpdate {
	return &NullableModelOrderPosUpdate{value: val, isSet: true}
}

func (v NullableModelOrderPosUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelOrderPosUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


