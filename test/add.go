package main

import (
	"fmt"

	"github.com/YongHaoWu/betterGo/enum"
)

func useAdd() {
	a, b := 1, 2
	out := enum.Add(a, b).(int)

	expect := a + b
	if expect != out {
		fmt.Printf("expected:%d, out:%d", expect, out)
	}
	fmt.Println("success, expect:", expect)
}
