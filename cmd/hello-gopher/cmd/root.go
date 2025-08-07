package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	// These variables are set at build time using ldflags
	version   = "dev"
	buildDate = "unknown"
	gitCommit = "unknown"
)

var rootCmd = &cobra.Command{
	Use:   "hello-gopher",
	Short: "A friendly CLI tool for Go enthusiasts",
	Long: `Hello-Gopher is a friendly command-line tool that demonstrates Go development best practices.
It provides greeting functionality and displays random Go proverbs, serving as a portfolio piece
that showcases idiomatic Go code, comprehensive testing, and professional distribution.

Examples:
  hello-gopher greet                    # Greet the default gopher
  hello-gopher greet --name Alice       # Greet Alice
  hello-gopher greet -n Bob             # Greet Bob (short flag)
  hello-gopher proverb                  # Display a random Go proverb
  hello-gopher --version                # Show version information`,
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		versionFlag, _ := cmd.Flags().GetBool("version")
		if versionFlag {
			cmd.Printf("hello-gopher version %s\n", version)
			cmd.Printf("Build date: %s\n", buildDate)
			cmd.Printf("Git commit: %s\n", gitCommit)
			cmd.Printf("Go version: %s\n", runtime.Version())
			cmd.Printf("OS/Arch: %s/%s\n", runtime.GOOS, runtime.GOARCH)
			return nil
		}

		// If unexpected arguments are provided, show error
		if len(args) > 0 {
			return NewUsageError(
				fmt.Sprintf("Unknown command: %s", args[0]),
				"Run 'hello-gopher --help' to see available commands",
			)
		}

		// If no subcommand is provided, show help
		cmd.Help()
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		HandleError(err)
	}
}

func init() {
	// Add version flag to root command
	rootCmd.Flags().BoolP("version", "v", false, "version for hello-gopher")

	// Set custom error handling for unknown flags
	rootCmd.SetFlagErrorFunc(func(cmd *cobra.Command, err error) error {
		return NewUsageError(
			err.Error(),
			fmt.Sprintf("Run '%s --help' for usage information", cmd.CommandPath()),
		)
	})
}