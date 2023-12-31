/*
API for course project

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the FilterRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilterRequest{}

// FilterRequest struct for FilterRequest
type FilterRequest struct {
	Filters []Filter `json:"filters"`
}

// NewFilterRequest instantiates a new FilterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilterRequest(filters []Filter) *FilterRequest {
	this := FilterRequest{}
	this.Filters = filters
	return &this
}

// NewFilterRequestWithDefaults instantiates a new FilterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilterRequestWithDefaults() *FilterRequest {
	this := FilterRequest{}
	return &this
}

// GetFilters returns the Filters field value
func (o *FilterRequest) GetFilters() []Filter {
	if o == nil {
		var ret []Filter
		return ret
	}

	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value
// and a boolean to check if the value has been set.
func (o *FilterRequest) GetFiltersOk() ([]Filter, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filters, true
}

// SetFilters sets field value
func (o *FilterRequest) SetFilters(v []Filter) {
	o.Filters = v
}

func (o FilterRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilterRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["filters"] = o.Filters
	return toSerialize, nil
}

type NullableFilterRequest struct {
	value *FilterRequest
	isSet bool
}

func (v NullableFilterRequest) Get() *FilterRequest {
	return v.value
}

func (v *NullableFilterRequest) Set(val *FilterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFilterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFilterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilterRequest(val *FilterRequest) *NullableFilterRequest {
	return &NullableFilterRequest{value: val, isSet: true}
}

func (v NullableFilterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
