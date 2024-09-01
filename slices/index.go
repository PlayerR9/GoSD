package slices

import (
	"errors"
	"iter"

	"github.com/PlayerR9/GoSD/pkg"
)

// Index is an index.
type Index[T pkg.Type] struct {
	// value is the index value.
	value int

	// ref is the slice reference.
	ref *Slice[T]
}

// Ensure implements the pkg.Type interface.
func (idx *Index[T]) Ensure() {
	pkg.ThrowIf(idx == nil, pkg.NewInvalidState("idx", pkg.NewNilValue()))
	pkg.ThrowIf(idx.ref == nil, pkg.NewInvalidState("idx.ref", pkg.NewNilValue()))
}

// Clean implements the pkg.Type interface.
func (idx *Index[T]) Clean() {
	if idx == nil {
		return
	}

	if idx.ref != nil {
		idx.ref = nil
	}
}

// Equals implements the pkg.Type interface.
//
// Two indexes are equal if they have the same value; regardless of the max value.
func (idx *Index[T]) Equals(other pkg.Type) bool {
	pkg.Ensure(false, idx)
	pkg.Ensure(false, other)

	switch other := other.(type) {
	case *Index[T]:
		return idx.value == other.value
	default:
		return false
	}
}

// NewIndex creates a new index.
//
// Parameters:
//   - slice: The slice reference.
//
// Returns:
//   - *Index: The new index. Never returns nil.
//
// Panics if the slice is nil.
func NewIndex[T pkg.Type](slice *Slice[T]) *Index[T] {
	pkg.Ensure(false, slice)

	return &Index[T]{
		value: 0,
		ref:   slice,
	}
}

// WithValue creates a new index with the given value.
//
// Parameters:
//   - value: The value.
//
// Returns:
//   - *Index: The new index. Never returns nil.
//
// Panics if the slice is nil.
func (idx *Index[T]) WithValue(value int) *Index[T] {
	pkg.Ensure(false, idx)

	pkg.ThrowIf(value < 0 || value >= idx.ref.Size(), pkg.NewIllegalArgument(errors.New("value must be less than slice size")))

	idx.value = value

	return idx
}

// Value returns the index value.
//
// Returns:
//   - int: The index value.
func (idx Index[T]) Value() int {
	return idx.value
}

// Set sets the index value.
//
// Parameters:
//   - value: The new value.
func (idx *Index[T]) Set(value int) {
	pkg.Ensure(false, idx)
	pkg.ThrowIf(value < 0 || value >= idx.ref.Size(), pkg.NewIllegalArgument(errors.New("value must be less than slice size")))

	idx.value = value
}

// Each creates an iterator that iterates over the index from 0 up to the maximum value.
//
// Returns:
//   - iter.Seq[*Index]: The iterator. Never returns nil.
func (idx Index[T]) Each() iter.Seq[*Index[T]] {
	pkg.Ensure(false, &idx)

	fn := func(yield func(*Index[T]) bool) {
		idx.value = 0

		for idx.value < idx.ref.Size() {
			if !yield(&idx) {
				return
			}

			idx.value++
		}
	}

	return fn
}
