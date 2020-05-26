package main

import (
	"fmt"

	"github.com/PioneerIncubator/betterGo/enum"
)

func randomFn(origin int) int {
	origin += 1
	var b = 10
	return origin * b
}

func varTestFn(origin int) int {
	a := 10
	return a * origin
}

func testMap(origin, expect []int, fn func(int) int) {
	enum.Map(origin, fn)
	flag := true
	for i := range origin {
		if !(expect[i] == origin[i]) {
			flag = false
			break
		}
	}
	if flag {
		fmt.Println("success, expect:", expect)
	} else {
		fmt.Printf("expected:%d, out:%d", expect, origin)
	}

}

func main() {
	origin := []int{2, 4, 6}
	expect := []int{30, 50, 70}
	testMap(origin, expect, randomFn)

	origin = []int{2}
	expect = []int{20}
	testMap(origin, expect, varTestFn)
}
