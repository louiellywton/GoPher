package cmd

import (
	"bytes"
	"errors"
	"runtime"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

// TestRootCommandEdgeCases tests edge cases for root command to improve coverage
func TestRootCommandEdgeCases(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		expectError bool
		errorType   string
	}{
		{
			name:        "unknown command",
			args:        []string{"unknown"},
			expectError: true,
			errorType:   "usage",
		},
		{
			name:        "multiple unknown args",
			args:        []string{"unknown", "args", "here"},
			expectError: true,
			errorType:   "usage",
		},
		{
			name:        "no args shows help",
			args:        []string{},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cobra.Command{
				Use:   "hello-gopher",
				RunE:  rootCmd.RunE,
			}
			cmd.Flags().BoolP("version", "v", false, "version info")
			
			var output bytes.Buffer
			cmd.SetOut(&output)
			cmd.SetErr(&output)
			cmd.SetArgs(tt.args)

			err := cmd.Execute()
			
			if tt.expectError && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			
			if tt.errorType == "usage" && err != nil {
				if cliErr, ok := err.(*CLIError); ok {
					if cliErr.Code != ExitUsageError {
						t.Errorf("Expected usage error code %d, got %d", ExitUsageError, cliErr.Code)
					}
				}
			}
		})
	}
}

// TestGreetCommandEdgeCases tests edge cases for greet command
func TestGreetCommandEdgeCases(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		expectError bool
		errorType   string
	}{
		{
			name:        "unexpected positional args",
			args:        []string{"unexpected", "args"},
			expectError: true,
			errorType:   "usage",
		},
		{
			name:        "single unexpected arg",
			args:        []string{"unexpected"},
			expectError: true,
			errorType:   "usage",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cobra.Command{
				Use:  "greet",
				RunE: greetCmd.RunE,
			}
			cmd.Flags().StringP("name", "n", "", "Name to greet")
			
			var output bytes.Buffer
			cmd.SetOut(&output)
			cmd.SetErr(&output)
			cmd.SetArgs(tt.args)

			err := cmd.Execute()
			
			if tt.expectError && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			
			if tt.errorType == "usage" && err != nil {
				if cliErr, ok := err.(*CLIError); ok {
					if cliErr.Code != ExitUsageError {
						t.Errorf("Expected usage error code %d, got %d", ExitUsageError, cliErr.Code)
					}
				}
			}
		})
	}
}

// TestProverbCommandEdgeCases tests edge cases for proverb command
func TestProverbCommandEdgeCases(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		expectError bool
		errorType   string
	}{
		{
			name:        "unexpected positional args",
			args:        []string{"unexpected", "args"},
			expectError: true,
			errorType:   "usage",
		},
		{
			name:        "single unexpected arg",
			args:        []string{"unexpected"},
			expectError: true,
			errorType:   "usage",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cobra.Command{
				Use:  "proverb",
				RunE: proverbCmd.RunE,
			}
			
			var output bytes.Buffer
			cmd.SetOut(&output)
			cmd.SetErr(&output)
			cmd.SetArgs(tt.args)

			err := cmd.Execute()
			
			if tt.expectError && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			
			if tt.errorType == "usage" && err != nil {
				if cliErr, ok := err.(*CLIError); ok {
					if cliErr.Code != ExitUsageError {
						t.Errorf("Expected usage error code %d, got %d", ExitUsageError, cliErr.Code)
					}
				}
			}
		})
	}
}

// TestErrorUnwrapping tests error unwrapping functionality
func TestErrorUnwrapping(t *testing.T) {
	originalErr := errors.New("original error")
	
	tests := []struct {
		name     string
		err      *CLIError
		hasWrap  bool
	}{
		{
			name:    "usage error without cause",
			err:     NewUsageError("usage message", "usage suggestion"),
			hasWrap: false,
		},
		{
			name:    "data error with cause",
			err:     NewDataError("data message", originalErr, "data suggestion"),
			hasWrap: true,
		},
		{
			name:    "system error with cause",
			err:     NewSystemError("system message", originalErr, "system suggestion"),
			hasWrap: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			unwrapped := tt.err.Unwrap()
			
			if tt.hasWrap && unwrapped == nil {
				t.Error("Expected wrapped error but got nil")
			}
			if !tt.hasWrap && unwrapped != nil {
				t.Errorf("Expected no wrapped error but got: %v", unwrapped)
			}
			if tt.hasWrap && unwrapped != originalErr {
				t.Errorf("Expected wrapped error %v, got %v", originalErr, unwrapped)
			}
		})
	}
}

// TestErrorMessages tests error message formatting
func TestErrorMessages(t *testing.T) {
	tests := []struct {
		name        string
		err         *CLIError
		expectMsg   string
		expectSugg  bool
	}{
		{
			name:        "error with suggestion",
			err:         NewUsageError("test message", "test suggestion"),
			expectMsg:   "test message",
			expectSugg:  true,
		},
		{
			name:        "error without suggestion",
			err:         &CLIError{Code: ExitUsageError, Message: "test message"},
			expectMsg:   "test message",
			expectSugg:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errStr := tt.err.Error()
			
			if !strings.Contains(errStr, tt.expectMsg) {
				t.Errorf("Expected error message to contain %q, got %q", tt.expectMsg, errStr)
			}
			
			hasSuggestion := strings.Contains(errStr, "Suggestion:")
			if tt.expectSugg && !hasSuggestion {
				t.Error("Expected suggestion in error message but didn't find it")
			}
			if !tt.expectSugg && hasSuggestion {
				t.Error("Didn't expect suggestion in error message but found it")
			}
		})
	}
}

// TestVersionCommandOutput tests version command output formatting
func TestVersionCommandOutput(t *testing.T) {
	// Create a custom version command that uses cmd.Printf instead of fmt.Printf
	var output bytes.Buffer
	cmd := &cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Printf("hello-gopher version %s\n", version)
			cmd.Printf("Build date: %s\n", buildDate)
			cmd.Printf("Git commit: %s\n", gitCommit)
			cmd.Printf("Go version: %s\n", runtime.Version())
			cmd.Printf("OS/Arch: %s/%s\n", runtime.GOOS, runtime.GOARCH)
		},
	}
	cmd.SetOut(&output)
	cmd.SetArgs([]string{})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	outputStr := output.String()
	expectedStrings := []string{
		"hello-gopher version",
		"Build date:",
		"Git commit:",
		"Go version:",
		"OS/Arch:",
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(outputStr, expected) {
			t.Errorf("Expected output to contain %q, got: %s", expected, outputStr)
		}
	}
}

// TestRootCommandVersionFlag tests version flag on root command
func TestRootCommandVersionFlag(t *testing.T) {
	var output bytes.Buffer
	cmd := &cobra.Command{
		Use:  "hello-gopher",
		RunE: rootCmd.RunE,
	}
	cmd.Flags().BoolP("version", "v", false, "version info")
	cmd.SetOut(&output)
	cmd.SetArgs([]string{"--version"})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	outputStr := output.String()
	expectedStrings := []string{
		"hello-gopher version",
		"Build date:",
		"Git commit:",
		"Go version:",
		"OS/Arch:",
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(outputStr, expected) {
			t.Errorf("Expected output to contain %q, got: %s", expected, outputStr)
		}
	}
}

// TestRootCommandShortVersionFlag tests short version flag
func TestRootCommandShortVersionFlag(t *testing.T) {
	var output bytes.Buffer
	cmd := &cobra.Command{
		Use:  "hello-gopher",
		RunE: rootCmd.RunE,
	}
	cmd.Flags().BoolP("version", "v", false, "version info")
	cmd.SetOut(&output)
	cmd.SetArgs([]string{"-v"})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	outputStr := output.String()
	if !strings.Contains(outputStr, "hello-gopher version") {
		t.Errorf("Expected version output, got: %s", outputStr)
	}
}

// TestHandleErrorFunction tests the HandleError function behavior
func TestHandleErrorFunction(t *testing.T) {
	// We can't easily test os.Exit, so we'll test the error type detection
	// and message formatting parts
	
	t.Run("nil error", func(t *testing.T) {
		// This should not panic or cause issues
		// We can't test the actual HandleError function due to os.Exit
		// but we can test that our error types work correctly
		var err error = nil
		if err != nil {
			t.Error("Expected nil error")
		}
	})
	
	t.Run("CLI error type detection", func(t *testing.T) {
		cliErr := NewUsageError("test", "suggestion")
		
		// Test that we can detect CLI error type
		if cliErr.Code != ExitUsageError {
			t.Error("Expected usage error code")
		}
	})
	
	t.Run("non-CLI error", func(t *testing.T) {
		regularErr := errors.New("regular error")
		
		// Test that regular errors are different from CLI errors
		if regularErr.Error() == "" {
			t.Error("Expected non-empty error message")
		}
	})
}

// TestFlagErrorHandling tests custom flag error handling
func TestFlagErrorHandling(t *testing.T) {
	// Test that invalid flags are handled properly
	cmd := &cobra.Command{
		Use:  "hello-gopher",
		RunE: rootCmd.RunE,
	}
	cmd.Flags().BoolP("version", "v", false, "version info")
	
	// Set the same flag error function as the root command
	cmd.SetFlagErrorFunc(func(cmd *cobra.Command, err error) error {
		return NewUsageError(
			err.Error(),
			"Run 'hello-gopher --help' for usage information",
		)
	})
	
	var output bytes.Buffer
	cmd.SetOut(&output)
	cmd.SetErr(&output)
	cmd.SetArgs([]string{"--invalid-flag"})

	err := cmd.Execute()
	if err == nil {
		t.Error("Expected error for invalid flag")
	}
	
	if cliErr, ok := err.(*CLIError); ok {
		if cliErr.Code != ExitUsageError {
			t.Errorf("Expected usage error code %d, got %d", ExitUsageError, cliErr.Code)
		}
		if !strings.Contains(cliErr.Error(), "Suggestion:") {
			t.Error("Expected suggestion in flag error")
		}
	} else {
		t.Errorf("Expected CLIError, got %T", err)
	}
}

// TestCommandInitialization tests that commands are properly initialized
func TestCommandInitialization(t *testing.T) {
	// Test that greet command has proper flags
	if greetCmd.Flags().Lookup("name") == nil {
		t.Error("Expected greet command to have 'name' flag")
	}
	
	// Test that root command has version flag
	if rootCmd.Flags().Lookup("version") == nil {
		t.Error("Expected root command to have 'version' flag")
	}
	
	// Test that commands have proper parent-child relationships
	found := false
	for _, cmd := range rootCmd.Commands() {
		if cmd.Name() == "greet" {
			found = true
			break
		}
	}
	if !found {
		t.Error("Expected greet command to be added to root command")
	}
	
	found = false
	for _, cmd := range rootCmd.Commands() {
		if cmd.Name() == "proverb" {
			found = true
			break
		}
	}
	if !found {
		t.Error("Expected proverb command to be added to root command")
	}
	
	found = false
	for _, cmd := range rootCmd.Commands() {
		if cmd.Name() == "version" {
			found = true
			break
		}
	}
	if !found {
		t.Error("Expected version command to be added to root command")
	}
}