package minDistance

import "fmt"

func work(word1, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := 0; i < len(word1)+1; i++ {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i
	}
	for i := 0; i < len(word2)+1; i++ {
		dp[0][i] = i
	}
	// 建立完成边界条件

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}
	fmt.Println(dp)
	return dp[len(word1)][len(word2)]
}

func min(x, y, z int) int {
	if x > y {
		if z > y {
			return y
		}
		return z
	}
	if x < z {
		return x
	}
	return z
}
