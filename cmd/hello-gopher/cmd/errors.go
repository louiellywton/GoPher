package cmd

import (
	"fmt"
	"os"
)

// Exit codes for different error scenarios
const (
	ExitSuccess    = 0
	ExitUsageError = 1
	ExitDataError  = 2
	ExitSystemError = 3
)

// CLIError represents a CLI-specific error with user guidance
type CLIError struct {
	Code       int
	Message    string
	Cause      error
	Suggestion string
}

// Error implements the error interface
func (e *CLIError) Error() string {
	if e.Suggestion != "" {
		return fmt.Sprintf("%s\nSuggestion: %s", e.Message, e.Suggestion)
	}
	return e.Message
}

// Unwrap returns the underlying cause error for error wrapping
func (e *CLIError) Unwrap() error {
	return e.Cause
}

// NewUsageError creates a new usage error with helpful suggestions
func NewUsageError(message string, suggestion string) *CLIError {
	return &CLIError{
		Code:       ExitUsageError,
		Message:    message,
		Suggestion: suggestion,
	}
}

// NewDataError creates a new data-related error
func NewDataError(message string, cause error, suggestion string) *CLIError {
	return &CLIError{
		Code:       ExitDataError,
		Message:    message,
		Cause:      cause,
		Suggestion: suggestion,
	}
}

// NewSystemError creates a new system-related error
func NewSystemError(message string, cause error, suggestion string) *CLIError {
	return &CLIError{
		Code:       ExitSystemError,
		Message:    message,
		Cause:      cause,
		Suggestion: suggestion,
	}
}

// HandleError processes CLI errors and exits with appropriate codes
func HandleError(err error) {
	if err == nil {
		return
	}

	if cliErr, ok := err.(*CLIError); ok {
		fmt.Fprintf(os.Stderr, "Error: %s\n", cliErr.Error())
		os.Exit(cliErr.Code)
	} else {
		// Handle non-CLI errors as generic system errors
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(ExitSystemError)
	}
}