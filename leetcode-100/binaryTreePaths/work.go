package binaryTreePaths

import (
	"fmt"
	tree "hufeng/demo/tree-Type"
)

// 这道题目就是使用一个前序遍历和回溯

func binaryTreePaths(root *tree.TreeNode) []string {
	if root == nil {
		return make([]string, 0)
	}
	ans := make([]string, 0)
	num := make([]int, 0)
	var dfs func(node *tree.TreeNode, num []int)
	dfs = func(node *tree.TreeNode, num []int) {
		if node == nil { // 这是判断当前节点为空节点的逻辑，不能在这里保存，而是在判断是叶子节点的时候进行保存
			return
		}
		num = append(num, node.Val)
		// 相比较回溯的源数据为非树形结构的情况，使用startIndex，这里不需要，直接按照递归和回溯进行
		dfs(node.Left, num)
		dfs(node.Right, num)
		if node.Left == nil && node.Right == nil {
			ans = append(ans, traversal(num))
		}
		num = num[:len(num)-1]
	}
	dfs(root, num)
	return ans
}

func traversal(nums []int) (ans string) {
	for i, v := range nums {
		if i == len(nums)-1 {
			ans += fmt.Sprintf("%d", v)
		} else {
			ans += fmt.Sprintf("%d->", v)
		}
	}
	return ans
}
