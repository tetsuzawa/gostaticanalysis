package main

import (
	"github.com/tetsuzawa/gostaticanalysis/erroriface"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(erroriface.Analyzer) }

