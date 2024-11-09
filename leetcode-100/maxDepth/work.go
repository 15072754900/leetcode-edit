package maxDepth

import tree "hufeng/demo/tree-Type"

// 找到二叉树的最大深度：从上向下是深度，从下向上是高度。水深、树高。
// 这道题需要使用回溯操作，在执行到一个步骤时发现不是最大值进行其他分支的递归。首先是中序遍历，然后进行处理。

func maxDepth(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	var dfs func(node *tree.TreeNode, ans int) int
	dfs = func(node *tree.TreeNode, ans int) int {
		if node == nil {
			return ans
		}
		left := dfs(node.Left, ans+1)
		right := dfs(node.Right, ans+1)
		return max(left, right)
	}
	return dfs(root, ans)
}
