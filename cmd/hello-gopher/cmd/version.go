package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Long:  `Print detailed version information including build date and git commit.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("hello-gopher version %s\n", version)
		fmt.Printf("Build date: %s\n", buildDate)
		fmt.Printf("Git commit: %s\n", gitCommit)
		fmt.Printf("Go version: %s\n", runtime.Version())
		fmt.Printf("OS/Arch: %s/%s\n", runtime.GOOS, runtime.GOARCH)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}