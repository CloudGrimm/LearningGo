package main

import "testing"

/**
Writing tests
Writing a test is just like writing a function, with a few rules
It needs to be in a file with a name like xxx_test.go
The test function must start with the word Test
The test function takes one argument only t *testing.T
TDD is a skill that needs practice to develop but by being able to
break problems down into smaller components that you can test you will have a much easier time writing software.
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
		got := Hello("Chris", "") //We're declaring some variables with the syntax varName := value, which lets us re-use
		//some values in our test for readability.
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
		got := Hello("Chris", "French")
		want := "Bonjour, Chris"
		assertCorrectMessage(t, got, want)
	})
	/***
		Let's go over the cycle again
		1. Write a test
		2. Make the compiler pass
		3. Run the test, see that it fails and check the error message is meaningful
		4. Write enough code to make the test pass
		5. Refactor
	On the face of it this may seem tedious but sticking to the feedback loop is important.
	Not only does it ensure that you have relevant tests,
	it helps ensure you design good software by refactoring with the safety of tests.
	*/

}
