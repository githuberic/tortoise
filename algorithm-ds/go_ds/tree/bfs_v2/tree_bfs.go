package bfs_v2

import (
	"algorithm-ds/go_ds/tree"
)

func DFSV2(root *tree.Node) interface{} {
	if root == nil {
		return nil
	}

	var res []interface{}
	linkedList := NewQueue()
	linkedList.Offer(root)
	for !linkedList.IsEmpty() {
		cur := linkedList.Poll().(*tree.Node)
		res = append(res, cur.Val)

		if cur.Left != nil {
			linkedList.Offer(cur.Left)
		}

		if cur.Right != nil {
			linkedList.Offer(cur.Right)
		}
	}
	return res
}
