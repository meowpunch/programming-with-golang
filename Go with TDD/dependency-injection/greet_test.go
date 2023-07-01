package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Jinhoon")

	got := buffer.String()
	want := "Hello Jinhoon"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
