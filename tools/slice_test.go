package tools_test

import (
	"reflect"
	"testing"

	"github.com/P-E-R-R-Y/golang-utils/tools"
)

func TestContains(t *testing.T) {
	ints := []int{1, 2, 3, 4}
	if !tools.Contains(ints, 3) {
		t.Errorf("Contains should return true for existing value")
	}
	if tools.Contains(ints, 5) {
		t.Errorf("Contains should return false for missing value")
	}

	strs := []string{"a", "b", "c"}
	if !tools.Contains(strs, "a") {
		t.Errorf(`Contains should return true for "a"`)
	}
	if tools.Contains(strs, "z") {
		t.Errorf(`Contains should return false for "z"`)
	}
}

func TestCloneSlice(t *testing.T) {
	original := []int{10, 20, 30}
	clone := tools.CloneSlice(original)

	if !reflect.DeepEqual(original, clone) {
		t.Errorf("CloneSlice did not copy slice correctly: got %v, want %v", clone, original)
	}
	if &original[0] == &clone[0] {
		t.Errorf("CloneSlice should return a new slice (shallow copy)")
	}
}

func TestFilter(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6}
	even := func(n int) bool { return n%2 == 0 }

	result := tools.Filter(numbers, even)
	expected := []int{2, 4, 6}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Filter failed: got %v, want %v", result, expected)
	}
}