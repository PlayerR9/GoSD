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

	// InvalidState happens when a method is called on an object that is in an invalid state.
	//
	// Example:
	// 	var obj *Slice[*Int]
	// 	obj.Append(NewInt().WithValue(42))
	InvalidState

	// IllegalArgument happens when a method is called with an illegal argument.
	//
	// Example:
	// 	foo.Call(42)
	IllegalArgument
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

// NewInvalidState creates a new error with the InvalidState error code.
//
// Parameters:
//   - state: The state that is invalid.
//   - msg: The reason for the error.
//
// Returns:
//   - *Err: The new error. Never returns nil.
func NewInvalidState(state string, msg error) *Err {
	return &Err{
		Code: InvalidState,
		Msg:  fmt.Errorf("state (%s) is invalid: %w", state, msg),
	}
}

// NewIllegalArgument creates a new error with the IllegalArgument error code.
//
// Parameters:
//   - msg: The reason for the error.
//
// Returns:
//   - *Err: The new error. Never returns nil.
func NewIllegalArgument(msg error) *Err {
	return &Err{
		Code: IllegalArgument,
		Msg:  msg,
	}
}
