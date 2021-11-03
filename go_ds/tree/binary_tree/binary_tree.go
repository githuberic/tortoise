package binary_tree

type TreeNode struct {
	Val   interface{}
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	LDepth := maxDepth(root.Left)
	RDepth := maxDepth(root.Right)
	if LDepth > RDepth {
		return LDepth + 1
	} else {
		return RDepth + 1
	}
}

func NewNode(data interface{}) *TreeNode {
	return &TreeNode{Val: data}
}

/**
是否满二叉树
*/
func isFull(root *TreeNode) bool {
	if root == nil {
		return true
	}
	LDepth := maxDepth(root.Left)
	RDepth := maxDepth(root.Right)
	return isFull(root.Left) && isFull(root.Right) && (LDepth == RDepth)
}

/**
前序遍历

中 -> 左 -> 右
*/
func preOrderTraversal(root *TreeNode) []interface{} {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []interface{}{}
	}

	var stack []*TreeNode
	stack = append(stack, root)

	var res []interface{}
	res = append(res, root.Val)
	res = append(res, preOrderTraversal(root.Left))
	res = append(res, preOrderTraversal(root.Right)...)

	return res
}

/**
中序遍历
 左 -> 中 -> 右
*/
func inOrderTraversal(root *TreeNode) []interface{} {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []interface{}{}
	}

	res := inOrderTraversal(root.Left)
	res = append(res, root.Val)
	res = append(res, inOrderTraversal(root.Right)...)

	return res
}

/**
后续遍历
左 -> 右 -> 中
*/
func postOrderTraversal(root *TreeNode) []interface{} {
	if root == nil {
		return nil
	}

	var res []interface{}
	if root.Left != nil {
		lRes := postOrderTraversal(root.Left)
		if len(lRes) > 0 {
			res = append(res, lRes...)
		}
	}
	if root.Right != nil {
		rRes := postOrderTraversal(root.Right)
		if len(rRes) > 0 {
			res = append(res, rRes...)
		}
	}
	res = append(res, root.Val)
	return res
}

/**
从上往下打印出二叉树的每个结点，同一层的结点按照从左到右的顺序打印。
*/
func levelOrderTraversal(tree *TreeNode) []interface{} {
	if tree == nil {
		return nil
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, tree)

	var res []interface{}

	if len(queue) > 0 {
		tree = queue[0]
		if tree.Left != nil {
			queue = append(queue, tree.Left)
		}
		if tree.Right != nil {
			queue = append(queue, tree.Right)
		}
		res = append(res, tree.Val)
		queue = queue[1:] // queue pop
	}
	return res
}
