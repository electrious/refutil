package refutil

import "reflect"

// Equalizer is way how to check if object
// can compare it self to another object. Used
// only with `Equal` method
type Equalizer interface {
	Equal(v interface{}) bool
}

// Kinder as interface for reflect.Value or reflect.Type Kind() method
type Kinder interface {
	Kind() reflect.Kind
}

// Interfacer is subset functionality
// reflect.Value is providing
type Interfacer interface {
	CanInterface() bool
	Interface() interface{}
}

// Comparator is used as abstraction for method
// like IsEqual which compare two values
type Comparator func(interface{}) bool

// Zeroer ...
type Zeroer interface {
	IsZero() bool
}
