package privatetag_test

import (
	"testing"

	"github.com/tetsuzawa/gostaticanalysis/privatetag"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, privatetag.Analyzer, "a")
}
