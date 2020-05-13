package main

/*
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

	// var arrayInt = []int{1, 2, 3}
	// lambda :=
	// enum.Map(arrayInt, func(a int) int {
	// 	return a + 1
	// })
}

*/
func Add(a, b int) int {
	return a + b
}

func main() {
	a, b := 1, 2
	out := Add(a, b)

	expect := a + b
	if expect != out {
		fmt.Printf("expected:%d, out:%d", expect, out)
	}
	fmt.Println("success, expect:", expect)
}
