package main

import (
	"github.com/PioneerIncubator/betterGo/enum"
)

func main() {
	//Enum.remove(["ab", "cdefg", "hijk"], func(x string) bool { return len(x) >2 }
	a := []string{"ab", "cdefg", "hijk"}
	_ = enum.Delete(a, func(x string) bool { return len(x) > 2 }).(bool)
}
