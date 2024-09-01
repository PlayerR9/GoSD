package types

import (
	"errors"
	"iter"

	"github.com/PlayerR9/GoSD/pkg"
)

// Index is an index.
type Index struct {
	// value is the index value.
	value int

	// max_value is the maximum index value.
	max_value int
}

// Clean implements the pkg.Type interface.
func (idx *Index) Clean() {}

// Equals implements the pkg.Type interface.
//
// Two indexes are equal if they have the same value; regardless of the max value.
func (idx *Index) Equals(other pkg.Type) bool {
	if other == nil {
		panic(pkg.NewNilComparison("other"))
	}

	other_val, ok := other.(*Index)
	if !ok {
		return false
	}

	return idx.value == other_val.value
}

// NewIndex creates a new index.
//
// Parameters:
//   - max_value: The maximum index value.
//
// Returns:
//   - Index: The new index.
//
// Panics if max_value is negative.
func NewIndex(max_value int) Index {
	if max_value < 0 {
		panic(pkg.NewInvalidCall("max_value", errors.New("max_value must be non-negative")))
	}

	return Index{
		value:     0,
		max_value: max_value,
	}
}

// WithValue creates a new index with the given value.
//
// Parameters:
//   - value: The value.
//
// Returns:
//   - Index: The new index.
func (idx Index) WithValue(value int) Index {
	if value >= idx.max_value {
		panic(pkg.NewInvalidCall("value", errors.New("value must be less than max_value")))
	}

	return Index{
		value:     value,
		max_value: idx.max_value,
	}
}

// Value returns the index value.
//
// Returns:
//   - int: The index value.
func (idx Index) Value() int {
	return idx.value
}

// Set sets the index value.
//
// Parameters:
//   - value: The new value.
func (idx *Index) Set(value int) {
	if value >= idx.max_value {
		panic(pkg.NewInvalidCall("value", errors.New("value must be less than max_value")))
	}

	idx.value = value
}

// Each creates an iterator that iterates over the index from 0 up to the maximum value.
//
// Returns:
//   - iter.Seq[*Index]: The iterator. Never returns nil.
func (idx Index) Each() iter.Seq[*Index] {
	fn := func(yield func(*Index) bool) {
		other_idx := NewIndex(idx.max_value).WithValue(0)

		for other_idx.value < idx.max_value {
			if !yield(&other_idx) {
				return
			}

			other_idx.value++
		}
	}

	return fn
}
