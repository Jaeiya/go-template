package main

import (
	"fmt"

	"github.com/jaeiya/go-template/internal"
)

func main() {
	fmt.Printf("   Version: %s\n CommitSha: %s\n", internal.Version, internal.CommitSha)
}
