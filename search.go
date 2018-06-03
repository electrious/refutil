package refutil

import (
	"reflect"
	"strings"
)

type comparator func(interface{}, interface{}) bool

// include try loop over the list check if the list includes the element.
// return (false, -1) if impossible.
// return (true, -1) if element was not found.
// return (true, ?) if element was found.
func include(list interface{}, element interface{}, c comparator) (index int, ok bool) {
	listValue := reflect.ValueOf(list)
	elementValue := reflect.ValueOf(element)
	defer func() {
		if e := recover(); e != nil {
			ok = false
			index = -1
		}
	}()
	if reflect.TypeOf(list).Kind() == reflect.String {
		return strings.Index(listValue.String(), elementValue.String()), true
	}
	if reflect.TypeOf(list).Kind() == reflect.Map {
		mapKeys := listValue.MapKeys()
		for i := 0; i < len(mapKeys); i++ {
			if c(mapKeys[i].Interface(), element) {
				return i, true
			}
		}
		return -1, true
	}
	for i := 0; i < listValue.Len(); i++ {
		if c(listValue.Index(i).Interface(), element) {
			return i, true
		}
	}
	return -1, true
}

// Index will search in list and return index / position if possible
// return (false, -1) if impossible.
// return (true, -1) if element was not found.
// return (true, ?) if element was found.
func Index(source interface{}, value interface{}) (index int, ok bool) {
	return include(source, value, IsEqual)
}

// IndexSame will search in list and return index / position if possible
// return (false, -1) if impossible.
// return (true, -1) if element was not found.
// return (true, ?) if element was found.
// note that this variant of search care about underlying type
func IndexSame(source interface{}, value interface{}) (index int, ok bool) {
	return include(source, value, IsDeepEqual)
}

// Contains will search for element in source. If possible to search
// first returned argument will be true and if found
// second argument will be true as well. This method is very similar
// to Index if index is > -1 found is true
func Contains(source interface{}, value interface{}) (found bool, ok bool) {
	index, done := include(source, value, IsEqual)
	if !done {
		return
	}
	if index == -1 {
		ok = true
		return
	}
	ok = true
	found = true
	return
}

// ContainsSame will search for element in source. If possible to search
// first returned argument will be true and if found
// second argument will be true as well. This method is very similar
// to Index if index is > -1 found is true
// note that this variant of search care about underlying type
func ContainsSame(source interface{}, value interface{}) (found bool, ok bool) {
	index, done := include(source, value, IsDeepEqual)
	if !done {
		return
	}
	if index == -1 {
		ok = true
		return
	}
	ok = true
	found = true
	return
}
