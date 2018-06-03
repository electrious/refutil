package refutil

import "reflect"

// Len try to get length of object.
// return (false, 0) if impossible.
func Len(x interface{}) (length int, ok bool) {
	v := reflect.ValueOf(x)
	defer func() {
		if e := recover(); e != nil {
			ok = false
			length = 0
		}
	}()
	return v.Len(), true
}
