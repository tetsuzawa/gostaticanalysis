package main

import (
	"github.com/tetsuzawa/gostaticanalysis/emptystring"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(emptystring.Analyzer) }
