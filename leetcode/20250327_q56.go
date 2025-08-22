package main

import (
	"fmt"
	"sort"
)

func main12() {
	var inter = [][]int{[]int{2, 3}, []int{5, 5}, []int{2, 2}, []int{3, 4}, []int{3, 4}}
	fmt.Println(merge(inter))
}
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var res [][]int
	for i := 0; i < len(intervals); i++ {
		var temp = make([]int, 2)
		copy(temp, intervals[i])
		for i+1 < len(intervals) && intervals[i+1][0] <= temp[1] {
			temp[1] = max(intervals[i+1][1], intervals[i][1])
			i++
		}
		if len(res) == 0 || res[len(res)-1][1] < temp[0] {
			res = append(res, temp)
		}
	}
	return res
}
