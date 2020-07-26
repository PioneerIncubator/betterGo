package enum

import (
	"reflect"
)

func Delete(slice, anonymousFunc interface{}) interface{} {
	in := reflect.ValueOf(slice)
	if in.Kind() != reflect.Slice {
		panic("delete: not slice")
	}
	n := in.Len()
	if n == 0 {
		return nil
	}

	// Get slice element's type
	elemType := in.Type().Elem()
	fn := reflect.ValueOf(anonymousFunc)
	if fn.Kind() != reflect.Func {
		str := elemType.String()
		panic("apply: function must be of type func(" + str + ", " + str + ")" + str)
	}

	var ins [1]reflect.Value

	count := 0
	for i := 0; i < n; i++ {
		ins[0] = in.Index(i)
		if fn.Call(ins[:])[0].Bool() {
			in.Index(count).Set(reflect.ValueOf(ins[i]))
			count++
		}
	}
	return in[:count].Interface()
}
