package tools_test

import (
	"testing"

	"github.com/P-E-R-R-Y/golang-utils/tools"
)

func TestDefault(t *testing.T) {
	tests := []struct {
		name     string
		val      int
		def      int
		expected int
	}{
		{"non-zero value", 42, 99, 42},
		{"zero value", 0, 99, 99},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tools.Default(tt.val, tt.def)
			if result != tt.expected {
				t.Errorf("Default(%v, %v) = %v; want %v", tt.val, tt.def, result, tt.expected)
			}
		})
	}
}

func TestPtr(t *testing.T) {
	val := 123
	ptr := tools.Ptr(val)
	if ptr == nil {
		t.Fatal("Ptr returned nil")
	}
	if *ptr != val {
		t.Errorf("Ptr(%d) = %d; want %d", val, *ptr, val)
	}
}

func TestDeref(t *testing.T) {
	val := 123
	ptr := &val
	if got := tools.Deref(ptr, 0); got != 123 {
		t.Errorf("Deref(ptr, 0) = %v; want 123", got)
	}

	if got := tools.Deref[int](nil, 456); got != 456 {
		t.Errorf("Deref(nil, 456) = %v; want 456", got)
	}
}