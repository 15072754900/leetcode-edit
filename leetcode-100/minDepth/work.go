package minDepth

import (
	tree "hufeng/demo/tree-Type"
)

// leetcode111 二叉树最小深度
// 进行编辑该树，首先是用BFS最好，然后是迭代，最后是递归
// 现在直接用递归进行分析， 首先是使用后序遍历，到叶子结点进行处理数据，从下向上，此题的核心要点是处理节点一定为非根节点
/*
三部曲：
1.递归参数和返回值：要返回一个最小值
2.递归退出条件：到达叶子节点退出
3.单层递归逻辑：进行比较和最小值之间的大小，使用min进行处理。
*/

func minDepth(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)

	// 此时可能为根节点进行，所以
	if root.Left == nil && root.Right != nil {
		return 1 + right
	}
	if root.Left != nil && root.Right == nil {
		return 1 + left
	}
	return min(left, right) + 1
}
