package main

import "fmt"

func test(arr []int) {
	arr[1] = 1
	arr = append(arr, 2)
}

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

func main() {
	arr := []int{0, 0}
	test(arr)
	fmt.Println(arr)
	fmt.Println(INT_MAX, INT_MIN)
	myNum := []int{10, 20, 30, 40, 50}
	// 创建新的切片，其长度为 2 个元素，容量为 4 个元素
	newNum := myNum[2:3]
	// 使用原有的容量来分配一个新元素
	// 将新元素赋值为 60
	newNum = append(newNum, 60)
	newNum = append(newNum, 70)
	newNum = append(newNum, 80)
	fmt.Println(myNum)
	fmt.Println(newNum)
}
