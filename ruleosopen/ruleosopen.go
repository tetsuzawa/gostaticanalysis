package ruleosopen

import (
	"fmt"
	"github.com/gostaticanalysis/analysisutil"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/ssa"
	"strconv"
	"strings"

	"golang.org/x/tools/go/analysis"
)

const doc = "ruleosopen is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "ruleosopen",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		buildssa.Analyzer,
	},
}

const ImportPath = "os"
const retTypeName = "*File"
const closeMethods = "Close"

func run(pass *analysis.Pass) (interface{}, error) {
	funcs := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA).SrcFuncs
	//cmaps := pass.ResultOf[commentmap.Analyzer].(comment.Maps)

	openTyp := analysisutil.TypeOf(pass, ImportPath, retTypeName)
	if openTyp == nil {
		// skip checking
		return nil, nil
	}

	var methods []*types.Func
	for _, s := range strings.Split(closeMethods, ",") {
		if m := analysisutil.MethodOf(openTyp, s); m != nil {
			// osパッケージからCloseという名前のmethodをとってくる
			methods = append(methods, m)
		}
	}

	skipFile := map[*ast.File]bool{}
	for _, f := range funcs {
		if Unimported(pass, f, skipFile, ImportPath) {
			// skip this
			continue
		}

		for k, b := range f.Blocks {
			fmt.Println(k)
			for i, instr := range b.Instrs {
				pos := instr.Pos()
				//line := pass.Fset.File(pos).Line(pos)

				// skip
				//if cmaps.IgnoreLine(pass.Fset, line, "") ||
				//	cmaps.IgnoreLine(pass.Fset, line, "unclosetx") {
				//	continue
				//}
				if ex, ok := instr.(*ssa.Extract); ok {
					pos = ex.Tuple.Pos()
				}

				called, ok := analysisutil.CalledFrom(b, i, openTyp, methods...)
				//ブロックの中で呼ばれかどうかを判定
				if ok && !called {
					pass.Reportf(pos, "NG")
				}
			}
		}
	}

	return nil, nil
}

func Unimported(pass *analysis.Pass, f *ssa.Function, skipFile map[*ast.File]bool, importPath string) (ret bool) {
	obj := f.Object()
	if obj == nil {
		return false
	}

	file := analysisutil.File(pass, obj.Pos())
	if file == nil {
		return false
	}

	if skip, has := skipFile[file]; has {
		return skip
	}
	defer func() {
		skipFile[file] = ret
	}()

	for _, impt := range file.Imports {
		path, err := strconv.Unquote(impt.Path.Value)
		if err != nil {
			continue
		}
		path = analysisutil.RemoveVendor(path)
		if path == importPath {
			return false
		}
	}

	return true
}
