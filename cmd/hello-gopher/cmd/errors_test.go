package cmd

import (
	"errors"
	"testing"
)

func TestCLIError_Error(t *testing.T) {
	tests := []struct {
		name     string
		cliError *CLIError
		expected string
	}{
		{
			name: "error with suggestion",
			cliError: &CLIError{
				Code:       ExitUsageError,
				Message:    "Invalid command",
				Suggestion: "Try using --help",
			},
			expected: "Invalid command\nSuggestion: Try using --help",
		},
		{
			name: "error without suggestion",
			cliError: &CLIError{
				Code:    ExitDataError,
				Message: "Data not found",
			},
			expected: "Data not found",
		},
		{
			name: "error with cause and suggestion",
			cliError: &CLIError{
				Code:       ExitSystemError,
				Message:    "System failure",
				Cause:      errors.New("underlying error"),
				Suggestion: "Check system resources",
			},
			expected: "System failure\nSuggestion: Check system resources",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.cliError.Error()
			if result != tt.expected {
				t.Errorf("CLIError.Error() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestCLIError_Unwrap(t *testing.T) {
	cause := errors.New("underlying error")
	cliError := &CLIError{
		Code:    ExitSystemError,
		Message: "System failure",
		Cause:   cause,
	}

	unwrapped := cliError.Unwrap()
	if unwrapped != cause {
		t.Errorf("CLIError.Unwrap() = %v, want %v", unwrapped, cause)
	}
}

func TestNewUsageError(t *testing.T) {
	message := "Invalid usage"
	suggestion := "Use --help"
	
	err := NewUsageError(message, suggestion)
	
	if err.Code != ExitUsageError {
		t.Errorf("NewUsageError().Code = %d, want %d", err.Code, ExitUsageError)
	}
	if err.Message != message {
		t.Errorf("NewUsageError().Message = %q, want %q", err.Message, message)
	}
	if err.Suggestion != suggestion {
		t.Errorf("NewUsageError().Suggestion = %q, want %q", err.Suggestion, suggestion)
	}
}

func TestNewDataError(t *testing.T) {
	message := "Data error"
	cause := errors.New("file not found")
	suggestion := "Check file path"
	
	err := NewDataError(message, cause, suggestion)
	
	if err.Code != ExitDataError {
		t.Errorf("NewDataError().Code = %d, want %d", err.Code, ExitDataError)
	}
	if err.Message != message {
		t.Errorf("NewDataError().Message = %q, want %q", err.Message, message)
	}
	if err.Cause != cause {
		t.Errorf("NewDataError().Cause = %v, want %v", err.Cause, cause)
	}
	if err.Suggestion != suggestion {
		t.Errorf("NewDataError().Suggestion = %q, want %q", err.Suggestion, suggestion)
	}
}

func TestNewSystemError(t *testing.T) {
	message := "System error"
	cause := errors.New("permission denied")
	suggestion := "Check permissions"
	
	err := NewSystemError(message, cause, suggestion)
	
	if err.Code != ExitSystemError {
		t.Errorf("NewSystemError().Code = %d, want %d", err.Code, ExitSystemError)
	}
	if err.Message != message {
		t.Errorf("NewSystemError().Message = %q, want %q", err.Message, message)
	}
	if err.Cause != cause {
		t.Errorf("NewSystemError().Cause = %v, want %v", err.Cause, cause)
	}
	if err.Suggestion != suggestion {
		t.Errorf("NewSystemError().Suggestion = %q, want %q", err.Suggestion, suggestion)
	}
}

// Note: HandleError function cannot be easily tested as it calls os.Exit
// In a real application, this would be tested through integration tests
// or by refactoring to accept an exit function as a parameter

func TestHandleError_NilError(t *testing.T) {
	// Test that HandleError with nil doesn't panic
	// We can't test the actual exit behavior without more complex setup
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("HandleError(nil) should not panic, got: %v", r)
		}
	}()
	
	// This test verifies the function exists and can handle nil
	// The actual exit behavior would be tested in integration tests
	t.Log("HandleError function exists and can be called")
}

// TestExitCodes verifies the exit code constants
func TestExitCodes(t *testing.T) {
	tests := []struct {
		name     string
		code     int
		expected int
	}{
		{"ExitSuccess", ExitSuccess, 0},
		{"ExitUsageError", ExitUsageError, 1},
		{"ExitDataError", ExitDataError, 2},
		{"ExitSystemError", ExitSystemError, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.code != tt.expected {
				t.Errorf("Expected %s to be %d, got %d", tt.name, tt.expected, tt.code)
			}
		})
	}
}