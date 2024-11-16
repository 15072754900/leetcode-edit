package hasPathSum

import tree "hufeng/demo/tree-Type"

// leetcode 112 路径总和

func hasPathSum(root *tree.TreeNode, targetSum int) bool {
	// 思路：一个前序遍历和回溯
	if root == nil {
		return false
	}
	// 退出条件一个是空节点，一个是叶子结点，其余退出条件不要使用
	if root.Left == nil && root.Right == nil && targetSum-root.Val == 0 {
		return true
	}
	var left, right bool
	if root.Left != nil {
		targetSum -= root.Val
		left = hasPathSum(root.Left, targetSum)
		targetSum += root.Val
	}
	if root.Right != nil {
		targetSum -= root.Val
		right = hasPathSum(root.Right, targetSum)
		targetSum += root.Val
	}
	return left || right
}
