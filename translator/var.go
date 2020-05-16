package translator

import (
	"fmt"
	"go/ast"
	"go/token"
)

var variableType = map[string]string{}
var assertPassCnt = 0
var assertType = ""

func RecordAssertType(input string) {
	fmt.Println("assertType is ", input)
	assertType = input
	assertPassCnt = 1
	// // gen = gen + assertType
	// fmt.Println("finally assertType is ", assertType)
}

func RecordDefineVarType(fset *token.FileSet, ret *ast.AssignStmt) {
	fmt.Println("---------------------")
	if len(ret.Lhs) == len(ret.Rhs) {
		for i, l := range ret.Lhs {
			assignVar := reflectType(fset, l)
			assignType := reflectType(fset, ret.Rhs[i])
			if assignType == "CallExpr" {
				expr := ret.Rhs[i].(*ast.CallExpr)
				if GetExprStr(fset, expr.Fun) == "make" {
					fmt.Println("[reflectType] this is make, type is ")
					switch x := expr.Args[0].(type) {
					case *ast.ArrayType:
						assignType = reflectType(fset, x.Elt)
						assignType = "[]" + assignType
					}
				}
			}
			if assignType == "BasicLit" {
				expr := ret.Rhs[i].(*ast.BasicLit)
				switch expr.Kind {
				// 12345
				case token.INT:
					assignType = "int"

					// FLOAT  // 123.45
					// IMAG   // 123.45i
					// CHAR   // 'a'
					// STRING // "abc"
					// literal_end
				}
			}

			fmt.Println("-- assignVar ", assignVar, " assign type ...... ", assignType)
			variableType[assignVar] = assignType
		}
		fmt.Println("[variableType] is ", variableType)
	}
	fmt.Println("---------------------")
}

func reflectType(fset *token.FileSet, arg interface{}) string {
	s := ""
	switch x := arg.(type) {
	case *ast.ArrayType:
		fmt.Println("[reflectType] ArrayType")
		return "[]"
	case *ast.CallExpr:
		s := GetExprStr(fset, x.Fun)
		fmt.Println("[reflectType] funName ", s, " is ast.CallExpr ")
		return "CallExpr"
	case *ast.ParenExpr:
		fmt.Println("[reflectType] ", s, " is ast.ParenExpr ")
	case *ast.FuncLit:
		// s = x.Value
		fmt.Println("[reflectType] ", s, " is ast.FuncLit ")
	case *ast.BasicLit:
		s = x.Value
		fmt.Println("[reflectType] ", s, " is ast.BasicLit ")
		return "BasicLit"
	case *ast.Ident:
		s = x.Name
		fmt.Println("[reflectType] ", s, " is ast.Ident ")
		// return "Ident"
	}
	return s
	// if s != "" {
	// 	fmt.Printf("[reflectType] :\t%s\n", s)
	// }

}
