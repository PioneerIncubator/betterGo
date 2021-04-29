package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"

	"github.com/PioneerIncubator/betterGo/fileoperations"
	log "github.com/sirupsen/logrus"

	"github.com/PioneerIncubator/betterGo/translator"
	"github.com/urfave/cli/v2"
	"golang.org/x/tools/go/ast/astutil"
)

func replaceOriginFunc(fset *token.FileSet, ret *ast.CallExpr, callFunExpr, newFunName, filePath string, isDir bool) {
	s := strings.Split(callFunExpr, ".")
	pkgName := s[0]
	newFunName = fmt.Sprintf("%s.%s", pkgName, newFunName)
	_, args, _ := translator.ExtractParamsTypeAndName(fset, ret.Args)

	originStr := fileoperations.GenCallExpr(callFunExpr, translator.GetAssertType(), args, false)
	targetStr := fileoperations.GenCallExpr(newFunName, translator.GetAssertType(), args, true)

	filePath = fmt.Sprintf("./%s", filePath)
	if !isDir {
		fileoperations.ReplaceOriginFuncByFile(filePath, originStr, targetStr)
	} else {
		fileoperations.ReplaceOriginFuncByDir(filePath, originStr, targetStr)
	}
}

func genTargetFuncImplement(fset *token.FileSet, ret *ast.CallExpr, callFunExpr, funDeclStr string) (bool, string) {
	s := strings.Split(callFunExpr, ".")
	pkgName := s[0]
	funName := s[1]
	genFilePath := fmt.Sprintf("./utils/%s", pkgName)
	genFileName := fmt.Sprintf("%s.go", funName)
	genFileName = strings.ToLower(genFileName)
	filePath := fmt.Sprintf("%s/%s", genFilePath, genFileName)

	_, _, listOfArgTypes := translator.ExtractParamsTypeAndName(fset, ret.Args)
	funcExists, previousFuncName := fileoperations.CheckFuncExists(filePath, listOfArgTypes)
	if funcExists {
		return true, previousFuncName
	}

	buffer := []byte(fmt.Sprintf("\n%s", funDeclStr))
	pkgStatement := fmt.Sprintf("package %s", pkgName)
	err := fileoperations.WriteFuncToFile(filePath, pkgStatement, buffer)
	if err != nil {
		log.Fatal(err)
	}

	return false, previousFuncName
}

// func isFunction() {

// }

func loopASTNode(fset *token.FileSet, node *ast.File, filePath string, isDir, rewriteAndGen bool) {
	for _, f := range node.Decls {
		// fmt.Println("loop node.Decls")
		// find a function declaration.
		fn, ok := f.(*ast.FuncDecl)
		if !ok {
			continue
		}
		astutil.Apply(fn, func(cr *astutil.Cursor) bool {
			n := cr.Node()
			switch ret := n.(type) {
			case *ast.GenDecl:
				log.WithFields(log.Fields{
					"astNodeType": ret,
				}).Info("Find a general declaration")
			case *ast.ValueSpec:
				log.WithFields(log.Fields{
					"astNodeType": ret,
				}).Info("Find a constant or variable declaration")
				translator.RecordDeclVarType(fset, ret)
			case *ast.AssignStmt:
				log.WithFields(log.Fields{
					"astNodeType": ret,
				}).Info("Find an assign statement")
				if ret.Tok == token.DEFINE { // a := 12
					translator.RecordAssignVarType(fset, ret)
				}
			case *ast.FuncDecl:
				if ret.Name.Name != "main" {
					log.WithFields(log.Fields{
						"astNodeType": ret,
						"funcName":    ret.Name.Name,
					}).Info("Find a function declaration")
					translator.GetFuncType(fset, ret)
				}
			case *ast.TypeAssertExpr:
				//TODO: expr lik out := enum.Reduce(a, mul, 1).(int)
				// Assert is parse before function call
				// which means we 'll parse (int) then enum.Reduce
				log.WithFields(log.Fields{
					"astNodeType": ret,
				}).Info("Find a type assert expression")
				assertType := translator.GetExprStr(fset, ret.Type)
				translator.RecordAssertType(assertType)
			case *ast.CallExpr:
				funName := translator.GetExprStr(fset, ret.Fun)
				// fmt.Println("[CallExpr] funName", funName)
				if strings.Contains(funName, "enum") {
					newFunName, funDeclStr := translator.GenEnumFunctionDecl(fset, funName, ret.Args)
					log.WithFields(log.Fields{
						"astNodeType":    ret,
						"newFunName":     newFunName,
						"newFunDeclStmt": funDeclStr,
					}).Info("Find a call expression")

					if rewriteAndGen {
						// Generate function to file
						funcExists, prevFuncName := genTargetFuncImplement(fset, ret, funName, funDeclStr)

						// Replace origin function call expression
						if funcExists {
							replaceOriginFunc(fset, ret, funName, prevFuncName, filePath, isDir)
						} else {
							replaceOriginFunc(fset, ret, funName, newFunName, filePath, isDir)
						}
					}
				}
			}
			return true
		}, nil)
	}
}

func loopASTFile(filePath string, rewriteAndGen bool) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		log.WithFields(log.Fields{
			"filePath": filePath,
			"err":      err,
		}).Fatal("Parse file fail")
	}
	loopASTNode(fset, node, filePath, false, rewriteAndGen)
}

func loopASTDir(filePath string, rewriteAndGen bool) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		log.WithFields(log.Fields{
			"dirPath": filePath,
			"err":     err,
		}).Fatal("Parse dir fail")
	}
	for _, v := range pkgs {
		for _, fileNode := range v.Files {
			loopASTNode(fset, fileNode, filePath, true, rewriteAndGen)
		}
	}

}

func main() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

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
			&cli.BoolFlag{
				Name:    "rewrite&gen",
				Aliases: []string{"w"},
				Value:   false,
				Usage:   "Rewrite files and generate files",
			},
		},
		Action: func(c *cli.Context) error {
			rewriteAndGen := c.Bool("rewrite&gen")
			if c.String("file") != "" {
				loopASTFile(c.String("file"), rewriteAndGen)
				return nil
			}
			if c.String("dir") != "" {
				loopASTDir(c.String("dir"), rewriteAndGen)
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
