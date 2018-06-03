package refutil

import "testing"

func TestLen(t *testing.T) {
	var header []string
	tests := []struct {
		x          interface{}
		wantLength int
		ok         bool
	}{
		{[]uint{}, 0, true},
		{[]uint{1}, 1, true},
		{"", 0, true},
		{"1", 1, true},
		{"10", 2, true},
		{map[string]string{}, 0, true},
		{map[string]string{"1": "1"}, 1, true},
		{map[string]string{"1": "1", "2": "2"}, 2, true},
		{header, 0, true},
		{struct{}{}, 0, false},
	}
	for _, tt := range tests {
		t.Run("test len", func(t *testing.T) {
			gotLength, isOk := Len(tt.x)
			if isOk != tt.ok {
				t.Errorf("Len() error = %v, wantErr %v", isOk, tt.ok)
				return
			}
			if gotLength != tt.wantLength {
				t.Errorf("Len() = %v, want %v", gotLength, tt.wantLength)
			}
		})
	}
}
