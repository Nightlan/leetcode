package main

import "fmt"

func rotate(matrix [][]int) [][]int {
	l := len(matrix)
	for i := 0; i < l/2; i++ {
		for j := 0; j < (l+1)/2; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[l-1-j][i]
			matrix[l-1-j][i] = matrix[l-1-i][l-1-j]
			matrix[l-1-i][l-1-j] = matrix[j][l-i-1]
			matrix[j][l-i-1] = tmp
		}
	}
	return matrix
}

func main() {
	testMatrix := [][]int{
		{5, 1, 9, 11, 33},
		{2, 4, 8, 10, 35},
		{13, 3, 6, 7, 55},
		{15, 14, 12, 16, 66},
		{15, 14, 12, 16, 66},
	}
	rotate(testMatrix)
	for _, v := range testMatrix {
		fmt.Println(v)
	}
	return
}
