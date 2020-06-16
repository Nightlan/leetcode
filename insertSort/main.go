package main

import (
	"fmt"
)

func insertSort(array []int) {
	for i := 1; i < len(array); i++ {
		for j := i; j > 0 && array[j-1] > array[j]; j-- {
			array[j-1], array[j] = array[j], array[j-1]
		}
	}
}

func main() {
	array := []int{2, 33, 333, 44, 33, 4, 77, 3, 1, 2}
	insertSort(array)
	fmt.Println(array)
}
