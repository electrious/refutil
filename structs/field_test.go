package structs

import (
	"reflect"
	"testing"
)

// A test struct that defines all cases
type Foo struct {
	A    string
	B    int    `structs:"y"`
	C    bool   `json:"c"`
	d    string // not exported
	E    *Baz
	x    string `xml:"x"` // not exported, with tag
	Y    []string
	Z    map[string]interface{}
	*Bar // embedded
}

type Baz struct {
	A string
	B int
}

type Bar struct {
	E string
	F int
	g []string
}

func newStruct() *Struct {
	b := &Bar{
		E: "example",
		F: 2,
		g: []string{"zeynep", "fatih"},
	}

	// B and x is not initialized for testing
	f := &Foo{
		A: "gopher",
		C: true,
		d: "small",
		E: nil,
		Y: []string{"example"},
		Z: nil,
	}
	f.Bar = b

	return New(f)
}

func TestField_Set(t *testing.T) {
	s := newStruct()

	f, _ := s.Field("A")
	err := f.Set("fatih")
	if err != nil {
		t.Error(err)
	}

	if f.Value().(string) != "fatih" {
		t.Errorf("Setted value is wrong: %s want: %s", f.Value().(string), "fatih")
	}

	f, _ = s.Field("Y")
	err = f.Set([]string{"override", "with", "this"})
	if err != nil {
		t.Error(err)
	}

	sliceLen := len(f.Value().([]string))
	if sliceLen != 3 {
		t.Errorf("Setted values slice length is wrong: %d, want: %d", sliceLen, 3)
	}

	f, _ = s.Field("C")
	err = f.Set(false)
	if err != nil {
		t.Error(err)
	}

	if f.Value().(bool) {
		t.Errorf("Setted value is wrong: %t want: %t", f.Value().(bool), false)
	}

	// let's pass a different type
	f, _ = s.Field("A")
	err = f.Set(123) // Field A is of type string, but we are going to pass an integer
	if err == nil {
		t.Error("Setting a field's value with a different type than the field's type should return an error")
	}

	// old value should be still there :)
	if f.Value().(string) != "fatih" {
		t.Errorf("Setted value is wrong: %s want: %s", f.Value().(string), "fatih")
	}

	// let's access an unexported field, which should give an error
	f, _ = s.Field("d")
	err = f.Set("large")
	if err != errNotExported {
		t.Error(err)
	}

	// let's set a pointer to struct
	b := &Bar{
		E: "gopher",
		F: 2,
	}

	f, _ = s.Field("Bar")
	err = f.Set(b)
	if err != nil {
		t.Error(err)
	}

	baz := &Baz{
		A: "helloWorld",
		B: 42,
	}

	f, _ = s.Field("E")
	err = f.Set(baz)
	if err != nil {
		t.Error(err)
	}

	ba, _ := s.Field("E")
	ba2 := ba.Value().(*Baz)

	if ba2.A != "helloWorld" {
		t.Errorf("could not set baz. Got: %s Want: helloWorld", ba2.A)
	}
}

func TestField_NotSettable(t *testing.T) {
	a := map[int]Baz{
		4: {
			A: "value",
		},
	}

	s := New(a[4])
	x, _ := s.Field("A")

	if err := x.Set("newValue"); err != errNotSettable {
		t.Errorf("Trying to set non-settable field should error with %q. Got %q instead.", errNotSettable, err)
	}
}

func TestField_Zero(t *testing.T) {
	s := newStruct()

	f, _ := s.Field("A")
	err := f.Zero()
	if err != nil {
		t.Error(err)
	}

	if f.Value().(string) != "" {
		t.Errorf("Zeroed value is wrong: %s want: %s", f.Value().(string), "")
	}

	f, _ = s.Field("Y")
	err = f.Zero()
	if err != nil {
		t.Error(err)
	}

	sliceLen := len(f.Value().([]string))
	if sliceLen != 0 {
		t.Errorf("Zeroed values slice length is wrong: %d, want: %d", sliceLen, 0)
	}

	f, _ = s.Field("C")
	err = f.Zero()
	if err != nil {
		t.Error(err)
	}

	if f.Value().(bool) {
		t.Errorf("Zeroed value is wrong: %t want: %t", f.Value().(bool), false)
	}

	// let's access an unexported field, which should give an error
	f, _ = s.Field("d")
	err = f.Zero()
	if err != errNotExported {
		t.Error(err)
	}

	f, _ = s.Field("Bar")
	err = f.Zero()
	if err != nil {
		t.Error(err)
	}

	f, _ = s.Field("E")
	err = f.Zero()
	if err != nil {
		t.Error(err)
	}

	v, _ := s.Field("E")
	if !v.value.IsNil() {
		t.Errorf("could not set baz. Got: %s Want: <nil>", v.value.Interface())
	}
}

func TestField_Kind(t *testing.T) {
	s := newStruct()

	f, _ := s.Field("A")
	if f.Kind() != reflect.String {
		t.Errorf("Field A has wrong kind: %s want: %s", f.Kind(), reflect.String)
	}

	f, _ = s.Field("B")
	if f.Kind() != reflect.Int {
		t.Errorf("Field B has wrong kind: %s want: %s", f.Kind(), reflect.Int)
	}

	// unexported
	f, _ = s.Field("d")
	if f.Kind() != reflect.String {
		t.Errorf("Field d has wrong kind: %s want: %s", f.Kind(), reflect.String)
	}
}

func TestField_Tag(t *testing.T) {
	s := newStruct()

	v, _ := s.Field("B")
	x := v.Tag("json")
	if x != "" {
		t.Errorf("Field's tag value of a non existing tag should return empty, got: %s", x)
	}

	v, _ = s.Field("C")
	x = v.Tag("json")
	if x != "c" {
		t.Errorf("Field's tag value of the existing field C should return 'c', got: %s", x)
	}

	v, _ = s.Field("d")
	x = v.Tag("json")
	if x != "" {
		t.Errorf("Field's tag value of a non exported field should return empty, got: %s", x)
	}

	v, _ = s.Field("x")
	x = v.Tag("xml")
	if x != "x" {
		t.Errorf("Field's tag value of a non exported field with a tag should return 'x', got: %s", x)
	}

	v, _ = s.Field("A")
	x = v.Tag("json")
	if x != "" {
		t.Errorf("Field's tag value of a existing field without a tag should return empty, got: %s", x)
	}
}

func TestField_Value(t *testing.T) {
	s := newStruct()

	a, _ := s.Field("A")
	x := a.Value()
	val, ok := x.(string)
	if !ok {
		t.Errorf("Field's value of a A should be string")
	}

	if val != "gopher" {
		t.Errorf("Field's value of a existing tag should return 'gopher', got: %s", val)
	}

	defer func() {
		err := recover()
		if err == nil {
			t.Error("Value of a non exported field from the field should panic")
		}
	}()

	// should panic
	d, _ := s.Field("d")
	d.Value()
}

func TestField_IsEmbedded(t *testing.T) {
	s := newStruct()
	b, _ := s.Field("Bar")
	if !b.IsEmbedded() {
		t.Errorf("Fields 'Bar' field is an embedded field")
	}
	d, _ := s.Field("d")
	if d.IsEmbedded() {
		t.Errorf("Fields 'd' field is not an embedded field")
	}
}

func TestField_IsExported(t *testing.T) {
	s := newStruct()
	b, _ := s.Field("Bar")
	if !b.IsExported() {
		t.Errorf("Fields 'Bar' field is an exported field")
	}
	a, _ := s.Field("A")
	if !a.IsExported() {
		t.Errorf("Fields 'A' field is an exported field")
	}
	d, _ := s.Field("d")
	if d.IsExported() {
		t.Errorf("Fields 'd' field is not an exported field")
	}
}

func TestField_IsZero(t *testing.T) {
	s := newStruct()

	if f, _ := s.Field("A"); f.IsZero() {
		t.Errorf("Fields 'A' field is an initialized field")
	}

	if f, _ := s.Field("B"); !f.IsZero() {
		t.Errorf("Fields 'B' field is not an initialized field")
	}
}

func TestField_Name(t *testing.T) {
	s := newStruct()

	if f, _ := s.Field("A"); f.Name() != "A" {
		t.Errorf("Fields 'A' field should have the name 'A'")
	}
}

func TestField_Field(t *testing.T) {
	s := newStruct()

	b, _ := s.Field("Bar")
	e, _ := b.Field("E")
	val, ok := e.Value().(string)
	if !ok {
		t.Error("The value of the field 'e' inside 'Bar' struct should be string")
	}

	if val != "example" {
		t.Errorf("The value of 'e' should be 'example, got: %s", val)
	}

}

func TestField_Fields(t *testing.T) {
	s := newStruct()
	f, _ := s.Field("Bar")
	fields := f.Fields()

	if len(fields) != 3 {
		t.Errorf("We expect 3 fields in embedded struct, was: %d", len(fields))
	}
}
