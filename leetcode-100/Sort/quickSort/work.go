package quickSort

import "fmt"

func quickSort(nums []int, low, high int) {
	if low < high {
		num := partition(nums, low, high)
		quickSort(nums, low, num-1)
		quickSort(nums, num+1, high)
	}
}

func partition(nums []int, left, right int) int {
	n := nums[right]
	i, j := left-1, left

	for ; j < right; j++ {
		if nums[j] <= n {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	nums[i+1], nums[right] = nums[right], nums[i+1]

	fmt.Println(nums)
	return i + 1
}
