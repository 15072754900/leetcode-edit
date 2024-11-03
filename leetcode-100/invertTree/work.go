package invertTree

import tree "hufeng/demo/tree-Type"

func invertTree(root *tree.TreeNode) *tree.TreeNode {
	// 参数和返回值
	// 终止条件
	if root == nil {
		return nil
	}
	// 单层递归的逻辑
	// 前序遍历进行处理就行
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
