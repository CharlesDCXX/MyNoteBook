package main

import "strings"

func isValid(s string) bool {
	mapKey := map[string]string{"(": ")", "[": "]", "{": "}"}
	var ss = "([{"
	var queue []string
	sKey := []rune(s)

	for i := 0; i < len(sKey); i++ {
		if strings.Contains(ss, string(sKey[i])) {
			queue = append(queue, string(sKey[i]))
			continue
		}
		if len(queue) == 0 {
			return false
		}
		if mapKey[queue[len(queue)-1]] != string(sKey[i]) {
			return false
		}
		queue = queue[:len(queue)-1]
	}
	if len(queue) == 0 {
		return true
	}
	return false
}
