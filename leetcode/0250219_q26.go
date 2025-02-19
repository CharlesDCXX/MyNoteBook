package main

func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	var temp, index = nums[0], 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != temp {
			nums[index] = nums[i]
			temp = nums[i]
			index++
		}
	}
	return index
}
