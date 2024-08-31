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
	// Each implementation must describe the behavior of the equals function.
	Equals(other Type) bool
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
//   - T: the result of the function.
type DoFunc[T Type] func() T
