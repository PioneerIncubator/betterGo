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

func GetAssertType() string {
	return assertType
}

func RecordAssignVarType(fset *token.FileSet, ret *ast.AssignStmt) {
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
				assignType = getBasicLitType(expr)
			}

			fmt.Println("-- assignVar ", assignVar, " assign type ...... ", assignType)
			variableType[assignVar] = assignType
		}
		fmt.Println("[variableType] is ", variableType)
	}
	fmt.Println("---------------------")
}

func getBasicLitType(expr *ast.BasicLit) string {
	switch expr.Kind {
	case token.INT:
		return "int"
	case token.FLOAT:
		return "float64"
	case token.STRING:
		return "string"
	case token.CHAR:
		return "char"
	}
	return ""
}

func RecordDeclVarType(fset *token.FileSet, ret *ast.ValueSpec) {
	fmt.Println("---------------------")
	for i, declVar := range ret.Names {
		switch lenOfValues := len(ret.Values); lenOfValues {
		case 0:
			declVarType := reflectType(fset, ret.Type)
			fmt.Println("-- declVar ", declVar, " declare type ...... ", declVarType)
			variableType[declVar.Name] = declVarType
			fmt.Println("[variableType] is ", variableType)
		default:
			value := ret.Values[i]
			declVarType := reflectType(fset, value)
			fmt.Println("-- declVar ", declVar, " declare type ...... ", declVarType)
			if declVarType == "BasicLit" {
				declVarType = getBasicLitType(value.(*ast.BasicLit))
			}
			variableType[declVar.Name] = declVarType
			fmt.Println("[variableType] is ", variableType)
		}
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
