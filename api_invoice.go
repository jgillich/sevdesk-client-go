/*
sevDesk API

<b>Contact:</b> To contact our support click <a href='https://landing.sevdesk.de/service-support-center-technik'>here</a><br><br>   # General information  Welcome to our API!<br>  sevDesk offers you the possibility of retrieving data using an interface, namely the sevDesk API, and making changes without having to use the web UI. The sevDesk interface is a REST-Full API. All sevDesk data and functions that are used in the web UI can also be controlled through the API.   # Cross-Origin Resource Sharing  This API features Cross-Origin Resource Sharing (CORS).<br>  It enables cross-domain communication from the browser.<br>  All responses have a wildcard same-origin which makes them completely public and accessible to everyone, including any code on any site.    # Embedding resources  When retrieving resources by using this API, you might encounter nested resources in the resources you requested.<br>  For example, an invoice always contains a contact, of which you can see the ID and the object name.<br>  This API gives you the possibility to embed these resources completely into the resources you originally requested.<br>  Taking our invoice example, this would mean, that you would not only see the ID and object name of a contact, but rather the complete contact resource.    To embed resources, all you need to do is to add the query parameter 'embed' to your GET request.<br>  As values, you can provide the name of the nested resource.<br>  Multiple nested resources are also possible by providing multiple names, separated by a comma.    # Authentication and Authorization  The sevDesk API uses a token authentication to authorize calls. For this purpose every sevDesk administrator has one API token, which is a <b>hexadecimal string</b> containing <b>32 characters</b>. The following clip shows where you can find the API token if this is your first time with our API.<br><br> <video src='OpenAPI/img/findingTheApiToken.mp4' controls width='640' height='360'> Ihr Browser kann dieses Video nicht wiedergeben.<br/> Dieses Video zeigt wie sie Ihr sevDesk API Token finden. </video> <br> The token will be needed in every request that you want to send and needs to be attached to the request url as a <b>Query Parameter</b><br> or provided as a value of an <b>Authorization Header</b>.<br> For security reasons, we suggest putting the API Token in the Authorization Header and not in the query string.<br> However, in the request examples in this documentation, we will keep it in the query string, as it is easier for you to copy them and try them yourself.<br> The following url is an example that shows where your token needs to be placed as a query parameter.<br> In this case, we used some random API token. <ul> <li><span>ht</span>tps://my.sevdesk.de/api/v1/Contact?token=<span style='color:red'>b7794de0085f5cd00560f160f290af38</span></li> </ul> The next example shows the token in the Authorization Header. <ul> <li>\"Authorization\" :<span style='color:red'>\"b7794de0085f5cd00560f160f290af38\"</span></li> </ul> The api tokens have an infinite lifetime and, in other words, exist as long as the sevDesk user exists.<br> For this reason, the user should <b>NEVER</b> be deleted.<br> If really necessary, it is advisable to save the api token as we will <b>NOT</b> be able to retrieve it afterwards!<br> It is also possible to generate a new API token, for example, if you want to prevent the usage of your sevDesk account by other people who got your current API token.<br> To achieve this, you just need to click on the \"generate new\" symbol to the right of your token and confirm it with your password.  # API News  To never miss API news and updates again, subscribe to our <b>free API newsletter</b> and get all relevant information to keep your sevDesk software running smoothly. To subscribe, simply click <a href = 'https://landing.sevdesk.de/api-newsletter'><b>here</b></a> and confirm the email address to which we may send all updates relevant to you.  # API Requests  In our case, REST API requests need to be build by combining the following components. <table> <tr> <th><b>Component</b></th> <th><b>Description</b></th> </tr> <tr> <td>HTTP-Methods</td> <td> <ul> <li>GET (retrieve a resource)</li> <li>POST (create a resource)</li> <li>PUT (update a resource)</li> <li>DELETE (delete a resource)</li> </ul> </td> </tr> <tr> <td>URL of the API</td> <td><span style='color: #2aa198'>ht</span><span style='color: #2aa198'>tps://my.sevdesk.de/api/v1</span></td> </tr> <tr> <td>URI of the resource</td> <td>The resource to query.<br>For example contacts in sevDesk:<br><br> <span style='color:red'>/Contact</span><br><br> Which will result in the following complete URL:<br><br> <span style='color: #2aa198'>ht</span><span style='color: #2aa198'>tps://my.sevdesk.de/api/v1</span><span style='color:red'>/Contact</span> </td> </tr> <tr> <td>Query parameters</td> <td>Are attached by using the connectives <b>?</b> and <b>&</b> in the URL.<br></td> </tr> <tr> <td>Request headers</td> <td>Typical request headers are for example:<br> <div> <br> <ul> <li>Content-type</li> <li>Authorization</li> <li>Accept-Encoding</li> <li>User-Agent</li> <li>...</li> </ul> </div> </td> </tr> <tr> <td>Request body</td> <td>Mostly required in POST and PUT requests.<br> Often the request body contains json, in our case, it also accepts url-encoded data. </td> </tr> </table><br> <span style='color:red'>Note</span>: please pass a meaningful entry at the header \"User-Agent\". If the \"User-Agent\" is set meaningfully, we can offer better support in case of queries from customers.<br> An example how such a \"User-Agent\" can look like: \"Integration-name by abc\". <br><br> This is a sample request for retrieving existing contacts in sevDesk using curl:<br><br> <img src='OpenAPI/img/GETRequest.PNG' alt='Get Request' height='150'><br><br> As you can see, the request contains all the components mentioned above.<br> It's HTTP method is GET, it has a correct endpoint (<span style='color: #2aa198'>ht</span><span style='color: #2aa198'>tps://my.sevdesk.de/api/v1</span><span style='color:red'>/Contact</span>), query parameters like <b>token</b> and additional <b>header</b> information!<br><br> <b>Query Parameters</b><br><br> As you might have seen in the sample request above, there are several other parameters besides \"token\", located in the url.<br> Those are mostly optional but prove to be very useful for a lot of requests as they can limit, extend, sort or filter the data you will get as a response.<br><br> These are the three most used query parameter for the sevDesk API. <table> <tr> <th><b>Parameter</b></th> <th><b>Description</b></th> </tr> <tr> <td>limit</td> <td>Limits the number of entries that are returned.<br> Most useful in GET requests which will most likely deliver big sets of data like country or currency lists.<br> In this case, you can bypass the default limitation on returned entries by providing a high number. </td> </tr> <tr> <td>offset</td> <td>Specifies a certain offset for the data that will be returned.<br> As an example, you can specify \"offset=2\" if you want all entries except for the first two.</td> </tr> <tr> <td>embed</td> <td>Will extend some of the returned data.<br> A brief example can be found below.</td> </tr> </table> This is an example for the usage of the embed parameter.<br> The following first request will return all company contact entries in sevDesk up to a limit of 100 without any additional information and no offset.<br><br> <img src='OpenAPI/img/ContactQueryWithoutEmbed.PNG' width='900' height='850'><br><br> Now have a look at the category attribute located in the response.<br> Naturally, it just contains the id and the object name of the object but no further information.<br> We will now use the parameter embed with the value \"category\".<br><br> <img src='OpenAPI/img/ContactQueryWithEmbed.PNG' width='900' height='850'><br><br> As you can see, the category object is now extended and shows all the attributes and their corresponding values.<br><br> There are lot of other query parameters that can be used to filter the returned data for objects that match a certain pattern, however, those will not be mentioned here and instead can be found at the detail documentation of the most used API endpoints like contact, invoice or voucher.<br><br> <b>Request Headers</b><br><br> The HTTP request (response) headers allow the client as well as the server to pass additional information with the request.<br> They transfer the parameters and arguments which are important for transmitting data over HTTP.<br> Three headers which are useful / necessary when using the sevDesk API are \"Authorization, \"Accept\" and \"Content-type\".<br> Underneath is a brief description of why and how they should be used.<br><br> <b>Authorization</b><br><br> Can be used if you want to provide your API token in the header instead of having it in the url. <ul> <li>Authorization:<span style='color:red'>yourApiToken</span></li> </ul> <b>Accept</b><br><br> Specifies the format of the response.<br> Required for operations with a response body. <ul> <li>Accept:application/<span style='color:red'>format</span> </li> </ul> In our case, <code><span style='color:red'>format</span></code> could be replaced with <code>json</code> or <code>xml</code><br><br> <b>Content-type</b><br><br> Specifies which format is used in the request.<br> Is required for operations with a request body. <ul> <li>Content-type:application/<span style='color:red'>format</span></li> </ul> In our case,<code><span style='color:red'>format</span></code>could be replaced with <code>json</code> or <code>x-www-form-urlencoded</code> <br><br><b>API Responses</b><br><br> HTTP status codes<br> When calling the sevDesk API it is very likely that you will get a HTTP status code in the response.<br> Some API calls will also return JSON response bodies which will contain information about the resource.<br> Each status code which is returned will either be a <b>success</b> code or an <b>error</b> code.<br><br> Success codes <table> <tr> <th><b>Status code</b></th> <th><b>Description</b></th> </tr> <tr> <td><code>200 OK</code></td> <td>The request was successful</td> </tr> <tr> <td><code>201 Created</code></td> <td>Most likely to be found in the response of a <b>POST</b> request.<br> This code indicates that the desired resource was successfully created.</td> </tr> </table> <br>Error codes <table> <tr> <th><b>Status code</b></th> <th><b>Description</b></th> </tr> <tr> <td><code>400 Bad request</code></td> <td>The request you sent is most likely syntactically incorrect.<br> You should check if the parameters in the request body or the url are correct.</td> </tr> <tr> <td><code>401 Unauthorized</code></td> <td>The authentication failed.<br> Most likely caused by a missing or wrong API token.</td> </tr> <tr> <td><code>403 Forbidden</code></td> <td>You do not have the permission the access the resource which is requested.</td> </tr> <tr> <td><code>404 Not found</code></td> <td>The resource you specified does not exist.</td> </tr> <tr> <td><code>500 Internal server error</code></td> <td>An internal server error has occurred.<br> Normally this means that something went wrong on our side.<br> However, sometimes this error will appear if we missed to catch an error which is normally a 400 status code! </td> </tr> </table>  # Your First Request  After reading the introduction to our API, you should now be able to make your first call.<br> For testing our API, we would always recommend to create a trial account for sevDesk to prevent unwanted changes to your main account.<br> A trial account will be in the highest tariff (materials management), so every sevDesk function can be tested! <br><br>To start testing we would recommend one of the following tools: <ul> <li><a href='https://www.getpostman.com/'>Postman</a></li> <li><a href='https://curl.haxx.se/download.html'>cURL</a></li> </ul> This example will illustrate your first request, which is creating a new Contact in sevDesk.<br> <ol> <li>Download <a href='https://www.getpostman.com/'><b>Postman</b></a> for your desired system and start the application</li> <li>Enter <span><b>ht</span>tps://my.sevdesk.de/api/v1/Contact</b> as the url</li> <li>Use the connective <b>?</b> to append <b>token=</b> to the end of the url, or create an authorization header. Insert your API token as the value</li> <li>For this test, select <b>POST</b> as the HTTP method</li> <li>Go to <b>Headers</b> and enter the key-value pair <b>Content-type</b> + <b>application/x-www-form-urlencoded</b><br> As an alternative, you can just go to <b>Body</b> and select <b>x-www-form-urlencoded</b></li> <li>Now go to <b>Body</b> (if you are not there yet) and enter the key-value pairs as shown in the following picture<br><br> <img src='OpenAPI/img/FirstRequestPostman.PNG' width='900'><br><br></li> <li>Click on <b>Send</b>. Your response should now look like this:<br><br> <img src='OpenAPI/img/FirstRequestResponse.PNG' width='900'></li> </ol> As you can see, a successful response in this case returns a JSON-formatted response body containing the contact you just created.<br> For keeping it simple, this was only a minimal example of creating a contact.<br> There are however numerous combinations of parameters that you can provide which add information to your contact.

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sevdesk

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


// InvoiceApiService InvoiceApi service
type InvoiceApiService service

type ApiBookInvoiceRequest struct {
	ctx context.Context
	ApiService *InvoiceApiService
	invoiceId int32
	bookInvoiceRequest *BookInvoiceRequest
}

// Booking data
func (r ApiBookInvoiceRequest) BookInvoiceRequest(bookInvoiceRequest BookInvoiceRequest) ApiBookInvoiceRequest {
	r.bookInvoiceRequest = &bookInvoiceRequest
	return r
}

func (r ApiBookInvoiceRequest) Execute() (*BookInvoice200Response, *http.Response, error) {
	return r.ApiService.BookInvoiceExecute(r)
}

/*
BookInvoice Book an invoice

Booking the invoice with a transaction is probably the most important part in the bookkeeping process.<br> There are several ways on correctly booking an invoice, all by using the same endpoint.<br> for more information look <a href='https://api.sevdesk.de/#section/How-to-book-an-invoice'>here</a>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param invoiceId ID of invoice to book
 @return ApiBookInvoiceRequest
*/
func (a *InvoiceApiService) BookInvoice(ctx context.Context, invoiceId int32) ApiBookInvoiceRequest {
	return ApiBookInvoiceRequest{
		ApiService: a,
		ctx: ctx,
		invoiceId: invoiceId,
	}
}

// Execute executes the request
//  @return BookInvoice200Response
func (a *InvoiceApiService) BookInvoiceExecute(r ApiBookInvoiceRequest) (*BookInvoice200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BookInvoice200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoiceApiService.BookInvoice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Invoice/{invoiceId}/bookAmount"
	localVarPath = strings.Replace(localVarPath, "{"+"invoiceId"+"}", url.PathEscape(parameterValueToString(r.invoiceId, "invoiceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.bookInvoiceRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCancelInvoiceRequest struct {
	ctx context.Context
	ApiService *InvoiceApiService
	invoiceId int32
}

func (r ApiCancelInvoiceRequest) Execute() (*GetInvoiceById200Response, *http.Response, error) {
	return r.ApiService.CancelInvoiceExecute(r)
}

/*
CancelInvoice Cancel an invoice / Create cancellation invoice

This endpoint will cancel the specified invoice therefor creating a cancellation invoice.<br>
     The cancellation invoice will be automatically paid and the source invoices status will change to 'cancelled'.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param invoiceId ID of invoice to be cancelled
 @return ApiCancelInvoiceRequest
*/
func (a *InvoiceApiService) CancelInvoice(ctx context.Context, invoiceId int32) ApiCancelInvoiceRequest {
	return ApiCancelInvoiceRequest{
		ApiService: a,
		ctx: ctx,
		invoiceId: invoiceId,
	}
}

// Execute executes the request
//  @return GetInvoiceById200Response
func (a *InvoiceApiService) CancelInvoiceExecute(r ApiCancelInvoiceRequest) (*GetInvoiceById200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetInvoiceById200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoiceApiService.CancelInvoice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Invoice/{invoiceId}/cancelInvoice"
	localVarPath = strings.Replace(localVarPath, "{"+"invoiceId"+"}", url.PathEscape(parameterValueToString(r.invoiceId, "invoiceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCreateInvoiceByFactoryRequest struct {
	ctx context.Context
	ApiService *InvoiceApiService
	saveInvoice *SaveInvoice
}

// Creation data. Please be aware, that you need to provide at least all required parameter      of the invoice model!
func (r ApiCreateInvoiceByFactoryRequest) SaveInvoice(saveInvoice SaveInvoice) ApiCreateInvoiceByFactoryRequest {
	r.saveInvoice = &saveInvoice
	return r
}

func (r ApiCreateInvoiceByFactoryRequest) Execute() (*SaveInvoiceResponse, *http.Response, error) {
	return r.ApiService.CreateInvoiceByFactoryExecute(r)
}

/*
CreateInvoiceByFactory Create a new invoice

This endpoint offers you the following functionality.
     <ul>
        <li>Create invoices together with positions and discounts</li>
        <li>Delete positions while adding new ones</li>
        <li>Delete or add discounts, or both at the same time</li>
        <li>Automatically fill the address of the supplied contact into the invoice address</li>
     </ul>
     To make your own request sample slimmer, you can omit all parameters which are not required and nullable.
     However, for a valid and logical bookkeeping document, you will also need some of them to ensure that all the necessary data is in the invoice.<br><br> The list of parameters starts with the invoice array.<br> This array contains all required attributes for a complete invoice.<br> Most of the attributes are covered in the invoice attribute list, there are only two parameters standing out, namely <b>mapAll</b> and <b>objectName</b>.<br> These are just needed for our system and you always need to provide them.<br><br> The list of parameters then continues with the invoice position array.<br> With this array you have the possibility to add multiple positions at once.<br> In the example it only contains one position, again together with the parameters <b>mapAll</b> and <b>objectName</b>, however, you can add more invoice positions by extending the array.<br> So if you wanted to add another position, you would add the same list of parameters with an incremented array index of "1" instead of "0".<br><br> The list ends with the four parameters invoicePosDelete, discountSave, discountDelete and takeDefaultAddress.<br> They only play a minor role if you only want to create an invoice but we will shortly explain what they can do.<br> With invoicePosDelete you have to option to delete invoice positions as this request can also be used to update invoices.<br> With discountSave you can add discounts to your invoice.<br> With discountDelete you can delete discounts from your invoice.<br> With takeDefaultAddress you can specify that the first address of the contact you are using for the invoice is taken for the invoice address attribute automatically, so you don't need to provide the address yourself.<br> If you want to know more about these parameters, for example if you want to use this request to update invoices, feel free to contact our support.<br><br> Finally, after covering all parameters, they only important information left, is that the order of the last four attributes always needs to be kept.<br> You will also always need to provide all of them, as otherwise the request won't work properly.<br><br> <b>Warning":"</b> You can not create a regular invoice with the <b>deliveryDate</b> being later than the <b>invoiceDate</b>.<br> To do that you will need to create a so called <b>Abschlagsrechnung</b> by setting the <b>invoiceType</b> parameter to <b>AR</b>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateInvoiceByFactoryRequest
*/
func (a *InvoiceApiService) CreateInvoiceByFactory(ctx context.Context) ApiCreateInvoiceByFactoryRequest {
	return ApiCreateInvoiceByFactoryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SaveInvoiceResponse
func (a *InvoiceApiService) CreateInvoiceByFactoryExecute(r ApiCreateInvoiceByFactoryRequest) (*SaveInvoiceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SaveInvoiceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoiceApiService.CreateInvoiceByFactory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Invoice/Factory/saveInvoice"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.saveInvoice
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCreateInvoiceFromOrderRequest struct {
	ctx context.Context
	ApiService *InvoiceApiService
	invoiceId *int32
	invoiceObjectName *string
	modelCreateInvoiceFromOrder *ModelCreateInvoiceFromOrder
}

// the id of the invoice
func (r ApiCreateInvoiceFromOrderRequest) InvoiceId(invoiceId int32) ApiCreateInvoiceFromOrderRequest {
	r.invoiceId = &invoiceId
	return r
}

// Model name, which is &#39;Invoice&#39;
func (r ApiCreateInvoiceFromOrderRequest) InvoiceObjectName(invoiceObjectName string) ApiCreateInvoiceFromOrderRequest {
	r.invoiceObjectName = &invoiceObjectName
	return r
}

// Create invoice
func (r ApiCreateInvoiceFromOrderRequest) ModelCreateInvoiceFromOrder(modelCreateInvoiceFromOrder ModelCreateInvoiceFromOrder) ApiCreateInvoiceFromOrderRequest {
	r.modelCreateInvoiceFromOrder = &modelCreateInvoiceFromOrder
	return r
}

func (r ApiCreateInvoiceFromOrderRequest) Execute() (*GetInvoiceById200Response, *http.Response, error) {
	return r.ApiService.CreateInvoiceFromOrderExecute(r)
}

/*
CreateInvoiceFromOrder Create invoice from order

Create an invoice from an order

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateInvoiceFromOrderRequest
*/
func (a *InvoiceApiService) CreateInvoiceFromOrder(ctx context.Context) ApiCreateInvoiceFromOrderRequest {
	return ApiCreateInvoiceFromOrderRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetInvoiceById200Response
func (a *InvoiceApiService) CreateInvoiceFromOrderExecute(r ApiCreateInvoiceFromOrderRequest) (*GetInvoiceById200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetInvoiceById200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoiceApiService.CreateInvoiceFromOrder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Invoice/Factory/createInvoiceFromOrder"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.invoiceId == nil {
		return localVarReturnValue, nil, reportError("invoiceId is required and must be specified")
	}
	if r.invoiceObjectName == nil {
		return localVarReturnValue, nil, reportError("invoiceObjectName is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "invoice[id]", r.invoiceId, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "invoice[objectName]", r.invoiceObjectName, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.modelCreateInvoiceFromOrder
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCreateInvoiceReminderRequest struct {
	ctx context.Context
	ApiService *InvoiceApiService
	invoiceId *int32
	invoiceObjectName *string
	createInvoiceReminderRequest *CreateInvoiceReminderRequest
}

// the id of the invoice
func (r ApiCreateInvoiceReminderRequest) InvoiceId(invoiceId int32) ApiCreateInvoiceReminderRequest {
	r.invoiceId = &invoiceId
	return r
}

// Model name, which is &#39;Invoice&#39;
func (r ApiCreateInvoiceReminderRequest) InvoiceObjectName(invoiceObjectName string) ApiCreateInvoiceReminderRequest {
	r.invoiceObjectName = &invoiceObjectName
	return r
}

// Create invoice
func (r ApiCreateInvoiceReminderRequest) CreateInvoiceReminderRequest(createInvoiceReminderRequest CreateInvoiceReminderRequest) ApiCreateInvoiceReminderRequest {
	r.createInvoiceReminderRequest = &createInvoiceReminderRequest
	return r
}

func (r ApiCreateInvoiceReminderRequest) Execute() (*GetInvoiceById200Response, *http.Response, error) {
	return r.ApiService.CreateInvoiceReminderExecute(r)
}

/*
CreateInvoiceReminder Create invoice reminder

Create an reminder from an invoice

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateInvoiceReminderRequest
*/
func (a *InvoiceApiService) CreateInvoiceReminder(ctx context.Context) ApiCreateInvoiceReminderRequest {
	return ApiCreateInvoiceReminderRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetInvoiceById200Response
func (a *InvoiceApiService) CreateInvoiceReminderExecute(r ApiCreateInvoiceReminderRequest) (*GetInvoiceById200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetInvoiceById200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoiceApiService.CreateInvoiceReminder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Invoice/Factory/createInvoiceReminder"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.invoiceId == nil {
		return localVarReturnValue, nil, reportError("invoiceId is required and must be specified")
	}
	if r.invoiceObjectName == nil {
		return localVarReturnValue, nil, reportError("invoiceObjectName is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "invoice[id]", r.invoiceId, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "invoice[objectName]", r.invoiceObjectName, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createInvoiceReminderRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetInvoiceByIdRequest struct {
	ctx context.Context
	ApiService *InvoiceApiService
	invoiceId int32
}

func (r ApiGetInvoiceByIdRequest) Execute() (*GetInvoiceById200Response, *http.Response, error) {
	return r.ApiService.GetInvoiceByIdExecute(r)
}

/*
GetInvoiceById Find invoice by ID

Returns a single invoice

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param invoiceId ID of invoice to return
 @return ApiGetInvoiceByIdRequest
*/
func (a *InvoiceApiService) GetInvoiceById(ctx context.Context, invoiceId int32) ApiGetInvoiceByIdRequest {
	return ApiGetInvoiceByIdRequest{
		ApiService: a,
		ctx: ctx,
		invoiceId: invoiceId,
	}
}

// Execute executes the request
//  @return GetInvoiceById200Response
func (a *InvoiceApiService) GetInvoiceByIdExecute(r ApiGetInvoiceByIdRequest) (*GetInvoiceById200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetInvoiceById200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoiceApiService.GetInvoiceById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Invoice/{invoiceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"invoiceId"+"}", url.PathEscape(parameterValueToString(r.invoiceId, "invoiceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetInvoicePositionsByIdRequest struct {
	ctx context.Context
	ApiService *InvoiceApiService
	invoiceId int32
	limit *int32
	offset *int32
	embed *[]string
}

// limits the number of entries returned
func (r ApiGetInvoicePositionsByIdRequest) Limit(limit int32) ApiGetInvoicePositionsByIdRequest {
	r.limit = &limit
	return r
}

// set the index where the returned entries start
func (r ApiGetInvoicePositionsByIdRequest) Offset(offset int32) ApiGetInvoicePositionsByIdRequest {
	r.offset = &offset
	return r
}

// Get some additional information. Embed can handle multiple values, they must be separated by comma.
func (r ApiGetInvoicePositionsByIdRequest) Embed(embed []string) ApiGetInvoicePositionsByIdRequest {
	r.embed = &embed
	return r
}

func (r ApiGetInvoicePositionsByIdRequest) Execute() (*GetInvoicePos200Response, *http.Response, error) {
	return r.ApiService.GetInvoicePositionsByIdExecute(r)
}

/*
GetInvoicePositionsById Find invoice positions

Returns all positions of an invoice

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param invoiceId ID of invoice to return the positions
 @return ApiGetInvoicePositionsByIdRequest
*/
func (a *InvoiceApiService) GetInvoicePositionsById(ctx context.Context, invoiceId int32) ApiGetInvoicePositionsByIdRequest {
	return ApiGetInvoicePositionsByIdRequest{
		ApiService: a,
		ctx: ctx,
		invoiceId: invoiceId,
	}
}

// Execute executes the request
//  @return GetInvoicePos200Response
func (a *InvoiceApiService) GetInvoicePositionsByIdExecute(r ApiGetInvoicePositionsByIdRequest) (*GetInvoicePos200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetInvoicePos200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoiceApiService.GetInvoicePositionsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Invoice/{invoiceId}/getPositions"
	localVarPath = strings.Replace(localVarPath, "{"+"invoiceId"+"}", url.PathEscape(parameterValueToString(r.invoiceId, "invoiceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.embed != nil {
		t := *r.embed
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "embed", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "embed", t, "multi")
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetInvoicesRequest struct {
	ctx context.Context
	ApiService *InvoiceApiService
	status *float32
	invoiceNumber *string
	startDate *int32
	endDate *int32
	contactId *int32
	contactObjectName *string
}

// Status of the invoices
func (r ApiGetInvoicesRequest) Status(status float32) ApiGetInvoicesRequest {
	r.status = &status
	return r
}

// Retrieve all invoices with this invoice number
func (r ApiGetInvoicesRequest) InvoiceNumber(invoiceNumber string) ApiGetInvoicesRequest {
	r.invoiceNumber = &invoiceNumber
	return r
}

// Retrieve all invoices with a date equal or higher
func (r ApiGetInvoicesRequest) StartDate(startDate int32) ApiGetInvoicesRequest {
	r.startDate = &startDate
	return r
}

// Retrieve all invoices with a date equal or lower
func (r ApiGetInvoicesRequest) EndDate(endDate int32) ApiGetInvoicesRequest {
	r.endDate = &endDate
	return r
}

// Retrieve all invoices with this contact. Must be provided with contact[objectName]
func (r ApiGetInvoicesRequest) ContactId(contactId int32) ApiGetInvoicesRequest {
	r.contactId = &contactId
	return r
}

// Only required if contact[id] was provided. &#39;Contact&#39; should be used as value.
func (r ApiGetInvoicesRequest) ContactObjectName(contactObjectName string) ApiGetInvoicesRequest {
	r.contactObjectName = &contactObjectName
	return r
}

func (r ApiGetInvoicesRequest) Execute() (*GetInvoiceById200Response, *http.Response, error) {
	return r.ApiService.GetInvoicesExecute(r)
}

/*
GetInvoices Retrieve invoices

There are a multitude of parameter which can be used to filter. A few of them are attached but
     for a complete list please check out <a href='https://api.sevdesk.de/#section/How-to-filter-for-certain-invoices'>this</a> list

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetInvoicesRequest
*/
func (a *InvoiceApiService) GetInvoices(ctx context.Context) ApiGetInvoicesRequest {
	return ApiGetInvoicesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetInvoiceById200Response
func (a *InvoiceApiService) GetInvoicesExecute(r ApiGetInvoicesRequest) (*GetInvoiceById200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetInvoiceById200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoiceApiService.GetInvoices")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Invoice"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "")
	}
	if r.invoiceNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "invoiceNumber", r.invoiceNumber, "")
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startDate", r.startDate, "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endDate", r.endDate, "")
	}
	if r.contactId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "contact[id]", r.contactId, "")
	}
	if r.contactObjectName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "contact[objectName]", r.contactObjectName, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetIsInvoicePartiallyPaidRequest struct {
	ctx context.Context
	ApiService *InvoiceApiService
	invoiceId int32
}

func (r ApiGetIsInvoicePartiallyPaidRequest) Execute() (*GetIsInvoicePartiallyPaid200Response, *http.Response, error) {
	return r.ApiService.GetIsInvoicePartiallyPaidExecute(r)
}

/*
GetIsInvoicePartiallyPaid Check if an invoice is already partially paid

Returns 'true' if the given invoice is partially paid - 'false' if it is not.
    Invoices which are completely paid are regarded as not partially paid.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param invoiceId ID of invoice to return
 @return ApiGetIsInvoicePartiallyPaidRequest
*/
func (a *InvoiceApiService) GetIsInvoicePartiallyPaid(ctx context.Context, invoiceId int32) ApiGetIsInvoicePartiallyPaidRequest {
	return ApiGetIsInvoicePartiallyPaidRequest{
		ApiService: a,
		ctx: ctx,
		invoiceId: invoiceId,
	}
}

// Execute executes the request
//  @return GetIsInvoicePartiallyPaid200Response
func (a *InvoiceApiService) GetIsInvoicePartiallyPaidExecute(r ApiGetIsInvoicePartiallyPaidRequest) (*GetIsInvoicePartiallyPaid200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetIsInvoicePartiallyPaid200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoiceApiService.GetIsInvoicePartiallyPaid")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Invoice/{invoiceId}/getIsPartiallyPaid"
	localVarPath = strings.Replace(localVarPath, "{"+"invoiceId"+"}", url.PathEscape(parameterValueToString(r.invoiceId, "invoiceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiInvoiceGetPdfRequest struct {
	ctx context.Context
	ApiService *InvoiceApiService
	invoiceId int32
	download *bool
	preventSendBy *bool
}

// If u want to download the pdf of the invoice.
func (r ApiInvoiceGetPdfRequest) Download(download bool) ApiInvoiceGetPdfRequest {
	r.download = &download
	return r
}

// Defines if u want to send the invoice.
func (r ApiInvoiceGetPdfRequest) PreventSendBy(preventSendBy bool) ApiInvoiceGetPdfRequest {
	r.preventSendBy = &preventSendBy
	return r
}

func (r ApiInvoiceGetPdfRequest) Execute() (*InvoiceGetPdf200Response, *http.Response, error) {
	return r.ApiService.InvoiceGetPdfExecute(r)
}

/*
InvoiceGetPdf Retrieve pdf document of an invoice

Retrieves the pdf document of an invoice with additional metadata.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param invoiceId ID of invoice from which you want the pdf
 @return ApiInvoiceGetPdfRequest
*/
func (a *InvoiceApiService) InvoiceGetPdf(ctx context.Context, invoiceId int32) ApiInvoiceGetPdfRequest {
	return ApiInvoiceGetPdfRequest{
		ApiService: a,
		ctx: ctx,
		invoiceId: invoiceId,
	}
}

// Execute executes the request
//  @return InvoiceGetPdf200Response
func (a *InvoiceApiService) InvoiceGetPdfExecute(r ApiInvoiceGetPdfRequest) (*InvoiceGetPdf200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InvoiceGetPdf200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoiceApiService.InvoiceGetPdf")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Invoice/{invoiceId}/getPdf"
	localVarPath = strings.Replace(localVarPath, "{"+"invoiceId"+"}", url.PathEscape(parameterValueToString(r.invoiceId, "invoiceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.download != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "download", r.download, "")
	}
	if r.preventSendBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "preventSendBy", r.preventSendBy, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiInvoiceRenderRequest struct {
	ctx context.Context
	ApiService *InvoiceApiService
	invoiceId int32
	invoiceRenderRequest *InvoiceRenderRequest
}

// Define if the document should be forcefully re-rendered.
func (r ApiInvoiceRenderRequest) InvoiceRenderRequest(invoiceRenderRequest InvoiceRenderRequest) ApiInvoiceRenderRequest {
	r.invoiceRenderRequest = &invoiceRenderRequest
	return r
}

func (r ApiInvoiceRenderRequest) Execute() (*InvoiceRender201Response, *http.Response, error) {
	return r.ApiService.InvoiceRenderExecute(r)
}

/*
InvoiceRender Render the pdf document of an invoice

Using this endpoint you can render the pdf document of an invoice.<br>
     Use cases for this are the retrieval of the pdf location or the forceful re-render of a already sent invoice.<br>
     Please be aware that changing an invoice after it has been sent to a customer is not an allowed bookkeeping process.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param invoiceId ID of invoice to render
 @return ApiInvoiceRenderRequest
*/
func (a *InvoiceApiService) InvoiceRender(ctx context.Context, invoiceId int32) ApiInvoiceRenderRequest {
	return ApiInvoiceRenderRequest{
		ApiService: a,
		ctx: ctx,
		invoiceId: invoiceId,
	}
}

// Execute executes the request
//  @return InvoiceRender201Response
func (a *InvoiceApiService) InvoiceRenderExecute(r ApiInvoiceRenderRequest) (*InvoiceRender201Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InvoiceRender201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoiceApiService.InvoiceRender")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Invoice/{invoiceId}/render"
	localVarPath = strings.Replace(localVarPath, "{"+"invoiceId"+"}", url.PathEscape(parameterValueToString(r.invoiceId, "invoiceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.invoiceRenderRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiInvoiceSendByRequest struct {
	ctx context.Context
	ApiService *InvoiceApiService
	invoiceId int32
	invoiceSendByRequest *InvoiceSendByRequest
}

// Specify the send type
func (r ApiInvoiceSendByRequest) InvoiceSendByRequest(invoiceSendByRequest InvoiceSendByRequest) ApiInvoiceSendByRequest {
	r.invoiceSendByRequest = &invoiceSendByRequest
	return r
}

func (r ApiInvoiceSendByRequest) Execute() (*InvoiceSendBy200Response, *http.Response, error) {
	return r.ApiService.InvoiceSendByExecute(r)
}

/*
InvoiceSendBy Mark invoice as sent

Marks an invoice as sent by a chosen send type.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param invoiceId ID of invoice to mark as sent
 @return ApiInvoiceSendByRequest
*/
func (a *InvoiceApiService) InvoiceSendBy(ctx context.Context, invoiceId int32) ApiInvoiceSendByRequest {
	return ApiInvoiceSendByRequest{
		ApiService: a,
		ctx: ctx,
		invoiceId: invoiceId,
	}
}

// Execute executes the request
//  @return InvoiceSendBy200Response
func (a *InvoiceApiService) InvoiceSendByExecute(r ApiInvoiceSendByRequest) (*InvoiceSendBy200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InvoiceSendBy200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoiceApiService.InvoiceSendBy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Invoice/{invoiceId}/sendBy"
	localVarPath = strings.Replace(localVarPath, "{"+"invoiceId"+"}", url.PathEscape(parameterValueToString(r.invoiceId, "invoiceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.invoiceSendByRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSendInvoiceViaEMailRequest struct {
	ctx context.Context
	ApiService *InvoiceApiService
	invoiceId int32
	sendInvoiceViaEMailRequest *SendInvoiceViaEMailRequest
}

// Mail data
func (r ApiSendInvoiceViaEMailRequest) SendInvoiceViaEMailRequest(sendInvoiceViaEMailRequest SendInvoiceViaEMailRequest) ApiSendInvoiceViaEMailRequest {
	r.sendInvoiceViaEMailRequest = &sendInvoiceViaEMailRequest
	return r
}

func (r ApiSendInvoiceViaEMailRequest) Execute() (*SendInvoiceViaEMail201Response, *http.Response, error) {
	return r.ApiService.SendInvoiceViaEMailExecute(r)
}

/*
SendInvoiceViaEMail Send invoice via email

This endpoint sends the specified invoice to a customer via email.<br>
    This will automatically mark the invoice as sent.<br>
    Please note, that in production an invoice is not allowed to be changed after this happened!

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param invoiceId ID of invoice to be sent via email
 @return ApiSendInvoiceViaEMailRequest
*/
func (a *InvoiceApiService) SendInvoiceViaEMail(ctx context.Context, invoiceId int32) ApiSendInvoiceViaEMailRequest {
	return ApiSendInvoiceViaEMailRequest{
		ApiService: a,
		ctx: ctx,
		invoiceId: invoiceId,
	}
}

// Execute executes the request
//  @return SendInvoiceViaEMail201Response
func (a *InvoiceApiService) SendInvoiceViaEMailExecute(r ApiSendInvoiceViaEMailRequest) (*SendInvoiceViaEMail201Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SendInvoiceViaEMail201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoiceApiService.SendInvoiceViaEMail")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Invoice/{invoiceId}/sendViaEmail"
	localVarPath = strings.Replace(localVarPath, "{"+"invoiceId"+"}", url.PathEscape(parameterValueToString(r.invoiceId, "invoiceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.sendInvoiceViaEMailRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
