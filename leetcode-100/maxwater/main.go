package maxwater

import "fmt"

func do() {
	// 定义双指针，进行在取值区间内移动，计算面积，并与当前最大值比较，获取更大的一个，然后返回
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	l := 0
	r := len(nums) - 1
	var ans int

	for l < r {
		cur := (r - l) * min(nums[l], nums[r])
		ans = max(ans, cur)
		if nums[l] <= nums[r] {
			l++
		} else {
			r--
		}
	}
	fmt.Println(ans)
}

// 处理力扣上11题，盛水的容器，使用的是双指针处理。

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
