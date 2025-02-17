/*
Intelligent DNS

Azion Intelligent DNS API

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package idns

import (
	"encoding/json"
)

// checks if the GetRecordsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRecordsResponse{}

// GetRecordsResponse Object returned by get zone record
type GetRecordsResponse struct {
	// The schema version
	SchemaVersion *int32 `json:"schema_version,omitempty"`
	// Number of records
	Count *int32 `json:"count,omitempty"`
	// The total pages
	TotalPages *int32 `json:"total_pages,omitempty"`
	Links *GetZonesResponseLinks `json:"links,omitempty"`
	Results *GetRecordsResponseResults `json:"results,omitempty"`
}

// NewGetRecordsResponse instantiates a new GetRecordsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRecordsResponse() *GetRecordsResponse {
	this := GetRecordsResponse{}
	return &this
}

// NewGetRecordsResponseWithDefaults instantiates a new GetRecordsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRecordsResponseWithDefaults() *GetRecordsResponse {
	this := GetRecordsResponse{}
	return &this
}

// GetSchemaVersion returns the SchemaVersion field value if set, zero value otherwise.
func (o *GetRecordsResponse) GetSchemaVersion() int32 {
	if o == nil || IsNil(o.SchemaVersion) {
		var ret int32
		return ret
	}
	return *o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecordsResponse) GetSchemaVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.SchemaVersion) {
		return nil, false
	}
	return o.SchemaVersion, true
}

// HasSchemaVersion returns a boolean if a field has been set.
func (o *GetRecordsResponse) HasSchemaVersion() bool {
	if o != nil && !IsNil(o.SchemaVersion) {
		return true
	}

	return false
}

// SetSchemaVersion gets a reference to the given int32 and assigns it to the SchemaVersion field.
func (o *GetRecordsResponse) SetSchemaVersion(v int32) {
	o.SchemaVersion = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *GetRecordsResponse) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecordsResponse) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *GetRecordsResponse) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *GetRecordsResponse) SetCount(v int32) {
	o.Count = &v
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *GetRecordsResponse) GetTotalPages() int32 {
	if o == nil || IsNil(o.TotalPages) {
		var ret int32
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecordsResponse) GetTotalPagesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalPages) {
		return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *GetRecordsResponse) HasTotalPages() bool {
	if o != nil && !IsNil(o.TotalPages) {
		return true
	}

	return false
}

// SetTotalPages gets a reference to the given int32 and assigns it to the TotalPages field.
func (o *GetRecordsResponse) SetTotalPages(v int32) {
	o.TotalPages = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GetRecordsResponse) GetLinks() GetZonesResponseLinks {
	if o == nil || IsNil(o.Links) {
		var ret GetZonesResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecordsResponse) GetLinksOk() (*GetZonesResponseLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GetRecordsResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GetZonesResponseLinks and assigns it to the Links field.
func (o *GetRecordsResponse) SetLinks(v GetZonesResponseLinks) {
	o.Links = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *GetRecordsResponse) GetResults() GetRecordsResponseResults {
	if o == nil || IsNil(o.Results) {
		var ret GetRecordsResponseResults
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecordsResponse) GetResultsOk() (*GetRecordsResponseResults, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *GetRecordsResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given GetRecordsResponseResults and assigns it to the Results field.
func (o *GetRecordsResponse) SetResults(v GetRecordsResponseResults) {
	o.Results = &v
}

func (o GetRecordsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRecordsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SchemaVersion) {
		toSerialize["schema_version"] = o.SchemaVersion
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.TotalPages) {
		toSerialize["total_pages"] = o.TotalPages
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableGetRecordsResponse struct {
	value *GetRecordsResponse
	isSet bool
}

func (v NullableGetRecordsResponse) Get() *GetRecordsResponse {
	return v.value
}

func (v *NullableGetRecordsResponse) Set(val *GetRecordsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRecordsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRecordsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRecordsResponse(val *GetRecordsResponse) *NullableGetRecordsResponse {
	return &NullableGetRecordsResponse{value: val, isSet: true}
}

func (v NullableGetRecordsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRecordsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


