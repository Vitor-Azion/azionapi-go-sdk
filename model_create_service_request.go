/*
Services API

Azion Services

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// CreateServiceRequest struct for CreateServiceRequest
type CreateServiceRequest struct {
	Name string `json:"name"`
}

// NewCreateServiceRequest instantiates a new CreateServiceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateServiceRequest(name string) *CreateServiceRequest {
	this := CreateServiceRequest{}
	this.Name = name
	return &this
}

// NewCreateServiceRequestWithDefaults instantiates a new CreateServiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateServiceRequestWithDefaults() *CreateServiceRequest {
	this := CreateServiceRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateServiceRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateServiceRequest) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateServiceRequest) SetName(v string) {
	o.Name = v
}

func (o CreateServiceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableCreateServiceRequest struct {
	value *CreateServiceRequest
	isSet bool
}

func (v NullableCreateServiceRequest) Get() *CreateServiceRequest {
	return v.value
}

func (v *NullableCreateServiceRequest) Set(val *CreateServiceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateServiceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateServiceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateServiceRequest(val *CreateServiceRequest) *NullableCreateServiceRequest {
	return &NullableCreateServiceRequest{value: val, isSet: true}
}

func (v NullableCreateServiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateServiceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


