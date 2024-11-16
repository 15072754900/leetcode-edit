package combinationSum2

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	// 首先进行排序，然后设计树层去重进行回溯处理
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	ans := make([][]int, 0)
	cur := make([]int, 0)
	// 进行树层去重，需要使用一个used进行设置
	used := make([]int, len(candidates)) // 初始化内容全为零，然后对其中的进行处理
	var backtracking func(nums, used []int, target, startIndex int)
	backtracking = func(nums, used []int, target, startIndex int) {
		if target < 0 {
			return
		}
		//if startIndex > 2 && nums[startIndex-1] == nums[startIndex-2] && used[startIndex-2] != 1 {
		//	return
		//}
		if target == 0 {
			fmt.Println(cur)
			temp := make([]int, len(cur))
			copy(temp, cur)
			ans = append(ans, temp)
			return
		}
		for i := startIndex; i < len(nums); i++ {
			if nums[i] > target {
				break
			}
			if i > 0 && nums[i] == nums[i-1] && used[i-1] != 1 {
				continue
			}
			cur = append(cur, nums[i])
			used[i] = 1
			backtracking(nums, used, target-nums[i], i+1)
			cur = cur[:len(cur)-1]
			used[i] = 0
		}
	}
	backtracking(candidates, used, target, 0)
	//fmt.Println(candidates, ans, used)
	return ans
}
