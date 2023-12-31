/*
 * API for course project
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type CreateEventTypeRequest struct {
	EventType EventType `json:"event_type"`
}

// AssertCreateEventTypeRequestRequired checks if the required fields are not zero-ed
func AssertCreateEventTypeRequestRequired(obj CreateEventTypeRequest) error {
	elements := map[string]interface{}{
		"event_type": obj.EventType,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertEventTypeRequired(obj.EventType); err != nil {
		return err
	}
	return nil
}

// AssertRecurseCreateEventTypeRequestRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of CreateEventTypeRequest (e.g. [][]CreateEventTypeRequest), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseCreateEventTypeRequestRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aCreateEventTypeRequest, ok := obj.(CreateEventTypeRequest)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertCreateEventTypeRequestRequired(aCreateEventTypeRequest)
	})
}
