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

func tongyi(nums []int, num int) bool {
	// 相较于我的代码增加了部分取值边界问题，修改了不具备循环观察的能力，进行循环处理
	if num >= len(nums) {
		return false
	}
	if num+nums[num] >= len(nums)-1 {
		return true
	}
	for i := 1; i <= nums[num]; i++ {
		if tongyi(nums, num+i) {
			return true
		}
	}

	return false
}

// 但是上述方法会出现空间复杂度过大的问题，即栈溢出的问题，所以我们需要使用迭代或者贪心算法实现

func canJump(nums []int) bool {
	// 边界条件也要设置好
	if len(nums) == 1 {
		return true
	}
	// 这道题我们重点在于一个覆盖距离，只要当前节点的覆盖距离到达了最后节点，就说明我们可以到达，反之不能
	// 还有一个之前一直没想到的点，就是可取值区间，满足最大覆盖距离的区间就是能到达的区间，这个
	// 是为了符合局部最优
	n := len(nums)
	cover := 0
	for i := 0; i <= cover; i++ {
		cover = max(cover, i+nums[i])
		if nums[i]+i+1 >= n {
			fmt.Println(true)
			return true
		}
	}
	return false
}
