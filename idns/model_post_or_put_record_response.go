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

// checks if the PostOrPutRecordResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostOrPutRecordResponse{}

// PostOrPutRecordResponse Object returned by create or update zone record
type PostOrPutRecordResponse struct {
	// The schema version
	SchemaVersion *int32 `json:"schema_version,omitempty"`
	Results *RecordPostOrPut `json:"results,omitempty"`
}

// NewPostOrPutRecordResponse instantiates a new PostOrPutRecordResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostOrPutRecordResponse() *PostOrPutRecordResponse {
	this := PostOrPutRecordResponse{}
	return &this
}

// NewPostOrPutRecordResponseWithDefaults instantiates a new PostOrPutRecordResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostOrPutRecordResponseWithDefaults() *PostOrPutRecordResponse {
	this := PostOrPutRecordResponse{}
	return &this
}

// GetSchemaVersion returns the SchemaVersion field value if set, zero value otherwise.
func (o *PostOrPutRecordResponse) GetSchemaVersion() int32 {
	if o == nil || IsNil(o.SchemaVersion) {
		var ret int32
		return ret
	}
	return *o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostOrPutRecordResponse) GetSchemaVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.SchemaVersion) {
		return nil, false
	}
	return o.SchemaVersion, true
}

// HasSchemaVersion returns a boolean if a field has been set.
func (o *PostOrPutRecordResponse) HasSchemaVersion() bool {
	if o != nil && !IsNil(o.SchemaVersion) {
		return true
	}

	return false
}

// SetSchemaVersion gets a reference to the given int32 and assigns it to the SchemaVersion field.
func (o *PostOrPutRecordResponse) SetSchemaVersion(v int32) {
	o.SchemaVersion = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PostOrPutRecordResponse) GetResults() RecordPostOrPut {
	if o == nil || IsNil(o.Results) {
		var ret RecordPostOrPut
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostOrPutRecordResponse) GetResultsOk() (*RecordPostOrPut, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PostOrPutRecordResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given RecordPostOrPut and assigns it to the Results field.
func (o *PostOrPutRecordResponse) SetResults(v RecordPostOrPut) {
	o.Results = &v
}

func (o PostOrPutRecordResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostOrPutRecordResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SchemaVersion) {
		toSerialize["schema_version"] = o.SchemaVersion
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullablePostOrPutRecordResponse struct {
	value *PostOrPutRecordResponse
	isSet bool
}

func (v NullablePostOrPutRecordResponse) Get() *PostOrPutRecordResponse {
	return v.value
}

func (v *NullablePostOrPutRecordResponse) Set(val *PostOrPutRecordResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePostOrPutRecordResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePostOrPutRecordResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostOrPutRecordResponse(val *PostOrPutRecordResponse) *NullablePostOrPutRecordResponse {
	return &NullablePostOrPutRecordResponse{value: val, isSet: true}
}

func (v NullablePostOrPutRecordResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostOrPutRecordResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


