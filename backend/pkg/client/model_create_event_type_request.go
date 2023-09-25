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

// checks if the CreateEventTypeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateEventTypeRequest{}

// CreateEventTypeRequest struct for CreateEventTypeRequest
type CreateEventTypeRequest struct {
	EventType EventType `json:"event_type"`
}

// NewCreateEventTypeRequest instantiates a new CreateEventTypeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateEventTypeRequest(eventType EventType) *CreateEventTypeRequest {
	this := CreateEventTypeRequest{}
	this.EventType = eventType
	return &this
}

// NewCreateEventTypeRequestWithDefaults instantiates a new CreateEventTypeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateEventTypeRequestWithDefaults() *CreateEventTypeRequest {
	this := CreateEventTypeRequest{}
	return &this
}

// GetEventType returns the EventType field value
func (o *CreateEventTypeRequest) GetEventType() EventType {
	if o == nil {
		var ret EventType
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *CreateEventTypeRequest) GetEventTypeOk() (*EventType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *CreateEventTypeRequest) SetEventType(v EventType) {
	o.EventType = v
}

func (o CreateEventTypeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateEventTypeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event_type"] = o.EventType
	return toSerialize, nil
}

type NullableCreateEventTypeRequest struct {
	value *CreateEventTypeRequest
	isSet bool
}

func (v NullableCreateEventTypeRequest) Get() *CreateEventTypeRequest {
	return v.value
}

func (v *NullableCreateEventTypeRequest) Set(val *CreateEventTypeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEventTypeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEventTypeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEventTypeRequest(val *CreateEventTypeRequest) *NullableCreateEventTypeRequest {
	return &NullableCreateEventTypeRequest{value: val, isSet: true}
}

func (v NullableCreateEventTypeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEventTypeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
