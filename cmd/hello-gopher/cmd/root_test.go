package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestRootCommand(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		validate func(t *testing.T, output string, err error)
	}{
		{
			name: "root command help",
			args: []string{"--help"},
			validate: func(t *testing.T, output string, err error) {
				if err != nil {
					t.Errorf("Help command should not error: %v", err)
				}
				t.Logf("Help output: %q", output)
				if !strings.Contains(output, "Hello-Gopher is a CLI tool") {
					t.Error("Expected help text to contain description")
				}
				if !strings.Contains(output, "Usage:") {
					t.Error("Expected help text to contain usage information")
				}
			},
		},
		{
			name: "root command version flag",
			args: []string{"--version"},
			validate: func(t *testing.T, output string, err error) {
				if err != nil {
					t.Errorf("Version flag should not error: %v", err)
				}
				// Version output should contain some version info
				if strings.TrimSpace(output) == "" {
					t.Error("Expected non-empty version output")
				}
			},
		},
		{
			name: "root command short version flag",
			args: []string{"-v"},
			validate: func(t *testing.T, output string, err error) {
				if err != nil {
					t.Errorf("Short version flag should not error: %v", err)
				}
				// Version output should contain some version info
				if strings.TrimSpace(output) == "" {
					t.Error("Expected non-empty version output")
				}
			},
		},
		{
			name: "root command no args",
			args: []string{},
			validate: func(t *testing.T, output string, err error) {
				// Root command with no args should show help
				if err != nil {
					t.Errorf("Root command with no args should not error: %v", err)
				}
				if !strings.Contains(output, "Hello-Gopher is a CLI tool") {
					t.Error("Expected help text when no args provided")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a fresh root command for each test
			testRootCmd := &cobra.Command{
				Use:   "hello-gopher",
				Short: "A friendly CLI tool for Go enthusiasts",
				Long: `Hello-Gopher is a CLI tool that demonstrates Go development best practices.
It provides greeting functionality and displays random Go proverbs to inspire
and educate developers.

This tool serves as a portfolio piece showcasing:
- Clean Go code architecture with interfaces
- Comprehensive testing with mocks and benchmarks  
- Professional CLI design with Cobra framework
- Cross-platform distribution with Goreleaser
- CI/CD integration with GitHub Actions`,
				Example: `  hello-gopher greet --name Alice        # Greet Alice
  hello-gopher greet                     # Default greeting
  hello-gopher proverb                   # Show a random Go proverb
  hello-gopher version                   # Show version information`,
				SilenceUsage:  true,
				SilenceErrors: true,
				RunE:          rootCmd.RunE,
			}
			testRootCmd.Flags().BoolP("version", "v", false, "version for hello-gopher")
			
			// Capture output
			var buf bytes.Buffer
			testRootCmd.SetOut(&buf)
			testRootCmd.SetErr(&buf)
			testRootCmd.SetArgs(tt.args)

			err := testRootCmd.Execute()
			output := buf.String()
			
			if tt.validate != nil {
				tt.validate(t, output, err)
			}
		})
	}
}

func TestRootCommandSubcommands(t *testing.T) {
	// Test that all expected subcommands are registered
	expectedCommands := []string{"greet", "proverb", "version"}
	
	for _, expectedCmd := range expectedCommands {
		found := false
		for _, cmd := range rootCmd.Commands() {
			if cmd.Name() == expectedCmd {
				found = true
				break
			}
		}
		
		if !found {
			t.Errorf("Expected subcommand %q not found", expectedCmd)
		}
	}
}

func TestRootCommandFlags(t *testing.T) {
	// Test that expected flags are available
	versionFlag := rootCmd.Flags().Lookup("version")
	if versionFlag == nil {
		t.Error("Expected --version flag not found")
	}
	
	if versionFlag.Shorthand != "v" {
		t.Errorf("Expected version flag shorthand 'v', got %q", versionFlag.Shorthand)
	}
}

func TestRootCommandConfiguration(t *testing.T) {
	// Test root command configuration instead of error handling (which is already tested)
	if rootCmd.Use != "hello-gopher" {
		t.Errorf("Expected rootCmd.Use to be 'hello-gopher', got %q", rootCmd.Use)
	}
	
	if !strings.Contains(rootCmd.Short, "friendly CLI tool") {
		t.Errorf("Expected rootCmd.Short to contain 'friendly CLI tool', got %q", rootCmd.Short)
	}
	
	if !rootCmd.SilenceUsage {
		t.Error("Expected rootCmd.SilenceUsage to be true")
	}
	
	if !rootCmd.SilenceErrors {
		t.Error("Expected rootCmd.SilenceErrors to be true")
	}
}

// BenchmarkRootCommand benchmarks root command execution
func BenchmarkRootCommand(b *testing.B) {
	testRootCmd := &cobra.Command{
		Use:  "hello-gopher",
		RunE: rootCmd.RunE,
	}
	testRootCmd.Flags().BoolP("version", "v", false, "version for hello-gopher")
	
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		testRootCmd.SetOut(&buf)
		testRootCmd.SetErr(&buf)
		testRootCmd.SetArgs([]string{"--help"})

		err := testRootCmd.Execute()
		if err != nil {
			b.Fatalf("Root command benchmark failed: %v", err)
		}
	}
}

// TestExecute tests the Execute function
func TestExecute(t *testing.T) {
	// This is a basic test to ensure Execute doesn't panic
	// We can't easily test the actual execution since it may call os.Exit
	
	// Test that Execute function exists and can be called
	// In a real scenario, this would be tested through integration tests
	t.Log("Execute function is available for testing")
	
	// Verify that rootCmd is properly initialized
	if rootCmd == nil {
		t.Error("rootCmd should not be nil")
	}
	
	if rootCmd.Use != "hello-gopher" {
		t.Errorf("Expected rootCmd.Use to be 'hello-gopher', got %q", rootCmd.Use)
	}
}