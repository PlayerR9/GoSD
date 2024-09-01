package slices

import (
	"errors"
	"fmt"
	"iter"
	"strconv"
	"strings"

	"github.com/PlayerR9/GoSD/pkg"
)

// Index is an index.
type Index[T pkg.Type] struct {
	// value is the index value.
	value int

	// max is the index max value. -1 indicates that the max value is the
	// size of the reference.
	max int

	// ref is the slice reference.
	ref *Slice[T]
}

// String implements the fmt.Stringer interface.
func (idx *Index[T]) String() string {
	var builder strings.Builder

	builder.WriteString("Index[value=")
	builder.WriteString(strconv.Itoa(idx.value))
	builder.WriteString(", max=")

	if idx.max == -1 {
		builder.WriteString("+Inf")
	} else {
		builder.WriteString(strconv.Itoa(idx.max))
	}

	builder.WriteString(", ref=")
	fmt.Fprintf(&builder, "%p", idx.ref)
	builder.WriteString("]")

	return builder.String()
}

// DeepCopy implements the pkg.Type interface.
func (idx *Index[T]) DeepCopy() pkg.Type {
	return &Index[T]{
		value: idx.value,
		ref:   idx.ref,
		max:   idx.max,
	}
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
		return idx.value == other.value && idx.ref == other.ref
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
		max:   -1,
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

	max := pkg.OrElse(idx.max == -1, idx.ref.Size(), idx.max)

	pkg.ThrowIf(value < 0 || value >= max, pkg.NewIllegalArgument(errors.New("value must be less than slice size")))

	idx.value = value

	return idx
}

// WithMax creates a new index with the given max value.
//
// Parameters:
//   - max: The max value.
//
// Returns:
//   - *Index: The new index. Never returns nil.
//
// Panics if the slice is nil.
func (idx *Index[T]) WithMax(max int) *Index[T] {
	pkg.Ensure(false, idx)

	pkg.ThrowIf(max < 0 || max > idx.ref.Size(), pkg.NewIllegalArgument(errors.New("max must be less than slice size")))

	idx.max = max

	return idx
}

// WithoutMax creates a new index without the max value.
//
// Returns:
//   - *Index: The new index. Never returns nil.
//
// Panics if the slice is nil.
func (idx *Index[T]) WithoutMax() *Index[T] {
	pkg.Ensure(false, idx)

	idx.max = -1

	return idx
}

// Value returns the index value.
//
// Returns:
//   - int: The index value.
func (idx Index[T]) Value() int {
	return idx.value
}

// Max returns the index max value.
//
// Returns:
//   - int: The index max value.
func (idx Index[T]) Max() int {
	return pkg.OrElse(idx.max == -1, idx.ref.Size(), idx.max)
}

// Set sets the index value.
//
// Parameters:
//   - value: The new value.
func (idx *Index[T]) Set(value int) {
	pkg.Ensure(false, idx)

	max := pkg.OrElse(idx.max == -1, idx.ref.Size(), idx.max)

	pkg.ThrowIf(value < 0 || value >= max, pkg.NewIllegalArgument(errors.New("value must be less than slice size")))

	idx.value = value
}

// Each creates an iterator that iterates over the index from the current value up to the maximum value.
//
// Returns:
//   - iter.Seq[*Index]: The iterator. Never returns nil.
func (idx Index[T]) Each() iter.Seq[*Index[T]] {
	pkg.Ensure(false, &idx)

	var fn func(yield func(*Index[T]) bool)

	if idx.max == -1 {
		fn = func(yield func(*Index[T]) bool) {
			for idx.value < idx.ref.Size() {
				if !yield(&idx) {
					return
				}

				idx.value++
			}
		}
	} else {
		fn = func(yield func(*Index[T]) bool) {
			for idx.value < idx.max {
				if !yield(&idx) {
					return
				}

				idx.value++
			}
		}
	}

	return fn
}
