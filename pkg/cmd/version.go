package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/CathalMullan/barb/pkg/validators"
	"github.com/CathalMullan/barb/pkg/version"
)

type versionCmd struct {
	cmd *cobra.Command
}

func newVersionCmd() *versionCmd {
	return &versionCmd{
		cmd: &cobra.Command{
			Use:   "version",
			Args:  validators.NoArgs,
			Short: "Check the version of barb",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Print(version.Template)
				version.CheckLatestVersion()
			},
			DisableFlagsInUseLine: true,
		},
	}
}
