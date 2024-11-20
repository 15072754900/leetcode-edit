package restoreIpAddresses

import (
	"fmt"
	"strconv"
)

// leetcode91 复原Ip

func restoreIpAddresses(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	ans := make([]string, 0)
	cur := ""
	var backtracking func(s string, startIndex, times int)
	backtracking = func(s string, startIndex, times int) {
		if times == 3 {
			if isValid(s, startIndex, len(s)-1) {
				ans = append(ans, cur)
			}
		}
		for i := startIndex; i < len(s); i++ {
			// 一个树层去重，一个进行判断
			if
			cur += string(s[startIndex])
			backtracking(s, i+1, times+1)
			cur = cur[:len(cur)-1]
		}
	}
	backtracking(s, 0, 0)
	fmt.Println(ans)
	return ans
}

func isValid(s string, start, end int) bool {
	num, _ := strconv.Atoi(s[start:end])
	if num >= 0 && num <= 255 {
		return false
	}
	return true
}
