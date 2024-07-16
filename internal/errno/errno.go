// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache-2.0 License. This product includes software developed at
// Datadog (https://www.datadoghq.com/).
// Copyright 2024-Present Datadog, Inc.

// Package errno provides a way to handle errors with exit codes.
package errno

// Exit codes for errors.
const (
	// ExitOK is the exit code for a successful operation.
	ExitOK uint8 = iota

	// ExitError is the exit code for a generic error.
	ExitError

	// ExitUsage is the exit code for an incorrect usage of a command.
	ExitUsage

	// ExitIO is the exit code for an I/O error.
	ExitIO

	// ExitNotFound is the exit code for a file not found error.
	ExitNotFound

	// ExitPermission is the exit code for an I/O permission denied error.
	ExitPermission

	// ExitAPIError is the exit code for errors related to API requests.
	ExitAPIError

	// ExitUnauthorized is the exit code for unauthorized requests to the OpenAI
	// API.
	ExitUnauthorized

	// ExitTimeout is the exit code for timeout errors.
	ExitTimeout

	// ExitInvalidInput is the exit code for invalid input errors.
	ExitInvalidInput
)

// Error is an error with an exit code.
type Error struct {
	// Err is the underlying error that caused this error.
	Err error

	// No is the machine-readable exit code.
	No uint8
}

// Error implements the error interface for Error.
func (e *Error) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}

	return "unknown error"
}

// Unwrap returns the underlying error.
func (e *Error) Unwrap() error {
	return e.Err
}

// Code returns the exit code of the error as an integer.
func (e *Error) Code() int {
	return int(e.No)
}

// New creates a new Error with the given exit code and error.
func New(code uint8, err error) *Error {
	return &Error{
		No:  code,
		Err: err,
	}
}
