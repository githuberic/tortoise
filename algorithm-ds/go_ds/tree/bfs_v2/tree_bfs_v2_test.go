package bfs_v2

import (
	"algorithm-ds/go_ds/tree"
	"fmt"
	"testing"
)

/**
验证是否满二叉树
*/
func TestBFS(t *testing.T) {
	root := tree.NewNode(1)
	n1 := tree.NewNode(2)
	n2 := tree.NewNode(3)
	n3 := tree.NewNode(4)
	n4 := tree.NewNode(5)
	root.Left = n1
	root.Right = n2
	n1.Left = n3
	n1.Right = n4
	n3.Left = tree.NewNode(41)
	n3.Right = tree.NewNode(42)

	values := DFSV2(root)
	fmt.Printf("BFS Order Value %v\n", values)
}
