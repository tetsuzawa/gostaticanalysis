package main

import (
	"github.com/tetsuzawa/gostaticanalysis/ruleosopen"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(ruleosopen.Analyzer) }

