package erroriface

import (
	"fmt"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

const doc = "erroriface is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "erroriface",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	errObj := types.Universe.Lookup("error")
	if errObj == nil {
		fmt.Println("could not find errObj")
		return nil, nil
	}
	errIface := errObj.Type().Underlying().(*types.Interface)

	s := pass.Pkg.Scope()
	for _, name := range s.Names() {
		obj, _ := s.Lookup(name).(*types.TypeName)
		if obj == nil {
			continue
		}
		typ := obj.Type()
		if types.Implements(typ, errIface) || types.Implements(types.NewPointer(typ), errIface) {
			fmt.Printf("%v implements error interface\n", obj.Name())
		}
	}

	return nil, nil
}
