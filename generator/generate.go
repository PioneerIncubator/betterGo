package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "../main.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Functions:")
	for _, f := range node.Decls {
		fn, ok := f.(*ast.FuncDecl)
		if !ok {
			continue
		}
		fmt.Println("fn.Name", fn.Name)
		// if fn.Name.Name == "APIDaemon" {
		// 	ast.Inspect(fn, func(n ast.Node) bool {
		// 		if ret, ok := n.(*ast.CallExpr); ok {
		// 			funName := getExprStr(fset, ret.Fun)
		// 			switch funName {
		// 			case "handlers.bindRedisGetURLToActionType":
		// 				genListFunction(fset, ret)
		// 			case "handlers.bindURLToActionType":
		// 				genReidsCUD(fset, ret)
		// 			}
		// 		}
		// 		return true
		// 	})

		// }

	}

}
