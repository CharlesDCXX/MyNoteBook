package main

//func letterCombinations(digits string) []string {
//	if len(digits) == 0{
//		return []string{}
//	}
//	var mapInt map[string]string = map[string]string{ "2": "abc", "3": "def", "4": "ghi", "5": "jkl", "6": "mno", "7": "pqrs", "8": "tuv", "9": "wxyz", }
//	var nums []int
//	for _, r := range digits {
//		num, _ := strconv.Atoi(string(r))
//		nums = append(nums, num)
//	}
//	var add func(nums []int) []string
//	add = func(nums []int) []string {
//		var result []string
//		res, _ := mapInt[nums[0]]
//		if len(nums) == 1 {
//			for i := 0; i < len(res); i++ {
//				result = append(result, res[i])
//			}
//			return result
//		}
//		resultTemp := add(nums[1:])
//		for i := 0; i < len(res); i++ {
//			for _, i2 := range resultTemp {
//				result = append(result, res[i]+i2)
//			}
//		}
//		return result
//	}
//	return add(nums)
//}
