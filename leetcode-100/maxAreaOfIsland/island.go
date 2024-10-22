package maxAreaOfIsland

func maxIsland(grid [][]int) int {
	maxArea := 0
	rows, cols := len(grid), len(grid[0])

	// 定义dfs函数进行最多邻域岛屿搜索
	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if r < 0 || c < 0 || r >= rows || c >= cols || grid[r][c] == 0 {
			return 0
		}
		grid[r][c] = 0 // 将访问到的数据标记为已访问
		area := 1
		dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 进行搜索的区间
		for _, d := range dirs {
			nr, nc := r+d[0], r+d[1]
			area += dfs(nr, nc)
		}
		return area
	}

	// 遍历整个网络，疑问：如何在进行判断的地方不去计算该位置从而减少复杂度
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				area := dfs(i, j)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}