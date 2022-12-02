package analyzer

import (
	"go/ast"
	_ "go/ast"
	_ "strings"

	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name:     "gostructlint",
	Doc:      "Checks if structs attributes are placed in the order big->small",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.StructType)(nil),
	}
	inspector.Preorder(nodeFilter, func(node ast.Node) {
		structType := node.(*ast.StructType)
		if structType.Incomplete {
			return
		}
		if len(structType.Fields.List) < 2 {
			return
		}
		pass.Reportf(node.Pos(), "%d struct fields aren't in correct order", structType.Pos)
	})

	return nil, nil
}
