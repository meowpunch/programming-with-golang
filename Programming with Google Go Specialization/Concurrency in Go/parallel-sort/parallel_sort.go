package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Sort(integers []int) {
	n := len(integers)

	if n < 4 {
		sort.Ints(integers)
	}

	partitions := [4][]int{}

	for i, _ := range partitions {
		partitions[i] = integers[i*n/4 : (i+1)*n/4]
	}

	done := make(chan bool)

	for _, partition := range partitions {
		go func(p []int) {
			sort.Ints(p)
			fmt.Println("Sorted partition is", p)
			done <- true
		}(partition)
	}

	for range partitions {
		<-done
	}

	mergedInteger := make([]int, n)
	merge(partitions, mergedInteger)
	copy(integers, mergedInteger)
}

func merge(partitions [4][]int, merged []int) {
	partitionSizes := [4]int{}

	for i := range partitions {
		partitionSizes[i] = len(partitions[i])
	}

	index := [4]int{}
	for i := 0; i < len(merged); i++ {
		minPartitionIndex := -1
		minValue := 0

		for j := 0; j < 4; j++ {
			if index[j] < partitionSizes[j] && (minPartitionIndex == -1 || partitions[j][index[j]] < minValue) {
				minPartitionIndex = j
				minValue = partitions[j][index[j]]
			}
		}
		merged[i] = minValue
		index[minPartitionIndex]++
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
	integers := ScanIntegers()
	Sort(integers)
	fmt.Println("Sorted integers is", integers)
}
