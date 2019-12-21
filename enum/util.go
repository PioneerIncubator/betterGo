package enum

import "reflect"

// goodFunc verifies that the function satisfies the signature, represented as a slice of types.
// The last type is the single result type; the others are the input types.
// A final type of nil means any result type is accepted.
func goodFunc(fn reflect.Value, types ...reflect.Type) bool {
	if fn.Kind() != reflect.Func {
		return false
	}
	// Last type is return, the rest are ins.
	if fn.Type().NumIn() != len(types)-1 || fn.Type().NumOut() != 1 {
		return false
	}
	for i := 0; i < len(types)-1; i++ {
		if fn.Type().In(i) != types[i] {
			return false
		}
	}
	outType := types[len(types)-1]
	if outType != nil && fn.Type().Out(0) != outType {
		return false
	}
	return true
}
