package slices

import (
	"errors"
	"fmt"
	"iter"
	"strings"

	"github.com/PlayerR9/GoSD/pkg"
)

// Slice is a slice.
type Slice[T pkg.Type] struct {
	// values is the slice values.
	values []T
}

// String implements the fmt.Stringer interface.
func (s *Slice[T]) String() string {
	var builder strings.Builder

	builder.WriteString("Slice[")

	values := make([]string, 0, len(s.values))
	for i := 0; i < len(s.values); i++ {
		values = append(values, s.values[i].String())
	}
	builder.WriteString(strings.Join(values, ", "))

	builder.WriteString("]")

	return builder.String()
}

// DeepCopy implements the pkg.Type interface.
func (s *Slice[T]) DeepCopy() pkg.Type {
	if s == nil {
		return nil
	}

	slice := make([]T, 0, len(s.values))

	for _, v := range s.values {
		v_copy := v.DeepCopy()

		tmp, ok := v_copy.(T)
		pkg.ThrowIf(!ok, pkg.NewInvalidState("v_copy", errors.New("invalid type")))

		slice = append(slice, tmp)
	}

	return &Slice[T]{
		values: slice,
	}
}

// Ensure implements the pkg.Type interface.
func (s *Slice[T]) Ensure() {
	pkg.ThrowIf(s == nil, pkg.NewInvalidState("s", pkg.NewNilValue()))
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
	pkg.Ensure(false, s)
	pkg.Ensure(false, other)

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
//   - Slice: The new slice. Never returns nil.
func NewSlice[T pkg.Type]() *Slice[T] {
	return &Slice[T]{
		values: make([]T, 0),
	}
}

// WithValue initializes a new slice with the given slice.
//
// Parameters:
//   - slice: The slice.
//
// Returns:
//   - *Slice: The new slice. Never returns nil.
func (s *Slice[T]) WithValue(slice []T) *Slice[T] {
	if s == nil {
		return &Slice[T]{
			values: slice,
		}
	}

	s.values = pkg.CleanSlice(s.values)
	s.values = slice

	return s
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
//
// Panics if the slice is nil.
func (s *Slice[T]) Append(elem T) {
	pkg.Ensure(false, s)

	s.values = append(s.values, elem)
}

// Merge merges the given slice into the slice.
//
// Parameters:
//   - slice: The slice to merge.
//
// Panics if the slice is nil.
func (s *Slice[T]) Merge(slice *Slice[T]) {
	pkg.Ensure(false, s)

	if slice == nil {
		return
	}

	s.values = append(s.values, slice.values...)
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

// Index returns the index type of the slice.
//
// Returns:
//   - *Index[T]: The index type. Never returns nil.
func (s *Slice[T]) Index() *Index[T] {
	pkg.Ensure(false, s)

	return &Index[T]{
		value: 0,
		ref:   s,
	}
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
func (s *Slice[T]) ElemAt(i *Index[T]) T {
	pkg.Ensure(false, s)
	pkg.Ensure(false, i)

	pkg.ThrowIf(i.ref != s, fmt.Errorf("index refers to a different slice: %p", i.ref))

	return s.values[i.value]
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

// DeleteFirst deletes the first element in the slice.
//
// Returns:
//   - T: The deleted element.
func (s *Slice[T]) DeleteFirst() T {
	pkg.Ensure(false, s)

	pkg.ThrowIf(len(s.values) == 0, fmt.Errorf("slice is empty"))

	top := s.values[0]
	s.values = s.values[1:]

	return top
}

// SetAt sets the element at the given index.
//
// Parameters:
//   - i: The index.
//   - elem: The element.
//
// Panics with the message "index out of range" if the index is out of range.
func (s *Slice[T]) SetAt(i *Index[T], elem T) {
	pkg.Ensure(false, s)
	pkg.Ensure(false, i)

	pkg.ThrowIf(i.ref != s, fmt.Errorf("index refers to a different slice: %p", i.ref))

	s.values[i.value] = elem
}

// Has checks whether the slice contains the given element.
//
// Parameters:
//   - elem: The element.
//
// Returns:
//   - bool: True if the slice contains the element, false otherwise.
func (s *Slice[T]) Has(elem T) bool {
	pkg.Ensure(false, s)

	for i := 0; i < len(s.values); i++ {
		if s.values[i].Equals(elem) {
			return true
		}
	}

	return false
}
