package enum

import "reflect"

type anonymousFuncType func(element interface{}) interface{}

func InterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

func Map(slice, anonymousFunc interface{}) {
	// s := InterfaceSlice(listOfElements)
	in := reflect.ValueOf(slice)
	if in.Kind() != reflect.Slice {
		panic("map: not slice")
	}

	elemType := in.Type().Elem()
	fn := reflect.ValueOf(anonymousFunc)
	if fn.Kind() != reflect.Func {
		return
	}
	ins := []reflect.Value{}
	for i := 0; i < in.Len(); i++ {
		ins[i] = fn.Call(in.Index(i))[0]
	}
	// for i, v := range in {
	// 	in[i] = fn.Call(v)[0]
	// }
}
