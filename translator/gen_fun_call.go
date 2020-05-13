package translator

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
)

func extractParamsName(listOfArgs []ast.Expr) string {
	var paramsName string
	for _, arg := range listOfArgs {
		switch x := arg.(type) {
		case *ast.BasicLit:
			switch x.Kind {
			case token.INT:
				paramsName = strings.Title(fmt.Sprintf("%s int", paramsName))
			}
		case *ast.Ident:
			paramsName = fmt.Sprintf("%s %s", paramsName, strings.Title(x.Name))
		}
	}
	return strings.ReplaceAll(paramsName, " ", "")
}

//  func Reduce(argname_1 []int, argname_2 func (int, int, string)int, argname_3 int) int
func genFunctionBody(funName string) string {
	var body string
	switch funName {
	case "Reduce":
		body = `
			lenSlice := len(argname_1)
			switch lenSlice {
			case 0:
				return 0
			case 1:
				return argname_1[1]
			}
			out := argname_2(argname_3, argname_1[0])
			next := argname_1[1]
			for i := 1; i < lenSlice; i++ {
				next = argname_1[i]
				out = argname_2(out, next)
			}
			return out
		`
	case "Add":
		body = `
			return argname_1 + argname_2
 		`
	}
	return body
}

func GenEnumFunctionDecl(funName string, listOfArgs []ast.Expr) (string, string) {
	paramsTypeDecl := extractParamsTypeAndName(listOfArgs)
	switch funName {
	case "enum.Reduce":
		// iterate function args to reveal the type
		//Reduce(slice, pairFunction, zero interface{}) interface{}
		funName = "Reduce"
	}
	functionBody := genFunctionBody(funName)

	funName += extractParamsName(listOfArgs)
	var funcitonDecl string
	if assertPassCnt == 1 {
		funcitonDecl = fmt.Sprintf(
			`func %s(%s) %s {
%s
		}`,
			funName,
			paramsTypeDecl,
			// TODO : Use sth bettor to record .assert
			assertType,
			functionBody,
		)
	}
	return funName, funcitonDecl
}
