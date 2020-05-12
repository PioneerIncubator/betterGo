package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"

	"github.com/YongHaoWu/betterGo/translator"
	"golang.org/x/tools/go/ast/astutil"
)

func replaceOriginFunc() {
}

func genTargetFuncImplement() {

}

func isFunction() {

}

func main() {
	fset := token.NewFileSet()
	//NOTE ParseDir later
	node, err := parser.ParseFile(fset, "./test/main.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range node.Decls {
		// fmt.Println("loop node.Decls")
		// find a function declaration.
		fn, ok := f.(*ast.FuncDecl)
		if !ok {
			continue
		}
		astutil.Apply(fn, func(cr *astutil.Cursor) bool {
			n := cr.Node()
			if ret, ok := n.(*ast.GenDecl); ok {
				fmt.Println("[GenDecl] is ", ret)
			}

			if ret, ok := n.(*ast.AssignStmt); ok {
				if ret.Tok == token.DEFINE {
					// a := 12
					translator.RecordDefineVarType(fset, ret)
				}
			}

			if ret, ok := n.(*ast.FuncDecl); ok {
				if ret.Name.Name != "main" {
					fmt.Println("find fucntion declar  ", ret.Name.Name)
					translator.GetFuncType(fset, ret)
				}
			}

			if ret, ok := n.(*ast.TypeAssertExpr); ok {
				//TODO: expr lik out := enum.Reduce(a, mul, 1).(int)
				// Assert is parse before function call
				// which means we 'll parse (int) then enum.Reduce
				assertType := translator.GetExprStr(fset, ret.Type)
				translator.RecordAssertType(assertType)
				return true
			}
			// call expr, find enum functions
			if ret, ok := n.(*ast.CallExpr); ok {
				funName := translator.GetExprStr(fset, ret.Fun)
				// fmt.Println("[CallExpr] funName", funName)
				if strings.Contains(funName, "enum") {
					newFunName, funDeclStr := translator.GenEnumFunctionDecl(funName, ret.Args)
					fmt.Println("[CallExpr] newfunName", newFunName)
					fmt.Println("gen funDeclStr:  ", funDeclStr)
				}

				// try rewrite the reduce function call
				/*TODO: useless
				switch x := ret.Fun.(type) {
				case *ast.Ident:
					x.Name = "targetpkg." + newFunName
					ret.Fun = x
				case *ast.SelectorExpr:
					x.X.(*ast.Ident).Name = "targetpkg"
					x.Sel.Name = newFunName
					fmt.Println("my..................", x.Sel.Name)
					ret.Fun = x
				}
				cr.Replace(ret)
				 if err := format.Node(os.Stdout, token.NewFileSet(), n); err != nil {
				 	log.Fatalln("Error:", err)
				 }
				*/
				fmt.Println("end=-=================================")
				return true
			}
			return true
		}, nil)

	}

}
