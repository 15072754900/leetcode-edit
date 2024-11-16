package canJump

import "fmt"

func work(nums []int) bool {
	// 还是使用动态规划解决问题，首先判断第一个是否能走，然后进行原数组的修改f(i)s
	return tongyi(nums, 0)

}

func still(nums []int, num int) bool {
	if num >= len(nums) {
		return false
	}
	if num+nums[num] >= len(nums)-1 {
		return true
	}
	return still(nums, num+nums[num])
}

func tongyi(nums []int, startIndex int) bool {
	// 相较于我的代码增加了部分取值边界问题，修改了不具备循环观察的能力，进行循环处理
	if startIndex >= len(nums) {
		return false
	}
	if startIndex+nums[startIndex] >= len(nums)-1 {
		return true
	}
	for i := 1; i <= nums[startIndex]; i++ {
		if tongyi(nums, startIndex+i) {
			return true
		}
	}

	return false
}

// 但是上述方法会出现空间复杂度过大的问题，即栈溢出的问题，所以我们需要使用迭代或者贪心算法实现

// 现在是11-10，使用动态规划解法进行

func DP(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	for i := 0; i < len(nums); i++ {
		fmt.Println(get(nums, i))
		if get(nums, i) >= len(nums) {
			return true
		}
	}
	return false
}
func get(nums []int, step int) int {
	ans := 0
	n := nums[step]
	for i := 1; i < n+1; i++ {
		if i+step >= len(nums) {
			continue
		}
		ans += nums[i+step]
	}
	return ans
}

// 覆盖距离，而不是到达距离
func tanxin(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	farest := 0
	for i := 0; i < len(nums); i++ {
		if i > farest {
			return false
		}
		farest = max(farest, i+nums[i])
		if farest >= len(nums)-1 {
			return true
		}
	}
	return false
}

// 用迷宫问题的dp强行解

func trueDP(nums []int) bool {

	dp := make([]bool, len(nums))

	dp[0] = true

	for i := 0; i < len(nums); i++ {
		if dp[i] {
			for j := 1; j <= nums[i] && i+j < len(nums); j++ {
				dp[i+j] = dp[i]
			}
		}

	}
	return dp[len(nums)-1]
}
