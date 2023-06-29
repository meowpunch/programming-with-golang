package main

import (
	"fmt"
	"time"
)

/*
- What is race condition?
A race condition is a problem that can happen as a result of all these possible interleaving that you have to consider.
More precisely, the result of program is non-deterministic. It means that the output can be different by execution even if the input is the same.

- How it can occur?
Race condition occurs due to communication. Especially in terms of go and goroutine, it happens when goroutine communicates on shared variable.

Let's look at the below code. Goroutine 1 and goroutine 2 communicates on shared variable x.
Mostly the result of `2 == findMax(integers)` is `true`. It can be `false`. And if there are more goroutines run function `increaseXAndPrintX`. The result is mostly `false`
*/
func main() {
	// shared variable
	x := 0
	integers := make([]int, 0, 2)

	increaseXAndPrintX := func() {
		x = x + 1
		integers = append(integers, x)
	}

	// goroutine 1
	go increaseXAndPrintX()
	// goroutine 2
	go increaseXAndPrintX()

	time.Sleep(100 * time.Millisecond)

	fmt.Println(2 == findMax(integers))
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
