package refutil

import "reflect"

// Equalizer is way how to check if object
// has way how to compare to another object
type Equalizer interface {
	IsEqual(v interface{}) bool
}

// IsDeepEqual determines if two Objects are considered equal.
func IsDeepEqual(expected, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}
	return reflect.DeepEqual(expected, actual)
}

// IsEqual gets whether two Objects are equal, or if their
// values are equal.
func IsEqual(expected, actual interface{}) bool {
	equalizer, ok := expected.(Equalizer)
	if ok {
		return equalizer.IsEqual(actual)
	}
	if IsDeepEqual(expected, actual) {
		return true
	}
	actualType := reflect.TypeOf(actual)
	if actualType == nil {
		return false
	}
	expectedValue := reflect.ValueOf(expected)
	if expectedValue.IsValid() && expectedValue.Type().ConvertibleTo(actualType) {
		// Attempt comparison after type conversion
		return reflect.DeepEqual(expectedValue.Convert(actualType).Interface(), actual)
	}
	return false
}
