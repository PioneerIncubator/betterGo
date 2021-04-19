package enum

import(
	"reflect"
	"fmt"
)

func Max(slice interface{}) interface{} {
	in := reflect.ValueOf(slice)
	if in.Kind() != reflect.Slice {
		panic("Max: not slice")
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
		var maxVal int64
		maxVal = in.Index(0).Int()
		for i := 1; i < n; i++ {
			if(in.Index(i).Int() > maxVal) {
				maxVal = in.Index(i).Int()
			}
		}
		return maxVal
	
	case []float64:
		var maxVal float64
		maxVal = in.Index(0).Float()
		for i := 1; i < n; i++ {
			if(in.Index(i).Float() > maxVal) {
				maxVal = in.Index(i).Float()
			}
		}
		return maxVal
}
}