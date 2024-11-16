package buildTree

import (
	"hufeng/demo/inorderTraversal"
	"hufeng/demo/postorderTraversal"
	"hufeng/demo/preorderTraversal"
	tree "hufeng/demo/tree-Type"
)

// leetcode-105、106 使用前序遍历、中序遍历和前序遍历、中序遍历进行处理。

func buildTreeInAndPost(inorder, postorder []int) *tree.TreeNode {
	// 使用前序遍历进行生成树，根节点信息按照组合slice进行处理
	// 只会进行后序的删减
	if len(postorder) == 0 {
		return nil
	}
	root := new(tree.TreeNode)
	top := postorder[len(postorder)-1]
	root.Val = top
	pla := getMid(inorder, top)
	// postorder = postorder[:len(postorder)-1] 这一步存在错误
	root.Left = buildTreeInAndPost(inorder[:pla], postorder[:pla])
	root.Right = buildTreeInAndPost(inorder[pla+1:], postorder[pla:len(postorder)-1])
	return root
}

// leetcode105 前序遍历、中序遍历构造二叉树

func buildTreePreAndIn(preorder, inorder []int) *tree.TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	node := new(tree.TreeNode)
	node.Val = preorder[0]
	index := getMid(inorder, node.Val)
	node.Left = buildTreePreAndIn(preorder[1:index+1], inorder[:index])
	node.Right = buildTreePreAndIn(preorder[index+1:], inorder[index+1:])
	return node
}

func getMid(inorder []int, num int) int {
	for i, v := range inorder {
		if v == num {
			num = i
			return num
		}
	}
	return -1
}

func getAnsPost(inorder, postorder []int) []int {
	return postorderTraversal.PostorderTraversal(buildTreePreAndIn(inorder, postorder))
}
func getAnsPre(inorder, postorder []int) []int {
	return preorderTraversal.PreorderTraversal(buildTreePreAndIn(inorder, postorder))
}
func getAnsIn(inorder, postorder []int) []int {
	return inorderTraversal.InorderTraversal(buildTreePreAndIn(inorder, postorder))
}
