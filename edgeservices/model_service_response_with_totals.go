/*
Services API

Azion Services

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edgeservices

import (
	"encoding/json"
)

// ServiceResponseWithTotals struct for ServiceResponseWithTotals
type ServiceResponseWithTotals struct {
	Services []ServiceResponse `json:"services"`
	Total int64 `json:"total"`
}

// NewServiceResponseWithTotals instantiates a new ServiceResponseWithTotals object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceResponseWithTotals(services []ServiceResponse, total int64) *ServiceResponseWithTotals {
	this := ServiceResponseWithTotals{}
	this.Services = services
	this.Total = total
	return &this
}

// NewServiceResponseWithTotalsWithDefaults instantiates a new ServiceResponseWithTotals object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceResponseWithTotalsWithDefaults() *ServiceResponseWithTotals {
	this := ServiceResponseWithTotals{}
	return &this
}

// GetServices returns the Services field value
func (o *ServiceResponseWithTotals) GetServices() []ServiceResponse {
	if o == nil {
		var ret []ServiceResponse
		return ret
	}

	return o.Services
}

// GetServicesOk returns a tuple with the Services field value
// and a boolean to check if the value has been set.
func (o *ServiceResponseWithTotals) GetServicesOk() (*[]ServiceResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Services, true
}

// SetServices sets field value
func (o *ServiceResponseWithTotals) SetServices(v []ServiceResponse) {
	o.Services = v
}

// GetTotal returns the Total field value
func (o *ServiceResponseWithTotals) GetTotal() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ServiceResponseWithTotals) GetTotalOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *ServiceResponseWithTotals) SetTotal(v int64) {
	o.Total = v
}

func (o ServiceResponseWithTotals) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["services"] = o.Services
	}
	if true {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableServiceResponseWithTotals struct {
	value *ServiceResponseWithTotals
	isSet bool
}

func (v NullableServiceResponseWithTotals) Get() *ServiceResponseWithTotals {
	return v.value
}

func (v *NullableServiceResponseWithTotals) Set(val *ServiceResponseWithTotals) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceResponseWithTotals) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceResponseWithTotals) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceResponseWithTotals(val *ServiceResponseWithTotals) *NullableServiceResponseWithTotals {
	return &NullableServiceResponseWithTotals{value: val, isSet: true}
}

func (v NullableServiceResponseWithTotals) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceResponseWithTotals) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


