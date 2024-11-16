package combinationSum

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	if target == 0 {
		return make([][]int, 0)
	}
	ans := make([][]int, 0)
	cur := make([]int, 0)
	total := 0
	// 正是存在可以重复的原因，所以才要设置startIndex，这样可以后面出现前面出现过的组合，由于不存在重复元素，不用树层去重
	// 现在我的理解去重，有序的树层去重，无序的树枝去重
	var backtracking func(nums []int, target int, startIndex int)
	backtracking = func(nums []int, target, startIndex int) {
		if target < 0 {
			return
		}
		if target == 0 {
			fmt.Println(cur)
			temp := make([]int, len(cur))
			copy(temp, cur)
			ans = append(ans, temp)
			return
		}
		if total >= 150 {
			return
		}
		for i := startIndex; i < len(nums); i++ {
			cur = append(cur, nums[i])
			total++
			backtracking(nums, target-nums[i], i)
			total--
			cur = cur[:len(cur)-1]
		}
	}
	backtracking(candidates, target, 0)
	fmt.Println(ans)
	return ans
}
