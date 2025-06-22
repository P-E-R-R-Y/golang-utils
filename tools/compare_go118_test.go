package tools_test

import (
	"testing"

	"github.com/P-E-R-R-Y/golang-utils/tools"
)

// TestMinValidTypes tests Min with valid Ordered types.
func TestMinValidTypes(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		got := tools.Min(10, 20)
		want := 10
		if got != want {
			t.Errorf("Min(10, 20) = %d; want %d", got, want)
		}
	})

	t.Run("float64", func(t *testing.T) {
		got := tools.Min(3.14, 2.71)
		want := 2.71
		if got != want {
			t.Errorf("Min(3.14, 2.71) = %v; want %v", got, want)
		}
	})

	t.Run("string", func(t *testing.T) {
		got := tools.Min("apple", "banana")
		want := "apple"
		if got != want {
			t.Errorf(`Min("apple", "banana") = %q; want %q`, got, want)
		}
	})
}

// TestMaxValidTypes tests Max with valid Ordered types.
func TestMaxValidTypes(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		got := tools.Max(10, 20)
		want := 20
		if got != want {
			t.Errorf("Max(10, 20) = %d; want %d", got, want)
		}
	})

	t.Run("float64", func(t *testing.T) {
		got := tools.Max(3.14, 2.71)
		want := 3.14
		if got != want {
			t.Errorf("Max(3.14, 2.71) = %v; want %v", got, want)
		}
	})

	t.Run("string", func(t *testing.T) {
		got := tools.Max("apple", "banana")
		want := "banana"
		if got != want {
			t.Errorf(`Max("apple", "banana") = %q; want %q`, got, want)
		}
	})
}