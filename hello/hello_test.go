package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Varun")
	want := "Hello, Varun"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
