package isSymmetric

import (
	tree "hufeng/demo/tree-Type"
)

// leetcode的101对称二叉树

func isSymmetric(root *tree.TreeNode) bool {
	// 使用后序遍历、进行的递归变量不能只有一个
	// 错误原因：进行处理时退出存在错误
	if root == nil {
		return true
	}
	var dfs func(left, right *tree.TreeNode) bool
	dfs = func(left, right *tree.TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		return dfs(left.Left, right.Right) && dfs(left.Right, right.Left)
	}
	return dfs(root.Left, root.Right)
}
