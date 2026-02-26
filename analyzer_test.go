package SelectelTest_test

import (
	"SelectelTest"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestSlogRules(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, SelectelTest.Analyzer, "slog_test")
}

func TestZapRules(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, SelectelTest.Analyzer, "zap_test")
}
