package types

import (
	"iter"

	"github.com/PlayerR9/GoSD/pkg"
)

// Slice is a slice.
type Slice[T pkg.Type] struct {
	// values is the slice values.
	values []T
}

// Clean implements the pkg.Type interface.
func (s *Slice[T]) Clean() {
	if s == nil {
		return
	}

	s.values = pkg.CleanSlice(s.values)
	s.values = nil
}

// Equals implements the pkg.Type interface.
//
// Two sets are equal if they have the same values.
func (s *Slice[T]) Equals(other pkg.Type) bool {
	if other == nil {
		panic(pkg.NewNilComparison("other"))
	}

	other_val, ok := other.(*Slice[T])
	if !ok {
		return false
	}

	if len(s.values) != len(other_val.values) {
		return false
	}

	for i := 0; i < len(s.values); i++ {
		if !s.values[i].Equals(other_val.values[i]) {
			return false
		}
	}

	return true
}

// NewSlice creates a new empty slice.
//
// Returns:
//   - Slice: The new slice.
func NewSlice[T pkg.Type]() Slice[T] {
	return Slice[T]{
		values: make([]T, 0),
	}
}

// WithValue initializes a new slice with the given slice.
//
// Parameters:
//   - slice: The slice.
//
// Returns:
//   - Slice: The new slice.
func (s Slice[T]) WithValue(slice []T) Slice[T] {
	return Slice[T]{
		values: slice,
	}
}

// IsEmpty checks whether the slice is empty.
//
// Returns:
//   - bool: True if the slice is empty, false otherwise.
func (s Slice[T]) IsEmpty() bool {
	return len(s.values) == 0
}

// Size returns the number of elements in the slice.
//
// Returns:
//   - int: The number of elements in the slice.
func (s Slice[T]) Size() int {
	return len(s.values)
}

// Append adds an element to the slice.
//
// Parameters:
//   - elem: The element to add.
func (s *Slice[T]) Append(elem T) {
	s.values = append(s.values, elem)
}

// Reset removes all elements from the slice.
func (s *Slice[T]) Reset() {
	if s == nil {
		return
	}

	for i := 0; i < len(s.values); i++ {
		s.values[i] = *new(T)
	}
	s.values = s.values[:0]
}

// Each returns an iterator that iterates over all elements in the slice.
//
// Returns:
//   - iter.Seq[T]: The iterator. Never returns nil.
func (s Slice[T]) Each() iter.Seq[T] {
	fn := func(yield func(T) bool) {
		for _, elem := range s.values {
			if !yield(elem) {
				return
			}
		}
	}

	return fn
}

// ElemAt returns the element at the given index.
//
// Parameters:
//   - i: The index.
//
// Returns:
//   - T: The element.
//
// Panic with the message "index out of range" if the index is out of range.
func (s Slice[T]) ElemAt(i *Int) T {
	if i == nil {
		panic(pkg.NewInvalidCall("i", pkg.NewNilValue()))
	}
	val := i.Value()

	if val < 0 || val >= len(s.values) {
		panic("index out of range")
	}

	return s.values[val]
}

// Copy creates a copy of the slice.
//
// Returns:
//   - Slice[T]: The copy.
func (s Slice[T]) Copy() Slice[T] {
	values := make([]T, len(s.values))
	copy(values, s.values)

	return Slice[T]{
		values: values,
	}
}
