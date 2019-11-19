package main

import "fmt"

func trap(height []int) (area int) {
	left, right := 0, len(height)-1
	leftMax, rightMax := height[left], height[right]
	for left < right {
		if leftMax <= rightMax {
			left++
			if leftMax <= height[left] {
				leftMax = height[left]
			} else {
				area += leftMax - height[left]
			}
		} else {
			right--
			if rightMax <= height[right] {
				rightMax = height[right]
			} else {
				area += rightMax - height[right]
			}
		}
	}
	return area
}

func main() {
	testNum := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(testNum))
	return
}
