package levelOrder

import tree "hufeng/demo/tree-Type"

// 层序遍历 BFS

/*
相较于递归和迭代，这个是使用队列进行数据处理，注重数据的先进先出
题目102
*/

func levelOrder(root *tree.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := make([]*tree.TreeNode, 0)
	ans := make([][]int, 0)
	queue = append(queue, root)
	length := 0
	// 包含变量二叉树深
	for len(queue) > 0 {
		total := len(queue)
		level := make([]int, 0)
		for i := 0; i < total; i++ {
			node := queue[0] // 还是要推出节点的
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// 直接进行ans[length] = level 是对一个空的slice进行处理，会发生panic
		ans = append(ans, level)
		length++
	}
	return ans
}
