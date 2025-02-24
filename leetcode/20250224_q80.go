package main

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	var l = 2
	for i := 2; i < len(nums); i++ {
		if nums[l-2] != nums[i] {
			nums[l] = nums[i]
			l++
		}
	}
	return l
}
