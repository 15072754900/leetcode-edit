package reverse

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mirror(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	mirror(root.Left)
	mirror(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

func mirror2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			queue = append(queue, queue[0].Left)
		}
		if node.Right != nil {
			queue = append(queue, queue[0].Right)
		}
		queue = queue[1:]
	}
	return root
}
