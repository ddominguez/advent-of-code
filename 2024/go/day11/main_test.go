package main

import "testing"

func TestRun(t *testing.T) {
	input := "125 17"
	count := 6
	expected := 22
	result := run(input, count)
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}
