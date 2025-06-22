//go:build go1.18
// +build go1.18

package tools

// Min returns the smaller of a or b.
// Requires Go 1.18+ and ordered types (numeric, string).
func Min[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Max returns the larger of a or b.
// Requires Go 1.18+ and ordered types (numeric, string).
func Max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}