package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestScanNameAndAddress(t *testing.T) {
	name, address := "John Smith", "123 Main St"

	// Use a pipe to redirect standard input for testing purposes
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	defer w.Close()
	oldStdin := os.Stdin
	os.Stdin = r
	defer func() { os.Stdin = oldStdin }()

	// Write test input to the write end of the pipe
	go func() {
		fmt.Fprintf(w, "%s\n%s\n", name, address)
	}()

	gotName, gotAddress := ScanNameAndAddress()

	if gotName != name {
		t.Errorf("ScanNameAndAddress() name = %q, want %q", gotName, name)
	}
	if gotAddress != address {
		t.Errorf("ScanNameAndAddress() address = %q, want %q", gotAddress, address)
	}
}

func TestCreateJSON(t *testing.T) {
	name, address := "John Smith", "123 Main St"
	want := `{"address":"123 Main St","name":"John Smith"}`

	got, err := CreateJSON(name, address)
	if err != nil {
		t.Errorf("CreateJSON() error = %v", err)
	}
	if !reflect.DeepEqual(string(got), want) {
		t.Errorf("CreateJSON() = %v, want %v", string(got), want)
	}
}
