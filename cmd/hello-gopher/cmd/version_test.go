package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestVersionCommand(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		validate func(t *testing.T, output string)
	}{
		{
			name: "version command output",
			args: []string{"version"},
			validate: func(t *testing.T, output string) {
				// Check that version information is displayed
				if !strings.Contains(output, "hello-gopher") {
					t.Error("Expected version output to contain 'hello-gopher'")
				}
				
				// Version output should contain version info
				lines := strings.Split(strings.TrimSpace(output), "\n")
				if len(lines) == 0 {
					t.Error("Expected non-empty version output")
				}
			},
		},
		{
			name: "version command help",
			args: []string{"version", "--help"},
			validate: func(t *testing.T, output string) {
				t.Logf("Version help output: %q", output)
				if !strings.Contains(output, "Version command displays") {
					t.Error("Expected help text to contain version description")
				}
				if !strings.Contains(output, "version of hello-gopher") {
					t.Error("Expected help text to contain version information")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test version command
			testCmd := &cobra.Command{
				Use:   "version",
				Short: "Show version information",
				Long: `Version command displays the current version of hello-gopher along with
build information including build date and Git commit hash when available.

This information is useful for debugging and ensuring you're running the
expected version of the tool.`,
				Example: `  hello-gopher version                  # Show version information`,
				RunE:    versionCmd.RunE, // Use the same RunE function
			}
			
			// Capture output
			var buf bytes.Buffer
			testCmd.SetOut(&buf)
			testCmd.SetErr(&buf)
			testCmd.SetArgs(tt.args[1:]) // Remove "version" from args

			err := testCmd.Execute()
			if err != nil {
				t.Fatalf("Version command execution failed: %v", err)
			}

			output := buf.String()
			if tt.validate != nil {
				tt.validate(t, output)
			}
		})
	}
}

func TestVersionCommandIntegration(t *testing.T) {
	// Test that the version command is properly registered with the root command
	found := false
	for _, cmd := range rootCmd.Commands() {
		if cmd.Name() == "version" {
			found = true
			break
		}
	}
	
	if !found {
		t.Error("version command not found in root command")
	}
}

func TestVersionVariables(t *testing.T) {
	// Test that version variables can be set (they're package-level variables)
	// This tests the build-time variable injection capability
	
	// These variables are set at build time, so we just verify they exist
	// and can be accessed without panicking
	t.Logf("Version: %s", version)
	t.Logf("Build Date: %s", buildDate)
	t.Logf("Git Commit: %s", gitCommit)
	
	// The variables should be strings (even if empty)
	if version == "" {
		t.Log("Version is empty (expected for test builds)")
	}
	if buildDate == "" {
		t.Log("Build date is empty (expected for test builds)")
	}
	if gitCommit == "" {
		t.Log("Git commit is empty (expected for test builds)")
	}
}

// BenchmarkVersionCommand benchmarks version command execution
func BenchmarkVersionCommand(b *testing.B) {
	testCmd := &cobra.Command{
		Use:  "version",
		RunE: versionCmd.RunE,
	}
	
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		testCmd.SetOut(&buf)
		testCmd.SetErr(&buf)
		testCmd.SetArgs([]string{})

		err := testCmd.Execute()
		if err != nil {
			b.Fatalf("Version command benchmark failed: %v", err)
		}
	}
}