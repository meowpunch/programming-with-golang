package main

import (
	"fmt"
	"sort"
	"strconv"
)

func SortIntSlice(intSlice []int) {
	sort.Ints(intSlice)
}

func main() {
	intSlice := make([]int, 0, 3)
	for {
		fmt.Print("Enter an integer (X to exit): ")
		var input string
		fmt.Scan(&input)
		if input == "X" {
			break
		}
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input, please try again.")
			continue
		}
		intSlice = append(intSlice, num)
	}

	SortIntSlice(intSlice)
	fmt.Println(intSlice)
}
