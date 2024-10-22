package nextPermutation

import (
	"fmt"
	"sort"
)

func nextPermutation(nums []int) {
	// 1.找拐点，2.找到拐点数据后面的比拐点数据大一点的数
	n := 0
	for i := len(nums) - 1; i > 0; i-- {
		fmt.Println(nums[i], nums[i-1])
		if nums[i] > nums[i-1] {
			fmt.Println(nums)

			for j := len(nums) - 1; j > 0; j-- {
				if nums[j] > nums[i-1] {
					nums[j], nums[i-1] = nums[i-1], nums[j]
					break
				}
			}
			n = i
			fmt.Println(nums)
			break
		}
	}

	// n是拐点，把n后面的数据进行降序排列就行
	sort.Ints(nums[n:])
	fmt.Println(nums)
}
