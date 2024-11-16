package isValidBST

import (
	tree "hufeng/demo/tree-Type"
)

// 中序遍历加pre节点遍历

func isValidBSTWrong(root *tree.TreeNode) bool {
	pre := new(tree.TreeNode)
	if root.Left == nil && root.Right == nil {
		return true
	}
	var dfs func(node *tree.TreeNode) bool
	dfs = func(node *tree.TreeNode) bool {
		if node == nil {
			return true
		}
		left := dfs(node.Left)
		if pre != nil && pre.Val >= node.Val {
			return false
		}
		right := dfs(node.Right)
		return left && right
	}
	return dfs(root)
}

func tongyi(root *tree.TreeNode) bool {
	var pre *tree.TreeNode = nil // 将 pre 初始化为 nil
	var dfs func(node *tree.TreeNode) bool
	dfs = func(node *tree.TreeNode) bool {
		if node == nil {
			return true
		}
		// 首先递归遍历左子树
		if !dfs(node.Left) {
			return false
		}
		// 检查当前节点值是否大于前驱节点值
		if pre != nil && pre.Val >= node.Val {
			return false
		}
		// 更新前驱节点
		pre = node
		// 最后递归遍历右子树
		return dfs(node.Right)
	}
	return dfs(root)
}
