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
			paramsName = strings.Title(fmt.Sprintf("%s %s", paramsName, GetBasicLitType(x)))
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
	case "Map":
		body = `
			lenSlice := len(argname_1)
			if lenSlice == 0 {
				return
			}
			for i := range argname_1 {
				argname_1[i] = argname_2(argname_1[i])
			}
		`
	case "Delete":
		body = `
			lenSlice := len(argname_1)
			if lenSlice == 0 {
				return false
			}
			count := 0
			for i := range argname_1 {
				if argname_2(argname_1[i]) {
					argname_1[count] = argname_1[i]
					count++
				}
			}
			argname_1 = argname_1[:count]
			return true
		`
	case "Find":
		body = `
			lenSlice := len(argname_1)
			if lenSlice == 0 {
				return nil
			}
			for i := range argname_1 {
				if argname_2(argname_1[i]) {
					return argname_1[i]
				}
			}
			return nil
		`
	case "Sum":
		body = `
			lenSlice := len(argname_1)
			if lenSlice == 0 {
				return nil
			}
			sum := argname_1[0]
			for i := range argname_1 {
				sum += i
			}
			sum = sum - argname_1[0]
			return sum
		`

	case "Max":
		body = `
			lenSlice := len(argname_1)
			if lenSlice == 0 {
				return nil
			}
			maxValue := argname_1[0]
			for i := range argname_1 {
				if(i > maxValue) {
					maxValue = i
				}
			}
			return maxValue
		`
	case "Min":
		body = `
			lenSlice := len(argname_1)
			if lenSlice == 0 {
				return nil
			}
			minValue := argname_1[0]
			for i := range argname_1 {
				if(i < minValue) {
					minValue = i
				}
			}
			return minValue
		`
	}

	return body
}

func GenEnumFunctionDecl(fset *token.FileSet, funName string, listOfArgs []ast.Expr) (string, string) {
	paramsTypeDecl, _, _ := ExtractParamsTypeAndName(fset, listOfArgs)
	switch funName {
	case "enum.Reduce":
		// iterate function args to reveal the type
		// Reduce(slice, pairFunction, zero interface{}) interface{}
		funName = "Reduce"
	case "enum.Add":
		funName = "Add"
	case "enum.Map":
		funName = "Map"
	case "enum.Delete":
		funName = "Delete"
	case "enum.Find":
		funName = "Find"
	case "enum.Sum":
		funName = "Sum"
	case "enum.Max":
		funName = "Max"
	case "enum.Min":
		funName = "Min"
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
	} else {
		funcitonDecl = fmt.Sprintf(
			`func %s(%s) {
%s
		}`,
			funName,
			paramsTypeDecl,
			functionBody,
		)

	}
	return funName, funcitonDecl
}
