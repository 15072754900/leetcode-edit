package partition

import "fmt"

// leetcode131 分隔回文串

func partition(s string) [][]string {
	if len(s) == 0 {
		return make([][]string, 0)
	}
	ans := make([][]string, 0)
	cur := make([]string, 0)
	var backtracking func(s string, startIndex int)
	backtracking = func(s string, startIndex int) {
		if startIndex >= len(s) {
			temp := make([]string, len(cur))
			copy(temp, cur)
			ans = append(ans, temp)
			return
		}
		for i := startIndex; i < len(s); i++ {
			fmt.Println(startIndex, i)
			if isPartition(s, startIndex, i) {
				fmt.Println("2-- ", startIndex, i)
				cur = append(cur, s[startIndex:i+1])
			} else {
				continue
			}
			backtracking(s, i+1)
			cur = cur[:len(cur)-1]
		}
	}
	backtracking(s, 0)
	return ans
}

func isPartition(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}
