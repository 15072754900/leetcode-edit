package sumOfThree

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	// 获取排序之后的slice
	// 使用多个指针进行处理

	//hash := make(map[int]bool)
	//res := []int{}
	//for _, value := range nums {
	//
	//	if _, ok := hash[value]; !ok {
	//		hash[value] = true
	//		res = append(res, value)
	//	}
	//}
	//nums = res
	n := len(nums)
	if nums[0] == 0 || nums[n-1] == 0 {
		return [][]int{}
	}
	ans := [][]int{}
	midRight := 0

	// 进行处理，例如大于0和小于0的比较，减少时间复杂度

	for i := 0; i < n-1; i++ {

		if nums[i] == 0 && nums[i+1] != 0 {
			midRight = i
		}
	}
	fmt.Println(nums)
	fmt.Println("midRight", midRight)
	for j := 0; j < midRight; j++ {
		for r := j + 1; r <= midRight; r++ {
			for k := midRight + 1; k < n; k++ {
				fmt.Println(nums[j], nums[r], nums[k], j, r, k, "nums[j]+nums[r]", "nums[k]", nums[j]+nums[r], nums[k], nums[j]+nums[r]+nums[k])
				if nums[j]+nums[r]+nums[k] == 0 {
					ans = append(ans, []int{nums[j], nums[r], nums[k]})
				}
			}
		}
	}
	ans = unique2DSlice(ans)
	return ans
}

func sliceToString(slice []int) string {
	sort.Ints(slice) // 对切片进行排序以确保一致性
	result := ""
	for _, v := range slice {
		result += fmt.Sprintf("%d,", v)
	}
	return result
}

func unique2DSlice(slices [][]int) [][]int {
	uniqueMap := make(map[string][]int)
	var result [][]int

	for _, slice := range slices {
		key := sliceToString(slice)
		if _, exists := uniqueMap[key]; !exists {
			uniqueMap[key] = slice
			result = append(result, slice)
		}
	}

	return result
}

func threeSum2(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 枚举 a
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]
		// 枚举 b
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans

}
