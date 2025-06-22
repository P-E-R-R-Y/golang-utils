package tools

// Ternary returns a if cond is true, else b.
func Ternary[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}
