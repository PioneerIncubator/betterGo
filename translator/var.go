package translator

import (
	"go/ast"
	"go/token"

	"github.com/PioneerIncubator/betterGo/types"
	log "github.com/sirupsen/logrus"
)

var variableType = map[string]string{}
var assertPassCnt = 0
var assertType = ""

func RecordAssertType(input string) {
	log.WithField("assertType", input).Info("Recording assert type")
	assertType = input
	assertPassCnt = 1
	// // gen = gen + assertType
	// fmt.Println("finally assertType is ", assertType)
}

func GetAssertType() string {
	return assertType
}

func RecordAssignVarType(fset *token.FileSet, ret *ast.AssignStmt) {
	if len(ret.Lhs) == len(ret.Rhs) {
		for i, l := range ret.Lhs {
			assignVar := reflectType(fset, l)
			assignType := reflectType(fset, ret.Rhs[i])
			if assignType == types.CallExprStr {
				expr := ret.Rhs[i].(*ast.CallExpr)
				if GetExprStr(fset, expr.Fun) == "make" {
					switch x := expr.Args[0].(type) {
					case *ast.ArrayType:
						assignType = reflectType(fset, x.Elt)
						assignType = "[]" + assignType
					}
				}
			}
			if assignType == types.BasicLitStr {
				expr := ret.Rhs[i].(*ast.BasicLit)
				assignType = GetBasicLitType(expr)
			}

			log.WithFields(log.Fields{
				"name":  assignVar,
				"value": assignType,
			}).Info("Mapping variable name and variable type")
			variableType[assignVar] = assignType
		}
	}
}

func GetBasicLitType(expr *ast.BasicLit) string {
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
	for i, declVar := range ret.Names {
		if len(ret.Values) == 0 {
			declVarType := reflectType(fset, ret.Type)
			log.WithFields(log.Fields{
				"name":  declVar,
				"value": declVarType,
			}).Info("Mapping variable name and variable type")
			variableType[declVar.Name] = declVarType
		} else {
			value := ret.Values[i]
			declVarType := reflectType(fset, value)
			log.WithFields(log.Fields{
				"name":  declVar,
				"value": declVarType,
			}).Info("Mapping variable name and variable type")
			if declVarType == types.BasicLitStr {
				declVarType = GetBasicLitType(value.(*ast.BasicLit))
			}
			variableType[declVar.Name] = declVarType
		}
	}
}

func reflectType(fset *token.FileSet, arg interface{}) string {
	s := ""
	switch x := arg.(type) {
	case *ast.ArrayType:
		return "[]"
	case *ast.CallExpr:
		return types.CallExprStr
	case *ast.ParenExpr:
	case *ast.FuncLit:
		// s = x.Value
	case *ast.BasicLit:
		s = x.Value
		return types.BasicLitStr
	case *ast.Ident:
		s = x.Name
		// return "Ident"
	}
	return s
	// if s != "" {
	// 	fmt.Printf("[reflectType] :\t%s\n", s)
	// }

}
