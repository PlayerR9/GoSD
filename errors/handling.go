package errors

import (
	assert "github.com/PlayerR9/GoSD/assert"
	pkg "github.com/PlayerR9/GoSD/pkg"
)

// ErrOf calls the given DoFunc and returns the result and error.
//
// Parameters:
//   - do: The DoFunc to call.
//
// Returns:
//   - T: The result of the DoFunc.
//   - error: The error that occurred.
func ErrOf[T pkg.Type](do pkg.DoFunc[T]) (T, error) {
	assert.AssertNotNil(do, "do")

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

// ErrHandler is a function that handles an error.
//
// Parameters:
//   - err: The error that occurred. Assume this is not nil.
//
// Returns:
//   - T: The result of the function.
type ErrHandler[T pkg.Type] func(err error) T

// Try calls the given DoFunc and executes the given exec function if an error occurs.
//
// Parameters:
//   - do: The DoFunc to call.
//   - exec: The function to execute if an error occurs. If nil, the error is rethrown.
//
// Returns:
//   - T: The result of the DoFunc.
func Try[T pkg.Type](do pkg.DoFunc[T], exec ErrHandler[T]) T {
	res, err := ErrOf(do)
	if err == nil {
		return res
	}

	if exec == nil {
		panic(err)
	}

	return exec(err)
}
