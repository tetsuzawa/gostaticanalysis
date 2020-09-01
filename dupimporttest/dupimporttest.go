package dupimporttest

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"strconv"
)

const doc = "dupimporttest is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "dupimporttest",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		paths := make(map[string]bool)
		for _, ip := range f.Imports {
			path, err := strconv.Unquote(ip.Path.Value)
			if err != nil {
				return nil, err
			}
			if paths[path] {
				pass.Reportf(ip.Pos(), "%s is duplicated import", path)
			} else {
				paths[path] = true
			}
		}
	}
	return nil, nil
}
