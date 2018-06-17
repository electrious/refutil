package refutil

// func TestComparator(t *tesingt.T) {

// }

// func TestIncludeElement(t *testing.T) {
// 	truthful := func(val bool) {
// 		if !val {
// 			t.Fail()
// 		}
// 	}
// 	falsey := func(val bool) {
// 		if val {
// 			t.Fail()
// 		}
// 	}
// 	list1 := []string{"Foo", "Bar"}
// 	list2 := []int{1, 2}
// 	simpleMap := map[interface{}]interface{}{"Foo": "Bar"}
// 	found, ok := include("Hello World", "World", IsEqual)
// 	truthful(ok)
// 	truthful(found != -1)
// 	found, ok = include(list1, "Foo", IsEqual)
// 	truthful(ok)
// 	truthful(found != -1)
// 	found, ok = include(list1, "Bar", IsEqual)
// 	truthful(ok)
// 	truthful(found != -1)
// 	found, ok = include(list2, 1, IsEqual)
// 	truthful(ok)
// 	truthful(found != -1)
// 	found, ok = include(list2, 2, IsEqual)
// 	truthful(ok)
// 	truthful(found != -1)
// 	found, ok = include(list1, "Foo!", IsEqual)
// 	truthful(ok)
// 	falsey(found != -1)
// 	found, ok = include(list2, 3, IsEqual)
// 	truthful(ok)
// 	falsey(found != -1)
// 	found, ok = include(list2, "1", IsEqual)
// 	truthful(ok)
// 	falsey(found != -1)
// 	found, ok = include(simpleMap, "Foo", IsEqual)
// 	truthful(ok)
// 	truthful(found != -1)
// 	found, ok = include(simpleMap, "Bar", IsEqual)
// 	truthful(ok)
// 	falsey(found != -1)
// 	found, ok = include(1433, "1", IsEqual)
// 	falsey(ok)
// 	falsey(found != -1)
// }
