package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	integers := []int{5, 2, 9, 1, -7, 3, 8, 4, 6, 10}
	expected := []int{-7, 1, 2, 3, 4, 5, 6, 8, 9, 10}

	BubbleSort(integers)

	if !reflect.DeepEqual(integers, expected) {
		t.Errorf("BubbleSort failed. Expected %v, got %v", expected, integers)
	}
}

func TestSwap(t *testing.T) {
	integers := []int{1, 2, 3, 4, 5}
	index := 2
	expected := []int{1, 2, 4, 3, 5}

	Swap(integers, index)

	if !reflect.DeepEqual(integers, expected) {
		t.Errorf("Swap failed. Expected %v, got %v", expected, integers)
	}
}

func TestScanTenIntegers(t *testing.T) {
	input := "5 2 9 1 7 3 8 4 6 10"
	expected := []int{5, 2, 9, 1, 7, 3, 8, 4, 6, 10}

	// Redirect standard input
	oldStdin := os.Stdin
	defer func() {
		os.Stdin = oldStdin
	}()

	r, w, _ := os.Pipe()
	os.Stdin = r

	// Write input to stdin
	go func() {
		fmt.Fprintln(w, input)
		w.Close()
	}()

	result := ScanIntegers()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ScanIntegers failed. Expected %v, got %v", expected, result)
	}
}
