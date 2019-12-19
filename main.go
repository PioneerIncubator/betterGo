package main

import (
	"fmt"

	"github.com/YongHaoWu/betterGo/enum"
)

func main() {
	a := []int{1, 2, 3}
	fmt.Println(a)
	enum.Map(a, func(x int) int {
		x = x + 100
		return x
	})
	fmt.Println(a)
}
