package main

import (
	"fmt"
	"os"

	"github.com/jaeiya/go-template/internal"
)

var (
	version   string
	commitSha string
	buildDate string
)

func main() {
	args := os.Args[1:]
	fmt.Printf(
		"\n\n      Args: %+v\n   Version: %s\n CommitSha: %s\n BuildDate: %s\n\n",
		args,
		version,
		commitSha,
		buildDate,
	)
	internal.LibFunc()
	fmt.Printf("\n")
}
