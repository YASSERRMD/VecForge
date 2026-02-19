package main

import "fmt"

var (
	Version   = "1.0.0"
	Commit    = "dev"
	Date      = "unknown"
)

func VersionString() string {
	return fmt.Sprintf("VecForge %s (commit: %s, date: %s)", Version, Commit, Date)
}
