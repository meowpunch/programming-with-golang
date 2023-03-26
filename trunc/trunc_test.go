package main

import (
	"testing"
)

/*
Write a program which prompts the user to enter a floating point number and prints the integer
which is a truncated version of the floating point number that was entered.
Truncation is the process of removing the digits to the right of the decimal place.
*/
func TestTrunc(t *testing.T) {
	tests := []struct {
		input    float64
		expected int
	}{
		{1.234, 1},
		{4.567, 4},
		{0.0, 0},
		{-1.234, -1},
		{-4.567, -4},
	}

	for _, test := range tests {
		got := Trunc(test.input)
		if got != test.expected {
			t.Errorf("Trunc(%f) = %d; expected %d", test.input, got, test.expected)
		}
	}
}
