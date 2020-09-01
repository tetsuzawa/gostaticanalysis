package numvar

import (
	"fmt"
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

const doc = "numvar is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "numvar",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		var count int
		for _, decl  := range f.Decls {
			decl, _ := decl.(*ast.GenDecl)
			if decl == nil {
				continue
			}
			for _, spec := range decl.Specs {
				spec , _ := spec.(*ast.ValueSpec)
				if spec == nil{
					continue
				}
				for _, id := range spec.Names {
					if id.Name != "_"{
						count += 1
					}
				}
			}
		}
		fmt.Println(count)
	}
	return nil, nil
}

