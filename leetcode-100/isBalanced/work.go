package isBalanced

import (
	tree "hufeng/demo/tree-Type"
)

// leetcode第110题，判断一棵树是否是二叉平衡树

/*
这里说明二叉树的类型
1.二叉树：包含左子节点、右子节点
2.二叉搜索树：树的任意一个节点都满足左子节点小于本身值，右子节点大于本身值
3.平衡二叉树：树的左右叶子节点的高度差绝对值不超过1
4.avl树：是一颗平衡二叉搜索树
*/

// 思路：进行前序遍历，首先是两个变量，作为最大深度和最小深度。第二种解法：分别获取最大和最小值，然后进行比较，但是这样就不好了
// 或者使用后序遍历，进行判断每一个节点是否为平衡二叉树

func isBalanced(root *tree.TreeNode) bool {
	if root == nil {
		return true
	}
	n1, n2 := 0, 0
	var dfs func(root *tree.TreeNode) int
	dfs = func(root *tree.TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		if root.Left == nil && root.Right != nil {
			return 1 + right
		}
		if root.Left != nil && root.Left == nil {
			return 1 + left
		}
		return min(left, right) + 1
	}
	n1 = dfs(root)
	dfs = func(root *tree.TreeNode) int {
		if root == nil {
			return 0
		}
		return max(dfs(root.Left)+1, dfs(root.Right)+1)
	}
	n2 = dfs(root)
	if n2-n1 > 1 {
		return false
	}
	return true
}

func carel(root *tree.TreeNode) bool {
	// 首先，还是后序遍历，但是输出的值包含处理方法，返回一个非-1值就是所在节点都是平衡二叉树，且数值为树高，返回-1就不是二叉平衡树
	if root == nil {
		return true
	}
	var dfs func(node *tree.TreeNode) int
	dfs = func(node *tree.TreeNode) int {
		if node == nil {
			return 0
		}
		// 这个地方一定要注意，对于一个数据的处理，当存在一个-1， 就直接一路退出
		leftDepth := dfs(node.Left)
		if leftDepth == -1 {
			return -1
		}
		rightDepth := dfs(node.Right)
		if rightDepth == -1 {
			return -1
		}
		if leftDepth-rightDepth > 1 || rightDepth-leftDepth > 1 {
			return -1
		}
		return max(leftDepth, rightDepth) + 1
	}
	if dfs(root) == -1 {
		return false
	}
	return true
}

func carl(root *tree.TreeNode) bool {
	if root == nil {
		return true
	}
	var dfs func(node *tree.TreeNode) int
	dfs = func(node *tree.TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		if left == -1 {
			return -1
		}
		right := dfs(node.Right)
		if right == -1 {
			return -1
		}
		if left-right > 1 || right-left > 1 {
			return -1
		}
		return max(left, right) + 1
	}
	ans := dfs(root.Left)
	if ans == -1 {
		return false
	}
	return true
}
