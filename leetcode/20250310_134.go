package main

import "math"

func canCompleteCircuit(gas []int, cost []int) int {
	maxN := 0
	minN := math.MaxInt
	index := 0
	sum := 0
	for i := 0; i < len(gas); i++ {
		if cost[i]-gas[i] > maxN {
			maxN = cost[i] - gas[i]
			index = i
		}
		if cost[i]-gas[i] < minN {
			minN = cost[i] - gas[i]
		}
		sum += gas[i] - cost[i]
	}
	if sum >= 0 {
		if index <= len(gas)-2 {
			return index + 1
		}
		return 0
	}
	return -1
}
