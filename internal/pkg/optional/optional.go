package optional

// Optional is a type for optional values.
type Optional[T any] struct {
	value   T
	isEmpty bool
}

// Of returns an Optional with a value.
func Of[T any](value T) Optional[T] {
	return Optional[T]{
		value:   value,
		isEmpty: false,
	}
}

// IsPresent checks if an Optional has a value.
// It returns true if an Optional is present.
func (o Optional[T]) IsPresent() bool {
	return !o.isEmpty
}

// IsEmpty checks if an Optional has no value.
// It returns true if an Optional is empty.
func (o Optional[T]) IsEmpty() bool {
	return o.isEmpty
}

// Empty returns an empty Optional.
func Empty[T any]() Optional[T] {
	return Optional[T]{
		isEmpty: true,
	}
}

// Set sets a value to an Optional.
func (o Optional[T]) Set(value T) Optional[T] {
	return Optional[T]{
		value:   value,
		isEmpty: false,
	}
}

// Value returns a value of an Optional.
func (o Optional[T]) Value() T {
	return o.value
}
