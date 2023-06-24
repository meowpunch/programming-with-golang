package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(integers []int, i int) {
	integers[i], integers[i+1] = integers[i+1], integers[i]
}

func BubbleSort(integers []int) {
	N := len(integers)
	for i := 0; i < N; i++ {
		for j := 0; j < N-i-1; j++ {
			if integers[j] > integers[j+1] {
				Swap(integers, j)
			}
		}
	}
}

func ScanIntegers() []int {
	fmt.Println("Enter up to ten integers separated by spaces:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	strIntegers := strings.Fields(input)

	integers := make([]int, 0, len(strIntegers))

	for _, str := range strIntegers {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("Invalid integer: %s\n", str)
			continue
		}
		integers = append(integers, num)
	}

	return integers
}

func main() {
	tenIntegers := ScanIntegers()
	BubbleSort(tenIntegers)
	fmt.Print(tenIntegers)
}
