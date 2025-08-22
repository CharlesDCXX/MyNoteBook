package main

func hIndex(citations []int) int {
	if len(citations) == 1 {
		if citations[0] >= 1 {
			return 1
		} else {
			return 0
		}
	}
	var length = len(citations)
	var hmap = make(map[int]int)
	hmax := 0
	for _, v := range citations {
		hmap[v] = hmap[v] + 1
		if v > hmax {
			hmax = v
		}
	}
	hsum := 0
	for i := hmax; i >= length; i-- {
		hsum += hmap[i]
	}
	if hsum >= length {
		return length
	}
	for i := length - 1; i >= 0; i-- {
		hsum += hmap[i]
		if hsum >= i {
			return i
		}
	}
	return 0
}
