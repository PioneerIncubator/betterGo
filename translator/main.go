package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
)

func getExprStr(fset *token.FileSet, expr interface{}) string {
	name := new(bytes.Buffer)
	printer.Fprint(name, fset, expr)
	return name.String()
}

//get function params
// for _, v := range ret.Type.Params.List {
// 	for _, ident := range v.Names {
// 		fmt.Println("[FuncDecl] name", ident.Name)
// 		paramsStr = paramsStr + ident.Name + ","
// 	}
// 	fmt.Println("[FuncDecl] Type.Params v", v)
// 	exprType := getExprStr(fset, v.Type)
// 	fmt.Println("[FuncDecl] Type.Params v.Type", v.Type, "exprType", exprType)
// 	paramsStr = paramsStr + exprType + ", "
// }

func reflectType(fset *token.FileSet, arg interface{}) string {
	s := ""
	switch x := arg.(type) {
	case *ast.ArrayType:
		fmt.Println("[reflectType] ArrayType")
		return "[]"
	case *ast.CallExpr:
		s := getExprStr(fset, x.Fun)
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

var variableType = map[string]string{}

func recordDefineVarType(fset *token.FileSet, ret *ast.AssignStmt) {
	fmt.Println("---------------------")
	// for _, l := range ret.Lhs {
	// 	assignVar := reflectType(fset, l)
	// 	fmt.Println("[GenDecl] assignVar is ", assignVar)
	// }
	// for _, r := range ret.Rhs {
	// 	assignType := reflectType(fset, r)
	// 	if assignType == "make" {
	// 		fmt.Println("[reflectType] this is make, type is ")
	// 		assignType = reflectType(fset, r.(*ast.CallExpr).Args[0])
	// 	}
	// 	fmt.Println("[GenDecl] AssignType is ", assignType)
	// }

	if len(ret.Lhs) == len(ret.Rhs) {
		for i, l := range ret.Lhs {
			assignVar := reflectType(fset, l)
			assignType := reflectType(fset, ret.Rhs[i])
			if assignType == "CallExpr" {
				expr := ret.Rhs[i].(*ast.CallExpr)
				if getExprStr(fset, expr.Fun) == "make" {
					fmt.Println("[reflectType] this is make, type is ")
					switch x := expr.Args[0].(type) {
					case *ast.ArrayType:
						assignType = reflectType(fset, x.Elt)
						// fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!", assignType)
						assignType = "[]" + assignType
					}
				}
			}
			if assignType == "BasicLit" {
				expr := ret.Rhs[i].(*ast.BasicLit)
				switch expr.Kind {
				// 12345
				case token.INT:
					assignType = "INT"

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
	// fmt.Println("[AssignStmt] is ", ret)
	// fmt.Println("[AssignStmt] is tok ", ret.Tok)

}

func getFuncType(fset *token.FileSet, ret *ast.FuncDecl) string {
	paramsStr := "func ("
	for i, v := range ret.Type.Params.List {
		exprType := getExprStr(fset, v.Type)
		for j := 0; j < len(v.Names); j++ {
			fmt.Println("[FuncDecl] paramStr", paramsStr)
			if i+1 == len(ret.Type.Params.List) {
				paramsStr = paramsStr + exprType + ")"
			} else {
				paramsStr = paramsStr + exprType + ", "
			}
		}
	}
	fmt.Println("[FuncDecl] paramStr", paramsStr)

	resultStr := ""
	//NOTE (ret1 int,ret2 double )
	for _, v := range ret.Type.Results.List {
		// fmt.Println("[FuncDecl] result v", v)
		exprType := getExprStr(fset, v.Type)
		if len(v.Names) == 0 {
			resultStr = exprType
		} else {
			for j := 0; j < len(v.Names); j++ {
				fmt.Println("[FuncDecl] resultStr", resultStr)
				resultStr = resultStr + exprType + ", "
			}
		}
	}
	fmt.Println("[FuncDecl] resultStr", resultStr)
	fmt.Println("[FuncDecl] func origin type is ", paramsStr+resultStr)
	return paramsStr + resultStr
	// fmt.Println("[FuncDecl] Type.Results", ret.Type.Results.List)
	// if ret.Tok == token.DEFINE {         Results
	// 	recordDefineVarType(fset, ret)
	// }
}

func generateEnumReduce() {

}

var gen string

func main() {
	fset := token.NewFileSet()
	//NOTE ParseDir later
	node, err := parser.ParseFile(fset, "./main.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Functions:")
	for _, f := range node.Decls {
		fn, ok := f.(*ast.FuncDecl)
		if !ok {
			continue
		}
		ast.Inspect(fn, func(n ast.Node) bool {
			if ret, ok := n.(*ast.GenDecl); ok {
				fmt.Println("[GenDecl] is ", ret)
			}

			if ret, ok := n.(*ast.AssignStmt); ok {
				if ret.Tok == token.DEFINE {
					recordDefineVarType(fset, ret)
				}
			}

			if ret, ok := n.(*ast.FuncDecl); ok {
				if ret.Name.Name != "main" {
					variableType[ret.Name.Name] = getFuncType(fset, ret)
				}
			}

			// call expr, find enum functions
			if ret, ok := n.(*ast.CallExpr); ok {
				funName := getExprStr(fset, ret.Fun)
				fmt.Println("[CallExpr] funName", funName)
				switch funName {
				case "enum.Reduce":
					// iterate function args to reveal the type
					//Reduce(slice, pairFunction, zero interface{}) interface{}
					gen = "func Reduce("
					for _, arg := range ret.Args {
						switch x := arg.(type) {
						case *ast.BasicLit:
							switch x.Kind {
							case token.INT:
								gen = gen + "int"
							}
						case *ast.Ident:
							argVarName := x.Name
							fmt.Println("argVarName is ", argVarName)
							fmt.Println("argVarType is ", variableType[argVarName])
							gen = gen + variableType[argVarName] + ","
						}
						// fmt.Println("args is ", arg)
						// reflectType(fset, arg)
					}
					gen = gen + ")"
					fmt.Println("gen is ", gen)
					return true
				}
			}
			if ret, ok := n.(*ast.TypeAssertExpr); ok {
				assertType := getExprStr(fset, ret.Type)
				fmt.Println("assertType is ", assertType)
				gen = gen + assertType
				fmt.Println("finally gen is ", gen)
				return true
			}
			return true
		})
	}

}
