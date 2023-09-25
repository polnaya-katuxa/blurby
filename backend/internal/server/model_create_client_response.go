/*
 * API for course project
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type CreateClientResponse struct {
	Created bool `json:"created"`
}

// AssertCreateClientResponseRequired checks if the required fields are not zero-ed
func AssertCreateClientResponseRequired(obj CreateClientResponse) error {
	elements := map[string]interface{}{
		"created": obj.Created,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseCreateClientResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of CreateClientResponse (e.g. [][]CreateClientResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseCreateClientResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aCreateClientResponse, ok := obj.(CreateClientResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertCreateClientResponseRequired(aCreateClientResponse)
	})
}