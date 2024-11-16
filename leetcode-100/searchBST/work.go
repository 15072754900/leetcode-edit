package searchBST

import tree "hufeng/demo/tree-Type"

// leetcode 700 这是对二叉搜索时的处理，二叉搜索树的特性：进行中序遍历时获取的数据都是有序的

func searchBST(root *tree.TreeNode, val int) *tree.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if root.Val > val {
		root = searchBST(root.Left, val)
	} else {
		root = searchBST(root.Right, val)
	}
	return root
}
