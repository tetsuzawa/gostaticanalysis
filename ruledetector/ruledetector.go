package ruledetector

import (
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"github.com/gostaticanalysis/analysisutil"
	"github.com/gostaticanalysis/comment"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

const doc = "ruledetector is ..."

const rulePrefix = "Rule_"

func Run() error {
	flag.Parse()
	mode := packages.NeedName |
		packages.NeedFiles |
		packages.NeedCompiledGoFiles |
		packages.NeedImports |
		packages.NeedTypes |
		packages.NeedTypesSizes |
		packages.NeedSyntax |
		packages.NeedTypesInfo |
		packages.NeedDeps
	cfg := &packages.Config{Mode: mode}
	pkgs, err := packages.Load(cfg, flag.Args()...)
	if err != nil {
		return err
	}
	if packages.PrintErrors(pkgs) > 0 {
		return errors.New("some errors occurred")
	}
	var maps comment.Maps
	for _, pkg := range pkgs {
		fmt.Println("analysisutil.LookupFromImports(pkg.Types.Imports(), os, Open):", analysisutil.LookupFromImports(pkg.Types.Imports(), "os", "Open"))
		//for _, f := range pkg.Syntax {
		//ast.Print(pkg.Fset, f)
		//}
		maps = comment.New(pkg.Fset, pkg.Syntax)
	}
	//fmt.Println("len pkgs", len(pkgs))

	_, ssaPkgs := ssautil.AllPackages(pkgs, 0)
	for pkgIdx, ssaPkg := range ssaPkgs {
		//fmt.Println("len ssa", len(ssaPkgs))
		ssaPkg.Build()
		for _, member := range (*ssaPkg).Members {
			// debug
			//fmt.Println(name, member)
			//fmt.Println(member.Type())
			//fmt.Println(member.Object())
			//fmt.Println(types.Identical(member.Type(), &types.Signature{}))
			//fmt.Println()

			// 取得したのが関数型か判定
			if !types.Identical(member.Type(), &types.Signature{}) {
				continue
			}
			//名前がRule_xxxか判定
			if !strings.HasPrefix(member.Name(), rulePrefix) {
				fmt.Println("no", member.Name())
				continue
			}
			f, ok := member.(*ssa.Function)
			if !ok {
				fmt.Println("ok")
			}
			fmt.Println("f", f.String())
			for _, b := range f.Blocks {
				fmt.Println("b", b.String())
				commentProcessed := map[*ast.CommentGroup]bool{}
				for _, instr := range b.Instrs {
					//pos := b.Instrs[i].Pos()
					//fmt.Println("maps.CommentsByPosLine", maps.CommentsByPosLine(pkgs[pkgIdx].Fset, instr.Pos()))
					for _, c := range maps.CommentsByPosLine(pkgs[pkgIdx].Fset, instr.Pos()) {
						if !commentProcessed[c] {
							// Commentを取得
							fmt.Println("CommentGroup", c.Text())
						}
						commentProcessed[c] = true
					}
					//fmt.Println("maps.CommentsByPosLine", maps.CommentsByPos(instr.Pos()))
					fmt.Println("instr", instr.String())
					//instrから関数の型を取りたい(osパッケージのOpen関数という型）
					call, ok := instr.(*ssa.Call)
					//callCommon := call.Common()
					if !ok {
						continue
					}
					if call.Common() == nil {
						fmt.Println("call.Common() is nil!!")
						continue
					}
					value := call.Common().Value
					method := call.Common().Method
					if method == nil {
						fmt.Println("method is nil!!")
						continue
					}
					fmt.Println("callCommon.Method", call.Common().Method)
					fmt.Println(lookupFromImports(pkgs[pkgIdx].Types.Imports(), method.Pkg().Name(), method.Name()))
					fmt.Println("call.Name", call.Name())
					fmt.Println("call.Type", call.Type())
					//fmt.Println(lookupFromImports(pkgs[pkgIdx].Types.Imports(), "os", "Open"))
				}
			}
			fmt.Printf("\n\n ----------------------------------------------------- \n\n")

			//if ssaPkg != nil {
			//	ssaPkg.Build()
			//	scope := ssaPkg.Pkg.Scope()
			//	fmt.Println(scope.String())
			//	for _, name := range scope.Names() {
			//		obj, _ := scope.Lookup(name).(*types.Func)
			//		if obj == nil {
			//			continue
			//		}
			//		fmt.Println(obj.Name())
			//
			//		// 名前がRule_xxxか判定
			//		if !hasRulePrefix(obj.Name()){
			//			continue
			//		}
			//
			//		typ := obj.Type()
			//		fmt.Println(typ.String())
			//		//if types.Implements(typ, errIface) || types.Implements(types.NewPointer(typ), errIface) {
			//		//	fmt.Printf("%member implements error interface\n", obj.Name())
			//		//}
			//	}

		}
	}
	return nil
}

func lookupFromImports(imports []*types.Package, pkg, name string) types.Object {
	if name == "" {
		return nil
	}

	if name[0] == '*' {
		return lookupFromImports(imports, pkg, name[1:])
	}

	return analysisutil.LookupFromImports(imports, pkg, name)
}
