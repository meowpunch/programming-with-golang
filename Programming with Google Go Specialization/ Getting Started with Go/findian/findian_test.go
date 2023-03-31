package main

import "testing"

func TestFindIan(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"ian", true},
		{"Ian", true},
		{"iuiygaygn", true},
		{"I d skd a efju N", true},
		{"ihhhhhn", false},
		{"ina", false},
		{"xian", false},
	}

	for _, test := range tests {
		output := FindIan(test.input)
		if output != test.expected {
			t.Errorf("For input %q, expected %t but got %t", test.input, test.expected, output)
		}
	}
}

func TestFindIanWithoutRegex(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"ian", true},
		{"Ian", true},
		{"iuiygaygn", true},
		{"I d skd a efju N", true},
		{"ihhhhhn", false},
		{"ina", false},
		{"xian", false},
	}

	for _, test := range tests {
		output := FindIanWithoutRegex(test.input)
		if output != test.expected {
			t.Errorf("For input %q, expected %t but got %t", test.input, test.expected, output)
		}
	}
}
