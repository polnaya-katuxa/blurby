/*
 * API for course project
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type FieldFilter struct {
	Field string `json:"field,omitempty"`

	Cmp string `json:"cmp,omitempty"`

	Value1 string `json:"value1,omitempty"`

	Value2 string `json:"value2,omitempty"`
}

// AssertFieldFilterRequired checks if the required fields are not zero-ed
func AssertFieldFilterRequired(obj FieldFilter) error {
	return nil
}

// AssertRecurseFieldFilterRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of FieldFilter (e.g. [][]FieldFilter), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseFieldFilterRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aFieldFilter, ok := obj.(FieldFilter)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertFieldFilterRequired(aFieldFilter)
	})
}
