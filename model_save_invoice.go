/*
sevDesk API

<b>Contact:</b> To contact our support click <a href='https://landing.sevdesk.de/service-support-center-technik'>here</a><br><br>   # General information  Welcome to our API!<br>  sevDesk offers you the possibility of retrieving data using an interface, namely the sevDesk API, and making changes without having to use the web UI. The sevDesk interface is a REST-Full API. All sevDesk data and functions that are used in the web UI can also be controlled through the API.   # Cross-Origin Resource Sharing  This API features Cross-Origin Resource Sharing (CORS).<br>  It enables cross-domain communication from the browser.<br>  All responses have a wildcard same-origin which makes them completely public and accessible to everyone, including any code on any site.    # Embedding resources  When retrieving resources by using this API, you might encounter nested resources in the resources you requested.<br>  For example, an invoice always contains a contact, of which you can see the ID and the object name.<br>  This API gives you the possibility to embed these resources completely into the resources you originally requested.<br>  Taking our invoice example, this would mean, that you would not only see the ID and object name of a contact, but rather the complete contact resource.    To embed resources, all you need to do is to add the query parameter 'embed' to your GET request.<br>  As values, you can provide the name of the nested resource.<br>  Multiple nested resources are also possible by providing multiple names, separated by a comma.    # Authentication and Authorization  The sevDesk API uses a token authentication to authorize calls. For this purpose every sevDesk administrator has one API token, which is a <b>hexadecimal string</b> containing <b>32 characters</b>. The following clip shows where you can find the API token if this is your first time with our API.<br><br> <video src='OpenAPI/img/findingTheApiToken.mp4' controls width='640' height='360'> Ihr Browser kann dieses Video nicht wiedergeben.<br/> Dieses Video zeigt wie sie Ihr sevDesk API Token finden. </video> <br> The token will be needed in every request that you want to send and needs to be attached to the request url as a <b>Query Parameter</b><br> or provided as a value of an <b>Authorization Header</b>.<br> For security reasons, we suggest putting the API Token in the Authorization Header and not in the query string.<br> However, in the request examples in this documentation, we will keep it in the query string, as it is easier for you to copy them and try them yourself.<br> The following url is an example that shows where your token needs to be placed as a query parameter.<br> In this case, we used some random API token. <ul> <li><span>ht</span>tps://my.sevdesk.de/api/v1/Contact?token=<span style='color:red'>b7794de0085f5cd00560f160f290af38</span></li> </ul> The next example shows the token in the Authorization Header. <ul> <li>\"Authorization\" :<span style='color:red'>\"b7794de0085f5cd00560f160f290af38\"</span></li> </ul> The api tokens have an infinite lifetime and, in other words, exist as long as the sevDesk user exists.<br> For this reason, the user should <b>NEVER</b> be deleted.<br> If really necessary, it is advisable to save the api token as we will <b>NOT</b> be able to retrieve it afterwards!<br> It is also possible to generate a new API token, for example, if you want to prevent the usage of your sevDesk account by other people who got your current API token.<br> To achieve this, you just need to click on the \"generate new\" symbol to the right of your token and confirm it with your password.  # API News  To never miss API news and updates again, subscribe to our <b>free API newsletter</b> and get all relevant information to keep your sevDesk software running smoothly. To subscribe, simply click <a href = 'https://landing.sevdesk.de/api-newsletter'><b>here</b></a> and confirm the email address to which we may send all updates relevant to you.  # API Requests  In our case, REST API requests need to be build by combining the following components. <table> <tr> <th><b>Component</b></th> <th><b>Description</b></th> </tr> <tr> <td>HTTP-Methods</td> <td> <ul> <li>GET (retrieve a resource)</li> <li>POST (create a resource)</li> <li>PUT (update a resource)</li> <li>DELETE (delete a resource)</li> </ul> </td> </tr> <tr> <td>URL of the API</td> <td><span style='color: #2aa198'>ht</span><span style='color: #2aa198'>tps://my.sevdesk.de/api/v1</span></td> </tr> <tr> <td>URI of the resource</td> <td>The resource to query.<br>For example contacts in sevDesk:<br><br> <span style='color:red'>/Contact</span><br><br> Which will result in the following complete URL:<br><br> <span style='color: #2aa198'>ht</span><span style='color: #2aa198'>tps://my.sevdesk.de/api/v1</span><span style='color:red'>/Contact</span> </td> </tr> <tr> <td>Query parameters</td> <td>Are attached by using the connectives <b>?</b> and <b>&</b> in the URL.<br></td> </tr> <tr> <td>Request headers</td> <td>Typical request headers are for example:<br> <div> <br> <ul> <li>Content-type</li> <li>Authorization</li> <li>Accept-Encoding</li> <li>User-Agent</li> <li>...</li> </ul> </div> </td> </tr> <tr> <td>Request body</td> <td>Mostly required in POST and PUT requests.<br> Often the request body contains json, in our case, it also accepts url-encoded data. </td> </tr> </table><br> <span style='color:red'>Note</span>: please pass a meaningful entry at the header \"User-Agent\". If the \"User-Agent\" is set meaningfully, we can offer better support in case of queries from customers.<br> An example how such a \"User-Agent\" can look like: \"Integration-name by abc\". <br><br> This is a sample request for retrieving existing contacts in sevDesk using curl:<br><br> <img src='OpenAPI/img/GETRequest.PNG' alt='Get Request' height='150'><br><br> As you can see, the request contains all the components mentioned above.<br> It's HTTP method is GET, it has a correct endpoint (<span style='color: #2aa198'>ht</span><span style='color: #2aa198'>tps://my.sevdesk.de/api/v1</span><span style='color:red'>/Contact</span>), query parameters like <b>token</b> and additional <b>header</b> information!<br><br> <b>Query Parameters</b><br><br> As you might have seen in the sample request above, there are several other parameters besides \"token\", located in the url.<br> Those are mostly optional but prove to be very useful for a lot of requests as they can limit, extend, sort or filter the data you will get as a response.<br><br> These are the three most used query parameter for the sevDesk API. <table> <tr> <th><b>Parameter</b></th> <th><b>Description</b></th> </tr> <tr> <td>limit</td> <td>Limits the number of entries that are returned.<br> Most useful in GET requests which will most likely deliver big sets of data like country or currency lists.<br> In this case, you can bypass the default limitation on returned entries by providing a high number. </td> </tr> <tr> <td>offset</td> <td>Specifies a certain offset for the data that will be returned.<br> As an example, you can specify \"offset=2\" if you want all entries except for the first two.</td> </tr> <tr> <td>embed</td> <td>Will extend some of the returned data.<br> A brief example can be found below.</td> </tr> </table> This is an example for the usage of the embed parameter.<br> The following first request will return all company contact entries in sevDesk up to a limit of 100 without any additional information and no offset.<br><br> <img src='OpenAPI/img/ContactQueryWithoutEmbed.PNG' width='900' height='850'><br><br> Now have a look at the category attribute located in the response.<br> Naturally, it just contains the id and the object name of the object but no further information.<br> We will now use the parameter embed with the value \"category\".<br><br> <img src='OpenAPI/img/ContactQueryWithEmbed.PNG' width='900' height='850'><br><br> As you can see, the category object is now extended and shows all the attributes and their corresponding values.<br><br> There are lot of other query parameters that can be used to filter the returned data for objects that match a certain pattern, however, those will not be mentioned here and instead can be found at the detail documentation of the most used API endpoints like contact, invoice or voucher.<br><br> <b>Request Headers</b><br><br> The HTTP request (response) headers allow the client as well as the server to pass additional information with the request.<br> They transfer the parameters and arguments which are important for transmitting data over HTTP.<br> Three headers which are useful / necessary when using the sevDesk API are \"Authorization, \"Accept\" and \"Content-type\".<br> Underneath is a brief description of why and how they should be used.<br><br> <b>Authorization</b><br><br> Can be used if you want to provide your API token in the header instead of having it in the url. <ul> <li>Authorization:<span style='color:red'>yourApiToken</span></li> </ul> <b>Accept</b><br><br> Specifies the format of the response.<br> Required for operations with a response body. <ul> <li>Accept:application/<span style='color:red'>format</span> </li> </ul> In our case, <code><span style='color:red'>format</span></code> could be replaced with <code>json</code> or <code>xml</code><br><br> <b>Content-type</b><br><br> Specifies which format is used in the request.<br> Is required for operations with a request body. <ul> <li>Content-type:application/<span style='color:red'>format</span></li> </ul> In our case,<code><span style='color:red'>format</span></code>could be replaced with <code>json</code> or <code>x-www-form-urlencoded</code> <br><br><b>API Responses</b><br><br> HTTP status codes<br> When calling the sevDesk API it is very likely that you will get a HTTP status code in the response.<br> Some API calls will also return JSON response bodies which will contain information about the resource.<br> Each status code which is returned will either be a <b>success</b> code or an <b>error</b> code.<br><br> Success codes <table> <tr> <th><b>Status code</b></th> <th><b>Description</b></th> </tr> <tr> <td><code>200 OK</code></td> <td>The request was successful</td> </tr> <tr> <td><code>201 Created</code></td> <td>Most likely to be found in the response of a <b>POST</b> request.<br> This code indicates that the desired resource was successfully created.</td> </tr> </table> <br>Error codes <table> <tr> <th><b>Status code</b></th> <th><b>Description</b></th> </tr> <tr> <td><code>400 Bad request</code></td> <td>The request you sent is most likely syntactically incorrect.<br> You should check if the parameters in the request body or the url are correct.</td> </tr> <tr> <td><code>401 Unauthorized</code></td> <td>The authentication failed.<br> Most likely caused by a missing or wrong API token.</td> </tr> <tr> <td><code>403 Forbidden</code></td> <td>You do not have the permission the access the resource which is requested.</td> </tr> <tr> <td><code>404 Not found</code></td> <td>The resource you specified does not exist.</td> </tr> <tr> <td><code>500 Internal server error</code></td> <td>An internal server error has occurred.<br> Normally this means that something went wrong on our side.<br> However, sometimes this error will appear if we missed to catch an error which is normally a 400 status code! </td> </tr> </table>  # Your First Request  After reading the introduction to our API, you should now be able to make your first call.<br> For testing our API, we would always recommend to create a trial account for sevDesk to prevent unwanted changes to your main account.<br> A trial account will be in the highest tariff (materials management), so every sevDesk function can be tested! <br><br>To start testing we would recommend one of the following tools: <ul> <li><a href='https://www.getpostman.com/'>Postman</a></li> <li><a href='https://curl.haxx.se/download.html'>cURL</a></li> </ul> This example will illustrate your first request, which is creating a new Contact in sevDesk.<br> <ol> <li>Download <a href='https://www.getpostman.com/'><b>Postman</b></a> for your desired system and start the application</li> <li>Enter <span><b>ht</span>tps://my.sevdesk.de/api/v1/Contact</b> as the url</li> <li>Use the connective <b>?</b> to append <b>token=</b> to the end of the url, or create an authorization header. Insert your API token as the value</li> <li>For this test, select <b>POST</b> as the HTTP method</li> <li>Go to <b>Headers</b> and enter the key-value pair <b>Content-type</b> + <b>application/x-www-form-urlencoded</b><br> As an alternative, you can just go to <b>Body</b> and select <b>x-www-form-urlencoded</b></li> <li>Now go to <b>Body</b> (if you are not there yet) and enter the key-value pairs as shown in the following picture<br><br> <img src='OpenAPI/img/FirstRequestPostman.PNG' width='900'><br><br></li> <li>Click on <b>Send</b>. Your response should now look like this:<br><br> <img src='OpenAPI/img/FirstRequestResponse.PNG' width='900'></li> </ol> As you can see, a successful response in this case returns a JSON-formatted response body containing the contact you just created.<br> For keeping it simple, this was only a minimal example of creating a contact.<br> There are however numerous combinations of parameters that you can provide which add information to your contact.

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sevdesk

import (
	"encoding/json"
	"os"
)

// checks if the SaveInvoice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SaveInvoice{}

// SaveInvoice struct for SaveInvoice
type SaveInvoice struct {
	Invoice ModelInvoice `json:"invoice"`
	InvoicePosSave *ModelInvoicePos `json:"invoicePosSave,omitempty"`
	InvoicePosDelete *SaveInvoiceInvoicePosDelete `json:"invoicePosDelete,omitempty"`
	// Filename of a previously upload file which should be attached.
	Filename **os.File `json:"filename,omitempty"`
	DiscountSave *SaveInvoiceDiscountSave `json:"discountSave,omitempty"`
	DiscountDelete *SaveInvoiceDiscountDelete `json:"discountDelete,omitempty"`
}

// NewSaveInvoice instantiates a new SaveInvoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaveInvoice(invoice ModelInvoice) *SaveInvoice {
	this := SaveInvoice{}
	this.Invoice = invoice
	var invoicePosDelete SaveInvoiceInvoicePosDelete
	this.InvoicePosDelete = &invoicePosDelete
	return &this
}

// NewSaveInvoiceWithDefaults instantiates a new SaveInvoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaveInvoiceWithDefaults() *SaveInvoice {
	this := SaveInvoice{}
	var invoicePosDelete SaveInvoiceInvoicePosDelete
	this.InvoicePosDelete = &invoicePosDelete
	return &this
}

// GetInvoice returns the Invoice field value
func (o *SaveInvoice) GetInvoice() ModelInvoice {
	if o == nil {
		var ret ModelInvoice
		return ret
	}

	return o.Invoice
}

// GetInvoiceOk returns a tuple with the Invoice field value
// and a boolean to check if the value has been set.
func (o *SaveInvoice) GetInvoiceOk() (*ModelInvoice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Invoice, true
}

// SetInvoice sets field value
func (o *SaveInvoice) SetInvoice(v ModelInvoice) {
	o.Invoice = v
}

// GetInvoicePosSave returns the InvoicePosSave field value if set, zero value otherwise.
func (o *SaveInvoice) GetInvoicePosSave() ModelInvoicePos {
	if o == nil || IsNil(o.InvoicePosSave) {
		var ret ModelInvoicePos
		return ret
	}
	return *o.InvoicePosSave
}

// GetInvoicePosSaveOk returns a tuple with the InvoicePosSave field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaveInvoice) GetInvoicePosSaveOk() (*ModelInvoicePos, bool) {
	if o == nil || IsNil(o.InvoicePosSave) {
		return nil, false
	}
	return o.InvoicePosSave, true
}

// HasInvoicePosSave returns a boolean if a field has been set.
func (o *SaveInvoice) HasInvoicePosSave() bool {
	if o != nil && !IsNil(o.InvoicePosSave) {
		return true
	}

	return false
}

// SetInvoicePosSave gets a reference to the given ModelInvoicePos and assigns it to the InvoicePosSave field.
func (o *SaveInvoice) SetInvoicePosSave(v ModelInvoicePos) {
	o.InvoicePosSave = &v
}

// GetInvoicePosDelete returns the InvoicePosDelete field value if set, zero value otherwise.
func (o *SaveInvoice) GetInvoicePosDelete() SaveInvoiceInvoicePosDelete {
	if o == nil || IsNil(o.InvoicePosDelete) {
		var ret SaveInvoiceInvoicePosDelete
		return ret
	}
	return *o.InvoicePosDelete
}

// GetInvoicePosDeleteOk returns a tuple with the InvoicePosDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaveInvoice) GetInvoicePosDeleteOk() (*SaveInvoiceInvoicePosDelete, bool) {
	if o == nil || IsNil(o.InvoicePosDelete) {
		return nil, false
	}
	return o.InvoicePosDelete, true
}

// HasInvoicePosDelete returns a boolean if a field has been set.
func (o *SaveInvoice) HasInvoicePosDelete() bool {
	if o != nil && !IsNil(o.InvoicePosDelete) {
		return true
	}

	return false
}

// SetInvoicePosDelete gets a reference to the given SaveInvoiceInvoicePosDelete and assigns it to the InvoicePosDelete field.
func (o *SaveInvoice) SetInvoicePosDelete(v SaveInvoiceInvoicePosDelete) {
	o.InvoicePosDelete = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *SaveInvoice) GetFilename() *os.File {
	if o == nil || IsNil(o.Filename) {
		var ret *os.File
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaveInvoice) GetFilenameOk() (**os.File, bool) {
	if o == nil || IsNil(o.Filename) {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *SaveInvoice) HasFilename() bool {
	if o != nil && !IsNil(o.Filename) {
		return true
	}

	return false
}

// SetFilename gets a reference to the given *os.File and assigns it to the Filename field.
func (o *SaveInvoice) SetFilename(v *os.File) {
	o.Filename = &v
}

// GetDiscountSave returns the DiscountSave field value if set, zero value otherwise.
func (o *SaveInvoice) GetDiscountSave() SaveInvoiceDiscountSave {
	if o == nil || IsNil(o.DiscountSave) {
		var ret SaveInvoiceDiscountSave
		return ret
	}
	return *o.DiscountSave
}

// GetDiscountSaveOk returns a tuple with the DiscountSave field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaveInvoice) GetDiscountSaveOk() (*SaveInvoiceDiscountSave, bool) {
	if o == nil || IsNil(o.DiscountSave) {
		return nil, false
	}
	return o.DiscountSave, true
}

// HasDiscountSave returns a boolean if a field has been set.
func (o *SaveInvoice) HasDiscountSave() bool {
	if o != nil && !IsNil(o.DiscountSave) {
		return true
	}

	return false
}

// SetDiscountSave gets a reference to the given SaveInvoiceDiscountSave and assigns it to the DiscountSave field.
func (o *SaveInvoice) SetDiscountSave(v SaveInvoiceDiscountSave) {
	o.DiscountSave = &v
}

// GetDiscountDelete returns the DiscountDelete field value if set, zero value otherwise.
func (o *SaveInvoice) GetDiscountDelete() SaveInvoiceDiscountDelete {
	if o == nil || IsNil(o.DiscountDelete) {
		var ret SaveInvoiceDiscountDelete
		return ret
	}
	return *o.DiscountDelete
}

// GetDiscountDeleteOk returns a tuple with the DiscountDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaveInvoice) GetDiscountDeleteOk() (*SaveInvoiceDiscountDelete, bool) {
	if o == nil || IsNil(o.DiscountDelete) {
		return nil, false
	}
	return o.DiscountDelete, true
}

// HasDiscountDelete returns a boolean if a field has been set.
func (o *SaveInvoice) HasDiscountDelete() bool {
	if o != nil && !IsNil(o.DiscountDelete) {
		return true
	}

	return false
}

// SetDiscountDelete gets a reference to the given SaveInvoiceDiscountDelete and assigns it to the DiscountDelete field.
func (o *SaveInvoice) SetDiscountDelete(v SaveInvoiceDiscountDelete) {
	o.DiscountDelete = &v
}

func (o SaveInvoice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SaveInvoice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["invoice"] = o.Invoice
	if !IsNil(o.InvoicePosSave) {
		toSerialize["invoicePosSave"] = o.InvoicePosSave
	}
	if !IsNil(o.InvoicePosDelete) {
		toSerialize["invoicePosDelete"] = o.InvoicePosDelete
	}
	if !IsNil(o.Filename) {
		toSerialize["filename"] = o.Filename
	}
	if !IsNil(o.DiscountSave) {
		toSerialize["discountSave"] = o.DiscountSave
	}
	if !IsNil(o.DiscountDelete) {
		toSerialize["discountDelete"] = o.DiscountDelete
	}
	return toSerialize, nil
}

type NullableSaveInvoice struct {
	value *SaveInvoice
	isSet bool
}

func (v NullableSaveInvoice) Get() *SaveInvoice {
	return v.value
}

func (v *NullableSaveInvoice) Set(val *SaveInvoice) {
	v.value = val
	v.isSet = true
}

func (v NullableSaveInvoice) IsSet() bool {
	return v.isSet
}

func (v *NullableSaveInvoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaveInvoice(val *SaveInvoice) *NullableSaveInvoice {
	return &NullableSaveInvoice{value: val, isSet: true}
}

func (v NullableSaveInvoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaveInvoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


