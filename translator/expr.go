package translator

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"

	"github.com/YongHaoWu/betterGo/utils"
)

func extractParamsTypeAndName(listOfArgs []ast.Expr) string {
	var paramsType string
	argname := "argname"
	argsNum := len(listOfArgs)
	for index, arg := range listOfArgs {
		argname = utils.IncrementString(argname, "", 1)
		switch x := arg.(type) {
		case *ast.BasicLit:
			switch x.Kind {
			case token.INT:
				paramsType = fmt.Sprintf("%s %s int", paramsType, argname)
			}
		case *ast.Ident:
			argVarName := x.Name
			fmt.Println("argVarName is ", argVarName)
			fmt.Println("argVarType is ", variableType[argVarName])
			paramsType = fmt.Sprintf("%s %s %s", paramsType, argname, variableType[argVarName])

			if index != argsNum-1 {
				paramsType += ","
			}
		}
	}
	return paramsType
}

func GetExprStr(fset *token.FileSet, expr interface{}) string {
	name := new(bytes.Buffer)
	printer.Fprint(name, fset, expr)
	return name.String()
}
