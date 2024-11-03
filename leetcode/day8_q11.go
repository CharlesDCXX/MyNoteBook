package main

func main11() {
	var height = []int{1, 2, 3, 4, 5, 25, 24, 3, 4}
	maxArea(height)
}
func maxArea(height []int) int {
	l, maxA, maxHigh := 0, 0, 0
	for l < len(height)-1 {
		r := len(height) - 1
		maxR := 0
		if l != 0 && height[l] < maxHigh {
			l++
			continue
		}
		high := height[l]
		if high > maxHigh {
			maxHigh = high
		}
		for r > l {
			if height[r] < maxR {
				r--
				continue
			}
			maxR = height[r]
			area := (r - l) * min(height[l], height[r])
			maxA = max(maxA, area)
			r--
		}
		l++
	}
	return maxA
}
