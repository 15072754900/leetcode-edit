package permute

import "fmt"

// leetcode46题，全排列，重点在与排列，所以是不进行startIndex的
//其中一个核心思想就是结果不包含重复的元素，算是树枝去重

func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	cur := make([]int, 0)
	var backtracking func(nums []int, startIndex int)
	uset := make(map[int]bool)
	backtracking = func(nums []int, startIndex int) {
		if len(cur) == len(nums) {
			temp := make([]int, len(cur))
			copy(temp, cur)
			ans = append(ans, temp)
			return
		}
		for i := 0; i < len(nums); i++ {
			fmt.Println(i, nums[i], uset)
			if uset[nums[i]] {
				fmt.Println(nums[i])
				continue
			}
			uset[nums[i]] = true
			cur = append(cur, nums[i])
			backtracking(nums, i+1)
			uset[nums[i]] = false
			cur = cur[:len(cur)-1]
		}
	}
	backtracking(nums, 0)
	fmt.Println(ans)
	return ans
}
