package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"

	"github.com/YongHaoWu/betterGo/translator"
	"github.com/YongHaoWu/betterGo/file_operations"
	"github.com/urfave/cli/v2"
	"golang.org/x/tools/go/ast/astutil"
)

// func replaceOriginFunc() {
// }

//func genTargetFuncImplement() {
//
//}

// func isFunction() {

// }

func loopASTNode(fset *token.FileSet, node *ast.File) {
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
					fmt.Println("find function declar  ", ret.Name.Name)
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

					s := strings.Split(funName, ".")
					filePath := "./utils/" + s[0]
					fileName := s[1] + ".go"
					tmpStr := "\n" + funDeclStr
					buffer := []byte(tmpStr)
					packageName := "package " + s[0]
					file_operations.WriteToFile(filePath + "/" + fileName, packageName, buffer)
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

func loopASTFile(filePath string) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	loopASTNode(fset, node)
}

func loopASTDir(filePath string) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		fmt.Println("parse dir fail", filePath)
		log.Fatal(err)
	}
	for k, v := range pkgs {
		fmt.Println("pkg k is ", k)
		for filename, fileNode := range v.Files {
			fmt.Println("filename  is ", filename)
			loopASTNode(fset, fileNode)
		}
	}

}

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "file",
				Aliases: []string{"f"},
				Usage:   "Generate and replace the file with Enum files",
			},
			&cli.StringFlag{
				Name:    "dir",
				Aliases: []string{"d"},
				Usage:   "Generate and replace the dirctory with Enum files",
			},
		},
		Action: func(c *cli.Context) error {
			if c.String("file") != "" {
				loopASTFile(c.String("file"))
				return nil
			}
			if c.String("dir") != "" {
				loopASTDir(c.String("dir"))
				return nil
			}

			log.Fatal("file or dir flag empty")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
