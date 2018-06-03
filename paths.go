package refutil

import (
	"reflect"
)

// PathTo will return actual path to package
// where interface{} is defined.
func PathTo(v interface{}) string {
	return reflect.TypeOf(Dereference(v)).PkgPath()
}
