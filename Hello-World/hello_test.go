package main

import "testing"

/**
Writing tests
Writing a test is just like writing a function, with a few rules
It needs to be in a file with a name like xxx_test.go
The test function must start with the word Test
The test function takes one argument only t *testing.T
**/
func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		//needed to tell the test suite that this method is a helper
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris") //We're declaring some variables with the syntax varName := value, which lets us re-use
		//some values in our test for readability.
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

}
