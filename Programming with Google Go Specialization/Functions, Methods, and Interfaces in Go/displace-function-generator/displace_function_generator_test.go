package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestGenDisplaceFn(t *testing.T) {
	// Test case 1
	a := 10.0
	v := 2.0
	s := 1.0
	fn := GenDisplaceFn(a, v, s)

	// Check the generated function with different time values
	t1 := 3.0
	expectedOutput1 := 52.0
	result1 := fn(t1)
	if result1 != expectedOutput1 {
		t.Errorf("GenDisplaceFn failed. For t = %f, expected %f, got %f", t1, expectedOutput1, result1)
	}

	t2 := 5.0
	expectedOutput2 := 136.0
	result2 := fn(t2)
	if result2 != expectedOutput2 {
		t.Errorf("GenDisplaceFn failed. For t = %f, expected %f, got %f", t2, expectedOutput2, result2)
	}
}

func TestScanThreeInitialFloatingNumbers(t *testing.T) {
	input := "3 2 5"
	var expectedAcceleration, expectedVelocity, expectedDisplacement float64
	expectedAcceleration = 3
	expectedVelocity = 2
	expectedDisplacement = 5

	// Redirect standard input
	oldStdin := os.Stdin
	defer func() {
		os.Stdin = oldStdin
	}()

	r, w, _ := os.Pipe()
	os.Stdin = r

	// Write input to stdin
	go func() {
		fmt.Fprintln(w, input)
		w.Close()
	}()

	a, v, s := ScanThreeInitialFloatingNumbers()
	if !reflect.DeepEqual(a, expectedAcceleration) {
		t.Errorf("TestScanThreeInitialFloatingNumbers failed. ExpectedAcceleration %v, got %v", expectedAcceleration, a)
	}
	if !reflect.DeepEqual(a, expectedAcceleration) {
		t.Errorf("TestScanThreeInitialFloatingNumbers failed. expectedVelocity %v, got %v", expectedVelocity, v)
	}
	if !reflect.DeepEqual(a, expectedAcceleration) {
		t.Errorf("TestScanThreeInitialFloatingNumbers failed. expectedDisplacement %v, got %v", expectedDisplacement, s)
	}
}
