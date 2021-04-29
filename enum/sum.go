package enum

import (
	"reflect"

	log "github.com/sirupsen/logrus"
)

func Sum(slice interface{}) interface{} {
	in := reflect.ValueOf(slice)
	if in.Kind() != reflect.Slice {
		log.Fatal("Input is not slice")
	}
	n := in.Len()
	if n == 0 {
		return nil
	}

	switch sliceType := slice.(type) {
	default:
		log.WithFields(log.Fields{
			"type": sliceType,
		}).Fatal("Unexpected type!")
		return nil

	//reflect.value only return int64 and float64
	case []int:
		var sum int64
		sum = 0
		for i := 0; i < n; i++ {
			sum += in.Index(i).Int()
		}
		return sum
	case []float64:
		var sum float64
		sum = 0.0
		for i := 0; i < n; i++ {
			sum += in.Index(i).Float()
		}
		return sum
	}

}
