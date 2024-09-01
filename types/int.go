package types

import (
	"strconv"

	"github.com/PlayerR9/GoSD/pkg"
)

// Int is an index.
type Int struct {
	// value is the index value.
	value int
}

// String implements the fmt.Stringer interface.
func (idx *Int) String() string {
	return strconv.Itoa(idx.value)
}

// DeepCopy implements the pkg.Type interface.
func (idx *Int) DeepCopy() pkg.Type {
	return &Int{
		value: idx.value,
	}
}

// Ensure implements the pkg.Type interface.
func (idx *Int) Ensure() {
	pkg.ThrowIf(idx == nil, pkg.NewInvalidState("idx", pkg.NewNilValue()))
}

// Clean implements the pkg.Type interface.
func (idx *Int) Clean() {}

// Equals implements the pkg.Type interface.
//
// Two indexes are equal if they have the same value; regardless of the max value.
func (idx *Int) Equals(other pkg.Type) bool {
	pkg.Ensure(false, idx)
	pkg.Ensure(false, other)

	switch other := other.(type) {
	case *Int:
		return idx.value == other.value
	default:
		return false
	}
}

// NewIndex creates a new index.
//
// Returns:
//   - *Int: The new index. Never returns nil.
func NewInt() *Int {
	return &Int{
		value: 0,
	}
}

// WithValue creates a new index with the given value.
//
// Parameters:
//   - value: The value.
//
// Returns:
//   - *Int: The new index. Never returns nil.
func (idx *Int) WithValue(value int) *Int {
	if idx == nil {
		return &Int{
			value: value,
		}
	}

	idx.value = value

	return idx
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
	pkg.Ensure(false, idx)

	idx.value = value
}
