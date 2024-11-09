package findMode

import (
	"fmt"
	tree "hufeng/demo/tree-Type"
)

func findMode(root *tree.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	// 使用中序遍历
	ans := make([]int, 0)
	hash := make(map[int]int)
	longest := 0
	var dfs func(node *tree.TreeNode)
	pre := new(tree.TreeNode)
	pre.Val = -1 // 初始化为一个不可能的值
	dfs = func(node *tree.TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre.Val == node.Val {
			hash[node.Val]++
		} else {
			hash[node.Val] = 1
		}
		if hash[node.Val] > longest {
			longest = hash[node.Val]
		}
		pre = node
		if pre != nil {
			fmt.Println(pre.Val)
		}
		dfs(node.Right)
	}
	dfs(root)
	fmt.Println(longest, hash)
	for val, times := range hash {
		if times == longest {
			ans = append(ans, val)
		}
	}

	return ans
}

func tongyi(root *tree.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ans := make([]int, 0)
	hash := make(map[int]int)
	longest := 0
	pre := new(tree.TreeNode)
	pre.Val = -1 // 初始化为一个不可能的值

	var dfs func(node *tree.TreeNode)
	dfs = func(node *tree.TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)

		if pre.Val == node.Val {
			hash[node.Val]++
		} else {
			hash[node.Val] = 1
		}

		if hash[node.Val] > longest {
			longest = hash[node.Val]
		}

		pre = node

		dfs(node.Right)
	}

	dfs(root)

	for val, times := range hash {
		if times == longest {
			ans = append(ans, val)
		}
	}

	return ans
}

// 时间、空间复杂度都最小

func carlFn(root *tree.TreeNode) []int {
	// 按照视频的方法进行处理，一次处理输出所有众数
	ans := make([]int, 0)
	if root == nil {
		return ans
	}

	longest := 0
	// 关于cur的取值问题很重要
	cur := 0
	pre := new(tree.TreeNode)
	pre.Val = -1
	var dfs func(root *tree.TreeNode)

	dfs = func(root *tree.TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		// 进行中序遍历处理数据
		// 下面就是核心步骤，进行cur在当前节点的不同数据情况下进行的分析。以及一定要有的动态节点转换，pre = root
		if pre == nil {
			cur = 1
		} else if pre.Val == root.Val {
			cur++
		} else {
			cur = 1
		}
		pre = root
		if cur == longest {
			ans = append(ans, root.Val)
		}
		if cur > longest {
			longest = cur
			ans = make([]int, 0)
			ans = append(ans, root.Val)
		}

		dfs(root.Right)
	}

	dfs(root)
	return ans
}
