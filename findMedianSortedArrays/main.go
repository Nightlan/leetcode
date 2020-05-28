package main

import "fmt"

const (
	maxInt = int((^uint(0)) >> 1)
	minInt = ^maxInt
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums2) < len(nums1) {
		return findMedianSortedArrays(nums2, nums1)
	}
	l1 := 0
	r1 := len(nums1) * 2
	mark := len(nums1) + len(nums2)
	lMax1 := 0
	rMin1 := 0
	lMax2 := 0
	rMin2 := 0
	for l1 <= r1 {
		pos1 := (l1 + r1) / 2
		if pos1 == 0 {
			lMax1 = minInt
		} else {
			lMax1 = nums1[(pos1-1)/2]
		}
		if pos1 == 2*len(nums1) {
			rMin1 = maxInt
		} else {
			rMin1 = nums1[(pos1)/2]
		}
		pos2 := mark - pos1
		if pos2 == 0 {
			lMax2 = minInt
		} else {
			lMax2 = nums2[(pos2-1)/2]
		}
		if pos2 == 2*len(nums2) {
			rMin2 = maxInt
		} else {
			rMin2 = nums2[(pos2)/2]
		}
		if rMin1 < lMax2 {
			l1 = pos1 + 1
		} else if rMin2 < lMax1 {
			r1 = pos1 - 1
		} else {
			break
		}
	}
	if lMax2 > lMax1 {
		lMax1 = lMax2
	}
	if rMin2 < rMin1 {
		rMin1 = rMin2
	}
	return (float64(lMax1) + float64(rMin1)) / 2
}

func main() {
	fmt.Println(maxInt, minInt)
	nums1 := []int{2, 3}
	nums2 := []int{1}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
