package combinationSum3

import "fmt"

func combinationSum3(k, n int) [][]int {
	// 这里可以用到used，但是没有重复的数字，所以直接不用去重
	//used := make([]int, k)
	if n == 0 || k == 0 {
		return make([][]int, 0)
	}
	ans := make([][]int, 0)
	num := make([]int, 0)
	var backtracking func(k, n, startIndex int)
	backtracking = func(k, n, startIndex int) {
		if k > (9-startIndex+1) || n < 0 {
			return
		}
		if n == 0 && k == 0 {
			// 需要注意，这里的k==0，需要使用num的长度进行初始化，否则只会是0的长度
			temp := make([]int, len(num))
			copy(temp, num)
			//fmt.Println("1", num, "2", temp)
			ans = append(ans, temp)
			return
		}

		for i := startIndex; i <= 9-len(num)+1; i++ {
			//if i > n/k {
			//	break
			//}
			num = append(num, i)
			backtracking(k-1, n-i, i+1)
			num = num[:len(num)-1]
		}
	}
	backtracking(k, n, 1)
	fmt.Println(ans)
	return ans
}
