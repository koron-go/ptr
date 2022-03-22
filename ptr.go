package ptr

// P obtains a pointer of a value.
func P[T any](v T) *T {
	return &v
}

// V obtains a value which refered by a pointer.
func V[T any](p *T) T {
	if p == nil {
		var zero T
		return zero
	}
	return *p
}
