package main

import (
	"fmt"

	"github.com/PioneerIncubator/betterGo/enum"
)

func TestMinInts() {
	origin := []int{2, 3, 4, 5}
	output := enum.Min(origin)
	if output == 2 {
		fmt.Printf("success, output is %d \n", output)
	} else {
		fmt.Printf("failed \n")
	}
}

func TestMinFloats() {
	origin := []float64{2.0, 3.0, 4.0, 5.0}
	output := enum.Min(origin)
	if output == 2.0 {
		fmt.Printf("success, output is %d \n", output)
	} else {
		fmt.Printf("failed \n")
	}
}

func main() {
	TestMinInts()
	TestMinFloats()
}