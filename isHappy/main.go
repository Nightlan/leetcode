package main

import "fmt"

type emptyType struct{}

var emptyVal emptyType

func isHappy(n int) bool {
	tmpVal := make(map[int]emptyType)
	for {
		sum := 0
		for val := n; val > 0; val = val / 10 {
			bitVal := val % 10
			sum += bitVal * bitVal
		}
		n = sum
		if n == 1 {
			return true
		}
		_, ok := tmpVal[n]
		if ok {
			break
		}
		tmpVal[n] = emptyVal
	}
	return false
}

func main() {
	n := 19444
	fmt.Println(isHappy(n))
}
