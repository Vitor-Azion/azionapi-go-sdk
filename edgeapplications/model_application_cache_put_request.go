/*
Edge Application

## Welcome to the Azion API!  With the following APIs you can check, remove or update existing settings, besides creating new ones.  * * *  ## Overview  The Azion API is a RESTful API, based on HTTPS requests, that allows you to integrate your systems with our platform, simply, quickly, and securely.  Here you will find instructions on how our API works and what functionality is available.  This documentation is being constantly updated. Make sure you verify if there are any updates or changes before you perform any development in your application, even if you are familiar with our API.  * * *  Both HTTPS requests and responses must be in JavaScript Object Notation (JSON) format. All HTTPS requests and responses must follow these conventions.  **Base URL**  The base URL of the API, which must be used, is:  [**https://api.azionapi.net/_**](https://api.azionapi.net/)  **HTTP Methods**  Each HTTP method defines the type of operation that will be run.  | HTTP Method | CRUD | Whole Collection (e.g. /items) | Specific Item (e.g. /items/:item_id) | | --- | --- | --- | --- | | GET | Read | It retrieves the list of items in a Collection. | It retrieves a specific item in a Collection. | | POST | Create | It creates a new item in the Collection. | \\- | | PUT | Update/Replace | It replaces a whole Collection with a new one. | It replaces an item in a Collection with a new one. | | PATCH | Update/Modify | It partially updates a Collection. | It partially updates an item in a Collection | | DELETE | Delete | It deletes a whole Collection. | It deletes an item in a Collection. |  * * *  **Status Codes**  The HTTP return code defines the status – successful or not – after the requested operation is run.  | Status Code | Meaning | | --- | --- | | 200 OK | General Status for a successful operation. | | 201 CREATED | Successfully created a collection or item, by means of POST or PUT. | | 204 NO CONTENT | Successful operation, but without any content to be return to the Body. Generally used for DELETE or PUT operations. | | 207 MULTI-STATUS | A batch of operations were triggered by a single request, which resulted in different return codes when it was run, for some of the operations. The codes are displayed in the “status” field, in the body of the response, for each sub-batch of operations for whichever are applicable. | | 400 BAD REQUEST | Request error, such as invalid parameters, missing mandatory parameters or others. | | 401 UNAUTHORIZED | Error caused by a missing HTTP Authentication header. | | 403 FORBIDDEN | User does not have the permissions to run the requested operation. | | 404 NOT FOUND | The requested resource does not exist. | | 405 METHOD NOT ALLOWED | The requested method is not applicable with the URL. | | 406 NOT ACCEPTABLE | Accept header missing from the HTTP request or the header contains formatting or the version is not supported by the API. | | 409 CONFLICT | Conflict generated in running the request, such as a request to create an already existing record. | | 429 TOO MANY REQUESTS | The request was temporarily rejected, due to exceeding the limit on operations. Wait for the time defined in the X-Ratelimit-Reset header and try again. | | 500 INTERNAL SERVER ERROR | Unintentional error, due to an unidentified failure in the request process. | | \\--- |  | | **HTTP Headers** |  |  All requests must be generated with the HTTP header specifying the desired code format for responses and the API version used. The current version of the API is 3 and the format is application/json.  ``` Accept: application/json; version=3  ```  * * *  **Rate Limit**  The limit of operations that can be run via the API is defined by 3 HTTP response headers:  *   **X-ratelimit-limit:** number of operations allowed in one time-frame; *   **X-ratelimit-remaining:** number of operations still available in one time-frame; *   **X-ratelimit-reset:** is the date and time, in IOS 8601 format, which represents the point in the future when the time-frame will be closed and when the limits will, therefore, be reset.       Example of HTTP response headers and a request:  ``` Date: Thu, 02 Nov 2017 12:30:28 GMT X-ratelimit-remaining: 199 X-ratelimit-limit: 200 X-ratelimit-reset: 2017-11-02T12:31:28.675446  ```  In the example provided, 200 operations are allowed within a 1-minute time frame. Of those, there are 199 still available, because one has already been run. The total limit will be reset after 1 minute.  When the X-ratelimit-limit is reached, or in other words, when the X-ratelimit-remaining reaches zero, the API will give the error, HTTP 429 TOO MANY REQUESTS. If your application receives this error, you will need to wait until the time defined in X-ratelimit-reset has passed, to make another request.  *   **X-ratelimit-limit by product**       The *X-ratelimit limit* variations by product are the following:  *   **Real-Time Metrics:** 20 requests per minute. *   **Real-Time Purge:** 200 requests per minute; except for the Wildcard, which is 2000 a day.       > The rate-limit values are based on the client_id.  * * *  **Optional Parameters**  In this version, it is possible to include some optional parameters as part of the GET method, which can help and modify the form in which your data will be returned.  You can combine these parameters to get the result you want.  They are:  | Parameter | Description | Accepted values: | | --- | --- | --- | | order_by | Identifies which field the return should be sorted by. The default ordering is ascending. | Depends on the fields available from the endpoint structure | | sort | Defines which ordering to be used: ascending or descending. | asc  <br>  <br>desc | | page | Identifies which page should be returned, if the return is paginated. The default value is 1. | Page number. | | page_size | Identifies how many items should be returned per page. The default value is 10. | Desired number of items. |  * * *  **Request Exemple**  You can use one parameter, or a combination. See the example below, which uses all of them in the same request.  ``` GET /domains?order_by=name&page_size=20&sort=desc&page=3 Accept: application/json; version=3 Authorization: token 2909f3932069047f4736dc87e72baaddd19c9f75  ```  * * *  # Authentication  Authentication and authorization of operations via Azion API is done through Tokens.  The first step is to create the Token through authenticating a user registered in [Real-Time Manager](https://sso.azion.com/login).  * * *  ## Encoding Username and Password in Base64  Only token creation and cancelling operations are performed through Basic Authentication, that is, with a username and password. You can create and cancel the token through the API itself, but for that you need to encode your username and password in base64.  Base64 encoding can be done from the command line on a Unix terminal:  ``` $ echo 'user@domain:password'|base64 dXNlckBkb21haW46cGFzc3dvcmQK  ```  If you do not have a Unix terminal available, you can use the free online service at [https://www.base64encode.org/](https://www.base64encode.org/)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edgeapplications

import (
	"encoding/json"
)

// checks if the ApplicationCachePutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationCachePutRequest{}

// ApplicationCachePutRequest struct for ApplicationCachePutRequest
type ApplicationCachePutRequest struct {
	Name string `json:"name"`
	BrowserCacheSettings *string `json:"browser_cache_settings,omitempty"`
	BrowserCacheSettingsMaximumTtl *int64 `json:"browser_cache_settings_maximum_ttl,omitempty"`
	CdnCacheSettings *string `json:"cdn_cache_settings,omitempty"`
	CdnCacheSettingsMaximumTtl *int64 `json:"cdn_cache_settings_maximum_ttl,omitempty"`
	CacheByQueryString *string `json:"cache_by_query_string,omitempty"`
	QueryStringFields []string `json:"query_string_fields,omitempty"`
	EnableQueryStringSort *bool `json:"enable_query_string_sort,omitempty"`
	CacheByCookies *string `json:"cache_by_cookies,omitempty"`
	CookieNames []string `json:"cookie_names,omitempty"`
	EnableCachingForPost *bool `json:"enable_caching_for_post,omitempty"`
	L2CachingEnabled *bool `json:"l2_caching_enabled,omitempty"`
	IsSliceConfigurationEnabled *bool `json:"is_slice_configuration_enabled,omitempty"`
	IsSliceEdgeCachingEnabled *bool `json:"is_slice_edge_caching_enabled,omitempty"`
	IsSliceL2CachingEnabled *bool `json:"is_slice_l2_caching_enabled,omitempty"`
	SliceConfigurationRange *int64 `json:"slice_configuration_range,omitempty"`
}

// NewApplicationCachePutRequest instantiates a new ApplicationCachePutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationCachePutRequest(name string) *ApplicationCachePutRequest {
	this := ApplicationCachePutRequest{}
	this.Name = name
	return &this
}

// NewApplicationCachePutRequestWithDefaults instantiates a new ApplicationCachePutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationCachePutRequestWithDefaults() *ApplicationCachePutRequest {
	this := ApplicationCachePutRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ApplicationCachePutRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationCachePutRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplicationCachePutRequest) SetName(v string) {
	o.Name = v
}

// GetBrowserCacheSettings returns the BrowserCacheSettings field value if set, zero value otherwise.
func (o *ApplicationCachePutRequest) GetBrowserCacheSettings() string {
	if o == nil || IsNil(o.BrowserCacheSettings) {
		var ret string
		return ret
	}
	return *o.BrowserCacheSettings
}

// GetBrowserCacheSettingsOk returns a tuple with the BrowserCacheSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCachePutRequest) GetBrowserCacheSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.BrowserCacheSettings) {
		return nil, false
	}
	return o.BrowserCacheSettings, true
}

// HasBrowserCacheSettings returns a boolean if a field has been set.
func (o *ApplicationCachePutRequest) HasBrowserCacheSettings() bool {
	if o != nil && !IsNil(o.BrowserCacheSettings) {
		return true
	}

	return false
}

// SetBrowserCacheSettings gets a reference to the given string and assigns it to the BrowserCacheSettings field.
func (o *ApplicationCachePutRequest) SetBrowserCacheSettings(v string) {
	o.BrowserCacheSettings = &v
}

// GetBrowserCacheSettingsMaximumTtl returns the BrowserCacheSettingsMaximumTtl field value if set, zero value otherwise.
func (o *ApplicationCachePutRequest) GetBrowserCacheSettingsMaximumTtl() int64 {
	if o == nil || IsNil(o.BrowserCacheSettingsMaximumTtl) {
		var ret int64
		return ret
	}
	return *o.BrowserCacheSettingsMaximumTtl
}

// GetBrowserCacheSettingsMaximumTtlOk returns a tuple with the BrowserCacheSettingsMaximumTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCachePutRequest) GetBrowserCacheSettingsMaximumTtlOk() (*int64, bool) {
	if o == nil || IsNil(o.BrowserCacheSettingsMaximumTtl) {
		return nil, false
	}
	return o.BrowserCacheSettingsMaximumTtl, true
}

// HasBrowserCacheSettingsMaximumTtl returns a boolean if a field has been set.
func (o *ApplicationCachePutRequest) HasBrowserCacheSettingsMaximumTtl() bool {
	if o != nil && !IsNil(o.BrowserCacheSettingsMaximumTtl) {
		return true
	}

	return false
}

// SetBrowserCacheSettingsMaximumTtl gets a reference to the given int64 and assigns it to the BrowserCacheSettingsMaximumTtl field.
func (o *ApplicationCachePutRequest) SetBrowserCacheSettingsMaximumTtl(v int64) {
	o.BrowserCacheSettingsMaximumTtl = &v
}

// GetCdnCacheSettings returns the CdnCacheSettings field value if set, zero value otherwise.
func (o *ApplicationCachePutRequest) GetCdnCacheSettings() string {
	if o == nil || IsNil(o.CdnCacheSettings) {
		var ret string
		return ret
	}
	return *o.CdnCacheSettings
}

// GetCdnCacheSettingsOk returns a tuple with the CdnCacheSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCachePutRequest) GetCdnCacheSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.CdnCacheSettings) {
		return nil, false
	}
	return o.CdnCacheSettings, true
}

// HasCdnCacheSettings returns a boolean if a field has been set.
func (o *ApplicationCachePutRequest) HasCdnCacheSettings() bool {
	if o != nil && !IsNil(o.CdnCacheSettings) {
		return true
	}

	return false
}

// SetCdnCacheSettings gets a reference to the given string and assigns it to the CdnCacheSettings field.
func (o *ApplicationCachePutRequest) SetCdnCacheSettings(v string) {
	o.CdnCacheSettings = &v
}

// GetCdnCacheSettingsMaximumTtl returns the CdnCacheSettingsMaximumTtl field value if set, zero value otherwise.
func (o *ApplicationCachePutRequest) GetCdnCacheSettingsMaximumTtl() int64 {
	if o == nil || IsNil(o.CdnCacheSettingsMaximumTtl) {
		var ret int64
		return ret
	}
	return *o.CdnCacheSettingsMaximumTtl
}

// GetCdnCacheSettingsMaximumTtlOk returns a tuple with the CdnCacheSettingsMaximumTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCachePutRequest) GetCdnCacheSettingsMaximumTtlOk() (*int64, bool) {
	if o == nil || IsNil(o.CdnCacheSettingsMaximumTtl) {
		return nil, false
	}
	return o.CdnCacheSettingsMaximumTtl, true
}

// HasCdnCacheSettingsMaximumTtl returns a boolean if a field has been set.
func (o *ApplicationCachePutRequest) HasCdnCacheSettingsMaximumTtl() bool {
	if o != nil && !IsNil(o.CdnCacheSettingsMaximumTtl) {
		return true
	}

	return false
}

// SetCdnCacheSettingsMaximumTtl gets a reference to the given int64 and assigns it to the CdnCacheSettingsMaximumTtl field.
func (o *ApplicationCachePutRequest) SetCdnCacheSettingsMaximumTtl(v int64) {
	o.CdnCacheSettingsMaximumTtl = &v
}

// GetCacheByQueryString returns the CacheByQueryString field value if set, zero value otherwise.
func (o *ApplicationCachePutRequest) GetCacheByQueryString() string {
	if o == nil || IsNil(o.CacheByQueryString) {
		var ret string
		return ret
	}
	return *o.CacheByQueryString
}

// GetCacheByQueryStringOk returns a tuple with the CacheByQueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCachePutRequest) GetCacheByQueryStringOk() (*string, bool) {
	if o == nil || IsNil(o.CacheByQueryString) {
		return nil, false
	}
	return o.CacheByQueryString, true
}

// HasCacheByQueryString returns a boolean if a field has been set.
func (o *ApplicationCachePutRequest) HasCacheByQueryString() bool {
	if o != nil && !IsNil(o.CacheByQueryString) {
		return true
	}

	return false
}

// SetCacheByQueryString gets a reference to the given string and assigns it to the CacheByQueryString field.
func (o *ApplicationCachePutRequest) SetCacheByQueryString(v string) {
	o.CacheByQueryString = &v
}

// GetQueryStringFields returns the QueryStringFields field value if set, zero value otherwise.
func (o *ApplicationCachePutRequest) GetQueryStringFields() []string {
	if o == nil || IsNil(o.QueryStringFields) {
		var ret []string
		return ret
	}
	return o.QueryStringFields
}

// GetQueryStringFieldsOk returns a tuple with the QueryStringFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCachePutRequest) GetQueryStringFieldsOk() ([]string, bool) {
	if o == nil || IsNil(o.QueryStringFields) {
		return nil, false
	}
	return o.QueryStringFields, true
}

// HasQueryStringFields returns a boolean if a field has been set.
func (o *ApplicationCachePutRequest) HasQueryStringFields() bool {
	if o != nil && !IsNil(o.QueryStringFields) {
		return true
	}

	return false
}

// SetQueryStringFields gets a reference to the given []string and assigns it to the QueryStringFields field.
func (o *ApplicationCachePutRequest) SetQueryStringFields(v []string) {
	o.QueryStringFields = v
}

// GetEnableQueryStringSort returns the EnableQueryStringSort field value if set, zero value otherwise.
func (o *ApplicationCachePutRequest) GetEnableQueryStringSort() bool {
	if o == nil || IsNil(o.EnableQueryStringSort) {
		var ret bool
		return ret
	}
	return *o.EnableQueryStringSort
}

// GetEnableQueryStringSortOk returns a tuple with the EnableQueryStringSort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCachePutRequest) GetEnableQueryStringSortOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableQueryStringSort) {
		return nil, false
	}
	return o.EnableQueryStringSort, true
}

// HasEnableQueryStringSort returns a boolean if a field has been set.
func (o *ApplicationCachePutRequest) HasEnableQueryStringSort() bool {
	if o != nil && !IsNil(o.EnableQueryStringSort) {
		return true
	}

	return false
}

// SetEnableQueryStringSort gets a reference to the given bool and assigns it to the EnableQueryStringSort field.
func (o *ApplicationCachePutRequest) SetEnableQueryStringSort(v bool) {
	o.EnableQueryStringSort = &v
}

// GetCacheByCookies returns the CacheByCookies field value if set, zero value otherwise.
func (o *ApplicationCachePutRequest) GetCacheByCookies() string {
	if o == nil || IsNil(o.CacheByCookies) {
		var ret string
		return ret
	}
	return *o.CacheByCookies
}

// GetCacheByCookiesOk returns a tuple with the CacheByCookies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCachePutRequest) GetCacheByCookiesOk() (*string, bool) {
	if o == nil || IsNil(o.CacheByCookies) {
		return nil, false
	}
	return o.CacheByCookies, true
}

// HasCacheByCookies returns a boolean if a field has been set.
func (o *ApplicationCachePutRequest) HasCacheByCookies() bool {
	if o != nil && !IsNil(o.CacheByCookies) {
		return true
	}

	return false
}

// SetCacheByCookies gets a reference to the given string and assigns it to the CacheByCookies field.
func (o *ApplicationCachePutRequest) SetCacheByCookies(v string) {
	o.CacheByCookies = &v
}

// GetCookieNames returns the CookieNames field value if set, zero value otherwise.
func (o *ApplicationCachePutRequest) GetCookieNames() []string {
	if o == nil || IsNil(o.CookieNames) {
		var ret []string
		return ret
	}
	return o.CookieNames
}

// GetCookieNamesOk returns a tuple with the CookieNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCachePutRequest) GetCookieNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.CookieNames) {
		return nil, false
	}
	return o.CookieNames, true
}

// HasCookieNames returns a boolean if a field has been set.
func (o *ApplicationCachePutRequest) HasCookieNames() bool {
	if o != nil && !IsNil(o.CookieNames) {
		return true
	}

	return false
}

// SetCookieNames gets a reference to the given []string and assigns it to the CookieNames field.
func (o *ApplicationCachePutRequest) SetCookieNames(v []string) {
	o.CookieNames = v
}

// GetEnableCachingForPost returns the EnableCachingForPost field value if set, zero value otherwise.
func (o *ApplicationCachePutRequest) GetEnableCachingForPost() bool {
	if o == nil || IsNil(o.EnableCachingForPost) {
		var ret bool
		return ret
	}
	return *o.EnableCachingForPost
}

// GetEnableCachingForPostOk returns a tuple with the EnableCachingForPost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCachePutRequest) GetEnableCachingForPostOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableCachingForPost) {
		return nil, false
	}
	return o.EnableCachingForPost, true
}

// HasEnableCachingForPost returns a boolean if a field has been set.
func (o *ApplicationCachePutRequest) HasEnableCachingForPost() bool {
	if o != nil && !IsNil(o.EnableCachingForPost) {
		return true
	}

	return false
}

// SetEnableCachingForPost gets a reference to the given bool and assigns it to the EnableCachingForPost field.
func (o *ApplicationCachePutRequest) SetEnableCachingForPost(v bool) {
	o.EnableCachingForPost = &v
}

// GetL2CachingEnabled returns the L2CachingEnabled field value if set, zero value otherwise.
func (o *ApplicationCachePutRequest) GetL2CachingEnabled() bool {
	if o == nil || IsNil(o.L2CachingEnabled) {
		var ret bool
		return ret
	}
	return *o.L2CachingEnabled
}

// GetL2CachingEnabledOk returns a tuple with the L2CachingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCachePutRequest) GetL2CachingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.L2CachingEnabled) {
		return nil, false
	}
	return o.L2CachingEnabled, true
}

// HasL2CachingEnabled returns a boolean if a field has been set.
func (o *ApplicationCachePutRequest) HasL2CachingEnabled() bool {
	if o != nil && !IsNil(o.L2CachingEnabled) {
		return true
	}

	return false
}

// SetL2CachingEnabled gets a reference to the given bool and assigns it to the L2CachingEnabled field.
func (o *ApplicationCachePutRequest) SetL2CachingEnabled(v bool) {
	o.L2CachingEnabled = &v
}

// GetIsSliceConfigurationEnabled returns the IsSliceConfigurationEnabled field value if set, zero value otherwise.
func (o *ApplicationCachePutRequest) GetIsSliceConfigurationEnabled() bool {
	if o == nil || IsNil(o.IsSliceConfigurationEnabled) {
		var ret bool
		return ret
	}
	return *o.IsSliceConfigurationEnabled
}

// GetIsSliceConfigurationEnabledOk returns a tuple with the IsSliceConfigurationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCachePutRequest) GetIsSliceConfigurationEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSliceConfigurationEnabled) {
		return nil, false
	}
	return o.IsSliceConfigurationEnabled, true
}

// HasIsSliceConfigurationEnabled returns a boolean if a field has been set.
func (o *ApplicationCachePutRequest) HasIsSliceConfigurationEnabled() bool {
	if o != nil && !IsNil(o.IsSliceConfigurationEnabled) {
		return true
	}

	return false
}

// SetIsSliceConfigurationEnabled gets a reference to the given bool and assigns it to the IsSliceConfigurationEnabled field.
func (o *ApplicationCachePutRequest) SetIsSliceConfigurationEnabled(v bool) {
	o.IsSliceConfigurationEnabled = &v
}

// GetIsSliceEdgeCachingEnabled returns the IsSliceEdgeCachingEnabled field value if set, zero value otherwise.
func (o *ApplicationCachePutRequest) GetIsSliceEdgeCachingEnabled() bool {
	if o == nil || IsNil(o.IsSliceEdgeCachingEnabled) {
		var ret bool
		return ret
	}
	return *o.IsSliceEdgeCachingEnabled
}

// GetIsSliceEdgeCachingEnabledOk returns a tuple with the IsSliceEdgeCachingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCachePutRequest) GetIsSliceEdgeCachingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSliceEdgeCachingEnabled) {
		return nil, false
	}
	return o.IsSliceEdgeCachingEnabled, true
}

// HasIsSliceEdgeCachingEnabled returns a boolean if a field has been set.
func (o *ApplicationCachePutRequest) HasIsSliceEdgeCachingEnabled() bool {
	if o != nil && !IsNil(o.IsSliceEdgeCachingEnabled) {
		return true
	}

	return false
}

// SetIsSliceEdgeCachingEnabled gets a reference to the given bool and assigns it to the IsSliceEdgeCachingEnabled field.
func (o *ApplicationCachePutRequest) SetIsSliceEdgeCachingEnabled(v bool) {
	o.IsSliceEdgeCachingEnabled = &v
}

// GetIsSliceL2CachingEnabled returns the IsSliceL2CachingEnabled field value if set, zero value otherwise.
func (o *ApplicationCachePutRequest) GetIsSliceL2CachingEnabled() bool {
	if o == nil || IsNil(o.IsSliceL2CachingEnabled) {
		var ret bool
		return ret
	}
	return *o.IsSliceL2CachingEnabled
}

// GetIsSliceL2CachingEnabledOk returns a tuple with the IsSliceL2CachingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCachePutRequest) GetIsSliceL2CachingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSliceL2CachingEnabled) {
		return nil, false
	}
	return o.IsSliceL2CachingEnabled, true
}

// HasIsSliceL2CachingEnabled returns a boolean if a field has been set.
func (o *ApplicationCachePutRequest) HasIsSliceL2CachingEnabled() bool {
	if o != nil && !IsNil(o.IsSliceL2CachingEnabled) {
		return true
	}

	return false
}

// SetIsSliceL2CachingEnabled gets a reference to the given bool and assigns it to the IsSliceL2CachingEnabled field.
func (o *ApplicationCachePutRequest) SetIsSliceL2CachingEnabled(v bool) {
	o.IsSliceL2CachingEnabled = &v
}

// GetSliceConfigurationRange returns the SliceConfigurationRange field value if set, zero value otherwise.
func (o *ApplicationCachePutRequest) GetSliceConfigurationRange() int64 {
	if o == nil || IsNil(o.SliceConfigurationRange) {
		var ret int64
		return ret
	}
	return *o.SliceConfigurationRange
}

// GetSliceConfigurationRangeOk returns a tuple with the SliceConfigurationRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCachePutRequest) GetSliceConfigurationRangeOk() (*int64, bool) {
	if o == nil || IsNil(o.SliceConfigurationRange) {
		return nil, false
	}
	return o.SliceConfigurationRange, true
}

// HasSliceConfigurationRange returns a boolean if a field has been set.
func (o *ApplicationCachePutRequest) HasSliceConfigurationRange() bool {
	if o != nil && !IsNil(o.SliceConfigurationRange) {
		return true
	}

	return false
}

// SetSliceConfigurationRange gets a reference to the given int64 and assigns it to the SliceConfigurationRange field.
func (o *ApplicationCachePutRequest) SetSliceConfigurationRange(v int64) {
	o.SliceConfigurationRange = &v
}

func (o ApplicationCachePutRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationCachePutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.BrowserCacheSettings) {
		toSerialize["browser_cache_settings"] = o.BrowserCacheSettings
	}
	if !IsNil(o.BrowserCacheSettingsMaximumTtl) {
		toSerialize["browser_cache_settings_maximum_ttl"] = o.BrowserCacheSettingsMaximumTtl
	}
	if !IsNil(o.CdnCacheSettings) {
		toSerialize["cdn_cache_settings"] = o.CdnCacheSettings
	}
	if !IsNil(o.CdnCacheSettingsMaximumTtl) {
		toSerialize["cdn_cache_settings_maximum_ttl"] = o.CdnCacheSettingsMaximumTtl
	}
	if !IsNil(o.CacheByQueryString) {
		toSerialize["cache_by_query_string"] = o.CacheByQueryString
	}
	if !IsNil(o.QueryStringFields) {
		toSerialize["query_string_fields"] = o.QueryStringFields
	}
	if !IsNil(o.EnableQueryStringSort) {
		toSerialize["enable_query_string_sort"] = o.EnableQueryStringSort
	}
	if !IsNil(o.CacheByCookies) {
		toSerialize["cache_by_cookies"] = o.CacheByCookies
	}
	if !IsNil(o.CookieNames) {
		toSerialize["cookie_names"] = o.CookieNames
	}
	if !IsNil(o.EnableCachingForPost) {
		toSerialize["enable_caching_for_post"] = o.EnableCachingForPost
	}
	if !IsNil(o.L2CachingEnabled) {
		toSerialize["l2_caching_enabled"] = o.L2CachingEnabled
	}
	if !IsNil(o.IsSliceConfigurationEnabled) {
		toSerialize["is_slice_configuration_enabled"] = o.IsSliceConfigurationEnabled
	}
	if !IsNil(o.IsSliceEdgeCachingEnabled) {
		toSerialize["is_slice_edge_caching_enabled"] = o.IsSliceEdgeCachingEnabled
	}
	if !IsNil(o.IsSliceL2CachingEnabled) {
		toSerialize["is_slice_l2_caching_enabled"] = o.IsSliceL2CachingEnabled
	}
	if !IsNil(o.SliceConfigurationRange) {
		toSerialize["slice_configuration_range"] = o.SliceConfigurationRange
	}
	return toSerialize, nil
}

type NullableApplicationCachePutRequest struct {
	value *ApplicationCachePutRequest
	isSet bool
}

func (v NullableApplicationCachePutRequest) Get() *ApplicationCachePutRequest {
	return v.value
}

func (v *NullableApplicationCachePutRequest) Set(val *ApplicationCachePutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationCachePutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationCachePutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationCachePutRequest(val *ApplicationCachePutRequest) *NullableApplicationCachePutRequest {
	return &NullableApplicationCachePutRequest{value: val, isSet: true}
}

func (v NullableApplicationCachePutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationCachePutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


