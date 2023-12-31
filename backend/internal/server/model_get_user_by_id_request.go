/*
 * API for course project
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type GetUserByIdRequest struct {
	Id string `json:"id"`
}

// AssertGetUserByIdRequestRequired checks if the required fields are not zero-ed
func AssertGetUserByIdRequestRequired(obj GetUserByIdRequest) error {
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

// AssertRecurseGetUserByIdRequestRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of GetUserByIdRequest (e.g. [][]GetUserByIdRequest), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseGetUserByIdRequestRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aGetUserByIdRequest, ok := obj.(GetUserByIdRequest)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertGetUserByIdRequestRequired(aGetUserByIdRequest)
	})
}
