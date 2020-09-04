package rulehttpclient

import (
	"errors"
	"fmt"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
	"io"
	"strconv"

	"github.com/gostaticanalysis/analysisutil"
	"github.com/tetsuzawa/gostaticanalysis/analysisutil/milestonequeue"
	"github.com/tetsuzawa/gostaticanalysis/analysisutil/usedtype"
)

const doc = "rulehttpclient is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "rulehttpclient",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		buildssa.Analyzer,
	},
}

//const ImportPath = "net/http"
//const retTypeName = "*Response"
//const closeMethods = "Close"

func run(pass *analysis.Pass) (interface{}, error) {
	obj1 := analysisutil.ObjectOf(pass, "net/http", "Client")

	obj2 := analysisutil.MethodOf(analysisutil.TypeOf(pass, "net/http", "Client"), "Do")
	//iocloser:=analysisutil.ObjectOf(pass, "io", "Closer")
	typ := analysisutil.TypeOf(pass, "io", "Closer")
	obj3 := analysisutil.MethodOf(typ, "Close")
	pkgs, ok := packages.Load(0, "io")
	if !ok {
		return nil,errors.New("failed to load io")
	}
	for _, pkg:=range pkgs{
		//pkg.TypesInfo.
	}
	//fmt.Println(iocloser)
	//obj3, _, _ := types.LookupFieldOrMethod(analysisutil.ObjectOf(pass, "io", "Closer").Type(), false, analysisutil.ObjectOf(pass, "io", "Closer").Pkg(), "Do")
	var mq = milestonequeue.MilestoneQueue{
		obj1,
		obj2,
		obj3,
	}
	funcs := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA).SrcFuncs
	//cmaps := pass.ResultOf[commentmap.Analyzer].(comment.Maps)

	//openTyp := analysisutil.TypeOf(pass, ImportPath, retTypeName)
	//if openTyp == nil {
	//	// skip checking
	//	return nil, nil
	//}
	//
	//var methods []*types.Func
	//for _, s := range strings.Split(closeMethods, ",") {
	//	if m := analysisutil.MethodOf(openTyp, s); m != nil {
	//		// osパッケージからCloseという名前のmethodをとってくる
	//		methods = append(methods, m)
	//	}
	//}

	//skipFile := map[*ast.File]bool{}
	for _, f := range funcs {
		//if Unimported(pass, f, skipFile, ImportPath) {
		//	// skip this
		//	continue
		//}

		for _, b := range f.Blocks {
			//poss := map[token.Pos]bool{}
			for _, instr := range b.Instrs {
				//pos := instr.Pos()
				//line := pass.Fset.File(pos).Line(pos)

				// skip
				//if cmaps.IgnoreLine(pass.Fset, line, "") ||
				//	cmaps.IgnoreLine(pass.Fset, line, "unclosetx") {
				//	continue
				//}
				switch instr := instr.(type) {
				case *ssa.Alloc:
					typeName, ok := mq.Head().(*types.TypeName)
					if !ok {
						continue
					}
					if usedtype.Alloc(instr, typeName) {
						mq.Pop()
						continue
					}
				case *ssa.Call:
					fn, ok := mq.Head().(*types.Func)
					if !ok {
						continue
					}
					//TODO methodのレシーバに対応
					if usedtype.Func(instr, nil, fn) {
						mq.Pop()
						continue
					}
				case *ssa.Defer:
					fn, ok := mq.Head().(*types.Func)
					if !ok {
						continue
					}
					//TODO methodのレシーバに対応
					if usedtype.Defer(instr, nil, fn) {
						mq.Pop()
						continue
					}
				default:
					continue
				}

				//called, ok := analysisutil.CalledFrom(b, i, openTyp, methods...)

				//ブロックの中で呼ばれかどうかを判定
				//if ok && !called && !poss[pos] {
				//	pass.Reportf(pos, "NG")
				//	poss[pos] = true
				//}
			}
		}
		if mq.Len() != 0 {
			fmt.Println("the function does not match the rule")
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
