package lengthOfLIS

import "fmt"

// 这题和完全平方数都是有多种实现方式的，需要进行多次训练

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	fmt.Println(dp)
	return dp[n-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 对于记忆化搜索方式
