package privatetag

import (
	"go/ast"
	"golang.org/x/tools/go/ast/inspector"
	"strconv"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

const doc = "privatetag is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "privatetag",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.StructType)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.StructType:
			for _, f := range n.Fields.List {
				tag, err := strconv.Unquote(f.Tag.Value)
				if err != nil {
					pass.Reportf(f.Pos(), err.Error())
				}
				if hasJSONTag(tag) {
					for _, name := range f.Names {
						if !name.IsExported() {
							pass.Reportf(f.Pos(), "%s field is unexported though it has a json tag", name.String())
						}
					}
				}
				if hasDBTag(tag) {
					for _, name := range f.Names {
						if !name.IsExported() {
							pass.Reportf(f.Pos(), "%s field is unexported though it has a db tag", name.String())
						}
					}
				}
			}
		}
	})
	return nil, nil
}

func hasJSONTag(s string) bool {
	tags := strings.Split(s, " ")
	for _, tag := range tags {
		t := strings.Split(tag, ":")
		if t[0] == "json" {
			return true
		}
	}
	return false
}

func hasDBTag(s string) bool {
	tags := strings.Split(s, " ")
	for _, tag := range tags {
		t := strings.Split(tag, ":")
		if t[0] == "db" {
			return true
		}
	}
	return false
}
