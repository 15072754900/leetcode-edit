package subsetsWithDup

import (
	"fmt"
	"sort"
)

// leetcode 90 子集2，去除重复元素的子集，就是进行一个树层去重

func subsetsWithDup(nums []int) [][]int {
	// 由于原数组内容数据各不相同，所以不用树层去重
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	fmt.Println(nums)
	ans := make([][]int, 0)
	cur := make([]int, 0)
	used := make([]int, len(nums))
	ans = append(ans, []int{})
	var backtracking func(nums, used []int, startIndex int)
	backtracking = func(nums, used []int, startIndex int) {
		for i := startIndex; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] && used[i-1] == 0 {
				// 很重要：这里的树层去重就是要处理当前元素和前一个元素相同的情况，但是不是在一个子树上。
				//注重的是前一个元素未被使用。
				continue
			}
			cur = append(cur, nums[i])
			temp := make([]int, len(cur))
			copy(temp, cur)
			ans = append(ans, temp)
			used[i] = 1
			backtracking(nums, used, i+1)
			used[i] = 0
			cur = cur[:len(cur)-1]
		}
	}
	backtracking(nums, used, 0)
	fmt.Println(ans)
	return ans
}
