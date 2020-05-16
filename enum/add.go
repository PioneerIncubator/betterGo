package enum

import (
	"fmt"
)

func Add(a, b interface{}) interface{} {

	switch typeAB := a.(type) {
	default:
		fmt.Printf("Unexpected type %T", typeAB)
		return nil
	case int:
		return a.(int) + b.(int)
	case float64:
		return a.(float64) + b.(float64)
	}
}
