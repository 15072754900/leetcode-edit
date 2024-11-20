package subsets

import "fmt"

// leetcode78子集问题

func subsets(nums []int) [][]int {
	// 由于原数组内容数据各不相同，所以不用树层去重
	ans := make([][]int, 0)
	cur := make([]int, 0)
	//ans = append(ans, []int{})
	var backtracking func(nums []int, startIndex int)
	backtracking = func(nums []int, startIndex int) {
		temp := make([]int, len(cur))
		copy(temp, cur)
		ans = append(ans, temp)
		for i := startIndex; i < len(nums); i++ {
			cur = append(cur, nums[i])
			backtracking(nums, i+1)
			cur = cur[:len(cur)-1]
		}
	}
	backtracking(nums, 0)
	fmt.Println(ans)
	return ans
}
