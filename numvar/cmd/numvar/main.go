package main

import (
	"github.com/tetsuzawa/gostaticanalysis/numvar"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(numvar.Analyzer) }
