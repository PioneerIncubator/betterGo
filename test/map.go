package main

import (
	"fmt"

	"github.com/YongHaoWu/betterGo/enum"
)

/*
func double(a interface{}) interface{}{
	in := reflect.ValueOf(a)
	switch in.Kind() {
	default:
		return nil
	case reflect.Int:
		return a.(int) * 2
	case reflect.Float64:
		return a.(float64) * 2
	}
}
*/

func double(a int) int {
	return a * 2
}

/*
func MapADouble(argname_1 []int, argname_2 func(int ) int) {
	lenSlice := len(argname_1)
	if lenSlice == 0 {
		return
	}
	for i := range argname_1 {
		argname_1[i] = argname_2(argname_1[i])
	}
}
*/

func main() {
	a := make([]int, 3)
	for i := range a {
		a[i] = i + 1
	}

	enum.Map(a, double)
	//realOut := make([]int, len(out))
	//for i := range out {
	//	realOut[i] = out[i].(int)
	//}
	//for i := range a {
	//	fmt.Println(a[i])
	//}
	expect := []int{2, 4, 6}

	flag := true
	for i := range a {
		if !(expect[i] == a[i]) {
			flag = false
			break
		}
	}
	if flag {
		fmt.Println("success, expect:", expect)
	} else {
		fmt.Printf("expected:%d, out:%d", expect, a)
	}
}
