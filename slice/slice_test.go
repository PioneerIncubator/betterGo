package main

import "testing"

func TestHelloWorld(t *testing.T) {
	// t.Fatal("not implemented")
	a := []int{0, 1, 2}
	got := slice.RemoveStringElement(a, 1)
	if got != []int{0, 2} {
		t.Errorf("Abs(-1) = %v; want 1", got)
	}
}
