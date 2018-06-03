package refutil

import (
	"reflect"
)

// IsNil checks if a specified object is nil or not
// also check underlying type if its nil or not
func IsNil(object interface{}) bool {
	if object == nil {
		return true
	}
	value := reflect.ValueOf(object)
	kind := value.Kind()
	if kind >= reflect.Chan && kind <= reflect.Slice && value.IsNil() {
		return true
	}
	if kind == reflect.Ptr {
		return IsNil(value.Elem().Interface())
	}
	return false
}
