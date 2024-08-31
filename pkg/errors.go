package pkg

import (
	"fmt"
	"strings"
)

// Err is an error.
type Err struct {
	// Code is the error code.
	Code ErrorCode

	// Msg is the error message.
	Msg error

	// Suggestions is the list of suggestions.
	Suggestions []string
}

// Error implements the error interface.
//
// Message: "{code}: {message}"
func (e Err) Error() string {
	return fmt.Sprintf("%s: %s", e.Code.String(), Error(e.Msg))
}

// Unwrap implements errors.Unwrap interface.
func (e Err) Unwrap() error {
	return e.Msg
}

// NewErr creates a new error.
//
// Parameters:
//   - code: The error code.
//   - message: The error message.
//
// Returns:
//   - *Err: The new error. Never returns nil.
func NewErr(code ErrorCode, message error) *Err {
	return &Err{
		Code:        code,
		Msg:         message,
		Suggestions: nil,
	}
}

// AddSuggestion adds a suggestion of the error.
//
// Parameters:
//   - suggestions: The suggestion of the error.
//
// Returns:
//   - *Err: The error. Never returns nil.
//
// Each element in the suggestion is separated by a space but each call to this function
// adds each suggestion on a new line.
func (e *Err) AddSuggestion(suggestions ...string) *Err {
	e.Suggestions = append(e.Suggestions, strings.Join(suggestions, " "))

	return e
}

// ChangeReason changes the reason of the error.
//
// Parameters:
//   - reason: The new reason of the error.
func (e *Err) ChangeReason(reason error) {
	e.Msg = reason
}
