package translator

import (
	"go/ast"
	"go/token"

	log "github.com/sirupsen/logrus"
)

func getFunParamListRawStr(fset *token.FileSet, ret *ast.FuncDecl) string {
	paramsStr := "func ("
	for _, v := range ret.Type.Params.List {
		exprType := GetExprStr(fset, v.Type)
		for j := 0; j < len(v.Names); j++ {
			// func mul(a, b int) int
			paramsStr = paramsStr + exprType + ", "
		}
	}
	paramsStr = paramsStr[:len(paramsStr)-2] + " )"
	return paramsStr
}

func getFunRetListRawStr(fset *token.FileSet, ret *ast.FuncDecl) string {
	retStr := ""
	// NOTE (ret1 int,ret2 double )
	if ret.Type.Results == nil { // could be no return value
		return retStr
	}
	for _, v := range ret.Type.Results.List {
		// fmt.Println("[getFunRetListRawStr] result v", v)
		exprType := GetExprStr(fset, v.Type)
		if len(v.Names) == 0 {
			retStr = exprType
		} else {
			for j := 0; j < len(v.Names); j++ {
				retStr += exprType
			}
		}
	}
	return retStr
}

func DecorateParamName(name string) string {
	return "BETTERGOPARAM" + name
}

func recordParamType(fset *token.FileSet, ret *ast.FuncType) {
	for _, v := range ret.Params.List {
		exprType := GetExprStr(fset, v.Type)
		for _, name := range v.Names {
			nameStr := GetExprStr(fset, name)
			log.WithFields(log.Fields{
				"name":   nameStr,
				"record": DecorateParamName(nameStr),
				"value":  exprType,
			}).Info("Mapping param name and param type")
			variableType[DecorateParamName(nameStr)] = exprType
		}
	}

}

func GetFuncType(fset *token.FileSet, ret *ast.FuncDecl) (string, string) {
	paramsStr := getFunParamListRawStr(fset, ret)
	retStr := getFunRetListRawStr(fset, ret)
	recordParamType(fset, ret.Type)

	log.WithFields(log.Fields{
		"name":  ret.Name.Name,
		"value": paramsStr + retStr,
	}).Info("Mapping param name and param type")
	variableType[ret.Name.Name] = paramsStr + retStr
	return paramsStr + retStr, retStr
	// fmt.Println("[FuncDecl] Type.Results", ret.Type.Results.List)
	// if ret.Tok == token.DEFINE {         Results
	// 	recordDefineVarType(fset, ret)
	// }
}
