package findSubsequences

import "fmt"

// leetcode491 递增子序列，不进行排序，执行过程中一定是要进行树层去除的，还是一个子集问题
// 但是这里根据题意还需要进行树枝去重

func findSubsequences(nums []int) [][]int {
	ans := make([][]int, 0)
	cur := make([]int, 0)
	// 还需要传递一个当前最大值，然后进行树枝去重
	var backtracking func(nums []int, startIndex int)
	backtracking = func(nums []int, startIndex int) {
		if len(cur) > 1 {
			temp := make([]int, len(cur))
			copy(temp, cur)
			ans = append(ans, temp)
		}
		// 设置单层元素的取值限制，去重都是在循环里面成立，但是定义是在回溯函数里面定义
		uset := make(map[int]bool)
		for i := startIndex; i < len(nums); i++ {
			if len(cur) > 0 && nums[i] < cur[len(cur)-1] || uset[nums[i]] {
				continue
			}
			cur = append(cur, nums[i])
			uset[nums[i]] = true
			backtracking(nums, i+1)
			cur = cur[:len(cur)-1]
		}
	}
	backtracking(nums, 0)
	fmt.Println(ans)
	return ans
}
