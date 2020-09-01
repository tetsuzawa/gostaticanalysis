package dupimporttest_test

import (
	"testing"

	"github.com/tetsuzawa/gostaticanalysis/dupimporttest"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, dupimporttest.Analyzer, "a")
}
