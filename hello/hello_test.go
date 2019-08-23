package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello to the world", func(t *testing.T) {
		got := Hello("Varun")
		want := "Hello, Varun"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello, even when empty name passed", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

}
