package main

import (
	"github.com/tetsuzawa/gostaticanalysis/rulehttpclient"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(rulehttpclient.Analyzer) }

