package main

import (
	"github.com/tetsuzawa/gostaticanalysis/privatetag"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(privatetag.Analyzer) }
