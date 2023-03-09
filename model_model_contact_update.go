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

// checks if the ModelContactUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelContactUpdate{}

// ModelContactUpdate Contact model
type ModelContactUpdate struct {
	// The organization name.<br> Be aware that the type of contact will depend on this attribute.<br> If it holds a value, the contact will be regarded as an organization.
	Name NullableString `json:"name,omitempty"`
	// Defines the status of the contact. 100 <-> Lead - 500 <-> Pending - 1000 <-> Active.
	Status NullableInt32 `json:"status,omitempty"`
	// The customer number
	CustomerNumber NullableString `json:"customerNumber,omitempty"`
	Parent NullableModelContactUpdateParent `json:"parent,omitempty"`
	// The <b>first</b> name of the contact.<br> Yeah... not quite right in literally every way. We know.<br> Not to be used for organizations.
	Surename NullableString `json:"surename,omitempty"`
	// The last name of the contact.<br> Not to be used for organizations.
	Familyname NullableString `json:"familyname,omitempty"`
	// A non-academic title for the contact. Not to be used for organizations.
	Titel NullableString `json:"titel,omitempty"`
	Category NullableModelContactUpdateCategory `json:"category,omitempty"`
	// A description for the contact.
	Description NullableString `json:"description,omitempty"`
	// A academic title for the contact. Not to be used for organizations.
	AcademicTitle NullableString `json:"academicTitle,omitempty"`
	// Gender of the contact.<br> Not to be used for organizations.
	Gender NullableString `json:"gender,omitempty"`
	// Second name of the contact.<br> Not to be used for organizations.
	Name2 NullableString `json:"name2,omitempty"`
	// Birthday of the contact.<br> Not to be used for organizations.
	Birthday NullableString `json:"birthday,omitempty"`
	// Vat number of the contact.
	VatNumber NullableString `json:"vatNumber,omitempty"`
	// Bank account number (IBAN) of the contact.
	BankAccount NullableString `json:"bankAccount,omitempty"`
	// Bank number of the bank used by the contact.
	BankNumber NullableString `json:"bankNumber,omitempty"`
	// Absolute time in days which the contact has to pay his invoices and subsequently get a cashback.
	DefaultCashbackTime NullableInt32 `json:"defaultCashbackTime,omitempty"`
	// Percentage of the invoice sum the contact gets back if he payed invoices in time.
	DefaultCashbackPercent NullableFloat32 `json:"defaultCashbackPercent,omitempty"`
	// The payment goal in days which is set for every invoice of the contact.
	DefaultTimeToPay NullableInt32 `json:"defaultTimeToPay,omitempty"`
	// The tax number of the contact.
	TaxNumber NullableString `json:"taxNumber,omitempty"`
	// The tax office of the contact (only for greek customers).
	TaxOffice NullableString `json:"taxOffice,omitempty"`
	// Defines if the contact is freed from paying vat.
	ExemptVat NullableBool `json:"exemptVat,omitempty"`
	// Defines which tax regulation the contact is using.
	TaxType NullableString `json:"taxType,omitempty"`
	TaxSet NullableModelContactTaxSet `json:"taxSet,omitempty"`
	// The default discount the contact gets for every invoice.<br> Depending on defaultDiscountPercentage attribute, in percent or absolute value.
	DefaultDiscountAmount NullableFloat32 `json:"defaultDiscountAmount,omitempty"`
	// Defines if the discount is a percentage (true) or an absolute value (false).
	DefaultDiscountPercentage NullableBool `json:"defaultDiscountPercentage,omitempty"`
	// Buyer reference of the contact.
	BuyerReference NullableString `json:"buyerReference,omitempty"`
	// Defines whether the contact is a government agency (true) or not (false).
	GovernmentAgency NullableBool `json:"governmentAgency,omitempty"`
}

// NewModelContactUpdate instantiates a new ModelContactUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelContactUpdate() *ModelContactUpdate {
	this := ModelContactUpdate{}
	var status int32 = 100
	this.Status = *NewNullableInt32(&status)
	return &this
}

// NewModelContactUpdateWithDefaults instantiates a new ModelContactUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelContactUpdateWithDefaults() *ModelContactUpdate {
	this := ModelContactUpdate{}
	var status int32 = 100
	this.Status = *NewNullableInt32(&status)
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelContactUpdate) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContactUpdate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ModelContactUpdate) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ModelContactUpdate) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ModelContactUpdate) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ModelContactUpdate) UnsetName() {
	o.Name.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelContactUpdate) GetStatus() int32 {
	if o == nil || IsNil(o.Status.Get()) {
		var ret int32
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContactUpdate) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *ModelContactUpdate) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableInt32 and assigns it to the Status field.
func (o *ModelContactUpdate) SetStatus(v int32) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *ModelContactUpdate) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *ModelContactUpdate) UnsetStatus() {
	o.Status.Unset()
}

// GetCustomerNumber returns the CustomerNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelContactUpdate) GetCustomerNumber() string {
	if o == nil || IsNil(o.CustomerNumber.Get()) {
		var ret string
		return ret
	}
	return *o.CustomerNumber.Get()
}

// GetCustomerNumberOk returns a tuple with the CustomerNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContactUpdate) GetCustomerNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerNumber.Get(), o.CustomerNumber.IsSet()
}

// HasCustomerNumber returns a boolean if a field has been set.
func (o *ModelContactUpdate) HasCustomerNumber() bool {
	if o != nil && o.CustomerNumber.IsSet() {
		return true
	}

	return false
}

// SetCustomerNumber gets a reference to the given NullableString and assigns it to the CustomerNumber field.
func (o *ModelContactUpdate) SetCustomerNumber(v string) {
	o.CustomerNumber.Set(&v)
}
// SetCustomerNumberNil sets the value for CustomerNumber to be an explicit nil
func (o *ModelContactUpdate) SetCustomerNumberNil() {
	o.CustomerNumber.Set(nil)
}

// UnsetCustomerNumber ensures that no value is present for CustomerNumber, not even an explicit nil
func (o *ModelContactUpdate) UnsetCustomerNumber() {
	o.CustomerNumber.Unset()
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelContactUpdate) GetParent() ModelContactUpdateParent {
	if o == nil || IsNil(o.Parent.Get()) {
		var ret ModelContactUpdateParent
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContactUpdate) GetParentOk() (*ModelContactUpdateParent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *ModelContactUpdate) HasParent() bool {
	if o != nil && o.Parent.IsSet() {
		return true
	}

	return false
}

// SetParent gets a reference to the given NullableModelContactUpdateParent and assigns it to the Parent field.
func (o *ModelContactUpdate) SetParent(v ModelContactUpdateParent) {
	o.Parent.Set(&v)
}
// SetParentNil sets the value for Parent to be an explicit nil
func (o *ModelContactUpdate) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil
func (o *ModelContactUpdate) UnsetParent() {
	o.Parent.Unset()
}

// GetSurename returns the Surename field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelContactUpdate) GetSurename() string {
	if o == nil || IsNil(o.Surename.Get()) {
		var ret string
		return ret
	}
	return *o.Surename.Get()
}

// GetSurenameOk returns a tuple with the Surename field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContactUpdate) GetSurenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Surename.Get(), o.Surename.IsSet()
}

// HasSurename returns a boolean if a field has been set.
func (o *ModelContactUpdate) HasSurename() bool {
	if o != nil && o.Surename.IsSet() {
		return true
	}

	return false
}

// SetSurename gets a reference to the given NullableString and assigns it to the Surename field.
func (o *ModelContactUpdate) SetSurename(v string) {
	o.Surename.Set(&v)
}
// SetSurenameNil sets the value for Surename to be an explicit nil
func (o *ModelContactUpdate) SetSurenameNil() {
	o.Surename.Set(nil)
}

// UnsetSurename ensures that no value is present for Surename, not even an explicit nil
func (o *ModelContactUpdate) UnsetSurename() {
	o.Surename.Unset()
}

// GetFamilyname returns the Familyname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelContactUpdate) GetFamilyname() string {
	if o == nil || IsNil(o.Familyname.Get()) {
		var ret string
		return ret
	}
	return *o.Familyname.Get()
}

// GetFamilynameOk returns a tuple with the Familyname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContactUpdate) GetFamilynameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Familyname.Get(), o.Familyname.IsSet()
}

// HasFamilyname returns a boolean if a field has been set.
func (o *ModelContactUpdate) HasFamilyname() bool {
	if o != nil && o.Familyname.IsSet() {
		return true
	}

	return false
}

// SetFamilyname gets a reference to the given NullableString and assigns it to the Familyname field.
func (o *ModelContactUpdate) SetFamilyname(v string) {
	o.Familyname.Set(&v)
}
// SetFamilynameNil sets the value for Familyname to be an explicit nil
func (o *ModelContactUpdate) SetFamilynameNil() {
	o.Familyname.Set(nil)
}

// UnsetFamilyname ensures that no value is present for Familyname, not even an explicit nil
func (o *ModelContactUpdate) UnsetFamilyname() {
	o.Familyname.Unset()
}

// GetTitel returns the Titel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelContactUpdate) GetTitel() string {
	if o == nil || IsNil(o.Titel.Get()) {
		var ret string
		return ret
	}
	return *o.Titel.Get()
}

// GetTitelOk returns a tuple with the Titel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContactUpdate) GetTitelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Titel.Get(), o.Titel.IsSet()
}

// HasTitel returns a boolean if a field has been set.
func (o *ModelContactUpdate) HasTitel() bool {
	if o != nil && o.Titel.IsSet() {
		return true
	}

	return false
}

// SetTitel gets a reference to the given NullableString and assigns it to the Titel field.
func (o *ModelContactUpdate) SetTitel(v string) {
	o.Titel.Set(&v)
}
// SetTitelNil sets the value for Titel to be an explicit nil
func (o *ModelContactUpdate) SetTitelNil() {
	o.Titel.Set(nil)
}

// UnsetTitel ensures that no value is present for Titel, not even an explicit nil
func (o *ModelContactUpdate) UnsetTitel() {
	o.Titel.Unset()
}

// GetCategory returns the Category field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelContactUpdate) GetCategory() ModelContactUpdateCategory {
	if o == nil || IsNil(o.Category.Get()) {
		var ret ModelContactUpdateCategory
		return ret
	}
	return *o.Category.Get()
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContactUpdate) GetCategoryOk() (*ModelContactUpdateCategory, bool) {
	if o == nil {
		return nil, false
	}
	return o.Category.Get(), o.Category.IsSet()
}

// HasCategory returns a boolean if a field has been set.
func (o *ModelContactUpdate) HasCategory() bool {
	if o != nil && o.Category.IsSet() {
		return true
	}

	return false
}

// SetCategory gets a reference to the given NullableModelContactUpdateCategory and assigns it to the Category field.
func (o *ModelContactUpdate) SetCategory(v ModelContactUpdateCategory) {
	o.Category.Set(&v)
}
// SetCategoryNil sets the value for Category to be an explicit nil
func (o *ModelContactUpdate) SetCategoryNil() {
	o.Category.Set(nil)
}

// UnsetCategory ensures that no value is present for Category, not even an explicit nil
func (o *ModelContactUpdate) UnsetCategory() {
	o.Category.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelContactUpdate) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContactUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ModelContactUpdate) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ModelContactUpdate) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ModelContactUpdate) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ModelContactUpdate) UnsetDescription() {
	o.Description.Unset()
}

// GetAcademicTitle returns the AcademicTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelContactUpdate) GetAcademicTitle() string {
	if o == nil || IsNil(o.AcademicTitle.Get()) {
		var ret string
		return ret
	}
	return *o.AcademicTitle.Get()
}

// GetAcademicTitleOk returns a tuple with the AcademicTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContactUpdate) GetAcademicTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AcademicTitle.Get(), o.AcademicTitle.IsSet()
}

// HasAcademicTitle returns a boolean if a field has been set.
func (o *ModelContactUpdate) HasAcademicTitle() bool {
	if o != nil && o.AcademicTitle.IsSet() {
		return true
	}

	return false
}

// SetAcademicTitle gets a reference to the given NullableString and assigns it to the AcademicTitle field.
func (o *ModelContactUpdate) SetAcademicTitle(v string) {
	o.AcademicTitle.Set(&v)
}
// SetAcademicTitleNil sets the value for AcademicTitle to be an explicit nil
func (o *ModelContactUpdate) SetAcademicTitleNil() {
	o.AcademicTitle.Set(nil)
}

// UnsetAcademicTitle ensures that no value is present for AcademicTitle, not even an explicit nil
func (o *ModelContactUpdate) UnsetAcademicTitle() {
	o.AcademicTitle.Unset()
}

// GetGender returns the Gender field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelContactUpdate) GetGender() string {
	if o == nil || IsNil(o.Gender.Get()) {
		var ret string
		return ret
	}
	return *o.Gender.Get()
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContactUpdate) GetGenderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gender.Get(), o.Gender.IsSet()
}

// HasGender returns a boolean if a field has been set.
func (o *ModelContactUpdate) HasGender() bool {
	if o != nil && o.Gender.IsSet() {
		return true
	}

	return false
}

// SetGender gets a reference to the given NullableString and assigns it to the Gender field.
func (o *ModelContactUpdate) SetGender(v string) {
	o.Gender.Set(&v)
}
// SetGenderNil sets the value for Gender to be an explicit nil
func (o *ModelContactUpdate) SetGenderNil() {
	o.Gender.Set(nil)
}

// UnsetGender ensures that no value is present for Gender, not even an explicit nil
func (o *ModelContactUpdate) UnsetGender() {
	o.Gender.Unset()
}

// GetName2 returns the Name2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelContactUpdate) GetName2() string {
	if o == nil || IsNil(o.Name2.Get()) {
		var ret string
		return ret
	}
	return *o.Name2.Get()
}

// GetName2Ok returns a tuple with the Name2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContactUpdate) GetName2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name2.Get(), o.Name2.IsSet()
}

// HasName2 returns a boolean if a field has been set.
func (o *ModelContactUpdate) HasName2() bool {
	if o != nil && o.Name2.IsSet() {
		return true
	}

	return false
}

// SetName2 gets a reference to the given NullableString and assigns it to the Name2 field.
func (o *ModelContactUpdate) SetName2(v string) {
	o.Name2.Set(&v)
}
// SetName2Nil sets the value for Name2 to be an explicit nil
func (o *ModelContactUpdate) SetName2Nil() {
	o.Name2.Set(nil)
}

// UnsetName2 ensures that no value is present for Name2, not even an explicit nil
func (o *ModelContactUpdate) UnsetName2() {
	o.Name2.Unset()
}

// GetBirthday returns the Birthday field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelContactUpdate) GetBirthday() string {
	if o == nil || IsNil(o.Birthday.Get()) {
		var ret string
		return ret
	}
	return *o.Birthday.Get()
}

// GetBirthdayOk returns a tuple with the Birthday field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContactUpdate) GetBirthdayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Birthday.Get(), o.Birthday.IsSet()
}

// HasBirthday returns a boolean if a field has been set.
func (o *ModelContactUpdate) HasBirthday() bool {
	if o != nil && o.Birthday.IsSet() {
		return true
	}

	return false
}

// SetBirthday gets a reference to the given NullableString and assigns it to the Birthday field.
func (o *ModelContactUpdate) SetBirthday(v string) {
	o.Birthday.Set(&v)
}
// SetBirthdayNil sets the value for Birthday to be an explicit nil
func (o *ModelContactUpdate) SetBirthdayNil() {
	o.Birthday.Set(nil)
}

// UnsetBirthday ensures that no value is present for Birthday, not even an explicit nil
func (o *ModelContactUpdate) UnsetBirthday() {
	o.Birthday.Unset()
}

// GetVatNumber returns the VatNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelContactUpdate) GetVatNumber() string {
	if o == nil || IsNil(o.VatNumber.Get()) {
		var ret string
		return ret
	}
	return *o.VatNumber.Get()
}

// GetVatNumberOk returns a tuple with the VatNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContactUpdate) GetVatNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VatNumber.Get(), o.VatNumber.IsSet()
}

// HasVatNumber returns a boolean if a field has been set.
func (o *ModelContactUpdate) HasVatNumber() bool {
	if o != nil && o.VatNumber.IsSet() {
		return true
	}

	return false
}

// SetVatNumber gets a reference to the given NullableString and assigns it to the VatNumber field.
func (o *ModelContactUpdate) SetVatNumber(v string) {
	o.VatNumber.Set(&v)
}
// SetVatNumberNil sets the value for VatNumber to be an explicit nil
func (o *ModelContactUpdate) SetVatNumberNil() {
	o.VatNumber.Set(nil)
}

// UnsetVatNumber ensures that no value is present for VatNumber, not even an explicit nil
func (o *ModelContactUpdate) UnsetVatNumber() {
	o.VatNumber.Unset()
}

// GetBankAccount returns the BankAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelContactUpdate) GetBankAccount() string {
	if o == nil || IsNil(o.BankAccount.Get()) {
		var ret string
		return ret
	}
	return *o.BankAccount.Get()
}

// GetBankAccountOk returns a tuple with the BankAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContactUpdate) GetBankAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BankAccount.Get(), o.BankAccount.IsSet()
}

// HasBankAccount returns a boolean if a field has been set.
func (o *ModelContactUpdate) HasBankAccount() bool {
	if o != nil && o.BankAccount.IsSet() {
		return true
	}

	return false
}

// SetBankAccount gets a reference to the given NullableString and assigns it to the BankAccount field.
func (o *ModelContactUpdate) SetBankAccount(v string) {
	o.BankAccount.Set(&v)
}
// SetBankAccountNil sets the value for BankAccount to be an explicit nil
func (o *ModelContactUpdate) SetBankAccountNil() {
	o.BankAccount.Set(nil)
}

// UnsetBankAccount ensures that no value is present for BankAccount, not even an explicit nil
func (o *ModelContactUpdate) UnsetBankAccount() {
	o.BankAccount.Unset()
}

// GetBankNumber returns the BankNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelContactUpdate) GetBankNumber() string {
	if o == nil || IsNil(o.BankNumber.Get()) {
		var ret string
		return ret
	}
	return *o.BankNumber.Get()
}

// GetBankNumberOk returns a tuple with the BankNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContactUpdate) GetBankNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BankNumber.Get(), o.BankNumber.IsSet()
}

// HasBankNumber returns a boolean if a field has been set.
func (o *ModelContactUpdate) HasBankNumber() bool {
	if o != nil && o.BankNumber.IsSet() {
		return true
	}

	return false
}

// SetBankNumber gets a reference to the given NullableString and assigns it to the BankNumber field.
func (o *ModelContactUpdate) SetBankNumber(v string) {
	o.BankNumber.Set(&v)
}
// SetBankNumberNil sets the value for BankNumber to be an explicit nil
func (o *ModelContactUpdate) SetBankNumberNil() {
	o.BankNumber.Set(nil)
}

// UnsetBankNumber ensures that no value is present for BankNumber, not even an explicit nil
func (o *ModelContactUpdate) UnsetBankNumber() {
	o.BankNumber.Unset()
}

// GetDefaultCashbackTime returns the DefaultCashbackTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelContactUpdate) GetDefaultCashbackTime() int32 {
	if o == nil || IsNil(o.DefaultCashbackTime.Get()) {
		var ret int32
		return ret
	}
	return *o.DefaultCashbackTime.Get()
}

// GetDefaultCashbackTimeOk returns a tuple with the DefaultCashbackTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContactUpdate) GetDefaultCashbackTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultCashbackTime.Get(), o.DefaultCashbackTime.IsSet()
}

// HasDefaultCashbackTime returns a boolean if a field has been set.
func (o *ModelContactUpdate) HasDefaultCashbackTime() bool {
	if o != nil && o.DefaultCashbackTime.IsSet() {
		return true
	}

	return false
}

// SetDefaultCashbackTime gets a reference to the given NullableInt32 and assigns it to the DefaultCashbackTime field.
func (o *ModelContactUpdate) SetDefaultCashbackTime(v int32) {
	o.DefaultCashbackTime.Set(&v)
}
// SetDefaultCashbackTimeNil sets the value for DefaultCashbackTime to be an explicit nil
func (o *ModelContactUpdate) SetDefaultCashbackTimeNil() {
	o.DefaultCashbackTime.Set(nil)
}

// UnsetDefaultCashbackTime ensures that no value is present for DefaultCashbackTime, not even an explicit nil
func (o *ModelContactUpdate) UnsetDefaultCashbackTime() {
	o.DefaultCashbackTime.Unset()
}

// GetDefaultCashbackPercent returns the DefaultCashbackPercent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelContactUpdate) GetDefaultCashbackPercent() float32 {
	if o == nil || IsNil(o.DefaultCashbackPercent.Get()) {
		var ret float32
		return ret
	}
	return *o.DefaultCashbackPercent.Get()
}

// GetDefaultCashbackPercentOk returns a tuple with the DefaultCashbackPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContactUpdate) GetDefaultCashbackPercentOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultCashbackPercent.Get(), o.DefaultCashbackPercent.IsSet()
}

// HasDefaultCashbackPercent returns a boolean if a field has been set.
func (o *ModelContactUpdate) HasDefaultCashbackPercent() bool {
	if o != nil && o.DefaultCashbackPercent.IsSet() {
		return true
	}

	return false
}

// SetDefaultCashbackPercent gets a reference to the given NullableFloat32 and assigns it to the DefaultCashbackPercent field.
func (o *ModelContactUpdate) SetDefaultCashbackPercent(v float32) {
	o.DefaultCashbackPercent.Set(&v)
}
// SetDefaultCashbackPercentNil sets the value for DefaultCashbackPercent to be an explicit nil
func (o *ModelContactUpdate) SetDefaultCashbackPercentNil() {
	o.DefaultCashbackPercent.Set(nil)
}

// UnsetDefaultCashbackPercent ensures that no value is present for DefaultCashbackPercent, not even an explicit nil
func (o *ModelContactUpdate) UnsetDefaultCashbackPercent() {
	o.DefaultCashbackPercent.Unset()
}

// GetDefaultTimeToPay returns the DefaultTimeToPay field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelContactUpdate) GetDefaultTimeToPay() int32 {
	if o == nil || IsNil(o.DefaultTimeToPay.Get()) {
		var ret int32
		return ret
	}
	return *o.DefaultTimeToPay.Get()
}

// GetDefaultTimeToPayOk returns a tuple with the DefaultTimeToPay field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContactUpdate) GetDefaultTimeToPayOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultTimeToPay.Get(), o.DefaultTimeToPay.IsSet()
}

// HasDefaultTimeToPay returns a boolean if a field has been set.
func (o *ModelContactUpdate) HasDefaultTimeToPay() bool {
	if o != nil && o.DefaultTimeToPay.IsSet() {
		return true
	}

	return false
}

// SetDefaultTimeToPay gets a reference to the given NullableInt32 and assigns it to the DefaultTimeToPay field.
func (o *ModelContactUpdate) SetDefaultTimeToPay(v int32) {
	o.DefaultTimeToPay.Set(&v)
}
// SetDefaultTimeToPayNil sets the value for DefaultTimeToPay to be an explicit nil
func (o *ModelContactUpdate) SetDefaultTimeToPayNil() {
	o.DefaultTimeToPay.Set(nil)
}

// UnsetDefaultTimeToPay ensures that no value is present for DefaultTimeToPay, not even an explicit nil
func (o *ModelContactUpdate) UnsetDefaultTimeToPay() {
	o.DefaultTimeToPay.Unset()
}

// GetTaxNumber returns the TaxNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelContactUpdate) GetTaxNumber() string {
	if o == nil || IsNil(o.TaxNumber.Get()) {
		var ret string
		return ret
	}
	return *o.TaxNumber.Get()
}

// GetTaxNumberOk returns a tuple with the TaxNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContactUpdate) GetTaxNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxNumber.Get(), o.TaxNumber.IsSet()
}

// HasTaxNumber returns a boolean if a field has been set.
func (o *ModelContactUpdate) HasTaxNumber() bool {
	if o != nil && o.TaxNumber.IsSet() {
		return true
	}

	return false
}

// SetTaxNumber gets a reference to the given NullableString and assigns it to the TaxNumber field.
func (o *ModelContactUpdate) SetTaxNumber(v string) {
	o.TaxNumber.Set(&v)
}
// SetTaxNumberNil sets the value for TaxNumber to be an explicit nil
func (o *ModelContactUpdate) SetTaxNumberNil() {
	o.TaxNumber.Set(nil)
}

// UnsetTaxNumber ensures that no value is present for TaxNumber, not even an explicit nil
func (o *ModelContactUpdate) UnsetTaxNumber() {
	o.TaxNumber.Unset()
}

// GetTaxOffice returns the TaxOffice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelContactUpdate) GetTaxOffice() string {
	if o == nil || IsNil(o.TaxOffice.Get()) {
		var ret string
		return ret
	}
	return *o.TaxOffice.Get()
}

// GetTaxOfficeOk returns a tuple with the TaxOffice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContactUpdate) GetTaxOfficeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxOffice.Get(), o.TaxOffice.IsSet()
}

// HasTaxOffice returns a boolean if a field has been set.
func (o *ModelContactUpdate) HasTaxOffice() bool {
	if o != nil && o.TaxOffice.IsSet() {
		return true
	}

	return false
}

// SetTaxOffice gets a reference to the given NullableString and assigns it to the TaxOffice field.
func (o *ModelContactUpdate) SetTaxOffice(v string) {
	o.TaxOffice.Set(&v)
}
// SetTaxOfficeNil sets the value for TaxOffice to be an explicit nil
func (o *ModelContactUpdate) SetTaxOfficeNil() {
	o.TaxOffice.Set(nil)
}

// UnsetTaxOffice ensures that no value is present for TaxOffice, not even an explicit nil
func (o *ModelContactUpdate) UnsetTaxOffice() {
	o.TaxOffice.Unset()
}

// GetExemptVat returns the ExemptVat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelContactUpdate) GetExemptVat() bool {
	if o == nil || IsNil(o.ExemptVat.Get()) {
		var ret bool
		return ret
	}
	return *o.ExemptVat.Get()
}

// GetExemptVatOk returns a tuple with the ExemptVat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContactUpdate) GetExemptVatOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExemptVat.Get(), o.ExemptVat.IsSet()
}

// HasExemptVat returns a boolean if a field has been set.
func (o *ModelContactUpdate) HasExemptVat() bool {
	if o != nil && o.ExemptVat.IsSet() {
		return true
	}

	return false
}

// SetExemptVat gets a reference to the given NullableBool and assigns it to the ExemptVat field.
func (o *ModelContactUpdate) SetExemptVat(v bool) {
	o.ExemptVat.Set(&v)
}
// SetExemptVatNil sets the value for ExemptVat to be an explicit nil
func (o *ModelContactUpdate) SetExemptVatNil() {
	o.ExemptVat.Set(nil)
}

// UnsetExemptVat ensures that no value is present for ExemptVat, not even an explicit nil
func (o *ModelContactUpdate) UnsetExemptVat() {
	o.ExemptVat.Unset()
}

// GetTaxType returns the TaxType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelContactUpdate) GetTaxType() string {
	if o == nil || IsNil(o.TaxType.Get()) {
		var ret string
		return ret
	}
	return *o.TaxType.Get()
}

// GetTaxTypeOk returns a tuple with the TaxType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContactUpdate) GetTaxTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxType.Get(), o.TaxType.IsSet()
}

// HasTaxType returns a boolean if a field has been set.
func (o *ModelContactUpdate) HasTaxType() bool {
	if o != nil && o.TaxType.IsSet() {
		return true
	}

	return false
}

// SetTaxType gets a reference to the given NullableString and assigns it to the TaxType field.
func (o *ModelContactUpdate) SetTaxType(v string) {
	o.TaxType.Set(&v)
}
// SetTaxTypeNil sets the value for TaxType to be an explicit nil
func (o *ModelContactUpdate) SetTaxTypeNil() {
	o.TaxType.Set(nil)
}

// UnsetTaxType ensures that no value is present for TaxType, not even an explicit nil
func (o *ModelContactUpdate) UnsetTaxType() {
	o.TaxType.Unset()
}

// GetTaxSet returns the TaxSet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelContactUpdate) GetTaxSet() ModelContactTaxSet {
	if o == nil || IsNil(o.TaxSet.Get()) {
		var ret ModelContactTaxSet
		return ret
	}
	return *o.TaxSet.Get()
}

// GetTaxSetOk returns a tuple with the TaxSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContactUpdate) GetTaxSetOk() (*ModelContactTaxSet, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxSet.Get(), o.TaxSet.IsSet()
}

// HasTaxSet returns a boolean if a field has been set.
func (o *ModelContactUpdate) HasTaxSet() bool {
	if o != nil && o.TaxSet.IsSet() {
		return true
	}

	return false
}

// SetTaxSet gets a reference to the given NullableModelContactTaxSet and assigns it to the TaxSet field.
func (o *ModelContactUpdate) SetTaxSet(v ModelContactTaxSet) {
	o.TaxSet.Set(&v)
}
// SetTaxSetNil sets the value for TaxSet to be an explicit nil
func (o *ModelContactUpdate) SetTaxSetNil() {
	o.TaxSet.Set(nil)
}

// UnsetTaxSet ensures that no value is present for TaxSet, not even an explicit nil
func (o *ModelContactUpdate) UnsetTaxSet() {
	o.TaxSet.Unset()
}

// GetDefaultDiscountAmount returns the DefaultDiscountAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelContactUpdate) GetDefaultDiscountAmount() float32 {
	if o == nil || IsNil(o.DefaultDiscountAmount.Get()) {
		var ret float32
		return ret
	}
	return *o.DefaultDiscountAmount.Get()
}

// GetDefaultDiscountAmountOk returns a tuple with the DefaultDiscountAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContactUpdate) GetDefaultDiscountAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultDiscountAmount.Get(), o.DefaultDiscountAmount.IsSet()
}

// HasDefaultDiscountAmount returns a boolean if a field has been set.
func (o *ModelContactUpdate) HasDefaultDiscountAmount() bool {
	if o != nil && o.DefaultDiscountAmount.IsSet() {
		return true
	}

	return false
}

// SetDefaultDiscountAmount gets a reference to the given NullableFloat32 and assigns it to the DefaultDiscountAmount field.
func (o *ModelContactUpdate) SetDefaultDiscountAmount(v float32) {
	o.DefaultDiscountAmount.Set(&v)
}
// SetDefaultDiscountAmountNil sets the value for DefaultDiscountAmount to be an explicit nil
func (o *ModelContactUpdate) SetDefaultDiscountAmountNil() {
	o.DefaultDiscountAmount.Set(nil)
}

// UnsetDefaultDiscountAmount ensures that no value is present for DefaultDiscountAmount, not even an explicit nil
func (o *ModelContactUpdate) UnsetDefaultDiscountAmount() {
	o.DefaultDiscountAmount.Unset()
}

// GetDefaultDiscountPercentage returns the DefaultDiscountPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelContactUpdate) GetDefaultDiscountPercentage() bool {
	if o == nil || IsNil(o.DefaultDiscountPercentage.Get()) {
		var ret bool
		return ret
	}
	return *o.DefaultDiscountPercentage.Get()
}

// GetDefaultDiscountPercentageOk returns a tuple with the DefaultDiscountPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContactUpdate) GetDefaultDiscountPercentageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultDiscountPercentage.Get(), o.DefaultDiscountPercentage.IsSet()
}

// HasDefaultDiscountPercentage returns a boolean if a field has been set.
func (o *ModelContactUpdate) HasDefaultDiscountPercentage() bool {
	if o != nil && o.DefaultDiscountPercentage.IsSet() {
		return true
	}

	return false
}

// SetDefaultDiscountPercentage gets a reference to the given NullableBool and assigns it to the DefaultDiscountPercentage field.
func (o *ModelContactUpdate) SetDefaultDiscountPercentage(v bool) {
	o.DefaultDiscountPercentage.Set(&v)
}
// SetDefaultDiscountPercentageNil sets the value for DefaultDiscountPercentage to be an explicit nil
func (o *ModelContactUpdate) SetDefaultDiscountPercentageNil() {
	o.DefaultDiscountPercentage.Set(nil)
}

// UnsetDefaultDiscountPercentage ensures that no value is present for DefaultDiscountPercentage, not even an explicit nil
func (o *ModelContactUpdate) UnsetDefaultDiscountPercentage() {
	o.DefaultDiscountPercentage.Unset()
}

// GetBuyerReference returns the BuyerReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelContactUpdate) GetBuyerReference() string {
	if o == nil || IsNil(o.BuyerReference.Get()) {
		var ret string
		return ret
	}
	return *o.BuyerReference.Get()
}

// GetBuyerReferenceOk returns a tuple with the BuyerReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContactUpdate) GetBuyerReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BuyerReference.Get(), o.BuyerReference.IsSet()
}

// HasBuyerReference returns a boolean if a field has been set.
func (o *ModelContactUpdate) HasBuyerReference() bool {
	if o != nil && o.BuyerReference.IsSet() {
		return true
	}

	return false
}

// SetBuyerReference gets a reference to the given NullableString and assigns it to the BuyerReference field.
func (o *ModelContactUpdate) SetBuyerReference(v string) {
	o.BuyerReference.Set(&v)
}
// SetBuyerReferenceNil sets the value for BuyerReference to be an explicit nil
func (o *ModelContactUpdate) SetBuyerReferenceNil() {
	o.BuyerReference.Set(nil)
}

// UnsetBuyerReference ensures that no value is present for BuyerReference, not even an explicit nil
func (o *ModelContactUpdate) UnsetBuyerReference() {
	o.BuyerReference.Unset()
}

// GetGovernmentAgency returns the GovernmentAgency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelContactUpdate) GetGovernmentAgency() bool {
	if o == nil || IsNil(o.GovernmentAgency.Get()) {
		var ret bool
		return ret
	}
	return *o.GovernmentAgency.Get()
}

// GetGovernmentAgencyOk returns a tuple with the GovernmentAgency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContactUpdate) GetGovernmentAgencyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.GovernmentAgency.Get(), o.GovernmentAgency.IsSet()
}

// HasGovernmentAgency returns a boolean if a field has been set.
func (o *ModelContactUpdate) HasGovernmentAgency() bool {
	if o != nil && o.GovernmentAgency.IsSet() {
		return true
	}

	return false
}

// SetGovernmentAgency gets a reference to the given NullableBool and assigns it to the GovernmentAgency field.
func (o *ModelContactUpdate) SetGovernmentAgency(v bool) {
	o.GovernmentAgency.Set(&v)
}
// SetGovernmentAgencyNil sets the value for GovernmentAgency to be an explicit nil
func (o *ModelContactUpdate) SetGovernmentAgencyNil() {
	o.GovernmentAgency.Set(nil)
}

// UnsetGovernmentAgency ensures that no value is present for GovernmentAgency, not even an explicit nil
func (o *ModelContactUpdate) UnsetGovernmentAgency() {
	o.GovernmentAgency.Unset()
}

func (o ModelContactUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelContactUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.CustomerNumber.IsSet() {
		toSerialize["customerNumber"] = o.CustomerNumber.Get()
	}
	if o.Parent.IsSet() {
		toSerialize["parent"] = o.Parent.Get()
	}
	if o.Surename.IsSet() {
		toSerialize["surename"] = o.Surename.Get()
	}
	if o.Familyname.IsSet() {
		toSerialize["familyname"] = o.Familyname.Get()
	}
	if o.Titel.IsSet() {
		toSerialize["titel"] = o.Titel.Get()
	}
	if o.Category.IsSet() {
		toSerialize["category"] = o.Category.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.AcademicTitle.IsSet() {
		toSerialize["academicTitle"] = o.AcademicTitle.Get()
	}
	if o.Gender.IsSet() {
		toSerialize["gender"] = o.Gender.Get()
	}
	if o.Name2.IsSet() {
		toSerialize["name2"] = o.Name2.Get()
	}
	if o.Birthday.IsSet() {
		toSerialize["birthday"] = o.Birthday.Get()
	}
	if o.VatNumber.IsSet() {
		toSerialize["vatNumber"] = o.VatNumber.Get()
	}
	if o.BankAccount.IsSet() {
		toSerialize["bankAccount"] = o.BankAccount.Get()
	}
	if o.BankNumber.IsSet() {
		toSerialize["bankNumber"] = o.BankNumber.Get()
	}
	if o.DefaultCashbackTime.IsSet() {
		toSerialize["defaultCashbackTime"] = o.DefaultCashbackTime.Get()
	}
	if o.DefaultCashbackPercent.IsSet() {
		toSerialize["defaultCashbackPercent"] = o.DefaultCashbackPercent.Get()
	}
	if o.DefaultTimeToPay.IsSet() {
		toSerialize["defaultTimeToPay"] = o.DefaultTimeToPay.Get()
	}
	if o.TaxNumber.IsSet() {
		toSerialize["taxNumber"] = o.TaxNumber.Get()
	}
	if o.TaxOffice.IsSet() {
		toSerialize["taxOffice"] = o.TaxOffice.Get()
	}
	if o.ExemptVat.IsSet() {
		toSerialize["exemptVat"] = o.ExemptVat.Get()
	}
	if o.TaxType.IsSet() {
		toSerialize["taxType"] = o.TaxType.Get()
	}
	if o.TaxSet.IsSet() {
		toSerialize["taxSet"] = o.TaxSet.Get()
	}
	if o.DefaultDiscountAmount.IsSet() {
		toSerialize["defaultDiscountAmount"] = o.DefaultDiscountAmount.Get()
	}
	if o.DefaultDiscountPercentage.IsSet() {
		toSerialize["defaultDiscountPercentage"] = o.DefaultDiscountPercentage.Get()
	}
	if o.BuyerReference.IsSet() {
		toSerialize["buyerReference"] = o.BuyerReference.Get()
	}
	if o.GovernmentAgency.IsSet() {
		toSerialize["governmentAgency"] = o.GovernmentAgency.Get()
	}
	return toSerialize, nil
}

type NullableModelContactUpdate struct {
	value *ModelContactUpdate
	isSet bool
}

func (v NullableModelContactUpdate) Get() *ModelContactUpdate {
	return v.value
}

func (v *NullableModelContactUpdate) Set(val *ModelContactUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableModelContactUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableModelContactUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelContactUpdate(val *ModelContactUpdate) *NullableModelContactUpdate {
	return &NullableModelContactUpdate{value: val, isSet: true}
}

func (v NullableModelContactUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelContactUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


