package main

import (
	"reflect"
	"testing"
)

func TestSortIntSlice(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{3}, []int{3}},
		{[]int{9, 5, 7}, []int{5, 7, 9}},
		{[]int{100, 50, 200, 1}, []int{1, 50, 100, 200}},
		{[]int{0, -1, -5}, []int{-5, -1, 0}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		SortIntSlice(test.input)
		if !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("For input %v, expected %v but got %v", test.input, test.expected, test.input)
		}
	}
}
