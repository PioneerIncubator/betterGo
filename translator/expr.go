package translator

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"

	"github.com/YongHaoWu/betterGo/utils"
)

func ExtractParamsTypeAndName(listOfArgs []ast.Expr) (string, []string, []string) {
	var paramsType string
	var listOfArgVarNames []string
	var listOfArgTypes []string
	argname := "argname"
	argsNum := len(listOfArgs)
	for i, arg := range listOfArgs {
		argname = utils.IncrementString(argname, "", 1)
		switch x := arg.(type) {
		case *ast.BasicLit:
			switch x.Kind {
			case token.INT:
				argVarName := x.Value
				listOfArgVarNames = append(listOfArgVarNames, argVarName)
				listOfArgTypes = append(listOfArgTypes, "int")
				paramsType = fmt.Sprintf("%s %s int", paramsType, argname)
			}
		case *ast.Ident:
			argVarName := x.Name
			listOfArgVarNames = append(listOfArgVarNames, argVarName)
			listOfArgTypes = append(listOfArgTypes, variableType[argVarName])
			fmt.Println("argVarName is ", argVarName)
			var argVarType string
			if paramType, ok := variableType[DecorateParamName(argVarName)]; ok {
				argVarType = paramType
			} else {
				argVarType = variableType[argVarName]
			}
			fmt.Println("argVarType is ", argVarType)
			paramsType = fmt.Sprintf("%s %s %s", paramsType, argname, argVarType)

			if i != argsNum-1 {
				paramsType += ","
			}
		}
	}
	listOfArgTypes = append(listOfArgTypes, assertType)
	return paramsType, listOfArgVarNames, listOfArgTypes
}

func GetExprStr(fset *token.FileSet, expr interface{}) string {
	name := new(bytes.Buffer)
	printer.Fprint(name, fset, expr)
	return name.String()
}
