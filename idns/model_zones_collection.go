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

// ZonesCollection struct for ZonesCollection
type ZonesCollection struct {
	// The schema version
	SchemaVersion *int32 `json:"schema_version,omitempty"`
	// Number of records
	Count *int32 `json:"count,omitempty"`
	// The total pages
	TotalPages *int32 `json:"total_pages,omitempty"`
	Links *ZonesCollectionLinks `json:"links,omitempty"`
	// Hosted zones collection
	Results []Zone `json:"results,omitempty"`
}

// NewZonesCollection instantiates a new ZonesCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZonesCollection() *ZonesCollection {
	this := ZonesCollection{}
	return &this
}

// NewZonesCollectionWithDefaults instantiates a new ZonesCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZonesCollectionWithDefaults() *ZonesCollection {
	this := ZonesCollection{}
	return &this
}

// GetSchemaVersion returns the SchemaVersion field value if set, zero value otherwise.
func (o *ZonesCollection) GetSchemaVersion() int32 {
	if o == nil || isNil(o.SchemaVersion) {
		var ret int32
		return ret
	}
	return *o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZonesCollection) GetSchemaVersionOk() (*int32, bool) {
	if o == nil || isNil(o.SchemaVersion) {
    return nil, false
	}
	return o.SchemaVersion, true
}

// HasSchemaVersion returns a boolean if a field has been set.
func (o *ZonesCollection) HasSchemaVersion() bool {
	if o != nil && !isNil(o.SchemaVersion) {
		return true
	}

	return false
}

// SetSchemaVersion gets a reference to the given int32 and assigns it to the SchemaVersion field.
func (o *ZonesCollection) SetSchemaVersion(v int32) {
	o.SchemaVersion = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ZonesCollection) GetCount() int32 {
	if o == nil || isNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZonesCollection) GetCountOk() (*int32, bool) {
	if o == nil || isNil(o.Count) {
    return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ZonesCollection) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ZonesCollection) SetCount(v int32) {
	o.Count = &v
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *ZonesCollection) GetTotalPages() int32 {
	if o == nil || isNil(o.TotalPages) {
		var ret int32
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZonesCollection) GetTotalPagesOk() (*int32, bool) {
	if o == nil || isNil(o.TotalPages) {
    return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *ZonesCollection) HasTotalPages() bool {
	if o != nil && !isNil(o.TotalPages) {
		return true
	}

	return false
}

// SetTotalPages gets a reference to the given int32 and assigns it to the TotalPages field.
func (o *ZonesCollection) SetTotalPages(v int32) {
	o.TotalPages = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ZonesCollection) GetLinks() ZonesCollectionLinks {
	if o == nil || isNil(o.Links) {
		var ret ZonesCollectionLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZonesCollection) GetLinksOk() (*ZonesCollectionLinks, bool) {
	if o == nil || isNil(o.Links) {
    return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ZonesCollection) HasLinks() bool {
	if o != nil && !isNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ZonesCollectionLinks and assigns it to the Links field.
func (o *ZonesCollection) SetLinks(v ZonesCollectionLinks) {
	o.Links = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ZonesCollection) GetResults() []Zone {
	if o == nil || isNil(o.Results) {
		var ret []Zone
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZonesCollection) GetResultsOk() ([]Zone, bool) {
	if o == nil || isNil(o.Results) {
    return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ZonesCollection) HasResults() bool {
	if o != nil && !isNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []Zone and assigns it to the Results field.
func (o *ZonesCollection) SetResults(v []Zone) {
	o.Results = v
}

func (o ZonesCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SchemaVersion) {
		toSerialize["schema_version"] = o.SchemaVersion
	}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !isNil(o.TotalPages) {
		toSerialize["total_pages"] = o.TotalPages
	}
	if !isNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !isNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableZonesCollection struct {
	value *ZonesCollection
	isSet bool
}

func (v NullableZonesCollection) Get() *ZonesCollection {
	return v.value
}

func (v *NullableZonesCollection) Set(val *ZonesCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableZonesCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableZonesCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZonesCollection(val *ZonesCollection) *NullableZonesCollection {
	return &NullableZonesCollection{value: val, isSet: true}
}

func (v NullableZonesCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZonesCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


