package main

import "fmt"

func main() {
	var num1 = []int{2, 0}
	var num2 = []int{1}
	var m = 1
	var n = 1
	merge(num1, m, num2, n)
	fmt.Println(num1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	var i1 = m - 1
	var j1 = n - 1
	if n == 0 {
		return
	}
	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
		return
	}
	for i := m + n - 1; i >= 0 && i1 >= 0 && j1 >= 0; i-- {
		if nums1[i1] > nums2[j1] {
			nums1[i] = nums1[i1]
			i1--
		} else {
			nums1[i] = nums2[j1]
			j1--
		}
	}
	if j1 < 0 {
		return
	}
	if i1 < 0 {
		for i := 0; i <= j1; i++ {
			nums1[i] = nums2[i]
		}
	}

}
