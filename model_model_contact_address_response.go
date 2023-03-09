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

// checks if the ModelContactAddressResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelContactAddressResponse{}

// ModelContactAddressResponse ContactAddress model
type ModelContactAddressResponse struct {
	// The contact address id
	Id *int32 `json:"id,omitempty"`
	// The contact address object name
	ObjectName *string `json:"objectName,omitempty"`
	// Date of contact address creation
	Create *time.Time `json:"create,omitempty"`
	// Date of last contact address update
	Update *time.Time `json:"update,omitempty"`
	Contact ModelContactAddressContact `json:"contact"`
	// Street name
	Street NullableString `json:"street,omitempty"`
	// Zib code
	Zip NullableString `json:"zip,omitempty"`
	// City name
	City NullableString `json:"city,omitempty"`
	Country ModelContactAddressCountry `json:"country"`
	Category NullableModelContactAddressCategory `json:"category,omitempty"`
	// Name in address
	Name NullableString `json:"name,omitempty"`
	SevClient *ModelContactAddressSevClient `json:"sevClient,omitempty"`
	// Second name in address
	Name2 *string `json:"name2,omitempty"`
	// Third name in address
	Name3 NullableString `json:"name3,omitempty"`
	// Fourth name in address
	Name4 NullableString `json:"name4,omitempty"`
}

// NewModelContactAddressResponse instantiates a new ModelContactAddressResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelContactAddressResponse(contact ModelContactAddressContact, country ModelContactAddressCountry) *ModelContactAddressResponse {
	this := ModelContactAddressResponse{}
	this.Contact = contact
	this.Country = country
	return &this
}

// NewModelContactAddressResponseWithDefaults instantiates a new ModelContactAddressResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelContactAddressResponseWithDefaults() *ModelContactAddressResponse {
	this := ModelContactAddressResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelContactAddressResponse) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContactAddressResponse) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelContactAddressResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ModelContactAddressResponse) SetId(v int32) {
	o.Id = &v
}

// GetObjectName returns the ObjectName field value if set, zero value otherwise.
func (o *ModelContactAddressResponse) GetObjectName() string {
	if o == nil || IsNil(o.ObjectName) {
		var ret string
		return ret
	}
	return *o.ObjectName
}

// GetObjectNameOk returns a tuple with the ObjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContactAddressResponse) GetObjectNameOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectName) {
		return nil, false
	}
	return o.ObjectName, true
}

// HasObjectName returns a boolean if a field has been set.
func (o *ModelContactAddressResponse) HasObjectName() bool {
	if o != nil && !IsNil(o.ObjectName) {
		return true
	}

	return false
}

// SetObjectName gets a reference to the given string and assigns it to the ObjectName field.
func (o *ModelContactAddressResponse) SetObjectName(v string) {
	o.ObjectName = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *ModelContactAddressResponse) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContactAddressResponse) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *ModelContactAddressResponse) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *ModelContactAddressResponse) SetCreate(v time.Time) {
	o.Create = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *ModelContactAddressResponse) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContactAddressResponse) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *ModelContactAddressResponse) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *ModelContactAddressResponse) SetUpdate(v time.Time) {
	o.Update = &v
}

// GetContact returns the Contact field value
func (o *ModelContactAddressResponse) GetContact() ModelContactAddressContact {
	if o == nil {
		var ret ModelContactAddressContact
		return ret
	}

	return o.Contact
}

// GetContactOk returns a tuple with the Contact field value
// and a boolean to check if the value has been set.
func (o *ModelContactAddressResponse) GetContactOk() (*ModelContactAddressContact, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contact, true
}

// SetContact sets field value
func (o *ModelContactAddressResponse) SetContact(v ModelContactAddressContact) {
	o.Contact = v
}

// GetStreet returns the Street field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelContactAddressResponse) GetStreet() string {
	if o == nil || IsNil(o.Street.Get()) {
		var ret string
		return ret
	}
	return *o.Street.Get()
}

// GetStreetOk returns a tuple with the Street field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContactAddressResponse) GetStreetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Street.Get(), o.Street.IsSet()
}

// HasStreet returns a boolean if a field has been set.
func (o *ModelContactAddressResponse) HasStreet() bool {
	if o != nil && o.Street.IsSet() {
		return true
	}

	return false
}

// SetStreet gets a reference to the given NullableString and assigns it to the Street field.
func (o *ModelContactAddressResponse) SetStreet(v string) {
	o.Street.Set(&v)
}
// SetStreetNil sets the value for Street to be an explicit nil
func (o *ModelContactAddressResponse) SetStreetNil() {
	o.Street.Set(nil)
}

// UnsetStreet ensures that no value is present for Street, not even an explicit nil
func (o *ModelContactAddressResponse) UnsetStreet() {
	o.Street.Unset()
}

// GetZip returns the Zip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelContactAddressResponse) GetZip() string {
	if o == nil || IsNil(o.Zip.Get()) {
		var ret string
		return ret
	}
	return *o.Zip.Get()
}

// GetZipOk returns a tuple with the Zip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContactAddressResponse) GetZipOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Zip.Get(), o.Zip.IsSet()
}

// HasZip returns a boolean if a field has been set.
func (o *ModelContactAddressResponse) HasZip() bool {
	if o != nil && o.Zip.IsSet() {
		return true
	}

	return false
}

// SetZip gets a reference to the given NullableString and assigns it to the Zip field.
func (o *ModelContactAddressResponse) SetZip(v string) {
	o.Zip.Set(&v)
}
// SetZipNil sets the value for Zip to be an explicit nil
func (o *ModelContactAddressResponse) SetZipNil() {
	o.Zip.Set(nil)
}

// UnsetZip ensures that no value is present for Zip, not even an explicit nil
func (o *ModelContactAddressResponse) UnsetZip() {
	o.Zip.Unset()
}

// GetCity returns the City field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelContactAddressResponse) GetCity() string {
	if o == nil || IsNil(o.City.Get()) {
		var ret string
		return ret
	}
	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContactAddressResponse) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// HasCity returns a boolean if a field has been set.
func (o *ModelContactAddressResponse) HasCity() bool {
	if o != nil && o.City.IsSet() {
		return true
	}

	return false
}

// SetCity gets a reference to the given NullableString and assigns it to the City field.
func (o *ModelContactAddressResponse) SetCity(v string) {
	o.City.Set(&v)
}
// SetCityNil sets the value for City to be an explicit nil
func (o *ModelContactAddressResponse) SetCityNil() {
	o.City.Set(nil)
}

// UnsetCity ensures that no value is present for City, not even an explicit nil
func (o *ModelContactAddressResponse) UnsetCity() {
	o.City.Unset()
}

// GetCountry returns the Country field value
func (o *ModelContactAddressResponse) GetCountry() ModelContactAddressCountry {
	if o == nil {
		var ret ModelContactAddressCountry
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *ModelContactAddressResponse) GetCountryOk() (*ModelContactAddressCountry, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *ModelContactAddressResponse) SetCountry(v ModelContactAddressCountry) {
	o.Country = v
}

// GetCategory returns the Category field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelContactAddressResponse) GetCategory() ModelContactAddressCategory {
	if o == nil || IsNil(o.Category.Get()) {
		var ret ModelContactAddressCategory
		return ret
	}
	return *o.Category.Get()
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContactAddressResponse) GetCategoryOk() (*ModelContactAddressCategory, bool) {
	if o == nil {
		return nil, false
	}
	return o.Category.Get(), o.Category.IsSet()
}

// HasCategory returns a boolean if a field has been set.
func (o *ModelContactAddressResponse) HasCategory() bool {
	if o != nil && o.Category.IsSet() {
		return true
	}

	return false
}

// SetCategory gets a reference to the given NullableModelContactAddressCategory and assigns it to the Category field.
func (o *ModelContactAddressResponse) SetCategory(v ModelContactAddressCategory) {
	o.Category.Set(&v)
}
// SetCategoryNil sets the value for Category to be an explicit nil
func (o *ModelContactAddressResponse) SetCategoryNil() {
	o.Category.Set(nil)
}

// UnsetCategory ensures that no value is present for Category, not even an explicit nil
func (o *ModelContactAddressResponse) UnsetCategory() {
	o.Category.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelContactAddressResponse) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContactAddressResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ModelContactAddressResponse) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ModelContactAddressResponse) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ModelContactAddressResponse) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ModelContactAddressResponse) UnsetName() {
	o.Name.Unset()
}

// GetSevClient returns the SevClient field value if set, zero value otherwise.
func (o *ModelContactAddressResponse) GetSevClient() ModelContactAddressSevClient {
	if o == nil || IsNil(o.SevClient) {
		var ret ModelContactAddressSevClient
		return ret
	}
	return *o.SevClient
}

// GetSevClientOk returns a tuple with the SevClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContactAddressResponse) GetSevClientOk() (*ModelContactAddressSevClient, bool) {
	if o == nil || IsNil(o.SevClient) {
		return nil, false
	}
	return o.SevClient, true
}

// HasSevClient returns a boolean if a field has been set.
func (o *ModelContactAddressResponse) HasSevClient() bool {
	if o != nil && !IsNil(o.SevClient) {
		return true
	}

	return false
}

// SetSevClient gets a reference to the given ModelContactAddressSevClient and assigns it to the SevClient field.
func (o *ModelContactAddressResponse) SetSevClient(v ModelContactAddressSevClient) {
	o.SevClient = &v
}

// GetName2 returns the Name2 field value if set, zero value otherwise.
func (o *ModelContactAddressResponse) GetName2() string {
	if o == nil || IsNil(o.Name2) {
		var ret string
		return ret
	}
	return *o.Name2
}

// GetName2Ok returns a tuple with the Name2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContactAddressResponse) GetName2Ok() (*string, bool) {
	if o == nil || IsNil(o.Name2) {
		return nil, false
	}
	return o.Name2, true
}

// HasName2 returns a boolean if a field has been set.
func (o *ModelContactAddressResponse) HasName2() bool {
	if o != nil && !IsNil(o.Name2) {
		return true
	}

	return false
}

// SetName2 gets a reference to the given string and assigns it to the Name2 field.
func (o *ModelContactAddressResponse) SetName2(v string) {
	o.Name2 = &v
}

// GetName3 returns the Name3 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelContactAddressResponse) GetName3() string {
	if o == nil || IsNil(o.Name3.Get()) {
		var ret string
		return ret
	}
	return *o.Name3.Get()
}

// GetName3Ok returns a tuple with the Name3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContactAddressResponse) GetName3Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name3.Get(), o.Name3.IsSet()
}

// HasName3 returns a boolean if a field has been set.
func (o *ModelContactAddressResponse) HasName3() bool {
	if o != nil && o.Name3.IsSet() {
		return true
	}

	return false
}

// SetName3 gets a reference to the given NullableString and assigns it to the Name3 field.
func (o *ModelContactAddressResponse) SetName3(v string) {
	o.Name3.Set(&v)
}
// SetName3Nil sets the value for Name3 to be an explicit nil
func (o *ModelContactAddressResponse) SetName3Nil() {
	o.Name3.Set(nil)
}

// UnsetName3 ensures that no value is present for Name3, not even an explicit nil
func (o *ModelContactAddressResponse) UnsetName3() {
	o.Name3.Unset()
}

// GetName4 returns the Name4 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelContactAddressResponse) GetName4() string {
	if o == nil || IsNil(o.Name4.Get()) {
		var ret string
		return ret
	}
	return *o.Name4.Get()
}

// GetName4Ok returns a tuple with the Name4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContactAddressResponse) GetName4Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name4.Get(), o.Name4.IsSet()
}

// HasName4 returns a boolean if a field has been set.
func (o *ModelContactAddressResponse) HasName4() bool {
	if o != nil && o.Name4.IsSet() {
		return true
	}

	return false
}

// SetName4 gets a reference to the given NullableString and assigns it to the Name4 field.
func (o *ModelContactAddressResponse) SetName4(v string) {
	o.Name4.Set(&v)
}
// SetName4Nil sets the value for Name4 to be an explicit nil
func (o *ModelContactAddressResponse) SetName4Nil() {
	o.Name4.Set(nil)
}

// UnsetName4 ensures that no value is present for Name4, not even an explicit nil
func (o *ModelContactAddressResponse) UnsetName4() {
	o.Name4.Unset()
}

func (o ModelContactAddressResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelContactAddressResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	// skip: objectName is readOnly
	// skip: create is readOnly
	// skip: update is readOnly
	toSerialize["contact"] = o.Contact
	if o.Street.IsSet() {
		toSerialize["street"] = o.Street.Get()
	}
	if o.Zip.IsSet() {
		toSerialize["zip"] = o.Zip.Get()
	}
	if o.City.IsSet() {
		toSerialize["city"] = o.City.Get()
	}
	toSerialize["country"] = o.Country
	if o.Category.IsSet() {
		toSerialize["category"] = o.Category.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.SevClient) {
		toSerialize["sevClient"] = o.SevClient
	}
	if !IsNil(o.Name2) {
		toSerialize["name2"] = o.Name2
	}
	if o.Name3.IsSet() {
		toSerialize["name3"] = o.Name3.Get()
	}
	if o.Name4.IsSet() {
		toSerialize["name4"] = o.Name4.Get()
	}
	return toSerialize, nil
}

type NullableModelContactAddressResponse struct {
	value *ModelContactAddressResponse
	isSet bool
}

func (v NullableModelContactAddressResponse) Get() *ModelContactAddressResponse {
	return v.value
}

func (v *NullableModelContactAddressResponse) Set(val *ModelContactAddressResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModelContactAddressResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModelContactAddressResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelContactAddressResponse(val *ModelContactAddressResponse) *NullableModelContactAddressResponse {
	return &NullableModelContactAddressResponse{value: val, isSet: true}
}

func (v NullableModelContactAddressResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelContactAddressResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


