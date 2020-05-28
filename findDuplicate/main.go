package main

func findDuplicate(nums []int) int {
	fast := nums[nums[0]]
	slow := nums[0]
	for fast != slow {
		fast = nums[nums[fast]]
		slow = nums[slow]
	}
	slow = 0
	for fast != slow {
		fast = nums[fast]
		slow = nums[slow]
	}
	return slow
}

func main() {

}
