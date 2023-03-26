package main

import "testing"

/*
Write a program which prompts the user to enter a string. The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’.
The program should print “Not Found!” otherwise. The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

Examples: The program should print “Found!” for the following example entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”. The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.
*/

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
