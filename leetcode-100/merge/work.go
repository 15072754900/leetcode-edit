package merge

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	ans := make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	//fmt.Println(intervals)
	for i := 1; i < len(intervals); {
		fmt.Println(intervals, ans)
		if intervals[i][0] > intervals[i-1][1] {
			ans = append(ans, intervals[i-1])
			intervals = intervals[1:]
		} else {
			cur := make([]int, 0)
			cur = append(cur, intervals[i-1][0])
			cur = append(cur, max(intervals[i][1], intervals[i-1][1]))
			ans = append(ans, cur)
			intervals = intervals[i+1:]
		}
	}
	ans = append(ans, intervals[0])
	fmt.Println(ans)
	return ans
}
