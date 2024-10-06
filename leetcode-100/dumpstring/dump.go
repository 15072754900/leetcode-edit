package dumpstring

import "fmt"

func LengthOfLongestSubstring(s string) int {
	left, right := 0, 1
	max := 0
	for left < right {
		if right < len(s) {
			if !judgeIn(s[right], s[left:right]) {
				left++
				right = left + 1
				max++
			}
			right++
		}
	}
	return max
}

func judgeIn(a byte, b string) bool {
	for i := 0; i < len(b); i++ {
		if a == b[i] {
			return false
		}
	}
	return true
}

func DoSomething() {
	fmt.Println("hello")
}

// 上面使用了一个函数进行处理重复元素，消耗过大，直接维护一个哈希表就行，然后不需要每次都进行right的修改。

func DoAgain(s string) int {
	if len(s) == 0 {
		return 0
	}
	left, right, max := 0, 0, 0
	hash := make(map[byte]bool)
	for right < len(s) {
		if !hash[s[right]] {
			hash[s[right]] = true
			right++
			// 进行更新max
			if max < right-left {
				max = right - left
			}
		} else {
			// 删除记录
			delete(hash, s[left])
			left++
		}
	}
	return max
}

func DoAgain2(s string) int {
	if len(s) == 0 {
		return 0
	}
	left, right, maxLen := 0, 0, 0
	charIndexMap := make(map[byte]int)

	for right < len(s) {
		if lastIdx, ok := charIndexMap[s[right]]; ok && lastIdx >= left {
			// 如果当前字符 s[right] 在当前窗口内已经出现过
			// 更新左指针到该字符上次出现位置的下一个位置
			left = lastIdx + 1
		}

		// 更新当前字符的最后出现位置
		charIndexMap[s[right]] = right

		// 计算当前窗口的长度，并更新最大长度
		currentLength := right - left + 1
		if currentLength > maxLen {
			maxLen = currentLength
		}

		// 移动右指针
		right++
	}

	return maxLen
}
