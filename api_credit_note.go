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
)


// CreditNoteApiService CreditNoteApi service
type CreditNoteApiService service

type ApiBookCreditNoteRequest struct {
	ctx context.Context
	ApiService *CreditNoteApiService
	creditNoteId int32
	bookCreditNoteRequest *BookCreditNoteRequest
}

// Booking data
func (r ApiBookCreditNoteRequest) BookCreditNoteRequest(bookCreditNoteRequest BookCreditNoteRequest) ApiBookCreditNoteRequest {
	r.bookCreditNoteRequest = &bookCreditNoteRequest
	return r
}

func (r ApiBookCreditNoteRequest) Execute() (*BookInvoice200Response, *http.Response, error) {
	return r.ApiService.BookCreditNoteExecute(r)
}

/*
BookCreditNote Book a credit note

Booking the credit note with a transaction is probably the most important part in the bookkeeping process.<br> There are several ways on correctly booking a credit note, all by using the same endpoint.<br> Conveniently, the booking process is exactly the same as the process for invoices and vouchers.<br> For this reason, you can have a look at it in the <a href='https://api.sevdesk.de/#section/How-to-book-an-invoice'>invoice chapter</a> and all you need to do is to change "Invoice" into "CreditNote" in the URL.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param creditNoteId ID of credit note to book
 @return ApiBookCreditNoteRequest
*/
func (a *CreditNoteApiService) BookCreditNote(ctx context.Context, creditNoteId int32) ApiBookCreditNoteRequest {
	return ApiBookCreditNoteRequest{
		ApiService: a,
		ctx: ctx,
		creditNoteId: creditNoteId,
	}
}

// Execute executes the request
//  @return BookInvoice200Response
func (a *CreditNoteApiService) BookCreditNoteExecute(r ApiBookCreditNoteRequest) (*BookInvoice200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BookInvoice200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CreditNoteApiService.BookCreditNote")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/CreditNote/{creditNoteId}/bookAmount"
	localVarPath = strings.Replace(localVarPath, "{"+"creditNoteId"+"}", url.PathEscape(parameterValueToString(r.creditNoteId, "creditNoteId")), -1)

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
	localVarPostBody = r.bookCreditNoteRequest
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

type ApiCreatecreditNoteRequest struct {
	ctx context.Context
	ApiService *CreditNoteApiService
	saveCreditNote *SaveCreditNote
}

// Creation data. Please be aware, that you need to provide at least all required parameter      of the credit note model!
func (r ApiCreatecreditNoteRequest) SaveCreditNote(saveCreditNote SaveCreditNote) ApiCreatecreditNoteRequest {
	r.saveCreditNote = &saveCreditNote
	return r
}

func (r ApiCreatecreditNoteRequest) Execute() (*SaveCreditNoteResponse, *http.Response, error) {
	return r.ApiService.CreatecreditNoteExecute(r)
}

/*
CreatecreditNote Create a new creditNote

The list of parameters starts with the credit note array.<br> This array contains all required attributes for a complete credit note.<br> Most of the attributes are covered in the credit note attribute list, there are only two parameters standing out, namely <b>mapAll</b> and <b>objectName</b>.<br> These are just needed for our system and you always need to provide them.<br> The list of parameters then continues with the credit note position array.<br> With this array you have the possibility to add multiple positions at once.<br> In the example it only contains one position, again together with the parameters <b>mapAll</b> and <b>objectName</b>, however, you can add more credit note positions by extending the array.<br> So if you wanted to add another position, you would add the same list of parameters with an incremented array index of "1" instead of "0".<br><br> The list ends with the five parameters creditNotePosDelete, discountSave, discountDelete, takeDefaultAddress and forCashRegister.<br> They only play a minor role if you only want to create a credit note but we will shortly explain what they can do.<br> With creditNotePosDelete you have to option to delete credit note positions as this request can also be used to update credit notes.<br> Both discount parameters are deprecated and have no use for credit notes, however they need to be provided in case you want to use the following two parameters.<br> With takeDefaultAddress you can specify that the first address of the contact you are using for the credit note is taken for the credit note address attribute automatically, so you don't need to provide the address yourself.<br> Finally, the forCashRegister parameter needs to be set to <b>true</b> if your credit note is to be booked on the cash register.<br> If you want to know more about these parameters, for example if you want to use this request to update credit notes, feel free to contact our support.<br> Finally, after covering all parameters, they only important information left, is that the order of the last five attributes always needs to be kept.<br> You will also always need to provide all of them, as otherwise the request won't work properly.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreatecreditNoteRequest
*/
func (a *CreditNoteApiService) CreatecreditNote(ctx context.Context) ApiCreatecreditNoteRequest {
	return ApiCreatecreditNoteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SaveCreditNoteResponse
func (a *CreditNoteApiService) CreatecreditNoteExecute(r ApiCreatecreditNoteRequest) (*SaveCreditNoteResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SaveCreditNoteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CreditNoteApiService.CreatecreditNote")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/CreditNote/Factory/saveCreditNote"

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
	localVarPostBody = r.saveCreditNote
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

type ApiCreditNoteGetPdfRequest struct {
	ctx context.Context
	ApiService *CreditNoteApiService
	creditNoteId int32
	download *bool
	preventSendBy *bool
}

// If u want to download the pdf of the credit note.
func (r ApiCreditNoteGetPdfRequest) Download(download bool) ApiCreditNoteGetPdfRequest {
	r.download = &download
	return r
}

// Defines if u want to send the credit note.
func (r ApiCreditNoteGetPdfRequest) PreventSendBy(preventSendBy bool) ApiCreditNoteGetPdfRequest {
	r.preventSendBy = &preventSendBy
	return r
}

func (r ApiCreditNoteGetPdfRequest) Execute() (*InvoiceGetPdf200Response, *http.Response, error) {
	return r.ApiService.CreditNoteGetPdfExecute(r)
}

/*
CreditNoteGetPdf Retrieve pdf document of a credit note

Retrieves the pdf document of a credit note with additional metadata.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param creditNoteId ID of credit note from which you want the pdf
 @return ApiCreditNoteGetPdfRequest
*/
func (a *CreditNoteApiService) CreditNoteGetPdf(ctx context.Context, creditNoteId int32) ApiCreditNoteGetPdfRequest {
	return ApiCreditNoteGetPdfRequest{
		ApiService: a,
		ctx: ctx,
		creditNoteId: creditNoteId,
	}
}

// Execute executes the request
//  @return InvoiceGetPdf200Response
func (a *CreditNoteApiService) CreditNoteGetPdfExecute(r ApiCreditNoteGetPdfRequest) (*InvoiceGetPdf200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InvoiceGetPdf200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CreditNoteApiService.CreditNoteGetPdf")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/CreditNote/{creditNoteId}/getPdf"
	localVarPath = strings.Replace(localVarPath, "{"+"creditNoteId"+"}", url.PathEscape(parameterValueToString(r.creditNoteId, "creditNoteId")), -1)

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

type ApiCreditNoteSendByRequest struct {
	ctx context.Context
	ApiService *CreditNoteApiService
	creditNoteId int32
	creditNoteSendByRequest *CreditNoteSendByRequest
}

// Specify the send type
func (r ApiCreditNoteSendByRequest) CreditNoteSendByRequest(creditNoteSendByRequest CreditNoteSendByRequest) ApiCreditNoteSendByRequest {
	r.creditNoteSendByRequest = &creditNoteSendByRequest
	return r
}

func (r ApiCreditNoteSendByRequest) Execute() (*CreditNoteSendBy200Response, *http.Response, error) {
	return r.ApiService.CreditNoteSendByExecute(r)
}

/*
CreditNoteSendBy Mark credit note as sent

Marks an credit note as sent by a chosen send type.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param creditNoteId ID of credit note to mark as sent
 @return ApiCreditNoteSendByRequest
*/
func (a *CreditNoteApiService) CreditNoteSendBy(ctx context.Context, creditNoteId int32) ApiCreditNoteSendByRequest {
	return ApiCreditNoteSendByRequest{
		ApiService: a,
		ctx: ctx,
		creditNoteId: creditNoteId,
	}
}

// Execute executes the request
//  @return CreditNoteSendBy200Response
func (a *CreditNoteApiService) CreditNoteSendByExecute(r ApiCreditNoteSendByRequest) (*CreditNoteSendBy200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreditNoteSendBy200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CreditNoteApiService.CreditNoteSendBy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/CreditNote/{creditNoteId}/sendBy"
	localVarPath = strings.Replace(localVarPath, "{"+"creditNoteId"+"}", url.PathEscape(parameterValueToString(r.creditNoteId, "creditNoteId")), -1)

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
	localVarPostBody = r.creditNoteSendByRequest
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

type ApiDeletecreditNoteRequest struct {
	ctx context.Context
	ApiService *CreditNoteApiService
	creditNoteId int32
}

func (r ApiDeletecreditNoteRequest) Execute() (*SendCreditNoteByPrinting200Response, *http.Response, error) {
	return r.ApiService.DeletecreditNoteExecute(r)
}

/*
DeletecreditNote Deletes an creditNote

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param creditNoteId Id of creditNote resource to delete
 @return ApiDeletecreditNoteRequest
*/
func (a *CreditNoteApiService) DeletecreditNote(ctx context.Context, creditNoteId int32) ApiDeletecreditNoteRequest {
	return ApiDeletecreditNoteRequest{
		ApiService: a,
		ctx: ctx,
		creditNoteId: creditNoteId,
	}
}

// Execute executes the request
//  @return SendCreditNoteByPrinting200Response
func (a *CreditNoteApiService) DeletecreditNoteExecute(r ApiDeletecreditNoteRequest) (*SendCreditNoteByPrinting200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SendCreditNoteByPrinting200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CreditNoteApiService.DeletecreditNote")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/creditNote/{creditNoteId}"
	localVarPath = strings.Replace(localVarPath, "{"+"creditNoteId"+"}", url.PathEscape(parameterValueToString(r.creditNoteId, "creditNoteId")), -1)

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

type ApiGetCreditNotesRequest struct {
	ctx context.Context
	ApiService *CreditNoteApiService
	status *string
	creditNoteNumber *string
	startDate *int32
	endDate *int32
	contactId *int32
	contactObjectName *string
}

// Status of the CreditNote
func (r ApiGetCreditNotesRequest) Status(status string) ApiGetCreditNotesRequest {
	r.status = &status
	return r
}

// Retrieve all CreditNotes with this creditNote number
func (r ApiGetCreditNotesRequest) CreditNoteNumber(creditNoteNumber string) ApiGetCreditNotesRequest {
	r.creditNoteNumber = &creditNoteNumber
	return r
}

// Retrieve all CreditNotes with a date equal or higher
func (r ApiGetCreditNotesRequest) StartDate(startDate int32) ApiGetCreditNotesRequest {
	r.startDate = &startDate
	return r
}

// Retrieve all CreditNotes with a date equal or lower
func (r ApiGetCreditNotesRequest) EndDate(endDate int32) ApiGetCreditNotesRequest {
	r.endDate = &endDate
	return r
}

// Retrieve all CreditNotes with this contact. Must be provided with contact[objectName]
func (r ApiGetCreditNotesRequest) ContactId(contactId int32) ApiGetCreditNotesRequest {
	r.contactId = &contactId
	return r
}

// Only required if contact[id] was provided. &#39;Contact&#39; should be used as value.
func (r ApiGetCreditNotesRequest) ContactObjectName(contactObjectName string) ApiGetCreditNotesRequest {
	r.contactObjectName = &contactObjectName
	return r
}

func (r ApiGetCreditNotesRequest) Execute() (*SendCreditNoteByPrinting200Response, *http.Response, error) {
	return r.ApiService.GetCreditNotesExecute(r)
}

/*
GetCreditNotes Retrieve CreditNote

There are a multitude of parameter which can be used to filter.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCreditNotesRequest
*/
func (a *CreditNoteApiService) GetCreditNotes(ctx context.Context) ApiGetCreditNotesRequest {
	return ApiGetCreditNotesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SendCreditNoteByPrinting200Response
func (a *CreditNoteApiService) GetCreditNotesExecute(r ApiGetCreditNotesRequest) (*SendCreditNoteByPrinting200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SendCreditNoteByPrinting200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CreditNoteApiService.GetCreditNotes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/CreditNote"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "")
	}
	if r.creditNoteNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "creditNoteNumber", r.creditNoteNumber, "")
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

type ApiGetcreditNoteByIdRequest struct {
	ctx context.Context
	ApiService *CreditNoteApiService
	creditNoteId int32
}

func (r ApiGetcreditNoteByIdRequest) Execute() (*SendCreditNoteByPrinting200Response, *http.Response, error) {
	return r.ApiService.GetcreditNoteByIdExecute(r)
}

/*
GetcreditNoteById Find creditNote by ID

Returns a single creditNote

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param creditNoteId ID of creditNote to return
 @return ApiGetcreditNoteByIdRequest
*/
func (a *CreditNoteApiService) GetcreditNoteById(ctx context.Context, creditNoteId int32) ApiGetcreditNoteByIdRequest {
	return ApiGetcreditNoteByIdRequest{
		ApiService: a,
		ctx: ctx,
		creditNoteId: creditNoteId,
	}
}

// Execute executes the request
//  @return SendCreditNoteByPrinting200Response
func (a *CreditNoteApiService) GetcreditNoteByIdExecute(r ApiGetcreditNoteByIdRequest) (*SendCreditNoteByPrinting200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SendCreditNoteByPrinting200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CreditNoteApiService.GetcreditNoteById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/creditNote/{creditNoteId}"
	localVarPath = strings.Replace(localVarPath, "{"+"creditNoteId"+"}", url.PathEscape(parameterValueToString(r.creditNoteId, "creditNoteId")), -1)

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

type ApiSendCreditNoteByPrintingRequest struct {
	ctx context.Context
	ApiService *CreditNoteApiService
	creditNoteId int32
	sendType *string
}

// the type you want to print.
func (r ApiSendCreditNoteByPrintingRequest) SendType(sendType string) ApiSendCreditNoteByPrintingRequest {
	r.sendType = &sendType
	return r
}

func (r ApiSendCreditNoteByPrintingRequest) Execute() (*SendCreditNoteByPrinting200Response, *http.Response, error) {
	return r.ApiService.SendCreditNoteByPrintingExecute(r)
}

/*
SendCreditNoteByPrinting Send credit note by printing

Sending a credit note to end-customers is an important part of the bookkeeping process.<br> Depending on the way you want to send the credit note, you need to use different endpoints.<br> Let's start with just printing out the credit note, meaning we only need to render the pdf.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param creditNoteId ID of creditNote to return
 @return ApiSendCreditNoteByPrintingRequest
*/
func (a *CreditNoteApiService) SendCreditNoteByPrinting(ctx context.Context, creditNoteId int32) ApiSendCreditNoteByPrintingRequest {
	return ApiSendCreditNoteByPrintingRequest{
		ApiService: a,
		ctx: ctx,
		creditNoteId: creditNoteId,
	}
}

// Execute executes the request
//  @return SendCreditNoteByPrinting200Response
func (a *CreditNoteApiService) SendCreditNoteByPrintingExecute(r ApiSendCreditNoteByPrintingRequest) (*SendCreditNoteByPrinting200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SendCreditNoteByPrinting200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CreditNoteApiService.SendCreditNoteByPrinting")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/creditNote/{creditNoteId}/sendByWithRender"
	localVarPath = strings.Replace(localVarPath, "{"+"creditNoteId"+"}", url.PathEscape(parameterValueToString(r.creditNoteId, "creditNoteId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.sendType == nil {
		return localVarReturnValue, nil, reportError("sendType is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "sendType", r.sendType, "")
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

type ApiSendCreditNoteViaEMailRequest struct {
	ctx context.Context
	ApiService *CreditNoteApiService
	creditNoteId int32
	sendInvoiceViaEMailRequest *SendInvoiceViaEMailRequest
}

// Mail data
func (r ApiSendCreditNoteViaEMailRequest) SendInvoiceViaEMailRequest(sendInvoiceViaEMailRequest SendInvoiceViaEMailRequest) ApiSendCreditNoteViaEMailRequest {
	r.sendInvoiceViaEMailRequest = &sendInvoiceViaEMailRequest
	return r
}

func (r ApiSendCreditNoteViaEMailRequest) Execute() (*SendCreditNoteByPrinting200Response, *http.Response, error) {
	return r.ApiService.SendCreditNoteViaEMailExecute(r)
}

/*
SendCreditNoteViaEMail Send credit note via email

This endpoint sends the specified credit note to a customer via email.<br>
    This will automatically mark the credit note as sent.<br>
    Please note, that in production an credit note is not allowed to be changed after this happened!

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param creditNoteId ID of credit note to be sent via email
 @return ApiSendCreditNoteViaEMailRequest
*/
func (a *CreditNoteApiService) SendCreditNoteViaEMail(ctx context.Context, creditNoteId int32) ApiSendCreditNoteViaEMailRequest {
	return ApiSendCreditNoteViaEMailRequest{
		ApiService: a,
		ctx: ctx,
		creditNoteId: creditNoteId,
	}
}

// Execute executes the request
//  @return SendCreditNoteByPrinting200Response
func (a *CreditNoteApiService) SendCreditNoteViaEMailExecute(r ApiSendCreditNoteViaEMailRequest) (*SendCreditNoteByPrinting200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SendCreditNoteByPrinting200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CreditNoteApiService.SendCreditNoteViaEMail")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/CreditNote/{creditNoteId}/sendViaEmail"
	localVarPath = strings.Replace(localVarPath, "{"+"creditNoteId"+"}", url.PathEscape(parameterValueToString(r.creditNoteId, "creditNoteId")), -1)

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

type ApiUpdatecreditNoteRequest struct {
	ctx context.Context
	ApiService *CreditNoteApiService
	creditNoteId int32
	modelCreditNoteUpdate *ModelCreditNoteUpdate
}

// Update data
func (r ApiUpdatecreditNoteRequest) ModelCreditNoteUpdate(modelCreditNoteUpdate ModelCreditNoteUpdate) ApiUpdatecreditNoteRequest {
	r.modelCreditNoteUpdate = &modelCreditNoteUpdate
	return r
}

func (r ApiUpdatecreditNoteRequest) Execute() (*SendCreditNoteByPrinting200Response, *http.Response, error) {
	return r.ApiService.UpdatecreditNoteExecute(r)
}

/*
UpdatecreditNote Update an existing creditNote

Update a creditNote

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param creditNoteId ID of creditNote to update
 @return ApiUpdatecreditNoteRequest
*/
func (a *CreditNoteApiService) UpdatecreditNote(ctx context.Context, creditNoteId int32) ApiUpdatecreditNoteRequest {
	return ApiUpdatecreditNoteRequest{
		ApiService: a,
		ctx: ctx,
		creditNoteId: creditNoteId,
	}
}

// Execute executes the request
//  @return SendCreditNoteByPrinting200Response
func (a *CreditNoteApiService) UpdatecreditNoteExecute(r ApiUpdatecreditNoteRequest) (*SendCreditNoteByPrinting200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SendCreditNoteByPrinting200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CreditNoteApiService.UpdatecreditNote")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/creditNote/{creditNoteId}"
	localVarPath = strings.Replace(localVarPath, "{"+"creditNoteId"+"}", url.PathEscape(parameterValueToString(r.creditNoteId, "creditNoteId")), -1)

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
	localVarPostBody = r.modelCreditNoteUpdate
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
