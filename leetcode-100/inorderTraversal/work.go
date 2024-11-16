package inorderTraversal

import (
	"fmt"
	tree "hufeng/demo/tree-Type"
)

// leetcode-94

func InorderTraversal(root *tree.TreeNode) []int {
	ans := make([]int, 0)
	var dfs func(node *tree.TreeNode, nums []int) []int
	dfs = func(node *tree.TreeNode, nums []int) []int {
		if node == nil {
			return nums
		}
		nums = dfs(node.Left, nums)
		// 中
		nums = append(nums, node.Val)
		nums = dfs(node.Right, nums)
		return nums
	}
	ans = dfs(root, ans)
	return ans
}

// 只进行输出数据但是不执行添加的代码，不输入到leetcode中进行分析

func printToTerminal(root *tree.TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	printToTerminal(root.Left)
	printToTerminal(root.Left)
}

func iterateFn(root *tree.TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	// 使用和之前不同的处理方式，首先，顺序是前-中-后，在处理数据之后就要获取一次右节点，而不是直接获取非空的左右子节点
	ans := make([]int, 0)
	stack := make([]*tree.TreeNode, 0)

	cur := root
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			// 说明此时没有坐子节点，直接输出子节点和该子节点的根节点
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans = append(ans, cur.Val)
			cur = cur.Right
		}
	}
	return ans
}
