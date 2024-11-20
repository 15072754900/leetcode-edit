package permuteUnique

import (
	"fmt"
	"sort"
)

// leetcode 47 全排列2 源slice中包含重复元素，进行树层去重，之前也包含一个数值去重

func permuteUnique(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	ans := make([][]int, 0)
	cur := make([]int, 0)

	uset2 := make([]int, len(nums))
	var backtracking func(nums []int, startIndex int)
	backtracking = func(nums []int, startIndex int) {
		//fmt.Println(len(nums), len(cur), startIndex)
		if len(nums) == len(cur) {
			temp := make([]int, len(cur))
			copy(temp, cur)
			ans = append(ans, temp)
			return
		}
		uset1 := make(map[int]bool)
		for i := 0; i < len(nums); i++ {
			if uset1[nums[i]] || uset2[i] == 1 {
				continue
			}
			if i > 0 && nums[i-1] == nums[i] && uset2[i-1] == 0 {
				continue
			}
			cur = append(cur, nums[i])
			uset1[nums[i]] = true
			uset2[i] = 1
			backtracking(nums, 0)
			cur = cur[:len(cur)-1]
			uset1[nums[i]] = false
			uset2[i] = 0
		}
	}
	backtracking(nums, 0)
	fmt.Println(ans)
	return ans
}
