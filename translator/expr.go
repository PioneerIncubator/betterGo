package translator

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"

	"github.com/PioneerIncubator/betterGo/utils"
	log "github.com/sirupsen/logrus"
)

func ExtractParamsTypeAndName(fset *token.FileSet, listOfArgs []ast.Expr) (string, []string, []string) {
	var paramsType string
	var listOfArgVarNames []string
	var listOfArgTypes []string
	argname := "argname"
	argsNum := len(listOfArgs)
	for i, arg := range listOfArgs {
		argname = utils.IncrementString(argname, "", 1)
		switch x := arg.(type) {
		case *ast.BasicLit:
			argVarName := x.Value
			listOfArgVarNames = append(listOfArgVarNames, argVarName)
			listOfArgTypes = append(listOfArgTypes, GetBasicLitType(x))
			paramsType = fmt.Sprintf("%s %s %s", paramsType, argname, GetBasicLitType(x))
		case *ast.Ident:
			argVarName := x.Name
			listOfArgVarNames = append(listOfArgVarNames, argVarName)
			listOfArgTypes = append(listOfArgTypes, variableType[argVarName])
			var argVarType string
			if paramType, ok := variableType[DecorateParamName(argVarName)]; ok {
				argVarType = paramType
			} else {
				argVarType = variableType[argVarName]
			}
			log.WithFields(log.Fields{
				"name": argVarName,
				"type": argVarType,
			}).Info("Find an identifier from function parameters")
			paramsType = fmt.Sprintf("%s %s %s", paramsType, argname, argVarType)
		case *ast.FuncLit:
			argDeclar, retDeclar := "", ""
			for _, v := range x.Type.Params.List {
				lenNames := len(v.Names)
				if argDeclar == "" {
					lenNames--
					argDeclar = GetExprStr(fset, v.Type)
				}
				for i := 0; i < lenNames; i++ {
					argDeclar = fmt.Sprintf("%s, %s", argDeclar, GetExprStr(fset, v.Type))
				}
			}
			for _, v := range x.Type.Results.List {
				lenNames := len(v.Names)
				if retDeclar == "" {
					lenNames--
					retDeclar = GetExprStr(fset, v.Type)
				}
				for i := 0; i < lenNames; i++ {
					retDeclar = fmt.Sprintf("%s,%s", retDeclar, GetExprStr(fset, v.Type))
				}
			}

			var lambdaTypeStr string
			if len(x.Type.Results.List) == 1 {
				lambdaTypeStr = fmt.Sprintf("func(%s) %s", argDeclar, retDeclar)
			} else {
				lambdaTypeStr = fmt.Sprintf("func(%s)(%s)", argDeclar, retDeclar)
			}
			paramsType = fmt.Sprintf("%s %s %s", paramsType, argname, lambdaTypeStr)
			listOfArgVarNames = append(listOfArgVarNames, "lambda")
			listOfArgTypes = append(listOfArgTypes, lambdaTypeStr)
		default:
			log.Error("Unknown type: ", x)
		}

		if i != argsNum-1 {
			paramsType += ","
		}
	}
	listOfArgTypes = append(listOfArgTypes, assertType)
	return paramsType, listOfArgVarNames, listOfArgTypes
}

func GetExprStr(fset *token.FileSet, expr interface{}) string {
	name := new(bytes.Buffer)
	err := printer.Fprint(name, fset, expr)
	if err != nil {
		log.Fatal(err)
	}
	return name.String()
}
