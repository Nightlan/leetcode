package main

func quickSort(array []int) {
	if len(array) <= 1 {
		return
	}
	pivot := array[0]
	mark := 0
	for pos := 1; pos < len(array); pos++ {
		if array[pos] < pivot {
			mark++
			array[mark], array[pos] = array[pos], array[mark]
		}
	}
	array[0], array[mark] = array[mark], array[0]
	quickSort(array[:mark])
	quickSort(array[mark+1:])
}

func main() {

}
