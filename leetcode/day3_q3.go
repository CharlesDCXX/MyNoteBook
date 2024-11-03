package main

// 一气呵成，美滋滋
func lengthOfLongestSubstring(s string) int {
	var longestSubstring int
	keyMap := make(map[byte]int)
	var l, r int
	for r < len(s) {
		key := s[r]
		if _, ok := keyMap[key]; !ok {
			keyMap[key] = 1
			r = r + 1
			if longestSubstring < (r - l) {
				longestSubstring = max(longestSubstring, r-l)
			}
			continue
		}
		delete(keyMap, s[l])
		l = l + 1
	}
	return longestSubstring
}
