package refutil

import "testing"

func TestIsEqual(t *testing.T) {
	if !IsDeepEqual("Hello World", "Hello World") {
		t.Error("objectsAreEqual should return true")
	}
	if !IsDeepEqual(123, 123) {
		t.Error("objectsAreEqual should return true")
	}
	if !IsDeepEqual(123.5, 123.5) {
		t.Error("objectsAreEqual should return true")
	}
	if !IsDeepEqual([]byte("Hello World"), []byte("Hello World")) {
		t.Error("objectsAreEqual should return true")
	}
	if !IsDeepEqual(nil, nil) {
		t.Error("objectsAreEqual should return true")
	}
	if IsDeepEqual(map[int]int{5: 10}, map[int]int{10: 20}) {
		t.Error("objectsAreEqual should return false")
	}
	if IsDeepEqual('x', "x") {
		t.Error("objectsAreEqual should return false")
	}
	if IsDeepEqual("x", 'x') {
		t.Error("objectsAreEqual should return false")
	}
	if IsDeepEqual(0, 0.1) {
		t.Error("objectsAreEqual should return false")
	}
	if IsDeepEqual(0.1, 0) {
		t.Error("objectsAreEqual should return false")
	}
	if IsDeepEqual(uint32(10), int32(10)) {
		t.Error("objectsAreEqual should return false")
	}
	if !IsEqual(uint32(10), int32(10)) {
		t.Error("IsEqual should return true")
	}
	if IsEqual(0, nil) {
		t.Fail()
	}
	if IsEqual(nil, 0) {
		t.Fail()
	}

}
