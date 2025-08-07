package cmd

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestRootCommandErrorHandling(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		expectError bool
		errorType   string
		errorMsg    string
	}{
		{
			name:        "unknown command",
			args:        []string{"unknown"},
			expectError: true,
			errorType:   "usage",
			errorMsg:    "Unknown command",
		},
		{
			name:        "unknown flag",
			args:        []string{"--unknown-flag"},
			expectError: true,
			errorType:   "usage",
			errorMsg:    "unknown flag",
		},
		{
			name:        "invalid flag format",
			args:        []string{"--name"}, // missing value
			expectError: true,
			errorType:   "usage",
		},
		{
			name:        "valid version flag",
			args:        []string{"--version"},
			expectError: false,
		},
		{
			name:        "valid short version flag",
			args:        []string{"-v"},
			expectError: false,
		},
		{
			name:        "valid help flag",
			args:        []string{"--help"},
			expectError: false,
		},
		{
			name:        "no arguments (should show help)",
			args:        []string{},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a fresh root command for each test
			testRootCmd := &cobra.Command{
				Use:   "hello-gopher",
				Short: "A friendly CLI tool for Go enthusiasts",
				SilenceUsage:  true,
				SilenceErrors: true,
				RunE: rootCmd.RunE, // Use the same RunE function
			}
			
			// Add the same flags as the real root command
			testRootCmd.Flags().BoolP("version", "v", false, "version for hello-gopher")
			
			// Set the same error handlers
			testRootCmd.SetFlagErrorFunc(func(cmd *cobra.Command, err error) error {
				return NewUsageError(
					err.Error(),
					fmt.Sprintf("Run '%s --help' for usage information", cmd.CommandPath()),
				)
			})
			
			// Capture output
			var output bytes.Buffer
			testRootCmd.SetOut(&output)
			testRootCmd.SetErr(&output)
			testRootCmd.SetArgs(tt.args)

			// Execute command
			err := testRootCmd.Execute()
			
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
					return
				}
				
				// Check if it's the right type of error
				if cliErr, ok := err.(*CLIError); ok {
					switch tt.errorType {
					case "usage":
						if cliErr.Code != ExitUsageError {
							t.Errorf("Expected usage error (code %d), got code %d", ExitUsageError, cliErr.Code)
						}
					}
					
					// Check error message contains expected text
					if tt.errorMsg != "" && !strings.Contains(strings.ToLower(cliErr.Error()), strings.ToLower(tt.errorMsg)) {
						t.Errorf("Expected error message to contain %q, got %q", tt.errorMsg, cliErr.Error())
					}
					
					// Verify suggestion is provided
					if cliErr.Suggestion == "" {
						t.Error("Expected error to include a suggestion")
					}
				} else {
					t.Errorf("Expected CLIError, got %T: %v", err, err)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got: %v", err)
				}
			}
		})
	}
}

func TestRootCommandVersionOutput(t *testing.T) {
	// Use the actual root command for testing
	var output bytes.Buffer
	rootCmd.SetOut(&output)
	rootCmd.SetErr(&output)
	rootCmd.SetArgs([]string{"--version"})

	// Execute command
	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	result := output.String()
	t.Logf("Version output: %q", result)
	
	// Check that version output contains expected elements
	expectedElements := []string{
		"hello-gopher version",
		"Build date:",
		"Git commit:",
		"Go version:",
		"OS/Arch:",
	}

	for _, element := range expectedElements {
		if !strings.Contains(result, element) {
			t.Errorf("Version output missing expected element: %q\nActual output: %q", element, result)
		}
	}
}

func TestRootCommandHelpOutput(t *testing.T) {
	// Use the actual root command for testing
	var output bytes.Buffer
	rootCmd.SetOut(&output)
	rootCmd.SetErr(&output)
	rootCmd.SetArgs([]string{"--help"})

	// Execute command
	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	result := output.String()
	t.Logf("Help output: %q", result)
	
	// Check that help output contains expected elements
	expectedElements := []string{
		"Hello-Gopher is a friendly command-line tool",
		"Usage:",
		"Examples:",
		"hello-gopher greet",
		"hello-gopher proverb",
		"Flags:",
		"-v, --version",
	}

	for _, element := range expectedElements {
		if !strings.Contains(result, element) {
			t.Errorf("Help output missing expected element: %q\nActual output: %q", element, result)
		}
	}
}