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

// CreateZone Object returned by create zone
type CreateZone struct {
	// The created hosted zone
	Results []Zone `json:"results,omitempty"`
	// The schema version
	SchemaVersion *int32 `json:"schema_version,omitempty"`
}

// NewCreateZone instantiates a new CreateZone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateZone() *CreateZone {
	this := CreateZone{}
	return &this
}

// NewCreateZoneWithDefaults instantiates a new CreateZone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateZoneWithDefaults() *CreateZone {
	this := CreateZone{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *CreateZone) GetResults() []Zone {
	if o == nil || isNil(o.Results) {
		var ret []Zone
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateZone) GetResultsOk() ([]Zone, bool) {
	if o == nil || isNil(o.Results) {
    return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *CreateZone) HasResults() bool {
	if o != nil && !isNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []Zone and assigns it to the Results field.
func (o *CreateZone) SetResults(v []Zone) {
	o.Results = v
}

// GetSchemaVersion returns the SchemaVersion field value if set, zero value otherwise.
func (o *CreateZone) GetSchemaVersion() int32 {
	if o == nil || isNil(o.SchemaVersion) {
		var ret int32
		return ret
	}
	return *o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateZone) GetSchemaVersionOk() (*int32, bool) {
	if o == nil || isNil(o.SchemaVersion) {
    return nil, false
	}
	return o.SchemaVersion, true
}

// HasSchemaVersion returns a boolean if a field has been set.
func (o *CreateZone) HasSchemaVersion() bool {
	if o != nil && !isNil(o.SchemaVersion) {
		return true
	}

	return false
}

// SetSchemaVersion gets a reference to the given int32 and assigns it to the SchemaVersion field.
func (o *CreateZone) SetSchemaVersion(v int32) {
	o.SchemaVersion = &v
}

func (o CreateZone) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !isNil(o.SchemaVersion) {
		toSerialize["schema_version"] = o.SchemaVersion
	}
	return json.Marshal(toSerialize)
}

type NullableCreateZone struct {
	value *CreateZone
	isSet bool
}

func (v NullableCreateZone) Get() *CreateZone {
	return v.value
}

func (v *NullableCreateZone) Set(val *CreateZone) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateZone) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateZone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateZone(val *CreateZone) *NullableCreateZone {
	return &NullableCreateZone{value: val, isSet: true}
}

func (v NullableCreateZone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateZone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


