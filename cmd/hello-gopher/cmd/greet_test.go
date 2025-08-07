package cmd

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/louiellywton/go-portfolio/01-hello-gopher/pkg/greeting"
	"github.com/spf13/cobra"
)

func TestGreetCommand(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "default greeting",
			args:     []string{"greet"},
			expected: "Hello, Gopher!",
		},
		{
			name:     "greeting with long flag",
			args:     []string{"greet", "--name", "Alice"},
			expected: "Hello, Alice!",
		},
		{
			name:     "greeting with short flag",
			args:     []string{"greet", "-n", "Bob"},
			expected: "Hello, Bob!",
		},
		{
			name:     "greeting with empty name",
			args:     []string{"greet", "--name", ""},
			expected: "Hello, Gopher!",
		},
		{
			name:     "greeting with special characters",
			args:     []string{"greet", "--name", "José"},
			expected: "Hello, José!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new command instance for testing
			testGreetCmd := &cobra.Command{
				Use:   "greet",
				Short: "Greet a gopher by name",
				RunE: func(cmd *cobra.Command, args []string) error {
					name, err := cmd.Flags().GetString("name")
					if err != nil {
						return err
					}

					// Create greeting service and generate greeting
					service := greeting.NewService()
					greeting := service.Greet(name)
					
					// Write to command output instead of stdout
					cmd.Print(greeting)
					return nil
				},
			}
			testGreetCmd.Flags().StringP("name", "n", "", "Name to greet (default: Gopher)")
			
			// Capture output
			var output bytes.Buffer
			testGreetCmd.SetOut(&output)
			testGreetCmd.SetErr(&output)
			testGreetCmd.SetArgs(tt.args[1:]) // Remove "greet" from args

			// Execute command
			err := testGreetCmd.Execute()
			if err != nil {
				t.Fatalf("Command execution failed: %v", err)
			}

			// Check output
			result := strings.TrimSpace(output.String())
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestGreetCommandHelp(t *testing.T) {
	// Create a new root command
	cmd := &cobra.Command{
		Use: "hello-gopher",
	}
	cmd.AddCommand(greetCmd)
	
	// Capture output
	var output bytes.Buffer
	cmd.SetOut(&output)
	cmd.SetErr(&output)
	cmd.SetArgs([]string{"greet", "--help"})

	// Execute command
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Help command execution failed: %v", err)
	}

	// Check that help output contains expected elements
	result := output.String()
	expectedElements := []string{
		"Greet command provides friendly greeting functionality",
		"Usage:",
		"hello-gopher greet [flags]",
		"Examples:",
		"hello-gopher greet --name Alice",
		"Flags:",
		"-n, --name string",
	}

	for _, element := range expectedElements {
		if !strings.Contains(result, element) {
			t.Errorf("Help output missing expected element: %q", element)
		}
	}
}

func TestGreetCommandIntegration(t *testing.T) {
	// Test that the greet command is properly registered with the root command
	found := false
	for _, cmd := range rootCmd.Commands() {
		if cmd.Name() == "greet" {
			found = true
			break
		}
	}
	
	if !found {
		t.Error("greet command not found in root command")
	}
}

func TestGreetCommandErrorHandling(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		expectError bool
		errorType   string
	}{
		{
			name:        "unexpected positional argument",
			args:        []string{"greet", "unexpected"},
			expectError: true,
			errorType:   "usage",
		},
		{
			name:        "multiple unexpected arguments",
			args:        []string{"greet", "arg1", "arg2"},
			expectError: true,
			errorType:   "usage",
		},
		{
			name:        "valid command with flag",
			args:        []string{"greet", "--name", "Test"},
			expectError: false,
		},
		{
			name:        "valid command without flag",
			args:        []string{"greet"},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a fresh root command for each test to avoid state issues
			testRootCmd := &cobra.Command{
				Use:   "hello-gopher",
				Short: "A friendly CLI tool for Go enthusiasts",
				SilenceUsage:  true,
				SilenceErrors: true,
				RunE: rootCmd.RunE,
			}
			testRootCmd.Flags().BoolP("version", "v", false, "version for hello-gopher")
			testRootCmd.SetFlagErrorFunc(func(cmd *cobra.Command, err error) error {
				return NewUsageError(
					err.Error(),
					fmt.Sprintf("Run '%s --help' for usage information", cmd.CommandPath()),
				)
			})
			
			// Add a fresh greet command
			testGreetCmd := &cobra.Command{
				Use:   "greet",
				Short: "Greet a gopher by name",
				RunE:  greetCmd.RunE,
			}
			testGreetCmd.Flags().StringP("name", "n", "", "Name to greet (default: Gopher)")
			testRootCmd.AddCommand(testGreetCmd)
			
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
				} else {
					t.Errorf("Expected CLIError, got %T", err)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got: %v", err)
				}
			}
		})
	}
}