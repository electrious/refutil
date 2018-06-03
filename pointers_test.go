package refutil

import (
	"reflect"
	"testing"
)

type teststruct struct {
	A int
}

func TestDereference(t *testing.T) {
	example1 := teststruct{A: 1}
	var a *bool
	tests := []struct {
		v    interface{}
		want interface{}
	}{
		{&example1, example1},
		{example1, example1},
		{nil, nil},
		{a, nil},
	}
	for _, tt := range tests {
		t.Run("test dereference", func(t *testing.T) {
			if got := Dereference(tt.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dereference() = %v, want %v", got, tt.want)
			}
		})
	}
}
