package combine

import "fmt"

func combine(n, k int) [][]int {
	if n == 0 {
		return make([][]int, 0)
	}
	ans := make([][]int, 0)
	num := make([]int, 0)

	var backtracking func(n, k, startIndex int)
	backtracking = func(n, k, startIndex int) {
		// 在golang里面很重要的一点是num作为引用类型，不能作为append的参数，需要建立一个临时数组记录值，每次进行初始化一个实例。
		//这里是不能将num每次作为一个初始化值进行使用
		if len(num) == k {
			temp := make([]int, k)
			copy(temp, num)
			ans = append(ans, temp)
			return
		}
		// 下面是修改了选取区间的剪枝操作
		for i := startIndex; i <= n-(k-len(num))+1; i++ {
			fmt.Println(i)
			num = append(num, i)
			backtracking(n, k, i+1)
			num = num[:len(num)-1] // 在完成、到达一个分支的根节点时进行推出。
		}
	}
	backtracking(n, k, 1)
	fmt.Println(ans)
	return ans
}
