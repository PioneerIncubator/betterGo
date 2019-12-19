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

func Map(listOfElements interface{}, anonymousFunc anonymousFuncType) {
	s := InterfaceSlice(listOfElements)
	for i, v := range s {
		s[i] = anonymousFunc(v)
	}
}
