package main

import "fmt"

func siftDown(array []int, index int) {
	for {
		child := index*2 + 1
		if child >= len(array) {
			break
		}
		if child+1 < len(array) && array[child+1] > array[child] {
			child++
		}
		if array[child] <= array[index] {
			break
		}
		array[child], array[index] = array[index], array[child]
		index = child
	}
}

func heapSort(array []int) {
	for i := len(array) / 2; i >= 0; i-- {
		siftDown(array, i)
	}
	for i := len(array) - 1; i >= 0; i-- {
		array[0], array[i] = array[i], array[0]
		siftDown(array[:i], 0)
	}
}

func main() {
	array := []int{2, 33, 333, 44, 33, 4, 77, 3, 1, 2}
	heapSort(array)
	fmt.Println(array)
}
