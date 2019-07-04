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
	got := Hello("Chris") //We're declaring some variables with the syntax varName := value, which lets us re-use
	//some values in our test for readability.
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
