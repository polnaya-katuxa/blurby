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

// checks if the DeleteUserRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteUserRequest{}

// DeleteUserRequest struct for DeleteUserRequest
type DeleteUserRequest struct {
	Login string `json:"login"`
}

// NewDeleteUserRequest instantiates a new DeleteUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteUserRequest(login string) *DeleteUserRequest {
	this := DeleteUserRequest{}
	this.Login = login
	return &this
}

// NewDeleteUserRequestWithDefaults instantiates a new DeleteUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteUserRequestWithDefaults() *DeleteUserRequest {
	this := DeleteUserRequest{}
	return &this
}

// GetLogin returns the Login field value
func (o *DeleteUserRequest) GetLogin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Login
}

// GetLoginOk returns a tuple with the Login field value
// and a boolean to check if the value has been set.
func (o *DeleteUserRequest) GetLoginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Login, true
}

// SetLogin sets field value
func (o *DeleteUserRequest) SetLogin(v string) {
	o.Login = v
}

func (o DeleteUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["login"] = o.Login
	return toSerialize, nil
}

type NullableDeleteUserRequest struct {
	value *DeleteUserRequest
	isSet bool
}

func (v NullableDeleteUserRequest) Get() *DeleteUserRequest {
	return v.value
}

func (v *NullableDeleteUserRequest) Set(val *DeleteUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteUserRequest(val *DeleteUserRequest) *NullableDeleteUserRequest {
	return &NullableDeleteUserRequest{value: val, isSet: true}
}

func (v NullableDeleteUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
