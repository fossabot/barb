package version

import (
	"context"
	"fmt"

	"github.com/google/go-github/v28/github"
)

// Version of the CLI.
// This is set to the actual version by GoReleaser by the git tag assigned to the release.
// Versions built from source will always show master.
// nolint:gochecknoglobals // TODO: Remove global.
var Version = "master"

// Template for the version string.
// nolint:gochecknoglobals // TODO: Remove global.
var Template = fmt.Sprintf("barb version: %s\n", Version)

// CheckLatestVersion makes a request to the GitHub API to pull the latest release of the CLI.
func CheckLatestVersion() {
	latest := getLatestVersion()

	if needsToUpgrade(Version, latest) {
		fmt.Println("A newer version of barb is available:", latest)
	}
}

func needsToUpgrade(version, latest string) bool {
	return latest != "" && (latest != version)
}

func getLatestVersion() string {
	client := github.NewClient(nil)
	rep, _, err := client.Repositories.GetLatestRelease(context.Background(), "CathalMullan", "barb")

	if err != nil {
		// We don't want to fail any functionality or display errors for this so fail silently.
		// TODO: Log error.
		return ""
	}

	return *rep.TagName
}
