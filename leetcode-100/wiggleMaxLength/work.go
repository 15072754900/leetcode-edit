package wiggleMaxLength

import "fmt"

func wiggleMaxLength(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	ans := 1
	pre := nums[1] - nums[0]
	if pre != 0 {
		ans += 1
	}
	for i := 2; i < len(nums); i++ {
		cur := nums[i] - nums[i-1]
		if (cur > 0 && pre <= 0) || (cur < 0 && pre >= 0) {
			ans += 1
			pre = cur
		}
		//fmt.Println(cur, pre, ans)
	}
	fmt.Println(ans)
	return ans
}
