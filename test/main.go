package main

import "fmt"

func main(){
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
