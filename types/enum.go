package types

import "github.com/PlayerR9/GoSD/pkg"

type Enumer interface {
	~int

	// String returns the literal representation of the enum value.
	//
	// Returns:
	//   - string: The literal representation of the enum value.
	String() string
}

// Enum is an enum.
type Enum[T Enumer] struct {
	// value is the enum value.
	value T
}

// String implements the pkg.Type interface.
func (e *Enum[T]) String() string {
	return e.value.String()
}

// DeepCopy implements the pkg.Type interface.
func (e *Enum[T]) DeepCopy() pkg.Type {
	return &Enum[T]{
		value: e.value,
	}
}

// Ensure implements the pkg.Type interface.
func (e *Enum[T]) Ensure() {
	pkg.ThrowIf(e == nil, pkg.NewInvalidState("e", pkg.NewNilValue()))
}

// Clean implements the pkg.Type interface.
func (e *Enum[T]) Clean() {}

// Equals implements the pkg.Type interface.
//
// Two enums are equal if they have the same value and are both enums.
func (e *Enum[T]) Equals(other pkg.Type) bool {
	pkg.Ensure(false, e)
	pkg.Ensure(false, other)

	switch other := other.(type) {
	case *Enum[T]:
		return e.value == other.value
	default:
		return false
	}
}

// NewEnum creates a new enum.
//
// Parameters:
//   - value: The value.
//
// Returns:
//   - *Enum: The new enum. Never returns nil.
func NewEnum[T Enumer](value T) *Enum[T] {
	return &Enum[T]{
		value: value,
	}
}

// Value returns the enum value.
//
// Returns:
//   - T: The enum value.
func (e Enum[T]) Value() T {
	return e.value
}

// Set sets the enum value.
//
// Parameters:
//   - value: The new value.
func (e *Enum[T]) Set(value T) {
	pkg.Ensure(false, e)

	e.value = value
}
