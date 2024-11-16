package mergeTrees

import tree "hufeng/demo/tree-Type"

// leetcode 617，合并二叉树

func mergeTrees(root1 *tree.TreeNode, root2 *tree.TreeNode) *tree.TreeNode {
	// 还是三部曲，参数和返回值：可以决定是否需要进行匿名函数的设计。退出条件，看是都不为空还是一个为空，亦或者都为空，然后后进行单层递归逻辑
	if root1 == nil && root2 == nil {
		return nil
	}
	node := new(tree.TreeNode)
	if root1 != nil && root2 != nil {
		node.Val = root1.Val + root2.Val
	} else if root1 != nil && root2 == nil {
		node.Val = root1.Val
	} else if root2 != nil && root1 == nil {
		node.Val = root2.Val
	}
	if root1 != nil && root2 != nil {
		node.Left = mergeTrees(root1.Left, root2.Left)
	} else if root1 != nil && root2 == nil {
		node.Left = mergeTrees(root1.Left, nil)
	} else if root1 == nil && root2 != nil {
		node.Left = mergeTrees(nil, root2.Left)
	}
	if root1 != nil && root2 != nil {
		node.Right = mergeTrees(root1.Right, root2.Right)
	} else if root1 != nil && root2 == nil {
		node.Right = mergeTrees(root1.Right, nil)
	} else if root1 == nil && root2 != nil {
		node.Right = mergeTrees(nil, root2.Right)
	}
	return node
}

func Carl(root1, root2 *tree.TreeNode) *tree.TreeNode {
	// 不去全新定义一个二叉树而是进行将原二叉树进行修改，补全而不是创建
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	root1.Left = Carl(root1.Left, root2.Left)
	root1.Right = Carl(root1.Right, root2.Right)
	return root1
}
