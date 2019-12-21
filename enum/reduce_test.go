package enum

import (
	"testing"
)

func mul(a, b int) int {
	return a * b
}

func TestReduce(t *testing.T) {
	a := make([]int, 10)
	for i := range a {
		a[i] = i + 1
	}
	// Compute 10!
	out := Reduce(a, mul, 1).(int)
	expect := 1
	for i := range a {
		expect *= a[i]
	}
	if expect != out {
		t.Fatalf("expected %d got %d", expect, out)
	}
}
