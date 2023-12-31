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

// checks if the UserInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserInfoResponse{}

// UserInfoResponse struct for UserInfoResponse
type UserInfoResponse struct {
	User User `json:"user"`
}

// NewUserInfoResponse instantiates a new UserInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserInfoResponse(user User) *UserInfoResponse {
	this := UserInfoResponse{}
	this.User = user
	return &this
}

// NewUserInfoResponseWithDefaults instantiates a new UserInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserInfoResponseWithDefaults() *UserInfoResponse {
	this := UserInfoResponse{}
	return &this
}

// GetUser returns the User field value
func (o *UserInfoResponse) GetUser() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *UserInfoResponse) GetUserOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *UserInfoResponse) SetUser(v User) {
	o.User = v
}

func (o UserInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user"] = o.User
	return toSerialize, nil
}

type NullableUserInfoResponse struct {
	value *UserInfoResponse
	isSet bool
}

func (v NullableUserInfoResponse) Get() *UserInfoResponse {
	return v.value
}

func (v *NullableUserInfoResponse) Set(val *UserInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserInfoResponse(val *UserInfoResponse) *NullableUserInfoResponse {
	return &NullableUserInfoResponse{value: val, isSet: true}
}

func (v NullableUserInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
