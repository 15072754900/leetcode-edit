package letterCombinations

import "fmt"

func letterCombinations(digits string) []string {
	if digits == "" {
		return make([]string, 0)
	}
	ans := make([]string, 0)
	cur := ""
	hash := map[string]string{"2": "abc", "3": "def", "4": "ghi", "5": "jkl", "6": "mno", "7": "pqrs", "8": "tuv", "9": "wxyz"}
	var backtracking func(s string, startIndex int)
	backtracking = func(s string, startIndex int) {
		if len(s) == startIndex { // 说明到达了尾部，相较于其他的问题：像内容加内容，或者向内容到达内容，这个是单个内容修改内容
			ans = append(ans, cur)
			fmt.Println(cur)
			return
		}
		letter := hash[string(s[startIndex])] // 最关键的是执行这一步骤，然后是向下取cur
		for i := 0; i < len(letter); i++ {
			cur += string(letter[i])
			backtracking(s, startIndex+1) // 这里一定是i，不是startIndex，因为i才是可以表示变化的剩余层数的变量
			// 要知道如何对string进行回溯，就是和slice一样，进行区间取值就行
			cur = cur[:len(cur)-1]
		}
	}
	backtracking(digits, 0)
	return ans
}
