package enum

import (
	"reflect"
)

func Map(slice, anonymousFunc interface{}) {
	in := reflect.ValueOf(slice)
	if in.Kind() != reflect.Slice {
		panic("map: not slice")
	}
	n := in.Len()
	if n == 0 {
		return
	}

	elemType := in.Type().Elem()
	fn := reflect.ValueOf(anonymousFunc)
	if fn.Kind() != reflect.Func {
		str := elemType.String()
		panic("apply: function must be of type func(" + str + ", " + str + ") " + str)
	}
	var ins [1]reflect.Value
	for i := 0; i < in.Len(); i++ {
		ins[0] = in.Index(i)
		tmp := fn.Call(ins[:])[0]
		in.Index(i).Set(tmp)
	}
}
