package shortpkgvar

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "shortpkgvar is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "shortpkgvar",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

var flagN int

func init() {
	Analyzer.Flags.IntVar(&flagN, "N", 1, "length")
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	inspect.Preorder(nil, func(n ast.Node) {
		s := pass.Pkg.Scope()
		for _, name := range s.Names() {
			if len(name) > flagN {
				continue
			}
			obj, _ := s.Lookup(name).(*types.Var)
			if obj == nil {
				continue
			}
			pass.Reportf(obj.Pos(), "NG")

		}
		//pkgName := pass.Pkg.Name()
		//fmt.Println(pkgName)

		//s := pass.TypesInfo.Scopes[n]
		//if s != nil {
		//	return
		//}

	})

	return nil, nil
}
