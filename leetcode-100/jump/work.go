package jump

import "fmt"

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	times := 0
	cover := 0
	cur := 0
	for i := 0; i < len(nums)-1; i++ {
		cover = max(cover, i+nums[i])
		// 当到达了当前节点可以到达的最远位置之后，进行下一步，而不是进行所有的都加一次
		// 然后更新最远位置
		if i >= cur {
			times++
			cur = cover
			if cover == len(nums)-1 {
				break
			}
		}
	}
	fmt.Println(times)
	return times
}
