package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/CathalMullan/barb/pkg/version"
)

// nolint:gochecknoglobals // TODO: Remove global.
var rootCmd = &cobra.Command{
	Use:           "barb",
	Short:         "A Docker Git Hook Manager",
	Long:          "barb is a Git Hook manager which uses Docker to run your hooks.",
	SilenceUsage:  true,
	SilenceErrors: true,
	Version:       version.Version,
}

// Execute - Entry point for the root command.
func Execute() {
	rootCmd.SetVersionTemplate(version.Template)

	if err := rootCmd.Execute(); err != nil {
		// Custom unknown command handling.
		if strings.Contains(err.Error(), "unknown command") {
			baseCmd := os.Args[1]
			unknownCmdStr := fmt.Sprintf("Unknown command \"%s\" for \"%s\".\n", baseCmd, rootCmd.CommandPath())

			suggestedCmdStr := ""
			suggestions := rootCmd.SuggestionsFor(baseCmd)
			if len(suggestions) > 0 {
				suggestedCmdStr = fmt.Sprintf("Did you mean \"%s\"?\n", suggestions[0])
			}

			helpCmdStr := fmt.Sprintf("\nSee \"barb --help\" for a list of available commands.")
			println(unknownCmdStr + suggestedCmdStr + helpCmdStr)
		} else {
			fmt.Println(err)
		}

		os.Exit(1)
	}
}

// nolint:gochecknoinits // TODO: Remove init.
func init() {
	rootCmd.Flags().BoolP("version", "v", false, "check the version of barb")

	rootCmd.AddCommand(newVersionCmd().cmd)
}
