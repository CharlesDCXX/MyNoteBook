package main

func removeElement(nums []int, val int) int {
	var l int
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[l] = nums[i]
			l++
		}
	}
	return l
}
