package main

import (
	"fmt"
	"github.com/tetsuzawa/gostaticanalysis/ruledetector"
	"os"
)

func main() {
	if err := ruledetector.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
