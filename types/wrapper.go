package types

import "github.com/PlayerR9/GoSD/pkg"

// Wrap is a wrapper.
type Wrap[T comparable] struct {
	// value is the wrapped value.
	value T
}

// Clean implements the pkg.Type interface.
func (w *Wrap[T]) Clean() {}

// Equals implements the pkg.Type interface.
//
// Two wraps are equal if they have the same value adn are both wraps.
func (w *Wrap[T]) Equals(other pkg.Type) bool {
	if other == nil {
		panic(pkg.NewNilComparison("other"))
	}

	other_val, ok := other.(*Wrap[T])
	if !ok {
		return false
	}

	return w.value == other_val.value
}

// NewWrap creates a new wrap.
//
// Parameters:
//   - value: The value to wrap.
//
// Returns:
//   - *Wrap: The new wrap. Never returns nil.
func NewWrap[T comparable](value T) *Wrap[T] {
	return &Wrap[T]{
		value: value,
	}
}

// Value returns the wrapped value.
//
// Returns:
//   - T: The wrapped value.
func (w Wrap[T]) Value() T {
	return w.value
}

// Set sets the wrapped value.
//
// Parameters:
//   - value: The new value.
func (w *Wrap[T]) Set(value T) {
	w.value = value
}
