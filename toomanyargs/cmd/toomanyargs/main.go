package main

import (
	"github.com/tetsuzawa/gostaticanalysis/toomanyargs"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(toomanyargs.Analyzer) }
