package main

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	var arrayJump = make([]int, len(nums), len(nums))
	for i, v := range arrayJump {
		for j := i + 1; j <= i+nums[i]; j++ {
			if j == len(nums)-1 {
				return v + 1
			}
			if arrayJump[j] == 0 || arrayJump[j] > v+1 {
				arrayJump[j] = v + 1
			}
		}
	}
	return 0
}
