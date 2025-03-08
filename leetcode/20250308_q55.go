package main

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	endLength := len(nums) - 1
	var arrayBool = make([]bool, len(nums), len(nums))
	arrayBool[0] = true
	for i, v := range arrayBool {
		if !v {
			continue
		}
		if i+nums[i] >= endLength {
			return true
		}
		for j := i; j <= i+nums[i]; j++ {
			arrayBool[j] = true
		}
	}
	return false
}
