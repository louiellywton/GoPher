package cmd

import (
	"bytes"
	"errors"
	"testing"

	"github.com/spf13/cobra"
)

// BenchmarkGreetCommand benchmarks greet command execution
func BenchmarkGreetCommand(b *testing.B) {
	b.Run("DefaultName", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			cmd := &cobra.Command{
				Use: "greet",
				RunE: greetCmd.RunE,
			}
			cmd.Flags().StringP("name", "n", "", "Name to greet")
			cmd.SetOut(bytes.NewBuffer(nil))
			cmd.SetErr(bytes.NewBuffer(nil))
			cmd.SetArgs([]string{})
			_ = cmd.Execute()
		}
	})
	
	b.Run("WithName", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			cmd := &cobra.Command{
				Use: "greet",
				RunE: greetCmd.RunE,
			}
			cmd.Flags().StringP("name", "n", "", "Name to greet")
			cmd.SetOut(bytes.NewBuffer(nil))
			cmd.SetErr(bytes.NewBuffer(nil))
			cmd.SetArgs([]string{"--name", "BenchUser"})
			_ = cmd.Execute()
		}
	})
	
	b.Run("ShortFlag", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			cmd := &cobra.Command{
				Use: "greet",
				RunE: greetCmd.RunE,
			}
			cmd.Flags().StringP("name", "n", "", "Name to greet")
			cmd.SetOut(bytes.NewBuffer(nil))
			cmd.SetErr(bytes.NewBuffer(nil))
			cmd.SetArgs([]string{"-n", "BenchUser"})
			_ = cmd.Execute()
		}
	})
}

// BenchmarkProverbCommand benchmarks proverb command execution
func BenchmarkProverbCommand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd := &cobra.Command{
			Use: "proverb",
			RunE: proverbCmd.RunE,
		}
		cmd.SetOut(bytes.NewBuffer(nil))
		cmd.SetErr(bytes.NewBuffer(nil))
		cmd.SetArgs([]string{})
		_ = cmd.Execute()
	}
}

// BenchmarkRootCommandVersion benchmarks root command version flag
func BenchmarkRootCommandVersion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd := &cobra.Command{
			Use: "hello-gopher",
			RunE: rootCmd.RunE,
		}
		cmd.Flags().BoolP("version", "v", false, "version info")
		cmd.SetOut(bytes.NewBuffer(nil))
		cmd.SetErr(bytes.NewBuffer(nil))
		cmd.SetArgs([]string{"--version"})
		_ = cmd.Execute()
	}
}

// BenchmarkRootCommandHelp benchmarks root command help
func BenchmarkRootCommandHelp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd := &cobra.Command{
			Use: "hello-gopher",
			RunE: rootCmd.RunE,
		}
		cmd.SetOut(bytes.NewBuffer(nil))
		cmd.SetErr(bytes.NewBuffer(nil))
		cmd.SetArgs([]string{"--help"})
		_ = cmd.Execute()
	}
}

// BenchmarkErrorHandling benchmarks error handling performance
func BenchmarkErrorHandling(b *testing.B) {
	b.Run("UsageError", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			err := NewUsageError("test message", "test suggestion")
			_ = err.Error()
		}
	})
	
	b.Run("DataError", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			err := NewDataError("test message", errors.New("cause"), "test suggestion")
			_ = err.Error()
		}
	})
	
	b.Run("SystemError", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			err := NewSystemError("test message", errors.New("cause"), "test suggestion")
			_ = err.Error()
		}
	})
}

// BenchmarkCommandCreation benchmarks command creation
func BenchmarkCommandCreation(b *testing.B) {
	b.Run("GreetCmd", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			cmd := &cobra.Command{
				Use: "greet",
				RunE: greetCmd.RunE,
			}
			cmd.Flags().StringP("name", "n", "", "Name to greet")
			_ = cmd
		}
	})
	
	b.Run("ProverbCmd", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			cmd := &cobra.Command{
				Use: "proverb",
				RunE: proverbCmd.RunE,
			}
			_ = cmd
		}
	})
	
	b.Run("VersionCmd", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			cmd := &cobra.Command{
				Use: "version",
				Run: versionCmd.Run,
			}
			_ = cmd
		}
	})
}

// BenchmarkFlagParsing benchmarks flag parsing performance
func BenchmarkFlagParsing(b *testing.B) {
	b.Run("GreetNameFlag", func(b *testing.B) {
		cmd := &cobra.Command{Use: "greet"}
		cmd.Flags().StringP("name", "n", "", "Name to greet")
		for i := 0; i < b.N; i++ {
			cmd.SetArgs([]string{"--name", "TestUser"})
			cmd.ParseFlags([]string{"--name", "TestUser"})
		}
	})
	
	b.Run("VersionFlag", func(b *testing.B) {
		cmd := &cobra.Command{Use: "hello-gopher"}
		cmd.Flags().BoolP("version", "v", false, "version info")
		for i := 0; i < b.N; i++ {
			cmd.SetArgs([]string{"--version"})
			cmd.ParseFlags([]string{"--version"})
		}
	})
}

// BenchmarkStringFormatting benchmarks string formatting in commands
func BenchmarkStringFormatting(b *testing.B) {
	b.Run("VersionOutput", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = formatVersionInfo("1.0.0", "2023-01-01", "abc123")
		}
	})
	
	b.Run("ErrorMessage", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			err := NewUsageError("test error", "test suggestion")
			_ = err.Error()
		}
	})
}

// Helper function for benchmarking version formatting
func formatVersionInfo(version, buildDate, gitCommit string) string {
	return "hello-gopher version " + version + "\n" +
		"Build date: " + buildDate + "\n" +
		"Git commit: " + gitCommit + "\n"
}

// BenchmarkCommandExecution benchmarks full command execution pipeline
func BenchmarkCommandExecution(b *testing.B) {
	b.Run("FullGreetPipeline", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			rootCmd := &cobra.Command{Use: "hello-gopher"}
			greetCmd := &cobra.Command{
				Use: "greet",
				RunE: greetCmd.RunE,
			}
			greetCmd.Flags().StringP("name", "n", "", "Name to greet")
			rootCmd.AddCommand(greetCmd)
			
			rootCmd.SetOut(bytes.NewBuffer(nil))
			rootCmd.SetErr(bytes.NewBuffer(nil))
			rootCmd.SetArgs([]string{"greet", "--name", "BenchUser"})
			_ = rootCmd.Execute()
		}
	})
	
	b.Run("FullProverbPipeline", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			rootCmd := &cobra.Command{Use: "hello-gopher"}
			proverbCmd := &cobra.Command{
				Use: "proverb",
				RunE: proverbCmd.RunE,
			}
			rootCmd.AddCommand(proverbCmd)
			
			rootCmd.SetOut(bytes.NewBuffer(nil))
			rootCmd.SetErr(bytes.NewBuffer(nil))
			rootCmd.SetArgs([]string{"proverb"})
			_ = rootCmd.Execute()
		}
	})
}