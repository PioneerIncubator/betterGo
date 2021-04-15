package translator

import (
	"fmt"
	"go/ast"
	"go/token"
)

func getFunParamListRawStr(fset *token.FileSet, ret *ast.FuncDecl) string {
	paramsStr := "func ("
	for i, v := range ret.Type.Params.List {
		fmt.Println("[getFunParamListRawStr] loop ret.Type.Params.List", i)
		exprType := GetExprStr(fset, v.Type)
		for j := 0; j < len(v.Names); j++ {
			fmt.Println("[getFunParamListRawStr] compose paramStr", paramsStr)
			// func mul(a, b int) int
			paramsStr = paramsStr + exprType + ", "
		}
	}
	paramsStr = paramsStr[:len(paramsStr)-2] + " )"
	fmt.Println("[getFunParamListRawStr] final paramStr", paramsStr)
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
				fmt.Println("[getFunRetListRawStr] retStr", retStr)
				retStr += exprType
			}
		}
	}
	fmt.Println("[getFunRetListRawStr] retStr", retStr)
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
			fmt.Printf("[recordParamType] param name %s record as %s, value %s\n", nameStr, DecorateParamName(nameStr), exprType)
			variableType[DecorateParamName(nameStr)] = exprType
		}
	}

}

func GetFuncType(fset *token.FileSet, ret *ast.FuncDecl) (string, string) {
	paramsStr := getFunParamListRawStr(fset, ret)
	retStr := getFunRetListRawStr(fset, ret)
	recordParamType(fset, ret.Type)

	fmt.Println("[GetFuncType] record ", ret.Name.Name, " func origin type is ", paramsStr+retStr)
	variableType[ret.Name.Name] = paramsStr + retStr
	return paramsStr + retStr, retStr
	// fmt.Println("[FuncDecl] Type.Results", ret.Type.Results.List)
	// if ret.Tok == token.DEFINE {         Results
	// 	recordDefineVarType(fset, ret)
	// }
}
