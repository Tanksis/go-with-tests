package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) { //subtests
		got := Hello("Chris", "") //reuse values
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Phil", "French")
		want := "Bonjour, Phil"
		assertCorrectMessage(t, got, want)
	})
}

//code to refactor and clean up functions above
func assertCorrectMessage(t testing.TB, got, want string) { //since got, want, are both strings we can write string once
	//t.Helper allows the test suite to know this is a helper function, so when fails occur
	//it reports in the respective function call as opposed to the test helper
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want) //f string
	}
}

//test files need to be in a file with name xxx_test.go
//test function must start with word Test
//test function takes one argument: t *testing.T
//you need to import "testing"
