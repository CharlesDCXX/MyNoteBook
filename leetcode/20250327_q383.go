package main

func canConstruct(ransomNote string, magazine string) bool {
	var magazineMap = make(map[rune]int)
	for _, v := range []rune(magazine) {
		magazineMap[v] = magazineMap[v] + 1
	}
	for _, v := range []rune(ransomNote) {
		if magazineMap[v] == 0 {
			return false
		}
		magazineMap[v] = magazineMap[v] - 1
	}
	return true
}
