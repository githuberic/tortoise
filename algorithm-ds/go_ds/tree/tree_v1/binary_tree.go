package tree_v1

import (
	"algorithm-ds/go_ds/tree"
	"algorithm-ds/go_ds/tree/stack/v1"
	"math"
)

// Height 二叉树深度
func Height(this *tree.Node) int {
	if this == nil {
		return 0
	}

	lHeight := Height(this.Left)
	rHeight := Height(this.Right)
	if lHeight > rHeight {
		return lHeight + 1
	} else {
		return rHeight + 1
	}
}

/**
满二叉树
*/
func isFull(root *tree.Node) bool {
	if root == nil {
		return true
	}

	lHeight := Height(root.Left)
	rHeight := Height(root.Right)
	return isFull(root.Left) && isFull(root.Right) && (lHeight == rHeight)
}

/**
是否平衡二叉树
*/
func isBalanced(root *tree.Node) bool {
	if root == nil {
		return true
	}

	lde := Height(root.Left)
	rde := Height(root.Right)
	flag := false
	if math.Abs(float64(lde-rde)) <= 1 {
		flag = true
	}
	return flag && isBalanced(root.Left) && isBalanced(root.Right)
}

/**
前序遍历

中 -> 左 -> 右
*/
func preOrderTraversal(root *tree.Node) []interface{} {
	if root == nil {
		return nil
	}

	var res []interface{}
	res = append(res, root.Val)
	res = append(res, preOrderTraversal(root.Left)...)
	res = append(res, preOrderTraversal(root.Right)...)

	return res
}

/**
中序遍历
 左 -> 中 -> 右
*/
func inOrderTraversal(root *tree.Node) []interface{} {
	if root == nil {
		return nil
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
func postOrderTraversal(root *tree.Node) []interface{} {
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

// 按层遍历
func LevelOrderTraversal(root *tree.Node) []interface{} {
	if root == nil {
		return nil
	}

	var res []interface{}

	// 采用队列实现
	queue := make([]*tree.Node, 0)
	// queue push
	queue = append(queue, root)
	for len(queue) > 0 {
		root = queue[0]
		res = append(res, root.Val)

		// queue pop
		queue = queue[1:]
		if root.Left != nil {
			queue = append(queue, root.Left)
		}
		if root.Right != nil {
			queue = append(queue, root.Right)
		}
	}
	return res
}

/**
Tree深度优先遍历
*/
func DFS(root *tree.Node) []interface{} {
	if root == nil {
		return nil
	}

	reS := v1.NewStack()
	s := v1.NewStack()
	s.Push(root)
	for !s.Empty() {
		cur := s.Pop().(*tree.Node)
		if cur.Left != nil {
			s.Push(cur.Left)
		}
		if cur.Right != nil {
			s.Push(cur.Right)
		}
		reS.Push(cur)
	}

	var res []interface{}
	for !reS.Empty() {
		res = append(res, reS.Pop().(*tree.Node).Val)
	}

	return res
}
