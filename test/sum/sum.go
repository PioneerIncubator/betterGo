package main

import (
	"fmt"

	"github.com/PioneerIncubator/betterGo/enum"
)

func TestSumInts() {
	origin := []int{1, 2, 3, 4}
	flag := enum.Sum(origin)
	if flag == 10 {
		fmt.Println("success")
	} else {
		fmt.Printf("failed")
	}
}

func TestSumFloats() {
	origin := []int{1.0, 2.0, 3.0, 4.0}
	flag := enum.Sum(origin)
	if flag == 10.0 {
		fmt.Println("success")
	} else {
		fmt.Printf("failed")
	}
}

func main() {
	TestSumInts()
	TestSumFloats()
}