package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestReadNamesFromFileSuccess(t *testing.T) {
	// Create a temporary test file
	file, err := ioutil.TempFile("", "test_file_")
	if err != nil {
		t.Fatalf("Failed to create temporary test file: %v", err)
	}
	defer os.Remove(file.Name())

	// Write test data to the file
	data := "John Smith\nJane Doe\n"
	if _, err := file.WriteString(data); err != nil {
		t.Fatalf("Failed to write test data to file: %v", err)
	}

	// Close the file to flush changes to disk
	if err := file.Close(); err != nil {
		t.Fatalf("Failed to close test file: %v", err)
	}

	// Call the function being tested
	got, err := ReadNamesFromFile(file.Name())
	if err != nil {
		t.Fatalf("ReadNamesFromFile failed: %v", err)
	}

	// Check the result
	want := []User{
		{Fname: "John", Lname: "Smith"},
		{Fname: "Jane", Lname: "Doe"},
	}
	if len(got) != len(want) {
		t.Errorf("ReadNamesFromFile returned %d names, want %d", len(got), len(want))
	}
	for i, name := range want {
		if got[i].Fname != name.Fname || got[i].Lname != name.Lname {
			t.Errorf("ReadNamesFromFile returned %v, want %v", got[i], name)
		}
	}
}

func TestReadNamesFromFileFailure(t *testing.T) {
	// Call the function being tested with a non-existent file
	_, err := ReadNamesFromFile("nonexistent_file.txt")
	if err == nil {
		t.Errorf("ReadNamesFromFile should have returned an error for non-existent file")
	}
}
