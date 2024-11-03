package main

func longestPalindrome(s string) string {
	var length = len(s)
	arr2 := make([][]bool, length)
	for i := range arr2 {
		arr2[i] = make([]bool, length)
	}
	for i := 0; i < length; i++ {
		arr2[i][i] = true
	}
	var maxL, begin = 1, 0
	for l := 2; l < length; l++ {
		for left := 0; left <= length; left++ {
			var right = left + l - 1
			if right >= length {
				break
			}
			if s[left] != s[right] {
				arr2[left][right] = false
			} else {
				if l < 3 {
					arr2[left][right] = true
				} else {
					arr2[left][right] = arr2[left+1][right-1]
				}
			}
			if arr2[left][right] && l > maxL {
				maxL = l
				begin = left
			}
		}
	}
	end := begin + maxL
	return s[begin:end]
}
func main6() {
	longestPalindrome("bb")

}
