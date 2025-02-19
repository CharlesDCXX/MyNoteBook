package main

import (
	"fmt"
	"regexp"
)

func extractEmail(path string) string {
	// 使用正则表达式提取邮箱，去掉数字和.xlsx后缀
	re := regexp.MustCompile(`\d+-([a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,})\.xlsx$`)
	match := re.FindStringSubmatch(path)
	if len(match) > 1 {
		return match[1] // 返回匹配到的邮箱部分
	}
	return ""
}

func main() {
	path := "/path/dasdas/dsad-das/ds_dsa/392817631278-duanchenxi@haique-tech.com.xlsx"
	email := extractEmail(path)
	fmt.Println("Extracted email:", email)
}
