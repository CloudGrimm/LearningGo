package main

import "testing"
func TestHello(t *testing.T) {
	got := Hello("Tinashe")
	want := "Hello, Tinashe"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}