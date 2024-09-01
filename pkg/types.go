package pkg

// Type is an interface that describes the behaviors of a SD type.
type Type interface {
	// Clean cleans up the type.
	Clean()

	// Equals checks if two types are equal. Nil values must panic.
	//
	// Parameters:
	//   - other: The other type.
	//
	// Returns:
	//   - bool: True if the types are equal, false otherwise.
	//
	// Throws:
	//   - *NilComparison: If the receiver or other is nil.
	//   - any other error: Depend on the implementation.
	//
	// Each implementation must describe the behavior of the equals function.
	Equals(other Type) bool

	// Ensure ensures that the type's state is valid. If not, it panics.
	Ensure()
}

// Ensure ensures that the type's state is valid. If not, it panics.
//
// Parameters:
//   - allow_nil: Whether to allow nil values.
//   - type_: The type to ensure.
func Ensure(allow_nil bool, type_ Type) {
	if type_ == nil {
		if allow_nil {
			return
		}

		Throw(NewInvalidCall("type_", NewNilValue()))
	}

	type_.Ensure()
}

// Clean cleans up the type.
//
// Parameters:
//   - type_: The type to clean up.
func Clean(type_ Type) {
	if type_ == nil {
		return
	}

	type_.Clean()
}

// CleanSlice cleans up the slice. Remember to set to nil after use.
//
// Parameters:
//   - slice: The slice to clean up.
//
// Returns:
//   - []T: The cleaned slice.
func CleanSlice[T Type](slice []T) []T {
	for i := 0; i < len(slice); i++ {
		slice[i].Clean()
	}

	return slice[:0:0]
}

// Equals checks if two types are equal. Nil values must panic.
//
// Parameters:
//   - type_: The other type.
//
// Returns:
//   - bool: True if the types are equal, false otherwise.
func Equals(a, b Type) bool {
	if a == nil {
		panic(NewNilComparison("a"))
	} else if b == nil {
		panic(NewNilComparison("b"))
	}

	return a.Equals(b)
}

// DoFunc is a function that does something. It must panic when error
// happens; regardless of the error.
//
// Returns:
//   - O: the result of the function.
type DoFunc[O Type] func() O

// DoWithArgFunc is a function that does something. It must panic when error
// happens; regardless of the error.
//
// Parameters:
//   - arg: the input of the function.
//
// Returns:
//   - O: the result of the function.
type DoWithArgFunc[I any, O Type] func(arg I) O
