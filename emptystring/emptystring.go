package emptystring

import (
	"go/ast"
	"go/token"
	"strconv"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

const doc = "emptystring is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "emptystring",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		for _, decl := range f.Decls {
			decl, _ := decl.(*ast.FuncDecl)
			if decl == nil {
				continue
			}
			for _, stmt := range decl.Body.List {
				ret, _ := stmt.(*ast.ReturnStmt)
				if ret == nil || len(ret.Results) != 1 {
					continue
				}
				lit, _ := ret.Results[0].(*ast.BasicLit)
				if lit == nil || lit.Kind != token.STRING {
					continue
				}
				str, err := strconv.Unquote(lit.Value)
				if err != nil {
					return nil, err
				}
				if str == "" {
					pass.Reportf(pass.Pos, "NG")
				}

			}
		}

	}
	return nil, nil
}
