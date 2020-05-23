package main

import (
	"fmt"

	"github.com/YongHaoWu/betterGo/enum"
)

func mul(a, b int) (c int) {
	c = a * b
	return
}

func testINT() {
	a, b := make([]int, 10), 12
	for i := range a {
		a[i] = i + 1
	}
	// Compute 10!
	out := enum.Reduce(a, mul, 1).(int)
	expect := 1
	for i := range a {
		expect *= a[i]
	}
	if expect != out {
		fmt.Printf("expected %d got %d , b %d", expect, out, b)
	}
	fmt.Println("success, ", expect)
}

func main() {
	testINT()

	c, d := make([]float32, 10), 12.3
	for i := range c {
		c[i] = float32(i) + 1.1
	}
	// Compute 10!
	floatOut := enum.Reduce(c, func(x, y float32) (z float32) {
		z = x * y
		return
	}, 1).(float32)
	var floatExpect float32 = 1.0
	for i := range c {
		floatExpect *= c[i]
	}
	if floatExpect != floatOut {
		fmt.Printf("expected %f got %f , b %f", floatExpect, floatOut, d)
	}
	fmt.Println("success, ", floatExpect)
	// var arrayInt = []int{1, 2, 3}
	// lambda :=
	// enum.Map(arrayInt, func(a int) int {
	// 	return a + 1
	// })
}
