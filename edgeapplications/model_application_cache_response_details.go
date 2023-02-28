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

// checks if the ApplicationCacheResponseDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationCacheResponseDetails{}

// ApplicationCacheResponseDetails struct for ApplicationCacheResponseDetails
type ApplicationCacheResponseDetails struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	BrowserCacheSettings string `json:"browser_cache_settings"`
	BrowserCacheSettingsMaximumTtl int64 `json:"browser_cache_settings_maximum_ttl"`
	CdnCacheSettings string `json:"cdn_cache_settings"`
	CdnCacheSettingsMaximumTtl int64 `json:"cdn_cache_settings_maximum_ttl"`
	CacheByQueryString string `json:"cache_by_query_string"`
	QueryStringFields []string `json:"query_string_fields"`
	EnableQueryStringSort bool `json:"enable_query_string_sort"`
	CacheByCookies string `json:"cache_by_cookies"`
	CookieNames []string `json:"cookie_names"`
	AdaptiveDeliveryAction *string `json:"adaptive_delivery_action,omitempty"`
	DeviceGroup []string `json:"device_group,omitempty"`
	EnableCachingForPost bool `json:"enable_caching_for_post"`
	EnableCachingForOptions *bool `json:"enable_caching_for_options,omitempty"`
	L2CachingEnabled bool `json:"l2_caching_enabled"`
}

// NewApplicationCacheResponseDetails instantiates a new ApplicationCacheResponseDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationCacheResponseDetails(id int64, name string, browserCacheSettings string, browserCacheSettingsMaximumTtl int64, cdnCacheSettings string, cdnCacheSettingsMaximumTtl int64, cacheByQueryString string, queryStringFields []string, enableQueryStringSort bool, cacheByCookies string, cookieNames []string, enableCachingForPost bool, l2CachingEnabled bool) *ApplicationCacheResponseDetails {
	this := ApplicationCacheResponseDetails{}
	this.Id = id
	this.Name = name
	this.BrowserCacheSettings = browserCacheSettings
	this.BrowserCacheSettingsMaximumTtl = browserCacheSettingsMaximumTtl
	this.CdnCacheSettings = cdnCacheSettings
	this.CdnCacheSettingsMaximumTtl = cdnCacheSettingsMaximumTtl
	this.CacheByQueryString = cacheByQueryString
	this.QueryStringFields = queryStringFields
	this.EnableQueryStringSort = enableQueryStringSort
	this.CacheByCookies = cacheByCookies
	this.CookieNames = cookieNames
	this.EnableCachingForPost = enableCachingForPost
	this.L2CachingEnabled = l2CachingEnabled
	return &this
}

// NewApplicationCacheResponseDetailsWithDefaults instantiates a new ApplicationCacheResponseDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationCacheResponseDetailsWithDefaults() *ApplicationCacheResponseDetails {
	this := ApplicationCacheResponseDetails{}
	return &this
}

// GetId returns the Id field value
func (o *ApplicationCacheResponseDetails) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApplicationCacheResponseDetails) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApplicationCacheResponseDetails) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ApplicationCacheResponseDetails) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationCacheResponseDetails) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplicationCacheResponseDetails) SetName(v string) {
	o.Name = v
}

// GetBrowserCacheSettings returns the BrowserCacheSettings field value
func (o *ApplicationCacheResponseDetails) GetBrowserCacheSettings() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BrowserCacheSettings
}

// GetBrowserCacheSettingsOk returns a tuple with the BrowserCacheSettings field value
// and a boolean to check if the value has been set.
func (o *ApplicationCacheResponseDetails) GetBrowserCacheSettingsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrowserCacheSettings, true
}

// SetBrowserCacheSettings sets field value
func (o *ApplicationCacheResponseDetails) SetBrowserCacheSettings(v string) {
	o.BrowserCacheSettings = v
}

// GetBrowserCacheSettingsMaximumTtl returns the BrowserCacheSettingsMaximumTtl field value
func (o *ApplicationCacheResponseDetails) GetBrowserCacheSettingsMaximumTtl() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.BrowserCacheSettingsMaximumTtl
}

// GetBrowserCacheSettingsMaximumTtlOk returns a tuple with the BrowserCacheSettingsMaximumTtl field value
// and a boolean to check if the value has been set.
func (o *ApplicationCacheResponseDetails) GetBrowserCacheSettingsMaximumTtlOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrowserCacheSettingsMaximumTtl, true
}

// SetBrowserCacheSettingsMaximumTtl sets field value
func (o *ApplicationCacheResponseDetails) SetBrowserCacheSettingsMaximumTtl(v int64) {
	o.BrowserCacheSettingsMaximumTtl = v
}

// GetCdnCacheSettings returns the CdnCacheSettings field value
func (o *ApplicationCacheResponseDetails) GetCdnCacheSettings() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CdnCacheSettings
}

// GetCdnCacheSettingsOk returns a tuple with the CdnCacheSettings field value
// and a boolean to check if the value has been set.
func (o *ApplicationCacheResponseDetails) GetCdnCacheSettingsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CdnCacheSettings, true
}

// SetCdnCacheSettings sets field value
func (o *ApplicationCacheResponseDetails) SetCdnCacheSettings(v string) {
	o.CdnCacheSettings = v
}

// GetCdnCacheSettingsMaximumTtl returns the CdnCacheSettingsMaximumTtl field value
func (o *ApplicationCacheResponseDetails) GetCdnCacheSettingsMaximumTtl() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CdnCacheSettingsMaximumTtl
}

// GetCdnCacheSettingsMaximumTtlOk returns a tuple with the CdnCacheSettingsMaximumTtl field value
// and a boolean to check if the value has been set.
func (o *ApplicationCacheResponseDetails) GetCdnCacheSettingsMaximumTtlOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CdnCacheSettingsMaximumTtl, true
}

// SetCdnCacheSettingsMaximumTtl sets field value
func (o *ApplicationCacheResponseDetails) SetCdnCacheSettingsMaximumTtl(v int64) {
	o.CdnCacheSettingsMaximumTtl = v
}

// GetCacheByQueryString returns the CacheByQueryString field value
func (o *ApplicationCacheResponseDetails) GetCacheByQueryString() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CacheByQueryString
}

// GetCacheByQueryStringOk returns a tuple with the CacheByQueryString field value
// and a boolean to check if the value has been set.
func (o *ApplicationCacheResponseDetails) GetCacheByQueryStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CacheByQueryString, true
}

// SetCacheByQueryString sets field value
func (o *ApplicationCacheResponseDetails) SetCacheByQueryString(v string) {
	o.CacheByQueryString = v
}

// GetQueryStringFields returns the QueryStringFields field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *ApplicationCacheResponseDetails) GetQueryStringFields() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.QueryStringFields
}

// GetQueryStringFieldsOk returns a tuple with the QueryStringFields field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCacheResponseDetails) GetQueryStringFieldsOk() ([]string, bool) {
	if o == nil || isNil(o.QueryStringFields) {
		return nil, false
	}
	return o.QueryStringFields, true
}

// SetQueryStringFields sets field value
func (o *ApplicationCacheResponseDetails) SetQueryStringFields(v []string) {
	o.QueryStringFields = v
}

// GetEnableQueryStringSort returns the EnableQueryStringSort field value
func (o *ApplicationCacheResponseDetails) GetEnableQueryStringSort() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EnableQueryStringSort
}

// GetEnableQueryStringSortOk returns a tuple with the EnableQueryStringSort field value
// and a boolean to check if the value has been set.
func (o *ApplicationCacheResponseDetails) GetEnableQueryStringSortOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnableQueryStringSort, true
}

// SetEnableQueryStringSort sets field value
func (o *ApplicationCacheResponseDetails) SetEnableQueryStringSort(v bool) {
	o.EnableQueryStringSort = v
}

// GetCacheByCookies returns the CacheByCookies field value
func (o *ApplicationCacheResponseDetails) GetCacheByCookies() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CacheByCookies
}

// GetCacheByCookiesOk returns a tuple with the CacheByCookies field value
// and a boolean to check if the value has been set.
func (o *ApplicationCacheResponseDetails) GetCacheByCookiesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CacheByCookies, true
}

// SetCacheByCookies sets field value
func (o *ApplicationCacheResponseDetails) SetCacheByCookies(v string) {
	o.CacheByCookies = v
}

// GetCookieNames returns the CookieNames field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *ApplicationCacheResponseDetails) GetCookieNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CookieNames
}

// GetCookieNamesOk returns a tuple with the CookieNames field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCacheResponseDetails) GetCookieNamesOk() ([]string, bool) {
	if o == nil || isNil(o.CookieNames) {
		return nil, false
	}
	return o.CookieNames, true
}

// SetCookieNames sets field value
func (o *ApplicationCacheResponseDetails) SetCookieNames(v []string) {
	o.CookieNames = v
}

// GetAdaptiveDeliveryAction returns the AdaptiveDeliveryAction field value if set, zero value otherwise.
func (o *ApplicationCacheResponseDetails) GetAdaptiveDeliveryAction() string {
	if o == nil || isNil(o.AdaptiveDeliveryAction) {
		var ret string
		return ret
	}
	return *o.AdaptiveDeliveryAction
}

// GetAdaptiveDeliveryActionOk returns a tuple with the AdaptiveDeliveryAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCacheResponseDetails) GetAdaptiveDeliveryActionOk() (*string, bool) {
	if o == nil || isNil(o.AdaptiveDeliveryAction) {
		return nil, false
	}
	return o.AdaptiveDeliveryAction, true
}

// HasAdaptiveDeliveryAction returns a boolean if a field has been set.
func (o *ApplicationCacheResponseDetails) HasAdaptiveDeliveryAction() bool {
	if o != nil && !isNil(o.AdaptiveDeliveryAction) {
		return true
	}

	return false
}

// SetAdaptiveDeliveryAction gets a reference to the given string and assigns it to the AdaptiveDeliveryAction field.
func (o *ApplicationCacheResponseDetails) SetAdaptiveDeliveryAction(v string) {
	o.AdaptiveDeliveryAction = &v
}

// GetDeviceGroup returns the DeviceGroup field value if set, zero value otherwise.
func (o *ApplicationCacheResponseDetails) GetDeviceGroup() []string {
	if o == nil || isNil(o.DeviceGroup) {
		var ret []string
		return ret
	}
	return o.DeviceGroup
}

// GetDeviceGroupOk returns a tuple with the DeviceGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCacheResponseDetails) GetDeviceGroupOk() ([]string, bool) {
	if o == nil || isNil(o.DeviceGroup) {
		return nil, false
	}
	return o.DeviceGroup, true
}

// HasDeviceGroup returns a boolean if a field has been set.
func (o *ApplicationCacheResponseDetails) HasDeviceGroup() bool {
	if o != nil && !isNil(o.DeviceGroup) {
		return true
	}

	return false
}

// SetDeviceGroup gets a reference to the given []string and assigns it to the DeviceGroup field.
func (o *ApplicationCacheResponseDetails) SetDeviceGroup(v []string) {
	o.DeviceGroup = v
}

// GetEnableCachingForPost returns the EnableCachingForPost field value
func (o *ApplicationCacheResponseDetails) GetEnableCachingForPost() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EnableCachingForPost
}

// GetEnableCachingForPostOk returns a tuple with the EnableCachingForPost field value
// and a boolean to check if the value has been set.
func (o *ApplicationCacheResponseDetails) GetEnableCachingForPostOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnableCachingForPost, true
}

// SetEnableCachingForPost sets field value
func (o *ApplicationCacheResponseDetails) SetEnableCachingForPost(v bool) {
	o.EnableCachingForPost = v
}

// GetEnableCachingForOptions returns the EnableCachingForOptions field value if set, zero value otherwise.
func (o *ApplicationCacheResponseDetails) GetEnableCachingForOptions() bool {
	if o == nil || isNil(o.EnableCachingForOptions) {
		var ret bool
		return ret
	}
	return *o.EnableCachingForOptions
}

// GetEnableCachingForOptionsOk returns a tuple with the EnableCachingForOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCacheResponseDetails) GetEnableCachingForOptionsOk() (*bool, bool) {
	if o == nil || isNil(o.EnableCachingForOptions) {
		return nil, false
	}
	return o.EnableCachingForOptions, true
}

// HasEnableCachingForOptions returns a boolean if a field has been set.
func (o *ApplicationCacheResponseDetails) HasEnableCachingForOptions() bool {
	if o != nil && !isNil(o.EnableCachingForOptions) {
		return true
	}

	return false
}

// SetEnableCachingForOptions gets a reference to the given bool and assigns it to the EnableCachingForOptions field.
func (o *ApplicationCacheResponseDetails) SetEnableCachingForOptions(v bool) {
	o.EnableCachingForOptions = &v
}

// GetL2CachingEnabled returns the L2CachingEnabled field value
func (o *ApplicationCacheResponseDetails) GetL2CachingEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.L2CachingEnabled
}

// GetL2CachingEnabledOk returns a tuple with the L2CachingEnabled field value
// and a boolean to check if the value has been set.
func (o *ApplicationCacheResponseDetails) GetL2CachingEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.L2CachingEnabled, true
}

// SetL2CachingEnabled sets field value
func (o *ApplicationCacheResponseDetails) SetL2CachingEnabled(v bool) {
	o.L2CachingEnabled = v
}

func (o ApplicationCacheResponseDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationCacheResponseDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["browser_cache_settings"] = o.BrowserCacheSettings
	toSerialize["browser_cache_settings_maximum_ttl"] = o.BrowserCacheSettingsMaximumTtl
	toSerialize["cdn_cache_settings"] = o.CdnCacheSettings
	toSerialize["cdn_cache_settings_maximum_ttl"] = o.CdnCacheSettingsMaximumTtl
	toSerialize["cache_by_query_string"] = o.CacheByQueryString
	if o.QueryStringFields != nil {
		toSerialize["query_string_fields"] = o.QueryStringFields
	}
	toSerialize["enable_query_string_sort"] = o.EnableQueryStringSort
	toSerialize["cache_by_cookies"] = o.CacheByCookies
	if o.CookieNames != nil {
		toSerialize["cookie_names"] = o.CookieNames
	}
	if !isNil(o.AdaptiveDeliveryAction) {
		toSerialize["adaptive_delivery_action"] = o.AdaptiveDeliveryAction
	}
	if !isNil(o.DeviceGroup) {
		toSerialize["device_group"] = o.DeviceGroup
	}
	toSerialize["enable_caching_for_post"] = o.EnableCachingForPost
	if !isNil(o.EnableCachingForOptions) {
		toSerialize["enable_caching_for_options"] = o.EnableCachingForOptions
	}
	toSerialize["l2_caching_enabled"] = o.L2CachingEnabled
	return toSerialize, nil
}

type NullableApplicationCacheResponseDetails struct {
	value *ApplicationCacheResponseDetails
	isSet bool
}

func (v NullableApplicationCacheResponseDetails) Get() *ApplicationCacheResponseDetails {
	return v.value
}

func (v *NullableApplicationCacheResponseDetails) Set(val *ApplicationCacheResponseDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationCacheResponseDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationCacheResponseDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationCacheResponseDetails(val *ApplicationCacheResponseDetails) *NullableApplicationCacheResponseDetails {
	return &NullableApplicationCacheResponseDetails{value: val, isSet: true}
}

func (v NullableApplicationCacheResponseDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationCacheResponseDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


