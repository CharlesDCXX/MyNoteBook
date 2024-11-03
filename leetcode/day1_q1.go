package main

func twoSum(nums []int, target int) []int {
	intMap := make(map[int]int)
	for i, v := range nums {
		intMap[v] = i
	}
	for i, v := range nums {
		vTemp := target - v
		if j, ok := intMap[vTemp]; ok {
			if i != j {
				return []int{i, j}
			}

		}
	}
	return []int{}
}

// 第一题，刷很多遍了，算重新开始，睡觉吧
