package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello Golang"

	if want != got {
		t.Fatalf("want %q, got %q\n", got, want)
	}
}
