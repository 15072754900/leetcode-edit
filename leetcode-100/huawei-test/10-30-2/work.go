package huawei

import "fmt"

func work(n int, nums1, nums2 []int) int {
	// 构建一个图，其中节点表示 nums2 中的元素，边表示元素之间的交换关系
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if nums1[i] != nums2[j] {
				graph[i] = append(graph[i], j)
			}
		}
	}

	// 匹配数组，初始值为 -1 表示未匹配
	match := make([]int, n)
	for i := 0; i < n; i++ {
		match[i] = -1
	}

	// 递归寻找匹配
	var findMatch func(int, []bool) bool
	findMatch = func(v int, used []bool) bool {
		for _, to := range graph[v] {
			if !used[to] {
				used[to] = true
				if match[to] == -1 || findMatch(match[to], used) {
					match[to] = v
					return true
				}
			}
		}
		return false
	}

	// 尝试为每个节点找到匹配
	for i := 0; i < n; i++ {
		used := make([]bool, n)
		if !findMatch(i, used) {
			return -1
		}
	}

	// 计算最小代价
	cost := 0
	for i := 0; i < n; i++ {
		if match[i] != -1 {
			cost += i + match[i]
		}
	}

	fmt.Println(cost)
	return cost
}
