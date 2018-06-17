package refutil

// FindIndex will search for index with `Comparator`. This is
// useful for your own comparsion
// func (v Data) FindIndex(element interface{}, compare Comparator) (index int) {
// 	arr := v.Indirect()
// 	if arr.InterfaceOrNil() == nil {
// 		return -1
// 	}
// 	if !arr.KindOneOf(reflect.Array, reflect.Slice) {
// 		panic(ErrArgumentNotSlice)
// 	}

// 	val := arr.Value()
// 	for i := 0; i < val.Len(); i++ {
// 		if e.CompareInterfacer(compare, val.Index(i)) {
// 			return i
// 		}
// 	}
// 	return -1
// }

// // Compare current value to another with Comparator interface
// func (v Data) Compare(comparator Comparator, element interface{}) bool {
// 	if !v.CanInterface() {
// 		return false
// 	}
// 	return comparator(element)
// }

// // CompareInterfacer current value to Value with Comparator interface
// func (v Data) CompareInterfacer(comparator Comparator, i Interfacer) bool {
// 	if !v.CanInterface() {
// 		return false
// 	}
// 	if !i.CanInterface() {
// 		return false
// 	}
// 	return comparator(i.Interface())
// }

// // CanIndex returns whether is possible to search in interface{}
// func CanIndex(source interface{}, element interface{}) bool {
// 	t := IndirectTypeOf(source)
// 	return KindOneOf(t, reflect.String, reflect.Map, reflect.Array, reflect.Slice)
// }

// // FindMapKeyByValue will search through map with `Comparator` and
// // return key if values are matching
// func FindMapKeyByValue(source interface{}, element interface{}, compare Comparator) (interface{}, bool) {
// 	s := NewData(source)
// 	if !s.KindOneOf(reflect.Map) {
// 		panic(ErrInvalidArgument)
// 	}
// 	e := NewData(element)
// 	val := s.Value()
// 	keys := val.MapKeys()
// 	for i := 0; i < len(keys); i++ {
// 		value := val.MapIndex(keys[i])
// 		if e.CompareInterfacer(compare, value) {
// 			return keys[i].Interface(), true
// 		}
// 	}
// 	return nil, false
// }

// // FindMapValueByKey will search through map with `Comparator` and
// // retrun value if keys matching
// func FindMapValueByKey(source interface{}, element interface{}, compare Comparator) (interface{}, bool) {
// 	s := NewData(source)
// 	if !s.KindOneOf(reflect.Map) {
// 		panic(ErrInvalidArgument)
// 	}
// 	e := NewData(element)
// 	val := s.Value()
// 	keys := val.MapKeys()
// 	for i := 0; i < len(keys); i++ {
// 		if e.CompareInterfacer(compare, keys[i]) {
// 			v := val.MapIndex(keys[i])
// 			if !v.CanInterface() {
// 				return nil, false
// 			}
// 			return v.Interface(), true
// 		}
// 	}
// 	return nil, false
// }

// // FindStringIndex will search for index in string. If type has
// // fmt.Stringer interface it will use it
// func FindStringIndex(source interface{}, element interface{}) (index int) {
// 	s := NewData(source)
// 	e := NewData(element)
// 	return strings.Index(s.String(), e.String())
// }

// // Index will search in source and look for same value using `IsEqual` method,
// // panics if source is not searchable  otherwise return int as index
// func Index(source interface{}, value interface{}) (index int) {
// 	return FindIndex(source, value, Equal)
// }

// // IndexSame will search in source and look for same value using `IsDeepEqual` method,
// // panics if source is not searchable  otherwise return int as index.
// // Note that this variant of search care about underlying type
// func IndexSame(source interface{}, value interface{}) (index int) {
// 	return FindIndex(source, value, DeepEqual)
// }

// // Contains will search in source and look for same value using `IsEqual` method,
// // panics if source is not searchable otherwise return bool if found
// func Contains(source interface{}, value interface{}) (found bool) {
// 	index := FindIndex(source, value, Equal)
// 	return index > -1
// }

// // ContainsSame will search in source and look for same value using `IsEqual` method,
// // panics if source is not searchable otherwise return bool if found.
// // Note that this variant of search care about underlying type
// func ContainsSame(source interface{}, value interface{}) (found bool) {
// 	index := FindIndex(source, value, DeepEqual)
// 	return index > -1
// }
