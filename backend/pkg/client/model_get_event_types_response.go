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

// checks if the GetEventTypesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEventTypesResponse{}

// GetEventTypesResponse struct for GetEventTypesResponse
type GetEventTypesResponse struct {
	EventTypes []EventType `json:"event_types"`
}

// NewGetEventTypesResponse instantiates a new GetEventTypesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEventTypesResponse(eventTypes []EventType) *GetEventTypesResponse {
	this := GetEventTypesResponse{}
	this.EventTypes = eventTypes
	return &this
}

// NewGetEventTypesResponseWithDefaults instantiates a new GetEventTypesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEventTypesResponseWithDefaults() *GetEventTypesResponse {
	this := GetEventTypesResponse{}
	return &this
}

// GetEventTypes returns the EventTypes field value
func (o *GetEventTypesResponse) GetEventTypes() []EventType {
	if o == nil {
		var ret []EventType
		return ret
	}

	return o.EventTypes
}

// GetEventTypesOk returns a tuple with the EventTypes field value
// and a boolean to check if the value has been set.
func (o *GetEventTypesResponse) GetEventTypesOk() ([]EventType, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventTypes, true
}

// SetEventTypes sets field value
func (o *GetEventTypesResponse) SetEventTypes(v []EventType) {
	o.EventTypes = v
}

func (o GetEventTypesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEventTypesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event_types"] = o.EventTypes
	return toSerialize, nil
}

type NullableGetEventTypesResponse struct {
	value *GetEventTypesResponse
	isSet bool
}

func (v NullableGetEventTypesResponse) Get() *GetEventTypesResponse {
	return v.value
}

func (v *NullableGetEventTypesResponse) Set(val *GetEventTypesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEventTypesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEventTypesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEventTypesResponse(val *GetEventTypesResponse) *NullableGetEventTypesResponse {
	return &NullableGetEventTypesResponse{value: val, isSet: true}
}

func (v NullableGetEventTypesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEventTypesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
