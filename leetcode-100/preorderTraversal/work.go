package preorderTraversal

import (
	"fmt"
	tree "hufeng/demo/tree-Type"
)

// 进行前序遍历（中左右）leetcode-144
/*
步骤：
1.确定递归函数的参数和返回值
2.确定终止条件
3.确定单层递归的逻辑
*/

func PreorderTraversal(root *tree.TreeNode) []int {
	ans := make([]int, 0)
	var dfs func(root *tree.TreeNode, nums []int) []int
	dfs = func(root *tree.TreeNode, nums []int) []int {
		if root == nil {
			return nums
		}
		nums = append(nums, root.Val)
		nums = dfs(root.Left, nums)
		nums = dfs(root.Right, nums)
		return nums
	}
	ans = dfs(root, ans)
	return ans
}

func do1(root *tree.TreeNode) []int {
	ans := make([]int, 0)
	var dfs func(node *tree.TreeNode) []int
	dfs = func(node *tree.TreeNode) []int {
		if node == nil {
			return ans
		}
		ans = append(ans, node.Val)
		dfs(root.Left)
		dfs(root.Right)
		return ans
	}
	return ans
}

func tongyi(root *tree.TreeNode) []int {
	var dfs func(root *tree.TreeNode, nums []int) []int
	dfs = func(root *tree.TreeNode, nums []int) []int {
		if root == nil {
			return nums
		}

		nums = append(nums, root.Val)
		fmt.Println(root.Val, nums)
		nums = dfs(root.Left, nums)
		nums = dfs(root.Right, nums)
		return nums
	}
	ans := dfs(root, make([]int, 0))
	fmt.Println(ans)
	return ans
}

// 作为一个压栈的解法进行处理二叉树，但是BFS是使用队列来执行，使用不同的数据结构的特性可以实现不同的功能
// 先出栈，获取节点，再进行后续节点入栈，需要先出了根节点

func iterateFn(root *tree.TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	stack := make([]*tree.TreeNode, 0)
	ans := make([]int, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
			fmt.Println(node.Right.Val)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
			fmt.Println(node.Left.Val)
		}
		printNums(stack)
	}
	fmt.Println(ans)
	return ans
}

func printNums(nums []*tree.TreeNode) {
	for _, val := range nums {
		fmt.Printf("%d", val.Val)
	}
}
