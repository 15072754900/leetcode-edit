package reverNums

func search(nums []int, target int) int {
	//思路：先进行数组恢复，然后对原始数组进行搜索，查找到数据位置，然后将数据进行处理就行
	// 但是这个有要求：时间复杂度必须为logn
	left, right := 0, len(nums)-1

	return Judge(left, right, target, nums)
}

// 判断那一边是有序的，进行处理

func Judge(left int, right int, target int, nums []int) int {
	for left <= right {
		mid := (right-left)>>1 + left
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] { // 左边有序
			if nums[left] <= target && nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // 右边有序
			if nums[mid] < target && nums[right] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

// 方法还是二分查找法，只是在实现上有些差别，需要加一些逻辑。这个区间选取是一个重点，要考虑的地方。

func myJudge(left, right, target int, nums []int) int {
	for left < right {
		mid := (right-left)>>2 + left
		if nums[mid] > target && nums[left] > target {
			left = mid + 1
		}
		if nums[mid] > target && nums[left] < target {
			right = mid
		}
		if nums[left] == target {
			return left
		}
		if nums[mid] == target {
			return mid
		}
		if nums[right] == target {
			return right
		}
	}
	return -1
}
