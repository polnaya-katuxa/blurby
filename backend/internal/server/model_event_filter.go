/*
 * API for course project
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type EventFilter struct {
	Alias string `json:"alias,omitempty"`

	Span string `json:"span,omitempty"`

	Rate int32 `json:"rate,omitempty"`
}

// AssertEventFilterRequired checks if the required fields are not zero-ed
func AssertEventFilterRequired(obj EventFilter) error {
	return nil
}

// AssertRecurseEventFilterRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of EventFilter (e.g. [][]EventFilter), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseEventFilterRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aEventFilter, ok := obj.(EventFilter)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertEventFilterRequired(aEventFilter)
	})
}
