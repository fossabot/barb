package validators

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func TestNoArgs(t *testing.T) {
	c := &cobra.Command{Use: "c"}
	var args []string

	result := NoArgs(c, args)
	require.Nil(t, result)
}

func TestNoArgsWithArgs(t *testing.T) {
	c := &cobra.Command{Use: "c"}
	args := []string{"foo"}

	result := NoArgs(c, args)
	require.EqualError(t, result, "\"c\" does not take any arguments. See \"barb c --help\" for the correct usage.")
}
