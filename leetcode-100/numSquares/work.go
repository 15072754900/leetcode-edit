package numSquares

import (
	"fmt"
	"math"
)

func numSquares(n int) int {
	// 思考之后，发现这种动态规划的题目核心在于：下一个的计算方式和前一个差不多，计算方法没区别，但是数据不一样，可以推出动态转移方程
	// 这道题：fn = 1 + min{sqrti~j}f[i-j^2]

	// 初始化取值slice
	fn := make([]int, n+1)
	for i := 1; i <= n; i++ {
		min := math.MaxInt
		for j := 1; j*j <= i; j++ {
			min = minI(min, fn[i-j*j])
		}
		fn[i] = min + 1
	}
	fmt.Println(fn)
	return fn[n]
}

func minI(x, y int) int {
	if x > y {
		return y
	}
	return x
}
