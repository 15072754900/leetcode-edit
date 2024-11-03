package postorderTraversal

import (
	"fmt"
	tree "hufeng/demo/tree-Type"
)

func postorderTraversal(root *tree.TreeNode) []int {
	ans := make([]int, 0)
	var dfs func(node *tree.TreeNode, nums []int) []int
	dfs = func(node *tree.TreeNode, nums []int) []int {
		if node == nil {
			return nums
		}
		nums = dfs(node.Left, nums)
		nums = dfs(node.Right, nums)
		nums = append(nums, node.Val)
		return nums
	}
	return dfs(root, ans)
}

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
		if node.Left != nil {
			stack = append(stack, node.Left)
			fmt.Println(node.Left.Val)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
			fmt.Println(node.Right.Val)
		}
		printNums(stack)
	}
	reverse(ans)
	fmt.Println(ans)
	return ans
}

func printNums(nums []*tree.TreeNode) {
	for _, val := range nums {
		fmt.Printf("%d", val.Val)
	}
}

func reverse(nums []int) {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
