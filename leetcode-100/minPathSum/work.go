package minPathSum

import "fmt"

// 处理矩阵中的最短路径
/*
首先，建立每个节点的到达之和
然后，定义可能的路径和，进行比较并加入到数组中
总的思路还是一个简单的动态规划，熟悉该问题，处理该问题使用到的状态转移方程
*/

// 在处理函数的定义上，可以实现很多种类，现在的一个简单的内容，包含n2、n的时间复杂度，使用的空间也不同，这是很重要的内容，必须学会

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	m := len(grid)
	n := len(grid[0])
	// 进行边缘的区域的计算
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[0]))
		if i == 0 {
			dp[i][0] = grid[i][0]
		} else {
			dp[i][0] = grid[i][0] + dp[i-1][0]
		}
	}
	for i := 0; i < len(grid[0]); i++ {
		if i == 0 {
			dp[0][i] = grid[0][i]
		} else {
			dp[0][i] = grid[0][i] + dp[0][i-1]
		}
	}
	fmt.Println(dp)

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	fmt.Println(dp)

	return dp[m-1][n-1]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
