package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ScanAndComputeAndPrintDisplacement(fn func(float64) float64) {
	for {
		fmt.Print("Enter the time t (X to exit): ")
		var input string
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Something wrong, please try again")
			continue
		}

		if input == "X" {
			break
		}
		time, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid input, please try again.")
			continue
		}

		fmt.Printf("The displayment computed by time %v: %v\n", time, fn(time))
	}
}

func ScanThreeInitialFloatingNumbers() (float64, float64, float64) {
	fmt.Println("Enter acceleration, initial velocity and initial displacement by spaces:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	strThreeFloatValues := strings.Fields(strings.TrimSpace(input))

	floats := make([]float64, 0, 3)

	for _, str := range strThreeFloatValues {
		num, err := strconv.ParseFloat(str, 64)
		if err != nil {
			fmt.Printf("Invalid floating number: %s\n", str)
			continue
		}
		floats = append(floats, num)
	}

	return floats[0], floats[1], floats[2]
}

func GenDisplaceFn(a float64, v float64, s float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*t*t*a + t*v + s
	}
}

func main() {
	a, v, s := ScanThreeInitialFloatingNumbers()
	displaceFn := GenDisplaceFn(a, v, s)
	ScanAndComputeAndPrintDisplacement(displaceFn)
}
