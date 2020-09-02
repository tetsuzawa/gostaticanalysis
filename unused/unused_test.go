package unused_test

import (
	"testing"

	"github.com/tetsuzawa/gostaticanalysis/unused"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, unused.Analyzer, "a")
}

