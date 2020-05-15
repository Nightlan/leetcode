package main

func twoSum(nums []int, target int) []int {
	existMap := make(map[int]int, 0)
	for i := range nums {
		diff := target - nums[i]
		pos, ok := existMap[diff]
		if ok {
			return []int{pos, i}
		}
		existMap[nums[i]] = i
	}
	return nil
}

func main() {

}
