package types

import (
	"iter"

	"github.com/PlayerR9/GoSD/pkg"
)

// Bool is a bool.
type Bool struct {
	// value is the bool value.
	value bool
}

// String implements the pkg.Type interface.
func (b *Bool) String() string {
	if b.value {
		return "T"
	} else {
		return "F"
	}
}

// DeepCopy implements the pkg.Type interface.
func (b *Bool) DeepCopy() pkg.Type {
	return &Bool{
		value: b.value,
	}
}

// Ensure implements the pkg.Type interface.
func (b *Bool) Ensure() {
	pkg.ThrowIf(b == nil, pkg.NewInvalidState("b", pkg.NewNilValue()))
}

// Clean implements the pkg.Type interface.
func (b *Bool) Clean() {}

// Equals implements the pkg.Type interface.
//
// Two bools are equal if they have the same value and are both bools.
func (b *Bool) Equals(other pkg.Type) bool {
	pkg.Ensure(false, b)
	pkg.Ensure(false, other)

	switch other := other.(type) {
	case *Bool:
		return b.value == other.value
	default:
		return false
	}
}

// NewBool creates a new bool initialized to false.
//
// Returns:
//   - *Bool: The new bool. Never returns nil.
func NewBool() *Bool {
	return &Bool{
		value: false,
	}
}

// WithValue creates a new bool with the given value.
//
// Parameters:
//   - value: The value.
//
// Returns:
//   - *Bool: The new bool. Never returns nil.
func (b *Bool) WithValue(value bool) *Bool {
	if b == nil {
		return &Bool{
			value: value,
		}
	}

	b.value = value

	return b
}

// Value returns the bool value.
//
// Returns:
//   - bool: The bool value.
func (b Bool) Value() bool {
	return b.value
}

// Set sets the bool value.
//
// Parameters:
//   - value: The new value.
func (b *Bool) Set(value bool) {
	pkg.Ensure(false, b)

	b.value = value
}

// Each creates an iterator that iterates over the bool.
//
// Returns:
//   - iter.Seq[*Bool]: The iterator. Never returns nil.
//
// The iterator always returns until a break, return statement is reached, or the value
// given by the iterator is set to false.
func (b *Bool) Each() iter.Seq[*Bool] {
	pkg.Ensure(false, b)

	fn := func(yield func(*Bool) bool) {
		for b.value {
			if !yield(b) {
				return
			}
		}
	}

	return fn
}
