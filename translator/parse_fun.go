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
	//NOTE (ret1 int,ret2 double )
	for _, v := range ret.Type.Results.List {
		// fmt.Println("[getFunRetListRawStr] result v", v)
		exprType := GetExprStr(fset, v.Type)
		if len(v.Names) == 0 {
			retStr = exprType
		} else {
			for j := 0; j < len(v.Names); j++ {
				fmt.Println("[getFunRetListRawStr] retStr", retStr)
				retStr = retStr + exprType
			}
		}
	}
	fmt.Println("[getFunRetListRawStr] retStr", retStr)
	return retStr
}

func GetFuncType(fset *token.FileSet, ret *ast.FuncDecl) (string, string) {
	paramsStr := getFunParamListRawStr(fset, ret)
	retStr := getFunRetListRawStr(fset, ret)

	fmt.Println("[GetFuncType] record ", ret.Name.Name, " func origin type is ", paramsStr+retStr)
	variableType[ret.Name.Name] = paramsStr + retStr
	return paramsStr + retStr, retStr
	// fmt.Println("[FuncDecl] Type.Results", ret.Type.Results.List)
	// if ret.Tok == token.DEFINE {         Results
	// 	recordDefineVarType(fset, ret)
	// }
}
