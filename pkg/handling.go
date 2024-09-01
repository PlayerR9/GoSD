package pkg

import (
	"errors"
	"fmt"
)

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

// IsPanic checks if the error is an ErrPanic.
//
// Parameters:
//   - err: The error to check.
//
// Returns:
//   - any: The value that caused the error.
//   - bool: True if the error is an ErrPanic, false otherwise.
func IsPanic(err error) (any, bool) {
	if err == nil {
		return nil, false
	}

	var panic_err *ErrPanic

	ok := errors.As(err, &panic_err)
	if ok {
		return panic_err.Value, true
	}

	return nil, false
}

// ErrOf calls the given DoFunc and returns the result and error.
//
// Parameters:
//   - do: The DoFunc to call.
//
// Returns:
//   - T: The result of the DoFunc.
//   - error: The error that occurred.
func ErrOf[O Type](do DoFunc[O]) (O, error) {
	ThrowIf(do == nil, NewInvalidCall("do", NewNilValue()))

	var reason error

	defer func() {
		r := recover()

		if r == nil {
			return
		}

		err, ok := r.(error)
		if !ok {
			err = NewErrPanic(r)
		}

		reason = err
	}()

	res := do()

	return res, reason
}

// ErrWithArgOf calls the given DoFunc and returns the result and error.
//
// Parameters:
//   - arg: The argument to the DoFunc.
//   - do: The DoFunc to call.
//
// Returns:
//   - T: The result of the DoFunc.
//   - error: The error that occurred.
func ErrWithArgOf[I any, O Type](arg I, do DoWithArgFunc[I, O]) (O, error) {
	ThrowIf(do == nil, NewInvalidCall("do", NewNilValue()))

	var reason error

	defer func() {
		r := recover()

		if r == nil {
			return
		}

		err, ok := r.(error)
		if !ok {
			err = NewErrPanic(r)
		}

		reason = err
	}()

	res := do(arg)

	return res, reason
}

// ErrHandler is a function that handles an error.
//
// Parameters:
//   - res: The result of the function.
//   - err: The error that occurred. Assume this is not nil.
//
// Returns:
//   - T: The result of the function.
type ErrHandler[T Type] func(res T, err error) T

// Try calls the given DoFunc and executes the given exec function if an error occurs.
//
// Parameters:
//   - arg: The argument to pass to the DoFunc.
//   - do: The DoFunc to call.
//   - exec: The function to execute if an error occurs. If nil, the error is rethrown.
//
// Returns:
//   - O	: The result of the DoFunc.
func Try[O Type](do DoFunc[O], exec ErrHandler[O]) O {
	res, err := ErrOf(do)
	if err == nil {
		return res
	}

	val, ok := IsPanic(err)
	if ok {
		panic(val)
	} else if exec == nil {
		panic(err)
	}

	return exec(res, err)
}

// TryWithArg calls the given DoFunc and executes the given exec function if an error occurs.
//
// Parameters:
//   - arg: The argument to pass to the DoFunc.
//   - do: The DoFunc to call.
//   - exec: The function to execute if an error occurs. If nil, the error is rethrown.
//
// Returns:
//   - O	: The result of the DoFunc.
func TryWithArg[I any, O Type](arg I, do DoWithArgFunc[I, O], exec ErrHandler[O]) O {
	res, err := ErrWithArgOf(arg, do)
	if err == nil {
		return res
	}

	val, ok := IsPanic(err)
	if ok {
		panic(val)
	} else if exec == nil {
		panic(err)
	}

	return exec(res, err)
}
