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

// checks if the ModelVoucherPos type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelVoucherPos{}

// ModelVoucherPos Voucher position model
type ModelVoucherPos struct {
	// The voucher position id
	Id *int32 `json:"id,omitempty"`
	// The voucher position object name
	ObjectName string `json:"objectName"`
	MapAll bool `json:"mapAll"`
	// Date of voucher position creation
	Create *string `json:"create,omitempty"`
	// Date of last voucher position update
	Update *string `json:"update,omitempty"`
	SevClient *ModelVoucherPosSevClient `json:"sevClient,omitempty"`
	Voucher ModelVoucherPosVoucher `json:"voucher"`
	AccountingType ModelVoucherPosAccountingType `json:"accountingType"`
	EstimatedAccountingType *ModelVoucherPosEstimatedAccountingType `json:"estimatedAccountingType,omitempty"`
	// Tax rate of the voucher position.
	TaxRate float32 `json:"taxRate"`
	// Determines whether 'sumNet' or 'sumGross' is regarded.<br>       If both are not given, 'sum' is regarded and treated as net or gross depending on 'net'.   All positions must be either net or gross, a mixture of the two is not possible.
	Net bool `json:"net"`
	// Determines whether position is regarded as an asset which can be depreciated.
	IsAsset *bool `json:"isAsset,omitempty"`
	// Net sum of the voucher position.<br>      Only regarded if 'net' is 'true', otherwise its readOnly.
	SumNet float32 `json:"sumNet"`
	// Tax sum of the voucher position.
	SumTax *float32 `json:"sumTax,omitempty"`
	// Gross sum of the voucher position.<br>      Only regarded if 'net' is 'false', otherwise its readOnly.
	SumGross float32 `json:"sumGross"`
	// Net accounting sum. Is equal to sumNet.
	SumNetAccounting *float32 `json:"sumNetAccounting,omitempty"`
	// Tax accounting sum. Is equal to sumTax.
	SumTaxAccounting *float32 `json:"sumTaxAccounting,omitempty"`
	// Gross accounting sum. Is equal to sumGross.
	SumGrossAccounting *float32 `json:"sumGrossAccounting,omitempty"`
	// Comment for the voucher position.
	Comment NullableString `json:"comment,omitempty"`
}

// NewModelVoucherPos instantiates a new ModelVoucherPos object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelVoucherPos(objectName string, mapAll bool, voucher ModelVoucherPosVoucher, accountingType ModelVoucherPosAccountingType, taxRate float32, net bool, sumNet float32, sumGross float32) *ModelVoucherPos {
	this := ModelVoucherPos{}
	this.ObjectName = objectName
	this.MapAll = mapAll
	this.Voucher = voucher
	this.AccountingType = accountingType
	this.TaxRate = taxRate
	this.Net = net
	this.SumNet = sumNet
	this.SumGross = sumGross
	return &this
}

// NewModelVoucherPosWithDefaults instantiates a new ModelVoucherPos object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelVoucherPosWithDefaults() *ModelVoucherPos {
	this := ModelVoucherPos{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelVoucherPos) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelVoucherPos) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelVoucherPos) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ModelVoucherPos) SetId(v int32) {
	o.Id = &v
}

// GetObjectName returns the ObjectName field value
func (o *ModelVoucherPos) GetObjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectName
}

// GetObjectNameOk returns a tuple with the ObjectName field value
// and a boolean to check if the value has been set.
func (o *ModelVoucherPos) GetObjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectName, true
}

// SetObjectName sets field value
func (o *ModelVoucherPos) SetObjectName(v string) {
	o.ObjectName = v
}

// GetMapAll returns the MapAll field value
func (o *ModelVoucherPos) GetMapAll() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MapAll
}

// GetMapAllOk returns a tuple with the MapAll field value
// and a boolean to check if the value has been set.
func (o *ModelVoucherPos) GetMapAllOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MapAll, true
}

// SetMapAll sets field value
func (o *ModelVoucherPos) SetMapAll(v bool) {
	o.MapAll = v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *ModelVoucherPos) GetCreate() string {
	if o == nil || IsNil(o.Create) {
		var ret string
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelVoucherPos) GetCreateOk() (*string, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *ModelVoucherPos) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given string and assigns it to the Create field.
func (o *ModelVoucherPos) SetCreate(v string) {
	o.Create = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *ModelVoucherPos) GetUpdate() string {
	if o == nil || IsNil(o.Update) {
		var ret string
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelVoucherPos) GetUpdateOk() (*string, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *ModelVoucherPos) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given string and assigns it to the Update field.
func (o *ModelVoucherPos) SetUpdate(v string) {
	o.Update = &v
}

// GetSevClient returns the SevClient field value if set, zero value otherwise.
func (o *ModelVoucherPos) GetSevClient() ModelVoucherPosSevClient {
	if o == nil || IsNil(o.SevClient) {
		var ret ModelVoucherPosSevClient
		return ret
	}
	return *o.SevClient
}

// GetSevClientOk returns a tuple with the SevClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelVoucherPos) GetSevClientOk() (*ModelVoucherPosSevClient, bool) {
	if o == nil || IsNil(o.SevClient) {
		return nil, false
	}
	return o.SevClient, true
}

// HasSevClient returns a boolean if a field has been set.
func (o *ModelVoucherPos) HasSevClient() bool {
	if o != nil && !IsNil(o.SevClient) {
		return true
	}

	return false
}

// SetSevClient gets a reference to the given ModelVoucherPosSevClient and assigns it to the SevClient field.
func (o *ModelVoucherPos) SetSevClient(v ModelVoucherPosSevClient) {
	o.SevClient = &v
}

// GetVoucher returns the Voucher field value
func (o *ModelVoucherPos) GetVoucher() ModelVoucherPosVoucher {
	if o == nil {
		var ret ModelVoucherPosVoucher
		return ret
	}

	return o.Voucher
}

// GetVoucherOk returns a tuple with the Voucher field value
// and a boolean to check if the value has been set.
func (o *ModelVoucherPos) GetVoucherOk() (*ModelVoucherPosVoucher, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Voucher, true
}

// SetVoucher sets field value
func (o *ModelVoucherPos) SetVoucher(v ModelVoucherPosVoucher) {
	o.Voucher = v
}

// GetAccountingType returns the AccountingType field value
func (o *ModelVoucherPos) GetAccountingType() ModelVoucherPosAccountingType {
	if o == nil {
		var ret ModelVoucherPosAccountingType
		return ret
	}

	return o.AccountingType
}

// GetAccountingTypeOk returns a tuple with the AccountingType field value
// and a boolean to check if the value has been set.
func (o *ModelVoucherPos) GetAccountingTypeOk() (*ModelVoucherPosAccountingType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountingType, true
}

// SetAccountingType sets field value
func (o *ModelVoucherPos) SetAccountingType(v ModelVoucherPosAccountingType) {
	o.AccountingType = v
}

// GetEstimatedAccountingType returns the EstimatedAccountingType field value if set, zero value otherwise.
func (o *ModelVoucherPos) GetEstimatedAccountingType() ModelVoucherPosEstimatedAccountingType {
	if o == nil || IsNil(o.EstimatedAccountingType) {
		var ret ModelVoucherPosEstimatedAccountingType
		return ret
	}
	return *o.EstimatedAccountingType
}

// GetEstimatedAccountingTypeOk returns a tuple with the EstimatedAccountingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelVoucherPos) GetEstimatedAccountingTypeOk() (*ModelVoucherPosEstimatedAccountingType, bool) {
	if o == nil || IsNil(o.EstimatedAccountingType) {
		return nil, false
	}
	return o.EstimatedAccountingType, true
}

// HasEstimatedAccountingType returns a boolean if a field has been set.
func (o *ModelVoucherPos) HasEstimatedAccountingType() bool {
	if o != nil && !IsNil(o.EstimatedAccountingType) {
		return true
	}

	return false
}

// SetEstimatedAccountingType gets a reference to the given ModelVoucherPosEstimatedAccountingType and assigns it to the EstimatedAccountingType field.
func (o *ModelVoucherPos) SetEstimatedAccountingType(v ModelVoucherPosEstimatedAccountingType) {
	o.EstimatedAccountingType = &v
}

// GetTaxRate returns the TaxRate field value
func (o *ModelVoucherPos) GetTaxRate() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TaxRate
}

// GetTaxRateOk returns a tuple with the TaxRate field value
// and a boolean to check if the value has been set.
func (o *ModelVoucherPos) GetTaxRateOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaxRate, true
}

// SetTaxRate sets field value
func (o *ModelVoucherPos) SetTaxRate(v float32) {
	o.TaxRate = v
}

// GetNet returns the Net field value
func (o *ModelVoucherPos) GetNet() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Net
}

// GetNetOk returns a tuple with the Net field value
// and a boolean to check if the value has been set.
func (o *ModelVoucherPos) GetNetOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Net, true
}

// SetNet sets field value
func (o *ModelVoucherPos) SetNet(v bool) {
	o.Net = v
}

// GetIsAsset returns the IsAsset field value if set, zero value otherwise.
func (o *ModelVoucherPos) GetIsAsset() bool {
	if o == nil || IsNil(o.IsAsset) {
		var ret bool
		return ret
	}
	return *o.IsAsset
}

// GetIsAssetOk returns a tuple with the IsAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelVoucherPos) GetIsAssetOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAsset) {
		return nil, false
	}
	return o.IsAsset, true
}

// HasIsAsset returns a boolean if a field has been set.
func (o *ModelVoucherPos) HasIsAsset() bool {
	if o != nil && !IsNil(o.IsAsset) {
		return true
	}

	return false
}

// SetIsAsset gets a reference to the given bool and assigns it to the IsAsset field.
func (o *ModelVoucherPos) SetIsAsset(v bool) {
	o.IsAsset = &v
}

// GetSumNet returns the SumNet field value
func (o *ModelVoucherPos) GetSumNet() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SumNet
}

// GetSumNetOk returns a tuple with the SumNet field value
// and a boolean to check if the value has been set.
func (o *ModelVoucherPos) GetSumNetOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SumNet, true
}

// SetSumNet sets field value
func (o *ModelVoucherPos) SetSumNet(v float32) {
	o.SumNet = v
}

// GetSumTax returns the SumTax field value if set, zero value otherwise.
func (o *ModelVoucherPos) GetSumTax() float32 {
	if o == nil || IsNil(o.SumTax) {
		var ret float32
		return ret
	}
	return *o.SumTax
}

// GetSumTaxOk returns a tuple with the SumTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelVoucherPos) GetSumTaxOk() (*float32, bool) {
	if o == nil || IsNil(o.SumTax) {
		return nil, false
	}
	return o.SumTax, true
}

// HasSumTax returns a boolean if a field has been set.
func (o *ModelVoucherPos) HasSumTax() bool {
	if o != nil && !IsNil(o.SumTax) {
		return true
	}

	return false
}

// SetSumTax gets a reference to the given float32 and assigns it to the SumTax field.
func (o *ModelVoucherPos) SetSumTax(v float32) {
	o.SumTax = &v
}

// GetSumGross returns the SumGross field value
func (o *ModelVoucherPos) GetSumGross() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SumGross
}

// GetSumGrossOk returns a tuple with the SumGross field value
// and a boolean to check if the value has been set.
func (o *ModelVoucherPos) GetSumGrossOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SumGross, true
}

// SetSumGross sets field value
func (o *ModelVoucherPos) SetSumGross(v float32) {
	o.SumGross = v
}

// GetSumNetAccounting returns the SumNetAccounting field value if set, zero value otherwise.
func (o *ModelVoucherPos) GetSumNetAccounting() float32 {
	if o == nil || IsNil(o.SumNetAccounting) {
		var ret float32
		return ret
	}
	return *o.SumNetAccounting
}

// GetSumNetAccountingOk returns a tuple with the SumNetAccounting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelVoucherPos) GetSumNetAccountingOk() (*float32, bool) {
	if o == nil || IsNil(o.SumNetAccounting) {
		return nil, false
	}
	return o.SumNetAccounting, true
}

// HasSumNetAccounting returns a boolean if a field has been set.
func (o *ModelVoucherPos) HasSumNetAccounting() bool {
	if o != nil && !IsNil(o.SumNetAccounting) {
		return true
	}

	return false
}

// SetSumNetAccounting gets a reference to the given float32 and assigns it to the SumNetAccounting field.
func (o *ModelVoucherPos) SetSumNetAccounting(v float32) {
	o.SumNetAccounting = &v
}

// GetSumTaxAccounting returns the SumTaxAccounting field value if set, zero value otherwise.
func (o *ModelVoucherPos) GetSumTaxAccounting() float32 {
	if o == nil || IsNil(o.SumTaxAccounting) {
		var ret float32
		return ret
	}
	return *o.SumTaxAccounting
}

// GetSumTaxAccountingOk returns a tuple with the SumTaxAccounting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelVoucherPos) GetSumTaxAccountingOk() (*float32, bool) {
	if o == nil || IsNil(o.SumTaxAccounting) {
		return nil, false
	}
	return o.SumTaxAccounting, true
}

// HasSumTaxAccounting returns a boolean if a field has been set.
func (o *ModelVoucherPos) HasSumTaxAccounting() bool {
	if o != nil && !IsNil(o.SumTaxAccounting) {
		return true
	}

	return false
}

// SetSumTaxAccounting gets a reference to the given float32 and assigns it to the SumTaxAccounting field.
func (o *ModelVoucherPos) SetSumTaxAccounting(v float32) {
	o.SumTaxAccounting = &v
}

// GetSumGrossAccounting returns the SumGrossAccounting field value if set, zero value otherwise.
func (o *ModelVoucherPos) GetSumGrossAccounting() float32 {
	if o == nil || IsNil(o.SumGrossAccounting) {
		var ret float32
		return ret
	}
	return *o.SumGrossAccounting
}

// GetSumGrossAccountingOk returns a tuple with the SumGrossAccounting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelVoucherPos) GetSumGrossAccountingOk() (*float32, bool) {
	if o == nil || IsNil(o.SumGrossAccounting) {
		return nil, false
	}
	return o.SumGrossAccounting, true
}

// HasSumGrossAccounting returns a boolean if a field has been set.
func (o *ModelVoucherPos) HasSumGrossAccounting() bool {
	if o != nil && !IsNil(o.SumGrossAccounting) {
		return true
	}

	return false
}

// SetSumGrossAccounting gets a reference to the given float32 and assigns it to the SumGrossAccounting field.
func (o *ModelVoucherPos) SetSumGrossAccounting(v float32) {
	o.SumGrossAccounting = &v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelVoucherPos) GetComment() string {
	if o == nil || IsNil(o.Comment.Get()) {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelVoucherPos) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *ModelVoucherPos) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *ModelVoucherPos) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *ModelVoucherPos) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *ModelVoucherPos) UnsetComment() {
	o.Comment.Unset()
}

func (o ModelVoucherPos) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelVoucherPos) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	toSerialize["objectName"] = o.ObjectName
	toSerialize["mapAll"] = o.MapAll
	// skip: create is readOnly
	// skip: update is readOnly
	if !IsNil(o.SevClient) {
		toSerialize["sevClient"] = o.SevClient
	}
	toSerialize["voucher"] = o.Voucher
	toSerialize["accountingType"] = o.AccountingType
	if !IsNil(o.EstimatedAccountingType) {
		toSerialize["estimatedAccountingType"] = o.EstimatedAccountingType
	}
	toSerialize["taxRate"] = o.TaxRate
	toSerialize["net"] = o.Net
	if !IsNil(o.IsAsset) {
		toSerialize["isAsset"] = o.IsAsset
	}
	toSerialize["sumNet"] = o.SumNet
	// skip: sumTax is readOnly
	toSerialize["sumGross"] = o.SumGross
	// skip: sumNetAccounting is readOnly
	// skip: sumTaxAccounting is readOnly
	// skip: sumGrossAccounting is readOnly
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	return toSerialize, nil
}

type NullableModelVoucherPos struct {
	value *ModelVoucherPos
	isSet bool
}

func (v NullableModelVoucherPos) Get() *ModelVoucherPos {
	return v.value
}

func (v *NullableModelVoucherPos) Set(val *ModelVoucherPos) {
	v.value = val
	v.isSet = true
}

func (v NullableModelVoucherPos) IsSet() bool {
	return v.isSet
}

func (v *NullableModelVoucherPos) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelVoucherPos(val *ModelVoucherPos) *NullableModelVoucherPos {
	return &NullableModelVoucherPos{value: val, isSet: true}
}

func (v NullableModelVoucherPos) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelVoucherPos) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


