package tools_test

import (
	"testing"

	"github.com/P-E-R-R-Y/golang-utils/tools"
)

func TestTernary_Int(t *testing.T) {
	result := tools.Ternary(true, 1, 2)
	if result != 1 {
		t.Errorf("Ternary(true, 1, 2) = %d; want 1", result)
	}

	result = tools.Ternary(false, 1, 2)
	if result != 2 {
		t.Errorf("Ternary(false, 1, 2) = %d; want 2", result)
	}
}

func TestTernary_String(t *testing.T) {
	result := tools.Ternary(true, "yes", "no")
	if result != "yes" {
		t.Errorf(`Ternary(true, "yes", "no") = %q; want "yes"`, result)
	}

	result = tools.Ternary(false, "yes", "no")
	if result != "no" {
		t.Errorf(`Ternary(false, "yes", "no") = %q; want "no"`, result)
	}
}

func TestTernary_Struct(t *testing.T) {
	type S struct{ Val int }

	a := S{Val: 10}
	b := S{Val: 20}

	result := tools.Ternary(true, a, b)
	if result != a {
		t.Errorf("Ternary(true, a, b) = %v; want %v", result, a)
	}

	result = tools.Ternary(false, a, b)
	if result != b {
		t.Errorf("Ternary(false, a, b) = %v; want %v", result, b)
	}
}