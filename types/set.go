package types

import (
	"iter"

	"github.com/PlayerR9/GoSD/pkg"
)

// Set is a set.
type Set[T pkg.Type] struct {
	// values is the set values.
	values []T
}

// Ensure implements the pkg.Type interface.
func (s *Set[T]) Ensure() {
	pkg.ThrowIf(s == nil, pkg.NewInvalidState("s", pkg.NewNilValue()))
}

// Clean implements the pkg.Type interface.
func (s *Set[T]) Clean() {
	if s == nil {
		return
	}

	s.values = pkg.CleanSlice(s.values)
	s.values = nil
}

// Equals implements the pkg.Type interface.
//
// Two sets are equal if they have the same values.
func (s *Set[T]) Equals(other pkg.Type) bool {
	pkg.Ensure(false, s)
	pkg.Ensure(false, other)

	switch other := other.(type) {
	case *Set[T]:
		if len(s.values) != len(other.values) {
			return false
		}

		for i := 0; i < len(s.values); i++ {
			if !s.values[i].Equals(other.values[i]) {
				return false
			}
		}

		return true
	default:
		return false
	}
}

// NewSet creates a new empty set.
//
// Returns:
//   - *Set: The new set. Never returns nil.
func NewSet[T pkg.Type]() *Set[T] {
	return &Set[T]{
		values: make([]T, 0),
	}
}

// WithValue set the values of the set.
//
// Parameters:
//   - slice: The slice.
//
// Returns:
//   - *Set: The new set. Never returns nil.
func (s *Set[T]) WithValue(slice []T) *Set[T] {
	var unique []T

	for i := 0; i < len(slice); i++ {
		if !pkg.Contains(unique, slice[i]) {
			unique = append(unique, slice[i])
		}
	}

	if s == nil {
		return &Set[T]{
			values: unique,
		}
	}

	s.values = pkg.CleanSlice(s.values)
	s.values = unique

	return s
}

// IsEmpty checks whether the set is empty.
//
// Returns:
//   - bool: True if the set is empty, false otherwise.
func (s Set[T]) IsEmpty() bool {
	return len(s.values) == 0
}

// Size returns the number of elements in the set.
//
// Returns:
//   - int: The number of elements in the set.
func (s Set[T]) Size() int {
	return len(s.values)
}

// Add adds an element to the set. If the element is already in the set, this method does nothing.
//
// Parameters:
//   - elem: The element to add.
//
// Returns:
//   - bool: True if the element was added, false otherwise.
func (s *Set[T]) Add(elem T) bool {
	pkg.Ensure(false, s)

	has := pkg.Contains(s.values, elem)
	if !has {
		s.values = append(s.values, elem)
	}

	return !has
}

// Union adds all elements from another set to the set.
//
// Parameters:
//   - other: The other set to add.
//
// Returns:
//   - int: The number of elements added.
func (s *Set[T]) Union(other *Set[T]) int {
	pkg.Ensure(false, s)

	if other == nil {
		return 0
	}

	var count int

	for i := 0; i < len(other.values); i++ {
		ok := pkg.Contains(s.values, other.values[i])
		if !ok {
			s.values = append(s.values, other.values[i])
			count++
		}
	}

	return count
}

// Reset removes all elements from the set.
func (s *Set[T]) Reset() {
	if s == nil {
		return
	}

	for i := 0; i < len(s.values); i++ {
		s.values[i] = *new(T)
	}
	s.values = s.values[:0]
}

// Each returns an iterator that iterates over all elements in the set.
//
// Returns:
//   - iter.Seq[T]: The iterator. Never returns nil.
func (s Set[T]) Each() iter.Seq[T] {
	fn := func(yield func(T) bool) {
		for _, elem := range s.values {
			if !yield(elem) {
				return
			}
		}
	}

	return fn
}
