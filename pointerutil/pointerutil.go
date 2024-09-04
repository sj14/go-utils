package pointerutil

func PointerTo[T any](v T) *T {
	return &v
}

// Returns the value of 'v' or the zero value if 'v' is nil.
func ValueFrom[T any](v *T) T {
	var zero T
	if v == nil {
		return zero
	}
	return *v
}
