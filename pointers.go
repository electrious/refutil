package refutil

import "reflect"

// Dereference will return actual value insted of pointer
// if values is passed value is returned
func Dereference(v interface{}) interface{} {
	if IsNil(v) {
		return nil
	}
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr {
		return v
	}
	return val.Elem().Interface()
}
