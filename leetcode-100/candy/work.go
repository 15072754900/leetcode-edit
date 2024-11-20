package candy

import "fmt"

// leetcode 135 分发糖果 困难

// 这道题的重点在于不是同时决定三个数的中间数，而是从左到右再从右到左，这种情况是三个数进行比较，类似于接雨水

func candy(ratings []int) int {
	if len(ratings) == 1 {
		return 1
	}
	n := len(ratings)
	left := make([]int, n)
	right := make([]int, n)
	if ratings[0] <= ratings[1] {
		left[0] = 1
	} else {
		left[0] = 2
	}
	if ratings[n-1] > ratings[n-2] {
		right[n-1] = 2
	} else {
		right[n-1] = 1
	}
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
		if ratings[n-i] < ratings[n-i-1] {
			right[n-i-1] = right[n-i] + 1
		} else {
			right[n-i-1] = 1
		}
	}
	fmt.Println(left, right)
	fmt.Println(sum(left, right))
	return sum(left, right)
}

func sum(nums1, nums2 []int) (ans int) {
	for i := 0; i < len(nums1); i++ {
		ans += max(nums1[i], nums2[i])
	}
	return
}
