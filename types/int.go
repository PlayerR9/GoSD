package types

import "github.com/PlayerR9/GoSD/pkg"

// Int is an index.
type Int struct {
	// value is the index value.
	value int
}

// Clean implements the pkg.Type interface.
func (idx *Int) Clean() {}

// Equals implements the pkg.Type interface.
//
// Two indexes are equal if they have the same value; regardless of the max value.
func (idx *Int) Equals(other pkg.Type) bool {
	if other == nil {
		panic(pkg.NewNilComparison("other"))
	}

	other_val, ok := other.(*Int)
	if !ok {
		return false
	}

	return idx.value == other_val.value
}

// NewIndex creates a new index.
//
// Returns:
//   - Int: The new index.
func NewInt() Int {
	return Int{
		value: 0,
	}
}

// WithValue creates a new index with the given value.
//
// Parameters:
//   - value: The value.
//
// Returns:
//   - Int: The new index.
func (idx Int) WithValue(value int) Int {
	return Int{
		value: value,
	}
}

// Value returns the index value.
//
// Returns:
//   - int: The index value.
func (idx Int) Value() int {
	return idx.value
}

// Set sets the index value.
//
// Parameters:
//   - value: The new value.
func (idx *Int) Set(value int) {
	idx.value = value
}
