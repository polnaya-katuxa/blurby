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

// checks if the GetUserRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUserRequest{}

// GetUserRequest struct for GetUserRequest
type GetUserRequest struct {
	Login string `json:"login"`
}

// NewGetUserRequest instantiates a new GetUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserRequest(login string) *GetUserRequest {
	this := GetUserRequest{}
	this.Login = login
	return &this
}

// NewGetUserRequestWithDefaults instantiates a new GetUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserRequestWithDefaults() *GetUserRequest {
	this := GetUserRequest{}
	return &this
}

// GetLogin returns the Login field value
func (o *GetUserRequest) GetLogin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Login
}

// GetLoginOk returns a tuple with the Login field value
// and a boolean to check if the value has been set.
func (o *GetUserRequest) GetLoginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Login, true
}

// SetLogin sets field value
func (o *GetUserRequest) SetLogin(v string) {
	o.Login = v
}

func (o GetUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["login"] = o.Login
	return toSerialize, nil
}

type NullableGetUserRequest struct {
	value *GetUserRequest
	isSet bool
}

func (v NullableGetUserRequest) Get() *GetUserRequest {
	return v.value
}

func (v *NullableGetUserRequest) Set(val *GetUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserRequest(val *GetUserRequest) *NullableGetUserRequest {
	return &NullableGetUserRequest{value: val, isSet: true}
}

func (v NullableGetUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
