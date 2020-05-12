package main

import (
	"fmt"

	"github.com/YongHaoWu/betterGo/enum"
)

func mul(a, b int) (c int) {
	c = a * b
	return
}

func main() {
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
