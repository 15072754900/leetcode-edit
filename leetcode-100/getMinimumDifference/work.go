package getMinimumDifference

import (
	tree "hufeng/demo/tree-Type"
	"math"
)

func getMinimumDifference(root *tree.TreeNode) int {
	var pre *tree.TreeNode = nil
	if root == nil {
		return 0
	}
	var dfs func(node *tree.TreeNode, num int) int
	dfs = func(node *tree.TreeNode, num int) int {
		if node == nil {
			return math.MaxInt64
		}
		left := dfs(node.Left, num)
		if pre != nil && node.Val-pre.Val < num {

		}
	}
}
