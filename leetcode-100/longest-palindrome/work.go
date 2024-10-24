package longest

type Judge struct {
	maxLength    int
	start        int
	nums         []int
	isPalindrome bool
}

func longestPalindrome1(s string) string {
	// 还要加一个判断语句，当小于2个的时候进行比较就行
	if len(s) < 2 {
		return s
	}
	judge := new(Judge)
	n := len(s)

	judge.maxLength = 1
	for x := 0; x < n; x++ {
		y := n - 1
		for y > x && y >= x+judge.maxLength/2 {
			if s[x] == s[y] {
				judge.isPalindrome = true
				left, right := x, y
				for left < right {
					if s[left] != s[right] {
						judge.isPalindrome = false
						break
					}
					left++
					right--
				}
				if judge.isPalindrome {
					judge.nums = append(judge.nums, y-x+1)
					if judge.maxLength < y-x+1 {
						judge.start = x
						judge.maxLength = y - x + 1
					}
				}
			} else {
				judge.nums = []int{}
			}
			y--
		}
	}
	return s[judge.start : judge.start+judge.maxLength]
}

// 设计三个循环，第一个循环开始元素，第二个内循环，循环从尾部指针开始到指定位置，第三个循环，判断是否“还是”相同的，获取当前最大值

// 包含一个内循环和一个外循环

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	maxLen := 1
	start := 0
	n := len(s)

	for x := 0; x < n; x++ {
		y := n - 1
		nums := []int{}

		for y > x && y >= x+maxLen/2 { // 只需要检查比当前最长回文子串更长的情况
			if s[x] == s[y] {
				isPalindrome := true
				left, right := x, y
				for left < right {
					if s[left] != s[right] {
						isPalindrome = false
						break
					}
					left++
					right--
				}
				if isPalindrome {
					length := y - x + 1
					nums = append(nums, length)
					if length > maxLen {
						maxLen = length
						start = x
					}
				} else {
					nums = []int{}
				}
			}
			y--
		}
	}

	return s[start : start+maxLen]
}
