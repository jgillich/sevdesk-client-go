# Go API client for sevdesk

<b>Contact:</b> To contact our support click <a href='https://landing.sevdesk.de/service-support-center-technik'>here</a><br><br> 
# General information
Welcome to our API!<br>
sevDesk offers you the possibility of retrieving data using an interface, namely the sevDesk API, and making changes without having to use the web UI. The sevDesk interface is a REST-Full API. All sevDesk data and functions that are used in the web UI can also be controlled through the API.

# Cross-Origin Resource Sharing
This API features Cross-Origin Resource Sharing (CORS).<br>
It enables cross-domain communication from the browser.<br>
All responses have a wildcard same-origin which makes them completely public and accessible to everyone, including any code on any site.

# Embedding resources
When retrieving resources by using this API, you might encounter nested resources in the resources you requested.<br>
For example, an invoice always contains a contact, of which you can see the ID and the object name.<br>
This API gives you the possibility to embed these resources completely into the resources you originally requested.<br>
Taking our invoice example, this would mean, that you would not only see the ID and object name of a contact, but rather the complete contact resource.

To embed resources, all you need to do is to add the query parameter 'embed' to your GET request.<br>
As values, you can provide the name of the nested resource.<br>
Multiple nested resources are also possible by providing multiple names, separated by a comma.
 
# Authentication and Authorization
 The sevDesk API uses a token authentication to authorize calls. For this purpose every sevDesk administrator has one API token, which is a <b>hexadecimal string</b> containing <b>32 characters</b>. The following clip shows where you can find the API token if this is your first time with our API.<br><br> <video src='OpenAPI/img/findingTheApiToken.mp4' controls width='640' height='360'> Ihr Browser kann dieses Video nicht wiedergeben.<br/> Dieses Video zeigt wie sie Ihr sevDesk API Token finden. </video> <br> The token will be needed in every request that you want to send and needs to be attached to the request url as a <b>Query Parameter</b><br> or provided as a value of an <b>Authorization Header</b>.<br> For security reasons, we suggest putting the API Token in the Authorization Header and not in the query string.<br> However, in the request examples in this documentation, we will keep it in the query string, as it is easier for you to copy them and try them yourself.<br> The following url is an example that shows where your token needs to be placed as a query parameter.<br> In this case, we used some random API token. <ul> <li><span>ht</span>tps://my.sevdesk.de/api/v1/Contact?token=<span style='color:red'>b7794de0085f5cd00560f160f290af38</span></li> </ul> The next example shows the token in the Authorization Header. <ul> <li>\"Authorization\" :<span style='color:red'>\"b7794de0085f5cd00560f160f290af38\"</span></li> </ul> The api tokens have an infinite lifetime and, in other words, exist as long as the sevDesk user exists.<br> For this reason, the user should <b>NEVER</b> be deleted.<br> If really necessary, it is advisable to save the api token as we will <b>NOT</b> be able to retrieve it afterwards!<br> It is also possible to generate a new API token, for example, if you want to prevent the usage of your sevDesk account by other people who got your current API token.<br> To achieve this, you just need to click on the \"generate new\" symbol to the right of your token and confirm it with your password. 
# API News
 To never miss API news and updates again, subscribe to our <b>free API newsletter</b> and get all relevant information to keep your sevDesk software running smoothly. To subscribe, simply click <a href = 'https://landing.sevdesk.de/api-newsletter'><b>here</b></a> and confirm the email address to which we may send all updates relevant to you. 
# API Requests
 In our case, REST API requests need to be build by combining the following components. <table> <tr> <th><b>Component</b></th> <th><b>Description</b></th> </tr> <tr> <td>HTTP-Methods</td> <td> <ul> <li>GET (retrieve a resource)</li> <li>POST (create a resource)</li> <li>PUT (update a resource)</li> <li>DELETE (delete a resource)</li> </ul> </td> </tr> <tr> <td>URL of the API</td> <td><span style='color: #2aa198'>ht</span><span style='color: #2aa198'>tps://my.sevdesk.de/api/v1</span></td> </tr> <tr> <td>URI of the resource</td> <td>The resource to query.<br>For example contacts in sevDesk:<br><br> <span style='color:red'>/Contact</span><br><br> Which will result in the following complete URL:<br><br> <span style='color: #2aa198'>ht</span><span style='color: #2aa198'>tps://my.sevdesk.de/api/v1</span><span style='color:red'>/Contact</span> </td> </tr> <tr> <td>Query parameters</td> <td>Are attached by using the connectives <b>?</b> and <b>&</b> in the URL.<br></td> </tr> <tr> <td>Request headers</td> <td>Typical request headers are for example:<br> <div> <br> <ul> <li>Content-type</li> <li>Authorization</li> <li>Accept-Encoding</li> <li>User-Agent</li> <li>...</li> </ul> </div> </td> </tr> <tr> <td>Request body</td> <td>Mostly required in POST and PUT requests.<br> Often the request body contains json, in our case, it also accepts url-encoded data. </td> </tr> </table><br> <span style='color:red'>Note</span>: please pass a meaningful entry at the header \"User-Agent\". If the \"User-Agent\" is set meaningfully, we can offer better support in case of queries from customers.<br> An example how such a \"User-Agent\" can look like: \"Integration-name by abc\". <br><br> This is a sample request for retrieving existing contacts in sevDesk using curl:<br><br> <img src='OpenAPI/img/GETRequest.PNG' alt='Get Request' height='150'><br><br> As you can see, the request contains all the components mentioned above.<br> It's HTTP method is GET, it has a correct endpoint (<span style='color: #2aa198'>ht</span><span style='color: #2aa198'>tps://my.sevdesk.de/api/v1</span><span style='color:red'>/Contact</span>), query parameters like <b>token</b> and additional <b>header</b> information!<br><br> <b>Query Parameters</b><br><br> As you might have seen in the sample request above, there are several other parameters besides \"token\", located in the url.<br> Those are mostly optional but prove to be very useful for a lot of requests as they can limit, extend, sort or filter the data you will get as a response.<br><br> These are the three most used query parameter for the sevDesk API. <table> <tr> <th><b>Parameter</b></th> <th><b>Description</b></th> </tr> <tr> <td>limit</td> <td>Limits the number of entries that are returned.<br> Most useful in GET requests which will most likely deliver big sets of data like country or currency lists.<br> In this case, you can bypass the default limitation on returned entries by providing a high number. </td> </tr> <tr> <td>offset</td> <td>Specifies a certain offset for the data that will be returned.<br> As an example, you can specify \"offset=2\" if you want all entries except for the first two.</td> </tr> <tr> <td>embed</td> <td>Will extend some of the returned data.<br> A brief example can be found below.</td> </tr> </table> This is an example for the usage of the embed parameter.<br> The following first request will return all company contact entries in sevDesk up to a limit of 100 without any additional information and no offset.<br><br> <img src='OpenAPI/img/ContactQueryWithoutEmbed.PNG' width='900' height='850'><br><br> Now have a look at the category attribute located in the response.<br> Naturally, it just contains the id and the object name of the object but no further information.<br> We will now use the parameter embed with the value \"category\".<br><br> <img src='OpenAPI/img/ContactQueryWithEmbed.PNG' width='900' height='850'><br><br> As you can see, the category object is now extended and shows all the attributes and their corresponding values.<br><br> There are lot of other query parameters that can be used to filter the returned data for objects that match a certain pattern, however, those will not be mentioned here and instead can be found at the detail documentation of the most used API endpoints like contact, invoice or voucher.<br><br> <b>Request Headers</b><br><br> The HTTP request (response) headers allow the client as well as the server to pass additional information with the request.<br> They transfer the parameters and arguments which are important for transmitting data over HTTP.<br> Three headers which are useful / necessary when using the sevDesk API are \"Authorization, \"Accept\" and \"Content-type\".<br> Underneath is a brief description of why and how they should be used.<br><br> <b>Authorization</b><br><br> Can be used if you want to provide your API token in the header instead of having it in the url. <ul> <li>Authorization:<span style='color:red'>yourApiToken</span></li> </ul> <b>Accept</b><br><br> Specifies the format of the response.<br> Required for operations with a response body. <ul> <li>Accept:application/<span style='color:red'>format</span> </li> </ul> In our case, <code><span style='color:red'>format</span></code> could be replaced with <code>json</code> or <code>xml</code><br><br> <b>Content-type</b><br><br> Specifies which format is used in the request.<br> Is required for operations with a request body. <ul> <li>Content-type:application/<span style='color:red'>format</span></li> </ul> In our case,<code><span style='color:red'>format</span></code>could be replaced with <code>json</code> or <code>x-www-form-urlencoded</code> <br><br><b>API Responses</b><br><br> HTTP status codes<br> When calling the sevDesk API it is very likely that you will get a HTTP status code in the response.<br> Some API calls will also return JSON response bodies which will contain information about the resource.<br> Each status code which is returned will either be a <b>success</b> code or an <b>error</b> code.<br><br> Success codes <table> <tr> <th><b>Status code</b></th> <th><b>Description</b></th> </tr> <tr> <td><code>200 OK</code></td> <td>The request was successful</td> </tr> <tr> <td><code>201 Created</code></td> <td>Most likely to be found in the response of a <b>POST</b> request.<br> This code indicates that the desired resource was successfully created.</td> </tr> </table> <br>Error codes <table> <tr> <th><b>Status code</b></th> <th><b>Description</b></th> </tr> <tr> <td><code>400 Bad request</code></td> <td>The request you sent is most likely syntactically incorrect.<br> You should check if the parameters in the request body or the url are correct.</td> </tr> <tr> <td><code>401 Unauthorized</code></td> <td>The authentication failed.<br> Most likely caused by a missing or wrong API token.</td> </tr> <tr> <td><code>403 Forbidden</code></td> <td>You do not have the permission the access the resource which is requested.</td> </tr> <tr> <td><code>404 Not found</code></td> <td>The resource you specified does not exist.</td> </tr> <tr> <td><code>500 Internal server error</code></td> <td>An internal server error has occurred.<br> Normally this means that something went wrong on our side.<br> However, sometimes this error will appear if we missed to catch an error which is normally a 400 status code! </td> </tr> </table> 
# Your First Request
 After reading the introduction to our API, you should now be able to make your first call.<br> For testing our API, we would always recommend to create a trial account for sevDesk to prevent unwanted changes to your main account.<br> A trial account will be in the highest tariff (materials management), so every sevDesk function can be tested! <br><br>To start testing we would recommend one of the following tools: <ul> <li><a href='https://www.getpostman.com/'>Postman</a></li> <li><a href='https://curl.haxx.se/download.html'>cURL</a></li> </ul> This example will illustrate your first request, which is creating a new Contact in sevDesk.<br> <ol> <li>Download <a href='https://www.getpostman.com/'><b>Postman</b></a> for your desired system and start the application</li> <li>Enter <span><b>ht</span>tps://my.sevdesk.de/api/v1/Contact</b> as the url</li> <li>Use the connective <b>?</b> to append <b>token=</b> to the end of the url, or create an authorization header. Insert your API token as the value</li> <li>For this test, select <b>POST</b> as the HTTP method</li> <li>Go to <b>Headers</b> and enter the key-value pair <b>Content-type</b> + <b>application/x-www-form-urlencoded</b><br> As an alternative, you can just go to <b>Body</b> and select <b>x-www-form-urlencoded</b></li> <li>Now go to <b>Body</b> (if you are not there yet) and enter the key-value pairs as shown in the following picture<br><br> <img src='OpenAPI/img/FirstRequestPostman.PNG' width='900'><br><br></li> <li>Click on <b>Send</b>. Your response should now look like this:<br><br> <img src='OpenAPI/img/FirstRequestResponse.PNG' width='900'></li> </ol> As you can see, a successful response in this case returns a JSON-formatted response body containing the contact you just created.<br> For keeping it simple, this was only a minimal example of creating a contact.<br> There are however numerous combinations of parameters that you can provide which add information to your contact.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sevdesk "github.com/jgillich/sevdesk-client-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sevdesk.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sevdesk.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), sevdesk.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sevdesk.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://my.sevdesk.de/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountingContactApi* | [**CreateAccountingContact**](docs/AccountingContactApi.md#createaccountingcontact) | **Post** /AccountingContact | Create a new accounting contact
*AccountingContactApi* | [**DeleteAccountingContact**](docs/AccountingContactApi.md#deleteaccountingcontact) | **Delete** /AccountingContact/{accountingContactId} | Deletes an accounting contact
*AccountingContactApi* | [**GetAccountingContact**](docs/AccountingContactApi.md#getaccountingcontact) | **Get** /AccountingContact | Retrieve accounting contact
*AccountingContactApi* | [**GetAccountingContactById**](docs/AccountingContactApi.md#getaccountingcontactbyid) | **Get** /AccountingContact/{accountingContactId} | Find accounting contact by ID
*AccountingContactApi* | [**UpdateAccountingContact**](docs/AccountingContactApi.md#updateaccountingcontact) | **Put** /AccountingContact/{accountingContactId} | Update an existing accounting contact
*CheckAccountApi* | [**CreateCheckAccount**](docs/CheckAccountApi.md#createcheckaccount) | **Post** /CheckAccount | Create a new check account
*CheckAccountApi* | [**DeleteCheckAccount**](docs/CheckAccountApi.md#deletecheckaccount) | **Delete** /CheckAccount/{checkAccountId} | Deletes a check account
*CheckAccountApi* | [**GetBalanceAtDate**](docs/CheckAccountApi.md#getbalanceatdate) | **Get** /CheckAccount/{checkAccountId}/getBalanceAtDate | Get the balance at a given date
*CheckAccountApi* | [**GetCheckAccountById**](docs/CheckAccountApi.md#getcheckaccountbyid) | **Get** /CheckAccount/{checkAccountId} | Find check account by ID
*CheckAccountApi* | [**GetCheckAccounts**](docs/CheckAccountApi.md#getcheckaccounts) | **Get** /CheckAccount | Retrieve check accounts
*CheckAccountApi* | [**UpdateCheckAccount**](docs/CheckAccountApi.md#updatecheckaccount) | **Put** /CheckAccount/{checkAccountId} | Update an existing check account
*CheckAccountTransactionApi* | [**CreateTransaction**](docs/CheckAccountTransactionApi.md#createtransaction) | **Post** /CheckAccountTransaction | Create a new transaction
*CheckAccountTransactionApi* | [**DeleteCheckAccountTransaction**](docs/CheckAccountTransactionApi.md#deletecheckaccounttransaction) | **Delete** /CheckAccountTransaction/{checkAccountTransactionId} | Deletes a check account transaction
*CheckAccountTransactionApi* | [**GetCheckAccountTransactionById**](docs/CheckAccountTransactionApi.md#getcheckaccounttransactionbyid) | **Get** /CheckAccountTransaction/{checkAccountTransactionId} | Find check account transaction by ID
*CheckAccountTransactionApi* | [**GetTransactions**](docs/CheckAccountTransactionApi.md#gettransactions) | **Get** /CheckAccountTransaction | Retrieve transactions
*CheckAccountTransactionApi* | [**UpdateCheckAccountTransaction**](docs/CheckAccountTransactionApi.md#updatecheckaccounttransaction) | **Put** /CheckAccountTransaction/{checkAccountTransactionId} | Update an existing check account transaction
*CommunicationWayApi* | [**CreateCommunicationWay**](docs/CommunicationWayApi.md#createcommunicationway) | **Post** /CommunicationWay | Create a new contact communication way
*CommunicationWayApi* | [**DeleteCommunicationWay**](docs/CommunicationWayApi.md#deletecommunicationway) | **Delete** /CommunicationWay/{communicationWayId} | Deletes a communication way
*CommunicationWayApi* | [**GetCommunicationWayById**](docs/CommunicationWayApi.md#getcommunicationwaybyid) | **Get** /CommunicationWay/{communicationWayId} | Find communication way by ID
*CommunicationWayApi* | [**GetCommunicationWayKeys**](docs/CommunicationWayApi.md#getcommunicationwaykeys) | **Get** /CommunicationWayKey | Retrieve communication way keys
*CommunicationWayApi* | [**GetCommunicationWays**](docs/CommunicationWayApi.md#getcommunicationways) | **Get** /CommunicationWay | Retrieve communication ways
*CommunicationWayApi* | [**UpdateCommunicationWay**](docs/CommunicationWayApi.md#updatecommunicationway) | **Put** /CommunicationWay/{communicationWayId} | Update a existing communication way
*ContactApi* | [**ContactCustomerNumberAvailabilityCheck**](docs/ContactApi.md#contactcustomernumberavailabilitycheck) | **Get** /Contact/Mapper/checkCustomerNumberAvailability | Check if a customer number is available
*ContactApi* | [**CreateContact**](docs/ContactApi.md#createcontact) | **Post** /Contact | Create a new contact
*ContactApi* | [**DeleteContact**](docs/ContactApi.md#deletecontact) | **Delete** /Contact/{contactId} | Deletes a contact
*ContactApi* | [**FindContactsByCustomFieldValue**](docs/ContactApi.md#findcontactsbycustomfieldvalue) | **Get** /Contact/Factory/findContactsByCustomFieldValue | Find contacts by custom field value
*ContactApi* | [**GetContactById**](docs/ContactApi.md#getcontactbyid) | **Get** /Contact/{contactId} | Find contact by ID
*ContactApi* | [**GetContactTabsItemCountById**](docs/ContactApi.md#getcontacttabsitemcountbyid) | **Get** /Contact/{contactId}/getTabsItemCount | Get number of all items
*ContactApi* | [**GetContacts**](docs/ContactApi.md#getcontacts) | **Get** /Contact | Retrieve contacts
*ContactApi* | [**GetNextCustomerNumber**](docs/ContactApi.md#getnextcustomernumber) | **Get** /Contact/Factory/getNextCustomerNumber | Get next free customer number
*ContactApi* | [**UpdateContact**](docs/ContactApi.md#updatecontact) | **Put** /Contact/{contactId} | Update a existing contact
*ContactAddressApi* | [**ContactAddressId**](docs/ContactAddressApi.md#contactaddressid) | **Get** /ContactAddress/{contactAddressId} | Find contact address by ID
*ContactAddressApi* | [**CreateContactAddress**](docs/ContactAddressApi.md#createcontactaddress) | **Post** /ContactAddress | Create a new contact address
*ContactAddressApi* | [**DeleteContactAddress**](docs/ContactAddressApi.md#deletecontactaddress) | **Delete** /ContactAddress/{contactAddressId} | Deletes a contact address
*ContactAddressApi* | [**GetContactAddresses**](docs/ContactAddressApi.md#getcontactaddresses) | **Get** /ContactAddress | Retrieve contact addresses
*ContactAddressApi* | [**UpdateContactAddress**](docs/ContactAddressApi.md#updatecontactaddress) | **Put** /ContactAddress/{contactAddressId} | update a existing contact address
*ContactFieldApi* | [**CreateContactField**](docs/ContactFieldApi.md#createcontactfield) | **Post** /ContactCustomField | Create contact field
*ContactFieldApi* | [**CreateContactFieldSetting**](docs/ContactFieldApi.md#createcontactfieldsetting) | **Post** /ContactCustomFieldSetting | Create contact field setting
*ContactFieldApi* | [**DeleteContactCustomFieldId**](docs/ContactFieldApi.md#deletecontactcustomfieldid) | **Delete** /ContactCustomField/{contactCustomFieldId} | delete a contact field
*ContactFieldApi* | [**DeleteContactFieldSetting**](docs/ContactFieldApi.md#deletecontactfieldsetting) | **Delete** /ContactCustomFieldSetting/{contactCustomFieldSettingId} | Deletes a contact field setting
*ContactFieldApi* | [**GetContactFieldSettingById**](docs/ContactFieldApi.md#getcontactfieldsettingbyid) | **Get** /ContactCustomFieldSetting/{contactCustomFieldSettingId} | Find contact field setting by ID
*ContactFieldApi* | [**GetContactFieldSettings**](docs/ContactFieldApi.md#getcontactfieldsettings) | **Get** /ContactCustomFieldSetting | Retrieve contact field settings
*ContactFieldApi* | [**GetContactFields**](docs/ContactFieldApi.md#getcontactfields) | **Get** /ContactCustomField | Retrieve contact fields
*ContactFieldApi* | [**GetContactFieldsById**](docs/ContactFieldApi.md#getcontactfieldsbyid) | **Get** /ContactCustomField/{contactCustomFieldId} | Retrieve contact fields
*ContactFieldApi* | [**GetPlaceholder**](docs/ContactFieldApi.md#getplaceholder) | **Get** /Textparser/fetchDictionaryEntriesByType | Retrieve Placeholders
*ContactFieldApi* | [**GetReferenceCount**](docs/ContactFieldApi.md#getreferencecount) | **Get** /ContactCustomFieldSetting/{contactCustomFieldSettingId}/getReferenceCount | Receive count reference
*ContactFieldApi* | [**UpdateContactFieldSetting**](docs/ContactFieldApi.md#updatecontactfieldsetting) | **Put** /ContactCustomFieldSetting/{contactCustomFieldSettingId} | Update contact field setting
*ContactFieldApi* | [**UpdateContactfield**](docs/ContactFieldApi.md#updatecontactfield) | **Put** /ContactCustomField/{contactCustomFieldId} | Update a contact field
*CreditNoteApi* | [**BookCreditNote**](docs/CreditNoteApi.md#bookcreditnote) | **Put** /CreditNote/{creditNoteId}/bookAmount | Book a credit note
*CreditNoteApi* | [**CreatecreditNote**](docs/CreditNoteApi.md#createcreditnote) | **Post** /CreditNote/Factory/saveCreditNote | Create a new creditNote
*CreditNoteApi* | [**CreditNoteGetPdf**](docs/CreditNoteApi.md#creditnotegetpdf) | **Get** /CreditNote/{creditNoteId}/getPdf | Retrieve pdf document of a credit note
*CreditNoteApi* | [**CreditNoteSendBy**](docs/CreditNoteApi.md#creditnotesendby) | **Put** /CreditNote/{creditNoteId}/sendBy | Mark credit note as sent
*CreditNoteApi* | [**DeletecreditNote**](docs/CreditNoteApi.md#deletecreditnote) | **Delete** /creditNote/{creditNoteId} | Deletes an creditNote
*CreditNoteApi* | [**GetCreditNotes**](docs/CreditNoteApi.md#getcreditnotes) | **Get** /CreditNote | Retrieve CreditNote
*CreditNoteApi* | [**GetcreditNoteById**](docs/CreditNoteApi.md#getcreditnotebyid) | **Get** /creditNote/{creditNoteId} | Find creditNote by ID
*CreditNoteApi* | [**SendCreditNoteByPrinting**](docs/CreditNoteApi.md#sendcreditnotebyprinting) | **Get** /creditNote/{creditNoteId}/sendByWithRender | Send credit note by printing
*CreditNoteApi* | [**SendCreditNoteViaEMail**](docs/CreditNoteApi.md#sendcreditnoteviaemail) | **Post** /CreditNote/{creditNoteId}/sendViaEmail | Send credit note via email
*CreditNoteApi* | [**UpdatecreditNote**](docs/CreditNoteApi.md#updatecreditnote) | **Put** /creditNote/{creditNoteId} | Update an existing creditNote
*CreditNotePosApi* | [**GetcreditNotePositions**](docs/CreditNotePosApi.md#getcreditnotepositions) | **Get** /creditNotePos | Retrieve creditNote positions
*ExportApi* | [**ExportContact**](docs/ExportApi.md#exportcontact) | **Get** /Export/contactListCsv | Export contact
*ExportApi* | [**ExportCreditNote**](docs/ExportApi.md#exportcreditnote) | **Get** /Export/creditNoteCsv | Export creditNote
*ExportApi* | [**ExportDatev**](docs/ExportApi.md#exportdatev) | **Get** /Export/datevCSV | Export datev
*ExportApi* | [**ExportInvoice**](docs/ExportApi.md#exportinvoice) | **Get** /Export/invoiceCsv | Export invoice
*ExportApi* | [**ExportInvoiceZip**](docs/ExportApi.md#exportinvoicezip) | **Get** /Export/invoiceZip | Export Invoice as zip
*ExportApi* | [**ExportTransactions**](docs/ExportApi.md#exporttransactions) | **Get** /Export/transactionsCsv | Export transaction
*ExportApi* | [**ExportVoucher**](docs/ExportApi.md#exportvoucher) | **Get** /Export/voucherListCsv | Export voucher as zip
*ExportApi* | [**ExportVoucherZip**](docs/ExportApi.md#exportvoucherzip) | **Get** /Export/voucherZip | Export voucher zip
*InvoiceApi* | [**BookInvoice**](docs/InvoiceApi.md#bookinvoice) | **Put** /Invoice/{invoiceId}/bookAmount | Book an invoice
*InvoiceApi* | [**CancelInvoice**](docs/InvoiceApi.md#cancelinvoice) | **Post** /Invoice/{invoiceId}/cancelInvoice | Cancel an invoice / Create cancellation invoice
*InvoiceApi* | [**CreateInvoiceByFactory**](docs/InvoiceApi.md#createinvoicebyfactory) | **Post** /Invoice/Factory/saveInvoice | Create a new invoice
*InvoiceApi* | [**CreateInvoiceFromOrder**](docs/InvoiceApi.md#createinvoicefromorder) | **Post** /Invoice/Factory/createInvoiceFromOrder | Create invoice from order
*InvoiceApi* | [**CreateInvoiceReminder**](docs/InvoiceApi.md#createinvoicereminder) | **Post** /Invoice/Factory/createInvoiceReminder | Create invoice reminder
*InvoiceApi* | [**GetInvoiceById**](docs/InvoiceApi.md#getinvoicebyid) | **Get** /Invoice/{invoiceId} | Find invoice by ID
*InvoiceApi* | [**GetInvoicePositionsById**](docs/InvoiceApi.md#getinvoicepositionsbyid) | **Get** /Invoice/{invoiceId}/getPositions | Find invoice positions
*InvoiceApi* | [**GetInvoices**](docs/InvoiceApi.md#getinvoices) | **Get** /Invoice | Retrieve invoices
*InvoiceApi* | [**GetIsInvoicePartiallyPaid**](docs/InvoiceApi.md#getisinvoicepartiallypaid) | **Get** /Invoice/{invoiceId}/getIsPartiallyPaid | Check if an invoice is already partially paid
*InvoiceApi* | [**InvoiceGetPdf**](docs/InvoiceApi.md#invoicegetpdf) | **Get** /Invoice/{invoiceId}/getPdf | Retrieve pdf document of an invoice
*InvoiceApi* | [**InvoiceRender**](docs/InvoiceApi.md#invoicerender) | **Post** /Invoice/{invoiceId}/render | Render the pdf document of an invoice
*InvoiceApi* | [**InvoiceSendBy**](docs/InvoiceApi.md#invoicesendby) | **Put** /Invoice/{invoiceId}/sendBy | Mark invoice as sent
*InvoiceApi* | [**SendInvoiceViaEMail**](docs/InvoiceApi.md#sendinvoiceviaemail) | **Post** /Invoice/{invoiceId}/sendViaEmail | Send invoice via email
*InvoicePosApi* | [**GetInvoicePos**](docs/InvoicePosApi.md#getinvoicepos) | **Get** /InvoicePos | Retrieve InvoicePos
*LayoutApi* | [**GetLetterpapersWithThumb**](docs/LayoutApi.md#getletterpaperswiththumb) | **Get** /DocServer/getLetterpapersWithThumb | Retrieve letterpapers
*LayoutApi* | [**GetTemplates**](docs/LayoutApi.md#gettemplates) | **Get** /DocServer/getTemplatesWithThumb | Retrieve templates
*LayoutApi* | [**UpdateCreditNoteTemplate**](docs/LayoutApi.md#updatecreditnotetemplate) | **Put** /CreditNote/{creditNoteId}/changeParameter | Update an of credit note template
*LayoutApi* | [**UpdateInvoiceTemplate**](docs/LayoutApi.md#updateinvoicetemplate) | **Put** /Invoice/{invoiceId}/changeParameter | Update an invoice template
*LayoutApi* | [**UpdateOrderTemplate**](docs/LayoutApi.md#updateordertemplate) | **Put** /Order/{orderId}/changeParameter | Update an order template
*OrderApi* | [**CreateContractNoteFromOrder**](docs/OrderApi.md#createcontractnotefromorder) | **Post** /Order/Factory/createContractNoteFromOrder | Create contract note from order
*OrderApi* | [**CreateOrder**](docs/OrderApi.md#createorder) | **Post** /Order/Factory/saveOrder | Create a new order
*OrderApi* | [**CreatePackingListFromOrder**](docs/OrderApi.md#createpackinglistfromorder) | **Post** /Order/Factory/createPackingListFromOrder | Create packing list from order
*OrderApi* | [**DeleteOrder**](docs/OrderApi.md#deleteorder) | **Delete** /Order/{orderId} | Deletes an order
*OrderApi* | [**GetDiscounts**](docs/OrderApi.md#getdiscounts) | **Get** /Order/{orderId}/getDiscounts | Find order discounts
*OrderApi* | [**GetOrderById**](docs/OrderApi.md#getorderbyid) | **Get** /Order/{orderId} | Find order by ID
*OrderApi* | [**GetOrderPositionsById**](docs/OrderApi.md#getorderpositionsbyid) | **Get** /Order/{orderId}/getPositions | Find order positions
*OrderApi* | [**GetOrders**](docs/OrderApi.md#getorders) | **Get** /Order | Retrieve orders
*OrderApi* | [**GetRelatedObjects**](docs/OrderApi.md#getrelatedobjects) | **Get** /Order/{orderId}/getRelatedObjects | Find related objects
*OrderApi* | [**OrderGetPdf**](docs/OrderApi.md#ordergetpdf) | **Get** /Order/{orderId}/getPdf | Retrieve pdf document of an order
*OrderApi* | [**OrderSendBy**](docs/OrderApi.md#ordersendby) | **Put** /Order/{orderId}/sendBy | Mark order as sent
*OrderApi* | [**SendorderViaEMail**](docs/OrderApi.md#sendorderviaemail) | **Post** /Order/{orderId}/sendViaEmail | Send order via email
*OrderApi* | [**UpdateOrder**](docs/OrderApi.md#updateorder) | **Put** /Order/{orderId} | Update an existing order
*OrderPosApi* | [**DeleteOrderPos**](docs/OrderPosApi.md#deleteorderpos) | **Delete** /OrderPos/{orderPosId} | Deletes an order Position
*OrderPosApi* | [**GetOrderPositionById**](docs/OrderPosApi.md#getorderpositionbyid) | **Get** /OrderPos/{orderPosId} | Find order position by ID
*OrderPosApi* | [**GetOrderPositions**](docs/OrderPosApi.md#getorderpositions) | **Get** /OrderPos | Retrieve order positions
*OrderPosApi* | [**UpdateOrderPosition**](docs/OrderPosApi.md#updateorderposition) | **Put** /OrderPos/{orderPosId} | Update an existing order position
*PartApi* | [**CreatePart**](docs/PartApi.md#createpart) | **Post** /Part | Create a new part
*PartApi* | [**GetPartById**](docs/PartApi.md#getpartbyid) | **Get** /Part/{partId} | Find part by ID
*PartApi* | [**GetParts**](docs/PartApi.md#getparts) | **Get** /Part | Retrieve parts
*PartApi* | [**PartGetStock**](docs/PartApi.md#partgetstock) | **Get** /Part/{partId}/getStock | Get stock of a part
*PartApi* | [**UpdatePart**](docs/PartApi.md#updatepart) | **Put** /Part/{partId} | Update an existing part
*ReportApi* | [**ReportContact**](docs/ReportApi.md#reportcontact) | **Get** /Report/contactlist | Export contact list
*ReportApi* | [**ReportInvoice**](docs/ReportApi.md#reportinvoice) | **Get** /Report/invoicelist | Export invoice list
*ReportApi* | [**ReportOrder**](docs/ReportApi.md#reportorder) | **Get** /Report/orderlist | Export order list
*ReportApi* | [**ReportVoucher**](docs/ReportApi.md#reportvoucher) | **Get** /Report/voucherlist | Export voucher list
*TagApi* | [**CreateTag**](docs/TagApi.md#createtag) | **Post** /Tag/Factory/create | Create a new tag
*TagApi* | [**DeleteTag**](docs/TagApi.md#deletetag) | **Delete** /Tag/{tagId} | Deletes a tag
*TagApi* | [**GetTagById**](docs/TagApi.md#gettagbyid) | **Get** /Tag/{tagId} | Find tag by ID
*TagApi* | [**GetTagRelations**](docs/TagApi.md#gettagrelations) | **Get** /TagRelation | Retrieve tag relations
*TagApi* | [**GetTags**](docs/TagApi.md#gettags) | **Get** /Tag | Retrieve tags
*TagApi* | [**UpdateTag**](docs/TagApi.md#updatetag) | **Put** /Tag/{tagId} | Update tag
*VoucherApi* | [**BookVoucher**](docs/VoucherApi.md#bookvoucher) | **Put** /Voucher/{voucherId}/bookAmount | Book a voucher
*VoucherApi* | [**CreateVoucherByFactory**](docs/VoucherApi.md#createvoucherbyfactory) | **Post** /Voucher/Factory/saveVoucher | Create a new voucher
*VoucherApi* | [**GetVoucherById**](docs/VoucherApi.md#getvoucherbyid) | **Get** /Voucher/{voucherId} | Find voucher by ID
*VoucherApi* | [**GetVouchers**](docs/VoucherApi.md#getvouchers) | **Get** /Voucher | Retrieve vouchers
*VoucherApi* | [**UpdateVoucher**](docs/VoucherApi.md#updatevoucher) | **Put** /Voucher/{voucherId} | Update an existing voucher
*VoucherApi* | [**VoucherUploadFile**](docs/VoucherApi.md#voucheruploadfile) | **Post** /Voucher/Factory/uploadTempFile | Upload voucher file
*VoucherPosApi* | [**GetVoucherPositions**](docs/VoucherPosApi.md#getvoucherpositions) | **Get** /VoucherPos | Retrieve voucher positions


## Documentation For Models

 - [BookCreditNoteRequest](docs/BookCreditNoteRequest.md)
 - [BookCreditNoteRequestCheckAccountTransaction](docs/BookCreditNoteRequestCheckAccountTransaction.md)
 - [BookInvoice200Response](docs/BookInvoice200Response.md)
 - [BookInvoiceRequest](docs/BookInvoiceRequest.md)
 - [BookInvoiceRequestCheckAccountTransaction](docs/BookInvoiceRequestCheckAccountTransaction.md)
 - [BookVoucher200Response](docs/BookVoucher200Response.md)
 - [BookVoucherRequest](docs/BookVoucherRequest.md)
 - [BookVoucherRequestCheckAccount](docs/BookVoucherRequestCheckAccount.md)
 - [BookVoucherRequestCheckAccountTransaction](docs/BookVoucherRequestCheckAccountTransaction.md)
 - [ContactAddressId200Response](docs/ContactAddressId200Response.md)
 - [CreateInvoiceReminderRequest](docs/CreateInvoiceReminderRequest.md)
 - [CreateInvoiceReminderRequestInvoice](docs/CreateInvoiceReminderRequestInvoice.md)
 - [CreatePackingListFromOrder200Response](docs/CreatePackingListFromOrder200Response.md)
 - [CreateTagRequest](docs/CreateTagRequest.md)
 - [CreateTagRequestObject](docs/CreateTagRequestObject.md)
 - [CreditNoteSendBy200Response](docs/CreditNoteSendBy200Response.md)
 - [CreditNoteSendByRequest](docs/CreditNoteSendByRequest.md)
 - [DeleteAccountingContact200Response](docs/DeleteAccountingContact200Response.md)
 - [ExportContactSevQueryParameter](docs/ExportContactSevQueryParameter.md)
 - [ExportCreditNoteSevQueryParameter](docs/ExportCreditNoteSevQueryParameter.md)
 - [ExportCreditNoteSevQueryParameterFilter](docs/ExportCreditNoteSevQueryParameterFilter.md)
 - [ExportCreditNoteSevQueryParameterFilterContact](docs/ExportCreditNoteSevQueryParameterFilterContact.md)
 - [ExportInvoiceSevQueryParameter](docs/ExportInvoiceSevQueryParameter.md)
 - [ExportInvoiceSevQueryParameterFilter](docs/ExportInvoiceSevQueryParameterFilter.md)
 - [ExportInvoiceSevQueryParameterFilterContact](docs/ExportInvoiceSevQueryParameterFilterContact.md)
 - [ExportTransactionsSevQueryParameter](docs/ExportTransactionsSevQueryParameter.md)
 - [ExportTransactionsSevQueryParameterFilter](docs/ExportTransactionsSevQueryParameterFilter.md)
 - [ExportTransactionsSevQueryParameterFilterCheckAccount](docs/ExportTransactionsSevQueryParameterFilterCheckAccount.md)
 - [ExportVoucherSevQueryParameter](docs/ExportVoucherSevQueryParameter.md)
 - [ExportVoucherSevQueryParameterFilter](docs/ExportVoucherSevQueryParameterFilter.md)
 - [ExportVoucherSevQueryParameterFilterContact](docs/ExportVoucherSevQueryParameterFilterContact.md)
 - [GetAccountingContactById200Response](docs/GetAccountingContactById200Response.md)
 - [GetBalanceAtDate200Response](docs/GetBalanceAtDate200Response.md)
 - [GetCheckAccountById200Response](docs/GetCheckAccountById200Response.md)
 - [GetCheckAccountTransactionById200Response](docs/GetCheckAccountTransactionById200Response.md)
 - [GetCommunicationWayById200Response](docs/GetCommunicationWayById200Response.md)
 - [GetCommunicationWayKeys200Response](docs/GetCommunicationWayKeys200Response.md)
 - [GetContactFieldSettings200Response](docs/GetContactFieldSettings200Response.md)
 - [GetContactFieldsById200Response](docs/GetContactFieldsById200Response.md)
 - [GetContactTabsItemCountById200Response](docs/GetContactTabsItemCountById200Response.md)
 - [GetContacts200Response](docs/GetContacts200Response.md)
 - [GetDiscounts200Response](docs/GetDiscounts200Response.md)
 - [GetInvoiceById200Response](docs/GetInvoiceById200Response.md)
 - [GetInvoicePos200Response](docs/GetInvoicePos200Response.md)
 - [GetIsInvoicePartiallyPaid200Response](docs/GetIsInvoicePartiallyPaid200Response.md)
 - [GetLetterpapersWithThumb200Response](docs/GetLetterpapersWithThumb200Response.md)
 - [GetLetterpapersWithThumb200ResponseLetterpapers](docs/GetLetterpapersWithThumb200ResponseLetterpapers.md)
 - [GetNextCustomerNumber200Response](docs/GetNextCustomerNumber200Response.md)
 - [GetOrderPositions200Response](docs/GetOrderPositions200Response.md)
 - [GetPartById200Response](docs/GetPartById200Response.md)
 - [GetPlaceholder200Response](docs/GetPlaceholder200Response.md)
 - [GetPlaceholder200ResponseValueInner](docs/GetPlaceholder200ResponseValueInner.md)
 - [GetReferenceCount200Response](docs/GetReferenceCount200Response.md)
 - [GetRelatedObjects200Response](docs/GetRelatedObjects200Response.md)
 - [GetTagById200Response](docs/GetTagById200Response.md)
 - [GetTagRelations200Response](docs/GetTagRelations200Response.md)
 - [GetTemplates200Response](docs/GetTemplates200Response.md)
 - [GetTemplates200ResponseTemplates](docs/GetTemplates200ResponseTemplates.md)
 - [GetVoucherById200Response](docs/GetVoucherById200Response.md)
 - [GetVoucherPositions200Response](docs/GetVoucherPositions200Response.md)
 - [GetcreditNotePositions200Response](docs/GetcreditNotePositions200Response.md)
 - [InvoiceGetPdf200Response](docs/InvoiceGetPdf200Response.md)
 - [InvoiceRender201Response](docs/InvoiceRender201Response.md)
 - [InvoiceRender201ResponseObjectsInner](docs/InvoiceRender201ResponseObjectsInner.md)
 - [InvoiceRenderRequest](docs/InvoiceRenderRequest.md)
 - [InvoiceSendBy200Response](docs/InvoiceSendBy200Response.md)
 - [InvoiceSendByRequest](docs/InvoiceSendByRequest.md)
 - [ModelAccountingContact](docs/ModelAccountingContact.md)
 - [ModelAccountingContactContact](docs/ModelAccountingContactContact.md)
 - [ModelAccountingContactResponse](docs/ModelAccountingContactResponse.md)
 - [ModelAccountingContactResponseContact](docs/ModelAccountingContactResponseContact.md)
 - [ModelAccountingContactResponseSevClient](docs/ModelAccountingContactResponseSevClient.md)
 - [ModelAccountingContactUpdate](docs/ModelAccountingContactUpdate.md)
 - [ModelAccountingContactUpdateContact](docs/ModelAccountingContactUpdateContact.md)
 - [ModelChangeLayout](docs/ModelChangeLayout.md)
 - [ModelChangeLayoutResponse](docs/ModelChangeLayoutResponse.md)
 - [ModelChangeLayoutResponseMetadaten](docs/ModelChangeLayoutResponseMetadaten.md)
 - [ModelCheckAccount](docs/ModelCheckAccount.md)
 - [ModelCheckAccountResponse](docs/ModelCheckAccountResponse.md)
 - [ModelCheckAccountResponseSevClient](docs/ModelCheckAccountResponseSevClient.md)
 - [ModelCheckAccountSevClient](docs/ModelCheckAccountSevClient.md)
 - [ModelCheckAccountTransaction](docs/ModelCheckAccountTransaction.md)
 - [ModelCheckAccountTransactionCheckAccount](docs/ModelCheckAccountTransactionCheckAccount.md)
 - [ModelCheckAccountTransactionResponse](docs/ModelCheckAccountTransactionResponse.md)
 - [ModelCheckAccountTransactionResponseCheckAccount](docs/ModelCheckAccountTransactionResponseCheckAccount.md)
 - [ModelCheckAccountTransactionResponseSevClient](docs/ModelCheckAccountTransactionResponseSevClient.md)
 - [ModelCheckAccountTransactionResponseSourceTransaction](docs/ModelCheckAccountTransactionResponseSourceTransaction.md)
 - [ModelCheckAccountTransactionResponseTargetTransaction](docs/ModelCheckAccountTransactionResponseTargetTransaction.md)
 - [ModelCheckAccountTransactionSevClient](docs/ModelCheckAccountTransactionSevClient.md)
 - [ModelCheckAccountTransactionSourceTransaction](docs/ModelCheckAccountTransactionSourceTransaction.md)
 - [ModelCheckAccountTransactionTargetTransaction](docs/ModelCheckAccountTransactionTargetTransaction.md)
 - [ModelCheckAccountTransactionUpdate](docs/ModelCheckAccountTransactionUpdate.md)
 - [ModelCheckAccountTransactionUpdateCheckAccount](docs/ModelCheckAccountTransactionUpdateCheckAccount.md)
 - [ModelCheckAccountUpdate](docs/ModelCheckAccountUpdate.md)
 - [ModelCommunicationWay](docs/ModelCommunicationWay.md)
 - [ModelCommunicationWayContact](docs/ModelCommunicationWayContact.md)
 - [ModelCommunicationWayKey](docs/ModelCommunicationWayKey.md)
 - [ModelCommunicationWayResponse](docs/ModelCommunicationWayResponse.md)
 - [ModelCommunicationWayResponseContact](docs/ModelCommunicationWayResponseContact.md)
 - [ModelCommunicationWayResponseKey](docs/ModelCommunicationWayResponseKey.md)
 - [ModelCommunicationWayResponseSevClient](docs/ModelCommunicationWayResponseSevClient.md)
 - [ModelCommunicationWaySevClient](docs/ModelCommunicationWaySevClient.md)
 - [ModelCommunicationWayUpdate](docs/ModelCommunicationWayUpdate.md)
 - [ModelCommunicationWayUpdateContact](docs/ModelCommunicationWayUpdateContact.md)
 - [ModelCommunicationWayUpdateKey](docs/ModelCommunicationWayUpdateKey.md)
 - [ModelContact](docs/ModelContact.md)
 - [ModelContactAddress](docs/ModelContactAddress.md)
 - [ModelContactAddressCategory](docs/ModelContactAddressCategory.md)
 - [ModelContactAddressContact](docs/ModelContactAddressContact.md)
 - [ModelContactAddressCountry](docs/ModelContactAddressCountry.md)
 - [ModelContactAddressResponse](docs/ModelContactAddressResponse.md)
 - [ModelContactAddressSevClient](docs/ModelContactAddressSevClient.md)
 - [ModelContactAddressUpdate](docs/ModelContactAddressUpdate.md)
 - [ModelContactAddressUpdateContact](docs/ModelContactAddressUpdateContact.md)
 - [ModelContactAddressUpdateCountry](docs/ModelContactAddressUpdateCountry.md)
 - [ModelContactCategory](docs/ModelContactCategory.md)
 - [ModelContactCustomField](docs/ModelContactCustomField.md)
 - [ModelContactCustomFieldContact](docs/ModelContactCustomFieldContact.md)
 - [ModelContactCustomFieldContactCustomFieldSetting](docs/ModelContactCustomFieldContactCustomFieldSetting.md)
 - [ModelContactCustomFieldResponse](docs/ModelContactCustomFieldResponse.md)
 - [ModelContactCustomFieldResponseContact](docs/ModelContactCustomFieldResponseContact.md)
 - [ModelContactCustomFieldResponseContactCustomFieldSetting](docs/ModelContactCustomFieldResponseContactCustomFieldSetting.md)
 - [ModelContactCustomFieldResponseSevClient](docs/ModelContactCustomFieldResponseSevClient.md)
 - [ModelContactCustomFieldSetting](docs/ModelContactCustomFieldSetting.md)
 - [ModelContactCustomFieldSettingResponse](docs/ModelContactCustomFieldSettingResponse.md)
 - [ModelContactCustomFieldSettingResponseSevClient](docs/ModelContactCustomFieldSettingResponseSevClient.md)
 - [ModelContactCustomFieldSettingUpdate](docs/ModelContactCustomFieldSettingUpdate.md)
 - [ModelContactCustomFieldUpdate](docs/ModelContactCustomFieldUpdate.md)
 - [ModelContactParent](docs/ModelContactParent.md)
 - [ModelContactResponse](docs/ModelContactResponse.md)
 - [ModelContactResponseCategory](docs/ModelContactResponseCategory.md)
 - [ModelContactResponseParent](docs/ModelContactResponseParent.md)
 - [ModelContactResponseSevClient](docs/ModelContactResponseSevClient.md)
 - [ModelContactResponseTaxSet](docs/ModelContactResponseTaxSet.md)
 - [ModelContactTaxSet](docs/ModelContactTaxSet.md)
 - [ModelContactUpdate](docs/ModelContactUpdate.md)
 - [ModelContactUpdateCategory](docs/ModelContactUpdateCategory.md)
 - [ModelContactUpdateParent](docs/ModelContactUpdateParent.md)
 - [ModelCreateInvoiceFromOrder](docs/ModelCreateInvoiceFromOrder.md)
 - [ModelCreateInvoiceFromOrderOrder](docs/ModelCreateInvoiceFromOrderOrder.md)
 - [ModelCreatePackingListFromOrder](docs/ModelCreatePackingListFromOrder.md)
 - [ModelCreditNote](docs/ModelCreditNote.md)
 - [ModelCreditNoteAddressCountry](docs/ModelCreditNoteAddressCountry.md)
 - [ModelCreditNoteContact](docs/ModelCreditNoteContact.md)
 - [ModelCreditNoteContactPerson](docs/ModelCreditNoteContactPerson.md)
 - [ModelCreditNoteCreateUser](docs/ModelCreditNoteCreateUser.md)
 - [ModelCreditNotePos](docs/ModelCreditNotePos.md)
 - [ModelCreditNotePosCreditNote](docs/ModelCreditNotePosCreditNote.md)
 - [ModelCreditNotePosResponse](docs/ModelCreditNotePosResponse.md)
 - [ModelCreditNotePosResponseCreditNote](docs/ModelCreditNotePosResponseCreditNote.md)
 - [ModelCreditNotePosResponsePart](docs/ModelCreditNotePosResponsePart.md)
 - [ModelCreditNotePosResponseSevClient](docs/ModelCreditNotePosResponseSevClient.md)
 - [ModelCreditNotePosResponseUnity](docs/ModelCreditNotePosResponseUnity.md)
 - [ModelCreditNotePosSevClient](docs/ModelCreditNotePosSevClient.md)
 - [ModelCreditNotePosUnity](docs/ModelCreditNotePosUnity.md)
 - [ModelCreditNoteResponse](docs/ModelCreditNoteResponse.md)
 - [ModelCreditNoteResponseAddressCountry](docs/ModelCreditNoteResponseAddressCountry.md)
 - [ModelCreditNoteResponseContact](docs/ModelCreditNoteResponseContact.md)
 - [ModelCreditNoteResponseContactPerson](docs/ModelCreditNoteResponseContactPerson.md)
 - [ModelCreditNoteResponseSevClient](docs/ModelCreditNoteResponseSevClient.md)
 - [ModelCreditNoteResponseTaxSet](docs/ModelCreditNoteResponseTaxSet.md)
 - [ModelCreditNoteSevClient](docs/ModelCreditNoteSevClient.md)
 - [ModelCreditNoteTaxSet](docs/ModelCreditNoteTaxSet.md)
 - [ModelCreditNoteUpdate](docs/ModelCreditNoteUpdate.md)
 - [ModelCreditNoteUpdateContact](docs/ModelCreditNoteUpdateContact.md)
 - [ModelCreditNoteUpdateContactPerson](docs/ModelCreditNoteUpdateContactPerson.md)
 - [ModelDiscount](docs/ModelDiscount.md)
 - [ModelDiscountObject](docs/ModelDiscountObject.md)
 - [ModelEmail](docs/ModelEmail.md)
 - [ModelEmailOrder](docs/ModelEmailOrder.md)
 - [ModelEmailOrderObject](docs/ModelEmailOrderObject.md)
 - [ModelEmailOrderSevClient](docs/ModelEmailOrderSevClient.md)
 - [ModelInvoice](docs/ModelInvoice.md)
 - [ModelInvoiceAddressCountry](docs/ModelInvoiceAddressCountry.md)
 - [ModelInvoiceContact](docs/ModelInvoiceContact.md)
 - [ModelInvoiceOrigin](docs/ModelInvoiceOrigin.md)
 - [ModelInvoicePaymentMethod](docs/ModelInvoicePaymentMethod.md)
 - [ModelInvoicePos](docs/ModelInvoicePos.md)
 - [ModelInvoicePosInvoice](docs/ModelInvoicePosInvoice.md)
 - [ModelInvoicePosPart](docs/ModelInvoicePosPart.md)
 - [ModelInvoicePosResponse](docs/ModelInvoicePosResponse.md)
 - [ModelInvoicePosResponseInvoice](docs/ModelInvoicePosResponseInvoice.md)
 - [ModelInvoicePosResponsePart](docs/ModelInvoicePosResponsePart.md)
 - [ModelInvoicePosResponseSevClient](docs/ModelInvoicePosResponseSevClient.md)
 - [ModelInvoicePosResponseUnity](docs/ModelInvoicePosResponseUnity.md)
 - [ModelInvoicePosSevClient](docs/ModelInvoicePosSevClient.md)
 - [ModelInvoicePosUnity](docs/ModelInvoicePosUnity.md)
 - [ModelInvoicePosUpdate](docs/ModelInvoicePosUpdate.md)
 - [ModelInvoicePosUpdateInvoice](docs/ModelInvoicePosUpdateInvoice.md)
 - [ModelInvoiceResponse](docs/ModelInvoiceResponse.md)
 - [ModelInvoiceResponseAddressCountry](docs/ModelInvoiceResponseAddressCountry.md)
 - [ModelInvoiceResponseContact](docs/ModelInvoiceResponseContact.md)
 - [ModelInvoiceResponseContactPerson](docs/ModelInvoiceResponseContactPerson.md)
 - [ModelInvoiceResponseCostCentre](docs/ModelInvoiceResponseCostCentre.md)
 - [ModelInvoiceResponseCreateUser](docs/ModelInvoiceResponseCreateUser.md)
 - [ModelInvoiceResponseOrigin](docs/ModelInvoiceResponseOrigin.md)
 - [ModelInvoiceResponsePaymentMethod](docs/ModelInvoiceResponsePaymentMethod.md)
 - [ModelInvoiceResponseTaxSet](docs/ModelInvoiceResponseTaxSet.md)
 - [ModelInvoiceTaxSet](docs/ModelInvoiceTaxSet.md)
 - [ModelInvoiceUpdate](docs/ModelInvoiceUpdate.md)
 - [ModelInvoiceUpdateContact](docs/ModelInvoiceUpdateContact.md)
 - [ModelInvoiceUpdateContactPerson](docs/ModelInvoiceUpdateContactPerson.md)
 - [ModelInvoiceUpdateCostCentre](docs/ModelInvoiceUpdateCostCentre.md)
 - [ModelInvoiceUpdateOrigin](docs/ModelInvoiceUpdateOrigin.md)
 - [ModelInvoiceUpdatePaymentMethod](docs/ModelInvoiceUpdatePaymentMethod.md)
 - [ModelInvoiceUpdateSevClient](docs/ModelInvoiceUpdateSevClient.md)
 - [ModelInvoiceUpdateTaxSet](docs/ModelInvoiceUpdateTaxSet.md)
 - [ModelOrder](docs/ModelOrder.md)
 - [ModelOrderAddressCountry](docs/ModelOrderAddressCountry.md)
 - [ModelOrderContact](docs/ModelOrderContact.md)
 - [ModelOrderContactPerson](docs/ModelOrderContactPerson.md)
 - [ModelOrderOrigin](docs/ModelOrderOrigin.md)
 - [ModelOrderPos](docs/ModelOrderPos.md)
 - [ModelOrderPosOrder](docs/ModelOrderPosOrder.md)
 - [ModelOrderPosResponse](docs/ModelOrderPosResponse.md)
 - [ModelOrderPosResponseOrder](docs/ModelOrderPosResponseOrder.md)
 - [ModelOrderPosResponseSevClient](docs/ModelOrderPosResponseSevClient.md)
 - [ModelOrderPosSevClient](docs/ModelOrderPosSevClient.md)
 - [ModelOrderPosUpdate](docs/ModelOrderPosUpdate.md)
 - [ModelOrderResponse](docs/ModelOrderResponse.md)
 - [ModelOrderResponseAddressCountry](docs/ModelOrderResponseAddressCountry.md)
 - [ModelOrderResponseContact](docs/ModelOrderResponseContact.md)
 - [ModelOrderResponseContactPerson](docs/ModelOrderResponseContactPerson.md)
 - [ModelOrderResponseOrigin](docs/ModelOrderResponseOrigin.md)
 - [ModelOrderResponseSevClient](docs/ModelOrderResponseSevClient.md)
 - [ModelOrderResponseTaxSet](docs/ModelOrderResponseTaxSet.md)
 - [ModelOrderTaxSet](docs/ModelOrderTaxSet.md)
 - [ModelOrderUpdate](docs/ModelOrderUpdate.md)
 - [ModelOrderUpdateAddressCountry](docs/ModelOrderUpdateAddressCountry.md)
 - [ModelOrderUpdateContact](docs/ModelOrderUpdateContact.md)
 - [ModelOrderUpdateContactPerson](docs/ModelOrderUpdateContactPerson.md)
 - [ModelOrderUpdateCreateUser](docs/ModelOrderUpdateCreateUser.md)
 - [ModelOrderUpdateSevClient](docs/ModelOrderUpdateSevClient.md)
 - [ModelOrderUpdateTaxSet](docs/ModelOrderUpdateTaxSet.md)
 - [ModelPart](docs/ModelPart.md)
 - [ModelPartCategory](docs/ModelPartCategory.md)
 - [ModelPartSevClient](docs/ModelPartSevClient.md)
 - [ModelPartUnity](docs/ModelPartUnity.md)
 - [ModelPartUpdate](docs/ModelPartUpdate.md)
 - [ModelTagCreateResponse](docs/ModelTagCreateResponse.md)
 - [ModelTagCreateResponseObject](docs/ModelTagCreateResponseObject.md)
 - [ModelTagCreateResponseSevClient](docs/ModelTagCreateResponseSevClient.md)
 - [ModelTagCreateResponseTag](docs/ModelTagCreateResponseTag.md)
 - [ModelTagResponse](docs/ModelTagResponse.md)
 - [ModelVoucher](docs/ModelVoucher.md)
 - [ModelVoucherCreateUser](docs/ModelVoucherCreateUser.md)
 - [ModelVoucherPos](docs/ModelVoucherPos.md)
 - [ModelVoucherPosAccountingType](docs/ModelVoucherPosAccountingType.md)
 - [ModelVoucherPosEstimatedAccountingType](docs/ModelVoucherPosEstimatedAccountingType.md)
 - [ModelVoucherPosResponse](docs/ModelVoucherPosResponse.md)
 - [ModelVoucherPosResponseAccountingType](docs/ModelVoucherPosResponseAccountingType.md)
 - [ModelVoucherPosResponseEstimatedAccountingType](docs/ModelVoucherPosResponseEstimatedAccountingType.md)
 - [ModelVoucherPosResponseSevClient](docs/ModelVoucherPosResponseSevClient.md)
 - [ModelVoucherPosResponseVoucher](docs/ModelVoucherPosResponseVoucher.md)
 - [ModelVoucherPosSevClient](docs/ModelVoucherPosSevClient.md)
 - [ModelVoucherPosVoucher](docs/ModelVoucherPosVoucher.md)
 - [ModelVoucherResponse](docs/ModelVoucherResponse.md)
 - [ModelVoucherResponseCostCentre](docs/ModelVoucherResponseCostCentre.md)
 - [ModelVoucherResponseCreateUser](docs/ModelVoucherResponseCreateUser.md)
 - [ModelVoucherResponseDocument](docs/ModelVoucherResponseDocument.md)
 - [ModelVoucherResponseSevClient](docs/ModelVoucherResponseSevClient.md)
 - [ModelVoucherResponseSupplier](docs/ModelVoucherResponseSupplier.md)
 - [ModelVoucherResponseTaxSet](docs/ModelVoucherResponseTaxSet.md)
 - [ModelVoucherSevClient](docs/ModelVoucherSevClient.md)
 - [ModelVoucherSupplier](docs/ModelVoucherSupplier.md)
 - [ModelVoucherUpdate](docs/ModelVoucherUpdate.md)
 - [ModelVoucherUpdateCostCentre](docs/ModelVoucherUpdateCostCentre.md)
 - [ModelVoucherUpdateDocument](docs/ModelVoucherUpdateDocument.md)
 - [ModelVoucherUpdateSupplier](docs/ModelVoucherUpdateSupplier.md)
 - [ModelVoucherUpdateTaxSet](docs/ModelVoucherUpdateTaxSet.md)
 - [OrderSendBy200Response](docs/OrderSendBy200Response.md)
 - [OrderSendByRequest](docs/OrderSendByRequest.md)
 - [PartGetStock200Response](docs/PartGetStock200Response.md)
 - [ReportContactSevQueryParameter](docs/ReportContactSevQueryParameter.md)
 - [ReportContactSevQueryParameterFilter](docs/ReportContactSevQueryParameterFilter.md)
 - [ReportContactSevQueryParameterFilterCountry](docs/ReportContactSevQueryParameterFilterCountry.md)
 - [ReportInvoiceSevQueryParameter](docs/ReportInvoiceSevQueryParameter.md)
 - [ReportOrderSevQueryParameter](docs/ReportOrderSevQueryParameter.md)
 - [ReportOrderSevQueryParameterFilter](docs/ReportOrderSevQueryParameterFilter.md)
 - [ReportOrderSevQueryParameterFilterContact](docs/ReportOrderSevQueryParameterFilterContact.md)
 - [ReportVoucherSevQueryParameter](docs/ReportVoucherSevQueryParameter.md)
 - [SaveCreditNote](docs/SaveCreditNote.md)
 - [SaveCreditNoteCreditNotePosDelete](docs/SaveCreditNoteCreditNotePosDelete.md)
 - [SaveCreditNoteDiscountDelete](docs/SaveCreditNoteDiscountDelete.md)
 - [SaveCreditNoteDiscountSave](docs/SaveCreditNoteDiscountSave.md)
 - [SaveCreditNoteResponse](docs/SaveCreditNoteResponse.md)
 - [SaveInvoice](docs/SaveInvoice.md)
 - [SaveInvoiceDiscountDelete](docs/SaveInvoiceDiscountDelete.md)
 - [SaveInvoiceDiscountSave](docs/SaveInvoiceDiscountSave.md)
 - [SaveInvoiceInvoicePosDelete](docs/SaveInvoiceInvoicePosDelete.md)
 - [SaveInvoiceResponse](docs/SaveInvoiceResponse.md)
 - [SaveOrder](docs/SaveOrder.md)
 - [SaveOrderOrderPosDelete](docs/SaveOrderOrderPosDelete.md)
 - [SaveOrderResponse](docs/SaveOrderResponse.md)
 - [SaveVoucher](docs/SaveVoucher.md)
 - [SaveVoucherResponse](docs/SaveVoucherResponse.md)
 - [SaveVoucherVoucherPosDelete](docs/SaveVoucherVoucherPosDelete.md)
 - [SendCreditNoteByPrinting200Response](docs/SendCreditNoteByPrinting200Response.md)
 - [SendInvoiceViaEMail201Response](docs/SendInvoiceViaEMail201Response.md)
 - [SendInvoiceViaEMailRequest](docs/SendInvoiceViaEMailRequest.md)
 - [SendorderViaEMail201Response](docs/SendorderViaEMail201Response.md)
 - [UpdateOrderTemplate200Response](docs/UpdateOrderTemplate200Response.md)
 - [UpdateTagRequest](docs/UpdateTagRequest.md)
 - [VoucherUploadFile201Response](docs/VoucherUploadFile201Response.md)
 - [VoucherUploadFile201ResponseObjectsInner](docs/VoucherUploadFile201ResponseObjectsInner.md)
 - [VoucherUploadFileRequest](docs/VoucherUploadFileRequest.md)


## Documentation For Authorization



### api_key

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



