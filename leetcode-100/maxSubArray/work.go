package maxSubArray

func maxSubArray(nums []int) int {
	// 该题目的思路：两个变量的修改，首先是left，然后是一个max，记录位置和最大值
	if len(nums) == 1 {
		return nums[0]
	}
	now := 0
	right := 0
	ans := 0
	cur := 0
	for left := 0; left < len(nums)-1; {
		now += nums[left]
		if now < 0 || now+nums[left+1] < 0 {
			now = 0
			left++
			continue
		} else {
			for right+left <= len(nums) {
				if right+left == len(nums) {
					return ans
				}
				cur += nums[right+left]
				if cur < 0 {
					break
				}
				if ans < cur {
					ans = cur
				}
				right++
			}
		}
	}
	if ans <= 0 {
		ans += nums[len(nums)-1]
	}
	return ans
}

func tongyi(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	// 初始化当前和与最大和
	curMax := nums[0]
	maxSum := nums[0]

	// 遍历数组从第二个元素开始
	for i := 1; i < len(nums); i++ {
		// 如果当前和为负，则重置为当前元素；否则加上当前元素
		if curMax > 0 {
			curMax += nums[i]
		} else {
			curMax = nums[i]
		}

		// 更新最大和
		if curMax > maxSum {
			maxSum = curMax
		}
	}

	return maxSum
}

func hufeng(nums []int) int {
	// 建立动态规划方程
	maxI := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i]
		}
		if maxI < nums[i] {
			maxI = nums[i]
		}
	}
	return maxI
}
