package main

import (
	"fmt"
)

func Trunc(x float64) int {
	return int(x)
}

func main() {
	var num float64

	// Prompt the user to enter a floating point number
	fmt.Print("Enter a floating point number: ")
	_, err := fmt.Scan(&num)

	// Print error message if needed
	if err != nil {
		fmt.Println(err)
	}

	// Convert the floating point number to an integer using type casting
	truncated := Trunc(num)

	// Print the truncated integer
	fmt.Println(truncated)
}
