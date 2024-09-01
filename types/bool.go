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

// Clean implements the pkg.Type interface.
func (b Bool) Clean() {}

// Equals implements the pkg.Type interface.
//
// Two bools are equal if they have the same value and are both bools.
func (b Bool) Equals(other pkg.Type) bool {
	if other == nil {
		panic(pkg.NewNilComparison("other"))
	}

	other_val, ok := other.(*Bool)
	if !ok {
		return false
	}

	return b.value == other_val.value
}

// NewBool creates a new bool initialized to false.
//
// Returns:
//   - Bool: The new bool.
func NewBool() Bool {
	return Bool{
		value: false,
	}
}

// WithValue creates a new bool with the given value.
//
// Parameters:
//   - value: The value.
//
// Returns:
//   - Bool: The new bool.
func (b Bool) WithValue(value bool) Bool {
	return Bool{
		value: value,
	}
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
	b.value = value
}

// Each creates an iterator that iterates over the bool.
//
// Returns:
//   - iter.Seq[*Bool]: The iterator. Never returns nil.
//
// The iterator always returns until a break, return statement is reached, or the value
// given by the iterator is set to false.
func (b Bool) Each() iter.Seq[*Bool] {
	fn := func(yield func(*Bool) bool) {
		other_b := NewBool().WithValue(b.value)

		for other_b.value {
			if !yield(&other_b) {
				return
			}
		}
	}

	return fn
}
