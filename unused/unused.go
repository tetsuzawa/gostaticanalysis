package unused

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "unused is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "unused",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.FuncDecl:
			if pass.Pkg.Path() == "main" && n.Name.Name == "main" || n.Name.Name == "init" {
				return
			}
			// 定義から持ってくる
			obj := pass.TypesInfo.Defs[n.Name]
			if obj.Exported() {
				return
			}
			if !used(pass, obj) {
				pass.Reportf(n.Pos(), "NG")
			}

		}
	})

	return nil, nil
}

func used(pass *analysis.Pass, obj types.Object) bool {
	for _, o := range pass.TypesInfo.Uses {
		if o == obj {
			return true
		}
	}
	return false
}
