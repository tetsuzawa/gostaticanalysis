package toomanyargs

import (
	"go/ast"
	"strconv"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "toomanyargs is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "toomanyargs",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

// passed by `go build -ldflags "-X toomanyargs.numArgs=5"`
var numArgs string

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncType)(nil),
	}

	numArgs, err := strconv.Atoi(numArgs)
	if err != nil {
		return nil, err
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.FuncType:
			var count int
			for _, field := range n.Params.List {
				count += len(field.Names)
			}
			if count >= numArgs {
				pass.Reportf(n.Pos(), "too many number of args in the func declaration. should be less than %d", numArgs)
			}
		}
	})

	return nil, nil
}
