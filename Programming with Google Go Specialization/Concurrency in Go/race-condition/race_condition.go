package main

import (
	"fmt"
	"time"
)

func main() {
	x := 0
	integers := make([]int, 0, 100000)

	increaseXAndPrintX := func() {
		x = x + 1
		integers = append(integers, x)
	}

	for i := 0; i < 100000; i++ {
		go increaseXAndPrintX()
	}

	time.Sleep(100 * time.Millisecond)

	fmt.Printf("%d != %d because of race condition\n", 99999, findMax(integers))
}

func findMax(numbers []int) int {
	if len(numbers) == 0 {
		panic("Empty slice")
	}

	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}

	return max
}
