package errors

import (
	"errors"
	"fmt"
)

var (
	// NilValue is the error returned when a pointer is nil. While readers are not expected to return this
	// error by itself, if it does, readers must not wrap it as callers will test this error using ==.
	NilValue error
)

func init() {
	NilValue = errors.New("pointer must not be nil")
}

// ErrPanic represents an error when a function panics.
type ErrPanic struct {
	// Value is the value that caused the error.
	Value any
}

// Error implements the error interface.
//
// Message: "panic: {value}"
func (e ErrPanic) Error() string {
	return fmt.Sprintf("panic: %v", e.Value)
}

// NewErrPanic creates a new ErrPanic error.
//
// Parameters:
//   - value: The value that caused the error.
//
// Returns:
//   - *ErrPanic: A pointer to the newly created ErrPanic. Never returns nil.
func NewErrPanic(value any) *ErrPanic {
	return &ErrPanic{
		Value: value,
	}
}
