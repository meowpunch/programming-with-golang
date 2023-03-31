package main

import (
	"fmt"
	"regexp"
	"strings"
)

func FindIan(s string) bool {
	re := regexp.MustCompile(`(?i)^i.*a.*n$`)
	return re.MatchString(s)
}

func FindIanWithoutRegex(s string) bool {
	return strings.HasPrefix(strings.ToLower(s), "i") &&
		strings.Contains(strings.ToLower(s), "a") &&
		strings.HasSuffix(strings.ToLower(s), "n")
}

func main() {
	// Prompt the user to enter a string
	var input string
	fmt.Print("Enter a string: ")
	_, err := fmt.Scanln(&input)

	// Print error message if needed
	if err != nil {
		fmt.Println(err)
	}

	// Check if the string starts with 'i', ends with 'n', and contains 'a'
	if FindIan(input) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
