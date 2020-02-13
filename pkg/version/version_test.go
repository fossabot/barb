package version

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNeedsToUpgrade(t *testing.T) {
	require.False(t, needsToUpgrade("4.2.4.2", "4.2.4.2"))
	require.True(t, needsToUpgrade("4.2.4.2", "4.2.4.3"))
}
