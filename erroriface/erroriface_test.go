package erroriface_test

import (
	"testing"

	"github.com/tetsuzawa/gostaticanalysis/erroriface"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, erroriface.Analyzer, "a")
}

