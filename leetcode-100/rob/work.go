package rob

import "fmt"

func rob(nums []int) int {
	// 首先建立边界条件
	if len(nums) == 1 {
		return nums[00]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], nums[i]+dp[i-2])
	}
	fmt.Println(dp)
	return dp[len(nums)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 还有一种记忆化递归方法：结合动态规划和DFS进行处理
