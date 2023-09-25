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

// checks if the DeleteUserResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteUserResponse{}

// DeleteUserResponse struct for DeleteUserResponse
type DeleteUserResponse struct {
	Deleted bool `json:"deleted"`
}

// NewDeleteUserResponse instantiates a new DeleteUserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteUserResponse(deleted bool) *DeleteUserResponse {
	this := DeleteUserResponse{}
	this.Deleted = deleted
	return &this
}

// NewDeleteUserResponseWithDefaults instantiates a new DeleteUserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteUserResponseWithDefaults() *DeleteUserResponse {
	this := DeleteUserResponse{}
	return &this
}

// GetDeleted returns the Deleted field value
func (o *DeleteUserResponse) GetDeleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value
// and a boolean to check if the value has been set.
func (o *DeleteUserResponse) GetDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deleted, true
}

// SetDeleted sets field value
func (o *DeleteUserResponse) SetDeleted(v bool) {
	o.Deleted = v
}

func (o DeleteUserResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteUserResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deleted"] = o.Deleted
	return toSerialize, nil
}

type NullableDeleteUserResponse struct {
	value *DeleteUserResponse
	isSet bool
}

func (v NullableDeleteUserResponse) Get() *DeleteUserResponse {
	return v.value
}

func (v *NullableDeleteUserResponse) Set(val *DeleteUserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteUserResponse(val *DeleteUserResponse) *NullableDeleteUserResponse {
	return &NullableDeleteUserResponse{value: val, isSet: true}
}

func (v NullableDeleteUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
