package main

import "fmt"

func subarraysDivByK(A []int, K int) int {
	surMap := make(map[int]int, len(A))
	surMap[0] = 1
	sum := 0
	res := 0
	for i := 0; i < len(A); i++ {
		sum += A[i]
		sur := (sum%K + K) % K
		_, ok := surMap[sur]
		if ok {
			res += surMap[sur]
		}
		surMap[sur]++
	}
	return res
}

func main() {
	fmt.Println(-1 % 4)
}
