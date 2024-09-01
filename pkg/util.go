package pkg

// Throw panic with the given error if it is not nil.
//
// Parameters:
//   - err: The error to throw.
func Throw(err error) {
	if err != nil {
		panic(err)
	}
}

// ThrowIf panics with the given error if the given condition is true.
//
// Parameters:
//   - cond: The condition to check.
//   - err: The error to throw.
func ThrowIf(cond bool, err error) {
	if cond && err != nil {
		panic(err)
	}
}

// Error returns the error message of an error.
//
// Parameters:
//   - err: The error to get the message of.
//
// Returns:
//   - string: The error message of the error.
//
// If the error is nil, the function returns "something went wrong" as the error message.
func Error(err error) string {
	if err == nil {
		return "something went wrong"
	}

	return err.Error()
}

// Contains checks if the given slice contains the given element.
//
// Parameters:
//   - elems: The slice to check.
//   - elem: The element to check.
//
// Returns:
//   - bool: True if the slice contains the element, false otherwise.
func Contains[T Type](elems []T, elem T) bool {
	if len(elems) == 0 {
		return false
	}

	for i := 0; i < len(elems); i++ {
		if elems[i].Equals(elem) {
			return true
		}
	}

	return false
}

// IndexOf gets the index of the given element in the given slice.
//
// Parameters:
//   - elems: The slice to check.
//   - elem: The element to check.
//
// Returns:
//   - int: The index of the element in the slice, or -1 if the element is not found.
func IndexOf[T Type](elems []T, elem T) int {
	if len(elems) == 0 {
		return -1
	}

	for i := 0; i < len(elems); i++ {
		if elems[i].Equals(elem) {
			return i
		}
	}

	return -1
}
