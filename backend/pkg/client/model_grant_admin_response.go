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

// checks if the GrantAdminResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GrantAdminResponse{}

// GrantAdminResponse struct for GrantAdminResponse
type GrantAdminResponse struct {
	Granted bool `json:"granted"`
}

// NewGrantAdminResponse instantiates a new GrantAdminResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrantAdminResponse(granted bool) *GrantAdminResponse {
	this := GrantAdminResponse{}
	this.Granted = granted
	return &this
}

// NewGrantAdminResponseWithDefaults instantiates a new GrantAdminResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrantAdminResponseWithDefaults() *GrantAdminResponse {
	this := GrantAdminResponse{}
	return &this
}

// GetGranted returns the Granted field value
func (o *GrantAdminResponse) GetGranted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Granted
}

// GetGrantedOk returns a tuple with the Granted field value
// and a boolean to check if the value has been set.
func (o *GrantAdminResponse) GetGrantedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Granted, true
}

// SetGranted sets field value
func (o *GrantAdminResponse) SetGranted(v bool) {
	o.Granted = v
}

func (o GrantAdminResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GrantAdminResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["granted"] = o.Granted
	return toSerialize, nil
}

type NullableGrantAdminResponse struct {
	value *GrantAdminResponse
	isSet bool
}

func (v NullableGrantAdminResponse) Get() *GrantAdminResponse {
	return v.value
}

func (v *NullableGrantAdminResponse) Set(val *GrantAdminResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantAdminResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantAdminResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantAdminResponse(val *GrantAdminResponse) *NullableGrantAdminResponse {
	return &NullableGrantAdminResponse{value: val, isSet: true}
}

func (v NullableGrantAdminResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantAdminResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
