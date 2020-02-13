package validators

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// NoArgs - Require no arguments.
func NoArgs(cmd *cobra.Command, args []string) error {
	errorMessage := fmt.Sprintf(
		"\"%s\" does not take any arguments. See \"barb %s --help\" for the correct usage.",
		cmd.Name(),
		cmd.Name(),
	)

	if len(args) > 0 {
		return errors.New(errorMessage)
	}

	return nil
}
