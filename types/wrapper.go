package types

import (
	"fmt"

	"github.com/PlayerR9/GoSD/pkg"
)

// Wrap is a wrapper.
type Wrap[T comparable] struct {
	// value is the wrapped value.
	value T
}

// String implements the fmt.Stringer interface.
func (w *Wrap[T]) String() string {
	return fmt.Sprint(w.value)
}

// DeepCopy implements the pkg.Type interface.
func (w *Wrap[T]) DeepCopy() pkg.Type {
	return &Wrap[T]{
		value: w.value,
	}
}

// Ensure implements the pkg.Type interface.
func (w *Wrap[T]) Ensure() {
	pkg.ThrowIf(w == nil, pkg.NewInvalidState("w", pkg.NewNilValue()))
}

// Clean implements the pkg.Type interface.
func (w *Wrap[T]) Clean() {}

// Equals implements the pkg.Type interface.
//
// Two wraps are equal if they have the same value adn are both wraps.
func (w *Wrap[T]) Equals(other pkg.Type) bool {
	pkg.Ensure(false, w)
	pkg.Ensure(false, other)

	switch other := other.(type) {
	case *Wrap[T]:
		return w.value == other.value
	default:
		return false
	}
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
