package ruledetector

import (
	"errors"
	"flag"
	"fmt"
	"go/types"
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

	_, ssaPkgs := ssautil.AllPackages(pkgs, 0)
	for _, ssaPkg := range ssaPkgs {
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
			//if !strings.HasPrefix(member.Name(), rulePrefix) {
			//	fmt.Println("no", member.Name())
			//	continue
			//}
			f, ok := member.(*ssa.Function)
			if !ok {
				fmt.Println("ok")
			}
			fmt.Println("f", f.String())
			for _, b := range f.Blocks {
				fmt.Println("b", b.String())
				for _, instr := range b.Instrs {
					//pos := b.Instrs[i].Pos()
					fmt.Println("instr", instr.String())
				}
			}
			fmt.Println()

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
