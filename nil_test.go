package refutil

import (
	"testing"
	"time"
)

func TestNil(t *testing.T) {
	var header []string
	var header2 = []string{"ok"}
	tests := []struct {
		x    interface{}
		want bool
	}{
		{&header, true},
		{nil, true},
		{(*time.Time)(nil), true},
		{header, true},
		{header2, false},
		{(*time.Time)(&time.Time{}), false},
		{1, false},
		{0, false},
	}
	for _, tt := range tests {
		t.Run("test nil", func(t *testing.T) {
			if got := IsNil(tt.x); got != tt.want {
				t.Errorf("Nil() = %v, want %v", got, tt.want)
			}
		})
	}
}
