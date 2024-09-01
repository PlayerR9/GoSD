package pkg

import (
	"errors"
	"fmt"
)

//go:generate stringer -type=ErrorCode

type ErrorCode int

const (
	// NilComparison happens when a nil value is used within a comparison.
	//
	// Example:
	// 	foo.Equals(nil)
	NilComparison ErrorCode = iota

	// InvalidCall happens whenever a function is called with an invalid argument.
	//
	// Example:
	// 	foo.Call()
	InvalidCall

	// NilValue happens when a nil value is used but it was expected a non-nil value.
	//
	// Example:
	// 	foo.Value(nil)
	NilValue
)

// NewNilComparison creates a new error with the NilComparison error code.
//
// Parameters:
//   - de_name: The name of the data entity that is nil.
//
// Returns:
//   - *Err: The new error. Never returns nil.
func NewNilComparison(de_name string) *Err {
	return &Err{
		Code:        NilComparison,
		Msg:         errors.New("nil values cannot be compared"),
		Suggestions: []string{fmt.Sprintf("ensure that data entity (%s) is not nil", de_name)},
	}
}

// NewInvalidCall creates a new error with the InvalidCall error code.
//
// Parameters:
//   - de_name: The name of the data entity that is invalid.
//   - reason: The reason for the error.
//
// Returns:
//   - *Err: The new error. Never returns nil.
func NewInvalidCall(de_name string, reason error) *Err {
	return &Err{
		Code:        InvalidCall,
		Msg:         reason,
		Suggestions: []string{fmt.Sprintf("ensure that data entity (%s) is valid", de_name)},
	}
}

// NewNilValue creates a new error with the NilValue error code.
//
// Returns:
//   - *Err: The new error. Never returns nil.
func NewNilValue() *Err {
	return &Err{
		Code: NilValue,
		Msg:  errors.New("value expected to be non-nil"),
	}
}
