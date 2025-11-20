package main

import "testing"

func TestGreet(t *testing.T) {
	want := "Hello world"

	input := greet()

	if input != want {
		// mark this test as failed
		t.Errorf("expected: %q, got: %q", want, input)
	}
}
