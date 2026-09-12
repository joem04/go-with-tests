package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Joe", "English") // what we put into function
		want := "Hello, Joe"           // expected output
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World" // empty name defaults to "World", empty language defaults to english
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Jose", "Spanish")
		want := "Hola, Jose"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Mathius", "French")
		want := "Bonjour, Mathius"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) { // created assertion function
	t.Helper() // marking as helper fucntion shows where error lays in sub test not the assesrtion fucntion
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
