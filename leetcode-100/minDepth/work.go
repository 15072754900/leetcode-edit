package minDepth

import (
	tree "hufeng/demo/tree-Type"
)

// leetcode111 二叉树最小深度

func minDepth(root *tree.TreeNode) (minLength int) {
	if root == nil {
		return 0
	}
	length := 0
	var dfs func(node *tree.TreeNode, num int) int
	dfs = func(node *tree.TreeNode, num int) int {
		if node == nil {
			return 0
		}
		if node.Left != nil {
			dfs(node.Left, num)
		}
		if node.Right != nil {
			dfs(node.Right, num)
		}
		return
	}

}
