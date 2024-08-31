package pkg

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
