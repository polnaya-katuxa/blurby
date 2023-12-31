/*
 * API for course project
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type GetClientsResponse struct {
	Clients []Client `json:"clients"`
}

// AssertGetClientsResponseRequired checks if the required fields are not zero-ed
func AssertGetClientsResponseRequired(obj GetClientsResponse) error {
	elements := map[string]interface{}{
		"clients": obj.Clients,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Clients {
		if err := AssertClientRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseGetClientsResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of GetClientsResponse (e.g. [][]GetClientsResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseGetClientsResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aGetClientsResponse, ok := obj.(GetClientsResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertGetClientsResponseRequired(aGetClientsResponse)
	})
}
