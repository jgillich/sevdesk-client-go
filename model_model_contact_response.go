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

// checks if the ModelContactResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelContactResponse{}

// ModelContactResponse Contact model
type ModelContactResponse struct {
	// The contact id
	Id *string `json:"id,omitempty"`
	// The contact object name
	ObjectName *string `json:"objectName,omitempty"`
	// Date of contact creation
	Create *time.Time `json:"create,omitempty"`
	// Date of last contact update
	Update *time.Time `json:"update,omitempty"`
	// The organization name.<br> Be aware that the type of contact will depend on this attribute.<br> If it holds a value, the contact will be regarded as an organization.
	Name *string `json:"name,omitempty"`
	// Defines the status of the contact. 100 <-> Lead - 500 <-> Pending - 1000 <-> Active.
	Status *string `json:"status,omitempty"`
	// The customer number
	CustomerNumber *string `json:"customerNumber,omitempty"`
	Parent *ModelContactResponseParent `json:"parent,omitempty"`
	// The <b>first</b> name of the contact.<br> Yeah... not quite right in literally every way. We know.<br> Not to be used for organizations.
	Surename *string `json:"surename,omitempty"`
	// The last name of the contact.<br> Not to be used for organizations.
	Familyname *string `json:"familyname,omitempty"`
	// A non-academic title for the contact. Not to be used for organizations.
	Titel *string `json:"titel,omitempty"`
	Category *ModelContactResponseCategory `json:"category,omitempty"`
	// A description for the contact.
	Description *string `json:"description,omitempty"`
	// A academic title for the contact. Not to be used for organizations.
	AcademicTitle *string `json:"academicTitle,omitempty"`
	// Gender of the contact.<br> Not to be used for organizations.
	Gender *string `json:"gender,omitempty"`
	SevClient *ModelContactResponseSevClient `json:"sevClient,omitempty"`
	// Second name of the contact.<br> Not to be used for organizations.
	Name2 *string `json:"name2,omitempty"`
	// Birthday of the contact.<br> Not to be used for organizations.
	Birthday *string `json:"birthday,omitempty"`
	// Vat number of the contact.
	VatNumber *string `json:"vatNumber,omitempty"`
	// Bank account number (IBAN) of the contact.
	BankAccount *string `json:"bankAccount,omitempty"`
	// Bank number of the bank used by the contact.
	BankNumber *string `json:"bankNumber,omitempty"`
	// Absolute time in days which the contact has to pay his invoices and subsequently get a cashback.
	DefaultCashbackTime *string `json:"defaultCashbackTime,omitempty"`
	// Percentage of the invoice sum the contact gets back if he payed invoices in time.
	DefaultCashbackPercent *float32 `json:"defaultCashbackPercent,omitempty"`
	// The payment goal in days which is set for every invoice of the contact.
	DefaultTimeToPay *string `json:"defaultTimeToPay,omitempty"`
	// The tax number of the contact.
	TaxNumber *string `json:"taxNumber,omitempty"`
	// The tax office of the contact (only for greek customers).
	TaxOffice *string `json:"taxOffice,omitempty"`
	// Defines if the contact is freed from paying vat.
	ExemptVat *string `json:"exemptVat,omitempty"`
	// Defines which tax regulation the contact is using.
	TaxType *string `json:"taxType,omitempty"`
	TaxSet *ModelContactResponseTaxSet `json:"taxSet,omitempty"`
	// The default discount the contact gets for every invoice.<br> Depending on defaultDiscountPercentage attribute, in percent or absolute value.
	DefaultDiscountAmount *float32 `json:"defaultDiscountAmount,omitempty"`
	// Defines if the discount is a percentage (true) or an absolute value (false).
	DefaultDiscountPercentage *string `json:"defaultDiscountPercentage,omitempty"`
	// Buyer reference of the contact.
	BuyerReference *string `json:"buyerReference,omitempty"`
	// Defines whether the contact is a government agency (true) or not (false).
	GovernmentAgency *string `json:"governmentAgency,omitempty"`
	// Additional information stored for the contact.
	// Deprecated
	AdditionalInformation *string `json:"additionalInformation,omitempty"`
}

// NewModelContactResponse instantiates a new ModelContactResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelContactResponse() *ModelContactResponse {
	this := ModelContactResponse{}
	return &this
}

// NewModelContactResponseWithDefaults instantiates a new ModelContactResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelContactResponseWithDefaults() *ModelContactResponse {
	this := ModelContactResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelContactResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContactResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelContactResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModelContactResponse) SetId(v string) {
	o.Id = &v
}

// GetObjectName returns the ObjectName field value if set, zero value otherwise.
func (o *ModelContactResponse) GetObjectName() string {
	if o == nil || IsNil(o.ObjectName) {
		var ret string
		return ret
	}
	return *o.ObjectName
}

// GetObjectNameOk returns a tuple with the ObjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContactResponse) GetObjectNameOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectName) {
		return nil, false
	}
	return o.ObjectName, true
}

// HasObjectName returns a boolean if a field has been set.
func (o *ModelContactResponse) HasObjectName() bool {
	if o != nil && !IsNil(o.ObjectName) {
		return true
	}

	return false
}

// SetObjectName gets a reference to the given string and assigns it to the ObjectName field.
func (o *ModelContactResponse) SetObjectName(v string) {
	o.ObjectName = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *ModelContactResponse) GetCreate() time.Time {
	if o == nil || IsNil(o.Create) {
		var ret time.Time
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContactResponse) GetCreateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *ModelContactResponse) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given time.Time and assigns it to the Create field.
func (o *ModelContactResponse) SetCreate(v time.Time) {
	o.Create = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *ModelContactResponse) GetUpdate() time.Time {
	if o == nil || IsNil(o.Update) {
		var ret time.Time
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContactResponse) GetUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *ModelContactResponse) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given time.Time and assigns it to the Update field.
func (o *ModelContactResponse) SetUpdate(v time.Time) {
	o.Update = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelContactResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContactResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelContactResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelContactResponse) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ModelContactResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContactResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ModelContactResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ModelContactResponse) SetStatus(v string) {
	o.Status = &v
}

// GetCustomerNumber returns the CustomerNumber field value if set, zero value otherwise.
func (o *ModelContactResponse) GetCustomerNumber() string {
	if o == nil || IsNil(o.CustomerNumber) {
		var ret string
		return ret
	}
	return *o.CustomerNumber
}

// GetCustomerNumberOk returns a tuple with the CustomerNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContactResponse) GetCustomerNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerNumber) {
		return nil, false
	}
	return o.CustomerNumber, true
}

// HasCustomerNumber returns a boolean if a field has been set.
func (o *ModelContactResponse) HasCustomerNumber() bool {
	if o != nil && !IsNil(o.CustomerNumber) {
		return true
	}

	return false
}

// SetCustomerNumber gets a reference to the given string and assigns it to the CustomerNumber field.
func (o *ModelContactResponse) SetCustomerNumber(v string) {
	o.CustomerNumber = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *ModelContactResponse) GetParent() ModelContactResponseParent {
	if o == nil || IsNil(o.Parent) {
		var ret ModelContactResponseParent
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContactResponse) GetParentOk() (*ModelContactResponseParent, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *ModelContactResponse) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given ModelContactResponseParent and assigns it to the Parent field.
func (o *ModelContactResponse) SetParent(v ModelContactResponseParent) {
	o.Parent = &v
}

// GetSurename returns the Surename field value if set, zero value otherwise.
func (o *ModelContactResponse) GetSurename() string {
	if o == nil || IsNil(o.Surename) {
		var ret string
		return ret
	}
	return *o.Surename
}

// GetSurenameOk returns a tuple with the Surename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContactResponse) GetSurenameOk() (*string, bool) {
	if o == nil || IsNil(o.Surename) {
		return nil, false
	}
	return o.Surename, true
}

// HasSurename returns a boolean if a field has been set.
func (o *ModelContactResponse) HasSurename() bool {
	if o != nil && !IsNil(o.Surename) {
		return true
	}

	return false
}

// SetSurename gets a reference to the given string and assigns it to the Surename field.
func (o *ModelContactResponse) SetSurename(v string) {
	o.Surename = &v
}

// GetFamilyname returns the Familyname field value if set, zero value otherwise.
func (o *ModelContactResponse) GetFamilyname() string {
	if o == nil || IsNil(o.Familyname) {
		var ret string
		return ret
	}
	return *o.Familyname
}

// GetFamilynameOk returns a tuple with the Familyname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContactResponse) GetFamilynameOk() (*string, bool) {
	if o == nil || IsNil(o.Familyname) {
		return nil, false
	}
	return o.Familyname, true
}

// HasFamilyname returns a boolean if a field has been set.
func (o *ModelContactResponse) HasFamilyname() bool {
	if o != nil && !IsNil(o.Familyname) {
		return true
	}

	return false
}

// SetFamilyname gets a reference to the given string and assigns it to the Familyname field.
func (o *ModelContactResponse) SetFamilyname(v string) {
	o.Familyname = &v
}

// GetTitel returns the Titel field value if set, zero value otherwise.
func (o *ModelContactResponse) GetTitel() string {
	if o == nil || IsNil(o.Titel) {
		var ret string
		return ret
	}
	return *o.Titel
}

// GetTitelOk returns a tuple with the Titel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContactResponse) GetTitelOk() (*string, bool) {
	if o == nil || IsNil(o.Titel) {
		return nil, false
	}
	return o.Titel, true
}

// HasTitel returns a boolean if a field has been set.
func (o *ModelContactResponse) HasTitel() bool {
	if o != nil && !IsNil(o.Titel) {
		return true
	}

	return false
}

// SetTitel gets a reference to the given string and assigns it to the Titel field.
func (o *ModelContactResponse) SetTitel(v string) {
	o.Titel = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *ModelContactResponse) GetCategory() ModelContactResponseCategory {
	if o == nil || IsNil(o.Category) {
		var ret ModelContactResponseCategory
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContactResponse) GetCategoryOk() (*ModelContactResponseCategory, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *ModelContactResponse) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given ModelContactResponseCategory and assigns it to the Category field.
func (o *ModelContactResponse) SetCategory(v ModelContactResponseCategory) {
	o.Category = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ModelContactResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContactResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ModelContactResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ModelContactResponse) SetDescription(v string) {
	o.Description = &v
}

// GetAcademicTitle returns the AcademicTitle field value if set, zero value otherwise.
func (o *ModelContactResponse) GetAcademicTitle() string {
	if o == nil || IsNil(o.AcademicTitle) {
		var ret string
		return ret
	}
	return *o.AcademicTitle
}

// GetAcademicTitleOk returns a tuple with the AcademicTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContactResponse) GetAcademicTitleOk() (*string, bool) {
	if o == nil || IsNil(o.AcademicTitle) {
		return nil, false
	}
	return o.AcademicTitle, true
}

// HasAcademicTitle returns a boolean if a field has been set.
func (o *ModelContactResponse) HasAcademicTitle() bool {
	if o != nil && !IsNil(o.AcademicTitle) {
		return true
	}

	return false
}

// SetAcademicTitle gets a reference to the given string and assigns it to the AcademicTitle field.
func (o *ModelContactResponse) SetAcademicTitle(v string) {
	o.AcademicTitle = &v
}

// GetGender returns the Gender field value if set, zero value otherwise.
func (o *ModelContactResponse) GetGender() string {
	if o == nil || IsNil(o.Gender) {
		var ret string
		return ret
	}
	return *o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContactResponse) GetGenderOk() (*string, bool) {
	if o == nil || IsNil(o.Gender) {
		return nil, false
	}
	return o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *ModelContactResponse) HasGender() bool {
	if o != nil && !IsNil(o.Gender) {
		return true
	}

	return false
}

// SetGender gets a reference to the given string and assigns it to the Gender field.
func (o *ModelContactResponse) SetGender(v string) {
	o.Gender = &v
}

// GetSevClient returns the SevClient field value if set, zero value otherwise.
func (o *ModelContactResponse) GetSevClient() ModelContactResponseSevClient {
	if o == nil || IsNil(o.SevClient) {
		var ret ModelContactResponseSevClient
		return ret
	}
	return *o.SevClient
}

// GetSevClientOk returns a tuple with the SevClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContactResponse) GetSevClientOk() (*ModelContactResponseSevClient, bool) {
	if o == nil || IsNil(o.SevClient) {
		return nil, false
	}
	return o.SevClient, true
}

// HasSevClient returns a boolean if a field has been set.
func (o *ModelContactResponse) HasSevClient() bool {
	if o != nil && !IsNil(o.SevClient) {
		return true
	}

	return false
}

// SetSevClient gets a reference to the given ModelContactResponseSevClient and assigns it to the SevClient field.
func (o *ModelContactResponse) SetSevClient(v ModelContactResponseSevClient) {
	o.SevClient = &v
}

// GetName2 returns the Name2 field value if set, zero value otherwise.
func (o *ModelContactResponse) GetName2() string {
	if o == nil || IsNil(o.Name2) {
		var ret string
		return ret
	}
	return *o.Name2
}

// GetName2Ok returns a tuple with the Name2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContactResponse) GetName2Ok() (*string, bool) {
	if o == nil || IsNil(o.Name2) {
		return nil, false
	}
	return o.Name2, true
}

// HasName2 returns a boolean if a field has been set.
func (o *ModelContactResponse) HasName2() bool {
	if o != nil && !IsNil(o.Name2) {
		return true
	}

	return false
}

// SetName2 gets a reference to the given string and assigns it to the Name2 field.
func (o *ModelContactResponse) SetName2(v string) {
	o.Name2 = &v
}

// GetBirthday returns the Birthday field value if set, zero value otherwise.
func (o *ModelContactResponse) GetBirthday() string {
	if o == nil || IsNil(o.Birthday) {
		var ret string
		return ret
	}
	return *o.Birthday
}

// GetBirthdayOk returns a tuple with the Birthday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContactResponse) GetBirthdayOk() (*string, bool) {
	if o == nil || IsNil(o.Birthday) {
		return nil, false
	}
	return o.Birthday, true
}

// HasBirthday returns a boolean if a field has been set.
func (o *ModelContactResponse) HasBirthday() bool {
	if o != nil && !IsNil(o.Birthday) {
		return true
	}

	return false
}

// SetBirthday gets a reference to the given string and assigns it to the Birthday field.
func (o *ModelContactResponse) SetBirthday(v string) {
	o.Birthday = &v
}

// GetVatNumber returns the VatNumber field value if set, zero value otherwise.
func (o *ModelContactResponse) GetVatNumber() string {
	if o == nil || IsNil(o.VatNumber) {
		var ret string
		return ret
	}
	return *o.VatNumber
}

// GetVatNumberOk returns a tuple with the VatNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContactResponse) GetVatNumberOk() (*string, bool) {
	if o == nil || IsNil(o.VatNumber) {
		return nil, false
	}
	return o.VatNumber, true
}

// HasVatNumber returns a boolean if a field has been set.
func (o *ModelContactResponse) HasVatNumber() bool {
	if o != nil && !IsNil(o.VatNumber) {
		return true
	}

	return false
}

// SetVatNumber gets a reference to the given string and assigns it to the VatNumber field.
func (o *ModelContactResponse) SetVatNumber(v string) {
	o.VatNumber = &v
}

// GetBankAccount returns the BankAccount field value if set, zero value otherwise.
func (o *ModelContactResponse) GetBankAccount() string {
	if o == nil || IsNil(o.BankAccount) {
		var ret string
		return ret
	}
	return *o.BankAccount
}

// GetBankAccountOk returns a tuple with the BankAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContactResponse) GetBankAccountOk() (*string, bool) {
	if o == nil || IsNil(o.BankAccount) {
		return nil, false
	}
	return o.BankAccount, true
}

// HasBankAccount returns a boolean if a field has been set.
func (o *ModelContactResponse) HasBankAccount() bool {
	if o != nil && !IsNil(o.BankAccount) {
		return true
	}

	return false
}

// SetBankAccount gets a reference to the given string and assigns it to the BankAccount field.
func (o *ModelContactResponse) SetBankAccount(v string) {
	o.BankAccount = &v
}

// GetBankNumber returns the BankNumber field value if set, zero value otherwise.
func (o *ModelContactResponse) GetBankNumber() string {
	if o == nil || IsNil(o.BankNumber) {
		var ret string
		return ret
	}
	return *o.BankNumber
}

// GetBankNumberOk returns a tuple with the BankNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContactResponse) GetBankNumberOk() (*string, bool) {
	if o == nil || IsNil(o.BankNumber) {
		return nil, false
	}
	return o.BankNumber, true
}

// HasBankNumber returns a boolean if a field has been set.
func (o *ModelContactResponse) HasBankNumber() bool {
	if o != nil && !IsNil(o.BankNumber) {
		return true
	}

	return false
}

// SetBankNumber gets a reference to the given string and assigns it to the BankNumber field.
func (o *ModelContactResponse) SetBankNumber(v string) {
	o.BankNumber = &v
}

// GetDefaultCashbackTime returns the DefaultCashbackTime field value if set, zero value otherwise.
func (o *ModelContactResponse) GetDefaultCashbackTime() string {
	if o == nil || IsNil(o.DefaultCashbackTime) {
		var ret string
		return ret
	}
	return *o.DefaultCashbackTime
}

// GetDefaultCashbackTimeOk returns a tuple with the DefaultCashbackTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContactResponse) GetDefaultCashbackTimeOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultCashbackTime) {
		return nil, false
	}
	return o.DefaultCashbackTime, true
}

// HasDefaultCashbackTime returns a boolean if a field has been set.
func (o *ModelContactResponse) HasDefaultCashbackTime() bool {
	if o != nil && !IsNil(o.DefaultCashbackTime) {
		return true
	}

	return false
}

// SetDefaultCashbackTime gets a reference to the given string and assigns it to the DefaultCashbackTime field.
func (o *ModelContactResponse) SetDefaultCashbackTime(v string) {
	o.DefaultCashbackTime = &v
}

// GetDefaultCashbackPercent returns the DefaultCashbackPercent field value if set, zero value otherwise.
func (o *ModelContactResponse) GetDefaultCashbackPercent() float32 {
	if o == nil || IsNil(o.DefaultCashbackPercent) {
		var ret float32
		return ret
	}
	return *o.DefaultCashbackPercent
}

// GetDefaultCashbackPercentOk returns a tuple with the DefaultCashbackPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContactResponse) GetDefaultCashbackPercentOk() (*float32, bool) {
	if o == nil || IsNil(o.DefaultCashbackPercent) {
		return nil, false
	}
	return o.DefaultCashbackPercent, true
}

// HasDefaultCashbackPercent returns a boolean if a field has been set.
func (o *ModelContactResponse) HasDefaultCashbackPercent() bool {
	if o != nil && !IsNil(o.DefaultCashbackPercent) {
		return true
	}

	return false
}

// SetDefaultCashbackPercent gets a reference to the given float32 and assigns it to the DefaultCashbackPercent field.
func (o *ModelContactResponse) SetDefaultCashbackPercent(v float32) {
	o.DefaultCashbackPercent = &v
}

// GetDefaultTimeToPay returns the DefaultTimeToPay field value if set, zero value otherwise.
func (o *ModelContactResponse) GetDefaultTimeToPay() string {
	if o == nil || IsNil(o.DefaultTimeToPay) {
		var ret string
		return ret
	}
	return *o.DefaultTimeToPay
}

// GetDefaultTimeToPayOk returns a tuple with the DefaultTimeToPay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContactResponse) GetDefaultTimeToPayOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultTimeToPay) {
		return nil, false
	}
	return o.DefaultTimeToPay, true
}

// HasDefaultTimeToPay returns a boolean if a field has been set.
func (o *ModelContactResponse) HasDefaultTimeToPay() bool {
	if o != nil && !IsNil(o.DefaultTimeToPay) {
		return true
	}

	return false
}

// SetDefaultTimeToPay gets a reference to the given string and assigns it to the DefaultTimeToPay field.
func (o *ModelContactResponse) SetDefaultTimeToPay(v string) {
	o.DefaultTimeToPay = &v
}

// GetTaxNumber returns the TaxNumber field value if set, zero value otherwise.
func (o *ModelContactResponse) GetTaxNumber() string {
	if o == nil || IsNil(o.TaxNumber) {
		var ret string
		return ret
	}
	return *o.TaxNumber
}

// GetTaxNumberOk returns a tuple with the TaxNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContactResponse) GetTaxNumberOk() (*string, bool) {
	if o == nil || IsNil(o.TaxNumber) {
		return nil, false
	}
	return o.TaxNumber, true
}

// HasTaxNumber returns a boolean if a field has been set.
func (o *ModelContactResponse) HasTaxNumber() bool {
	if o != nil && !IsNil(o.TaxNumber) {
		return true
	}

	return false
}

// SetTaxNumber gets a reference to the given string and assigns it to the TaxNumber field.
func (o *ModelContactResponse) SetTaxNumber(v string) {
	o.TaxNumber = &v
}

// GetTaxOffice returns the TaxOffice field value if set, zero value otherwise.
func (o *ModelContactResponse) GetTaxOffice() string {
	if o == nil || IsNil(o.TaxOffice) {
		var ret string
		return ret
	}
	return *o.TaxOffice
}

// GetTaxOfficeOk returns a tuple with the TaxOffice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContactResponse) GetTaxOfficeOk() (*string, bool) {
	if o == nil || IsNil(o.TaxOffice) {
		return nil, false
	}
	return o.TaxOffice, true
}

// HasTaxOffice returns a boolean if a field has been set.
func (o *ModelContactResponse) HasTaxOffice() bool {
	if o != nil && !IsNil(o.TaxOffice) {
		return true
	}

	return false
}

// SetTaxOffice gets a reference to the given string and assigns it to the TaxOffice field.
func (o *ModelContactResponse) SetTaxOffice(v string) {
	o.TaxOffice = &v
}

// GetExemptVat returns the ExemptVat field value if set, zero value otherwise.
func (o *ModelContactResponse) GetExemptVat() string {
	if o == nil || IsNil(o.ExemptVat) {
		var ret string
		return ret
	}
	return *o.ExemptVat
}

// GetExemptVatOk returns a tuple with the ExemptVat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContactResponse) GetExemptVatOk() (*string, bool) {
	if o == nil || IsNil(o.ExemptVat) {
		return nil, false
	}
	return o.ExemptVat, true
}

// HasExemptVat returns a boolean if a field has been set.
func (o *ModelContactResponse) HasExemptVat() bool {
	if o != nil && !IsNil(o.ExemptVat) {
		return true
	}

	return false
}

// SetExemptVat gets a reference to the given string and assigns it to the ExemptVat field.
func (o *ModelContactResponse) SetExemptVat(v string) {
	o.ExemptVat = &v
}

// GetTaxType returns the TaxType field value if set, zero value otherwise.
func (o *ModelContactResponse) GetTaxType() string {
	if o == nil || IsNil(o.TaxType) {
		var ret string
		return ret
	}
	return *o.TaxType
}

// GetTaxTypeOk returns a tuple with the TaxType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContactResponse) GetTaxTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TaxType) {
		return nil, false
	}
	return o.TaxType, true
}

// HasTaxType returns a boolean if a field has been set.
func (o *ModelContactResponse) HasTaxType() bool {
	if o != nil && !IsNil(o.TaxType) {
		return true
	}

	return false
}

// SetTaxType gets a reference to the given string and assigns it to the TaxType field.
func (o *ModelContactResponse) SetTaxType(v string) {
	o.TaxType = &v
}

// GetTaxSet returns the TaxSet field value if set, zero value otherwise.
func (o *ModelContactResponse) GetTaxSet() ModelContactResponseTaxSet {
	if o == nil || IsNil(o.TaxSet) {
		var ret ModelContactResponseTaxSet
		return ret
	}
	return *o.TaxSet
}

// GetTaxSetOk returns a tuple with the TaxSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContactResponse) GetTaxSetOk() (*ModelContactResponseTaxSet, bool) {
	if o == nil || IsNil(o.TaxSet) {
		return nil, false
	}
	return o.TaxSet, true
}

// HasTaxSet returns a boolean if a field has been set.
func (o *ModelContactResponse) HasTaxSet() bool {
	if o != nil && !IsNil(o.TaxSet) {
		return true
	}

	return false
}

// SetTaxSet gets a reference to the given ModelContactResponseTaxSet and assigns it to the TaxSet field.
func (o *ModelContactResponse) SetTaxSet(v ModelContactResponseTaxSet) {
	o.TaxSet = &v
}

// GetDefaultDiscountAmount returns the DefaultDiscountAmount field value if set, zero value otherwise.
func (o *ModelContactResponse) GetDefaultDiscountAmount() float32 {
	if o == nil || IsNil(o.DefaultDiscountAmount) {
		var ret float32
		return ret
	}
	return *o.DefaultDiscountAmount
}

// GetDefaultDiscountAmountOk returns a tuple with the DefaultDiscountAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContactResponse) GetDefaultDiscountAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.DefaultDiscountAmount) {
		return nil, false
	}
	return o.DefaultDiscountAmount, true
}

// HasDefaultDiscountAmount returns a boolean if a field has been set.
func (o *ModelContactResponse) HasDefaultDiscountAmount() bool {
	if o != nil && !IsNil(o.DefaultDiscountAmount) {
		return true
	}

	return false
}

// SetDefaultDiscountAmount gets a reference to the given float32 and assigns it to the DefaultDiscountAmount field.
func (o *ModelContactResponse) SetDefaultDiscountAmount(v float32) {
	o.DefaultDiscountAmount = &v
}

// GetDefaultDiscountPercentage returns the DefaultDiscountPercentage field value if set, zero value otherwise.
func (o *ModelContactResponse) GetDefaultDiscountPercentage() string {
	if o == nil || IsNil(o.DefaultDiscountPercentage) {
		var ret string
		return ret
	}
	return *o.DefaultDiscountPercentage
}

// GetDefaultDiscountPercentageOk returns a tuple with the DefaultDiscountPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContactResponse) GetDefaultDiscountPercentageOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultDiscountPercentage) {
		return nil, false
	}
	return o.DefaultDiscountPercentage, true
}

// HasDefaultDiscountPercentage returns a boolean if a field has been set.
func (o *ModelContactResponse) HasDefaultDiscountPercentage() bool {
	if o != nil && !IsNil(o.DefaultDiscountPercentage) {
		return true
	}

	return false
}

// SetDefaultDiscountPercentage gets a reference to the given string and assigns it to the DefaultDiscountPercentage field.
func (o *ModelContactResponse) SetDefaultDiscountPercentage(v string) {
	o.DefaultDiscountPercentage = &v
}

// GetBuyerReference returns the BuyerReference field value if set, zero value otherwise.
func (o *ModelContactResponse) GetBuyerReference() string {
	if o == nil || IsNil(o.BuyerReference) {
		var ret string
		return ret
	}
	return *o.BuyerReference
}

// GetBuyerReferenceOk returns a tuple with the BuyerReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContactResponse) GetBuyerReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.BuyerReference) {
		return nil, false
	}
	return o.BuyerReference, true
}

// HasBuyerReference returns a boolean if a field has been set.
func (o *ModelContactResponse) HasBuyerReference() bool {
	if o != nil && !IsNil(o.BuyerReference) {
		return true
	}

	return false
}

// SetBuyerReference gets a reference to the given string and assigns it to the BuyerReference field.
func (o *ModelContactResponse) SetBuyerReference(v string) {
	o.BuyerReference = &v
}

// GetGovernmentAgency returns the GovernmentAgency field value if set, zero value otherwise.
func (o *ModelContactResponse) GetGovernmentAgency() string {
	if o == nil || IsNil(o.GovernmentAgency) {
		var ret string
		return ret
	}
	return *o.GovernmentAgency
}

// GetGovernmentAgencyOk returns a tuple with the GovernmentAgency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContactResponse) GetGovernmentAgencyOk() (*string, bool) {
	if o == nil || IsNil(o.GovernmentAgency) {
		return nil, false
	}
	return o.GovernmentAgency, true
}

// HasGovernmentAgency returns a boolean if a field has been set.
func (o *ModelContactResponse) HasGovernmentAgency() bool {
	if o != nil && !IsNil(o.GovernmentAgency) {
		return true
	}

	return false
}

// SetGovernmentAgency gets a reference to the given string and assigns it to the GovernmentAgency field.
func (o *ModelContactResponse) SetGovernmentAgency(v string) {
	o.GovernmentAgency = &v
}

// GetAdditionalInformation returns the AdditionalInformation field value if set, zero value otherwise.
// Deprecated
func (o *ModelContactResponse) GetAdditionalInformation() string {
	if o == nil || IsNil(o.AdditionalInformation) {
		var ret string
		return ret
	}
	return *o.AdditionalInformation
}

// GetAdditionalInformationOk returns a tuple with the AdditionalInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ModelContactResponse) GetAdditionalInformationOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalInformation) {
		return nil, false
	}
	return o.AdditionalInformation, true
}

// HasAdditionalInformation returns a boolean if a field has been set.
func (o *ModelContactResponse) HasAdditionalInformation() bool {
	if o != nil && !IsNil(o.AdditionalInformation) {
		return true
	}

	return false
}

// SetAdditionalInformation gets a reference to the given string and assigns it to the AdditionalInformation field.
// Deprecated
func (o *ModelContactResponse) SetAdditionalInformation(v string) {
	o.AdditionalInformation = &v
}

func (o ModelContactResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelContactResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	// skip: objectName is readOnly
	// skip: create is readOnly
	// skip: update is readOnly
	// skip: name is readOnly
	// skip: status is readOnly
	// skip: customerNumber is readOnly
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	// skip: surename is readOnly
	// skip: familyname is readOnly
	// skip: titel is readOnly
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	// skip: description is readOnly
	// skip: academicTitle is readOnly
	// skip: gender is readOnly
	if !IsNil(o.SevClient) {
		toSerialize["sevClient"] = o.SevClient
	}
	// skip: name2 is readOnly
	// skip: birthday is readOnly
	// skip: vatNumber is readOnly
	// skip: bankAccount is readOnly
	// skip: bankNumber is readOnly
	// skip: defaultCashbackTime is readOnly
	// skip: defaultCashbackPercent is readOnly
	// skip: defaultTimeToPay is readOnly
	// skip: taxNumber is readOnly
	// skip: taxOffice is readOnly
	// skip: exemptVat is readOnly
	// skip: taxType is readOnly
	if !IsNil(o.TaxSet) {
		toSerialize["taxSet"] = o.TaxSet
	}
	// skip: defaultDiscountAmount is readOnly
	// skip: defaultDiscountPercentage is readOnly
	// skip: buyerReference is readOnly
	// skip: governmentAgency is readOnly
	// skip: additionalInformation is readOnly
	return toSerialize, nil
}

type NullableModelContactResponse struct {
	value *ModelContactResponse
	isSet bool
}

func (v NullableModelContactResponse) Get() *ModelContactResponse {
	return v.value
}

func (v *NullableModelContactResponse) Set(val *ModelContactResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModelContactResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModelContactResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelContactResponse(val *ModelContactResponse) *NullableModelContactResponse {
	return &NullableModelContactResponse{value: val, isSet: true}
}

func (v NullableModelContactResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelContactResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


