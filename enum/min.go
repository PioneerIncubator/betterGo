package enum

import(
	"reflect"
	"fmt"
)

func Min(slice interface{}) interface{} {
	in := reflect.ValueOf(slice)
	if in.Kind() != reflect.Slice {
		panic("Min: not slice")
	}
	n := in.Len()
	if n == 0 {
		return nil
	}

	switch sliceType := slice.(type) {
	default:
		fmt.Printf("Unexpected type %T", sliceType)
		return nil
	
	//reflect.value only return int64 and float64
	case []int:
		var minVal int64
		minVal = in.Index(0).Int()
		for i := 1; i < n; i++ {
			if(in.Index(i).Int() < minVal) {
				minVal = in.Index(i).Int()
			}
		}
		return minVal
	
	case []float64:
		var minVal float64
		minVal = in.Index(0).Float()
		for i := 1; i < n; i++ {
			if(in.Index(i).Float() < minVal) {
				minVal = in.Index(i).Float()
			}
		}
		return minVal
}
}