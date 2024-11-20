package findMinArrowShots

import (
	"fmt"
	"sort"
)

// leetcode 452 射爆气球数量

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	fmt.Println(points)
	ans := 1
	// 选取的方案是将其右边界直接转为上一个的右边界，前提是当前左边界小于前一个的右边界，这就是一个局部最优解
	for i := 1; i < len(points); i++ {
		if points[i][0] > points[i-1][1] {
			ans += 1
		} else {
			points[i][1] = min(points[i-1][1], points[i][1])
		}
	}
	fmt.Println(ans)
	return ans
}
