/*
 * API for course project
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type GetClientRequest struct {
	Id string `json:"id"`
}

// AssertGetClientRequestRequired checks if the required fields are not zero-ed
func AssertGetClientRequestRequired(obj GetClientRequest) error {
	elements := map[string]interface{}{
		"id": obj.Id,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseGetClientRequestRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of GetClientRequest (e.g. [][]GetClientRequest), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseGetClientRequestRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aGetClientRequest, ok := obj.(GetClientRequest)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertGetClientRequestRequired(aGetClientRequest)
	})
}
