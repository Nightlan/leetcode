package main

import "fmt"

func spiralOrder(matrix [][]int) (res []int) {
	height := len(matrix)
	if height == 0 {
		return
	}
	width := len(matrix[0])
	if width == 0 {
		return
	}
	res = make([]int, 0, height*width)
	m := (width + 1) / 2
	n := (height + 1) / 2
	for i := 0; i < m && i < n; i++ {
		for j := i; j < width-i; j++ { // 上
			res = append(res, matrix[i][j])
		}
		for j := i + 1; j < height-i; j++ { // 右
			res = append(res, matrix[j][width-1-i])
		}
		for j := width - i - 1 - 1; j >= i && height-1-i > i; j-- { // 下
			res = append(res, matrix[height-1-i][j])
		}
		for j := height - i - 1 - 1; j >= i+1 && i < width-1-i; j-- { // 左
			res = append(res, matrix[j][i])
		}
	}
	return
}

func main() {
	testMatrix := [][]int{
		{5, 1, 9, 11, 33},
		{2, 4, 8, 10, 35},
		{13, 3, 6, 7, 55},
		{15, 14, 12, 16, 66},
		{15, 14, 12, 16, 66},
	}
	fmt.Println(spiralOrder(testMatrix))
	return
}
