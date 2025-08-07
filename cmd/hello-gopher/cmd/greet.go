package cmd

import (
	"fmt"

	"github.com/louiellywton/go-portfolio/01-hello-gopher/pkg/greeting"
	"github.com/spf13/cobra"
)

var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "Greet a gopher by name",
	Long: `Greet command provides friendly greeting functionality.
By default, it greets "Gopher", but you can specify a custom name using the --name flag.

This command demonstrates basic CLI functionality with flag support and integration
with the greeting package interfaces.`,
	Example: `  hello-gopher greet                    # Greet the default gopher
  hello-gopher greet --name Alice       # Greet Alice
  hello-gopher greet -n Bob             # Greet Bob using short flag`,
	RunE: func(cmd *cobra.Command, args []string) error {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return NewSystemError(
				"Failed to parse command flags",
				err,
				"Try running 'hello-gopher greet --help' for usage information",
			)
		}

		// Validate that no unexpected arguments were provided
		if len(args) > 0 {
			return NewUsageError(
				fmt.Sprintf("Unexpected argument(s): %v", args),
				"The greet command doesn't accept positional arguments. Use --name flag instead",
			)
		}

		// Create greeting service and generate greeting
		service := greeting.NewService()
		greeting := service.Greet(name)
		
		fmt.Println(greeting)
		return nil
	},
}

func init() {
	// Add greet command to root command
	rootCmd.AddCommand(greetCmd)
	
	// Add name flag with both long and short versions
	greetCmd.Flags().StringP("name", "n", "", "Name to greet (default: Gopher)")
}