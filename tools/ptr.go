package tools

// Default returns 'val' if it's not the zero value for T,
// otherwise returns the fallback value 'def'.
func Default[T comparable](val, def T) T {
	var zero T
	if val == zero {
		return def
	}
	return val
}

// Ptr returns a pointer to the given value.
func Ptr[T any](val T) *T {
	return &val
}

// Deref returns the value pointed to by ptr if ptr is not nil,
// otherwise returns the fallback value def.
func Deref[T any](ptr *T, def T) T {
	if ptr == nil {
		return def
	}
	return *ptr
}