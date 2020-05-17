package enum

import (
	"reflect"
)

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
	n := in.Len()
	if n == 0 {
		return
	}

	elemType := in.Type().Elem()
	fn := reflect.ValueOf(anonymousFunc)
	//if !goodFunc(fn, elemType, elemType, elemType) {
	//	 str := elemType.String()
	//	 panic("apply: function must be of type func(" + str + ", " + str + ") " + str)
	//}
	if fn.Kind() != reflect.Func {
		str := elemType.String()
		panic("apply: function must be of type func(" + str + ", " + str + ") " + str)
	}
	var ins [1]reflect.Value
	for i := 0; i < in.Len(); i++ {
		ins[0] = in.Index(i)
		tmp := fn.Call(ins[:])[0]
	}
	// for i, v := range in {
	// 	in[i] = fn.Call(v)[0]
	// }
}

/*
func Map(slice, anonymousFunc interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = anonymousFunc(s.Index(i))
	}

	return ret
}
*/

/*
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

func Map(listOfElements interface{}, anonymousFunc func(element interface{}) interface{}) {
	s := InterfaceSlice(listOfElements)
	for i, v := range s {
		s[i] = anonymousFunc(v)
	}
}
*/
