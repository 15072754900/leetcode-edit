package uniquePaths

func uniquePaths(m, n int) int {
	// 按照动态规划的思路进行处理，首先两个问题，1.建立动态转移方程。2.设置边界条件
	ans := make([][]int, m)
	for index := range ans {
		ans[index] = make([]int, n)
		ans[index][0] = 1
	}
	for j := 0; j < n; j++ {
		ans[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			ans[i][j] = ans[i-1][j] + ans[i][j-1]
		}
	}
	return ans[m-1][n-1]

	// 按照之前的思路：建立一个二维slice，然后定义一个二维slice、bool类型，进行每一次标注已经完成
	//way := [][]int{{1, 0}, {0, 1}}
	//done := make([][]bool, len(way))
	//
	//// 一个转移方程、一个边界条件，这两个如何进行处理
	//// 本质上是通过遍历所有节点，判断是否能从该节点到达终点，可以结果加一，不能下一个继续遍历
	//for i := 0; i < m; i++ {
	//	for j := 0; j < n; j++ {
	//
	//	}
	//}

}

//func judgeToEnd() {
//
//}
