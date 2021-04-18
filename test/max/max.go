package main

import (
	"fmt"

	"github.com/PioneerIncubator/betterGo/enum"
)
func TestMaxInts() {
	origin := []int{2, 3, 4, 5}
	output := enum.Max(origin)
	if output == 5 {
		fmt.Printf("success, output is %d \n", output)
	} else {
		fmt.Printf("failed \n")
	}
}

func TestMaxFloats() {
	origin := []float64{2.0, 3.0, 4.0, 5.0}
	output := enum.Max(origin)
	if output == 5.0 {
		fmt.Printf("success, output is %d \n", output)
	} else {
		fmt.Printf("failed \n")
	}
}
func main() {
	TestMaxInts()
	TestMaxFloats()
}