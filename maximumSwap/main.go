package main

import (
	"fmt"
	"strconv"
)

func maximumSwap(num int) int {
	numStr := []byte(strconv.Itoa(num))
	lastVal := numStr[0]
	curMax := 0
	curMin := 0
	for i := 1; i < len(numStr); i++ {
		if numStr[i] > lastVal {
			curMax = i
			curMin = i - 1
			for j := curMax + 1; j < len(numStr); j++ {
				if numStr[j] >= numStr[curMax] {
					curMax = j
				}
			}
			for k := curMin - 1; k >= 0; k-- {
				if numStr[k] < numStr[curMax] {
					curMin = k
				}
			}
			break
		}
		lastVal = numStr[i]
	}
	numStr[curMin], numStr[curMax] = numStr[curMax], numStr[curMin]
	res, _ := strconv.Atoi(string(numStr))
	return res
}

func main() {
	fmt.Println(maximumSwap(10909091))
}
