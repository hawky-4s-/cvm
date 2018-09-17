package internal

import "fmt"

var (
	version   = "unknown"
	commit    = "unknown"
	goVersion = "unknown"
)

func Version() string {
	return fmt.Sprintf("CVM version: %s\nGit commit: %s\nCompiled with Go: %s", version, commit, goVersion)
}
