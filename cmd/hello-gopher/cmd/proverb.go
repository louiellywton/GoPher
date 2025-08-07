package cmd

import (
	"fmt"

	"github.com/louiellywton/go-portfolio/01-hello-gopher/pkg/greeting"
	"github.com/spf13/cobra"
)

var proverbCmd = &cobra.Command{
	Use:   "proverb",
	Short: "Display a random Go proverb",
	Long: `Proverb command displays random Go proverbs to inspire and educate.
Each execution shows a different proverb from a curated collection of Go programming
wisdom and best practices.

This command demonstrates integration with the ProverbProvider interface and
proper error handling for data loading failures.`,
	Example: `  hello-gopher proverb                  # Display a random Go proverb`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Validate that no unexpected arguments were provided
		if len(args) > 0 {
			return NewUsageError(
				fmt.Sprintf("Unexpected argument(s): %v", args),
				"The proverb command doesn't accept any arguments",
			)
		}

		// Create greeting service and get a random proverb
		service := greeting.NewService()
		
		// Load proverbs first to handle any loading errors
		if err := service.LoadProverbs(); err != nil {
			return NewDataError(
				"Failed to load Go proverbs",
				err,
				"This appears to be a data issue. Please check if the application was built correctly",
			)
		}
		
		proverb := service.RandomProverb()
		cmd.Println(proverb)
		return nil
	},
}

func init() {
	// Add proverb command to root command
	rootCmd.AddCommand(proverbCmd)
}