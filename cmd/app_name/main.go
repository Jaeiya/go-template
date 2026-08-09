package main

import (
	"fmt"

	"github.com/jaeiya/go-template/internal"
)

var (
	version   string
	commitSha string
	buildDate string
)

func main() {
	fmt.Printf(
		"\n\n   Version: %s\n CommitSha: %s\nBuildDate: %s\n\n",
		version,
		commitSha,
		buildDate,
	)
	internal.LibFunc()
	fmt.Printf("\n")
}
