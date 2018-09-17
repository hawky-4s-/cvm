package internal

import "fmt"

var (
	version   = "dev"
	commit    = "none"
	date      = "unknown"
	goversion = "unknown"
)

func Version() string {
	return fmt.Sprintf("CVM version %s \"%s\"\nGit commit hash: %s\nBuilt at: %s\nGo version: %s", version, commit, date, goversion)
}
