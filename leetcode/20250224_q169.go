package main

func majorityElement(nums []int) int {
	mapK := make(map[int]int)
	for _, v := range nums {
		index, _ := mapK[v]
		mapK[v] = index + 1
	}
	for _, inde := range mapK {
		if inde >= len(nums)/2 {
			return inde
		}
	}
	return 0
}
