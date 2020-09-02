package main

import (
	"github.com/tetsuzawa/gostaticanalysis/unused"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(unused.Analyzer) }

