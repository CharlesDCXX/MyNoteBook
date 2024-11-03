package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var re []int
	var a, b int
	length := len(nums1) + len(nums2)

	for len(nums1) > a && len(nums2) > b {
		if nums1[a] < nums2[b] {
			re = append(re, nums1[a])
			a = a + 1
		} else {
			re = append(re, nums2[b])
			b = b + 1
		}

	}
	for len(nums1) > a {
		re = append(re, nums1[a])
		a = a + 1
	}
	for len(nums2) > b {
		re = append(re, nums2[b])
		b = b + 1
	}
	if length%2 == 0 {
		a := float64(re[length/2-1]+re[length/2]) / 2
		return a
	} else {
		return float64(re[(length-1)/2])
	}
}
func main5() {
	findMedianSortedArrays([]int{1, 2}, []int{3, 4})
}
