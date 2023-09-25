/*
 * API for course project
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type ClientStats struct {
	Num int32 `json:"num,omitempty"`

	AvgAge int32 `json:"avgAge,omitempty"`
}

// AssertClientStatsRequired checks if the required fields are not zero-ed
func AssertClientStatsRequired(obj ClientStats) error {
	return nil
}

// AssertRecurseClientStatsRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ClientStats (e.g. [][]ClientStats), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseClientStatsRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aClientStats, ok := obj.(ClientStats)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertClientStatsRequired(aClientStats)
	})
}
