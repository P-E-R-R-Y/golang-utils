package tools

// Contains returns true if val is in slice.
func Contains[T comparable](slice []T, val T) bool {
	for _, s := range slice {
		if s == val {
			return true
		}
	}
	return false
}

// CloneSlice returns a shallow copy of src slice.
func CloneSlice[T any](src []T) []T {
	dst := make([]T, len(src))
	copy(dst, src)
	return dst
}

func Filter[T any](in []T, keep func(T) bool) []T {
	out := make([]T, 0, len(in))
	for _, v := range in {
		if keep(v) {
			out = append(out, v)
		}
	}
	return out
}