package main

import (
	"testing"
)

func add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	a, b := 1, 2
	// Compute 1 + 2
	out := add(a, b)
	expect := a + b
	if expect != out {
		//fmt.Printf("expected:%d, out:%d", expect, out)
		t.Fatalf("expected %d got %d", expect, out)
	}
}
