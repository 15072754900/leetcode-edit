package groupAnagrams

import "fmt"

func groupAnagrams(ss []string) [][]string {
	ans := make([]string, len(ss))
	copy(ans, ss)

	// 定义一个没定义过的map，这个的意思是定义结构的所有数据和一个对应的排完序的string，这样就可以不是string对索引，而是直接进行数据对所有数据
	hash := make(map[string][]string)

	for i := 0; i < len(ans); i++ {
		cur := []byte(ans[i])
		quickSort(cur, 0, len(ans[i])-1)
		ans[i] = string(cur)
		fmt.Println(ans[i])
		hash[ans[i]] = append(hash[ans[i]], ss[i])
	}

	// 然后定义一个返回的slice
	ansS := make([][]string, 0, len(hash)) // 不明确长度，但是明确体积，避免内存指针重复分配
	for _, value := range hash {
		ansS = append(ansS, value)
	}
	fmt.Println(hash, ansS)
	return ansS
}

func quickSort(nums []uint8, low, high int) {
	// 在内部进行slice数据的修改位置，使用快速排序的思想，进行处理。现在还没有实现堆排序。
	if low < high {
		num := partition(nums, low, high)
		quickSort(nums, low, num-1)
		quickSort(nums, num+1, high)
	}
}

func partition(nums []uint8, low, high int) int {
	i, j := low-1, low
	target := nums[high]

	for ; j < high; j++ {
		if nums[j] <= target {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[high] = nums[high], nums[i+1]
	return i + 1
}

func quickSortRune(nums []rune, low, high int) {
	// 在内部进行slice数据的修改位置，使用快速排序的思想，进行处理。现在还没有实现堆排序。
	if low < high {
		num := partitionRune(nums, low, high)
		quickSortRune(nums, low, num-1)
		quickSortRune(nums, num+1, high)
	}
}

func partitionRune(nums []rune, low, high int) int {
	i, j := low-1, low
	target := nums[high]

	for ; j < high; j++ {
		if nums[j] <= target {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[high] = nums[high], nums[i+1]
	return i + 1
}
