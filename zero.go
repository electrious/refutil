package refutil

import "reflect"

// Zeroer ...
type Zeroer interface {
	IsZero() bool
}

// isZeroSlice will check if underlying slice type is equal to
// zero of its own type
func isZeroSlice(compare interface{}) bool {
	typ := reflect.TypeOf(compare)
	sliceType := reflect.SliceOf(typ.Elem())
	zeroLenSlice := reflect.MakeSlice(sliceType, 0, 0)
	convertedZeroLenSlice := zeroLenSlice.Convert(typ)
	return reflect.DeepEqual(compare, convertedZeroLenSlice.Interface())
}

// IsZero return true if underlying type is equal to its zero value
func IsZero(compare interface{}) bool {
	if IsNil(compare) {
		return true
	}
	if zeroer, k := compare.(Zeroer); k {
		return zeroer.IsZero()
	}
	typ := reflect.TypeOf(compare)
	value := reflect.ValueOf(compare)
	zero := reflect.Zero(typ).Interface()
	isZero := reflect.DeepEqual(compare, zero)
	if isZero {
		return true
	}
	switch typ.Kind() {
	case reflect.Map:
		return value.Len() == 0
	case reflect.Chan:
		return value.Len() == 0
	case reflect.Slice:
		return isZeroSlice(compare)
	case reflect.Ptr:
		{
			if value.IsNil() {
				return true
			}
			el := value.Elem().Interface()
			emptyEl := reflect.New(typ.Elem()).Elem().Interface()
			return el == emptyEl
		}
	default:
		return false
	}
}
