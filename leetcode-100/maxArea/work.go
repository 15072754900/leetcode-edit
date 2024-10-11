package maxArea

import "fmt"

func maxArea(height []int) int {
	// 这道题目也应该是动态双指针处理
	if len(height) == 2 {
		return min(height[0], height[1])
	}
	ans := 0
	n := len(height)
	left, right := 0, n-1
	for left < right {
		fmt.Println(left, right)
		x := min(height[left], height[right])
		cur := (right - left) * x
		ans = max(ans, cur)
		fmt.Println("right-left,height[left], height[right],ans", right-left, height[left], height[right], ans)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
