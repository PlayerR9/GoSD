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
}

// Clean implements the pkg.Type interface.
func (idx *Index) Clean() {}

// Equals implements the pkg.Type interface.
//
// Two indexes are equal if they have the same value.
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
//   - value: The index value.
//
// Returns:
//   - Index: The new index.
//
// Panics if the value is negative.
func NewIndex(value int) Index {
	if value < 0 {
		panic(pkg.NewInvalidCall("value", errors.New("value must be non-negative")))
	}

	return Index{
		value: value,
	}
}

// NewIndexOfSlice creates a new index of the length of the slice.
//
// Parameters:
//   - slice: The slice.
//
// Returns:
//   - Index: The new index.
func NewIndexOfSlice[T any](slice []T) Index {
	return Index{
		value: len(slice),
	}
}

// Each creates an iterator that iterates over the index.
//
// Returns:
//   - iter.Seq[int]: The iterator. Never returns nil.
func (idx *Index) Each() iter.Seq[int] {
	fn := func(yield func(int) bool) {
		for i := 0; i < idx.value; i++ {
			if !yield(i) {
				return
			}
		}
	}

	return fn
}
