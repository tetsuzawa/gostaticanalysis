package main

import (
	"github.com/tetsuzawa/gostaticanalysis/dupimporttest"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(dupimporttest.Analyzer) }
