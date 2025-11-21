package main

import "testing"

func TestGreet_English(t *testing.T) {
	lang := language("en")
	want := "Hello world"

	input := greet(lang)

	if input != want {
		// mark this test as failed
		t.Errorf("expected: %q, got: %q", want, input)
	}
}

func TestGreet_French(t *testing.T) {
	lang := language("fr")
	want := "Bonjour le monde"

	input := greet(lang)

	if input != want {
		// mark this test as failed
		t.Errorf("expected: %q, got: %q", want, input)
	}
}

func TestGreet_Indonesia(t *testing.T) {
	lang := language("id")
	want := ""

	input := greet(lang)

	if input != want {
		// mark this test as failed
		t.Errorf("expected: %q, got: %q", want, input)
	}
}
