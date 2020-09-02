package main

import (
	"github.com/tetsuzawa/gostaticanalysis/shortpkgvar"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(shortpkgvar.Analyzer) }

