package main

import (
	"fmt"

	"github.com/PioneerIncubator/betterGo/enum"
)

func main() {
	TestFindInts()
	TestFindStrings()
}

func TestFindStrings() {
	origin := []string{"ab", "cdefg", "hijk"}
	output := enum.Find(origin, func(x string) bool { return len(x) > 2 })
	if output != nil {
		fmt.Printf("success, output is %s \n", output)
	} else {
		fmt.Printf("failed \n")
	}
}

func TestFindInts() {
	origin := []int{2, 3, 4, 5}
	output := enum.Find(origin, func(x int) bool { return x%2 == 0 })
	if output != nil {
		fmt.Printf("success, output is %d \n", output)
	} else {
		fmt.Printf("failed \n")
	}
}
