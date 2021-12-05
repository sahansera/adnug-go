package main

import "testing"

func TestAdd(t *testing.T) {
	x, y := 1, 1
	want := 2
	got := add(x, y)
	if got != want {
		t.Fatalf("want %d, but got %d", want, got)
	}
}
