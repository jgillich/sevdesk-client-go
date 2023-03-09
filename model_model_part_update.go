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

// checks if the ModelPartUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelPartUpdate{}

// ModelPartUpdate Part model
type ModelPartUpdate struct {
	// The part id
	Id NullableInt32 `json:"id,omitempty"`
	// The part object name
	ObjectName NullableString `json:"objectName,omitempty"`
	// Date of part creation
	Create *time.Time `json:"create,omitempty"`
	// Date of last part update
	Update *time.Time `json:"update,omitempty"`
	// Name of the part
	Name *string `json:"name,omitempty"`
	// The part number
	PartNumber *string `json:"partNumber,omitempty"`
	// A text describing the part
	Text NullableString `json:"text,omitempty"`
	Category NullableModelPartCategory `json:"category,omitempty"`
	// The stock of the part
	Stock *float32 `json:"stock,omitempty"`
	// Defines if the stock should be enabled
	StockEnabled NullableBool `json:"stockEnabled,omitempty"`
	Unity *ModelPartUnity `json:"unity,omitempty"`
	// Net price for which the part is sold. we will change this parameter so that the gross price is calculated automatically, until then the priceGross parameter must be used.
	Price NullableFloat32 `json:"price,omitempty"`
	// Net price for which the part is sold
	PriceNet NullableFloat32 `json:"priceNet,omitempty"`
	// Gross price for which the part is sold
	PriceGross NullableFloat32 `json:"priceGross,omitempty"`
	SevClient *ModelPartSevClient `json:"sevClient,omitempty"`
	// Purchase price of the part
	PricePurchase NullableFloat32 `json:"pricePurchase,omitempty"`
	// Tax rate of the part
	TaxRate *float32 `json:"taxRate,omitempty"`
	// Status of the part. 50 <-> Inactive - 100 <-> Active
	Status NullableInt32 `json:"status,omitempty"`
	// An internal comment for the part.<br>       Does not appear on invoices and orders.
	InternalComment NullableString `json:"internalComment,omitempty"`
}

// NewModelPartUpdate instantiates a new ModelPartUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelPartUpdate() *ModelPartUpdate {
	this := ModelPartUpdate{}
	return &this
}

// NewModelPartUpdateWithDefaults instantiates a new ModelPartUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelPartUpdateWithDefaults() *ModelPartUpdate {
	this := ModelPartUpdate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPartUpdate) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPartUpdate) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ModelPartUpdate) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *ModelPartUpdate) SetId(v int32) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ModelPartUpdate) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ModelPartUpdate) UnsetId() {
	o.Id.Unset()
}

// GetObjectName returns the ObjectName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPartUpdate) GetObjectName() string {
	if o == nil || IsNil(o.ObjectName.Get()) {
		var ret string
		return ret
	}
	return *o.ObjectName.Get()
}

// GetObjectNameOk returns a tuple with the ObjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPartUpdate) GetObjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObjectName.Get(), o.ObjectName.IsSet()
}

// HasObjectName returns a boolean if a field has been set.
func (o *ModelPartUpdate) HasObjectName() bool {
	if o != nil && o.ObjectName.IsSet() {
		return true
	}

	return false
}

// SetObjectName gets a reference to the given NullableString and assigns it to the ObjectName field.
func (o *ModelPartUpdate) SetObjectName(v string) {
	o.ObjectName.Set(&v)
}
// SetObjectNameNil sets the value for ObjectName to be an explicit nil
func (o *ModelPartUpdate) SetObjectNameNil() {
	o.ObjectName.Set(nil)
}

// UnsetObjectName ensures that no value is present for ObjectName, not even an explicit nil
func (o *ModelPartUpdate) UnsetObjectName() {
	o.ObjectName.Unset()
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *ModelPartUpdate) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPartUpdate) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *ModelPartUpdate) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *ModelPartUpdate) SetCreate(v time.Time) {
	o.Create = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *ModelPartUpdate) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPartUpdate) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *ModelPartUpdate) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *ModelPartUpdate) SetUpdate(v time.Time) {
	o.Update = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelPartUpdate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPartUpdate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelPartUpdate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelPartUpdate) SetName(v string) {
	o.Name = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *ModelPartUpdate) GetPartNumber() string {
	if o == nil || IsNil(o.PartNumber) {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPartUpdate) GetPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PartNumber) {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *ModelPartUpdate) HasPartNumber() bool {
	if o != nil && !IsNil(o.PartNumber) {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *ModelPartUpdate) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetText returns the Text field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPartUpdate) GetText() string {
	if o == nil || IsNil(o.Text.Get()) {
		var ret string
		return ret
	}
	return *o.Text.Get()
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPartUpdate) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Text.Get(), o.Text.IsSet()
}

// HasText returns a boolean if a field has been set.
func (o *ModelPartUpdate) HasText() bool {
	if o != nil && o.Text.IsSet() {
		return true
	}

	return false
}

// SetText gets a reference to the given NullableString and assigns it to the Text field.
func (o *ModelPartUpdate) SetText(v string) {
	o.Text.Set(&v)
}
// SetTextNil sets the value for Text to be an explicit nil
func (o *ModelPartUpdate) SetTextNil() {
	o.Text.Set(nil)
}

// UnsetText ensures that no value is present for Text, not even an explicit nil
func (o *ModelPartUpdate) UnsetText() {
	o.Text.Unset()
}

// GetCategory returns the Category field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPartUpdate) GetCategory() ModelPartCategory {
	if o == nil || IsNil(o.Category.Get()) {
		var ret ModelPartCategory
		return ret
	}
	return *o.Category.Get()
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPartUpdate) GetCategoryOk() (*ModelPartCategory, bool) {
	if o == nil {
		return nil, false
	}
	return o.Category.Get(), o.Category.IsSet()
}

// HasCategory returns a boolean if a field has been set.
func (o *ModelPartUpdate) HasCategory() bool {
	if o != nil && o.Category.IsSet() {
		return true
	}

	return false
}

// SetCategory gets a reference to the given NullableModelPartCategory and assigns it to the Category field.
func (o *ModelPartUpdate) SetCategory(v ModelPartCategory) {
	o.Category.Set(&v)
}
// SetCategoryNil sets the value for Category to be an explicit nil
func (o *ModelPartUpdate) SetCategoryNil() {
	o.Category.Set(nil)
}

// UnsetCategory ensures that no value is present for Category, not even an explicit nil
func (o *ModelPartUpdate) UnsetCategory() {
	o.Category.Unset()
}

// GetStock returns the Stock field value if set, zero value otherwise.
func (o *ModelPartUpdate) GetStock() float32 {
	if o == nil || IsNil(o.Stock) {
		var ret float32
		return ret
	}
	return *o.Stock
}

// GetStockOk returns a tuple with the Stock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPartUpdate) GetStockOk() (*float32, bool) {
	if o == nil || IsNil(o.Stock) {
		return nil, false
	}
	return o.Stock, true
}

// HasStock returns a boolean if a field has been set.
func (o *ModelPartUpdate) HasStock() bool {
	if o != nil && !IsNil(o.Stock) {
		return true
	}

	return false
}

// SetStock gets a reference to the given float32 and assigns it to the Stock field.
func (o *ModelPartUpdate) SetStock(v float32) {
	o.Stock = &v
}

// GetStockEnabled returns the StockEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPartUpdate) GetStockEnabled() bool {
	if o == nil || IsNil(o.StockEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.StockEnabled.Get()
}

// GetStockEnabledOk returns a tuple with the StockEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPartUpdate) GetStockEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.StockEnabled.Get(), o.StockEnabled.IsSet()
}

// HasStockEnabled returns a boolean if a field has been set.
func (o *ModelPartUpdate) HasStockEnabled() bool {
	if o != nil && o.StockEnabled.IsSet() {
		return true
	}

	return false
}

// SetStockEnabled gets a reference to the given NullableBool and assigns it to the StockEnabled field.
func (o *ModelPartUpdate) SetStockEnabled(v bool) {
	o.StockEnabled.Set(&v)
}
// SetStockEnabledNil sets the value for StockEnabled to be an explicit nil
func (o *ModelPartUpdate) SetStockEnabledNil() {
	o.StockEnabled.Set(nil)
}

// UnsetStockEnabled ensures that no value is present for StockEnabled, not even an explicit nil
func (o *ModelPartUpdate) UnsetStockEnabled() {
	o.StockEnabled.Unset()
}

// GetUnity returns the Unity field value if set, zero value otherwise.
func (o *ModelPartUpdate) GetUnity() ModelPartUnity {
	if o == nil || IsNil(o.Unity) {
		var ret ModelPartUnity
		return ret
	}
	return *o.Unity
}

// GetUnityOk returns a tuple with the Unity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPartUpdate) GetUnityOk() (*ModelPartUnity, bool) {
	if o == nil || IsNil(o.Unity) {
		return nil, false
	}
	return o.Unity, true
}

// HasUnity returns a boolean if a field has been set.
func (o *ModelPartUpdate) HasUnity() bool {
	if o != nil && !IsNil(o.Unity) {
		return true
	}

	return false
}

// SetUnity gets a reference to the given ModelPartUnity and assigns it to the Unity field.
func (o *ModelPartUpdate) SetUnity(v ModelPartUnity) {
	o.Unity = &v
}

// GetPrice returns the Price field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPartUpdate) GetPrice() float32 {
	if o == nil || IsNil(o.Price.Get()) {
		var ret float32
		return ret
	}
	return *o.Price.Get()
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPartUpdate) GetPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Price.Get(), o.Price.IsSet()
}

// HasPrice returns a boolean if a field has been set.
func (o *ModelPartUpdate) HasPrice() bool {
	if o != nil && o.Price.IsSet() {
		return true
	}

	return false
}

// SetPrice gets a reference to the given NullableFloat32 and assigns it to the Price field.
func (o *ModelPartUpdate) SetPrice(v float32) {
	o.Price.Set(&v)
}
// SetPriceNil sets the value for Price to be an explicit nil
func (o *ModelPartUpdate) SetPriceNil() {
	o.Price.Set(nil)
}

// UnsetPrice ensures that no value is present for Price, not even an explicit nil
func (o *ModelPartUpdate) UnsetPrice() {
	o.Price.Unset()
}

// GetPriceNet returns the PriceNet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPartUpdate) GetPriceNet() float32 {
	if o == nil || IsNil(o.PriceNet.Get()) {
		var ret float32
		return ret
	}
	return *o.PriceNet.Get()
}

// GetPriceNetOk returns a tuple with the PriceNet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPartUpdate) GetPriceNetOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PriceNet.Get(), o.PriceNet.IsSet()
}

// HasPriceNet returns a boolean if a field has been set.
func (o *ModelPartUpdate) HasPriceNet() bool {
	if o != nil && o.PriceNet.IsSet() {
		return true
	}

	return false
}

// SetPriceNet gets a reference to the given NullableFloat32 and assigns it to the PriceNet field.
func (o *ModelPartUpdate) SetPriceNet(v float32) {
	o.PriceNet.Set(&v)
}
// SetPriceNetNil sets the value for PriceNet to be an explicit nil
func (o *ModelPartUpdate) SetPriceNetNil() {
	o.PriceNet.Set(nil)
}

// UnsetPriceNet ensures that no value is present for PriceNet, not even an explicit nil
func (o *ModelPartUpdate) UnsetPriceNet() {
	o.PriceNet.Unset()
}

// GetPriceGross returns the PriceGross field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPartUpdate) GetPriceGross() float32 {
	if o == nil || IsNil(o.PriceGross.Get()) {
		var ret float32
		return ret
	}
	return *o.PriceGross.Get()
}

// GetPriceGrossOk returns a tuple with the PriceGross field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPartUpdate) GetPriceGrossOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PriceGross.Get(), o.PriceGross.IsSet()
}

// HasPriceGross returns a boolean if a field has been set.
func (o *ModelPartUpdate) HasPriceGross() bool {
	if o != nil && o.PriceGross.IsSet() {
		return true
	}

	return false
}

// SetPriceGross gets a reference to the given NullableFloat32 and assigns it to the PriceGross field.
func (o *ModelPartUpdate) SetPriceGross(v float32) {
	o.PriceGross.Set(&v)
}
// SetPriceGrossNil sets the value for PriceGross to be an explicit nil
func (o *ModelPartUpdate) SetPriceGrossNil() {
	o.PriceGross.Set(nil)
}

// UnsetPriceGross ensures that no value is present for PriceGross, not even an explicit nil
func (o *ModelPartUpdate) UnsetPriceGross() {
	o.PriceGross.Unset()
}

// GetSevClient returns the SevClient field value if set, zero value otherwise.
func (o *ModelPartUpdate) GetSevClient() ModelPartSevClient {
	if o == nil || IsNil(o.SevClient) {
		var ret ModelPartSevClient
		return ret
	}
	return *o.SevClient
}

// GetSevClientOk returns a tuple with the SevClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPartUpdate) GetSevClientOk() (*ModelPartSevClient, bool) {
	if o == nil || IsNil(o.SevClient) {
		return nil, false
	}
	return o.SevClient, true
}

// HasSevClient returns a boolean if a field has been set.
func (o *ModelPartUpdate) HasSevClient() bool {
	if o != nil && !IsNil(o.SevClient) {
		return true
	}

	return false
}

// SetSevClient gets a reference to the given ModelPartSevClient and assigns it to the SevClient field.
func (o *ModelPartUpdate) SetSevClient(v ModelPartSevClient) {
	o.SevClient = &v
}

// GetPricePurchase returns the PricePurchase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPartUpdate) GetPricePurchase() float32 {
	if o == nil || IsNil(o.PricePurchase.Get()) {
		var ret float32
		return ret
	}
	return *o.PricePurchase.Get()
}

// GetPricePurchaseOk returns a tuple with the PricePurchase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPartUpdate) GetPricePurchaseOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PricePurchase.Get(), o.PricePurchase.IsSet()
}

// HasPricePurchase returns a boolean if a field has been set.
func (o *ModelPartUpdate) HasPricePurchase() bool {
	if o != nil && o.PricePurchase.IsSet() {
		return true
	}

	return false
}

// SetPricePurchase gets a reference to the given NullableFloat32 and assigns it to the PricePurchase field.
func (o *ModelPartUpdate) SetPricePurchase(v float32) {
	o.PricePurchase.Set(&v)
}
// SetPricePurchaseNil sets the value for PricePurchase to be an explicit nil
func (o *ModelPartUpdate) SetPricePurchaseNil() {
	o.PricePurchase.Set(nil)
}

// UnsetPricePurchase ensures that no value is present for PricePurchase, not even an explicit nil
func (o *ModelPartUpdate) UnsetPricePurchase() {
	o.PricePurchase.Unset()
}

// GetTaxRate returns the TaxRate field value if set, zero value otherwise.
func (o *ModelPartUpdate) GetTaxRate() float32 {
	if o == nil || IsNil(o.TaxRate) {
		var ret float32
		return ret
	}
	return *o.TaxRate
}

// GetTaxRateOk returns a tuple with the TaxRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPartUpdate) GetTaxRateOk() (*float32, bool) {
	if o == nil || IsNil(o.TaxRate) {
		return nil, false
	}
	return o.TaxRate, true
}

// HasTaxRate returns a boolean if a field has been set.
func (o *ModelPartUpdate) HasTaxRate() bool {
	if o != nil && !IsNil(o.TaxRate) {
		return true
	}

	return false
}

// SetTaxRate gets a reference to the given float32 and assigns it to the TaxRate field.
func (o *ModelPartUpdate) SetTaxRate(v float32) {
	o.TaxRate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPartUpdate) GetStatus() int32 {
	if o == nil || IsNil(o.Status.Get()) {
		var ret int32
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPartUpdate) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *ModelPartUpdate) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableInt32 and assigns it to the Status field.
func (o *ModelPartUpdate) SetStatus(v int32) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *ModelPartUpdate) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *ModelPartUpdate) UnsetStatus() {
	o.Status.Unset()
}

// GetInternalComment returns the InternalComment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPartUpdate) GetInternalComment() string {
	if o == nil || IsNil(o.InternalComment.Get()) {
		var ret string
		return ret
	}
	return *o.InternalComment.Get()
}

// GetInternalCommentOk returns a tuple with the InternalComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPartUpdate) GetInternalCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InternalComment.Get(), o.InternalComment.IsSet()
}

// HasInternalComment returns a boolean if a field has been set.
func (o *ModelPartUpdate) HasInternalComment() bool {
	if o != nil && o.InternalComment.IsSet() {
		return true
	}

	return false
}

// SetInternalComment gets a reference to the given NullableString and assigns it to the InternalComment field.
func (o *ModelPartUpdate) SetInternalComment(v string) {
	o.InternalComment.Set(&v)
}
// SetInternalCommentNil sets the value for InternalComment to be an explicit nil
func (o *ModelPartUpdate) SetInternalCommentNil() {
	o.InternalComment.Set(nil)
}

// UnsetInternalComment ensures that no value is present for InternalComment, not even an explicit nil
func (o *ModelPartUpdate) UnsetInternalComment() {
	o.InternalComment.Unset()
}

func (o ModelPartUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelPartUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.ObjectName.IsSet() {
		toSerialize["objectName"] = o.ObjectName.Get()
	}
	// skip: create is readOnly
	// skip: update is readOnly
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PartNumber) {
		toSerialize["partNumber"] = o.PartNumber
	}
	if o.Text.IsSet() {
		toSerialize["text"] = o.Text.Get()
	}
	if o.Category.IsSet() {
		toSerialize["category"] = o.Category.Get()
	}
	if !IsNil(o.Stock) {
		toSerialize["stock"] = o.Stock
	}
	if o.StockEnabled.IsSet() {
		toSerialize["stockEnabled"] = o.StockEnabled.Get()
	}
	if !IsNil(o.Unity) {
		toSerialize["unity"] = o.Unity
	}
	if o.Price.IsSet() {
		toSerialize["price"] = o.Price.Get()
	}
	if o.PriceNet.IsSet() {
		toSerialize["priceNet"] = o.PriceNet.Get()
	}
	if o.PriceGross.IsSet() {
		toSerialize["priceGross"] = o.PriceGross.Get()
	}
	if !IsNil(o.SevClient) {
		toSerialize["sevClient"] = o.SevClient
	}
	if o.PricePurchase.IsSet() {
		toSerialize["pricePurchase"] = o.PricePurchase.Get()
	}
	if !IsNil(o.TaxRate) {
		toSerialize["taxRate"] = o.TaxRate
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.InternalComment.IsSet() {
		toSerialize["internalComment"] = o.InternalComment.Get()
	}
	return toSerialize, nil
}

type NullableModelPartUpdate struct {
	value *ModelPartUpdate
	isSet bool
}

func (v NullableModelPartUpdate) Get() *ModelPartUpdate {
	return v.value
}

func (v *NullableModelPartUpdate) Set(val *ModelPartUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableModelPartUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableModelPartUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelPartUpdate(val *ModelPartUpdate) *NullableModelPartUpdate {
	return &NullableModelPartUpdate{value: val, isSet: true}
}

func (v NullableModelPartUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelPartUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


