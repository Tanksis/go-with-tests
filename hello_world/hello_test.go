package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chris") //reuse values
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want) //f strings, using .Errorf
	}
}

//test files need to be in a file with name xxx_test.go
//test function must start with word Test
//test function takes one argument: t *testing.T
//you need to import "testing"
