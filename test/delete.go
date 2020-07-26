package main

import (
	"fmt"

	"github.com/PioneerIncubator/betterGo/enum"
)

func main() {
	TestDeleteInts()
	TestDeleteStrings()
}

func TestDeleteStrings() {
	origin := []string{"ab", "cdefg", "hijk"}
	flag := enum.Delete(origin, func(x string) bool { return len(x) > 2 })
	if flag {
		fmt.Println("success")
	} else {
		fmt.Printf("failed")
	}
}

func TestDeleteInts() {
	origin := []int{2, 3, 4, 5}
	flag := enum.Delete(origin, func(x int) bool { return x%2 == 0 })
	if flag {
		fmt.Println("success")
	} else {
		fmt.Printf("failed")
	}
}
